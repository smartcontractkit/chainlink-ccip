// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {IMailbox} from "../../pools/Lombard/interfaces/IMailbox.sol";

contract MockMailbox is IMailbox {
  bytes32 internal s_payloadHash;
  bool internal s_executed = true;
  bytes internal s_executionResult;

  bytes public s_lastRawPayload;
  bytes public s_lastProof;

  function setResult(bytes32 payloadHash, bool executed, bytes calldata executionResult) external {
    s_payloadHash = payloadHash;
    s_executed = executed;
    s_executionResult = executionResult;
  }

  function deliverAndHandle(
    bytes calldata rawPayload,
    bytes calldata proof
  ) external returns (bytes32, bool, bytes memory) {
    s_lastRawPayload = rawPayload;
    s_lastProof = proof;
    bytes32 payloadHash = s_payloadHash != bytes32(0) ? s_payloadHash : keccak256(rawPayload);
    return (payloadHash, s_executed, s_executionResult);
  }
}
