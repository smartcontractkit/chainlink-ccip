// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {Client} from "../libraries/Client.sol";
import {MessageV1Codec} from "../libraries/MessageV1Codec.sol";

interface ICCVOnRampV1 {
  function getFee(
    uint64 destChainSelector,
    Client.EVM2AnyMessage memory message,
    bytes memory extraArgs
  ) external view returns (uint256);

  /// @notice Message sending, verifier hook
  /// @param message Decoded MessageV1 structure for the message being sent
  /// @param messageId keccak256(encodedMessage) of the MessageV1 encoding
  /// @param feeToken Fee token used for this message
  /// @param feeTokenAmount Amount of fee token provided
  /// @param verifierArgs Opaque verifier-specific arguments from the sender
  /// @return verifierData Verifier-specific return data blob
  function forwardToVerifier(
    MessageV1Codec.MessageV1 calldata message,
    bytes32 messageId,
    address feeToken,
    uint256 feeTokenAmount,
    bytes calldata verifierArgs
  ) external returns (bytes memory verifierData);
}
