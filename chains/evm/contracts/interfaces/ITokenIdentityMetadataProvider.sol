// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/// @notice Gives the token identity metadata of a token on send, and checks it on receive. The default identity of a
/// token is its address. Token identity metadata is the extra identity data that the issuer defines, for example an ID that
/// is the same on every chain. A pool that supports this feature sends it in the pool data.
/// @dev Both functions are called as views, so a provider cannot change state.
interface ITokenIdentityMetadataProvider {
  /// @notice Returns the token identity metadata of a local token for a send.
  /// @param localToken The local token that is being locked or burned.
  /// @param remoteChainSelector The destination chain.
  /// @return tokenIdentityMetadata The token identity metadata to validate on the destination chain.
  function getTokenIdentityMetadata(
    address localToken,
    uint64 remoteChainSelector
  ) external view returns (bytes memory tokenIdentityMetadata);

  /// @notice Returns whether the token identity metadata from the source chain matches the local token.
  /// @param localToken The local token that is being released or minted.
  /// @param remoteChainSelector The source chain.
  /// @param tokenIdentityMetadata The token identity metadata from the source chain. Empty if none was provided.
  /// @return isValid True if the metadata matches the local token, false otherwise.
  function isTokenIdentityMetadataValid(
    address localToken,
    uint64 remoteChainSelector,
    bytes calldata tokenIdentityMetadata
  ) external view returns (bool isValid);
}
