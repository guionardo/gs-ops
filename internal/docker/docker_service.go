package docker

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/docker/docker/client"
)

type DockerService struct {
	client *client.Client
}

const SocketPrefix = "unix://"

var ErrInvalidSocketFile = fmt.Errorf("socket file must be like %s/path/to/file.socket", SocketPrefix)

func NewDockerService(dockerSocketFile string) (service *DockerService, err error) {
	if err = validateDockerSocketfile(dockerSocketFile); err != nil {
		return
	}
	os.Setenv(client.EnvOverrideHost, dockerSocketFile)
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	service = &DockerService{
		client: cli,
	}
	return
}

func (s *DockerService) DockerRunning(ctx context.Context) error {
	info, err := s.client.Info(ctx)
	if err == nil {
		fmt.Printf("Docker Info: %v\n", info)
	}
	return err
}

func validateDockerSocketfile(dockerSocketFile string) (err error) {
	if _, err = os.Stat(strings.TrimPrefix(dockerSocketFile, SocketPrefix)); err != nil {
		err = ErrInvalidSocketFile
	}
	return
}
