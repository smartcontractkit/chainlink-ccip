package sequences

import (
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	adapters_v1_0_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_1_0/operations/burn_mint_multi_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_1_0/operations/connection_policy_store"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_1_0/operations/fee_policy_store"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_1_0/operations/pool_config_manager"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_1_0/operations/security_policy_store"
	"github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

// ValidateDeployMultiTokenPoolInput checks the pool version, type, qualifier, and admin address.
func ValidateDeployMultiTokenPoolInput(input tokens.DeployMultiTokenPoolInput) error {
	if input.PoolVersion == nil {
		return errors.New("pool version is required")
	}
	if !input.PoolVersion.Equal(burn_mint_multi_token_pool.Version) {
		return fmt.Errorf("unsupported multi-token pool version %s: only %s is supported", input.PoolVersion, burn_mint_multi_token_pool.Version)
	}
	if input.PoolType != string(burn_mint_multi_token_pool.ContractType) {
		return fmt.Errorf("unsupported multi-token pool type %q: only %s is supported", input.PoolType, burn_mint_multi_token_pool.ContractType)
	}
	if input.PoolQualifier == "" {
		return errors.New("pool qualifier is required")
	}
	if input.Admin != "" {
		if !common.IsHexAddress(input.Admin) {
			return fmt.Errorf("invalid PoolConfigManager admin address %q", input.Admin)
		}
		if common.HexToAddress(input.Admin) == (common.Address{}) {
			return errors.New("PoolConfigManager admin must be non-zero")
		}
	}
	return nil
}

// multiTokenPoolDependencies holds the existing contracts the stack is wired to.
type multiTokenPoolDependencies struct {
	tokenAdminRegistry common.Address
	rmnProxy           common.Address
	router             common.Address
}

// ValidateDeployMultiTokenPoolState checks that the TokenAdminRegistry, RMNProxy, and Router
// resolve, that the TokenAdminRegistry matches the recorded ramps', and that any existing stack
// is consistent. It sends no transactions.
func ValidateDeployMultiTokenPoolState(b cldf_ops.Bundle, chain evm.Chain, ds datastore.DataStore, input tokens.DeployMultiTokenPoolInput) error {
	_, err := validateDeployMultiTokenPoolState(b, chain, ds, input)
	return err
}

func validateDeployMultiTokenPoolState(b cldf_ops.Bundle, chain evm.Chain, ds datastore.DataStore, input tokens.DeployMultiTokenPoolInput) (multiTokenPoolDependencies, error) {
	if ds == nil {
		return multiTokenPoolDependencies{}, errors.New("existing datastore is required")
	}
	// Resolve the dependencies here so a missing ref fails verification.
	poolUtils := adapters_v1_0_0.EVMTokenBase{}
	var deps multiTokenPoolDependencies
	var err error
	if deps.tokenAdminRegistry, err = poolUtils.GetTokenAdminRegistryAddress(ds, input.ChainSelector); err != nil {
		return multiTokenPoolDependencies{}, err
	}
	if deps.rmnProxy, err = poolUtils.GetRMNProxyAddress(ds, input.ChainSelector); err != nil {
		return multiTokenPoolDependencies{}, err
	}
	if deps.router, err = poolUtils.ResolveRouterAddress(ds, input.ChainSelector, nil); err != nil {
		return multiTokenPoolDependencies{}, err
	}
	existing := ds.Addresses().Filter(datastore.AddressRefByChainSelector(input.ChainSelector))
	if err := checkReusedStackIsConsistent(existing, input.PoolQualifier, input.ChainSelector); err != nil {
		return multiTokenPoolDependencies{}, err
	}
	if err := checkZkSyncMultiTokenPoolDeployment(chain, existing, input.PoolQualifier); err != nil {
		return multiTokenPoolDependencies{}, err
	}
	if err := checkTokenAdminRegistryMatchesRamps(b, chain, existing, deps.tokenAdminRegistry); err != nil {
		return multiTokenPoolDependencies{}, err
	}
	if err := checkReusedStackOnChain(b, chain, existing, input.PoolQualifier, deps.tokenAdminRegistry); err != nil {
		return multiTokenPoolDependencies{}, err
	}
	return deps, nil
}

// DeployMultiTokenPool deploys a BurnMintMultiTokenPool and its configuration stack on one
// EVM chain. The pool is deliberately not registered for any token by this sequence.
var DeployMultiTokenPool = cldf_ops.NewSequence(
	"deploy-multi-token-pool",
	burn_mint_multi_token_pool.Version,
	"Deploy the EVM multi-token pool contracts",
	func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input tokens.DeployMultiTokenPoolInput) (sequences.OnChainOutput, error) {
		chain, ok := chains.EVMChains()[input.ChainSelector]
		if !ok {
			return sequences.OnChainOutput{}, fmt.Errorf("EVM chain with selector %d not found", input.ChainSelector)
		}
		if err := ValidateDeployMultiTokenPoolInput(input); err != nil {
			return sequences.OnChainOutput{}, err
		}
		// Validate before the first deployment so a failure leaves nothing orphaned.
		deps, err := validateDeployMultiTokenPoolState(b, chain, input.ExistingDataStore, input)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}

		admin := common.Address{}
		if chain.DeployerKey != nil {
			admin = chain.DeployerKey.From
		}
		if input.Admin != "" {
			admin = common.HexToAddress(input.Admin)
		}
		if admin == (common.Address{}) {
			return sequences.OnChainOutput{}, errors.New("PoolConfigManager admin must be non-zero")
		}

		existingAddresses := input.ExistingDataStore.Addresses().Filter(datastore.AddressRefByChainSelector(input.ChainSelector))
		connectionStoreRef, err := maybeDeployMultiTokenPoolContract(b, connection_policy_store.Deploy, chain, contract.DeployInput[connection_policy_store.ConstructorArgs]{
			ChainSelector:  input.ChainSelector,
			TypeAndVersion: cldf_deployment.NewTypeAndVersion(connection_policy_store.ContractType, *connection_policy_store.Version),
			Args:           connection_policy_store.ConstructorArgs{},
			Qualifier:      &input.PoolQualifier,
		}, existingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy ConnectionPolicyStore: %w", err)
		}
		feeStoreRef, err := maybeDeployMultiTokenPoolContract(b, fee_policy_store.Deploy, chain, contract.DeployInput[fee_policy_store.ConstructorArgs]{
			ChainSelector:  input.ChainSelector,
			TypeAndVersion: cldf_deployment.NewTypeAndVersion(fee_policy_store.ContractType, *fee_policy_store.Version),
			Args:           fee_policy_store.ConstructorArgs{},
			Qualifier:      &input.PoolQualifier,
		}, existingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy FeePolicyStore: %w", err)
		}
		securityStoreRef, err := maybeDeployMultiTokenPoolContract(b, security_policy_store.Deploy, chain, contract.DeployInput[security_policy_store.ConstructorArgs]{
			ChainSelector:  input.ChainSelector,
			TypeAndVersion: cldf_deployment.NewTypeAndVersion(security_policy_store.ContractType, *security_policy_store.Version),
			Args:           security_policy_store.ConstructorArgs{},
			Qualifier:      &input.PoolQualifier,
		}, existingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy SecurityPolicyStore: %w", err)
		}

		connectionStore := common.HexToAddress(connectionStoreRef.Address)
		feeStore := common.HexToAddress(feeStoreRef.Address)
		securityStore := common.HexToAddress(securityStoreRef.Address)
		poolConfigManagerRef, err := maybeDeployMultiTokenPoolContract(b, pool_config_manager.Deploy, chain, contract.DeployInput[pool_config_manager.ConstructorArgs]{
			ChainSelector:  input.ChainSelector,
			TypeAndVersion: cldf_deployment.NewTypeAndVersion(pool_config_manager.ContractType, *pool_config_manager.Version),
			Args: pool_config_manager.ConstructorArgs{
				TokenAdminRegistry: deps.tokenAdminRegistry,
				ConnectionStore:    connectionStore,
				FeeStore:           feeStore,
				SecurityStore:      securityStore,
				Admin:              admin,
			},
			Qualifier: &input.PoolQualifier,
		}, existingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy PoolConfigManager: %w", err)
		}
		poolConfigManagerAddress := common.HexToAddress(poolConfigManagerRef.Address)

		poolRef, err := maybeDeployMultiTokenPoolContract(b, burn_mint_multi_token_pool.Deploy, chain, contract.DeployInput[burn_mint_multi_token_pool.ConstructorArgs]{
			ChainSelector:  input.ChainSelector,
			TypeAndVersion: cldf_deployment.NewTypeAndVersion(burn_mint_multi_token_pool.ContractType, *burn_mint_multi_token_pool.Version),
			Args: burn_mint_multi_token_pool.ConstructorArgs{
				RmnProxy:          deps.rmnProxy,
				Router:            deps.router,
				PoolConfigManager: poolConfigManagerAddress,
			},
			Qualifier: &input.PoolQualifier,
		}, existingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy BurnMintMultiTokenPool: %w", err)
		}

		return sequences.OnChainOutput{
			Addresses: []datastore.AddressRef{connectionStoreRef, feeStoreRef, securityStoreRef, poolConfigManagerRef, poolRef},
		}, nil
	},
)

