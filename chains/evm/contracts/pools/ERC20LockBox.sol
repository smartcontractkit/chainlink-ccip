// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IOwner} from "../interfaces/IOwner.sol";

import {TokenAdminRegistry} from "../tokenAdminRegistry/TokenAdminRegistry.sol";

import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";
import {IERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/utils/SafeERC20.sol";

/// @title ERC20 Lock Box.
/// @notice A contract that holds ERC20 tokens for a token pool to simplify pool upgrades without requiring a manual
/// liquidity migration. If a token pool is being modified, the token pool administrator can simply set the new token pool
/// in the token admin registry, and the tokens will be automatically allowed to be withdrawn by the new token pool on
/// incoming messages.
/// @dev This contract is designed to support ERC20-tokens permissionlessly, as any compatible token pool can use it as
/// storage for their token liquidity. As a result many different tokens will be stored in this contract, but can
/// only be withdrawn by the associated token pool as defined in the token admin registry or an allowed caller.
/// @dev Only token pools which implement IOwnable are supported. If a token pool uses an alternative access control
/// mechanism, such as RBAC, it will not be able to use this lockbox, and should instead use a custom implementation
/// specific to their access control mechanism.
contract ERC20LockBox is ITypeAndVersion {
  using SafeERC20 for IERC20;

  error Unauthorized(address caller);
  error InsufficientBalance(uint256 requested, uint256 available);
  error TokenAmountCannotBeZero();
  error RecipientCannotBeZeroAddress();
  error TokenAddressCannotBeZero();
  error ZeroAddressNotAllowed();

  event AllowedCallerUpdated(address indexed token, address indexed caller, bool allowed);
  event Deposit(address indexed token, address indexed depositor, uint256 amount);
  event Withdrawal(address indexed token, address indexed recipient, uint256 amount);

  struct AllowedCallerConfigArgs {
    address token;
    address caller;
    bool allowed;
  }

  /// @notice The token admin registry is used to determine if the caller is the administrator of the token
  /// or the token pool.
  TokenAdminRegistry public immutable i_tokenAdminRegistry;

  /// @notice The lockbox allows for multiple authorized callers for a token. This allows support for
  /// complex token pool designs, such as USDC, which uses a child pool to interact with the lockbox rather than
  /// the contract registered with the token admin registry. Without this, it would not be possible to support
  /// such designs as the contract which actually handles the tokens would not be able to interact with this contract.
  /// It is also necessary as it enables liquidity providers to handle tokens, but must be managed carefully to ensure
  /// that unauthorized entities are not configured as allowed callers, as they would be able to withdraw tokens
  /// without requiring the token pool owner's approval.
  mapping(address token => mapping(address caller => bool isAllowed)) internal s_allowedCallers;

  string public constant typeAndVersion = "ERC20LockBox 1.6.2-dev";

  constructor(
    address tokenAdminRegistry
  ) {
    if (tokenAdminRegistry == address(0)) {
      revert ZeroAddressNotAllowed();
    }
    i_tokenAdminRegistry = TokenAdminRegistry(tokenAdminRegistry);
  }

  /// @notice Deposits tokens for a specific remote chain selector. This eases the process of migrating tokens
  /// from a legacy token pool to a new one, since only the allowedCaller needs to be changed. Without it, the tokens
  /// would need to be manually withdrawn and re-deposited into the new token pool from a legacy pool, which is a
  /// time-consuming and error-prone process.
  /// @param token The address of the ERC20 token to deposit.
  /// @param amount The amount of tokens to deposit.
  /// @dev This function does NOT support storing native tokens, as the token pool which handles native is expected to
  /// have wrapped it into an ERC20-compatibletoken first.
  function deposit(address token, uint256 amount) external {
    _validateDepositWithdraw(token, amount);

    IERC20(token).safeTransferFrom(msg.sender, address(this), amount);

    emit Deposit(token, msg.sender, amount);
  }

  /// @notice Withdraws tokens for a specific remote chain selector.
  /// @param token The address of the ERC20 token to withdraw.
  /// @param amount The amount of tokens to withdraw.
  /// @param recipient The address that will receive the withdrawn tokens.
  function withdraw(address token, uint256 amount, address recipient) external {
    _validateDepositWithdraw(token, amount);

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

  /// @notice Configures the allowed callers for deposit and withdraw functions.
  /// @dev Only the administrator of the token in the token admin registry can configure allowed callers.
  /// @dev Can add or remove multiple callers in a single transaction.
  /// @param configArgs Array of configuration arguments for allowed callers.
  function configureAllowedCallers(
    AllowedCallerConfigArgs[] calldata configArgs
  ) external virtual {
    for (uint256 i = 0; i < configArgs.length; ++i) {
      address token = configArgs[i].token;
      if (token == address(0)) {
        revert TokenAddressCannotBeZero();
      }

      _validateCallerIsTokenPoolOwner(token);

      address caller = configArgs[i].caller;
      bool allowed = configArgs[i].allowed;

      if (s_allowedCallers[token][caller] != allowed) {
        // Allowing for external callers is critical to enabling more complex proxy-pool mechanisms such as USDC.
        // In these designs, the contract registered with the token admin registry may not be directly responsible
        // for handling tokens, and acts only as a proxy to another pool. Without allowing external callers,
        // it would not be possible, as only the proxy would be able to interact with this contract.
        s_allowedCallers[token][caller] = allowed;
        emit AllowedCallerUpdated(token, caller, allowed);
      }
    }
  }

  /// @notice Validates that the caller is the owner of the token pool for a given token.
  /// @param token The address of the ERC20 token.
  /// @dev This function is only configured to support token pools which implement IOwnable. If a token pool
  /// uses an alternative access control mechanism, such as RBAC, it will not be able to use this lockbox, and should
  /// instead use a custom implementation which overrides this function.
  function _validateCallerIsTokenPoolOwner(
    address token
  ) internal virtual {
    // Only the owner of the token pool itself, which MAY NOT be the administrator of the token in the token
    // admin registry, can configure allowed callers. Currently, the token pool owner manages liquidity providers,
    // who are allowed to withdraw liquidity at will. Since an allowed caller can do the same on this contract, limiting
    // who can manage that configuration ensures there are no additional trust assumptions for managing liquidity.
    if (msg.sender != IOwner(i_tokenAdminRegistry.getPool(token)).owner()) {
      revert Unauthorized(msg.sender);
    }
  }

  /// @notice Validates the deposit and withdraw functions.
  /// @param token The address of the ERC20 token.
  /// @param amount The amount of tokens to deposit or withdraw.
  function _validateDepositWithdraw(address token, uint256 amount) internal view {
    if (token == address(0)) {
      revert TokenAddressCannotBeZero();
    }

    if (amount == 0) {
      revert TokenAmountCannotBeZero();
    }

    if (!isAllowedCaller(token, msg.sender)) {
      revert Unauthorized(msg.sender);
    }
  }

  /// @notice Checks if an address is allowed to call deposit and withdraw functions.
  /// @param token The address of the ERC20 token.
  /// @param caller The address to check.
  /// @return allowed True if the address is allowed, false otherwise.
  function isAllowedCaller(address token, address caller) public view returns (bool allowed) {
    TokenAdminRegistry.TokenConfig memory tokenConfig = i_tokenAdminRegistry.getTokenConfig(token);

    // The caller is allowed if they are the token pool or a specially allowed caller.
    return (caller == tokenConfig.tokenPool || s_allowedCallers[token][caller]);
  }
}
