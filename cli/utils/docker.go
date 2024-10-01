package utils

import (
	"context"
	"fmt"
	"os"

	"github.com/docker/cli/cli/command"
	configtypes "github.com/docker/cli/cli/config/types"
	cliflags "github.com/docker/cli/cli/flags"
	"github.com/docker/docker/api/types/registry"
	"github.com/smartcontractkit/crib/cli/wrappers"
)

func InitializeDockerCLI() (wrappers.DockerCLI, error) {
	cliWrapper, err := wrappers.NewDockerCliWrapper(
		command.WithInputStream(os.Stdin),
		command.WithOutputStream(os.Stdout),
		command.WithErrorStream(os.Stderr),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create Docker CLI: %v", err)
	}

	clientOpts := cliflags.NewClientOptions()
	err = cliWrapper.Initialize(clientOpts)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Docker CLI: %v", err)
	}
	return cliWrapper, nil
}

// DockerLogin performs docker login for a registry using the provided info, and storing the credentials for later use by the docker CLI outside the CRIB CLI.
// its a simplified copy of the unexposed method described here https://github.com/docker/cli/blob/b1ae218f605181100bc1bf65aebe456e7d85928b/cli/command/registry/login.go#L183-L204
func DockerLogin(dockerCli wrappers.DockerCLI, username string, password string, registryURL string) (*registry.AuthenticateOKBody, error) {
	authConfig := registry.AuthConfig{Username: username, Password: password, ServerAddress: registryURL}
	response, err := dockerCli.Client().RegistryLogin(context.TODO(), authConfig)
	if err != nil {
		return &response, fmt.Errorf("failed to docker login: %v", err)
	}

	if response.IdentityToken != "" {
		authConfig.Password = ""
		authConfig.IdentityToken = response.IdentityToken
	}

	creds := dockerCli.ConfigFile().GetCredentialsStore(authConfig.ServerAddress)
	if err := creds.Store(configtypes.AuthConfig(authConfig)); err != nil {
		return &response, fmt.Errorf("error saving credentials: %v", err)
	}

	return &response, nil
}
