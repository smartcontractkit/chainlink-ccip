// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";

import {IFastTransferPool} from "../interfaces/IFastTransferPool.sol";
import {IRouterClient} from "../interfaces/IRouterClient.sol";
import {IWrappedNative} from "../interfaces/IWrappedNative.sol";

import {CCIPReceiver} from "../applications/CCIPReceiver.sol";
import {Client} from "../libraries/Client.sol";
import {Ownable2StepMsgSender} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2StepMsgSender.sol";

import {IERC20} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";

/// @title Abstract Fast-Transfer Pool
/// @notice Base contract for fast-transfer pools that provides common functionality
///         for quoting, fill-tracking, and CCIP send helpers.
abstract contract FastTransferTokenPoolAbstract is CCIPReceiver, ITypeAndVersion, IFastTransferPool {
  error WhitelistNotEnabled();
  error InvalidSourcePoolAddress(bytes sender);

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
    uint256 amountToTransfer;
    uint256 fastTransferFee;
    bytes receiver;
  }

  // Storage layout
  mapping(uint64 destinationChainSelector => LaneConfig laneConfig) public s_fastTransferLaneConfig;
  mapping(bytes32 fillId => address filler) private s_fills;
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

    bytes memory data = abi.encode(MintMessage(amount, fastFee, receiver));
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
  /// @return messageId The CCIP message ID
  function ccipSendToken(
    address feeToken,
    uint64 destinationChainSelector,
    uint256 amount,
    bytes calldata receiver,
    bytes calldata extraArgs
  ) external payable virtual returns (bytes32 messageId) {
    LaneConfig storage laneConfig = s_fastTransferLaneConfig[destinationChainSelector];
    if (!laneConfig.enabled) revert InvalidLaneConfig();

    bool slow = extraArgs.length > 0 && (extraArgs[0] & bytes1(0x01)) == bytes1(0x01);
    uint256 fastFee = slow ? 0 : (amount * laneConfig.bpsFastFee) / 10_000;

    // Lock/burn tokens (actual logic will live in the TokenPool implementation)
    _handleTokenToTransfer(msg.sender, amount + fastFee);

    // Get CCIP fee and transfer it
    bytes memory data = abi.encode(MintMessage(amount, fastFee, receiver));
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

    messageId = IRouterClient(getRouter()).ccipSend(destinationChainSelector, message);

    emit FastFillRequest(messageId, destinationChainSelector, amount, fastFee, receiver);
    return messageId;
  }

  /// @notice Fast fills a transfer using liquidity provider funds
  /// @param fillRequestId The fill request ID
  /// @param amount The amount to fill
  /// @param receiver The receiver address
  function fastFill(bytes32 fillRequestId, uint256 amount, address receiver) public virtual {
    bytes32 fillId = keccak256(abi.encodePacked(fillRequestId, amount, receiver));
    address filler = s_fills[fillId];
    if (filler != address(0)) revert AlreadyFilled(fillRequestId);

    // Optional whitelist check
    if (s_whitelistEnabled) {
      LaneConfig storage laneConfig = s_fastTransferLaneConfig[0]; // Use appropriate chain selector
      if (!laneConfig.fillerWhitelist[msg.sender]) revert WhitelistNotEnabled();
    }

    // Transfer tokens from filler to receiver
    _transferFromFiller(msg.sender, receiver, amount);

    // Record fill
    s_fills[fillId] = msg.sender;
    emit FastFill(fillRequestId, msg.sender, amount, receiver);
  }

  // @inheritdoc CCIPReceiver
  function _ccipReceive(
    Client.Any2EVMMessage memory message
  ) internal override onlyRouter {
    // Decode message
    MintMessage memory mintMsg = abi.decode(message.data, (MintMessage));
    bytes32 fillId = keccak256(abi.encodePacked(message.messageId, mintMsg.amountToTransfer, mintMsg.receiver));
    address filler = s_fills[fillId];

    // not fast-filled
    if (filler == address(0)) {
      _settle(message.sourceChainSelector, address(uint160(uint256(bytes32(mintMsg.receiver)))), mintMsg.amountToTransfer + mintMsg.fastTransferFee, true);
    }
    // already finalized
    else if (filler == address(1)) {
      revert MessageAlreadySettled(message.messageId);
    }
    // fast-filled; verify amount
    else {
      // Honest filler -> pay them back + fee
      _settle(message.sourceChainSelector, filler, filler, mintMsg.amountToTransfer + mintMsg.fastTransferFee, false);
    }
    // Mark completed
    s_fills[fillId] = address(1); //sentinel value

    emit FastFillCompleted(message.messageId);
  }

  /// @notice Handles the token to transfer on fast fill request at source chain
  /// @param destinationChainSelector The destination chain selector to which the fast fill request is sent
  /// @param sender The sender address
  /// @param amount The amount to transfer
  function _handleTokenToTransfer(uint64 destinationChainSelector, address sender, uint256 amount) internal virtual;

  /// @notice Transfers tokens from the filler to the receiver
  /// @param filler The address of the filler
  /// @param receiver The address of the receiver
  /// @param amount The amount to transfer
  function _transferFromFiller(address filler, address receiver, uint256 amount) internal virtual;

  /// @notice Handles the settlement of a fast fill request at destination chain
  /// @param sourceChainSelector The source chain of the fast fill request
  /// @param filler The filler of the fast fill request
  /// @param settlementReceiver The receiver of settlement on destination chain
  /// @param amount The amount to settle
  function _settle(uint64 sourceChainSelector, address settlementReceiver, uint256 amount, bool shouldConsumeRateLimit) internal virtual;

  /// @notice Override this function in your implementation.
  /// @dev The check is dependent on the ownership implementation of the pool, we do not enforce the ownership implementation here
  function _checkAdmin() internal view virtual;
}
