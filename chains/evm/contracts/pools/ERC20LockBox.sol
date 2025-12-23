// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ILockBox} from "../interfaces/ILockBox.sol";
import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";

import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

import {IERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/utils/SafeERC20.sol";

/// @title ERC20 Lock Box
/// @notice Per-token/per-domain lockbox that holds ERC20 liquidity so pools can be upgraded without migrating
/// funds.
/// @dev Each lockbox is bound to one token and one liquidity domain id. This implementation supports only a single token to be deposited in the lockbox.
/// Only the owner can manage the allowlist; allowed callers (or the owner) can deposit/withdraw after the domain check.
contract ERC20LockBox is ITypeAndVersion, ILockBox, AuthorizedCallers {
  using SafeERC20 for IERC20;

  error InsufficientBalance(uint256 requested, uint256 available);
  error TokenAmountCannotBeZero();
  error RecipientCannotBeZeroAddress();
  error UnsupportedLiquidityDomain(bytes32 liquidityDomainId);
  error UnsupportedToken(address token);

  event Deposit(address indexed token, address indexed depositor, uint256 amount);
  event Withdrawal(address indexed token, address indexed recipient, uint256 amount);

  /// @notice The token supported by this lockbox.
  IERC20 internal immutable i_token;
  /// @notice User defined liquidity domain this lockbox is bound to.
  /// @dev This value helps pools route deposits/withdrawals to the correct lockbox for a given domain and prevents
  /// mismatched calls by enforcing the expected domain id on each operation.
  bytes32 internal immutable i_liquidityDomainId;

  string public constant typeAndVersion = "ERC20LockBox 1.7.0-dev";

  constructor(
    address token,
    bytes32 liquidityDomainId
  ) AuthorizedCallers(new address[](0)) {
    if (token == address(0)) {
      revert ZeroAddressNotAllowed();
    }

    i_token = IERC20(token);
    i_liquidityDomainId = liquidityDomainId;
  }

  /// @notice Deposits tokens into this contract. This eases the process of migrating tokens
  /// from a legacy token pool to a new one, since only the allowedCaller needs to be changed. Without it, the tokens
  /// would need to be manually withdrawn and re-deposited into the new token pool from a legacy pool, which is a
  /// time-consuming and error-prone process.
  /// @inheritdoc ILockBox
  function deposit(
    uint64, // remoteChainSelector
    address token,
    bytes32 liquidityDomainId,
    uint256 amount
  ) external {
    _validateDepositWithdraw(token, liquidityDomainId, amount);

    IERC20(token).safeTransferFrom(msg.sender, address(this), amount);

    emit Deposit(token, msg.sender, amount);
  }

  /// @inheritdoc ILockBox
  function withdraw(
    uint64, // remoteChainSelector
    address token,
    bytes32 liquidityDomainId,
    uint256 amount,
    address recipient
  ) external {
    _validateDepositWithdraw(token, liquidityDomainId, amount);

    if (recipient == address(0)) {
      revert RecipientCannotBeZeroAddress();
    }

    uint256 balance = IERC20(token).balanceOf(address(this));
    if (amount > balance) {
      revert InsufficientBalance(amount, balance);
    }

    IERC20(token).safeTransfer(recipient, amount);

    emit Withdrawal(token, recipient, amount);
  }

  /// @notice Validates the deposit and withdraw functions.
  /// @param token The token being deposited/withdrawn.
  /// @param amount The amount of tokens to deposit or withdraw.
  function _validateDepositWithdraw(
    address token,
    bytes32 liquidityDomainId,
    uint256 amount
  ) internal view {
    if (amount == 0) {
      revert TokenAmountCannotBeZero();
    }
    if (token != address(i_token)) {
      revert UnsupportedToken(token);
    }
    if (liquidityDomainId != i_liquidityDomainId) {
      revert UnsupportedLiquidityDomain(liquidityDomainId);
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

  /// @notice Gets the liquidity domain id this lockbox is bound to.
  /// @return liquidityDomainId The liquidity domain id.
  function getLiquidityDomainId() external view returns (bytes32) {
    return i_liquidityDomainId;
  }
}
