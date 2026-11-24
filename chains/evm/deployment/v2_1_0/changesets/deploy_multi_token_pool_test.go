package changesets_test

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/onchain"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/rmn_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/testsetup"
	_ "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_1_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_1_0/operations/burn_mint_multi_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_1_0/operations/connection_policy_store"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_1_0/operations/fee_policy_store"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_1_0/operations/pool_config_manager"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_1_0/operations/security_policy_store"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_1_0/sequences"
	burn_mint_multi_token_pool_bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/burn_mint_multi_token_pool"
	connection_policy_store_bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/connection_policy_store"
	fee_policy_store_bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/fee_policy_store"
	pool_config_manager_bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/pool_config_manager"
	security_policy_store_bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/security_policy_store"
	"github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	changeset_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
)

const multiTokenPoolEVMTestChainSel = uint64(5009297550715157269)

func deployMultiTokenPoolPrerequisites(t *testing.T, e *cldf.Environment) []datastore.AddressRef {
	t.Helper()
	chain := e.BlockChains.EVMChains()[multiTokenPoolEVMTestChainSel]

	rmnProxyRef, err := contract.MaybeDeployContract(e.OperationsBundle, rmn_proxy.Deploy, chain, contract.DeployInput[rmn_proxy.ConstructorArgs]{
		ChainSelector:  multiTokenPoolEVMTestChainSel,
		TypeAndVersion: cldf.NewTypeAndVersion(rmn_proxy.ContractType, *rmn_proxy.Version),
		Args:           rmn_proxy.ConstructorArgs{RMN: common.HexToAddress("0x1")},
	}, nil)
	require.NoError(t, err)
	routerRef, err := contract.MaybeDeployContract(e.OperationsBundle, router.Deploy, chain, contract.DeployInput[router.ConstructorArgs]{
		ChainSelector:  multiTokenPoolEVMTestChainSel,
		TypeAndVersion: cldf.NewTypeAndVersion(router.ContractType, *router.Version),
		Args: router.ConstructorArgs{
			WrappedNative: common.HexToAddress("0x2"),
			RMNProxy:      common.HexToAddress(rmnProxyRef.Address),
		},
	}, nil)
	require.NoError(t, err)
	tokenAdminRegistryRef, err := contract.MaybeDeployContract(e.OperationsBundle, token_admin_registry.Deploy, chain, contract.DeployInput[token_admin_registry.ConstructorArgs]{
		ChainSelector:  multiTokenPoolEVMTestChainSel,
		TypeAndVersion: cldf.NewTypeAndVersion(token_admin_registry.ContractType, *token_admin_registry.Version),
	}, nil)
	require.NoError(t, err)

	refs := []datastore.AddressRef{rmnProxyRef, routerRef, tokenAdminRegistryRef}
	ds := datastore.NewMemoryDataStore()
	for _, ref := range refs {
		require.NoError(t, ds.Addresses().Add(ref))
	}
	e.DataStore = ds.Seal()
	return refs
}

func multiTokenPoolTestConfig(admin string) tokens.DeployMultiTokenPoolsConfig {
	return tokens.DeployMultiTokenPoolsConfig{Chains: map[uint64]tokens.DeployMultiTokenPoolChainConfig{
		multiTokenPoolEVMTestChainSel: {
			PoolType:      string(burn_mint_multi_token_pool.ContractType),
			PoolVersion:   burn_mint_multi_token_pool.Version,
			PoolQualifier: "customer-a",
			Admin:         admin,
		},
	}}
}

func findMultiTokenPoolRef(t *testing.T, refs []datastore.AddressRef, contractType cldf.ContractType) datastore.AddressRef {
	t.Helper()
	for _, ref := range refs {
		if ref.Type == datastore.ContractType(contractType) {
			return ref
		}
	}
	require.Failf(t, "missing MTTP address ref", "contract type %s", contractType)
	return datastore.AddressRef{}
}

func applyMultiTokenPoolDeployment(t *testing.T, e *cldf.Environment, admin string) cldf.ChangesetOutput {
	t.Helper()
	changeset := tokens.DeployMultiTokenPools(tokens.GetMultiTokenPoolAdapterRegistry(), changeset_utils.NewMCMSReaderRegistry())
	cfg := multiTokenPoolTestConfig(admin)
	require.NoError(t, changeset.VerifyPreconditions(*e, cfg))
	out, err := changeset.Apply(*e, cfg)
	require.NoError(t, err)
	return out
}

