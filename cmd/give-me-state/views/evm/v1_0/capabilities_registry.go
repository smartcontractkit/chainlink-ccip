package v1_0

import (
	"call-orchestrator-demo/views"
	"call-orchestrator-demo/views/evm/common"
	"encoding/hex"
)

// Function selectors for CapabilitiesRegistry v1.0
var (
	// getNodes() returns (INodeInfoProvider.NodeInfo[] memory)
	selectorGetNodes = common.HexToSelector("e29581aa")
	// getNodeOperators() returns (INodeInfoProvider.NodeOperator[] memory)
	selectorGetNodeOperators = common.HexToSelector("66acaa33")
)

// ViewCapabilitiesRegistry generates a view of the CapabilitiesRegistry contract (v1.0.0).
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

	// Get node operators
	nodeOperators, err := getNodeOperators(ctx)
	if err != nil {
		result["nodeOperators_error"] = err.Error()
	} else {
		result["nodeOperators"] = nodeOperators
	}

	// Get nodes
	nodes, err := getNodes(ctx)
	if err != nil {
		result["nodes_error"] = err.Error()
	} else {
		result["nodes"] = nodes
	}

	return result, nil
}

// getNodeOperators fetches all node operators.
// Returns: INodeInfoProvider.NodeOperator[] memory
// struct NodeOperator { address admin; string name; }
func getNodeOperators(ctx *views.ViewContext) ([]map[string]any, error) {
	data, err := common.ExecuteCall(ctx, selectorGetNodeOperators)
	if err != nil {
		return nil, err
	}

	return decodeNodeOperatorsArray(data)
}

// getNodes fetches all nodes.
// Returns: INodeInfoProvider.NodeInfo[] memory
//
//	struct NodeInfo {
//	  uint32 nodeOperatorId;
//	  uint32 configCount;
//	  uint32 workflowDONId;
//	  bytes32 signer;
//	  bytes32 p2pId;
//	  bytes32 encryptionPublicKey;
//	  bytes32[] hashedCapabilityIds;
//	  uint256[] capabilitiesDONIds;
//	}
func getNodes(ctx *views.ViewContext) ([]map[string]any, error) {
	data, err := common.ExecuteCall(ctx, selectorGetNodes)
	if err != nil {
		return nil, err
	}

	return decodeNodesArray(data)
}

// decodeNodeOperatorsArray decodes NodeOperator[] memory
// struct NodeOperator { address admin; string name; }
//
// ABI layout for dynamic array of structs with dynamic members:
// [0x00]: offset to array data (typically 0x20)
// [arrayOffset]: length of array
// [arrayOffset + 0x20 + i*0x20]: offset to element[i] (relative to offset table start, i.e., arrayOffset + 0x20)
//
// Each element (struct with dynamic member):
// [elemPos + 0x00]: address admin (32 bytes)
// [elemPos + 0x20]: offset to string data (relative to elemPos)
// [elemPos + strOffset]: string length
// [elemPos + strOffset + 0x20]: string bytes
func decodeNodeOperatorsArray(data []byte) ([]map[string]any, error) {
	if len(data) < 64 {
		return []map[string]any{}, nil
	}

	// Step 1: Get offset to array data
	arrayOffset := common.DecodeUint64FromBytes(data[0:32])
	if arrayOffset+32 > uint64(len(data)) {
		return []map[string]any{}, nil
	}

	// Step 2: Get array length
	length := common.DecodeUint64FromBytes(data[arrayOffset : arrayOffset+32])
	if length == 0 || length > 10000 {
		return []map[string]any{}, nil
	}

	operators := make([]map[string]any, 0, length)

	// Step 3: Iterate through each element
	for i := uint64(0); i < length; i++ {
		// Get offset to this element (relative to arrayOffset)
		elemOffsetPos := arrayOffset + 32 + i*32
		if elemOffsetPos+32 > uint64(len(data)) {
			break
		}
		elemOffsetRel := common.DecodeUint64FromBytes(data[elemOffsetPos : elemOffsetPos+32])
		// CRITICAL: offsets are relative to the offset table start (arrayOffset + 32), not arrayOffset!
		elemPos := arrayOffset + 32 + elemOffsetRel

		if elemPos+64 > uint64(len(data)) {
			break
		}

		result := make(map[string]any)

		// Slot 0: address admin
		adminAddr, err := common.DecodeAddress(data[elemPos : elemPos+32])
		if err == nil {
			result["admin"] = adminAddr
		}

		// Slot 1: offset to string (relative to elemPos)
		stringOffsetRel := common.DecodeUint64FromBytes(data[elemPos+32 : elemPos+64])
		stringPos := elemPos + stringOffsetRel

		// Decode string: [length][data]
		if stringPos+32 <= uint64(len(data)) {
			strLen := common.DecodeUint64FromBytes(data[stringPos : stringPos+32])
			if strLen > 0 && strLen < 1000 && stringPos+32+strLen <= uint64(len(data)) {
				result["name"] = string(data[stringPos+32 : stringPos+32+strLen])
			} else {
				result["name"] = ""
			}
		} else {
			result["name"] = ""
		}

		operators = append(operators, result)
	}

	return operators, nil
}

