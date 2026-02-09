package v1_0

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	"give-me-state-v2/views"
	"give-me-state-v2/views/evm/common"

	gethCommon "github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

const capabilitiesRegistryABIJson = `[{"inputs":[],"name":"getNodes","outputs":[{"components":[{"internalType":"uint32","name":"nodeOperatorId","type":"uint32"},{"internalType":"uint32","name":"configCount","type":"uint32"},{"internalType":"uint32","name":"workflowDONId","type":"uint32"},{"internalType":"bytes32","name":"signer","type":"bytes32"},{"internalType":"bytes32","name":"p2pId","type":"bytes32"},{"internalType":"bytes32","name":"encryptionPublicKey","type":"bytes32"},{"internalType":"bytes32[]","name":"hashedCapabilityIds","type":"bytes32[]"},{"internalType":"uint256[]","name":"capabilitiesDONIds","type":"uint256[]"}],"internalType":"struct INodeInfoProvider.NodeInfo[]","name":"","type":"tuple[]"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"getNodeOperators","outputs":[{"components":[{"internalType":"address","name":"admin","type":"address"},{"internalType":"string","name":"name","type":"string"}],"internalType":"struct INodeInfoProvider.NodeOperator[]","name":"","type":"tuple[]"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"owner","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"typeAndVersion","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"}]`

var capRegABI abi.ABI

func init() {
	var err error
	capRegABI, err = abi.JSON(strings.NewReader(capabilitiesRegistryABIJson))
	if err != nil {
		panic("Failed to parse CapabilitiesRegistry ABI: " + err.Error())
	}
}

// executeCapRegCall packs a call, executes it, and returns raw response bytes.
func executeCapRegCall(ctx *views.ViewContext, method string, args ...interface{}) ([]byte, error) {
	calldata, err := capRegABI.Pack(method, args...)
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

// getNodeOperators fetches all node operators using ABI bindings.
func getNodeOperators(ctx *views.ViewContext) ([]map[string]any, error) {
	data, err := executeCapRegCall(ctx, "getNodeOperators")
	if err != nil {
		return nil, err
	}

	results, err := capRegABI.Unpack("getNodeOperators", data)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack getNodeOperators: %w", err)
	}
	if len(results) == 0 {
		return []map[string]any{}, nil
	}

	// The result is []struct{Admin common.Address; Name string}
	operators, ok := results[0].([]struct {
		Admin gethCommon.Address `json:"admin"`
		Name  string            `json:"name"`
	})
	if !ok {
		return nil, fmt.Errorf("unexpected type for node operators: %T", results[0])
	}

	out := make([]map[string]any, len(operators))
	for i, op := range operators {
		out[i] = map[string]any{
			"admin": op.Admin.Hex(),
			"name":  op.Name,
		}
	}
	return out, nil
}

// getNodes fetches all nodes using ABI bindings.
func getNodes(ctx *views.ViewContext) ([]map[string]any, error) {
	data, err := executeCapRegCall(ctx, "getNodes")
	if err != nil {
		return nil, err
	}

	results, err := capRegABI.Unpack("getNodes", data)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack getNodes: %w", err)
	}
	if len(results) == 0 {
		return []map[string]any{}, nil
	}

	// The result is []NodeInfo
	nodes, ok := results[0].([]struct {
		NodeOperatorId      uint32     `json:"nodeOperatorId"`
		ConfigCount         uint32     `json:"configCount"`
		WorkflowDONId       uint32     `json:"workflowDONId"`
		Signer              [32]byte   `json:"signer"`
		P2pId               [32]byte   `json:"p2pId"`
		EncryptionPublicKey [32]byte   `json:"encryptionPublicKey"`
		HashedCapabilityIds [][32]byte `json:"hashedCapabilityIds"`
		CapabilitiesDONIds  []*big.Int `json:"capabilitiesDONIds"`
	})
	if !ok {
		return nil, fmt.Errorf("unexpected type for nodes: %T", results[0])
	}

	out := make([]map[string]any, len(nodes))
	for i, n := range nodes {
		hashedCapIds := make([]string, len(n.HashedCapabilityIds))
		for j, h := range n.HashedCapabilityIds {
			hashedCapIds[j] = "0x" + hex.EncodeToString(h[:])
		}

		capDonIds := make([]uint64, len(n.CapabilitiesDONIds))
		for j, id := range n.CapabilitiesDONIds {
			capDonIds[j] = id.Uint64()
		}

		out[i] = map[string]any{
			"nodeOperatorId":      n.NodeOperatorId,
			"configCount":         n.ConfigCount,
			"workflowDONId":       n.WorkflowDONId,
			"signer":              "0x" + hex.EncodeToString(n.Signer[:]),
			"p2pId":               "0x" + hex.EncodeToString(n.P2pId[:]),
			"encryptionPublicKey": "0x" + hex.EncodeToString(n.EncryptionPublicKey[:]),
			"hashedCapabilityIds": hashedCapIds,
			"capabilitiesDONIds":  capDonIds,
		}
	}
	return out, nil
}

// ViewCapabilitiesRegistry generates a view of the CapabilitiesRegistry contract (v1.0.0).
// Uses bespoke ABI JSON for proper struct decoding (no Go bindings available).
func ViewCapabilitiesRegistry(ctx *views.ViewContext) (map[string]any, error) {
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

	nodeOperators, err := getNodeOperators(ctx)
	if err != nil {
		result["nodeOperators_error"] = err.Error()
	} else {
		result["nodeOperators"] = nodeOperators
	}

	nodes, err := getNodes(ctx)
	if err != nil {
		result["nodes_error"] = err.Error()
	} else {
		result["nodes"] = nodes
	}

	return result, nil
}