// AuthorizeMultiTokenPool authorizes the recorded PoolConfigManager on the three recorded stores.
// It is separate from deployment so a fresh authorization run never discards replayed deployment
// reports. The generic changeset adds newly deployed refs to ExistingDataStore before calling it.
var AuthorizeMultiTokenPool = cldf_ops.NewSequence(
	"authorize-multi-token-pool",
	burn_mint_multi_token_pool.Version,
	"Authorize the EVM multi-token pool's PoolConfigManager on its policy stores",
	func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input tokens.DeployMultiTokenPoolInput) (sequences.OnChainOutput, error) {
		chain, ok := chains.EVMChains()[input.ChainSelector]
		if !ok {
			return sequences.OnChainOutput{}, fmt.Errorf("EVM chain with selector %d not found", input.ChainSelector)
		}
		if err := ValidateDeployMultiTokenPoolInput(input); err != nil {
			return sequences.OnChainOutput{}, err
		}
		// Authorization must revalidate fresh on-chain wiring even when deployment was replayed.
		if _, err := validateDeployMultiTokenPoolState(b, chain, input.ExistingDataStore, input); err != nil {
			return sequences.OnChainOutput{}, err
		}
		existingAddresses := input.ExistingDataStore.Addresses().Filter(datastore.AddressRefByChainSelector(input.ChainSelector))
		stack, err := resolveMultiTokenPoolStackAddresses(existingAddresses, input.PoolQualifier, input.ChainSelector)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}

		writes := make([]contract.WriteOutput, 0, 3)
		callers := func(addr common.Address) []common.Address { return []common.Address{addr} }
		authorizations := []struct {
			name      string
			authorize func() (contract.WriteOutput, bool, error)
		}{
			{"ConnectionPolicyStore", func() (contract.WriteOutput, bool, error) {
				return authorizePoolConfigManager(b, chain, stack.connectionStore, stack.poolConfigManager,
					connection_policy_store.GetAllAuthorizedCallers, connection_policy_store.ApplyAuthorizedCallerUpdates,
					connection_policy_store.Owner, connection_policy_store.AuthorizedCallerArgs{AddedCallers: callers(stack.poolConfigManager)},
					input.SelectedTimelockAddress)
			}},
			{"FeePolicyStore", func() (contract.WriteOutput, bool, error) {
				return authorizePoolConfigManager(b, chain, stack.feeStore, stack.poolConfigManager,
					fee_policy_store.GetAllAuthorizedCallers, fee_policy_store.ApplyAuthorizedCallerUpdates,
					fee_policy_store.Owner, fee_policy_store.AuthorizedCallerArgs{AddedCallers: callers(stack.poolConfigManager)},
					input.SelectedTimelockAddress)
			}},
			{"SecurityPolicyStore", func() (contract.WriteOutput, bool, error) {
				return authorizePoolConfigManager(b, chain, stack.securityStore, stack.poolConfigManager,
					security_policy_store.GetAllAuthorizedCallers, security_policy_store.ApplyAuthorizedCallerUpdates,
					security_policy_store.Owner, security_policy_store.AuthorizedCallerArgs{AddedCallers: callers(stack.poolConfigManager)},
					input.SelectedTimelockAddress)
			}},
		}
		for _, a := range authorizations {
			write, changed, err := a.authorize()
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to authorize PoolConfigManager on %s: %w", a.name, err)
			}
			if changed {
				writes = append(writes, write)
			}
		}

		batchOps := make([]mcms_types.BatchOperation, 0, 1)
		if len(writes) > 0 {
			batchOp, err := contract.NewBatchOperationFromWrites(writes)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to build store authorization batch: %w", err)
			}
			if len(batchOp.Transactions) > 0 {
				batchOps = append(batchOps, batchOp)
			}
		}
		return sequences.OnChainOutput{BatchOps: batchOps}, nil
	},
)

