// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {Internal} from "./Internal.sol";

import {Identifier} from "../vendor/optimism/interop-lib/v0/src/interfaces/IIdentifier.sol";

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

  /// @notice Generate an unique hash for an Any2EVMRampMessage.
  /// @dev This is similiar to how messageId is generated in the OnRamp, but using the Any2EVMRampMessage
  /// type, and OffRamp metadata hash. This gives a unique identifier for the message that can be derived
  /// in both the On/OffRampOverSuperchainInterop.
  /// @param message The interop message to hash.
  /// @return messageHash The hash of the interop message.
  function _hashInteropMessage(
    Internal.Any2EVMRampMessage memory message,
    address onRamp
  ) internal pure returns (bytes32) {
    bytes32 offRampMetaDataHash = keccak256(
      abi.encode(
        Internal.ANY_2_EVM_MESSAGE_HASH,
        message.header.sourceChainSelector,
        message.header.destChainSelector,
        keccak256(abi.encode(onRamp))
      )
    );

    return Internal._hash(message, offRampMetaDataHash);
  }
}
