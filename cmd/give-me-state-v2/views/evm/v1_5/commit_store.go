package v1_5

import (
	"give-me-state-v2/views"
	"give-me-state-v2/views/evm/common"
	"encoding/hex"
	"math/big"
)

// Function selectors for CommitStore v1.5
var (
	// getStaticConfig() returns (StaticConfig)
	selectorGetStaticConfig = common.HexToSelector("06285c69")
	// getDynamicConfig() returns (DynamicConfig)
	selectorGetDynamicConfig = common.HexToSelector("7437ff9f")
	// getExpectedNextSequenceNumber() returns (uint64)
	selectorGetExpectedNextSeqNum = common.HexToSelector("4120fccd")
	// getLatestPriceEpochAndRound() returns (uint64)
	selectorGetLatestPriceEpochAndRound = common.HexToSelector("10c374ed")
	// getTransmitters() returns (address[])
	selectorGetTransmitters = common.HexToSelector("666cab8d")
	// isUnpausedAndNotCursed() returns (bool)
	selectorIsUnpausedAndNotCursed = common.HexToSelector("e89d039f")
	// latestConfigDetails() returns (uint32 configCount, uint32 blockNumber, bytes32 configDigest)
	selectorLatestConfigDetails = common.HexToSelector("81ff7048")
	// latestConfigDigestAndEpoch() returns (bool scanLogs, bytes32 configDigest, uint64 epoch)
	selectorLatestConfigDigestAndEpoch = common.HexToSelector("afcb95d7")
	// paused() returns (bool)
	selectorPaused = common.HexToSelector("5c975abb")
)

// ViewCommitStore generates a view of the CommitStore contract (v1.5.0).
func ViewCommitStore(ctx *views.ViewContext) (map[string]any, error) {
	result := make(map[string]any)

	result["address"] = ctx.AddressHex
	result["chainSelector"] = ctx.ChainSelector
	result["version"] = "1.5.0"

	owner, err := common.GetOwner(ctx)
	if err != nil {
		result["owner_error"] = err.Error()
	} else {
		result["owner"] = owner
	}

	typeAndVersion, err := common.GetTypeAndVersion(ctx)
	if err != nil {
		result["typeAndVersion_error"] = err.Error()
	} else {
		result["typeAndVersion"] = typeAndVersion
	}

	// Get static config
	staticConfig, err := getCommitStoreStaticConfig(ctx)
	if err != nil {
		result["staticConfig_error"] = err.Error()
	} else {
		result["staticConfig"] = staticConfig
	}

	// Get dynamic config
	dynamicConfig, err := getCommitStoreDynamicConfig(ctx)
	if err != nil {
		result["dynamicConfig_error"] = err.Error()
	} else {
		result["dynamicConfig"] = dynamicConfig
	}

	// Get expected next sequence number
	nextSeqNum, err := getExpectedNextSequenceNumber(ctx)
	if err != nil {
		result["expectedNextSequenceNumber_error"] = err.Error()
	} else {
		result["expectedNextSequenceNumber"] = nextSeqNum
	}

	// Get latest price epoch and round
	priceEpochAndRound, err := getLatestPriceEpochAndRound(ctx)
	if err != nil {
		result["latestPriceEpochAndRound_error"] = err.Error()
	} else {
		result["latestPriceEpochAndRound"] = priceEpochAndRound
	}

	// Get transmitters
	transmitters, err := getTransmitters(ctx)
	if err != nil {
		result["transmitters_error"] = err.Error()
	} else {
		result["transmitters"] = transmitters
	}

	// Get isUnpausedAndNotCursed
	isUnpausedAndNotCursed, err := getIsUnpausedAndNotCursed(ctx)
	if err != nil {
		result["isUnpausedAndNotCursed_error"] = err.Error()
	} else {
		result["isUnpausedAndNotCursed"] = isUnpausedAndNotCursed
	}

	// Get latestConfigDetails
	configDetails, err := getLatestConfigDetails(ctx)
	if err != nil {
		result["latestConfigDetails_error"] = err.Error()
	} else {
		result["latestConfigDetails"] = configDetails
	}

	// Get latestConfigDigestAndEpoch
	digestAndEpoch, err := getLatestConfigDigestAndEpoch(ctx)
	if err != nil {
		result["latestConfigDigestAndEpoch_error"] = err.Error()
	} else {
		result["latestConfigDigestAndEpoch"] = digestAndEpoch
	}

	// Get paused status
	paused, err := getPaused(ctx)
	if err != nil {
		result["paused_error"] = err.Error()
	} else {
		result["paused"] = paused
	}

	return result, nil
}

