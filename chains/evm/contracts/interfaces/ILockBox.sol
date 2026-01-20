// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

interface ILockBox {
  /// @notice Deposits the token into the lockbox.
  /// @param token The address of the token to deposit.
  /// @param remoteChainSelector The chain selector of the remote chain.
  /// @param amount The amount of tokens to deposit.
  function deposit(
    address token,
    uint64 remoteChainSelector,
    uint256 amount
  ) external;

  /// @notice Withdraws tokens to a specific recipient.
  /// @param token The address of the token to withdraw.
  /// @param remoteChainSelector The chain selector of the remote chain.
  /// @param amount The amount of tokens to withdraw. If set to max uint256, withdraws the entire balance.
  /// @param recipient The address of the recipient to receive the withdrawn tokens.
  function withdraw(
    address token,
    uint64 remoteChainSelector,
    uint256 amount,
    address recipient
  ) external;

  /// @notice Returns whether the lockbox supports a token.
  /// @param token The address of the token.
  /// @return supported True if the token is supported.
  function isTokenSupported(
    address token
  ) external view returns (bool);
}
