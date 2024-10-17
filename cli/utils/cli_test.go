package utils

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ecr"
	ecrtypes "github.com/aws/aws-sdk-go-v2/service/ecr/types"
	"github.com/docker/cli/cli/config/configfile"
	"github.com/docker/docker/api/types/registry"
	"github.com/go-git/go-git/v5"
	"github.com/smartcontractkit/crib/cli/wrappers"
	wrappermocks "github.com/smartcontractkit/crib/cli/wrappers/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestGetGitTopLevelDir(t *testing.T) {
	t.Parallel()

	// Create a temporary directory for testing
	tempDir := t.TempDir()

	// Initialize a new git repository in the temporary directory
	gitRepoInTempDir := filepath.Join(tempDir, "git_repo")
	_, err := git.PlainInit(gitRepoInTempDir, false)
	require.NoError(t, err)

	// Create a subdirectory within the git repository
	subDir := filepath.Join(gitRepoInTempDir, "subdir", "subsubdir")
	err = os.MkdirAll(subDir, 0o755)
	require.NoError(t, err)

	testCases := []struct {
		name        string
		dir         string
		expectedDir string
		expectedErr error
	}{
		{
			name:        "ValidGitRepository",
			dir:         gitRepoInTempDir,
			expectedDir: gitRepoInTempDir,
			expectedErr: nil,
		},
		{
			name:        "NonExistentDirectory",
			dir:         filepath.Join(tempDir, "nonexistent"),
			expectedDir: "",
			expectedErr: git.ErrRepositoryNotExists,
		},
		{
			name:        "DirectoryWithoutGitRepository",
			dir:         filepath.Join(tempDir, "without_git"),
			expectedDir: "",
			expectedErr: git.ErrRepositoryNotExists,
		},
		{
			name:        "SubdirectoryWithinGitRepository",
			dir:         subDir,
			expectedDir: gitRepoInTempDir,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			if tc.name == "DirectoryWithoutGitRepository" {
				require.NoError(t, os.Mkdir(tc.dir, 0o750))
			}

			topLevelDir, err := GetGitTopLevelDir(tc.dir)
			if tc.expectedErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedErr, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedDir, topLevelDir)
			}
		})
	}
}

func TestListFiles(t *testing.T) {
	t.Parallel()

	// Create a temporary directory for testing
	tempDir := t.TempDir()

	// Create some files in the temporary directory
	// Initialize files in a slice and create them in a loop
	files := []string{"file1.txt", "file2.txt", "file3.txt"}
	for _, file := range files {
		_, err := os.Create(filepath.Join(tempDir, file))
		require.NoError(t, err)
	}
	require.NoError(t, os.Mkdir(filepath.Join(tempDir, "subdir"), 0o750))

	testCases := []struct {
		name        string
		dir         string
		expected    []string
		expectedErr error
	}{
		{
			name:        "ValidDirectory",
			dir:         tempDir,
			expected:    []string{"file1.txt", "file2.txt", "file3.txt", "subdir"},
			expectedErr: nil,
		},
		{
			name:        "NonExistentDirectory",
			dir:         filepath.Join(t.TempDir(), "nonexistent"),
			expected:    nil,
			expectedErr: &os.PathError{},
		},
		{
			name:        "EmptyDirectory",
			dir:         filepath.Join(t.TempDir(), "empty"),
			expected:    []string{},
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			if tc.name == "EmptyDirectory" {
				require.NoError(t, os.Mkdir(tc.dir, 0o750))
			}

			files, err := ListFiles(tc.dir)
			if tc.expectedErr != nil {
				assert.Error(t, err)
				assert.IsType(t, tc.expectedErr, err)
			} else {
				assert.NoError(t, err)
				assert.ElementsMatch(t, tc.expected, files)
			}
		})
	}
}

