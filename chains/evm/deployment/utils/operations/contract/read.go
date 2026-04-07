package contract

import (
	cldf_evm "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	upstream "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

// ReadParams contains the parameters to create a read operation.
//
// Deprecated: Import github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract directly.
type ReadParams[ARGS any, RET any, C any] = upstream.ReadParams[ARGS, RET, C]

// NewRead creates a new read operation.
//
// Deprecated: Import github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract directly.
func NewRead[ARGS any, RET any, C any](params ReadParams[ARGS, RET, C]) *operations.Operation[FunctionInput[ARGS], RET, cldf_evm.Chain] {
	return upstream.NewRead(params)
}
