// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {CCVAggregator} from "../../../offRamp/CCVAggregator.sol";
import {CCVAggregatorHelper, CCVAggregatorSetup} from "./CCVAggregatorSetup.t.sol";

contract CCVAggregator_getAllSourceChainConfigs is CCVAggregatorSetup {
  function test_getAllSourceChainConfigs_ReturnsSingleChain() public {
    (uint64[] memory selectors, CCVAggregator.SourceChainConfig[] memory configs) = s_agg.getAllSourceChainConfigs();

    assertEq(selectors.length, 1);
    assertEq(configs.length, 1);

    assertEq(selectors[0], SOURCE_CHAIN_SELECTOR);
    assertEq(address(configs[0].router), address(s_sourceRouter));
    assertEq(configs[0].isEnabled, true);
    assertEq(configs[0].defaultCCV.length, 1);
    assertEq(configs[0].defaultCCV[0], s_defaultCCV);
  }

  function test_getAllSourceChainConfigs_ReturnsMultipleChains() public {
    // Add a second source chain
    uint64 chain2 = SOURCE_CHAIN_SELECTOR + 1;
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

    (uint64[] memory selectors, CCVAggregator.SourceChainConfig[] memory chainConfigs) =
      s_agg.getAllSourceChainConfigs();

    assertEq(selectors.length, 2);
    assertEq(chainConfigs.length, 2);

    // Check first chain
    assertEq(selectors[0], SOURCE_CHAIN_SELECTOR);
    assertEq(chainConfigs[0].defaultCCV[0], s_defaultCCV);

    // Check second chain
    assertEq(selectors[1], chain2);
    assertEq(chainConfigs[1].defaultCCV[0], makeAddr("ccv2"));
  }

  function test_getAllSourceChainConfigs_ReturnsEmptyArrays_WhenNoChainsConfigured() public {
    // Create a new aggregator without any source chain configs
    CCVAggregator.StaticConfig memory staticConfig = CCVAggregator.StaticConfig({
      localChainSelector: DEST_CHAIN_SELECTOR,
      gasForCallExactCheck: 5000,
      rmnRemote: s_mockRMNRemote,
      tokenAdminRegistry: address(s_tokenAdminRegistry)
    });

    CCVAggregator newAgg = new CCVAggregator(staticConfig);

    (uint64[] memory selectors, CCVAggregator.SourceChainConfig[] memory configs) = newAgg.getAllSourceChainConfigs();

    assertEq(selectors.length, 0);
    assertEq(configs.length, 0);
  }

  function test_getAllSourceChainConfigs_ReturnsCorrectOrder() public {
    // Add multiple chains in different order
    uint64 chain2 = SOURCE_CHAIN_SELECTOR + 1;
    uint64 chain3 = SOURCE_CHAIN_SELECTOR + 2;

    CCVAggregator.SourceChainConfigArgs[] memory configs = new CCVAggregator.SourceChainConfigArgs[](2);
    configs[0] = CCVAggregator.SourceChainConfigArgs({
      router: s_sourceRouter,
      sourceChainSelector: chain3,
      isEnabled: true,
      onRamp: abi.encode(makeAddr("onRamp3")),
      defaultCCV: new address[](1),
      laneMandatedCCVs: new address[](0)
    });
    configs[0].defaultCCV[0] = makeAddr("ccv3");

    configs[1] = CCVAggregator.SourceChainConfigArgs({
      router: s_sourceRouter,
      sourceChainSelector: chain2,
      isEnabled: true,
      onRamp: abi.encode(makeAddr("onRamp2")),
      defaultCCV: new address[](1),
      laneMandatedCCVs: new address[](0)
    });
    configs[1].defaultCCV[0] = makeAddr("ccv2");

    s_agg.applySourceChainConfigUpdates(configs);

    (uint64[] memory selectors, CCVAggregator.SourceChainConfig[] memory chainConfigs) =
      s_agg.getAllSourceChainConfigs();

    assertEq(selectors.length, 3);
    assertEq(chainConfigs.length, 3);

    // Should maintain the order they were added
    assertEq(selectors[0], SOURCE_CHAIN_SELECTOR);
    assertEq(selectors[1], chain3);
    assertEq(selectors[2], chain2);
  }

  function test_getAllSourceChainConfigs_HandlesDisabledChains() public {
    // Disable the first chain
    CCVAggregator.SourceChainConfigArgs[] memory configs = new CCVAggregator.SourceChainConfigArgs[](1);
    configs[0] = CCVAggregator.SourceChainConfigArgs({
      router: s_sourceRouter,
      sourceChainSelector: SOURCE_CHAIN_SELECTOR,
      isEnabled: false,
      onRamp: abi.encode(makeAddr("onRamp")),
      defaultCCV: new address[](1),
      laneMandatedCCVs: new address[](0)
    });
    configs[0].defaultCCV[0] = s_defaultCCV;

    s_agg.applySourceChainConfigUpdates(configs);

    (uint64[] memory selectors, CCVAggregator.SourceChainConfig[] memory chainConfigs) =
      s_agg.getAllSourceChainConfigs();

    assertEq(selectors.length, 1);
    assertEq(chainConfigs.length, 1);
    assertEq(chainConfigs[0].isEnabled, false);
  }
}
