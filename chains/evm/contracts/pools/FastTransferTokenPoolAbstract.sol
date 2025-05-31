// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IFastTransferPool} from "../interfaces/IFastTransferPool.sol";
import {IRMN} from "../interfaces/IRMN.sol";
import {IRouterClient} from "../interfaces/IRouterClient.sol";
import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";

import {CCIPReceiver} from "../applications/CCIPReceiver.sol";
import {Client} from "../libraries/Client.sol";
import {Internal} from "../libraries/Internal.sol";
import {TokenPool} from "./TokenPool.sol";

import {IERC20} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/utils/SafeERC20.sol";
import {EnumerableSet} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v4.8.3/contracts/utils/structs/EnumerableSet.sol";

/// @title Abstract Fast-Transfer Pool
/// @notice Base contract for fast-transfer pools that provides common functionality
/// for quoting, fill-tracking, and CCIP send helpers.
abstract contract FastTransferTokenPoolAbstract is TokenPool, CCIPReceiver, ITypeAndVersion, IFastTransferPool {
  using EnumerableSet for EnumerableSet.AddressSet;
  using SafeERC20 for IERC20;

  error InvalidDestChainConfig();
  error FillerNotAllowlisted(uint64 remoteChainSelector, address filler);
  error TransferAmountExceedsMaxFillAmount(uint64 remoteChainSelector, uint256 amount);

  event DestChainConfigUpdated(
    uint64 indexed destinationChainSelector,
    uint16 fastTransferBpsFee,
    uint256 maxFillAmountPerRequest,
    bytes destinationPool,
    bytes4 chainFamilySelector,
    uint256 settlementOverheadGas,
    bool fillerAllowlistEnabled
  );
  event FillerAllowListUpdated(uint64 indexed destChainSelector, address[] addFillers, address[] removeFillers);
  event DestinationPoolUpdated(uint64 indexed destChainSelector, address destinationPool);

  struct DestChainConfig {
    uint256 maxFillAmountPerRequest; //  Max amount that can be filled per request.
    bool fillerAllowlistEnabled; // ───╮ Allowlist for fillers.
    uint16 fastTransferBpsFee; //      │ Allowed range of [0-10_000].
    //                                 │ Settlement overhead gas for the destination chain. Used as a toggle for
    uint32 settlementOverheadGas; // ──╯ either custom ExtraArgs or GenericExtraArgsV2.
    bytes destinationPool; // Address of the destination pool.
    bytes customExtraArgs; // Pre-encoded extra args for EVM to Any message. Only used if settlementOverheadGas is 0.
  }

  struct DestChainConfigUpdateArgs {
    bool fillerAllowlistEnabled; // ──╮ Allowlist for fillers.
    uint16 fastTransferBpsFee; //     │ Allowed range of [0-10_000].
    uint32 settlementOverheadGas; //  │ Settlement overhead gas for the destination chain.
    uint64 remoteChainSelector; //    │ Remote chain selector. ABI encoded in the case of an EVM pool.
    bytes4 chainFamilySelector; // ───╯ Selector that identifies the destination chain's family.
    uint256 maxFillAmountPerRequest; // Maximum amount that can be filled per request.
    bytes destinationPool; // Address of the destination pool.
    bytes customExtraArgs; // Pre-encoded extra args for EVM to Any message. Only used if settlementOverheadGas is 0.
  }

  struct MintMessage {
    uint256 sourceAmount; // Amount to fill in the source token denomination.
    uint16 fastTransferFeeBps; // ─╮ Fast transfer fee in the source token.
    uint8 sourceDecimals; // ──────╯ Decimals of the source token.
    bytes receiver; // Receiver address on the destination chain. ABI encoded in the case of an EVM address.
  }

  /// @notice Enum representing the state of a fill request.
  enum FillState {
    NOT_FILLED, // Request has not been filled yet.
    FILLED, // Request has been filled by a filler.
    SETTLED // Request has been settled via CCIP.

  }

  /// @notice Struct to track fill request information.
  struct FillInfo {
    FillState state; // Current state of the fill request.
    address filler; // Address of the filler, 0x0 until filled. If 0x0 after filled, it means the request was not fast-filled.
  }

  /// @notice The division factor for basis points (BPS). This also represents the maximum BPS fee for fast transfer.
  uint256 internal constant BPS_DIVIDER = 10_000;

  /// @dev Mapping of remote chain selector to destinationChain configuration
  mapping(uint64 remoteChainSelector => DestChainConfig destinationChainConfig) private s_fastTransferDestChainConfig;

  /// @dev Mapping of remote chain selector to filler allowlist
  mapping(uint64 remoteChainSelector => EnumerableSet.AddressSet fillerAllowList) private s_fillerAllowLists;

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

  /// @notice Gets the fill information for a given fill ID
  /// @param fillId The fill ID to query
  /// @return fillInfo The fill information including state and filler address
  function getFillInfo(
    bytes32 fillId
  ) external view returns (FillInfo memory) {
    return s_fills[fillId];
  }

  /// @inheritdoc IFastTransferPool
  function ccipSendToken(
    address feeToken,
    uint64 destinationChainSelector,
    uint256 amount,
    bytes calldata receiver,
    bytes calldata extraArgs
  ) external payable virtual override returns (bytes32 fillRequestId) {
    (Quote memory quote, Client.EVM2AnyMessage memory message) =
      _getFeeQuoteAndCCIPMessage(feeToken, destinationChainSelector, amount, receiver, extraArgs);
    _consumeOutboundRateLimit(destinationChainSelector, amount);
    _handleFastTransferLockOrBurn(msg.sender, amount);

    // If the user is not paying in native, we need to transfer the fee token to the contract.
    if (feeToken != address(0)) {
      IERC20(feeToken).safeTransferFrom(msg.sender, address(this), quote.ccipSettlementFee);
      IERC20(feeToken).safeApprove(i_ccipRouter, quote.ccipSettlementFee);
    }

    fillRequestId = IRouterClient(getRouter()).ccipSend{value: msg.value}(destinationChainSelector, message);

    emit FastTransferRequested({
      fillRequestId: fillRequestId,
      destinationChainSelector: destinationChainSelector,
      amount: amount,
      fastTransferFee: quote.fastTransferFee,
      receiver: receiver
    });

    return fillRequestId;
  }

  /// @inheritdoc IFastTransferPool
  function computeFillId(
    bytes32 fillRequestId,
    uint256 amount,
    uint8 decimals,
    address receiver
  ) public pure override returns (bytes32) {
    return keccak256(abi.encode(fillRequestId, amount, decimals, receiver));
  }

  // ================================================================
  // │                      Fee calculation                         │
  // ================================================================

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

  /// @notice Pulls out all of the fee‐quotation + message‐build logic
  function _getFeeQuoteAndCCIPMessage(
    address settlementFeeToken,
    uint64 destinationChainSelector,
    uint256 amount,
    bytes calldata receiver,
    bytes calldata
  ) internal view virtual returns (IFastTransferPool.Quote memory quote, Client.EVM2AnyMessage memory message) {
    _validateSendRequest(destinationChainSelector);

    // TODO not use storage
    DestChainConfig storage destChainConfig = s_fastTransferDestChainConfig[destinationChainSelector];
    if (amount > destChainConfig.maxFillAmountPerRequest) {
      revert TransferAmountExceedsMaxFillAmount(destinationChainSelector, amount);
    }
    quote.fastTransferFee = (amount * destChainConfig.fastTransferBpsFee) / BPS_DIVIDER;

    bytes memory extraArgs;

    // We use 0 as a toggle for whether the destination chain requires custom ExtraArgs. Zero would not be a sensible
    // value for settlementOverheadGas, so we can use it as a toggle.
    if (destChainConfig.settlementOverheadGas == 0) {
      extraArgs = destChainConfig.customExtraArgs;
    } else {
      // If the value is not zero, we encode it as GenericExtraArgsV2.
      extraArgs = Client._argsToBytes(
        Client.GenericExtraArgsV2({gasLimit: destChainConfig.settlementOverheadGas, allowOutOfOrderExecution: true})
      );
    }

    message = Client.EVM2AnyMessage({
      receiver: destChainConfig.destinationPool,
      // pack the MintMessage
      data: abi.encode(
        MintMessage({
          sourceAmount: amount,
          sourceDecimals: i_tokenDecimals,
          fastTransferFeeBps: destChainConfig.fastTransferBpsFee,
          receiver: receiver
        })
      ),
      tokenAmounts: new Client.EVMTokenAmount[](0),
      feeToken: settlementFeeToken,
      extraArgs: extraArgs
    });

    quote.ccipSettlementFee = IRouterClient(getRouter()).getFee(destinationChainSelector, message);
    return (quote, message);
  }

  // ================================================================
  // │                           Filling                            │
  // ================================================================

  /// @notice Fast fills a transfer using liquidity provider funds based on CCIP settlement
  /// @param fillRequestId The fill request ID
  /// @param srcAmountToFill The amount to fill
  /// @param sourceDecimals The decimals of the source token
  /// @param receiver The receiver address
  function fastFill(
    bytes32 fillRequestId,
    uint64 sourceChainSelector,
    uint256 srcAmountToFill,
    uint8 sourceDecimals,
    address receiver
  ) public virtual {
    // Calculate the local amount.
    uint256 localAmount = _calculateLocalAmount(srcAmountToFill, sourceDecimals);
    // We rate limit when there are funds going to an end user, not when they are going to a filler.
    _consumeInboundRateLimit(sourceChainSelector, localAmount);
    // Transfer tokens from filler to receiver
    _transferFromFiller(msg.sender, receiver, localAmount);

    bytes32 fillId = computeFillId(fillRequestId, srcAmountToFill, sourceDecimals, receiver);
    FillInfo memory fillInfo = s_fills[fillId];
    if (fillInfo.state != FillState.NOT_FILLED) revert AlreadyFilled(fillRequestId);

    {
      DestChainConfig storage destChainConfig = s_fastTransferDestChainConfig[sourceChainSelector];
      if (destChainConfig.fillerAllowlistEnabled) {
        if (!s_fillerAllowLists[sourceChainSelector].contains(msg.sender)) {
          revert FillerNotAllowlisted(sourceChainSelector, msg.sender);
        }
      }
    }
    // Record fill
    s_fills[fillId] = FillInfo({state: FillState.FILLED, filler: msg.sender});
    emit FastTransferFilled(fillRequestId, fillId, msg.sender, localAmount, receiver);
  }

  // @inheritdoc CCIPReceiver
  function _ccipReceive(
    Client.Any2EVMMessage memory message
  ) internal virtual override onlyRouter {
    _settle(message.sourceChainSelector, message.messageId, message.sender, abi.decode(message.data, (MintMessage)));

    emit FastTransferSettled(message.messageId);
  }

  function _settle(
    uint64 sourceChainSelector,
    bytes32 fillRequestId,
    bytes memory sourcePoolAddress,
    MintMessage memory mintMessage
  ) internal virtual {
    address receiver = address(uint160(uint256(bytes32(mintMessage.receiver))));

    _validateSettlement(sourceChainSelector, sourcePoolAddress);

    // Inputs are in the source chain denomination, so we need to convert them to the local token denomination.
    uint256 localAmount = _calculateLocalAmount(mintMessage.sourceAmount, mintMessage.sourceDecimals);
    uint256 sourceAmountToFill = localAmount - (localAmount * mintMessage.fastTransferFeeBps) / BPS_DIVIDER;

    bytes32 fillId = computeFillId(fillRequestId, sourceAmountToFill, mintMessage.sourceDecimals, receiver);
    FillInfo memory fillInfo = s_fills[fillId];

    if (fillInfo.state == FillState.NOT_FILLED) {
      // Rate limits should be consumed only when the request was not fast-filled. During fast fill, the rate limit is
      // consumed by the filler.
      _consumeInboundRateLimit(sourceChainSelector, localAmount);
      // When no filler is involved, we send the entire amount to the receiver.
      _handleSlowFill(localAmount, receiver);
    } else if (fillInfo.state == FillState.FILLED) {
      _handleFastFilledReimbursement(fillInfo.filler, localAmount);
    } else {
      // The catch all assertion for already settled fills ensures that any wrong value will revert.
      revert AlreadySettled(fillRequestId);
    }

    s_fills[fillId].state = FillState.SETTLED;
  }

  /// @notice Validates settlement prerequisites. Can be overridden by derived contracts to add additional checks.
  /// @param sourceChainSelector The source chain selector
  /// @param sourcePoolAddress The source pool address
  function _validateSettlement(uint64 sourceChainSelector, bytes memory sourcePoolAddress) internal view virtual {
    if (IRMN(i_rmnProxy).isCursed(bytes16(uint128(sourceChainSelector)))) revert CursedByRMN();
    //Validates that the source pool address is configured on this pool.
    if (!isRemotePool(sourceChainSelector, sourcePoolAddress)) {
      revert InvalidSourcePoolAddress(sourcePoolAddress);
    }
  }

  /// @notice Validates the send request parameters. Can be overridden by derived contracts to add additional checks.
  /// @param destinationChainSelector The destination chain selector.
  /// @dev Checks if the destination chain is allowed, if the sender is allowed, and if the RMN curse applies.
  function _validateSendRequest(
    uint64 destinationChainSelector
  ) internal view virtual {
    if (IRMN(i_rmnProxy).isCursed(bytes16(uint128(destinationChainSelector)))) revert CursedByRMN();
    _checkAllowList(msg.sender);
    if (!isSupportedChain(destinationChainSelector)) revert ChainNotAllowed(destinationChainSelector);
  }

  // ================================================================
  // │                      Filling Hooks                           │
  // ================================================================

  /// @notice Handles the token to transfer on fast fill request at source chain
  /// @param sender The sender address
  /// @param amount The amount to transfer
  function _handleFastTransferLockOrBurn(address sender, uint256 amount) internal virtual {
    // Since this is a fast transfer, the Router doesn't forward the tokens to the pool.
    i_token.safeTransferFrom(sender, address(this), amount);
    // Use the normal burn logic once the tokens are in the pool.
    _lockOrBurn(amount);
  }

  /// @notice Transfers tokens from the filler to the receiver.
  /// @param filler The address of the filler.
  /// @param receiver The address of the receiver.
  /// @param amount The amount to transfer in local denomination.
  function _transferFromFiller(address filler, address receiver, uint256 amount) internal virtual {
    getToken().safeTransferFrom(filler, receiver, amount);
  }

  /// @notice Handles settlement when the request was not fast-filled
  /// @param settlementAmountLocal The amount to settle in local token
  /// @param receiver The receiver address
  function _handleSlowFill(uint256 settlementAmountLocal, address receiver) internal virtual {
    _releaseOrMint(receiver, settlementAmountLocal);
  }

  /// @notice Handles reimbursement when the request was fast-filled
  /// @param filler The filler address to reimburse
  /// @param settlementAmountLocal The amount to reimburse in local token
  function _handleFastFilledReimbursement(address filler, uint256 settlementAmountLocal) internal virtual {
    // Honest filler -> pay them back + fee
    _releaseOrMint(filler, settlementAmountLocal);
  }

  // ================================================================
  // │                          Config                              │
  // ================================================================

  /// @notice Override getRouter to resolve diamond inheritance
  function getRouter() public view virtual override(TokenPool, CCIPReceiver) returns (address) {
    return TokenPool.getRouter();
  }

  /// @notice Gets the destChain configuration for a given destination chain selector
  /// @param remoteChainSelector The remote chain selector
  /// @return destChainConfig The destChain configuration for the given destination chain selector
  function getDestChainConfig(
    uint64 remoteChainSelector
  ) external view virtual returns (DestChainConfig memory, address[] memory) {
    return (s_fastTransferDestChainConfig[remoteChainSelector], s_fillerAllowLists[remoteChainSelector].values());
  }

  /// @notice Updates the destination chain configuration
  /// @param destChainConfigArgs The destChain configuration arguments
  function updateDestChainConfig(
    DestChainConfigUpdateArgs[] calldata destChainConfigArgs
  ) external virtual onlyOwner {
    for (uint256 i = 0; i < destChainConfigArgs.length; ++i) {
      _updateDestChainConfig(destChainConfigArgs[i]);
    }
  }

  function _updateDestChainConfig(
    DestChainConfigUpdateArgs calldata destChainConfigArgs
  ) internal virtual {
    // We know Solana requires custom args, if they are not provided, we revert.
    if (destChainConfigArgs.chainFamilySelector == Internal.CHAIN_FAMILY_SELECTOR_SVM) {
      if (destChainConfigArgs.settlementOverheadGas != 0) {
        revert InvalidDestChainConfig();
      }
    }

    if (destChainConfigArgs.fastTransferBpsFee > BPS_DIVIDER) revert InvalidDestChainConfig();

    DestChainConfig storage destChainConfig = s_fastTransferDestChainConfig[destChainConfigArgs.remoteChainSelector];
    destChainConfig.destinationPool = destChainConfigArgs.destinationPool;
    destChainConfig.fastTransferBpsFee = destChainConfigArgs.fastTransferBpsFee;
    destChainConfig.fillerAllowlistEnabled = destChainConfigArgs.fillerAllowlistEnabled;
    destChainConfig.maxFillAmountPerRequest = destChainConfigArgs.maxFillAmountPerRequest;
    destChainConfig.settlementOverheadGas = destChainConfigArgs.settlementOverheadGas;
    destChainConfig.customExtraArgs = destChainConfigArgs.customExtraArgs;

    emit DestChainConfigUpdated(
      destChainConfigArgs.remoteChainSelector,
      destChainConfigArgs.fastTransferBpsFee,
      destChainConfigArgs.maxFillAmountPerRequest,
      destChainConfigArgs.destinationPool,
      destChainConfigArgs.chainFamilySelector,
      destChainConfigArgs.settlementOverheadGas,
      destChainConfigArgs.fillerAllowlistEnabled
    );
  }

  /// @notice Override supportsInterface to resolve diamond inheritance
  function supportsInterface(
    bytes4 interfaceId
  ) public pure virtual override(TokenPool, CCIPReceiver) returns (bool) {
    return interfaceId == type(IFastTransferPool).interfaceId || TokenPool.supportsInterface(interfaceId)
      || CCIPReceiver.supportsInterface(interfaceId);
  }

  // ================================================================
  // │                      Filler allowlist                        │
  // ================================================================

  /// @notice Gets all allowlisted fillers for a given destination chain
  /// @param remoteChainSelector The remote chain selector
  /// @return fillers Array of allowlisted filler addresses
  function getAllowedFillers(
    uint64 remoteChainSelector
  ) external view virtual returns (address[] memory) {
    return s_fillerAllowLists[remoteChainSelector].values();
  }

  /// @notice Checks if a filler is allowlisted for a given destChain.
  /// @param remoteChainSelector The remote chain selector.
  /// @param filler The filler address to check.
  /// @return True if the filler is allowed, false otherwise.
  function isAllowedFiller(uint64 remoteChainSelector, address filler) external view virtual returns (bool) {
    return s_fillerAllowLists[remoteChainSelector].contains(filler);
  }

  /// @notice Updates the filler allowlist configuration for a given lane.
  /// @param destinationChainSelector The destination chain selector.
  /// @param fillersToAdd The addresses to add to the allowlist.
  /// @param fillersToRemove The addresses to remove from the allowlist.
  function updateFillerAllowList(
    uint64 destinationChainSelector,
    address[] memory fillersToAdd,
    address[] memory fillersToRemove
  ) external virtual onlyOwner {
    EnumerableSet.AddressSet storage allowList = s_fillerAllowLists[destinationChainSelector];

    for (uint256 i = 0; i < fillersToAdd.length; ++i) {
      allowList.add(fillersToAdd[i]);
    }
    for (uint256 i = 0; i < fillersToRemove.length; ++i) {
      allowList.remove(fillersToRemove[i]);
    }

    emit FillerAllowListUpdated(destinationChainSelector, fillersToAdd, fillersToRemove);
  }
}