// getCommitStoreStaticConfig fetches and decodes the static config.
func getCommitStoreStaticConfig(ctx *views.ViewContext) (map[string]any, error) {
	data, err := common.ExecuteCall(ctx, selectorGetStaticConfig)
	if err != nil {
		return nil, err
	}

	// StaticConfig: (uint64 chainSelector, uint64 sourceChainSelector, address onRamp, address commitStore, address rmnProxy)
	// Note: Some fields may vary, decode what we can
	if len(data) < 160 {
		return map[string]any{"raw": "0x" + hex.EncodeToString(data)}, nil
	}

	return map[string]any{
		"chainSelector":       new(big.Int).SetBytes(data[0:32]).Uint64(),
		"sourceChainSelector": new(big.Int).SetBytes(data[32:64]).Uint64(),
		"onRamp":              "0x" + hex.EncodeToString(data[76:96]),
		"commitStore":         "0x" + hex.EncodeToString(data[108:128]),
		"rmnProxy":            "0x" + hex.EncodeToString(data[140:160]),
	}, nil
}

// getCommitStoreDynamicConfig fetches and decodes the dynamic config.
func getCommitStoreDynamicConfig(ctx *views.ViewContext) (map[string]any, error) {
	data, err := common.ExecuteCall(ctx, selectorGetDynamicConfig)
	if err != nil {
		return nil, err
	}

	// DynamicConfig: (address priceRegistry)
	if len(data) < 32 {
		return map[string]any{"raw": "0x" + hex.EncodeToString(data)}, nil
	}

	return map[string]any{
		"priceRegistry": "0x" + hex.EncodeToString(data[12:32]),
	}, nil
}

// getExpectedNextSequenceNumber fetches the expected next sequence number.
func getExpectedNextSequenceNumber(ctx *views.ViewContext) (uint64, error) {
	data, err := common.ExecuteCall(ctx, selectorGetExpectedNextSeqNum)
	if err != nil {
		return 0, err
	}
	return common.DecodeUint64(data)
}

// getLatestPriceEpochAndRound fetches the latest price epoch and round.
func getLatestPriceEpochAndRound(ctx *views.ViewContext) (uint64, error) {
	data, err := common.ExecuteCall(ctx, selectorGetLatestPriceEpochAndRound)
	if err != nil {
		return 0, err
	}
	return common.DecodeUint64(data)
}

// getTransmitters fetches the list of transmitter addresses.
func getTransmitters(ctx *views.ViewContext) ([]string, error) {
	data, err := common.ExecuteCall(ctx, selectorGetTransmitters)
	if err != nil {
		return nil, err
	}

	// Dynamic array of addresses
	if len(data) < 64 {
		return []string{}, nil
	}

	// First 32 bytes: offset to array
	// At offset: length
	// Then: addresses
	length := common.DecodeUint64FromBytes(data[32:64])
	if length == 0 {
		return []string{}, nil
	}

	transmitters := make([]string, 0, length)
	for i := uint64(0); i < length; i++ {
		offset := 64 + i*32
		if offset+32 > uint64(len(data)) {
			break
		}
		addr := "0x" + hex.EncodeToString(data[offset+12:offset+32])
		transmitters = append(transmitters, addr)
	}

	return transmitters, nil
}

// getIsUnpausedAndNotCursed fetches the isUnpausedAndNotCursed flag.
func getIsUnpausedAndNotCursed(ctx *views.ViewContext) (bool, error) {
	data, err := common.ExecuteCall(ctx, selectorIsUnpausedAndNotCursed)
	if err != nil {
		return false, err
	}
	return common.DecodeBool(data)
}

// getLatestConfigDetails fetches the latest config details.
func getLatestConfigDetails(ctx *views.ViewContext) (map[string]any, error) {
	data, err := common.ExecuteCall(ctx, selectorLatestConfigDetails)
	if err != nil {
		return nil, err
	}

	// Returns (uint32 configCount, uint32 blockNumber, bytes32 configDigest)
	if len(data) < 96 {
		return map[string]any{"raw": "0x" + hex.EncodeToString(data)}, nil
	}

	return map[string]any{
		"configCount":  new(big.Int).SetBytes(data[0:32]).Uint64(),
		"blockNumber":  new(big.Int).SetBytes(data[32:64]).Uint64(),
		"configDigest": "0x" + hex.EncodeToString(data[64:96]),
	}, nil
}

// getLatestConfigDigestAndEpoch fetches the latest config digest and epoch.
func getLatestConfigDigestAndEpoch(ctx *views.ViewContext) (map[string]any, error) {
	data, err := common.ExecuteCall(ctx, selectorLatestConfigDigestAndEpoch)
	if err != nil {
		return nil, err
	}

	// Returns (bool scanLogs, bytes32 configDigest, uint64 epoch)
	if len(data) < 96 {
		return map[string]any{"raw": "0x" + hex.EncodeToString(data)}, nil
	}

	scanLogs := data[31] != 0

	return map[string]any{
		"scanLogs":     scanLogs,
		"configDigest": "0x" + hex.EncodeToString(data[32:64]),
		"epoch":        new(big.Int).SetBytes(data[64:96]).Uint64(),
	}, nil
}

// getPaused fetches the paused status.
func getPaused(ctx *views.ViewContext) (bool, error) {
	data, err := common.ExecuteCall(ctx, selectorPaused)
	if err != nil {
		return false, err
	}
	return common.DecodeBool(data)
}
