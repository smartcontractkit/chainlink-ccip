// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IBurnMintERC20} from "../../interfaces/IBurnMintERC20.sol";

import {MultiTokenPool} from "./MultiTokenPool.sol";

import {IERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/IERC20.sol";

contract BurnMintMultiTokenPool is MultiTokenPool {
  constructor(
    IERC20[] memory tokens,
    uint8 localTokenDecimals,
    address rmnProxy,
    address router
  ) MultiTokenPool(tokens, localTokenDecimals, rmnProxy, router) {}

  function _lockOrBurn(
    address token,
    uint64,
    uint256 amount
  ) internal virtual override {
    IBurnMintERC20(token).burn(msg.sender, amount);
  }

  function _releaseOrMint(
    address token,
    address receiver,
    uint256 amount,
    uint64
  ) internal virtual override {
    IBurnMintERC20(token).mint(receiver, amount);
  }
}
