package contract

import (
	"encoding/json"
	"errors"
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

// ExecInfo contains information about an executed transaction.
// Defined as a struct in case we want to add more fields in the future without breaking existing usage.
type ExecInfo struct {
	// Hash is the transaction hash.
	Hash string
}

// WriteOutput is the output of a write operation.
type WriteOutput struct {
	// ChainSelector is the selector of the target chain.
	ChainSelector uint64 `json:"chainSelector"`
	// Tx is the prepared transaction (in MCMS format).
	Tx mcms_types.Transaction `json:"tx"`
	// ExecInfo is populated if the write was executed, contains info about the executed transaction.
	ExecInfo *ExecInfo `json:"execInfo,omitempty"`
}

func (o WriteOutput) Executed() bool {
	return o.ExecInfo != nil
}

type WriteParams[ARGS any, C any] struct {
	// Name is the name of the operation.
	Name string
	// Version is the version of the operation.
	Version *semver.Version
	// Description is a brief description of the operation.
	Description string
	// ContractType is the type of the target contract.
	ContractType deployment.ContractType
	// ContractABI is the ABI of the target contract.
	ContractABI string
	// NewContract is a function that creates a new instance of the contract binding.
	NewContract func(address common.Address, backend bind.ContractBackend) (C, error)
	// IsAllowedCaller is a function that checks if the caller is allowed to call the function.
	IsAllowedCaller func(contract C, opts *bind.CallOpts, caller common.Address, input ARGS) (bool, error)
	// Validate is a function that validates the input arguments.
	Validate func(input ARGS) error
	// CallContract is a function that calls the desired write method on the contract.
	CallContract func(contract C, opts *bind.TransactOpts, input ARGS) (*eth_types.Transaction, error)
}

// NewWrite creates a new write operation.
// Any interfacing with gethwrappers should live in the callContract function.
// If the deployer key is an allowed caller, the transaction will be signed and sent.
// Otherwise, the MCMS transaction will be returned for alternative handling.
func NewWrite[ARGS any, C any](params WriteParams[ARGS, C]) *operations.Operation[FunctionInput[ARGS], WriteOutput, evm.Chain] {
	return operations.NewOperation(
		params.Name,
		params.Version,
		params.Description,
		func(b operations.Bundle, chain evm.Chain, input FunctionInput[ARGS]) (WriteOutput, error) {
			// BEGIN Validation
			if params.Validate != nil {
				if err := params.Validate(input.Args); err != nil {
					return WriteOutput{}, fmt.Errorf("invalid args for %s: %w", params.Name, err)
				}
			}
			if input.ChainSelector != chain.Selector {
				return WriteOutput{}, fmt.Errorf("mismatch between inputted chain selector and selector defined within dependencies: %d != %d", input.ChainSelector, chain.Selector)
			}
			if params.ContractType == "" {
				return WriteOutput{}, fmt.Errorf("contract type must be specified for %s", params.Name)
			}
			if params.ContractABI == "" {
				return WriteOutput{}, fmt.Errorf("contract ABI must be specified for %s", params.Name)
			}
			if params.NewContract == nil {
				return WriteOutput{}, fmt.Errorf("newContract function must be defined for %s", params.Name)
			}
			if params.CallContract == nil {
				return WriteOutput{}, fmt.Errorf("callContract function must be defined for %s", params.Name)
			}
			if params.IsAllowedCaller == nil {
				return WriteOutput{}, fmt.Errorf("isAllowedCaller function must be defined for %s", params.Name)
			}
			// END Validation

			boundContract, err := params.NewContract(input.Address, chain.Client)
			if err != nil {
				return WriteOutput{}, fmt.Errorf("failed to create contract instance for %s at %s on %s: %w", params.Name, input.Address, chain, err)
			}
			allowed, err := params.IsAllowedCaller(boundContract, &bind.CallOpts{Context: b.GetContext()}, chain.DeployerKey.From, input.Args)
			if err != nil {
				return WriteOutput{}, fmt.Errorf("failed to check if %s is an allowed caller of %s against %s on %s: %w", chain.DeployerKey.From, params.Name, input.Address, chain, err)
			}
			opts := deployment.SimTransactOpts()
			if allowed {
				opts = chain.DeployerKey
			}
			var execInfo *ExecInfo
			tx, callErr := params.CallContract(boundContract, opts, input.Args)
			if allowed {
				// If the call has actually been sent, we need check the call error and confirm the transaction.
				_, confirmErr := deployment.ConfirmIfNoErrorWithABI(chain, tx, params.ContractABI, callErr)
				if confirmErr != nil {
					return WriteOutput{}, fmt.Errorf("failed to confirm %s tx against %s on %s with args %+v: %w", params.Name, input.Address, chain, input.Args, confirmErr)
				}
				execInfo = &ExecInfo{Hash: tx.Hash().Hex()}
				b.Logger.Debugw(fmt.Sprintf("Confirmed %s tx against %s on %s", params.Name, input.Address, chain), "hash", tx.Hash().Hex(), "args", input.Args)
			} else if callErr != nil {
				// If we didn't execute the transaction, but there was an error preparing it, return the error.
				return WriteOutput{}, fmt.Errorf("failed to prepare %s tx against %s on %s with args %+v: %w", params.Name, input.Address, chain, input.Args, callErr)
			} else {
				b.Logger.Debugw(fmt.Sprintf("Prepared %s tx against %s on %s", params.Name, input.Address, chain), "args", input.Args)
			}

			return WriteOutput{
				ChainSelector: input.ChainSelector,
				ExecInfo:      execInfo,
				Tx: mcms_types.Transaction{
					OperationMetadata: mcms_types.OperationMetadata{
						ContractType: string(params.ContractType),
					},
					To:               input.Address.Hex(),
					Data:             tx.Data(),
					AdditionalFields: json.RawMessage(`{"value": 0}`),
				},
			}, nil
		},
	)
}

type ownableContract interface {
	Address() common.Address
	Owner(opts *bind.CallOpts) (common.Address, error)
}

func OnlyOwner[C ownableContract, ARGS any](contract C, opts *bind.CallOpts, caller common.Address, args ARGS) (bool, error) {
	owner, err := contract.Owner(opts)
	if err != nil {
		return false, fmt.Errorf("failed to get owner of %s: %w", contract.Address(), err)
	}
	return owner == caller, nil
}

func AllCallersAllowed[C any, ARGS any](contract C, opts *bind.CallOpts, caller common.Address, args ARGS) (bool, error) {
	return true, nil
}

// NewBatchOperation constructs an MCMS BatchOperation from a slice of WriteOutputs.
// It filters out any WriteOutputs that have already been executed.
// Returns an error if the WriteOutputs target multiple chains.
// If all WriteOutputs are executed, it returns an empty BatchOperation and no error.
func NewBatchOperationFromWrites(outs []WriteOutput) (mcms_types.BatchOperation, error) {
	if len(outs) == 0 {
		return mcms_types.BatchOperation{}, nil
	}

	batchOps := make(map[uint64]mcms_types.BatchOperation)
	var chainSelector uint64
	for i, out := range outs {
		if out.Executed() {
			continue // Skip executed transactions, they should not be included.
		}
		if batchOp, exists := batchOps[out.ChainSelector]; !exists {
			if i != 0 {
				return mcms_types.BatchOperation{}, errors.New("failed to make batch operation: writes target multiple chains")
			}
			batchOps[out.ChainSelector] = mcms_types.BatchOperation{
				ChainSelector: mcms_types.ChainSelector(out.ChainSelector),
				Transactions:  []mcms_types.Transaction{out.Tx},
			}
			chainSelector = out.ChainSelector
		} else {
			batchOp.Transactions = append(batchOp.Transactions, out.Tx)
			batchOps[out.ChainSelector] = batchOp
		}
	}

	// If there are no unexecuted writes, return an empty BatchOperation.
	if len(batchOps) == 0 {
		return mcms_types.BatchOperation{}, nil
	}

	return batchOps[chainSelector], nil
}
