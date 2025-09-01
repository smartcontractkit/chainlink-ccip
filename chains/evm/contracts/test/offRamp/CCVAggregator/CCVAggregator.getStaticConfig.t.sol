// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {CCVAggregator} from "../../../offRamp/CCVAggregator.sol";
import {CCVAggregatorSetup} from "./CCVAggregatorSetup.t.sol";

contract CCVAggregator_getStaticConfig is CCVAggregatorSetup {
  function test_getStaticConfig_MatchesConstructorValues() public {
    CCVAggregator.StaticConfig memory newConfig = CCVAggregator.StaticConfig({
      localChainSelector: 999999,
      gasForCallExactCheck: 10000,
      rmnRemote: s_mockRMNRemote,
      tokenAdminRegistry: makeAddr("newRegistry")
    });

    CCVAggregator newAgg = new CCVAggregator(newConfig);
    CCVAggregator.StaticConfig memory returnedConfig = newAgg.getStaticConfig();

    assertEq(returnedConfig.localChainSelector, newConfig.localChainSelector);
    assertEq(returnedConfig.gasForCallExactCheck, newConfig.gasForCallExactCheck);
    assertEq(address(returnedConfig.rmnRemote), address(newConfig.rmnRemote));
    assertEq(returnedConfig.tokenAdminRegistry, newConfig.tokenAdminRegistry);
  }
}
