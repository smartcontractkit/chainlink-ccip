// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {BaseOnRampSetup} from "./BaseOnRampSetup.t.sol";

contract BaseOnRamp_getDestChainConfig is BaseOnRampSetup {
  function test_getDestChainConfig() public view {
    // Get config for the default destination chain set in setup
    (bool allowlistEnabled, address ccvProxy, address[] memory allowedSenders) =
      s_baseOnRamp.getDestChainConfig(DEST_CHAIN_SELECTOR);

    assertEq(ccvProxy, s_ccvProxy);
    assertFalse(allowlistEnabled);
    assertEq(allowedSenders.length, 0);
  }
}
