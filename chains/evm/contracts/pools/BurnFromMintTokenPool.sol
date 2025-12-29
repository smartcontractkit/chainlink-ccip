// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IBurnMintERC20} from "../interfaces/IBurnMintERC20.sol";
import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";

import {BurnMintTokenPoolAbstract} from "./BurnMintTokenPoolAbstract.sol";
import {TokenPool} from "./TokenPool.sol";

/// @notice This pool mints and burns a 3rd-party token.
/// @dev Pool whitelisting mode is set in the constructor and cannot be modified later.
/// It either accepts any address as originalSender, or only accepts whitelisted originalSender.
/// The only way to change whitelisting mode is to deploy a new pool.
/// If that is expected, please make sure the token's burner/minter roles are adjustable.
/// @dev This contract is a variant of BurnMintTokenPool that uses `burnFrom(from, amount)`.
contract BurnFromMintTokenPool is BurnMintTokenPoolAbstract, ITypeAndVersion {
  string public constant override typeAndVersion = "BurnFromMintTokenPool 1.7.0-dev";

  constructor(
    IBurnMintERC20 token,
    uint8 localTokenDecimals,
    address advancedPoolHooks,
    address rmnProxy,
    address router,
    address feeAggregator
  ) TokenPool(token, localTokenDecimals, advancedPoolHooks, rmnProxy, router, feeAggregator) {
    // Some tokens allow burning from the sender without approval, but not all do.
    // To be safe, we approve the pool to burn from the pool.
    token.approve(address(this), type(uint256).max);
  }

  /// @inheritdoc TokenPool
  /// @param amount The amount of tokens to burn from the pool.
  function _lockOrBurn(
    uint256 amount
  ) internal virtual override {
    IBurnMintERC20(address(i_token)).burnFrom(address(this), amount);
  }
}
