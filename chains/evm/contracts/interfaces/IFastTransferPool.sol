// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.4;

/// @title IFastTransferPool
/// @notice Interface for the CCIP Fast-Transfer Pool
interface IFastTransferPool {
  /// @notice Quote struct containing fee information
  struct Quote {
    uint256 ccipSettlementFee; // Fee paid to for CCIP settlement in CCIP supported fee tokens.
    uint256 fastTransferFee; // Fee paid to the fast transfer filler in the same asset as requested.
  }

  error AlreadyFilled(bytes32 fillId);
  error AlreadySettled(bytes32 fillId);

  /// @notice Emitted when a fast transfer is requested
  event FastTransferRequested(
    uint64 indexed destinationChainSelector,
    bytes32 indexed fillId,
    bytes32 indexed settlementId,
    /// @param sourceAmountNetFee The amount being transferred, excluding the fast fill fee, expressed in source token decimals.
    uint256 sourceAmountNetFee,
    uint256 fastTransferFee,
    bytes receiver
  );
  /// @notice Emitted when a fast transfer is filled. This means the end user has received the tokens but the slow
  /// transfer is still in progress.
  event FastTransferFilled(
    bytes32 indexed fillId, bytes32 indexed settlementId, address indexed filler, uint256 destAmount, address receiver
  );
  /// @notice Emitted when a fast transfer is settled. This means the slow transfer has completed and the filler has
  /// received their fast transfer tokens and fee.
  event FastTransferSettled(bytes32 indexed fillId, bytes32 indexed settlementId);

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
  /// @param receiver The receiver address.
  /// @param settlementFeeToken The token used to pay the settlement fee.
  /// @param extraArgs Extra arguments for the transfer.
  /// @return settlementId The fill request ID.
  function ccipSendToken(
    uint64 destinationChainSelector,
    uint256 amount,
    bytes calldata receiver,
    address settlementFeeToken,
    bytes calldata extraArgs
  ) external payable returns (bytes32 settlementId);

  /// @notice Fast fills a transfer using liquidity provider funds
  /// @param settlementId The settlement ID, which under normal circumstances is the same as the CCIP message ID.
  /// @param fillId The fill ID, computed from the fill request parameters.
  /// @param sourceChainSelector The source chain selector.
  /// @param sourceAmountNetFee The amount being filled, excluding the fast fill fee, expressed in source token decimals.
  /// @param sourceDecimals The decimals of the token on the source token.
  /// @param receiver The receiver on the destination chain. ABI encoded in the case of an EVM destination chain.
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
  /// @param sourceAmountNetFee The amount being filled, excluding the fast fill fee, expressed in source token decimals.
  /// @param sourceDecimals The decimals of the token on the source token.
  /// @param receiver The receiver on the destination chain. ABI encoded in the case of an EVM destination chain.
  /// @return fillId The computed fill ID.
  function computeFillId(
    bytes32 settlementId,
    uint256 sourceAmountNetFee,
    uint8 sourceDecimals,
    bytes memory receiver
  ) external pure returns (bytes32);
}
