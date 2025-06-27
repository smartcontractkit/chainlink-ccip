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

/// @notice Base contract for fast-transfer pools that provides common functionality
/// for quoting, fill-tracking, and CCIP send helpers.
/// @dev To make this contract usable, it must be inherited by a concrete implementation that implements:
/// - `_lockOrBurn` - handles both the TokenPool lock/burn and the fast transfer lock/burn.
/// - `_releaseOrMint` - handles both the TokenPool release/mint and the fast transfer release/mint.
/// Additionally, there are some hooks that can optionally be overridden:
/// - `_handleFastTransferLockOrBurn`
/// - `_handleFastFill`
/// - `_handleSlowFill`
/// - `_handleFastFillReimbursement`
/// There are also validation functions that can optionally be overridden:
/// - `_validateSendRequest` - called before sending a fast transfer.
/// - `_validateSettlement` - called before settling.
abstract contract FastTransferTokenPoolAbstract is TokenPool, CCIPReceiver, ITypeAndVersion, IFastTransferPool {
  using EnumerableSet for EnumerableSet.AddressSet;
  using SafeERC20 for IERC20;

  error InvalidDestChainConfig();
  error FillerNotAllowlisted(uint64 remoteChainSelector, address filler);
  error InvalidFillId(bytes32 fillId);
  error TransferAmountExceedsMaxFillAmount(uint64 remoteChainSelector, uint256 amount);
  error InsufficientPoolFees(uint256 requested, uint256 available);
  error QuoteFeeExceedsUserMaxLimit(uint256 quoteFee, uint256 maxFastTransferFee);

  event DestChainConfigUpdated(
    uint64 indexed destinationChainSelector,
    uint16 fastTransferFillerFeeBps,
    uint16 fastTransferPoolFeeBps,
    uint256 maxFillAmountPerRequest,
    bytes destinationPool,
    bytes4 chainFamilySelector,
    uint256 settlementOverheadGas,
    bool fillerAllowlistEnabled
  );
  event FillerAllowListUpdated(address[] addFillers, address[] removeFillers);
  event DestinationPoolUpdated(uint64 indexed destChainSelector, address destinationPool);

  struct DestChainConfig {
    uint256 maxFillAmountPerRequest; // Maximum amount that can be filled per request.
    bool fillerAllowlistEnabled; // ────╮ Allowlist for fillers.
    uint16 fastTransferFillerFeeBps; // │ Basis points fee going to filler [0-10_000].
    uint16 fastTransferPoolFeeBps; //   │ Basis points fee going to pool [0-10_000].
    //                                  │ Settlement overhead gas for the destination chain. Used as a toggle for
    uint32 settlementOverheadGas; // ───╯ either custom ExtraArgs or GenericExtraArgsV2.
    bytes destinationPool; // Address of the destination pool.
    bytes customExtraArgs; // Pre-encoded extra args for EVM to Any message. Only used if settlementOverheadGas is 0.
  }

  struct DestChainConfigUpdateArgs {
    bool fillerAllowlistEnabled; // ────╮ Allowlist for fillers.
    uint16 fastTransferFillerFeeBps; // │ Basis points fee going to filler [0-10_000].
    uint16 fastTransferPoolFeeBps; //   │ Basis points fee going to pool [0-10_000].
    uint32 settlementOverheadGas; //    │ Settlement overhead gas for the destination chain.
    uint64 remoteChainSelector; //      │ Remote chain selector. ABI encoded in the case of an EVM pool.
    bytes4 chainFamilySelector; // ─────╯ Selector that identifies the destination chain's family.
    uint256 maxFillAmountPerRequest; // Maximum amount that can be filled per request.
    bytes destinationPool; // Address of the destination pool.
    bytes customExtraArgs; // Pre-encoded extra args for EVM to Any message. Only used if settlementOverheadGas is 0.
  }

  struct MintMessage {
    uint256 sourceAmount; // Amount to fill in the source token denomination.
    uint16 fastTransferFillerFeeBps; // ─╮ Basis points fee going to filler.
    uint16 fastTransferPoolFeeBps; //    │ Basis points fee going to pool.
    uint8 sourceDecimals; // ────────────╯ Decimals of the source token.
    bytes receiver; // Receiver address on the destination chain. ABI encoded in the case of an EVM address.
  }

  /// @notice Struct to track fill request information.
  struct FillInfo {
    IFastTransferPool.FillState state; // Current state of the fill request.
    // Address of the filler, 0x0 until filled. If 0x0 after filled, it means the request was not fast-filled.
    address filler;
  }

  /// @notice The division factor for basis points (BPS). This also represents the maximum BPS fee for fast transfer.
  uint256 internal constant BPS_DIVIDER = 10_000;

  /// @dev Mapping of remote chain selector to destinationChain configuration.
  mapping(uint64 remoteChainSelector => DestChainConfig destinationChainConfig) internal s_fastTransferDestChainConfig;

  /// @dev Only addresses present in this list are able to fill.
  EnumerableSet.AddressSet internal s_fillerAllowLists;

  /// @dev Mapping of fill request ID to fill information.
  /// This is used to track the state and filler of each fill request.
  mapping(bytes32 fillId => FillInfo fillInfo) internal s_fills;

  /// @param token The token this pool manages.
  /// @param localTokenDecimals The decimals of the local token.
  /// @param allowlist The allowlist of addresses.
  /// @param rmnProxy The RMN proxy address.
  /// @param router Address of the CCIP router.
  constructor(
    IERC20 token,
    uint8 localTokenDecimals,
    address[] memory allowlist,
    address rmnProxy,
    address router
  ) TokenPool(token, localTokenDecimals, allowlist, rmnProxy, router) CCIPReceiver(router) {}

  /// @notice Gets the fill information for a given fill ID.
  /// @return fillInfo The fill information including state and filler address.
  function getFillInfo(
    bytes32 fillId
  ) external view returns (FillInfo memory) {
    return s_fills[fillId];
  }

  /// @inheritdoc IFastTransferPool
  function ccipSendToken(
    uint64 destinationChainSelector,
    uint256 amount,
    uint256 maxFastTransferFee,
    bytes calldata receiver,
    address settlementFeeToken,
    bytes calldata extraArgs
  ) external payable virtual override returns (bytes32 settlementId) {
    (Quote memory quote, Client.EVM2AnyMessage memory message) =
      _getFeeQuoteAndCCIPMessage(destinationChainSelector, amount, receiver, settlementFeeToken, extraArgs);
    _consumeOutboundRateLimit(destinationChainSelector, amount);
    if (quote.fastTransferFee > maxFastTransferFee) {
      revert QuoteFeeExceedsUserMaxLimit(quote.fastTransferFee, maxFastTransferFee);
    }
    _handleFastTransferLockOrBurn(msg.sender, amount);

    // If the user is not paying in native, we need to transfer the fee token to the contract.
    if (settlementFeeToken != address(0)) {
      IERC20(settlementFeeToken).safeTransferFrom(msg.sender, address(this), quote.ccipSettlementFee);
      IERC20(settlementFeeToken).safeApprove(i_ccipRouter, quote.ccipSettlementFee);
    }

    settlementId = IRouterClient(getRouter()).ccipSend{value: msg.value}(destinationChainSelector, message);
    uint256 amountNetFee = amount - quote.fastTransferFee;

    emit FastTransferRequested({
      destinationChainSelector: destinationChainSelector,
      fillId: computeFillId(settlementId, amountNetFee, i_tokenDecimals, receiver),
      settlementId: settlementId,
      sourceAmountNetFee: amountNetFee,
      sourceDecimals: i_tokenDecimals,
      fastTransferFee: quote.fastTransferFee,
      receiver: receiver
    });

    return settlementId;
  }

  /// @inheritdoc IFastTransferPool
  function computeFillId(
    bytes32 settlementId,
    uint256 sourceAmountNetFee,
    uint8 sourceDecimals,
    bytes memory receiver
  ) public pure override returns (bytes32) {
    return keccak256(abi.encode(settlementId, sourceAmountNetFee, sourceDecimals, receiver));
  }

  // ================================================================
  // │                      Fee calculation                         │
  // ================================================================

  /// @notice Calculates the filler and pool fees for a fast transfer.
  /// @dev Common function to ensure consistent fee calculation
  /// @param amount The transfer amount
  /// @param fillerFeeBps Filler fee in basis points
  /// @param poolFeeBps Pool fee in basis points
  /// @return fillerFee The calculated filler fee
  /// @return poolFee The calculated pool fee
  function _calculateFastTransferFees(
    uint256 amount,
    uint16 fillerFeeBps,
    uint16 poolFeeBps
  ) internal pure returns (uint256 fillerFee, uint256 poolFee) {
    // Calculate individual fees using separate divisions to ensure consistency
    fillerFee = (amount * fillerFeeBps) / BPS_DIVIDER;
    poolFee = (amount * poolFeeBps) / BPS_DIVIDER;
    return (fillerFee, poolFee);
  }

  /// @inheritdoc IFastTransferPool
  function getCcipSendTokenFee(
    uint64 destinationChainSelector,
    uint256 amount,
    bytes calldata receiver,
    address settlementFeeToken,
    bytes calldata extraArgs
  ) public view virtual override returns (Quote memory) {
    (Quote memory quote,) =
      _getFeeQuoteAndCCIPMessage(destinationChainSelector, amount, receiver, settlementFeeToken, extraArgs);
    return quote;
  }

  function _getFeeQuoteAndCCIPMessage(
    uint64 destinationChainSelector,
    uint256 amount,
    bytes calldata receiver,
    address settlementFeeToken,
    bytes calldata
  ) internal view virtual returns (IFastTransferPool.Quote memory quote, Client.EVM2AnyMessage memory message) {
    _validateSendRequest(destinationChainSelector);

    // Using storage here appears to be cheaper.
    DestChainConfig storage destChainConfig = s_fastTransferDestChainConfig[destinationChainSelector];
    if (amount > destChainConfig.maxFillAmountPerRequest) {
      revert TransferAmountExceedsMaxFillAmount(destinationChainSelector, amount);
    }

    (uint256 fillerFee, uint256 poolFee) = _calculateFastTransferFees(
      amount, destChainConfig.fastTransferFillerFeeBps, destChainConfig.fastTransferPoolFeeBps
    );
    quote.fastTransferFee = fillerFee + poolFee;
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
          fastTransferFillerFeeBps: destChainConfig.fastTransferFillerFeeBps,
          fastTransferPoolFeeBps: destChainConfig.fastTransferPoolFeeBps,
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

  /// @notice Fast fills a transfer using liquidity provider funds based on CCIP settlement.
  /// @param fillId The fill ID, computed from the fill request parameters.
  /// @param settlementId The settlement ID, which under normal circumstances is the same as the CCIP message ID.
  /// @param sourceAmountNetFee The amount to fill, calculated as the amount sent in `ccipSendToken` minus
  /// the fast fill fee, expressed in source token decimals.
  /// @param sourceDecimals The decimals of the source token.
  /// @param receiver The receiver address.
  function fastFill(
    bytes32 fillId,
    bytes32 settlementId,
    uint64 sourceChainSelector,
    uint256 sourceAmountNetFee,
    uint8 sourceDecimals,
    address receiver
  ) public virtual {
    if (s_fastTransferDestChainConfig[sourceChainSelector].fillerAllowlistEnabled) {
      if (!s_fillerAllowLists.contains(msg.sender)) {
        revert FillerNotAllowlisted(sourceChainSelector, msg.sender);
      }
    }

    if (fillId != computeFillId(settlementId, sourceAmountNetFee, sourceDecimals, abi.encode(receiver))) {
      revert InvalidFillId(fillId);
    }

    FillInfo memory fillInfo = s_fills[fillId];
    if (fillInfo.state != IFastTransferPool.FillState.NOT_FILLED) revert AlreadyFilledOrSettled(fillId);

    // Calculate the local amount.
    uint256 localAmount = _calculateLocalAmount(sourceAmountNetFee, sourceDecimals);

    s_fills[fillId] = FillInfo({state: IFastTransferPool.FillState.FILLED, filler: msg.sender});

    emit FastTransferFilled(fillId, settlementId, msg.sender, localAmount, receiver);

    _handleFastFill(fillId, msg.sender, receiver, localAmount);
  }

  // @inheritdoc CCIPReceiver
  function _ccipReceive(
    Client.Any2EVMMessage memory message
  ) internal virtual override onlyRouter {
    _settle(message.sourceChainSelector, message.messageId, message.sender, abi.decode(message.data, (MintMessage)));
  }

  function _settle(
    uint64 sourceChainSelector,
    bytes32 settlementId,
    bytes memory sourcePoolAddress,
    MintMessage memory mintMessage
  ) internal virtual {
    _validateSettlement(sourceChainSelector, sourcePoolAddress);

    // Calculate the fast transfer inputs
    (uint256 sourceFillerFee, uint256 sourcePoolFee) = _calculateFastTransferFees(
      mintMessage.sourceAmount, mintMessage.fastTransferFillerFeeBps, mintMessage.fastTransferPoolFeeBps
    );
    // Inputs are in the source chain denomination, so we need to convert them to the local token denomination.
    uint256 localAmount = _calculateLocalAmount(mintMessage.sourceAmount, mintMessage.sourceDecimals);
    uint256 localPoolFee = _calculateLocalAmount(sourcePoolFee, mintMessage.sourceDecimals);
    bytes32 fillId = computeFillId(
      settlementId,
      // sourceAmountNetFee is the amount minus the fast fill fee, so we need to subtract both fees.
      mintMessage.sourceAmount - sourceFillerFee - sourcePoolFee,
      mintMessage.sourceDecimals,
      mintMessage.receiver
    );

    // Cache current fill info to decide which hook to call.
    FillInfo memory fillInfo = s_fills[fillId];
    /// Mark the fill as SETTLED before any value transfers or external calls.
    /// This makes the new state visible immediately, preventing the same fill
    /// from being settled twice even if execution re-enters this contract.
    s_fills[fillId].state = IFastTransferPool.FillState.SETTLED;
    // Rate limiting should apply to the full sourceAmount regardless of whether the request was fast-filled or not.
    // This ensures that the rate limit controls the overall rate of release/mint operations.
    _consumeInboundRateLimit(sourceChainSelector, localAmount);

    // The amount to reimburse to the filler in local denomination.
    uint256 fillerReimbursementAmount = 0;
    if (fillInfo.state == IFastTransferPool.FillState.NOT_FILLED) {
      // Set the local pool fee to zero, as fees are only applied for fast-fill operations
      localPoolFee = 0;
      // When no filler is involved, we send the entire amount to the receiver.
      _handleSlowFill(fillId, localAmount, abi.decode(mintMessage.receiver, (address)));
    } else if (fillInfo.state == IFastTransferPool.FillState.FILLED) {
      fillerReimbursementAmount = localAmount - localPoolFee;
      _handleFastFillReimbursement(fillId, fillInfo.filler, fillerReimbursementAmount, localPoolFee);
    } else {
      // The catch all assertion for already settled fills ensures that any wrong value will revert.
      revert AlreadySettled(fillId);
    }
    emit FastTransferSettled(fillId, settlementId, fillerReimbursementAmount, localPoolFee, fillInfo.state);
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

  /// @notice Validates settlement prerequisites. Can be overridden by derived contracts to add additional checks.
  /// @param sourceChainSelector The source chain selector.
  /// @param sourcePoolAddress The source pool address.
  function _validateSettlement(uint64 sourceChainSelector, bytes memory sourcePoolAddress) internal view virtual {
    if (IRMN(i_rmnProxy).isCursed(bytes16(uint128(sourceChainSelector)))) revert CursedByRMN();
    //Validates that the source pool address is configured on this pool.
    if (!isRemotePool(sourceChainSelector, sourcePoolAddress)) {
      revert InvalidSourcePoolAddress(sourcePoolAddress);
    }
  }

  // ================================================================
  // │                      Filling Hooks                           │
  // ================================================================

  /// @notice Handles the token to transfer on fast fill request at source chain.
  /// @param sender The sender address.
  /// @param amount The amount to transfer.
  function _handleFastTransferLockOrBurn(address sender, uint256 amount) internal virtual {
    // Since this is a fast transfer, the Router doesn't forward the tokens to the pool.
    getToken().safeTransferFrom(sender, address(this), amount);
    // Use the normal burn logic once the tokens are in the pool.
    _lockOrBurn(amount);
  }

  /// @notice Transfers tokens from the filler to the receiver.
  /// @dev The first param is the fillId. It's unused in this implementation, but kept to allow overriding this function
  /// to handle the transfer in a different way.
  /// @param filler The address of the filler.
  /// @param receiver The address of the receiver.
  /// @param amount The amount to transfer in local denomination.
  function _handleFastFill(bytes32, address filler, address receiver, uint256 amount) internal virtual {
    getToken().safeTransferFrom(filler, receiver, amount);
  }

  /// @notice Handles settlement when the request was not fast-filled
  /// @dev The first param is the fillId. It's unused in this implementation, but kept to allow overriding this function
  /// to handle the slow fill in a different way.
  /// @param localSettlementAmount The amount to settle in local token
  /// @param receiver The receiver address
  function _handleSlowFill(bytes32, uint256 localSettlementAmount, address receiver) internal virtual {
    _releaseOrMint(receiver, localSettlementAmount);
  }

  /// @notice Handles reimbursement when the request was fast-filled.
  /// @dev The first param is the fillId. It's unused in this implementation, but kept to allow overriding this function
  /// to handle the reimbursement in a different way.
  ///
  /// Burn/Mint token pools:
  /// This default implementation mints pool fee rewards directly to the pool itself (address(this)).
  /// The pool contract itself holds the reward tokens and they can be managed through standard token operations.
  ///
  /// Lock/Release pools:
  /// Lock/Release pools should override this function to implement accounting-based fee management since they
  /// cannot mint new tokens. They should keep track of accumulated pool fees in a storage variable (e.g., s_accumulatedPoolFees)
  /// @param filler The filler address to reimburse.
  /// @param fillerReimbursementAmount The amount to reimburse (what they provided + their fee).
  /// @param poolReimbursementAmount The amount to reimburse to the pool (the pool fee).
  function _handleFastFillReimbursement(
    bytes32,
    address filler,
    uint256 fillerReimbursementAmount,
    uint256 poolReimbursementAmount
  ) internal virtual {
    // Mint entire amount to pool first
    _releaseOrMint(address(this), fillerReimbursementAmount + poolReimbursementAmount);

    // Then transfer filler's share to them
    if (fillerReimbursementAmount > 0) {
      getToken().safeTransfer(filler, fillerReimbursementAmount);
    }
  }

  // ================================================================
  // │                          Config                              │
  // ================================================================

  /// @notice Override getRouter to resolve both TokenPool and CCIPReceiver implementing getRouter().
  function getRouter() public view virtual override(TokenPool, CCIPReceiver) returns (address) {
    return TokenPool.getRouter();
  }

  /// @notice Gets the destChain configuration for a given destination chain selector.
  /// @param remoteChainSelector The remote chain selector.
  /// @return destChainConfig The destChain configuration for the given destination chain selector.
  function getDestChainConfig(
    uint64 remoteChainSelector
  ) external view virtual returns (DestChainConfig memory, address[] memory) {
    return (s_fastTransferDestChainConfig[remoteChainSelector], s_fillerAllowLists.values());
  }

  /// @notice Gets the accumulated pool fees that can be withdrawn.
  /// @dev This is an abstract function that must be implemented by derived contracts.
  /// Burn/Mint pools : Should return the contract's token balance since pool fees
  /// are minted directly to the pool contract (e.g., `return getToken().balanceOf(address(this))`).
  /// Lock/Release pools : Should implement their own accounting mechanism for pool fees
  /// by adding a storage variable (e.g., `s_accumulatedPoolFees`) since they cannot mint
  /// additional tokens for pool fee rewards.
  /// Note: Fee accounting can be obscured by sending tokens directly to the pool.
  /// This does not introduce security issues but will need to be handled operationally.
  /// @return The amount of accumulated pool fees available for withdrawal.
  function getAccumulatedPoolFees() public view virtual returns (uint256);

  /// @notice Withdraws all accumulated pool fees to the specified recipient.
  /// @dev For BURN/MINT pools, this transfers the entire token balance of the pool contract.
  /// LOCK/RELEASE pools should override this function with their own accounting mechanism.
  /// @param recipient The address to receive the withdrawn fees.
  function withdrawPoolFees(
    address recipient
  ) external virtual onlyOwner {
    uint256 amount = getAccumulatedPoolFees();
    if (amount > 0) {
      getToken().safeTransfer(recipient, amount);
      emit PoolFeeWithdrawn(recipient, amount);
    }
  }

  /// @notice Updates the destination chain configuration.
  /// @param destChainConfigArgs The destChain configuration arguments.
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

    // Ensure total fees is below 100%
    if (destChainConfigArgs.fastTransferFillerFeeBps + destChainConfigArgs.fastTransferPoolFeeBps >= BPS_DIVIDER) {
      revert InvalidDestChainConfig();
    }

    DestChainConfig storage destChainConfig = s_fastTransferDestChainConfig[destChainConfigArgs.remoteChainSelector];
    destChainConfig.destinationPool = destChainConfigArgs.destinationPool;
    destChainConfig.fastTransferFillerFeeBps = destChainConfigArgs.fastTransferFillerFeeBps;
    destChainConfig.fastTransferPoolFeeBps = destChainConfigArgs.fastTransferPoolFeeBps;
    destChainConfig.fillerAllowlistEnabled = destChainConfigArgs.fillerAllowlistEnabled;
    destChainConfig.maxFillAmountPerRequest = destChainConfigArgs.maxFillAmountPerRequest;
    destChainConfig.settlementOverheadGas = destChainConfigArgs.settlementOverheadGas;
    destChainConfig.customExtraArgs = destChainConfigArgs.customExtraArgs;

    emit DestChainConfigUpdated(
      destChainConfigArgs.remoteChainSelector,
      destChainConfigArgs.fastTransferFillerFeeBps,
      destChainConfigArgs.fastTransferPoolFeeBps,
      destChainConfigArgs.maxFillAmountPerRequest,
      destChainConfigArgs.destinationPool,
      destChainConfigArgs.chainFamilySelector,
      destChainConfigArgs.settlementOverheadGas,
      destChainConfigArgs.fillerAllowlistEnabled
    );
  }

  /// @notice Override supportsInterface to resolve the double inheritance.
  function supportsInterface(
    bytes4 interfaceId
  ) public pure virtual override(TokenPool, CCIPReceiver) returns (bool) {
    return interfaceId == type(IFastTransferPool).interfaceId || TokenPool.supportsInterface(interfaceId)
      || CCIPReceiver.supportsInterface(interfaceId);
  }

  // ================================================================
  // │                      Filler allowlist                        │
  // ================================================================

  /// @notice Gets all allowlisted fillers for a given destination chain.
  /// @return fillers Array of allowlisted filler addresses.
  function getAllowedFillers() external view virtual returns (address[] memory) {
    return s_fillerAllowLists.values();
  }

  /// @notice Checks if a filler is allowlisted for a given destChain.
  /// @param filler The filler address to check.
  /// @return True if the filler is allowed, false otherwise.
  function isAllowedFiller(
    address filler
  ) external view virtual returns (bool) {
    return s_fillerAllowLists.contains(filler);
  }

  /// @notice Updates the filler allowlist configuration for a given lane.
  /// @param fillersToAdd The addresses to add to the allowlist.
  /// @param fillersToRemove The addresses to remove from the allowlist.
  function updateFillerAllowList(
    address[] memory fillersToAdd,
    address[] memory fillersToRemove
  ) external virtual onlyOwner {
    for (uint256 i = 0; i < fillersToAdd.length; ++i) {
      s_fillerAllowLists.add(fillersToAdd[i]);
    }
    for (uint256 i = 0; i < fillersToRemove.length; ++i) {
      s_fillerAllowLists.remove(fillersToRemove[i]);
    }

    emit FillerAllowListUpdated(fillersToAdd, fillersToRemove);
  }
}
