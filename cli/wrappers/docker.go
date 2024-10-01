package wrappers

import (
	"context"

	"github.com/docker/cli/cli/command"
	"github.com/docker/cli/cli/config/configfile"
	cliflags "github.com/docker/cli/cli/flags"
	"github.com/docker/docker/api/types/registry"
)

type DockerAPI interface {
	RegistryLogin(ctx context.Context, auth registry.AuthConfig) (registry.AuthenticateOKBody, error)
}
type DockerCLI interface {
	Initialize(opts *cliflags.ClientOptions, ops ...command.CLIOption) error
	Client() DockerAPI
	ConfigFile() *configfile.ConfigFile
}

type DockerCliWrapper struct {
	cli *command.DockerCli
}

func NewDockerCliWrapper(ops ...command.CLIOption) (*DockerCliWrapper, error) {
	cli, err := command.NewDockerCli(ops...)
	if err != nil {
		return nil, err
	}

	return &DockerCliWrapper{cli: cli}, nil
}

func (c *DockerCliWrapper) Initialize(opts *cliflags.ClientOptions, ops ...command.CLIOption) error {
	return c.cli.Initialize(opts, ops...)
}

func (c *DockerCliWrapper) Client() DockerAPI {
	return c.cli.Client()
}

func (c *DockerCliWrapper) ConfigFile() *configfile.ConfigFile {
	return c.cli.ConfigFile()
}
