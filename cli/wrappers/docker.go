package wrappers

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/docker/cli/cli/command"
	"github.com/docker/cli/cli/config/configfile"
	configtypes "github.com/docker/cli/cli/config/types"
	cliflags "github.com/docker/cli/cli/flags"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/api/types/registry"
	"github.com/docker/docker/client"
)

var (
	DefaultRunContainerTimeout    time.Duration = 10 * time.Second
	DefaultDeleteContainerTimeout time.Duration = 10 * time.Second
)

type DockerCLI interface {
	Client() client.APIClient
	ConfigFile() *configfile.ConfigFile
	Login(username string, password string, registryURL string) (*registry.AuthenticateOKBody, error)
	ConnectContainerToNetwork(ctx context.Context, containerID string, networkID string) (bool, error)
	RunContainer(ctx context.Context, config *container.Config, hostConfig *container.HostConfig, networkingConfig *network.NetworkingConfig, containerName string, forceRecreate bool, timeout *time.Duration) (bool, error)
	DeleteContainer(ctx context.Context, containerName string, removeOptions container.RemoveOptions, timeout *time.Duration) error
}

type DockerCli struct {
	Cli command.Cli
}

func NewDockerCli(ops ...command.CLIOption) (*DockerCli, error) {
	if len(ops) == 0 {
		ops = append(ops, command.WithInputStream(os.Stdin), command.WithOutputStream(os.Stdout), command.WithErrorStream(os.Stderr))
	}

	cli, err := command.NewDockerCli(ops...)
	if err != nil {
		return nil, err
	}

	clientOpts := cliflags.NewClientOptions()
	if err := cli.Initialize(clientOpts); err != nil {
		return nil, err
	}

	return &DockerCli{Cli: cli}, nil
}

func (c *DockerCli) Client() client.APIClient {
	return c.Cli.Client()
}

func (c *DockerCli) ConfigFile() *configfile.ConfigFile {
	return c.Cli.ConfigFile()
}

// Login performs docker login for a registry using the provided info, and storing the credentials for later use by the docker CLI outside the CRIB CLI.
// its a simplified copy of the unexposed method described here https://github.com/docker/cli/blob/b1ae218f605181100bc1bf65aebe456e7d85928b/cli/command/registry/login.go#L183-L204
func (c *DockerCli) Login(username string, password string, registryURL string) (*registry.AuthenticateOKBody, error) {
	authConfig := registry.AuthConfig{Username: username, Password: password, ServerAddress: registryURL}
	response, err := c.Cli.Client().RegistryLogin(context.TODO(), authConfig)
	if err != nil {
		return &response, fmt.Errorf("failed to docker login: %w", err)
	}

	if response.IdentityToken != "" {
		authConfig.Password = ""
		authConfig.IdentityToken = response.IdentityToken
	}

	creds := c.Cli.ConfigFile().GetCredentialsStore(authConfig.ServerAddress)
	if err := creds.Store(configtypes.AuthConfig(authConfig)); err != nil {
		return &response, fmt.Errorf("error saving credentials: %w", err)
	}

	return &response, nil
}

