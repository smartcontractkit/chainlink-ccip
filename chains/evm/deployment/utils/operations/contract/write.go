package contract

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	eth_types "github.com/ethereum/go-ethereum/core/types"
	mcms_types "github.com/smartcontractkit/mcms/types"

	cldf_evm "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	upstream "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

// ExecInfo and WriteOutput match the framework contract package for interoperability.
type (
	ExecInfo    = upstream.ExecInfo
	WriteOutput = upstream.WriteOutput
)

// WriteParams contains parameters to create a write operation. Either set Contract
// to a bound instance, or set NewContract to bind at execution time using
// FunctionInput.Address and the chain client.
type WriteParams[ARGS any, C interface{ Address() common.Address }] struct {
	Name            string
	Version         *semver.Version
	Description     string
	ContractType    deployment.ContractType
	ContractABI     string
	Contract        C
	NewContract     func(common.Address, bind.ContractBackend) (C, error)
	IsAllowedCaller func(contract C, opts *bind.CallOpts, caller common.Address, input ARGS) (bool, error)
	Validate        func(input ARGS) error
	CallContract    func(contract C, opts *bind.TransactOpts, input ARGS) (*eth_types.Transaction, error)
}

// NewWrite creates a write operation. When NewContract is set, the contract is
// instantiated from FunctionInput.Address at execution time.
func NewWrite[ARGS any, C interface{ Address() common.Address }](params WriteParams[ARGS, C]) *operations.Operation[FunctionInput[ARGS], WriteOutput, cldf_evm.Chain] {
	return operations.NewOperation(
		params.Name,
		params.Version,
		params.Description,
		func(b operations.Bundle, chain cldf_evm.Chain, input FunctionInput[ARGS]) (WriteOutput, error) {
			if params.Validate != nil {
				if err := params.Validate(input.Args); err != nil {
					return WriteOutput{}, fmt.Errorf("invalid args for %s: %w", params.Name, err)
				}
			}
			if params.ContractType == "" {
				return WriteOutput{}, fmt.Errorf("contract type must be specified for %s", params.Name)
			}
			if params.ContractABI == "" {
				return WriteOutput{}, fmt.Errorf("contract ABI must be specified for %s", params.Name)
			}
			if params.CallContract == nil {
				return WriteOutput{}, fmt.Errorf("callContract function must be defined for %s", params.Name)
			}
			if params.IsAllowedCaller == nil {
				return WriteOutput{}, fmt.Errorf("isAllowedCaller function must be defined for %s", params.Name)
			}

			c, err := resolveContract(params.Contract, params.NewContract, input.Address, chain)
			if err != nil {
				return WriteOutput{}, err
			}
			if input.ChainSelector != 0 && input.ChainSelector != chain.Selector {
				return WriteOutput{}, fmt.Errorf("chain selector mismatch for %s: input %d vs chain %d", params.Name, input.ChainSelector, chain.Selector)
			}

			allowed, err := params.IsAllowedCaller(c, &bind.CallOpts{Context: b.GetContext()}, chain.DeployerKey.From, input.Args)
			if err != nil {
				return WriteOutput{}, fmt.Errorf("failed to check if %s is an allowed caller of %s against %s on %s: %w", chain.DeployerKey.From, params.Name, c.Address(), chain, err)
			}
			opts := deployment.SimTransactOpts()
			if allowed {
				opts = chain.DeployerKey
			}
			var execInfo *upstream.ExecInfo
			tx, callErr := params.CallContract(c, opts, input.Args)
			if callErr == nil && tx == nil {
				return WriteOutput{}, fmt.Errorf("contract call returned nil transaction for %s against %s on %s", params.Name, c.Address(), chain)
			}
			if allowed {
				_, confirmErr := deployment.ConfirmIfNoErrorWithABI(chain, tx, params.ContractABI, callErr)
				if confirmErr != nil {
					return WriteOutput{}, fmt.Errorf("failed to confirm %s tx against %s on %s with args %+v: %w", params.Name, c.Address(), chain, input.Args, confirmErr)
				}
				execInfo = &upstream.ExecInfo{Hash: tx.Hash().Hex()}
				b.Logger.Debugw(fmt.Sprintf("Confirmed %s tx against %s on %s", params.Name, c.Address(), chain), "hash", tx.Hash().Hex(), "args", input.Args)
			} else if callErr != nil {
				return WriteOutput{}, fmt.Errorf("failed to prepare %s tx against %s on %s with args %+v: %w", params.Name, c.Address(), chain, input.Args, callErr)
			} else {
				b.Logger.Debugw(fmt.Sprintf("Prepared %s tx against %s on %s", params.Name, c.Address(), chain), "args", input.Args)
			}

			return WriteOutput{
				ChainSelector: chain.Selector,
				ExecInfo:      execInfo,
				Tx: mcms_types.Transaction{
					OperationMetadata: mcms_types.OperationMetadata{
						ContractType: string(params.ContractType),
					},
					To:               c.Address().Hex(),
					Data:             tx.Data(),
					AdditionalFields: json.RawMessage(`{"value": 0}`),
				},
			}, nil
		},
	)
}

