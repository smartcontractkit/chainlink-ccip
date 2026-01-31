// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {AdvancedPoolHooks} from "../../../pools/AdvancedPoolHooks.sol";
import {
  MockPolicyEngine,
  MockPolicyEngineNoDetach,
  MockPolicyEngineRevertingDetach
} from "../../mocks/MockPolicyEngine.sol";
import {AdvancedPoolHooksSetup} from "./AdvancedPoolHooksSetup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract AdvancedPoolHooks_setPolicyEngineAllowFailedDetach is AdvancedPoolHooksSetup {
  MockPolicyEngine internal s_mockPolicyEngine;
  MockPolicyEngine internal s_mockPolicyEngine2;

  function setUp() public virtual override {
    super.setUp();
    s_mockPolicyEngine = new MockPolicyEngine();
    s_mockPolicyEngine2 = new MockPolicyEngine();
  }

  function test_setPolicyEngineAllowFailedDetach_Swap() public {
    s_advancedPoolHooks.setPolicyEngineAllowFailedDetach(address(s_mockPolicyEngine));
    assertTrue(s_mockPolicyEngine.isAttached(address(s_advancedPoolHooks)));
    assertFalse(s_mockPolicyEngine2.isAttached(address(s_advancedPoolHooks)));

    vm.expectEmit();
    emit AdvancedPoolHooks.PolicyEngineSet(address(s_mockPolicyEngine), address(s_mockPolicyEngine2));

    s_advancedPoolHooks.setPolicyEngineAllowFailedDetach(address(s_mockPolicyEngine2));

    assertEq(address(s_mockPolicyEngine2), s_advancedPoolHooks.getPolicyEngine());
    assertFalse(s_mockPolicyEngine.isAttached(address(s_advancedPoolHooks)));
    assertTrue(s_mockPolicyEngine2.isAttached(address(s_advancedPoolHooks)));
  }

  function test_setPolicyEngineAllowFailedDetach_OldEngineDetachReverts() public {
    MockPolicyEngineRevertingDetach revertingEngine = new MockPolicyEngineRevertingDetach();
    s_advancedPoolHooks.setPolicyEngineAllowFailedDetach(address(revertingEngine));
    assertEq(address(revertingEngine), s_advancedPoolHooks.getPolicyEngine());

    vm.expectEmit();
    emit AdvancedPoolHooks.PolicyEngineSet(address(revertingEngine), address(s_mockPolicyEngine));

    s_advancedPoolHooks.setPolicyEngineAllowFailedDetach(address(s_mockPolicyEngine));

    assertEq(address(s_mockPolicyEngine), s_advancedPoolHooks.getPolicyEngine());
    assertTrue(s_mockPolicyEngine.isAttached(address(s_advancedPoolHooks)));
  }

  function test_setPolicyEngineAllowFailedDetach_OldEngineDoesNotImplementDetach() public {
    MockPolicyEngineNoDetach noDetachEngine = new MockPolicyEngineNoDetach();
    s_advancedPoolHooks.setPolicyEngineAllowFailedDetach(address(noDetachEngine));
    assertEq(address(noDetachEngine), s_advancedPoolHooks.getPolicyEngine());

    vm.expectEmit();
    emit AdvancedPoolHooks.PolicyEngineSet(address(noDetachEngine), address(s_mockPolicyEngine));

    s_advancedPoolHooks.setPolicyEngineAllowFailedDetach(address(s_mockPolicyEngine));

    assertEq(address(s_mockPolicyEngine), s_advancedPoolHooks.getPolicyEngine());
    assertTrue(s_mockPolicyEngine.isAttached(address(s_advancedPoolHooks)));
  }

  // Reverts

  function test_setPolicyEngineAllowFailedDetach_RevertWhen_OnlyCallableByOwner() public {
    vm.stopPrank();
    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);

    s_advancedPoolHooks.setPolicyEngineAllowFailedDetach(address(s_mockPolicyEngine));
  }
}

