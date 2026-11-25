package changesets_test

import (
	"fmt"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
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

func installMultiTokenPoolDataStore(t *testing.T, e *cldf.Environment, refs []datastore.AddressRef) {
	t.Helper()
	ds := datastore.NewMemoryDataStore()
	for _, ref := range refs {
		require.NoError(t, ds.Addresses().Add(ref))
	}
	e.DataStore = ds.Seal()
	e.OperationsBundle = testsetup.BundleWithFreshReporter(e.OperationsBundle)
}

func multiTokenPoolSequenceInput(e *cldf.Environment, admin string) tokens.DeployMultiTokenPoolInput {
	return tokens.DeployMultiTokenPoolInput{
		ChainSelector:     multiTokenPoolEVMTestChainSel,
		PoolType:          string(burn_mint_multi_token_pool.ContractType),
		PoolVersion:       burn_mint_multi_token_pool.Version,
		PoolQualifier:     "customer-a",
		Admin:             admin,
		ExistingDataStore: e.DataStore,
	}
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
	require.Equal(t, "BurnMintMultiTokenPool 2.1.0", poolTypeAndVersion)
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
	require.Equal(t, "PoolConfigManager 2.1.0", managerTypeAndVersion)
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

func TestDeployMultiTokenPoolsSameReporterReplaysDeploymentAndRechecksAuthorization(t *testing.T) {
	e, err := environment.New(t.Context(), environment.WithEVMSimulated(t, []uint64{multiTokenPoolEVMTestChainSel}))
	require.NoError(t, err)
	prerequisites := deployMultiTokenPoolPrerequisites(t, e)
	chain := e.BlockChains.EVMChains()[multiTokenPoolEVMTestChainSel]
	firstOut := applyMultiTokenPoolDeployment(t, e, chain.DeployerKey.From.Hex())
	firstRefs, err := firstOut.DataStore.Addresses().Fetch()
	require.NoError(t, err)
	managerAddress := common.HexToAddress(findMultiTokenPoolRef(t, firstRefs, pool_config_manager.ContractType).Address)
	connectionStore, err := connection_policy_store_bindings.NewConnectionPolicyStore(
		common.HexToAddress(findMultiTokenPoolRef(t, firstRefs, connection_policy_store.ContractType).Address), chain.Client)
	require.NoError(t, err)

	// Change authorization out of band, then rerun with the same reporter. Deployment must replay
	// its original report while the separately keyed authorization sequence restores the caller.
	removeTx, err := connectionStore.ApplyAuthorizedCallerUpdates(chain.DeployerKey, connection_policy_store_bindings.AuthorizedCallersAuthorizedCallerArgs{
		RemovedCallers: []common.Address{managerAddress},
	})
	require.NoError(t, err)
	_, err = chain.Confirm(removeTx)
	require.NoError(t, err)
	ds := datastore.NewMemoryDataStore()
	for _, ref := range append(prerequisites, firstRefs...) {
		require.NoError(t, ds.Addresses().Add(ref))
	}
	e.DataStore = ds.Seal()

	nonceBefore, err := chain.Client.PendingNonceAt(t.Context(), chain.DeployerKey.From)
	require.NoError(t, err)
	secondOut := applyMultiTokenPoolDeployment(t, e, chain.DeployerKey.From.Hex())
	nonceAfter, err := chain.Client.PendingNonceAt(t.Context(), chain.DeployerKey.From)
	require.NoError(t, err)
	require.Equal(t, nonceBefore+1, nonceAfter, "only the missing authorization should execute")
	authorizedCallers, err := connectionStore.GetAllAuthorizedCallers(&bind.CallOpts{Context: t.Context()})
	require.NoError(t, err)
	require.Contains(t, authorizedCallers, managerAddress)

	var firstDeploymentReportID, secondDeploymentReportID string
	var firstAuthorizationKey, secondAuthorizationKey string
	for _, report := range firstOut.Reports {
		switch report.Def.ID {
		case "deploy-multi-token-pool":
			firstDeploymentReportID = report.ID
		case "authorize-multi-token-pool":
			firstAuthorizationKey = report.IdempotencyKey
		}
	}
	for _, report := range secondOut.Reports {
		switch report.Def.ID {
		case "deploy-multi-token-pool":
			secondDeploymentReportID = report.ID
		case "authorize-multi-token-pool":
			secondAuthorizationKey = report.IdempotencyKey
		}
	}
	require.NotEmpty(t, firstDeploymentReportID)
	require.Equal(t, firstDeploymentReportID, secondDeploymentReportID, "deployment must be replayed from the same report")
	require.NotEmpty(t, firstAuthorizationKey)
	require.NotEmpty(t, secondAuthorizationKey)
	require.NotEqual(t, firstAuthorizationKey, secondAuthorizationKey, "authorization must not be replayed")
}

func TestDeployMultiTokenPoolRejectsInvalidReusedPoolConfigManager(t *testing.T) {
	tests := []struct {
		name        string
		mutate      func(t *testing.T, chain evm.Chain, refs []datastore.AddressRef)
		expectedErr string
	}{
		{
			name: "miswired recorded store",
			mutate: func(t *testing.T, _ evm.Chain, refs []datastore.AddressRef) {
				t.Helper()
				for i := range refs {
					if refs[i].Type == datastore.ContractType(connection_policy_store.ContractType) {
						refs[i].Address = common.HexToAddress("0xdead").Hex()
						return
					}
				}
				require.Fail(t, "missing ConnectionPolicyStore ref")
			},
			expectedErr: "static config ConnectionStore",
		},
		{
			name: "recorded manager is an EOA",
			mutate: func(t *testing.T, chain evm.Chain, refs []datastore.AddressRef) {
				t.Helper()
				for i := range refs {
					if refs[i].Type == datastore.ContractType(pool_config_manager.ContractType) {
						refs[i].Address = chain.DeployerKey.From.Hex()
						return
					}
				}
				require.Fail(t, "missing PoolConfigManager ref")
			},
			expectedErr: "failed to read static config of reused PoolConfigManager",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			e, err := environment.New(t.Context(), environment.WithEVMSimulated(t, []uint64{multiTokenPoolEVMTestChainSel}))
			require.NoError(t, err)
			prerequisites := deployMultiTokenPoolPrerequisites(t, e)
			chain := e.BlockChains.EVMChains()[multiTokenPoolEVMTestChainSel]
			firstOut := applyMultiTokenPoolDeployment(t, e, chain.DeployerKey.From.Hex())
			firstRefs, err := firstOut.DataStore.Addresses().Fetch()
			require.NoError(t, err)
			test.mutate(t, chain, firstRefs)
			installMultiTokenPoolDataStore(t, e, append(prerequisites, firstRefs...))

			nonceBefore, err := chain.Client.PendingNonceAt(t.Context(), chain.DeployerKey.From)
			require.NoError(t, err)
			_, err = cldf_ops.ExecuteSequence(e.OperationsBundle, sequences.DeployMultiTokenPool, e.BlockChains, multiTokenPoolSequenceInput(e, chain.DeployerKey.From.Hex()))
			require.ErrorContains(t, err, test.expectedErr)
			nonceAfter, err := chain.Client.PendingNonceAt(t.Context(), chain.DeployerKey.From)
			require.NoError(t, err)
			require.Equal(t, nonceBefore, nonceAfter)
		})
	}
}

func TestDeployMultiTokenPoolRejectsRetiredManagerOnReusedPool(t *testing.T) {
	e, err := environment.New(t.Context(), environment.WithEVMSimulated(t, []uint64{multiTokenPoolEVMTestChainSel}))
	require.NoError(t, err)
	prerequisites := deployMultiTokenPoolPrerequisites(t, e)
	chain := e.BlockChains.EVMChains()[multiTokenPoolEVMTestChainSel]
	firstOut := applyMultiTokenPoolDeployment(t, e, chain.DeployerKey.From.Hex())
	firstRefs, err := firstOut.DataStore.Addresses().Fetch()
	require.NoError(t, err)

	// Deploy a replacement PCM with the same valid static config, but leave the reused pool
	// pointing at the original PCM. The retry must not authorize the retired recorded manager.
	replacementQualifier := "replacement-pcm"
	replacementRef, err := contract.MaybeDeployContract(e.OperationsBundle, pool_config_manager.Deploy, chain, contract.DeployInput[pool_config_manager.ConstructorArgs]{
		ChainSelector:  multiTokenPoolEVMTestChainSel,
		TypeAndVersion: cldf.NewTypeAndVersion(pool_config_manager.ContractType, *pool_config_manager.Version),
		Args: pool_config_manager.ConstructorArgs{
			TokenAdminRegistry: common.HexToAddress(findMultiTokenPoolRef(t, prerequisites, token_admin_registry.ContractType).Address),
			ConnectionStore:    common.HexToAddress(findMultiTokenPoolRef(t, firstRefs, connection_policy_store.ContractType).Address),
			FeeStore:           common.HexToAddress(findMultiTokenPoolRef(t, firstRefs, fee_policy_store.ContractType).Address),
			SecurityStore:      common.HexToAddress(findMultiTokenPoolRef(t, firstRefs, security_policy_store.ContractType).Address),
			Admin:              chain.DeployerKey.From,
		},
		Qualifier: &replacementQualifier,
	}, nil)
	require.NoError(t, err)
	replacementRef.Qualifier = "customer-a"
	for i := range firstRefs {
		if firstRefs[i].Type == datastore.ContractType(pool_config_manager.ContractType) {
			firstRefs[i] = replacementRef
		}
	}
	installMultiTokenPoolDataStore(t, e, append(prerequisites, firstRefs...))

	nonceBefore, err := chain.Client.PendingNonceAt(t.Context(), chain.DeployerKey.From)
	require.NoError(t, err)
	_, err = cldf_ops.ExecuteSequence(e.OperationsBundle, sequences.AuthorizeMultiTokenPool, e.BlockChains, multiTokenPoolSequenceInput(e, chain.DeployerKey.From.Hex()))
	require.ErrorContains(t, err, "refusing to authorize a retired manager")
	nonceAfter, err := chain.Client.PendingNonceAt(t.Context(), chain.DeployerKey.From)
	require.NoError(t, err)
	require.Equal(t, nonceBefore, nonceAfter)

	connectionStore, err := connection_policy_store_bindings.NewConnectionPolicyStore(
		common.HexToAddress(findMultiTokenPoolRef(t, firstRefs, connection_policy_store.ContractType).Address), chain.Client)
	require.NoError(t, err)
	authorizedCallers, err := connectionStore.GetAllAuthorizedCallers(&bind.CallOpts{Context: t.Context()})
	require.NoError(t, err)
	require.NotContains(t, authorizedCallers, common.HexToAddress(replacementRef.Address))
}

func TestDeployMultiTokenPoolZkSyncRequiresFullyReusedStack(t *testing.T) {
	t.Run("missing deployment is rejected during preflight", func(t *testing.T) {
		e, err := environment.New(t.Context(), environment.WithEVMSimulated(t, []uint64{multiTokenPoolEVMTestChainSel}))
		require.NoError(t, err)
		deployMultiTokenPoolPrerequisites(t, e)
		chain := e.BlockChains.EVMChains()[multiTokenPoolEVMTestChainSel]
		chain.IsZkSyncVM = true
		zkSyncChains := cldf_chain.NewBlockChains(map[uint64]cldf_chain.BlockChain{multiTokenPoolEVMTestChainSel: chain})

		nonceBefore, err := chain.Client.PendingNonceAt(t.Context(), chain.DeployerKey.From)
		require.NoError(t, err)
		_, err = cldf_ops.ExecuteSequence(e.OperationsBundle, sequences.DeployMultiTokenPool, zkSyncChains, multiTokenPoolSequenceInput(e, chain.DeployerKey.From.Hex()))
		require.ErrorContains(t, err, "zkSync VM")
		require.ErrorContains(t, err, "only EVM bytecode")
		nonceAfter, err := chain.Client.PendingNonceAt(t.Context(), chain.DeployerKey.From)
		require.NoError(t, err)
		require.Equal(t, nonceBefore, nonceAfter)
	})

	t.Run("fully reused stack is allowed", func(t *testing.T) {
		e, err := environment.New(t.Context(), environment.WithEVMSimulated(t, []uint64{multiTokenPoolEVMTestChainSel}))
		require.NoError(t, err)
		prerequisites := deployMultiTokenPoolPrerequisites(t, e)
		chain := e.BlockChains.EVMChains()[multiTokenPoolEVMTestChainSel]
		firstOut := applyMultiTokenPoolDeployment(t, e, chain.DeployerKey.From.Hex())
		firstRefs, err := firstOut.DataStore.Addresses().Fetch()
		require.NoError(t, err)
		installMultiTokenPoolDataStore(t, e, append(prerequisites, firstRefs...))
		chain.IsZkSyncVM = true
		zkSyncChains := cldf_chain.NewBlockChains(map[uint64]cldf_chain.BlockChain{multiTokenPoolEVMTestChainSel: chain})

		nonceBefore, err := chain.Client.PendingNonceAt(t.Context(), chain.DeployerKey.From)
		require.NoError(t, err)
		report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, sequences.DeployMultiTokenPool, zkSyncChains, multiTokenPoolSequenceInput(e, chain.DeployerKey.From.Hex()))
		require.NoError(t, err)
		require.Empty(t, report.Output.BatchOps)
		authorizationReport, err := cldf_ops.ExecuteSequence(e.OperationsBundle, sequences.AuthorizeMultiTokenPool, zkSyncChains, multiTokenPoolSequenceInput(e, chain.DeployerKey.From.Hex()))
		require.NoError(t, err)
		require.Empty(t, authorizationReport.Output.BatchOps)
		nonceAfter, err := chain.Client.PendingNonceAt(t.Context(), chain.DeployerKey.From)
		require.NoError(t, err)
		require.Equal(t, nonceBefore, nonceAfter)
	})
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
	report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, sequences.AuthorizeMultiTokenPool, e.BlockChains, tokens.DeployMultiTokenPoolInput{
		ChainSelector:           multiTokenPoolEVMTestChainSel,
		PoolType:                string(burn_mint_multi_token_pool.ContractType),
		PoolVersion:             burn_mint_multi_token_pool.Version,
		PoolQualifier:           "customer-a",
		Admin:                   chain.DeployerKey.From.Hex(),
		SelectedTimelockAddress: newOwner.From.Hex(),
		ExistingDataStore:       e.DataStore,
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

func TestDeployMultiTokenPoolRejectsWrongSelectedTimelockForPreparedAuthorization(t *testing.T) {
	e, err := environment.New(t.Context(), environment.WithEVMSimulatedWithConfig(t, []uint64{multiTokenPoolEVMTestChainSel}, onchain.EVMSimLoaderConfig{NumAdditionalAccounts: 1}))
	require.NoError(t, err)
	prerequisites := deployMultiTokenPoolPrerequisites(t, e)
	chain := e.BlockChains.EVMChains()[multiTokenPoolEVMTestChainSel]
	firstOut := applyMultiTokenPoolDeployment(t, e, chain.DeployerKey.From.Hex())
	firstRefs, err := firstOut.DataStore.Addresses().Fetch()
	require.NoError(t, err)
	managerAddress := common.HexToAddress(findMultiTokenPoolRef(t, firstRefs, pool_config_manager.ContractType).Address)
	connectionStoreAddress := common.HexToAddress(findMultiTokenPoolRef(t, firstRefs, connection_policy_store.ContractType).Address)
	connectionStore, err := connection_policy_store_bindings.NewConnectionPolicyStore(connectionStoreAddress, chain.Client)
	require.NoError(t, err)

	removeTx, err := connectionStore.ApplyAuthorizedCallerUpdates(chain.DeployerKey, connection_policy_store_bindings.AuthorizedCallersAuthorizedCallerArgs{
		RemovedCallers: []common.Address{managerAddress},
	})
	require.NoError(t, err)
	_, err = chain.Confirm(removeTx)
	require.NoError(t, err)
	actualOwner := chain.Users[0]
	transferTx, err := connectionStore.TransferOwnership(chain.DeployerKey, actualOwner.From)
	require.NoError(t, err)
	_, err = chain.Confirm(transferTx)
	require.NoError(t, err)
	acceptTx, err := connectionStore.AcceptOwnership(actualOwner)
	require.NoError(t, err)
	_, err = chain.Confirm(acceptTx)
	require.NoError(t, err)
	installMultiTokenPoolDataStore(t, e, append(prerequisites, firstRefs...))

	input := multiTokenPoolSequenceInput(e, chain.DeployerKey.From.Hex())
	input.SelectedTimelockAddress = common.HexToAddress("0xdead").Hex()
	nonceBefore, err := chain.Client.PendingNonceAt(t.Context(), chain.DeployerKey.From)
	require.NoError(t, err)
	_, err = cldf_ops.ExecuteSequence(e.OperationsBundle, sequences.AuthorizeMultiTokenPool, e.BlockChains, input)
	require.ErrorContains(t, err, "not the selected timelock")
	require.ErrorContains(t, err, actualOwner.From.Hex())
	nonceAfter, err := chain.Client.PendingNonceAt(t.Context(), chain.DeployerKey.From)
	require.NoError(t, err)
	require.Equal(t, nonceBefore, nonceAfter)
	authorizedCallers, err := connectionStore.GetAllAuthorizedCallers(&bind.CallOpts{Context: t.Context()})
	require.NoError(t, err)
	require.NotContains(t, authorizedCallers, managerAddress)
}

func TestDeployMultiTokenPoolRequiresAnExplicitAdmin(t *testing.T) {
	e, err := environment.New(t.Context(), environment.WithEVMSimulated(t, []uint64{multiTokenPoolEVMTestChainSel}))
	require.NoError(t, err)
	deployMultiTokenPoolPrerequisites(t, e)

	_, err = cldf_ops.ExecuteSequence(e.OperationsBundle, sequences.DeployMultiTokenPool, e.BlockChains, tokens.DeployMultiTokenPoolInput{
		ChainSelector:     multiTokenPoolEVMTestChainSel,
		PoolType:          string(burn_mint_multi_token_pool.ContractType),
		PoolVersion:       burn_mint_multi_token_pool.Version,
		PoolQualifier:     "customer-a",
		ExistingDataStore: e.DataStore,
	})
	require.ErrorContains(t, err, "PoolConfigManager admin is required")
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

func TestDeployMultiTokenPoolVerifySelectsNewestRecordedRampBeforeReaderCompatibility(t *testing.T) {
	wrongTAR := common.HexToAddress("0xdead")
	t.Run("newer minor uses a compatible same-major reader and ignores retired ramps", func(t *testing.T) {
		err := verifyWithRamps(t, func(e *cldf.Environment, tar common.Address) []datastore.AddressRef {
			retired := deployTestOnRampV160(t, e, wrongTAR)
			newest := deployTestOnRampV200(t, e, tar)
			newest.Version = semver.MustParse("2.1.0")
			return []datastore.AddressRef{retired, newest}
		})
		require.NoError(t, err)
	})
	t.Run("unknown newer major is rejected", func(t *testing.T) {
		err := verifyWithRamps(t, func(e *cldf.Environment, tar common.Address) []datastore.AddressRef {
			newest := deployTestOnRampV200(t, e, tar)
			newest.Version = semver.MustParse("3.0.0")
			return []datastore.AddressRef{newest}
		})
		require.ErrorContains(t, err, "unsupported newest recorded ramp OnRamp 3.0.0")
		require.ErrorContains(t, err, "no compatible static-config reader")
	})
	t.Run("malformed address at the newest version is rejected", func(t *testing.T) {
		err := verifyWithRamps(t, func(e *cldf.Environment, tar common.Address) []datastore.AddressRef {
			newest := deployTestOnRampV200(t, e, tar)
			newest.Version = semver.MustParse("2.1.0")
			newest.Address = "not-an-address"
			return []datastore.AddressRef{newest}
		})
		require.ErrorContains(t, err, "newest recorded ramp OnRamp 2.1.0")
		require.ErrorContains(t, err, "invalid address")
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

// reusedStackEnvironment deploys the stack once and returns an environment whose datastore holds
// the prerequisites and the deployed refs, so a second verification reuses the stack.
func reusedStackEnvironment(t *testing.T, mutateRefs func(e *cldf.Environment, refs []datastore.AddressRef) []datastore.AddressRef) (*cldf.Environment, evm.Chain) {
	t.Helper()
	e, err := environment.New(t.Context(), environment.WithEVMSimulatedWithConfig(t, []uint64{multiTokenPoolEVMTestChainSel}, onchain.EVMSimLoaderConfig{NumAdditionalAccounts: 1}))
	require.NoError(t, err)
	prerequisites := deployMultiTokenPoolPrerequisites(t, e)
	chain := e.BlockChains.EVMChains()[multiTokenPoolEVMTestChainSel]
	firstOut := applyMultiTokenPoolDeployment(t, e, chain.DeployerKey.From.Hex())
	firstRefs, err := firstOut.DataStore.Addresses().Fetch()
	require.NoError(t, err)

	refs := append(prerequisites, firstRefs...)
	if mutateRefs != nil {
		refs = mutateRefs(e, refs)
	}
	ds := datastore.NewMemoryDataStore()
	for _, ref := range refs {
		require.NoError(t, ds.Addresses().Add(ref))
	}
	e.DataStore = ds.Seal()
	e.OperationsBundle = cldf_ops.NewBundle(e.OperationsBundle.GetContext, e.Logger, cldf_ops.NewMemoryReporter())
	return e, chain
}

func TestDeployMultiTokenPoolVerifyComparesConfiguredAdminWithReusedPoolConfigManager(t *testing.T) {
	e, chain := reusedStackEnvironment(t, nil)
	changeset := tokens.DeployMultiTokenPools(tokens.GetMultiTokenPoolAdapterRegistry(), changeset_utils.NewMCMSReaderRegistry())

	require.NoError(t, changeset.VerifyPreconditions(*e, multiTokenPoolTestConfig(chain.DeployerKey.From.Hex())), "the configured admin matches the existing default admin")

	err := changeset.VerifyPreconditions(*e, multiTokenPoolTestConfig(chain.Users[0].From.Hex()))
	require.ErrorContains(t, err, "has default admin")
	require.ErrorContains(t, err, "the admin is not changed on reuse")
}

func TestDeployMultiTokenPoolVerifyComparesReusedPoolRouterAndRMNProxy(t *testing.T) {
	tests := []struct {
		name        string
		replace     cldf.ContractType
		redeploy    func(t *testing.T, e *cldf.Environment, chain evm.Chain, refs []datastore.AddressRef) datastore.AddressRef
		expectedErr string
	}{
		{
			name:    "router",
			replace: router.ContractType,
			redeploy: func(t *testing.T, e *cldf.Environment, chain evm.Chain, refs []datastore.AddressRef) datastore.AddressRef {
				ref, err := contract.MaybeDeployContract(e.OperationsBundle, router.Deploy, chain, contract.DeployInput[router.ConstructorArgs]{
					ChainSelector:  multiTokenPoolEVMTestChainSel,
					TypeAndVersion: cldf.NewTypeAndVersion(router.ContractType, *router.Version),
					Args: router.ConstructorArgs{
						// Different arguments from the original router: identical deploy input would
						// replay the cached deployment instead of creating a new contract.
						WrappedNative: common.HexToAddress("0x5"),
						RMNProxy:      common.HexToAddress(findMultiTokenPoolRef(t, refs, rmn_proxy.ContractType).Address),
					},
				}, nil)
				require.NoError(t, err)
				return ref
			},
			expectedErr: "uses Router",
		},
		{
			name:    "RMNProxy",
			replace: rmn_proxy.ContractType,
			redeploy: func(t *testing.T, e *cldf.Environment, chain evm.Chain, _ []datastore.AddressRef) datastore.AddressRef {
				ref, err := contract.MaybeDeployContract(e.OperationsBundle, rmn_proxy.Deploy, chain, contract.DeployInput[rmn_proxy.ConstructorArgs]{
					ChainSelector:  multiTokenPoolEVMTestChainSel,
					TypeAndVersion: cldf.NewTypeAndVersion(rmn_proxy.ContractType, *rmn_proxy.Version),
					Args:           rmn_proxy.ConstructorArgs{RMN: common.HexToAddress("0x3")},
				}, nil)
				require.NoError(t, err)
				return ref
			},
			expectedErr: "uses RMNProxy",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			e, chain := reusedStackEnvironment(t, func(e *cldf.Environment, refs []datastore.AddressRef) []datastore.AddressRef {
				replacement := test.redeploy(t, e, e.BlockChains.EVMChains()[multiTokenPoolEVMTestChainSel], refs)
				kept := make([]datastore.AddressRef, 0, len(refs))
				for _, ref := range refs {
					if ref.Type != datastore.ContractType(test.replace) {
						kept = append(kept, ref)
					}
				}
				return append(kept, replacement)
			})
			changeset := tokens.DeployMultiTokenPools(tokens.GetMultiTokenPoolAdapterRegistry(), changeset_utils.NewMCMSReaderRegistry())
			err := changeset.VerifyPreconditions(*e, multiTokenPoolTestConfig(chain.DeployerKey.From.Hex()))
			require.ErrorContains(t, err, test.expectedErr)
		})
	}
}
