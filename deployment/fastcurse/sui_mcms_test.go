package fastcurse

import (
	"testing"

	cselectors "github.com/smartcontractkit/chain-selectors"
	mcms_types "github.com/smartcontractkit/mcms/types"
	"github.com/stretchr/testify/require"

	mcms_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
)

func TestMcmsInputForUncurseProposal_SuiChainsUseSlowQualifier(t *testing.T) {
	t.Parallel()

	suiSelector := cselectors.SUI_TESTNET.Selector
	evmSelector := cselectors.ETHEREUM_MAINNET.Selector

	cfg := mcms_utils.Input{Qualifier: "RMNMCMS"}
	batchOps := []mcms_types.BatchOperation{
		{ChainSelector: mcms_types.ChainSelector(suiSelector)},
		{ChainSelector: mcms_types.ChainSelector(evmSelector)},
	}

	out := mcmsInputForUncurseProposal(cfg, batchOps)

	require.Equal(t, "RMNMCMS", out.Qualifier)
	require.Equal(t, "", out.ChainQualifiers[suiSelector])
	_, hasEVMOverride := out.ChainQualifiers[evmSelector]
	require.False(t, hasEVMOverride)
}

func TestMcmsInputForUncurseProposal_NoSuiChainsLeavesInputUnchanged(t *testing.T) {
	t.Parallel()

	evmSelector := cselectors.ETHEREUM_MAINNET.Selector
	cfg := mcms_utils.Input{Qualifier: "RMNMCMS"}
	batchOps := []mcms_types.BatchOperation{
		{ChainSelector: mcms_types.ChainSelector(evmSelector)},
	}

	out := mcmsInputForUncurseProposal(cfg, batchOps)

	require.Equal(t, cfg, out)
}
