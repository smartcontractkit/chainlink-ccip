package v1_6

import (
	"call-orchestrator-demo/views"
	"call-orchestrator-demo/views/evm/common"
	"math/big"
	"sync"
)

// FeeQuoter selectors
var (
	// getAllAuthorizedCallers() returns (address[])
	selectorFQGetAllAuthorizedCallers = common.HexToSelector("2451a627")
	// getFeeTokens() returns (address[])
	selectorFQGetFeeTokens = common.HexToSelector("cdc73d51")
	// getDestChainConfig(uint64) returns (DestChainConfig)
	selectorFQGetDestChainConfig = common.HexToSelector("6def4ce7")
)

// ViewFeeQuoter generates a view of the FeeQuoter contract (v1.6.0).
func ViewFeeQuoter(ctx *views.ViewContext) (map[string]any, error) {
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

	// Get authorized callers
	authorizedCallers, err := getFeeQuoterAuthorizedCallers(ctx)
	if err == nil && len(authorizedCallers) > 0 {
		result["authorizedCallers"] = authorizedCallers
	}

	// Get fee tokens
	feeTokens, err := getFeeQuoterFeeTokens(ctx)
	if err == nil && len(feeTokens) > 0 {
		result["feeTokens"] = feeTokens
	}

	// Get static config
	staticConfig, err := getFeeQuoterStaticConfig(ctx)
	if err != nil {
		result["staticConfig_error"] = err.Error()
	} else {
		result["staticConfig"] = staticConfig
	}

	// Get destination chain configs - we need to discover which chains are configured
	// This is expensive, so we'll get it from OnRamp dest chains if available
	// For now, we'll try to get configs for common chain selectors
	destChainConfigs := getFeeQuoterDestChainConfigs(ctx)
	if len(destChainConfigs) > 0 {
		result["destinationChainConfig"] = destChainConfigs
	}

	return result, nil
}

// getFeeQuoterAuthorizedCallers fetches the authorized callers.
func getFeeQuoterAuthorizedCallers(ctx *views.ViewContext) ([]string, error) {
	data, err := common.ExecuteCall(ctx, selectorFQGetAllAuthorizedCallers)
	if err != nil {
		return nil, err
	}

	// Dynamic array: offset (32) + length (32) + addresses
	if len(data) < 64 {
		return []string{}, nil
	}

	length := common.DecodeUint64FromBytes(data[32:64])
	if length == 0 {
		return []string{}, nil
	}

	callers := make([]string, 0, length)
	for i := uint64(0); i < length; i++ {
		offset := 64 + i*32
		if offset+32 > uint64(len(data)) {
			break
		}
		addr, _ := common.DecodeAddress(data[offset : offset+32])
		callers = append(callers, addr)
	}

	return callers, nil
}

// getFeeQuoterFeeTokens fetches the fee tokens.
func getFeeQuoterFeeTokens(ctx *views.ViewContext) ([]string, error) {
	data, err := common.ExecuteCall(ctx, selectorFQGetFeeTokens)
	if err != nil {
		return nil, err
	}

	// Dynamic array: offset (32) + length (32) + addresses
	if len(data) < 64 {
		return []string{}, nil
	}

	length := common.DecodeUint64FromBytes(data[32:64])
	if length == 0 {
		return []string{}, nil
	}

	tokens := make([]string, 0, length)
	for i := uint64(0); i < length; i++ {
		offset := 64 + i*32
		if offset+32 > uint64(len(data)) {
			break
		}
		addr, _ := common.DecodeAddress(data[offset : offset+32])
		tokens = append(tokens, addr)
	}

	return tokens, nil
}

// getFeeQuoterStaticConfig fetches the static configuration.
func getFeeQuoterStaticConfig(ctx *views.ViewContext) (map[string]any, error) {
	data, err := common.ExecuteCall(ctx, common.SelectorGetStaticConfig)
	if err != nil {
		return nil, err
	}

	config := make(map[string]any)

	if len(data) >= 32 {
		maxFee := new(big.Int).SetBytes(data[0:32])
		config["maxFeeJuelsPerMsg"] = maxFee.String()
	}
	if len(data) >= 64 {
		linkToken, _ := common.DecodeAddress(data[32:64])
		config["linkToken"] = linkToken
	}
	if len(data) >= 96 {
		staleness := common.DecodeUint64FromBytes(data[64:96])
		config["tokenPriceStalenessThreshold"] = staleness
	}

	return config, nil
}

// getFeeQuoterDestChainConfigs fetches destination chain configs concurrently.
// Uses AllChainSelectors from ctx to discover configurations.
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

			// Check if config is enabled (indicates this chain is configured)
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

