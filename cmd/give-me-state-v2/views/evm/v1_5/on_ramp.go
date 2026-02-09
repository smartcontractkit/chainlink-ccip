package v1_5

import (
	"give-me-state-v2/views"
	"give-me-state-v2/views/evm/common"
)

// ViewOnRamp generates a view of the OnRamp (EVM2EVMOnRamp) contract (v1.5.0).
// Uses basic fields only due to complex struct decoding requirements for static/dynamic configs.
func ViewOnRamp(ctx *views.ViewContext) (map[string]any, error) {
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
