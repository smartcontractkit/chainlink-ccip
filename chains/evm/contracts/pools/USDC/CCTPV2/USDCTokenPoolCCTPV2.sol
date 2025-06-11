// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ITokenMessenger} from "../interfaces/ITokenMessenger.sol";

import {CCTPV2} from "../../../libraries/CCTPV2.sol";
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

  // Value: 2000 indicates "slow burn" mode where attestation waits for source chain finality
  // Alternative: 1000 would indicate "fast burn" mode (not used by this pool)  uint32 public constant FINALITY_THRESHOLD = 2000;
  address public immutable i_previousPool;

  constructor(
    ITokenMessenger tokenMessenger,
    CCTPMessageTransmitterProxy cctpMessageTransmitterProxy,
    IERC20 token,
    address[] memory allowlist,
    address rmnProxy,
    address router,
    address previousPool
  ) USDCTokenPool(tokenMessenger, cctpMessageTransmitterProxy, token, allowlist, rmnProxy, router, 1) {
    i_previousPool = previousPool;

    CCTPV2._validateConfig(tokenMessenger, cctpMessageTransmitterProxy);
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

    USDCTokenPool.Domain storage domain = s_chainToDomain[lockOrBurnIn.remoteChainSelector];

    if (!domain.enabled) revert UnknownDomain(lockOrBurnIn.remoteChainSelector);

    if (lockOrBurnIn.receiver.length != 32) {
      revert InvalidReceiver(lockOrBurnIn.receiver);
    }

    // To support certain non-EVM chains, the mint recipient may be overridden to be a token pool which then
    // forwards the tokens to the receiver. The message itself will not be changed and the destination token pool will
    // still receive the correct address of the final token receiver.
    bytes32 decodedReceiver;
    if (domain.mintRecipient != bytes32(0)) {
      decodedReceiver = domain.mintRecipient;
    } else {
      decodedReceiver = abi.decode(lockOrBurnIn.receiver, (bytes32));
    }

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
      CCTPV2.MAX_FEE, // maxFee
      CCTPV2.FINALITY_THRESHOLD // minFinalityThreshold
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

    // This supports legacy inflight messages which were sent using CCTP V1. The sourcePoolData will be 64 bytes long, as the
    // cctpVersion field will not be present. The message must be proxied to the previous pool to satisfy the allowedCaller.
    // Any messages sent after the migration will have the cctpVersion field, and will be handled by the CCTP V2 functionality.
    // There should not be any messages sent on V1 that include the cctpVersion field, as it was not supported initially,
    // and the field was added after the migration, therefore the only messages that should be sent on V1 are legacy messages.
    if (releaseOrMintIn.sourcePoolData.length == 64) {
      return USDCTokenPool(i_previousPool).releaseOrMint(releaseOrMintIn);
    }

    SourceTokenDataPayload memory sourceTokenData = abi.decode(releaseOrMintIn.sourcePoolData, (SourceTokenDataPayload));

    MessageAndAttestation memory msgAndAttestation =
      abi.decode(releaseOrMintIn.offchainTokenData, (MessageAndAttestation));

    CCTPV2._validateMessage(msgAndAttestation.message, sourceTokenData, i_localDomainIdentifier);

    if (
      !i_messageTransmitterProxy.receiveMessage(
        msgAndAttestation.message, msgAndAttestation.attestation, sourceTokenData.cctpVersion
      )
    ) {
      revert UnlockingUSDCFailed();
    }

    emit ReleasedOrMinted({
      remoteChainSelector: releaseOrMintIn.remoteChainSelector,
      token: address(i_token),
      sender: msg.sender,
      recipient: releaseOrMintIn.receiver,
      amount: releaseOrMintIn.amount
    });

    return Pool.ReleaseOrMintOutV1({destinationAmount: releaseOrMintIn.amount});
  }
}
