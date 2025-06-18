// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {Internal} from "./Internal.sol";

import {Identifier} from "../interfaces/optimism/IIdentifier.sol";

library SuperchainInterop {
  /// @notice The custom event used to relay messages over superchain.
  /// @param destChainSelector The destination chain selector.
  /// @param sequenceNumber The sequence number of the message.
  /// @param message The message to relay, already converted to the Any2EVM format.
  event CCIPSuperchainMessageSent(
    uint64 indexed destChainSelector, uint64 indexed sequenceNumber, Internal.Any2EVMRampMessage message
  );

  /// @notice Exec report data used for OffRampOverSuperchainInterop message execution.
  struct ExecutionReport {
    /// @notice The complete log data of the CCIPSuperchainMessageSent event.
    /// 1st 32-byte word is selector, then indexed 32-byte words, then abi-encoded event data.
    bytes logData;
    Identifier identifier; // The metadata of the CCIPSuperchainMessageSent event.
    bytes[] offchainTokenData; // Offchain token attestation data for each token transfer.
  }
}
