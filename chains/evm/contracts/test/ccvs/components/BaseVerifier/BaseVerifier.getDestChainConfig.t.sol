// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {BaseVerifierSetup} from "./BaseVerifierSetup.t.sol";

contract BaseVerifier_getRemoteChainConfig is BaseVerifierSetup {
  function test_getRemoteChainConfig() public view {
    // Get config for the default destination chain set in setup.
    (bool allowlistEnabled, address router, address[] memory allowedSenders) =
      s_baseVerifier.getRemoteChainConfig(DEST_CHAIN_SELECTOR);

    assertEq(router, address(s_router));
    assertFalse(allowlistEnabled);
    assertEq(allowedSenders.length, 0);
  }
}
