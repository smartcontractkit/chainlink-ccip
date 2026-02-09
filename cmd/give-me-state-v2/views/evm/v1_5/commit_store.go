package v1_5

import (
	"encoding/hex"
	"fmt"
	"math/big"

	"give-me-state-v2/views"

	"github.com/ethereum/go-ethereum/common"
)

// packCommitStoreCall packs a method call using the CommitStore v1.5 ABI.
func packCommitStoreCall(method string, args ...interface{}) ([]byte, error) {
	return CommitStoreABI.Pack(method, args...)
}

// executeCommitStoreCall packs a call, executes it, and returns raw response bytes.
func executeCommitStoreCall(ctx *views.ViewContext, method string, args ...interface{}) ([]byte, error) {
	calldata, err := packCommitStoreCall(method, args...)
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

// getCommitStoreOwner fetches the owner address.
func getCommitStoreOwner(ctx *views.ViewContext) (string, error) {
	data, err := executeCommitStoreCall(ctx, "owner")
	if err != nil {
		return "", err
	}
	results, err := CommitStoreABI.Unpack("owner", data)
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

// getCommitStoreTypeAndVersion fetches the typeAndVersion string.
func getCommitStoreTypeAndVersion(ctx *views.ViewContext) (string, error) {
	data, err := executeCommitStoreCall(ctx, "typeAndVersion")
	if err != nil {
		return "", err
	}
	results, err := CommitStoreABI.Unpack("typeAndVersion", data)
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

// getCommitStoreStaticConfig fetches and decodes the static config.
func getCommitStoreStaticConfig(ctx *views.ViewContext) (map[string]any, error) {
	data, err := executeCommitStoreCall(ctx, "getStaticConfig")
	if err != nil {
		return nil, err
	}

	results, err := CommitStoreABI.Unpack("getStaticConfig", data)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack getStaticConfig: %w", err)
	}
	if len(results) == 0 {
		return nil, fmt.Errorf("no results from getStaticConfig call")
	}

	cfg, ok := results[0].(struct {
		ChainSelector       uint64         `json:"chainSelector"`
		SourceChainSelector uint64         `json:"sourceChainSelector"`
		OnRamp              common.Address `json:"onRamp"`
		RmnProxy            common.Address `json:"rmnProxy"`
	})
	if !ok {
		return nil, fmt.Errorf("unexpected type for StaticConfig: %T", results[0])
	}

	return map[string]any{
		"chainSelector":       cfg.ChainSelector,
		"sourceChainSelector": cfg.SourceChainSelector,
		"onRamp":              cfg.OnRamp.Hex(),
		"rmnProxy":            cfg.RmnProxy.Hex(),
	}, nil
}

// getCommitStoreDynamicConfig fetches and decodes the dynamic config.
func getCommitStoreDynamicConfig(ctx *views.ViewContext) (map[string]any, error) {
	data, err := executeCommitStoreCall(ctx, "getDynamicConfig")
	if err != nil {
		return nil, err
	}

	results, err := CommitStoreABI.Unpack("getDynamicConfig", data)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack getDynamicConfig: %w", err)
	}
	if len(results) == 0 {
		return nil, fmt.Errorf("no results from getDynamicConfig call")
	}

	cfg, ok := results[0].(struct {
		PriceRegistry common.Address `json:"priceRegistry"`
	})
	if !ok {
		return nil, fmt.Errorf("unexpected type for DynamicConfig: %T", results[0])
	}

	return map[string]any{
		"priceRegistry": cfg.PriceRegistry.Hex(),
	}, nil
}

// getExpectedNextSequenceNumber fetches the expected next sequence number.
func getExpectedNextSequenceNumber(ctx *views.ViewContext) (uint64, error) {
	data, err := executeCommitStoreCall(ctx, "getExpectedNextSequenceNumber")
	if err != nil {
		return 0, err
	}
	results, err := CommitStoreABI.Unpack("getExpectedNextSequenceNumber", data)
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

// getLatestPriceEpochAndRound fetches the latest price epoch and round.
func getLatestPriceEpochAndRound(ctx *views.ViewContext) (uint64, error) {
	data, err := executeCommitStoreCall(ctx, "getLatestPriceEpochAndRound")
	if err != nil {
		return 0, err
	}
	results, err := CommitStoreABI.Unpack("getLatestPriceEpochAndRound", data)
	if err != nil {
		return 0, fmt.Errorf("failed to unpack getLatestPriceEpochAndRound: %w", err)
	}
	if len(results) == 0 {
		return 0, fmt.Errorf("no results from getLatestPriceEpochAndRound call")
	}
	val, ok := results[0].(uint64)
	if !ok {
		return 0, fmt.Errorf("unexpected type for epoch/round: %T", results[0])
	}
	return val, nil
}

// getTransmitters fetches the list of transmitter addresses.
func getTransmitters(ctx *views.ViewContext) ([]string, error) {
	data, err := executeCommitStoreCall(ctx, "getTransmitters")
	if err != nil {
		return nil, err
	}
	results, err := CommitStoreABI.Unpack("getTransmitters", data)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack getTransmitters: %w", err)
	}
	if len(results) == 0 {
		return []string{}, nil
	}
	addrs, ok := results[0].([]common.Address)
	if !ok {
		return nil, fmt.Errorf("unexpected type for transmitters: %T", results[0])
	}
	transmitters := make([]string, len(addrs))
	for i, a := range addrs {
		transmitters[i] = a.Hex()
	}
	return transmitters, nil
}

// getIsUnpausedAndNotCursed fetches the isUnpausedAndNotCursed flag.
func getIsUnpausedAndNotCursed(ctx *views.ViewContext) (bool, error) {
	data, err := executeCommitStoreCall(ctx, "isUnpausedAndNotCursed")
	if err != nil {
		return false, err
	}
	results, err := CommitStoreABI.Unpack("isUnpausedAndNotCursed", data)
	if err != nil {
		return false, fmt.Errorf("failed to unpack isUnpausedAndNotCursed: %w", err)
	}
	if len(results) == 0 {
		return false, fmt.Errorf("no results from isUnpausedAndNotCursed call")
	}
	val, ok := results[0].(bool)
	if !ok {
		return false, fmt.Errorf("unexpected type for isUnpausedAndNotCursed: %T", results[0])
	}
	return val, nil
}

// getLatestConfigDetails fetches the latest config details.
func getLatestConfigDetails(ctx *views.ViewContext) (map[string]any, error) {
	data, err := executeCommitStoreCall(ctx, "latestConfigDetails")
	if err != nil {
		return nil, err
	}
	results, err := CommitStoreABI.Unpack("latestConfigDetails", data)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack latestConfigDetails: %w", err)
	}
	if len(results) < 3 {
		return nil, fmt.Errorf("expected 3 results from latestConfigDetails, got %d", len(results))
	}

	configCount, _ := results[0].(uint32)
	blockNumber, _ := results[1].(uint32)
	configDigest, _ := results[2].([32]byte)

	return map[string]any{
		"configCount":  configCount,
		"blockNumber":  blockNumber,
		"configDigest": "0x" + hex.EncodeToString(configDigest[:]),
	}, nil
}

