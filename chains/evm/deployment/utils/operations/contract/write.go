package contract

import (
	"fmt"

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
	ChainSelector uint64 `json:"chainSelector"`
	// Tx is the prepared transaction (in MCMS format)
	Tx mcms_types.Transaction `json:"tx"`
	// TimelockAddress is the address of the timelock contract, if applicable for MCMS proposals
	TimelockAddress common.Address `json:"timelockAddress,omitempty"`
	// MCMAddress is the address of the MCMS contract, if applicable for MCMS proposals
	MCMAddress common.Address `json:"mcmAddress,omitempty"`
	// Executed indicates whether the transaction was executed (signed and sent) or not
	Executed bool `json:"executed"`
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
	isAllowedCaller func(contract C, opts *bind.CallOpts, caller common.Address) (bool, error),
	validate func(input ARGS) error,
	callContract func(contract C, opts *bind.TransactOpts, input ARGS) (*eth_types.Transaction, error),
) *operations.Operation[FunctionInput[ARGS], WriteOutput, evm.Chain] {
	return operations.NewOperation(
		name,
		version,
		description,
		func(b operations.Bundle, chain evm.Chain, input FunctionInput[ARGS]) (WriteOutput, error) {
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
			var allowed bool
			if isAllowedCaller != nil {
				allowed, err = isAllowedCaller(contract, &bind.CallOpts{Context: b.GetContext()}, chain.DeployerKey.From)
				if err != nil {
					return WriteOutput{}, fmt.Errorf("failed to check if %s is an allowed caller of %s against %s on %s: %w", chain.DeployerKey.From, name, input.Address, chain, err)
				}
			}
			opts := deployment.SimTransactOpts()
			if allowed {
				opts = chain.DeployerKey
			}
			tx, callErr := callContract(contract, opts, input.Args)
			if allowed {
				// If the call has actually been sent, we need check the call error and confirm the transaction.
				_, confirmErr := deployment.ConfirmIfNoErrorWithABI(chain, tx, contractABI, callErr)
				if confirmErr != nil {
					return WriteOutput{}, fmt.Errorf("failed to confirm %s tx against %s on %s with args %+v: %w", name, input.Address, chain, input.Args, confirmErr)
				}
				b.Logger.Debugw(fmt.Sprintf("Confirmed %s tx against %s on %s", name, input.Address, chain), "hash", tx.Hash().Hex(), "args", input.Args)
			} else if callErr != nil {
				// If we didn't execute the transaction, but there was an error preparing it, return the error.
				return WriteOutput{}, fmt.Errorf("failed to prepare %s tx against %s on %s with args %+v: %w", name, input.Address, chain, input.Args, callErr)
			} else {
				b.Logger.Debugw(fmt.Sprintf("Prepared %s tx against %s on %s", name, input.Address, chain), "args", input.Args)
			}

			return WriteOutput{
				Executed:      allowed,
				ChainSelector: input.ChainSelector,
				Tx: mcms_types.Transaction{
					OperationMetadata: mcms_types.OperationMetadata{
						ContractType: string(contractType),
					},
					To:               input.Address.Hex(),
					Data:             tx.Data(),
					AdditionalFields: []byte{0x7B, 0x7D}, // "{}" in bytes
				},
			}, nil
		},
	)
}

type ownableContract interface {
	Address() common.Address
	Owner(opts *bind.CallOpts) (common.Address, error)
}

func OnlyOwner[C ownableContract](contract C, opts *bind.CallOpts, caller common.Address) (bool, error) {
	owner, err := contract.Owner(opts)
	if err != nil {
		return false, fmt.Errorf("failed to get owner of %s: %w", contract.Address(), err)
	}
	return owner == caller, nil
}
