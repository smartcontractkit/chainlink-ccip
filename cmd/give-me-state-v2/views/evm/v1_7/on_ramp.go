package v1_7

import (
	"fmt"

	"give-me-state-v2/views"

	"github.com/ethereum/go-ethereum/common"

	onramp "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/onramp"
)

// packOnRampCall packs a method call using the OnRamp ABI and returns the calldata bytes.
func packOnRampCall(method string, args ...interface{}) ([]byte, error) {
	return OnRampABI.Pack(method, args...)
}

// executeOnRampCall packs a call, executes it via the orchestrator, and returns raw response bytes.
func executeOnRampCall(ctx *views.ViewContext, method string, args ...interface{}) ([]byte, error) {
	calldata, err := packOnRampCall(method, args...)
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

// getOnRampOwner fetches the owner address using the OnRamp bindings.
func getOnRampOwner(ctx *views.ViewContext) (string, error) {
	data, err := executeOnRampCall(ctx, "owner")
	if err != nil {
		return "", err
	}

	results, err := OnRampABI.Unpack("owner", data)
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

// getOnRampTypeAndVersion fetches the typeAndVersion string using the OnRamp bindings.
func getOnRampTypeAndVersion(ctx *views.ViewContext) (string, error) {
	data, err := executeOnRampCall(ctx, "typeAndVersion")
	if err != nil {
		return "", err
	}

	results, err := OnRampABI.Unpack("typeAndVersion", data)
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

// getOnRampAllDestChainConfigs fetches all destination chain configs using the OnRamp bindings.
func getOnRampAllDestChainConfigs(ctx *views.ViewContext) ([]uint64, []onramp.OnRampDestChainConfig, error) {
	data, err := executeOnRampCall(ctx, "getAllDestChainConfigs")
	if err != nil {
		return nil, nil, err
	}

	results, err := OnRampABI.Unpack("getAllDestChainConfigs", data)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to unpack getAllDestChainConfigs response: %w", err)
	}

	if len(results) < 2 {
		return nil, nil, fmt.Errorf("expected 2 results from getAllDestChainConfigs, got %d", len(results))
	}

	selectors, ok := results[0].([]uint64)
	if !ok {
		return nil, nil, fmt.Errorf("unexpected type for destChainSelectors: %T", results[0])
	}

	configs, ok := results[1].([]struct {
		Router                    common.Address   `json:"router"`
		MessageNumber             uint64           `json:"messageNumber"`
		AddressBytesLength        uint8            `json:"addressBytesLength"`
		TokenReceiverAllowed      bool             `json:"tokenReceiverAllowed"`
		MessageNetworkFeeUSDCents uint16           `json:"messageNetworkFeeUSDCents"`
		TokenNetworkFeeUSDCents   uint16           `json:"tokenNetworkFeeUSDCents"`
		BaseExecutionGasCost      uint32           `json:"baseExecutionGasCost"`
		DefaultExecutor           common.Address   `json:"defaultExecutor"`
		LaneMandatedCCVs          []common.Address `json:"laneMandatedCCVs"`
		DefaultCCVs               []common.Address `json:"defaultCCVs"`
		OffRamp                   []byte           `json:"offRamp"`
	})
	if !ok {
		return nil, nil, fmt.Errorf("unexpected type for destChainConfigs: %T", results[1])
	}

	destConfigs := make([]onramp.OnRampDestChainConfig, len(configs))
	for i, cfg := range configs {
		destConfigs[i] = onramp.OnRampDestChainConfig{
			Router:                    cfg.Router,
			MessageNumber:             cfg.MessageNumber,
			AddressBytesLength:        cfg.AddressBytesLength,
			TokenReceiverAllowed:      cfg.TokenReceiverAllowed,
			MessageNetworkFeeUSDCents: cfg.MessageNetworkFeeUSDCents,
			TokenNetworkFeeUSDCents:   cfg.TokenNetworkFeeUSDCents,
			BaseExecutionGasCost:      cfg.BaseExecutionGasCost,
			DefaultExecutor:           cfg.DefaultExecutor,
			LaneMandatedCCVs:          cfg.LaneMandatedCCVs,
			DefaultCCVs:               cfg.DefaultCCVs,
			OffRamp:                   cfg.OffRamp,
		}
	}

	return selectors, destConfigs, nil
}

// collectDestChainConfigs collects all destination chain configs and formats them for the view.
func collectDestChainConfigs(ctx *views.ViewContext) ([]map[string]any, error) {
	selectors, configs, err := getOnRampAllDestChainConfigs(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get dest chain configs: %w", err)
	}

	destChainConfigsList := make([]map[string]any, 0, len(selectors))

	for i, destChainSelector := range selectors {
		config := configs[i]

		defaultCCVsHex := make([]string, len(config.DefaultCCVs))
		for j, ccv := range config.DefaultCCVs {
			defaultCCVsHex[j] = ccv.Hex()
		}

		laneMandatedCCVsHex := make([]string, len(config.LaneMandatedCCVs))
		for j, ccv := range config.LaneMandatedCCVs {
			laneMandatedCCVsHex[j] = ccv.Hex()
		}

		destChainConfigsList = append(destChainConfigsList, map[string]any{
			"destChainSelector":         destChainSelector,
			"router":                    config.Router.Hex(),
			"messageNumber":             config.MessageNumber,
			"addressBytesLength":        config.AddressBytesLength,
			"tokenReceiverAllowed":      config.TokenReceiverAllowed,
			"messageNetworkFeeUSDCents": config.MessageNetworkFeeUSDCents,
			"tokenNetworkFeeUSDCents":   config.TokenNetworkFeeUSDCents,
			"baseExecutionGasCost":      config.BaseExecutionGasCost,
			"defaultExecutor":           config.DefaultExecutor.Hex(),
			"laneMandatedCCVs":          laneMandatedCCVsHex,
			"defaultCCVs":               defaultCCVsHex,
			"offRamp":                   fmt.Sprintf("0x%x", config.OffRamp),
		})
	}

	return destChainConfigsList, nil
}

// ViewOnRamp generates a view of the OnRamp contract (v1.7.0).
// Uses the generated bindings to pack/unpack calls.
func ViewOnRamp(ctx *views.ViewContext) (map[string]any, error) {
	result := make(map[string]any)

	result["address"] = ctx.AddressHex
	result["chainSelector"] = ctx.ChainSelector
	result["version"] = "1.7.0"

	owner, err := getOnRampOwner(ctx)
	if err != nil {
		result["owner_error"] = err.Error()
	} else {
		result["owner"] = owner
	}

	typeAndVersion, err := getOnRampTypeAndVersion(ctx)
	if err != nil {
		result["typeAndVersion_error"] = err.Error()
	} else {
		result["typeAndVersion"] = typeAndVersion
	}

	destChainConfigs, err := collectDestChainConfigs(ctx)
	if err != nil {
		result["destChainConfigs_error"] = err.Error()
	} else {
		result["destChainConfigs"] = destChainConfigs
	}

	return result, nil
}
