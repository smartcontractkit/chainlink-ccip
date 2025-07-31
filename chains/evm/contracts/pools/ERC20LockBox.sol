// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {TokenAdminRegistry} from "../tokenAdminRegistry/TokenAdminRegistry.sol";

import {Ownable2StepMsgSender} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2StepMsgSender.sol";
import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";
import {IERC20} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/utils/SafeERC20.sol";

/// @title ERC20 Lock Box.
/// @notice A contract that holds ERC20 tokens for a token pool to simplify pool upgrades without requiring a manual
/// liquidity migration. If a token pool is being modified, the token pool administrator can simply set the new token pool
/// in the token admin registry, and the tokens will be automatically allowed to be withdrawn by the new token pool on
/// incoming messages.
/// @dev This contract is designed to support any ERC20-token permissionlessly, as any token pool can use it as
/// storage for their token liquidity. As a result many different tokens will be stored in this contract, but can
/// only be withdrawn by the associated token pool as defined in the token admin registry.
contract ERC20LockBox is Ownable2StepMsgSender, ITypeAndVersion {
  using SafeERC20 for IERC20;

  error Unauthorized(address caller);
  error InsufficientBalance(uint64 remoteChainSelector, uint256 requested, uint256 available);
  error TokenAmountCannotBeZero();
  error RecipientCannotBeZeroAddress();
  error TokenAddressCannotBeZero();
  error ZeroAddressNotAllowed();

  event AllowedCallerUpdated(address indexed token, address indexed caller, bool allowed);
  event Deposit(address indexed token, uint64 indexed remoteChainSelector, address indexed depositor, uint256 amount);
  event Withdrawal(
    address indexed token, uint64 indexed remoteChainSelector, address indexed recipient, uint256 amount
  );

  struct AllowedCallerConfigArgs {
    address token;
    address caller;
    bool allowed;
  }

  /// @notice The token admin registry is used to determine if the caller is the administrator of the token
  /// or the token pool.
  TokenAdminRegistry public immutable i_tokenAdminRegistry;

  /// @notice The lockbox allows for multiple authorized callers for a token. This allows support for
  /// liquidity providers to manage tokens without requiring the token pool owner to interact directly.
  /// Only the owner in the token admin registry can configure allowed callers for a token, and
  /// once a caller is allowed, they can call deposit and withdraw functions for the given token.
  mapping(address token => mapping(address caller => bool isAllowed)) internal s_allowedCallers;

  /// @notice The balance of available tokens for a given chain selector. This is a security measure to ensure that
  /// tokens which may be siloed by a specific chain selector are not available to other chains. If a token does not
  /// need to be siloed, then the chain selector 0 can be used to track the total balance of the token instead.
  mapping(address token => mapping(uint64 remoteChainSelector => uint256 balance)) internal s_tokenBalances;

  string public constant typeAndVersion = "ERC20LockBox 1.0.0-dev";

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
  /// @param remoteChainSelector The chain selector for which to deposit tokens. 0 is a valid chain selector. If token
  /// balances do not need to be Siloed by chain selector, then this is the recommended value.
  /// @dev If a user sends tokens to the lockbox directly, without calling this function, then the tokens
  /// will be unrecoverable since the internal s_tokenBalances mapping is will not be updated. It is NOT advised
  /// to send tokens to the lockbox directly, and instead tokens should be deposited via a liquidity provider or
  /// provideLiquidity function in the token pool itself.
  function deposit(address token, uint256 amount, uint64 remoteChainSelector) external {
    _validateDepositWithdraw(token, amount);

    IERC20(token).safeTransferFrom(msg.sender, address(this), amount);

    s_tokenBalances[token][remoteChainSelector] += amount;

    emit Deposit(token, remoteChainSelector, msg.sender, amount);
  }

  /// @notice Withdraws tokens for a specific remote chain selector.
  /// @param token The address of the ERC20 token to withdraw.
  /// @param amount The amount of tokens to withdraw.
  /// @param recipient The address that will receive the withdrawn tokens.
  /// @param remoteChainSelector The chain selector for which to withdraw tokens.
  function withdraw(address token, uint256 amount, address recipient, uint64 remoteChainSelector) external {
    _validateDepositWithdraw(token, amount);

    if (recipient == address(0)) {
      revert RecipientCannotBeZeroAddress();
    }

    // An explicit balance check returns a more useful error message than a silent underflow with no useful information.
    uint256 balance = s_tokenBalances[token][remoteChainSelector];
    if (balance < amount) {
      revert InsufficientBalance(remoteChainSelector, amount, balance);
    }

    // Decrease the balance for the specified chain selector.
    s_tokenBalances[token][remoteChainSelector] -= amount;

    IERC20(token).safeTransfer(recipient, amount);

    emit Withdrawal(token, remoteChainSelector, recipient, amount);
  }

  /// @notice Configures the allowed callers for deposit and withdraw functions.
  /// @dev Only the administrator of the token in the token admin registry can configure allowed callers.
  /// @dev Can add or remove multiple callers in a single transaction.
  /// @param configArgs Array of configuration arguments for allowed callers.
  function configureAllowedCallers(
    AllowedCallerConfigArgs[] calldata configArgs
  ) external {
    for (uint256 i = 0; i < configArgs.length; ++i) {
      address token = configArgs[i].token;
      if (token == address(0)) {
        revert TokenAddressCannotBeZero();
      }

      // Only the token administrator set in the token admin registry can modify the allowed callers.
      // Note: This gives the token pool administrator full control over the tokens in the lockbox, as they
      // will be able to determine who can call withdraw() at any time for the given token and amount.
      // Special care should be taken to ensure that the token pool administrator is not malicious or compromised.
      if (!i_tokenAdminRegistry.isAdministrator(token, msg.sender)) {
        revert Unauthorized(msg.sender);
      }

      address caller = configArgs[i].caller;
      bool allowed = configArgs[i].allowed;

      if (s_allowedCallers[token][caller] != allowed) {
        s_allowedCallers[token][caller] = allowed;
        emit AllowedCallerUpdated(token, caller, allowed);
      }
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
  /// @dev This gives the token pool administrator full control over the tokens in the lockbox, as they
  /// will be able to call withdraw() at any time for the given token and amount. Special care should be taken
  /// to ensure that the token pool administrator is not malicious or compromised.
  function isAllowedCaller(address token, address caller) public view returns (bool allowed) {
    TokenAdminRegistry.TokenConfig memory tokenConfig = i_tokenAdminRegistry.getTokenConfig(token);

    // The caller is allowed if they are the token pool, the administrator, or a designated allowed caller.
    return (caller == tokenConfig.tokenPool || caller == tokenConfig.administrator || s_allowedCallers[token][caller]);
  }

  /// @notice Get the balance for a specific token and remote chain selector.
  /// @param token The address of the ERC20 token.
  /// @param remoteChainSelector The remote chain selector to query.
  /// @return balance The balance of tokens for the specified token and remote chain selector.
  function getBalance(address token, uint64 remoteChainSelector) external view returns (uint256 balance) {
    return s_tokenBalances[token][remoteChainSelector];
  }
}
