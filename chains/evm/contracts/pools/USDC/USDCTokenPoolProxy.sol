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
import {IERC165} from "@openzeppelin/contracts@4.8.3/utils/introspection/IERC165.sol";

/// @notice A token pool proxy for USDC that allows for routing of messages to the correct pool based on the correct
/// lock or burn mechanism. This includes CCTP v1, CCTP v2, and lock release.
/// @dev This contract will be listed in the Token Admin Registry as a token pool. All of the child pools which
/// receive the messages should have this contract set as an authorized caller. It does not inherit from the base
/// TokenPool contract but still implements the IPoolV1 interface.
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
contract USDCTokenPoolProxy is Ownable2StepMsgSender, IPoolV1, ITypeAndVersion {
  using SafeERC20 for IERC20;
  using ERC165CheckerReverting for address;

  error AddressCannotBeZero();
  error InvalidLockOrBurnMechanism(LockOrBurnMechanism mechanism);
  error InvalidMessageVersion(bytes4 version);
  error InvalidMessageLength(uint256 length);
  error MismatchedArrayLengths();
  error NoLockOrBurnMechanismSet(uint64 remoteChainSelector);
  error CallerIsNotARampOnRouter(address caller);
  error TokenPoolUnsupported(address pool);

  event LockOrBurnMechanismUpdated(uint64 indexed remoteChainSelector, LockOrBurnMechanism mechanism);
  event PoolAddressesUpdated(PoolAddresses pools);
  event LockReleasePoolUpdated(uint64 indexed remoteChainSelector, address lockReleasePool);

  struct PoolAddresses {
    address legacyCctpV1Pool; // A CCTP V1 token pool that did not utilize a message transmitter proxy.
    address cctpV1Pool;
    address cctpV2Pool;
  }

  enum LockOrBurnMechanism {
    INVALID_MECHANISM,
    CCTP_V1,
    CCTP_V2,
    LOCK_RELEASE
  }

  IERC20 internal immutable i_token;
  IRouter internal immutable i_router;

  mapping(uint64 remoteChainSelector => LockOrBurnMechanism mechanism) internal s_lockOrBurnMechanism;
  mapping(uint64 remoteChainSelector => address lockReleasePool) internal s_lockReleasePools;

  /// @dev The legacy CCTP V1, CCTP V1, and CCTP V2 pools which interact with CCTP contracts.
  PoolAddresses internal s_pools;

  string public constant override typeAndVersion = "USDCTokenPoolProxy 1.6.x-dev";

  constructor(IERC20 token, PoolAddresses memory pools, address router) {
    // Note: It is not required that every pool address be set, as this proxy may be deployed on a chain which does not
    // support a specific version of CCTP. As a result only the token and router are enforced to be non-zero.
    if (address(token) == address(0) || router == address(0)) {
      revert AddressCannotBeZero();
    }

    i_token = token;
    s_pools = pools;
    i_router = IRouter(router);
  }

  /// @notice Lock or Burn outgoing tokens to the correct pool based on the lock or burn mechanism.
  function lockOrBurn(
    Pool.LockOrBurnInV1 calldata lockOrBurnIn
  ) public virtual returns (Pool.LockOrBurnOutV1 memory) {
    // Since this contract does not inherit from the TokenPool contract, it must manually validate the caller as an onRamp.
    if (i_router.getOnRamp(lockOrBurnIn.remoteChainSelector) != msg.sender) {
      revert CallerIsNotARampOnRouter(msg.sender);
    }

    LockOrBurnMechanism mechanism = s_lockOrBurnMechanism[lockOrBurnIn.remoteChainSelector];

    // If a mechanism has not been configured for the remote chain selector, revert.
    if (mechanism == LockOrBurnMechanism.INVALID_MECHANISM) {
      revert InvalidLockOrBurnMechanism(mechanism);
    }

    PoolAddresses memory pools = s_pools;

    // The child pool which will perform the lock/burn operation.
    address childPool;

    if (mechanism == LockOrBurnMechanism.CCTP_V2) {
      childPool = pools.cctpV2Pool;
    } else if (mechanism == LockOrBurnMechanism.CCTP_V1) {
      childPool = pools.cctpV1Pool;
    } else if (mechanism == LockOrBurnMechanism.LOCK_RELEASE) {
      childPool = s_lockReleasePools[lockOrBurnIn.remoteChainSelector];
    }

    // If the destination pool is the zero address, then no mechanism has been configured for the outgoing tokens
    // and thus the destination chain is not supported and should revert.
    if (childPool == address(0)) {
      revert NoLockOrBurnMechanismSet(lockOrBurnIn.remoteChainSelector);
    }

    // Transfer the tokens to the proper child pool, as this contract is only a proxy and will not perform
    // the lock/burn itself.
    i_token.safeTransfer(childPool, lockOrBurnIn.amount);

    return USDCTokenPool(childPool).lockOrBurn(lockOrBurnIn);
  }

  /// @inheritdoc IPoolV1
  function isSupportedChain(
    uint64 remoteChainSelector
  ) external view returns (bool) {
    // If the outgoing mechanism is not set for a chain, then the chain is not supported because there cannot be a lock
    // or burn operation.
    return s_lockOrBurnMechanism[remoteChainSelector] != LockOrBurnMechanism.INVALID_MECHANISM;
  }

  /// @inheritdoc IPoolV1
  function isSupportedToken(
    address token
  ) external view returns (bool) {
    return address(i_token) == token;
  }

  /// @notice Signals which version of the pool interface is supported
  function supportsInterface(
    bytes4 interfaceId
  ) public pure override returns (bool) {
    return interfaceId == type(IPoolV1).interfaceId || interfaceId == Pool.CCIP_POOL_V1
      || interfaceId == type(IERC165).interfaceId;
  }

  /// @inheritdoc IPoolV1
  function releaseOrMint(
    Pool.ReleaseOrMintInV1 calldata releaseOrMintIn
  ) public virtual returns (Pool.ReleaseOrMintOutV1 memory) {
    // Since this proxy does not inherit from the TokenPool contract, it must manually validate the caller as an offRamp.
    if (!i_router.isOffRamp(releaseOrMintIn.remoteChainSelector, msg.sender)) {
      revert CallerIsNotARampOnRouter(msg.sender);
    }

    // The first 4 bytes of source pool data are the version which can be extracted directly and cast into a uint32.
    bytes4 version = bytes4(releaseOrMintIn.sourcePoolData[:4]);

    // If the source pool data is the lock release flag, use the lock release pool set for the remote chain selector.
    if (version == USDCSourcePoolDataCodec.LOCK_RELEASE_FLAG) {
      return USDCTokenPool(s_lockReleasePools[releaseOrMintIn.remoteChainSelector]).releaseOrMint(releaseOrMintIn);
    }

    if (version == USDCSourcePoolDataCodec.CCTP_VERSION_1_TAG) {
      return USDCTokenPool(s_pools.cctpV1Pool).releaseOrMint(releaseOrMintIn);
    }

    // Both tags will route to the same CCTP V2 pool, but will allow for pools to have greater granularity in deciding
    // the type of transfer (slow or fast) to use when depositing into CCTP.
    if (
      version == USDCSourcePoolDataCodec.CCTP_VERSION_2_TAG || version == USDCSourcePoolDataCodec.CCTP_VERSION_2_CCV_TAG
    ) {
      return USDCTokenPool(s_pools.cctpV2Pool).releaseOrMint(releaseOrMintIn);
    }

    // In previous versions of the USDC Token Pool, the sourcePoolData only contained two fields, a uint64 and uint32.
    // For structs stored only in memory, the compiler assigns each field to its own 32-byte slot, instead of tightly
    // packing like in storage. This means that a message originating from a previous version of the pool will have a
    // sourcePoolData that is 64 bytes long, indicating an inflight message originating from a previous version of
    // the USDC Token pool.
    // This branch must come before a version check, because the first field would be a uint64 and thus if a version
    // was attempted to be extracted from the first 4-bytes of a uint64, it would be 0, and thus the message would be
    // routed to the CCTP V1 pool without first sanitizing the source pool data for proper formatting.
    // Note: It is possible for a future version of the source pool data to also be 64 bytes long. However, any future
    // version will have a version number in the first 4 bytes and will be routed to the proper pool before this check
    // is reached. Therefore this branch will only be triggerd for messages using the legacy source pool data format.
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
        // need to turn the source pool data into the new format, otherwise the decoding scheme will fail. Once there is
        // confidence that no more messages are inflight, these branches can be safely removed.

        // Since the CCTP v1 pool will have this contract set as an allowed caller, no additional configurations are
        // needed to route the message to the v1 pool.
        return USDCTokenPool(s_pools.cctpV1Pool).releaseOrMint(_generateNewReleaseOrMintIn(releaseOrMintIn));
      }
    }

    revert InvalidMessageVersion(version);
  }

  /// @notice Update the pool addresses that this token pool will route a message to.
  /// @param pools The new pool addresses to update the token pool proxy with. Since the legacy CCTP V1 pool may not be
  /// used, the zero address is a valid input and therefore input sanitization for it is not required.
  function updatePoolAddresses(
    PoolAddresses calldata pools
  ) external onlyOwner {
    if (pools.cctpV1Pool != address(0) && !pools.cctpV1Pool._supportsInterfaceReverting(type(IPoolV1).interfaceId)) {
      revert TokenPoolUnsupported(pools.cctpV1Pool);
    }

    if (pools.cctpV2Pool != address(0) && !pools.cctpV2Pool._supportsInterfaceReverting(type(IPoolV1).interfaceId)) {
      revert TokenPoolUnsupported(pools.cctpV2Pool);
    }

    // If the legacy CCTP V1 Pool is being used, then it must support the IPoolV1 interface. If it is not, don't check it.
    if (
      pools.legacyCctpV1Pool != address(0)
        && !pools.legacyCctpV1Pool._supportsInterfaceReverting(type(IPoolV1).interfaceId)
    ) {
      revert TokenPoolUnsupported(pools.legacyCctpV1Pool);
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
        revert TokenPoolUnsupported(lockReleasePools[i]);
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
    // individually, which would be more gas intensive, as only the sourcePoolData field is going to be modified, as well
    // as the releaseOrMintIn struct is calldata, which cannot be modified in place.
    newReleaseOrMintIn = releaseOrMintIn;

    // While the legacy source pool data struct uses the same fields as the current source pool data struct, it is
    // was initially encoded using abi.encode(sourceTokenDataPayload) instead of the encoding scheme used in the
    // USDCSourcePoolDataCodec library, and without a version tag. Therefore, we need to decode the source pool data
    // into a SourceTokenDataPayloadV1 struct and then re-encode it into a format that using the proper versioning
    // scheme whereby the CCTP V1 pool can process the message.
    newReleaseOrMintIn.sourcePoolData = USDCSourcePoolDataCodec._encodeSourceTokenDataPayloadV1(
      abi.decode(releaseOrMintIn.sourcePoolData, (USDCSourcePoolDataCodec.SourceTokenDataPayloadV1))
    );

    return newReleaseOrMintIn;
  }
}
