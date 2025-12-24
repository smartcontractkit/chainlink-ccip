// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/IERC20.sol";

/// @notice Minimal ERC20 interface with mint/burn extensions used across CCIP.
/// @dev Mirrors the Chainlink `IBurnMintERC20` interface but targets OpenZeppelin Contracts v5.3.0.
interface IBurnMintERC20 is IERC20 {
  /// @notice Mints new tokens for a given address.
  /// @param account The address to mint the new tokens to.
  /// @param amount The number of tokens to be minted.
  /// @dev This function increases the total supply.
  function mint(
    address account,
    uint256 amount
  ) external;

  /// @notice Burns tokens from the sender.
  /// @param amount The number of tokens to be burned.
  /// @dev This function decreases the total supply.
  function burn(
    uint256 amount
  ) external;

  /// @notice Burns tokens from a given address.
  /// @param account The address to burn tokens from.
  /// @param amount The number of tokens to be burned.
  /// @dev This function decreases the total supply.
  function burn(
    address account,
    uint256 amount
  ) external;

  /// @notice Burns tokens from a given address.
  /// @param account The address to burn tokens from.
  /// @param amount The number of tokens to be burned.
  /// @dev This function decreases the total supply.
  function burnFrom(
    address account,
    uint256 amount
  ) external;
}

