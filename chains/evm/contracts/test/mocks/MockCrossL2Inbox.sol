// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ICrossL2Inbox} from "../../vendor/optimism/interop-lib/v0/src/interfaces/ICrossL2Inbox.sol";
import {Identifier} from "../../vendor/optimism/interop-lib/v0/src/interfaces/IIdentifier.sol";

contract MockCrossL2Inbox is ICrossL2Inbox {
  error ValidationFailed(string reason);
  error UnexpectedCall(bytes32 idHash, bytes32 msgHash);

  struct ValidateMessageCall {
    Identifier identifier;
    bytes32 msgHash;
  }

  ValidateMessageCall[] public s_validateMessageCalls;
  mapping(bytes32 identifierHash => mapping(bytes32 msgHash => bool success)) public s_validationSuccesses;
  mapping(bytes32 identifierHash => mapping(bytes32 msgHash => string revertMessage)) public s_validationErrors;

  function validateMessage(Identifier calldata identifier, bytes32 msgHash) external {
    s_validateMessageCalls.push(ValidateMessageCall({identifier: identifier, msgHash: msgHash}));

    bytes32 idHash = keccak256(abi.encode(identifier));

    // Check if there's a specific error set for this identifier and msgHash
    if (bytes(s_validationErrors[idHash][msgHash]).length > 0) {
      revert ValidationFailed(s_validationErrors[idHash][msgHash]);
    }

    // Check if there's a specific success result
    if (s_validationSuccesses[idHash][msgHash]) {
      return;
    }

    revert UnexpectedCall(idHash, msgHash);
  }

  function setValidationSuccess(Identifier memory identifier, bytes32 msgHash) external {
    bytes32 idHash = keccak256(abi.encode(identifier));
    s_validationSuccesses[idHash][msgHash] = true;
  }

  function setValidationError(Identifier memory identifier, bytes32 msgHash, string memory revertMessage) external {
    bytes32 idHash = keccak256(abi.encode(identifier));
    s_validationErrors[idHash][msgHash] = revertMessage;
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

  function calculateChecksum(Identifier memory identifier, bytes32 msgHash) external pure returns (bytes32 checksum_) {
    return keccak256(abi.encode(identifier, msgHash));
  }
}
