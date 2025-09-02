// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {CCVAggregator} from "../../../offRamp/CCVAggregator.sol";
import {CCVAggregatorSetup} from "./CCVAggregatorSetup.t.sol";

contract CCVAggregator_getAllSourceChainConfigs is CCVAggregatorSetup {
  function test_getAllSourceChainConfigs_ReturnsSingleChain() public view {
    (uint64[] memory selectors, CCVAggregator.SourceChainConfig[] memory configs) = s_agg.getAllSourceChainConfigs();

    assertEq(selectors.length, 1);
    assertEq(configs.length, 1);

    assertEq(selectors[0], SOURCE_CHAIN_SELECTOR);
    assertEq(address(configs[0].router), address(s_sourceRouter));
    assertEq(configs[0].isEnabled, true);
    assertEq(configs[0].defaultCCVs.length, 1);
    assertEq(configs[0].defaultCCVs[0], s_defaultCCV);
  }

  function test_getAllSourceChainConfigs_ReturnsMultipleChains() public {
    // Add a second source chain.
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

    // Check first chain.
    assertEq(selectors[0], SOURCE_CHAIN_SELECTOR);
    assertEq(chainConfigs[0].defaultCCVs[0], s_defaultCCV);

    // Check second chain.
    assertEq(selectors[1], chain2);
    assertEq(chainConfigs[1].defaultCCVs[0], makeAddr("ccv2"));
  }
}
