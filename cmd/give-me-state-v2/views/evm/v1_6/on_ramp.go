package v1_6

import (
	"fmt"
	"sync"

	"give-me-state-v2/views"

	"github.com/ethereum/go-ethereum/common"
)

// packOnRampCall packs a method call using the OnRamp v1.6 ABI.
func packOnRampCall(method string, args ...interface{}) ([]byte, error) {
	return OnRampABI.Pack(method, args...)
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

// getOnRampOwner fetches the owner address.
func getOnRampOwner(ctx *views.ViewContext) (string, error) {
	data, err := executeOnRampCall(ctx, "owner")
	if err != nil {
		return "", err
	}

	results, err := OnRampABI.Unpack("owner", data)
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

// getOnRampTypeAndVersion fetches the typeAndVersion string.
func getOnRampTypeAndVersion(ctx *views.ViewContext) (string, error) {
	data, err := executeOnRampCall(ctx, "typeAndVersion")
	if err != nil {
		return "", err
	}

	results, err := OnRampABI.Unpack("typeAndVersion", data)
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

// getOnRampStaticConfig fetches the static configuration using ABI bindings.
func getOnRampStaticConfig(ctx *views.ViewContext) (map[string]any, error) {
	data, err := executeOnRampCall(ctx, "getStaticConfig")
	if err != nil {
		return nil, err
	}

	results, err := OnRampABI.Unpack("getStaticConfig", data)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack getStaticConfig: %w", err)
	}
	if len(results) == 0 {
		return nil, fmt.Errorf("no results from getStaticConfig call")
	}

	// The result is a struct: OnRampStaticConfig{ChainSelector, RmnRemote, NonceManager, TokenAdminRegistry}
	cfg, ok := results[0].(struct {
		ChainSelector      uint64         `json:"chainSelector"`
		RmnRemote          common.Address `json:"rmnRemote"`
		NonceManager       common.Address `json:"nonceManager"`
		TokenAdminRegistry common.Address `json:"tokenAdminRegistry"`
	})
	if !ok {
		return nil, fmt.Errorf("unexpected type for StaticConfig: %T", results[0])
	}

	return map[string]any{
		"chainSelector":      cfg.ChainSelector,
		"rmnRemote":          cfg.RmnRemote.Hex(),
		"nonceManager":       cfg.NonceManager.Hex(),
		"tokenAdminRegistry": cfg.TokenAdminRegistry.Hex(),
	}, nil
}

// getOnRampDynamicConfig fetches the dynamic configuration using ABI bindings.
func getOnRampDynamicConfig(ctx *views.ViewContext) (map[string]any, error) {
	data, err := executeOnRampCall(ctx, "getDynamicConfig")
	if err != nil {
		return nil, err
	}

	results, err := OnRampABI.Unpack("getDynamicConfig", data)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack getDynamicConfig: %w", err)
	}
	if len(results) == 0 {
		return nil, fmt.Errorf("no results from getDynamicConfig call")
	}

	cfg, ok := results[0].(struct {
		FeeQuoter              common.Address `json:"feeQuoter"`
		ReentrancyGuardEntered bool           `json:"reentrancyGuardEntered"`
		MessageInterceptor     common.Address `json:"messageInterceptor"`
		FeeAggregator          common.Address `json:"feeAggregator"`
		AllowlistAdmin         common.Address `json:"allowlistAdmin"`
	})
	if !ok {
		return nil, fmt.Errorf("unexpected type for DynamicConfig: %T", results[0])
	}

	return map[string]any{
		"feeQuoter":              cfg.FeeQuoter.Hex(),
		"reentrancyGuardEntered": cfg.ReentrancyGuardEntered,
		"messageInterceptor":     cfg.MessageInterceptor.Hex(),
		"feeAggregator":          cfg.FeeAggregator.Hex(),
		"allowlistAdmin":         cfg.AllowlistAdmin.Hex(),
	}, nil
}

// getDestChainConfigForChain fetches the DestChainConfig for a specific chain using ABI bindings.
func getDestChainConfigForChain(ctx *views.ViewContext, chainSel uint64) (map[string]any, error) {
	data, err := executeOnRampCall(ctx, "getDestChainConfig", chainSel)
	if err != nil {
		return nil, err
	}

	results, err := OnRampABI.Unpack("getDestChainConfig", data)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack getDestChainConfig: %w", err)
	}
	if len(results) == 0 {
		return nil, fmt.Errorf("no results from getDestChainConfig call")
	}

	cfg, ok := results[0].(struct {
		SequenceNumber   uint64         `json:"sequenceNumber"`
		AllowlistEnabled bool           `json:"allowlistEnabled"`
		Router           common.Address `json:"router"`
	})
	if !ok {
		return nil, fmt.Errorf("unexpected type for DestChainConfig: %T", results[0])
	}

	return map[string]any{
		"sequenceNumber":   cfg.SequenceNumber,
		"allowlistEnabled": cfg.AllowlistEnabled,
		"router":           cfg.Router.Hex(),
	}, nil
}