// decodeNodesArray decodes NodeInfo[] memory
func decodeNodesArray(data []byte) ([]map[string]any, error) {
	if len(data) < 64 {
		return []map[string]any{}, nil
	}

	// First 32 bytes is offset to array data
	offset := common.DecodeUint64FromBytes(data[0:32])
	if offset+32 > uint64(len(data)) {
		return []map[string]any{}, nil
	}

	// Length of the array
	length := common.DecodeUint64FromBytes(data[offset : offset+32])
	if length == 0 {
		return []map[string]any{}, nil
	}

	nodes := make([]map[string]any, 0, length)

	// After length, we have offsets to each NodeInfo element
	offsetsStart := offset + 32
	for i := uint64(0); i < length; i++ {
		if offsetsStart+i*32+32 > uint64(len(data)) {
			break
		}

		// Get the offset to this element (relative to the start of the array data)
		elementOffset := common.DecodeUint64FromBytes(data[offsetsStart+i*32 : offsetsStart+i*32+32])
		actualOffset := offset + elementOffset

		if actualOffset+256 > uint64(len(data)) {
			break
		}

		node, err := decodeNodeInfo(data, actualOffset)
		if err == nil && node != nil {
			nodes = append(nodes, node)
		}
	}

	return nodes, nil
}

// decodeNodeInfo decodes a single NodeInfo from the data at the given offset.
//
//	struct NodeInfo {
//	  uint32 nodeOperatorId;
//	  uint32 configCount;
//	  uint32 workflowDONId;
//	  bytes32 signer;
//	  bytes32 p2pId;
//	  bytes32 encryptionPublicKey;
//	  bytes32[] hashedCapabilityIds;
//	  uint256[] capabilitiesDONIds;
//	}
func decodeNodeInfo(data []byte, offset uint64) (map[string]any, error) {
	if offset+256 > uint64(len(data)) {
		return nil, nil
	}

	result := make(map[string]any)

	// nodeOperatorId (uint32 - padded to 32 bytes)
	result["nodeOperatorId"] = uint32(common.DecodeUint64FromBytes(data[offset : offset+32]))

	// configCount (uint32 - padded to 32 bytes)
	result["configCount"] = uint32(common.DecodeUint64FromBytes(data[offset+32 : offset+64]))

	// workflowDONId (uint32 - padded to 32 bytes)
	result["workflowDONId"] = uint32(common.DecodeUint64FromBytes(data[offset+64 : offset+96]))

	// signer (bytes32)
	result["signer"] = "0x" + hex.EncodeToString(data[offset+96:offset+128])

	// p2pId (bytes32)
	result["p2pId"] = "0x" + hex.EncodeToString(data[offset+128:offset+160])

	// encryptionPublicKey (bytes32)
	result["encryptionPublicKey"] = "0x" + hex.EncodeToString(data[offset+160:offset+192])

	// hashedCapabilityIds (bytes32[] - offset)
	hashedCapIdsOffset := common.DecodeUint64FromBytes(data[offset+192 : offset+224])
	actualHashedCapIdsOffset := offset + hashedCapIdsOffset
	hashedCapIds, err := decodeBytes32Array(data, actualHashedCapIdsOffset)
	if err == nil {
		result["hashedCapabilityIds"] = hashedCapIds
	}

	// capabilitiesDONIds (uint256[] - offset)
	capDONIdsOffset := common.DecodeUint64FromBytes(data[offset+224 : offset+256])
	actualCapDONIdsOffset := offset + capDONIdsOffset
	capDONIds, err := decodeUint256Array(data, actualCapDONIdsOffset)
	if err == nil {
		result["capabilitiesDONIds"] = capDONIds
	}

	return result, nil
}

// decodeBytes32Array decodes a bytes32[] array.
func decodeBytes32Array(data []byte, offset uint64) ([]string, error) {
	if offset+32 > uint64(len(data)) {
		return []string{}, nil
	}

	length := common.DecodeUint64FromBytes(data[offset : offset+32])
	if length == 0 {
		return []string{}, nil
	}

	result := make([]string, 0, length)
	for i := uint64(0); i < length; i++ {
		elemOffset := offset + 32 + i*32
		if elemOffset+32 > uint64(len(data)) {
			break
		}
		result = append(result, "0x"+hex.EncodeToString(data[elemOffset:elemOffset+32]))
	}

	return result, nil
}

// decodeUint256Array decodes a uint256[] array.
func decodeUint256Array(data []byte, offset uint64) ([]uint64, error) {
	if offset+32 > uint64(len(data)) {
		return []uint64{}, nil
	}

	length := common.DecodeUint64FromBytes(data[offset : offset+32])
	if length == 0 {
		return []uint64{}, nil
	}

	result := make([]uint64, 0, length)
	for i := uint64(0); i < length; i++ {
		elemOffset := offset + 32 + i*32
		if elemOffset+32 > uint64(len(data)) {
			break
		}
		// Decode as uint64 (safe for DON IDs which are typically small)
		result = append(result, common.DecodeUint64FromBytes(data[elemOffset:elemOffset+32]))
	}

	return result, nil
}

// decodeStringAtOffset decodes a string at a specific offset.
// The offset points to the length field of the string.
func decodeStringAtOffset(data []byte, offset uint64) (string, error) {
	if offset+32 > uint64(len(data)) {
		return "", nil
	}

	// Read length
	length := common.DecodeUint64FromBytes(data[offset : offset+32])
	if length == 0 {
		return "", nil
	}

	if offset+32+length > uint64(len(data)) {
		return "", nil
	}

	// Read string data
	return string(data[offset+32 : offset+32+length]), nil
}
