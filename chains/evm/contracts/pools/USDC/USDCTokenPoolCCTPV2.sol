// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ITokenMessenger} from "./interfaces/ITokenMessenger.sol";

import {Pool} from "../../libraries/Pool.sol";

import {USDCSourcePoolDataCodec} from "../../libraries/USDCSourcePoolDataCodec.sol";
import {CCTPMessageTransmitterProxy} from "./CCTPMessageTransmitterProxy.sol";
import {USDCTokenPool} from "./USDCTokenPool.sol";

import {IERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/IERC20.sol";

/// @notice This pool mints and burns USDC tokens through the Cross Chain Transfer
/// Protocol (CCTP) V2.
/// @dev This pool inherits from the USDCTokenPool contract, but is not used for CCTP V1. It overrides
/// only the functions which are different for CCTP V2, due to a different message format and
/// deposit function. Since both pools use a message transmitter proxy, which will use the same
/// function selector, the releaseOrMint function does not need to be overridden.
contract USDCTokenPoolCCTPV2 is USDCTokenPool {
  error InvalidMinFinalityThreshold(uint32 expected, uint32 got);
  error InvalidExecutionFinalityThreshold(uint32 expected, uint32 got);
  error InvalidDepositHash(bytes32 expected, bytes32 got);
  error InvalidBurnToken(address expected, address got);
  error InvalidMinFee(uint256 maxAcceptableFee, uint256 actualFee);

  /// @dev CCTP's max fee is based on the use of fast-burn. Since this pool does not utilize that feature, max fee should be 0.
  uint32 public constant MAX_FEE = 0;

  /// @dev 2000 indicates that finality must be reached before attestation is possible in CCTP V2.
  uint32 public constant FINALITY_THRESHOLD = 2000;

  /// @dev The minimum length of a valid USDC message where all the required fields are present and capable of being parsed. While a real USDC message will be longer, only the first 280 bytes are needed to be parsed to extract the required fields.
  uint256 public constant MIN_USDC_MESSAGE_LENGTH = 280;

  function typeAndVersion() external pure virtual override returns (string memory) {
    return "USDCTokenPoolCCTPV2 1.6.x-dev";
  }

  /// @dev This contract is only used for CCTP V2, which is why the supportedUSDCVersion field of the parent
  /// constructor is set to 1. CCTP V1 used a version number of 0, so the version number is incremented by 1 for V2.
  constructor(
    ITokenMessenger tokenMessenger,
    CCTPMessageTransmitterProxy cctpMessageTransmitterProxy,
    IERC20 token,
    address[] memory allowlist,
    address rmnProxy,
    address router
  ) USDCTokenPool(tokenMessenger, cctpMessageTransmitterProxy, token, allowlist, rmnProxy, router, 1) {}

  /// @notice Burn tokens from the pool to initiate cross-chain transfer.
  /// @notice Outgoing messages (burn operations) are routed via `i_tokenMessenger.depositForBurn()`.
  /// The allowedCaller is preconfigured per destination domain and token pool version refer Domain struct.
  /// @dev Emits ITokenMessenger.DepositForBurn event.
  /// @dev Assumes caller has validated the destinationReceiver.
  function lockOrBurn(
    Pool.LockOrBurnInV1 calldata lockOrBurnIn
  ) public virtual override returns (Pool.LockOrBurnOutV1 memory) {
    _validateLockOrBurn(lockOrBurnIn);

    Domain memory domain = s_chainToDomain[lockOrBurnIn.remoteChainSelector];
    if (!domain.enabled) revert UnknownDomain(lockOrBurnIn.remoteChainSelector);

    if (lockOrBurnIn.receiver.length != 32) {
      revert InvalidReceiver(lockOrBurnIn.receiver);
    }

    // Some CCTP-V2 chains support a configurable fee switch, but not all. It is therefore
    // necessary to check via a try-catch block. If the call reverts, then the fee switch is not supported and the
    // standard transfer fee will be zero, and no further action is required.
    try i_tokenMessenger.getMinFeeAmount(lockOrBurnIn.amount) returns (uint256 minFee) {
      // This token pool only supports zero-fee standard transfers. If the minFee is non-zero
      // then the function should revert as the message may not be able to be successfully
      // delivered on destination due to unexpected minting fees.
      if (minFee > MAX_FEE) {
        revert InvalidMinFee(MAX_FEE, minFee);
      }
    } catch {}

    bytes32 decodedReceiver;
    // For EVM chains, the mintRecipient is not used, but is needed for Solana, where the mintRecipient will
    // be a PDA owned by the pool, and will forward the tokens to its final destination after minting.
    if (domain.mintRecipient != bytes32(0)) {
      decodedReceiver = domain.mintRecipient;
    } else {
      decodedReceiver = abi.decode(lockOrBurnIn.receiver, (bytes32));
    }

    // Deposit the tokens for burn into CCTP.
    i_tokenMessenger.depositForBurn(
      lockOrBurnIn.amount,
      domain.domainIdentifier,
      decodedReceiver,
      address(i_token),
      domain.allowedCaller,
      MAX_FEE,
      FINALITY_THRESHOLD
    );

    // In CCTP v2, the nonce is not returned to this contract upon sending the message request, and will instead be
    // acquired off-chain and included in the destination-message's offchainTokenData. However, to ensure that the
    // correct attestation can be matched to a specific CCIP message, an identifier is needed. This hash is used to
    // identify the CCIP message which corresponds to an attestation.
    bytes32 depositHash = USDCSourcePoolDataCodec._calculateDepositHash(
      i_localDomainIdentifier, // sourceDomain
      lockOrBurnIn.amount, // amount
      domain.domainIdentifier, // destinationDomain
      decodedReceiver, // mintRecipient
      bytes32(uint256(uint160(address(i_token)))), // burnToken
      domain.allowedCaller, // destinationCaller
      MAX_FEE, // maxFee
      FINALITY_THRESHOLD // minFinalityThreshold
    );

    // Encode the source pool data with its version number. The version number is hard-coded to 1 to maintain
    // parity with the CCTP V2 version number.
    bytes memory sourcePoolData = USDCSourcePoolDataCodec._encodeSourceTokenDataPayloadV2(
      USDCSourcePoolDataCodec.SourceTokenDataPayloadV2({sourceDomain: i_localDomainIdentifier, depositHash: depositHash})
    );

    emit LockedOrBurned({
      remoteChainSelector: lockOrBurnIn.remoteChainSelector,
      token: address(i_token),
      sender: msg.sender,
      amount: lockOrBurnIn.amount
    });

    return Pool.LockOrBurnOutV1({
      destTokenAddress: getRemoteToken(lockOrBurnIn.remoteChainSelector),
      destPoolData: sourcePoolData
    });
  }

  /// @inheritdoc USDCTokenPool
  function releaseOrMint(
    Pool.ReleaseOrMintInV1 calldata releaseOrMintIn
  ) public virtual override returns (Pool.ReleaseOrMintOutV1 memory) {
    _validateReleaseOrMint(releaseOrMintIn, releaseOrMintIn.sourceDenominatedAmount);

    MessageAndAttestation memory msgAndAttestation =
      abi.decode(releaseOrMintIn.offchainTokenData, (MessageAndAttestation));

    // Decode the source pool data from its raw bytes into a SourceTokenDataPayloadV1 struct that can be
    // more easily validated below.
    USDCSourcePoolDataCodec.SourceTokenDataPayloadV2 memory sourceTokenDataPayload =
      USDCSourcePoolDataCodec._decodeSourceTokenDataPayloadV2(releaseOrMintIn.sourcePoolData);

    _validateMessage(msgAndAttestation.message, sourceTokenDataPayload);

    // Proxy the message to the message transmitter, which will mint the tokens through CCTP's contracts.
    if (!i_messageTransmitterProxy.receiveMessage(msgAndAttestation.message, msgAndAttestation.attestation)) {
      revert UnlockingUSDCFailed();
    }

    emit ReleasedOrMinted({
      remoteChainSelector: releaseOrMintIn.remoteChainSelector,
      token: address(i_token),
      sender: msg.sender,
      recipient: releaseOrMintIn.receiver,
      amount: releaseOrMintIn.sourceDenominatedAmount
    });

    return Pool.ReleaseOrMintOutV1({destinationAmount: releaseOrMintIn.sourceDenominatedAmount});
  }

  /// @notice Validates the USDC encoded message against the given parameters.
  /// @param usdcMessage The USDC encoded message
  /// @param sourceTokenData The expected source chain CCTP identifier as provided by the CCIP-Source-Pool.
  /// @dev Only supports version SUPPORTED_USDC_VERSION of the CCTP V2 message format
  /// which is documented at https://developers.circle.com/cctp/technical-guide#cctp-v2-message-format.
  /// @dev The circle documentation clarifies that this top level message header format is standard for all messages passing
  /// through CCTP.
  /// @dev Message format for USDC:
  ///     * Field                      Bytes      Type       Index
  ///     * version                    4          uint32     0
  ///     * sourceDomain               4          uint32     4
  ///     * destinationDomain          4          uint32     8
  ///     * nonce                      32         bytes32   12
  ///     * sender                     32         bytes32   44
  ///     * recipient                  32         bytes32   76
  ///     * destinationCaller          32         bytes32   108
  ///     * minFinalityThreshold       4         uint32     140
  ///     * finalityThresholdExecuted  4         uint32     144
  ///     * messageBody                dynamic    bytes     148

  /// @dev Message Body for USDC.
  ///     * Field                 Bytes      Type       Index
  ///     * version               4          uint32     0
  ///     * burnToken             32         bytes32    4
  ///     * mintRecipient         32         bytes32    36
  ///     * amount                32         uint256    68
  ///     * messageSender         32         bytes32    100
  ///     * maxFee                32         uint256    132
  ///     * feeExecuted           32         uint256    164
  ///     * expirationBlock       32         uint256    196
  ///     * hookData              dynamic    bytes      228
  /// @dev The CCTP documentation does not explicitly state that this message body format will be used for all messages passing
  /// through CCTP, including for Non-EVM chains. As a result if in the future this format is not used, parsing logic may
  /// need to be modified accordingly.
  function _validateMessage(
    bytes memory usdcMessage,
    USDCSourcePoolDataCodec.SourceTokenDataPayloadV2 memory sourceTokenData
  ) internal view {
    if (usdcMessage.length < MIN_USDC_MESSAGE_LENGTH) revert InvalidMessageLength(usdcMessage.length);

    uint32 version;
    // solhint-disable-next-line no-inline-assembly
    assembly {
      // We truncate using the datatype of the version variable, meaning
      // we will only be left with the first 4 bytes of the message when we cast it to uint32. We want the lower 4 bytes
      // to be the version when casted to a uint32 , so we only add 4. If you added 32, attempting to skip the first word
      // containing the length, then version would be in the upper-4 bytes of the corresponding slot, which
      // would not be as easily parsed into a uint32.
      version := mload(add(usdcMessage, 4)) // 0 + 4 = 4
    }

    // Note: Even though the CCTP Version is V2, it's on-chain version number is 1, since V1 used a version number of 0.
    // This is different from the sourceDomain field of sourceTokenData, which is used by off-chain code and the token
    // pools, rather than being a formal part of the CCTP message format.
    if (version != i_supportedUSDCVersion) revert InvalidMessageVersion(i_supportedUSDCVersion, version);

    // Fields from the message header
    uint32 messageSourceDomain;
    uint32 destinationDomain;
    uint32 minFinalityThreshold;
    uint32 finalityThresholdExecuted;
    bytes32 destinationCaller;

    // Fields from the message body
    uint256 amount;
    bytes32 burnToken;
    bytes32 mintRecipient;

    // solhint-disable-next-line no-inline-assembly
    assembly {
      messageSourceDomain := mload(add(usdcMessage, 8)) // 4 + 4 = 8
      destinationDomain := mload(add(usdcMessage, 12)) // 8 + 4 = 12
      destinationCaller := mload(add(usdcMessage, 140)) // 32 + 108 = 140
      minFinalityThreshold := mload(add(usdcMessage, 144)) // 140 + 4 = 144
      finalityThresholdExecuted := mload(add(usdcMessage, 148)) // 144 + 4 = 148

      // The message body starts at index 148 and because it is a dynamic byte array, contains a 32-byte
      // length field prefixing the data.
      burnToken := mload(add(usdcMessage, 184)) // 148 + 32 + 4 = 184
      mintRecipient := mload(add(usdcMessage, 216)) // 148 + 32 + 36 = 216
      amount := mload(add(usdcMessage, 248)) // 148 + 32 + 68 = 248
    }

    // Check that the source domain included in the CCTP Message matches the one forwarded by the source pool.
    if (messageSourceDomain != sourceTokenData.sourceDomain) {
      revert InvalidSourceDomain(sourceTokenData.sourceDomain, messageSourceDomain);
    }

    // Check that the destination domain in the CCTP message matches the immutable domain of this pool.
    if (destinationDomain != i_localDomainIdentifier) {
      revert InvalidDestinationDomain(i_localDomainIdentifier, destinationDomain);
    }

    // This pool only supports slow transfers on CCTP, so ensure that the message matches the same requirements.
    if (minFinalityThreshold != FINALITY_THRESHOLD) {
      revert InvalidMinFinalityThreshold(FINALITY_THRESHOLD, minFinalityThreshold);
    }

    if (finalityThresholdExecuted != FINALITY_THRESHOLD) {
      revert InvalidExecutionFinalityThreshold(FINALITY_THRESHOLD, finalityThresholdExecuted);
    }

    // Calculate the deposit hash for the message using the locally parsed fields.
    bytes32 derivedDepositHash = USDCSourcePoolDataCodec._calculateDepositHash(
      messageSourceDomain,
      amount,
      destinationDomain,
      mintRecipient,
      burnToken,
      destinationCaller,
      MAX_FEE,
      minFinalityThreshold
    );

    // Check that the locally calculated deposit hash matches the one provided by the source pool. This is critical to
    // ensuring that the correct attestation is used for the given message. Without it, a user may be able to use an
    // attestation for a different message to mint the tokens, thus preventing the legitimate user from minting the
    // tokens on destination.
    if (derivedDepositHash != sourceTokenData.depositHash) {
      revert InvalidDepositHash(sourceTokenData.depositHash, derivedDepositHash);
    }
  }
}