// getFQDestChainConfigForChain fetches the FeeQuoter DestChainConfig for a specific chain.
func getFQDestChainConfigForChain(ctx *views.ViewContext, chainSel uint64) (map[string]any, error) {
	data, err := common.ExecuteCall(ctx, selectorFQGetDestChainConfig, common.EncodeUint64(chainSel))
	if err != nil {
		return nil, err
	}

	if len(data) < 32 {
		return nil, nil
	}

	config := make(map[string]any)

	// DestChainConfig structure (approximate - field order may vary):
	// bool isEnabled, uint16 maxNumberOfTokensPerMsg, uint32 maxDataBytes, uint32 maxPerMsgGasLimit,
	// uint32 destGasOverhead, uint16 destGasPerPayloadByteBase, uint16 destGasPerPayloadByteHigh,
	// uint16 destGasPerPayloadByteThreshold, uint32 destDataAvailabilityOverheadGas,
	// uint16 destGasPerDataAvailabilityByte, uint16 destDataAvailabilityMultiplierBps,
	// uint16 defaultTokenFeeUSDCents, uint32 defaultTokenDestGasOverhead, uint32 defaultTxGasLimit,
	// uint64 gasMultiplierWeiPerEth, uint32 networkFeeUSDCents, bool enforceOutOfOrder, bytes4 chainFamilySelector

	offset := 0

	if len(data) >= offset+32 {
		isEnabled, _ := common.DecodeBool(data[offset : offset+32])
		config["isEnabled"] = isEnabled
		offset += 32
	}
	if len(data) >= offset+32 {
		maxTokens := common.DecodeUint64FromBytes(data[offset : offset+32])
		config["maxNumberOfTokensPerMsg"] = maxTokens
		offset += 32
	}
	if len(data) >= offset+32 {
		maxDataBytes := common.DecodeUint64FromBytes(data[offset : offset+32])
		config["maxDataBytes"] = maxDataBytes
		offset += 32
	}
	if len(data) >= offset+32 {
		maxGas := common.DecodeUint64FromBytes(data[offset : offset+32])
		config["maxPerMsgGasLimit"] = maxGas
		offset += 32
	}
	if len(data) >= offset+32 {
		destGasOverhead := common.DecodeUint64FromBytes(data[offset : offset+32])
		config["destGasOverhead"] = destGasOverhead
		offset += 32
	}
	if len(data) >= offset+32 {
		config["destGasPerPayloadByteBase"] = common.DecodeUint64FromBytes(data[offset : offset+32])
		offset += 32
	}
	if len(data) >= offset+32 {
		config["destGasPerPayloadByteHigh"] = common.DecodeUint64FromBytes(data[offset : offset+32])
		offset += 32
	}
	if len(data) >= offset+32 {
		config["destGasPerPayloadByteThreshold"] = common.DecodeUint64FromBytes(data[offset : offset+32])
		offset += 32
	}
	if len(data) >= offset+32 {
		config["destDataAvailabilityOverheadGas"] = common.DecodeUint64FromBytes(data[offset : offset+32])
		offset += 32
	}
	if len(data) >= offset+32 {
		config["destGasPerDataAvailabilityByte"] = common.DecodeUint64FromBytes(data[offset : offset+32])
		offset += 32
	}
	if len(data) >= offset+32 {
		config["destDataAvailabilityMultiplierBps"] = common.DecodeUint64FromBytes(data[offset : offset+32])
		offset += 32
	}
	if len(data) >= offset+32 {
		config["defaultTokenFeeUSDCents"] = common.DecodeUint64FromBytes(data[offset : offset+32])
		offset += 32
	}
	if len(data) >= offset+32 {
		config["defaultTokenDestGasOverhead"] = common.DecodeUint64FromBytes(data[offset : offset+32])
		offset += 32
	}
	if len(data) >= offset+32 {
		config["defaultTxGasLimit"] = common.DecodeUint64FromBytes(data[offset : offset+32])
		offset += 32
	}
	if len(data) >= offset+32 {
		gasMultiplier := new(big.Int).SetBytes(data[offset : offset+32])
		config["gasMultiplierWeiPerEth"] = gasMultiplier.Uint64()
		offset += 32
	}
	if len(data) >= offset+32 {
		config["networkFeeUSDCents"] = common.DecodeUint64FromBytes(data[offset : offset+32])
		offset += 32
	}
	if len(data) >= offset+32 {
		enforceOOO, _ := common.DecodeBool(data[offset : offset+32])
		config["enforceOutOfOrder"] = enforceOOO
		offset += 32
	}
	if len(data) >= offset+32 {
		// chainFamilySelector is bytes4 in the last 4 bytes of a 32-byte slot
		selector := data[offset+28 : offset+32]
		config["chainFamilySelector"] = views.BytesToHex(selector)[2:] // Remove 0x prefix
	}

	return config, nil
}
