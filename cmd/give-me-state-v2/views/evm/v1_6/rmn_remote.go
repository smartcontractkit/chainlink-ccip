package v1_6

import (
	"encoding/hex"
	"fmt"

	"give-me-state-v2/views"

	"github.com/ethereum/go-ethereum/common"
)

// packRMNRemoteCall packs a method call using the RMNRemote v1.6 ABI.
func packRMNRemoteCall(method string, args ...interface{}) ([]byte, error) {
	return RMNRemoteABI.Pack(method, args...)
}

// executeRMNRemoteCall packs a call, executes it, and returns raw response bytes.
func executeRMNRemoteCall(ctx *views.ViewContext, method string, args ...interface{}) ([]byte, error) {
	calldata, err := packRMNRemoteCall(method, args...)
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

// getRMNRemoteOwner fetches the owner address.
func getRMNRemoteOwner(ctx *views.ViewContext) (string, error) {
	data, err := executeRMNRemoteCall(ctx, "owner")
	if err != nil {
		return "", err
	}
	results, err := RMNRemoteABI.Unpack("owner", data)
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

// getRMNRemoteTypeAndVersion fetches the typeAndVersion string.
func getRMNRemoteTypeAndVersion(ctx *views.ViewContext) (string, error) {
	data, err := executeRMNRemoteCall(ctx, "typeAndVersion")
	if err != nil {
		return "", err
	}
	results, err := RMNRemoteABI.Unpack("typeAndVersion", data)
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

// getRMNRemoteIsCursed fetches the global curse status using ABI bindings.
func getRMNRemoteIsCursed(ctx *views.ViewContext) (bool, error) {
	data, err := executeRMNRemoteCall(ctx, "isCursed0")
	if err != nil {
		return false, err
	}
	results, err := RMNRemoteABI.Unpack("isCursed0", data)
	if err != nil {
		return false, fmt.Errorf("failed to unpack isCursed: %w", err)
	}
	if len(results) == 0 {
		return false, fmt.Errorf("no results from isCursed call")
	}
	cursed, ok := results[0].(bool)
	if !ok {
		return false, fmt.Errorf("unexpected type for isCursed: %T", results[0])
	}
	return cursed, nil
}

// getRMNRemoteVersionedConfig fetches the versioned config (signers, fSign, etc).
func getRMNRemoteVersionedConfig(ctx *views.ViewContext) (map[string]any, error) {
	data, err := executeRMNRemoteCall(ctx, "getVersionedConfig")
	if err != nil {
		return nil, err
	}

	results, err := RMNRemoteABI.Unpack("getVersionedConfig", data)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack getVersionedConfig: %w", err)
	}
	if len(results) < 2 {
		return nil, fmt.Errorf("expected 2 results from getVersionedConfig, got %d", len(results))
	}

	version, ok := results[0].(uint32)
	if !ok {
		return nil, fmt.Errorf("unexpected type for version: %T", results[0])
	}

	cfg, ok := results[1].(struct {
		RmnHomeContractConfigDigest [32]byte `json:"rmnHomeContractConfigDigest"`
		Signers                     []struct {
			OnchainPublicKey common.Address `json:"onchainPublicKey"`
			NodeIndex        uint64         `json:"nodeIndex"`
		} `json:"signers"`
		FSign uint64 `json:"fSign"`
	})
	if !ok {
		return nil, fmt.Errorf("unexpected type for RMNRemoteConfig: %T", results[1])
	}

	signers := make([]map[string]any, len(cfg.Signers))
	for i, s := range cfg.Signers {
		signers[i] = map[string]any{
			"onchainPublicKey": s.OnchainPublicKey.Hex(),
			"nodeIndex":        s.NodeIndex,
		}
	}

	return map[string]any{
		"version":                     version,
		"rmnHomeContractConfigDigest": "0x" + hex.EncodeToString(cfg.RmnHomeContractConfigDigest[:]),
		"signers":                     signers,
		"fSign":                       cfg.FSign,
	}, nil
}

// ViewRMNRemote generates a view of the RMNRemote contract (v1.6.0).
// Uses ABI bindings for proper struct decoding.
func ViewRMNRemote(ctx *views.ViewContext) (map[string]any, error) {
	result := make(map[string]any)

	result["address"] = ctx.AddressHex
	result["chainSelector"] = ctx.ChainSelector
	result["version"] = "1.6.0"

	owner, err := getRMNRemoteOwner(ctx)
	if err != nil {
		result["owner_error"] = err.Error()
	} else {
		result["owner"] = owner
	}

	typeAndVersion, err := getRMNRemoteTypeAndVersion(ctx)
	if err != nil {
		result["typeAndVersion_error"] = err.Error()
	} else {
		result["typeAndVersion"] = typeAndVersion
	}

	isCursed, err := getRMNRemoteIsCursed(ctx)
	if err != nil {
		result["isCursed_error"] = err.Error()
	} else {
		result["isCursed"] = isCursed
	}

	versionedConfig, err := getRMNRemoteVersionedConfig(ctx)
	if err != nil {
		result["versionedConfig_error"] = err.Error()
	} else {
		result["versionedConfig"] = versionedConfig
	}

	return result, nil
}