func TestDeployMultiTokenPoolsApply(t *testing.T) {
	e, err := environment.New(t.Context(), environment.WithEVMSimulated(t, []uint64{multiTokenPoolEVMTestChainSel}))
	require.NoError(t, err)
	prerequisites := deployMultiTokenPoolPrerequisites(t, e)
	chain := e.BlockChains.EVMChains()[multiTokenPoolEVMTestChainSel]
	admin := common.HexToAddress("0x100")

	out := applyMultiTokenPoolDeployment(t, e, admin.Hex())
	refs, err := out.DataStore.Addresses().Fetch()
	require.NoError(t, err)
	require.Len(t, refs, 5)
	for _, ref := range refs {
		require.Equal(t, multiTokenPoolEVMTestChainSel, ref.ChainSelector)
		require.Equal(t, burn_mint_multi_token_pool.Version.String(), ref.Version.String())
		require.Equal(t, "customer-a", ref.Qualifier)
	}

	poolRef := findMultiTokenPoolRef(t, refs, burn_mint_multi_token_pool.ContractType)
	managerRef := findMultiTokenPoolRef(t, refs, pool_config_manager.ContractType)
	connectionStoreRef := findMultiTokenPoolRef(t, refs, connection_policy_store.ContractType)
	feeStoreRef := findMultiTokenPoolRef(t, refs, fee_policy_store.ContractType)
	securityStoreRef := findMultiTokenPoolRef(t, refs, security_policy_store.ContractType)
	rmnProxyRef := findMultiTokenPoolRef(t, prerequisites, rmn_proxy.ContractType)
	routerRef := findMultiTokenPoolRef(t, prerequisites, router.ContractType)
	tokenAdminRegistryRef := findMultiTokenPoolRef(t, prerequisites, token_admin_registry.ContractType)

	pool, err := burn_mint_multi_token_pool_bindings.NewBurnMintMultiTokenPool(common.HexToAddress(poolRef.Address), chain.Client)
	require.NoError(t, err)
	callOpts := &bind.CallOpts{Context: t.Context()}
	poolTypeAndVersion, err := pool.TypeAndVersion(callOpts)
	require.NoError(t, err)
	require.Equal(t, "BurnMintMultiTokenPool 2.1.0-dev", poolTypeAndVersion, "datastore refs intentionally use the non-prerelease 2.1.0 alias")
	gotManager, err := pool.GetPoolConfigManager(callOpts)
	require.NoError(t, err)
	require.Equal(t, managerRef.Address, gotManager.Hex())
	gotRouter, err := pool.GetRouter(callOpts)
	require.NoError(t, err)
	require.Equal(t, routerRef.Address, gotRouter.Hex())
	gotRMNProxy, err := pool.GetRmnProxy(callOpts)
	require.NoError(t, err)
	require.Equal(t, rmnProxyRef.Address, gotRMNProxy.Hex())

	manager, err := pool_config_manager_bindings.NewPoolConfigManager(common.HexToAddress(managerRef.Address), chain.Client)
	require.NoError(t, err)
	managerTypeAndVersion, err := manager.TypeAndVersion(callOpts)
	require.NoError(t, err)
	require.Equal(t, "PoolConfigManager 2.1.0-dev", managerTypeAndVersion)
	staticConfig, err := manager.GetStaticConfig(callOpts)
	require.NoError(t, err)
	require.Equal(t, tokenAdminRegistryRef.Address, staticConfig.TokenAdminRegistry.Hex())
	require.Equal(t, connectionStoreRef.Address, staticConfig.ConnectionStore.Hex())
	require.Equal(t, feeStoreRef.Address, staticConfig.FeeStore.Hex())
	require.Equal(t, securityStoreRef.Address, staticConfig.SecurityStore.Hex())
	defaultAdmin, err := manager.DefaultAdmin(callOpts)
	require.NoError(t, err)
	require.Equal(t, admin, defaultAdmin)

	connectionStore, err := connection_policy_store_bindings.NewConnectionPolicyStore(common.HexToAddress(connectionStoreRef.Address), chain.Client)
	require.NoError(t, err)
	connectionCallers, err := connectionStore.GetAllAuthorizedCallers(callOpts)
	require.NoError(t, err)
	require.Equal(t, []common.Address{common.HexToAddress(managerRef.Address)}, connectionCallers)
	feeStore, err := fee_policy_store_bindings.NewFeePolicyStore(common.HexToAddress(feeStoreRef.Address), chain.Client)
	require.NoError(t, err)
	feeCallers, err := feeStore.GetAllAuthorizedCallers(callOpts)
	require.NoError(t, err)
	require.Equal(t, []common.Address{common.HexToAddress(managerRef.Address)}, feeCallers)
	securityStore, err := security_policy_store_bindings.NewSecurityPolicyStore(common.HexToAddress(securityStoreRef.Address), chain.Client)
	require.NoError(t, err)
	securityCallers, err := securityStore.GetAllAuthorizedCallers(callOpts)
	require.NoError(t, err)
	require.Equal(t, []common.Address{common.HexToAddress(managerRef.Address)}, securityCallers)
}

