package mcms

import (
	"call-orchestrator-demo/views"
	"call-orchestrator-demo/views/evm/common"
	"encoding/hex"
)

// Function selectors for MCMS contracts (ManyChainMultiSig)
var (
	// owner() returns (address)
	selectorOwner = common.HexToSelector("8da5cb5b")
	// getConfig() returns (Signer[], uint8[32] groupQuorums, uint8[32] groupParents)
	selectorGetConfig = common.HexToSelector("c3f909d4")
)

// Signer represents a signer in the MCMS config
type Signer struct {
	Address string `json:"addr"`
	Index   uint8  `json:"index"`
	Group   uint8  `json:"group"`
}

// MCMSConfig represents the MCMS contract configuration
type MCMSConfig struct {
	Signers      []Signer `json:"signers"`
	GroupQuorums []uint8  `json:"groupQuorums"`
	GroupParents []uint8  `json:"groupParents"`
}

// ViewMCMS generates a view of an MCMS contract (bypasser, canceller, proposer).
func ViewMCMS(ctx *views.ViewContext) (map[string]any, error) {
	result := make(map[string]any)

	result["address"] = ctx.AddressHex
	result["chainSelector"] = ctx.ChainSelector

	// Get owner
	if owner, err := getMCMSOwner(ctx); err == nil {
		result["owner"] = owner
	} else {
		result["owner_error"] = err.Error()
	}

	// Get config
	if config, err := getMCMSConfig(ctx); err == nil {
		result["config"] = config
	} else {
		result["config_error"] = err.Error()
	}

	return result, nil
}

// getMCMSOwner fetches the owner address.
func getMCMSOwner(ctx *views.ViewContext) (string, error) {
	data, err := common.ExecuteCall(ctx, selectorOwner)
	if err != nil {
		return "", err
	}
	return common.DecodeAddress(data)
}

// getMCMSConfig fetches and decodes the MCMS configuration.
func getMCMSConfig(ctx *views.ViewContext) (*MCMSConfig, error) {
	data, err := common.ExecuteCall(ctx, selectorGetConfig)
	if err != nil {
		return nil, err
	}

	return decodeMCMSConfig(data)
}

// ViewCallProxy generates a view of a CallProxy contract.
// CallProxy has no public read functions, so we just return the address.
func ViewCallProxy(ctx *views.ViewContext) (map[string]any, error) {
	result := make(map[string]any)

	result["address"] = ctx.AddressHex
	result["chainSelector"] = ctx.ChainSelector
	// CallProxy has no owner or other readable state
	result["owner"] = "0x0000000000000000000000000000000000000000"

	return result, nil
}

// decodeMCMSConfig decodes the getConfig() return data.
// Returns: (Signer[] signers, uint8[32] groupQuorums, uint8[32] groupParents)
// Signer struct: (address addr, uint8 index, uint8 group)
func decodeMCMSConfig(data []byte) (*MCMSConfig, error) {
	config := &MCMSConfig{
		Signers:      []Signer{},
		GroupQuorums: make([]uint8, 32),
		GroupParents: make([]uint8, 32),
	}

	if len(data) < 96 {
		return config, nil
	}

	// Layout:
	// [0:32]   - offset to signers array
	// [32:64]  - offset to groupQuorums array (fixed size 32)
	// [64:96]  - offset to groupParents array (fixed size 32)
	// Then the actual data at those offsets

	signersOffset := common.DecodeUint64FromBytes(data[0:32])
	groupQuorumsOffset := common.DecodeUint64FromBytes(data[32:64])
	groupParentsOffset := common.DecodeUint64FromBytes(data[64:96])

	// Decode signers array
	if signersOffset+32 <= uint64(len(data)) {
		signersLen := common.DecodeUint64FromBytes(data[signersOffset : signersOffset+32])
		signersDataStart := signersOffset + 32

		for i := uint64(0); i < signersLen; i++ {
			// Each signer is a tuple (address, uint8, uint8) = 3 * 32 bytes
			signerOffset := signersDataStart + i*96
			if signerOffset+96 > uint64(len(data)) {
				break
			}

			// address (20 bytes, right-padded in 32 bytes)
			addr := "0x" + hex.EncodeToString(data[signerOffset+12:signerOffset+32])
			// uint8 index
			index := uint8(data[signerOffset+63])
			// uint8 group
			group := uint8(data[signerOffset+95])

			config.Signers = append(config.Signers, Signer{
				Address: addr,
				Index:   index,
				Group:   group,
			})
		}
	}

	// Decode groupQuorums (fixed array of 32 uint8s)
	// uint8[32] is packed into 32 bytes (1 byte each), but in ABI encoding each uint8 takes 32 bytes
	if groupQuorumsOffset+32*32 <= uint64(len(data)) {
		for i := 0; i < 32; i++ {
			offset := groupQuorumsOffset + uint64(i)*32
			config.GroupQuorums[i] = uint8(data[offset+31])
		}
	}

	// Decode groupParents (fixed array of 32 uint8s)
	if groupParentsOffset+32*32 <= uint64(len(data)) {
		for i := 0; i < 32; i++ {
			offset := groupParentsOffset + uint64(i)*32
			config.GroupParents[i] = uint8(data[offset+31])
		}
	}

	return config, nil
}
