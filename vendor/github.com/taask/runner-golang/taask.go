package taask

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"path"
	"time"

	"github.com/cohix/simplcrypto"

	log "github.com/cohix/simplog"
	"github.com/pkg/errors"
	cconfig "github.com/taask/client-golang/config"
	"github.com/taask/runner-golang/config"
	"github.com/taask/taask-server/auth"
	"github.com/taask/taask-server/model"
	"github.com/taask/taask-server/service"
	"google.golang.org/grpc"
)

// TaskHandler represents a handler that can handle a task
type TaskHandler func([]byte) (interface{}, error)

// SpecTaskHandler represents a handler that can handle a spec task
type SpecTaskHandler func(*DecryptedTask) (interface{}, error)

// Runner describes a runner
type Runner struct {
	runner      *model.Runner
	client      service.RunnerServiceClient
	localAuth   *cconfig.LocalAuthConfig
	handler     TaskHandler
	specHandler SpecTaskHandler
}

// NewRunner creates a new runner
func NewRunner(kind string, tags []string, handler TaskHandler) (*Runner, error) {
	modelRunner := &model.Runner{
		UUID: model.NewRunnerUUID(),
		Kind: kind,
		Tags: tags,
	}

	runner := &Runner{
		runner:  modelRunner,
		handler: handler,
	}

	return runner, nil
}

// NewSpecRunner creates a new runner that handles spec tasks
func NewSpecRunner(kind string, tags []string, handler SpecTaskHandler) (*Runner, error) {
	modelRunner := &model.Runner{
		UUID: model.NewRunnerUUID(),
		Kind: kind,
		Tags: tags,
	}

	runner := &Runner{
		runner:      modelRunner,
		specHandler: handler,
	}

	return runner, nil
}

// ConnectAndStreamTasks connects to a taask-server, registers the runner, and runs
func (r *Runner) ConnectAndStreamTasks(addr, port string) error {
	log.LogInfo(fmt.Sprintf("starting runner of kind %s", r.runner.Kind))

	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", addr, port), grpc.WithInsecure())
	if err != nil {
		return errors.Wrap(err, "failed to Dial")
	}

	r.client = service.NewRunnerServiceClient(conn)

	if err := r.authWithDefaultConfig(); err != nil {
		return errors.Wrap(err, "failed to auth")
	}

	if err := r.run(); err != nil {
		return errors.Wrap(err, "failed to run")
	}

	return nil
}

// ConnectAndRunTaskFromFile connects to a taask-server, registers the runner, and runs one task
func (r *Runner) ConnectAndRunTaskFromFile(filepath, addr, port string) error {
	log.LogInfo(fmt.Sprintf("starting runner of kind %s", r.runner.Kind))

	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", addr, port), grpc.WithInsecure())
	if err != nil {
		return errors.Wrap(err, "failed to Dial")
	}

	r.client = service.NewRunnerServiceClient(conn)

	if err := r.authWithDefaultConfig(); err != nil {
		return errors.Wrap(err, "failed to auth")
	}

	task, err := readTaskFile(filepath)
	if err != nil {
		return errors.Wrap(err, "failed to readTaskFile")
	}

	r.runTask(task)

	return nil
}

func (r *Runner) authWithDefaultConfig() error {
	filepath := path.Join(config.DefaultRunnerConfigDir(), config.ConfigRunnerDefaultFilename)
	localAuth, err := cconfig.LocalAuthConfigFromFile(filepath)
	if err != nil {
		return errors.Wrap(err, "failed to LocalAuthConfigFromFile")
	}

	return r.authenticate(localAuth)
}

// Authenticate auths with the taask server and saves the session
func (r *Runner) authenticate(localAuth *cconfig.LocalAuthConfig) error {
	r.localAuth = localAuth

	memberUUID := model.NewRunnerUUID()

	keypair, err := simplcrypto.GenerateNewKeyPair()
	if err != nil {
		return errors.Wrap(err, "failed to GenerateNewKeyPair")
	}

	timestamp := time.Now().Unix()

	nonce := make([]byte, 8)
	binary.LittleEndian.PutUint64(nonce, uint64(timestamp))
	hashWithNonce := append(r.localAuth.MemberGroup.AuthHash, nonce...)

	authHashSig, err := keypair.Sign(hashWithNonce)
	if err != nil {
		return errors.Wrap(err, "failed to Sign")
	}

	attempt := &service.AuthMemberRequest{
		UUID:              memberUUID,
		GroupUUID:         r.localAuth.MemberGroup.UUID,
		PubKey:            keypair.SerializablePubKey(),
		AuthHashSignature: authHashSig,
		Timestamp:         timestamp,
	}

	authResp, err := r.client.AuthRunner(context.Background(), attempt)
	if err != nil {
		return errors.Wrap(err, "failed to AuthClient")
	}

	challengeBytes, err := keypair.Decrypt(authResp.EncChallenge)
	if err != nil {
		return errors.Wrap(err, "failed to Decrypt challenge")
	}

	challengeSig, err := keypair.Sign(challengeBytes)
	if err != nil {
		return errors.Wrap(err, "failed to Sign challenge")
	}

	session := cconfig.ActiveSession{
		Session: &auth.Session{
			MemberUUID:          memberUUID,
			GroupUUID:           r.localAuth.MemberGroup.UUID,
			SessionChallengeSig: challengeSig,
		},
		Keypair: keypair,
	}

	r.localAuth.ActiveSession = session

	return nil
}