// getExpectedNextSeqNumForChain fetches the expected next sequence number for a chain.
func getExpectedNextSeqNumForChain(ctx *views.ViewContext, chainSel uint64) (uint64, error) {
	data, err := executeOnRampCall(ctx, "getExpectedNextSequenceNumber", chainSel)
	if err != nil {
		return 0, err
	}

	results, err := OnRampABI.Unpack("getExpectedNextSequenceNumber", data)
	if err != nil {
		return 0, fmt.Errorf("failed to unpack getExpectedNextSequenceNumber: %w", err)
	}
	if len(results) == 0 {
		return 0, fmt.Errorf("no results from getExpectedNextSequenceNumber call")
	}
	seqNum, ok := results[0].(uint64)
	if !ok {
		return 0, fmt.Errorf("unexpected type for sequence number: %T", results[0])
	}
	return seqNum, nil
}

// getAllowedSendersForChain fetches the allowed senders list for a chain.
func getAllowedSendersForChain(ctx *views.ViewContext, chainSel uint64) ([]string, error) {
	data, err := executeOnRampCall(ctx, "getAllowedSendersList", chainSel)
	if err != nil {
		return nil, err
	}

	results, err := OnRampABI.Unpack("getAllowedSendersList", data)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack getAllowedSendersList: %w", err)
	}
	if len(results) == 0 {
		return []string{}, nil
	}
	addrs, ok := results[0].([]common.Address)
	if !ok {
		return nil, fmt.Errorf("unexpected type for allowed senders: %T", results[0])
	}
	senders := make([]string, len(addrs))
	for i, a := range addrs {
		senders[i] = a.Hex()
	}
	return senders, nil
}

// getOnRampDestChainData fetches destination chain specific data concurrently.
func getOnRampDestChainData(ctx *views.ViewContext) (map[string]any, error) {
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

			destConfig, err := getDestChainConfigForChain(ctx, cs)
			if err != nil {
				return
			}

			router, ok := destConfig["router"].(string)
			if !ok || router == "" || router == "0x0000000000000000000000000000000000000000" {
				return
			}

			chainData := make(map[string]any)
			chainData["destChainConfig"] = destConfig

			nextSeq, err := getExpectedNextSeqNumForChain(ctx, cs)
			if err == nil {
				chainData["expectedNextSeqNum"] = nextSeq
			}

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

// ViewOnRamp generates a view of the OnRamp contract (v1.6.0).
// Uses ABI bindings for proper struct decoding.
func ViewOnRamp(ctx *views.ViewContext) (map[string]any, error) {
	result := make(map[string]any)

	result["address"] = ctx.AddressHex
	result["chainSelector"] = ctx.ChainSelector
	result["version"] = "1.6.0"

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

	staticConfig, err := getOnRampStaticConfig(ctx)
	if err != nil {
		result["staticConfig_error"] = err.Error()
	} else {
		result["staticConfig"] = staticConfig
	}

	dynamicConfig, err := getOnRampDynamicConfig(ctx)
	if err != nil {
		result["dynamicConfig_error"] = err.Error()
	} else {
		result["dynamicConfig"] = dynamicConfig
	}

	destChainData, err := getOnRampDestChainData(ctx)
	if err != nil {
		result["destChainSpecificData_error"] = err.Error()
	} else if len(destChainData) > 0 {
		result["destChainSpecificData"] = destChainData
	}

	return result, nil
}
