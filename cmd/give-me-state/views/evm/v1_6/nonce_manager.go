package v1_6

import (
	"encoding/hex"

	"call-orchestrator-demo/views"
	"call-orchestrator-demo/views/evm/common"
)

// Function selectors for NonceManager
var (
	// getAllAuthorizedCallers() returns (address[])
	selectorGetAllAuthorizedCallers = common.HexToSelector("2451a627")
)

// getNonceManagerAuthorizedCallers fetches all authorized callers using manual decoding.
func getNonceManagerAuthorizedCallers(ctx *views.ViewContext) ([]string, error) {
	data, err := common.ExecuteCall(ctx, selectorGetAllAuthorizedCallers)
	if err != nil {
		return nil, err
	}

	// Decode dynamic array of addresses
	// ABI encoding: offset (32 bytes) + length (32 bytes) + elements (32 bytes each)
	if len(data) < 64 {
		return []string{}, nil
	}

	// Read length from offset 32
	length := common.DecodeUint64FromBytes(data[32:64])
	if length == 0 {
		return []string{}, nil
	}

	callers := make([]string, length)
	for i := uint64(0); i < length; i++ {
		offset := 64 + i*32
		if offset+32 > uint64(len(data)) {
			break
		}
		// Address is in the last 20 bytes of the 32-byte slot
		addr := data[offset+12 : offset+32]
		callers[i] = "0x" + hex.EncodeToString(addr)
	}
	return callers, nil
}

// ViewNonceManager generates a view of the NonceManager contract (v1.6.0).
func ViewNonceManager(ctx *views.ViewContext) (map[string]any, error) {
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
	authorizedCallers, err := getNonceManagerAuthorizedCallers(ctx)
	if err != nil {
		result["authorizedCallers_error"] = err.Error()
	} else {
		result["authorizedCallers"] = authorizedCallers
	}

	return result, nil
}
