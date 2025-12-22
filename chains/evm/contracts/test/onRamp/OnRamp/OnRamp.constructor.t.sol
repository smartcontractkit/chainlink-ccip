// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IRMNRemote} from "../../../interfaces/IRMNRemote.sol";
import {OnRamp} from "../../../onRamp/OnRamp.sol";
import {OnRampSetup} from "./OnRampSetup.t.sol";

contract OnRamp_constructor is OnRampSetup {
  function test_constructor() public {
    OnRamp.StaticConfig memory staticConfig = OnRamp.StaticConfig({
      chainSelector: SOURCE_CHAIN_SELECTOR,
      rmnRemote: s_mockRMNRemote,
      tokenAdminRegistry: address(s_tokenAdminRegistry)
    });

    OnRamp.DynamicConfig memory dynamicConfig = OnRamp.DynamicConfig({
      feeQuoter: address(s_feeQuoter),
      reentrancyGuardEntered: false,
      feeAggregator: FEE_AGGREGATOR
    });

    OnRamp proxy = new OnRamp(staticConfig, dynamicConfig);

    OnRamp.StaticConfig memory gotStaticConfig = proxy.getStaticConfig();
    assertEq(gotStaticConfig.chainSelector, staticConfig.chainSelector);
    assertEq(address(gotStaticConfig.rmnRemote), address(staticConfig.rmnRemote));
    assertEq(gotStaticConfig.tokenAdminRegistry, staticConfig.tokenAdminRegistry);

    OnRamp.DynamicConfig memory gotDynamicConfig = proxy.getDynamicConfig();
    assertEq(gotDynamicConfig.feeQuoter, dynamicConfig.feeQuoter);
    assertEq(gotDynamicConfig.feeAggregator, dynamicConfig.feeAggregator);
    assertFalse(gotDynamicConfig.reentrancyGuardEntered);
  }

  function test_constructor_RevertWhen_StaticConfigInvalid() public {
    // Zero chainSelector.
    OnRamp.StaticConfig memory staticConfigZeroChainSelector = OnRamp.StaticConfig({
      chainSelector: 0,
      rmnRemote: s_mockRMNRemote,
      tokenAdminRegistry: address(s_tokenAdminRegistry)
    });
    OnRamp.DynamicConfig memory dynamicConfigValid = OnRamp.DynamicConfig({
      feeQuoter: address(s_feeQuoter),
      reentrancyGuardEntered: false,
      feeAggregator: FEE_AGGREGATOR
    });
    vm.expectRevert(OnRamp.InvalidConfig.selector);
    new OnRamp(staticConfigZeroChainSelector, dynamicConfigValid);

    // Zero rmnRemote.
    OnRamp.StaticConfig memory staticConfigZeroRMNRemote = OnRamp.StaticConfig({
      chainSelector: SOURCE_CHAIN_SELECTOR,
      rmnRemote: IRMNRemote(address(0)),
      tokenAdminRegistry: address(s_tokenAdminRegistry)
    });
    vm.expectRevert(OnRamp.InvalidConfig.selector);
    new OnRamp(staticConfigZeroRMNRemote, dynamicConfigValid);

    // Zero tokenAdminRegistry.
    OnRamp.StaticConfig memory staticConfigZeroTokenAdminRegistry = OnRamp.StaticConfig({
      chainSelector: SOURCE_CHAIN_SELECTOR,
      rmnRemote: s_mockRMNRemote,
      tokenAdminRegistry: address(0)
    });
    vm.expectRevert(OnRamp.InvalidConfig.selector);
    new OnRamp(staticConfigZeroTokenAdminRegistry, dynamicConfigValid);
  }

  function test_constructor_RevertWhen_DynamicConfigInvalid() public {
    OnRamp.StaticConfig memory staticConfig = OnRamp.StaticConfig({
      chainSelector: SOURCE_CHAIN_SELECTOR,
      rmnRemote: s_mockRMNRemote,
      tokenAdminRegistry: address(s_tokenAdminRegistry)
    });

    // feeQuoter == address(0)
    OnRamp.DynamicConfig memory dynamicConfig0 =
      OnRamp.DynamicConfig({feeQuoter: address(0), reentrancyGuardEntered: false, feeAggregator: FEE_AGGREGATOR});
    vm.expectRevert(OnRamp.InvalidConfig.selector);
    new OnRamp(staticConfig, dynamicConfig0);

    // feeAggregator == address(0)
    OnRamp.DynamicConfig memory dynamicConfig1 =
      OnRamp.DynamicConfig({feeQuoter: address(s_feeQuoter), reentrancyGuardEntered: false, feeAggregator: address(0)});
    vm.expectRevert(OnRamp.InvalidConfig.selector);
    new OnRamp(staticConfig, dynamicConfig1);

    // reentrancyGuardEntered == true
    OnRamp.DynamicConfig memory dynamicConfig2 = OnRamp.DynamicConfig({
      feeQuoter: address(s_feeQuoter),
      reentrancyGuardEntered: true,
      feeAggregator: FEE_AGGREGATOR
    });
    vm.expectRevert(OnRamp.InvalidConfig.selector);
    new OnRamp(staticConfig, dynamicConfig2);
  }
}
