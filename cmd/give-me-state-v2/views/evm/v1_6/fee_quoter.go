package v1_6

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
	"sync"

	"give-me-state-v2/views"
	"give-me-state-v2/views/evm/common"

	gethCommon "github.com/ethereum/go-ethereum/common"
)

// packFeeQuoterCall packs a method call using the FeeQuoter v1.6 ABI.
func packFeeQuoterCall(method string, args ...interface{}) ([]byte, error) {
	return FeeQuoterABI.Pack(method, args...)
}

// executeFeeQuoterCall packs a call, executes it, and returns raw response bytes.
func executeFeeQuoterCall(ctx *views.ViewContext, method string, args ...interface{}) ([]byte, error) {
	calldata, err := packFeeQuoterCall(method, args...)
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

// getFeeQuoterOwner fetches the owner address.
func getFeeQuoterOwner(ctx *views.ViewContext) (string, error) {
	data, err := executeFeeQuoterCall(ctx, "owner")
	if err != nil {
		return "", err
	}
	results, err := FeeQuoterABI.Unpack("owner", data)
	if err != nil {
		return "", fmt.Errorf("failed to unpack owner: %w", err)
	}
	if len(results) == 0 {
		return "", fmt.Errorf("no results from owner call")
	}
	owner, ok := results[0].(gethCommon.Address)
	if !ok {
		return "", fmt.Errorf("unexpected type for owner: %T", results[0])
	}
	return owner.Hex(), nil
}

// getFeeQuoterTypeAndVersion fetches the typeAndVersion string.
func getFeeQuoterTypeAndVersion(ctx *views.ViewContext) (string, error) {
	data, err := executeFeeQuoterCall(ctx, "typeAndVersion")
	if err != nil {
		return "", err
	}
	results, err := FeeQuoterABI.Unpack("typeAndVersion", data)
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

// getFeeQuoterAuthorizedCallers fetches the authorized callers using ABI bindings.
func getFeeQuoterAuthorizedCallers(ctx *views.ViewContext) ([]string, error) {
	data, err := executeFeeQuoterCall(ctx, "getAllAuthorizedCallers")
	if err != nil {
		return nil, err
	}

	results, err := FeeQuoterABI.Unpack("getAllAuthorizedCallers", data)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack getAllAuthorizedCallers: %w", err)
	}
	if len(results) == 0 {
		return []string{}, nil
	}

	addrs, ok := results[0].([]gethCommon.Address)
	if !ok {
		return nil, fmt.Errorf("unexpected type for authorized callers: %T", results[0])
	}

	callers := make([]string, len(addrs))
	for i, a := range addrs {
		callers[i] = a.Hex()
	}
	return callers, nil
}

// getFeeQuoterFeeTokens fetches the fee token addresses using ABI bindings.
func getFeeQuoterFeeTokens(ctx *views.ViewContext) ([]string, error) {
	data, err := executeFeeQuoterCall(ctx, "getFeeTokens")
	if err != nil {
		return nil, err
	}

	results, err := FeeQuoterABI.Unpack("getFeeTokens", data)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack getFeeTokens: %w", err)
	}
	if len(results) == 0 {
		return []string{}, nil
	}

	addrs, ok := results[0].([]gethCommon.Address)
	if !ok {
		return nil, fmt.Errorf("unexpected type for fee tokens: %T", results[0])
	}

	tokens := make([]string, len(addrs))
	for i, a := range addrs {
		tokens[i] = a.Hex()
	}
	return tokens, nil
}

// getFeeQuoterPremiumMultiplierWeiPerEth calls FeeQuoter.getPremiumMultiplierWeiPerEth(feeToken).
func getFeeQuoterPremiumMultiplierWeiPerEth(ctx *views.ViewContext, tokenAddrHex string) (string, error) {
	addr := gethCommon.HexToAddress(tokenAddrHex)

	data, err := executeFeeQuoterCall(ctx, "getPremiumMultiplierWeiPerEth", addr)
	if err != nil {
		return "", err
	}

	results, err := FeeQuoterABI.Unpack("getPremiumMultiplierWeiPerEth", data)
	if err != nil {
		return "", fmt.Errorf("failed to unpack getPremiumMultiplierWeiPerEth: %w", err)
	}
	if len(results) == 0 {
		return "0", nil
	}

	premium, ok := results[0].(uint64)
	if !ok {
		return "", fmt.Errorf("unexpected type for premiumMultiplierWeiPerEth: %T", results[0])
	}
	return fmt.Sprintf("%d", premium), nil
}

