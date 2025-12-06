// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

interface IMailbox {
  function deliverAndHandle(
    bytes calldata rawPayload,
    bytes calldata proof
  ) external returns (bytes32, bool executed, bytes memory optionalMessage);
}
