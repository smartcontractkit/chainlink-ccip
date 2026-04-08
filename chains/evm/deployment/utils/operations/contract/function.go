package contract

import upstream "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"

// FunctionInput is the input structure for all reads and writes.
//
// Deprecated: Import github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract directly.
type FunctionInput[ARGS any] = upstream.FunctionInput[ARGS]
