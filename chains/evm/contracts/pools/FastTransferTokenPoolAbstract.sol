// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IAny2EVMMessageReceiver} from "../interfaces/IAny2EVMMessageReceiver.sol";
import {IFastTransferPool} from "../interfaces/IFastTransferPool.sol";
import {IRMN} from "../interfaces/IRMN.sol";
import {IRouterClient} from "../interfaces/IRouterClient.sol";
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
import {EnumerableSet} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v4.8.3/contracts/utils/structs/EnumerableSet.sol";

/// @title Abstract Fast-Transfer Pool
/// @notice Base contract for fast-transfer pools that provides common functionality
/// for quoting, fill-tracking, and CCIP send helpers.
abstract contract FastTransferTokenPoolAbstract is TokenPool, CCIPReceiver, ITypeAndVersion, IFastTransferPool {
  using EnumerableSet for EnumerableSet.AddressSet;
  using SafeERC20 for IERC20;

  error WhitelistNotEnabled();
  error InvalidDestChainConfig();
  error FillerNotAllowlisted(uint64 remoteChainSelector, address filler);
  error TransferAmountExceedsMaxFillAmount(uint64 remoteChainSelector, uint256 amount);

  event LaneUpdated(
    uint64 indexed destinationChainSelector,
    uint16 bps,
    uint256 maxFillAmountPerRequest,
    bytes destinationPool,
    address[] addFillers,
    address[] removeFillers
  );
  event FillerAllowListUpdated(uint64 indexed dst, address[] addFillers, address[] removeFillers);
  event DestinationPoolUpdated(uint64 indexed dst, address destinationPool);

  struct DestChainConfig {
    uint256 maxFillAmountPerRequest; // Max amount that can be filled per request.
    bytes destinationPool; //           Destination pool address.
    uint16 fastTransferBpsFee; // ────╮ Allowed range of [0-10_000].
    bool fillerAllowlistEnabled; // ──╯ Allowlist for fillers.
    EnumerableSet.AddressSet fillerAllowList; // Enumerable set of allowed fillers.
  }

  struct DestChainConfigView {
    uint256 maxFillAmountPerRequest; // Max amount that can be filled per request.
    uint16 fastTransferBpsFee; // ────╮ Allowed range of [0-10_000].
    bool fillerAllowlistEnabled; // ──╯ Allowlist for fillers.
    bytes destinationPool; //           Address of the destination pool. ABI encoded in the case of an EVM pool.
  }

  struct DestChainConfigUpdateArgs {
    uint256 maxFillAmountPerRequest; // Maximum amount that can be filled per request.
    bytes destinationPool; //           Address of the destination pool. ABI encoded in the case of an EVM pool.
    uint64 remoteChainSelector; // ───╮ Remote chain selector.
    uint16 fastTransferBpsFee; //     | Allowed range of [0-10_000].
    bool fillerAllowlistEnabled; // ──╯ True is filler allowlist is enabled.
    address[] addFillers; //            Address allowed to fill.
    address[] removeFillers; //         Addresses to remove from the allowlist.
  }

  struct MintMessage {
    uint256 sourceAmountToTransfer; // Amount to fill in the source token denomination.
    uint256 fastTransferFee; // Fast transfer fee in the source token.
    bytes receiver; // Receiver address on the destination chain. ABI encoded in the case of an EVM address.
    uint8 sourceDecimals; // Decimals of the source token.
  }

  /// @notice Enum representing the state of a fill request.
  enum FillState {
    NOT_FILLED, // Request has not been filled yet.
    FILLED, // Request has been filled by a filler.
    SETTLED // Request has been settled via CCIP.

  }

  /// @notice Struct to track fill request information
  struct FillInfo {
    FillState state; // Current state of the fill request
    address filler; // Address of the filler (only relevant when state is Filled)
  }

  /// @dev Mapping of remote chain selector to lane configuration
  mapping(uint64 remoteChainSelector => DestChainConfig laneConfig) private s_fastTransferDestChainConfig;

  /// @dev Mapping of fill request ID to fill information
  /// This is used to track the state and filler of each fill request
  mapping(bytes32 fillId => FillInfo fillInfo) internal s_fills;

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
  function updateDestChainConfig(
    DestChainConfigUpdateArgs calldata laneConfigArgs
  ) external virtual onlyOwner {
    _updateDestChainConfig(laneConfigArgs);
  }

  /// @notice Internal function to update the lane configuration
  /// @param laneConfigArgs The lane configuration arguments
  function _updateDestChainConfig(
    DestChainConfigUpdateArgs memory laneConfigArgs
  ) internal virtual {
    if (laneConfigArgs.fastTransferBpsFee > 10_000) revert InvalidDestChainConfig();
    DestChainConfig storage laneConfig = s_fastTransferDestChainConfig[laneConfigArgs.remoteChainSelector];
    laneConfig.destinationPool = laneConfigArgs.destinationPool;
    laneConfig.fastTransferBpsFee = laneConfigArgs.fastTransferBpsFee;
    laneConfig.fillerAllowlistEnabled = laneConfigArgs.fillerAllowlistEnabled;
    laneConfig.maxFillAmountPerRequest = laneConfigArgs.maxFillAmountPerRequest;

    for (uint256 i = 0; i < laneConfigArgs.removeFillers.length; ++i) {
      laneConfig.fillerAllowList.remove(laneConfigArgs.removeFillers[i]);
    }
    for (uint256 i = 0; i < laneConfigArgs.addFillers.length; ++i) {
      laneConfig.fillerAllowList.add(laneConfigArgs.addFillers[i]);
    }
    emit LaneUpdated(
      laneConfigArgs.remoteChainSelector,
      laneConfigArgs.fastTransferBpsFee,
      laneConfigArgs.maxFillAmountPerRequest,
      laneConfigArgs.destinationPool,
      laneConfigArgs.addFillers,
      laneConfigArgs.removeFillers
    );
  }

  /// @notice Updates the filler allowlist configuration for a given lane.
  /// @param destinationChainSelector The destination chain selector.
  /// @param addFillers The addresses to add to the allowlist.
  /// @param removeFillers The addresses to remove from the allowlist.
  function updateFillerAllowList(
    uint64 destinationChainSelector,
    address[] memory addFillers,
    address[] memory removeFillers
  ) external virtual onlyOwner {
    DestChainConfig storage laneConfig = s_fastTransferDestChainConfig[destinationChainSelector];
    for (uint256 i = 0; i < addFillers.length; ++i) {
      laneConfig.fillerAllowList.add(addFillers[i]);
    }
    for (uint256 i = 0; i < removeFillers.length; ++i) {
      laneConfig.fillerAllowList.remove(removeFillers[i]);
    }
    emit FillerAllowListUpdated(destinationChainSelector, addFillers, removeFillers);
  }

  /// @notice Gets the lane configuration for a given destination chain selector
  /// @param remoteChainSelector The remote chain selector
  /// @return laneConfig The lane configuration for the given destination chain selector
  function getDestChainConfig(
    uint64 remoteChainSelector
  ) external view returns (DestChainConfigView memory) {
    DestChainConfig storage config = s_fastTransferDestChainConfig[remoteChainSelector];
    return DestChainConfigView({
      fastTransferBpsFee: config.fastTransferBpsFee,
      fillerAllowlistEnabled: config.fillerAllowlistEnabled,
      destinationPool: config.destinationPool,
      maxFillAmountPerRequest: config.maxFillAmountPerRequest
    });
  }

  /// @notice Checks if a filler is whitelisted for a given lane
  /// @param remoteChainSelector The remote chain selector
  /// @param filler The filler address to check
  /// @return isWhitelisted Whether the filler is whitelisted
  function isFillerAllowListed(uint64 remoteChainSelector, address filler) external view returns (bool) {
    return s_fastTransferDestChainConfig[remoteChainSelector].fillerAllowList.contains(filler);
  }

  /// @notice Gets all allowlisted fillers for a given destination chain
  /// @param remoteChainSelector The remote chain selector
  /// @return fillers Array of allowlisted filler addresses
  function getAllowListedFillers(
    uint64 remoteChainSelector
  ) external view returns (address[] memory) {
    return s_fastTransferDestChainConfig[remoteChainSelector].fillerAllowList.values();
  }

  /// @notice Gets the fill information for a given fill ID
  /// @param fillId The fill ID to query
  /// @return fillInfo The fill information including state and filler address
  function getFillInfo(
    bytes32 fillId
  ) external view returns (FillInfo memory) {
    return s_fills[fillId];
  }

  /// @notice Helper function to generate fill ID from request parameters
  /// @param fillRequestId The original fill request ID
  /// @param amount The amount being filled
  /// @param receiver The receiver address
  /// @return fillId The computed fill ID
  function computeFillId(
    bytes32 fillRequestId,
    uint256 amount,
    address receiver
  ) public pure override returns (bytes32) {
    return _computeFillId(fillRequestId, amount, receiver);
  }

  /// @inheritdoc IFastTransferPool
  function getCcipSendTokenFee(
    address settlementFeeToken,
    uint64 destinationChainSelector,
    uint256 amount,
    bytes calldata receiver,
    bytes calldata extraArgs
  ) public view virtual override returns (Quote memory) {
    (Quote memory quote,) =
      _getFeeQuoteAndCCIPMessage(settlementFeeToken, destinationChainSelector, amount, receiver, extraArgs);
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
    _consumeOutboundRateLimit(destinationChainSelector, amount);
    _handleTokenToTransfer(destinationChainSelector, msg.sender, amount + quote.fastTransferFee);

    if (feeToken != address(0)) {
      IERC20(feeToken).safeTransferFrom(msg.sender, address(this), quote.ccipSettlementFee);
      IERC20(feeToken).safeApprove(i_ccipRouter, quote.ccipSettlementFee);
    }
    fillRequestId = IRouterClient(getRouter()).ccipSend{value: msg.value}(destinationChainSelector, message);
    emit FastTransferRequested(fillRequestId, destinationChainSelector, amount, quote.fastTransferFee, receiver);
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
    _validateSendRequest(destinationChainSelector);
    DestChainConfig storage lane = s_fastTransferDestChainConfig[destinationChainSelector];
    if (amount > lane.maxFillAmountPerRequest) {
      revert TransferAmountExceedsMaxFillAmount(destinationChainSelector, amount);
    }
    quote.fastTransferFee = (amount * lane.fastTransferBpsFee) / 10_000;
    // pack the MintMessage
    bytes memory data = abi.encode(
      MintMessage({
        sourceAmountToTransfer: amount,
        sourceDecimals: i_tokenDecimals,
        fastTransferFee: quote.fastTransferFee,
        receiver: receiver
      })
    );

    message = Client.EVM2AnyMessage({
      receiver: lane.destinationPool,
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
  /// @param sourceDecimals The decimals of the source token
  /// @param receiver The receiver address
  function fastFill(
    bytes32 fillRequestId,
    uint64 sourceChainSelector,
    uint256 srcAmount,
    uint8 sourceDecimals,
    address receiver
  ) public virtual {
    bytes32 fillId = computeFillId(fillRequestId, srcAmount, receiver);
    FillInfo memory fillInfo = s_fills[fillId];
    if (fillInfo.state != FillState.NOT_FILLED) revert AlreadyFilled(fillRequestId);

    {
      DestChainConfig storage laneConfig = s_fastTransferDestChainConfig[sourceChainSelector];
      if (laneConfig.fillerAllowlistEnabled) {
        if (!laneConfig.fillerAllowList.contains(msg.sender)) {
          revert FillerNotAllowlisted(sourceChainSelector, msg.sender);
        }
      }
    }
    // Transfer tokens from filler to receiver
    uint256 destAmount = _transferFromFiller(sourceChainSelector, msg.sender, receiver, srcAmount, sourceDecimals);
    // Record fill
    s_fills[fillId] = FillInfo({state: FillState.FILLED, filler: msg.sender});
    emit FastTransferFilled(fillRequestId, fillId, msg.sender, destAmount, receiver);
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
        mintMessage.sourceAmountToTransfer,
        mintMessage.sourceDecimals,
        mintMessage.fastTransferFee,
        address(uint160(uint256(bytes32(mintMessage.receiver))))
      );
    }
    emit FastTransferSettled(message.messageId);
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
  /// @param sourceDecimals The decimals of the source token
  /// @return destAmount The amount transferred to the receiver on the destination chain
  function _transferFromFiller(
    uint64 sourceChainSelector,
    address filler,
    address receiver,
    uint256 srcAmount,
    uint8 sourceDecimals
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

    bytes32 fillId = computeFillId(fillRequestId, srcAmount, receiver);
    FillInfo memory fillInfo = s_fills[fillId];

    // Handle settlement based on fill state
    if (fillInfo.state == FillState.NOT_FILLED) {
      // Not fast-filled - mint/release to receiver
      _handleSlowFill(sourceChainSelector, settlementAmountLocal, receiver);
    } else if (fillInfo.state == FillState.SETTLED) {
      // Already settled
      revert AlreadySettled(fillRequestId);
    } else {
      // Fast-filled - reimburse filler
      _handleFastFilledReimbursement(fillInfo.filler, settlementAmountLocal);
    }

    // Mark as settled
    s_fills[fillId] = FillInfo({state: FillState.SETTLED, filler: address(0)});
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

  /// @notice Validates the send request parameters
  /// @param destinationChainSelector The destination chain selector
  /// @dev Checks if the destination chain is allowed, if the sender is whitelisted, and if the RMN curse applies
  function _validateSendRequest(
    uint64 destinationChainSelector
  ) internal view virtual {
    if (IRMN(i_rmnProxy).isCursed(bytes16(uint128(destinationChainSelector)))) revert CursedByRMN();
    _checkAllowList(msg.sender);
    if (!isSupportedChain(destinationChainSelector)) revert ChainNotAllowed(destinationChainSelector);
  }

  /// @notice Handles settlement when the request was not fast-filled
  /// @param sourceChainSelector The source chain selector
  /// @param settlementAmountLocal The amount to settle in local token
  /// @param receiver The receiver address
  function _handleSlowFill(
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
    return interfaceId == type(IFastTransferPool).interfaceId
      || interfaceId == type(IAny2EVMMessageReceiver).interfaceId || interfaceId == type(IERC165).interfaceId
      || super.supportsInterface(interfaceId);
  }

  function _computeFillId(
    bytes32 fillRequestId,
    uint256 amount,
    address receiver
  ) internal pure virtual returns (bytes32) {
    return keccak256(abi.encode(fillRequestId, amount, receiver));
  }
}
