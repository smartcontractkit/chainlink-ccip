// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IAdvancedPoolHooks} from "../../interfaces/IAdvancedPoolHooks.sol";
import {IGetPoolConfigManager} from "../../interfaces/IGetPoolConfigManager.sol";
import {IMultiTokenPool} from "../../interfaces/IMultiTokenPool.sol";
import {IPoolV1} from "../../interfaces/IPool.sol";
import {IPoolConfigManager} from "../../interfaces/IPoolConfigManager.sol";
import {IPoolV2} from "../../interfaces/IPoolV2.sol";
import {IRMN} from "../../interfaces/IRMN.sol";
import {IRouter} from "../../interfaces/IRouter.sol";

import {FeeTokenHandler} from "../../libraries/FeeTokenHandler.sol";
import {FinalityCodec} from "../../libraries/FinalityCodec.sol";
import {Pool} from "../../libraries/Pool.sol";
import {Ownable2StepMsgSender} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2StepMsgSender.sol";

import {IERC165} from "@openzeppelin/contracts@5.3.0/utils/introspection/IERC165.sol";

/// @notice Base contract for CCIP pools that serve many tokens. All token and lane config lives in the pool config
/// manager, which the pool calls through IPoolConfigManager.
/// @dev The pool runs the checks that a pool config manager cannot skip: the RMN curse, the ramp caller and the lock or
/// burn flavor. The pool config manager resolves the lane, builds and checks the pool data, including the source token
/// and the token identity metadata, charges the token fee and consumes the rate limits.
abstract contract MultiTokenPool is IMultiTokenPool, Ownable2StepMsgSender {
  error ZeroAddressInvalid();
  error CursedByRMN();
  error CallerIsNotARampOnRouter(address caller);
  error CallerIsNotPoolConfigManager(address caller);
  error UnsupportedLockOrBurnFlavor(address token, IPoolConfigManager.LockOrBurnFlavor flavor);

  event LockedOrBurned(uint64 indexed remoteChainSelector, address token, address sender, uint256 amount);
  event ReleasedOrMinted(
    uint64 indexed remoteChainSelector, address token, address sender, address recipient, uint256 amount
  );
  event RouterSet(address oldRouter, address newRouter);
  event PoolConfigManagerSet(address oldPoolConfigManager, address newPoolConfigManager);

  /// @dev The address of the RMN proxy.
  address internal immutable i_rmnProxy;

  /// @dev The address of the router.
  IRouter internal s_router;
  /// @dev The pool config manager that holds the config of every token of this pool.
  IPoolConfigManager internal s_poolConfigManager;

  constructor(
    address rmnProxy,
    address router,
    address poolConfigManager
  ) {
    if (rmnProxy == address(0) || router == address(0) || poolConfigManager == address(0)) {
      revert ZeroAddressInvalid();
    }
    i_rmnProxy = rmnProxy;
    s_router = IRouter(router);
    s_poolConfigManager = IPoolConfigManager(poolConfigManager);

    emit RouterSet(address(0), router);
    emit PoolConfigManagerSet(address(0), poolConfigManager);
  }

  // ================================================================
  // │                         Config                               │
  // ================================================================

  /// @notice Gets the RMN proxy address.
  function getRmnProxy() public view virtual returns (address rmnProxy) {
    return i_rmnProxy;
  }

  /// @notice Gets the router address.
  function getRouter() public view virtual returns (address router) {
    return address(s_router);
  }

  /// @inheritdoc IGetPoolConfigManager
  function getPoolConfigManager() public view virtual returns (address poolConfigManager) {
    return address(s_poolConfigManager);
  }

  /// @notice Sets the router.
  /// @param router The new router.
  function setRouter(
    address router
  ) external virtual onlyOwner {
    if (router == address(0)) revert ZeroAddressInvalid();

    emit RouterSet(address(s_router), router);
    s_router = IRouter(router);
  }

  /// @notice Sets the pool config manager. The new pool config manager applies to the next transfer.
  /// @param poolConfigManager The new pool config manager.
  function setPoolConfigManager(
    address poolConfigManager
  ) external virtual onlyOwner {
    if (poolConfigManager == address(0)) revert ZeroAddressInvalid();

    emit PoolConfigManagerSet(address(s_poolConfigManager), poolConfigManager);
    s_poolConfigManager = IPoolConfigManager(poolConfigManager);
  }

  /// @notice Signals which version of the pool interface is supported.
  /// @param interfaceId The interface identifier, as specified in ERC-165.
  function supportsInterface(
    bytes4 interfaceId
  ) public view virtual override returns (bool) {
    return interfaceId == Pool.CCIP_POOL_V1 || interfaceId == type(IPoolV2).interfaceId
      || interfaceId == type(IPoolV1).interfaceId || interfaceId == type(IMultiTokenPool).interfaceId
      || interfaceId == type(IGetPoolConfigManager).interfaceId || interfaceId == type(IERC165).interfaceId;
  }

  // ================================================================
  // │                        Lock or Burn                          │
  // ================================================================

  /// @inheritdoc IPoolV2
  function lockOrBurn(
    Pool.LockOrBurnInV1 calldata lockOrBurnIn,
    bytes4 requestedFinalityConfig,
    bytes calldata tokenArgs
  ) public virtual returns (Pool.LockOrBurnOutV1 memory lockOrBurnOut, uint256 destTokenAmount) {
    return _lockOrBurnWithConfig(lockOrBurnIn, requestedFinalityConfig, tokenArgs, true);
  }

  /// @inheritdoc IPoolV1
  /// @dev Uses default finality and charges no token fee.
  function lockOrBurn(
    Pool.LockOrBurnInV1 calldata lockOrBurnIn
  ) public virtual returns (Pool.LockOrBurnOutV1 memory lockOrBurnOut) {
    (lockOrBurnOut,) = _lockOrBurnWithConfig(lockOrBurnIn, FinalityCodec.WAIT_FOR_FINALITY_FLAG, "", false);
    return lockOrBurnOut;
  }

  /// @notice Runs the outbound checks, then locks or burns the amount after the token fee.
  /// @param lockOrBurnIn The original lock or burn request.
  /// @param requestedFinalityConfig The requested finality, encoded with FinalityCodec.
  /// @param tokenArgs Additional token arguments, passed to the hook.
  /// @param chargeTokenFee False for the V1 path, which charges no token fee.
  /// @return lockOrBurnOut The remote token and the pool data.
  /// @return destTokenAmount The amount after the token fee.
  function _lockOrBurnWithConfig(
    Pool.LockOrBurnInV1 calldata lockOrBurnIn,
    bytes4 requestedFinalityConfig,
    bytes memory tokenArgs,
    bool chargeTokenFee
  ) internal virtual returns (Pool.LockOrBurnOutV1 memory lockOrBurnOut, uint256 destTokenAmount) {
    _checkCurse(lockOrBurnIn.remoteChainSelector);
    _onlyOnRamp(lockOrBurnIn.remoteChainSelector);

    IPoolConfigManager.OutboundLane memory lane = s_poolConfigManager.validateLockOrBurn(
      lockOrBurnIn.localToken,
      lockOrBurnIn.remoteChainSelector,
      lockOrBurnIn.amount,
      requestedFinalityConfig,
      chargeTokenFee
    );
    _checkFlavor(lockOrBurnIn.localToken, lane.lockOrBurnFlavor);

    destTokenAmount = lockOrBurnIn.amount - lane.feeAmount;

    if (lane.hook != address(0)) {
      IAdvancedPoolHooks(lane.hook).preflightCheck(lockOrBurnIn, requestedFinalityConfig, tokenArgs, destTokenAmount);
    }

    _lockOrBurn(
      lockOrBurnIn.localToken,
      lockOrBurnIn.remoteChainSelector,
      destTokenAmount,
      lane.lockOrBurnFlavor,
      lane.lockOrBurnAddress
    );

    emit LockedOrBurned({
      remoteChainSelector: lockOrBurnIn.remoteChainSelector,
      token: lockOrBurnIn.localToken,
      sender: msg.sender,
      amount: destTokenAmount
    });

    return
      (Pool.LockOrBurnOutV1({destTokenAddress: lane.remoteToken, destPoolData: lane.destPoolData}), destTokenAmount);
  }

  /// @notice Moves the tokens out with the flavor of the token.
  /// @param token The local token.
  /// @param remoteChainSelector The destination chain.
  /// @param amount The amount after the token fee.
  /// @param flavor The lock or burn flavor of the token. The caller checked that the pool serves it.
  /// @param lockOrBurnAddress The BurnToAddress target or the lock box, zero otherwise.
  function _lockOrBurn(
    address token,
    uint64 remoteChainSelector,
    uint256 amount,
    IPoolConfigManager.LockOrBurnFlavor flavor,
    address lockOrBurnAddress
  ) internal virtual;

  // ================================================================
  // │                      Release or Mint                         │
  // ================================================================

  /// @inheritdoc IPoolV2
  function releaseOrMint(
    Pool.ReleaseOrMintInV1 calldata releaseOrMintIn,
    bytes4 requestedFinalityConfig
  ) public virtual returns (Pool.ReleaseOrMintOutV1 memory) {
    _checkCurse(releaseOrMintIn.remoteChainSelector);
    _onlyOffRamp(releaseOrMintIn.remoteChainSelector);

    IPoolConfigManager.ResolvedInboundTransfer memory transfer =
      s_poolConfigManager.validateReleaseOrMint(releaseOrMintIn, requestedFinalityConfig);
    _checkFlavor(releaseOrMintIn.localToken, transfer.lockOrBurnFlavor);

    if (transfer.hook != address(0)) {
      IAdvancedPoolHooks(transfer.hook).postflightCheck(releaseOrMintIn, transfer.localAmount, requestedFinalityConfig);
    }

    _releaseOrMint(
      releaseOrMintIn.localToken,
      releaseOrMintIn.receiver,
      transfer.localAmount,
      releaseOrMintIn.remoteChainSelector,
      transfer.lockOrBurnFlavor,
      transfer.lockOrBurnAddress
    );

    emit ReleasedOrMinted({
      remoteChainSelector: releaseOrMintIn.remoteChainSelector,
      token: releaseOrMintIn.localToken,
      sender: msg.sender,
      recipient: releaseOrMintIn.receiver,
      amount: transfer.localAmount
    });

    return Pool.ReleaseOrMintOutV1({destinationAmount: transfer.localAmount});
  }

  /// @inheritdoc IPoolV1
  /// @dev Uses default finality.
  function releaseOrMint(
    Pool.ReleaseOrMintInV1 calldata releaseOrMintIn
  ) public virtual returns (Pool.ReleaseOrMintOutV1 memory) {
    return releaseOrMint(releaseOrMintIn, FinalityCodec.WAIT_FOR_FINALITY_FLAG);
  }

  /// @notice Moves the tokens in with the flavor of the token.
  /// @param token The local token.
  /// @param receiver The receiver of the tokens.
  /// @param amount The amount in local decimals.
  /// @param remoteChainSelector The source chain.
  /// @param flavor The lock or burn flavor of the token. The caller checked that the pool serves it.
  /// @param lockOrBurnAddress The BurnToAddress target or the lock box, zero otherwise.
  function _releaseOrMint(
    address token,
    address receiver,
    uint256 amount,
    uint64 remoteChainSelector,
    IPoolConfigManager.LockOrBurnFlavor flavor,
    address lockOrBurnAddress
  ) internal virtual;

  // ================================================================
  // │                         Validation                           │
  // ================================================================

  /// @notice Returns whether the pool serves a lock or burn flavor.
  /// @param flavor The lock or burn flavor.
  function _isSupportedFlavor(
    IPoolConfigManager.LockOrBurnFlavor flavor
  ) internal pure virtual returns (bool);

  /// @notice Reverts if the pool does not serve the flavor of the token.
  /// @param token The local token.
  /// @param flavor The lock or burn flavor of the token.
  function _checkFlavor(
    address token,
    IPoolConfigManager.LockOrBurnFlavor flavor
  ) internal pure virtual {
    if (!_isSupportedFlavor(flavor)) revert UnsupportedLockOrBurnFlavor(token, flavor);
  }

  /// @notice Reverts if the remote chain is cursed.
  /// @param remoteChainSelector The remote chain.
  function _checkCurse(
    uint64 remoteChainSelector
  ) internal view virtual {
    if (IRMN(i_rmnProxy).isCursed(bytes16(uint128(remoteChainSelector)))) revert CursedByRMN();
  }

  /// @notice Reverts unless the caller is the OnRamp for the remote chain on the router.
  /// @dev The pool does not check the lane. The pool config manager does, so a token without a lane reverts there.
  /// @param remoteChainSelector The destination chain.
  function _onlyOnRamp(
    uint64 remoteChainSelector
  ) internal view virtual {
    if (msg.sender != s_router.getOnRamp(remoteChainSelector)) revert CallerIsNotARampOnRouter(msg.sender);
  }

  /// @notice Reverts unless the caller is an OffRamp for the remote chain on the router.
  /// @param remoteChainSelector The source chain.
  function _onlyOffRamp(
    uint64 remoteChainSelector
  ) internal view virtual {
    if (!s_router.isOffRamp(remoteChainSelector, msg.sender)) revert CallerIsNotARampOnRouter(msg.sender);
  }

  // ================================================================
  // │                           Views                              │
  // ================================================================

  /// @inheritdoc IPoolV1
  /// @dev True for a registered token with a flavor that the pool serves. Ignores the pause and the lanes.
  function isSupportedToken(
    address token
  ) public view virtual returns (bool) {
    IPoolConfigManager.TokenConfig memory config = _getTokenConfig(token);
    return config.isRegistered && _isSupportedFlavor(config.lockOrBurnFlavor);
  }

  /// @inheritdoc IMultiTokenPool
  function isSupportedChain(
    address localToken,
    uint64 remoteChainSelector
  ) public view virtual returns (bool supported) {
    return isSupportedToken(localToken) && _getLaneConnection(localToken, remoteChainSelector).hasConnection;
  }

  /// @inheritdoc IPoolV1
  /// @dev Always reverts. A multi-token pool has no lane without a local token.
  function isSupportedChain(
    uint64
  ) public pure virtual returns (bool) {
    revert LocalTokenRequired();
  }

  /// @inheritdoc IMultiTokenPool
  function getRemoteToken(
    address localToken,
    uint64 remoteChainSelector
  ) public view virtual returns (bytes memory remoteToken) {
    return _getLaneConnection(localToken, remoteChainSelector).remoteToken;
  }

  /// @inheritdoc IPoolV2
  /// @dev Always reverts. A multi-token pool has no remote token without a local token.
  function getRemoteToken(
    uint64
  ) public pure virtual returns (bytes memory) {
    revert LocalTokenRequired();
  }

  /// @inheritdoc IMultiTokenPool
  function getTokenDecimals(
    address localToken
  ) public view virtual returns (uint8 decimals) {
    return _getTokenConfig(localToken).localDecimals;
  }

  /// @inheritdoc IPoolV2
  /// @dev The CCV lists do not depend on the amount, the finality or the extra data.
  function getRequiredCCVs(
    address localToken,
    uint64 remoteChainSelector,
    uint256, // sourceAmount
    bytes4, // requestedFinalityConfig
    bytes calldata, // extraData
    IPoolV2.MessageDirection direction
  ) public view virtual returns (address[] memory requiredCCVs) {
    return s_poolConfigManager.getRequiredCCVs(localToken, remoteChainSelector, direction);
  }

  /// @inheritdoc IPoolV2
  function getTokenTransferFeeConfig(
    address localToken,
    uint64 destChainSelector,
    bytes4, // requestedFinalityConfig
    bytes calldata // tokenArgs
  ) public view virtual returns (TokenTransferFeeConfig memory feeConfig) {
    IPoolConfigManager.LaneKey[] memory lanes = new IPoolConfigManager.LaneKey[](1);
    lanes[0] = IPoolConfigManager.LaneKey({localToken: localToken, remoteChainSelector: destChainSelector});
    return s_poolConfigManager.getTokenTransferFeeConfigBatch(lanes)[0];
  }

  /// @inheritdoc IPoolV2
  /// @dev Reverts if the token profile does not allow the requested finality.
  function getFee(
    address localToken,
    uint64 destChainSelector,
    uint256, // amount
    address, // feeToken
    bytes4 requestedFinalityConfig,
    bytes calldata // tokenArgs
  )
    public
    view
    virtual
    returns (uint256 feeUSDCents, uint32 destGasOverhead, uint32 destBytesOverhead, uint16 tokenFeeBps, bool isEnabled)
  {
    IPoolConfigManager.LaneFee memory fee =
      s_poolConfigManager.getFee(localToken, destChainSelector, requestedFinalityConfig);
    return (fee.feeUSDCents, fee.destGasOverhead, fee.destBytesOverhead, fee.tokenFeeBps, fee.isEnabled);
  }

  /// @notice Returns the resolved config of one token.
  function _getTokenConfig(
    address token
  ) internal view virtual returns (IPoolConfigManager.TokenConfig memory) {
    address[] memory tokens = new address[](1);
    tokens[0] = token;
    return s_poolConfigManager.getTokenConfigBatch(tokens)[0];
  }

  /// @notice Returns the resolved connection of one lane.
  function _getLaneConnection(
    address localToken,
    uint64 remoteChainSelector
  ) internal view virtual returns (IPoolConfigManager.LaneConnection memory) {
    IPoolConfigManager.LaneKey[] memory lanes = new IPoolConfigManager.LaneKey[](1);
    lanes[0] = IPoolConfigManager.LaneKey({localToken: localToken, remoteChainSelector: remoteChainSelector});
    return s_poolConfigManager.getLaneConnectionBatch(lanes)[0];
  }

  // ================================================================
  // │                            Fees                              │
  // ================================================================

  /// @inheritdoc IMultiTokenPool
  /// @dev The pool balance of a token holds only fees, because a send locks or burns the amount after the fee.
  function withdrawFeeTokens(
    address[] calldata feeTokens,
    address recipient
  ) external virtual {
    if (msg.sender != address(s_poolConfigManager)) revert CallerIsNotPoolConfigManager(msg.sender);
    FeeTokenHandler._withdrawFeeTokens(feeTokens, recipient);
  }
}
