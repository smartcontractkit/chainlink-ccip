// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {Client} from "../libraries/Client.sol";
import {MessageV1Codec} from "../libraries/MessageV1Codec.sol";

interface ICCVOnRampV1 {
  /// @notice Quotes the fee for a CCIP message to a destination chain.
  /// @dev This takes EVM2AnyMessage (instead of MessageV1) because
  /// the router client API that user contracts interact with (IRouterClient.getFee)
  /// exposes EVM2AnyMessage. The on-ramp can translate to MessageV1 internally
  /// where required (e.g., verifier hooks), but using EVM2AnyMessage here keeps the
  /// interface aligned with what clients construct and pass to the router.
  function getFee(
    address originalCaller,
    Client.EVM2AnyMessage memory message,
    bytes memory extraArgs
  ) external view returns (uint256);

  /// @notice Message sending, verifier hook.
  /// @param originalCaller The original caller of forwardToVerifier.
  /// @param message Decoded MessageV1 structure for the message being sent.
  /// @param messageId The message ID of the message being sent.
  /// @param feeToken Fee token used for this message.
  /// @param feeTokenAmount Amount of fee token provided.
  /// @param verifierArgs Opaque verifier-specific arguments from the sender.
  /// @return verifierData Verifier-specific return data blob.
  function forwardToVerifier(
    address originalCaller,
    MessageV1Codec.MessageV1 calldata message,
    bytes32 messageId,
    address feeToken,
    uint256 feeTokenAmount,
    bytes calldata verifierArgs
  ) external returns (bytes memory verifierData);
}
