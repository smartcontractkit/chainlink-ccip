package v1_6

import (
	"give-me-state-v2/views"
	"give-me-state-v2/views/evm/common"
	"encoding/hex"
	"math/big"
)

// Function selectors for RMNHome v1.6
var (
	// getActiveDigest() returns (bytes32)
	selectorGetActiveDigest = common.HexToSelector("123e65db")
	// getCandidateDigest() returns (bytes32)
	selectorGetCandidateDigest = common.HexToSelector("38354c5c")
	// getConfig(bytes32 configDigest) returns (VersionedConfig, bool ok)
	selectorGetConfig = common.HexToSelector("6dd5b69d")
)

// ViewRMNHome generates a view of the RMNHome contract (v1.6.0).
func ViewRMNHome(ctx *views.ViewContext) (map[string]any, error) {
	result := make(map[string]any)

	result["address"] = ctx.AddressHex
	result["chainSelector"] = ctx.ChainSelector
	result["version"] = "1.6.0"

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

	// Get active config
	activeDigest, err := getDigest(ctx, selectorGetActiveDigest)
	if err != nil {
		result["activeDigest_error"] = err.Error()
	} else {
		result["activeDigest"] = activeDigest
		if activeDigest != "0x0000000000000000000000000000000000000000000000000000000000000000" {
			activeConfig, err := getRMNHomeConfig(ctx, activeDigest)
			if err != nil {
				result["activeConfig_error"] = err.Error()
			} else {
				result["activeConfig"] = activeConfig
			}
		}
	}

	// Get candidate config
	candidateDigest, err := getDigest(ctx, selectorGetCandidateDigest)
	if err != nil {
		result["candidateDigest_error"] = err.Error()
	} else {
		result["candidateDigest"] = candidateDigest
		if candidateDigest != "0x0000000000000000000000000000000000000000000000000000000000000000" {
			candidateConfig, err := getRMNHomeConfig(ctx, candidateDigest)
			if err != nil {
				result["candidateConfig_error"] = err.Error()
			} else {
				result["candidateConfig"] = candidateConfig
			}
		}
	}

	return result, nil
}

// getDigest fetches a bytes32 digest.
func getDigest(ctx *views.ViewContext, selector []byte) (string, error) {
	data, err := common.ExecuteCall(ctx, selector)
	if err != nil {
		return "", err
	}
	if len(data) < 32 {
		return "", nil
	}
	return "0x" + hex.EncodeToString(data[:32]), nil
}

// getRMNHomeConfig fetches the config for a given digest.
func getRMNHomeConfig(ctx *views.ViewContext, digestHex string) (map[string]any, error) {
	// Decode the digest hex string to bytes
	digestBytes, err := hex.DecodeString(digestHex[2:]) // Remove "0x" prefix
	if err != nil {
		return nil, err
	}

	// Pad to 32 bytes if needed
	digest := make([]byte, 32)
	copy(digest[32-len(digestBytes):], digestBytes)

	data, err := common.ExecuteCall(ctx, selectorGetConfig, digest)
	if err != nil {
		return nil, err
	}

	// The return is (VersionedConfig versionedConfig, bool ok)
	// VersionedConfig is a complex struct, decode what we can
	return decodeRMNHomeVersionedConfig(data)
}

