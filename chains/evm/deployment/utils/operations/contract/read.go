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

// ReadParams contains parameters to create a read operation. Either set Contract
// to a bound instance, or set NewContract to bind at execution time using
// FunctionInput.Address and the chain client.
type ReadParams[ARGS any, RET any, C interface{ Address() common.Address }] struct {
	Name         string
	Version      *semver.Version
	Description  string
	ContractType cldf_deployment.ContractType
	Contract     C
	NewContract  func(common.Address, bind.ContractBackend) (C, error)
	CallContract func(contract C, opts *bind.CallOpts, input ARGS) (RET, error)
}

// NewRead creates a read operation that resolves the contract from input.Address
// when NewContract is set; otherwise it uses Contract.
func NewRead[ARGS any, RET any, C interface{ Address() common.Address }](params ReadParams[ARGS, RET, C]) *operations.Operation[FunctionInput[ARGS], RET, cldf_evm.Chain] {
	return operations.NewOperation(
		params.Name,
		params.Version,
		params.Description,
		func(b operations.Bundle, chain cldf_evm.Chain, input FunctionInput[ARGS]) (RET, error) {
			var empty RET
			if params.ContractType == "" {
				return empty, fmt.Errorf("contract type must be specified for %s", params.Name)
			}
			if params.CallContract == nil {
				return empty, fmt.Errorf("callContract function must be defined for %s", params.Name)
			}
			c, err := resolveContract(params.Contract, params.NewContract, input.Address, chain)
			if err != nil {
				return empty, err
			}
			if input.ChainSelector != 0 && input.ChainSelector != chain.Selector {
				return empty, fmt.Errorf("chain selector mismatch for %s: input %d vs chain %d", params.Name, input.ChainSelector, chain.Selector)
			}
			return params.CallContract(c, &bind.CallOpts{Context: b.GetContext()}, input.Args)
		},
	)
}