// findExistingDeployment mirrors the type/version/qualifier matching of MaybeDeployContract.
func findExistingDeployment(existing []datastore.AddressRef, contractType cldf_deployment.ContractType, version *semver.Version, qualifier string) (datastore.AddressRef, bool) {
	index := slices.IndexFunc(existing, func(ref datastore.AddressRef) bool {
		return ref.Type == datastore.ContractType(contractType) &&
			ref.Version != nil && ref.Version.String() == version.String() &&
			ref.Qualifier == qualifier
	})
	if index < 0 {
		return datastore.AddressRef{}, false
	}
	return existing[index], true
}

func hasExistingDeployment(existing []datastore.AddressRef, contractType cldf_deployment.ContractType, version *semver.Version, qualifier string) bool {
	_, ok := findExistingDeployment(existing, contractType, version, qualifier)
	return ok
}

type multiTokenPoolStackAddresses struct {
	connectionStore   common.Address
	feeStore          common.Address
	securityStore     common.Address
	poolConfigManager common.Address
	pool              common.Address
}

// resolveMultiTokenPoolStackAddresses resolves all five refs instead of deploying missing contracts.
func resolveMultiTokenPoolStackAddresses(existing []datastore.AddressRef, qualifier string, selector uint64) (multiTokenPoolStackAddresses, error) {
	var stack multiTokenPoolStackAddresses
	contracts := []struct {
		name         string
		contractType cldf_deployment.ContractType
		version      *semver.Version
		destination  *common.Address
	}{
		{"ConnectionPolicyStore", connection_policy_store.ContractType, connection_policy_store.Version, &stack.connectionStore},
		{"FeePolicyStore", fee_policy_store.ContractType, fee_policy_store.Version, &stack.feeStore},
		{"SecurityPolicyStore", security_policy_store.ContractType, security_policy_store.Version, &stack.securityStore},
		{"PoolConfigManager", pool_config_manager.ContractType, pool_config_manager.Version, &stack.poolConfigManager},
		{"BurnMintMultiTokenPool", burn_mint_multi_token_pool.ContractType, burn_mint_multi_token_pool.Version, &stack.pool},
	}
	for _, c := range contracts {
		ref, ok := findExistingDeployment(existing, c.contractType, c.version, qualifier)
		if !ok {
			return multiTokenPoolStackAddresses{}, fmt.Errorf("missing %s ref on chain %d for qualifier %q", c.name, selector, qualifier)
		}
		if !common.IsHexAddress(ref.Address) {
			return multiTokenPoolStackAddresses{}, fmt.Errorf("recorded %s address %q on chain %d is invalid", c.name, ref.Address, selector)
		}
		*c.destination = common.HexToAddress(ref.Address)
	}
	return stack, nil
}

