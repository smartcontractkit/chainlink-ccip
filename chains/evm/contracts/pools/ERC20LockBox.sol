// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {TokenAdminRegistry} from "../tokenAdminRegistry/TokenAdminRegistry.sol";

import {Ownable2StepMsgSender} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2StepMsgSender.sol";
import {IERC20} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/utils/SafeERC20.sol";
import {EnumerableSet} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v5.0.2/contracts/utils/structs/EnumerableSet.sol";

/// @title ERC20 Lock Box
/// @notice A contract that holds ERC20 tokens for a token pool to simplify pool upgrades without requiring a liquidity migration.
contract ERC20LockBox is Ownable2StepMsgSender {
  using SafeERC20 for IERC20;
  using EnumerableSet for EnumerableSet.AddressSet;

  error Unauthorized(address caller);
  error InsufficientBalance(uint64 remoteChainSelector, uint256 requested, uint256 available);
  error TokenAmountCannotBeZero();
  error RecipientCannotBeZeroAddress();
  error TokenAddressCannotBeZero();
  error TokenAdminRegistryCannotBeZeroAddress();

  event AllowedCallerAdded(address indexed caller);
  event AllowedCallerRemoved(address indexed caller);
  event Deposit(uint64 indexed remoteChainSelector, address indexed depositor, uint256 amount);
  event Withdrawal(uint64 indexed remoteChainSelector, address indexed recipient, uint256 amount);

  struct AllowedCallerConfigArgs {
    address token;
    address caller;
    bool allowed;
  }

  TokenAdminRegistry public immutable i_tokenAdminRegistry;

  mapping(address => mapping(address => bool)) public s_allowedCallers;

  mapping(address => mapping(uint64 => uint256)) public s_tokenBalances;

  constructor(
    address tokenAdminRegistry
  ) {
    if (tokenAdminRegistry == address(0)) {
      revert TokenAdminRegistryCannotBeZeroAddress();
    }
    i_tokenAdminRegistry = TokenAdminRegistry(tokenAdminRegistry);
  }

  /// @notice Deposits tokens for a specific remote chain selector
  /// @param token The address of the ERC20 token to deposit
  /// @param amount The amount of tokens to deposit
  /// @param remoteChainSelector The chain selector for which to deposit tokens
  function deposit(address token, uint256 amount, uint64 remoteChainSelector) external {
    // Validate the token address. It must be checked before the caller is authorized otherwise the allowed caller check
    // will revert.
    if (token == address(0)) {
      revert TokenAddressCannotBeZero();
    }

    // Note: The caller may be a token pool or a liquidity provider.
    if (!isAllowedCaller(token, msg.sender)) {
      revert Unauthorized(msg.sender);
    }

    if (amount == 0) {
      revert TokenAmountCannotBeZero();
    }

    // Transfer tokens from the caller to this contract
    IERC20(token).safeTransferFrom(msg.sender, address(this), amount);

    // Increase the balance for the specified chain selector
    s_tokenBalances[token][remoteChainSelector] += amount;

    emit Deposit(remoteChainSelector, msg.sender, amount);
  }

  /// @notice Withdraws tokens for a specific remote chain selector
  /// @param token The address of the ERC20 token to withdraw
  /// @param amount The amount of tokens to withdraw
  /// @param recipient The address that will receive the withdrawn tokens
  /// @param remoteChainSelector The chain selector for which to withdraw tokens
  function withdraw(address token, uint256 amount, address recipient, uint64 remoteChainSelector) external {
    // Validate the token address. It must be checked before the caller is authorized otherwise the allowed caller check
    // will revert.
    if (token == address(0)) {
      revert TokenAddressCannotBeZero();
    }

    // Check if the caller is authorized
    if (!isAllowedCaller(token, msg.sender)) {
      revert Unauthorized(msg.sender);
    }

    if (recipient == address(0)) {
      revert RecipientCannotBeZeroAddress();
    }

    if (amount == 0) {
      revert TokenAmountCannotBeZero();
    }

    // Check if sufficient balance exists for the chain selector
    if (s_tokenBalances[token][remoteChainSelector] < amount) {
      revert InsufficientBalance(remoteChainSelector, amount, s_tokenBalances[token][remoteChainSelector]);
    }

    // Decrease the balance for the specified chain selector
    s_tokenBalances[token][remoteChainSelector] -= amount;

    // Transfer tokens from this contract to the recipient
    IERC20(token).safeTransfer(recipient, amount);

    // Emit the withdrawal event
    emit Withdrawal(remoteChainSelector, recipient, amount);
  }

  /// @notice Configures the allowed callers for deposit and withdraw functions
  /// @dev Only the owner can configure allowed callers
  /// @dev Can add or remove multiple callers in a single transaction
  /// @param configArgs Array of configuration arguments for allowed callers
  function configureAllowedCallers(
    AllowedCallerConfigArgs[] calldata configArgs
  ) external {
    // Iterate through all configuration arguments
    for (uint256 i = 0; i < configArgs.length; ++i) {
      // Validate the token address
      address token = configArgs[i].token;
      if (token == address(0)) {
        revert TokenAddressCannotBeZero();
      }

      // Only the owner of the token pool can configure allowed callers
      address tokenPool = i_tokenAdminRegistry.getPool(token);
      address poolOwner = Ownable2StepMsgSender(tokenPool).owner();
      if (msg.sender != poolOwner) {
        revert Unauthorized(msg.sender);
      }

      address caller = configArgs[i].caller;
      bool allowed = configArgs[i].allowed;

      if (allowed) {
        // Add the caller to the allowed set
        if (!s_allowedCallers[token][caller]) {
          s_allowedCallers[token][caller] = true;
          emit AllowedCallerAdded(caller);
        }
      } else {
        // Remove the caller from the allowed set
        if (s_allowedCallers[token][caller]) {
          delete s_allowedCallers[token][caller];
          emit AllowedCallerRemoved(caller);
        }
      }
    }
  }

  /// @notice Checks if an address is allowed to call deposit and withdraw functions
  /// @param token The address of the ERC20 token
  /// @param caller The address to check
  /// @return allowed True if the address is allowed, false otherwise
  function isAllowedCaller(address token, address caller) public view returns (bool allowed) {
    return (msg.sender == i_tokenAdminRegistry.getPool(token) || s_allowedCallers[token][caller]);
  }

  /// @notice Get the balance for a specific token and remote chain selector
  /// @param token The address of the ERC20 token
  /// @param remoteChainSelector The remote chain selector to query
  /// @return balance The balance of tokens for the specified token and remote chain selector
  function getBalance(address token, uint64 remoteChainSelector) external view returns (uint256 balance) {
    return s_tokenBalances[token][remoteChainSelector];
  }
}
