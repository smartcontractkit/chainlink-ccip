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

type ReadParams[ARGS any, RET any, C any] struct {
	// Name is the name of the operation.
	Name string
	// Version is the version of the operation.
	Version *semver.Version
	// Description is a brief description of the operation.
	Description string
	// ContractType is the type of contract to read from.
	ContractType cldf_deployment.ContractType
	// NewContract is a function that creates a new instance of the contract binding.
	NewContract func(address common.Address, backend bind.ContractBackend) (C, error)
	// CallContract is a function that calls the desired read function on the contract.
	CallContract func(contract C, opts *bind.CallOpts, input ARGS) (RET, error)
}

// NewRead creates a new read operation.
// Any interfacing with gethwrappers should live in the callContract function.
func NewRead[ARGS any, RET any, C any](params ReadParams[ARGS, RET, C]) *operations.Operation[FunctionInput[ARGS], RET, cldf_evm.Chain] {
	return operations.NewOperation(
		params.Name,
		params.Version,
		params.Description,
		func(b operations.Bundle, chain cldf_evm.Chain, input FunctionInput[ARGS]) (RET, error) {
			var empty RET
			// BEGIN Validation
			if input.ChainSelector != chain.Selector {
				return empty, fmt.Errorf("mismatch between inputted chain selector and selector defined within dependencies: %d != %d", input.ChainSelector, chain.Selector)
			}
			if params.ContractType == "" {
				return empty, fmt.Errorf("contract type must be specified for %s", params.Name)
			}
			if input.Address == (common.Address{}) {
				return empty, fmt.Errorf("address must be specified for %s", params.Name)
			}
			if params.NewContract == nil {
				return empty, fmt.Errorf("newContract function must be defined for %s", params.Name)
			}
			if params.CallContract == nil {
				return empty, fmt.Errorf("callContract function must be defined for %s", params.Name)
			}
			// END Validation

			contract, err := params.NewContract(input.Address, chain.Client)
			if err != nil {
				return empty, fmt.Errorf("failed to create contract instance for %s at %s on %s: %w", params.Name, input.Address, chain, err)
			}
			return params.CallContract(contract, &bind.CallOpts{Context: b.GetContext()}, input.Args)
		},
	)
}
