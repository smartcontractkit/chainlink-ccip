package adapters_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	chain_selectors "github.com/smartcontractkit/chain-selectors"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/deployment/fees"
)

// TestFeesAdapter_GetDefaultTokenTransferFeeConfig_NoUnrequestedFee guards against the CCIP 2.0 double
// fee on the FeeQuoter path: the OnRamp already charges the lane's token network fee, so the default
// FeeQuoter token transfer fee must be zero. The worst case is an Ethereum destination, where the old
// 1.x default was 150c.
func TestFeesAdapter_GetDefaultTokenTransferFeeConfig_NoUnrequestedFee(t *testing.T) {
	adapter := adapters.NewFeesAdapter(nil)

	cfg := adapter.GetDefaultTokenTransferFeeConfig(12345, chain_selectors.ETHEREUM_MAINNET.Selector)
	require.Equal(t, uint32(0), cfg.MinFeeUSDCents)
	require.True(t, cfg.IsEnabled)
	require.NotZero(t, cfg.DestGasOverhead)

	// Sanity check: the shared 1.x default is still non-zero so the pre-2.0 adapters keep charging it.
	legacy := fees.GetDefaultChainAgnosticTokenTransferFeeConfig(12345, chain_selectors.ETHEREUM_MAINNET.Selector)
	require.NotEqual(t, uint32(0), legacy.MinFeeUSDCents)
}
