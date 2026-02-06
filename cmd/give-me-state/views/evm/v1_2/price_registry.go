package v1_2

import (
	"encoding/hex"
	"fmt"
	"math/big"

	"call-orchestrator-demo/views"
	"call-orchestrator-demo/views/evm/common"

	"github.com/ethereum/go-ethereum/accounts/abi"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_2_0/price_registry"
)

// Function selectors for PriceRegistry
var (
	PriceRegistryABI abi.ABI

	// getFeeTokens() returns (address[])
	selectorGetFeeTokens []byte
	// getStalenessThreshold() returns (uint128)
	selectorGetStalenessThreshold []byte
	// getPriceUpdaters() returns (address[])
	selectorGetPriceUpdaters []byte
)

func init() {
	// Parse the PriceRegistry ABI once at startup
	parsedPriceRegistry, err := price_registry.PriceRegistryMetaData.GetAbi()
	if err != nil {
		panic(fmt.Sprintf("failed to parse PriceRegistry ABI: %v", err))
	}
	PriceRegistryABI = *parsedPriceRegistry

	selectorGetFeeTokens = PriceRegistryABI.Methods["getFeeTokens"].ID
	selectorGetStalenessThreshold = PriceRegistryABI.Methods["getStalenessThreshold"].ID
	selectorGetPriceUpdaters = PriceRegistryABI.Methods["getPriceUpdaters"].ID
}

// getPriceRegistryFeeTokens fetches the list of fee tokens using manual decoding.
func getPriceRegistryFeeTokens(ctx *views.ViewContext) ([]string, error) {
	data, err := common.ExecuteCall(ctx, selectorGetFeeTokens)
	if err != nil {
		return nil, err
	}
	return decodeAddressArray(data)
}

// getPriceRegistryStalenessThreshold fetches the staleness threshold.
func getPriceRegistryStalenessThreshold(ctx *views.ViewContext) (string, error) {
	data, err := common.ExecuteCall(ctx, selectorGetStalenessThreshold)
	if err != nil {
		return "", err
	}
	if len(data) < 32 {
		return "0", nil
	}
	n := new(big.Int).SetBytes(data[:32])
	return n.String(), nil
}

// getPriceRegistryUpdaters fetches the list of price updaters.
func getPriceRegistryUpdaters(ctx *views.ViewContext) ([]string, error) {
	data, err := common.ExecuteCall(ctx, selectorGetPriceUpdaters)
	if err != nil {
		return nil, err
	}
	return decodeAddressArray(data)
}

// decodeAddressArray decodes an ABI-encoded dynamic array of addresses.
func decodeAddressArray(data []byte) ([]string, error) {
	if len(data) < 64 {
		return []string{}, nil
	}
	// Read length from offset 32
	length := common.DecodeUint64FromBytes(data[32:64])
	if length == 0 {
		return []string{}, nil
	}
	addresses := make([]string, length)
	for i := uint64(0); i < length; i++ {
		offset := 64 + i*32
		if offset+32 > uint64(len(data)) {
			break
		}
		// Address is in the last 20 bytes of the 32-byte slot
		addr := data[offset+12 : offset+32]
		addresses[i] = "0x" + hex.EncodeToString(addr)
	}
	return addresses, nil
}

// ViewPriceRegistry generates a view of the PriceRegistry contract (v1.2.0).
func ViewPriceRegistry(ctx *views.ViewContext) (map[string]any, error) {
	result := make(map[string]any)

	result["address"] = ctx.AddressHex
	result["chainSelector"] = ctx.ChainSelector
	result["version"] = "1.2.0"

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

	// Get fee tokens
	feeTokens, err := getPriceRegistryFeeTokens(ctx)
	if err != nil {
		result["feeTokens_error"] = err.Error()
	} else {
		result["feeTokens"] = feeTokens
	}

	// Get staleness threshold
	stalenessThreshold, err := getPriceRegistryStalenessThreshold(ctx)
	if err != nil {
		result["stalenessThreshold_error"] = err.Error()
	} else {
		result["stalenessThreshold"] = stalenessThreshold
	}

	// Get price updaters
	updaters, err := getPriceRegistryUpdaters(ctx)
	if err != nil {
		result["updaters_error"] = err.Error()
	} else {
		result["updaters"] = updaters
	}

	return result, nil
}
