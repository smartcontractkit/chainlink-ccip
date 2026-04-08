// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {BaseVerifier} from "../../../../ccvs/components/BaseVerifier.sol";
import {BaseVerifierSetup} from "./BaseVerifierSetup.t.sol";

contract BaseVerifier_getRemoteChainConfig is BaseVerifierSetup {
  function test_getRemoteChainConfig() public view {
    // Get config for the default destination chain set in setup.
    (BaseVerifier.RemoteChainConfigArgs memory config, address[] memory allowedSenders) =
      s_baseVerifier.getRemoteChainConfig(DEST_CHAIN_SELECTOR);

    assertEq(address(config.router), address(s_router));
    assertFalse(config.allowlistEnabled);
    assertEq(allowedSenders.length, 0);
  }
}
