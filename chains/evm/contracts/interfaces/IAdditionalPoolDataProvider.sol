// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/// @notice Gives the additional pool data of a token on send, and checks it on receive. The token issuer defines what
/// the data means, for example an ID that is the same on every chain. A pool that supports this feature sends it in the
/// pool data.
/// @dev Both functions are called as views, so a provider cannot change state.
interface IAdditionalPoolDataProvider {
  /// @notice Returns the additional pool data of a local token for a send.
  /// @param localToken The local token that is being locked or burned.
  /// @param remoteChainSelector The destination chain.
  /// @return additionalPoolData The additional pool data to validate on the destination chain.
  function getAdditionalPoolData(
    address localToken,
    uint64 remoteChainSelector
  ) external view returns (bytes memory additionalPoolData);

  /// @notice Returns whether the additional pool data from the source chain is valid for the local token.
  /// @param localToken The local token that is being released or minted.
  /// @param remoteChainSelector The source chain.
  /// @param additionalPoolData The additional pool data from the source chain. Empty if none was provided.
  /// @return isValid True if the data is valid for the local token, false otherwise.
  function isAdditionalPoolDataValid(
    address localToken,
    uint64 remoteChainSelector,
    bytes calldata additionalPoolData
  ) external view returns (bool isValid);
}
