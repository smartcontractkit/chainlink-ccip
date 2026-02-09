package v1_6

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"

	"give-me-state-v2/views"

	"github.com/ethereum/go-ethereum/common"
)

// packCCIPHomeCall packs a method call using the CCIPHome v1.6 ABI.
func packCCIPHomeCall(method string, args ...interface{}) ([]byte, error) {
	return CCIPHomeABI.Pack(method, args...)
}

// executeCCIPHomeCall packs a call, executes it, and returns raw response bytes.
func executeCCIPHomeCall(ctx *views.ViewContext, method string, args ...interface{}) ([]byte, error) {
	calldata, err := packCCIPHomeCall(method, args...)
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

// getCCIPHomeOwner fetches the owner address.
func getCCIPHomeOwner(ctx *views.ViewContext) (string, error) {
	data, err := executeCCIPHomeCall(ctx, "owner")
	if err != nil {
		return "", err
	}
	results, err := CCIPHomeABI.Unpack("owner", data)
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

// getCCIPHomeTypeAndVersion fetches the typeAndVersion string.
func getCCIPHomeTypeAndVersion(ctx *views.ViewContext) (string, error) {
	data, err := executeCCIPHomeCall(ctx, "typeAndVersion")
	if err != nil {
		return "", err
	}
	results, err := CCIPHomeABI.Unpack("typeAndVersion", data)
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

// getCapabilityRegistry fetches the capability registry address.
func getCapabilityRegistry(ctx *views.ViewContext) (string, error) {
	data, err := executeCCIPHomeCall(ctx, "getCapabilityRegistry")
	if err != nil {
		return "", err
	}
	results, err := CCIPHomeABI.Unpack("getCapabilityRegistry", data)
	if err != nil {
		return "", fmt.Errorf("failed to unpack getCapabilityRegistry: %w", err)
	}
	if len(results) == 0 {
		return "", fmt.Errorf("no results from getCapabilityRegistry call")
	}
	addr, ok := results[0].(common.Address)
	if !ok {
		return "", fmt.Errorf("unexpected type for capabilityRegistry: %T", results[0])
	}
	return addr.Hex(), nil
}

// getNumChainConfigurations fetches the number of chain configurations.
func getNumChainConfigurations(ctx *views.ViewContext) (uint64, error) {
	data, err := executeCCIPHomeCall(ctx, "getNumChainConfigurations")
	if err != nil {
		return 0, err
	}
	results, err := CCIPHomeABI.Unpack("getNumChainConfigurations", data)
	if err != nil {
		return 0, fmt.Errorf("failed to unpack getNumChainConfigurations: %w", err)
	}
	if len(results) == 0 {
		return 0, fmt.Errorf("no results from getNumChainConfigurations call")
	}
	val, ok := results[0].(*big.Int)
	if !ok {
		return 0, fmt.Errorf("unexpected type for numChainConfigurations: %T", results[0])
	}
	return val.Uint64(), nil
}

// getAllChainConfigs fetches all chain configurations using ABI bindings.
func getAllChainConfigs(ctx *views.ViewContext, numChains uint64) ([]map[string]any, error) {
	data, err := executeCCIPHomeCall(ctx, "getAllChainConfigs", new(big.Int).SetUint64(0), new(big.Int).SetUint64(numChains))
	if err != nil {
		return nil, err
	}

	results, err := CCIPHomeABI.Unpack("getAllChainConfigs", data)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack getAllChainConfigs: %w", err)
	}
	if len(results) == 0 {
		return []map[string]any{}, nil
	}

	// The result is []CCIPHomeChainConfigArgs
	configs, ok := results[0].([]struct {
		ChainSelector uint64 `json:"chainSelector"`
		ChainConfig   struct {
			Readers [][32]byte `json:"readers"`
			FChain  uint8      `json:"fChain"`
			Config  []byte     `json:"config"`
		} `json:"chainConfig"`
	})
	if !ok {
		return nil, fmt.Errorf("unexpected type for chain configs: %T", results[0])
	}

	out := make([]map[string]any, 0, len(configs))
	for _, cfg := range configs {
		readers := make([]string, len(cfg.ChainConfig.Readers))
		for i, r := range cfg.ChainConfig.Readers {
			readers[i] = "0x" + hex.EncodeToString(r[:])
		}

		chainConfig := map[string]any{
			"readers": readers,
			"fChain":  cfg.ChainConfig.FChain,
		}

		// Try to decode config bytes
		decodedConfig := decodeChainConfigBytes(cfg.ChainConfig.Config)
		if decodedConfig != nil {
			chainConfig["config"] = decodedConfig
		} else if len(cfg.ChainConfig.Config) > 0 {
			chainConfig["configRaw"] = "0x" + hex.EncodeToString(cfg.ChainConfig.Config)
		}

		out = append(out, map[string]any{
			"chainSelector": cfg.ChainSelector,
			"chainConfig":   chainConfig,
		})
	}

	return out, nil
}

// decodeChainConfigBytes attempts to decode the chain config bytes.
func decodeChainConfigBytes(configBytes []byte) map[string]any {
	if len(configBytes) == 0 {
		return nil
	}

	// Try JSON first
	if configBytes[0] == 0x7b {
		result := make(map[string]any)
		if err := json.Unmarshal(configBytes, &result); err == nil {
			return result
		}
	}

	// Fallback: packed binary values
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

// ViewCCIPHome generates a view of the CCIPHome contract (v1.6.0).
// Uses ABI bindings for proper struct decoding.
func ViewCCIPHome(ctx *views.ViewContext) (map[string]any, error) {
	result := make(map[string]any)

	result["address"] = ctx.AddressHex
	result["chainSelector"] = ctx.ChainSelector
	result["version"] = "1.6.0"

	owner, err := getCCIPHomeOwner(ctx)
	if err != nil {
		result["owner_error"] = err.Error()
	} else {
		result["owner"] = owner
	}

	typeAndVersion, err := getCCIPHomeTypeAndVersion(ctx)
	if err != nil {
		result["typeAndVersion_error"] = err.Error()
	} else {
		result["typeAndVersion"] = typeAndVersion
	}

	capRegistry, err := getCapabilityRegistry(ctx)
	if err != nil {
		result["capabilityRegistry_error"] = err.Error()
	} else {
		result["capabilityRegistry"] = capRegistry
	}

	numChains, err := getNumChainConfigurations(ctx)
	if err != nil {
		result["numChainConfigurations_error"] = err.Error()
	} else {
		result["numChainConfigurations"] = numChains

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
