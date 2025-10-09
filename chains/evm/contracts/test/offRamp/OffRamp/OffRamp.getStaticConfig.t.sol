// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {OffRamp} from "../../../offRamp/OffRamp.sol";
import {OffRampSetup} from "./OffRampSetup.t.sol";

contract OffRamp_getStaticConfig is OffRampSetup {
  function test_getStaticConfig_MatchesConstructorValues() public {
    OffRamp.StaticConfig memory newConfig = OffRamp.StaticConfig({
      localChainSelector: 999999,
      gasForCallExactCheck: 10000,
      rmnRemote: s_mockRMNRemote,
      tokenAdminRegistry: makeAddr("newRegistry")
    });

    OffRamp newAgg = new OffRamp(newConfig);
    OffRamp.StaticConfig memory returnedConfig = newAgg.getStaticConfig();

    assertEq(returnedConfig.localChainSelector, newConfig.localChainSelector);
    assertEq(returnedConfig.gasForCallExactCheck, newConfig.gasForCallExactCheck);
    assertEq(address(returnedConfig.rmnRemote), address(newConfig.rmnRemote));
    assertEq(returnedConfig.tokenAdminRegistry, newConfig.tokenAdminRegistry);
  }
}
