package taask

import (
	"encoding/json"
	"io/ioutil"

	"github.com/pkg/errors"
	"github.com/taask/taask-server/model"
)

// DefaultTaskFilename and others are consts for reading task files
const (
	DefaultTaskFilename = "task.json"
)

func readTaskFile(filepath string) (*model.Task, error) {
	raw, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, errors.Wrap(err, "failed to ReadFile")
	}

	task := &model.Task{}
	if err := json.Unmarshal(raw, task); err != nil {
		return nil, errors.Wrap(err, errors.Wrap(err, "failed to yaml and json Unmarshal").Error()) // stupid, but whatever
	}

	return task, nil
}
