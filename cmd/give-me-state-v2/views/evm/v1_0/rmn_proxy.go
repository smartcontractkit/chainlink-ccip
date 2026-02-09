package v1_0

import (
	"give-me-state-v2/views"
	"give-me-state-v2/views/evm/common"
)

// Function selectors for RMNProxy (ARMProxy)
var (
	// getARM() returns (address)
	selectorGetARM = common.HexToSelector("2e90aa21")
)

// getRMNProxyARM fetches the ARM (RMN) address using manual decoding.
func getRMNProxyARM(ctx *views.ViewContext) (string, error) {
	data, err := common.ExecuteCall(ctx, selectorGetARM)
	if err != nil {
		return "", err
	}
	return common.DecodeAddress(data)
}

// ViewRMNProxy generates a view of the RMNProxy (ARMProxy) contract (v1.0.0).
func ViewRMNProxy(ctx *views.ViewContext) (map[string]any, error) {
	result := make(map[string]any)

	result["address"] = ctx.AddressHex
	result["chainSelector"] = ctx.ChainSelector
	result["version"] = "1.0.0"

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

	// Get ARM (RMN) address
	arm, err := getRMNProxyARM(ctx)
	if err != nil {
		result["arm_error"] = err.Error()
	} else {
		result["arm"] = arm
	}

	return result, nil
}
