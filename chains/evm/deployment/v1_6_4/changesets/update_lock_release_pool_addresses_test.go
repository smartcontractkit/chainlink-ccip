package changesets_test

import (
	"math/big"
	"testing"

	"github.com/aws/smithy-go/ptr"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	"github.com/smartcontractkit/chainlink-ccip/deployment/testhelpers"
	deploymentutils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"

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

	datastore "github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/stretchr/testify/require"

	changesets "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/changesets"
	usdc_token_pool_proxy_bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_4/usdc_token_pool_proxy"
	mcms "github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	mcms_types "github.com/smartcontractkit/mcms/types"

	changesets_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"

	burn_mint_token_pool_bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/burn_mint_token_pool"
)

func TestUpdateLockReleasePoolAddressesChangeset(t *testing.T) {
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
		"TEST6",
		6,
		big.NewInt(0),             // maxSupply (0 = unlimited)
		big.NewInt(0),             // preMint
		evmChain.DeployerKey.From, // newOwner
	)

	require.NoError(t, err, "Failed to deploy FactoryBurnMintERC20 token")
	_, err = evmChain.Confirm(tx)
	require.NoError(t, err, "Failed to confirm token deployment transaction")

	routerAddress := common.Address{1}

	// Deploy ERC20LockBox using the deploy operation

	usdc_token_pool_proxy_ref, err := contract_utils.MaybeDeployContract(e.OperationsBundle, usdc_token_pool_proxy.Deploy, evmChain, contract.DeployInput[usdc_token_pool_proxy.ConstructorArgs]{
		TypeAndVersion: cldf_deployment.NewTypeAndVersion(usdc_token_pool_proxy.ContractType, *semver.MustParse("1.6.4")),
		ChainSelector:  chainSelector,
		Args: usdc_token_pool_proxy.ConstructorArgs{
			Token: tokenAddress,
			PoolAddresses: usdc_token_pool_proxy.PoolAddresses{
				LegacyCctpV1Pool: common.Address{},
				CctpV1Pool:       common.Address{},
				CctpV2Pool:       common.Address{},
			},
			Router: routerAddress,
		},
	}, nil)

	require.NoError(t, err, "Failed to deploy ERC20LockBox")

	// import binding for burn_mint_token_pool and deploy using tokenAddress
	burnMintTokenPoolAddress, tx, _, err := burn_mint_token_pool_bindings.DeployBurnMintTokenPool(
		evmChain.DeployerKey,
		evmChain.Client,
		tokenAddress,
		6,
		[]common.Address{},
		common.Address{1},
		common.Address{2},
	)
	require.NoError(t, err, "Failed to deploy BurnMintTokenPool contract")
	_, err = evmChain.Confirm(tx)
	require.NoError(t, err, "Failed to confirm BurnMintTokenPool contract deployment")

	updateLockReleasePoolAddressesInput := changesets.UpdateLockReleasePoolAddressesInput{
		ChainInputs: []changesets.UpdateLockReleasePoolAddressesPerChainInput{
			{
				ChainSelector: chainSelector,
				Address:       common.HexToAddress(usdc_token_pool_proxy_ref.Address),
				LockReleasePoolAddrs: usdc_token_pool_proxy.UpdateLockReleasePoolAddressesArgs{
					RemoteChainSelectors: []uint64{chainSelector},
					LockReleasePools:     []common.Address{burnMintTokenPoolAddress},
				},
			},
		},
		MCMS: mcms.Input{
			OverridePreviousRoot: false,
			ValidUntil:           3759765795,
			TimelockDelay:        mcms_types.MustParseDuration("0s"),
			TimelockAction:       mcms_types.TimelockActionSchedule,
			Qualifier:            "test",
			Description:          "Update lock release pool addresses",
		},
	}

	// deploy mcms
	evmDeployer := &adapters.EVMDeployer{}
	dReg := deploy.GetRegistry()
	dReg.RegisterDeployer(chain_selectors.FamilyEVM, deploy.MCMSVersion, evmDeployer)
	cs := deploy.DeployMCMS(dReg, nil)
	output, err := cs.Apply(*e, deploy.MCMSDeploymentConfig{
		AdapterVersion: semver.MustParse("1.0.0"),
		Chains: map[uint64]deploy.MCMSDeploymentConfigPerChain{
			uint64(chain_selectors.TEST_90000001.Selector): {
				Canceller:        testhelpers.SingleGroupMCMS(),
				Bypasser:         testhelpers.SingleGroupMCMS(),
				Proposer:         testhelpers.SingleGroupMCMS(),
				TimelockMinDelay: big.NewInt(0),
				Qualifier:        ptr.String("test"),
				TimelockAdmin:    evmChain.DeployerKey.From,
			},
		},
	})
	allAddrRefs, err := output.DataStore.Addresses().Fetch()
	require.NoError(t, err)
	timelockAddrs := make(map[uint64]string)
	for _, addrRef := range allAddrRefs {
		require.NoError(t, ds.Addresses().Add(addrRef))
		if addrRef.Type == datastore.ContractType(deploymentutils.RBACTimelock) {
			timelockAddrs[addrRef.ChainSelector] = addrRef.Address
		}
	}
	// update env datastore
	e.DataStore = ds.Seal()

	// Register the MCMS Reader
	mcmsRegistry := changesets_utils.GetRegistry()
	evmMCMSReader := &adapters.EVMMCMSReader{}
	mcmsRegistry.RegisterMCMSReader(chain_selectors.FamilyEVM, evmMCMSReader)

	updateLockReleasePoolAddressesChangeset := changesets.UpdateLockReleasePoolAddressesChangeset(mcmsRegistry)
	deployChangesetOutput, err := updateLockReleasePoolAddressesChangeset.Apply(*e, updateLockReleasePoolAddressesInput)
	require.NoError(t, err, "UpdateLockReleasePoolAddressesChangeset should not error")
	require.Greater(t, len(deployChangesetOutput.Reports), 0)

	// Verify the lock release pool address
	usdcTokenPoolProxy, err := usdc_token_pool_proxy_bindings.NewUSDCTokenPoolProxy(common.HexToAddress(usdc_token_pool_proxy_ref.Address), evmChain.Client)
	require.NoError(t, err, "Failed to create USDCTokenPoolProxy")
	lockReleasePoolAddr, err := usdcTokenPoolProxy.GetLockReleasePoolAddress(&bind.CallOpts{Context: t.Context()}, chainSelector)
	require.NoError(t, err, "Failed to get lock release pool address")
	require.Equal(t, lockReleasePoolAddr, burnMintTokenPoolAddress, "Lock release pool address should match")
}
