// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {MultiTokenPoolSetup} from "./MultiTokenPoolSetup.t.sol";

contract MultiTokenPool_getTokenDecimals is MultiTokenPoolSetup {
  function test_getTokenDecimals() public view {
    assertEq(DEFAULT_TOKEN_DECIMALS, s_pool.getTokenDecimals(address(s_token)));
  }
}
