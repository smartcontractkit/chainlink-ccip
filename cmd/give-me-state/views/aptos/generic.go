package aptos

import (
	"encoding/json"
	"fmt"

	"call-orchestrator-demo/views"
)

// ===== Aptos View Functions Using CallManager =====
// All calls go through the rate-limited, cached CallManager.

// ViewCCIP generates a view of the CCIP module.
func ViewCCIP(ctx *views.ViewContext) (map[string]any, error) {
	result := make(map[string]any)
	result["address"] = ctx.AddressHex
	result["chainSelector"] = ctx.ChainSelector

	moduleAddr := ctx.AddressHex

	// Get owner via auth module
	if owner, err := callViewString(ctx, moduleAddr, "auth", "owner", nil, nil); err == nil {
		result["owner"] = owner
	}

	// FeeQuoter
	feeQuoter := make(map[string]any)
	if typeAndVersion, err := callViewString(ctx, moduleAddr, "fee_quoter", "type_and_version", nil, nil); err == nil {
		feeQuoter["typeAndVersion"] = typeAndVersion
	}
	if staticConfig, err := callViewJSON(ctx, moduleAddr, "fee_quoter", "get_static_config", nil, nil); err == nil {
		feeQuoter["staticConfig"] = staticConfig
	}
	if feeTokens, err := callViewJSON(ctx, moduleAddr, "fee_quoter", "get_fee_tokens", nil, nil); err == nil {
		feeQuoter["feeTokens"] = feeTokens
	}
	if len(feeQuoter) > 0 {
		result["feeQuoter"] = feeQuoter
	}

	// RMNRemote
	rmnRemote := make(map[string]any)
	if typeAndVersion, err := callViewString(ctx, moduleAddr, "rmn_remote", "type_and_version", nil, nil); err == nil {
		rmnRemote["typeAndVersion"] = typeAndVersion
	}
	if cursedSubjects, err := callViewJSON(ctx, moduleAddr, "rmn_remote", "get_cursed_subjects", nil, nil); err == nil {
		rmnRemote["cursedSubjects"] = cursedSubjects
	}
	if versionedConfig, err := callViewJSON(ctx, moduleAddr, "rmn_remote", "get_versioned_config", nil, nil); err == nil {
		rmnRemote["versionedConfig"] = versionedConfig
	}
	if len(rmnRemote) > 0 {
		result["rmnRemote"] = rmnRemote
	}

	// TokenAdminRegistry
	if typeAndVersion, err := callViewString(ctx, moduleAddr, "token_admin_registry", "type_and_version", nil, nil); err == nil {
		result["tokenAdminRegistry"] = map[string]any{
			"typeAndVersion": typeAndVersion,
		}
	}

	// NonceManager
	if typeAndVersion, err := callViewString(ctx, moduleAddr, "nonce_manager", "type_and_version", nil, nil); err == nil {
		result["nonceManager"] = map[string]any{
			"typeAndVersion": typeAndVersion,
		}
	}

	return result, nil
}

// ViewRouter generates a view of the Router.
func ViewRouter(ctx *views.ViewContext) (map[string]any, error) {
	result := make(map[string]any)
	result["address"] = ctx.AddressHex
	result["chainSelector"] = ctx.ChainSelector

	moduleAddr := ctx.AddressHex

	if typeAndVersion, err := callViewString(ctx, moduleAddr, "router", "type_and_version", nil, nil); err == nil {
		result["typeAndVersion"] = typeAndVersion
	}
	if owner, err := callViewString(ctx, moduleAddr, "router", "owner", nil, nil); err == nil {
		result["owner"] = owner
	}
	if destChains, err := callViewJSON(ctx, moduleAddr, "router", "get_dest_chains", nil, nil); err == nil {
		result["destChains"] = destChains
	}

	return result, nil
}

// ViewOnRamp generates a view of the OnRamp.
func ViewOnRamp(ctx *views.ViewContext) (map[string]any, error) {
	result := make(map[string]any)
	result["address"] = ctx.AddressHex
	result["chainSelector"] = ctx.ChainSelector

	moduleAddr := ctx.AddressHex

	if typeAndVersion, err := callViewString(ctx, moduleAddr, "onramp", "type_and_version", nil, nil); err == nil {
		result["typeAndVersion"] = typeAndVersion
	}
	if owner, err := callViewString(ctx, moduleAddr, "onramp", "owner", nil, nil); err == nil {
		result["owner"] = owner
	}
	if staticConfig, err := callViewJSON(ctx, moduleAddr, "onramp", "get_static_config", nil, nil); err == nil {
		result["staticConfig"] = staticConfig
	}
	if dynamicConfig, err := callViewJSON(ctx, moduleAddr, "onramp", "get_dynamic_config", nil, nil); err == nil {
		result["dynamicConfig"] = dynamicConfig
	}

	return result, nil
}

