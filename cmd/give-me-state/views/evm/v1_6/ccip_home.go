package v1_6

import (
	"call-orchestrator-demo/views"
	"call-orchestrator-demo/views/evm/common"
	"encoding/hex"
	"encoding/json"
	"math/big"
)

// Function selectors for CCIPHome v1.6
var (
	// getNumChainConfigurations() returns (uint256)
	selectorGetNumChainConfigurations = common.HexToSelector("7ac0d41e")
	// getAllChainConfigs(uint256 pageIndex, uint256 pageSize) returns (ChainConfigArgs[])
	selectorGetAllChainConfigs = common.HexToSelector("b74b2356")
	// getCapabilityRegistry() returns (address)
	selectorGetCapabilityRegistry = common.HexToSelector("020330e6")
)

// ViewCCIPHome generates a view of the CCIPHome contract (v1.6.0).
func ViewCCIPHome(ctx *views.ViewContext) (map[string]any, error) {
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

	// Get capability registry
	capRegistry, err := getCapabilityRegistry(ctx)
	if err != nil {
		result["capabilityRegistry_error"] = err.Error()
	} else {
		result["capabilityRegistry"] = capRegistry
	}

	// Get number of chain configurations
	numChains, err := getNumChainConfigurations(ctx)
	if err != nil {
		result["numChainConfigurations_error"] = err.Error()
	} else {
		result["numChainConfigurations"] = numChains

		// Get all chain configs if there are any
		if numChains > 0 {
			chainConfigs, err := getAllChainConfigs(ctx, numChains)
			if err != nil {
				result["chainConfigs_error"] = err.Error()
			} else {
				result["chainConfigs"] = chainConfigs
			}
		} else {
			result["chainConfigs"] = []map[string]any{}
		}
	}

	return result, nil
}

// getCapabilityRegistry fetches the capability registry address.
func getCapabilityRegistry(ctx *views.ViewContext) (string, error) {
	data, err := common.ExecuteCall(ctx, selectorGetCapabilityRegistry)
	if err != nil {
		return "", err
	}
	return common.DecodeAddress(data)
}

// getNumChainConfigurations fetches the number of chain configurations.
func getNumChainConfigurations(ctx *views.ViewContext) (uint64, error) {
	data, err := common.ExecuteCall(ctx, selectorGetNumChainConfigurations)
	if err != nil {
		return 0, err
	}
	return common.DecodeUint64(data)
}

// getAllChainConfigs fetches all chain configurations.
func getAllChainConfigs(ctx *views.ViewContext, numChains uint64) ([]map[string]any, error) {
	// Call getAllChainConfigs(0, numChains)
	pageIndex := common.EncodeUint64(0)
	pageSize := common.EncodeUint64(numChains)

	data, err := common.ExecuteCall(ctx, selectorGetAllChainConfigs, pageIndex, pageSize)
	if err != nil {
		return nil, err
	}

	return decodeChainConfigsArray(data)
}

// decodeChainConfigsArray decodes the ChainConfigArgs[] return value.
// ChainConfigArgs: (uint64 chainSelector, ChainConfig chainConfig)
// ChainConfig: (bytes32[] readers, uint8 fChain, bytes config)
func decodeChainConfigsArray(data []byte) ([]map[string]any, error) {
	if len(data) < 64 {
		return []map[string]any{}, nil
	}

	// Dynamic array: first 32 bytes is offset to array data
	offset := common.DecodeUint64FromBytes(data[0:32])
	if offset+32 > uint64(len(data)) {
		return []map[string]any{}, nil
	}

	// Length of the array
	length := common.DecodeUint64FromBytes(data[offset : offset+32])
	if length == 0 {
		return []map[string]any{}, nil
	}

	configs := make([]map[string]any, 0, length)

	// After length, we have offsets to each ChainConfigArgs element
	offsetsStart := offset + 32
	for i := uint64(0); i < length; i++ {
		if offsetsStart+i*32+32 > uint64(len(data)) {
			break
		}

		// Get the offset to this element (relative to the start of the offsets section, i.e., after the length)
		elementOffset := common.DecodeUint64FromBytes(data[offsetsStart+i*32 : offsetsStart+i*32+32])
		actualOffset := offsetsStart + elementOffset

		if actualOffset+64 > uint64(len(data)) {
			break
		}

		config, err := decodeChainConfigArgs(data, actualOffset)
		if err != nil {
			continue
		}
		configs = append(configs, config)
	}

	return configs, nil
}

// decodeChainConfigArgs decodes a single ChainConfigArgs from the data at the given offset.
func decodeChainConfigArgs(data []byte, offset uint64) (map[string]any, error) {
	if offset+64 > uint64(len(data)) {
		return nil, nil
	}

	result := make(map[string]any)

	// ChainConfigArgs: (uint64 chainSelector, ChainConfig chainConfig)
	// chainSelector is a uint64 (padded to 32 bytes)
	chainSelector := common.DecodeUint64FromBytes(data[offset : offset+32])
	result["chainSelector"] = chainSelector

	// chainConfig offset (relative to this struct)
	chainConfigOffset := common.DecodeUint64FromBytes(data[offset+32 : offset+64])
	actualConfigOffset := offset + chainConfigOffset

	// Decode ChainConfig
	chainConfig, err := decodeChainConfig(data, actualConfigOffset)
	if err == nil && chainConfig != nil {
		result["chainConfig"] = chainConfig
	}

	return result, nil
}

