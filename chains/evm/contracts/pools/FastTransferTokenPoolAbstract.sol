// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";

import {IFastTransferPool} from "../interfaces/IFastTransferPool.sol";
import {IRouterClient} from "../interfaces/IRouterClient.sol";
import {IWrappedNative} from "../interfaces/IWrappedNative.sol";

import {CCIPReceiver} from "../applications/CCIPReceiver.sol";
import {Client} from "../libraries/Client.sol";

import {IERC20} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";

/// @title Abstract Fast-Transfer Pool
/// @notice Base contract for fast-transfer pools that provides common functionality
/// for quoting, fill-tracking, and CCIP send helpers.
abstract contract FastTransferTokenPoolAbstract is CCIPReceiver, ITypeAndVersion, IFastTransferPool {
  error WhitelistNotEnabled();

  event LaneUpdated(
    uint64 indexed dst,
    uint16 bps,
    bool enabled,
    uint256 fillAmountMaxPerRequest,
    address destinationPool,
    address[] addFillers,
    address[] removeFillers
  );
  event FillerWhitelistUpdated(uint64 indexed dst, address[] addFillers, address[] removeFillers);
  event DestinationPoolUpdated(uint64 indexed dst, address destinationPool);
  event FastFillCompleted(bytes32 indexed fillRequestId);

  struct LaneConfig {
    uint16 bpsFastFee; // 0-10_000
    bool enabled; // pause per lane
    address destinationPool;
    uint256 fillAmountMaxPerRequest; // max amount that can be filled per request
    mapping(address filler => bool isWhitelisted) fillerWhitelist; // whitelist for fillers
  }

  struct MintMessage {
    uint256 srcAmountToTransfer;
    uint8 srcDecimals;
    uint256 fastTransferFee;
    bytes receiver;
  }

  // Storage layout
  mapping(uint64 destinationChainSelector => LaneConfig laneConfig) public s_fastTransferLaneConfig;
  mapping(bytes32 fillId => address filler) internal s_fills;
  bool public s_whitelistEnabled;
  address private s_wrappedNative;

  /// @notice Gets the remote pool address for a given chain selector
  /// @param chainSelector The chain selector
  /// @return address The remote pool address
  function _getRemotePool(
    uint64 chainSelector
  ) internal view returns (address) {
    return s_fastTransferLaneConfig[chainSelector].destinationPool;
  }

  /// @notice Sets the lane configuration
  /// @param dst The destination chain selector
  /// @param bps The fee basis points (0-10000)
  /// @param enabled Whether the lane is enabled
  /// @param fillAmountMaxPerRequest The maximum amount that can be filled per request
  /// @param addFillers The addresses to add to the whitelist
  /// @param removeFillers The addresses to remove from the whitelist
  function updateLaneConfig(
    uint64 dst,
    uint16 bps,
    bool enabled,
    address destinationPool,
    uint256 fillAmountMaxPerRequest,
    address[] memory addFillers,
    address[] memory removeFillers
  ) external virtual {
    _checkAdmin();
    if (bps > 10_000) revert InvalidLaneConfig();
    LaneConfig storage laneConfig = s_fastTransferLaneConfig[dst];
    laneConfig.destinationPool = destinationPool;
    laneConfig.bpsFastFee = bps;
    laneConfig.enabled = enabled;
    laneConfig.fillAmountMaxPerRequest = fillAmountMaxPerRequest;
    for (uint256 i; i < addFillers.length; ++i) {
      laneConfig.fillerWhitelist[addFillers[i]] = true;
    }
    for (uint256 i; i < removeFillers.length; ++i) {
      laneConfig.fillerWhitelist[removeFillers[i]] = false;
    }
    emit LaneUpdated(dst, bps, enabled, fillAmountMaxPerRequest, destinationPool, addFillers, removeFillers);
  }

  /// @notice Sets the filler whitelist configuration for a given lane
  /// @param dst The destination chain selector
  /// @param addFillers The addresses to add to the whitelist
  /// @param removeFillers The addresses to remove from the whitelist
  function updateFillerWhitelist(
    uint64 dst,
    address[] memory addFillers,
    address[] memory removeFillers
  ) external virtual {
    _checkAdmin();
    LaneConfig storage laneConfig = s_fastTransferLaneConfig[dst];
    for (uint256 i; i < addFillers.length; ++i) {
      laneConfig.fillerWhitelist[addFillers[i]] = true;
    }
    for (uint256 i; i < removeFillers.length; ++i) {
      laneConfig.fillerWhitelist[removeFillers[i]] = false;
    }
    emit FillerWhitelistUpdated(dst, addFillers, removeFillers);
  }

  /// @notice Gets the CCIP send token fee and fast transfer fee for a given transfer
  /// @param feeToken The token used to pay the CCIP fee
  /// @param dstChainSelector The destination chain selector
  /// @param amount The amount to transfer
  /// @param receiver The receiver address
  /// @param extraArgs Extra arguments for the transfer
  /// @return Quote containing the CCIP fee and fast transfer fee
  function getCcipSendTokenFee(
    address feeToken,
    uint64 dstChainSelector,
    uint256 amount,
    bytes calldata receiver,
    bytes calldata extraArgs
  ) public view virtual returns (Quote memory) {
    LaneConfig storage laneConfig = s_fastTransferLaneConfig[dstChainSelector];
    if (!laneConfig.enabled) revert InvalidLaneConfig();

    bool slow = extraArgs.length > 0 && (extraArgs[0] & bytes1(0x01)) == bytes1(0x01);
    uint256 fastFee = slow ? 0 : amount * laneConfig.bpsFastFee / 10_000;

    bytes memory data = abi.encode(MintMessage(amount, 18, fastFee, receiver));
    Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
      receiver: abi.encode(_getRemotePool(dstChainSelector)),
      data: data,
      tokenAmounts: new Client.EVMTokenAmount[](0),
      feeToken: feeToken,
      extraArgs: ""
    });
    // TODO: add extraArgs, mentioning the gas limit for settlement
    uint256 sendFee = IRouterClient(getRouter()).getFee(dstChainSelector, message);

    return Quote(sendFee, fastFee);
  }

  /// @notice Sends tokens via CCIP with optional fast transfer
  /// @param feeToken The token used to pay the CCIP fee
  /// @param destinationChainSelector The destination chain selector
  /// @param amount The amount to transfer
  /// @param receiver The receiver address
  /// @param extraArgs Extra arguments for the transfer
  /// @return fillRequestId The fill request ID
  function ccipSendToken(
    address feeToken,
    uint64 destinationChainSelector,
    uint256 amount,
    bytes calldata receiver,
    bytes calldata extraArgs
  ) external payable virtual returns (bytes32 fillRequestId) {
    LaneConfig storage laneConfig = s_fastTransferLaneConfig[destinationChainSelector];
    if (!laneConfig.enabled) revert InvalidLaneConfig();

    bool slow = extraArgs.length > 0 && (extraArgs[0] & bytes1(0x01)) == bytes1(0x01);
    uint256 fastFee = slow ? 0 : (amount * laneConfig.bpsFastFee) / 10_000;

    // Lock/burn tokens (actual logic will live in the TokenPool implementation)
    _handleTokenToTransfer(destinationChainSelector, msg.sender, amount + fastFee);

    // Get CCIP fee and transfer it
    bytes memory data = abi.encode(MintMessage(amount, 18, fastFee, receiver));
    // Prepare CCIP message
    Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
      receiver: abi.encode(_getRemotePool(destinationChainSelector)),
      data: data,
      tokenAmounts: new Client.EVMTokenAmount[](0),
      feeToken: feeToken,
      extraArgs: ""
    });
    // TODO: add extraArgs, mentioning the gas limit for settlement
    if (feeToken == address(0)) {
      feeToken = s_wrappedNative;
      uint256 feeTokenAmount = IRouterClient(i_ccipRouter).getFee(destinationChainSelector, message);
      if (msg.value < feeTokenAmount) revert IRouterClient.InsufficientFeeTokenAmount();
      feeTokenAmount = msg.value;
      IWrappedNative(feeToken).deposit{value: feeTokenAmount}();
    } else {
      if (msg.value > 0) revert IRouterClient.InvalidMsgValue();
      uint256 feeTokenAmount = IRouterClient(i_ccipRouter).getFee(destinationChainSelector, message);
      IERC20(feeToken).transferFrom(msg.sender, address(this), feeTokenAmount);
      IERC20(feeToken).approve(i_ccipRouter, feeTokenAmount);
    }

    fillRequestId = IRouterClient(getRouter()).ccipSend(destinationChainSelector, message);

    emit FastFillRequest(fillRequestId, destinationChainSelector, amount, fastFee, receiver);
    return fillRequestId;
  }

  /// @notice Fast fills a transfer using liquidity provider funds based on CCIP settlement
  /// @param fillRequestId The fill request ID
  /// @param srcAmount The amount to fill
  /// @param srcDecimals The decimals of the source token
  /// @param receiver The receiver address
  function fastFill(
    bytes32 fillRequestId,
    uint64 sourceChainSelector,
    uint256 srcAmount,
    uint8 srcDecimals,
    address receiver
  ) public virtual {
    bytes32 fillId = keccak256(abi.encodePacked(fillRequestId, srcAmount, receiver));
    address filler = s_fills[fillId];
    if (filler != address(0)) revert AlreadyFilled(fillRequestId);

    // Optional whitelist check
    if (s_whitelistEnabled) {
      LaneConfig storage laneConfig = s_fastTransferLaneConfig[0]; // Use appropriate chain selector
      if (!laneConfig.fillerWhitelist[msg.sender]) revert WhitelistNotEnabled();
    }

    // Transfer tokens from filler to receiver
    uint256 destAmount = _transferFromFiller(sourceChainSelector, msg.sender, receiver, srcAmount, srcDecimals);

    // Record fill
    s_fills[fillId] = msg.sender;
    emit FastFill(fillRequestId, fillId, msg.sender, destAmount, receiver);
  }

  // @inheritdoc CCIPReceiver
  function _ccipReceive(
    Client.Any2EVMMessage memory message
  ) internal override onlyRouter {
    // Decode message
    MintMessage memory mintMsg = abi.decode(message.data, (MintMessage));
    _settle(
      message.sourceChainSelector,
      message.messageId,
      message.sender,
      mintMsg.srcAmountToTransfer,
      mintMsg.srcDecimals,
      mintMsg.fastTransferFee,
      address(uint160(uint256(bytes32(mintMsg.receiver))))
    );
    emit FastFillCompleted(message.messageId);
  }

  /// @notice Handles the token to transfer on fast fill request at source chain
  /// @param destinationChainSelector The destination chain selector to which the fast fill request is sent
  /// @param sender The sender address
  /// @param amount The amount to transfer
  function _handleTokenToTransfer(uint64 destinationChainSelector, address sender, uint256 amount) internal virtual;

  /// @notice Transfers tokens from the filler to the receiver
  /// @param sourceChainSelector The source chain selector
  /// @param filler The address of the filler
  /// @param receiver The address of the receiver
  /// @param srcAmount The amount to transfer
  /// @param srcDecimals The decimals of the source token
  /// @return destAmount The amount transferred to the receiver on the destination chain
  function _transferFromFiller(
    uint64 sourceChainSelector,
    address filler,
    address receiver,
    uint256 srcAmount,
    uint8 srcDecimals
  ) internal virtual returns (uint256 destAmount);

  /// @notice Handles the settlement of a fast fill request at destination chain
  /// @param sourceChainSelector The source chain selector
  /// @param fillRequestId The fill request ID
  /// @param sourcePoolAddress The source pool address
  /// @param srcAmount The amount to transfer
  /// @param srcDecimal The decimals of the source token
  /// @param fastTransferFee The fast transfer fee
  /// @param receiver The receiver address
  function _settle(
    uint64 sourceChainSelector,
    bytes32 fillRequestId,
    bytes memory sourcePoolAddress,
    uint256 srcAmount,
    uint8 srcDecimal,
    uint256 fastTransferFee,
    address receiver
  ) internal virtual;

  /// @notice Override this function in your implementation.
  /// @dev The check is dependent on the ownership implementation of the implementation contract of the pool
  /// we do not enforce the ownership implementation in this abstract contract
  function _checkAdmin() internal view virtual;
}
