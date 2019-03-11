package main

import (
	"flag"
	"os"
	"strings"

	log "github.com/cohix/simplog"
	"github.com/pkg/errors"
	"github.com/taask/runner-golang"
)

const (
	imageAnnotation = "io.taask.container.image"
)

var serverHost = flag.String("host", "taask-server-internal", "host for taask-server")
var serverPort = flag.String("port", "3687", "port for taask-server")

func main() {
	flag.Parse()

	runner, err := taask.NewSpecRunner("io.taask.docker", []string{}, func(task *taask.DecryptedTask) (interface{}, error) {
		image, err := getImageAnnotation(task.Meta.Annotations)
		if err != nil {
			log.LogError(err)
			return nil, errors.Wrap(err, "failed to getImageAnnotation")
		}

		result, err := GetDockerTaskResult(image, task)
		if err != nil {
			log.LogError(err)
			return nil, errors.Wrap(err, "failed to GetDockerTaskResult")
		}

		return result, nil
	})

	if err != nil {
		log.LogError(errors.Wrap(err, "failed to NewRunner"))
		os.Exit(1)
	}

	if err := runner.ConnectAndStreamTasks(*serverHost, *serverPort); err != nil {
		log.LogError(errors.Wrap(err, "failed to ConnectAndRun"))
		os.Exit(1)
	}
}

func getImageAnnotation(annotations []string) (string, error) {
	for _, a := range annotations {
		parts := strings.Split(a, ":")
		if len(parts) != 2 {
			return "", errors.New("couldn't parse image annotation")
		}

		ann := parts[0]
		val := parts[1]

		if ann == imageAnnotation {
			return val, nil
		}
	}

	return "", errors.New("image annotation not set")
}