// decodeChainConfig decodes a ChainConfig from the data at the given offset.
// ChainConfig: (bytes32[] readers, uint8 fChain, bytes config)
func decodeChainConfig(data []byte, offset uint64) (map[string]any, error) {
	if offset+96 > uint64(len(data)) {
		return nil, nil
	}

	result := make(map[string]any)

	// readers array offset (relative to this struct)
	readersOffset := common.DecodeUint64FromBytes(data[offset : offset+32])
	actualReadersOffset := offset + readersOffset

	// fChain (uint8 padded to 32 bytes)
	fChain := data[offset+63] // Last byte of the 32-byte slot
	result["fChain"] = fChain

	// config bytes offset
	configOffset := common.DecodeUint64FromBytes(data[offset+64 : offset+96])
	actualConfigOffset := offset + configOffset

	// Decode readers array (bytes32[])
	readers, err := decodeBytes32Array(data, actualReadersOffset)
	if err == nil {
		result["readers"] = readers
	}

	// Decode config bytes
	configBytes, err := decodeBytesField(data, actualConfigOffset)
	if err == nil {
		// Try to decode the config bytes as chain config parameters
		decodedConfig := decodeChainConfigBytes(configBytes)
		if decodedConfig != nil {
			result["config"] = decodedConfig
		} else {
			result["configRaw"] = "0x" + hex.EncodeToString(configBytes)
		}
	}

	return result, nil
}

// decodeBytes32Array decodes a bytes32[] array.
func decodeBytes32Array(data []byte, offset uint64) ([]string, error) {
	if offset+32 > uint64(len(data)) {
		return []string{}, nil
	}

	length := common.DecodeUint64FromBytes(data[offset : offset+32])
	if length == 0 {
		return []string{}, nil
	}

	result := make([]string, 0, length)
	for i := uint64(0); i < length; i++ {
		elemOffset := offset + 32 + i*32
		if elemOffset+32 > uint64(len(data)) {
			break
		}
		// These are P2P peer IDs encoded as bytes32
		// Convert to the 12D3KooW... format if possible, otherwise hex
		peerIdBytes := data[elemOffset : elemOffset+32]
		result = append(result, formatPeerID(peerIdBytes))
	}

	return result, nil
}

// decodeBytesField decodes a dynamic bytes field.
func decodeBytesField(data []byte, offset uint64) ([]byte, error) {
	if offset+32 > uint64(len(data)) {
		return nil, nil
	}

	length := common.DecodeUint64FromBytes(data[offset : offset+32])
	if length == 0 {
		return []byte{}, nil
	}

	if offset+32+length > uint64(len(data)) {
		return nil, nil
	}

	return data[offset+32 : offset+32+length], nil
}

// formatPeerID formats a bytes32 peer ID.
// In libp2p, peer IDs are typically base58-encoded multihash values.
// For now, we'll just return the hex representation.
func formatPeerID(peerIdBytes []byte) string {
	// The peer ID is stored as a 32-byte value
	// In the goal_state, these appear as "12D3KooW..." format (base58)
	// For simplicity, we'll return hex for now
	// A full implementation would use multibase/multihash decoding
	return "0x" + hex.EncodeToString(peerIdBytes)
}

// decodeChainConfigBytes attempts to decode the chain config bytes.
// The config bytes are typically JSON-encoded.
func decodeChainConfigBytes(configBytes []byte) map[string]any {
	if len(configBytes) == 0 {
		return nil
	}

	// Try to decode as JSON first (most common case)
	// Check if it starts with '{' (0x7b)
	if configBytes[0] == 0x7b {
		result := make(map[string]any)
		if err := json.Unmarshal(configBytes, &result); err == nil {
			return result
		}
	}

	// Fallback: try to decode as packed binary values
	// (uint32 gasPriceDeviationPPB, uint32 daGasPriceDeviationPPB, uint32 optimisticConfirmations, bool chainFeeDeviationDisabled)
	if len(configBytes) >= 16 {
		result := make(map[string]any)
		gasPriceDeviationPPB := new(big.Int).SetBytes(configBytes[0:4]).Uint64()
		daGasPriceDeviationPPB := new(big.Int).SetBytes(configBytes[4:8]).Uint64()
		optimisticConfirmations := new(big.Int).SetBytes(configBytes[8:12]).Uint64()

		result["gasPriceDeviationPPB"] = gasPriceDeviationPPB
		result["daGasPriceDeviationPPB"] = daGasPriceDeviationPPB
		result["optimisticConfirmations"] = optimisticConfirmations

		if len(configBytes) >= 13 {
			result["chainFeeDeviationDisabled"] = configBytes[12] != 0
		}
		return result
	}

	return nil
}
