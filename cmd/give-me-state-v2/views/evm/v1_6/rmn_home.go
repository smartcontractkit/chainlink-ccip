package v1_6

import (
	"encoding/hex"
	"fmt"
	"math/big"

	"give-me-state-v2/views"

	"github.com/ethereum/go-ethereum/common"
)

// packRMNHomeCall packs a method call using the RMNHome v1.6 ABI.
func packRMNHomeCall(method string, args ...interface{}) ([]byte, error) {
	return RMNHomeABI.Pack(method, args...)
}

// executeRMNHomeCall packs a call, executes it, and returns raw response bytes.
func executeRMNHomeCall(ctx *views.ViewContext, method string, args ...interface{}) ([]byte, error) {
	calldata, err := packRMNHomeCall(method, args...)
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

// getRMNHomeOwner fetches the owner address.
func getRMNHomeOwner(ctx *views.ViewContext) (string, error) {
	data, err := executeRMNHomeCall(ctx, "owner")
	if err != nil {
		return "", err
	}
	results, err := RMNHomeABI.Unpack("owner", data)
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

// getRMNHomeTypeAndVersion fetches the typeAndVersion string.
func getRMNHomeTypeAndVersion(ctx *views.ViewContext) (string, error) {
	data, err := executeRMNHomeCall(ctx, "typeAndVersion")
	if err != nil {
		return "", err
	}
	results, err := RMNHomeABI.Unpack("typeAndVersion", data)
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

// getDigest fetches a bytes32 digest using ABI bindings.
func getDigest(ctx *views.ViewContext, method string) (string, error) {
	data, err := executeRMNHomeCall(ctx, method)
	if err != nil {
		return "", err
	}
	results, err := RMNHomeABI.Unpack(method, data)
	if err != nil {
		return "", fmt.Errorf("failed to unpack %s: %w", method, err)
	}
	if len(results) == 0 {
		return "", fmt.Errorf("no results from %s call", method)
	}
	digest, ok := results[0].([32]byte)
	if !ok {
		return "", fmt.Errorf("unexpected type for digest: %T", results[0])
	}
	return "0x" + hex.EncodeToString(digest[:]), nil
}

// getRMNHomeConfig fetches the config for a given digest using ABI bindings.
func getRMNHomeConfig(ctx *views.ViewContext, digestHex string) (map[string]any, error) {
	digestBytes, err := hex.DecodeString(digestHex[2:])
	if err != nil {
		return nil, err
	}
	var digest [32]byte
	copy(digest[32-len(digestBytes):], digestBytes)

	data, err := executeRMNHomeCall(ctx, "getConfig", digest)
	if err != nil {
		return nil, err
	}

	results, err := RMNHomeABI.Unpack("getConfig", data)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack getConfig: %w", err)
	}
	if len(results) < 2 {
		return nil, fmt.Errorf("expected 2 results from getConfig, got %d", len(results))
	}

	// Check ok flag
	ok, isBool := results[1].(bool)
	if isBool && !ok {
		return nil, fmt.Errorf("config not found for digest %s", digestHex)
	}

	// The result is a VersionedConfig struct
	vCfg, isOk := results[0].(struct {
		Version       uint32   `json:"version"`
		ConfigDigest  [32]byte `json:"configDigest"`
		StaticConfig  struct {
			Nodes []struct {
				PeerId            [32]byte `json:"peerId"`
				OffchainPublicKey [32]byte `json:"offchainPublicKey"`
			} `json:"nodes"`
			OffchainConfig []byte `json:"offchainConfig"`
		} `json:"staticConfig"`
		DynamicConfig struct {
			SourceChains []struct {
				ChainSelector       uint64   `json:"chainSelector"`
				FObserve            uint64   `json:"fObserve"`
				ObserverNodesBitmap *big.Int `json:"observerNodesBitmap"`
			} `json:"sourceChains"`
			OffchainConfig []byte `json:"offchainConfig"`
		} `json:"dynamicConfig"`
	})
	if !isOk {
		return nil, fmt.Errorf("unexpected type for VersionedConfig: %T", results[0])
	}

	// Build nodes
	nodes := make([]map[string]any, len(vCfg.StaticConfig.Nodes))
	for i, n := range vCfg.StaticConfig.Nodes {
		nodes[i] = map[string]any{
			"peerId":            "0x" + hex.EncodeToString(n.PeerId[:]),
			"offchainPublicKey": "0x" + hex.EncodeToString(n.OffchainPublicKey[:]),
		}
	}

	// Build source chains
	sourceChains := make([]map[string]any, len(vCfg.DynamicConfig.SourceChains))
	for i, sc := range vCfg.DynamicConfig.SourceChains {
		sourceChains[i] = map[string]any{
			"chainSelector":       sc.ChainSelector,
			"fObserve":            sc.FObserve,
			"observerNodesBitmap": sc.ObserverNodesBitmap.String(),
		}
	}

	config := map[string]any{
		"version":      vCfg.Version,
		"configDigest": "0x" + hex.EncodeToString(vCfg.ConfigDigest[:]),
		"staticConfig": map[string]any{
			"nodes": nodes,
		},
		"dynamicConfig": map[string]any{
			"sourceChains": sourceChains,
		},
	}

	return config, nil
}

// ViewRMNHome generates a view of the RMNHome contract (v1.6.0).
// Uses ABI bindings for proper struct decoding.
func ViewRMNHome(ctx *views.ViewContext) (map[string]any, error) {
	result := make(map[string]any)

	result["address"] = ctx.AddressHex
	result["chainSelector"] = ctx.ChainSelector
	result["version"] = "1.6.0"

	owner, err := getRMNHomeOwner(ctx)
	if err != nil {
		result["owner_error"] = err.Error()
	} else {
		result["owner"] = owner
	}

	typeAndVersion, err := getRMNHomeTypeAndVersion(ctx)
	if err != nil {
		result["typeAndVersion_error"] = err.Error()
	} else {
		result["typeAndVersion"] = typeAndVersion
	}

	activeDigest, err := getDigest(ctx, "getActiveDigest")
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

	candidateDigest, err := getDigest(ctx, "getCandidateDigest")
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
