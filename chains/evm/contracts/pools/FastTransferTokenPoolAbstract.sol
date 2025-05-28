// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.10;

// Local interfaces
import {IFastTransferPool} from "../interfaces/IFastTransferPool.sol";
import {IRouterClient} from "../interfaces/IRouterClient.sol";

// Chainlink interfaces
import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";

// Local libraries and applications
import {CCIPReceiver} from "../applications/CCIPReceiver.sol";
import {Client} from "../libraries/Client.sol";

// OpenZeppelin imports
import {IERC20} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/utils/SafeERC20.sol";

/// @title Abstract Fast-Transfer Pool
/// @notice Base contract for fast-transfer pools that provides common functionality
/// for quoting, fill-tracking, and CCIP send helpers.
abstract contract FastTransferTokenPoolAbstract is CCIPReceiver, ITypeAndVersion, IFastTransferPool {
  using SafeERC20 for IERC20;

  error WhitelistNotEnabled();
  error InvalidLaneConfig();
  error FillerNotWhitelisted(uint64 remoteChainSelector, address filler);

  event LaneUpdated(
    uint64 indexed dst,
    uint16 bps,
    bool enabled,
    uint256 fillAmountMaxPerRequest,
    address destinationPool,
    address[] addFillers,
    address[] removeFillers
  );
  event FillerAllowListUpdated(uint64 indexed dst, address[] addFillers, address[] removeFillers);
  event DestinationPoolUpdated(uint64 indexed dst, address destinationPool);

  struct LaneConfig {
    uint256 fillAmountMaxPerRequest; //    max amount that can be filled per request
    address destinationPool; // ─────────╮
    uint16 bpsFastFee; //                │ 0-10_000
    bool enabled; //                     │ pause per lane
    bool fillerAllowlistEnabled; //──────╯ whitelist for fillers
    mapping(address filler => bool isAllowed) fillerAllowList; // whitelist of fillers
  }

  struct LaneConfigView {
    uint256 fillAmountMaxPerRequest; //    max amount that can be filled per request
    address destinationPool; // ─────────╮
    uint16 bpsFastFee; //                │ 0-10_000
    bool enabled; //                     │ pause per lane
    bool fillerAllowlistEnabled; //──────╯ whitelist for fillers
  }

  struct LaneConfigArgs {
    uint256 fillAmountMaxPerRequest; //    max amount that can be filled per request
    address[] addFillers; //               address allowed to fill
    address[] removeFillers; //            addresses to remove from the whitelist
    address destinationPool; // ─────────╮
    uint64 remoteChainSelector; //       │
    uint16 bpsFastFee; //                │ 0-10_000
    bool enabled; //                     │ pause per lane
    bool fillerAllowlistEnabled; //──────╯
  }

  struct MintMessage {
    uint256 srcAmountToTransfer; // source amount from fill request
    uint256 fastTransferFee; // fast transfer fee in the source token
    bytes receiver; // receiver address on the destination chain
    uint8 srcDecimals; // decimals of the source token
  }
  //

  /// @dev Mapping of remote chain selector to lane configuration
  mapping(uint64 remoteChainSelector => LaneConfig laneConfig) private s_fastTransferLaneConfig;

  /// @dev Mapping of fill request ID to filler address
  /// This is used to track which filler has filled a request
  /// @dev The filler address is set to address(0) if the request has not been filled
  /// @dev The filler address is set to address(1) if the request has been settled
  /// @dev The filler address is set to the filler address if the request has been filled by that filler
  mapping(bytes32 fillId => address filler) internal s_fills;

  /// @notice Initializes the fast transfer pool
  /// @param router Address of the CCIP router

  constructor(
    address router
  )
    //LaneConfigArgs[] memory laneConfigArgs
    CCIPReceiver(router)
  {}

  /// @notice Updates the lane configuration
  /// @param laneConfigArgs The lane configuration arguments
  function updateLaneConfig(
    LaneConfigArgs calldata laneConfigArgs
  ) external virtual {
    _checkAdmin();
    _updateLaneConfig(laneConfigArgs);
  }

  /// @notice Internal function to update the lane configuration
  /// @param laneConfigArgs The lane configuration arguments
  function _updateLaneConfig(
    LaneConfigArgs memory laneConfigArgs
  ) internal virtual {
    if (laneConfigArgs.bpsFastFee > 10_000) revert InvalidLaneConfig();
    LaneConfig storage laneConfig = s_fastTransferLaneConfig[laneConfigArgs.remoteChainSelector];
    laneConfig.destinationPool = laneConfigArgs.destinationPool;
    laneConfig.bpsFastFee = laneConfigArgs.bpsFastFee;
    laneConfig.enabled = laneConfigArgs.enabled;
    laneConfig.fillerAllowlistEnabled = laneConfigArgs.fillerAllowlistEnabled;
    laneConfig.fillAmountMaxPerRequest = laneConfigArgs.fillAmountMaxPerRequest;
    for (uint256 i; i < laneConfigArgs.addFillers.length; ++i) {
      laneConfig.fillerAllowList[laneConfigArgs.addFillers[i]] = true;
    }
    for (uint256 i; i < laneConfigArgs.removeFillers.length; ++i) {
      laneConfig.fillerAllowList[laneConfigArgs.removeFillers[i]] = false;
    }
    emit LaneUpdated(
      laneConfigArgs.remoteChainSelector,
      laneConfigArgs.bpsFastFee,
      laneConfigArgs.enabled,
      laneConfigArgs.fillAmountMaxPerRequest,
      laneConfigArgs.destinationPool,
      laneConfigArgs.addFillers,
      laneConfigArgs.removeFillers
    );
  }

  /// @notice Updates the filler whitelist configuration for a given lane
  /// @param destinationChainSelector The destination chain selector
  /// @param addFillers The addresses to add to the whitelist
  /// @param removeFillers The addresses to remove from the whitelist
  function updatefillerAllowList(
    uint64 destinationChainSelector,
    address[] memory addFillers,
    address[] memory removeFillers
  ) external virtual {
    _checkAdmin();
    LaneConfig storage laneConfig = s_fastTransferLaneConfig[destinationChainSelector];
    for (uint256 i; i < addFillers.length; ++i) {
      laneConfig.fillerAllowList[addFillers[i]] = true;
    }
    for (uint256 i; i < removeFillers.length; ++i) {
      laneConfig.fillerAllowList[removeFillers[i]] = false;
    }
    emit FillerAllowListUpdated(destinationChainSelector, addFillers, removeFillers);
  }

  /// @notice Gets the lane configuration for a given destination chain selector
  /// @param remoteChainSelector The remote chain selector
  /// @return laneConfig The lane configuration for the given destination chain selector
  function getLaneConfig(
    uint64 remoteChainSelector
  ) external view returns (LaneConfigView memory) {
    LaneConfig storage config = s_fastTransferLaneConfig[remoteChainSelector];
    return LaneConfigView({
      bpsFastFee: config.bpsFastFee,
      enabled: config.enabled,
      fillerAllowlistEnabled: config.fillerAllowlistEnabled,
      destinationPool: config.destinationPool,
      fillAmountMaxPerRequest: config.fillAmountMaxPerRequest
    });
  }

  /// @notice Checks if a filler is whitelisted for a given lane
  /// @param remoteChainSelector The remote chain selector
  /// @param filler The filler address to check
  /// @return isWhitelisted Whether the filler is whitelisted
  function isfillerAllowListed(uint64 remoteChainSelector, address filler) external view returns (bool) {
    return s_fastTransferLaneConfig[remoteChainSelector].fillerAllowList[filler];
  }

  /// @inheritdoc IFastTransferPool
  function getCcipSendTokenFee(
    address feeToken,
    uint64 destinationChainSelector,
    uint256 amount,
    bytes calldata receiver,
    bytes calldata extraArgs
  ) public view virtual override returns (Quote memory) {
    (Quote memory quote,) = _getFeeQuoteAndCCIPMessage(feeToken, destinationChainSelector, amount, receiver, extraArgs);
    return quote;
  }

  /// @inheritdoc IFastTransferPool
  function ccipSendToken(
    address feeToken,
    uint64 destinationChainSelector,
    uint256 amount,
    bytes calldata receiver,
    bytes calldata extraArgs
  ) external payable virtual override returns (bytes32 fillRequestId) {
    // burn/lock tokens + pay fastFee (in _handleTokenToTransfer)
    (Quote memory quote, Client.EVM2AnyMessage memory message) =
      _getFeeQuoteAndCCIPMessage(feeToken, destinationChainSelector, amount, receiver, extraArgs);
    _handleTokenToTransfer(destinationChainSelector, msg.sender, amount + quote.fastTransferFee);

    if (feeToken != address(0)) {
      IERC20(feeToken).safeTransferFrom(msg.sender, address(this), quote.sendTokenFee);
      IERC20(feeToken).safeApprove(i_ccipRouter, quote.sendTokenFee);
    }
    fillRequestId = IRouterClient(getRouter()).ccipSend{value: msg.value}(destinationChainSelector, message);
    emit FastFillRequest(fillRequestId, destinationChainSelector, amount, quote.fastTransferFee, receiver);
    return fillRequestId;
  }

  /// @notice Pulls out all of the fee‐quotation + message‐build logic
  function _getFeeQuoteAndCCIPMessage(
    address feeToken,
    uint64 destinationChainSelector,
    uint256 amount,
    bytes calldata receiver,
    bytes calldata extraArgs
  ) internal view returns (IFastTransferPool.Quote memory quote, Client.EVM2AnyMessage memory message) {
    LaneConfig storage lane = s_fastTransferLaneConfig[destinationChainSelector];
    if (!lane.enabled) revert IFastTransferPool.LaneDisabled();

    // compute fastFee
    // bool slow = extraArgs.length > 0 && (extraArgs[0] & bytes1(0x01)) == bytes1(0x01);
    // quote.fastTransferFee = slow ? 0 : (amount * lane.bpsFastFee) / 10_000;
    quote.fastTransferFee = (amount * lane.bpsFastFee) / 10_000;
    // pack the MintMessage
    bytes memory data = abi.encode(
      MintMessage({
        srcAmountToTransfer: amount,
        srcDecimals: 18,
        fastTransferFee: quote.fastTransferFee,
        receiver: receiver
      })
    );

    message = Client.EVM2AnyMessage({
      receiver: abi.encode(lane.destinationPool),
      data: data,
      tokenAmounts: new Client.EVMTokenAmount[](0),
      feeToken: feeToken,
      extraArgs: ""
    });

    quote.sendTokenFee = IRouterClient(getRouter()).getFee(destinationChainSelector, message);
    return (quote, message);
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

    {
      LaneConfig storage laneConfig = s_fastTransferLaneConfig[sourceChainSelector];
      if (laneConfig.fillerAllowlistEnabled) {
        if (!laneConfig.fillerAllowList[msg.sender]) revert FillerNotWhitelisted(sourceChainSelector, msg.sender);
      }
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
    // Decode message data directly into variables
    MintMessage memory mintMessage = abi.decode(message.data, (MintMessage));
    {
      _settle(
        message.sourceChainSelector,
        message.messageId,
        message.sender,
        mintMessage.srcAmountToTransfer,
        mintMessage.srcDecimals,
        mintMessage.fastTransferFee,
        address(uint160(uint256(bytes32(mintMessage.receiver))))
      );
    }
    emit FastFillSettled(message.messageId);
  }

  /// @notice Handles the token to transfer on fast fill request at source chain
  /// @param destinationChainSelector The destination chain selector
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
