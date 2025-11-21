package changesets_test

import (
	"math/big"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/aws/smithy-go/ptr"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/stretchr/testify/require"

	common "github.com/ethereum/go-ethereum/common"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/adapters"
	changesets "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/changesets"
	siloed_usdc_token_pool_bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_4/siloed_usdc_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	"github.com/smartcontractkit/chainlink-ccip/deployment/testhelpers"
	deploymentutils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	mcms "github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	datastore "github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	mcms_types "github.com/smartcontractkit/mcms/types"

	factory_burn_mint_erc20_bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/factory_burn_mint_erc20"

	changesets_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
)

func TestSiloedUSDCTokenPoolDeployChangeset(t *testing.T) {
	chainSelector := uint64(chain_selectors.TEST_90000001.Selector)
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSelector}),
	)
	require.NoError(t, err, "Failed to create environment")
	require.NotNil(t, e, "Environment should be created")
	ds := datastore.NewMemoryDataStore()
	e.DataStore = ds.Seal()

	evmChain := e.BlockChains.EVMChains()[chainSelector]

	tokenAddress, tx, _, err := factory_burn_mint_erc20_bindings.DeployFactoryBurnMintERC20(
		evmChain.DeployerKey,
		evmChain.Client,
		"TestToken",
		"TEST6",
		6,
		big.NewInt(0),
		big.NewInt(0),
		evmChain.DeployerKey.From,
	)
	require.NoError(t, err, "Failed to deploy FactoryBurnMintERC20 token")
	_, err = evmChain.Confirm(tx)
	require.NoError(t, err, "Failed to confirm token deployment transaction")

	allowlist := []common.Address{common.Address{2}}
	rmnProxyAddress := common.Address{3}
	routerAddress := common.Address{4}
	lockBoxAddress := common.Address{5}

	changesetInput := changesets.SiloedUSDCTokenPoolDeployInput{
		ChainInputs: []changesets.SiloedUSDCTokenPoolDeployInputPerChain{
			{
				ChainSelector: chainSelector,
				Token:         tokenAddress,
				Allowlist:     allowlist,
				RMNProxy:      rmnProxyAddress,
				Router:        routerAddress,
				LockBox:       lockBoxAddress,
			},
		},
		MCMS: mcms.Input{
			OverridePreviousRoot: false,
			ValidUntil:           3759765795,
			TimelockDelay:        mcms_types.MustParseDuration("0s"),
			TimelockAction:       mcms_types.TimelockActionSchedule,
			Qualifier:            "test",
			Description:          "Deploy SiloedUSDCTokenPool",
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
	deployChangesetOutput, err := changesets.SiloedUSDCTokenPoolDeployChangeset(mcmsRegistry).Apply(*e, changesetInput)
	require.NoError(t, err, "Failed to apply SiloedUSDCTokenPoolDeployChangeset")
	require.NotNil(t, deployChangesetOutput, "Changeset output should not be nil")
	require.Greater(t, len(deployChangesetOutput.Reports), 0)

	siloedUSDCTokenPoolAddress, err := deployChangesetOutput.DataStore.Addresses().Get(datastore.NewAddressRefKey(chainSelector, "SiloedUSDCTokenPool", semver.MustParse("1.6.4"), ""))
	require.NoError(t, err, "Failed to get SiloedUSDCTokenPool address")
	require.Equal(t, siloedUSDCTokenPoolAddress.Address, siloedUSDCTokenPoolAddress.Address, "Expected SiloedUSDCTokenPool address to be in changeset output")

	siloedUSDCTokenPool, err := siloed_usdc_token_pool_bindings.NewSiloedUSDCTokenPool(common.HexToAddress(siloedUSDCTokenPoolAddress.Address), evmChain.Client)
	require.NoError(t, err, "Failed to create SiloedUSDCTokenPool")

	// Check that isSupportedToken(tokenAddress) is true
	isSupported, err := siloedUSDCTokenPool.IsSupportedToken(&bind.CallOpts{Context: t.Context()}, tokenAddress)
	require.NoError(t, err, "Failed to call IsSupportedToken")
	require.True(t, isSupported, "Expected IsSupportedToken to return true for the deployed token address")

	// Check that getToken() returns the correct token address
	gotToken, err := siloedUSDCTokenPool.GetToken(&bind.CallOpts{Context: t.Context()})
	require.NoError(t, err, "Failed to call GetToken")
	require.Equal(t, tokenAddress, gotToken, "Expected GetToken to return the deployed token address")

	// Check that getRmnProxy() returns the correct RMN proxy address
	gotRMNProxy, err := siloedUSDCTokenPool.GetRmnProxy(&bind.CallOpts{Context: t.Context()})
	require.NoError(t, err, "Failed to call GetRmnProxy")
	require.Equal(t, rmnProxyAddress, gotRMNProxy, "Expected GetRmnProxy to return the deployed RMN proxy address")

	// Check that getRouter() returns the correct router address
	gotRouter, err := siloedUSDCTokenPool.GetRouter(&bind.CallOpts{Context: t.Context()})
	require.NoError(t, err, "Failed to call GetRouter")
	require.Equal(t, routerAddress, gotRouter, "Expected GetRouter to return the deployed router address")
}