// ViewOffRamp generates a view of the OffRamp.
func ViewOffRamp(ctx *views.ViewContext) (map[string]any, error) {
	result := make(map[string]any)
	result["address"] = ctx.AddressHex
	result["chainSelector"] = ctx.ChainSelector

	moduleAddr := ctx.AddressHex

	if typeAndVersion, err := callViewString(ctx, moduleAddr, "offramp", "type_and_version", nil, nil); err == nil {
		result["typeAndVersion"] = typeAndVersion
	}
	if owner, err := callViewString(ctx, moduleAddr, "offramp", "owner", nil, nil); err == nil {
		result["owner"] = owner
	}
	if staticConfig, err := callViewJSON(ctx, moduleAddr, "offramp", "get_static_config", nil, nil); err == nil {
		result["staticConfig"] = staticConfig
	}
	if dynamicConfig, err := callViewJSON(ctx, moduleAddr, "offramp", "get_dynamic_config", nil, nil); err == nil {
		result["dynamicConfig"] = dynamicConfig
	}
	if sourceChainConfigs, err := callViewJSON(ctx, moduleAddr, "offramp", "get_all_source_chain_configs", nil, nil); err == nil {
		result["sourceChainConfigs"] = sourceChainConfigs
	}
	if latestPriceSeq, err := callViewJSON(ctx, moduleAddr, "offramp", "get_latest_price_sequence_number", nil, nil); err == nil {
		result["latestPriceSequenceNumber"] = latestPriceSeq
	}

	return result, nil
}

// ViewMCMS generates a view of the MCMS module.
func ViewMCMS(ctx *views.ViewContext) (map[string]any, error) {
	result := make(map[string]any)
	result["address"] = ctx.AddressHex
	result["chainSelector"] = ctx.ChainSelector

	moduleAddr := ctx.AddressHex

	if owner, err := callViewString(ctx, moduleAddr, "mcms_account", "owner", nil, nil); err == nil {
		result["owner"] = owner
	}
	if minDelay, err := callViewJSON(ctx, moduleAddr, "mcms", "timelock_min_delay", nil, nil); err == nil {
		result["timelockMinDelay"] = minDelay
	}

	// Get configs for each role (bypasser=0, proposer=1, canceller=2)
	if bypasserConfig, err := callViewJSON(ctx, moduleAddr, "mcms", "get_config", nil, []any{"0"}); err == nil {
		result["bypasser"] = bypasserConfig
	}
	if proposerConfig, err := callViewJSON(ctx, moduleAddr, "mcms", "get_config", nil, []any{"1"}); err == nil {
		result["proposer"] = proposerConfig
	}
	if cancellerConfig, err := callViewJSON(ctx, moduleAddr, "mcms", "get_config", nil, []any{"2"}); err == nil {
		result["canceller"] = cancellerConfig
	}

	return result, nil
}

// ViewToken generates a view of a managed token.
func ViewToken(ctx *views.ViewContext) (map[string]any, error) {
	result := make(map[string]any)
	result["address"] = ctx.AddressHex
	result["chainSelector"] = ctx.ChainSelector

	moduleAddr := ctx.AddressHex

	if typeAndVersion, err := callViewString(ctx, moduleAddr, "managed_token", "type_and_version", nil, nil); err == nil {
		result["typeAndVersion"] = typeAndVersion
	}
	if tokenMetadata, err := callViewJSON(ctx, moduleAddr, "managed_token", "token_metadata", nil, nil); err == nil {
		result["tokenMetadata"] = tokenMetadata
	}
	if burners, err := callViewJSON(ctx, moduleAddr, "managed_token", "get_allowed_burners", nil, nil); err == nil {
		result["burners"] = burners
	}
	if minters, err := callViewJSON(ctx, moduleAddr, "managed_token", "get_allowed_minters", nil, nil); err == nil {
		result["minters"] = minters
	}

	return result, nil
}

// ViewTokenPool generates a view of a token pool.
func ViewTokenPool(ctx *views.ViewContext) (map[string]any, error) {
	result := make(map[string]any)
	result["address"] = ctx.AddressHex
	result["chainSelector"] = ctx.ChainSelector

	// Token pools in Aptos may have different structures
	// Add basic view function calls as needed

	return result, nil
}

// ViewGenericAccount provides a basic view for any Aptos account.
func ViewGenericAccount(ctx *views.ViewContext) (map[string]any, error) {
	result := make(map[string]any)
	result["address"] = ctx.AddressHex
	result["chainSelector"] = ctx.ChainSelector
	result["type"] = "AptosGeneric"

	return result, nil
}

// ===== Helper Functions =====

// callViewString executes a view function and returns a string result
func callViewString(ctx *views.ViewContext, moduleAddr, moduleName, funcName string, typeArgs []string, args []any) (string, error) {
	data, err := views.ExecuteAptosView(ctx, moduleAddr, moduleName, funcName, typeArgs, args)
	if err != nil {
		return "", err
	}

	// Aptos view returns JSON array, first element is the result
	var result []any
	if err := json.Unmarshal(data, &result); err != nil {
		return "", fmt.Errorf("failed to parse response: %w", err)
	}
	if len(result) == 0 {
		return "", fmt.Errorf("empty response")
	}

	str, ok := result[0].(string)
	if !ok {
		return fmt.Sprintf("%v", result[0]), nil
	}
	return str, nil
}

// callViewJSON executes a view function and returns the raw JSON result
func callViewJSON(ctx *views.ViewContext, moduleAddr, moduleName, funcName string, typeArgs []string, args []any) (any, error) {
	data, err := views.ExecuteAptosView(ctx, moduleAddr, moduleName, funcName, typeArgs, args)
	if err != nil {
		return nil, err
	}

	// Aptos view returns JSON array
	var result []any
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	// Return first element if single value, otherwise return array
	if len(result) == 1 {
		return result[0], nil
	}
	return result, nil
}
