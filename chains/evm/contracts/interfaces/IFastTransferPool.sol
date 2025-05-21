// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.0;

/// @title IFastTransferPool
/// @notice Interface for the CCIP Fast-Transfer Pool
interface IFastTransferPool {
  /// @notice Quote struct containing fee information
  struct Quote {
    uint256 sendTokenFee; // paid in feeToken
    uint256 fastTransferFee; // paid in asset
  }

  error AlreadyFilled(bytes32 fillRequestId);
  error MessageAlreadySettled(bytes32 fillRequestId);
  error InvalidLaneConfig();

  event FastFillRequest(
    bytes32 indexed fillRequestId,
    uint64 indexed dstChainSelector,
    uint256 amount,
    uint256 fastTransferFee,
    bytes receiver
  );
  event FastFillSettled(bytes32 indexed fillRequestId);
  event FastFill(bytes32 indexed fillRequestId, address indexed filler, uint256 amount, address indexed receiver);
  event InvalidFill(
    bytes32 indexed fillRequestId, address indexed filler, uint256 filledAmount, uint256 expectedAmount
  );

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
    address receiver,
    bytes calldata extraArgs
  ) external view returns (Quote memory);

  /// @notice Sends tokens via CCIP with optional fast transfer
  /// @param feeToken The token used to pay the CCIP fee
  /// @param dstChainSelector The destination chain selector
  /// @param amount The amount to transfer
  /// @param receiver The receiver address
  /// @param extraArgs Extra arguments for the transfer
  /// @return fillRequestId The fill request ID
  function ccipSendToken(
    address feeToken,
    uint64 dstChainSelector,
    uint256 amount,
    address receiver,
    bytes calldata extraArgs
  ) external returns (bytes32 fillRequestId);

  /// @notice Fast fills a transfer using liquidity provider funds
  /// @param messageId The CCIP message ID
  /// @param amount The amount to fill
  /// @param receiver The receiver address
  function fastFill(bytes32 messageId, uint256 amount, address receiver) external;
}
