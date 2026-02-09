package v1_6

import (
	"give-me-state-v2/views"
	"give-me-state-v2/views/evm/common"
)

// Function selectors for RMNRemote
var (
	// isCursed() returns (bool)
	selectorIsCursed = common.HexToSelector("397796f7")
)

// getRMNRemoteIsCursed fetches the global curse status using manual decoding.
func getRMNRemoteIsCursed(ctx *views.ViewContext) (bool, error) {
	data, err := common.ExecuteCall(ctx, selectorIsCursed)
	if err != nil {
		return false, err
	}
	return common.DecodeBool(data)
}

// ViewRMNRemote generates a view of the RMNRemote contract (v1.6.0).
// Uses basic fields only due to complex struct decoding requirements.
func ViewRMNRemote(ctx *views.ViewContext) (map[string]any, error) {
	result := make(map[string]any)

	result["address"] = ctx.AddressHex
	result["chainSelector"] = ctx.ChainSelector
	result["version"] = "1.6.0"

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

	// Get isCursed
	isCursed, err := getRMNRemoteIsCursed(ctx)
	if err != nil {
		result["isCursed_error"] = err.Error()
	} else {
		result["isCursed"] = isCursed
	}

	return result, nil
}
