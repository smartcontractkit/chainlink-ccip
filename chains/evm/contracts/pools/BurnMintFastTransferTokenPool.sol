// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IBurnMintERC20} from "@chainlink/contracts/src/v0.8/shared/token/ERC20/IBurnMintERC20.sol";

import {FastTransferTokenPoolAbstract} from "./FastTransferTokenPoolAbstract.sol";

import {IERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/utils/SafeERC20.sol";

/// @notice A token pool that supports burn-mint operations and fast transfers
contract BurnMintFastTransferTokenPool is FastTransferTokenPoolAbstract {
  using SafeERC20 for IERC20;

  string public constant override typeAndVersion = "BurnMintFastTransferTokenPool 1.6.3-dev";

  constructor(
    IBurnMintERC20 token,
    uint8 localTokenDecimals,
    address[] memory allowlist,
    address rmnProxy,
    address router,
    uint64 sourceChainSelector
  ) FastTransferTokenPoolAbstract(token, localTokenDecimals, allowlist, rmnProxy, router, sourceChainSelector) {}

  /// @notice Handles the locking or burning of tokens for both fast and slow transfers. Regardless of the transfer
  /// type, all the tokens are always burned.
  function _lockOrBurn(
    uint256 amount
  ) internal virtual override {
    IBurnMintERC20(address(i_token)).burn(amount);
  }

  /// @notice Handles the release or minting of tokens for both fast and slow transfers.
  /// @param receiver The address that will receive the tokens.
  /// In the case of a fast transfer this will depend on the fill status.
  /// - NOT_FILLED - the receiver is the end user.
  /// - FILLED - the receiver is the filler.
  /// @param amount The amount is always the entire amount, including the fee. That means the fee will go back to the
  /// requester of the transfer is the transfer status was NOT_FILLED, or to the filler if the status was FILLED.
  function _releaseOrMint(address receiver, uint256 amount) internal virtual override {
    IBurnMintERC20(address(i_token)).mint(receiver, amount);
  }

  /// @notice Returns the accumulated pool fees
  /// @return The total accumulated pool fees, which is the balance of the token in the pool contract.
  function getAccumulatedPoolFees() public view override returns (uint256) {
    return getToken().balanceOf(address(this));
  }
}
