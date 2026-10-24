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

contract MultiTokenPoolHooks_setPolicyEngineAllowFailedDetach is MultiTokenPoolHooksSetup {
  function test_setPolicyEngineAllowFailedDetach() public {
    MockPolicyEngine newPolicyEngine = new MockPolicyEngine();
    s_hooks.setPolicyEngineAllowFailedDetach(s_policyEngine);

    vm.expectEmit();
    emit MultiTokenPoolHooks.PolicyEngineAttached(newPolicyEngine);
    s_hooks.setPolicyEngineAllowFailedDetach(newPolicyEngine);

    assertEq(address(newPolicyEngine), address(s_hooks.getPolicyEngine()));
    assertFalse(s_policyEngine.isAttached(address(s_hooks)));
    assertTrue(newPolicyEngine.isAttached(address(s_hooks)));
  }

  function test_setPolicyEngineAllowFailedDetach_OldPolicyEngineDetachReverts() public {
    IPolicyEngine revertingPolicyEngine = IPolicyEngine(address(new MockPolicyEngineRevertingDetach()));
    s_hooks.setPolicyEngine(revertingPolicyEngine);

    vm.expectEmit();
    emit MultiTokenPoolHooks.PolicyEngineDetachFailed(
      revertingPolicyEngine, abi.encodeWithSelector(MockPolicyEngineRevertingDetach.DetachNotSupported.selector)
    );
    vm.expectEmit();
    emit MultiTokenPoolHooks.PolicyEngineAttached(s_policyEngine);
    s_hooks.setPolicyEngineAllowFailedDetach(s_policyEngine);

    assertEq(address(s_policyEngine), address(s_hooks.getPolicyEngine()));
    assertTrue(s_policyEngine.isAttached(address(s_hooks)));
  }

  function test_setPolicyEngineAllowFailedDetach_OldPolicyEngineHasNoDetach() public {
    IPolicyEngine noDetachPolicyEngine = IPolicyEngine(address(new MockPolicyEngineNoDetach()));
    s_hooks.setPolicyEngine(noDetachPolicyEngine);

    vm.expectEmit();
    emit MultiTokenPoolHooks.PolicyEngineDetachFailed(noDetachPolicyEngine, "");
    vm.expectEmit();
    emit MultiTokenPoolHooks.PolicyEngineAttached(s_policyEngine);
    s_hooks.setPolicyEngineAllowFailedDetach(s_policyEngine);

    assertEq(address(s_policyEngine), address(s_hooks.getPolicyEngine()));
    assertTrue(s_policyEngine.isAttached(address(s_hooks)));
  }

  // Reverts

  function test_setPolicyEngineAllowFailedDetach_RevertWhen_OnlyCallableByOwner() public {
    vm.startPrank(STRANGER);

    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_hooks.setPolicyEngineAllowFailedDetach(s_policyEngine);
  }
}
