// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ITokenMessenger} from "../interfaces/ITokenMessenger.sol";

import {Pool} from "../../../libraries/Pool.sol";
import {CCTPMessageTransmitterProxy} from "../CCTPMessageTransmitterProxy.sol";
import {USDCTokenPool} from "../USDCTokenPool.sol";

import {IERC20} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/utils/SafeERC20.sol";

import {EnumerableSet} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v5.0.2/contracts/utils/structs/EnumerableSet.sol";

/// @notice This pool mints and burns USDC tokens through the Cross Chain Transfer
/// Protocol (CCTP) V2, which uses a different contract and message format as V1.
/// @dev The code for the message transmitter proxy does NOT need to be modified since both CCTP V1 and V2 utilize the same
/// interface for its MessageTransmitter, but the CCTP-controlled address that the proxy points to will be different
/// than its V1 predecessor
contract USDCTokenPoolCCTPV2 is USDCTokenPool {
  using SafeERC20 for IERC20;
  using EnumerableSet for EnumerableSet.UintSet;

  error InvalidMinFinalityThreshold(uint32 expected, uint32 actual);
  error InvalidExecutionFinalityThreshold(uint32 expected, uint32 actual);

  // CCTP's max fee is based on the use of fast-burn. Since this pool does not utilize that feature, max fee should be 0.
  uint32 public constant MAX_FEE = 0;

  // CCTP V2 uses 2000 to indicate that attestations should not occur until finality is achieved on the source chain.
  uint32 public constant FINALITY_THRESHOLD = 2000;

  constructor(
    ITokenMessenger tokenMessenger,
    CCTPMessageTransmitterProxy cctpMessageTransmitterProxy,
    IERC20 token,
    address[] memory allowlist,
    address rmnProxy,
    address router
  ) USDCTokenPool(tokenMessenger, cctpMessageTransmitterProxy, token, allowlist, rmnProxy, router, 1) {}

  /// @notice Burn tokens from the pool to initiate cross-chain transfer.
  /// @notice Outgoing messages (burn operations) are routed via `i_tokenMessenger.depositForBurnWithCaller`.
  /// The allowedCaller is preconfigured per destination domain and token pool version refer Domain struct.
  /// @dev Emits ITokenMessenger.DepositForBurn event.
  /// @dev Assumes caller has validated the destinationReceiver.
  function lockOrBurn(
    Pool.LockOrBurnInV1 calldata lockOrBurnIn
  ) public virtual override returns (Pool.LockOrBurnOutV1 memory) {
    _validateLockOrBurn(lockOrBurnIn);

    USDCTokenPool.Domain memory domain = s_chainToDomain[lockOrBurnIn.remoteChainSelector];

    if (!domain.enabled) revert UnknownDomain(lockOrBurnIn.remoteChainSelector);

    if (lockOrBurnIn.receiver.length != 32) {
      revert InvalidReceiver(lockOrBurnIn.receiver);
    }
    bytes32 decodedReceiver = abi.decode(lockOrBurnIn.receiver, (bytes32));

    // Since this pool is the msg sender of the CCTP transaction, only this contract
    // is able to call replaceDepositForBurn. Since this contract does not implement
    // replaceDepositForBurn, the tokens cannot be maliciously re-routed to another address.

    // Since the CCTP message will use slow-burn, the maxFee is 0, and the finality threshold is standard (2000).
    // Using fast-burn would require a maxFee and a finality threshold of 1000, which may be added in the future.
   
    // In CCTP V2, nonces are deterministic and not sequential. As a result the nonce is not returned to this contract
    // upon sending the message, and will therefore not be included in the destPoolData. It will instead be
    // acquired off-chain and included in the destination-message's offchainTokenData.
    i_tokenMessenger.depositForBurn(
      lockOrBurnIn.amount, // amount
      domain.domainIdentifier, // destinationDomain
      decodedReceiver, // mintRecipient
      address(i_token), // burnToken
      domain.allowedCaller, // destinationCaller
      MAX_FEE, // maxFee
      FINALITY_THRESHOLD // minFinalityThreshold
    );

    emit LockedOrBurned({
      remoteChainSelector: lockOrBurnIn.remoteChainSelector,
      token: address(i_token),
      sender: msg.sender,
      amount: lockOrBurnIn.amount
    });

    // Since CCTP V2 does not return a nonce during the deposit call, we can just use zero to satisfy the struct field.
    return Pool.LockOrBurnOutV1({
      destTokenAddress: getRemoteToken(lockOrBurnIn.remoteChainSelector),
      destPoolData: abi.encode(
        SourceTokenDataPayload({nonce: 0, sourceDomain: i_localDomainIdentifier, cctpVersion: CCTPVersion.VERSION_2})
      )
    });
  }

  /// @notice Mint tokens from the pool to the recipient
  /// * sourceTokenData is part of the verified message and passed directly from
  /// the offRamp so it is guaranteed to be what the lockOrBurn pool released on the
  /// source chain. It contains (nonce, sourceDomain) which is guaranteed by CCTP
  /// to be unique.
  /// * offchainTokenData is untrusted (can be supplied by manual execution), but we assert
  /// that (sourceDomain) is equal to the message's (sourceDomain) and
  /// receiveMessage will assert that Attestation contains a valid attestation signature
  /// for that message, including its (nonce, sourceDomain). This way, the only
  /// non-reverting offchainTokenData that can be supplied is a valid attestation for the
  /// specific message that was sent on source.
  function releaseOrMint(
    Pool.ReleaseOrMintInV1 calldata releaseOrMintIn
  ) public virtual override returns (Pool.ReleaseOrMintOutV1 memory) {
    _validateReleaseOrMint(releaseOrMintIn);

    SourceTokenDataPayload memory sourceTokenData = abi.decode(releaseOrMintIn.sourcePoolData, (SourceTokenDataPayload));

    MessageAndAttestation memory msgAndAttestation =
      abi.decode(releaseOrMintIn.offchainTokenData, (MessageAndAttestation));

    _validateMessage(msgAndAttestation.message, sourceTokenData);

    if (
      !i_messageTransmitterProxy.receiveMessage(
        msgAndAttestation.message, msgAndAttestation.attestation, sourceTokenData.cctpVersion
      )
    ) {
      revert UnlockingUSDCFailed();
    }

    // emit Minted(msg.sender, releaseOrMintIn.receiver, releaseOrMintIn.amount);
    emit ReleasedOrMinted({
      remoteChainSelector: releaseOrMintIn.remoteChainSelector,
      token: address(i_token),
      sender: msg.sender,
      recipient: releaseOrMintIn.receiver,
      amount: releaseOrMintIn.amount
    });

    return Pool.ReleaseOrMintOutV1({destinationAmount: releaseOrMintIn.amount});
  }

  /// @notice Validates the USDC encoded message against the given parameters.
  /// @param usdcMessage The USDC encoded message
  /// @param sourcePoolTokenData The expected source chain CCTP identifier as provided by the CCIP-Source-Pool.
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
    SourceTokenDataPayload memory sourcePoolTokenData
  ) internal view override {
    uint32 version;
    // solhint-disable-next-line no-inline-assembly
    assembly {
      // We truncate using the datatype of the version variable, meaning
      // we will only be left with the first 4 bytes of the message.
      version := mload(add(usdcMessage, 4)) // 0 + 4 = 4
    }

    // This token pool only supports version 1 of the CCTP message format
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
    if (messageSourceDomain != sourcePoolTokenData.sourceDomain) {
      revert InvalidSourceDomain(sourcePoolTokenData.sourceDomain, messageSourceDomain);
    }

    // Check that the destination domain in the CCTP message matches the immutable domain of this pool.
    if (destinationDomain != i_localDomainIdentifier) {
      revert InvalidDestinationDomain(i_localDomainIdentifier, destinationDomain);
    }

    if (minFinalityThreshold != FINALITY_THRESHOLD) {
      revert InvalidMinFinalityThreshold(FINALITY_THRESHOLD, minFinalityThreshold);
    }

    if (finalityThresholdExecuted != FINALITY_THRESHOLD) {
      revert InvalidExecutionFinalityThreshold(FINALITY_THRESHOLD, finalityThresholdExecuted);
    }

    // TODO: Add check for valid version of sourcePool.cctpVersion is VERSION_2
  }
}
