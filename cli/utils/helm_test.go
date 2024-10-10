package utils

import (
	"fmt"
	"testing"

	"github.com/smartcontractkit/crib/cli/wrappers"
	wrappermocks "github.com/smartcontractkit/crib/cli/wrappers/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestHelmRegistryLogin(t *testing.T) {
	t.Parallel()

	type testCases struct {
		description            string
		mockHelmRegistryClient wrappers.HelmRegistryAPI
		username               string
		password               string
		serverAddress          string
		wantError              string
	}

	mockHelmRegistryClient := wrappermocks.NewHelmRegistryAPI(t)
	mockHelmRegistryClient.EXPECT().
		Login(
			"registryURL", mock.AnythingOfType("registry.LoginOption"),
		).Return(nil)

	mockHelmRegistryClientFailed := wrappermocks.NewHelmRegistryAPI(t)
	mockHelmRegistryClientFailed.EXPECT().
		Login(
			"registryURL", mock.AnythingOfType("registry.LoginOption"),
		).Return(fmt.Errorf("login attempt failed with status: 400 Bad Request"))

	for _, scenario := range []testCases{
		{
			description:            "successful helm registry login",
			mockHelmRegistryClient: mockHelmRegistryClient,
			username:               "user",
			password:               "password",
			serverAddress:          "registryURL",
			wantError:              "",
		},
		{
			description:            "failed helm registry login",
			mockHelmRegistryClient: mockHelmRegistryClientFailed,
			username:               "user",
			password:               "wrongpassword",
			serverAddress:          "registryURL",
			wantError:              "login attempt failed",
		},
	} {
		t.Run(scenario.description, func(t *testing.T) {
			t.Parallel()

			err := HelmRegistryLogin(scenario.mockHelmRegistryClient, scenario.username, scenario.password, scenario.serverAddress)
			if scenario.wantError == "" {
				require.NoError(t, err)
			} else {
				assert.ErrorContains(t, err, scenario.wantError)
			}
		})
	}
}
