// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ICrossChainVerifierResolver} from "../../interfaces/ICrossChainVerifierResolver.sol";
import {IPoolV1, IPoolV1V2, IPoolV2} from "../../interfaces/IPoolV1V2.sol";
import {IRouter} from "../../interfaces/IRouter.sol";
import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";

import {Pool} from "../../libraries/Pool.sol";
import {USDCSourcePoolDataCodec} from "../../libraries/USDCSourcePoolDataCodec.sol";
import {Ownable2StepMsgSender} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2StepMsgSender.sol";

import {IERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/utils/SafeERC20.sol";
import {ERC165Checker} from "@openzeppelin/contracts@5.3.0/utils/introspection/ERC165Checker.sol";
import {IERC165} from "@openzeppelin/contracts@5.3.0/utils/introspection/IERC165.sol";

/// @notice A token pool proxy for USDC that allows for routing of messages to the correct pool based on the correct
/// lock or burn mechanism. This includes CCTP v1, CCTP v2, CCTP v2 with CCV, and lock release.
/// @dev This contract will be listed in the Token Admin Registry as a token pool. All of the child pools which
/// receive the messages should have this contract set as an authorized caller. It does not inherit from the base
/// TokenPool contract but still implements the IPoolV2 interface.
contract USDCTokenPoolProxy is Ownable2StepMsgSender, IPoolV1V2, ITypeAndVersion {
  using SafeERC20 for IERC20;
  using ERC165Checker for address;

  error AddressCannotBeZero();
  error CCVCompatiblePoolNotSet();
  error ChainNotSupportedByVerifier(uint64 remoteChainSelector);
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

  struct MessageAndAttestation {
    bytes message;
    bytes attestation;
  }

  struct PoolAddresses {
    address cctpV1Pool;
    address cctpV2Pool;
    address cctpTokenPool;
    address siloedUsdcTokenPool;
  }

  enum LockOrBurnMechanism {
    INVALID_MECHANISM,
    CCTP_V1,
    CCTP_V2,
    LOCK_RELEASE,
    CCTP_V2_WITH_CCV
  }

  IERC20 internal immutable i_token;
  IRouter internal immutable i_router;
  ICrossChainVerifierResolver private immutable i_cctpVerifier;

  mapping(uint64 remoteChainSelector => LockOrBurnMechanism mechanism) internal s_lockOrBurnMechanism;

  /// @dev This token pool should have minimal state, as it is only used to route messages to the correct
  /// pool. If more mechanisms are needed, such as a new CCTP version, then this contract should be updated
  /// to include the proper routing logic and reference the appropriate child pool.
  /// On/OffRamp
  ///     ↓
  /// USDCPoolProxy
  ///     ├──→ CCTPV1Pool → MessageTransmitterProxy/TokenMessenger V1 → CCTPV1
  ///     ├──→ CCTPV2Pool → MessageTransmitterProxy/TokenMessenger V2 → CCTPV2
  ///     ├──→ CCTPTokenPool → CCTPVerifier → MessageTransmitterProxy/TokenMessenger V2 → CCTPV2
  ///     └──→ SiloedUSDCTokenPool → ERC20LockBox
  address internal s_cctpV1Pool;
  address internal s_cctpV2Pool;
  address internal s_cctpTokenPool;
  address internal s_siloedUsdcTokenPool;

  /// @dev Constant representing the default finality.
  uint16 internal constant WAIT_FOR_FINALITY = 0;

  string public constant override typeAndVersion = "USDCTokenPoolProxy 1.7.0-dev";

  constructor(
    IERC20 token,
    PoolAddresses memory pools,
    address router,
    address cctpVerifier
  ) {
    // Note: It is not required that every pool address be set, as this proxy may be deployed on a chain which does not support a specific version of CCTP.
    // As a result only the token, router, and cctpVerifier are enforced to be non-zero.
    if (address(token) == address(0) || router == address(0) || cctpVerifier == address(0)) {
      revert AddressCannotBeZero();
    }

    i_token = token;
    i_router = IRouter(router);
    i_cctpVerifier = ICrossChainVerifierResolver(cctpVerifier);

    s_cctpV1Pool = pools.cctpV1Pool;
    s_cctpV2Pool = pools.cctpV2Pool;
    s_cctpTokenPool = pools.cctpTokenPool;
    s_siloedUsdcTokenPool = pools.siloedUsdcTokenPool;
  }

  /// @inheritdoc IPoolV1
  /// @notice Lock or burn outgoing tokens to the correct pool based on the lock or burn mechanism.
  /// @param lockOrBurnIn Encoded data fields for the processing of tokens on the source chain.
  function lockOrBurn(
    Pool.LockOrBurnInV1 calldata lockOrBurnIn
  ) public virtual override returns (Pool.LockOrBurnOutV1 memory) {
    (Pool.LockOrBurnOutV1 memory lockOrBurnOut,) = lockOrBurn(lockOrBurnIn, WAIT_FOR_FINALITY, "");
    return lockOrBurnOut;
  }

  /// @inheritdoc IPoolV2
  /// @notice Lock or burn outgoing tokens to the correct pool based on the lock or burn mechanism.
  /// @param lockOrBurnIn Encoded data fields for the processing of tokens on the source chain.
  /// @param blockConfirmationRequested Requested block confirmation.
  /// @param tokenArgs Additional token arguments.
  function lockOrBurn(
    Pool.LockOrBurnInV1 calldata lockOrBurnIn,
    uint16 blockConfirmationRequested,
    bytes memory tokenArgs
  ) public virtual returns (Pool.LockOrBurnOutV1 memory lockOrBurnOut, uint256 destTokenAmount) {
    // Since this contract does not inherit from the TokenPool contract, it must manually validate the caller as an onRamp.
    if (i_router.getOnRamp(lockOrBurnIn.remoteChainSelector) != msg.sender) {
      revert CallerIsNotARampOnRouter(msg.sender);
    }

    LockOrBurnMechanism mechanism = s_lockOrBurnMechanism[lockOrBurnIn.remoteChainSelector];

    // If a mechanism has not been configured for the remote chain selector, revert.
    if (mechanism == LockOrBurnMechanism.INVALID_MECHANISM) {
      revert InvalidLockOrBurnMechanism(mechanism);
    }

    if (mechanism == LockOrBurnMechanism.CCTP_V2_WITH_CCV) {
      // CCV-compatible lockOrBurn path is completed within this if statement to avoid redundant checks.
      address ccvPool = s_cctpTokenPool;
      if (ccvPool == address(0)) {
        revert NoLockOrBurnMechanismSet(lockOrBurnIn.remoteChainSelector);
      }
      // If using the CCTP verifier, transfer funds to the verifier instead of the pool.
      // First ensure that the chain is supported by the verifier.
      address verifierImpl = i_cctpVerifier.getOutboundImplementation(lockOrBurnIn.remoteChainSelector, tokenArgs);
      if (verifierImpl == address(0)) {
        revert ChainNotSupportedByVerifier(lockOrBurnIn.remoteChainSelector);
      }
      i_token.safeTransfer(verifierImpl, lockOrBurnIn.amount);

      return IPoolV2(ccvPool).lockOrBurn(lockOrBurnIn, blockConfirmationRequested, tokenArgs);
    }

    // The child pool which will perform the lock/burn operation.
    address childPool;

    if (mechanism == LockOrBurnMechanism.CCTP_V2) {
      childPool = s_cctpV2Pool;
    } else if (mechanism == LockOrBurnMechanism.CCTP_V1) {
      childPool = s_cctpV1Pool;
    } else if (mechanism == LockOrBurnMechanism.LOCK_RELEASE) {
      childPool = s_siloedUsdcTokenPool;
    }

    // If the destination pool is the zero address, then no mechanism has been configured for the outgoing tokens
    // and thus the destination chain is not supported and should revert.
    if (childPool == address(0)) {
      revert NoLockOrBurnMechanismSet(lockOrBurnIn.remoteChainSelector);
    }

    // Transfer the tokens to the correct address, as this contract is only a proxy and will not perform the lock/burn itself.
    i_token.safeTransfer(childPool, lockOrBurnIn.amount);

    return (IPoolV1(childPool).lockOrBurn(lockOrBurnIn), lockOrBurnIn.amount);
  }

  /// @inheritdoc IPoolV1
  /// @dev If the outgoing mechanism is not set for a chain, then the chain is not supported because there cannot be a
  /// lock or burn operation.
  function isSupportedChain(
    uint64 remoteChainSelector
  ) external view returns (bool) {
    return s_lockOrBurnMechanism[remoteChainSelector] != LockOrBurnMechanism.INVALID_MECHANISM;
  }

  /// @inheritdoc IPoolV1
  function isSupportedToken(
    address token
  ) external view returns (bool) {
    return address(i_token) == token;
  }

  /// @inheritdoc IPoolV1
  /// @param releaseOrMintIn Encoded data fields for the processing of tokens on the destination chain.
  function releaseOrMint(
    Pool.ReleaseOrMintInV1 calldata releaseOrMintIn
  ) public virtual override returns (Pool.ReleaseOrMintOutV1 memory) {
    return releaseOrMint(releaseOrMintIn, WAIT_FOR_FINALITY);
  }

  /// @inheritdoc IPoolV2
  /// @param releaseOrMintIn Encoded data fields for the processing of tokens on the destination chain.
  /// @param blockConfirmationRequested Requested block confirmation.
  function releaseOrMint(
    Pool.ReleaseOrMintInV1 calldata releaseOrMintIn,
    uint16 blockConfirmationRequested
  ) public virtual returns (Pool.ReleaseOrMintOutV1 memory) {
    // Since this proxy does not inherit from the TokenPool contract, it must manually validate the caller as an offRamp.
    if (!i_router.isOffRamp(releaseOrMintIn.remoteChainSelector, msg.sender)) {
      revert CallerIsNotARampOnRouter(msg.sender);
    }

    // The first 4 bytes of source pool data are the version which can be extracted directly and cast into a uint32.
    bytes4 version = bytes4(releaseOrMintIn.sourcePoolData[:4]);

    // If the source pool data is the lock release flag, use the lock release pool set for the remote chain selector.
    if (version == USDCSourcePoolDataCodec.LOCK_RELEASE_FLAG) {
      return IPoolV1(s_siloedUsdcTokenPool).releaseOrMint(releaseOrMintIn);
    }

    if (version == USDCSourcePoolDataCodec.CCTP_VERSION_1_TAG) {
      return IPoolV1(s_cctpV1Pool).releaseOrMint(releaseOrMintIn);
    }

    if (version == USDCSourcePoolDataCodec.CCTP_VERSION_2_TAG) {
      return IPoolV1(s_cctpV2Pool).releaseOrMint(releaseOrMintIn);
    }

    if (version == USDCSourcePoolDataCodec.CCTP_VERSION_2_CCV_TAG) {
      return IPoolV2(s_cctpTokenPool).releaseOrMint(releaseOrMintIn, blockConfirmationRequested);
    }

    // In previous versions of the USDC Token Pool, the sourcePoolData only contained two fields, a uint64 and uint32.
    // For structs stored only in memory, the compiler assigns each field to its own 32-byte slot, instead of tightly
    // packing like in storage. This means that a message originating from a previous version of the pool will have a
    // sourcePoolData that is 64 bytes long, indicating an inflight message originating from a previous version of
    // the USDC Token pool.
    // Note: It is possible for a future version of the source pool data to also be 64 bytes long. However, any future
    // version will have a version number in the first 4 bytes and will be routed to the proper pool before this check
    // is reached. Therefore this branch will only be triggered for messages using the legacy source pool data format.
    if (releaseOrMintIn.sourcePoolData.length == 64) {
      // Since the CCTP v1 pool will have this contract set as an allowed caller, no additional configurations are
      // needed to route the message to the v1 pool.

      Pool.ReleaseOrMintInV1 memory newReleaseOrMintIn = releaseOrMintIn;
      // While the legacy source pool data struct uses the same fields as the current source pool data struct, it is
      // was initially encoded using abi.encode(sourceTokenDataPayload) instead of the encoding scheme used in the
      // USDCSourcePoolDataCodec library, and without a version tag. Therefore, we need to decode the source pool data
      // into a SourceTokenDataPayloadV1 struct and then re-encode it into a format that using the proper versioning
      // scheme whereby the CCTP V1 pool can process the message.
      newReleaseOrMintIn.sourcePoolData = USDCSourcePoolDataCodec._encodeSourceTokenDataPayloadV1(
        abi.decode(releaseOrMintIn.sourcePoolData, (USDCSourcePoolDataCodec.SourceTokenDataPayloadV1))
      );

      return IPoolV1(s_cctpV1Pool).releaseOrMint(newReleaseOrMintIn);
    }

    revert InvalidMessageVersion(version);
  }

  /// @notice Update the pool addresses that this token pool will route a message to.
  /// @param pools The new pool addresses to update the token pool proxy with. Since the legacy CCTP V1 pool may not be
  /// used, the zero address is a valid input and therefore input sanitization for it is not required.
  function updatePoolAddresses(
    PoolAddresses calldata pools
  ) external onlyOwner {
    if (pools.cctpV1Pool != address(0) && !pools.cctpV1Pool.supportsInterface(type(IPoolV1).interfaceId)) {
      revert TokenPoolUnsupported(pools.cctpV1Pool);
    }

    if (pools.cctpV2Pool != address(0) && !pools.cctpV2Pool.supportsInterface(type(IPoolV1).interfaceId)) {
      revert TokenPoolUnsupported(pools.cctpV2Pool);
    }

    if (pools.cctpTokenPool != address(0)) {
      if (!pools.cctpTokenPool.supportsInterface(type(IPoolV2).interfaceId)) {
        revert TokenPoolUnsupported(pools.cctpTokenPool);
      }
    }

    // If the siloed USDC pool is being used, then it must support the IPoolV1 interface. If it is not, don't check it.
    if (
      pools.siloedUsdcTokenPool != address(0) && !pools.siloedUsdcTokenPool.supportsInterface(type(IPoolV1).interfaceId)
    ) {
      revert TokenPoolUnsupported(pools.siloedUsdcTokenPool);
    }

    s_cctpV1Pool = pools.cctpV1Pool;
    s_cctpV2Pool = pools.cctpV2Pool;
    s_cctpTokenPool = pools.cctpTokenPool;
    s_siloedUsdcTokenPool = pools.siloedUsdcTokenPool;

    emit PoolAddressesUpdated(pools);
  }

  /// @notice Get the current pool addresses that this token pool will route a message to.
  /// @return The current pool addresses that this token pool will route a message to.
  function getPools() public view returns (PoolAddresses memory) {
    return PoolAddresses({
      cctpV1Pool: s_cctpV1Pool,
      cctpV2Pool: s_cctpV2Pool,
      cctpTokenPool: s_cctpTokenPool,
      siloedUsdcTokenPool: s_siloedUsdcTokenPool
    });
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

  /// @inheritdoc IPoolV2
  /// @param localToken The local asset being transferred.
  /// @param destChainSelector The destination lane selector.
  /// @param amount The amount of tokens being bridged on this lane.
  /// @param feeToken The token used to pay feeUSDCents.
  /// @param blockConfirmationRequested Requested block confirmation.
  /// @param tokenArgs Opaque token arguments supplied by the caller.
  function getFee(
    address localToken,
    uint64 destChainSelector,
    uint256 amount,
    address feeToken,
    uint16 blockConfirmationRequested,
    bytes calldata tokenArgs
  )
    external
    view
    onlyWithCCVCompatiblePool
    returns (uint256 feeUSDCents, uint32 destGasOverhead, uint32 destBytesOverhead, uint16 tokenFeeBps, bool isEnabled)
  {
    return IPoolV2(s_cctpTokenPool)
      .getFee(localToken, destChainSelector, amount, feeToken, blockConfirmationRequested, tokenArgs);
  }

  /// @inheritdoc IPoolV2
  /// @param localToken The local asset being transferred.
  /// @param destChainSelector The chain selector of the destination chain.
  /// @param blockConfirmationRequested Requested block confirmation.
  /// @param tokenArgs Additional token argument from the CCIP message.
  function getTokenTransferFeeConfig(
    address localToken,
    uint64 destChainSelector,
    uint16 blockConfirmationRequested,
    bytes calldata tokenArgs
  ) external view onlyWithCCVCompatiblePool returns (TokenTransferFeeConfig memory feeConfig) {
    return IPoolV2(s_cctpTokenPool)
      .getTokenTransferFeeConfig(localToken, destChainSelector, blockConfirmationRequested, tokenArgs);
  }

  /// @inheritdoc IPoolV2
  /// @param remoteChainSelector Remote chain selector.
  function getRemoteToken(
    uint64 remoteChainSelector
  ) external view onlyWithCCVCompatiblePool returns (bytes memory) {
    return IPoolV2(s_cctpTokenPool).getRemoteToken(remoteChainSelector);
  }

  /// @inheritdoc IPoolV2
  /// @dev Instead of calling the pool, we take a shortcut and return the CCTPVerifier as required directly.
  function getRequiredCCVs(
    address, // localToken
    uint64 remoteChainSelector,
    uint256, // amount
    uint16, // blockConfirmationRequested
    bytes calldata, // extraData
    MessageDirection // direction
  ) external view onlyWithCCVCompatiblePool returns (address[] memory requiredCCVs) {
    if (s_lockOrBurnMechanism[remoteChainSelector] == LockOrBurnMechanism.INVALID_MECHANISM) {
      revert NoLockOrBurnMechanismSet(remoteChainSelector);
    }

    // Common case: The lockOrBurn mechanism is CCTP V2 with CCV.
    // In this case, we simply need to return the CCTP CCV.
    address[] memory ccvs = new address[](1);
    if (s_lockOrBurnMechanism[remoteChainSelector] == LockOrBurnMechanism.CCTP_V2_WITH_CCV) {
      ccvs[0] = address(i_cctpVerifier);
      return ccvs;
    }

    // If using lock-release, we can't specify CCTP because CCTP won't ultimately be called.
    // Other CCTP mechanisms will never rely on CCVs and have no impact on the return value.
    // Therefore, we return address(0) to indicate that default CCVs should be used for the lock-release mechanism.
    return ccvs;
  }

  /// @notice Ensures that a CCV-compatible pool is set.
  modifier onlyWithCCVCompatiblePool() {
    if (s_cctpTokenPool == address(0)) {
      revert CCVCompatiblePoolNotSet();
    }
    _;
  }

  /// @inheritdoc IERC165
  function supportsInterface(
    bytes4 interfaceId
  ) public pure override returns (bool) {
    return interfaceId == type(IPoolV2).interfaceId || interfaceId == type(IPoolV1).interfaceId
      || interfaceId == Pool.CCIP_POOL_V1 || interfaceId == type(IERC165).interfaceId;
  }
}
