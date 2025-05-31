// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IBurnMintERC20} from "@chainlink/contracts/src/v0.8/shared/token/ERC20/IBurnMintERC20.sol";

import {TokenPool} from "./TokenPool.sol";

abstract contract BurnMintTokenPoolAbstract is TokenPool {
  /// @notice Contains the specific release or mint token logic for a pool.
  /// @dev overriding this method allows us to create pools with different release/mint signatures
  /// without duplicating the underlying logic.
  function _releaseOrMint(address receiver, uint256 amount) internal virtual override {
    IBurnMintERC20(address(i_token)).mint(receiver, amount);
  }
}
