// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {OnRamp} from "../../../onRamp/OnRamp.sol";
import {RouterSetup} from "./RouterSetup.t.sol";

contract Router_getSupportedTokens is RouterSetup {
  function test_RevertWhen_GetSupportedTokens() public {
    vm.expectRevert(OnRamp.GetSupportedTokensFunctionalityRemovedCheckAdminRegistry.selector);
    s_onRamp.getSupportedTokens(DEST_CHAIN_SELECTOR);
  }
}
