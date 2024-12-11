package model

import (
	"os"
	"testing"

	"github.com/smartcontractkit/crib/dependencies/ccip-v2-scripts/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSaveJSONOutputFile(t *testing.T) {
	t.Parallel()
	// Create a temporary file
	tmpDir := t.TempDir()

	env := config.DevspaceEnv{
		Namespace: "test-ns",
		TmpDir:    tmpDir,
	}
	state := NewEnvState(nil, env)

	savedFile := state.SaveJSONOutputFile("test_file.json", env)
	file, err := os.ReadFile(savedFile)
	require.NoError(t, err)

	assert.Greater(t, len(file), 0, "saved file is not empty")
}
