package cctp

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/stretchr/testify/require"

	hybrid_lock_release_usdc_token_pool_bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_2/hybrid_lock_release_usdc_token_pool"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

func TestMigrateHybridLockReleaseLiquidityErrorsWhenLockBoxMissing(t *testing.T) {
	chainSelector := chain_selectors.ETHEREUM_TESTNET_SEPOLIA.Selector
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSelector}),
	)
	require.NoError(t, err, "Failed to create environment")
	require.NotNil(t, e, "Environment should be created")
	e.DataStore = datastore.NewMemoryDataStore().Seal()

	setup := setupCCTPTestEnvironment(t, e, chainSelector)
	chain := e.BlockChains.EVMChains()[chainSelector]

	siloedReport, err := operations.ExecuteSequence(
		e.OperationsBundle,
		DeploySiloedUSDCLockRelease,
		e.BlockChains,
		DeploySiloedUSDCLockReleaseInput{
			ChainSelector:             chainSelector,
			USDCToken:                 setup.USDCToken.Hex(),
			Router:                    setup.Router.Hex(),
			RMN:                       setup.RMN.Hex(),
			SiloedUSDCTokenPool:       "",
			LockReleaseChainSelectors: nil,
		},
	)
	require.NoError(t, err, "ExecuteSequence should not error")

	hybridPoolAddr, tx, hybridPool, err := hybrid_lock_release_usdc_token_pool_bindings.DeployHybridLockReleaseUSDCTokenPool(
		chain.DeployerKey,
		chain.Client,
		setup.TokenMessenger,
		setup.MessageTransmitter,
		setup.USDCToken,
		[]common.Address{},
		setup.RMN,
		setup.Router,
		common.Address{},
	)
	require.NoErrorf(t, err, "Failed to confirm HybridLockReleaseUSDCTokenPool deployment: %v", err)
	require.NoError(t, err, "Failed to confirm HybridLockReleaseUSDCTokenPool deployment: %v", err)

	lockReleaseChainSelector := chain_selectors.ETHEREUM_MAINNET_ARBITRUM_1.Selector
	tx, err = hybridPool.UpdateChainSelectorMechanisms(chain.DeployerKey, nil, []uint64{lockReleaseChainSelector})
	require.NoError(t, err, "Failed to configure lock-release chain on hybrid pool")
	_, err = chain.Confirm(tx)
	require.NoError(t, err, "Failed to confirm lock-release chain update")

	_, err = operations.ExecuteSequence(
		e.OperationsBundle,
		MigrateHybridLockReleaseLiquidity,
		e.BlockChains,
		MigrateHybridLockReleaseLiquidityInput{
			ChainSelector:              chainSelector,
			HybridLockReleaseTokenPool: hybridPoolAddr.Hex(),
			SiloedUSDCTokenPool:        siloedReport.Output.SiloedPoolAddress,
			USDCToken:                  setup.USDCToken.Hex(),
			LockReleaseChainSelectors:  []uint64{lockReleaseChainSelector},
			LiquidityWithdrawPercent:   50,
		},
	)
	require.ErrorContains(t, err, "lockbox not configured for chain")
}
