package v1_5

import (
	"give-me-state-v2/views"
	"give-me-state-v2/views/evm/common"
)

// ViewRegistryModuleOwnerCustom generates a view of the RegistryModuleOwnerCustom contract (v1.5.0).
// Note: This contract does NOT have an owner() method, only typeAndVersion().
func ViewRegistryModuleOwnerCustom(ctx *views.ViewContext) (map[string]any, error) {
	result := make(map[string]any)

	result["address"] = ctx.AddressHex
	result["chainSelector"] = ctx.ChainSelector
	result["version"] = "1.5.0"

	// Get typeAndVersion (this contract does NOT have owner())
	typeAndVersion, err := common.GetTypeAndVersion(ctx)
	if err != nil {
		result["typeAndVersion_error"] = err.Error()
	} else {
		result["typeAndVersion"] = typeAndVersion
	}

	return result, nil
}
