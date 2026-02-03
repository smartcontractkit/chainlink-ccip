package v1_5_1

import (
	"call-orchestrator-demo/views"
	"call-orchestrator-demo/views/evm/common"
)

// ViewTokenPoolFactory generates a view of the TokenPoolFactory contract (v1.5.1).
// Note: This contract does NOT have an owner() method.
func ViewTokenPoolFactory(ctx *views.ViewContext) (map[string]any, error) {
	result := make(map[string]any)

	result["address"] = ctx.AddressHex
	result["chainSelector"] = ctx.ChainSelector
	result["version"] = "1.5.1"

	// Get typeAndVersion (this contract does NOT have owner())
	typeAndVersion, err := common.GetTypeAndVersion(ctx)
	if err != nil {
		result["typeAndVersion_error"] = err.Error()
	} else {
		result["typeAndVersion"] = typeAndVersion
	}

	return result, nil
}
