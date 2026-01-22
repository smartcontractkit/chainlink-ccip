// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {AdvancedPoolHooks} from "../../../pools/AdvancedPoolHooks.sol";
import {MockPolicyEngine} from "../../mocks/MockPolicyEngine.sol";
import {AdvancedPoolHooksSetup} from "./AdvancedPoolHooksSetup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract AdvancedPoolHooks_setPolicyEngine is AdvancedPoolHooksSetup {
  MockPolicyEngine internal s_mockPolicyEngine;
  MockPolicyEngine internal s_mockPolicyEngine2;

  function setUp() public virtual override {
    super.setUp();
    s_mockPolicyEngine = new MockPolicyEngine();
    s_mockPolicyEngine2 = new MockPolicyEngine();
  }

  function test_setPolicyEngine() public {
    assertEq(s_advancedPoolHooks.getPolicyEngine(), address(0));

    vm.expectEmit();
    emit AdvancedPoolHooks.PolicyEngineSet(address(0), address(s_mockPolicyEngine));

    s_advancedPoolHooks.setPolicyEngine(address(s_mockPolicyEngine));

    assertEq(s_advancedPoolHooks.getPolicyEngine(), address(s_mockPolicyEngine));
    assertTrue(s_mockPolicyEngine.isAttached(address(s_advancedPoolHooks)));
  }

  function test_setPolicyEngine_ToZeroAddress() public {
    // First set a policy engine
    s_advancedPoolHooks.setPolicyEngine(address(s_mockPolicyEngine));
    assertEq(s_advancedPoolHooks.getPolicyEngine(), address(s_mockPolicyEngine));
    assertTrue(s_mockPolicyEngine.isAttached(address(s_advancedPoolHooks)));

    // Now disable by setting to address(0)
    vm.expectEmit();
    emit AdvancedPoolHooks.PolicyEngineSet(address(s_mockPolicyEngine), address(0));

    s_advancedPoolHooks.setPolicyEngine(address(0));

    assertEq(s_advancedPoolHooks.getPolicyEngine(), address(0));
    assertFalse(s_mockPolicyEngine.isAttached(address(s_advancedPoolHooks)));
  }

  function test_setPolicyEngine_Change() public {
    // First set a policy engine
    s_advancedPoolHooks.setPolicyEngine(address(s_mockPolicyEngine));
    assertTrue(s_mockPolicyEngine.isAttached(address(s_advancedPoolHooks)));
    assertFalse(s_mockPolicyEngine2.isAttached(address(s_advancedPoolHooks)));

    // Change to a different policy engine
    vm.expectEmit();
    emit AdvancedPoolHooks.PolicyEngineSet(address(s_mockPolicyEngine), address(s_mockPolicyEngine2));

    s_advancedPoolHooks.setPolicyEngine(address(s_mockPolicyEngine2));

    assertEq(s_advancedPoolHooks.getPolicyEngine(), address(s_mockPolicyEngine2));
    assertFalse(s_mockPolicyEngine.isAttached(address(s_advancedPoolHooks)));
    assertTrue(s_mockPolicyEngine2.isAttached(address(s_advancedPoolHooks)));
  }

  function test_setPolicyEngine_SameValue() public {
    // First set a policy engine
    s_advancedPoolHooks.setPolicyEngine(address(s_mockPolicyEngine));
    assertTrue(s_mockPolicyEngine.isAttached(address(s_advancedPoolHooks)));

    // Setting the same value should be a no-op (no event emitted)
    vm.recordLogs();
    s_advancedPoolHooks.setPolicyEngine(address(s_mockPolicyEngine));

    // Verify no events were emitted
    assertEq(vm.getRecordedLogs().length, 0);

    // State should remain unchanged
    assertEq(s_advancedPoolHooks.getPolicyEngine(), address(s_mockPolicyEngine));
    assertTrue(s_mockPolicyEngine.isAttached(address(s_advancedPoolHooks)));
  }

  function test_getPolicyEngine() public {
    // Initially should be address(0)
    assertEq(s_advancedPoolHooks.getPolicyEngine(), address(0));

    // After setting, should return the correct address
    s_advancedPoolHooks.setPolicyEngine(address(s_mockPolicyEngine));
    assertEq(s_advancedPoolHooks.getPolicyEngine(), address(s_mockPolicyEngine));
  }

  // Reverts

  function test_setPolicyEngine_RevertWhen_OnlyCallableByOwner() public {
    vm.stopPrank();
    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);

    s_advancedPoolHooks.setPolicyEngine(address(s_mockPolicyEngine));
  }
}
