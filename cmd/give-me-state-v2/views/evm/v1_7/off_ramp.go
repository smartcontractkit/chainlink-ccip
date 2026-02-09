package v1_7

import (
	"fmt"

	"give-me-state-v2/views"

	"github.com/ethereum/go-ethereum/common"

	offramp "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/offramp"
)

// packOffRampCall packs a method call using the OffRamp ABI and returns the calldata bytes.
func packOffRampCall(method string, args ...interface{}) ([]byte, error) {
	return OffRampABI.Pack(method, args...)
}

// executeOffRampCall packs a call, executes it via the orchestrator, and returns raw response bytes.
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

// getOffRampOwner fetches the owner address using the OffRamp bindings.
func getOffRampOwner(ctx *views.ViewContext) (string, error) {
	data, err := executeOffRampCall(ctx, "owner")
	if err != nil {
		return "", err
	}

	results, err := OffRampABI.Unpack("owner", data)
	if err != nil {
		return "", fmt.Errorf("failed to unpack owner response: %w", err)
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

// getOffRampTypeAndVersion fetches the typeAndVersion string using the OffRamp bindings.
func getOffRampTypeAndVersion(ctx *views.ViewContext) (string, error) {
	data, err := executeOffRampCall(ctx, "typeAndVersion")
	if err != nil {
		return "", err
	}

	results, err := OffRampABI.Unpack("typeAndVersion", data)
	if err != nil {
		return "", fmt.Errorf("failed to unpack typeAndVersion response: %w", err)
	}

	if len(results) == 0 {
		return "", fmt.Errorf("no results from typeAndVersion call")
	}

	typeAndVersion, ok := results[0].(string)
	if !ok {
		return "", fmt.Errorf("unexpected type for typeAndVersion: %T", results[0])
	}

	return typeAndVersion, nil
}

// getOffRampAllSourceChainConfigs fetches all source chain configs using the OffRamp bindings.
func getOffRampAllSourceChainConfigs(ctx *views.ViewContext) ([]uint64, []offramp.OffRampSourceChainConfig, error) {
	data, err := executeOffRampCall(ctx, "getAllSourceChainConfigs")
	if err != nil {
		return nil, nil, err
	}

	results, err := OffRampABI.Unpack("getAllSourceChainConfigs", data)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to unpack getAllSourceChainConfigs response: %w", err)
	}

	if len(results) < 2 {
		return nil, nil, fmt.Errorf("expected 2 results from getAllSourceChainConfigs, got %d", len(results))
	}

	selectors, ok := results[0].([]uint64)
	if !ok {
		return nil, nil, fmt.Errorf("unexpected type for sourceChainSelectors: %T", results[0])
	}

	configs, ok := results[1].([]struct {
		Router           common.Address   `json:"router"`
		IsEnabled        bool             `json:"isEnabled"`
		OnRamps          [][]byte         `json:"onRamps"`
		DefaultCCVs      []common.Address `json:"defaultCCVs"`
		LaneMandatedCCVs []common.Address `json:"laneMandatedCCVs"`
	})
	if !ok {
		return nil, nil, fmt.Errorf("unexpected type for sourceChainConfigs: %T", results[1])
	}

	sourceConfigs := make([]offramp.OffRampSourceChainConfig, len(configs))
	for i, cfg := range configs {
		sourceConfigs[i] = offramp.OffRampSourceChainConfig{
			Router:           cfg.Router,
			IsEnabled:        cfg.IsEnabled,
			OnRamps:          cfg.OnRamps,
			DefaultCCVs:      cfg.DefaultCCVs,
			LaneMandatedCCVs: cfg.LaneMandatedCCVs,
		}
	}

	return selectors, sourceConfigs, nil
}

// collectSourceChainConfigs collects all source chain configs and formats them for the view.
func collectSourceChainConfigs(ctx *views.ViewContext) ([]map[string]any, error) {
	selectors, configs, err := getOffRampAllSourceChainConfigs(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get source chain configs: %w", err)
	}

	sourceChainConfigsList := make([]map[string]any, 0, len(selectors))

	for i, sourceChainSelector := range selectors {
		config := configs[i]

		defaultCCVsHex := make([]string, len(config.DefaultCCVs))
		for j, ccv := range config.DefaultCCVs {
			defaultCCVsHex[j] = ccv.Hex()
		}

		laneMandatedCCVsHex := make([]string, len(config.LaneMandatedCCVs))
		for j, ccv := range config.LaneMandatedCCVs {
			laneMandatedCCVsHex[j] = ccv.Hex()
		}

		onRampsHex := make([]string, len(config.OnRamps))
		for j, onRampBytes := range config.OnRamps {
			onRampsHex[j] = fmt.Sprintf("0x%x", onRampBytes)
		}

		sourceChainConfigsList = append(sourceChainConfigsList, map[string]any{
			"sourceChainSelector": sourceChainSelector,
			"router":              config.Router.Hex(),
			"isEnabled":           config.IsEnabled,
			"onRamps":             onRampsHex,
			"defaultCCVs":         defaultCCVsHex,
			"laneMandatedCCVs":    laneMandatedCCVsHex,
		})
	}

	return sourceChainConfigsList, nil
}

// ViewOffRamp generates a view of the OffRamp contract (v1.7.0).
// Uses the generated bindings to pack/unpack calls.
func ViewOffRamp(ctx *views.ViewContext) (map[string]any, error) {
	result := make(map[string]any)

	result["address"] = ctx.AddressHex
	result["chainSelector"] = ctx.ChainSelector
	result["version"] = "1.7.0"

	owner, err := getOffRampOwner(ctx)
	if err != nil {
		result["owner_error"] = err.Error()
	} else {
		result["owner"] = owner
	}

	typeAndVersion, err := getOffRampTypeAndVersion(ctx)
	if err != nil {
		result["typeAndVersion_error"] = err.Error()
	} else {
		result["typeAndVersion"] = typeAndVersion
	}

	sourceChainConfigs, err := collectSourceChainConfigs(ctx)
	if err != nil {
		result["sourceChainConfigs_error"] = err.Error()
	} else {
		result["sourceChainConfigs"] = sourceChainConfigs
	}

	return result, nil
}
