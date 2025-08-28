// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {CCVAggregator} from "../../../offRamp/CCVAggregator.sol";
import {CCVAggregatorSetup} from "./CCVAggregatorSetup.t.sol";

contract CCVAggregator_getSourceChainConfig is CCVAggregatorSetup {
  function test_getSourceChainConfig_ReturnsCorrectConfig() public view {
    CCVAggregator.SourceChainConfig memory config = s_agg.getSourceChainConfig(SOURCE_CHAIN_SELECTOR);

    assertEq(address(config.router), address(s_sourceRouter));
    assertEq(config.isEnabled, true);
    assertEq(config.defaultCCVs.length, 1);
    assertEq(config.defaultCCVs[0], s_defaultCCV);
    assertEq(config.laneMandatedCCVs.length, 0);
  }
}
