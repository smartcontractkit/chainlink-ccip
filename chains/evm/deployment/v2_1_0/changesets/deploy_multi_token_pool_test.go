package changesets_test

import (
	"fmt"
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
	ownership_sequences "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
	onramp_v1_6_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/onramp"
	onramp_v2_0_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/onramp"
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
	require.Empty(t, secondRefs, "reused refs must not be re-emitted in the changeset output")
	for _, firstRef := range firstRefs {
		existing, err := e.DataStore.Addresses().Get(firstRef.Key())
		require.NoError(t, err)
		require.Equal(t, firstRef.Address, existing.Address)
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

func TestDeployMultiTokenPoolRejectsMixedReuseBeforeDeploying(t *testing.T) {
	tests := []struct {
		name        string
		missing     cldf.ContractType
		expectedErr string
	}{
		{name: "reused PoolConfigManager with missing store", missing: fee_policy_store.ContractType, expectedErr: "FeePolicyStore does not"},
		{name: "reused pool with missing PoolConfigManager", missing: pool_config_manager.ContractType, expectedErr: "PoolConfigManager does not"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			e, err := environment.New(t.Context(), environment.WithEVMSimulated(t, []uint64{multiTokenPoolEVMTestChainSel}))
			require.NoError(t, err)
			prerequisites := deployMultiTokenPoolPrerequisites(t, e)
			chain := e.BlockChains.EVMChains()[multiTokenPoolEVMTestChainSel]
			admin := chain.DeployerKey.From.Hex()

			firstOut := applyMultiTokenPoolDeployment(t, e, admin)
			firstRefs, err := firstOut.DataStore.Addresses().Fetch()
			require.NoError(t, err)

			ds := datastore.NewMemoryDataStore()
			for _, ref := range append(prerequisites, firstRefs...) {
				if ref.Type == datastore.ContractType(test.missing) {
					continue
				}
				require.NoError(t, ds.Addresses().Add(ref))
			}
			e.DataStore = ds.Seal()
			// A fresh bundle stops the first run's cached sequence report from being replayed.
			e.OperationsBundle = cldf_ops.NewBundle(e.OperationsBundle.GetContext, e.Logger, cldf_ops.NewMemoryReporter())

			nonceBefore, err := chain.Client.PendingNonceAt(t.Context(), chain.DeployerKey.From)
			require.NoError(t, err)
			changeset := tokens.DeployMultiTokenPools(tokens.GetMultiTokenPoolAdapterRegistry(), changeset_utils.NewMCMSReaderRegistry())
			_, err = changeset.Apply(*e, multiTokenPoolTestConfig(admin))
			require.ErrorContains(t, err, test.expectedErr)
			nonceAfter, err := chain.Client.PendingNonceAt(t.Context(), chain.DeployerKey.From)
			require.NoError(t, err)
			require.Equal(t, nonceBefore, nonceAfter, "no transaction may be sent when the existing stack is inconsistent")
		})
	}
}

func TestDeployMultiTokenPoolVerifyRejectsInvalidAdminBeforeDeploying(t *testing.T) {
	e, err := environment.New(t.Context(), environment.WithEVMSimulated(t, []uint64{multiTokenPoolEVMTestChainSel}))
	require.NoError(t, err)
	deployMultiTokenPoolPrerequisites(t, e)

	changeset := tokens.DeployMultiTokenPools(tokens.GetMultiTokenPoolAdapterRegistry(), changeset_utils.NewMCMSReaderRegistry())
	err = changeset.VerifyPreconditions(*e, multiTokenPoolTestConfig("not-an-address"))
	require.ErrorContains(t, err, "invalid PoolConfigManager admin address")
}

func TestDeployMultiTokenPoolVerifyRejectsZeroAdminBeforeDeploying(t *testing.T) {
	e, err := environment.New(t.Context(), environment.WithEVMSimulated(t, []uint64{multiTokenPoolEVMTestChainSel}))
	require.NoError(t, err)
	deployMultiTokenPoolPrerequisites(t, e)

	changeset := tokens.DeployMultiTokenPools(tokens.GetMultiTokenPoolAdapterRegistry(), changeset_utils.NewMCMSReaderRegistry())
	err = changeset.VerifyPreconditions(*e, multiTokenPoolTestConfig(common.Address{}.Hex()))
	require.ErrorContains(t, err, "must be non-zero")
}

func deployTestOnRampV160(t *testing.T, e *cldf.Environment, tar common.Address) datastore.AddressRef {
	t.Helper()
	ref, err := contract.MaybeDeployContract(e.OperationsBundle, onramp_v1_6_0.Deploy, e.BlockChains.EVMChains()[multiTokenPoolEVMTestChainSel], contract.DeployInput[onramp_v1_6_0.ConstructorArgs]{
		ChainSelector:  multiTokenPoolEVMTestChainSel,
		TypeAndVersion: cldf.NewTypeAndVersion(onramp_v1_6_0.ContractType, *onramp_v1_6_0.Version),
		Args: onramp_v1_6_0.ConstructorArgs{
			StaticConfig: onramp_v1_6_0.StaticConfig{
				ChainSelector:      multiTokenPoolEVMTestChainSel,
				RmnRemote:          common.HexToAddress("0x1"),
				NonceManager:       common.HexToAddress("0x2"),
				TokenAdminRegistry: tar,
			},
			DynamicConfig: onramp_v1_6_0.DynamicConfig{FeeQuoter: common.HexToAddress("0x3"), FeeAggregator: common.HexToAddress("0x4")},
		},
	}, nil)
	require.NoError(t, err)
	return ref
}

func deployTestOnRampV200(t *testing.T, e *cldf.Environment, tar common.Address) datastore.AddressRef {
	t.Helper()
	ref, err := contract.MaybeDeployContract(e.OperationsBundle, onramp_v2_0_0.Deploy, e.BlockChains.EVMChains()[multiTokenPoolEVMTestChainSel], contract.DeployInput[onramp_v2_0_0.ConstructorArgs]{
		ChainSelector:  multiTokenPoolEVMTestChainSel,
		TypeAndVersion: cldf.NewTypeAndVersion(onramp_v2_0_0.ContractType, *onramp_v2_0_0.Version),
		Args: onramp_v2_0_0.ConstructorArgs{
			StaticConfig: onramp_v2_0_0.StaticConfig{
				ChainSelector:         multiTokenPoolEVMTestChainSel,
				RmnRemote:             common.HexToAddress("0x1"),
				MaxUSDCentsPerMessage: 1,
				TokenAdminRegistry:    tar,
			},
			DynamicConfig: onramp_v2_0_0.DynamicConfig{FeeQuoter: common.HexToAddress("0x3"), FeeAggregator: common.HexToAddress("0x4")},
		},
	}, nil)
	require.NoError(t, err)
	return ref
}

// verifyWithRamps verifies the test config on a chain whose datastore holds the prerequisites
// and the ramps returned by deployRamps.
func verifyWithRamps(t *testing.T, deployRamps func(e *cldf.Environment, tar common.Address) []datastore.AddressRef) error {
	t.Helper()
	e, err := environment.New(t.Context(), environment.WithEVMSimulated(t, []uint64{multiTokenPoolEVMTestChainSel}))
	require.NoError(t, err)
	prerequisites := deployMultiTokenPoolPrerequisites(t, e)
	chain := e.BlockChains.EVMChains()[multiTokenPoolEVMTestChainSel]
	tar := common.HexToAddress(findMultiTokenPoolRef(t, prerequisites, token_admin_registry.ContractType).Address)

	ds := datastore.NewMemoryDataStore()
	for _, ref := range append(prerequisites, deployRamps(e, tar)...) {
		require.NoError(t, ds.Addresses().Add(ref))
	}
	e.DataStore = ds.Seal()

	changeset := tokens.DeployMultiTokenPools(tokens.GetMultiTokenPoolAdapterRegistry(), changeset_utils.NewMCMSReaderRegistry())
	return changeset.VerifyPreconditions(*e, multiTokenPoolTestConfig(chain.DeployerKey.From.Hex()))
}

func TestDeployMultiTokenPoolVerifyChecksTokenAdminRegistryAgainstRamps(t *testing.T) {
	wrongTAR := common.HexToAddress("0xdead")
	deployers := map[string]func(t *testing.T, e *cldf.Environment, tar common.Address) datastore.AddressRef{
		"OnRamp 1.6.0": deployTestOnRampV160,
		"OnRamp 2.0.0": deployTestOnRampV200,
	}
	for rampName, deployRamp := range deployers {
		t.Run(rampName+" mismatched", func(t *testing.T) {
			err := verifyWithRamps(t, func(e *cldf.Environment, _ common.Address) []datastore.AddressRef {
				return []datastore.AddressRef{deployRamp(t, e, wrongTAR)}
			})
			require.ErrorContains(t, err, "must use the ramps' TokenAdminRegistry")
			require.ErrorContains(t, err, rampName)
		})
		t.Run(rampName+" matching", func(t *testing.T) {
			err := verifyWithRamps(t, func(e *cldf.Environment, tar common.Address) []datastore.AddressRef {
				return []datastore.AddressRef{deployRamp(t, e, tar)}
			})
			require.NoError(t, err)
		})
	}
}

func TestDeployMultiTokenPoolVerifyOnlyChecksNewestRampVersion(t *testing.T) {
	wrongTAR := common.HexToAddress("0xdead")
	t.Run("retired older ramp with a different TAR is ignored", func(t *testing.T) {
		err := verifyWithRamps(t, func(e *cldf.Environment, tar common.Address) []datastore.AddressRef {
			return []datastore.AddressRef{deployTestOnRampV160(t, e, wrongTAR), deployTestOnRampV200(t, e, tar)}
		})
		require.NoError(t, err)
	})
	t.Run("newest ramp with a different TAR is rejected", func(t *testing.T) {
		err := verifyWithRamps(t, func(e *cldf.Environment, tar common.Address) []datastore.AddressRef {
			return []datastore.AddressRef{deployTestOnRampV160(t, e, tar), deployTestOnRampV200(t, e, wrongTAR)}
		})
		require.ErrorContains(t, err, "OnRamp 2.0.0")
	})
}

func TestDeployMultiTokenPoolVerifyFailsBeforeAnyChainDeploysWhenALaterChainIsInvalid(t *testing.T) {
	const invalidChainSel = uint64(15971525489660198786) // has no prerequisites in the datastore
	e, err := environment.New(t.Context(), environment.WithEVMSimulated(t, []uint64{multiTokenPoolEVMTestChainSel, invalidChainSel}))
	require.NoError(t, err)
	deployMultiTokenPoolPrerequisites(t, e) // prerequisites for the valid chain only
	chain := e.BlockChains.EVMChains()[multiTokenPoolEVMTestChainSel]

	cfg := multiTokenPoolTestConfig(chain.DeployerKey.From.Hex())
	cfg.Chains[invalidChainSel] = cfg.Chains[multiTokenPoolEVMTestChainSel]
	nonceBefore, err := chain.Client.PendingNonceAt(t.Context(), chain.DeployerKey.From)
	require.NoError(t, err)

	changeset := tokens.DeployMultiTokenPools(tokens.GetMultiTokenPoolAdapterRegistry(), changeset_utils.NewMCMSReaderRegistry())
	err = changeset.VerifyPreconditions(*e, cfg)
	require.ErrorContains(t, err, "failed to find token admin registry address")
	require.ErrorContains(t, err, fmt.Sprint(invalidChainSel))

	nonceAfter, err := chain.Client.PendingNonceAt(t.Context(), chain.DeployerKey.From)
	require.NoError(t, err)
	require.Equal(t, nonceBefore, nonceAfter, "verification must not deploy anything on the valid chain")
}

// The pool and stores stay owned by the deployer after deployment. The generic ownership
// sequence transfers them.
func TestDeployedMultiTokenPoolContractsSupportGenericOwnershipTransfer(t *testing.T) {
	e, err := environment.New(t.Context(), environment.WithEVMSimulatedWithConfig(t, []uint64{multiTokenPoolEVMTestChainSel}, onchain.EVMSimLoaderConfig{NumAdditionalAccounts: 1}))
	require.NoError(t, err)
	deployMultiTokenPoolPrerequisites(t, e)
	chain := e.BlockChains.EVMChains()[multiTokenPoolEVMTestChainSel]
	out := applyMultiTokenPoolDeployment(t, e, chain.DeployerKey.From.Hex())
	refs, err := out.DataStore.Addresses().Fetch()
	require.NoError(t, err)

	// A non-deployer address stands in for the timelock. The transfer executes and the accept
	// step is returned as a proposal.
	timelock := chain.Users[0].From
	ownable := []datastore.AddressRef{
		findMultiTokenPoolRef(t, refs, connection_policy_store.ContractType),
		findMultiTokenPoolRef(t, refs, fee_policy_store.ContractType),
		findMultiTokenPoolRef(t, refs, security_policy_store.ContractType),
		findMultiTokenPoolRef(t, refs, burn_mint_multi_token_pool.ContractType),
	}
	batchOps, err := ownership_sequences.TransferAndAcceptOwnership(e.OperationsBundle, chain, ownership_sequences.OwnershipInputsFromRefs(multiTokenPoolEVMTestChainSel, timelock, ownable))
	require.NoError(t, err)
	require.Len(t, batchOps, len(ownable))

	acceptOwnershipSelector := []byte{0x79, 0xba, 0x50, 0x97} // acceptOwnership()
	for i, ref := range ownable {
		require.Len(t, batchOps[i].Transactions, 1, "%s must yield exactly the accept proposal", ref.Type)
		tx := batchOps[i].Transactions[0]
		require.Equal(t, common.HexToAddress(ref.Address), common.HexToAddress(tx.To), "%s proposal targets the contract", ref.Type)
		require.Equal(t, acceptOwnershipSelector, tx.Data, "%s proposal calls acceptOwnership()", ref.Type)

		// Ownership stays with the deployer until the accept executes.
		owner, _, err := ownership_sequences.LoadOwnableContract(common.HexToAddress(ref.Address), chain.Client)
		require.NoError(t, err)
		require.Equal(t, chain.DeployerKey.From, owner, "%s ownership is pending, not yet accepted", ref.Type)
	}
}

func TestDeployMultiTokenPoolVerifyRequiresRMNProxyAndRouter(t *testing.T) {
	tests := []struct {
		name        string
		missing     cldf.ContractType
		expectedErr string
	}{
		{name: "missing RMNProxy", missing: rmn_proxy.ContractType, expectedErr: "failed to find RMNProxy address"},
		{name: "missing Router", missing: router.ContractType, expectedErr: "failed to resolve router address"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			e, err := environment.New(t.Context(), environment.WithEVMSimulated(t, []uint64{multiTokenPoolEVMTestChainSel}))
			require.NoError(t, err)
			prerequisites := deployMultiTokenPoolPrerequisites(t, e)
			chain := e.BlockChains.EVMChains()[multiTokenPoolEVMTestChainSel]

			ds := datastore.NewMemoryDataStore()
			for _, ref := range prerequisites {
				if ref.Type == datastore.ContractType(test.missing) {
					continue
				}
				require.NoError(t, ds.Addresses().Add(ref))
			}
			e.DataStore = ds.Seal()

			changeset := tokens.DeployMultiTokenPools(tokens.GetMultiTokenPoolAdapterRegistry(), changeset_utils.NewMCMSReaderRegistry())
			err = changeset.VerifyPreconditions(*e, multiTokenPoolTestConfig(chain.DeployerKey.From.Hex()))
			require.ErrorContains(t, err, test.expectedErr)
		})
	}
}
