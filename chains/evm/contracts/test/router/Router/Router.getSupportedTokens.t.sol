// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {CCVProxy} from "../../../onRamp/CCVProxy.sol";
import {OnRampSetup} from "../../onRamp/OnRamp/OnRampSetup.t.sol";

contract Router_getSupportedTokens is OnRampSetup {
  function test_RevertWhen_GetSupportedTokens() public {
    vm.expectRevert(CCVProxy.GetSupportedTokensFunctionalityRemovedCheckAdminRegistry.selector);
    s_onRamp.getSupportedTokens(DEST_CHAIN_SELECTOR);
  }
}
