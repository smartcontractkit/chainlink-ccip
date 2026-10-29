// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {PolicyConfig} from "../../../libraries/PolicyConfig.sol";
import {PolicyConfigSetup} from "./PolicyConfigSetup.t.sol";

contract PolicyConfig__removePolicy is PolicyConfigSetup {
  function setUp() public override {
    super.setUp();
    s_helper.setPolicyRecord(POLICY, CHAIN_SELECTOR);
  }

  function test__removePolicy() public {
    s_helper.removePolicyRecord(POLICY, CHAIN_SELECTOR);

    assertTrue(s_helper.removePolicy(POLICY));
    assertEq(0, s_helper.getPolicyNames().length);

    // A removed name can be used again, and it starts as a new policy.
    s_helper.setPolicyRecord(POLICY, OTHER_CHAIN_SELECTOR);
    assertEq(1, s_helper.getPolicyNames().length);
    uint256[] memory policyChains = s_helper.getPolicyChains(POLICY);
    assertEq(1, policyChains.length);
    assertEq(OTHER_CHAIN_SELECTOR, policyChains[0]);
    assertEq(0, s_helper.getTokenCount(POLICY));
  }

  function test__removePolicy_PolicyNotFound() public {
    assertFalse(s_helper.removePolicy(OTHER_POLICY));
    assertFalse(s_helper.removePolicy(bytes32(0)));

    assertEq(1, s_helper.getPolicyNames().length);
  }

  // Reverts

  function test__removePolicy_RevertWhen_PolicyInUse() public {
    s_helper.moveToken(bytes32(0), POLICY);
    s_helper.removePolicyRecord(POLICY, CHAIN_SELECTOR);

    vm.expectRevert(abi.encodeWithSelector(PolicyConfig.PolicyInUse.selector, POLICY));
    s_helper.removePolicy(POLICY);
  }

  function test__removePolicy_RevertWhen_PolicyHasRecords() public {
    vm.expectRevert(abi.encodeWithSelector(PolicyConfig.PolicyHasRecords.selector, POLICY));
    s_helper.removePolicy(POLICY);
  }
}
