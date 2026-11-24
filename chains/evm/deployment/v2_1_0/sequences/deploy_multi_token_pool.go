package sequences

import (
	"errors"
	"fmt"
	"slices"

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

// DeployMultiTokenPool deploys a BurnMintMultiTokenPool and its configuration stack on one
// EVM chain. The pool is deliberately not registered for any token by this sequence.
var DeployMultiTokenPool = cldf_ops.NewSequence(
	"deploy-multi-token-pool",
	burn_mint_multi_token_pool.Version,
	"Deploy the EVM multi-token pool contracts and authorize their PoolConfigManager",
	func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input tokens.DeployMultiTokenPoolInput) (sequences.OnChainOutput, error) {
		chain, ok := chains.EVMChains()[input.ChainSelector]
		if !ok {
			return sequences.OnChainOutput{}, fmt.Errorf("EVM chain with selector %d not found", input.ChainSelector)
		}
		if input.PoolVersion == nil {
			return sequences.OnChainOutput{}, errors.New("pool version is required")
		}
		if !input.PoolVersion.Equal(burn_mint_multi_token_pool.Version) {
			return sequences.OnChainOutput{}, fmt.Errorf("unsupported multi-token pool version %s: only %s is supported", input.PoolVersion, burn_mint_multi_token_pool.Version)
		}
		if input.PoolType != string(burn_mint_multi_token_pool.ContractType) {
			return sequences.OnChainOutput{}, fmt.Errorf("unsupported multi-token pool type %q: only %s is supported", input.PoolType, burn_mint_multi_token_pool.ContractType)
		}
		if input.PoolQualifier == "" {
			return sequences.OnChainOutput{}, errors.New("pool qualifier is required")
		}
		if input.ExistingDataStore == nil {
			return sequences.OnChainOutput{}, errors.New("existing datastore is required")
		}

		admin := common.Address{}
		if chain.DeployerKey != nil {
			admin = chain.DeployerKey.From
		}
		if input.Admin != "" {
			if !common.IsHexAddress(input.Admin) {
				return sequences.OnChainOutput{}, fmt.Errorf("invalid PoolConfigManager admin address %q", input.Admin)
			}
			admin = common.HexToAddress(input.Admin)
		}
		if admin == (common.Address{}) {
			return sequences.OnChainOutput{}, errors.New("PoolConfigManager admin must be non-zero")
		}

		poolUtils := adapters_v1_0_0.EVMTokenBase{}
		tokenAdminRegistry, err := poolUtils.GetTokenAdminRegistryAddress(input.ExistingDataStore, input.ChainSelector)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		rmnProxy, err := poolUtils.GetRMNProxyAddress(input.ExistingDataStore, input.ChainSelector)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		router, err := poolUtils.ResolveRouterAddress(input.ExistingDataStore, input.ChainSelector, nil)
		if err != nil {
			return sequences.OnChainOutput{}, err
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
				TokenAdminRegistry: tokenAdminRegistry,
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
				RmnProxy:          rmnProxy,
				Router:            router,
				PoolConfigManager: poolConfigManagerAddress,
			},
			Qualifier: &input.PoolQualifier,
		}, existingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy BurnMintMultiTokenPool: %w", err)
		}

		writes := make([]contract.WriteOutput, 0, 3)
		connectionWrite, changed, err := authorizePoolConfigManager(
			b, chain, connectionStore, poolConfigManagerAddress,
			connection_policy_store.GetAllAuthorizedCallers,
			connection_policy_store.ApplyAuthorizedCallerUpdates,
			connection_policy_store.AuthorizedCallerArgs{AddedCallers: []common.Address{poolConfigManagerAddress}},
		)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to authorize PoolConfigManager on ConnectionPolicyStore: %w", err)
		}
		if changed {
			writes = append(writes, connectionWrite)
		}
		feeWrite, changed, err := authorizePoolConfigManager(
			b, chain, feeStore, poolConfigManagerAddress,
			fee_policy_store.GetAllAuthorizedCallers,
			fee_policy_store.ApplyAuthorizedCallerUpdates,
			fee_policy_store.AuthorizedCallerArgs{AddedCallers: []common.Address{poolConfigManagerAddress}},
		)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to authorize PoolConfigManager on FeePolicyStore: %w", err)
		}
		if changed {
			writes = append(writes, feeWrite)
		}
		securityWrite, changed, err := authorizePoolConfigManager(
			b, chain, securityStore, poolConfigManagerAddress,
			security_policy_store.GetAllAuthorizedCallers,
			security_policy_store.ApplyAuthorizedCallerUpdates,
			security_policy_store.AuthorizedCallerArgs{AddedCallers: []common.Address{poolConfigManagerAddress}},
		)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to authorize PoolConfigManager on SecurityPolicyStore: %w", err)
		}
		if changed {
			writes = append(writes, securityWrite)
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

		return sequences.OnChainOutput{
			Addresses: []datastore.AddressRef{connectionStoreRef, feeStoreRef, securityStoreRef, poolConfigManagerRef, poolRef},
			BatchOps:  batchOps,
		}, nil
	},
)

// maybeDeployMultiTokenPoolContract reuses an existing deployment and makes the skip visible
// to operators. The actual type/version/qualifier matching remains in MaybeDeployContract.
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
	for _, existing := range existingAddresses {
		if existing.Key() == ref.Key() && existing.Address == ref.Address {
			b.Logger.Infof("%s already deployed on chain %d at %s; skipping deployment", ref.Type, chain.Selector, ref.Address)
			break
		}
	}
	return ref, nil
}

// authorizePoolConfigManager adds the PCM to a store only when it is not already authorized.
// The read and write are forced so a retry cannot reuse a stale authorization list or an
// already-executed write report after out-of-band ownership or authorization changes.
func authorizePoolConfigManager[ARGS any](
	b cldf_ops.Bundle,
	chain evm.Chain,
	storeAddress common.Address,
	poolConfigManagerAddress common.Address,
	getAllAuthorizedCallers *cldf_ops.Operation[contract.FunctionInput[struct{}], []common.Address, evm.Chain],
	applyAuthorizedCallerUpdates *cldf_ops.Operation[contract.FunctionInput[ARGS], contract.WriteOutput, evm.Chain],
	args ARGS,
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
	return write.Output, true, nil
}
