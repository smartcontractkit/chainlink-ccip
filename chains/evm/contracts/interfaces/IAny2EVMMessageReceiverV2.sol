// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IAny2EVMMessageReceiver} from "./IAny2EVMMessageReceiver.sol";

interface IAny2EVMMessageReceiverV2 is IAny2EVMMessageReceiver {
  /// @notice Get the CCV configuration for a source chain and validate the finality requirement.
  /// @dev The requestedBlockDepth parameter provided MUST be checked, or anyone will be able to send messages with any
  /// level of finality to the receiver. In most cases, the receiver will want to require a certain level of finality.
  /// When a trusted sender is used (and verified by the receiver), this is less critical as the trusted sender will
  /// only send messages with a certain level of finality. The simplest way to implement this is to either allow FTF
  /// messages when sender-verification is used, or require finality for all messages. That means the config can be a
  /// simple boolean instead of n^2 config where for each source, some safe block depth must be chosen.
  ///
  /// A few methods to check the block depth requirement are:
  /// - Only allow trusted senders, never revert based on block depth
  /// - Revert if requestedBlockDepth is less than a single threshold (e.g. 10 blocks) for all chains
  /// - Revert if requestedBlockDepth is less than a threshold specific to the source chain (e.g. 10 blocks for chain A, 20 blocks for chain B)
  /// - Do not allow FTF messages at all, revert if requestedBlockDepth is not 0
  ///
  /// @param sourceChainSelector The source chain selector of the incoming message. This can be used to specify
  /// different CCV requirements for different source chains, and provides context for the requestedBlockDepth parameter.
  /// @param requestedBlockDepth The block depth of the incoming message. This is the number of blocks that have been
  /// mined on the source chain since the message was sent. A value of 0 indicates that the message will wait for
  /// finality and is the safest option.
  /// @dev Messages are executable when either the required block depth has been reached, or the chain has marked the
  /// block as finalized. Whichever one comes first will allow the message to be executed.
  /// @return requiredCCVs The list of required CCVs for messages from this source chain. All of these CCVs must pass
  /// verification for a message to be accepted.
  /// @return optionalCCVs The list of optional CCVs for messages from this source chain. These CCVs can be used to
  /// increase the security of messages from this source chain, but are not strictly required. If any optional CCVs are
  /// included, the optionalThreshold parameter must also be set to indicate how many of the optional CCVs must pass
  /// verification for a message to be accepted.
  /// @return optionalThreshold The number of optional CCVs that must pass verification for a message to be accepted.
  function getCCVs(
    uint64 sourceChainSelector,
    uint16 requestedBlockDepth
  ) external view returns (address[] memory requiredCCVs, address[] memory optionalCCVs, uint8 optionalThreshold);
}
