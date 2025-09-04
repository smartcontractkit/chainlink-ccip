// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IRMNRemote} from "../../../interfaces/IRMNRemote.sol";
import {CCVProxy} from "../../../onRamp/CCVProxy.sol";
import {CCVProxySetup} from "./CCVProxySetup.t.sol";

contract CCVProxy_constructor is CCVProxySetup {
  function test_constructor() public {
    CCVProxy.StaticConfig memory s = CCVProxy.StaticConfig({
      chainSelector: SOURCE_CHAIN_SELECTOR,
      rmnRemote: s_mockRMNRemote,
      tokenAdminRegistry: address(s_tokenAdminRegistry)
    });

    CCVProxy.DynamicConfig memory d = CCVProxy.DynamicConfig({
      feeQuoter: address(s_feeQuoter),
      reentrancyGuardEntered: false,
      feeAggregator: FEE_AGGREGATOR
    });

    CCVProxy proxy = new CCVProxy(s, d);

    CCVProxy.StaticConfig memory gotS = proxy.getStaticConfig();
    assertEq(gotS.chainSelector, s.chainSelector);
    assertEq(address(gotS.rmnRemote), address(s.rmnRemote));
    assertEq(gotS.tokenAdminRegistry, s.tokenAdminRegistry);

    CCVProxy.DynamicConfig memory gotD = proxy.getDynamicConfig();
    assertEq(gotD.feeQuoter, d.feeQuoter);
    assertEq(gotD.feeAggregator, d.feeAggregator);
    assertFalse(gotD.reentrancyGuardEntered);
  }

  function test_constructor_RevertWhen_StaticConfigInvalid() public {
    // Zero chainSelector
    CCVProxy.StaticConfig memory s0 = CCVProxy.StaticConfig({
      chainSelector: 0,
      rmnRemote: s_mockRMNRemote,
      tokenAdminRegistry: address(s_tokenAdminRegistry)
    });
    CCVProxy.DynamicConfig memory d = CCVProxy.DynamicConfig({
      feeQuoter: address(s_feeQuoter),
      reentrancyGuardEntered: false,
      feeAggregator: FEE_AGGREGATOR
    });
    vm.expectRevert(CCVProxy.InvalidConfig.selector);
    new CCVProxy(s0, d);

    // Zero rmnRemote
    CCVProxy.StaticConfig memory s1 = CCVProxy.StaticConfig({
      chainSelector: SOURCE_CHAIN_SELECTOR,
      rmnRemote: IRMNRemote(address(0)),
      tokenAdminRegistry: address(s_tokenAdminRegistry)
    });
    vm.expectRevert(CCVProxy.InvalidConfig.selector);
    new CCVProxy(s1, d);

    // Zero tokenAdminRegistry
    CCVProxy.StaticConfig memory s2 = CCVProxy.StaticConfig({
      chainSelector: SOURCE_CHAIN_SELECTOR,
      rmnRemote: s_mockRMNRemote,
      tokenAdminRegistry: address(0)
    });
    vm.expectRevert(CCVProxy.InvalidConfig.selector);
    new CCVProxy(s2, d);
  }

  function test_constructor_RevertWhen_DynamicConfigInvalid() public {
    CCVProxy.StaticConfig memory s = CCVProxy.StaticConfig({
      chainSelector: SOURCE_CHAIN_SELECTOR,
      rmnRemote: s_mockRMNRemote,
      tokenAdminRegistry: address(s_tokenAdminRegistry)
    });

    // feeQuoter == address(0)
    CCVProxy.DynamicConfig memory dynamicConfig0 =
      CCVProxy.DynamicConfig({feeQuoter: address(0), reentrancyGuardEntered: false, feeAggregator: FEE_AGGREGATOR});
    vm.expectRevert(CCVProxy.InvalidConfig.selector);
    new CCVProxy(s, dynamicConfig0);

    // feeAggregator == address(0)
    CCVProxy.DynamicConfig memory dynamicConfig1 = CCVProxy.DynamicConfig({
      feeQuoter: address(s_feeQuoter),
      reentrancyGuardEntered: false,
      feeAggregator: address(0)
    });
    vm.expectRevert(CCVProxy.InvalidConfig.selector);
    new CCVProxy(s, dynamicConfig1);

    // reentrancyGuardEntered == true
    CCVProxy.DynamicConfig memory dynamicConfig2 = CCVProxy.DynamicConfig({
      feeQuoter: address(s_feeQuoter),
      reentrancyGuardEntered: true,
      feeAggregator: FEE_AGGREGATOR
    });
    vm.expectRevert(CCVProxy.InvalidConfig.selector);
    new CCVProxy(s, dynamicConfig2);
  }
}
