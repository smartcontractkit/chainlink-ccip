// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IMessageTransmitter} from "./interfaces/IMessageTransmitter.sol";
import {ITokenMessenger} from "./interfaces/ITokenMessenger.sol";
import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";

import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

/// @title CCTP Message Transmitter Proxy
/// @notice A proxy contract for handling messages transmitted via the Cross Chain Transfer Protocol (CCTP).
/// @dev This contract is responsible for sending messages to the `IMessageTransmitter` and ensuring only allowed callers can invoke it.
contract CCTPMessageTransmitterProxy is AuthorizedCallers, ITypeAndVersion {
  string public constant override typeAndVersion = "CCTPMessageTransmitterProxy 1.7.0-dev";

  /// @notice Immutable reference to the `IMessageTransmitter` contract.
  IMessageTransmitter public immutable i_cctpTransmitter;

  /// @notice One-time cyclic dependency between TokenPool and MessageTransmitter.
  /// @param tokenMessenger The Circle TokenMessenger used to resolve the local MessageTransmitter.
  constructor(
    ITokenMessenger tokenMessenger
  ) AuthorizedCallers(new address[](0)) {
    i_cctpTransmitter = IMessageTransmitter(tokenMessenger.localMessageTransmitter());
  }

  /// @notice Receives a message from the `IMessageTransmitter` contract and validates it.
  /// @dev Can only be called by an allowed caller to process incoming messages.
  /// @param message The payload of the message being received.
  /// @param attestation The cryptographic proof validating the message.
  /// @return success A boolean indicating if the message was successfully processed.
  function receiveMessage(
    bytes calldata message,
    bytes calldata attestation
  ) external onlyAuthorizedCallers returns (bool success) {
    return i_cctpTransmitter.receiveMessage(message, attestation);
  }
}