// checkReusedStackIsConsistent rejects a stack where a dependent contract exists but one of its
// constructor dependencies does not: PoolConfigManager without all three stores, or the pool
// without its PoolConfigManager. Constructor wiring of reused contracts is not reconciled.
func checkReusedStackIsConsistent(existing []datastore.AddressRef, qualifier string, selector uint64) error {
	stores := []struct {
		name string
		ok   bool
	}{
		{"ConnectionPolicyStore", hasExistingDeployment(existing, connection_policy_store.ContractType, connection_policy_store.Version, qualifier)},
		{"FeePolicyStore", hasExistingDeployment(existing, fee_policy_store.ContractType, fee_policy_store.Version, qualifier)},
		{"SecurityPolicyStore", hasExistingDeployment(existing, security_policy_store.ContractType, security_policy_store.Version, qualifier)},
	}
	managerExists := hasExistingDeployment(existing, pool_config_manager.ContractType, pool_config_manager.Version, qualifier)
	poolExists := hasExistingDeployment(existing, burn_mint_multi_token_pool.ContractType, burn_mint_multi_token_pool.Version, qualifier)
	if managerExists {
		for _, store := range stores {
			if !store.ok {
				return fmt.Errorf("PoolConfigManager exists on chain %d for qualifier %q but %s does not; its store wiring would be stale. Remove the stale refs or use a new pool qualifier", selector, qualifier, store.name)
			}
		}
	}
	if poolExists && !managerExists {
		return fmt.Errorf("BurnMintMultiTokenPool exists on chain %d for qualifier %q but PoolConfigManager does not; its manager wiring would be stale. Remove the stale refs or use a new pool qualifier", selector, qualifier)
	}
	return nil
}

