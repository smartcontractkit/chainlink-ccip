// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IRouter} from "../../../interfaces/IRouter.sol";
import {CCVAggregator} from "../../../offRamp/CCVAggregator.sol";
import {CCVAggregatorSetup} from "./CCVAggregatorSetup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract CCVAggregator_applySourceChainConfigUpdates is CCVAggregatorSetup {
  function test_applySourceChainConfigUpdates_multipleChains() public {
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
      laneMandatedCCVs: new address[](2)
    });
    configs[1].defaultCCV[0] = makeAddr("ccv2");
    configs[1].laneMandatedCCVs[0] = makeAddr("mandatedCCV1");
    configs[1].laneMandatedCCVs[1] = makeAddr("mandatedCCV2");

    vm.expectEmit();
    emit CCVAggregator.SourceChainConfigSet(
      chain1,
      CCVAggregator.SourceChainConfig({
        router: configs[0].router,
        isEnabled: configs[0].isEnabled,
        onRamp: configs[0].onRamp,
        defaultCCVs: configs[0].defaultCCV,
        laneMandatedCCVs: configs[0].laneMandatedCCVs
      })
    );

    s_agg.applySourceChainConfigUpdates(configs);

    CCVAggregator.SourceChainConfig memory config1 = s_agg.getSourceChainConfig(chain1);
    CCVAggregator.SourceChainConfig memory config2 = s_agg.getSourceChainConfig(chain2);

    assertEq(address(config1.router), address(configs[0].router));
    assertEq(address(config2.router), address(configs[1].router));

    assertEq(config1.isEnabled, configs[0].isEnabled);
    assertEq(config2.isEnabled, configs[1].isEnabled);

    assertEq(chain1, configs[0].sourceChainSelector);
    assertEq(chain2, configs[1].sourceChainSelector);
    assertEq(config1.onRamp, configs[0].onRamp);
    assertEq(config2.onRamp, configs[1].onRamp);

    assertEq(config1.defaultCCVs[0], configs[0].defaultCCV[0]);
    assertEq(config2.defaultCCVs[0], configs[1].defaultCCV[0]);

    assertEq(config1.laneMandatedCCVs.length, 0);
    assertEq(config2.laneMandatedCCVs.length, 2);
    assertEq(config2.laneMandatedCCVs[0], configs[1].laneMandatedCCVs[0]);
    assertEq(config2.laneMandatedCCVs[1], configs[1].laneMandatedCCVs[1]);
  }

  function test_applySourceChainConfigUpdates_updateExistingChain() public {
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
    assertEq(config.defaultCCVs.length, 2);
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

  function test_applySourceChainConfigUpdates_RevertWhen_OnlyCallableByOwner() public {
    vm.stopPrank();
    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);

    s_agg.applySourceChainConfigUpdates(new CCVAggregator.SourceChainConfigArgs[](0));
  }
}
