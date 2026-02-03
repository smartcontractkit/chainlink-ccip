package v1_6

import (
	"call-orchestrator-demo/views"
	"call-orchestrator-demo/views/evm/common"
	"sync"
)

// OnRamp selectors
var (
	// getDestChainConfig(uint64) returns (DestChainConfig)
	selectorGetDestChainConfig = common.HexToSelector("6def4ce7")
	// getExpectedNextSequenceNumber(uint64) returns (uint64)
	selectorGetExpectedNextSeqNum = common.HexToSelector("9041be3d")
	// getAllowedSendersList(uint64) returns (AllowedSendersInfo)
	selectorGetAllowedSendersList = common.HexToSelector("972b4612")
)


// ViewOnRamp generates a view of the OnRamp contract (v1.6.0).
func ViewOnRamp(ctx *views.ViewContext) (map[string]any, error) {
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
	staticConfig, err := getOnRampStaticConfig(ctx)
	if err != nil {
		result["staticConfig_error"] = err.Error()
	} else {
		result["staticConfig"] = staticConfig
	}

	// Get dynamic config
	dynamicConfig, err := getOnRampDynamicConfig(ctx)
	if err != nil {
		result["dynamicConfig_error"] = err.Error()
	} else {
		result["dynamicConfig"] = dynamicConfig
	}

	// Get dest chain specific data (concurrent)
	destChainData, err := getOnRampDestChainData(ctx)
	if err != nil {
		result["destChainSpecificData_error"] = err.Error()
	} else if len(destChainData) > 0 {
		result["destChainSpecificData"] = destChainData
	}

	return result, nil
}

// getOnRampStaticConfig fetches the static configuration (v1.6.0 style).
func getOnRampStaticConfig(ctx *views.ViewContext) (map[string]any, error) {
	data, err := common.ExecuteCall(ctx, common.SelectorGetStaticConfig)
	if err != nil {
		return nil, err
	}

	config := make(map[string]any)
	config["rawData"] = views.BytesToHex(data)

	// Parse known fields (structure varies by version)
	offset := 0
	if len(data) >= offset+32 {
		cs, _ := common.DecodeUint64(data[offset : offset+32])
		config["chainSelector"] = cs
		offset += 32
	}
	if len(data) >= offset+32 {
		rmn, _ := common.DecodeAddress(data[offset : offset+32])
		config["rmn"] = rmn
		offset += 32
	}
	if len(data) >= offset+32 {
		nonceManager, _ := common.DecodeAddress(data[offset : offset+32])
		config["nonceManager"] = nonceManager
		offset += 32
	}
	if len(data) >= offset+32 {
		tokenAdminRegistry, _ := common.DecodeAddress(data[offset : offset+32])
		config["tokenAdminRegistry"] = tokenAdminRegistry
	}

	return config, nil
}

// getOnRampDynamicConfig fetches the dynamic configuration (v1.6.0 style).
func getOnRampDynamicConfig(ctx *views.ViewContext) (map[string]any, error) {
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
		messageValidator, _ := common.DecodeAddress(data[offset : offset+32])
		config["messageValidator"] = messageValidator
		offset += 32
	}
	if len(data) >= offset+32 {
		feeAggregator, _ := common.DecodeAddress(data[offset : offset+32])
		config["feeAggregator"] = feeAggregator
	}

	return config, nil
}

// getOnRampDestChainData fetches destination chain specific data concurrently.
// In v1.6, we use AllChainSelectors from ctx since getAllDestChainConfigs() doesn't exist.
func getOnRampDestChainData(ctx *views.ViewContext) (map[string]any, error) {
	// Use chain selectors from deployment to probe for configured destinations
	// A chain is configured if getDestChainConfig returns valid data with a non-zero router

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

			// Get dest chain config first to check if this chain is configured
			destConfig, err := getDestChainConfigForChain(ctx, cs)
			if err != nil {
				return // Chain not configured or error
			}

			// Check if router is set (indicates chain is configured)
			router, ok := destConfig["router"].(string)
			if !ok || router == "" || router == "0x0000000000000000000000000000000000000000" {
				return // Not configured
			}

			chainData := make(map[string]any)
			chainData["destChainConfig"] = destConfig

			// Get expected next sequence number
			nextSeq, err := getExpectedNextSeqNumForChain(ctx, cs)
			if err == nil {
				chainData["expectedNextSeqNum"] = nextSeq
			}

			// Get allowed senders list
			senders, err := getAllowedSendersForChain(ctx, cs)
			if err == nil {
				chainData["allowedSendersList"] = senders
			}

			mu.Lock()
			result[views.Uint64ToString(cs)] = chainData
			mu.Unlock()
		}(chainSel)
	}

	wg.Wait()
	return result, nil
}

// getDestChainConfigForChain fetches the DestChainConfig for a specific chain.
func getDestChainConfigForChain(ctx *views.ViewContext, chainSel uint64) (map[string]any, error) {
	data, err := common.ExecuteCall(ctx, selectorGetDestChainConfig, common.EncodeUint64(chainSel))
	if err != nil {
		return nil, err
	}

	config := make(map[string]any)

	// DestChainConfig: uint64 sequenceNumber, bool allowlistEnabled, address router
	offset := 0
	if len(data) >= offset+32 {
		seqNum := common.DecodeUint64FromBytes(data[offset : offset+32])
		config["sequenceNumber"] = seqNum
		offset += 32
	}
	if len(data) >= offset+32 {
		allowlistEnabled, _ := common.DecodeBool(data[offset : offset+32])
		config["allowlistEnabled"] = allowlistEnabled
		offset += 32
	}
	if len(data) >= offset+32 {
		router, _ := common.DecodeAddress(data[offset : offset+32])
		config["router"] = router
	}

	return config, nil
}

// getExpectedNextSeqNumForChain fetches the expected next sequence number for a chain.
func getExpectedNextSeqNumForChain(ctx *views.ViewContext, chainSel uint64) (uint64, error) {
	data, err := common.ExecuteCall(ctx, selectorGetExpectedNextSeqNum, common.EncodeUint64(chainSel))
	if err != nil {
		return 0, err
	}
	return common.DecodeUint64(data)
}

// getAllowedSendersForChain fetches the allowed senders list for a chain.
func getAllowedSendersForChain(ctx *views.ViewContext, chainSel uint64) ([]string, error) {
	data, err := common.ExecuteCall(ctx, selectorGetAllowedSendersList, common.EncodeUint64(chainSel))
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

	senders := make([]string, 0, length)
	for i := uint64(0); i < length; i++ {
		offset := 64 + i*32
		if offset+32 > uint64(len(data)) {
			break
		}
		addr, _ := common.DecodeAddress(data[offset : offset+32])
		senders = append(senders, addr)
	}

	return senders, nil
}
