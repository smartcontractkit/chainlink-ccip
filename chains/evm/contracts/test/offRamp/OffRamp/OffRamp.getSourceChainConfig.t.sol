// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {OffRamp} from "../../../offRamp/OffRamp.sol";
import {OffRampSetup} from "./OffRampSetup.t.sol";

contract OffRamp_getSourceChainConfig is OffRampSetup {
  function test_getSourceChainConfig_ReturnsCorrectConfig() public view {
    OffRamp.SourceChainConfig memory config = s_agg.getSourceChainConfig(SOURCE_CHAIN_SELECTOR);

    assertEq(address(config.router), address(s_sourceRouter));
    assertEq(config.isEnabled, true);
    assertEq(config.defaultCCVs.length, 1);
    assertEq(config.defaultCCVs[0], s_defaultCCV);
    assertEq(config.laneMandatedCCVs.length, 0);
  }
}