// checkZkSyncMultiTokenPoolDeployment rejects partial stacks before deployment: these operations
// carry EVM bytecode only. A complete set of existing refs can still be reused and authorized.
func checkZkSyncMultiTokenPoolDeployment(chain evm.Chain, existing []datastore.AddressRef, qualifier string) error {
	if !chain.IsZkSyncVM {
		return nil
	}
	contracts := []struct {
		name         string
		contractType cldf_deployment.ContractType
		version      *semver.Version
	}{
		{"ConnectionPolicyStore", connection_policy_store.ContractType, connection_policy_store.Version},
		{"FeePolicyStore", fee_policy_store.ContractType, fee_policy_store.Version},
		{"SecurityPolicyStore", security_policy_store.ContractType, security_policy_store.Version},
		{"PoolConfigManager", pool_config_manager.ContractType, pool_config_manager.Version},
		{"BurnMintMultiTokenPool", burn_mint_multi_token_pool.ContractType, burn_mint_multi_token_pool.Version},
	}
	missing := make([]string, 0, len(contracts))
	for _, c := range contracts {
		if !hasExistingDeployment(existing, c.contractType, c.version, qualifier) {
			missing = append(missing, c.name)
		}
	}
	if len(missing) > 0 {
		return fmt.Errorf(
			"cannot deploy multi-token pool stack on zkSync VM chain %d: operations contain only EVM bytecode; missing existing refs for %s",
			chain.Selector, strings.Join(missing, ", "),
		)
	}
	return nil
}

