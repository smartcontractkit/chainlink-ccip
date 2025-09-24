// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolV1} from "../../interfaces/IPool.sol";
import {IRouter} from "../../interfaces/IRouter.sol";
import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";

import {ERC165CheckerReverting} from "../../libraries/ERC165CheckerReverting.sol";
import {Pool} from "../../libraries/Pool.sol";
import {USDCSourcePoolDataCodec} from "../../libraries/USDCSourcePoolDataCodec.sol";
import {USDCTokenPool} from "./USDCTokenPool.sol";

import {Ownable2StepMsgSender} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2StepMsgSender.sol";
import {IERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/utils/SafeERC20.sol";

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
///     ├──→ LegacyCCTPV1Pool → CCTPV1
///     ├──→ CCTPV1Pool → MessageTransmitterProxy/TokenMessenger V1 → CCTPV1
///     ├──→ CCTPV2Pool → MessageTransmitterProxy/TokenMessenger V2 → CCTPV2
///     └──→ SiloedUSDCTokenPool → ERC20LockBox
contract USDCTokenPoolProxy is Ownable2StepMsgSender, ITypeAndVersion {
  using SafeERC20 for IERC20;
  using ERC165CheckerReverting for address;

  error AddressCannotBeZero();
  error InvalidLockOrBurnMechanism(LockOrBurnMechanism mechanism);
  error InvalidMessageVersion(uint32 version);
  error InvalidMessageLength(uint256 length);
  error MismatchedArrayLengths();
  error InvalidDestinationPool();
  error Unauthorized();
  error TokenPoolUnsupported();

  event LockOrBurnMechanismUpdated(uint64 indexed remoteChainSelector, LockOrBurnMechanism mechanism);
  event PoolAddressesUpdated(PoolAddresses pools);
  event LockReleasePoolUpdated(uint64 indexed remoteChainSelector, address lockReleasePool);

  struct PoolAddresses {
    address legacyCctpV1Pool; // A v1 token pool that did not utilize a message transmitter proxy.
    address cctpV1Pool;
    address cctpV2Pool;
  }

  enum LockOrBurnMechanism {
    INVALID_MECHANISM,
    CCTP_V1,
    CCTP_V2,
    LOCK_RELEASE
  }

  mapping(uint64 remoteChainSelector => LockOrBurnMechanism mechanism) internal s_lockOrBurnMechanism;
  mapping(uint64 remoteChainSelector => address lockReleasePool) internal s_lockReleasePools;

  PoolAddresses internal s_pools;

  IERC20 internal immutable i_token;
  IRouter internal immutable i_router;

  string public constant override typeAndVersion = "USDCTokenPoolProxy 1.6.3-dev";

  constructor(IERC20 token, PoolAddresses memory pools, address router) {
    // Note: The legacy pool is allowed to be zero, as it is not requireed if this proxy is being deployed
    // on a chain which has already migrated to a pool that utilizes a message transmitter proxy.
    if (
      address(token) == address(0) || pools.cctpV1Pool == address(0) || pools.cctpV2Pool == address(0)
        || router == address(0)
    ) {
      revert AddressCannotBeZero();
    }

    i_token = token;
    s_pools = pools;
    i_router = IRouter(router);
  }

  function lockOrBurn(
    Pool.LockOrBurnInV1 calldata lockOrBurnIn
  ) public virtual returns (Pool.LockOrBurnOutV1 memory) {
    if (i_router.getOnRamp(lockOrBurnIn.remoteChainSelector) != msg.sender) {
      revert Unauthorized();
    }

    LockOrBurnMechanism mechanism = s_lockOrBurnMechanism[lockOrBurnIn.remoteChainSelector];

    // If a mechanism has not been configured for the remote chain selector, revert.
    if (mechanism == LockOrBurnMechanism.INVALID_MECHANISM) {
      revert InvalidLockOrBurnMechanism(mechanism);
    }

    PoolAddresses memory pools = s_pools;

    address destinationPool;

    // The order of the branches is based on the expected frequency of each mechanism being used, in order to save
    // gas on every call.
    if (mechanism == LockOrBurnMechanism.CCTP_V2) {
      destinationPool = pools.cctpV2Pool;
    } else if (mechanism == LockOrBurnMechanism.CCTP_V1) {
      destinationPool = pools.cctpV1Pool;
    } else if (mechanism == LockOrBurnMechanism.LOCK_RELEASE) {
      destinationPool = s_lockReleasePools[lockOrBurnIn.remoteChainSelector];
    }

    if (destinationPool == address(0)) {
      revert InvalidDestinationPool();
    }

    // Transfer the tokens to the proper child pool, as this contract is only a proxy and will not perform
    // the lock/burn itself.
    i_token.safeTransfer(destinationPool, lockOrBurnIn.amount);

    return USDCTokenPool(destinationPool).lockOrBurn(lockOrBurnIn);
  }

  function releaseOrMint(
    Pool.ReleaseOrMintInV1 calldata releaseOrMintIn
  ) public virtual returns (Pool.ReleaseOrMintOutV1 memory) {
    if (!i_router.isOffRamp(releaseOrMintIn.remoteChainSelector, msg.sender)) {
      revert Unauthorized();
    }

    // If the source pool data is the lock release flag, we use the lock release pool set for the remote chain selector.
    if (bytes4(releaseOrMintIn.sourcePoolData) == LOCK_RELEASE_FLAG) {
      return USDCTokenPool(s_lockReleasePools[releaseOrMintIn.remoteChainSelector]).releaseOrMint(releaseOrMintIn);
    }

    // In previous versions of the USDC Token Pool, the sourcePoolData only contained two fields, a uint64 and uint32.
    // For structs stored only in memory, the compiler assigns each field to its own 32-byte slot, instead of tightly
    // packing like in storage. This means that a message originating from a previous version of the pool will have a
    // sourcePoolData that is 64 bytes long, indicating an inflight message originating from a previous version of
    // the USDC Token pool.
    if (releaseOrMintIn.sourcePoolData.length == 64) {
      // There are two possible scenarios for the legacy inflight messages:
      // 1. The legacy pool did not utilize a message transmitter proxy.
      // 2. The legacy pool utilized a message transmitter proxy, but the format of the sourcePoolTokenData was as described
      // in the comments above.

      // In the first scenario, only the message's destinationCaller, I.E the legacy pool, can execute the mint, and so
      // the message needs to be routed to the legacy pool. In the second scenario, the destinationCaller will be the
      // message transmitter proxy, and the message needs to be routed to the appropriate V1-compatible pool.
      if (_checkForLegacyInflightMessages(releaseOrMintIn.offchainTokenData)) {
        // Note: Supporting this branch will require this proxy to be set as an offRamp in the router, which is a design
        // decision that is not ideal, but allows for a direct upgrade from the first version of the USDC Token Pool to
        // this version.
        return USDCTokenPool(s_pools.legacyCctpV1Pool).releaseOrMint(releaseOrMintIn);
      } else {
        // Since the new pool and the inflight message should utilize the same version of CCTP, and would have the same
        // destinationCaller (the message transmitter proxy), we can route the message to the v1 pool, but we first
        // need to turn the source pool data into the new format, otherwise abi-decoding will revert. Once there is
        // confidence that no more messages are inflight, these branches can be safely removed.

        // Since the CCTP v1 pool will have this contract set as an allowed caller, no additional configurations are
        // needed to route the message to the v1 pool.
        return USDCTokenPool(s_pools.cctpV1Pool).releaseOrMint(_generateNewReleaseOrMintIn(releaseOrMintIn));
      }
    }

    // According to the CCTP specification, the first 4 bytes of the message are the version, which we can extract
    // directly and cast into a uint32.
    uint32 version = uint32(bytes4(releaseOrMintIn.sourcePoolData[0:4]));

    if (version == 0) {
      return USDCTokenPool(s_pools.cctpV1Pool).releaseOrMint(releaseOrMintIn);
    } else if (version == 1) {
      return USDCTokenPool(s_pools.cctpV2Pool).releaseOrMint(releaseOrMintIn);
    } else {
      revert InvalidMessageVersion(version);
    }
  }

  /// @notice Update the pool addresses that this token pool will route a message to.
  /// @param pools The new pool addresses to update the token pool proxy with. Since the legacy CCTP V1 pool may not be
  /// used, the zero address is a valid input and therefore input sanitization for it is not required.
  function updatePoolAddresses(
    PoolAddresses calldata pools
  ) external onlyOwner {
    if (pools.cctpV1Pool == address(0) || pools.cctpV2Pool == address(0)) {
      revert AddressCannotBeZero();
    }

    // If the V1 or V2 Pool does not support the IPoolV1 interface, revert.
    // If the legacy CCTP V1 Pool is being used, then it must support the IPoolV1 interface.
    if (
      !pools.cctpV1Pool._supportsInterfaceReverting(type(IPoolV1).interfaceId)
        || !pools.cctpV2Pool._supportsInterfaceReverting(type(IPoolV1).interfaceId)
        || (
          pools.legacyCctpV1Pool != address(0)
            && !pools.legacyCctpV1Pool._supportsInterfaceReverting(type(IPoolV1).interfaceId)
        )
    ) {
      revert TokenPoolUnsupported();
    }

    s_pools = pools;

    emit PoolAddressesUpdated(pools);
  }

  /// @notice Get the current pool addresses that this token pool will route a message to.
  /// @return The current pool addresses that this token pool will route a message to.
  function getPools() public view returns (PoolAddresses memory) {
    return s_pools;
  }

  /// @notice Get the lock or burn mechanism for a given remote chain selector.
  /// @param remoteChainSelector The remote chain selector to get the mechanism for.
  /// @return The lock or burn mechanism for the given remote chain selector, including CCTP V1/V2 and Lock/Release
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

    for (uint256 i = 0; i < remoteChainSelectors.length; ++i) {
      s_lockOrBurnMechanism[remoteChainSelectors[i]] = mechanisms[i];
      emit LockOrBurnMechanismUpdated(remoteChainSelectors[i], mechanisms[i]);
    }
  }

  /// @notice Update the lock release pool addresses for a list of remote chain selectors.
  /// @param remoteChainSelectors The remote chain selectors to update the lock release pool addresses for.
  /// @param lockReleasePools The new lock release pool addresses for the given remote chain selectors.
  function updateLockReleasePoolAddresses(
    uint64[] calldata remoteChainSelectors,
    address[] calldata lockReleasePools
  ) external onlyOwner {
    if (remoteChainSelectors.length != lockReleasePools.length) {
      revert MismatchedArrayLengths();
    }

    for (uint256 i = 0; i < remoteChainSelectors.length; ++i) {
      // If the token pool is being added, ensure that it supports the token pool v1 interface. If the pool is the zero address,
      // then it is being removed, as a migration from L/R to CCTP, and therefore no check is needed, as it was
      // already performed when originally added.
      if (
        lockReleasePools[i] != address(0) && !lockReleasePools[i]._supportsInterfaceReverting(type(IPoolV1).interfaceId)
      ) {
        revert TokenPoolUnsupported();
      }

      // Note: Since the lock release pool is only used for chains that do not have CCTP support, after a successful
      // migration to CCTP, the lock release pool may no longer be needed, and therefore the zero address is
      // a valid input and input validation is not required. It is also why no check for the mechanism being
      // LOCK_RELEASE is performed either, as after a migration this may no longer be the case.
      s_lockReleasePools[remoteChainSelectors[i]] = lockReleasePools[i];
      emit LockReleasePoolUpdated(remoteChainSelectors[i], lockReleasePools[i]);
    }
  }

  /// @notice Get the lock release pool address for a given remote chain selector.
  /// @param remoteChainSelector The remote chain selector to get the lock release pool address for.
  /// @return The lock release pool address for the given remote chain selector.
  function getLockReleasePoolAddress(
    uint64 remoteChainSelector
  ) public view returns (address) {
    return s_lockReleasePools[remoteChainSelector];
  }

  /// @notice Check if the releaseOrMintIn struct is an inflight message from a legacy pool that did not utilize a
  /// message transmitter proxy.
  /// @param offChainTokenData The off chain message and attestation needed to check for destinationCaller.
  /// @return True if the releaseOrMintIn struct is an inflight message from a legacy pool that did not utilize a
  /// message transmitter proxy, false otherwise.
  function _checkForLegacyInflightMessages(
    bytes calldata offChainTokenData
  ) internal view virtual returns (bool) {
    // Cache the legacy pool address to avoid multiple SLOADs.
    address legacyPool = s_pools.legacyCctpV1Pool;

    // If the legacy pool without a proxy is not set, then there is no need to check the destinationCaller.
    if (legacyPool == address(0)) {
      return false;
    }

    bytes memory messageBytes = abi.decode(offChainTokenData, (USDCTokenPool.MessageAndAttestation)).message;

    bytes32 destinationCallerBytes32;
    assembly {
      // destinationCaller is a 32-byte word starting at position 84 in messageBytes body, so add 32 to skip the 1st word
      // representing bytes length.
      destinationCallerBytes32 := mload(add(messageBytes, 116)) // 84 + 32 = 116
    }
    address destinationCaller = address(uint160(uint256(destinationCallerBytes32)));

    return destinationCaller == legacyPool;
  }

  /// @notice Converts a legacy sourcePoolData struct into a new format that can be used to release or mint USDC on the
  /// previous pool. This is necessary because the sourcePoolData is stored in a different format in the previous pool,
  /// and must be in a properly decodable format.
  /// @param releaseOrMintIn The releaseOrMintIn struct to generate a new struct for.
  /// @return newReleaseOrMintIn The new releaseOrMintIn struct.
  function _generateNewReleaseOrMintIn(
    Pool.ReleaseOrMintInV1 calldata releaseOrMintIn
  ) internal pure returns (Pool.ReleaseOrMintInV1 memory newReleaseOrMintIn) {
    // Copy the releaseOrMintIn struct to the newReleaseOrMintIn struct. We do this to avoid having to copy each field
    // individually, which would be more gas intensive, as only the sourcePoolData field is going to be modified.
    newReleaseOrMintIn = releaseOrMintIn;

    uint32 nonce = uint32(bytes4(releaseOrMintIn.sourcePoolData[0:4]));
    uint32 sourceDomain = uint32(bytes4(releaseOrMintIn.sourcePoolData[4:8]));

    // Since this is a legacy message, it should only operate on CCTP V1 messages. As a result it is safe to hard
    // code the version to 0.
    newReleaseOrMintIn.sourcePoolData = USDCSourcePoolDataCodec._encodeSourceTokenDataPayloadV0(
      bytes4(0), USDCTokenPool.SourceTokenDataPayloadV0({nonce: nonce, sourceDomain: sourceDomain})
    );

    return newReleaseOrMintIn;
  }
}
