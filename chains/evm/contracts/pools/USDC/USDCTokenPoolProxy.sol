// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";

import {Pool} from "../../libraries/Pool.sol";
import {TokenPool} from "../TokenPool.sol";
import {USDCTokenPool} from "./USDCTokenPool.sol";

import {IERC20} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/utils/SafeERC20.sol";

/// @dev The flag used to indicate that the source pool data is coming from a chain that does not have CCTP Support,
/// and so the lock release pool should be used. The BurnMintWithLockReleaseTokenPool uses this flag as its source pool
/// data to indicate that the tokens should be released from the lock release pool rather than attempting to be minted
/// through CCTP.
/// @dev The preimage is bytes4(keccak256("NO_CCTP_USE_LOCK_RELEASE")).
bytes4 constant LOCK_RELEASE_FLAG = 0xfa7c07de;

/// @notice A token pool proxy for USDC that allows for routing of messages to the correct pool based on the correct
/// lock or burn mechanism. This includes CCTP v1, CCTP v2, and lock release.
/// @dev This contract will be listed in the Token Admin Registry as a token pool. All of the child pools which
/// receive the messages should have this contract set as an authorized caller.
/// @dev This token pool should have minimal state, as it is only used to route messages to the correct
/// pool. If more mechanisms are needed, such as a new CCTP version, then this contract should be updated
/// to include the proper routing logic and reference the appropriate child pool.
/// On/OffRamp
///     ↓
/// USDCPoolProxy
///     ├──→ CCTPV1Pool → MessageTransmitterProxy/TokenMessenger V1 → CCTPV1
///     ├──→ CCTPV2Pool → MessageTransmitterProxy/TokenMessenger V2 → CCTPV2
///     └──→ SiloedUSDCTokenPool → ERC20LockBox
contract USDCTokenPoolProxy is TokenPool, ITypeAndVersion {
  using SafeERC20 for IERC20;

  event LockOrBurnMechanismUpdated(uint64 indexed remoteChainSelector, LockOrBurnMechanism mechanism);
  event PoolAddressesUpdated(PoolAddresses pools);

  error PoolAddressCannotBeZero();
  error InvalidLockOrBurnMechanism(LockOrBurnMechanism mechanism);
  error InvalidMessageVersion(uint32 version);
  error InvalidMessageLength(uint256 length);

  struct PoolAddresses {
    address cctpV1Pool;
    address cctpV2Pool;
    address lockReleasePool;
  }

  enum LockOrBurnMechanism {
    INVALID_MECHANISM,
    CCTP_V1,
    CCTP_V2,
    LOCK_RELEASE
  }

  PoolAddresses internal s_pools;

  mapping(uint64 remoteChainSelector => LockOrBurnMechanism) internal s_lockOrBurnMechanism;

  string public constant override typeAndVersion = "USDCTokenPoolProxy 1.6.2-dev";

  constructor(
    IERC20 token,
    address router,
    address[] memory allowlist,
    address rmnProxy,
    PoolAddresses memory pools
  ) TokenPool(token, 6, allowlist, rmnProxy, router) {
    if (pools.cctpV1Pool == address(0) || pools.cctpV2Pool == address(0) || pools.lockReleasePool == address(0)) {
      revert PoolAddressCannotBeZero();
    }

    s_pools = pools;
  }

  /// @inheritdoc TokenPool
  function lockOrBurn(
    Pool.LockOrBurnInV1 calldata lockOrBurnIn
  ) public virtual override returns (Pool.LockOrBurnOutV1 memory) {
    _validateLockOrBurn(lockOrBurnIn);

    LockOrBurnMechanism mechanism = s_lockOrBurnMechanism[lockOrBurnIn.remoteChainSelector];

    // If a mechanism has not been configured for the remote chain selector, revert.
    if (mechanism == LockOrBurnMechanism.INVALID_MECHANISM) {
      revert InvalidLockOrBurnMechanism(mechanism);
    }

    PoolAddresses memory pools = s_pools;

    address destinationPool;

    if (mechanism == LockOrBurnMechanism.LOCK_RELEASE) {
      destinationPool = pools.lockReleasePool;
    } else if (mechanism == LockOrBurnMechanism.CCTP_V1) {
      destinationPool = pools.cctpV1Pool;
    } else if (mechanism == LockOrBurnMechanism.CCTP_V2) {
      destinationPool = pools.cctpV2Pool;
    }

    // Transfer the tokens to the destination pool otherwise any burn or transfer will revert due to insufficient balance.
    i_token.safeTransfer(destinationPool, lockOrBurnIn.amount);

    return USDCTokenPool(destinationPool).lockOrBurn(lockOrBurnIn);
  }

  /// @inheritdoc TokenPool
  function releaseOrMint(
    Pool.ReleaseOrMintInV1 calldata releaseOrMintIn
  ) public virtual override returns (Pool.ReleaseOrMintOutV1 memory) {
    _validateReleaseOrMint(releaseOrMintIn, releaseOrMintIn.sourceDenominatedAmount);

    PoolAddresses memory pools = s_pools;

    // If the source pool data is the lock release flag, we use the lock release pool.
    if (bytes4(releaseOrMintIn.sourcePoolData) == LOCK_RELEASE_FLAG) {
      return USDCTokenPool(pools.lockReleasePool).releaseOrMint(releaseOrMintIn);
    }

    // In previous versions of the USDC Token Pool, the sourcePoolData only contained two fields, a uint64 and uint32.
    // For structs stored only in memory, the compiler assigns each field to its own 32-byte slot, instead of tightly
    // packing like in storage. This means that a message originating from a previous version of the pool will have a
    // sourcePoolData that is 64 bytes long, indicating an inflight message originating from a previous version of
    // the USDC Token pool.
    if (releaseOrMintIn.sourcePoolData.length == 64) {
      // Since the new pool and the inflight message should utilize the same version of CCTP, and would have the same
      // destinationCaller (the message transmitter proxy), we can route the message to the v1 pool, but we first
      // need to turn the source pool data into the new format so that the abi-decoding will succeed. Once there is
      // confidence that no more messages are inflight, these branches can be safely removed.

      // While adding some complexity to the code, this is preferable than having to maintain support for the legacy token
      // pool. Since that legacy pool has an Only-OffRamp check, this proxy would have to be set as an offRamp in the
      // router, which is a design decision that is not ideal and presents additional security risks. This mechanism
      // instead allows for the legacy pool to be removed, and the proxy to be updated to only have to utilize the new
      // version of the USDC Token Pool.
      Pool.ReleaseOrMintInV1 memory legacyReleaseOrMintIn = _generateNewReleaseOrMintIn(releaseOrMintIn);

      // Since the CCTP v1 pool will have this contract set as an allowed caller, no additional configurations are
      // needed to route the message to the v1 pool.
      return USDCTokenPool(pools.cctpV1Pool).releaseOrMint(legacyReleaseOrMintIn);
    }

    // In both version 1 and 2 of CCTP, the version is stored in the first 4 bytes of the message, so this check is
    // valid for both versions. If this changes in future versions, this will need to be updated.
    bytes calldata usdcMessage = releaseOrMintIn.offchainTokenData;
    // Check the first 4 bytes of the message to prevent an out-of-bounds read.
    if (usdcMessage.length < 4) {
      revert InvalidMessageLength(usdcMessage.length);
    }

    // According to the CCTP spec, the first 4 bytes of the message are the version, which we can extract
    // directly and cast into a uint32.
    uint32 version = uint32(bytes4(usdcMessage[0:4]));

    if (version == 0) {
      return USDCTokenPool(pools.cctpV1Pool).releaseOrMint(releaseOrMintIn);
    } else if (version == 1) {
      return USDCTokenPool(pools.cctpV2Pool).releaseOrMint(releaseOrMintIn);
    } else {
      revert InvalidMessageVersion(version);
    }
  }

  /// @notice Update the pool addresses that this token pool will route a message to.
  /// @param pools The new pool addresses to update the token pool proxy with.
  function updatePoolAddresses(
    PoolAddresses calldata pools
  ) external onlyOwner {
    if (pools.cctpV1Pool == address(0) || pools.cctpV2Pool == address(0) || pools.lockReleasePool == address(0)) {
      revert PoolAddressCannotBeZero();
    }

    s_pools = pools;

    emit PoolAddressesUpdated(pools);
  }

  function getPools() public view returns (PoolAddresses memory) {
    return s_pools;
  }

  /// @notice Get the lock or burn mechanism for a given remote chain selector.
  /// @param remoteChainSelector The remote chain selector to get the lock or burn mechanism for.
  /// @return The lock or burn mechanism for the given remote chain selector.
  function getLockOrBurnMechanism(
    uint64 remoteChainSelector
  ) public view returns (LockOrBurnMechanism) {
    return s_lockOrBurnMechanism[remoteChainSelector];
  }

  /// @notice Update the lock or burn mechanism for a list of remote chain selectors.
  /// @param remoteChainSelectors The remote chain selectors to update the lock or burn mechanism for.
  /// @param mechanisms The new lock or burn mechanisms for the given remote chain selectors.
  /// @dev Only callable by the owner.
  function updateLockOrBurnMechanisms(
    uint64[] calldata remoteChainSelectors,
    LockOrBurnMechanism[] calldata mechanisms
  ) external onlyOwner {
    if (remoteChainSelectors.length != mechanisms.length) {
      revert MismatchedArrayLengths();
    }

    for (uint256 i = 0; i < remoteChainSelectors.length; i++) {
      if (!isSupportedChain(remoteChainSelectors[i])) {
        revert TokenPool.NonExistentChain(remoteChainSelectors[i]);
      }

      s_lockOrBurnMechanism[remoteChainSelectors[i]] = mechanisms[i];
      emit LockOrBurnMechanismUpdated(remoteChainSelectors[i], mechanisms[i]);
    }
  }

  /// @notice Converts a legacy sourcePoolData struct into a new format that can be used to release or mint USDC on the
  /// previous pool. This is necessary because the sourcePoolData is stored in a different format in the previous pool,
  /// and must be in a properly decodable format.
  /// @param releaseOrMintIn The releaseOrMintIn struct to generate a new struct for.
  /// @return newReleaseOrMintIn The new releaseOrMintIn struct.
  function _generateNewReleaseOrMintIn(
    Pool.ReleaseOrMintInV1 calldata releaseOrMintIn
  ) internal view returns (Pool.ReleaseOrMintInV1 memory newReleaseOrMintIn) {
    // Copy the releaseOrMintIn struct to the newReleaseOrMintIn struct. We do this to avoid having to copy each field
    // individually, which would be more gas intensive, as only the sourcePoolData field is going to be modified.
    newReleaseOrMintIn = releaseOrMintIn;

    // Get the local domain identifier and the transmitter proxy address from the previous pool.
    uint32 localDomain = USDCTokenPool(s_pools.cctpV1Pool).i_localDomainIdentifier();
    bytes32 allowedCaller =
      bytes32(uint256(uint160(address(USDCTokenPool(s_pools.cctpV1Pool).i_messageTransmitterProxy()))));

    // Decode the legacy sourcePoolData to get the nonce and sourceDomain. The original sourcePoolData was a struct
    // with two fields, a uint64 and a uint32. We can decode it into two variables, directly, nonce and sourceDomain.
    (uint64 nonce, uint32 sourceDomain) = abi.decode(releaseOrMintIn.sourcePoolData, (uint64, uint32));

    // Create the new payload out of the legacy sourcePoolData.
    USDCTokenPool.SourceTokenDataPayload memory newPayload = USDCTokenPool.SourceTokenDataPayload({
      nonce: nonce,
      sourceDomain: sourceDomain,
      cctpVersion: USDCTokenPool.CCTPVersion.CCTP_V1,
      amount: releaseOrMintIn.sourceDenominatedAmount,
      destinationDomain: localDomain,
      mintRecipient: bytes32(uint256(uint160(releaseOrMintIn.receiver))), // Cast the receiver address to a bytes32.
      burnToken: releaseOrMintIn.localToken,
      destinationCaller: allowedCaller,
      maxFee: 0, // Since maxFee is not used in CCTP V1, we set it to 0.
      minFinalityThreshold: 0 // Since minFinalityThreshold is not used in CCTP V1, we set it to 0.
    });

    // Encode the new payload into the sourcePoolData field of the newReleaseOrMintIn struct.
    newReleaseOrMintIn.sourcePoolData = abi.encode(newPayload);

    return newReleaseOrMintIn;
  }
}
