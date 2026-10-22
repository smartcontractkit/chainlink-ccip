package adapters_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/testsetup"
	"github.com/smartcontractkit/chainlink-ccip/deployment/fees"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	ccvadapters "github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/adapters"
	v2changesets "github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/changesets"
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

// TestFeesAdapter_GetOnchainOnRampDestChainConfig verifies the v2.0 adapter reads the OnRamp's
// destination chain config for a configured lane. The lane is seeded with the default token network
// fee, which is what the migration import subtracts to avoid double charging.
func TestFeesAdapter_GetOnchainOnRampDestChainConfig(t *testing.T) {
	localSelector := chain_selectors.ETHEREUM_MAINNET.Selector
	remoteSelector := chain_selectors.ETHEREUM_MAINNET_ARBITRUM_1.Selector

	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{localSelector, remoteSelector}),
	)
	require.NoError(t, err)

	ds := datastore.NewMemoryDataStore()
	e.OperationsBundle = testsetup.BundleWithFreshReporter(e.OperationsBundle)
	deployLaneContractsToDatastore(t, e, localSelector, ds)
	e.OperationsBundle = testsetup.BundleWithFreshReporter(e.OperationsBundle)
	deployLaneContractsToDatastore(t, e, remoteSelector, ds)
	e.DataStore = ds.Seal()

	deployer := e.BlockChains.EVMChains()[localSelector].DeployerKey.From.Hex()
	e.OperationsBundle = testsetup.BundleWithFreshReporter(e.OperationsBundle)
	_, err = v2changesets.ConfigureChainsForLanesFromTopology(
		ccvadapters.GetCommitteeVerifierContractRegistry(),
		ccvadapters.GetChainFamilyRegistry(),
		changesets.GetRegistry(),
	).Apply(*e, v2changesets.ConfigureChainsForLanesFromTopologyConfig{
		Topology: bidirectionalLaneTopology(deployer, localSelector, remoteSelector),
		BuildLanesCrossFamilyConfig: v2changesets.BuildLanesCrossFamilyConfig{
			Lanes: []v2changesets.CrossFamilyLanePair{laneWithSeededGasPrice(localSelector, remoteSelector)},
			MCMS:  mcms.Input{},
		},
	})
	require.NoError(t, err)

	// Lane configuration caches Router reads from before the OnRamp was wired; refresh the bundle so
	// the Router's OnRamp mapping is observed when resolving the fee adapter.
	e.OperationsBundle = testsetup.BundleWithFreshReporter(e.OperationsBundle)

	feeAdapter, onRampRef, _, err := fees.ResolveFeeAdapterAndOnRamp(e.OperationsBundle, e.BlockChains, e.DataStore, localSelector, remoteSelector)
	require.NoError(t, err)

	reader, ok := feeAdapter.(fees.OnRampDestChainConfigReader)
	require.True(t, ok, "v2.0 fee adapter must implement OnRampDestChainConfigReader")

	cfg, err := reader.GetOnchainOnRampDestChainConfig(e.OperationsBundle, e.BlockChains, onRampRef, localSelector, remoteSelector)
	require.NoError(t, err)
	require.NotNil(t, cfg.Ref.Version)
	require.True(t, cfg.Ref.Version.Equal(utils.Version_2_0_0), "expected v2.0 OnRamp, got %s", cfg.Ref.Version)
	require.Equal(t,
		sequences.GetDefaultTokenFeeUSDCentsFor1_6Lane(localSelector, remoteSelector),
		cfg.TokenNetworkFeeUSDCents,
	)
}