// getFeeQuoterERC20Decimals fetches decimals from an ERC20 token.
func getFeeQuoterERC20Decimals(ctx *views.ViewContext, tokenAddrHex string) (uint8, error) {
	tokenAddrHex = strings.TrimPrefix(tokenAddrHex, "0x")
	tokenAddr, err := hex.DecodeString(tokenAddrHex)
	if err != nil {
		return 0, err
	}
	if len(tokenAddr) < 20 {
		padded := make([]byte, 20)
		copy(padded[20-len(tokenAddr):], tokenAddr)
		tokenAddr = padded
	}
	calldata := views.ABIEncodeCall(common.HexToSelector("313ce567")) // decimals()
	call := views.Call{ChainID: ctx.ChainSelector, Target: tokenAddr, Data: calldata}
	result := ctx.TypedOrchestrator.Execute(call)
	if result.Error != nil {
		return 0, result.Error
	}
	if len(result.Data) < 32 {
		return 0, nil
	}
	return uint8(result.Data[31]), nil
}

// getFeeQuoterFeeTokensEnriched returns fee tokens with name, symbol, decimals, and premiumMultiplierWeiPerEth.
func getFeeQuoterFeeTokensEnriched(ctx *views.ViewContext) ([]map[string]any, error) {
	addresses, err := getFeeQuoterFeeTokens(ctx)
	if err != nil {
		return nil, err
	}
	out := make([]map[string]any, 0, len(addresses))
	for _, addr := range addresses {
		entry := map[string]any{"address": addr}

		if name, err := common.GetERC20Name(ctx, addr); err != nil {
			entry["name_error"] = err.Error()
		} else {
			entry["name"] = name
		}
		if symbol, err := common.GetERC20Symbol(ctx, addr); err != nil {
			entry["symbol_error"] = err.Error()
		} else {
			entry["symbol"] = symbol
		}
		if decimals, err := getFeeQuoterERC20Decimals(ctx, addr); err != nil {
			entry["decimals_error"] = err.Error()
		} else {
			entry["decimals"] = decimals
		}
		if premium, err := getFeeQuoterPremiumMultiplierWeiPerEth(ctx, addr); err != nil {
			entry["premiumMultiplierWeiPerEth_error"] = err.Error()
		} else {
			entry["premiumMultiplierWeiPerEth"] = premium
		}

		out = append(out, entry)
	}
	return out, nil
}

// getFeeQuoterStaticConfig fetches the static configuration using ABI bindings.
func getFeeQuoterStaticConfig(ctx *views.ViewContext) (map[string]any, error) {
	data, err := executeFeeQuoterCall(ctx, "getStaticConfig")
	if err != nil {
		return nil, err
	}

	results, err := FeeQuoterABI.Unpack("getStaticConfig", data)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack getStaticConfig: %w", err)
	}
	if len(results) == 0 {
		return nil, fmt.Errorf("no results from getStaticConfig call")
	}

	cfg, ok := results[0].(struct {
		MaxFeeJuelsPerMsg            *big.Int           `json:"maxFeeJuelsPerMsg"`
		LinkToken                    gethCommon.Address `json:"linkToken"`
		TokenPriceStalenessThreshold uint32             `json:"tokenPriceStalenessThreshold"`
	})
	if !ok {
		return nil, fmt.Errorf("unexpected type for StaticConfig: %T", results[0])
	}

	return map[string]any{
		"maxFeeJuelsPerMsg":            cfg.MaxFeeJuelsPerMsg.String(),
		"linkToken":                    cfg.LinkToken.Hex(),
		"tokenPriceStalenessThreshold": cfg.TokenPriceStalenessThreshold,
	}, nil
}

