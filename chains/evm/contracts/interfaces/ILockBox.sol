// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

interface ILockBox {
  /// @notice Deposits the token into the lockbox.
  function deposit(
    uint64 remoteChainSelector,
    address token,
    bytes32 liquidityDomainId,
    uint256 amount
  ) external;

  /// @notice Withdraws tokens to a specific recipient.
  function withdraw(
    uint64 remoteChainSelector,
    address token,
    bytes32 liquidityDomainId,
    uint256 amount,
    address recipient
  ) external;
}
