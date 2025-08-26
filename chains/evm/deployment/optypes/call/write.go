package call

import (
	"fmt"
	"slices"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	eth_types "github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

// WriteOutput is the output of a write operation.
type WriteOutput struct {
	// ChainSelector is the selector of the target chain
	ChainSelector uint64
	// Tx is the prepared transaction (in MCMS format)
	Tx mcms_types.Transaction
	// Executed indicates whether the transaction was executed (signed and sent) or not
	Executed bool
}

// NewWrite creates a new write operation.
// Any interfacing with gethwrappers should live in the callContract function.
// If the deployer key is an allowed caller, the transaction will be signed and sent.
// Otherwise, the MCMS transaction will be returned for alternative handling.
func NewWrite[ARGS any, C any](
	name string,
	version *semver.Version,
	description string,
	contractType deployment.ContractType,
	contractABI string,
	newContract func(address common.Address, backend bind.ContractBackend) (C, error),
	allowedCallers func(contract C, opts *bind.CallOpts) ([]common.Address, error),
	validate func(input ARGS) error,
	callContract func(contract C, opts *bind.TransactOpts, input ARGS) (*eth_types.Transaction, error),
) *operations.Operation[Input[ARGS], WriteOutput, evm.Chain] {
	return operations.NewOperation(
		name,
		version,
		description,
		func(b operations.Bundle, chain evm.Chain, input Input[ARGS]) (WriteOutput, error) {
			if validate != nil {
				if err := validate(input.Args); err != nil {
					return WriteOutput{}, fmt.Errorf("invalid args for %s: %w", name, err)
				}
			}
			if input.ChainSelector != chain.Selector {
				return WriteOutput{}, fmt.Errorf("mismatch between inputted chain selector and selector defined within dependencies: %d != %d", input.ChainSelector, chain.Selector)
			}
			contract, err := newContract(input.Address, chain.Client)
			if err != nil {
				return WriteOutput{}, fmt.Errorf("failed to create contract instance for %s at %s on %s: %w", name, input.Address, chain, err)
			}
			var callers []common.Address
			if allowedCallers != nil {
				callers, err = allowedCallers(contract, &bind.CallOpts{Context: b.GetContext()})
				if err != nil {
					return WriteOutput{}, fmt.Errorf("failed to get allowed callers for %s at %s on %s: %w", name, input.Address, chain, err)
				}
			}
			opts := deployment.SimTransactOpts()
			var executed bool
			if callers != nil && slices.Contains(callers, chain.DeployerKey.From) {
				opts = chain.DeployerKey
				executed = true // Won't be returned if execution fails, so we can update it here
			}
			tx, err := callContract(contract, opts, input.Args)
			if err != nil {
				return WriteOutput{}, fmt.Errorf("failed to prepare %s tx against %s on %s", name, input.Address, chain)
			}
			b.Logger.Debugw(fmt.Sprintf("Prepared %s tx against %s on %s", name, input.Address, chain), "args", input.Args)
			if executed {
				// If the call has actually been sent, we need check the call error and confirm the transaction.
				_, err := deployment.ConfirmIfNoErrorWithABI(chain, tx, contractABI, err)
				if err != nil {
					return WriteOutput{}, fmt.Errorf("failed to confirm %s tx against %s on %s: %w", name, input.Address, chain, err)
				}
				b.Logger.Debugw(fmt.Sprintf("Confirmed %s tx against %s on %s", name, input.Address, chain), "hash", tx.Hash().Hex())
			}

			return WriteOutput{
				Executed:      executed,
				ChainSelector: input.ChainSelector,
				Tx: mcms_types.Transaction{
					OperationMetadata: mcms_types.OperationMetadata{
						ContractType: string(contractType),
					},
					To:   input.Address.Hex(),
					Data: tx.Data(),
				},
			}, err
		},
	)
}

type ownableContract interface {
	Address() common.Address
	Owner(opts *bind.CallOpts) (common.Address, error)
}

func OnlyOwner[C ownableContract](contract C, opts *bind.CallOpts) ([]common.Address, error) {
	owner, err := contract.Owner(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to get owner of %s: %w", contract.Address(), err)
	}
	return []common.Address{owner}, nil
}
