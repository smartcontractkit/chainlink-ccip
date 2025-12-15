// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";

import {Ownable2StepMsgSender} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2StepMsgSender.sol";

import {IERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/utils/SafeERC20.sol";

/// @title ERC20 Lock Box.
/// @notice A per-token lockbox that holds ERC20 liquidity so pools can be upgraded without migrating funds.
/// @dev One token per lockbox. Only the owner can update the allowed caller list; allowed callers (or the owner) can
/// deposit and withdraw the supported token.
contract ERC20LockBox is ITypeAndVersion, Ownable2StepMsgSender {
  using SafeERC20 for IERC20;

  error Unauthorized(address caller);
  error InsufficientBalance(uint256 requested, uint256 available);
  error TokenAmountCannotBeZero();
  error RecipientCannotBeZeroAddress();
  error ZeroAddressNotAllowed();
  error UnsupportedChainSelector(uint64 chainSelector);

  event AllowedCallerUpdated(address indexed token, address indexed caller, bool allowed);
  event Deposit(address indexed token, address indexed depositor, uint256 amount);
  event Withdrawal(address indexed token, address indexed recipient, uint256 amount);

  struct AllowedCallerConfigArgs {
    address caller;
    bool allowed;
  }

  /// @notice The token supported by this lockbox.
  IERC20 public immutable i_token;
  uint64 public immutable i_remoteChainSelector;

  /// @notice Allowed callers that can deposit and withdraw.
  mapping(address caller => bool isAllowed) internal s_allowedCallers;

  string public constant typeAndVersion = "ERC20LockBox 1.7.0-dev";

  constructor(address token, uint64 remoteChainSelector) {
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

  /// @notice Configures the allowed callers for deposit and withdraw functions.
  /// @dev Only the owner can configure allowed callers.
  /// @dev Can add or remove multiple callers in a single transaction.
  /// @param configArgs Array of configuration arguments for allowed callers.
  function configureAllowedCallers(
    AllowedCallerConfigArgs[] calldata configArgs
  ) external virtual {
    _validateOwner();

    for (uint256 i = 0; i < configArgs.length; ++i) {
      address caller = configArgs[i].caller;
      if (caller == address(0)) {
        revert ZeroAddressNotAllowed();
      }
      bool allowed = configArgs[i].allowed;

      if (s_allowedCallers[caller] != allowed) {
        s_allowedCallers[caller] = allowed;
        emit AllowedCallerUpdated(address(i_token), caller, allowed);
      }
    }
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
    if (!isAllowedCaller(msg.sender)) {
      revert Unauthorized(msg.sender);
    }
  }

  /// @notice Checks if an address is allowed to call deposit and withdraw functions.
  /// @param caller The address to check.
  /// @return allowed True if the address is allowed, false otherwise.
  function isAllowedCaller(
    address caller
  ) public view returns (bool allowed) {
    return caller == owner() || s_allowedCallers[caller];
  }

  /// @notice Gets the token supported by this lockbox.
  /// @return token The ERC20 token.
  function getToken() external view returns (IERC20 token) {
    return i_token;
  }

  function _validateOwner() internal view {
    if (msg.sender != owner()) {
      revert Unauthorized(msg.sender);
    }
  }
}
