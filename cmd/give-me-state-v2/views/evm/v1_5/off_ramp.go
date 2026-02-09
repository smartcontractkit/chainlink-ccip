package v1_5

import (
	"fmt"

	"give-me-state-v2/views"

	"github.com/ethereum/go-ethereum/common"
)

// packOffRampCall packs a method call using the EVM2EVMOffRamp v1.5 ABI.
func packOffRampCall(method string, args ...interface{}) ([]byte, error) {
	return EVM2EVMOffRampABI.Pack(method, args...)
}

// executeOffRampCall packs a call, executes it, and returns raw response bytes.
func executeOffRampCall(ctx *views.ViewContext, method string, args ...interface{}) ([]byte, error) {
	calldata, err := packOffRampCall(method, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to pack %s call: %w", method, err)
	}

	call := views.Call{
		ChainID: ctx.ChainSelector,
		Target:  ctx.Address,
		Data:    calldata,
	}

	result := ctx.TypedOrchestrator.Execute(call)
	if result.Error != nil {
		return nil, fmt.Errorf("%s call failed: %w", method, result.Error)
	}

	return result.Data, nil
}

// getOffRampV15Owner fetches the owner address.
func getOffRampV15Owner(ctx *views.ViewContext) (string, error) {
	data, err := executeOffRampCall(ctx, "owner")
	if err != nil {
		return "", err
	}
	results, err := EVM2EVMOffRampABI.Unpack("owner", data)
	if err != nil {
		return "", fmt.Errorf("failed to unpack owner: %w", err)
	}
	if len(results) == 0 {
		return "", fmt.Errorf("no results from owner call")
	}
	owner, ok := results[0].(common.Address)
	if !ok {
		return "", fmt.Errorf("unexpected type for owner: %T", results[0])
	}
	return owner.Hex(), nil
}

// getOffRampV15TypeAndVersion fetches the typeAndVersion string.
func getOffRampV15TypeAndVersion(ctx *views.ViewContext) (string, error) {
	data, err := executeOffRampCall(ctx, "typeAndVersion")
	if err != nil {
		return "", err
	}
	results, err := EVM2EVMOffRampABI.Unpack("typeAndVersion", data)
	if err != nil {
		return "", fmt.Errorf("failed to unpack typeAndVersion: %w", err)
	}
	if len(results) == 0 {
		return "", fmt.Errorf("no results from typeAndVersion call")
	}
	tv, ok := results[0].(string)
	if !ok {
		return "", fmt.Errorf("unexpected type for typeAndVersion: %T", results[0])
	}
	return tv, nil
}

// getOffRampV15StaticConfig fetches the static configuration using ABI bindings.
func getOffRampV15StaticConfig(ctx *views.ViewContext) (map[string]any, error) {
	data, err := executeOffRampCall(ctx, "getStaticConfig")
	if err != nil {
		return nil, err
	}

	results, err := EVM2EVMOffRampABI.Unpack("getStaticConfig", data)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack getStaticConfig: %w", err)
	}
	if len(results) == 0 {
		return nil, fmt.Errorf("no results from getStaticConfig call")
	}

	cfg, ok := results[0].(struct {
		CommitStore         common.Address `json:"commitStore"`
		ChainSelector       uint64         `json:"chainSelector"`
		SourceChainSelector uint64         `json:"sourceChainSelector"`
		OnRamp              common.Address `json:"onRamp"`
		PrevOffRamp         common.Address `json:"prevOffRamp"`
		RmnProxy            common.Address `json:"rmnProxy"`
		TokenAdminRegistry  common.Address `json:"tokenAdminRegistry"`
	})
	if !ok {
		return nil, fmt.Errorf("unexpected type for StaticConfig: %T", results[0])
	}

	return map[string]any{
		"commitStore":         cfg.CommitStore.Hex(),
		"chainSelector":       cfg.ChainSelector,
		"sourceChainSelector": cfg.SourceChainSelector,
		"onRamp":              cfg.OnRamp.Hex(),
		"prevOffRamp":         cfg.PrevOffRamp.Hex(),
		"rmnProxy":            cfg.RmnProxy.Hex(),
		"tokenAdminRegistry":  cfg.TokenAdminRegistry.Hex(),
	}, nil
}

// getOffRampV15DynamicConfig fetches the dynamic configuration using ABI bindings.
func getOffRampV15DynamicConfig(ctx *views.ViewContext) (map[string]any, error) {
	data, err := executeOffRampCall(ctx, "getDynamicConfig")
	if err != nil {
		return nil, err
	}

	results, err := EVM2EVMOffRampABI.Unpack("getDynamicConfig", data)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack getDynamicConfig: %w", err)
	}
	if len(results) == 0 {
		return nil, fmt.Errorf("no results from getDynamicConfig call")
	}

	cfg, ok := results[0].(struct {
		PermissionLessExecutionThresholdSeconds uint32         `json:"permissionLessExecutionThresholdSeconds"`
		MaxDataBytes                            uint32         `json:"maxDataBytes"`
		MaxNumberOfTokensPerMsg                 uint16         `json:"maxNumberOfTokensPerMsg"`
		Router                                  common.Address `json:"router"`
		PriceRegistry                           common.Address `json:"priceRegistry"`
	})
	if !ok {
		return nil, fmt.Errorf("unexpected type for DynamicConfig: %T", results[0])
	}

	return map[string]any{
		"permissionLessExecutionThresholdSeconds": cfg.PermissionLessExecutionThresholdSeconds,
		"maxDataBytes":                            cfg.MaxDataBytes,
		"maxNumberOfTokensPerMsg":                 cfg.MaxNumberOfTokensPerMsg,
		"router":                                  cfg.Router.Hex(),
		"priceRegistry":                           cfg.PriceRegistry.Hex(),
	}, nil
}

// ViewOffRamp generates a view of the OffRamp (EVM2EVMOffRamp) contract (v1.5.0).
// Uses ABI bindings for proper struct decoding including static and dynamic configs.
func ViewOffRamp(ctx *views.ViewContext) (map[string]any, error) {
	result := make(map[string]any)

	result["address"] = ctx.AddressHex
	result["chainSelector"] = ctx.ChainSelector
	result["version"] = "1.5.0"

	owner, err := getOffRampV15Owner(ctx)
	if err != nil {
		result["owner_error"] = err.Error()
	} else {
		result["owner"] = owner
	}

	typeAndVersion, err := getOffRampV15TypeAndVersion(ctx)
	if err != nil {
		result["typeAndVersion_error"] = err.Error()
	} else {
		result["typeAndVersion"] = typeAndVersion
	}

	staticConfig, err := getOffRampV15StaticConfig(ctx)
	if err != nil {
		result["staticConfig_error"] = err.Error()
	} else {
		result["staticConfig"] = staticConfig
	}

	dynamicConfig, err := getOffRampV15DynamicConfig(ctx)
	if err != nil {
		result["dynamicConfig_error"] = err.Error()
	} else {
		result["dynamicConfig"] = dynamicConfig
	}

	return result, nil
}
