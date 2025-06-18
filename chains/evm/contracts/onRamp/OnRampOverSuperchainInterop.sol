// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Internal} from "../libraries/Internal.sol";
import {SuperchainInterop} from "../libraries/SuperchainInterop.sol";
import {OnRamp} from "./OnRamp.sol";

/// @notice OnRamp that supports superchain interoperability by storing sent messages for re-emission
contract OnRampOverSuperchainInterop is OnRamp {
  error MessageDoesNotExist(uint64 destChainSelector, uint64 sequenceNumber, bytes32 messageHash);
  error ExtraArgsTooShort(uint256 length);
  error InvalidSourceChainSelector(uint64 sourceChainSelector);

  /// @notice Using a function because constant state variables cannot be overridden by child contracts.
  function typeAndVersion() public pure virtual override returns (string memory) {
    return "OnRampOverSuperchainInterop 1.6.1-dev";
  }

  /// @dev stores sent interop message hashes to facilitate re-emission.
  mapping(uint64 destChainSelector => mapping(uint64 sequenceNumber => bytes32 messageHash)) internal
    s_sentInteropMessageHashes;

  constructor(
    StaticConfig memory staticConfig,
    DynamicConfig memory dynamicConfig,
    DestChainConfigArgs[] memory destChainConfigs
  ) OnRamp(staticConfig, dynamicConfig, destChainConfigs) {}

  /// @notice Override _postProcessMessage hook to calculate messageId and store the message
  /// @param message The message to process.
  /// @return processedMessage The processed message.
  function _postProcessMessage(
    Internal.EVM2AnyRampMessage memory message
  ) internal virtual override returns (Internal.EVM2AnyRampMessage memory) {
    Internal.EVM2AnyRampMessage memory processedMessage = super._postProcessMessage(message);

    // Get the gas limit from the extraArgs
    uint256 gasLimit = extractGasLimit(processedMessage.extraArgs);

    Internal.Any2EVMTokenTransfer[] memory destTokenTranfers =
      new Internal.Any2EVMTokenTransfer[](processedMessage.tokenAmounts.length);
    for (uint256 i = 0; i < processedMessage.tokenAmounts.length; ++i) {
      Internal.EVM2AnyTokenTransfer memory tokenTransfer = processedMessage.tokenAmounts[i];
      destTokenTranfers[i] = Internal.Any2EVMTokenTransfer({
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
        sourceChainSelector: processedMessage.header.sourceChainSelector,
        destChainSelector: processedMessage.header.destChainSelector,
        sequenceNumber: processedMessage.header.sequenceNumber,
        nonce: processedMessage.header.nonce
      }),
      sender: abi.encode(processedMessage.sender),
      data: processedMessage.data,
      receiver: abi.decode(processedMessage.receiver, (address)),
      gasLimit: gasLimit,
      tokenAmounts: destTokenTranfers
    });

    // Calculate the messageId the same way it's done in forwardFromRouter
    bytes32 messageId = generateMessageId(processedMessage);

    // Parent OnRamp has not set the messageId yet, need to set it here.
    interopMessage.header.messageId = messageId;

    // Interop hash uniquely identifies the interop message.
    bytes32 interopMessageHash = hashInteropMessage(interopMessage);
    s_sentInteropMessageHashes[interopMessage.header.destChainSelector][interopMessage.header.sequenceNumber] =
      interopMessageHash;

    emit SuperchainInterop.CCIPSuperchainMessageSent(
      interopMessage.header.destChainSelector, interopMessage.header.sequenceNumber, interopMessage
    );

    return processedMessage;
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

  /// @notice Hashes the interop message.
  /// @dev This uniquely identifies the interop message using the same logic as the offRamp.
  /// @param message The interop message to hash.
  /// @return messageHash The hash of the interop message.
  function hashInteropMessage(
    Internal.Any2EVMRampMessage memory message
  ) public view returns (bytes32) {
    bytes32 offRampMetaDataHash = keccak256(
      abi.encode(
        Internal.ANY_2_EVM_MESSAGE_HASH,
        i_chainSelector,
        message.header.destChainSelector,
        keccak256(abi.encode(address(this)))
      )
    );

    return Internal._hash(message, offRampMetaDataHash);
  }

  /// @notice Re-emits the CCIPSuperchainMessageSent event for a previously sent interop message.
  /// @dev This is necessary because Superchain Interop dest chain does not persist events forever.
  /// A typical persistance time is 7 days.
  /// @param interopMessage The previously-sent interop message to re-emit.
  function reemitInteropMessage(
    Internal.Any2EVMRampMessage calldata interopMessage
  ) external {
    // Validate that the message is meant for this chain.
    if (interopMessage.header.sourceChainSelector != i_chainSelector) {
      revert InvalidSourceChainSelector(interopMessage.header.sourceChainSelector);
    }

    bytes32 interopMessageHash = hashInteropMessage(interopMessage);

    uint64 destChainSelector = interopMessage.header.destChainSelector;
    uint64 sequenceNumber = interopMessage.header.sequenceNumber;

    // Validates that the message had been sent before.
    if (s_sentInteropMessageHashes[destChainSelector][sequenceNumber] != interopMessageHash) {
      revert MessageDoesNotExist(destChainSelector, sequenceNumber, interopMessageHash);
    }

    // Re-emit the CCIPMessageSent event with the same data
    emit SuperchainInterop.CCIPSuperchainMessageSent(
      interopMessage.header.destChainSelector, interopMessage.header.sequenceNumber, interopMessage
    );
  }
}
