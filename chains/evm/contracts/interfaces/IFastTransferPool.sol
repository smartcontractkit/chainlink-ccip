// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.4;

/// @title IFastTransferPool
/// @notice Interface for the CCIP Fast-Transfer Pool.
interface IFastTransferPool {
  /// @notice Enum representing the state of a fill request.
  enum FillState {
    NOT_FILLED, // Request has not been filled yet.
    FILLED, // Request has been filled by a filler.
    SETTLED // Request has been settled via CCIP.

  }

  /// @notice Quote struct containing fee information.
  struct Quote {
    uint256 ccipSettlementFee; // Fee paid to for CCIP settlement in CCIP supported fee tokens.
    uint256 fastTransferFee; // Fee paid to the fast transfer filler in the same asset as requested.
  }

  error AlreadyFilledOrSettled(bytes32 fillId);
  error AlreadySettled(bytes32 fillId);

  /// @notice Emitted when a fast transfer is requested.
  event FastTransferRequested(
    uint64 indexed destinationChainSelector,
    bytes32 indexed fillId,
    bytes32 indexed settlementId,
    /// @param sourceAmountNetFee The amount being transferred, excluding the fast fill fee, expressed in source token
    /// decimals.
    uint256 sourceAmountNetFee,
    uint8 sourceDecimals,
    uint256 fillerFee,
    uint256 poolFee,
    /// @param destinationPool The destination chain pool where both the fill and settlement processes occur.
    /// @dev Fillers must invoke `fill` on the exact `destinationPool` address specified in the event tied to a fast transfer request.
    /// This ensures proper handling, as the active destination pool address can be updated in the token admin registry during token pool upgrades.
    /// In such cases, inflight messages are routed to the pool address specified in the event, where settlement takes place.
    /// To ensure accurate compensation during settlement, the fast fill must also occur at the pool address specified in the event.
    /// @notice Observability tools or indexing components should observe the `FastTransferFilled` and `FastTransferSettled` events from the `destinationPool`
    /// emitted in this event to monitor both fill and settlement actions for accurate status and metrics.
    bytes destinationPool,
    bytes receiver
  );
  /// @notice Emitted when a fast transfer is filled. This means the end user has received the tokens but the slow
  /// transfer is still in progress.
  event FastTransferFilled(
    bytes32 indexed fillId, bytes32 indexed settlementId, address indexed filler, uint256 destAmount, address receiver
  );
  /// @notice Emitted when a fast transfer is settled. This means the slow transfer has completed and the filler has
  /// received their fast transfer tokens and fee.
  event FastTransferSettled(
    bytes32 indexed fillId,
    bytes32 indexed settlementId,
    uint256 fillerReimbursementAmount,
    uint256 poolFeeAccumulated,
    FillState prevState
  );

  /// @notice Emitted when pool fees are withdrawn.
  event PoolFeeWithdrawn(address indexed recipient, uint256 amount);

  /// @notice Gets the CCIP send token fee and fast transfer fee for a given transfer.
  /// @param destinationChainSelector The destination chain selector.
  /// @param amount The amount to transfer.
  /// @param receiver The receiver address.
  /// @param settlementFeeToken The token used to pay the CCIP settlement fee.
  /// @param extraArgs Extra arguments for the transfer.
  /// @return Quote containing the CCIP fee and fast transfer fee.
  function getCcipSendTokenFee(
    uint64 destinationChainSelector,
    uint256 amount,
    bytes calldata receiver,
    address settlementFeeToken,
    bytes calldata extraArgs
  ) external view returns (Quote memory);

  /// @notice Sends tokens via CCIP with optional fast transfer.
  /// @param destinationChainSelector The destination chain selector.
  /// @param amount The amount to transfer.
  /// @param maxFastTransferFee The maximum allowable fee deducted when computing
  /// the sourceAmountNetFee, expressed in lock/burn token units.
  /// @param receiver The receiver address.
  /// @param settlementFeeToken The token used to pay the settlement fee.
  /// @param extraArgs Extra arguments for the transfer.
  /// @return settlementId The fill request ID.
  function ccipSendToken(
    uint64 destinationChainSelector,
    uint256 amount,
    uint256 maxFastTransferFee,
    bytes calldata receiver,
    address settlementFeeToken,
    bytes calldata extraArgs
  ) external payable returns (bytes32 settlementId);

  /// @notice Fast fills a transfer using liquidity provider funds.
  /// @notice Fillers must ensure that the parameters provided here exactly match the
  /// values emitted in the `FastTransferRequested` event on the source chain.
  /// It is recommended that the fillId must be passed directly from the event to avoid any mismatch
  /// as encoding differences can cause divergence between fast fill and slow path settlement IDs.
  /// @param settlementId The settlement ID, which under normal circumstances is the same as the CCIP message ID.
  /// @param fillId The fill ID, computed from the fill request parameters.
  /// @param sourceChainSelector The source chain selector.
  /// @param sourceAmountNetFee The amount being filled, excluding the fast fill fee, expressed in source token decimals.
  /// @param sourceDecimals The decimals of the token on the source token.
  /// @param receiver The receiver on the destination chain.
  function fastFill(
    bytes32 settlementId,
    bytes32 fillId,
    uint64 sourceChainSelector,
    uint256 sourceAmountNetFee,
    uint8 sourceDecimals,
    address receiver
  ) external;

  /// @notice Helper function to generate fill ID from request parameters.
  /// @param settlementId The settlement ID, which under normal circumstances is the same as the CCIP message ID.
  /// @param sourceChainSelector The chain selector where the fill request originated.
  /// @param sourceAmountNetFee The amount being filled, excluding the fast fill fee, expressed in source token decimals.
  /// @param sourceDecimals The decimals of the token on the source token.
  /// @param receiver The receiver on the destination chain. ABI encoded in the case of an EVM destination chain.
  /// @return fillId The computed fill ID.
  function computeFillId(
    bytes32 settlementId,
    uint64 sourceChainSelector,
    uint256 sourceAmountNetFee,
    uint8 sourceDecimals,
    bytes memory receiver
  ) external pure returns (bytes32);

  /// @notice Gets the accumulated pool fees that can be withdrawn.
  /// @return The amount of accumulated pool fees.
  function getAccumulatedPoolFees() external view returns (uint256);

  /// @notice Withdraws all accumulated pool fees to the specified recipient.
  /// @param recipient The address to receive the withdrawn fees.
  function withdrawPoolFees(
    address recipient
  ) external;
}