// checkReusedStackOnChain verifies constructor wiring that the datastore refs alone cannot prove.
// All reads bypass the operations cache so retries cannot authorize a stale or retired manager.
func checkReusedStackOnChain(
	b cldf_ops.Bundle,
	chain evm.Chain,
	existing []datastore.AddressRef,
	qualifier string,
	tokenAdminRegistry common.Address,
) error {
	managerRef, managerExists := findExistingDeployment(existing, pool_config_manager.ContractType, pool_config_manager.Version, qualifier)
	if !managerExists {
		return nil
	}
	if !common.IsHexAddress(managerRef.Address) {
		return fmt.Errorf("recorded PoolConfigManager address %q on chain %d is invalid", managerRef.Address, chain.Selector)
	}
	if b.GetContext == nil {
		return errors.New("an operations bundle is required to read the reused multi-token pool stack")
	}
	managerAddress := common.HexToAddress(managerRef.Address)
	staticConfig, err := cldf_ops.ExecuteOperation(
		b,
		pool_config_manager.GetStaticConfig,
		chain,
		contract.FunctionInput[struct{}]{ChainSelector: chain.Selector, Address: managerAddress},
		cldf_ops.WithForceExecute[contract.FunctionInput[struct{}], evm.Chain](),
	)
	if err != nil {
		return fmt.Errorf("failed to read static config of reused PoolConfigManager at %s on chain %d: %w", managerRef.Address, chain.Selector, err)
	}
	if staticConfig.Output.TokenAdminRegistry != tokenAdminRegistry {
		return fmt.Errorf(
			"reused PoolConfigManager at %s on chain %d has TokenAdminRegistry %s but the recorded TokenAdminRegistry is %s",
			managerRef.Address, chain.Selector, staticConfig.Output.TokenAdminRegistry, tokenAdminRegistry,
		)
	}

	stores := []struct {
		name         string
		contractType cldf_deployment.ContractType
		version      *semver.Version
		configured   common.Address
	}{
		{"ConnectionStore", connection_policy_store.ContractType, connection_policy_store.Version, staticConfig.Output.ConnectionStore},
		{"FeeStore", fee_policy_store.ContractType, fee_policy_store.Version, staticConfig.Output.FeeStore},
		{"SecurityStore", security_policy_store.ContractType, security_policy_store.Version, staticConfig.Output.SecurityStore},
	}
	for _, store := range stores {
		storeRef, ok := findExistingDeployment(existing, store.contractType, store.version, qualifier)
		if !ok {
			return fmt.Errorf("reused PoolConfigManager at %s on chain %d has no recorded %s ref", managerRef.Address, chain.Selector, store.name)
		}
		if !common.IsHexAddress(storeRef.Address) {
			return fmt.Errorf("recorded %s address %q on chain %d is invalid", store.name, storeRef.Address, chain.Selector)
		}
		recordedAddress := common.HexToAddress(storeRef.Address)
		if store.configured != recordedAddress {
			return fmt.Errorf(
				"reused PoolConfigManager at %s on chain %d has static config %s %s but the recorded address is %s",
				managerRef.Address, chain.Selector, store.name, store.configured, recordedAddress,
			)
		}
	}

	poolRef, poolExists := findExistingDeployment(existing, burn_mint_multi_token_pool.ContractType, burn_mint_multi_token_pool.Version, qualifier)
	if !poolExists {
		return nil
	}
	if !common.IsHexAddress(poolRef.Address) {
		return fmt.Errorf("recorded BurnMintMultiTokenPool address %q on chain %d is invalid", poolRef.Address, chain.Selector)
	}
	poolManager, err := cldf_ops.ExecuteOperation(
		b,
		burn_mint_multi_token_pool.GetPoolConfigManager,
		chain,
		contract.FunctionInput[struct{}]{ChainSelector: chain.Selector, Address: common.HexToAddress(poolRef.Address)},
		cldf_ops.WithForceExecute[contract.FunctionInput[struct{}], evm.Chain](),
	)
	if err != nil {
		return fmt.Errorf("failed to read PoolConfigManager of reused BurnMintMultiTokenPool at %s on chain %d: %w", poolRef.Address, chain.Selector, err)
	}
	if poolManager.Output != managerAddress {
		return fmt.Errorf(
			"reused BurnMintMultiTokenPool at %s on chain %d uses PoolConfigManager %s but the recorded PoolConfigManager is %s; refusing to authorize a retired manager",
			poolRef.Address, chain.Selector, poolManager.Output, managerAddress,
		)
	}
	return nil
}

// maybeDeployMultiTokenPoolContract deploys the contract unless a matching one already exists.
// MaybeDeployContract does the matching.
func maybeDeployMultiTokenPoolContract[ARGS any](
	b cldf_ops.Bundle,
	op *cldf_ops.Operation[contract.DeployInput[ARGS], datastore.AddressRef, evm.Chain],
	chain evm.Chain,
	input contract.DeployInput[ARGS],
	existingAddresses []datastore.AddressRef,
) (datastore.AddressRef, error) {
	ref, err := contract.MaybeDeployContract(b, op, chain, input, existingAddresses)
	if err != nil {
		return datastore.AddressRef{}, err
	}
	reused := slices.ContainsFunc(existingAddresses, func(existing datastore.AddressRef) bool {
		return existing.Key() == ref.Key() && existing.Address == ref.Address
	})
	if reused {
		b.Logger.Infof("%s already deployed on chain %d at %s; skipping deployment", ref.Type, chain.Selector, ref.Address)
	}
	return ref, nil
}

