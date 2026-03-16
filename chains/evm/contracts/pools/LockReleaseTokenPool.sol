// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ILockBox} from "../interfaces/ILockBox.sol";
import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";

import {TokenPool} from "./TokenPool.sol";

import {IERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/utils/SafeERC20.sol";

/// @notice Token pool used for tokens on their native chain. This uses a lock and release mechanism.
/// @dev One token per LockReleaseTokenPool.
contract LockReleaseTokenPool is TokenPool, ITypeAndVersion {
  using SafeERC20 for IERC20;

  function typeAndVersion() external pure virtual override returns (string memory) {
    return "LockReleaseTokenPool 2.0.0-dev";
  }

  /// @notice The lock box for the token pool.
  ILockBox internal immutable i_lockBox;

  constructor(
    IERC20 token,
    uint8 localTokenDecimals,
    address advancedPoolHooks,
    address rmnProxy,
    address router,
    address lockBox
  ) TokenPool(token, localTokenDecimals, advancedPoolHooks, rmnProxy, router) {
    if (lockBox == address(0)) revert ZeroAddressInvalid();

    ILockBox lockBoxContract = ILockBox(lockBox);
    if (!lockBoxContract.isTokenSupported(address(token))) {
      revert InvalidToken(address(token));
    }
    token.forceApprove(lockBox, type(uint256).max);
    i_lockBox = lockBoxContract;
  }

  /// @notice Gets the lock box address.
  function getLockBox() external view returns (address) {
    return address(i_lockBox);
  }

  /// @inheritdoc TokenPool
  /// @dev The router has already transferred the full amount to this contract before calling lockOrBurn.
  /// For V1 the amount = full amount. For V2 the amount = destTokenAmount (after fees), and fees remain on this contract.
  function _lockOrBurn(
    uint64 remoteChainSelector,
    uint256 amount
  ) internal override {
    i_lockBox.deposit(address(i_token), remoteChainSelector, amount);
  }

  /// @inheritdoc TokenPool
  /// @dev Releases tokens from the lock box to the receiver.
  function _releaseOrMint(
    address receiver,
    uint256 amount,
    uint64 remoteChainSelector
  ) internal override {
    i_lockBox.withdraw(address(i_token), remoteChainSelector, amount, receiver);
  }
}
