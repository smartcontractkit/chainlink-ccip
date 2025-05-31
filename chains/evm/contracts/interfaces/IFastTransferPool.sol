// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.0;

/// @title IFastTransferPool
/// @notice Interface for the CCIP Fast-Transfer Pool
interface IFastTransferPool {
  /// @notice Quote struct containing fee information
  struct Quote {
    uint256 ccipSettlementFee; // Fee paid to for CCIP settlement in CCIP supported fee tokens.
    uint256 fastTransferFee; // Fee paid to the fast transfer filler in the same asset as requested.
  }

  error AlreadyFilled(bytes32 fillRequestId);
  error AlreadySettled(bytes32 fillRequestId);

  /// @notice Emitted when a fast transfer is requested
  event FastTransferRequested(
    bytes32 indexed fillRequestId,
    uint64 indexed destinationChainSelector,
    uint256 amount,
    uint256 fastTransferFee,
    bytes receiver
  );
  /// @notice Emitted when a fast transfer is filled. This means the end user has received the tokens but the slow
  /// transfer is still in progress.
  event FastTransferFilled(
    bytes32 indexed fillRequestId, bytes32 indexed fillId, address indexed filler, uint256 destAmount, address receiver
  );
  /// @notice Emitted when a fast transfer is settled. This means the slow transfer has completed and the filler has
  /// received their fast transfer tokens and fee.
  event FastTransferSettled(bytes32 indexed fillRequestId);

  /// @notice Gets the CCIP send token fee and fast transfer fee for a given transfer
  /// @param settlementFeeToken The token used to pay the CCIP settlement fee
  /// @param destinationChainSelector The destination chain selector
  /// @param amount The amount to transfer
  /// @param receiver The receiver address
  /// @param extraArgs Extra arguments for the transfer
  /// @return Quote containing the CCIP fee and fast transfer fee
  function getCcipSendTokenFee(
    address settlementFeeToken,
    uint64 destinationChainSelector,
    uint256 amount,
    bytes calldata receiver,
    bytes calldata extraArgs
  ) external view returns (Quote memory);

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
  ) external payable returns (bytes32 fillRequestId);

  /// @notice Fast fills a transfer using liquidity provider funds
  /// @param fillRequestId The fill request ID
  /// @param fillId The fill ID, computed from the fill request parameters
  /// @param sourceChainSelector The source chain selector
  /// @param srcAmount The amount to fill
  /// @param sourceDecimals The decimals of the source token
  /// @param receiver The receiver address
  function fastFill(
    bytes32 fillRequestId,
    bytes32 fillId,
    uint64 sourceChainSelector,
    uint256 srcAmount,
    uint8 sourceDecimals,
    address receiver
  ) external;

  /// @notice Helper function to generate fill ID from request parameters
  /// @param fillRequestId The original fill request ID
  /// @param amount The amount being filled
  /// @param receiver The receiver address
  /// @return fillId The computed fill ID
  function computeFillId(
    bytes32 fillRequestId,
    uint256 amount,
    uint8 decimals,
    address receiver
  ) external pure returns (bytes32);
}
