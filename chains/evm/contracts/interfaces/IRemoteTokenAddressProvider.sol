// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/// @notice Gives the remote token address of each token of a connection policy.
interface IRemoteTokenAddressProvider {
  /// @notice Returns the remote token address of a local token on a remote chain.
  /// @dev A wrong answer on send burns the amount, and the message can never succeed. For a token or chain that this
  /// provider does not know, it must revert or return empty bytes. DO NOT return abi.encode(address(0)) as default.
  /// On receive, the source token must equal the return value byte for byte.
  /// @param localToken The local token.
  /// @param remoteChainSelector The remote chain.
  /// @return remoteToken The token on the remote chain. For an EVM chain, this is abi.encode(address).
  function getRemoteTokenAddress(
    address localToken,
    uint64 remoteChainSelector
  ) external view returns (bytes memory remoteToken);
}
