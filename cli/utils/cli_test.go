package utils_test

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/ecr"
	ecrtypes "github.com/aws/aws-sdk-go-v2/service/ecr/types"
	"github.com/docker/docker/api/types/registry"
	"github.com/go-git/go-git/v5"
	k8smocks "github.com/smartcontractkit/crib/cli/mocks/external/kubernetes"
	"github.com/smartcontractkit/crib/cli/utils"
	"github.com/smartcontractkit/crib/cli/wrappers"
	wrappermocks "github.com/smartcontractkit/crib/cli/wrappers/mocks"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	networkingv1api "k8s.io/api/networking/v1"
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

			topLevelDir, err := utils.GetGitTopLevelDir(tc.dir)
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

			files, err := utils.ListFiles(tc.dir)
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
		{
			name:      "SimplyReadUserInput",
			prompt:    "Please enter a value: ",
			choices:   []string{},
			userInput: "whatever\n",
			expected:  "whatever",
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

			result := utils.PresentPrompt(tc.prompt, tc.choices)
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

			err = utils.WriteConfig(tempFile.Name(), tc.kv)
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

			result, err := utils.ExtractHostFromUrl(tc.input)
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

	mockDockerRegistryHost := "012345678910.dkr.ecr.us-east-1.amazonaws.com"
	mockDockerRegistries := map[string]string{
		"mocked-env": mockDockerRegistryHost,
	}
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

	mockDockerCli := wrappermocks.NewDockerCLI(t)
	mockDockerCli.EXPECT().
		Login(
			"user", "password", strings.TrimPrefix(mockDockerRegistryHost, "https://"),
		).Return(&registry.AuthenticateOKBody{IdentityToken: "", Status: "Login Succeeded"}, nil)

	mockDockerCliFailed := wrappermocks.NewDockerCLI(t)
	mockDockerCliFailedErr := errors.New("failed to docker login")
	mockDockerCliFailed.EXPECT().
		Login(
			"user", "password", strings.TrimPrefix(mockDockerRegistryHost, "https://"),
		).Return(&registry.AuthenticateOKBody{}, mockDockerCliFailedErr)

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
		wantOutput               *utils.RefreshRegistriesECRCredentialsOutput
	}{
		{
			name:                     "SuccessfulRefreshDockerAndHelmRegistries",
			mockEcrClient:            mockEcrClient,
			mockDockerCli:            mockDockerCli,
			mockHelmRegistryClient:   mockHelmRegistryClient,
			chainlinkHelmRegistryUri: mockHelmRegistryHost,
			wantOutput: &utils.RefreshRegistriesECRCredentialsOutput{
				RegistryLoginAttempts: &[]utils.RegistryLoginAttempt{
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
			wantOutput: &utils.RefreshRegistriesECRCredentialsOutput{
				RegistryLoginAttempts: &[]utils.RegistryLoginAttempt{
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
			wantOutput: &utils.RefreshRegistriesECRCredentialsOutput{
				RegistryLoginAttempts: &[]utils.RegistryLoginAttempt{
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
			wantOutput: &utils.RefreshRegistriesECRCredentialsOutput{
				RegistryLoginAttempts: &[]utils.RegistryLoginAttempt{},
			},
		},
		{
			name:                     "FailedRefreshDockerRegistry",
			mockEcrClient:            mockEcrClient,
			mockDockerCli:            mockDockerCliFailed,
			mockHelmRegistryClient:   mockHelmRegistryClient,
			chainlinkHelmRegistryUri: mockHelmRegistryHost,
			wantOutput: &utils.RefreshRegistriesECRCredentialsOutput{
				RegistryLoginAttempts: &[]utils.RegistryLoginAttempt{
					{
						RegistryType: "docker",
						RegistryHost: strings.TrimPrefix(mockDockerRegistryHost, "https://"),
						LoginErr:     mockDockerCliFailedErr,
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
			wantOutput: &utils.RefreshRegistriesECRCredentialsOutput{
				RegistryLoginAttempts: &[]utils.RegistryLoginAttempt{
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
			wantOutput: &utils.RefreshRegistriesECRCredentialsOutput{
				RegistryLoginAttempts: &[]utils.RegistryLoginAttempt{
					{
						RegistryType: "docker",
						RegistryHost: strings.TrimPrefix(mockDockerRegistryHost, "https://"),
						LoginErr:     mockDockerCliFailedErr,
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
			wantOutput: &utils.RefreshRegistriesECRCredentialsOutput{
				ECRGetAuthorizationTokenError: fmt.Errorf("unable to fetch ECR authorization token, %v", mockEcrClientFailedErr),
				RegistryLoginAttempts:         &[]utils.RegistryLoginAttempt{},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tc.wantOutput, utils.RefreshRegistriesECRCredentials(tc.mockEcrClient, tc.mockDockerCli, tc.mockHelmRegistryClient, tc.chainlinkHelmRegistryUri, mockDockerRegistries))
		})
	}
}

func TestIsValidCribNamespace(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name            string
		namespace       string
		provider        string
		skipPrefixCheck bool
		expectedErr     error
	}{
		{
			name:            "ValidNamespaceWithPrefix",
			namespace:       "crib-validnamespace",
			provider:        "aws",
			skipPrefixCheck: false,
			expectedErr:     nil,
		},
		{
			name:            "InvalidNamespaceWithoutPrefix",
			namespace:       "invalidnamespace",
			provider:        "aws",
			skipPrefixCheck: false,
			expectedErr:     fmt.Errorf("DEVSPACE_NAMESPACE must begin with 'crib-' prefix"),
		},
		{
			name:            "ValidNamespaceWithoutPrefixButSkipCheck",
			namespace:       "validnamespace",
			provider:        "aws",
			skipPrefixCheck: true,
			expectedErr:     nil,
		},
		{
			name:            "EmptyNamespaceWithPrefixCheck",
			namespace:       "",
			provider:        "aws",
			skipPrefixCheck: false,
			expectedErr:     fmt.Errorf("DEVSPACE_NAMESPACE must begin with 'crib-' prefix"),
		},
		{
			name:            "EmptyNamespaceWithoutPrefixCheck",
			namespace:       "",
			provider:        "aws",
			skipPrefixCheck: true,
			expectedErr:     nil,
		},
		{
			name:            "ValidNamespaceForKindProvider",
			namespace:       "crib-local",
			provider:        "kind",
			skipPrefixCheck: false,
			expectedErr:     nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			err := utils.IsValidCribNamespace(tc.namespace, tc.provider, tc.skipPrefixCheck)
			if tc.expectedErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedErr.Error(), err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestEnsureCribNamespaceReady(t *testing.T) {
	t.Parallel()

	defaultWaitTimeout := 3 * time.Second
	defaultsleepBetweenAttempts := 100 * time.Millisecond
	// Set only the required environment variables
	viper.Set("CHAINLINK_TEAM", "CRIB")
	viper.Set("CHAINLINK_PRODUCT", "TestProduct")
	viper.Set("CHAINLINK_COMPONENT", "CRIB")
	viper.Set("CHAINLINK_COST_CENTER", "CRIB")

	testCases := []struct {
		name                    string
		namespace               string
		provider                string
		waitTimeout             *time.Duration
		sleepBetweenAttempts    *time.Duration
		skipRoleBindingCheck    bool
		applyK8sClientMockCalls func(*wrappermocks.K8sCLI)
		expectedErr             error
	}{
		{
			name:                 "NamespaceExistsAWSProvider",
			namespace:            "crib-test",
			provider:             "aws",
			waitTimeout:          &defaultWaitTimeout,
			sleepBetweenAttempts: &defaultsleepBetweenAttempts,
			skipRoleBindingCheck: false,
			applyK8sClientMockCalls: func(m *wrappermocks.K8sCLI) {
				m.EXPECT().
					EnsureNamespaceExists(context.TODO(), "crib-test").
					Return(true, nil)
				m.EXPECT().
					WaitForResource(context.TODO(), mock.Anything, "crib-test-crib-poweruser", defaultsleepBetweenAttempts, defaultWaitTimeout).
					Return(nil)
				m.EXPECT().
					LabelNamespace(context.TODO(), "crib-test", map[string]string{"chain.link/component": "CRIB", "chain.link/cost-center": "CRIB", "chain.link/product": "TestProduct", "chain.link/team": "CRIB"}).
					Return(nil)
			},
			expectedErr: nil,
		},
		{
			name:                 "NamespaceDoesNotExistAWSProvider",
			namespace:            "crib-test",
			provider:             "aws",
			waitTimeout:          &defaultWaitTimeout,
			sleepBetweenAttempts: &defaultsleepBetweenAttempts,
			skipRoleBindingCheck: false,
			applyK8sClientMockCalls: func(m *wrappermocks.K8sCLI) {
				m.EXPECT().
					EnsureNamespaceExists(context.TODO(), "crib-test").
					Return(false, nil)
				m.EXPECT().
					WaitForResource(context.TODO(), mock.Anything, "crib-test-crib-poweruser", defaultsleepBetweenAttempts, defaultWaitTimeout).
					Return(nil)
				m.EXPECT().
					LabelNamespace(context.TODO(), "crib-test", map[string]string{"chain.link/component": "CRIB", "chain.link/cost-center": "CRIB", "chain.link/product": "TestProduct", "chain.link/team": "CRIB"}).
					Return(nil)
			},
			expectedErr: nil,
		},
		{
			name:                 "NamespaceExistsKindProvider",
			namespace:            "crib-test",
			provider:             "kind",
			waitTimeout:          &defaultWaitTimeout,
			sleepBetweenAttempts: &defaultsleepBetweenAttempts,
			skipRoleBindingCheck: false,
			applyK8sClientMockCalls: func(m *wrappermocks.K8sCLI) {
				m.EXPECT().
					EnsureNamespaceExists(context.TODO(), "crib-test").
					Return(true, nil)
			},
			expectedErr: nil,
		},
		{
			name:                 "NamespaceCreationFails",
			namespace:            "crib-test",
			provider:             "aws",
			waitTimeout:          &defaultWaitTimeout,
			sleepBetweenAttempts: &defaultsleepBetweenAttempts,
			skipRoleBindingCheck: false,
			applyK8sClientMockCalls: func(m *wrappermocks.K8sCLI) {
				m.EXPECT().
					EnsureNamespaceExists(context.TODO(), "crib-test").
					Return(false, errors.New("failed to create namespace crib-test: failed to create namespace"))
			},
			expectedErr: fmt.Errorf("failed to ensure namespace existence: failed to create namespace crib-test: %w", errors.New("failed to create namespace")),
		},
		{
			name:                 "RoleBindingCreationFails",
			namespace:            "crib-test",
			provider:             "aws",
			waitTimeout:          &defaultWaitTimeout,
			sleepBetweenAttempts: &defaultsleepBetweenAttempts,
			skipRoleBindingCheck: false,
			applyK8sClientMockCalls: func(m *wrappermocks.K8sCLI) {
				m.EXPECT().
					EnsureNamespaceExists(context.TODO(), "crib-test").
					Return(false, nil)
				m.EXPECT().
					WaitForResource(context.TODO(), mock.Anything, "crib-test-crib-poweruser", defaultsleepBetweenAttempts, defaultWaitTimeout).
					Return(errors.New("timed out waiting for resource crib-test-crib-poweruser"))
			},
			expectedErr: fmt.Errorf("failed to wait for crib-power-user role binding to be created: %w", errors.New("timed out waiting for resource crib-test-crib-poweruser")),
		},
		{
			name:                 "LabelNamespaceFails",
			namespace:            "crib-test",
			provider:             "aws",
			waitTimeout:          &defaultWaitTimeout,
			sleepBetweenAttempts: &defaultsleepBetweenAttempts,
			skipRoleBindingCheck: false,
			applyK8sClientMockCalls: func(m *wrappermocks.K8sCLI) {
				m.EXPECT().
					EnsureNamespaceExists(context.TODO(), "crib-test").
					Return(false, nil)
				m.EXPECT().
					WaitForResource(context.TODO(), mock.Anything, "crib-test-crib-poweruser", defaultsleepBetweenAttempts, defaultWaitTimeout).
					Return(nil)
				m.EXPECT().
					LabelNamespace(context.TODO(), "crib-test", map[string]string{"chain.link/component": "CRIB", "chain.link/cost-center": "CRIB", "chain.link/product": "TestProduct", "chain.link/team": "CRIB"}).
					Return(errors.New("LabelNamespace failed"))
			},
			expectedErr: errors.New("failed to label namespace: LabelNamespace failed"),
		},
		{
			name:                 "SkipRoleBindingCheck",
			namespace:            "crib-test",
			provider:             "aws",
			waitTimeout:          &defaultWaitTimeout,
			sleepBetweenAttempts: &defaultsleepBetweenAttempts,
			skipRoleBindingCheck: true,
			applyK8sClientMockCalls: func(m *wrappermocks.K8sCLI) {
				m.EXPECT().
					EnsureNamespaceExists(context.TODO(), "crib-test").
					Return(false, nil)
				m.EXPECT().
					LabelNamespace(context.TODO(), "crib-test", map[string]string{"chain.link/component": "CRIB", "chain.link/cost-center": "CRIB", "chain.link/product": "TestProduct", "chain.link/team": "CRIB"}).
					Return(nil)
			},
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			mockK8sClient := wrappermocks.NewK8sCLI(t)
			tc.applyK8sClientMockCalls(mockK8sClient)
			err := utils.EnsureCribNamespaceReady(context.TODO(), mockK8sClient, k8smocks.NewResourceInterface(t), tc.namespace, tc.provider, tc.waitTimeout, tc.sleepBetweenAttempts, tc.skipRoleBindingCheck)
			if tc.expectedErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedErr.Error(), err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

// nolint: paralleltest, nolintlint
func TestPrintIngressHosts(t *testing.T) {
	namespace := "test-namespace"

	testCases := []struct {
		name                    string
		applyK8sClientMockCalls func(*wrappermocks.K8sCLI)
		provider                string
		expectedErr             string
		expectedHosts           []string
	}{
		{
			name:     "ValidIngressesProviderAWS",
			provider: "aws",
			applyK8sClientMockCalls: func(m *wrappermocks.K8sCLI) {
				m.EXPECT().
					ListIngresses(context.TODO(), namespace).
					Return(&networkingv1api.IngressList{Items: []networkingv1api.Ingress{
						{
							Spec: networkingv1api.IngressSpec{
								Rules: []networkingv1api.IngressRule{
									{Host: "example.com"},
									{Host: "test.com"},
								},
							},
						},
					}}, nil)
			},
			expectedErr:   "",
			expectedHosts: []string{"https://example.com", "https://test.com"},
		},
		{
			name:     "ValidIngressesProviderKind",
			provider: "kind",
			applyK8sClientMockCalls: func(m *wrappermocks.K8sCLI) {
				m.EXPECT().
					ListIngresses(context.TODO(), namespace).
					Return(&networkingv1api.IngressList{Items: []networkingv1api.Ingress{
						{
							Spec: networkingv1api.IngressSpec{
								Rules: []networkingv1api.IngressRule{
									{Host: "example.com"},
									{Host: "test.com"},
								},
							},
						},
					}}, nil)
			},
			expectedErr:   "",
			expectedHosts: []string{"http://example.com", "http://test.com"},
		},
		{
			name:     "NoIngresses",
			provider: string(mock.AnythingOfType("string")),
			applyK8sClientMockCalls: func(m *wrappermocks.K8sCLI) {
				m.EXPECT().
					ListIngresses(context.TODO(), namespace).
					Return(&networkingv1api.IngressList{}, nil)
			},
			expectedErr:   "",
			expectedHosts: []string{},
		},
		{
			name:     "ListIngressesError",
			provider: string(mock.AnythingOfType("string")),
			applyK8sClientMockCalls: func(m *wrappermocks.K8sCLI) {
				m.EXPECT().
					ListIngresses(context.TODO(), namespace).
					Return(&networkingv1api.IngressList{}, fmt.Errorf("error listing ingresses"))
			},
			expectedErr:   "error listing ingresses",
			expectedHosts: []string{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockK8sClient := wrappermocks.NewK8sCLI(t)
			tc.applyK8sClientMockCalls(mockK8sClient)

			// Capture the output
			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			err := utils.PrintIngressHosts(context.TODO(), mockK8sClient, namespace, tc.provider)
			if tc.expectedErr != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedErr)
			} else {
				assert.NoError(t, err)
			}

			// Restore stdout and read the captured output
			w.Close()
			os.Stdout = oldStdout
			var output []byte
			output, _ = io.ReadAll(r)

			// Check the output
			capturedOutput := string(output)
			if tc.name != "ListIngressesError" {
				assert.Contains(t, capturedOutput, "Ingress hostnames")
			}

			for _, host := range tc.expectedHosts {
				assert.Contains(t, capturedOutput, host)
			}
		})
	}
}

// nolint: paralleltest, nolintlint
// TestGetNamespaceLabels tests the GetNamespaceLabels function
func TestGetNamespaceLabels(t *testing.T) {
	tests := []struct {
		name           string
		envVars        map[string]string
		expectedLabels map[string]string
		expectedErr    error
	}{
		{
			name: "All required ENVs provided",
			envVars: map[string]string{
				"CHAINLINK_TEAM":        "infra",
				"CHAINLINK_PRODUCT":     "kubernetes",
				"CHAINLINK_COMPONENT":   "gap",
				"CHAINLINK_COST_CENTER": "platform",
			},
			expectedLabels: map[string]string{
				"chain.link/team":        "infra",
				"chain.link/product":     "kubernetes",
				"chain.link/component":   "gap",
				"chain.link/cost-center": "platform",
			},
			expectedErr: nil,
		},
		{
			name: "Missing optional ENVs",
			envVars: map[string]string{
				"CHAINLINK_TEAM":    "infra",
				"CHAINLINK_PRODUCT": "kubernetes",
			},
			expectedLabels: map[string]string{
				"chain.link/team":        "infra",
				"chain.link/product":     "kubernetes",
				"chain.link/component":   "crib", // Default value
				"chain.link/cost-center": "crib", // Default value
			},
			expectedErr: nil,
		},
		{
			name: "Missing required ENVs",
			envVars: map[string]string{
				"CHAINLINK_COMPONENT":   "gap",
				"CHAINLINK_COST_CENTER": "platform",
			},
			expectedLabels: nil,
			expectedErr:    fmt.Errorf("one or more required environment variables are missing: CHAINLINK_TEAM, CHAINLINK_PRODUCT"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Reset Viper before each test
			viper.Reset()

			// Set environment variables
			for key, value := range tt.envVars {
				os.Setenv(key, value)
			}

			// Tell Viper to read from environment variables
			viper.AutomaticEnv()

			// Call the function
			labels, err := utils.GetNamespaceLabels()

			// Cleanup environment variables
			for key := range tt.envVars {
				os.Unsetenv(key)
			}

			// Check for errors
			if tt.expectedErr != nil {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedLabels, labels)
			}
		})
	}
}
