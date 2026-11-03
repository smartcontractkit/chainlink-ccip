// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/// @notice Resolves remote token addresses for tokens assigned to a connection policy.
interface IRemoteTokenAddressProvider {
  /// @notice Returns the remote token address of a local token on a remote chain.
  /// @dev An incorrect destination token can cause loss of funds. Revert or return empty bytes for unsupported tokens
  /// or chains. On receive, the source token in the pool data must match the returned address byte for byte.
  /// @param localToken The local token.
  /// @param remoteChainSelector The remote chain selector.
  /// @return remoteToken The token on the remote chain. For an EVM chain, this is abi.encode(address).
  function getRemoteTokenAddress(
    address localToken,
    uint64 remoteChainSelector
  ) external view returns (bytes memory remoteToken);
}
