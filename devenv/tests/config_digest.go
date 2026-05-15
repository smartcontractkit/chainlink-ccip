package tests

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	"github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/stretchr/testify/require"
)

func RunUpdateConfigDigestTests(t *testing.T, e *deployment.Environment, selectors []uint64) {
	selectorsToImpl := buildImplsMap(t, e, selectors)

	// get two distinct selectors
	fromImpl, toImpl := selectorsToImpl[selectors[0]], selectorsToImpl[selectors[1]]
	require.NotEqual(t, fromImpl, toImpl)

	destFamily, err := chainsel.GetSelectorFamily(toImpl.ChainSelector())
	require.NoError(t, err)

	deployAdapter, ok := deploy.GetRegistry().GetDeployer(destFamily, semver.MustParse("1.6.0"))
	require.True(t, ok)
	require.NotNil(t, deployAdapter)

	_, err = operations.ExecuteSequence(e.OperationsBundle, deployAdapter.SetOCR3Config(), e.BlockChains, deploy.SetOCR3ConfigInput{
		ChainSelector: toImpl.ChainSelector(),
		Datastore:     e.DataStore,
		Configs: map[ccipocr3.PluginType]deploy.OCR3ConfigArgs{
			ccipocr3.PluginTypeCCIPCommit: {
				ConfigDigest:                   [32]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32},
				PluginType:                     ccipocr3.PluginTypeCCIPCommit,
				F:                              1,
				IsSignatureVerificationEnabled: true,
				Signers:                        [][]byte{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10, 11, 12}, {13, 14, 15}},
				Transmitters:                   [][]byte{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10, 11, 12}, {13, 14, 15}},
			},
		},
	})
	require.NoError(t, err)
}
