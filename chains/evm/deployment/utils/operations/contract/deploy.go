package contract

import (
	cldf_evm "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	upstream "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

// DeployInput is the input structure for the Deploy operation.
//
// Deprecated: Import github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract directly.
type DeployInput[ARGS any] = upstream.DeployInput[ARGS]

// Bytecode specifies the exact bytecode to deploy for each supported VM.
//
// Deprecated: Import github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract directly.
type Bytecode = upstream.Bytecode

// DeployParams encapsulates all parameters required to create a deploy operation for an EVM contract.
//
// Deprecated: Import github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract directly.
type DeployParams[ARGS any] = upstream.DeployParams[ARGS]

// NewDeploy creates a new operation that deploys an EVM contract.
//
// Deprecated: Import github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract directly.
func NewDeploy[ARGS any](params DeployParams[ARGS]) *operations.Operation[DeployInput[ARGS], datastore.AddressRef, cldf_evm.Chain] {
	return upstream.NewDeploy(params)
}

// MaybeDeployContract deploys a contract if no matching address ref already exists.
//
// Deprecated: Import github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract directly.
func MaybeDeployContract[ARGS any](
	b operations.Bundle,
	op *operations.Operation[DeployInput[ARGS], datastore.AddressRef, cldf_evm.Chain],
	chain cldf_evm.Chain,
	input DeployInput[ARGS],
	existingAddresses []datastore.AddressRef,
) (datastore.AddressRef, error) {
	return upstream.MaybeDeployContract(b, op, chain, input, existingAddresses)
}
