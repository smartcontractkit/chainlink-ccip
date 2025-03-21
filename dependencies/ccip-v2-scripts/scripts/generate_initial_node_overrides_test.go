package scripts

import (
	"testing"

	"github.com/smartcontractkit/chainlink/deployment/environment/crib"
	"github.com/smartcontractkit/chainlink/system-tests/lib/crypto"
	"github.com/smartcontractkit/crib/dependencies/ccip-v2-scripts/config"
	"github.com/smartcontractkit/crib/dependencies/ccip-v2-scripts/logging"
	"github.com/smartcontractkit/crib/dependencies/ccip-v2-scripts/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateConfigOverrides_shouldCreateNodeDetails(t *testing.T) {
	t.Parallel()

	// given
	tmpDir := t.TempDir()
	env := config.DevspaceEnv{
		Namespace:         "crib-local",
		Provider:          "kind",
		DonBootNodeCount:  1,
		DonNodeCount:      4,
		IngressBaseDomain: "main.stage.cldev.sh",
		TmpDir:            tmpDir,
		GethChainsCount:   2,
	}
	logger := logging.NewConsoleLogger()

	keys, err := crypto.GenerateP2PKeys("", env.DonBootNodeCount+env.DonNodeCount)
	require.NoError(t, err)

	envState := model.NewEnvState(logger, env)

	// when
	GenerateConfigOverrides(env, envState, keys)

	// then
	reader := crib.NewOutputReader(env.TmpDir)
	nodesDetails := reader.ReadNodesDetails()

	assert.Equal(t, []string{}, nodesDetails.NodeIDs, "Node IDs should remain unchanged")
	assert.Equal(t, "ccip-bt-0", nodesDetails.BootstrapNode.InternalHost)
	assert.Equal(t, "5001", nodesDetails.BootstrapNode.Port)
	assert.Len(t, nodesDetails.BootstrapNode.P2PID, 52)
}