func (r *Runner) run() error {
	req := &service.RegisterRunnerRequest{
		Kind:    r.runner.Kind,
		Tags:    r.runner.Tags,
		Session: r.localAuth.ActiveSession.Session,
	}

	log.LogInfo("registering with server...")

	stream, err := r.client.RegisterRunner(context.Background(), req)
	if err != nil {
		return errors.Wrap(err, "failed to RegisterRunner")
	}

	log.LogInfo("ready to receive tasks")

	for {
		task, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				log.LogError(errors.New("stream broken; terminating"))
				break
			}

			log.LogError(errors.Wrap(err, "stream error"))
			break
		}

		if task.UUID == "" {
			// an empty task is like a heartbeat, ignore it
			continue
		}

		log.LogInfo(fmt.Sprintf("received task with uuid %s", task.UUID))

		go r.runTask(task)
	}

	return nil
}

func (r *Runner) runTask(task *model.Task) {
	// set task status to active
	// sendUpdate calls task.Update, so have to do this synchronously
	if err := r.sendUpdate(task, nil, nil, nil); err != nil {
		log.LogError(errors.Wrap(err, "failed to sendUpdate"))
		return
	}

	taskKeyJSON, err := r.localAuth.ActiveSession.Keypair.Decrypt(task.Meta.RunnerEncTaskKey)
	if err != nil {
		log.LogError(errors.Wrap(err, "failed to Decrypt task key"))
		return
	}

	taskKey, err := simplcrypto.SymKeyFromJSON(taskKeyJSON)
	if err != nil {
		log.LogError(errors.Wrap(err, "failed to SymKeyFromJSON"))
		return
	}

	taskBodyJSON, err := taskKey.Decrypt(task.EncBody)
	if err != nil {
		log.LogError(errors.Wrap(err, "failed to Decrypt task body"))
		return
	}

	var result interface{}
	var handlerErr error
	if r.handler != nil {
		result, handlerErr = r.handler(taskBodyJSON)
	} else if r.specHandler != nil {
		decryptedTask := DecryptedTask{
			Task: *task,
			Body: taskBodyJSON,
		}

		result, handlerErr = r.specHandler(&decryptedTask)
	} else {
		log.LogError(errors.New("handler and specHandler are nil, bailing out"))
		return
	}

	if handlerErr != nil {
		// sendUpdate calls task.Update
		if err := r.sendUpdate(task, taskKey, nil, handlerErr); err != nil {
			log.LogError(errors.Wrap(err, "failed to sendUpdate"))
		}

		return
	}

	// sendUpdate calls task.Update... just making sure you know.
	if err := r.sendUpdate(task, taskKey, result, nil); err != nil {
		log.LogError(errors.Wrap(err, "failed to sendUpdate"))
	}
}

func (r *Runner) sendUpdate(task *model.Task, taskKey *simplcrypto.SymKey, result interface{}, taskErr error) error {
	update := model.TaskUpdate{}

	if result == nil && taskErr == nil {
		update.Status = model.TaskStatusRunning
	} else {
		var encResult *simplcrypto.Message

		if result == nil && taskErr != nil {
			update.Status = model.TaskStatusFailed

			var err error
			encResult, err = taskKey.Encrypt([]byte(taskErr.Error()))
			if err != nil {
				return errors.Wrap(err, "failed to Encrypt error result")
			}
		} else if result != nil && taskErr == nil {
			update.Status = model.TaskStatusCompleted

			resultJSON, err := json.Marshal(result)
			if err != nil {
				return errors.Wrap(err, "failed to Marshal result")
			}

			encResult, err = taskKey.Encrypt(resultJSON)
			if err != nil {
				return errors.Wrap(err, "failed to Encrypt result")
			}
		}

		update.EncResult = encResult
	}

	realUpdate, err := task.Update(update)
	if err != nil {
		return errors.Wrap(err, "failed to task.Update")
	}

	req := &service.UpdateTaskRequest{
		Update:  &realUpdate,
		Session: r.localAuth.ActiveSession.Session,
	}

	if _, err := r.client.UpdateTask(context.Background(), req); err != nil {
		return errors.Wrap(err, "failed to UpdateTask")
	}

	return nil
}

func signedAuthHashAttempt(keypair *simplcrypto.KeyPair, joinCode string) (*simplcrypto.Signature, int64, error) {
	groupAuthHash := auth.GroupAuthHash(joinCode, "")

	now := time.Now().Unix()

	nonce := make([]byte, 8)
	binary.LittleEndian.PutUint64(nonce, uint64(now))
	hashWithNonce := append(groupAuthHash, nonce...)

	joinSig, err := keypair.Sign(hashWithNonce)
	if err != nil {
		return nil, 0, errors.Wrap(err, "failed to Sign")
	}

	return joinSig, now, nil
}
