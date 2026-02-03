package v1_5

import (
	"call-orchestrator-demo/views"
	"call-orchestrator-demo/views/evm/common"
	"encoding/hex"
	"math/big"
)

// Function selectors for RMN v1.5
var (
	// getConfigDetails() returns (uint32 version, uint32 blockNumber, RMNConfig config)
	selectorGetConfigDetails = common.HexToSelector("3f42ab73")
)

// RMNSigner represents a signer in the RMN config
type RMNSigner struct {
	OnchainPublicKey string `json:"onchainPublicKey"`
	NodeIndex        uint64 `json:"nodeIndex"`
}

// RMNConfig represents the config returned by getConfigDetails
type RMNConfig struct {
	Signers   []RMNSigner `json:"signers"`
	Threshold uint64      `json:"threshold"`
}

// ViewRMN generates a view of the RMN contract (v1.5.0).
func ViewRMN(ctx *views.ViewContext) (map[string]any, error) {
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

	configDetails, err := getRMNConfigDetails(ctx)
	if err != nil {
		result["configDetails_error"] = err.Error()
	} else {
		result["configDetails"] = configDetails
	}

	return result, nil
}

// getRMNConfigDetails fetches and decodes the config details.
func getRMNConfigDetails(ctx *views.ViewContext) (map[string]any, error) {
	data, err := common.ExecuteCall(ctx, selectorGetConfigDetails)
	if err != nil {
		return nil, err
	}

	// The return is: (uint32 version, uint32 blockNumber, RMNConfig config)
	// RMNConfig is a struct with (Signer[] signers, uint64 threshold)
	// This is complex ABI encoding, let's decode what we can

	if len(data) < 96 {
		return map[string]any{
			"raw": "0x" + hex.EncodeToString(data),
		}, nil
	}

	configVersion := new(big.Int).SetBytes(data[0:32]).Uint64()
	blockNumber := new(big.Int).SetBytes(data[32:64]).Uint64()

	configDetails := map[string]any{
		"version":     configVersion,
		"blockNumber": blockNumber,
	}

	// The config struct starts at offset indicated by data[64:96]
	// Decode signers array and threshold
	if len(data) >= 128 {
		// Offset to config struct
		configOffset := new(big.Int).SetBytes(data[64:96]).Uint64()
		if configOffset < uint64(len(data)) {
			// Try to decode signers and threshold from the config
			config, err := decodeRMNConfig(data, configOffset)
			if err == nil {
				configDetails["config"] = config
			} else {
				// Fall back to raw data
				configDetails["config_raw"] = "0x" + hex.EncodeToString(data[configOffset:])
			}
		}
	}

	return configDetails, nil
}

// decodeRMNConfig decodes the RMNConfig struct from the data at the given offset.
func decodeRMNConfig(data []byte, offset uint64) (map[string]any, error) {
	if offset+64 > uint64(len(data)) {
		return nil, nil
	}

	// RMNConfig: (Signer[] signers, uint64 threshold)
	// signers is a dynamic array, so first 32 bytes is offset to signers array
	// next 32 bytes is threshold

	signersOffset := new(big.Int).SetBytes(data[offset : offset+32]).Uint64()
	threshold := new(big.Int).SetBytes(data[offset+32 : offset+64]).Uint64()

	config := map[string]any{
		"threshold": threshold,
	}

	// Decode signers array
	signersArrayOffset := offset + signersOffset
	if signersArrayOffset+32 <= uint64(len(data)) {
		signersLength := new(big.Int).SetBytes(data[signersArrayOffset : signersArrayOffset+32]).Uint64()
		signers := make([]map[string]any, 0, signersLength)

		// Each Signer is (address onchainPublicKey, uint64 nodeIndex)
		// Packed as 64 bytes per signer (32 for address, 32 for nodeIndex)
		for i := uint64(0); i < signersLength; i++ {
			signerOffset := signersArrayOffset + 32 + i*64
			if signerOffset+64 > uint64(len(data)) {
				break
			}

			pubKey := "0x" + hex.EncodeToString(data[signerOffset+12:signerOffset+32])
			nodeIndex := new(big.Int).SetBytes(data[signerOffset+32 : signerOffset+64]).Uint64()

			signers = append(signers, map[string]any{
				"onchainPublicKey": pubKey,
				"nodeIndex":        nodeIndex,
			})
		}
		config["signers"] = signers
	}

	return config, nil
}
