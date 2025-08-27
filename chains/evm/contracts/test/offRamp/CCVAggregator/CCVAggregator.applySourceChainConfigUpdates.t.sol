// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IRouter} from "../../../interfaces/IRouter.sol";
import {CCVAggregator} from "../../../offRamp/CCVAggregator.sol";
import {CCVAggregatorHelper, CCVAggregatorSetup} from "./CCVAggregatorSetup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract CCVAggregator_applySourceChainConfigUpdates is CCVAggregatorSetup {
  function test_applySourceChainConfigUpdates_Success_SingleChain() public {
    uint64 newChain = SOURCE_CHAIN_SELECTOR + 1;
    address newCCV = makeAddr("newCCV");

    CCVAggregator.SourceChainConfigArgs[] memory configs = new CCVAggregator.SourceChainConfigArgs[](1);
    configs[0] = CCVAggregator.SourceChainConfigArgs({
      router: s_sourceRouter,
      sourceChainSelector: newChain,
      isEnabled: true,
      onRamp: abi.encode(makeAddr("onRamp")),
      defaultCCV: new address[](1),
      laneMandatedCCVs: new address[](0)
    });
    configs[0].defaultCCV[0] = newCCV;

    s_agg.applySourceChainConfigUpdates(configs);

    CCVAggregator.SourceChainConfig memory config = s_agg.getSourceChainConfig(newChain);
    assertEq(address(config.router), address(s_sourceRouter));
    assertEq(config.isEnabled, true);
    assertEq(config.defaultCCV.length, 1);
    assertEq(config.defaultCCV[0], newCCV);
  }

  function test_applySourceChainConfigUpdates_Success_MultipleChains() public {
    uint64 chain1 = SOURCE_CHAIN_SELECTOR + 1;
    uint64 chain2 = SOURCE_CHAIN_SELECTOR + 2;

    CCVAggregator.SourceChainConfigArgs[] memory configs = new CCVAggregator.SourceChainConfigArgs[](2);
    configs[0] = CCVAggregator.SourceChainConfigArgs({
      router: s_sourceRouter,
      sourceChainSelector: chain1,
      isEnabled: true,
      onRamp: abi.encode(makeAddr("onRamp1")),
      defaultCCV: new address[](1),
      laneMandatedCCVs: new address[](0)
    });
    configs[0].defaultCCV[0] = makeAddr("ccv1");

    configs[1] = CCVAggregator.SourceChainConfigArgs({
      router: s_sourceRouter,
      sourceChainSelector: chain2,
      isEnabled: false,
      onRamp: abi.encode(makeAddr("onRamp2")),
      defaultCCV: new address[](1),
      laneMandatedCCVs: new address[](0)
    });
    configs[1].defaultCCV[0] = makeAddr("ccv2");

    s_agg.applySourceChainConfigUpdates(configs);

    CCVAggregator.SourceChainConfig memory config1 = s_agg.getSourceChainConfig(chain1);
    CCVAggregator.SourceChainConfig memory config2 = s_agg.getSourceChainConfig(chain2);

    assertEq(config1.isEnabled, true);
    assertEq(config2.isEnabled, false);
    assertEq(config1.defaultCCV[0], makeAddr("ccv1"));
    assertEq(config2.defaultCCV[0], makeAddr("ccv2"));
  }

  function test_applySourceChainConfigUpdates_Success_UpdateExistingChain() public {
    // Update existing chain configuration
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
    assertEq(config.isEnabled, false);
    assertEq(config.defaultCCV.length, 2);
    assertEq(config.laneMandatedCCVs.length, 1);
  }

  function test_applySourceChainConfigUpdates_RevertWhen_ZeroChainSelectorNotAllowed() public {
    CCVAggregator.SourceChainConfigArgs[] memory configs = new CCVAggregator.SourceChainConfigArgs[](1);
    configs[0] = CCVAggregator.SourceChainConfigArgs({
      router: s_sourceRouter,
      sourceChainSelector: 0,
      isEnabled: true,
      onRamp: abi.encode(makeAddr("onRamp")),
      defaultCCV: new address[](1),
      laneMandatedCCVs: new address[](0)
    });
    configs[0].defaultCCV[0] = makeAddr("ccv");

    vm.expectRevert(CCVAggregator.ZeroChainSelectorNotAllowed.selector);
    s_agg.applySourceChainConfigUpdates(configs);
  }

  function test_applySourceChainConfigUpdates_RevertWhen_ZeroAddressNotAllowed_Router() public {
    CCVAggregator.SourceChainConfigArgs[] memory configs = new CCVAggregator.SourceChainConfigArgs[](1);
    configs[0] = CCVAggregator.SourceChainConfigArgs({
      router: IRouter(address(0)),
      sourceChainSelector: SOURCE_CHAIN_SELECTOR + 1,
      isEnabled: true,
      onRamp: abi.encode(makeAddr("onRamp")),
      defaultCCV: new address[](1),
      laneMandatedCCVs: new address[](0)
    });
    configs[0].defaultCCV[0] = makeAddr("ccv");

    vm.expectRevert(CCVAggregator.ZeroAddressNotAllowed.selector);
    s_agg.applySourceChainConfigUpdates(configs);
  }

  function test_applySourceChainConfigUpdates_RevertWhen_ZeroAddressNotAllowed_DefaultCCV() public {
    CCVAggregator.SourceChainConfigArgs[] memory configs = new CCVAggregator.SourceChainConfigArgs[](1);
    configs[0] = CCVAggregator.SourceChainConfigArgs({
      router: s_sourceRouter,
      sourceChainSelector: SOURCE_CHAIN_SELECTOR + 1,
      isEnabled: true,
      onRamp: abi.encode(makeAddr("onRamp")),
      defaultCCV: new address[](0),
      laneMandatedCCVs: new address[](0)
    });

    vm.expectRevert(CCVAggregator.ZeroAddressNotAllowed.selector);
    s_agg.applySourceChainConfigUpdates(configs);
  }

  function test_applySourceChainConfigUpdates_RevertWhen_ZeroAddressNotAllowed_OnRamp() public {
    CCVAggregator.SourceChainConfigArgs[] memory configs = new CCVAggregator.SourceChainConfigArgs[](1);
    configs[0] = CCVAggregator.SourceChainConfigArgs({
      router: s_sourceRouter,
      sourceChainSelector: SOURCE_CHAIN_SELECTOR + 1,
      isEnabled: true,
      onRamp: "",
      defaultCCV: new address[](1),
      laneMandatedCCVs: new address[](0)
    });
    configs[0].defaultCCV[0] = makeAddr("ccv");

    vm.expectRevert(CCVAggregator.ZeroAddressNotAllowed.selector);
    s_agg.applySourceChainConfigUpdates(configs);
  }

  function test_applySourceChainConfigUpdates_EmitsSourceChainConfigSet() public {
    uint64 newChain = SOURCE_CHAIN_SELECTOR + 1;
    address newCCV = makeAddr("newCCV");

    CCVAggregator.SourceChainConfigArgs[] memory configs = new CCVAggregator.SourceChainConfigArgs[](1);
    configs[0] = CCVAggregator.SourceChainConfigArgs({
      router: s_sourceRouter,
      sourceChainSelector: newChain,
      isEnabled: true,
      onRamp: abi.encode(makeAddr("onRamp")),
      defaultCCV: new address[](1),
      laneMandatedCCVs: new address[](0)
    });
    configs[0].defaultCCV[0] = newCCV;

    vm.expectEmit();
    emit CCVAggregator.SourceChainConfigSet(
      newChain,
      CCVAggregator.SourceChainConfig({
        router: s_sourceRouter,
        isEnabled: true,
        defaultCCV: configs[0].defaultCCV,
        laneMandatedCCVs: new address[](0)
      })
    );

    s_agg.applySourceChainConfigUpdates(configs);
  }

  function test_applySourceChainConfigUpdates_OnlyCallableByOwner() public {
    uint64 newChain = SOURCE_CHAIN_SELECTOR + 1;

    CCVAggregator.SourceChainConfigArgs[] memory configs = new CCVAggregator.SourceChainConfigArgs[](1);
    configs[0] = CCVAggregator.SourceChainConfigArgs({
      router: s_sourceRouter,
      sourceChainSelector: newChain,
      isEnabled: true,
      onRamp: abi.encode(makeAddr("onRamp")),
      defaultCCV: new address[](1),
      laneMandatedCCVs: new address[](0)
    });
    configs[0].defaultCCV[0] = makeAddr("ccv");

    vm.stopPrank();
    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);

    s_agg.applySourceChainConfigUpdates(configs);
  }
}
