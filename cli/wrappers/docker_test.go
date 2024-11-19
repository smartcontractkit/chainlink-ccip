package wrappers_test

import (
	"context"
	"errors"
	"fmt"
	"path/filepath"
	"testing"
	"time"

	"github.com/docker/cli/cli/config/configfile"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/api/types/registry"
	"github.com/docker/docker/errdefs"
	dockermocks "github.com/smartcontractkit/crib/cli/mocks/external/docker"
	"github.com/smartcontractkit/crib/cli/wrappers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestDockerCli_Login(t *testing.T) {
	t.Parallel()

	username := "user"
	password := "password"
	serverAddress := "registryURL"

	testCases := []struct {
		name                    string
		applyAPIClientMockCalls func(m *dockermocks.APIClient)
		expectedOutput          registry.AuthenticateOKBody
		expectedErr             string
	}{
		{
			name: "Success",
			applyAPIClientMockCalls: func(m *dockermocks.APIClient) {
				m.EXPECT().
					RegistryLogin(
						context.TODO(), registry.AuthConfig{Username: "user", Password: "password", ServerAddress: "registryURL"},
					).Return(registry.AuthenticateOKBody{IdentityToken: "", Status: "Login Succeeded"}, nil)
			},
			expectedOutput: registry.AuthenticateOKBody{Status: "Login Succeeded"},
			expectedErr:    "",
		},
		{
			name: "Failed",
			applyAPIClientMockCalls: func(m *dockermocks.APIClient) {
				m.EXPECT().
					RegistryLogin(
						context.TODO(), registry.AuthConfig{Username: "user", Password: "password", ServerAddress: "registryURL"},
					).Return(registry.AuthenticateOKBody{}, fmt.Errorf("login attempt failed with status: 400 Bad Request"))
			},
			expectedOutput: registry.AuthenticateOKBody{},
			expectedErr:    "login attempt failed",
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockAPIClient := dockermocks.NewAPIClient(t)
			tt.applyAPIClientMockCalls(mockAPIClient)
			mockDockerCli := dockermocks.NewDockerCli(t)
			mockDockerCli.EXPECT().Client().Return(mockAPIClient)
			if tt.name == "Success" {
				mockDockerCli.EXPECT().ConfigFile().Return(configfile.New(filepath.Join(t.TempDir(), "config.json")))
			}

			dockerCli := &wrappers.DockerCli{Cli: mockDockerCli}

			got, err := dockerCli.Login(username, password, serverAddress)
			if tt.expectedErr == "" {
				require.NoError(t, err)
			} else {
				assert.ErrorContains(t, err, tt.expectedErr)
			}
			assert.Equal(t, &tt.expectedOutput, got)
		})
	}
}

