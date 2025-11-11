// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface IExecutor {
  /// @notice Returns the minimum number of block confirmations that's allowed to be requested. The actual waiting for
  /// the block confirmations is handled by the CCVs. This value is only here to gate the value a user can request from
  /// a verifier.
  function getMinBlockConfirmations() external view returns (uint16);

  /// @notice Validates whether or not the executor can process the message and returns the fee required to do so.
  /// @param destChainSelector The destination chain selector.
  /// @param requestedBlockDepth The requested block depth for finality.
  /// @param ccvAddresses Array of CCV addresses that will be used for the message.
  /// @param extraArgs Extra arguments for the executor.
  function getFee(
    uint64 destChainSelector,
    uint16 requestedBlockDepth,
    address[] memory ccvAddresses,
    bytes memory extraArgs
  ) external view returns (uint16 usdCents);
}
