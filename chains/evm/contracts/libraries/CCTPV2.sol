pragma solidity ^0.8.25;

import {CCTPMessageTransmitterProxy} from "../pools/USDC/CCTPMessageTransmitterProxy.sol";
import {IMessageTransmitter} from "../pools/USDC/interfaces/IMessageTransmitter.sol";
import {ITokenMessenger} from "../pools/USDC/interfaces/ITokenMessenger.sol";

import {USDCTokenPool} from "../pools/USDC/USDCTokenPool.sol";

library CCTPV2 {
  uint32 public constant FINALITY_THRESHOLD = 2000;
  uint32 public constant MESSAGE_VERSION = 1;

  // CCTP's max fee is based on the use of fast-burn. Since this pool does not utilize that feature, max fee should be 0.
  uint32 public constant MAX_FEE = 0;

  error InvalidMessageVersion(uint32 version);
  error InvalidSourceDomain(uint32 expected, uint32 got);
  error InvalidDestinationDomain(uint32 expected, uint32 got);
  error InvalidMinFinalityThreshold(uint32 expected, uint32 got);
  error InvalidExecutionFinalityThreshold(uint32 expected, uint32 got);
  error InvalidConfig();
  error InvalidTransmitterInProxy();
  error InvalidTokenMessengerVersion(uint32 version);

  event ConfigSet(address tokenMessenger);

  function _validateConfig(
    ITokenMessenger tokenMessenger,
    CCTPMessageTransmitterProxy cctpMessageTransmitterProxy
  ) internal {
    // NOTE: Even though it is officially referred to as Version 1, CCTP V1 contracts
    // use the version #0, and CCTP V2 contracts return a version number #1, so a contract
    // interacting with CCTP'V2 will look for it to return the version number of 1.

    if (address(tokenMessenger) == address(0)) revert InvalidConfig();

    // Get the Local Message Transmitter from the tokenMessenger
    IMessageTransmitter transmitter = IMessageTransmitter(tokenMessenger.localMessageTransmitter());

    // Check that the two contracts are using the same expected version of USDC/CCTP.
    uint32 transmitterVersion = transmitter.version();
    uint32 tokenMessengerVersion = tokenMessenger.messageBodyVersion();
    if (transmitterVersion != 1) revert InvalidMessageVersion(transmitterVersion);
    if (tokenMessengerVersion != 1) revert InvalidTokenMessengerVersion(tokenMessengerVersion);

    // Check that the transmitter called by the TransmitterProxy is the same as the one called by the TokenMessenger
    if (cctpMessageTransmitterProxy.i_cctpTransmitterV2() != transmitter) revert InvalidTransmitterInProxy();

    emit ConfigSet(address(tokenMessenger));
  }

  /// @notice Validates the USDC encoded message against the given parameters.
  /// @param usdcMessage The USDC encoded message
  /// @param sourcetokenDataPayload The expected source chain CCTP identifier as provided by the CCIP-Source-Pool.
  /// @param localDomainIdentifier The local domain identifier of the pool.
  /// @dev Only supports version SUPPORTED_USDC_VERSION of the CCTP V2 message format
  /// @dev Message format for USDC:
  ///     * Field                      Bytes      Type       Index
  ///     * version                    4          uint32     0
  ///     * sourceDomain               4          uint32     4
  ///     * destinationDomain          4          uint32     8
  ///     * nonce                      32         bytes32   12
  ///     * sender                     32         bytes32   44
  ///     * recipient                  32         bytes32   76
  ///     * destinationCaller          32         bytes32   108
  ///     * minFinalityThreshold       32         uint32    140
  ///     * finalityThresholdExecuted  32         uint32    144
  ///     * messageBody                dynamic    bytes     148
  function _validateMessage(
    bytes memory usdcMessage,
    USDCTokenPool.SourceTokenDataPayload memory sourcetokenDataPayload,
    uint32 localDomainIdentifier
  ) internal pure {
    uint32 version;
    // solhint-disable-next-line no-inline-assembly
    assembly {
      // We truncate using the datatype of the version variable, meaning
      // we will only be left with the first 4 bytes of the message.
      version := mload(add(usdcMessage, 4)) // 0 + 4 = 4
    }

    // This token pool only supports CCTP V2 with message format version being 1
    // We check the version prior to loading the rest of the message
    // to avoid unexpected reverts due to out-of-bounds reads.
    if (version != 1) revert InvalidMessageVersion(version);

    uint32 messageSourceDomain;
    uint32 destinationDomain;
    uint32 minFinalityThreshold;
    uint32 finalityThresholdExecuted;

    // solhint-disable-next-line no-inline-assembly
    assembly {
      messageSourceDomain := mload(add(usdcMessage, 8)) // 4 + 4 = 8
      destinationDomain := mload(add(usdcMessage, 12)) // 8 + 4 = 12
      minFinalityThreshold := mload(add(usdcMessage, 144)) // 140 + 4 = 144
      finalityThresholdExecuted := mload(add(usdcMessage, 148)) // 144 + 4 = 148
    }

    // Check that the source domain included in the CCTP Message matches the one forwarded by the source pool.
    if (messageSourceDomain != sourcetokenDataPayload.sourceDomain) {
      revert InvalidSourceDomain(sourcetokenDataPayload.sourceDomain, messageSourceDomain);
    }

    // Check that the destination domain in the CCTP message matches the immutable domain of this pool.
    if (destinationDomain != localDomainIdentifier) {
      revert InvalidDestinationDomain(localDomainIdentifier, destinationDomain);
    }

    // This pool only supports slow transfers on CCTP, so ensure that the message matches the same requirements.
    if (minFinalityThreshold != FINALITY_THRESHOLD) {
      revert InvalidMinFinalityThreshold(FINALITY_THRESHOLD, minFinalityThreshold);
    }

    if (finalityThresholdExecuted != FINALITY_THRESHOLD) {
      revert InvalidExecutionFinalityThreshold(FINALITY_THRESHOLD, finalityThresholdExecuted);
    }

    if (sourcetokenDataPayload.cctpVersion != USDCTokenPool.CCTPVersion.VERSION_2) {
      revert USDCTokenPool.InvalidCCTPVersion(sourcetokenDataPayload.sourceDomain, sourcetokenDataPayload.cctpVersion);
    }
  }
}