// RunContainer creates and starts a container with the given config, hostConfig and networkingConfig.
// If the container is already running, it returns true.
func (c *DockerCli) RunContainer(ctx context.Context, config *container.Config, hostConfig *container.HostConfig, networkingConfig *network.NetworkingConfig, containerName string, forceRecreate bool, timeout *time.Duration) (bool, error) {
	if timeout == nil {
		timeout = &DefaultRunContainerTimeout
	}

	alreadyRunning, err := c.isContainerRunning(ctx, containerName)
	if err != nil {
		return alreadyRunning, fmt.Errorf("could not check if container is running: %w", err)
	}

	if alreadyRunning {
		if !forceRecreate {
			return alreadyRunning, nil
		}
		err := c.DeleteContainer(ctx, containerName, container.RemoveOptions{Force: true}, nil)
		if err != nil {
			return alreadyRunning, fmt.Errorf("failed to delete container: %w", err)
		}
	}

	if networkingConfig == nil {
		networkingConfig = &network.NetworkingConfig{}
	}

	containerID := containerName
	if config != nil && hostConfig != nil {
		resp, err := c.Cli.Client().ContainerCreate(ctx, config, hostConfig, networkingConfig, nil, containerName)
		if err != nil {
			return alreadyRunning, fmt.Errorf("failed to create container: %w", err)
		}
		containerID = resp.ID
	}

	if err := c.Cli.Client().ContainerStart(ctx, containerID, container.StartOptions{}); err != nil {
		return alreadyRunning, fmt.Errorf("failed to start container: %w", err)
	}

	now := time.Now()
	for {
		if time.Since(now) > *timeout {
			return alreadyRunning, fmt.Errorf("timed out waiting for container %s to start", containerID)
		}

		containerJSON, err := c.inspectContainer(ctx, containerID)
		if err != nil {
			return alreadyRunning, fmt.Errorf("failed to inspect container: %w", err)
		}
		if containerJSON.State.Running {
			break
		}
	}

	return alreadyRunning, nil
}

// isContainerRunning checks if a container with the given name is running.
func (c *DockerCli) isContainerRunning(ctx context.Context, containerName string) (bool, error) {
	containerJSON, err := c.inspectContainer(ctx, containerName)
	if err != nil {
		if client.IsErrNotFound(err) {
			return false, nil
		}
		return false, err
	}

	return containerJSON.State.Running, nil
}

// inspectContainer returns the json resulting from docker inspect for a given container.
func (c *DockerCli) inspectContainer(ctx context.Context, containerName string) (types.ContainerJSON, error) {
	containerJSON, err := c.Cli.Client().ContainerInspect(ctx, containerName)
	if err != nil {
		return types.ContainerJSON{}, fmt.Errorf("unable to inspect container %s: %w", containerName, err)
	}

	return containerJSON, nil
}

// DeleteContainer deletes a container with the given name.
func (c *DockerCli) DeleteContainer(ctx context.Context, containerName string, removeOptions container.RemoveOptions, timeout *time.Duration) error {
	if timeout == nil {
		timeout = &DefaultDeleteContainerTimeout
	}

	containerCreated, err := c.isContainerCreated(ctx, containerName)
	if err != nil {
		return fmt.Errorf("could not check if container is created: %w", err)
	}
	if !containerCreated {
		// container doesn't exist, nothing to delete
		return nil
	}

	if !removeOptions.Force {
		containerRunning, err := c.isContainerRunning(ctx, containerName)
		if err != nil {
			return fmt.Errorf("could not check if container is running: %w", err)
		}

		if containerRunning {
			timeoutSeconds := int(timeout.Seconds())
			if err := c.Cli.Client().ContainerStop(ctx, containerName, container.StopOptions{Timeout: &timeoutSeconds}); err != nil {
				return fmt.Errorf("failed to stop container: %w", err)
			}
		}
	}

	return c.Cli.Client().ContainerRemove(ctx, containerName, removeOptions)
}

// isContainerCreated checks if a container with the given name is created.
func (c *DockerCli) isContainerCreated(ctx context.Context, containerName string) (bool, error) {
	if _, err := c.inspectContainer(ctx, containerName); err != nil {
		if client.IsErrNotFound(err) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

// ConnectContainerToNetwork connects a container to a network.
// If the container is already connected to the network, it returns true.
func (c *DockerCli) ConnectContainerToNetwork(ctx context.Context, containerID, networkID string) (bool, error) {
	containerJSON, err := c.inspectContainer(ctx, containerID)
	if err != nil {
		return false, err
	}

	if _, ok := containerJSON.NetworkSettings.Networks[networkID]; ok {
		return true, nil
	}

	if err := c.Cli.Client().NetworkConnect(ctx, networkID, containerID, &network.EndpointSettings{}); err != nil {
		return false, fmt.Errorf("failed to connect container %s to network %s: %w", containerID, networkID, err)
	}

	return false, nil
}
