package contract

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	cldf_evm "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

// NewRead creates a new read operation.
// Any interfacing with gethwrappers should live in the callContract function.
func NewRead[ARGS any, RET any, C any](
	name string,
	version *semver.Version,
	description string,
	contractType cldf_deployment.ContractType,
	newContract func(address common.Address, backend bind.ContractBackend) (C, error),
	callContract func(contract C, opts *bind.CallOpts, input ARGS) (RET, error),
) *operations.Operation[FunctionInput[ARGS], RET, cldf_evm.Chain] {
	return operations.NewOperation(
		name,
		version,
		description,
		func(b operations.Bundle, chain cldf_evm.Chain, input FunctionInput[ARGS]) (RET, error) {
			var empty RET
			if input.ChainSelector != chain.Selector {
				return empty, fmt.Errorf("mismatch between inputted chain selector and selector defined within dependencies: %d != %d", input.ChainSelector, chain.Selector)
			}
			contract, err := newContract(input.Address, chain.Client)
			if err != nil {
				return empty, fmt.Errorf("failed to create contract instance for %s at %s on %s: %w", name, input.Address, chain, err)
			}
			return callContract(contract, &bind.CallOpts{Context: b.GetContext()}, input.Args)
		},
	)
}
