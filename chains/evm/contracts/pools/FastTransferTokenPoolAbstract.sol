// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

// Local interfaces

import {IAny2EVMMessageReceiver} from "../interfaces/IAny2EVMMessageReceiver.sol";
import {IFastTransferPool} from "../interfaces/IFastTransferPool.sol";
import {IRouterClient} from "../interfaces/IRouterClient.sol";

import {IRMN} from "../interfaces/IRMN.sol";
import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";

import {CCIPReceiver} from "../applications/CCIPReceiver.sol";
import {Client} from "../libraries/Client.sol";
import {TokenPool} from "./TokenPool.sol";

import {IERC20} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/utils/SafeERC20.sol";
import {IERC165} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v4.8.3/contracts/utils/introspection/IERC165.sol";

/// @title Abstract Fast-Transfer Pool
/// @notice Base contract for fast-transfer pools that provides common functionality
/// for quoting, fill-tracking, and CCIP send helpers.
abstract contract FastTransferTokenPoolAbstract is TokenPool, CCIPReceiver, ITypeAndVersion, IFastTransferPool {
  using SafeERC20 for IERC20;

  error WhitelistNotEnabled();
  error InvalidLaneConfig();
  error FillerNotWhitelisted(uint64 remoteChainSelector, address filler);

  event LaneUpdated(
    uint64 indexed dst,
    uint16 bps,
    uint256 fillAmountMaxPerRequest,
    bytes destinationPool,
    address[] addFillers,
    address[] removeFillers
  );
  event FillerAllowListUpdated(uint64 indexed dst, address[] addFillers, address[] removeFillers);
  event DestinationPoolUpdated(uint64 indexed dst, address destinationPool);

  struct LaneConfig {
    uint256 fillAmountMaxPerRequest; //    max amount that can be filled per request
    uint16 bpsFastFee; // ─────────────╮ 0-10_000
    bool fillerAllowlistEnabled; // ───╯ whitelist for fillers
    bytes destinationPool; // destination pool address
    mapping(address filler => bool isAllowed) fillerAllowList; // whitelist of fillers
  }

  struct LaneConfigView {
    uint256 fillAmountMaxPerRequest; //    max amount that can be filled per request
    uint16 bpsFastFee; // ────────────╮ 0-10_000
    bool fillerAllowlistEnabled; // ──╯ whitelist for fillers
    bytes destinationPool; // destination pool address
  }

  struct LaneConfigArgs {
    uint256 fillAmountMaxPerRequest; //    max amount that can be filled per request
    address[] addFillers; //               address allowed to fill
    address[] removeFillers; //            addresses to remove from the whitelist
    uint64 remoteChainSelector; // ──────╮
    uint16 bpsFastFee; //                │ 0-10_000
    bool fillerAllowlistEnabled; // ─────╯
    bytes destinationPool;
  }

  struct MintMessage {
    uint256 srcAmountToTransfer; // source amount from fill request
    uint256 fastTransferFee; // fast transfer fee in the source token
    bytes receiver; // receiver address on the destination chain
    uint8 srcDecimals; // decimals of the source token
  }

  /// @dev Mapping of remote chain selector to lane configuration
  mapping(uint64 remoteChainSelector => LaneConfig laneConfig) private s_fastTransferLaneConfig;

  /// @dev Mapping of fill request ID to filler address
  /// This is used to track which filler has filled a request
  /// @dev The filler address is set to address(0) if the request has not been filled
  /// @dev The filler address is set to address(1) if the request has been settled
  /// @dev The filler address is set to the filler address if the request has been filled by that filler
  mapping(bytes32 fillId => address filler) internal s_fills;

  /// @notice Initializes the fast transfer pool
  /// @param token The token this pool manages
  /// @param localTokenDecimals The decimals of the local token
  /// @param allowlist The allowlist of addresses
  /// @param rmnProxy The RMN proxy address
  /// @param router Address of the CCIP router
  constructor(
    IERC20 token,
    uint8 localTokenDecimals,
    address[] memory allowlist,
    address rmnProxy,
    address router
  ) TokenPool(token, localTokenDecimals, allowlist, rmnProxy, router) CCIPReceiver(router) {}

  /// @notice Updates the lane configuration
  /// @param laneConfigArgs The lane configuration arguments
  function updateLaneConfig(
    LaneConfigArgs calldata laneConfigArgs
  ) external virtual onlyOwner {
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
  ) external virtual onlyOwner {
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
    address settlementFeeToken,
    uint64 destinationChainSelector,
    uint256 amount,
    bytes calldata receiver,
    bytes calldata extraArgs
  ) public view virtual override returns (Quote memory) {
    (Quote memory quote,) = _getFeeQuoteAndCCIPMessage(settlementFeeToken, destinationChainSelector, amount, receiver, extraArgs);
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
    if (IRMN(i_rmnProxy).isCursed(bytes16(uint128(destinationChainSelector)))) revert CursedByRMN();
    _checkAllowList(msg.sender);
    if (!isSupportedChain(destinationChainSelector)) revert ChainNotAllowed(destinationChainSelector);
    _consumeOutboundRateLimit(destinationChainSelector, amount);
    _handleTokenToTransfer(destinationChainSelector, msg.sender, amount + quote.fastTransferFee);

    if (feeToken != address(0)) {
      IERC20(feeToken).safeTransferFrom(msg.sender, address(this), quote.ccipSettlementFee);
      IERC20(feeToken).safeApprove(i_ccipRouter, quote.ccipSettlementFee);
    }
    fillRequestId = IRouterClient(getRouter()).ccipSend{value: msg.value}(destinationChainSelector, message);
    emit FastFillRequest(fillRequestId, destinationChainSelector, amount, quote.fastTransferFee, receiver);
    return fillRequestId;
  }

  /// @notice Pulls out all of the fee‐quotation + message‐build logic
  function _getFeeQuoteAndCCIPMessage(
    address settlementFeeToken,
    uint64 destinationChainSelector,
    uint256 amount,
    bytes calldata receiver,
    bytes calldata
  ) internal view returns (IFastTransferPool.Quote memory quote, Client.EVM2AnyMessage memory message) {
    LaneConfig storage lane = s_fastTransferLaneConfig[destinationChainSelector];

    quote.fastTransferFee = (amount * lane.bpsFastFee) / 10_000;
    // pack the MintMessage
    bytes memory data = abi.encode(
      MintMessage({
        srcAmountToTransfer: amount,
        srcDecimals: i_tokenDecimals,
        fastTransferFee: quote.fastTransferFee,
        receiver: receiver
      })
    );

    message = Client.EVM2AnyMessage({
      receiver: abi.encode(lane.destinationPool),
      data: data,
      tokenAmounts: new Client.EVMTokenAmount[](0),
      feeToken: settlementFeeToken,
      extraArgs: ""
    });

    quote.ccipSettlementFee = IRouterClient(getRouter()).getFee(destinationChainSelector, message);
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
  ) internal virtual {
    _validateSettlement(sourceChainSelector, sourcePoolAddress);

    uint256 localAmount = _calculateLocalAmount(srcAmount, srcDecimal);
    uint256 settlementAmountLocal = localAmount + _calculateLocalAmount(fastTransferFee, srcDecimal);

    bytes32 fillId = keccak256(abi.encodePacked(fillRequestId, localAmount, receiver));
    address filler = s_fills[fillId];

    // Handle settlement based on fill state
    if (filler == address(0)) {
      // Not fast-filled - mint/release to receiver
      _handleNotFastFilled(sourceChainSelector, settlementAmountLocal, receiver);
    } else if (filler == address(1)) {
      // Already settled
      revert MessageAlreadySettled(fillRequestId);
    } else {
      // Fast-filled - reimburse filler
      _handleFastFilledReimbursement(filler, settlementAmountLocal);
    }

    // Mark as settled
    s_fills[fillId] = address(1);
  }

  /// @notice Validates settlement prerequisites
  /// @param sourceChainSelector The source chain selector
  /// @param sourcePoolAddress The source pool address
  function _validateSettlement(uint64 sourceChainSelector, bytes memory sourcePoolAddress) internal view virtual {
    if (IRMN(i_rmnProxy).isCursed(bytes16(uint128(sourceChainSelector)))) revert CursedByRMN();
    //Validates that the source pool address is configured on this pool.
    if (!isRemotePool(sourceChainSelector, sourcePoolAddress)) {
      revert InvalidSourcePoolAddress(sourcePoolAddress);
    }
  }

  /// @notice Handles settlement when the request was not fast-filled
  /// @param sourceChainSelector The source chain selector
  /// @param settlementAmountLocal The amount to settle in local token
  /// @param receiver The receiver address
  function _handleNotFastFilled(
    uint64 sourceChainSelector,
    uint256 settlementAmountLocal,
    address receiver
  ) internal virtual;

  /// @notice Handles reimbursement when the request was fast-filled
  /// @param filler The filler address to reimburse
  /// @param settlementAmountLocal The amount to reimburse in local token
  function _handleFastFilledReimbursement(address filler, uint256 settlementAmountLocal) internal virtual;

  /// @notice Override getRouter to resolve diamond inheritance
  function getRouter() public view virtual override(TokenPool, CCIPReceiver) returns (address) {
    return TokenPool.getRouter();
  }

  /// @notice Override supportsInterface to resolve diamond inheritance
  function supportsInterface(
    bytes4 interfaceId
  ) public pure virtual override(TokenPool, CCIPReceiver) returns (bool) {
    return interfaceId == type(IFastTransferPool).interfaceId || interfaceId == type(ITypeAndVersion).interfaceId
      || interfaceId == type(IAny2EVMMessageReceiver).interfaceId || interfaceId == type(IERC165).interfaceId
      || super.supportsInterface(interfaceId);
  }
}
