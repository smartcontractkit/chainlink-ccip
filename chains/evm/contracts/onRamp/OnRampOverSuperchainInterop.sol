// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Internal} from "../libraries/Internal.sol";
import {SuperchainInterop} from "../libraries/SuperchainInterop.sol";
import {OnRamp} from "./OnRamp.sol";

/// @notice This OnRamp supports OP Superchain Interop by emitting an interop-friendly event
/// and storing sent messages for re-emission when old source logs get pruned on remote chains.
contract OnRampOverSuperchainInterop is OnRamp {
  error MessageDoesNotExist(uint64 destChainSelector, uint64 sequenceNumber, bytes32 messageHash);
  error ExtraArgsTooShort(uint256 length);
  error InvalidSourceChainSelector(uint64 sourceChainSelector);

  // STATIC CONFIG
  string public constant override typeAndVersion = "OnRampOverSuperchainInterop 1.6.2-dev";

  /// @notice Stores previously sent interop message hashes to facilitate re-emission.
  /// @dev destChainSelector and sequenceNumber uniquely identify a message for a given onramp.
  mapping(uint64 destChainSelector => mapping(uint64 sequenceNumber => bytes32 messageHash)) internal
    s_sentInteropMessageHashes;

  constructor(
    StaticConfig memory staticConfig,
    DynamicConfig memory dynamicConfig,
    DestChainConfigArgs[] memory destChainConfigs
  ) OnRamp(staticConfig, dynamicConfig, destChainConfigs) {}

  /// @notice Override _postProcessMessage hook to emit interop-friendly event and store the message hash.
  /// @dev Superchain Interop, and most trustless cross-chain verifier solutions, rely on the fact that
  /// source storage or event data can be verified on the destination chain. This requires source and dest
  /// leveraging the same data, without offchain translation from EVM2AnyRampMessage to Any2EVMRampMessage.
  /// Here, we convert EVM2AnyRampMessage to Any2EVMRampMessage before emitting CCIPSuperchainMessageSent event.
  /// This makes CCIPSuperchainMessageSent event data readily usable by the OffRamp and for future re-emissions.
  /// @param message The message to process.
  /// @return The processed message, in this case the unaltered, original message being passed in.
  function _postProcessMessage(
    Internal.EVM2AnyRampMessage memory message
  ) internal virtual override returns (Internal.EVM2AnyRampMessage memory) {
    // Get the gas limit from the extraArgs
    uint256 gasLimit = extractGasLimit(message.extraArgs);

    Internal.Any2EVMTokenTransfer[] memory destTokenTransfers =
      new Internal.Any2EVMTokenTransfer[](message.tokenAmounts.length);
    for (uint256 i = 0; i < message.tokenAmounts.length; ++i) {
      Internal.EVM2AnyTokenTransfer memory tokenTransfer = message.tokenAmounts[i];

      // TokenPool owners can return arbitrary destTokenAddress values, so we validate
      // early to catch malicious/malformed data before abi.decode() potentially fails with unclear errors.
      Internal._validateEVMAddress(tokenTransfer.destTokenAddress);

      destTokenTransfers[i] = Internal.Any2EVMTokenTransfer({
        sourcePoolAddress: abi.encode(tokenTransfer.sourcePoolAddress),
        destTokenAddress: abi.decode(tokenTransfer.destTokenAddress, (address)),
        destGasAmount: abi.decode(tokenTransfer.destExecData, (uint32)),
        extraData: tokenTransfer.extraData,
        amount: tokenTransfer.amount
      });
    }

    // This conversion was done off chain in regular CCIP, with native interop it must happen onchain.
    // Conversion happens at OnRamp as opposed to OffRamp to reduce the distance of logical conversions
    // between validated logData and message content.
    // What is emitted at source is verified as the final message format used in execution.
    Internal.Any2EVMRampMessage memory interopMessage = Internal.Any2EVMRampMessage({
      header: Internal.RampMessageHeader({ // deep copy, because we will be setting the messageId later
        messageId: "",
        sourceChainSelector: message.header.sourceChainSelector,
        destChainSelector: message.header.destChainSelector,
        sequenceNumber: message.header.sequenceNumber,
        nonce: message.header.nonce
      }),
      sender: abi.encode(message.sender),
      data: message.data,
      receiver: abi.decode(message.receiver, (address)),
      gasLimit: gasLimit,
      tokenAmounts: destTokenTransfers
    });

    // Parent OnRamp has not set the messageId yet, calculate messageId the same way it's done in parent OnRamp
    // and set it in the interop message.
    interopMessage.header.messageId = generateMessageId(message);

    // Interop hash uniquely identifies the interop message.
    s_sentInteropMessageHashes[interopMessage.header.destChainSelector][interopMessage.header.sequenceNumber] =
      SuperchainInterop._hashInteropMessage(interopMessage, address(this));

    emit SuperchainInterop.CCIPSuperchainMessageSent(
      interopMessage.header.destChainSelector, interopMessage.header.sequenceNumber, interopMessage
    );

    return message;
  }

  /// @notice Extracts the gas limit from the extraArgs
  /// @dev This assume extraArgs is encoded with abi.encodeWithSelector,
  /// and the 1st element is always the gas limit.
  /// This assumption can be broken if EVM extraArgs are extended with dynamic types.
  /// It is likely in future ramp versions, which should sunset this superchain-specific ramp.
  /// @param extraArgs The extraArgs to extract the gas limit from
  /// @return gasLimit The gas limit
  function extractGasLimit(
    bytes memory extraArgs
  ) public pure returns (uint256 gasLimit) {
    if (extraArgs.length < 36) {
      revert ExtraArgsTooShort(extraArgs.length);
    }
    // We are limited to bytes memory as opposed to calldata, we cannot do data[4:] slice
    // memory layout of extraArgs array is: [ len 32 ][ selector 4 ] [gasLimit 32 ] [...]
    // mload always reads 32 bytes, we apply an offset of 36 bytes to fetch the gas limit
    assembly {
      gasLimit := mload(add(extraArgs, 0x24))
    }

    return gasLimit;
  }

  /// @notice Re-emits the CCIPSuperchainMessageSent event for a previously sent interop message.
  /// @dev This is necessary because Superchain does not persist cross-chain events on dest forever.
  /// A typical persistence window is 7 days. After that, the event needs to be re-emitted at source.
  /// @param interopMessage The previously-sent interop message to re-emit.
  function reemitInteropMessage(
    Internal.Any2EVMRampMessage calldata interopMessage
  ) external {
    // Validate that the message came from this chain.
    if (interopMessage.header.sourceChainSelector != i_chainSelector) {
      revert InvalidSourceChainSelector(interopMessage.header.sourceChainSelector);
    }

    bytes32 interopMessageHash = SuperchainInterop._hashInteropMessage(interopMessage, address(this));

    // Validates that the message had been sent before from this OnRamp.
    if (
      s_sentInteropMessageHashes[interopMessage.header.destChainSelector][interopMessage.header.sequenceNumber]
        != interopMessageHash
    ) {
      revert MessageDoesNotExist(
        interopMessage.header.destChainSelector, interopMessage.header.sequenceNumber, interopMessageHash
      );
    }

    // Re-emit the CCIPMessageSent event with the same data
    // Note we are not checking allowlist here. For the message to be re-emitted, it must have been sent
    // at a time when sender is allowed. To remain consistent with regular CCIP manual exec behavior,
    // we allow re-emissions regardless of current allowlist state at the source.
    emit SuperchainInterop.CCIPSuperchainMessageSent(
      interopMessage.header.destChainSelector, interopMessage.header.sequenceNumber, interopMessage
    );
  }
}
