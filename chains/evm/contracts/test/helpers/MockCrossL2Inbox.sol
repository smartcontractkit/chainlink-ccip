// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ICrossL2Inbox} from "../../vendor/optimism/interop-lib/v0/src/interfaces/ICrossL2Inbox.sol";
import {Identifier} from "../../vendor/optimism/interop-lib/v0/src/interfaces/IIdentifier.sol";

contract MockCrossL2Inbox is ICrossL2Inbox {
  mapping(bytes32 msgHash => bool isValid) public s_validMessages;
  bool public s_shouldRevert;

  function validateMessage(Identifier calldata _id, bytes32 _msgHash) external {
    if (s_shouldRevert) {
      revert("CrossL2Inbox: validation failed");
    }
    if (!s_validMessages[_msgHash]) {
      revert("Invalid message");
    }
    emit ExecutingMessage(_msgHash, _id);
  }

  function setValidMessage(bytes32 _msgHash, bool _valid) external {
    s_validMessages[_msgHash] = _valid;
  }

  function setShouldRevert(bool _shouldRevert) external {
    s_shouldRevert = _shouldRevert;
  }

  function calculateChecksum(Identifier memory _id, bytes32 _msgHash) external pure returns (bytes32) {
    return keccak256(abi.encode(_id, _msgHash));
  }
}