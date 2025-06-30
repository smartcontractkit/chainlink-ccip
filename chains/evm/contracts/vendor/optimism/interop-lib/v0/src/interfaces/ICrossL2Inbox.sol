// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {Identifier} from "./IIdentifier.sol";

/// @title ICrossL2Inbox
/// @notice Interface for the CrossL2Inbox contract.
/// @dev This is a copy of the ICrossL2Inbox interface from the ethereum-optimism/interop-lib repo.
/// This interface is post-audit and has been stable for many months, but has not been officially
/// included in a versioned release.
interface ICrossL2Inbox {
  error ReentrantCall();

  /// @notice Thrown when the caller is not DEPOSITOR_ACCOUNT when calling `setInteropStart()`
  error NotDepositor();

  /// @notice Thrown when attempting to set interop start when it's already set.
  error InteropStartAlreadySet();

  /// @notice Thrown when a non-written transient storage slot is attempted to be read from.
  error NotEntered();

  /// @notice Thrown when trying to execute a cross chain message on a deposit transaction.
  error NoExecutingDeposits();

  event ExecutingMessage(bytes32 indexed msgHash, Identifier id);

  /// @notice Validates a cross chain message on the destination chain
  ///         and emits an ExecutingMessage event. This function is useful
  ///         for applications that understand the schema of the _message payload and want to
  ///         process it in a custom way.
  /// @param _id      Identifier of the message.
  /// @param _msgHash Hash of the message payload to call target with.
  function validateMessage(Identifier calldata _id, bytes32 _msgHash) external;

  /// @notice Calculates a custom checksum for a cross chain message `Identifier` and `msgHash`.
  /// @param _id The identifier of the message.
  /// @param _msgHash The hash of the message.
  /// @return checksum_ The checksum of the message.
  function calculateChecksum(Identifier memory _id, bytes32 _msgHash) external pure returns (bytes32 checksum_);
}
