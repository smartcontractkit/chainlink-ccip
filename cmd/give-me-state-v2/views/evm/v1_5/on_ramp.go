package v1_5

import (
	"fmt"
	"math/big"

	"give-me-state-v2/views"

	"github.com/ethereum/go-ethereum/common"
)

// packOnRampCall packs a method call using the EVM2EVMOnRamp v1.5 ABI.
func packOnRampCall(method string, args ...interface{}) ([]byte, error) {
	return EVM2EVMOnRampABI.Pack(method, args...)
}

// executeOnRampCall packs a call, executes it, and returns raw response bytes.
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

// getOnRampV15Owner fetches the owner address.
func getOnRampV15Owner(ctx *views.ViewContext) (string, error) {
	data, err := executeOnRampCall(ctx, "owner")
	if err != nil {
		return "", err
	}
	results, err := EVM2EVMOnRampABI.Unpack("owner", data)
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

// getOnRampV15TypeAndVersion fetches the typeAndVersion string.
func getOnRampV15TypeAndVersion(ctx *views.ViewContext) (string, error) {
	data, err := executeOnRampCall(ctx, "typeAndVersion")
	if err != nil {
		return "", err
	}
	results, err := EVM2EVMOnRampABI.Unpack("typeAndVersion", data)
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

// getOnRampV15StaticConfig fetches the static configuration using ABI bindings.
func getOnRampV15StaticConfig(ctx *views.ViewContext) (map[string]any, error) {
	data, err := executeOnRampCall(ctx, "getStaticConfig")
	if err != nil {
		return nil, err
	}

	results, err := EVM2EVMOnRampABI.Unpack("getStaticConfig", data)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack getStaticConfig: %w", err)
	}
	if len(results) == 0 {
		return nil, fmt.Errorf("no results from getStaticConfig call")
	}

	cfg, ok := results[0].(struct {
		LinkToken          common.Address `json:"linkToken"`
		ChainSelector      uint64         `json:"chainSelector"`
		DestChainSelector  uint64         `json:"destChainSelector"`
		DefaultTxGasLimit  uint64         `json:"defaultTxGasLimit"`
		MaxNopFeesJuels    *big.Int       `json:"maxNopFeesJuels"`
		PrevOnRamp         common.Address `json:"prevOnRamp"`
		RmnProxy           common.Address `json:"rmnProxy"`
		TokenAdminRegistry common.Address `json:"tokenAdminRegistry"`
	})
	if !ok {
		return nil, fmt.Errorf("unexpected type for StaticConfig: %T", results[0])
	}

	return map[string]any{
		"linkToken":          cfg.LinkToken.Hex(),
		"chainSelector":      cfg.ChainSelector,
		"destChainSelector":  cfg.DestChainSelector,
		"defaultTxGasLimit":  cfg.DefaultTxGasLimit,
		"maxNopFeesJuels":    cfg.MaxNopFeesJuels.String(),
		"prevOnRamp":         cfg.PrevOnRamp.Hex(),
		"rmnProxy":           cfg.RmnProxy.Hex(),
		"tokenAdminRegistry": cfg.TokenAdminRegistry.Hex(),
	}, nil
}

// getOnRampV15DynamicConfig fetches the dynamic configuration using ABI bindings.
func getOnRampV15DynamicConfig(ctx *views.ViewContext) (map[string]any, error) {
	data, err := executeOnRampCall(ctx, "getDynamicConfig")
	if err != nil {
		return nil, err
	}

	results, err := EVM2EVMOnRampABI.Unpack("getDynamicConfig", data)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack getDynamicConfig: %w", err)
	}
	if len(results) == 0 {
		return nil, fmt.Errorf("no results from getDynamicConfig call")
	}

	cfg, ok := results[0].(struct {
		Router                            common.Address `json:"router"`
		MaxNumberOfTokensPerMsg           uint16         `json:"maxNumberOfTokensPerMsg"`
		DestGasOverhead                   uint32         `json:"destGasOverhead"`
		DestGasPerPayloadByte             uint16         `json:"destGasPerPayloadByte"`
		DestDataAvailabilityOverheadGas   uint32         `json:"destDataAvailabilityOverheadGas"`
		DestGasPerDataAvailabilityByte    uint16         `json:"destGasPerDataAvailabilityByte"`
		DestDataAvailabilityMultiplierBps uint16         `json:"destDataAvailabilityMultiplierBps"`
		PriceRegistry                     common.Address `json:"priceRegistry"`
		MaxDataBytes                      uint32         `json:"maxDataBytes"`
		MaxPerMsgGasLimit                 uint32         `json:"maxPerMsgGasLimit"`
		DefaultTokenFeeUSDCents           uint16         `json:"defaultTokenFeeUSDCents"`
		DefaultTokenDestGasOverhead       uint32         `json:"defaultTokenDestGasOverhead"`
		EnforceOutOfOrder                 bool           `json:"enforceOutOfOrder"`
	})
	if !ok {
		return nil, fmt.Errorf("unexpected type for DynamicConfig: %T", results[0])
	}

	return map[string]any{
		"router":                            cfg.Router.Hex(),
		"maxNumberOfTokensPerMsg":           cfg.MaxNumberOfTokensPerMsg,
		"destGasOverhead":                   cfg.DestGasOverhead,
		"destGasPerPayloadByte":             cfg.DestGasPerPayloadByte,
		"destDataAvailabilityOverheadGas":   cfg.DestDataAvailabilityOverheadGas,
		"destGasPerDataAvailabilityByte":    cfg.DestGasPerDataAvailabilityByte,
		"destDataAvailabilityMultiplierBps": cfg.DestDataAvailabilityMultiplierBps,
		"priceRegistry":                     cfg.PriceRegistry.Hex(),
		"maxDataBytes":                      cfg.MaxDataBytes,
		"maxPerMsgGasLimit":                 cfg.MaxPerMsgGasLimit,
		"defaultTokenFeeUSDCents":           cfg.DefaultTokenFeeUSDCents,
		"defaultTokenDestGasOverhead":       cfg.DefaultTokenDestGasOverhead,
		"enforceOutOfOrder":                 cfg.EnforceOutOfOrder,
	}, nil
}

// ViewOnRamp generates a view of the OnRamp (EVM2EVMOnRamp) contract (v1.5.0).
// Uses ABI bindings for proper struct decoding including static and dynamic configs.
func ViewOnRamp(ctx *views.ViewContext) (map[string]any, error) {
	result := make(map[string]any)

	result["address"] = ctx.AddressHex
	result["chainSelector"] = ctx.ChainSelector
	result["version"] = "1.5.0"

	owner, err := getOnRampV15Owner(ctx)
	if err != nil {
		result["owner_error"] = err.Error()
	} else {
		result["owner"] = owner
	}

	typeAndVersion, err := getOnRampV15TypeAndVersion(ctx)
	if err != nil {
		result["typeAndVersion_error"] = err.Error()
	} else {
		result["typeAndVersion"] = typeAndVersion
	}

	staticConfig, err := getOnRampV15StaticConfig(ctx)
	if err != nil {
		result["staticConfig_error"] = err.Error()
	} else {
		result["staticConfig"] = staticConfig
	}

	dynamicConfig, err := getOnRampV15DynamicConfig(ctx)
	if err != nil {
		result["dynamicConfig_error"] = err.Error()
	} else {
		result["dynamicConfig"] = dynamicConfig
	}

	return result, nil
}
