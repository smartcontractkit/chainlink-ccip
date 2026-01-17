// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {Internal} from "./Internal.sol";

library SuperchainInterop {
  /// @notice The custom event used to relay messages over superchain.
  /// @param destChainSelector The destination chain selector.
  /// @param sequenceNumber The sequence number of the message.
  /// @param message The message to relay, already converted to the Any2EVM format.
  event CCIPSuperchainMessageSent(
    uint64 indexed destChainSelector, uint64 indexed sequenceNumber, Internal.Any2EVMRampMessage message
  );

  /// @notice Generate an unique hash for an Any2EVMRampMessage.
  /// @dev This is similar to how messageId is generated in the OnRamp, but using the Any2EVMRampMessage
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