func TestDockerCli_ConnectContainerToNetwork(t *testing.T) {
	t.Parallel()

	containerID := "test-container"
	networkID := "test-network"

	testCases := []struct {
		name                     string
		applyAPIClientMockCalls  func(m *dockermocks.APIClient)
		expectedErr              string
		expectedAlreadyConnected bool
	}{
		{
			name: "ContainerExistsNetworkExistsConnectionSucceeds",
			applyAPIClientMockCalls: func(m *dockermocks.APIClient) {
				m.EXPECT().
					ContainerInspect(
						context.TODO(), containerID,
					).Return(types.ContainerJSON{NetworkSettings: &types.NetworkSettings{}}, nil)
				m.EXPECT().
					NetworkConnect(
						context.TODO(), networkID, containerID, &network.EndpointSettings{},
					).Return(nil)
			},
			expectedErr:              "",
			expectedAlreadyConnected: false,
		},
		{
			name: "ContainerExistsNetworkExistsConnectionFails",
			applyAPIClientMockCalls: func(m *dockermocks.APIClient) {
				m.EXPECT().
					ContainerInspect(
						context.TODO(), containerID,
					).Return(types.ContainerJSON{NetworkSettings: &types.NetworkSettings{}}, nil)
				m.EXPECT().
					NetworkConnect(
						context.TODO(), networkID, containerID, &network.EndpointSettings{},
					).Return(errors.New("error connecting to network"))
			},
			expectedErr:              "error connecting to network",
			expectedAlreadyConnected: false,
		},
		{
			name: "ContainerDoesNotExist",
			applyAPIClientMockCalls: func(m *dockermocks.APIClient) {
				m.EXPECT().
					ContainerInspect(
						context.TODO(), containerID,
					).Return(types.ContainerJSON{}, errors.New("container does not exist"))
			},
			expectedErr:              "container does not exist",
			expectedAlreadyConnected: false,
		},
		{
			name: "AlreadyConnected",
			applyAPIClientMockCalls: func(m *dockermocks.APIClient) {
				m.EXPECT().
					ContainerInspect(
						context.TODO(), containerID,
					).Return(types.ContainerJSON{NetworkSettings: &types.NetworkSettings{Networks: map[string]*network.EndpointSettings{networkID: nil}}}, nil)
			},
			expectedErr:              "",
			expectedAlreadyConnected: true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockAPIClient := dockermocks.NewAPIClient(t)
			tt.applyAPIClientMockCalls(mockAPIClient)
			mockDockerCli := dockermocks.NewDockerCli(t)
			mockDockerCli.EXPECT().Client().Return(mockAPIClient)
			dockerCli := &wrappers.DockerCli{Cli: mockDockerCli}

			alreadyConnected, err := dockerCli.ConnectContainerToNetwork(context.TODO(), containerID, networkID)
			assert.Equal(t, tt.expectedAlreadyConnected, alreadyConnected)
			if tt.expectedErr != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestDockerCli_RunContainer(t *testing.T) {
	t.Parallel()

	containerName := "test-container"
	config := &container.Config{}
	hostConfig := &container.HostConfig{}
	networkingConfig := &network.NetworkingConfig{}
	timeout := 1 * time.Second

	testCases := []struct {
		name                    string
		forceRecreate           bool
		applyAPIClientMockCalls func(m *dockermocks.APIClient)
		expectedErr             string
		expectedAlreadyRunning  bool
	}{
		{
			name:          "AlreadyRunning",
			forceRecreate: false,
			applyAPIClientMockCalls: func(m *dockermocks.APIClient) {
				m.EXPECT().
					ContainerInspect(
						context.TODO(), containerName,
					).Return(types.ContainerJSON{
					ContainerJSONBase: &types.ContainerJSONBase{
						State: &types.ContainerState{
							Running: true,
						},
					},
				}, nil)
			},
			expectedErr:            "",
			expectedAlreadyRunning: true,
		},
		{
			name:          "Success",
			forceRecreate: false,
			applyAPIClientMockCalls: func(m *dockermocks.APIClient) {
				m.EXPECT().
					ContainerInspect(
						context.TODO(), containerName,
					).Return(types.ContainerJSON{}, errdefs.NotFound(errors.New("container not found"))).Times(1)
				m.EXPECT().
					ContainerCreate(
						context.TODO(), config, hostConfig, networkingConfig, mock.Anything, containerName,
					).Return(container.CreateResponse{ID: containerName}, nil)
				m.EXPECT().
					ContainerStart(
						context.TODO(), containerName, container.StartOptions{},
					).Return(nil)
				m.EXPECT().
					ContainerInspect(
						context.TODO(), containerName,
					).Return(types.ContainerJSON{
					ContainerJSONBase: &types.ContainerJSONBase{
						State: &types.ContainerState{
							Running: true,
						},
					},
				}, nil).Times(1)
			},
			expectedErr:            "",
			expectedAlreadyRunning: false,
		},
		{
			name:          "CreateContainerFails",
			forceRecreate: false,
			applyAPIClientMockCalls: func(m *dockermocks.APIClient) {
				m.EXPECT().
					ContainerInspect(
						context.TODO(), containerName,
					).Return(types.ContainerJSON{}, errdefs.NotFound(errors.New("container not found"))).Times(1)
				m.EXPECT().
					ContainerCreate(
						context.TODO(), config, hostConfig, networkingConfig, mock.Anything, containerName,
					).Return(container.CreateResponse{}, errors.New("error creating container"))
			},
			expectedErr:            "error creating container",
			expectedAlreadyRunning: false,
		},
		{
			name:          "StartContainerFails",
			forceRecreate: false,
			applyAPIClientMockCalls: func(m *dockermocks.APIClient) {
				m.EXPECT().
					ContainerInspect(
						context.TODO(), containerName,
					).Return(types.ContainerJSON{}, errdefs.NotFound(errors.New("container not found"))).Times(1)
				m.EXPECT().
					ContainerCreate(
						context.TODO(), config, hostConfig, networkingConfig, mock.Anything, containerName,
					).Return(container.CreateResponse{ID: containerName}, nil)
				m.EXPECT().
					ContainerStart(
						context.TODO(), containerName, container.StartOptions{},
					).Return(errors.New("error starting container"))
			},
			expectedErr:            "error starting container",
			expectedAlreadyRunning: false,
		},
		{
			name:          "TimedOutWaitingForContainerToStart",
			forceRecreate: false,
			applyAPIClientMockCalls: func(m *dockermocks.APIClient) {
				m.EXPECT().
					ContainerInspect(
						context.TODO(), containerName,
					).Return(types.ContainerJSON{}, errdefs.NotFound(errors.New("container not found"))).Times(1)
				m.EXPECT().
					ContainerCreate(
						context.TODO(), config, hostConfig, networkingConfig, mock.Anything, containerName,
					).Return(container.CreateResponse{ID: containerName}, nil)
				m.EXPECT().
					ContainerStart(
						context.TODO(), containerName, container.StartOptions{},
					).Return(nil)
				m.EXPECT().
					ContainerInspect(
						context.TODO(), containerName,
					).Return(types.ContainerJSON{
					ContainerJSONBase: &types.ContainerJSONBase{
						State: &types.ContainerState{
							Running: false,
						},
					},
				}, nil)
			},
			expectedErr:            "timed out waiting for container test-container to start",
			expectedAlreadyRunning: false,
		},
		{
			name:          "ForceRecreateSucceeds",
			forceRecreate: true,
			applyAPIClientMockCalls: func(m *dockermocks.APIClient) {
				m.EXPECT().
					ContainerInspect(
						context.TODO(), containerName,
					).Return(types.ContainerJSON{
					ContainerJSONBase: &types.ContainerJSONBase{
						State: &types.ContainerState{
							Running: true,
						},
					},
				}, nil)
				m.EXPECT().
					ContainerRemove(
						context.TODO(), containerName, container.RemoveOptions{Force: true},
					).Return(nil)
				m.EXPECT().
					ContainerCreate(
						context.TODO(), config, hostConfig, networkingConfig, mock.Anything, containerName,
					).Return(container.CreateResponse{ID: containerName}, nil)
				m.EXPECT().
					ContainerStart(
						context.TODO(), containerName, container.StartOptions{},
					).Return(nil)
			},
			expectedErr:            "",
			expectedAlreadyRunning: true,
		},
		{
			name:          "ForceRecreateFails",
			forceRecreate: true,
			applyAPIClientMockCalls: func(m *dockermocks.APIClient) {
				m.EXPECT().
					ContainerInspect(
						context.TODO(), containerName,
					).Return(types.ContainerJSON{
					ContainerJSONBase: &types.ContainerJSONBase{
						State: &types.ContainerState{
							Running: true,
						},
					},
				}, nil)
				m.EXPECT().
					ContainerRemove(
						context.TODO(), containerName, container.RemoveOptions{Force: true},
					).Return(nil)
				m.EXPECT().
					ContainerCreate(
						context.TODO(), config, hostConfig, networkingConfig, mock.Anything, containerName,
					).Return(container.CreateResponse{ID: containerName}, nil)
				m.EXPECT().
					ContainerStart(
						context.TODO(), containerName, container.StartOptions{},
					).Return(errors.New("error starting container"))
			},
			expectedErr:            "error starting container",
			expectedAlreadyRunning: true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockAPIClient := dockermocks.NewAPIClient(t)
			tt.applyAPIClientMockCalls(mockAPIClient)
			mockDockerCli := dockermocks.NewDockerCli(t)
			mockDockerCli.EXPECT().Client().Return(mockAPIClient)
			dockerCli := &wrappers.DockerCli{Cli: mockDockerCli}

			alreadyRunning, err := dockerCli.RunContainer(context.TODO(), config, hostConfig, networkingConfig, containerName, tt.forceRecreate, &timeout)
			assert.Equal(t, tt.expectedAlreadyRunning, alreadyRunning)
			if tt.expectedErr != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestDockerCli_DeleteContainer(t *testing.T) {
	t.Parallel()

	containerName := "test-container"
	timeout := 1 * time.Second
	timeoutSeconds := int(timeout.Seconds())

	testCases := []struct {
		name                    string
		force                   bool
		applyAPIClientMockCalls func(m *dockermocks.APIClient)
		expectedErr             string
	}{
		{
			name:  "ContainerRunningDeleteSucceeds",
			force: false,
			applyAPIClientMockCalls: func(m *dockermocks.APIClient) {
				m.EXPECT().
					ContainerInspect(
						context.TODO(), containerName,
					).Return(types.ContainerJSON{
					ContainerJSONBase: &types.ContainerJSONBase{
						State: &types.ContainerState{
							Running: true,
						},
					},
				}, nil)
				m.EXPECT().
					ContainerStop(
						context.TODO(), containerName, container.StopOptions{Timeout: &timeoutSeconds},
					).Return(nil)
				m.EXPECT().
					ContainerRemove(
						context.TODO(), containerName, container.RemoveOptions{},
					).Return(nil)
			},
			expectedErr: "",
		},
		{
			name:  "ContainerStoppedDeleteSucceeds",
			force: false,
			applyAPIClientMockCalls: func(m *dockermocks.APIClient) {
				m.EXPECT().
					ContainerInspect(
						context.TODO(), containerName,
					).Return(types.ContainerJSON{
					ContainerJSONBase: &types.ContainerJSONBase{
						State: &types.ContainerState{
							Running: false,
						},
					},
				}, nil)
				m.EXPECT().
					ContainerRemove(
						context.TODO(), containerName, container.RemoveOptions{},
					).Return(nil)
			},
			expectedErr: "",
		},
		{
			name:  "ErrorStoppingContainer",
			force: false,
			applyAPIClientMockCalls: func(m *dockermocks.APIClient) {
				m.EXPECT().
					ContainerInspect(
						context.TODO(), containerName,
					).Return(types.ContainerJSON{
					ContainerJSONBase: &types.ContainerJSONBase{
						State: &types.ContainerState{
							Running: true,
						},
					},
				}, nil)
				m.EXPECT().
					ContainerStop(
						context.TODO(), containerName, container.StopOptions{Timeout: &timeoutSeconds},
					).Return(errors.New("error stopping container"))
			},
			expectedErr: "error stopping container",
		},
		{
			name:  "ErrorDeleteContainer",
			force: false,
			applyAPIClientMockCalls: func(m *dockermocks.APIClient) {
				m.EXPECT().
					ContainerInspect(
						context.TODO(), containerName,
					).Return(types.ContainerJSON{
					ContainerJSONBase: &types.ContainerJSONBase{
						State: &types.ContainerState{
							Running: false,
						},
					},
				}, nil)
				m.EXPECT().
					ContainerRemove(
						context.TODO(), containerName, container.RemoveOptions{},
					).Return(errors.New("error deleting container"))
			},
			expectedErr: "error deleting container",
		},
		{
			name:  "ContainerDoesNotExist",
			force: false,
			applyAPIClientMockCalls: func(m *dockermocks.APIClient) {
				m.EXPECT().
					ContainerInspect(
						context.TODO(), containerName,
					).Return(types.ContainerJSON{}, errdefs.NotFound(errors.New("container not found")))
			},
			expectedErr: "",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockAPIClient := dockermocks.NewAPIClient(t)
			tt.applyAPIClientMockCalls(mockAPIClient)
			mockDockerCli := dockermocks.NewDockerCli(t)
			mockDockerCli.EXPECT().Client().Return(mockAPIClient)
			dockerCli := &wrappers.DockerCli{Cli: mockDockerCli}

			err := dockerCli.DeleteContainer(context.TODO(), containerName, container.RemoveOptions{Force: tt.force}, &timeout)
			if tt.expectedErr != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
