// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IRMNRemote} from "../../../interfaces/IRMNRemote.sol";
import {OffRamp} from "../../../offRamp/OffRamp.sol";
import {BaseTest} from "../../BaseTest.t.sol";

contract OffRamp_constructor is BaseTest {
  function test_constructor() public {
    OffRamp.StaticConfig memory config = OffRamp.StaticConfig({
      localChainSelector: DEST_CHAIN_SELECTOR,
      gasForCallExactCheck: 5000,
      rmnRemote: s_mockRMNRemote,
      tokenAdminRegistry: address(0x123)
    });

    vm.expectEmit();
    emit OffRamp.StaticConfigSet(config);

    OffRamp offRamp = new OffRamp(config);

    OffRamp.StaticConfig memory returnedConfig = offRamp.getStaticConfig();
    assertEq(returnedConfig.localChainSelector, config.localChainSelector);
    assertEq(returnedConfig.gasForCallExactCheck, config.gasForCallExactCheck);
    assertEq(address(returnedConfig.rmnRemote), address(config.rmnRemote));
    assertEq(returnedConfig.tokenAdminRegistry, config.tokenAdminRegistry);
  }

  function test_constructor_RevertWhen_ZeroAddressNotAllowed_RMNRemote() public {
    OffRamp.StaticConfig memory config = OffRamp.StaticConfig({
      localChainSelector: DEST_CHAIN_SELECTOR,
      gasForCallExactCheck: 5000,
      rmnRemote: IRMNRemote(address(0)),
      tokenAdminRegistry: address(0x123)
    });

    vm.expectRevert(OffRamp.ZeroAddressNotAllowed.selector);
    new OffRamp(config);
  }

  function test_constructor_RevertWhen_ZeroAddressNotAllowed_TokenAdminRegistry() public {
    OffRamp.StaticConfig memory config = OffRamp.StaticConfig({
      localChainSelector: DEST_CHAIN_SELECTOR,
      gasForCallExactCheck: 5000,
      rmnRemote: s_mockRMNRemote,
      tokenAdminRegistry: address(0)
    });

    vm.expectRevert(OffRamp.ZeroAddressNotAllowed.selector);
    new OffRamp(config);
  }

  function test_constructor_RevertWhen_ZeroChainSelectorNotAllowed() public {
    OffRamp.StaticConfig memory config = OffRamp.StaticConfig({
      localChainSelector: 0,
      gasForCallExactCheck: 5000,
      rmnRemote: s_mockRMNRemote,
      tokenAdminRegistry: address(0x123)
    });

    vm.expectRevert(OffRamp.ZeroChainSelectorNotAllowed.selector);
    new OffRamp(config);
  }
}