func TestDeployMultiTokenPoolsApplyIsIdempotent(t *testing.T) {
	e, err := environment.New(t.Context(), environment.WithEVMSimulated(t, []uint64{multiTokenPoolEVMTestChainSel}))
	require.NoError(t, err)
	prerequisites := deployMultiTokenPoolPrerequisites(t, e)
	admin := e.BlockChains.EVMChains()[multiTokenPoolEVMTestChainSel].DeployerKey.From.Hex()

	firstOut := applyMultiTokenPoolDeployment(t, e, admin)
	firstRefs, err := firstOut.DataStore.Addresses().Fetch()
	require.NoError(t, err)

	ds := datastore.NewMemoryDataStore()
	for _, ref := range append(prerequisites, firstRefs...) {
		require.NoError(t, ds.Addresses().Add(ref))
	}
	e.DataStore = ds.Seal()
	e.OperationsBundle = testsetup.BundleWithFreshReporter(e.OperationsBundle)
	secondOut := applyMultiTokenPoolDeployment(t, e, admin)
	secondRefs, err := secondOut.DataStore.Addresses().Fetch()
	require.NoError(t, err)
	require.Len(t, secondRefs, len(firstRefs))
	for _, firstRef := range firstRefs {
		secondRef := findMultiTokenPoolRef(t, secondRefs, cldf.ContractType(firstRef.Type))
		require.Equal(t, firstRef.Address, secondRef.Address)
	}
	require.Empty(t, secondOut.MCMSTimelockProposals)
}

func TestDeployMultiTokenPoolPreparesStoreAuthorizationForNonDeployerOwner(t *testing.T) {
	e, err := environment.New(t.Context(), environment.WithEVMSimulatedWithConfig(t, []uint64{multiTokenPoolEVMTestChainSel}, onchain.EVMSimLoaderConfig{NumAdditionalAccounts: 1}))
	require.NoError(t, err)
	prerequisites := deployMultiTokenPoolPrerequisites(t, e)
	chain := e.BlockChains.EVMChains()[multiTokenPoolEVMTestChainSel]
	require.NotEmpty(t, chain.Users)

	firstOut := applyMultiTokenPoolDeployment(t, e, chain.DeployerKey.From.Hex())
	firstRefs, err := firstOut.DataStore.Addresses().Fetch()
	require.NoError(t, err)
	connectionStoreRef := findMultiTokenPoolRef(t, firstRefs, connection_policy_store.ContractType)
	managerRef := findMultiTokenPoolRef(t, firstRefs, pool_config_manager.ContractType)
	managerAddress := common.HexToAddress(managerRef.Address)
	connectionStoreAddress := common.HexToAddress(connectionStoreRef.Address)
	connectionStore, err := connection_policy_store_bindings.NewConnectionPolicyStore(connectionStoreAddress, chain.Client)
	require.NoError(t, err)

	// Remove the manager and hand the store to another owner. A retry must prepare the missing
	// authorization for that owner instead of sending it from the deployer.
	removeTx, err := connectionStore.ApplyAuthorizedCallerUpdates(chain.DeployerKey, connection_policy_store_bindings.AuthorizedCallersAuthorizedCallerArgs{
		RemovedCallers: []common.Address{managerAddress},
	})
	require.NoError(t, err)
	_, err = chain.Confirm(removeTx)
	require.NoError(t, err)
	newOwner := chain.Users[0]
	transferTx, err := connectionStore.TransferOwnership(chain.DeployerKey, newOwner.From)
	require.NoError(t, err)
	_, err = chain.Confirm(transferTx)
	require.NoError(t, err)
	acceptTx, err := connectionStore.AcceptOwnership(newOwner)
	require.NoError(t, err)
	_, err = chain.Confirm(acceptTx)
	require.NoError(t, err)
	owner, err := connectionStore.Owner(&bind.CallOpts{Context: t.Context()})
	require.NoError(t, err)
	require.Equal(t, newOwner.From, owner)
	authorizationRead, err := cldf_ops.ExecuteOperation(
		e.OperationsBundle,
		connection_policy_store.GetAllAuthorizedCallers,
		chain,
		contract.FunctionInput[struct{}]{ChainSelector: multiTokenPoolEVMTestChainSel, Address: connectionStoreAddress},
		cldf_ops.WithForceExecute[contract.FunctionInput[struct{}], evm.Chain](),
	)
	require.NoError(t, err)
	require.NotContains(t, authorizationRead.Output, managerAddress)

	ds := datastore.NewMemoryDataStore()
	for _, ref := range append(prerequisites, firstRefs...) {
		require.NoError(t, ds.Addresses().Add(ref))
	}
	e.DataStore = ds.Seal()
	e.OperationsBundle = testsetup.BundleWithFreshReporter(e.OperationsBundle)
	report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, sequences.DeployMultiTokenPool, e.BlockChains, tokens.DeployMultiTokenPoolInput{
		ChainSelector:     multiTokenPoolEVMTestChainSel,
		PoolType:          string(burn_mint_multi_token_pool.ContractType),
		PoolVersion:       burn_mint_multi_token_pool.Version,
		PoolQualifier:     "customer-a",
		Admin:             chain.DeployerKey.From.Hex(),
		ExistingDataStore: e.DataStore,
	})
	require.NoError(t, err)
	require.Len(t, report.Output.BatchOps, 1)
	require.Len(t, report.Output.BatchOps[0].Transactions, 1)
	preparedTx := report.Output.BatchOps[0].Transactions[0]
	require.Equal(t, connectionStoreAddress.Hex(), common.HexToAddress(preparedTx.To).Hex())

	storeABI, err := connection_policy_store_bindings.ConnectionPolicyStoreMetaData.GetAbi()
	require.NoError(t, err)
	expectedCalldata, err := storeABI.Pack("applyAuthorizedCallerUpdates", connection_policy_store_bindings.AuthorizedCallersAuthorizedCallerArgs{
		AddedCallers: []common.Address{managerAddress},
	})
	require.NoError(t, err)
	require.Equal(t, expectedCalldata, []byte(preparedTx.Data))
	authorizedCallers, err := connectionStore.GetAllAuthorizedCallers(&bind.CallOpts{Context: t.Context()})
	require.NoError(t, err)
	require.NotContains(t, authorizedCallers, managerAddress, "the prepared write must not execute directly")
}

