// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {AdvancedPoolHooks} from "../../../pools/AdvancedPoolHooks.sol";
import {MockPolicyEngine, MockPolicyEngineNoDetach, MockPolicyEngineRevertingDetach} from "../../mocks/MockPolicyEngine.sol";
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

  function test_setPolicyEngine_FromZeroAddress() public {
    assertEq(address(0), s_advancedPoolHooks.getPolicyEngine());

    vm.expectEmit();
    emit AdvancedPoolHooks.PolicyEngineSet(address(0), address(s_mockPolicyEngine));

    s_advancedPoolHooks.setPolicyEngine(address(s_mockPolicyEngine));

    assertEq(address(s_mockPolicyEngine), s_advancedPoolHooks.getPolicyEngine());
    assertTrue(s_mockPolicyEngine.isAttached(address(s_advancedPoolHooks)));
  }

  function test_setPolicyEngine_ToZeroAddress() public {
    s_advancedPoolHooks.setPolicyEngine(address(s_mockPolicyEngine));
    assertEq(address(s_mockPolicyEngine), s_advancedPoolHooks.getPolicyEngine());
    assertTrue(s_mockPolicyEngine.isAttached(address(s_advancedPoolHooks)));

    vm.expectEmit();
    emit AdvancedPoolHooks.PolicyEngineSet(address(s_mockPolicyEngine), address(0));

    s_advancedPoolHooks.setPolicyEngine(address(0));

    assertEq(address(0), s_advancedPoolHooks.getPolicyEngine());
    assertFalse(s_mockPolicyEngine.isAttached(address(s_advancedPoolHooks)));
  }

  function test_setPolicyEngine_Swap() public {
    s_advancedPoolHooks.setPolicyEngine(address(s_mockPolicyEngine));
    assertTrue(s_mockPolicyEngine.isAttached(address(s_advancedPoolHooks)));
    assertFalse(s_mockPolicyEngine2.isAttached(address(s_advancedPoolHooks)));

    vm.expectEmit();
    emit AdvancedPoolHooks.PolicyEngineSet(address(s_mockPolicyEngine), address(s_mockPolicyEngine2));

    s_advancedPoolHooks.setPolicyEngine(address(s_mockPolicyEngine2));

    assertEq(address(s_mockPolicyEngine2), s_advancedPoolHooks.getPolicyEngine());
    assertFalse(s_mockPolicyEngine.isAttached(address(s_advancedPoolHooks)));
    assertTrue(s_mockPolicyEngine2.isAttached(address(s_advancedPoolHooks)));
  }

  function test_setPolicyEngine_SameValue() public {
    s_advancedPoolHooks.setPolicyEngine(address(s_mockPolicyEngine));
    assertTrue(s_mockPolicyEngine.isAttached(address(s_advancedPoolHooks)));

    vm.recordLogs();
    s_advancedPoolHooks.setPolicyEngine(address(s_mockPolicyEngine));

    assertEq(0, vm.getRecordedLogs().length);
    assertEq(address(s_mockPolicyEngine), s_advancedPoolHooks.getPolicyEngine());
    assertTrue(s_mockPolicyEngine.isAttached(address(s_advancedPoolHooks)));
  }

  function test_getPolicyEngine() public {
    assertEq(address(0), s_advancedPoolHooks.getPolicyEngine());

    s_advancedPoolHooks.setPolicyEngine(address(s_mockPolicyEngine));
    assertEq(address(s_mockPolicyEngine), s_advancedPoolHooks.getPolicyEngine());
  }

  // Reverts

  function test_setPolicyEngine_RevertWhen_OnlyCallableByOwner() public {
    vm.stopPrank();
    vm.prank(STRANGER);
    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_advancedPoolHooks.setPolicyEngine(address(s_mockPolicyEngine));
  }

  function test_setPolicyEngine_RevertWhen_OldEngineDetachReverts() public {
    MockPolicyEngineRevertingDetach revertingEngine = new MockPolicyEngineRevertingDetach();
    s_advancedPoolHooks.setPolicyEngine(address(revertingEngine));
    assertEq(address(revertingEngine), s_advancedPoolHooks.getPolicyEngine());

    vm.expectRevert(
      abi.encodeWithSelector(
        AdvancedPoolHooks.PolicyEngineDetachFailed.selector,
        address(revertingEngine),
        abi.encodeWithSignature("Error(string)", "detach not supported")
      )
    );
    s_advancedPoolHooks.setPolicyEngine(address(s_mockPolicyEngine));

    assertEq(address(revertingEngine), s_advancedPoolHooks.getPolicyEngine());
  }

  function test_setPolicyEngine_RevertWhen_OldEngineDoesNotImplementDetach() public {
    MockPolicyEngineNoDetach noDetachEngine = new MockPolicyEngineNoDetach();
    s_advancedPoolHooks.setPolicyEngine(address(noDetachEngine));
    assertEq(address(noDetachEngine), s_advancedPoolHooks.getPolicyEngine());

    vm.expectRevert(
      abi.encodeWithSelector(AdvancedPoolHooks.PolicyEngineDetachFailed.selector, address(noDetachEngine), "")
    );
    s_advancedPoolHooks.setPolicyEngine(address(s_mockPolicyEngine));

    assertEq(address(noDetachEngine), s_advancedPoolHooks.getPolicyEngine());
  }
}
