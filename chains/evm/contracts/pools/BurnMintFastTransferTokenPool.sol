// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IBurnMintERC20} from "@chainlink/contracts/src/v0.8/shared/token/ERC20/IBurnMintERC20.sol";

import {FastTransferTokenPoolAbstract} from "./FastTransferTokenPoolAbstract.sol";

import {IERC20} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/utils/SafeERC20.sol";

/// @title BurnMintFastTransferTokenPool
/// @notice A token pool that supports burn-mint operations and fast transfers
/// @dev Inherits from BurnMintTokenPoolAbstract and FastTransferTokenPoolAbstract
contract BurnMintFastTransferTokenPool is FastTransferTokenPoolAbstract {
  using SafeERC20 for IERC20;

  string public constant override typeAndVersion = "BurnMintFastTransferTokenPool 1.6.1";

  constructor(
    IBurnMintERC20 token,
    uint8 localTokenDecimals,
    address[] memory allowlist,
    address rmnProxy,
    address router
  ) FastTransferTokenPoolAbstract(token, localTokenDecimals, allowlist, rmnProxy, router) {}

  /// @notice Handles the transfer of tokens when a fast transfer is initiated
  function _handleFastTransferLockOrBurn(uint64, address, uint256 amount) internal override {
    // Since this is a fast transfer, the Router doesn't forward the tokens to the pool.
    i_token.safeTransferFrom(msg.sender, address(this), amount);
    // Use the normal burn logic once the tokens are in the pool.
    _lockOrBurn(amount);
  }

  /// @notice Handles the locking or burning of tokens for both fast and slow transfers.
  function _lockOrBurn(
    uint256 amount
  ) internal virtual override {
    IBurnMintERC20(address(i_token)).burn(amount);
  }

  /// @notice Handles the release or minting of tokens for both fast and slow transfers.
  /// @param receiver The address that will receive the tokens.
  /// In the case of a fast transfer this will depend on the fill status.
  /// - NOT_FILLED - the receiver is the end user
  /// - FILLED - the receiver is the filler
  function _releaseOrMint(address receiver, uint256 amount) internal virtual override {
    IBurnMintERC20(address(i_token)).mint(receiver, amount);
  }

  /// @inheritdoc FastTransferTokenPoolAbstract
  function _transferFromFiller(address filler, address receiver, uint256 amount) internal override {
    getToken().safeTransferFrom(filler, receiver, amount);
  }

  /// @inheritdoc FastTransferTokenPoolAbstract
  function _handleSlowFill(uint256 settlementAmountLocal, address receiver) internal override {
    IBurnMintERC20(address(i_token)).mint(receiver, settlementAmountLocal);
  }

  /// @inheritdoc FastTransferTokenPoolAbstract
  function _handleFastFilledReimbursement(address filler, uint256 settlementAmountLocal) internal override {
    // Honest filler -> pay them back + fee
    IBurnMintERC20(address(i_token)).mint(filler, settlementAmountLocal);
  }
}
