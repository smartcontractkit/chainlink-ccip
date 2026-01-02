// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IRouter} from "../../../interfaces/IRouter.sol";
import {OffRamp} from "../../../offRamp/OffRamp.sol";
import {OffRampSetup} from "./OffRampSetup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract OffRamp_applySourceChainConfigUpdates is OffRampSetup {
  function test_applySourceChainConfigUpdates_multipleChains() public {
    uint64 chain1 = SOURCE_CHAIN_SELECTOR + 1;
    uint64 chain2 = SOURCE_CHAIN_SELECTOR + 2;

    bytes[] memory onRamps1 = new bytes[](1);
    onRamps1[0] = abi.encode(makeAddr("onRamp1"));

    bytes[] memory onRamps2 = new bytes[](1);
    onRamps2[0] = abi.encode(makeAddr("onRamp2"));

    OffRamp.SourceChainConfigArgs[] memory configs = new OffRamp.SourceChainConfigArgs[](2);
    configs[0] = OffRamp.SourceChainConfigArgs({
      router: s_sourceRouter,
      sourceChainSelector: chain1,
      isEnabled: true,
      onRamps: onRamps1,
      defaultCCVs: new address[](1),
      laneMandatedCCVs: new address[](0)
    });
    configs[0].defaultCCVs[0] = makeAddr("ccv1");

    configs[1] = OffRamp.SourceChainConfigArgs({
      router: s_sourceRouter,
      sourceChainSelector: chain2,
      isEnabled: false,
      onRamps: onRamps2,
      defaultCCVs: new address[](1),
      laneMandatedCCVs: new address[](2)
    });
    configs[1].defaultCCVs[0] = makeAddr("ccv2");
    configs[1].laneMandatedCCVs[0] = makeAddr("mandatedCCV1");
    configs[1].laneMandatedCCVs[1] = makeAddr("mandatedCCV2");

    vm.expectEmit();
    emit OffRamp.SourceChainConfigSet(chain1, configs[0]);

    s_offRamp.applySourceChainConfigUpdates(configs);

    OffRamp.SourceChainConfig memory config1 = s_offRamp.getSourceChainConfig(chain1);
    OffRamp.SourceChainConfig memory config2 = s_offRamp.getSourceChainConfig(chain2);

    assertEq(address(config1.router), address(configs[0].router));
    assertEq(address(config2.router), address(configs[1].router));

    assertEq(config1.isEnabled, configs[0].isEnabled);
    assertEq(config2.isEnabled, configs[1].isEnabled);

    assertEq(chain1, configs[0].sourceChainSelector);
    assertEq(chain2, configs[1].sourceChainSelector);
    assertEq(config1.onRamps[0], configs[0].onRamps[0]);
    assertEq(config2.onRamps[0], configs[1].onRamps[0]);

    assertEq(config1.defaultCCVs[0], configs[0].defaultCCVs[0]);
    assertEq(config2.defaultCCVs[0], configs[1].defaultCCVs[0]);

    assertEq(config1.laneMandatedCCVs.length, 0);
    assertEq(config2.laneMandatedCCVs.length, 2);
    assertEq(config2.laneMandatedCCVs[0], configs[1].laneMandatedCCVs[0]);
    assertEq(config2.laneMandatedCCVs[1], configs[1].laneMandatedCCVs[1]);
  }

  function test_applySourceChainConfigUpdates_updateExistingChain() public {
    bytes[] memory onRamps = new bytes[](1);
    onRamps[0] = abi.encode(makeAddr("onRamp"));

    OffRamp.SourceChainConfigArgs[] memory configs = new OffRamp.SourceChainConfigArgs[](1);
    configs[0] = OffRamp.SourceChainConfigArgs({
      router: s_sourceRouter,
      sourceChainSelector: SOURCE_CHAIN_SELECTOR,
      isEnabled: false,
      onRamps: onRamps,
      defaultCCVs: new address[](2),
      laneMandatedCCVs: new address[](1)
    });
    configs[0].defaultCCVs[0] = makeAddr("ccv1");
    configs[0].defaultCCVs[1] = makeAddr("ccv2");
    configs[0].laneMandatedCCVs[0] = makeAddr("mandatedCCV");

    s_offRamp.applySourceChainConfigUpdates(configs);

    OffRamp.SourceChainConfig memory config = s_offRamp.getSourceChainConfig(SOURCE_CHAIN_SELECTOR);
    assertEq(config.isEnabled, false);
    assertEq(config.defaultCCVs.length, 2);
    assertEq(config.laneMandatedCCVs.length, 1);
  }

  function test_applySourceChainConfigUpdates_RevertWhen_ZeroChainSelectorNotAllowed() public {
    OffRamp.SourceChainConfigArgs[] memory configs = new OffRamp.SourceChainConfigArgs[](1);
    configs[0] = OffRamp.SourceChainConfigArgs({
      router: s_sourceRouter,
      sourceChainSelector: 0,
      isEnabled: true,
      onRamps: new bytes[](0),
      defaultCCVs: new address[](1),
      laneMandatedCCVs: new address[](0)
    });
    configs[0].defaultCCVs[0] = makeAddr("ccv");

    vm.expectRevert(OffRamp.ZeroChainSelectorNotAllowed.selector);
    s_offRamp.applySourceChainConfigUpdates(configs);
  }

  function test_applySourceChainConfigUpdates_RevertWhen_ZeroAddressNotAllowed_Router() public {
    OffRamp.SourceChainConfigArgs[] memory configs = new OffRamp.SourceChainConfigArgs[](1);
    configs[0] = OffRamp.SourceChainConfigArgs({
      router: IRouter(address(0)),
      sourceChainSelector: SOURCE_CHAIN_SELECTOR + 1,
      isEnabled: true,
      onRamps: new bytes[](0),
      defaultCCVs: new address[](1),
      laneMandatedCCVs: new address[](0)
    });
    configs[0].defaultCCVs[0] = makeAddr("ccv");

    vm.expectRevert(OffRamp.ZeroAddressNotAllowed.selector);
    s_offRamp.applySourceChainConfigUpdates(configs);
  }

  function test_applySourceChainConfigUpdates_RevertWhen_ZeroAddressNotAllowed_DefaultCCV() public {
    OffRamp.SourceChainConfigArgs[] memory configs = new OffRamp.SourceChainConfigArgs[](1);
    configs[0] = OffRamp.SourceChainConfigArgs({
      router: s_sourceRouter,
      sourceChainSelector: SOURCE_CHAIN_SELECTOR + 1,
      isEnabled: true,
      onRamps: new bytes[](0),
      defaultCCVs: new address[](0),
      laneMandatedCCVs: new address[](0)
    });

    vm.expectRevert(OffRamp.ZeroAddressNotAllowed.selector);
    s_offRamp.applySourceChainConfigUpdates(configs);
  }

  function test_applySourceChainConfigUpdates_RevertWhen_ZeroAddressNotAllowed_ZeroAddressInDefaultCCVsArray() public {
    OffRamp.SourceChainConfigArgs[] memory configs = new OffRamp.SourceChainConfigArgs[](1);
    configs[0] = OffRamp.SourceChainConfigArgs({
      router: s_sourceRouter,
      sourceChainSelector: SOURCE_CHAIN_SELECTOR + 1,
      isEnabled: true,
      onRamps: new bytes[](0),
      defaultCCVs: new address[](1),
      laneMandatedCCVs: new address[](0)
    });
    configs[0].defaultCCVs[0] = address(0); // Zero address in array

    vm.expectRevert(OffRamp.ZeroAddressNotAllowed.selector);
    s_offRamp.applySourceChainConfigUpdates(configs);
  }

  function test_applySourceChainConfigUpdates_RevertWhen_ZeroAddressNotAllowed_ZeroInLaneMandatedCCVsArray() public {
    OffRamp.SourceChainConfigArgs[] memory configs = new OffRamp.SourceChainConfigArgs[](1);
    configs[0] = OffRamp.SourceChainConfigArgs({
      router: s_sourceRouter,
      sourceChainSelector: SOURCE_CHAIN_SELECTOR + 1,
      isEnabled: true,
      onRamps: new bytes[](0),
      defaultCCVs: new address[](1),
      laneMandatedCCVs: new address[](1)
    });
    configs[0].defaultCCVs[0] = makeAddr("ccv1");
    configs[0].laneMandatedCCVs[0] = address(0); // Zero address in array

    vm.expectRevert(OffRamp.ZeroAddressNotAllowed.selector);
    s_offRamp.applySourceChainConfigUpdates(configs);
  }

  function test_applySourceChainConfigUpdates_RevertWhen_ZeroAddressNotAllowed_OnRamp() public {
    bytes[] memory onRamps = new bytes[](1);
    onRamps[0] = "";

    OffRamp.SourceChainConfigArgs[] memory configs = new OffRamp.SourceChainConfigArgs[](1);
    configs[0] = OffRamp.SourceChainConfigArgs({
      router: s_sourceRouter,
      sourceChainSelector: SOURCE_CHAIN_SELECTOR + 1,
      isEnabled: true,
      onRamps: onRamps,
      defaultCCVs: new address[](1),
      laneMandatedCCVs: new address[](0)
    });
    configs[0].defaultCCVs[0] = makeAddr("ccv");

    vm.expectRevert(OffRamp.ZeroAddressNotAllowed.selector);
    s_offRamp.applySourceChainConfigUpdates(configs);
  }

  function test_applySourceChainConfigUpdates_RevertWhen_OnlyCallableByOwner() public {
    vm.stopPrank();
    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);

    s_offRamp.applySourceChainConfigUpdates(new OffRamp.SourceChainConfigArgs[](0));
  }
}
