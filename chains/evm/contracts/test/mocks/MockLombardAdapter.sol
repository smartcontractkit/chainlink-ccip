// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {IERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/IERC20.sol";

contract MockLombardAdapter {
  address internal immutable i_lombardBridge;
  address internal immutable i_underlyingToken;

  constructor(
    address lombardBridge,
    address underlyingToken
  ) {
    i_underlyingToken = underlyingToken;
    i_lombardBridge = lombardBridge;
  }

  function transferFrom(
    address from,
    address to,
    uint256 amount
  ) external returns (bool) {
    IERC20(i_underlyingToken).transferFrom(from, to, amount);
    return true;
  }
}
