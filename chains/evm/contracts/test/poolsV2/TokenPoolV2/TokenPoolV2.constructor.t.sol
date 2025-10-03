// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {TokenPoolV2Setup} from "./TokenPoolV2Setup.t.sol";

contract TokenPoolV2_constructor is TokenPoolV2Setup {
  function test_constructor() public view {
    assertEq(address(s_token), address(s_tokenPool.getToken()));
    assertEq(address(s_mockRMNRemote), s_tokenPool.getRmnProxy());
    assertFalse(s_tokenPool.getAllowListEnabled());
    assertEq(address(s_sourceRouter), s_tokenPool.getRouter());
    assertEq(DEFAULT_TOKEN_DECIMALS, s_tokenPool.getTokenDecimals());
  }
}