// nolint: paralleltest,nolintlint
func TestPromptForInput(t *testing.T) {
	testCases := []struct {
		name         string
		key          string
		defaultValue string
		userInput    string
		expected     string
		expectedErr  error
	}{
		{
			name:         "UserProvidesInput",
			key:          "username",
			defaultValue: "defaultUser",
			userInput:    "testUser\n",
			expected:     "testUser",
			expectedErr:  nil,
		},
		{
			name:         "UserProvidesNoInput",
			key:          "username",
			defaultValue: "defaultUser",
			userInput:    "\n",
			expected:     "defaultUser",
			expectedErr:  nil,
		},
		{
			name:         "UserProvidesInputWithoutDefault",
			key:          "username",
			defaultValue: "",
			userInput:    "testUser\n",
			expected:     "testUser",
			expectedErr:  nil,
		},
		{
			name:         "UserProvidesNoInputWithoutDefault",
			key:          "username",
			defaultValue: "",
			userInput:    "\n",
			expected:     "",
			expectedErr:  nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Mock stdin
			oldStdin := os.Stdin
			defer func() { os.Stdin = oldStdin }()
			r, w, _ := os.Pipe()
			os.Stdin = r

			// Write user input to the pipe
			_, err := w.WriteString(tc.userInput)
			require.NoError(t, err)
			w.Close()

			result, err := PromptForInput(tc.key, tc.defaultValue)
			assert.Equal(t, tc.expected, result)
			if tc.expectedErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedErr.Error(), err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

// nolint: paralleltest,nolintlint
func TestPresentPrompt(t *testing.T) {
	testCases := []struct {
		name      string
		prompt    string
		choices   []string
		userInput string
		expected  string
	}{
		{
			name:      "ValidChoice",
			prompt:    "Choose an option",
			choices:   []string{"option1", "option2", "option3"},
			userInput: "option2\n",
			expected:  "option2",
		},
		{
			name:      "InvalidThenValidChoice",
			prompt:    "Choose an option",
			choices:   []string{"option1", "option2", "option3"},
			userInput: "invalid\noption3\n",
			expected:  "option3",
		},
		{
			name:      "WhitespaceInput",
			prompt:    "Choose an option",
			choices:   []string{"option1", "option2", "option3"},
			userInput: "  option1  \n",
			expected:  "option1",
		},
		{
			name:      "EmptyInputThenValidChoice",
			prompt:    "Choose an option",
			choices:   []string{"option1", "option2", "option3"},
			userInput: "\noption1\n",
			expected:  "option1",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Mock stdin
			oldStdin := os.Stdin
			defer func() { os.Stdin = oldStdin }()
			r, w, _ := os.Pipe()
			os.Stdin = r

			// Write user input to the pipe
			_, err := w.WriteString(tc.userInput)
			require.NoError(t, err)
			w.Close()

			result := PresentPrompt(tc.prompt, tc.choices)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestWriteConfig(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name          string
		initialLines  []string
		kv            map[string]string
		expectedLines []string
		expectedErr   error
	}{
		{
			name: "UpdateExistingKeys",
			initialLines: []string{
				"KEY1=oldvalue1",
				"KEY2=oldvalue2",
			},
			kv: map[string]string{
				"KEY1": "newvalue1",
				"KEY2": "newvalue2",
			},
			expectedLines: []string{
				"KEY1=newvalue1",
				"KEY2=newvalue2",
			},
			expectedErr: nil,
		},
		{
			name: "AddNewKeys",
			initialLines: []string{
				"KEY1=oldvalue1",
			},
			kv: map[string]string{
				"KEY2": "newvalue2",
				"KEY3": "newvalue3",
			},
			expectedLines: []string{
				"KEY1=oldvalue1",
				"# Added by CRIB CLI",
				"KEY2=newvalue2",
				"KEY3=newvalue3",
			},
			expectedErr: nil,
		},
		{
			name: "MixedUpdateAndAddKeys",
			initialLines: []string{
				"KEY1=oldvalue1",
				"KEY2=oldvalue2",
			},
			kv: map[string]string{
				"KEY2": "newvalue2",
				"KEY3": "newvalue3",
			},
			expectedLines: []string{
				"KEY1=oldvalue1",
				"KEY2=newvalue2",
				"# Added by CRIB CLI",
				"KEY3=newvalue3",
			},
			expectedErr: nil,
		},
		{
			name:         "EmptyFile",
			initialLines: []string{},
			kv: map[string]string{
				"KEY1": "newvalue1",
			},
			expectedLines: []string{
				"# Added by CRIB CLI",
				"KEY1=newvalue1",
			},
			expectedErr: nil,
		},
		{
			name: "NoChanges",
			initialLines: []string{
				"KEY1=oldvalue1",
			},
			kv: map[string]string{},
			expectedLines: []string{
				"KEY1=oldvalue1",
			},
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			// Create a temporary file for testing
			tempFile, err := os.CreateTemp("", "test.env")
			require.NoError(t, err)
			defer os.Remove(tempFile.Name())

			// Write initial lines to the temporary file
			for _, line := range tc.initialLines {
				_, err := tempFile.WriteString(line + "\n")
				require.NoError(t, err)
			}
			tempFile.Close()

			err = WriteConfig(tempFile.Name(), tc.kv)
			if tc.expectedErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedErr.Error(), err.Error())
			} else {
				assert.NoError(t, err)

				// Read the file and check the contents
				content, err := os.ReadFile(tempFile.Name())
				require.NoError(t, err)
				lines := strings.Split(strings.TrimSpace(string(content)), "\n")
				assert.ElementsMatch(t, tc.expectedLines, lines)
			}
		})
	}
}

func TestExtractHostFromUrl(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name        string
		input       string
		expected    string
		expectedErr error
	}{
		{
			name:        "ValidURLWithHost",
			input:       "https://example.com/path",
			expected:    "example.com",
			expectedErr: nil,
		},
		{
			name:        "ValidURLWithPort",
			input:       "https://example.com:8080/path",
			expected:    "example.com:8080",
			expectedErr: nil,
		},
		{
			name:        "ValidURLWithoutPath",
			input:       "https://example.com",
			expected:    "example.com",
			expectedErr: nil,
		},
		{
			name:        "InvalidURL",
			input:       "://invalid-url",
			expected:    "",
			expectedErr: &url.Error{},
		},
		{
			name:        "EmptyURL",
			input:       "",
			expected:    "",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			result, err := ExtractHostFromUrl(tc.input)
			if tc.expectedErr != nil {
				assert.Error(t, err)
				assert.IsType(t, tc.expectedErr, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, result)
			}
		})
	}
}