func resolveContract[C interface{ Address() common.Address }](
	static C,
	newContract func(common.Address, bind.ContractBackend) (C, error),
	addr common.Address,
	chain cldf_evm.Chain,
) (C, error) {
	var zero C
	if newContract != nil {
		if addr == (common.Address{}) {
			return zero, fmt.Errorf("contract address is required in operation input")
		}
		return newContract(addr, chain.Client)
	}
	return static, nil
}

type ownableContract interface {
	Address() common.Address
	Owner(opts *bind.CallOpts) (common.Address, error)
}

// RetryContractCall retries contract read calls that can briefly fail after deployment.
func RetryContractCall[T any](
	opts *bind.CallOpts,
	waitLabel string,
	failureLabel string,
	contractAddress common.Address,
	check func() (T, error),
) (T, error) {
	const (
		timeout    = 5 * time.Second
		retryDelay = 500 * time.Millisecond
	)

	ctx := context.Background()
	if opts != nil && opts.Context != nil {
		ctx = opts.Context
	}

	deadline := time.Now().Add(timeout)
	var lastErr error
	var zero T

	for time.Now().Before(deadline) {
		result, err := check()
		if err == nil {
			return result, nil
		}

		if strings.Contains(err.Error(), "empty string") || strings.Contains(err.Error(), "no contract code") {
			lastErr = err
			select {
			case <-ctx.Done():
				return zero, fmt.Errorf("context cancelled while waiting for %s of %s: %w", waitLabel, contractAddress, ctx.Err())
			case <-time.After(retryDelay):
			}
			continue
		}

		return zero, fmt.Errorf("failed to %s of %s: %w", failureLabel, contractAddress, err)
	}

	return zero, fmt.Errorf("failed to %s of %s after %v: %w", failureLabel, contractAddress, timeout, lastErr)
}

func OnlyOwner[C ownableContract, ARGS any](contract C, opts *bind.CallOpts, caller common.Address, args ARGS) (bool, error) {
	owner, err := RetryContractCall(opts, "owner", "get owner", contract.Address(), func() (common.Address, error) {
		return contract.Owner(opts)
	})
	if err != nil {
		return false, err
	}
	return owner == caller, nil
}

type AccessControlContract interface {
	Address() common.Address
	HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error)
}

func HasRole[C AccessControlContract](
	contract C,
	opts *bind.CallOpts,
	role [32]byte,
	account common.Address,
) (bool, error) {
	return RetryContractCall(opts, "role check", "check role", contract.Address(), func() (bool, error) {
		return contract.HasRole(opts, role, account)
	})
}

func AllCallersAllowed[C any, ARGS any](contract C, opts *bind.CallOpts, caller common.Address, args ARGS) (bool, error) {
	return true, nil
}

func NoCallersAllowed[C any, ARGS any](contract C, opts *bind.CallOpts, caller common.Address, args ARGS) (bool, error) {
	return false, nil
}

func NewBatchOperationFromWrites(outs []WriteOutput) (mcms_types.BatchOperation, error) {
	if len(outs) == 0 {
		return mcms_types.BatchOperation{}, nil
	}

	var (
		chainSelector uint64
		txs           []mcms_types.Transaction
	)
	for _, out := range outs {
		if out.Executed() {
			continue
		}
		if len(txs) == 0 {
			chainSelector = out.ChainSelector
			txs = append(txs, out.Tx)
			continue
		}
		if out.ChainSelector != chainSelector {
			return mcms_types.BatchOperation{}, errors.New("failed to make batch operation: writes target multiple chains")
		}
		txs = append(txs, out.Tx)
	}

	if len(txs) == 0 {
		return mcms_types.BatchOperation{}, nil
	}

	return mcms_types.BatchOperation{
		ChainSelector: mcms_types.ChainSelector(chainSelector),
		Transactions:  txs,
	}, nil
}
