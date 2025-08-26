package call

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

// NewWrite creates a new write operation.
// Any interfacing with gethwrappers should live in the callContract function.
// This operation does not actually send the transaction, leaving execution method up to the caller (i.e. deployer key, MCMS)
// The return type is an MCMS transaction, which is a chain-agnostic representation of an executable.
func NewWrite[ARGS any, C any](
	name string,
	version *semver.Version,
	description string,
	contractType deployment.ContractType,
	newContract func(address common.Address, backend bind.ContractBackend) (C, error),
	callContract func(contract C, opts *bind.TransactOpts, input ARGS) (*eth_types.Transaction, error),
) *operations.Operation[Input[ARGS], mcms_types.Transaction, evm.Chain] {
	return operations.NewOperation(
		name,
		version,
		description,
		func(b operations.Bundle, chain evm.Chain, input Input[ARGS]) (mcms_types.Transaction, error) {
			if input.ChainSelector != chain.Selector {
				return mcms_types.Transaction{}, fmt.Errorf("mismatch between inputted chain selector and selector defined within dependencies: %d != %d", input.ChainSelector, chain.Selector)
			}
			contract, err := newContract(input.Address, chain.Client)
			if err != nil {
				return mcms_types.Transaction{}, fmt.Errorf("failed to create contract instance for %s at %s on %s: %w", name, input.Address, chain, err)
			}
			tx, err := callContract(contract, deployment.SimTransactOpts(), input.Args)
			if err != nil {
				return mcms_types.Transaction{}, fmt.Errorf("failed to prepare %s tx against %s on %s", name, input.Address, chain)
			}
			b.Logger.Debugw(fmt.Sprintf("Prepared %s tx against %s on %s", name, input.Address, chain), "args", input.Args)

			return mcms_types.Transaction{
				OperationMetadata: mcms_types.OperationMetadata{
					ContractType: string(contractType),
				},
				To:   input.Address.Hex(),
				Data: tx.Data(),
			}, err
		},
	)
}
