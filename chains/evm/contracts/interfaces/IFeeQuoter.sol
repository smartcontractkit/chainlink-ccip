// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

import {Internal} from "../libraries/Internal.sol";

interface IFeeQuoter {
  /// @notice Get the list of fee tokens.
  /// @return feeTokens The tokens set as fee tokens.
  function getFeeTokens() external view returns (address[] memory);

  /// @notice Get the `tokenPrice` for a given token.
  /// @param token The token to get the price for.
  /// @return tokenPrice The tokenPrice for the given token.
  function getTokenPrice(
    address token
  ) external view returns (Internal.TimestampedPackedUint224 memory);

  /// @notice Get the `tokenPrice` for a given token, checks if the price is valid.
  /// @param token The token to get the price for.
  /// @return tokenPrice The tokenPrice for the given token if it exists and is valid.
  function getValidatedTokenPrice(
    address token
  ) external view returns (uint224);

  /// @notice Get the `tokenPrice` for an array of tokens.
  /// @param tokens The tokens to get prices for.
  /// @return tokenPrices The tokenPrices for the given tokens.
  function getTokenPrices(
    address[] calldata tokens
  ) external view returns (Internal.TimestampedPackedUint224[] memory);

  /// @notice Update the price for given tokens and gas prices for given chains.
  /// @param priceUpdates The price updates to apply.
  function updatePrices(
    Internal.PriceUpdates memory priceUpdates
  ) external;

  /// @notice Get an encoded `gasPrice` for a given destination chain ID.
  /// The 224-bit result encodes necessary gas price components.
  /// On L1 chains like Ethereum or Avax, the only component is the gas price.
  /// On Optimistic Rollups, there are two components - the L2 gas price, and L1 base fee for data availability.
  /// On future chains, there could be more or differing price components.
  /// @param destChainSelector The destination chain to get the price for.
  /// @return gasPrice The encoded gasPrice for the given destination chain ID.
  function getDestinationChainGasPrice(
    uint64 destChainSelector
  ) external view returns (Internal.TimestampedPackedUint224 memory);

  // ================================================================
  // │                 Not needed for new 1.7 chains                │
  // ================================================================

  /// @notice Gets the resolved token transfer fee components for a token transfer.
  /// @dev This function will check token-specific config first, then fall back to destination chain defaults.
  /// @param destChainSelector The destination chain selector.
  /// @param token The token address.
  /// @return feeUSDCents The fee in USD cents (multiples of 0.01 USD).
  /// @return destGasOverhead The gas charged to execute the token transfer on the destination chain.
  /// @return destBytesOverhead The bytes overhead for the token transfer on the destination chain.
  function getTokenTransferFee(
    uint64 destChainSelector,
    address token
  ) external view returns (uint32 feeUSDCents, uint32 destGasOverhead, uint32 destBytesOverhead);

  /// @notice Quotes the total gas and gas cost in USD cents.
  /// @param destChainSelector The destination chain selector.
  /// @param nonCalldataGas The non-calldata gas to be used for the message.
  /// @param calldataSize The size of the calldata in bytes.
  /// @return totalGas The total gas needed for the message.
  /// @return gasCostInUsdCents The gas cost in USD cents, taking into account the calldata cost as well.
  function quoteGasForExec(
    uint64 destChainSelector,
    uint32 nonCalldataGas,
    uint32 calldataSize
  ) external view returns (uint32 totalGas, uint256 gasCostInUsdCents);

  /// @notice Resolves legacy extra args for backward compatibility. Only has to support EVM, SVM, Aptos and SUI chain
  /// families as all future families have to use the new extraArgs format.
  /// @param destChainSelector The destination chain selector.
  /// @param extraArgs The extra args bytes.
  /// @return tokenReceiver The token receiver address encoded as bytes. Always length 32 or 0.
  /// @return gasLimit The gas limit to use for the message.
  /// @return executorArgs The executor args encoded as bytes. These are transformed into the new format.
  function resolveLegacyArgs(
    uint64 destChainSelector,
    bytes calldata extraArgs
  ) external view returns (bytes memory tokenReceiver, uint32 gasLimit, bytes memory executorArgs);
}
