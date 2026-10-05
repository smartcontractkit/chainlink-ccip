package changesets

import (
	"testing"

	"github.com/stretchr/testify/require"

	chainsel "github.com/smartcontractkit/chain-selectors"

	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/shared"
)

func TestValidateRemoteOnRamp(t *testing.T) {
	t.Parallel()
	evmSel := chainsel.ETHEREUM_MAINNET.Selector
	solSel := chainsel.SOLANA_MAINNET.Selector
	aptosSel := chainsel.APTOS_MAINNET.Selector
	suiSel := chainsel.SUI_MAINNET.Selector
	tonSel := chainsel.TON_MAINNET.Selector

	e := cldf.Environment{DataStore: sealedDataStore(t,
		testRef(evmSel, "0x1", string(shared.OnRamp), "1.6.0", ""),
		testRef(solSel, "router", string(shared.Router), "1.6.0", ""),
		testRef(aptosSel, "0x2", string(shared.AptosCCIPType), "1.6.0", ""),
		testRef(suiSel, "0x3", "SuiCCIP", "1.6.0", ""),
		testRef(tonSel, "EQ", string(shared.Router), "1.6.0", ""),
	)}
	for _, sel := range []uint64{evmSel, solSel, aptosSel, suiSel, tonSel} {
		require.NoError(t, validateRemoteOnRamp(e, sel), "chain %d", sel)
	}

	// the EVM onramp is only loaded at 1.6.0
	e = cldf.Environment{DataStore: sealedDataStore(t, testRef(evmSel, "0x1", string(shared.OnRamp), "1.5.0", ""))}
	require.ErrorContains(t, validateRemoteOnRamp(e, evmSel), "onramp contract does not exist on evm chain")

	// a superseded ref doesn't count
	e = cldf.Environment{DataStore: sealedDataStore(t, testRef(solSel, "router", string(shared.Router), "1.6.0", "", shared.SupersededLabel))}
	require.ErrorContains(t, validateRemoteOnRamp(e, solSel), "router contract does not exist on solana chain")

	// another chain's onramp doesn't count
	e = cldf.Environment{DataStore: sealedDataStore(t, testRef(evmSel, "0x1", string(shared.OnRamp), "1.6.0", ""))}
	require.Error(t, validateRemoteOnRamp(e, chainsel.ETHEREUM_MAINNET_ARBITRUM_1.Selector))

	require.Error(t, validateRemoteOnRamp(cldf.Environment{}, evmSel))
	require.Error(t, validateRemoteOnRamp(e, 1))
}
