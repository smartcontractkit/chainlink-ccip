package v1_5

import (
	"fmt"

	"give-me-state-v2/views"

	"github.com/ethereum/go-ethereum/common"
)

// packRMNCall packs a method call using the RMNContract v1.5 ABI.
func packRMNCall(method string, args ...interface{}) ([]byte, error) {
	return RMNContractABI.Pack(method, args...)
}

// executeRMNCall packs a call, executes it, and returns raw response bytes.
func executeRMNCall(ctx *views.ViewContext, method string, args ...interface{}) ([]byte, error) {
	calldata, err := packRMNCall(method, args...)
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

// getRMNOwner fetches the owner address.
func getRMNOwner(ctx *views.ViewContext) (string, error) {
	data, err := executeRMNCall(ctx, "owner")
	if err != nil {
		return "", err
	}
	results, err := RMNContractABI.Unpack("owner", data)
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

// getRMNTypeAndVersion fetches the typeAndVersion string.
func getRMNTypeAndVersion(ctx *views.ViewContext) (string, error) {
	data, err := executeRMNCall(ctx, "typeAndVersion")
	if err != nil {
		return "", err
	}
	results, err := RMNContractABI.Unpack("typeAndVersion", data)
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

// getRMNConfigDetails fetches and decodes the config details using ABI bindings.
func getRMNConfigDetails(ctx *views.ViewContext) (map[string]any, error) {
	data, err := executeRMNCall(ctx, "getConfigDetails")
	if err != nil {
		return nil, err
	}

	results, err := RMNContractABI.Unpack("getConfigDetails", data)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack getConfigDetails: %w", err)
	}
	// Returns (uint32 version, uint32 blockNumber, Config memory config)
	if len(results) < 3 {
		return nil, fmt.Errorf("expected 3 results from getConfigDetails, got %d", len(results))
	}

	configVersion, _ := results[0].(uint32)
	blockNumber, _ := results[1].(uint32)

	// Config struct: {Voter[] voters, uint16 blessWeightThreshold, uint16 curseWeightThreshold}
	cfg, ok := results[2].(struct {
		Voters []struct {
			BlessVoteAddr common.Address `json:"blessVoteAddr"`
			CurseVoteAddr common.Address `json:"curseVoteAddr"`
			BlessWeight   uint8          `json:"blessWeight"`
			CurseWeight   uint8          `json:"curseWeight"`
		} `json:"voters"`
		BlessWeightThreshold uint16 `json:"blessWeightThreshold"`
		CurseWeightThreshold uint16 `json:"curseWeightThreshold"`
	})
	if !ok {
		return nil, fmt.Errorf("unexpected type for RMN Config: %T", results[2])
	}

	voters := make([]map[string]any, len(cfg.Voters))
	for i, v := range cfg.Voters {
		voters[i] = map[string]any{
			"blessVoteAddr": v.BlessVoteAddr.Hex(),
			"curseVoteAddr": v.CurseVoteAddr.Hex(),
			"blessWeight":   v.BlessWeight,
			"curseWeight":   v.CurseWeight,
		}
	}

	return map[string]any{
		"version":     configVersion,
		"blockNumber": blockNumber,
		"config": map[string]any{
			"voters":               voters,
			"blessWeightThreshold": cfg.BlessWeightThreshold,
			"curseWeightThreshold": cfg.CurseWeightThreshold,
		},
	}, nil
}

// ViewRMN generates a view of the RMN contract (v1.5.0).
// Uses ABI bindings for proper struct decoding.
func ViewRMN(ctx *views.ViewContext) (map[string]any, error) {
	result := make(map[string]any)

	result["address"] = ctx.AddressHex
	result["chainSelector"] = ctx.ChainSelector
	result["version"] = "1.5.0"

	owner, err := getRMNOwner(ctx)
	if err != nil {
		result["owner_error"] = err.Error()
	} else {
		result["owner"] = owner
	}

	typeAndVersion, err := getRMNTypeAndVersion(ctx)
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
