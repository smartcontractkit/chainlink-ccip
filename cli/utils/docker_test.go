package utils

import (
	"context"
	"fmt"
	"path/filepath"
	"testing"

	"github.com/docker/cli/cli/config/configfile"
	"github.com/docker/docker/api/types/registry"
	"github.com/smartcontractkit/crib/cli/wrappers"
	wrappermocks "github.com/smartcontractkit/crib/cli/wrappers/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDockerLogin(t *testing.T) {
	type testCases struct {
		description   string
		mockDockerCli wrappers.DockerCLI
		username      string
		password      string
		serverAddress string
		wantOutput    registry.AuthenticateOKBody
		wantError     string
	}

	mockDockerClient := wrappermocks.NewDockerAPI(t)
	mockDockerCli := wrappermocks.NewDockerCLI(t)
	mockDockerClient.EXPECT().
		RegistryLogin(
			context.TODO(), registry.AuthConfig{Username: "user", Password: "password", ServerAddress: "registryURL"},
		).Return(registry.AuthenticateOKBody{IdentityToken: "", Status: "Login Succeeded"}, nil)
	mockDockerCli.EXPECT().Client().Return(mockDockerClient)
	mockDockerCli.EXPECT().ConfigFile().Return(configfile.New(filepath.Join(t.TempDir(), "config.json")))

	mockDockerClientFailed := wrappermocks.NewDockerAPI(t)
	mockDockerCliFailed := wrappermocks.NewDockerCLI(t)
	mockDockerClientFailed.EXPECT().
		RegistryLogin(
			context.TODO(), registry.AuthConfig{Username: "user", Password: "wrongpassword", ServerAddress: "registryURL"},
		).Return(registry.AuthenticateOKBody{}, fmt.Errorf("login attempt failed with status: 400 Bad Request"))
	mockDockerCliFailed.EXPECT().Client().Return(mockDockerClientFailed)

	for _, scenario := range []testCases{
		{
			description:   "successful docker login",
			mockDockerCli: mockDockerCli,
			username:      "user",
			password:      "password",
			serverAddress: "registryURL",
			wantOutput:    registry.AuthenticateOKBody{Status: "Login Succeeded"},
			wantError:     "",
		},
		{
			description:   "failed docker login",
			mockDockerCli: mockDockerCliFailed,
			username:      "user",
			password:      "wrongpassword",
			serverAddress: "registryURL",
			wantOutput:    registry.AuthenticateOKBody{},
			wantError:     "login attempt failed",
		},
	} {
		t.Run(scenario.description, func(t *testing.T) {
			got, err := DockerLogin(scenario.mockDockerCli, scenario.username, scenario.password, scenario.serverAddress)
			if scenario.wantError == "" {
				require.NoError(t, err)
			} else {
				assert.ErrorContains(t, err, scenario.wantError)
			}
			assert.Equal(t, &scenario.wantOutput, got)
		})
	}
}
