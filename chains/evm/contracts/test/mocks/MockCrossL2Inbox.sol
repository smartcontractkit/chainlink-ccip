// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ICrossL2Inbox} from "../../vendor/optimism/interop-lib/v0/src/interfaces/ICrossL2Inbox.sol";
import {Identifier} from "../../vendor/optimism/interop-lib/v0/src/interfaces/IIdentifier.sol";

contract MockCrossL2Inbox is ICrossL2Inbox {
  error ValidationFailed(string reason);

  struct ValidateMessageCall {
    Identifier identifier;
    bytes32 msgHash;
  }

  ValidateMessageCall[] public s_validateMessageCalls;
  mapping(bytes32 => mapping(bytes32 => bool)) public s_validationResults;
  mapping(bytes32 => mapping(bytes32 => string)) public s_validationErrors;
  bool public s_defaultValidationResult = true;
  string public s_defaultValidationError = "Validation failed";

  function validateMessage(Identifier calldata _id, bytes32 _msgHash) external {
    s_validateMessageCalls.push(ValidateMessageCall({identifier: _id, msgHash: _msgHash}));

    bytes32 idHash = keccak256(abi.encode(_id));

    // Check if there's a specific result set for this identifier and msgHash
    if (bytes(s_validationErrors[idHash][_msgHash]).length > 0) {
      revert ValidationFailed(s_validationErrors[idHash][_msgHash]);
    }

    // Check if there's a specific success result
    if (s_validationResults[idHash][_msgHash]) {
      return;
    }

    // Use default behavior
    if (!s_defaultValidationResult) {
      revert ValidationFailed(s_defaultValidationError);
    }
  }

  function setValidationResult(Identifier memory _id, bytes32 _msgHash, bool _result) external {
    bytes32 idHash = keccak256(abi.encode(_id));
    s_validationResults[idHash][_msgHash] = _result;
  }

  function setValidationError(Identifier memory _id, bytes32 _msgHash, string memory _error) external {
    bytes32 idHash = keccak256(abi.encode(_id));
    s_validationErrors[idHash][_msgHash] = _error;
  }

  function setDefaultValidationResult(
    bool _result
  ) external {
    s_defaultValidationResult = _result;
  }

  function setDefaultValidationError(
    string memory _error
  ) external {
    s_defaultValidationError = _error;
  }

  function getValidateMessageCallCount() external view returns (uint256) {
    return s_validateMessageCalls.length;
  }

  function getValidateMessageCall(
    uint256 index
  ) external view returns (ValidateMessageCall memory) {
    return s_validateMessageCalls[index];
  }

  function resetValidateMessageCalls() external {
    delete s_validateMessageCalls;
  }

  function calculateChecksum(Identifier memory _id, bytes32 _msgHash) external pure returns (bytes32 checksum_) {
    return keccak256(abi.encode(_id, _msgHash));
  }
}