// authorizePoolConfigManager adds the PCM to a store only when it is not already authorized.
// The read and write bypass the operations cache so retries see current on-chain state.
func authorizePoolConfigManager[ARGS any](
	b cldf_ops.Bundle,
	chain evm.Chain,
	storeAddress common.Address,
	poolConfigManagerAddress common.Address,
	getAllAuthorizedCallers *cldf_ops.Operation[contract.FunctionInput[struct{}], []common.Address, evm.Chain],
	applyAuthorizedCallerUpdates *cldf_ops.Operation[contract.FunctionInput[ARGS], contract.WriteOutput, evm.Chain],
	getStoreOwner *cldf_ops.Operation[contract.FunctionInput[struct{}], common.Address, evm.Chain],
	args ARGS,
	selectedTimelockAddress string,
) (contract.WriteOutput, bool, error) {
	authorizedCallers, err := cldf_ops.ExecuteOperation(
		b,
		getAllAuthorizedCallers,
		chain,
		contract.FunctionInput[struct{}]{ChainSelector: chain.Selector, Address: storeAddress},
		cldf_ops.WithForceExecute[contract.FunctionInput[struct{}], evm.Chain](),
	)
	if err != nil {
		return contract.WriteOutput{}, false, err
	}
	if slices.Contains(authorizedCallers.Output, poolConfigManagerAddress) {
		b.Logger.Infof(
			"PoolConfigManager %s already authorized on store %s on chain %d; skipping authorization",
			poolConfigManagerAddress,
			storeAddress,
			chain.Selector,
		)
		return contract.WriteOutput{}, false, nil
	}

	write, err := cldf_ops.ExecuteOperation(
		b,
		applyAuthorizedCallerUpdates,
		chain,
		contract.FunctionInput[ARGS]{
			ChainSelector: chain.Selector,
			Address:       storeAddress,
			Args:          args,
		},
		cldf_ops.WithForceExecute[contract.FunctionInput[ARGS], evm.Chain](),
	)
	if err != nil {
		return contract.WriteOutput{}, false, err
	}
	if !write.Output.Executed() {
		if err := validatePreparedStoreAuthorizationOwner(b, chain, storeAddress, getStoreOwner, selectedTimelockAddress); err != nil {
			return contract.WriteOutput{}, false, err
		}
	}
	return write.Output, true, nil
}

// validatePreparedStoreAuthorizationOwner prevents an authorization proposal from being built for
// an owner other than the timelock selected by the changeset. Deployer-executed writes skip it.
func validatePreparedStoreAuthorizationOwner(
	b cldf_ops.Bundle,
	chain evm.Chain,
	storeAddress common.Address,
	getStoreOwner *cldf_ops.Operation[contract.FunctionInput[struct{}], common.Address, evm.Chain],
	selectedTimelockAddress string,
) error {
	if selectedTimelockAddress == "" {
		return fmt.Errorf("selected timelock address is required to prepare authorization on store %s on chain %d", storeAddress, chain.Selector)
	}
	if !common.IsHexAddress(selectedTimelockAddress) {
		return fmt.Errorf("selected timelock address %q is invalid for store %s on chain %d", selectedTimelockAddress, storeAddress, chain.Selector)
	}
	owner, err := cldf_ops.ExecuteOperation(
		b,
		getStoreOwner,
		chain,
		contract.FunctionInput[struct{}]{ChainSelector: chain.Selector, Address: storeAddress},
		cldf_ops.WithForceExecute[contract.FunctionInput[struct{}], evm.Chain](),
	)
	if err != nil {
		return fmt.Errorf("failed to read owner of store %s on chain %d for selected timelock validation: %w", storeAddress, chain.Selector, err)
	}
	selectedTimelock := common.HexToAddress(selectedTimelockAddress)
	if owner.Output != selectedTimelock {
		return fmt.Errorf(
			"store %s on chain %d is owned by %s, not the selected timelock %s; refusing to prepare an unexecutable authorization",
			storeAddress, chain.Selector, owner.Output, selectedTimelock,
		)
	}
	return nil
}
