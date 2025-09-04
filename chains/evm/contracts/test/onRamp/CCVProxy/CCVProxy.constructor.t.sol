// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IRMNRemote} from "../../../interfaces/IRMNRemote.sol";
import {CCVProxy} from "../../../onRamp/CCVProxy.sol";
import {CCVProxySetup} from "./CCVProxySetup.t.sol";

contract CCVProxy_constructor is CCVProxySetup {
  function test_constructor() public {
    CCVProxy.StaticConfig memory staticConfig = CCVProxy.StaticConfig({
      chainSelector: SOURCE_CHAIN_SELECTOR,
      rmnRemote: s_mockRMNRemote,
      tokenAdminRegistry: address(s_tokenAdminRegistry)
    });

    CCVProxy.DynamicConfig memory dynamicConfig = CCVProxy.DynamicConfig({
      feeQuoter: address(s_feeQuoter),
      reentrancyGuardEntered: false,
      feeAggregator: FEE_AGGREGATOR
    });

    CCVProxy proxy = new CCVProxy(staticConfig, dynamicConfig);

    CCVProxy.StaticConfig memory gotStaticConfig = proxy.getStaticConfig();
    assertEq(gotStaticConfig.chainSelector, staticConfig.chainSelector);
    assertEq(address(gotStaticConfig.rmnRemote), address(staticConfig.rmnRemote));
    assertEq(gotStaticConfig.tokenAdminRegistry, staticConfig.tokenAdminRegistry);

    CCVProxy.DynamicConfig memory gotDynamicConfig = proxy.getDynamicConfig();
    assertEq(gotDynamicConfig.feeQuoter, dynamicConfig.feeQuoter);
    assertEq(gotDynamicConfig.feeAggregator, dynamicConfig.feeAggregator);
    assertFalse(gotDynamicConfig.reentrancyGuardEntered);
  }

  function test_constructor_RevertWhen_StaticConfigInvalid() public {
    // Zero chainSelector.
    CCVProxy.StaticConfig memory staticConfigZeroChainSelector = CCVProxy.StaticConfig({
      chainSelector: 0,
      rmnRemote: s_mockRMNRemote,
      tokenAdminRegistry: address(s_tokenAdminRegistry)
    });
    CCVProxy.DynamicConfig memory dynamicConfigValid = CCVProxy.DynamicConfig({
      feeQuoter: address(s_feeQuoter),
      reentrancyGuardEntered: false,
      feeAggregator: FEE_AGGREGATOR
    });
    vm.expectRevert(CCVProxy.InvalidConfig.selector);
    new CCVProxy(staticConfigZeroChainSelector, dynamicConfigValid);

    // Zero rmnRemote.
    CCVProxy.StaticConfig memory staticConfigZeroRMNRemote = CCVProxy.StaticConfig({
      chainSelector: SOURCE_CHAIN_SELECTOR,
      rmnRemote: IRMNRemote(address(0)),
      tokenAdminRegistry: address(s_tokenAdminRegistry)
    });
    vm.expectRevert(CCVProxy.InvalidConfig.selector);
    new CCVProxy(staticConfigZeroRMNRemote, dynamicConfigValid);

    // Zero tokenAdminRegistry.
    CCVProxy.StaticConfig memory staticConfigZeroTokenAdminRegistry = CCVProxy.StaticConfig({
      chainSelector: SOURCE_CHAIN_SELECTOR,
      rmnRemote: s_mockRMNRemote,
      tokenAdminRegistry: address(0)
    });
    vm.expectRevert(CCVProxy.InvalidConfig.selector);
    new CCVProxy(staticConfigZeroTokenAdminRegistry, dynamicConfigValid);
  }

  function test_constructor_RevertWhen_DynamicConfigInvalid() public {
    CCVProxy.StaticConfig memory staticConfig = CCVProxy.StaticConfig({
      chainSelector: SOURCE_CHAIN_SELECTOR,
      rmnRemote: s_mockRMNRemote,
      tokenAdminRegistry: address(s_tokenAdminRegistry)
    });

    // feeQuoter == address(0)
    CCVProxy.DynamicConfig memory dynamicConfig0 =
      CCVProxy.DynamicConfig({feeQuoter: address(0), reentrancyGuardEntered: false, feeAggregator: FEE_AGGREGATOR});
    vm.expectRevert(CCVProxy.InvalidConfig.selector);
    new CCVProxy(staticConfig, dynamicConfig0);

    // feeAggregator == address(0)
    CCVProxy.DynamicConfig memory dynamicConfig1 = CCVProxy.DynamicConfig({
      feeQuoter: address(s_feeQuoter),
      reentrancyGuardEntered: false,
      feeAggregator: address(0)
    });
    vm.expectRevert(CCVProxy.InvalidConfig.selector);
    new CCVProxy(staticConfig, dynamicConfig1);

    // reentrancyGuardEntered == true
    CCVProxy.DynamicConfig memory dynamicConfig2 = CCVProxy.DynamicConfig({
      feeQuoter: address(s_feeQuoter),
      reentrancyGuardEntered: true,
      feeAggregator: FEE_AGGREGATOR
    });
    vm.expectRevert(CCVProxy.InvalidConfig.selector);
    new CCVProxy(staticConfig, dynamicConfig2);
  }
}
