package v1_6

import (
	"give-me-state-v2/views"
	"give-me-state-v2/views/evm/common"
	"sync"
)

// OffRamp selectors
var (
	// getSourceChainConfig(uint64) returns (SourceChainConfig)
	selectorGetSourceChainConfig = common.HexToSelector("e9d68a8e")
	// getLatestPriceSequenceNumber() returns (uint64)
	selectorGetLatestPriceSeqNum = common.HexToSelector("3f4b04aa")
)

// ViewOffRamp generates a view of the OffRamp contract (v1.6.0).
func ViewOffRamp(ctx *views.ViewContext) (map[string]any, error) {
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

	// Get static config
	staticConfig, err := getOffRampStaticConfig(ctx)
	if err != nil {
		result["staticConfig_error"] = err.Error()
	} else {
		result["staticConfig"] = staticConfig
	}

	// Get dynamic config
	dynamicConfig, err := getOffRampDynamicConfig(ctx)
	if err != nil {
		result["dynamicConfig_error"] = err.Error()
	} else {
		result["dynamicConfig"] = dynamicConfig
	}

	// Get latest price sequence number
	latestPriceSeqNum, err := getOffRampLatestPriceSeqNum(ctx)
	if err == nil {
		result["latestPriceSequenceNumber"] = latestPriceSeqNum
	}

	// Get source chain configs (concurrent)
	sourceChainConfigs, err := getOffRampSourceChainConfigs(ctx)
	if err != nil {
		result["sourceChainConfigs_error"] = err.Error()
	} else if len(sourceChainConfigs) > 0 {
		result["sourceChainConfigs"] = sourceChainConfigs
	}

	return result, nil
}

// getOffRampStaticConfig fetches the static configuration (v1.6.0 style).
func getOffRampStaticConfig(ctx *views.ViewContext) (map[string]any, error) {
	data, err := common.ExecuteCall(ctx, common.SelectorGetStaticConfig)
	if err != nil {
		return nil, err
	}

	config := make(map[string]any)
	config["rawData"] = views.BytesToHex(data)

	// Parse known fields per v1.6.0 ABI:
	// struct StaticConfig {
	//   uint64 chainSelector;
	//   uint16 gasForCallExactCheck;
	//   address rmnRemote;
	//   address tokenAdminRegistry;
	//   address nonceManager;
	// }
	offset := 0
	if len(data) >= offset+32 {
		cs, _ := common.DecodeUint64(data[offset : offset+32])
		config["chainSelector"] = cs
		offset += 32
	}
	if len(data) >= offset+32 {
		gasForCallExactCheck := common.DecodeUint64FromBytes(data[offset : offset+32])
		config["gasForCallExactCheck"] = uint16(gasForCallExactCheck)
		offset += 32
	}
	if len(data) >= offset+32 {
		rmnRemote, _ := common.DecodeAddress(data[offset : offset+32])
		config["rmnRemote"] = rmnRemote
		offset += 32
	}
	if len(data) >= offset+32 {
		tokenAdminRegistry, _ := common.DecodeAddress(data[offset : offset+32])
		config["tokenAdminRegistry"] = tokenAdminRegistry
		offset += 32
	}
	if len(data) >= offset+32 {
		nonceManager, _ := common.DecodeAddress(data[offset : offset+32])
		config["nonceManager"] = nonceManager
	}

	return config, nil
}

// getOffRampDynamicConfig fetches the dynamic configuration (v1.6.0 style).
func getOffRampDynamicConfig(ctx *views.ViewContext) (map[string]any, error) {
	data, err := common.ExecuteCall(ctx, common.SelectorGetDynamicConfig)
	if err != nil {
		return nil, err
	}

	config := make(map[string]any)

	// Parse known fields
	offset := 0
	if len(data) >= offset+32 {
		feeQuoter, _ := common.DecodeAddress(data[offset : offset+32])
		config["feeQuoter"] = feeQuoter
		offset += 32
	}
	if len(data) >= offset+32 {
		permissionlessExec := common.DecodeUint64FromBytes(data[offset : offset+32])
		config["permissionLessExecutionThresholdSeconds"] = permissionlessExec
	}

	return config, nil
}

// getOffRampLatestPriceSeqNum fetches the latest price sequence number.
func getOffRampLatestPriceSeqNum(ctx *views.ViewContext) (uint64, error) {
	data, err := common.ExecuteCall(ctx, selectorGetLatestPriceSeqNum)
	if err != nil {
		return 0, err
	}
	return common.DecodeUint64(data)
}

// getOffRampSourceChainConfigs fetches source chain configs concurrently.
// In v1.6, we use AllChainSelectors from ctx since getAllSourceChainConfigs() may not exist.
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
				return // Chain not configured or error
			}

			// Check if router is set and isEnabled (indicates chain is configured)
			isEnabled, ok := config["isEnabled"].(bool)
			if !ok || !isEnabled {
				return // Not enabled
			}

			mu.Lock()
			result[views.Uint64ToString(cs)] = config
			mu.Unlock()
		}(chainSel)
	}

	wg.Wait()
	return result, nil
}

// getSourceChainConfigForChain fetches the SourceChainConfig for a specific chain.
func getSourceChainConfigForChain(ctx *views.ViewContext, chainSel uint64) (map[string]any, error) {
	data, err := common.ExecuteCall(ctx, selectorGetSourceChainConfig, common.EncodeUint64(chainSel))
	if err != nil {
		return nil, err
	}

	config := make(map[string]any)

	// SourceChainConfig: address router, bool isEnabled, uint64 minSeqNr, bool isRMNVerificationDisabled, bytes onRamp
	offset := 0
	if len(data) >= offset+32 {
		router, _ := common.DecodeAddress(data[offset : offset+32])
		config["router"] = router
		offset += 32
	}
	if len(data) >= offset+32 {
		isEnabled, _ := common.DecodeBool(data[offset : offset+32])
		config["isEnabled"] = isEnabled
		offset += 32
	}
	if len(data) >= offset+32 {
		minSeqNr := common.DecodeUint64FromBytes(data[offset : offset+32])
		config["minSeqNr"] = minSeqNr
		offset += 32
	}
	if len(data) >= offset+32 {
		isRMNDisabled, _ := common.DecodeBool(data[offset : offset+32])
		config["isRMNVerificationDisabled"] = isRMNDisabled
		offset += 32
	}
	// onRamp is a dynamic bytes field - decode it
	if len(data) >= offset+32 {
		// This is an offset to the bytes data
		bytesOffset := common.DecodeUint64FromBytes(data[offset : offset+32])
		if bytesOffset > 0 && int(bytesOffset+32) <= len(data) {
			bytesLen := common.DecodeUint64FromBytes(data[bytesOffset : bytesOffset+32])
			if bytesLen > 0 && int(bytesOffset+32+bytesLen) <= len(data) {
				onRampBytes := data[bytesOffset+32 : bytesOffset+32+bytesLen]
				// Format based on length - if 20 bytes, it's an EVM address, if 32 it might be Solana
				if len(onRampBytes) == 20 {
					config["onRamp"] = "0x" + views.BytesToHex(onRampBytes)[2:]
				} else {
					config["onRamp"] = views.BytesToHex(onRampBytes)
				}
			}
		}
	}

	return config, nil
}
