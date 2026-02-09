package v1_6

import (
	"fmt"

	"give-me-state-v2/views"

	"github.com/ethereum/go-ethereum/common"
)

// packNonceManagerCall packs a method call using the NonceManager v1.6 ABI.
func packNonceManagerCall(method string, args ...interface{}) ([]byte, error) {
	return NonceManagerABI.Pack(method, args...)
}

// executeNonceManagerCall packs a call, executes it, and returns raw response bytes.
func executeNonceManagerCall(ctx *views.ViewContext, method string, args ...interface{}) ([]byte, error) {
	calldata, err := packNonceManagerCall(method, args...)
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

// getNonceManagerOwner fetches the owner address.
func getNonceManagerOwner(ctx *views.ViewContext) (string, error) {
	data, err := executeNonceManagerCall(ctx, "owner")
	if err != nil {
		return "", err
	}
	results, err := NonceManagerABI.Unpack("owner", data)
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

// getNonceManagerTypeAndVersion fetches the typeAndVersion string.
func getNonceManagerTypeAndVersion(ctx *views.ViewContext) (string, error) {
	data, err := executeNonceManagerCall(ctx, "typeAndVersion")
	if err != nil {
		return "", err
	}
	results, err := NonceManagerABI.Unpack("typeAndVersion", data)
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

// getNonceManagerAuthorizedCallers fetches all authorized callers using ABI bindings.
func getNonceManagerAuthorizedCallers(ctx *views.ViewContext) ([]string, error) {
	data, err := executeNonceManagerCall(ctx, "getAllAuthorizedCallers")
	if err != nil {
		return nil, err
	}

	results, err := NonceManagerABI.Unpack("getAllAuthorizedCallers", data)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack getAllAuthorizedCallers: %w", err)
	}
	if len(results) == 0 {
		return []string{}, nil
	}

	addrs, ok := results[0].([]common.Address)
	if !ok {
		return nil, fmt.Errorf("unexpected type for authorized callers: %T", results[0])
	}

	callers := make([]string, len(addrs))
	for i, a := range addrs {
		callers[i] = a.Hex()
	}
	return callers, nil
}

// ViewNonceManager generates a view of the NonceManager contract (v1.6.0).
// Uses ABI bindings for proper decoding.
func ViewNonceManager(ctx *views.ViewContext) (map[string]any, error) {
	result := make(map[string]any)

	result["address"] = ctx.AddressHex
	result["chainSelector"] = ctx.ChainSelector
	result["version"] = "1.6.0"

	owner, err := getNonceManagerOwner(ctx)
	if err != nil {
		result["owner_error"] = err.Error()
	} else {
		result["owner"] = owner
	}

	typeAndVersion, err := getNonceManagerTypeAndVersion(ctx)
	if err != nil {
		result["typeAndVersion_error"] = err.Error()
	} else {
		result["typeAndVersion"] = typeAndVersion
	}

	authorizedCallers, err := getNonceManagerAuthorizedCallers(ctx)
	if err != nil {
		result["authorizedCallers_error"] = err.Error()
	} else {
		result["authorizedCallers"] = authorizedCallers
	}

	return result, nil
}
