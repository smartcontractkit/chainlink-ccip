// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IMessageTransmitter} from "./interfaces/IMessageTransmitter.sol";
import {ITokenMessenger} from "./interfaces/ITokenMessenger.sol";

import {USDCTokenPool} from "./USDCTokenPool.sol";
import {Ownable2StepMsgSender} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2StepMsgSender.sol";

import {EnumerableSet} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v5.0.2/contracts/utils/structs/EnumerableSet.sol";

/// @title CCTP Message Transmitter Proxy
/// @notice A proxy contract for handling messages transmitted via the Cross Chain Transfer Protocol (CCTP).
/// @dev This contract is responsible for sending messages to the `IMessageTransmitter` and ensuring only allowed callers can invoke it.
contract CCTPMessageTransmitterProxy is Ownable2StepMsgSender {
  using EnumerableSet for EnumerableSet.AddressSet;

  /// @notice Error thrown when a function is called by an unauthorized address.
  error Unauthorized(address caller);

  /// @notice Emitted when an allowed caller is added.
  event AllowedCallerAdded(address indexed caller);
  /// @notice Emitted when an allowed caller is removed.
  event AllowedCallerRemoved(address indexed caller);

  struct AllowedCallerConfigArgs {
    address caller;
    bool allowed;
  }

  /// @notice Immutable reference to the `IMessageTransmitter` contract.
  IMessageTransmitter public immutable i_cctpTransmitterV1;

  IMessageTransmitter public immutable i_cctpTransmitterV2;

  /// @notice Enumerable set of addresses allowed to call `receiveMessage`.
  EnumerableSet.AddressSet private s_allowedCallers;

  /// @notice One-time cyclic dependency between TokenPool and MessageTransmitter.
  /// @dev In CCTP V1 and V2 the transmitter addresses are different, so we must pass in two different ones
  /// to ensure that the message can be dispatched correctly on each version of a CCTP Message. This allows
  /// the contract to interact with both versions while only having to maintain a single proxy contract.
  /// @param cctpV1_tokenMessenger The token messenger contract for CCTP V1, Can be address(0) if CCTP-V1 is not
  /// deployed on the current chain and does not need to be supported.
  /// @param cctpV2_tokenMessenger The token messenger contract for CCTP V2, Can be address(0) if CCTP-V2 is not
  /// deployed on the current chain and does not need to be supported.
  constructor(ITokenMessenger cctpV1_tokenMessenger, ITokenMessenger cctpV2_tokenMessenger) {
    if (address(cctpV1_tokenMessenger) != address(0)) {
      i_cctpTransmitterV1 = IMessageTransmitter(cctpV1_tokenMessenger.localMessageTransmitter());
    }

    if (address(cctpV2_tokenMessenger) != address(0)) {
      i_cctpTransmitterV2 = IMessageTransmitter(cctpV2_tokenMessenger.localMessageTransmitter());
    }
  }

  /// @notice Receives a message from the `IMessageTransmitter` contract and validates it.
  /// @dev Can only be called by an allowed caller to process incoming messages.
  /// @param message The payload of the message being received.
  /// @param attestation The cryptographic proof validating the message.
  /// @param cctpVersion The version of the CCTP Message transmitter to use, as an enum as defined in USDCTokenPool
  /// and passed in by the corresponding token pool.
  /// @return success A boolean indicating if the message was successfully processed.
  function receiveMessage(
    bytes calldata message,
    bytes calldata attestation,
    USDCTokenPool.CCTPVersion cctpVersion
  ) external returns (bool success) {
    if (!s_allowedCallers.contains(msg.sender)) {
      revert Unauthorized(msg.sender);
    }

    // Get the correct transmitter address based on the CCTP Version. The else-if pattern simplifies adding future support
    // for a CCTP Version-3 if needed.
    IMessageTransmitter transmitter;
    if (cctpVersion == USDCTokenPool.CCTPVersion.VERSION_1) {
      transmitter = i_cctpTransmitterV1;
    } else if (cctpVersion == USDCTokenPool.CCTPVersion.VERSION_2) {
      transmitter = i_cctpTransmitterV2;
    }
    // If the message version is anything other than the supported ones, return false so that the tx will revert in the
    // USDC Token Pool due to USDCUnlockingFailed()
    else {
      return false;
    }

    // Dispatch the message to the transmitter to mint tokens and return the result.
    return transmitter.receiveMessage(message, attestation);
  }

  /// @notice Configures the allowed callers for the `receiveMessage` function.
  /// @param configArgs An array of `AllowedCallerConfigArgs` structs.
  function configureAllowedCallers(
    AllowedCallerConfigArgs[] calldata configArgs
  ) external onlyOwner {
    for (uint256 i = 0; i < configArgs.length; ++i) {
      if (configArgs[i].allowed) {
        if (s_allowedCallers.add(configArgs[i].caller)) {
          emit AllowedCallerAdded(configArgs[i].caller);
        }
      } else {
        if (s_allowedCallers.remove(configArgs[i].caller)) {
          emit AllowedCallerRemoved(configArgs[i].caller);
        }
      }
    }
  }

  /// @notice Checks if the caller is allowed to call the `receiveMessage` function.
  /// @param caller The address to check.
  /// @return allowed A boolean indicating if the caller is allowed.
  function isAllowedCaller(
    address caller
  ) external view returns (bool allowed) {
    return s_allowedCallers.contains(caller);
  }

  /// @notice Returns an array of all allowed callers.
  /// @return allowedCallers An array of allowed caller addresses.
  function getAllowedCallers() external view returns (address[] memory allowedCallers) {
    return s_allowedCallers.values();
  }
}
