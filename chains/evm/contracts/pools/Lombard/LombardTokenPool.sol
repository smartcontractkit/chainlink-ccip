// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {ICrossChainVerifierResolver} from "../../interfaces/ICrossChainVerifierResolver.sol";
import {IBridgeV1} from "./interfaces/IBridgeV1.sol";
import {IMailbox} from "./interfaces/IMailbox.sol";
import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";

import {Pool} from "../../libraries/Pool.sol";
import {TokenPool} from "../TokenPool.sol";

import {IERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/IERC20.sol";
import {IERC20Metadata} from "@openzeppelin/contracts@4.8.3/token/ERC20/extensions/IERC20Metadata.sol";
import {SafeERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/utils/SafeERC20.sol";

/// @notice Lombard CCIP token pool.
/// For v2 flows, token movement (burn/mint) is handled by the Lombard verifier,
/// the pool performs validation, rate limiting, accounting and event emission.
/// IPoolV2.lockOrBurn forwards tokens to the verifier.
/// IPoolV2.releaseOrMint does not move tokens, _releaseOrMint is a no-op.
/// IPoolV1.lockOrBurn and IPoolV1.releaseOrMint make this pool backwards compatible with old lanes.
contract LombardTokenPool is TokenPool, ITypeAndVersion {
  using SafeERC20 for IERC20;
  using SafeERC20 for IERC20Metadata;

  error ZeroVerifierNotAllowed();
  error OutboundImplementationNotFoundForVerifier();
  error ZeroBridge();
  error ZeroLombardChainId();
  error PathNotExist(uint64 remoteChainSelector);
  error InvalidMessageVersion(uint8 expected, uint8 received);
  error RemoteTokenMismatch(bytes32 bridge, bytes32 pool);
  error InvalidReceiver(bytes receiver);
  error ChainNotSupported(uint64 remoteChainSelector);
  error InvalidAllowedCaller(bytes allowedCaller);
  error ExecutionError();
  error HashMismatch();

  /// The following events are emitted for Lombard-specific configuration updates and are utilized by Lombard.
  /// @param remoteChainSelector CCIP selector of destination chain.
  /// @param lChainId The chain ID according to Lombard Multi Chain ID convention.
  /// @param allowedCaller The address that's allowed to call the bridge on the destination chain.
  event PathSet(uint64 indexed remoteChainSelector, bytes32 indexed lChainId, bytes32 allowedCaller);
  /// @param remoteChainSelector CCIP selector of destination chain.
  /// @param lChainId The chain id of destination chain by Lombard Multi Chain Id conversion.
  /// @param allowedCaller The address that's allowed to call the bridge on the destination chain.
  event PathRemoved(uint64 indexed remoteChainSelector, bytes32 indexed lChainId, bytes32 allowedCaller);
  event LombardConfigurationSet(address indexed verifier, address indexed bridge, address indexed tokenAdapter);

  struct Path {
    /// @notice The address that's allowed to call the bridge on the destination chain.
    bytes32 allowedCaller;
    /// @notice Lombard chain id of destination chain.
    bytes32 lChainId;
  }

  string public constant override typeAndVersion = "LombardTokenPool 1.7.0-dev";

  /// @notice Supported bridge message version.
  uint8 internal constant SUPPORTED_BRIDGE_MSG_VERSION = 1;
  /// @notice The address of bridge contract.
  IBridgeV1 public immutable i_bridge;
  /// @notice Lombard verifier resolver address. lockOrBurn fetches the outbound implementation and forwards tokens to it.
  address internal immutable i_lombardVerifierResolver;
  /// @notice Optional token adapter used for chains like Avalanche BTC.b. Since each pool manages a single token,
  /// and the adapter is a source-chain-level replacement for that token, there can only be one adapter per pool.
  address internal immutable i_tokenAdapter;

  /// @notice Mapping of CCIP chain selector to chain specific config.
  mapping(uint64 chainSelector => Path path) internal s_chainSelectorToPath;

  /// @param verifier The address of Lombard verifier resolver. Used in V2 flows to fetch the outbound
  /// implementation that handles token burns and cross-chain attestations.
  /// @param bridge The Lombard BridgeV2 contract that handles cross-chain token transfers.
  /// @param adapter Optional source-chain token address override. Used for non-upgradeable tokens like BTC.b
  /// on Avalanche where an adapter contract performs mint/burn on behalf of the actual token. When set, this
  /// address is passed to bridge.deposit() instead of the pool's token address. Set to address(0) if not needed.
  constructor(
    IERC20Metadata token,
    address verifier,
    IBridgeV1 bridge,
    address adapter,
    address advancedPoolHooks,
    address rmnProxy,
    address router,
    uint8 fallbackDecimals
  ) TokenPool(token, _getTokenDecimals(token, fallbackDecimals), advancedPoolHooks, rmnProxy, router) {
    if (address(bridge) == address(0)) {
      revert ZeroBridge();
    }
    uint8 bridgeMsgVersion = bridge.MSG_VERSION();
    if (bridgeMsgVersion != SUPPORTED_BRIDGE_MSG_VERSION) {
      revert InvalidMessageVersion(SUPPORTED_BRIDGE_MSG_VERSION, bridgeMsgVersion);
    }
    if (verifier == address(0)) {
      revert ZeroVerifierNotAllowed();
    }
    i_bridge = bridge;
    i_lombardVerifierResolver = verifier;
    i_tokenAdapter = adapter;
    emit LombardConfigurationSet(verifier, address(bridge), adapter);
  }

  // ================================================================
  // │                        Lock or Burn                          │
  // ================================================================

  /// @notice For IPoolV2.lockOrBurn call, this contract only forwards tokens to the verifier.
  /// @dev Forward the net amount to the verifier; actual burn/bridge is done there.
  function lockOrBurn(
    Pool.LockOrBurnInV1 calldata lockOrBurnIn,
    uint16 blockConfirmationRequested,
    bytes calldata tokenArgs
  ) public override returns (Pool.LockOrBurnOutV1 memory lockOrBurnOut, uint256 destTokenAmount) {
    address verifierImpl = ICrossChainVerifierResolver(i_lombardVerifierResolver).getOutboundImplementation(
      lockOrBurnIn.remoteChainSelector, ""
    );
    if (verifierImpl == address(0)) {
      revert OutboundImplementationNotFoundForVerifier();
    }
    i_token.safeTransfer(verifierImpl, lockOrBurnIn.amount);
    return super.lockOrBurn(lockOrBurnIn, blockConfirmationRequested, tokenArgs);
  }

  /// @notice Backwards compatible lockOrBurn for lanes using the V1 flow.
  /// @dev Token minting is performed by the Lombard bridge's mailbox during deliverAndHandle.
  /// This pool only validates the proof and emits events; no _lockOrBurn call is needed.
  function lockOrBurn(
    Pool.LockOrBurnInV1 calldata lockOrBurnIn
  ) public override(TokenPool) returns (Pool.LockOrBurnOutV1 memory lockOrBurnOut) {
    _validateLockOrBurn(lockOrBurnIn, WAIT_FOR_FINALITY, "");

    Path memory path = s_chainSelectorToPath[lockOrBurnIn.remoteChainSelector];
    if (path.allowedCaller == bytes32(0)) {
      revert PathNotExist(lockOrBurnIn.remoteChainSelector);
    }

    // For some tokens we need to override the source token with an adapter
    address sourceTokenOrAdapter = i_tokenAdapter != address(0) ? i_tokenAdapter : address(i_token);
    // verify bridge destination token equal to pool
    bytes32 bridgeDestToken = i_bridge.getAllowedDestinationToken(path.lChainId, sourceTokenOrAdapter);
    bytes32 poolDestToken = abi.decode(getRemoteToken(lockOrBurnIn.remoteChainSelector), (bytes32));
    if (bridgeDestToken != poolDestToken) {
      revert RemoteTokenMismatch(bridgeDestToken, poolDestToken);
    }

    if (lockOrBurnIn.receiver.length != 32) {
      revert InvalidReceiver(lockOrBurnIn.receiver);
    }

    (, bytes32 payloadHash) = i_bridge.deposit({
      destinationChain: path.lChainId,
      token: sourceTokenOrAdapter,
      sender: lockOrBurnIn.originalSender,
      recipient: abi.decode(lockOrBurnIn.receiver, (bytes32)),
      amount: lockOrBurnIn.amount,
      destinationCaller: path.allowedCaller
    });

    emit LockedOrBurned({
      remoteChainSelector: lockOrBurnIn.remoteChainSelector,
      token: address(i_token),
      sender: lockOrBurnIn.originalSender,
      amount: lockOrBurnIn.amount
    });

    return Pool.LockOrBurnOutV1({
      destTokenAddress: getRemoteToken(lockOrBurnIn.remoteChainSelector),
      destPoolData: abi.encode(payloadHash)
    });
  }

  // ================================================================
  // │                      Release or Mint                         │
  // ================================================================

  /// @notice Backwards compatible releaseOrMint for CCIP 1.5/1.6 lanes. Verifies the bridge payload proof.
  function releaseOrMint(
    Pool.ReleaseOrMintInV1 calldata releaseOrMintIn
  ) public virtual override returns (Pool.ReleaseOrMintOutV1 memory) {
    _validateReleaseOrMint(releaseOrMintIn, releaseOrMintIn.sourceDenominatedAmount, WAIT_FOR_FINALITY);

    (bytes memory rawPayload, bytes memory proof) = abi.decode(releaseOrMintIn.offchainTokenData, (bytes, bytes));

    (bytes32 payloadHash, bool executed,) = IMailbox(i_bridge.mailbox()).deliverAndHandle(rawPayload, proof);
    if (!executed) {
      revert ExecutionError();
    }
    // we know payload hash returned on source chain.
    if (payloadHash != abi.decode(releaseOrMintIn.sourcePoolData, (bytes32))) {
      revert HashMismatch();
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

  // ================================================================
  // │                         Path config                          │
  // ================================================================

  /// @notice Gets the path for a given CCIP chain selector.
  /// @param remoteChainSelector CCIP chain selector of remote chain.
  /// @return Path struct containing lChainId and allowedCaller.
  function getPath(
    uint64 remoteChainSelector
  ) external view returns (Path memory) {
    return s_chainSelectorToPath[remoteChainSelector];
  }

  /// @notice Sets the Lombard chain id and allowed caller for a CCIP chain selector.
  /// @param remoteChainSelector CCIP chain selector of remote chain.
  /// @param lChainId Lombard chain id of remote chain.
  /// @param allowedCaller The address of TokenPool on destination chain.
  function setPath(uint64 remoteChainSelector, bytes32 lChainId, bytes calldata allowedCaller) external onlyOwner {
    if (!isSupportedChain(remoteChainSelector)) {
      revert ChainNotSupported(remoteChainSelector);
    }

    if (lChainId == bytes32(0)) {
      revert ZeroLombardChainId();
    }

    // only remote pool is expected allowed caller.
    if (!isRemotePool(remoteChainSelector, allowedCaller)) {
      revert InvalidRemotePoolForChain(remoteChainSelector, allowedCaller);
    }

    if (allowedCaller.length != 32) {
      revert InvalidAllowedCaller(allowedCaller);
    }
    bytes32 decodedAllowedCaller = abi.decode(allowedCaller, (bytes32));

    s_chainSelectorToPath[remoteChainSelector] = Path({lChainId: lChainId, allowedCaller: decodedAllowedCaller});

    emit PathSet(remoteChainSelector, lChainId, decodedAllowedCaller);
  }

  /// @notice Removes path mapping for a destination chain.
  /// @param remoteChainSelector CCIP chain selector of destination chain.
  function removePath(
    uint64 remoteChainSelector
  ) external onlyOwner {
    Path memory path = s_chainSelectorToPath[remoteChainSelector];

    if (path.allowedCaller == bytes32(0)) {
      revert PathNotExist(remoteChainSelector);
    }

    delete s_chainSelectorToPath[remoteChainSelector];

    emit PathRemoved(remoteChainSelector, path.lChainId, path.allowedCaller);
  }

  // ================================================================
  // │                        Internal utils                        │
  // ================================================================

  function _getTokenDecimals(IERC20Metadata token, uint8 fallbackDecimals) internal view returns (uint8) {
    try token.decimals() returns (uint8 dec) {
      return dec;
    } catch {
      return fallbackDecimals;
    }
  }

  /// @notice Returns the Lombard-specific configuration for this pool.
  /// @return verifierResolver The address of the Lombard verifier resolver.
  /// @return bridge The address of the Lombard bridge contract.
  /// @return tokenAdapter The optional token adapter address (address(0) if not used).
  function getLombardConfig() external view returns (address verifierResolver, address bridge, address tokenAdapter) {
    return (i_lombardVerifierResolver, address(i_bridge), i_tokenAdapter);
  }
}
