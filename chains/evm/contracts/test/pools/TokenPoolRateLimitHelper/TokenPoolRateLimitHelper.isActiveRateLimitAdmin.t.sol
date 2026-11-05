// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {TokenPoolRateLimitHelperSetup} from "./TokenPoolRateLimitHelperSetup.t.sol";

contract TokenPoolRateLimitHelper_isActiveRateLimitAdmin is TokenPoolRateLimitHelperSetup {
  function test_isActiveRateLimitAdmin() public {
    assertTrue(s_helper.isActiveRateLimitAdmin());

    s_pool.setDynamicConfig(address(s_sourceRouter), STRANGER, address(0));
    assertFalse(s_helper.isActiveRateLimitAdmin());
  }
}
