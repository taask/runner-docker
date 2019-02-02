package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	log "github.com/cohix/simplog"
	"github.com/pkg/errors"
	taask "github.com/taask/runner-golang"
)

const hostHomeDirKey = "TAASK_HOST_HOMEDIR"

// GetDockerTaskResult runs the specified docker image and returns the last line of output
func GetDockerTaskResult(image string, task *taask.DecryptedTask) (interface{}, error) {
	output, err := runDockerImageWithTask(image, task)
	if err != nil {
		log.LogError(errors.Wrap(errors.New(string(output)), "failed to runDockerImageWithTask"))
		return nil, errors.Wrap(err, "failed to runDockerImageWithTask")
	}

	outputString := string(output)

	log.LogInfo("container output:")
	fmt.Println(outputString)

	lines := strings.Split(outputString, "\n")

	var resultBytes []byte

	for i := len(lines) - 1; i >= 0; i-- {
		if lines[i] != "" {
			resultBytes = []byte(lines[i])
			break
		}
	}

	if resultBytes == nil {
		return nil, errors.New("container produced no output")
	}

	resultMap := make(map[string]interface{})
	if err := json.Unmarshal(resultBytes, &resultMap); err != nil {
		return nil, errors.Wrap(err, "failed to Marshal result")
	}

	return resultMap, nil
}

// runDockerImage runs a docker image and then returns its output
func runDockerImageWithTask(image string, task *taask.DecryptedTask) ([]byte, error) {
	if err := writeTaskBodyToTmpDir(task); err != nil {
		return nil, errors.Wrap(err, "failed to writeTaskBodyToTmpDir")
	}

	vol := volumeStringForUUID(task.UUID)
	cmd := exec.Command("docker", "run", "-v", vol, image)

	log.LogInfo(fmt.Sprintf("running docker image %s for task %s", image, task.UUID))
	return cmd.CombinedOutput()
}

// returns the tmpdir that the file was written to
func writeTaskBodyToTmpDir(task *taask.DecryptedTask) error {
	tmpDir := tmpDirForUUID(task.UUID)

	if err := os.Mkdir(tmpDir, 0777); err != nil {
		return errors.Wrap(err, "failed to Mkdir")
	}

	filepath := filepath.Join(tmpDir, "input")

	if err := ioutil.WriteFile(filepath, task.Body, 0777); err != nil {
		return errors.Wrap(err, "failed to WriteFile")
	}

	return nil
}

func tmpDirForUUID(uuid string) string {
	// this is the path in the runner container,
	// mapped to a tmp dir on the host by k8s/docker
	root := "/root/.taask/runner/data"

	return filepath.Join(root, uuid)
}

func volumeStringForUUID(uuid string) string {
	home := "/root"
	if hostHome, ok := os.LookupEnv(hostHomeDirKey); ok {
		home = hostHome
	}

	// this is the path on the _host_ (TODO: Make it configurable), that maps to the input for this task
	return fmt.Sprintf("%s/.taask/runner/data/%s:/root/.taask", home, uuid)
}
