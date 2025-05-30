// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ILiquidityContainer} from "../../interfaces/ILiquidityContainer.sol";

import {IMessageTransmitter} from "../USDC/IMessageTransmitter.sol";
import {ITokenMessenger} from "../USDC/ITokenMessenger.sol";

import {Pool} from "../../libraries/Pool.sol";
import {TokenPool} from "../TokenPool.sol";
import {CCTPMessageTransmitterProxy} from "../USDC/CCTPMessageTransmitterProxy.sol";
import {USDCTokenPool} from "../USDC/USDCTokenPool.sol";
import {USDCBridgeMigrator} from "./USDCBridgeMigrator.sol";

import {IERC20} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/utils/SafeERC20.sol";
import {EnumerableSet} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v4.8.3/contracts/utils/structs/EnumerableSet.sol";

// bytes4(keccak256("NO_CCTP_USE_LOCK_RELEASE"))
bytes4 constant LOCK_RELEASE_FLAG = 0xfa7c07de;

/// @notice A token pool for USDC which uses CCTP for supported chains and Lock/Release for all others
/// @dev The functionality from LockReleaseTokenPool.sol has been duplicated due to lack of compiler support for shared
/// constructors between parents
/// @dev The primary token mechanism in this pool is Burn/Mint with CCTP, with Lock/Release as the
/// secondary, opt in mechanism for chains not currently supporting CCTP.
contract HybridLockReleaseUSDCTokenPool is USDCTokenPool, USDCBridgeMigrator {
  using SafeERC20 for IERC20;
  using EnumerableSet for EnumerableSet.UintSet;

  event LiquidityTransferred(address indexed from, uint64 indexed remoteChainSelector, uint256 amount);
  event LiquidityProviderSet(
    address indexed oldProvider, address indexed newProvider, uint64 indexed remoteChainSelector
  );

  event LockReleaseEnabled(uint64 indexed remoteChainSelector);
  event LockReleaseDisabled(uint64 indexed remoteChainSelector);

  error LanePausedForCCTPMigration(uint64 remoteChainSelector);
  error TokenLockingNotAllowedAfterMigration(uint64 remoteChainSelector);

  error InvalidMinFinalityThreshold(uint32 expected, uint32 actual);
  error InvalidExecutionFinalityThreshold(uint32 expected, uint32 actual);
  error CCTPNotEnabledForRemoteChainSelector(uint64 remoteChainSelector);

  /// @notice The address of the liquidity provider for a specific chain.
  /// External liquidity is not required when there is one canonical token deployed to a chain,
  /// and CCIP is facilitating mint/burn on all the other chains, in which case the invariant
  /// balanceOf(pool) on home chain >= sum(totalSupply(mint/burn "wrapped" token) on all remote chains) should always hold
  mapping(uint64 remoteChainSelector => address liquidityProvider) internal s_liquidityProvider;

  ITokenMessenger internal immutable i_tokenMessengerV2;
  CCTPMessageTransmitterProxy internal immutable i_cctpMessageTransmitterProxyV2;

  // CCTP's max fee is based on the use of fast-burn. Since this pool does not utilize that feature, max fee should be 0.
  uint32 public constant MAX_FEE = 0;

  // CCTP V2 uses 2000 to indicate that attestations should not occur until finality is achieved on the source chain.
  uint32 public constant FINALITY_THRESHOLD = 2000;

  // TODO: Change this and the tests to go with it.
  enum CCTPVersion {
    // NOT_ENABLED,
    VERSION_0,
    VERSION_1
  }

  ITokenMessenger public immutable i_tokenMessengerCCTPV2;
  CCTPMessageTransmitterProxy public immutable i_cctpMessageTransmitterProxyCCTPV2;

  mapping(uint64 remoteChainSelector => CCTPVersion version) internal s_cctpVersion;

  constructor(
    ITokenMessenger tokenMessenger,
    ITokenMessenger tokenMessengerV2,
    CCTPMessageTransmitterProxy cctpMessageTransmitterProxy,
    CCTPMessageTransmitterProxy cctpMessageTransmitterProxyV2,
    IERC20 token,
    address[] memory allowlist,
    address rmnProxy,
    address router
  )
    // Since this inherits from the CCTP-V1 supported contract, we use USDC-V1 in the parent and manually check for V2
    USDCTokenPool(tokenMessenger, cctpMessageTransmitterProxy, token, allowlist, rmnProxy, router, 0)
    USDCBridgeMigrator(address(token))
  {

    // The following code has been duplicated exactly from USDC Token Pool but with the only difference being
    // checks for version 1 and uses tokenMessengerV2 and cctpMessageTransmitterProxyV2 for CCTP V2. 
    // This is because of Solidity's limited inheritance capabilities makes it impossible to run the constructor
    // for USDCTokenPool twice with different parameters, so it has been duplicated here.

    if (address(tokenMessengerV2) == address(0)) revert InvalidConfig();
    IMessageTransmitter transmitter = IMessageTransmitter(tokenMessengerV2.localMessageTransmitter());
    uint32 transmitterVersion = transmitter.version();
    if (transmitterVersion != 1) revert InvalidMessageVersion(transmitterVersion);
    uint32 tokenMessengerVersion = tokenMessengerV2.messageBodyVersion();
    if (tokenMessengerVersion != 1) revert InvalidTokenMessengerVersion(tokenMessengerVersion);
    if (cctpMessageTransmitterProxyV2.i_cctpTransmitter() != transmitter) revert InvalidTransmitterInProxy();

    emit ConfigSet(address(tokenMessenger));

    // Since CCTPV2 uses different messengers and transmitter addresses as CCTPV1, the addresses must be stored
    // as separate immutable variables from those inherited as part of USDCTokenPool.
    i_tokenMessengerCCTPV2 = tokenMessengerV2;
    i_cctpMessageTransmitterProxyCCTPV2 = cctpMessageTransmitterProxyV2;
    i_token.safeIncreaseAllowance(address(i_tokenMessengerCCTPV2), type(uint256).max);
  }

  // ================================================================
  // │                   Incoming/Outgoing Mechanisms               |
  // ================================================================

  /// @notice Locks the token in the pool
  /// @dev The _validateLockOrBurn check is an essential security check
  function lockOrBurn(
    Pool.LockOrBurnInV1 calldata lockOrBurnIn
  ) public virtual override returns (Pool.LockOrBurnOutV1 memory) {
    // // If the alternative mechanism (L/R) for chains which have it enabled
    if (!shouldUseLockRelease(lockOrBurnIn.remoteChainSelector)) {
      CCTPVersion cctpVersion = s_cctpVersion[lockOrBurnIn.remoteChainSelector];
      if (cctpVersion == CCTPVersion.VERSION_0) {
        return super.lockOrBurn(lockOrBurnIn);
      } else if (cctpVersion == CCTPVersion.VERSION_1) {
        return _lockOrBurnCCTPV2(lockOrBurnIn);
      }
      // Additional Safety Mechanism
      else {
        revert CCTPNotEnabledForRemoteChainSelector(lockOrBurnIn.remoteChainSelector);
      }
    }

    // Circle requires a supply-lock to prevent outgoing messages once the migration process begins.
    // This prevents new outgoing messages once the migration has begun to ensure any the procedure runs as expected
    if (s_proposedUSDCMigrationChain == lockOrBurnIn.remoteChainSelector) {
      revert LanePausedForCCTPMigration(s_proposedUSDCMigrationChain);
    }

    return _lockReleaseOutgoingMessage(lockOrBurnIn);
  }

  /// @notice Release tokens from the pool to the recipient
  /// @dev The _validateReleaseOrMint check is an essential security check
  function releaseOrMint(
    Pool.ReleaseOrMintInV1 calldata releaseOrMintIn
  ) public virtual override returns (Pool.ReleaseOrMintOutV1 memory) {
    // Use CCTP Burn/Mint mechanism for chains which have it enabled. The LOCK_RELEASE_FLAG is used in sourcePoolData to
    // discern this, since the source-chain will not be a hybrid-pool but a standard burn-mint. In the event of a
    // stuck message after a migration has occurred, and the message was not executed properly before the migration
    // began, and locked tokens were not released until now, the message will already have been committed to with this
    // flag so it is safe to release the tokens. The source USDC pool is trusted to send messages with the correct
    // flag as well.
    if (bytes4(releaseOrMintIn.sourcePoolData) != LOCK_RELEASE_FLAG) {
      CCTPVersion cctpVersion = s_cctpVersion[releaseOrMintIn.remoteChainSelector];
      if (cctpVersion == CCTPVersion.VERSION_0) {
        return super.releaseOrMint(releaseOrMintIn);
      } else if (cctpVersion == CCTPVersion.VERSION_1) {
        return _releaseOrMintCCTPV2(releaseOrMintIn);
      }
      // Additional Safety Mechanism
      else {
        revert CCTPNotEnabledForRemoteChainSelector(releaseOrMintIn.remoteChainSelector);
      }
    }
    return _lockReleaseIncomingMessage(releaseOrMintIn);
  }

  /// @notice Contains the alternative mechanism for incoming tokens, in this implementation is "Release" incoming tokens
  function _lockReleaseIncomingMessage(
    Pool.ReleaseOrMintInV1 calldata releaseOrMintIn
  ) internal virtual returns (Pool.ReleaseOrMintOutV1 memory) {
    _validateReleaseOrMint(releaseOrMintIn);

    // Circle requires a supply-lock to prevent incoming messages once the migration process begins.
    // This prevents new incoming messages once the migration has begun to ensure any the procedure runs as expected
    if (s_proposedUSDCMigrationChain == releaseOrMintIn.remoteChainSelector) {
      revert LanePausedForCCTPMigration(s_proposedUSDCMigrationChain);
    }

    // Decrease internal tracking of locked tokens to ensure accurate accounting for burnLockedUSDC() migration
    // If the chain has already been migrated, then this mapping would be zero, and the operation would underflow.
    // This branch ensures that we're subtracting from the correct mapping. It is also safe to subtract from the
    // excluded tokens mapping, as this function would only be invoked in the event of a stuck tx after a migration
    if (s_lockedTokensByChainSelector[releaseOrMintIn.remoteChainSelector] == 0) {
      s_tokensExcludedFromBurn[releaseOrMintIn.remoteChainSelector] -= releaseOrMintIn.amount;
    } else {
      s_lockedTokensByChainSelector[releaseOrMintIn.remoteChainSelector] -= releaseOrMintIn.amount;
    }

    getToken().safeTransfer(releaseOrMintIn.receiver, releaseOrMintIn.amount);

    emit Released(msg.sender, releaseOrMintIn.receiver, releaseOrMintIn.amount);

    return Pool.ReleaseOrMintOutV1({destinationAmount: releaseOrMintIn.amount});
  }

  /// @notice Contains the alternative mechanism, in this implementation is "Lock" on outgoing tokens
  function _lockReleaseOutgoingMessage(
    Pool.LockOrBurnInV1 calldata lockOrBurnIn
  ) internal virtual returns (Pool.LockOrBurnOutV1 memory) {
    _validateLockOrBurn(lockOrBurnIn);

    // Increase internal accounting of locked tokens for burnLockedUSDC() migration
    s_lockedTokensByChainSelector[lockOrBurnIn.remoteChainSelector] += lockOrBurnIn.amount;

    emit Locked(msg.sender, lockOrBurnIn.amount);

    return Pool.LockOrBurnOutV1({
      destTokenAddress: getRemoteToken(lockOrBurnIn.remoteChainSelector),
      destPoolData: abi.encode(LOCK_RELEASE_FLAG)
    });
  }

  // ================================================================
  // │                   Liquidity Management                       |
  // ================================================================

  /// @notice Gets LiquidityManager, can be address(0) if none is configured.
  /// @return The current liquidity manager for the given chain selector
  function getLiquidityProvider(
    uint64 remoteChainSelector
  ) external view returns (address) {
    return s_liquidityProvider[remoteChainSelector];
  }

  /// @notice Sets the LiquidityManager address.
  /// @dev Only callable by the owner.
  function setLiquidityProvider(uint64 remoteChainSelector, address liquidityProvider) external onlyOwner {
    address oldProvider = s_liquidityProvider[remoteChainSelector];

    s_liquidityProvider[remoteChainSelector] = liquidityProvider;

    emit LiquidityProviderSet(oldProvider, liquidityProvider, remoteChainSelector);
  }

  /// @notice Adds liquidity to the pool for a specific chain. The tokens should be approved first.
  /// @dev Liquidity is expected to be added on a per chain basis. Parties are expected to provide liquidity for their
  /// own chain which implements non canonical USDC and liquidity is not shared across lanes.
  /// @dev Once liquidity is added, it is locked in the pool until it is removed by an incoming message on the
  /// lock release mechanism. This is a hard requirement by Circle to ensure parity with the destination chain
  /// supply is maintained.
  /// @param amount The amount of tokens to provide as liquidity.
  /// @param remoteChainSelector The chain for which liquidity is provided to. Necessary to ensure there's accurate
  /// parity between locked USDC in this contract and the circulating supply on the remote chain
  function provideLiquidity(uint64 remoteChainSelector, uint256 amount) external {
    if (s_liquidityProvider[remoteChainSelector] != msg.sender) revert TokenPool.Unauthorized(msg.sender);

    // Prevent adding liquidity to a chain which has already been migrated
    if (s_migratedChains.contains(remoteChainSelector)) {
      revert TokenLockingNotAllowedAfterMigration(remoteChainSelector);
    }

    // prevent adding liquidity to a chain which has been proposed for migration
    if (remoteChainSelector == s_proposedUSDCMigrationChain) {
      revert LanePausedForCCTPMigration(remoteChainSelector);
    }

    s_lockedTokensByChainSelector[remoteChainSelector] += amount;

    i_token.safeTransferFrom(msg.sender, address(this), amount);

    emit ILiquidityContainer.LiquidityAdded(msg.sender, amount);
  }

  /// @notice Removed liquidity to the pool. The tokens will be sent to msg.sender.
  /// @param remoteChainSelector The chain where liquidity is being released.
  /// @param amount The amount of liquidity to remove.
  /// @dev The function should only be called if non canonical USDC on the remote chain has been burned and is not being
  /// withdrawn on this chain, otherwise a mismatch may occur between locked token balance and remote circulating supply
  /// which may block a potential future migration of the chain to CCTP.
  function withdrawLiquidity(uint64 remoteChainSelector, uint256 amount) external onlyOwner {
    // A supply-lock is required to prevent outgoing messages once the migration process begins.
    // This prevents new outgoing messages once the migration has begun to ensure any the procedure runs as expected
    if (remoteChainSelector == s_proposedUSDCMigrationChain) {
      revert LanePausedForCCTPMigration(remoteChainSelector);
    }

    s_lockedTokensByChainSelector[remoteChainSelector] -= amount;

    i_token.safeTransfer(msg.sender, amount);

    emit ILiquidityContainer.LiquidityRemoved(msg.sender, amount);
  }

  // ================================================================
  // │                   CCTPV2 Logic
  // ================================================================

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
  function _releaseOrMintCCTPV2(
    Pool.ReleaseOrMintInV1 calldata releaseOrMintIn
  ) internal virtual returns (Pool.ReleaseOrMintOutV1 memory) {
    _validateReleaseOrMint(releaseOrMintIn);

    uint32 sourceDomainIdentifier = abi.decode(releaseOrMintIn.sourcePoolData, (uint32));

    MessageAndAttestation memory msgAndAttestation =
      abi.decode(releaseOrMintIn.offchainTokenData, (MessageAndAttestation));

    _validateMessageCCTPV2(msgAndAttestation.message, sourceDomainIdentifier);

    if (!i_cctpMessageTransmitterProxyCCTPV2.receiveMessage(msgAndAttestation.message, msgAndAttestation.attestation)) {
      revert UnlockingUSDCFailed();
    }

    emit Minted(msg.sender, releaseOrMintIn.receiver, releaseOrMintIn.amount);
    return Pool.ReleaseOrMintOutV1({destinationAmount: releaseOrMintIn.amount});
  }

  /// @notice Validates the USDC encoded message against the given parameters.
  /// @param usdcMessage The USDC encoded message
  /// @param sourcePoolDomain The expected source chain CCTP identifier as provided by the CCIP-Source-Pool.
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
  function _validateMessageCCTPV2(bytes memory usdcMessage, uint32 sourcePoolDomain) internal view {
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
    if (messageSourceDomain != sourcePoolDomain) {
      revert InvalidSourceDomain(sourcePoolDomain, messageSourceDomain);
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
  }

  /// @notice Burn tokens from the pool to initiate cross-chain transfer.
  /// @notice Outgoing messages (burn operations) are routed via `i_tokenMessenger.depositForBurnWithCaller`.
  /// The allowedCaller is preconfigured per destination domain and token pool version refer Domain struct.
  /// @dev Emits ITokenMessenger.DepositForBurn event.
  /// @dev Assumes caller has validated the destinationReceiver.
  function _lockOrBurnCCTPV2(
    Pool.LockOrBurnInV1 calldata lockOrBurnIn
  ) internal virtual returns (Pool.LockOrBurnOutV1 memory) {
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
    i_tokenMessengerV2.depositForBurn(
      lockOrBurnIn.amount, // amount
      domain.domainIdentifier, // destinationDomain
      decodedReceiver, // mintRecipient
      address(i_token), // burnToken
      domain.allowedCaller, // destinationCaller
      MAX_FEE, // maxFee
      FINALITY_THRESHOLD // minFinalityThreshold
    );

    emit Burned(msg.sender, lockOrBurnIn.amount);

    return Pool.LockOrBurnOutV1({
      destTokenAddress: getRemoteToken(lockOrBurnIn.remoteChainSelector),
      destPoolData: abi.encode(i_localDomainIdentifier)
    });
  }

  // TODO: Natspec.
  function updateCCTPVersion(
    uint64[] calldata remoteChainSelectors,
    CCTPVersion[] calldata versions
  ) external onlyOwner {
    for (uint256 i = 0; i < remoteChainSelectors.length; ++i) {
      s_cctpVersion[remoteChainSelectors[i]] = versions[i];
      // TODO: Emit Event
    }
  }

  // ================================================================
  // │                   Alt Mechanism Logic                        |
  // ================================================================

  /// @notice Return whether a lane should use the alternative L/R mechanism in the token pool.
  /// @param remoteChainSelector the remote chain the lane is interacting with
  /// @return bool Return true if the alternative L/R mechanism should be used, and is decided by the Owner
  function shouldUseLockRelease(
    uint64 remoteChainSelector
  ) public view virtual returns (bool) {
    return s_shouldUseLockRelease[remoteChainSelector];
  }

  /// @notice Updates designations for chains on whether to use primary or alt mechanism on CCIP messages
  /// @param removes A list of chain selectors to disable Lock-Release, and enforce BM
  /// @param adds A list of chain selectors to enable LR instead of BM. These chains must not have been migrated
  /// to CCTP yet or the transaction will revert
  function updateChainSelectorMechanisms(uint64[] calldata removes, uint64[] calldata adds) external onlyOwner {
    for (uint256 i = 0; i < removes.length; ++i) {
      delete s_shouldUseLockRelease[removes[i]];
      emit LockReleaseDisabled(removes[i]);
    }

    for (uint256 i = 0; i < adds.length; ++i) {
      // Prevent enabling lock release on chains which have already been migrated
      if (s_migratedChains.contains(adds[i])) {
        revert TokenLockingNotAllowedAfterMigration(adds[i]);
      }
      s_shouldUseLockRelease[adds[i]] = true;
      emit LockReleaseEnabled(adds[i]);
    }
  }
}
