// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {CCVAggregator} from "../../../offRamp/CCVAggregator.sol";
import {CCVAggregatorHelper, CCVAggregatorSetup} from "./CCVAggregatorSetup.t.sol";

contract CCVAggregator_getSourceChainConfig is CCVAggregatorSetup {
  function test_getSourceChainConfig_ReturnsCorrectConfig() public view {
    CCVAggregator.SourceChainConfig memory config = s_agg.getSourceChainConfig(SOURCE_CHAIN_SELECTOR);

    assertEq(address(config.router), address(s_sourceRouter));
    assertEq(config.isEnabled, true);
    assertEq(config.defaultCCV.length, 1);
    assertEq(config.defaultCCV[0], s_defaultCCV);
    assertEq(config.laneMandatedCCVs.length, 0);
  }

  function test_getSourceChainConfig_ReturnsEmptyConfig_WhenChainNotConfigured() public view {
    uint64 unconfiguredChain = SOURCE_CHAIN_SELECTOR + 1;
    CCVAggregator.SourceChainConfig memory config = s_agg.getSourceChainConfig(unconfiguredChain);

    assertEq(address(config.router), address(0));
    assertEq(config.isEnabled, false);
    assertEq(config.defaultCCV.length, 0);
    assertEq(config.laneMandatedCCVs.length, 0);
  }

  function test_getSourceChainConfig_ReturnsUpdatedConfig_AfterConfigurationUpdate() public {
    // Update the source chain configuration
    CCVAggregator.SourceChainConfigArgs[] memory configs = new CCVAggregator.SourceChainConfigArgs[](1);
    configs[0] = CCVAggregator.SourceChainConfigArgs({
      router: s_sourceRouter,
      sourceChainSelector: SOURCE_CHAIN_SELECTOR,
      isEnabled: false,
      onRamp: abi.encode(makeAddr("onRamp")),
      defaultCCV: new address[](2),
      laneMandatedCCVs: new address[](1)
    });
    configs[0].defaultCCV[0] = makeAddr("ccv1");
    configs[0].defaultCCV[1] = makeAddr("ccv2");
    configs[0].laneMandatedCCVs[0] = makeAddr("mandatedCCV");

    s_agg.applySourceChainConfigUpdates(configs);

    CCVAggregator.SourceChainConfig memory config = s_agg.getSourceChainConfig(SOURCE_CHAIN_SELECTOR);

    assertEq(address(config.router), address(s_sourceRouter));
    assertEq(config.isEnabled, false);
    assertEq(config.defaultCCV.length, 2);
    assertEq(config.defaultCCV[0], makeAddr("ccv1"));
    assertEq(config.defaultCCV[1], makeAddr("ccv2"));
    assertEq(config.laneMandatedCCVs.length, 1);
    assertEq(config.laneMandatedCCVs[0], makeAddr("mandatedCCV"));
  }

  function test_getSourceChainConfig_HandlesMultipleChains() public {
    uint64 chain1 = SOURCE_CHAIN_SELECTOR;
    uint64 chain2 = SOURCE_CHAIN_SELECTOR + 1;

    // Configure second chain
    CCVAggregator.SourceChainConfigArgs[] memory configs = new CCVAggregator.SourceChainConfigArgs[](1);
    configs[0] = CCVAggregator.SourceChainConfigArgs({
      router: s_sourceRouter,
      sourceChainSelector: chain2,
      isEnabled: true,
      onRamp: abi.encode(makeAddr("onRamp2")),
      defaultCCV: new address[](1),
      laneMandatedCCVs: new address[](0)
    });
    configs[0].defaultCCV[0] = makeAddr("ccv2");

    s_agg.applySourceChainConfigUpdates(configs);

    // Check both chains
    CCVAggregator.SourceChainConfig memory config1 = s_agg.getSourceChainConfig(chain1);
    CCVAggregator.SourceChainConfig memory config2 = s_agg.getSourceChainConfig(chain2);

    assertEq(config1.isEnabled, true);
    assertEq(config2.isEnabled, true);
    assertEq(config1.defaultCCV[0], s_defaultCCV);
    assertEq(config2.defaultCCV[0], makeAddr("ccv2"));
  }
}