func TestRefreshRegistriesECRCredentials(t *testing.T) {
	t.Parallel()

	mockDockerRegistryHost := "https://012345678910.dkr.ecr.us-east-1.amazonaws.com"
	mockHelmRegistryHost := "oci://chainlink-helm-registry.com"
	ecrAuthToken := base64.StdEncoding.EncodeToString([]byte("user:password"))

	mockEcrClient := wrappermocks.NewECRAPI(t)
	mockEcrClient.EXPECT().
		GetAuthorizationToken(
			context.TODO(), &ecr.GetAuthorizationTokenInput{},
		).Return(
		&ecr.GetAuthorizationTokenOutput{
			AuthorizationData: []ecrtypes.AuthorizationData{
				{AuthorizationToken: &ecrAuthToken, ProxyEndpoint: &mockDockerRegistryHost},
			},
		}, nil,
	)
	mockEcrClientFailed := wrappermocks.NewECRAPI(t)
	mockEcrClientFailedErr := errors.New("ecr.GetAuthorizationToken failed")
	mockEcrClientFailed.EXPECT().
		GetAuthorizationToken(
			context.TODO(), &ecr.GetAuthorizationTokenInput{},
		).Return(nil, mockEcrClientFailedErr)

	mockDockerClient := wrappermocks.NewDockerAPI(t)
	mockDockerCli := wrappermocks.NewDockerCLI(t)
	mockDockerClient.EXPECT().
		RegistryLogin(
			context.TODO(), registry.AuthConfig{Username: "user", Password: "password", ServerAddress: strings.TrimPrefix(mockDockerRegistryHost, "https://")},
		).Return(registry.AuthenticateOKBody{IdentityToken: "", Status: "Login Succeeded"}, nil)
	mockDockerCli.EXPECT().Client().Return(mockDockerClient)
	mockDockerCli.EXPECT().ConfigFile().Return(configfile.New(filepath.Join(t.TempDir(), "config.json"))) // TODO: revisit

	mockDockerClientFailed := wrappermocks.NewDockerAPI(t)
	mockDockerCliFailed := wrappermocks.NewDockerCLI(t)
	mockDockerCliFailedErr := errors.New("failed to docker login")
	mockDockerClientFailed.EXPECT().
		RegistryLogin(
			context.TODO(), registry.AuthConfig{Username: "user", Password: "password", ServerAddress: strings.TrimPrefix(mockDockerRegistryHost, "https://")},
		).Return(registry.AuthenticateOKBody{}, mockDockerCliFailedErr)
	mockDockerCliFailed.EXPECT().Client().Return(mockDockerClientFailed)

	mockHelmRegistryClient := wrappermocks.NewHelmRegistryAPI(t)
	mockHelmRegistryClient.EXPECT().
		Login(
			strings.TrimPrefix(mockHelmRegistryHost, "oci://"), mock.AnythingOfType("registry.LoginOption"),
		).Return(nil)

	mockHelmRegistryClientFailed := wrappermocks.NewHelmRegistryAPI(t)
	mockHelmRegistryClientFailedErr := errors.New("failed to helm registry login")
	mockHelmRegistryClientFailed.EXPECT().
		Login(
			strings.TrimPrefix(mockHelmRegistryHost, "oci://"), mock.AnythingOfType("registry.LoginOption"),
		).Return(mockHelmRegistryClientFailedErr)

	testCases := []struct {
		name                     string
		mockEcrClient            wrappers.ECRAPI
		mockDockerCli            wrappers.DockerCLI
		mockHelmRegistryClient   wrappers.HelmRegistryAPI
		chainlinkHelmRegistryUri string
		wantOutput               *RefreshRegistriesECRCredentialsOutput
	}{
		{
			name:                     "SuccessfulRefreshDockerAndHelmRegistries",
			mockEcrClient:            mockEcrClient,
			mockDockerCli:            mockDockerCli,
			mockHelmRegistryClient:   mockHelmRegistryClient,
			chainlinkHelmRegistryUri: mockHelmRegistryHost,
			wantOutput: &RefreshRegistriesECRCredentialsOutput{
				RegistryLoginAttempts: &[]RegistryLoginAttempt{
					{
						RegistryType: "docker",
						RegistryHost: strings.TrimPrefix(mockDockerRegistryHost, "https://"),
						LoginErr:     nil,
					},
					{
						RegistryType: "helm",
						RegistryHost: strings.TrimPrefix(mockHelmRegistryHost, "oci://"),
						LoginErr:     nil,
					},
				},
			},
		},
		{
			name:                     "SuccessfulRefreshOnlyDockerRegistry",
			mockEcrClient:            mockEcrClient,
			mockDockerCli:            mockDockerCli,
			mockHelmRegistryClient:   nil,
			chainlinkHelmRegistryUri: "",
			wantOutput: &RefreshRegistriesECRCredentialsOutput{
				RegistryLoginAttempts: &[]RegistryLoginAttempt{
					{
						RegistryType: "docker",
						RegistryHost: strings.TrimPrefix(mockDockerRegistryHost, "https://"),
						LoginErr:     nil,
					},
				},
			},
		},
		{
			name:                     "SuccessfulRefreshOnlyHelmRegistry",
			mockEcrClient:            mockEcrClient,
			mockDockerCli:            nil,
			mockHelmRegistryClient:   mockHelmRegistryClient,
			chainlinkHelmRegistryUri: mockHelmRegistryHost,
			wantOutput: &RefreshRegistriesECRCredentialsOutput{
				RegistryLoginAttempts: &[]RegistryLoginAttempt{
					{
						RegistryType: "helm",
						RegistryHost: strings.TrimPrefix(mockHelmRegistryHost, "oci://"),
						LoginErr:     nil,
					},
				},
			},
		},
		{
			name:                     "NothingToRefresh",
			mockEcrClient:            mockEcrClient,
			mockDockerCli:            nil,
			mockHelmRegistryClient:   nil,
			chainlinkHelmRegistryUri: mockHelmRegistryHost, // shouldn't make a difference, as helmRegistryClient is nil
			wantOutput: &RefreshRegistriesECRCredentialsOutput{
				RegistryLoginAttempts: &[]RegistryLoginAttempt{},
			},
		},
		{
			name:                     "FailedRefreshDockerRegistry",
			mockEcrClient:            mockEcrClient,
			mockDockerCli:            mockDockerCliFailed,
			mockHelmRegistryClient:   mockHelmRegistryClient,
			chainlinkHelmRegistryUri: mockHelmRegistryHost,
			wantOutput: &RefreshRegistriesECRCredentialsOutput{
				RegistryLoginAttempts: &[]RegistryLoginAttempt{
					{
						RegistryType: "docker",
						RegistryHost: strings.TrimPrefix(mockDockerRegistryHost, "https://"),
						LoginErr:     fmt.Errorf("failed to docker login: %v", mockDockerCliFailedErr),
					},
					{
						RegistryType: "helm",
						RegistryHost: strings.TrimPrefix(mockHelmRegistryHost, "oci://"),
						LoginErr:     nil,
					},
				},
			},
		},
		{
			name:                     "FailedRefreshHelmRegistry",
			mockEcrClient:            mockEcrClient,
			mockDockerCli:            mockDockerCli,
			mockHelmRegistryClient:   mockHelmRegistryClientFailed,
			chainlinkHelmRegistryUri: mockHelmRegistryHost,
			wantOutput: &RefreshRegistriesECRCredentialsOutput{
				RegistryLoginAttempts: &[]RegistryLoginAttempt{
					{
						RegistryType: "docker",
						RegistryHost: strings.TrimPrefix(mockDockerRegistryHost, "https://"),
						LoginErr:     nil,
					},
					{
						RegistryType: "helm",
						RegistryHost: strings.TrimPrefix(mockHelmRegistryHost, "oci://"),
						LoginErr:     mockHelmRegistryClientFailedErr,
					},
				},
			},
		},
		{
			name:                     "FailedRefreshDockerAndHelmRegistries",
			mockEcrClient:            mockEcrClient,
			mockDockerCli:            mockDockerCliFailed,
			mockHelmRegistryClient:   mockHelmRegistryClientFailed,
			chainlinkHelmRegistryUri: mockHelmRegistryHost,
			wantOutput: &RefreshRegistriesECRCredentialsOutput{
				RegistryLoginAttempts: &[]RegistryLoginAttempt{
					{
						RegistryType: "docker",
						RegistryHost: strings.TrimPrefix(mockDockerRegistryHost, "https://"),
						LoginErr:     fmt.Errorf("failed to docker login: %v", mockDockerCliFailedErr),
					},
					{
						RegistryType: "helm",
						RegistryHost: strings.TrimPrefix(mockHelmRegistryHost, "oci://"),
						LoginErr:     mockHelmRegistryClientFailedErr,
					},
				},
			},
		},
		{
			name:                     "FailedRefreshDockerAndHelmRegistriesDueToEcrClientIssue",
			mockEcrClient:            mockEcrClientFailed,
			mockDockerCli:            mockDockerCli,
			mockHelmRegistryClient:   mockHelmRegistryClient,
			chainlinkHelmRegistryUri: mockHelmRegistryHost,
			wantOutput: &RefreshRegistriesECRCredentialsOutput{
				ECRGetAuthorizationTokenError: fmt.Errorf("unable to fetch ECR authorization token, %v", mockEcrClientFailedErr),
				RegistryLoginAttempts:         &[]RegistryLoginAttempt{},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tc.wantOutput, RefreshRegistriesECRCredentials(tc.mockEcrClient, tc.mockDockerCli, tc.mockHelmRegistryClient, tc.chainlinkHelmRegistryUri))
		})
	}
}
