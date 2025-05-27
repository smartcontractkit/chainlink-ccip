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
  error LaneDisabled();

  event FastFillRequest(
    bytes32 indexed fillRequestId,
    uint64 indexed dstChainSelector,
    uint256 amount,
    uint256 fastTransferFee,
    bytes receiver
  );
  event FastFillSettled(bytes32 indexed fillRequestId);
  event FastFill(
    bytes32 indexed fillRequestId, bytes32 indexed fillId, address indexed filler, uint256 destAmount, address receiver
  );
  event InvalidFill(
    bytes32 indexed fillRequestId, address indexed filler, uint256 filledAmount, uint256 expectedAmount
  );

  /// @notice Gets the CCIP send token fee and fast transfer fee for a given transfer
  /// @param feeToken The token used to pay the CCIP fee
  /// @param destinationChainSelector The destination chain selector
  /// @param amount The amount to transfer
  /// @param receiver The receiver address
  /// @param extraArgs Extra arguments for the transfer
  /// @return Quote containing the CCIP fee and fast transfer fee
  function getCcipSendTokenFee(
    address feeToken,
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
  /// @param sourceChainSelector The source chain selector
  /// @param srcAmount The amount to fill
  /// @param srcDecimals The decimals of the source token
  /// @param receiver The receiver address
  function fastFill(
    bytes32 fillRequestId,
    uint64 sourceChainSelector,
    uint256 srcAmount,
    uint8 srcDecimals,
    address receiver
  ) external;
}
