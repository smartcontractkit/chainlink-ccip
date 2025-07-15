// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ITokenMessenger} from "./interfaces/ITokenMessenger.sol";

import {Pool} from "../../libraries/Pool.sol";
import {CCTPMessageTransmitterProxy} from "./CCTPMessageTransmitterProxy.sol";
import {USDCTokenPool} from "./USDCTokenPool.sol";

import {IERC20} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/utils/SafeERC20.sol";

/// @notice This pool mints and burns USDC tokens through the Cross Chain Transfer
/// Protocol (CCTP).
contract USDCTokenPoolCCTPV2 is USDCTokenPool {
  using SafeERC20 for IERC20;

  error InvalidMinFinalityThreshold(uint32 expected, uint32 got);
  error InvalidExecutionFinalityThreshold(uint32 expected, uint32 got);

  // CCTP's max fee is based on the use of fast-burn. Since this pool does not utilize that feature, max fee should be 0.
  uint32 public constant MAX_FEE = 0;

  // TODO: Add Comment
  uint32 public constant FINALITY_THRESHOLD = 2000;

  ITokenMessenger public immutable i_legacyTokenMessenger;

  // TODO: Fix Comments
  // Note: This constructor is only used for CCTP V2, which is why the supportedUSDCVersion is set to 1.
  constructor(
    ITokenMessenger legacyTokenMessenger,
    ITokenMessenger tokenMessenger,
    CCTPMessageTransmitterProxy cctpMessageTransmitterProxy,
    IERC20 token,
    address[] memory allowlist,
    address rmnProxy,
    address router,
    address previousPool
  ) USDCTokenPool(tokenMessenger, cctpMessageTransmitterProxy, token, allowlist, rmnProxy, router, previousPool, 1) {
    if (previousPool != address(0)) {
      // If the previous pool exists, we need to acquire the previous pool's message transmitter proxy so that a 
      // messages' destinationCaller can be checked against it.
      try USDCTokenPool(previousPool).i_messageTransmitterProxy() returns (CCTPMessageTransmitterProxy proxy) {
        i_previousMessageTransmitterProxy = address(proxy);
      } catch {
        revert InvalidPreviousPool();
      }
    } else {
      i_previousMessageTransmitterProxy = address(0);
    }

    // Increase allowance for the legacy token messenger to allow for the migration of tokens.
    i_token.safeIncreaseAllowance(address(legacyTokenMessenger), type(uint256).max);
    i_legacyTokenMessenger = legacyTokenMessenger;

    emit ConfigSet(address(tokenMessenger));
  }

  /// @notice Burn tokens from the pool to initiate cross-chain transfer.
  /// @notice Outgoing messages (burn operations) are routed via `i_tokenMessenger.depositForBurnWithCaller`.
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

    bytes32 decodedReceiver;
    // For EVM chains, the mintRecipient is not used, but is needed for Solana, where the mintRecipient will
    // be a PDA owned by the pool, and will forward the tokens to its final destination after minting.
    if (domain.mintRecipient != bytes32(0)) {
      decodedReceiver = domain.mintRecipient;
    } else {
      decodedReceiver = abi.decode(lockOrBurnIn.receiver, (bytes32));
    }

    uint64 nonce;
    CCTPVersion cctpVersion;
    // Since this pool is the msg sender of the CCTP transaction, only this contract
    // is able to call replaceDepositForBurn. Since this contract does not implement
    // replaceDepositForBurn, the tokens cannot be maliciously re-routed to another address.

    // If the CCTP version is CCTP_V1, we use the legacy token messenger to deposit for burn.
    if (domain.cctpVersion == CCTPVersion.CCTP_V1) {
      cctpVersion = CCTPVersion.CCTP_V1;

      nonce = i_legacyTokenMessenger.depositForBurnWithCaller(
        lockOrBurnIn.amount,
        domain.domainIdentifier,
        decodedReceiver,
        address(i_token),
        domain.allowedCaller
      );

    // If the CCTP version is CCTP_V2, we use the new token messenger to deposit for burn.
    } else if (domain.cctpVersion == CCTPVersion.CCTP_V2) {
      cctpVersion = CCTPVersion.CCTP_V2;

      i_tokenMessenger.depositForBurn(
        lockOrBurnIn.amount,
        domain.domainIdentifier,
        decodedReceiver,
        address(i_token),
        domain.allowedCaller,
        MAX_FEE,
        FINALITY_THRESHOLD
      );
    }

    emit LockedOrBurned({
      remoteChainSelector: lockOrBurnIn.remoteChainSelector,
      token: address(i_token),
      sender: msg.sender,
      amount: lockOrBurnIn.amount
    });

    SourceTokenDataPayload memory sourceTokenDataPayload = SourceTokenDataPayload({
      nonce: nonce,
      sourceDomain: i_localDomainIdentifier,
      cctpVersion: cctpVersion,
      amount: lockOrBurnIn.amount,
      destinationDomain: i_localDomainIdentifier,
      mintRecipient: decodedReceiver,
      burnToken: address(i_token),
      destinationCaller: domain.allowedCaller,
      maxFee: MAX_FEE,
      minFinalityThreshold: FINALITY_THRESHOLD
    });

    // As of CCTP v2, the nonce is not returned to this contract upon sending the message, and will instead be
    // acquired off-chain and included in the destination-message's offchainTokenData, so we set it to 0.
    return Pool.LockOrBurnOutV1({
      destTokenAddress: getRemoteToken(lockOrBurnIn.remoteChainSelector),
      destPoolData: abi.encode(sourceTokenDataPayload)
    });
  }

  /// @notice Mint tokens from the pool to the recipient
  /// * sourceTokenData is part of the verified message and passed directly from
  /// the offRamp so it is guaranteed to be what the lockOrBurn pool released on the
  /// source chain. It contains (nonce, sourceDomain) which is guaranteed by CCTP
  /// to be unique.
  /// * offchainTokenData is untrusted (can be supplied by manual execution), but we assert
  /// that (nonce, sourceDomain) is equal to the message's (nonce, sourceDomain) and
  /// receiveMessage will assert that Attestation contains a valid attestation signature
  /// for that message, including its (nonce, sourceDomain). This way, the only
  /// non-reverting offchainTokenData that can be supplied is a valid attestation for the
  /// specific message that was sent on source.
  function releaseOrMint(
    Pool.ReleaseOrMintInV1 calldata releaseOrMintIn
  ) public virtual override returns (Pool.ReleaseOrMintOutV1 memory) {
    _validateReleaseOrMint(releaseOrMintIn, releaseOrMintIn.sourceDenominatedAmount);

    MessageAndAttestation memory msgAndAttestation =
      abi.decode(releaseOrMintIn.offchainTokenData, (MessageAndAttestation));

    // If the destinationCaller is the previous pool, indicating an inflight message during the migration, we need to
    // route the message to the previous pool to satisfy the allowedCaller.
    bytes32 destinationCallerBytes32;
    bytes memory messageBytes = msgAndAttestation.message;
    assembly {
      // destinationCaller is a 32-byte word starting at position 84 in messageBytes body, so add 32 to skip the 1st word
      // representing bytes length
      destinationCallerBytes32 := mload(add(messageBytes, 140)) // 108 + 32 = 140
    }
    address destinationCaller = address(uint160(uint256(destinationCallerBytes32)));

    // If the destinationCaller is the previous pool's message transmitter proxy, we can use this
    // as an indication that CCTP V1 was used to send the message, and route it to the previous pool for minting.
    // In previous versions, the sourcePoolData only contained two fields, a uint64 and uint32. For structs stored only
    // in memory, the compiler assigns each field to its only 32-byte slot, instead of tightly packing line in storage.
    // This means that the sourcePoolData will be 64 bytes long. This indicates an inflight message during the
    // migration, and needs to be routed to the previous pool, otherwise the parsing will revert.
    if (
      (i_previousPool != address(0) && destinationCaller == i_previousMessageTransmitterProxy)
        || releaseOrMintIn.sourcePoolData.length == 64
    ) {
      // If the destinationCaller is the previous pool's message transmitter proxy, we can use this
      // as an indication that CCTP V1 was used to send the message, and route it to the previous pool for minting.
      return USDCTokenPool(i_previousPool).releaseOrMint(releaseOrMintIn);
    }

    // This decoding is done after the check for the previous pool to avoid issues with decoding the previous pool's
    // sourcePoolData into a struct with a different number of fields.
    SourceTokenDataPayload memory sourceTokenDataPayload =
      abi.decode(releaseOrMintIn.sourcePoolData, (SourceTokenDataPayload));

    // We call this after the destinationCaller check to ensure that the message is valid for CCTP V2. If it was called
    // before, then a V1 message which should be forwarded to the previous pool will be rejected.
    _validateMessage(msgAndAttestation.message, sourceTokenDataPayload);

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
    SourceTokenDataPayload memory sourceTokenData
  ) internal view override {
    // 116 is the minimum length of a valid USDC message. Since destinationCaller needs to be checked for the previous
    // pool, this ensures that it can be parsed correctly and that the message is not too short. Since messageBody is
    // dynamic and not always used, it is not checked.
    if (usdcMessage.length < 148) revert InvalidMessageLength(usdcMessage.length);

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
    // This token pool only supports version 0 of the CCTP message format
    // We check the version prior to loading the rest of the message
    // to avoid unexpected reverts due to out-of-bounds reads.
    if (version != i_supportedUSDCVersion) revert InvalidMessageVersion(version);

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
    if (messageSourceDomain != sourceTokenData.sourceDomain) {
      revert InvalidSourceDomain(sourceTokenData.sourceDomain, messageSourceDomain);
    }

    // Check that the destination domain in the CCTP message matches the immutable domain of this pool.
    if (destinationDomain != i_localDomainIdentifier) {
      revert InvalidDestinationDomain(i_localDomainIdentifier, destinationDomain);
    }

    // Check that the CCTP version in the CCTP message matches the expected version.
    if (sourceTokenData.cctpVersion != CCTPVersion.CCTP_V2) {
      revert USDCTokenPool.InvalidCCTPVersion(CCTPVersion.CCTP_V2, sourceTokenData.cctpVersion);
    }

    // This pool only supports slow transfers on CCTP, so ensure that the message matches the same requirements.
    if (minFinalityThreshold != FINALITY_THRESHOLD) {
      revert InvalidMinFinalityThreshold(FINALITY_THRESHOLD, minFinalityThreshold);
    }

    if (finalityThresholdExecuted != FINALITY_THRESHOLD) {
      revert InvalidExecutionFinalityThreshold(FINALITY_THRESHOLD, finalityThresholdExecuted);
    }
  }
}