// getLatestConfigDigestAndEpoch fetches the latest config digest and epoch.
func getLatestConfigDigestAndEpoch(ctx *views.ViewContext) (map[string]any, error) {
	data, err := executeCommitStoreCall(ctx, "latestConfigDigestAndEpoch")
	if err != nil {
		return nil, err
	}
	results, err := CommitStoreABI.Unpack("latestConfigDigestAndEpoch", data)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack latestConfigDigestAndEpoch: %w", err)
	}
	// Returns (bool scanLogs, bytes32 configDigest, uint32 epoch)
	if len(results) < 3 {
		return nil, fmt.Errorf("expected 3 results from latestConfigDigestAndEpoch, got %d", len(results))
	}

	scanLogs, _ := results[0].(bool)
	configDigest, _ := results[1].([32]byte)
	// epoch could be uint32 or uint64 depending on the binding
	var epoch uint64
	switch v := results[2].(type) {
	case uint32:
		epoch = uint64(v)
	case uint64:
		epoch = v
	case *big.Int:
		epoch = v.Uint64()
	}

	return map[string]any{
		"scanLogs":     scanLogs,
		"configDigest": "0x" + hex.EncodeToString(configDigest[:]),
		"epoch":        epoch,
	}, nil
}

// getPaused fetches the paused status.
func getPaused(ctx *views.ViewContext) (bool, error) {
	data, err := executeCommitStoreCall(ctx, "paused")
	if err != nil {
		return false, err
	}
	results, err := CommitStoreABI.Unpack("paused", data)
	if err != nil {
		return false, fmt.Errorf("failed to unpack paused: %w", err)
	}
	if len(results) == 0 {
		return false, fmt.Errorf("no results from paused call")
	}
	val, ok := results[0].(bool)
	if !ok {
		return false, fmt.Errorf("unexpected type for paused: %T", results[0])
	}
	return val, nil
}

