package v1_0

import (
	"fmt"

	"give-me-state-v2/views"
	"give-me-state-v2/views/evm/common"

	gethCommon "github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/accounts/abi"

	rmn_proxy "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/rmn_proxy_contract"
)

// Parsed ABI for RMNProxy
var RMNProxyABI abi.ABI

func init() {
	parsed, err := rmn_proxy.RMNProxyMetaData.GetAbi()
	if err != nil {
		panic(fmt.Sprintf("failed to parse RMNProxy ABI: %v", err))
	}
	RMNProxyABI = *parsed
}

// executeRMNProxyCall packs a call, executes it, and returns raw response bytes.
func executeRMNProxyCall(ctx *views.ViewContext, method string, args ...interface{}) ([]byte, error) {
	calldata, err := RMNProxyABI.Pack(method, args...)
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

// getRMNProxyARM fetches the ARM (RMN) address using ABI bindings.
func getRMNProxyARM(ctx *views.ViewContext) (string, error) {
	data, err := executeRMNProxyCall(ctx, "getARM")
	if err != nil {
		return "", err
	}
	results, err := RMNProxyABI.Unpack("getARM", data)
	if err != nil {
		return "", fmt.Errorf("failed to unpack getARM: %w", err)
	}
	if len(results) == 0 {
		return "", fmt.Errorf("no results from getARM call")
	}
	addr, ok := results[0].(gethCommon.Address)
	if !ok {
		return "", fmt.Errorf("unexpected type for ARM: %T", results[0])
	}
	return addr.Hex(), nil
}

// ViewRMNProxy generates a view of the RMNProxy (ARMProxy) contract (v1.0.0).
// Uses ABI bindings for proper decoding.
func ViewRMNProxy(ctx *views.ViewContext) (map[string]any, error) {
	result := make(map[string]any)

	result["address"] = ctx.AddressHex
	result["chainSelector"] = ctx.ChainSelector
	result["version"] = "1.0.0"

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

	arm, err := getRMNProxyARM(ctx)
	if err != nil {
		result["arm_error"] = err.Error()
	} else {
		result["arm"] = arm
	}

	return result, nil
}