// getFQDestChainConfigForChain fetches the FeeQuoter DestChainConfig for a specific chain.
func getFQDestChainConfigForChain(ctx *views.ViewContext, chainSel uint64) (map[string]any, error) {
	data, err := executeFeeQuoterCall(ctx, "getDestChainConfig", chainSel)
	if err != nil {
		return nil, err
	}

	results, err := FeeQuoterABI.Unpack("getDestChainConfig", data)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack getDestChainConfig: %w", err)
	}
	if len(results) == 0 {
		return nil, fmt.Errorf("no results from getDestChainConfig call")
	}

	cfg, ok := results[0].(struct {
		IsEnabled                         bool     `json:"isEnabled"`
		MaxNumberOfTokensPerMsg           uint16   `json:"maxNumberOfTokensPerMsg"`
		MaxDataBytes                      uint32   `json:"maxDataBytes"`
		MaxPerMsgGasLimit                 uint32   `json:"maxPerMsgGasLimit"`
		DestGasOverhead                   uint32   `json:"destGasOverhead"`
		DestGasPerPayloadByteBase         uint8    `json:"destGasPerPayloadByteBase"`
		DestGasPerPayloadByteHigh         uint8    `json:"destGasPerPayloadByteHigh"`
		DestGasPerPayloadByteThreshold    uint16   `json:"destGasPerPayloadByteThreshold"`
		DestDataAvailabilityOverheadGas   uint32   `json:"destDataAvailabilityOverheadGas"`
		DestGasPerDataAvailabilityByte    uint16   `json:"destGasPerDataAvailabilityByte"`
		DestDataAvailabilityMultiplierBps uint16   `json:"destDataAvailabilityMultiplierBps"`
		ChainFamilySelector               [4]byte  `json:"chainFamilySelector"`
		EnforceOutOfOrder                 bool     `json:"enforceOutOfOrder"`
		DefaultTokenFeeUSDCents           uint16   `json:"defaultTokenFeeUSDCents"`
		DefaultTokenDestGasOverhead       uint32   `json:"defaultTokenDestGasOverhead"`
		DefaultTxGasLimit                 uint32   `json:"defaultTxGasLimit"`
		GasMultiplierWeiPerEth            uint64   `json:"gasMultiplierWeiPerEth"`
		GasPriceStalenessThreshold        uint32   `json:"gasPriceStalenessThreshold"`
		NetworkFeeUSDCents                uint32   `json:"networkFeeUSDCents"`
	})
	if !ok {
		return nil, fmt.Errorf("unexpected type for DestChainConfig: %T", results[0])
	}

	return map[string]any{
		"isEnabled":                         cfg.IsEnabled,
		"maxNumberOfTokensPerMsg":           cfg.MaxNumberOfTokensPerMsg,
		"maxDataBytes":                      cfg.MaxDataBytes,
		"maxPerMsgGasLimit":                 cfg.MaxPerMsgGasLimit,
		"destGasOverhead":                   cfg.DestGasOverhead,
		"destGasPerPayloadByteBase":         cfg.DestGasPerPayloadByteBase,
		"destGasPerPayloadByteHigh":         cfg.DestGasPerPayloadByteHigh,
		"destGasPerPayloadByteThreshold":    cfg.DestGasPerPayloadByteThreshold,
		"destDataAvailabilityOverheadGas":   cfg.DestDataAvailabilityOverheadGas,
		"destGasPerDataAvailabilityByte":    cfg.DestGasPerDataAvailabilityByte,
		"destDataAvailabilityMultiplierBps": cfg.DestDataAvailabilityMultiplierBps,
		"chainFamilySelector":               hex.EncodeToString(cfg.ChainFamilySelector[:]),
		"enforceOutOfOrder":                 cfg.EnforceOutOfOrder,
		"defaultTokenFeeUSDCents":           cfg.DefaultTokenFeeUSDCents,
		"defaultTokenDestGasOverhead":       cfg.DefaultTokenDestGasOverhead,
		"defaultTxGasLimit":                 cfg.DefaultTxGasLimit,
		"gasMultiplierWeiPerEth":            cfg.GasMultiplierWeiPerEth,
		"gasPriceStalenessThreshold":        cfg.GasPriceStalenessThreshold,
		"networkFeeUSDCents":                cfg.NetworkFeeUSDCents,
	}, nil
}

// getFeeQuoterDestChainConfigs fetches destination chain configs concurrently.
func getFeeQuoterDestChainConfigs(ctx *views.ViewContext) map[string]any {
	if len(ctx.AllChainSelectors) == 0 {
		return map[string]any{}
	}

	result := make(map[string]any)
	var mu sync.Mutex
	var wg sync.WaitGroup

	for _, chainSel := range ctx.AllChainSelectors {
		wg.Add(1)
		go func(cs uint64) {
			defer wg.Done()

			config, err := getFQDestChainConfigForChain(ctx, cs)
			if err != nil || config == nil {
				return
			}

			if enabled, ok := config["isEnabled"].(bool); ok && enabled {
				mu.Lock()
				result[views.Uint64ToString(cs)] = config
				mu.Unlock()
			}
		}(chainSel)
	}

	wg.Wait()
	return result
}

// ViewFeeQuoter generates a view of the FeeQuoter contract (v1.6.0).
// Uses ABI bindings for proper struct decoding.
func ViewFeeQuoter(ctx *views.ViewContext) (map[string]any, error) {
	result := make(map[string]any)

	result["address"] = ctx.AddressHex
	result["chainSelector"] = ctx.ChainSelector
	result["version"] = "1.6.0"

	owner, err := getFeeQuoterOwner(ctx)
	if err != nil {
		result["owner_error"] = err.Error()
	} else {
		result["owner"] = owner
	}

	typeAndVersion, err := getFeeQuoterTypeAndVersion(ctx)
	if err != nil {
		result["typeAndVersion_error"] = err.Error()
	} else {
		result["typeAndVersion"] = typeAndVersion
	}

	authorizedCallers, err := getFeeQuoterAuthorizedCallers(ctx)
	if err == nil && len(authorizedCallers) > 0 {
		result["authorizedCallers"] = authorizedCallers
	}

	feeTokensEnriched, err := getFeeQuoterFeeTokensEnriched(ctx)
	if err == nil && len(feeTokensEnriched) > 0 {
		result["feeTokens"] = feeTokensEnriched
	}

	staticConfig, err := getFeeQuoterStaticConfig(ctx)
	if err != nil {
		result["staticConfig_error"] = err.Error()
	} else {
		result["staticConfig"] = staticConfig
	}

	destChainConfigs := getFeeQuoterDestChainConfigs(ctx)
	if len(destChainConfigs) > 0 {
		result["destinationChainConfig"] = destChainConfigs
	}

	return result, nil
}
