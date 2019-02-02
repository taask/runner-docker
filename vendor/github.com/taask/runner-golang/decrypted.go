package taask

import "github.com/taask/taask-server/model"

// DecryptedTask represents a decrypted task
type DecryptedTask struct {
	model.Task
	Body []byte
}
