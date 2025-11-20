package changesets_test

import (
	"math/big"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	contract_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/operations/usdc_token_pool_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/factory_burn_mint_erc20"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"

	router_bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/router"

	datastore "github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/stretchr/testify/require"

	changesets "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/changesets"
	usdc_token_pool_proxy_bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_4/usdc_token_pool_proxy"
	mcms "github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	mcms_types "github.com/smartcontractkit/mcms/types"

	changesets_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
)

func TestUpdateLockOrBurnMechanismChangeset(t *testing.T) {
	chainSelector := uint64(chain_selectors.TEST_90000001.Selector)
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSelector}),
	)
	require.NoError(t, err, "Failed to create environment")
	require.NotNil(t, e, "Environment should be created")
	ds := datastore.NewMemoryDataStore()
	e.DataStore = ds.Seal()

	evmChain := e.BlockChains.EVMChains()[chainSelector]

	// Update the env datastore from the datastore
	e.DataStore = ds.Seal()

	// Deploy a real ERC20 token using factory_burn_mint_erc20
	tokenAddress, tx, _, err := factory_burn_mint_erc20.DeployFactoryBurnMintERC20(
		evmChain.DeployerKey,
		evmChain.Client,
		"TestToken",
		"TEST18",
		18,
		big.NewInt(0),             // maxSupply (0 = unlimited)
		big.NewInt(0),             // preMint
		evmChain.DeployerKey.From, // newOwner
	)
	require.NoError(t, err, "Failed to deploy FactoryBurnMintERC20 token")
	_, err = evmChain.Confirm(tx)
	require.NoError(t, err, "Failed to confirm token deployment transaction")

	routerAddress, tx, _, err := router_bindings.DeployRouter(evmChain.DeployerKey, evmChain.Client, common.Address{1}, common.Address{2})
	require.NoError(t, err, "Failed to deploy Router")
	_, err = evmChain.Confirm(tx)
	require.NoError(t, err, "Failed to confirm Router deployment transaction")

	// Deploy ERC20LockBox using the deploy operation

	usdc_token_pool_proxy_ref, err := contract_utils.MaybeDeployContract(e.OperationsBundle, usdc_token_pool_proxy.Deploy, evmChain, contract.DeployInput[usdc_token_pool_proxy.ConstructorArgs]{
		TypeAndVersion: cldf_deployment.NewTypeAndVersion(usdc_token_pool_proxy.ContractType, *semver.MustParse("1.6.4")),
		ChainSelector:  chainSelector,
		Args: usdc_token_pool_proxy.ConstructorArgs{
			Token: tokenAddress,
			PoolAddresses: usdc_token_pool_proxy.PoolAddresses{
				LegacyCctpV1Pool: common.Address{1},
				CctpV1Pool:       common.Address{2},
				CctpV2Pool:       common.Address{3},
			},
			Router: routerAddress,
		},
	}, nil)

	require.NoError(t, err, "Failed to deploy ERC20LockBox")

	updateLockOrBurnMechanismInput := changesets.UpdateLockOrBurnMechanismInput{
		ChainInputs: []changesets.UpdateLockOrBurnMechanismPerChainInput{
			{
				ChainSelector: chainSelector,
				Address:       common.HexToAddress(usdc_token_pool_proxy_ref.Address),
				Mechanisms: usdc_token_pool_proxy.UpdateLockOrBurnMechanismsArgs{
					RemoteChainSelectors: []uint64{chainSelector},
					Mechanisms:           []uint8{1},
				},
			},
		},
		MCMS: mcms.Input{
			OverridePreviousRoot: false,
			ValidUntil:           3759765795,
			TimelockDelay:        mcms_types.MustParseDuration("0s"),
			TimelockAction:       mcms_types.TimelockActionSchedule,
			Qualifier:            "test",
			Description:          "Update lock or burn mechanism",
		},
	}

	mcmsRegistry := changesets_utils.GetRegistry()
	updateLockOrBurnMechanismChangeset := changesets.UpdateLockOrBurnMechanismChangeset(mcmsRegistry)
	output, err := updateLockOrBurnMechanismChangeset.Apply(*e, updateLockOrBurnMechanismInput)
	require.NoError(t, err, "UpdateLockOrBurnMechanismChangeset should not error")
	require.Greater(t, len(output.Reports), 0)

	// Verify the lock or burn mechanism
	usdcTokenPoolProxy, err := usdc_token_pool_proxy_bindings.NewUSDCTokenPoolProxy(common.HexToAddress(usdc_token_pool_proxy_ref.Address), evmChain.Client)
	require.NoError(t, err, "Failed to create USDCTokenPoolProxy")
	mechanism, err := usdcTokenPoolProxy.GetLockOrBurnMechanism(&bind.CallOpts{Context: t.Context()}, chainSelector)
	require.NoError(t, err, "Failed to get lock or burn mechanism")
	require.Equal(t, mechanism, uint8(1), "Lock or burn mechanism should be 1")
}
