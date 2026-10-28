// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/// @notice Provides the additional pool data of a token in lockOrBurn, and checks it in releaseOrMint.
/// A pool that supports this feature sends it in the pool data.
/// @dev Both functions are called as views, so a provider cannot change state.
interface IAdditionalPoolDataProvider {
  /// @notice Returns the additional pool data for a lockOrBurn.
  /// @param localToken The local token that is being locked or burned.
  /// @param remoteChainSelector The destination chain.
  /// @param requestedFinalityConfig The requested finality, encoded with FinalityCodec.
  /// @param tokenArgs The token arguments of the message. The sender chooses them.
  /// @return additionalPoolData The additional pool data to validate on the destination chain.
  function getAdditionalPoolData(
    address localToken,
    uint64 remoteChainSelector,
    bytes4 requestedFinalityConfig,
    bytes calldata tokenArgs
  ) external view returns (bytes memory additionalPoolData);

  /// @notice Returns if the additional pool data from the source chain is valid for the local token on dest chain.
  /// @param localToken The local token that is being released or minted.
  /// @param remoteChainSelector The source chain.
  /// @param requestedFinalityConfig The requested finality, encoded with FinalityCodec.
  /// @param additionalPoolData The additional pool data from the source chain. Empty if none was provided.
  /// @return isValid True if the data is valid for the local token, false otherwise.
  function isAdditionalPoolDataValid(
    address localToken,
    uint64 remoteChainSelector,
    bytes4 requestedFinalityConfig,
    bytes calldata additionalPoolData
  ) external view returns (bool isValid);
}
