package v1_6

import (
	"encoding/hex"
	"fmt"
	"sync"

	"give-me-state-v2/views"

	"github.com/ethereum/go-ethereum/common"
)

// packOffRampCall packs a method call using the OffRamp v1.6 ABI.
func packOffRampCall(method string, args ...interface{}) ([]byte, error) {
	return OffRampABI.Pack(method, args...)
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

// getOffRampOwner fetches the owner address.
func getOffRampOwner(ctx *views.ViewContext) (string, error) {
	data, err := executeOffRampCall(ctx, "owner")
	if err != nil {
		return "", err
	}
	results, err := OffRampABI.Unpack("owner", data)
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

// getOffRampTypeAndVersion fetches the typeAndVersion string.
func getOffRampTypeAndVersion(ctx *views.ViewContext) (string, error) {
	data, err := executeOffRampCall(ctx, "typeAndVersion")
	if err != nil {
		return "", err
	}
	results, err := OffRampABI.Unpack("typeAndVersion", data)
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

// getOffRampStaticConfig fetches the static configuration using ABI bindings.
func getOffRampStaticConfig(ctx *views.ViewContext) (map[string]any, error) {
	data, err := executeOffRampCall(ctx, "getStaticConfig")
	if err != nil {
		return nil, err
	}

	results, err := OffRampABI.Unpack("getStaticConfig", data)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack getStaticConfig: %w", err)
	}
	if len(results) == 0 {
		return nil, fmt.Errorf("no results from getStaticConfig call")
	}

	cfg, ok := results[0].(struct {
		ChainSelector        uint64         `json:"chainSelector"`
		GasForCallExactCheck uint16         `json:"gasForCallExactCheck"`
		RmnRemote            common.Address `json:"rmnRemote"`
		TokenAdminRegistry   common.Address `json:"tokenAdminRegistry"`
		NonceManager         common.Address `json:"nonceManager"`
	})
	if !ok {
		return nil, fmt.Errorf("unexpected type for StaticConfig: %T", results[0])
	}

	return map[string]any{
		"chainSelector":        cfg.ChainSelector,
		"gasForCallExactCheck": cfg.GasForCallExactCheck,
		"rmnRemote":            cfg.RmnRemote.Hex(),
		"tokenAdminRegistry":   cfg.TokenAdminRegistry.Hex(),
		"nonceManager":         cfg.NonceManager.Hex(),
	}, nil
}

// getOffRampDynamicConfig fetches the dynamic configuration using ABI bindings.
func getOffRampDynamicConfig(ctx *views.ViewContext) (map[string]any, error) {
	data, err := executeOffRampCall(ctx, "getDynamicConfig")
	if err != nil {
		return nil, err
	}

	results, err := OffRampABI.Unpack("getDynamicConfig", data)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack getDynamicConfig: %w", err)
	}
	if len(results) == 0 {
		return nil, fmt.Errorf("no results from getDynamicConfig call")
	}

	cfg, ok := results[0].(struct {
		FeeQuoter                               common.Address `json:"feeQuoter"`
		PermissionLessExecutionThresholdSeconds uint32         `json:"permissionLessExecutionThresholdSeconds"`
		MessageInterceptor                      common.Address `json:"messageInterceptor"`
	})
	if !ok {
		return nil, fmt.Errorf("unexpected type for DynamicConfig: %T", results[0])
	}

	return map[string]any{
		"feeQuoter":                               cfg.FeeQuoter.Hex(),
		"permissionLessExecutionThresholdSeconds": cfg.PermissionLessExecutionThresholdSeconds,
		"messageInterceptor":                      cfg.MessageInterceptor.Hex(),
	}, nil
}

// getOffRampLatestPriceSeqNum fetches the latest price sequence number.
func getOffRampLatestPriceSeqNum(ctx *views.ViewContext) (uint64, error) {
	data, err := executeOffRampCall(ctx, "getLatestPriceSequenceNumber")
	if err != nil {
		return 0, err
	}
	results, err := OffRampABI.Unpack("getLatestPriceSequenceNumber", data)
	if err != nil {
		return 0, fmt.Errorf("failed to unpack getLatestPriceSequenceNumber: %w", err)
	}
	if len(results) == 0 {
		return 0, fmt.Errorf("no results from getLatestPriceSequenceNumber call")
	}
	seqNum, ok := results[0].(uint64)
	if !ok {
		return 0, fmt.Errorf("unexpected type for sequence number: %T", results[0])
	}
	return seqNum, nil
}

// getSourceChainConfigForChain fetches the SourceChainConfig for a specific chain.
func getSourceChainConfigForChain(ctx *views.ViewContext, chainSel uint64) (map[string]any, error) {
	data, err := executeOffRampCall(ctx, "getSourceChainConfig", chainSel)
	if err != nil {
		return nil, err
	}

	results, err := OffRampABI.Unpack("getSourceChainConfig", data)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack getSourceChainConfig: %w", err)
	}
	if len(results) == 0 {
		return nil, fmt.Errorf("no results from getSourceChainConfig call")
	}

	cfg, ok := results[0].(struct {
		Router                    common.Address `json:"router"`
		IsEnabled                 bool           `json:"isEnabled"`
		MinSeqNr                  uint64         `json:"minSeqNr"`
		IsRMNVerificationDisabled bool           `json:"isRMNVerificationDisabled"`
		OnRamp                    []byte         `json:"onRamp"`
	})
	if !ok {
		return nil, fmt.Errorf("unexpected type for SourceChainConfig: %T", results[0])
	}

	config := map[string]any{
		"router":                    cfg.Router.Hex(),
		"isEnabled":                 cfg.IsEnabled,
		"minSeqNr":                  cfg.MinSeqNr,
		"isRMNVerificationDisabled": cfg.IsRMNVerificationDisabled,
	}

	// Format onRamp bytes based on length
	if len(cfg.OnRamp) == 20 {
		config["onRamp"] = common.BytesToAddress(cfg.OnRamp).Hex()
	} else if len(cfg.OnRamp) > 0 {
		config["onRamp"] = "0x" + hex.EncodeToString(cfg.OnRamp)
	}

	return config, nil
}

