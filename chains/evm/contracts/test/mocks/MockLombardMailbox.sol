// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {IMailbox} from "../../interfaces/lombard/IMailbox.sol";

contract MockLombardMailbox is IMailbox {
  bytes32 internal s_payloadHash;
  bool internal s_executed = true;
  bytes internal s_executionResult;

  bytes public s_lastRawPayload;

  function setResult(
    bytes32 payloadHash,
    bool executed,
    bytes calldata executionResult
  ) external {
    s_payloadHash = payloadHash;
    s_executed = executed;
    s_executionResult = executionResult;
  }

  /// @dev Alias used by MockLombardBridge to set the execution result (optional message).
  function setMessageId(
    bytes calldata optionalMessage
  ) external {
    s_executionResult = optionalMessage;
  }

  function setShouldSucceed(
    bool shouldSucceed
  ) external {
    s_executed = shouldSucceed;
  }

  function deliverAndHandle(
    bytes calldata rawPayload,
    bytes calldata
  ) external returns (bytes32, bool, bytes memory) {
    s_lastRawPayload = rawPayload;
    // It means that the bridge did not set any expectations for the payload, so we return a hash of the raw payload and
    // the raw payload itself as the optional message.
    if (s_payloadHash == bytes32(0) && s_executionResult.length == 0) {
      return (keccak256(rawPayload), true, rawPayload);
    }

    bytes32 payloadHash = s_payloadHash != bytes32(0) ? s_payloadHash : keccak256(rawPayload);
    return (payloadHash, s_executed, s_executionResult);
  }
}