// ViewCommitStore generates a view of the CommitStore contract (v1.5.0).
// Uses ABI bindings for proper struct decoding.
func ViewCommitStore(ctx *views.ViewContext) (map[string]any, error) {
	result := make(map[string]any)

	result["address"] = ctx.AddressHex
	result["chainSelector"] = ctx.ChainSelector
	result["version"] = "1.5.0"

	owner, err := getCommitStoreOwner(ctx)
	if err != nil {
		result["owner_error"] = err.Error()
	} else {
		result["owner"] = owner
	}

	typeAndVersion, err := getCommitStoreTypeAndVersion(ctx)
	if err != nil {
		result["typeAndVersion_error"] = err.Error()
	} else {
		result["typeAndVersion"] = typeAndVersion
	}

	staticConfig, err := getCommitStoreStaticConfig(ctx)
	if err != nil {
		result["staticConfig_error"] = err.Error()
	} else {
		result["staticConfig"] = staticConfig
	}

	dynamicConfig, err := getCommitStoreDynamicConfig(ctx)
	if err != nil {
		result["dynamicConfig_error"] = err.Error()
	} else {
		result["dynamicConfig"] = dynamicConfig
	}

	nextSeqNum, err := getExpectedNextSequenceNumber(ctx)
	if err != nil {
		result["expectedNextSequenceNumber_error"] = err.Error()
	} else {
		result["expectedNextSequenceNumber"] = nextSeqNum
	}

	priceEpochAndRound, err := getLatestPriceEpochAndRound(ctx)
	if err != nil {
		result["latestPriceEpochAndRound_error"] = err.Error()
	} else {
		result["latestPriceEpochAndRound"] = priceEpochAndRound
	}

	transmitters, err := getTransmitters(ctx)
	if err != nil {
		result["transmitters_error"] = err.Error()
	} else {
		result["transmitters"] = transmitters
	}

	isUnpausedAndNotCursed, err := getIsUnpausedAndNotCursed(ctx)
	if err != nil {
		result["isUnpausedAndNotCursed_error"] = err.Error()
	} else {
		result["isUnpausedAndNotCursed"] = isUnpausedAndNotCursed
	}

	configDetails, err := getLatestConfigDetails(ctx)
	if err != nil {
		result["latestConfigDetails_error"] = err.Error()
	} else {
		result["latestConfigDetails"] = configDetails
	}

	digestAndEpoch, err := getLatestConfigDigestAndEpoch(ctx)
	if err != nil {
		result["latestConfigDigestAndEpoch_error"] = err.Error()
	} else {
		result["latestConfigDigestAndEpoch"] = digestAndEpoch
	}

	paused, err := getPaused(ctx)
	if err != nil {
		result["paused_error"] = err.Error()
	} else {
		result["paused"] = paused
	}

	return result, nil
}
