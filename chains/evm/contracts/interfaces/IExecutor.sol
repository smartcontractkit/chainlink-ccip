// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface IExecutor {
  /// @notice Returns the allowed finality config according to the FinalityCodec encoding.
  function getAllowedFinalityConfig() external view returns (bytes4);

  /// @notice Validates whether or not the executor can process the message and returns the fee required to do so.
  /// @param destChainSelector The destination chain selector.
  /// @param requestedFinalityConfig The requested finality encoding for the message (see `FinalityCodec`).
  /// @param ccvAddresses Array of CCV addresses that will be used for the message.
  /// @param extraArgs Extra arguments for the executor.
  function getFee(
    uint64 destChainSelector,
    bytes4 requestedFinalityConfig,
    address[] memory ccvAddresses,
    bytes memory extraArgs,
    address feeToken
  ) external view returns (uint16 usdCents);
}
