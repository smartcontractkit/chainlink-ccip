package sequences

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	cldf_evm "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/smartcontractkit/chainlink-evm/gethwrappers/shared/generated/initial/burn_mint_erc677"
	"github.com/smartcontractkit/mcms/sdk/evm/bindings"
	"github.com/smartcontractkit/mcms/types"

	ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations"
)

// OwnableContractsInput describes one or more ownable contracts whose ownership should be updated.
type OwnableContractsInput struct {
	ChainSelector uint64
	Contracts     []ops.OpTransferOwnershipInput
}

// SeqTransferMCMOwnershipToTimelockInput is kept for backward compatibility.
type SeqTransferMCMOwnershipToTimelockInput = OwnableContractsInput

type ownershipStep uint8

const (
	ownershipStepTransfer ownershipStep = 1 << iota
	ownershipStepAccept
)

// OwnershipInputsFromRefs builds ownership inputs for transferring contract ownership to a timelock.
func OwnershipInputsFromRefs(
	chainSelector uint64,
	timelock common.Address,
	refs []datastore.AddressRef,
) []ops.OpTransferOwnershipInput {
	inputs := make([]ops.OpTransferOwnershipInput, 0, len(refs))
	for _, ref := range refs {
		inputs = append(inputs, ops.OpTransferOwnershipInput{
			ChainSelector:   chainSelector,
			Address:         common.HexToAddress(ref.Address),
			ProposedOwner:   timelock,
			ContractType:    cldf.ContractType(ref.Type),
			TimelockAddress: timelock,
		})
	}
	return inputs
}

// TransferOwnership transfers Ownable2Step ownership of the given contracts.
func TransferOwnership(
	b cldf_ops.Bundle,
	chain cldf_evm.Chain,
	contracts []ops.OpTransferOwnershipInput,
) ([]types.BatchOperation, error) {
	return executeOwnershipSteps(b, chain, contracts, ownershipStepTransfer)
}

// AcceptOwnership accepts Ownable2Step ownership of the given contracts.
func AcceptOwnership(
	b cldf_ops.Bundle,
	chain cldf_evm.Chain,
	contracts []ops.OpTransferOwnershipInput,
) ([]types.BatchOperation, error) {
	return executeOwnershipSteps(b, chain, contracts, ownershipStepAccept)
}

// TransferAndAcceptOwnership transfers then accepts ownership for the given contracts.
// Transfer is executed on-chain when the deployer is the current owner; accept is returned
// as MCMS batch operations for the timelock when the proposed owner is the timelock.
func TransferAndAcceptOwnership(
	b cldf_ops.Bundle,
	chain cldf_evm.Chain,
	contracts []ops.OpTransferOwnershipInput,
) ([]types.BatchOperation, error) {
	return executeOwnershipSteps(b, chain, contracts, ownershipStepTransfer|ownershipStepAccept)
}

func executeOwnershipSteps(
	b cldf_ops.Bundle,
	chain cldf_evm.Chain,
	contracts []ops.OpTransferOwnershipInput,
	steps ownershipStep,
) ([]types.BatchOperation, error) {
	var batchOps []types.BatchOperation
	for _, contractInput := range contracts {
		contractAddr := contractInput.Address.Hex()
		owner, ownable, err := LoadOwnableContract(contractInput.Address, chain.Client)
		if err != nil {
			return nil, fmt.Errorf("failed to load ownable contract %s: %w", contractAddr, err)
		}
		if owner == contractInput.ProposedOwner {
			b.Logger.Infof("ownership of contract %s (%s) on chain %s is already set to %s, skipping",
				contractAddr, contractInput.ContractType, chain.Name(), contractInput.ProposedOwner.Hex())
			continue
		}

		deps := ops.OpEVMOwnershipDeps{Chain: chain, OwnableC: ownable}
		var writes []contract.WriteOutput

		if steps&ownershipStepTransfer != 0 {
			report, err := cldf_ops.ExecuteOperation(b, ops.OpTransferOwnership, deps, contractInput)
			if err != nil {
				b.Logger.Errorf("failed to transfer ownership of contract %s on chain %s: %v", contractAddr, chain.Name(), err)
				return nil, fmt.Errorf("error transferring ownership of contract %s on chain %s: %w", contractAddr, chain.Name(), err)
			}
			writes = append(writes, report.Output)
		}
		if steps&ownershipStepAccept != 0 {
			report, err := cldf_ops.ExecuteOperation(b, ops.OpAcceptOwnership, deps, contractInput)
			if err != nil {
				b.Logger.Errorf("failed to accept ownership of contract %s on chain %s: %v", contractAddr, chain.Name(), err)
				return nil, fmt.Errorf("error accepting ownership of contract %s on chain %s: %w", contractAddr, chain.Name(), err)
			}
			writes = append(writes, report.Output)
		}
		if len(writes) == 0 {
			continue
		}
		batch, err := contract.NewBatchOperationFromWrites(writes)
		if err != nil {
			return nil, fmt.Errorf("failed to create batch operation from ownership writes: %w", err)
		}
		batchOps = append(batchOps, batch)
	}
	return batchOps, nil
}

func LoadOwnableContract(addr common.Address, client bind.ContractBackend) (common.Address, ops.OwnershipTranferable, error) {
	c, err := burn_mint_erc677.NewBurnMintERC677(addr, client)
	if err != nil {
		return common.Address{}, nil, fmt.Errorf("failed to create contract: %w", err)
	}
	owner, err := c.Owner(nil)
	if err != nil {
		return common.Address{}, nil, fmt.Errorf("failed to get owner of contract %s: %w", c.Address(), err)
	}

	return owner, c, nil
}

func LoadTimelockContract(addr common.Address, client bind.ContractBackend) (*bindings.RBACTimelock, error) {
	c, err := bindings.NewRBACTimelock(addr, client)
	if err != nil {
		return nil, fmt.Errorf("failed to load timelock contract: %w", err)
	}

	return c, nil
}
