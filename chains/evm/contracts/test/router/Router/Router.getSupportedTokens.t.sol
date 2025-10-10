// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {CCVProxy} from "../../../onRamp/CCVProxy.sol";
import {RouterSetup} from "../../onRamp/OnRamp/RouterSetup.t.sol";

contract Router_getSupportedTokens is RouterSetup {
  function test_RevertWhen_GetSupportedTokens() public {
    vm.expectRevert(CCVProxy.GetSupportedTokensFunctionalityRemovedCheckAdminRegistry.selector);
    s_onRamp.getSupportedTokens(DEST_CHAIN_SELECTOR);
  }
}
