// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IRMNRemote} from "../../../interfaces/IRMNRemote.sol";
import {CCVAggregator} from "../../../offRamp/CCVAggregator.sol";
import {BaseTest} from "../../BaseTest.t.sol";

contract CCVAggregator_constructor is BaseTest {
  function test_constructor() public {
    CCVAggregator.StaticConfig memory config = CCVAggregator.StaticConfig({
      localChainSelector: DEST_CHAIN_SELECTOR,
      gasForCallExactCheck: 5000,
      rmnRemote: s_mockRMNRemote,
      tokenAdminRegistry: address(0x123)
    });

    vm.expectEmit();
    emit CCVAggregator.StaticConfigSet(config);

    CCVAggregator aggregator = new CCVAggregator(config);

    CCVAggregator.StaticConfig memory returnedConfig = aggregator.getStaticConfig();
    assertEq(returnedConfig.localChainSelector, config.localChainSelector);
    assertEq(returnedConfig.gasForCallExactCheck, config.gasForCallExactCheck);
    assertEq(address(returnedConfig.rmnRemote), address(config.rmnRemote));
    assertEq(returnedConfig.tokenAdminRegistry, config.tokenAdminRegistry);
  }

  function test_constructor_RevertWhen_ZeroAddressNotAllowed_RMNRemote() public {
    CCVAggregator.StaticConfig memory config = CCVAggregator.StaticConfig({
      localChainSelector: DEST_CHAIN_SELECTOR,
      gasForCallExactCheck: 5000,
      rmnRemote: IRMNRemote(address(0)),
      tokenAdminRegistry: address(0x123)
    });

    vm.expectRevert(CCVAggregator.ZeroAddressNotAllowed.selector);
    new CCVAggregator(config);
  }

  function test_constructor_RevertWhen_ZeroAddressNotAllowed_TokenAdminRegistry() public {
    CCVAggregator.StaticConfig memory config = CCVAggregator.StaticConfig({
      localChainSelector: DEST_CHAIN_SELECTOR,
      gasForCallExactCheck: 5000,
      rmnRemote: s_mockRMNRemote,
      tokenAdminRegistry: address(0)
    });

    vm.expectRevert(CCVAggregator.ZeroAddressNotAllowed.selector);
    new CCVAggregator(config);
  }

  function test_constructor_RevertWhen_ZeroChainSelectorNotAllowed() public {
    CCVAggregator.StaticConfig memory config = CCVAggregator.StaticConfig({
      localChainSelector: 0,
      gasForCallExactCheck: 5000,
      rmnRemote: s_mockRMNRemote,
      tokenAdminRegistry: address(0x123)
    });

    vm.expectRevert(CCVAggregator.ZeroChainSelectorNotAllowed.selector);
    new CCVAggregator(config);
  }
}
