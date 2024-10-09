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
	// Create a temporary directory for testing
	tempDir := t.TempDir()

	// Initialize a new git repository in the temporary directory
	gitRepoInTempDir := filepath.Join(tempDir, "git_repo")
	_, err := git.PlainInit(gitRepoInTempDir, false)
	require.NoError(t, err)

	// Create a subdirectory within the git repository
	subDir := filepath.Join(gitRepoInTempDir, "subdir", "subsubdir")
	err = os.MkdirAll(subDir, 0755)
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
			if tc.name == "DirectoryWithoutGitRepository" {
				require.NoError(t, os.Mkdir(tc.dir, 0750))
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
