package changesets_test

import (
	"math/big"
	"testing"

	"github.com/aws/smithy-go/ptr"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/changesets"
	usdc_token_pool_proxy_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/operations/usdc_token_pool_proxy"
	usdc_token_pool_proxy_bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_4/usdc_token_pool_proxy"
	"github.com/smartcontractkit/chainlink-ccip/deployment/testhelpers"
	changesets_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	mcms_types "github.com/smartcontractkit/mcms/types"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/adapters"

	"github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	deploymentutils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
)

func TestDeployUSDCTokenPoolProxyChangeset_BothPoolsDeployed(t *testing.T) {
	chainSelector := uint64(chain_selectors.TEST_90000001.Selector)
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSelector}),
	)
	require.NoError(t, err, "Failed to create environment")
	require.NotNil(t, e, "Environment should be created")
	ds := datastore.NewMemoryDataStore()
	e.DataStore = ds.Seal()

	evmChain := e.BlockChains.EVMChains()[chainSelector]

	tokenAddress := common.Address{1}
	err = ds.Addresses().Add(datastore.AddressRef{
		Type:          datastore.ContractType("USDCToken"),
		Version:       semver.MustParse("1.0.0"),
		Address:       tokenAddress.Hex(),
		ChainSelector: chainSelector,
	})
	require.NoError(t, err, "Failed to add USDCToken address to datastore")

	routerAddress := common.Address{2}
	err = ds.Addresses().Add(datastore.AddressRef{
		Type:          datastore.ContractType("Router"),
		Version:       semver.MustParse("1.2.0"),
		Address:       routerAddress.Hex(),
		ChainSelector: chainSelector,
	})
	require.NoError(t, err, "Failed to add router address to datastore")

	// Simulate a deployed CCTP V1 pool
	cctpV1PoolAddress := common.Address{3}
	err = ds.Addresses().Add(datastore.AddressRef{
		Type:          datastore.ContractType("USDCTokenPool"),
		Version:       semver.MustParse("1.6.4"),
		Address:       cctpV1PoolAddress.Hex(),
		ChainSelector: chainSelector,
	})
	require.NoError(t, err, "Failed to add USDCTokenPool address to datastore")

	// Simulate a deployed CCTP V2 pool
	cctpV2PoolAddress := common.Address{4}
	err = ds.Addresses().Add(datastore.AddressRef{
		Type:          datastore.ContractType("USDCTokenPoolCCTPV2"),
		Version:       semver.MustParse("1.6.4"),
		Address:       cctpV2PoolAddress.Hex(),
		ChainSelector: chainSelector,
	})
	require.NoError(t, err, "Failed to add USDCTokenPoolCCTPV2 address to datastore")

	// Define pool addresses for the input
	legacyPoolAddress := common.Address{3}

	// Create the input to the changeset
	deployInput := changesets.DeployUSDCTokenPoolProxyInput{
		ChainInputs: []changesets.DeployUSDCTokenPoolProxyPerChainInput{
			{
				ChainSelector:     chainSelector,
				LegacyPoolAddress: legacyPoolAddress,
			},
		},
		MCMS: mcms.Input{
			OverridePreviousRoot: false,
			ValidUntil:           3759765795,
			TimelockDelay:        mcms_types.MustParseDuration("0s"),
			TimelockAction:       mcms_types.TimelockActionSchedule,
			Qualifier:            "test",
			Description:          "Deploy USDC Token Pool Proxy",
		},
	}

	// deploy mcms
	evmDeployer := &adapters.EVMDeployer{}
	dReg := deploy.GetRegistry()
	dReg.RegisterDeployer(chain_selectors.FamilyEVM, deploy.MCMSVersion, evmDeployer)
	cs := deploy.DeployMCMS(dReg)
	output, err := cs.Apply(*e, deploy.MCMSDeploymentConfig{
		AdapterVersion: semver.MustParse("1.0.0"),
		Chains: map[uint64]deploy.MCMSDeploymentConfigPerChain{
			chainSelector: {
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

	deployChangeset := changesets.DeployUSDCTokenPoolProxyChangeset(mcmsRegistry)
	deployChangesetOutput, err := deployChangeset.Apply(*e, deployInput)
	require.NoError(t, err, "DeployUSDCTokenPoolProxyChangeset should not error")
	require.Greater(t, len(deployChangesetOutput.Reports), 0)

	// Extract the deployed contract address from the output DataStore
	addressRef, err := deployChangesetOutput.DataStore.Addresses().Get(
		datastore.NewAddressRefKey(
			chainSelector,
			datastore.ContractType(usdc_token_pool_proxy_ops.ContractType),
			semver.MustParse("1.6.4"),
			"",
		),
	)
	require.NoError(t, err, "Failed to get deployed contract address from DataStore")
	require.NotEmpty(t, addressRef.Address, "Deployed contract address should not be empty")

	deployedAddress := common.HexToAddress(addressRef.Address)

	// Create a gobinding instance from the deployed address
	usdcTokenPoolProxy, err := usdc_token_pool_proxy_bindings.NewUSDCTokenPoolProxy(deployedAddress, evmChain.Client)
	require.NoError(t, err, "Failed to create USDCTokenPoolProxy gobinding instance")

	// Call getPools() and verify the pool addresses match the input
	pools, err := usdcTokenPoolProxy.GetPools(&bind.CallOpts{Context: t.Context()})
	require.NoError(t, err, "Failed to call getPools()")
	require.Equal(t, legacyPoolAddress, pools.LegacyCctpV1Pool, "LegacyCctpV1Pool should match input")
	require.Equal(t, cctpV1PoolAddress, pools.CctpV1Pool, "CctpV1Pool should match input")
	require.Equal(t, cctpV2PoolAddress, pools.CctpV2Pool, "CctpV2Pool should match input")

	// Call isSupportedToken() and verify it returns true for the input token
	isSupported, err := usdcTokenPoolProxy.IsSupportedToken(&bind.CallOpts{Context: t.Context()}, tokenAddress)
	require.NoError(t, err, "Failed to call isSupportedToken()")
	require.True(t, isSupported, "Token should be supported")
}

func TestDeployUSDCTokenPoolProxyChangeset_OnlyCCTPV2PoolDeployed(t *testing.T) {
	chainSelector := uint64(chain_selectors.TEST_90000001.Selector)
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSelector}),
	)
	require.NoError(t, err, "Failed to create environment")
	require.NotNil(t, e, "Environment should be created")
	ds := datastore.NewMemoryDataStore()
	e.DataStore = ds.Seal()

	evmChain := e.BlockChains.EVMChains()[chainSelector]

	tokenAddress := common.Address{1}
	err = ds.Addresses().Add(datastore.AddressRef{
		Type:          datastore.ContractType("USDCToken"),
		Version:       semver.MustParse("1.0.0"),
		Address:       tokenAddress.Hex(),
		ChainSelector: chainSelector,
	})
	require.NoError(t, err, "Failed to add USDCToken address to datastore")

	routerAddress := common.Address{2}
	err = ds.Addresses().Add(datastore.AddressRef{
		Type:          datastore.ContractType("Router"),
		Version:       semver.MustParse("1.2.0"),
		Address:       routerAddress.Hex(),
		ChainSelector: chainSelector,
	})
	require.NoError(t, err, "Failed to add router address to datastore")

	cctpV2PoolAddress := common.Address{4}
	err = ds.Addresses().Add(datastore.AddressRef{
		Type:          datastore.ContractType("USDCTokenPoolCCTPV2"),
		Version:       semver.MustParse("1.6.4"),
		Address:       cctpV2PoolAddress.Hex(),
		ChainSelector: chainSelector,
	})
	require.NoError(t, err, "Failed to add USDCTokenPoolCCTPV2 address to datastore")

	legacyPoolAddress := common.Address{3}

	// Create the input to the changeset
	deployInput := changesets.DeployUSDCTokenPoolProxyInput{
		ChainInputs: []changesets.DeployUSDCTokenPoolProxyPerChainInput{
			{
				ChainSelector:     chainSelector,
				LegacyPoolAddress: legacyPoolAddress,
			},
		},
		MCMS: mcms.Input{
			OverridePreviousRoot: false,
			ValidUntil:           3759765795,
			TimelockDelay:        mcms_types.MustParseDuration("0s"),
			TimelockAction:       mcms_types.TimelockActionSchedule,
			Qualifier:            "test",
			Description:          "Deploy USDC Token Pool Proxy",
		},
	}

	// deploy mcms
	evmDeployer := &adapters.EVMDeployer{}
	dReg := deploy.GetRegistry()
	dReg.RegisterDeployer(chain_selectors.FamilyEVM, deploy.MCMSVersion, evmDeployer)
	cs := deploy.DeployMCMS(dReg)
	output, err := cs.Apply(*e, deploy.MCMSDeploymentConfig{
		AdapterVersion: semver.MustParse("1.0.0"),
		Chains: map[uint64]deploy.MCMSDeploymentConfigPerChain{
			chainSelector: {
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

	deployChangeset := changesets.DeployUSDCTokenPoolProxyChangeset(mcmsRegistry)
	deployChangesetOutput, err := deployChangeset.Apply(*e, deployInput)
	require.NoError(t, err, "DeployUSDCTokenPoolProxyChangeset should not error")
	require.Greater(t, len(deployChangesetOutput.Reports), 0)

	// Extract the deployed contract address from the output DataStore
	addressRef, err := deployChangesetOutput.DataStore.Addresses().Get(
		datastore.NewAddressRefKey(
			chainSelector,
			datastore.ContractType(usdc_token_pool_proxy_ops.ContractType),
			semver.MustParse("1.6.4"),
			"",
		),
	)
	require.NoError(t, err, "Failed to get deployed contract address from DataStore")
	require.NotEmpty(t, addressRef.Address, "Deployed contract address should not be empty")

	deployedAddress := common.HexToAddress(addressRef.Address)

	// Create a gobinding instance from the deployed address
	usdcTokenPoolProxy, err := usdc_token_pool_proxy_bindings.NewUSDCTokenPoolProxy(deployedAddress, evmChain.Client)
	require.NoError(t, err, "Failed to create USDCTokenPoolProxy gobinding instance")

	// Call getPools() and verify the pool addresses match the input
	pools, err := usdcTokenPoolProxy.GetPools(&bind.CallOpts{Context: t.Context()})
	require.NoError(t, err, "Failed to call getPools()")
	require.Equal(t, legacyPoolAddress, pools.LegacyCctpV1Pool, "LegacyCctpV1Pool should be empty")
	require.Equal(t, common.Address{}, pools.CctpV1Pool, "CctpV1Pool should be empty")
	require.Equal(t, cctpV2PoolAddress, pools.CctpV2Pool, "CctpV2Pool should match input")

	// Call isSupportedToken() and verify it returns true for the input token
	isSupported, err := usdcTokenPoolProxy.IsSupportedToken(&bind.CallOpts{Context: t.Context()}, tokenAddress)
	require.NoError(t, err, "Failed to call isSupportedToken()")
	require.True(t, isSupported, "Token should be supported")
}

func TestDeployUSDCTokenPoolProxyChangeset_OnlyCCTPV1PoolDeployed(t *testing.T) {
	chainSelector := uint64(chain_selectors.TEST_90000001.Selector)
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSelector}),
	)
	require.NoError(t, err, "Failed to create environment")
	require.NotNil(t, e, "Environment should be created")
	ds := datastore.NewMemoryDataStore()
	e.DataStore = ds.Seal()

	evmChain := e.BlockChains.EVMChains()[chainSelector]

	tokenAddress := common.Address{1}
	err = ds.Addresses().Add(datastore.AddressRef{
		Type:          datastore.ContractType("USDCToken"),
		Version:       semver.MustParse("1.0.0"),
		Address:       tokenAddress.Hex(),
		ChainSelector: chainSelector,
	})
	require.NoError(t, err, "Failed to add USDCToken address to datastore")

	routerAddress := common.Address{2}
	err = ds.Addresses().Add(datastore.AddressRef{
		Type:          datastore.ContractType("Router"),
		Version:       semver.MustParse("1.2.0"),
		Address:       routerAddress.Hex(),
		ChainSelector: chainSelector,
	})
	require.NoError(t, err, "Failed to add router address to datastore")

	cctpV1PoolAddress := common.Address{3}
	err = ds.Addresses().Add(datastore.AddressRef{
		Type:          datastore.ContractType("USDCTokenPool"),
		Version:       semver.MustParse("1.6.4"),
		Address:       cctpV1PoolAddress.Hex(),
		ChainSelector: chainSelector,
	})
	require.NoError(t, err, "Failed to add USDCTokenPool address to datastore")

	legacyPoolAddress := common.Address{3}

	// Create the input to the changeset
	deployInput := changesets.DeployUSDCTokenPoolProxyInput{
		ChainInputs: []changesets.DeployUSDCTokenPoolProxyPerChainInput{
			{
				ChainSelector:     chainSelector,
				LegacyPoolAddress: legacyPoolAddress,
			},
		},
		MCMS: mcms.Input{
			OverridePreviousRoot: false,
			ValidUntil:           3759765795,
			TimelockDelay:        mcms_types.MustParseDuration("0s"),
			TimelockAction:       mcms_types.TimelockActionSchedule,
			Qualifier:            "test",
			Description:          "Deploy USDC Token Pool Proxy",
		},
	}

	// deploy mcms
	evmDeployer := &adapters.EVMDeployer{}
	dReg := deploy.GetRegistry()
	dReg.RegisterDeployer(chain_selectors.FamilyEVM, deploy.MCMSVersion, evmDeployer)
	cs := deploy.DeployMCMS(dReg)
	output, err := cs.Apply(*e, deploy.MCMSDeploymentConfig{
		AdapterVersion: semver.MustParse("1.0.0"),
		Chains: map[uint64]deploy.MCMSDeploymentConfigPerChain{
			chainSelector: {
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

	deployChangeset := changesets.DeployUSDCTokenPoolProxyChangeset(mcmsRegistry)
	deployChangesetOutput, err := deployChangeset.Apply(*e, deployInput)
	require.NoError(t, err, "DeployUSDCTokenPoolProxyChangeset should not error")
	require.Greater(t, len(deployChangesetOutput.Reports), 0)

	// Extract the deployed contract address from the output DataStore
	addressRef, err := deployChangesetOutput.DataStore.Addresses().Get(
		datastore.NewAddressRefKey(
			chainSelector,
			datastore.ContractType(usdc_token_pool_proxy_ops.ContractType),
			semver.MustParse("1.6.4"),
			"",
		),
	)
	require.NoError(t, err, "Failed to get deployed contract address from DataStore")
	require.NotEmpty(t, addressRef.Address, "Deployed contract address should not be empty")

	deployedAddress := common.HexToAddress(addressRef.Address)

	// Create a gobinding instance from the deployed address
	usdcTokenPoolProxy, err := usdc_token_pool_proxy_bindings.NewUSDCTokenPoolProxy(deployedAddress, evmChain.Client)
	require.NoError(t, err, "Failed to create USDCTokenPoolProxy gobinding instance")

	// Call getPools() and verify the pool addresses match the input
	pools, err := usdcTokenPoolProxy.GetPools(&bind.CallOpts{Context: t.Context()})
	require.NoError(t, err, "Failed to call getPools()")
	require.Equal(t, legacyPoolAddress, pools.LegacyCctpV1Pool, "LegacyCctpV1Pool should be empty")
	require.Equal(t, cctpV1PoolAddress, pools.CctpV1Pool, "CctpV1Pool should match input")
	require.Equal(t, common.Address{}, pools.CctpV2Pool, "CctpV2Pool should be empty")

	// Call isSupportedToken() and verify it returns true for the input token
	isSupported, err := usdcTokenPoolProxy.IsSupportedToken(&bind.CallOpts{Context: t.Context()}, tokenAddress)
	require.NoError(t, err, "Failed to call isSupportedToken()")
	require.True(t, isSupported, "Token should be supported")
}
