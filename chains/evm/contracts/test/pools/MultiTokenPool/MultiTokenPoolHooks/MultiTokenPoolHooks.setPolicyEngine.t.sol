// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPolicyEngine} from "@chainlink/policy-management/interfaces/IPolicyEngine.sol";

import {MultiTokenPoolHooks} from "../../../../pools/MultiTokenPool/MultiTokenPoolHooks.sol";
import {
  MockPolicyEngine,
  MockPolicyEngineNoDetach,
  MockPolicyEngineRevertingDetach
} from "../../../mocks/MockPolicyEngine.sol";
import {MultiTokenPoolHooksSetup} from "./MultiTokenPoolHooksSetup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract MultiTokenPoolHooks_setPolicyEngine is MultiTokenPoolHooksSetup {
  function test_setPolicyEngine() public {
    vm.expectEmit();
    emit MultiTokenPoolHooks.PolicyEngineAttached(s_policyEngine);
    s_hooks.setPolicyEngine(s_policyEngine);

    assertEq(address(s_policyEngine), address(s_hooks.getPolicyEngine()));
    assertTrue(s_policyEngine.isAttached(address(s_hooks)));
  }

  function test_setPolicyEngine_Swap() public {
    MockPolicyEngine newPolicyEngine = new MockPolicyEngine();
    s_hooks.setPolicyEngine(s_policyEngine);

    vm.expectEmit();
    emit MultiTokenPoolHooks.PolicyEngineAttached(newPolicyEngine);
    s_hooks.setPolicyEngine(newPolicyEngine);

    assertEq(address(newPolicyEngine), address(s_hooks.getPolicyEngine()));
    assertFalse(s_policyEngine.isAttached(address(s_hooks)));
    assertTrue(newPolicyEngine.isAttached(address(s_hooks)));
  }

  function test_setPolicyEngine_ToZero() public {
    s_hooks.setPolicyEngine(s_policyEngine);

    vm.expectEmit();
    emit MultiTokenPoolHooks.PolicyEngineAttached(IPolicyEngine(address(0)));
    s_hooks.setPolicyEngine(IPolicyEngine(address(0)));

    assertEq(address(0), address(s_hooks.getPolicyEngine()));
    assertFalse(s_policyEngine.isAttached(address(s_hooks)));
  }

  function test_setPolicyEngine_SameValue() public {
    s_hooks.setPolicyEngine(s_policyEngine);

    vm.recordLogs();
    s_hooks.setPolicyEngine(s_policyEngine);

    assertEq(0, vm.getRecordedLogs().length);
    assertEq(address(s_policyEngine), address(s_hooks.getPolicyEngine()));
    assertTrue(s_policyEngine.isAttached(address(s_hooks)));
  }

  // Reverts

  function test_setPolicyEngine_RevertWhen_OnlyCallableByOwner() public {
    vm.startPrank(STRANGER);

    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_hooks.setPolicyEngine(s_policyEngine);
  }

  function test_setPolicyEngine_RevertWhen_OldPolicyEngineDetachReverts() public {
    IPolicyEngine revertingPolicyEngine = IPolicyEngine(address(new MockPolicyEngineRevertingDetach()));
    s_hooks.setPolicyEngine(revertingPolicyEngine);

    vm.expectRevert(
      abi.encodeWithSelector(
        MultiTokenPoolHooks.PolicyEngineDetachReverted.selector,
        revertingPolicyEngine,
        abi.encodeWithSelector(MockPolicyEngineRevertingDetach.DetachNotSupported.selector)
      )
    );
    s_hooks.setPolicyEngine(s_policyEngine);
  }

  function test_setPolicyEngine_RevertWhen_OldPolicyEngineHasNoDetach() public {
    IPolicyEngine noDetachPolicyEngine = IPolicyEngine(address(new MockPolicyEngineNoDetach()));
    s_hooks.setPolicyEngine(noDetachPolicyEngine);

    vm.expectRevert(
      abi.encodeWithSelector(MultiTokenPoolHooks.PolicyEngineDetachReverted.selector, noDetachPolicyEngine, "")
    );
    s_hooks.setPolicyEngine(s_policyEngine);
  }
}