// decodeRMNHomeVersionedConfig decodes the VersionedConfig struct from getConfig return data.
// VersionedConfig: (uint32 version, bytes32 configDigest, StaticConfig staticConfig, DynamicConfig dynamicConfig)
// StaticConfig: (Node[] nodes)
// DynamicConfig: (SourceChain[] sourceChains)
// Node: (bytes32 peerId, bytes32 offchainPublicKey)
// SourceChain: (uint64 chainSelector, uint64 fObserve, uint256 observerNodesBitmap)
func decodeRMNHomeVersionedConfig(data []byte) (map[string]any, error) {
	if len(data) < 128 {
		return map[string]any{"raw": "0x" + hex.EncodeToString(data)}, nil
	}

	// The struct is ABI encoded with dynamic components
	// First decode the static parts

	config := make(map[string]any)

	// First 32 bytes: offset to versionedConfig tuple (usually 0x40 = 64)
	// We need to navigate the ABI encoding carefully

	// For a tuple return, the data is:
	// - offset to VersionedConfig (32 bytes)
	// - ok bool (32 bytes)
	// - Then the VersionedConfig data at the offset

	// Check if ok is true (last part of the return)
	// Actually the layout is more complex. Let's try to decode version and digest directly

	// Simplified: just extract version from the data
	// The VersionedConfig starts with version (uint32) which is padded to 32 bytes
	version := common.DecodeUint64FromBytes(data[0:32])
	config["version"] = version

	// Next 32 bytes should be configDigest
	if len(data) >= 64 {
		config["configDigest"] = "0x" + hex.EncodeToString(data[32:64])
	}

	// The staticConfig and dynamicConfig are dynamic types with offsets
	// For simplicity, we'll just indicate we have them without full decoding
	if len(data) >= 128 {
		staticConfigOffset := common.DecodeUint64FromBytes(data[64:96])
		dynamicConfigOffset := common.DecodeUint64FromBytes(data[96:128])

		// Try to decode static config (nodes array)
		if staticConfigOffset > 0 && staticConfigOffset < uint64(len(data)) {
			nodes, err := decodeRMNNodes(data, staticConfigOffset)
			if err == nil {
				config["staticConfig"] = map[string]any{"nodes": nodes}
			}
		}

		// Try to decode dynamic config (source chains array)
		if dynamicConfigOffset > 0 && dynamicConfigOffset < uint64(len(data)) {
			sourceChains, err := decodeRMNSourceChains(data, dynamicConfigOffset)
			if err == nil {
				config["dynamicConfig"] = map[string]any{"sourceChains": sourceChains}
			}
		}
	}

	return config, nil
}

// decodeRMNNodes decodes the nodes array from the data.
func decodeRMNNodes(data []byte, offset uint64) ([]map[string]any, error) {
	if offset+32 > uint64(len(data)) {
		return []map[string]any{}, nil
	}

	// At offset, there's another offset to the array
	arrayOffset := common.DecodeUint64FromBytes(data[offset : offset+32])
	actualOffset := offset + arrayOffset

	if actualOffset+32 > uint64(len(data)) {
		return []map[string]any{}, nil
	}

	// Length of the array
	length := common.DecodeUint64FromBytes(data[actualOffset : actualOffset+32])
	if length == 0 {
		return []map[string]any{}, nil
	}

	nodes := make([]map[string]any, 0, length)

	// Each node is (bytes32 peerId, bytes32 offchainPublicKey) = 64 bytes
	for i := uint64(0); i < length; i++ {
		nodeOffset := actualOffset + 32 + i*64
		if nodeOffset+64 > uint64(len(data)) {
			break
		}

		peerId := "0x" + hex.EncodeToString(data[nodeOffset:nodeOffset+32])
		offchainPubKey := "0x" + hex.EncodeToString(data[nodeOffset+32:nodeOffset+64])

		nodes = append(nodes, map[string]any{
			"peerId":            peerId,
			"offchainPublicKey": offchainPubKey,
		})
	}

	return nodes, nil
}

// decodeRMNSourceChains decodes the source chains array from the data.
func decodeRMNSourceChains(data []byte, offset uint64) ([]map[string]any, error) {
	if offset+32 > uint64(len(data)) {
		return []map[string]any{}, nil
	}

	// At offset, there's another offset to the array
	arrayOffset := common.DecodeUint64FromBytes(data[offset : offset+32])
	actualOffset := offset + arrayOffset

	if actualOffset+32 > uint64(len(data)) {
		return []map[string]any{}, nil
	}

	// Length of the array
	length := common.DecodeUint64FromBytes(data[actualOffset : actualOffset+32])
	if length == 0 {
		return []map[string]any{}, nil
	}

	chains := make([]map[string]any, 0, length)

	// Each source chain is (uint64 chainSelector, uint64 fObserve, uint256 observerNodesBitmap) = 96 bytes
	for i := uint64(0); i < length; i++ {
		chainOffset := actualOffset + 32 + i*96
		if chainOffset+96 > uint64(len(data)) {
			break
		}

		chainSelector := common.DecodeUint64FromBytes(data[chainOffset : chainOffset+32])
		fObserve := common.DecodeUint64FromBytes(data[chainOffset+32 : chainOffset+64])
		observerBitmap := new(big.Int).SetBytes(data[chainOffset+64 : chainOffset+96])

		chains = append(chains, map[string]any{
			"chainSelector":       chainSelector,
			"fObserve":            fObserve,
			"observerNodesBitmap": observerBitmap.String(),
		})
	}

	return chains, nil
}
