// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";

import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

import {IERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/utils/SafeERC20.sol";

/// @title ERC20 Lock Box.
/// @notice A per-token lockbox that holds ERC20 liquidity so pools can be upgraded without migrating funds.
/// @dev One token per lockbox. Only the owner can update the allowed caller list; allowed callers (or the owner) can
/// deposit and withdraw the supported token.
contract ERC20LockBox is ITypeAndVersion, AuthorizedCallers {
  using SafeERC20 for IERC20;

  error InsufficientBalance(uint256 requested, uint256 available);
  error TokenAmountCannotBeZero();
  error RecipientCannotBeZeroAddress();
  error UnsupportedChainSelector(uint64 chainSelector);

  event Deposit(address indexed token, address indexed depositor, uint256 amount);
  event Withdrawal(address indexed token, address indexed recipient, uint256 amount);

  /// @notice The token supported by this lockbox.
  IERC20 internal immutable i_token;
  uint64 internal immutable i_remoteChainSelector;

  string public constant typeAndVersion = "ERC20LockBox 1.7.0-dev";

  constructor(address token, uint64 remoteChainSelector) AuthorizedCallers(new address[](0)) {
    if (token == address(0)) {
      revert ZeroAddressNotAllowed();
    }

    i_token = IERC20(token);
    i_remoteChainSelector = remoteChainSelector;
  }

  /// @notice Deposits tokens into this contract. This eases the process of migrating tokens
  /// from a legacy token pool to a new one, since only the allowedCaller needs to be changed. Without it, the tokens
  /// would need to be manually withdrawn and re-deposited into the new token pool from a legacy pool, which is a
  /// time-consuming and error-prone process.
  /// @param amount The amount of tokens to deposit.
  /// @param remoteChainSelector The chain selector this lockbox instance is bound to (0 for unsiloed).
  /// @dev This function does NOT support storing native tokens, as the token pool which handles native is expected to
  /// have wrapped it into an ERC20-compatibletoken first.
  function deposit(uint64 remoteChainSelector, uint256 amount) external {
    _validateDepositWithdraw(remoteChainSelector, amount);

    i_token.safeTransferFrom(msg.sender, address(this), amount);

    emit Deposit(address(i_token), msg.sender, amount);
  }

  /// @notice Withdraws tokens to a specific recipient.
  /// @param amount The amount of tokens to withdraw.
  /// @param recipient The address that will receive the withdrawn tokens.
  /// @param remoteChainSelector The chain selector this lockbox instance is bound to (0 for unsiloed).
  function withdraw(uint64 remoteChainSelector, uint256 amount, address recipient) external {
    _validateDepositWithdraw(remoteChainSelector, amount);

    if (recipient == address(0)) {
      revert RecipientCannotBeZeroAddress();
    }

    uint256 balance = i_token.balanceOf(address(this));
    if (amount > balance) {
      revert InsufficientBalance(amount, balance);
    }

    i_token.safeTransfer(recipient, amount);

    emit Withdrawal(address(i_token), recipient, amount);
  }

  /// @notice Validates the deposit and withdraw functions.
  /// @param amount The amount of tokens to deposit or withdraw.
  function _validateDepositWithdraw(uint64 remoteChainSelector, uint256 amount) internal view {
    if (amount == 0) {
      revert TokenAmountCannotBeZero();
    }
    if (remoteChainSelector != i_remoteChainSelector) {
      revert UnsupportedChainSelector(remoteChainSelector);
    }
    if (msg.sender != owner()) {
      _validateCaller();
    }
  }

  /// @notice Gets the token supported by this lockbox.
  /// @return token The ERC20 token.
  function getToken() external view returns (IERC20 token) {
    return i_token;
  }

  /// @notice Gets the remote chain selector this lockbox is bound to.
  /// @return remoteChainSelector The remote chain selector.
  function getRemoteChainSelector() external view returns (uint64) {
    return i_remoteChainSelector;
  }
}
