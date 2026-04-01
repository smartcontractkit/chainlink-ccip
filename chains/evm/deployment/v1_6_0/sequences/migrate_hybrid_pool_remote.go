package sequences

import (
	"bytes"
	"fmt"
	"math/big"
	"slices"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	hybrid_pool_binding "github.com/smartcontractkit/ccip-contract-examples/chains/evm/gobindings/generated/latest/hybrid_with_external_minter_token_pool"
	evm_contract "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	tar_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
	hybrid_pool_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/hybrid_with_external_minter_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_0/token_admin_registry"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

type MigrateHybridPoolRemoteInput struct {
	HubChainSelector     uint64
	HubPoolAddress       common.Address
	RemoteChainSelector  uint64
	NewRemotePoolAddress common.Address
	OldRemotePoolAddress common.Address
	RemoteChainSupply    *big.Int
	TargetGroup          uint8
	RemoteTARAddress     common.Address
	RemoteTokenAddress   common.Address
}

// Migration-specific operation wrappers with NoCallersAllowed to guarantee all writes
// are collected for a single atomic MCMS proposal rather than executed directly.

var migrationAddRemotePool = evm_contract.NewWrite(evm_contract.WriteParams[hybrid_pool_ops.AddRemotePoolArgs, *hybrid_pool_binding.HybridWithExternalMinterTokenPool]{
	Name:            "migration:hybrid-pool:add-remote-pool",
	Version:         hybrid_pool_ops.Version,
	Description:     "Proposal-only addRemotePool for token-pool migration",
	ContractType:    hybrid_pool_ops.ContractType,
	ContractABI:     hybrid_pool_binding.HybridWithExternalMinterTokenPoolABI,
	NewContract:     hybrid_pool_binding.NewHybridWithExternalMinterTokenPool,
	IsAllowedCaller: evm_contract.NoCallersAllowed[*hybrid_pool_binding.HybridWithExternalMinterTokenPool, hybrid_pool_ops.AddRemotePoolArgs],
	Validate:        func(hybrid_pool_ops.AddRemotePoolArgs) error { return nil },
	CallContract: func(c *hybrid_pool_binding.HybridWithExternalMinterTokenPool, opts *bind.TransactOpts, args hybrid_pool_ops.AddRemotePoolArgs) (*types.Transaction, error) {
		return c.AddRemotePool(opts, args.RemoteChainSelector, args.RemotePoolAddress)
	},
})

var migrationRemoveRemotePool = evm_contract.NewWrite(evm_contract.WriteParams[hybrid_pool_ops.RemoveRemotePoolArgs, *hybrid_pool_binding.HybridWithExternalMinterTokenPool]{
	Name:            "migration:hybrid-pool:remove-remote-pool",
	Version:         hybrid_pool_ops.Version,
	Description:     "Proposal-only removeRemotePool for token-pool migration",
	ContractType:    hybrid_pool_ops.ContractType,
	ContractABI:     hybrid_pool_binding.HybridWithExternalMinterTokenPoolABI,
	NewContract:     hybrid_pool_binding.NewHybridWithExternalMinterTokenPool,
	IsAllowedCaller: evm_contract.NoCallersAllowed[*hybrid_pool_binding.HybridWithExternalMinterTokenPool, hybrid_pool_ops.RemoveRemotePoolArgs],
	Validate:        func(hybrid_pool_ops.RemoveRemotePoolArgs) error { return nil },
	CallContract: func(c *hybrid_pool_binding.HybridWithExternalMinterTokenPool, opts *bind.TransactOpts, args hybrid_pool_ops.RemoveRemotePoolArgs) (*types.Transaction, error) {
		return c.RemoveRemotePool(opts, args.RemoteChainSelector, args.RemotePoolAddress)
	},
})

var migrationUpdateGroups = evm_contract.NewWrite(evm_contract.WriteParams[hybrid_pool_ops.UpdateGroupsArgs, *hybrid_pool_binding.HybridWithExternalMinterTokenPool]{
	Name:            "migration:hybrid-pool:update-groups",
	Version:         hybrid_pool_ops.Version,
	Description:     "Proposal-only updateGroups for token-pool migration",
	ContractType:    hybrid_pool_ops.ContractType,
	ContractABI:     hybrid_pool_binding.HybridWithExternalMinterTokenPoolABI,
	NewContract:     hybrid_pool_binding.NewHybridWithExternalMinterTokenPool,
	IsAllowedCaller: evm_contract.NoCallersAllowed[*hybrid_pool_binding.HybridWithExternalMinterTokenPool, hybrid_pool_ops.UpdateGroupsArgs],
	Validate:        func(hybrid_pool_ops.UpdateGroupsArgs) error { return nil },
	CallContract: func(c *hybrid_pool_binding.HybridWithExternalMinterTokenPool, opts *bind.TransactOpts, args hybrid_pool_ops.UpdateGroupsArgs) (*types.Transaction, error) {
		return c.UpdateGroups(opts, args.GroupUpdates)
	},
})

var migrationSetPool = evm_contract.NewWrite(evm_contract.WriteParams[tar_ops.SetPoolArgs, *token_admin_registry.TokenAdminRegistry]{
	Name:            "migration:token-admin-registry:set-pool",
	Version:         tar_ops.Version,
	Description:     "Proposal-only setPool for token-pool migration",
	ContractType:    tar_ops.ContractType,
	ContractABI:     token_admin_registry.TokenAdminRegistryABI,
	NewContract:     token_admin_registry.NewTokenAdminRegistry,
	IsAllowedCaller: evm_contract.NoCallersAllowed[*token_admin_registry.TokenAdminRegistry, tar_ops.SetPoolArgs],
	Validate:        func(tar_ops.SetPoolArgs) error { return nil },
	CallContract: func(c *token_admin_registry.TokenAdminRegistry, opts *bind.TransactOpts, args tar_ops.SetPoolArgs) (*types.Transaction, error) {
		return c.SetPool(opts, args.TokenAddress, args.TokenPoolAddress)
	},
})

var MigrateHybridPoolRemote = cldf_ops.NewSequence(
	"migrate-hybrid-pool-remote",
	semver.MustParse("1.6.0"),
	"Migrates a remote chain token pool from lock-release to burn-mint on a hybrid hub pool",
	func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input MigrateHybridPoolRemoteInput) (sequences.OnChainOutput, error) {
		if input.RemoteChainSupply == nil {
			return sequences.OnChainOutput{}, fmt.Errorf("RemoteChainSupply must not be nil")
		}

		hubChain, ok := chains.EVMChains()[input.HubChainSelector]
		if !ok {
			return sequences.OnChainOutput{}, fmt.Errorf("hub chain with selector %d not defined", input.HubChainSelector)
		}
		remoteChain, ok := chains.EVMChains()[input.RemoteChainSelector]
		if !ok {
			return sequences.OnChainOutput{}, fmt.Errorf("remote chain with selector %d not defined", input.RemoteChainSelector)
		}

		oldPoolBytes := common.LeftPadBytes(input.OldRemotePoolAddress.Bytes(), 32)
		newPoolBytes := common.LeftPadBytes(input.NewRemotePoolAddress.Bytes(), 32)

		hubWrites := make([]evm_contract.WriteOutput, 0, 3)

		remotePoolsReport, err := cldf_ops.ExecuteOperation(b, hybrid_pool_ops.GetRemotePools, hubChain, evm_contract.FunctionInput[uint64]{
			ChainSelector: input.HubChainSelector,
			Address:       input.HubPoolAddress,
			Args:          input.RemoteChainSelector,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to read hub remote pools for remote chain %d: %w", input.RemoteChainSelector, err)
		}
		for _, pool := range remotePoolsReport.Output {
			if !bytes.Equal(pool, oldPoolBytes) && !bytes.Equal(pool, newPoolBytes) {
				return sequences.OnChainOutput{}, fmt.Errorf(
					"unexpected pool %x in remote pool set for chain %d",
					pool,
					input.RemoteChainSelector,
				)
			}
		}

		oldPresent := containsBytes(remotePoolsReport.Output, oldPoolBytes)
		newPresent := containsBytes(remotePoolsReport.Output, newPoolBytes)
		if !oldPresent && !newPresent {
			return sequences.OnChainOutput{}, fmt.Errorf(
				"neither old pool %s nor new pool %s registered for chain %d",
				input.OldRemotePoolAddress,
				input.NewRemotePoolAddress,
				input.RemoteChainSelector,
			)
		}

		if !newPresent {
			addReport, err := cldf_ops.ExecuteOperation(b, migrationAddRemotePool, hubChain, evm_contract.FunctionInput[hybrid_pool_ops.AddRemotePoolArgs]{
				ChainSelector: input.HubChainSelector,
				Address:       input.HubPoolAddress,
				Args: hybrid_pool_ops.AddRemotePoolArgs{
					RemoteChainSelector: input.RemoteChainSelector,
					RemotePoolAddress:   newPoolBytes,
				},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to add new remote pool %s on hub chain %d: %w", input.NewRemotePoolAddress, input.HubChainSelector, err)
			}
			hubWrites = append(hubWrites, addReport.Output)
		}

		if oldPresent {
			removeReport, err := cldf_ops.ExecuteOperation(b, migrationRemoveRemotePool, hubChain, evm_contract.FunctionInput[hybrid_pool_ops.RemoveRemotePoolArgs]{
				ChainSelector: input.HubChainSelector,
				Address:       input.HubPoolAddress,
				Args: hybrid_pool_ops.RemoveRemotePoolArgs{
					RemoteChainSelector: input.RemoteChainSelector,
					RemotePoolAddress:   oldPoolBytes,
				},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to remove old remote pool %s on hub chain %d: %w", input.OldRemotePoolAddress, input.HubChainSelector, err)
			}
			hubWrites = append(hubWrites, removeReport.Output)
		}

		groupReport, err := cldf_ops.ExecuteOperation(b, hybrid_pool_ops.GetGroup, hubChain, evm_contract.FunctionInput[uint64]{
			ChainSelector: input.HubChainSelector,
			Address:       input.HubPoolAddress,
			Args:          input.RemoteChainSelector,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to read hub group for remote chain %d: %w", input.RemoteChainSelector, err)
		}
		if groupReport.Output != input.TargetGroup {
			updateReport, err := cldf_ops.ExecuteOperation(b, migrationUpdateGroups, hubChain, evm_contract.FunctionInput[hybrid_pool_ops.UpdateGroupsArgs]{
				ChainSelector: input.HubChainSelector,
				Address:       input.HubPoolAddress,
				Args: hybrid_pool_ops.UpdateGroupsArgs{
					GroupUpdates: []hybrid_pool_binding.HybridTokenPoolAbstractGroupUpdate{
						{
							RemoteChainSelector: input.RemoteChainSelector,
							Group:               input.TargetGroup,
							RemoteChainSupply:   input.RemoteChainSupply,
						},
					},
				},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to update group for remote chain %d on hub chain %d: %w", input.RemoteChainSelector, input.HubChainSelector, err)
			}
			hubWrites = append(hubWrites, updateReport.Output)
		}

		remoteWrites := make([]evm_contract.WriteOutput, 0, 1)

		tarConfigReport, err := cldf_ops.ExecuteOperation(b, tar_ops.GetTokenConfig, remoteChain, evm_contract.FunctionInput[common.Address]{
			ChainSelector: input.RemoteChainSelector,
			Address:       input.RemoteTARAddress,
			Args:          input.RemoteTokenAddress,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to read token config from TAR %s on chain %d: %w", input.RemoteTARAddress, input.RemoteChainSelector, err)
		}

		currentPool := tarConfigReport.Output.TokenPool
		if currentPool == (common.Address{}) {
			return sequences.OnChainOutput{}, fmt.Errorf("token %s has no pool set in TAR on chain %d", input.RemoteTokenAddress, input.RemoteChainSelector)
		}
		if currentPool != input.NewRemotePoolAddress {
			if currentPool != input.OldRemotePoolAddress {
				return sequences.OnChainOutput{}, fmt.Errorf(
					"TAR pool %s is neither old %s nor new %s for token %s on chain %d",
					currentPool,
					input.OldRemotePoolAddress,
					input.NewRemotePoolAddress,
					input.RemoteTokenAddress,
					input.RemoteChainSelector,
				)
			}

			setPoolReport, err := cldf_ops.ExecuteOperation(b, migrationSetPool, remoteChain, evm_contract.FunctionInput[tar_ops.SetPoolArgs]{
				ChainSelector: input.RemoteChainSelector,
				Address:       input.RemoteTARAddress,
				Args: tar_ops.SetPoolArgs{
					TokenAddress:     input.RemoteTokenAddress,
					TokenPoolAddress: input.NewRemotePoolAddress,
				},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to set TAR pool for token %s on chain %d: %w", input.RemoteTokenAddress, input.RemoteChainSelector, err)
			}
			remoteWrites = append(remoteWrites, setPoolReport.Output)
		}

		batchOps := make([]mcms_types.BatchOperation, 0, 2)
		if len(hubWrites) > 0 {
			hubBatch, err := evm_contract.NewBatchOperationFromWrites(hubWrites)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to build hub chain batch operation: %w", err)
			}
			if len(hubBatch.Transactions) > 0 {
				batchOps = append(batchOps, hubBatch)
			}
		}
		if len(remoteWrites) > 0 {
			remoteBatch, err := evm_contract.NewBatchOperationFromWrites(remoteWrites)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to build remote chain batch operation: %w", err)
			}
			if len(remoteBatch.Transactions) > 0 {
				batchOps = append(batchOps, remoteBatch)
			}
		}

		return sequences.OnChainOutput{BatchOps: batchOps}, nil
	},
)

func containsBytes(haystack [][]byte, needle []byte) bool {
	return slices.ContainsFunc(haystack, func(candidate []byte) bool {
		return bytes.Equal(candidate, needle)
	})
}
