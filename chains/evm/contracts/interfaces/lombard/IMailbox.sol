// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

/// @notice Minimal interface to deliver and handle Lombard bridge payloads.
interface IMailbox {
  /// @notice Verifies and executes a bridged payload.
  /// @param payload Raw payload emitted on the source chain.
  /// @param proof Bridging proof for the payload.
  /// @return payloadHash Hash of the payload.
  /// @return executed True if the payload was successfully handled.
  /// @return returnData Optional data returned by the handler.
  function deliverAndHandle(
    bytes calldata payload,
    bytes calldata proof
  ) external returns (bytes32 payloadHash, bool executed, bytes memory returnData);
}
