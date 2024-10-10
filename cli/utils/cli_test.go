package utils

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/go-git/go-git/v5"
	"github.com/stretchr/testify/assert"
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

func TestPresentPrompt(t *testing.T) {
	t.Parallel()

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
			t.Parallel()

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