// getOffRampSourceChainConfigs fetches source chain configs concurrently.
func getOffRampSourceChainConfigs(ctx *views.ViewContext) (map[string]any, error) {
	if len(ctx.AllChainSelectors) == 0 {
		return map[string]any{}, nil
	}

	result := make(map[string]any)
	var mu sync.Mutex
	var wg sync.WaitGroup

	for _, chainSel := range ctx.AllChainSelectors {
		wg.Add(1)
		go func(cs uint64) {
			defer wg.Done()

			config, err := getSourceChainConfigForChain(ctx, cs)
			if err != nil {
				return
			}

			isEnabled, ok := config["isEnabled"].(bool)
			if !ok || !isEnabled {
				return
			}

			mu.Lock()
			result[views.Uint64ToString(cs)] = config
			mu.Unlock()
		}(chainSel)
	}

	wg.Wait()
	return result, nil
}

// ViewOffRamp generates a view of the OffRamp contract (v1.6.0).
// Uses ABI bindings for proper struct decoding.
func ViewOffRamp(ctx *views.ViewContext) (map[string]any, error) {
	result := make(map[string]any)

	result["address"] = ctx.AddressHex
	result["chainSelector"] = ctx.ChainSelector
	result["version"] = "1.6.0"

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

	staticConfig, err := getOffRampStaticConfig(ctx)
	if err != nil {
		result["staticConfig_error"] = err.Error()
	} else {
		result["staticConfig"] = staticConfig
	}

	dynamicConfig, err := getOffRampDynamicConfig(ctx)
	if err != nil {
		result["dynamicConfig_error"] = err.Error()
	} else {
		result["dynamicConfig"] = dynamicConfig
	}

	latestPriceSeqNum, err := getOffRampLatestPriceSeqNum(ctx)
	if err == nil {
		result["latestPriceSequenceNumber"] = latestPriceSeqNum
	}

	sourceChainConfigs, err := getOffRampSourceChainConfigs(ctx)
	if err != nil {
		result["sourceChainConfigs_error"] = err.Error()
	} else if len(sourceChainConfigs) > 0 {
		result["sourceChainConfigs"] = sourceChainConfigs
	}

	return result, nil
}