func TestDeployMultiTokenPoolDefaultsAdminToDeployer(t *testing.T) {
	e, err := environment.New(t.Context(), environment.WithEVMSimulated(t, []uint64{multiTokenPoolEVMTestChainSel}))
	require.NoError(t, err)
	deployMultiTokenPoolPrerequisites(t, e)
	chain := e.BlockChains.EVMChains()[multiTokenPoolEVMTestChainSel]

	report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, sequences.DeployMultiTokenPool, e.BlockChains, tokens.DeployMultiTokenPoolInput{
		ChainSelector:     multiTokenPoolEVMTestChainSel,
		PoolType:          string(burn_mint_multi_token_pool.ContractType),
		PoolVersion:       burn_mint_multi_token_pool.Version,
		PoolQualifier:     "customer-a",
		ExistingDataStore: e.DataStore,
	})
	require.NoError(t, err)
	managerRef := findMultiTokenPoolRef(t, report.Output.Addresses, pool_config_manager.ContractType)
	manager, err := pool_config_manager_bindings.NewPoolConfigManager(common.HexToAddress(managerRef.Address), chain.Client)
	require.NoError(t, err)
	defaultAdmin, err := manager.DefaultAdmin(&bind.CallOpts{Context: t.Context()})
	require.NoError(t, err)
	require.Equal(t, chain.DeployerKey.From, defaultAdmin)
}

func TestDeployMultiTokenPoolRejectsUnsupportedInputsBeforeDeployment(t *testing.T) {
	e, err := environment.New(t.Context(), environment.WithEVMSimulated(t, []uint64{multiTokenPoolEVMTestChainSel}))
	require.NoError(t, err)

	tests := []struct {
		name        string
		input       tokens.DeployMultiTokenPoolInput
		expectedErr string
	}{
		{
			name: "missing version",
			input: tokens.DeployMultiTokenPoolInput{
				ChainSelector: multiTokenPoolEVMTestChainSel,
			},
			expectedErr: "pool version is required",
		},
		{
			name: "unsupported version",
			input: tokens.DeployMultiTokenPoolInput{
				ChainSelector: multiTokenPoolEVMTestChainSel,
				PoolVersion:   semver.MustParse("2.0.0"),
			},
			expectedErr: "unsupported multi-token pool version",
		},
		{
			name: "unsupported type",
			input: tokens.DeployMultiTokenPoolInput{
				ChainSelector: multiTokenPoolEVMTestChainSel,
				PoolVersion:   burn_mint_multi_token_pool.Version,
				PoolType:      "BurnMintTokenPool",
			},
			expectedErr: "unsupported multi-token pool type",
		},
		{
			name: "missing qualifier",
			input: tokens.DeployMultiTokenPoolInput{
				ChainSelector: multiTokenPoolEVMTestChainSel,
				PoolVersion:   burn_mint_multi_token_pool.Version,
				PoolType:      string(burn_mint_multi_token_pool.ContractType),
			},
			expectedErr: "pool qualifier is required",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := cldf_ops.ExecuteSequence(e.OperationsBundle, sequences.DeployMultiTokenPool, e.BlockChains, test.input)
			require.ErrorContains(t, err, test.expectedErr)
		})
	}
}
