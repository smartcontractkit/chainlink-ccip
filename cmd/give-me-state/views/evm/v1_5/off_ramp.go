package v1_5

import (
	"call-orchestrator-demo/views"
	"call-orchestrator-demo/views/evm/common"
)

// ViewOffRamp generates a view of the OffRamp (EVM2EVMOffRamp) contract (v1.5.0).
// Uses basic fields only due to complex struct decoding requirements for static/dynamic configs.
func ViewOffRamp(ctx *views.ViewContext) (map[string]any, error) {
	result := make(map[string]any)

	result["address"] = ctx.AddressHex
	result["chainSelector"] = ctx.ChainSelector
	result["version"] = "1.5.0"

	// Get owner
	owner, err := common.GetOwner(ctx)
	if err != nil {
		result["owner_error"] = err.Error()
	} else {
		result["owner"] = owner
	}

	// Get typeAndVersion
	typeAndVersion, err := common.GetTypeAndVersion(ctx)
	if err != nil {
		result["typeAndVersion_error"] = err.Error()
	} else {
		result["typeAndVersion"] = typeAndVersion
	}

	// Note: staticConfig and dynamicConfig require complex struct decoding
	// that isn't easily done without ABI bindings

	return result, nil
}
