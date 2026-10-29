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

    s_helper.removePolicy(POLICY);
    assertEq(0, s_helper.getPolicyNames().length);

    // A removed name can be used again, and it starts as a new policy.
    s_helper.setPolicyRecord(POLICY, OTHER_CHAIN_SELECTOR);
    assertEq(1, s_helper.getPolicyNames().length);
    uint256[] memory policyChainSelectors = s_helper.getPolicyChainSelectors(POLICY);
    assertEq(1, policyChainSelectors.length);
    assertEq(OTHER_CHAIN_SELECTOR, policyChainSelectors[0]);
    assertEq(0, s_helper.getTokenCount(POLICY));
  }

  function test__removePolicy_AfterLastTokenLeaves() public {
    s_helper.moveToken(bytes32(0), POLICY);
    s_helper.moveToken(POLICY, bytes32(0));
    s_helper.removePolicyRecord(POLICY, CHAIN_SELECTOR);

    s_helper.removePolicy(POLICY);
    assertEq(0, s_helper.getPolicyNames().length);
  }

  function test__removePolicy_KeepsOtherPolicies() public {
    s_helper.setPolicyRecord(OTHER_POLICY, CHAIN_SELECTOR);
    s_helper.removePolicyRecord(POLICY, CHAIN_SELECTOR);

    s_helper.removePolicy(POLICY);

    bytes32[] memory policyNames = s_helper.getPolicyNames();
    assertEq(1, policyNames.length);
    assertEq(OTHER_POLICY, policyNames[0]);
  }

  // Reverts

  function test__removePolicy_RevertWhen_PolicyNotFound() public {
    vm.expectRevert(abi.encodeWithSelector(PolicyConfig.PolicyNotFound.selector, OTHER_POLICY));
    s_helper.removePolicy(OTHER_POLICY);

    vm.expectRevert(abi.encodeWithSelector(PolicyConfig.PolicyNotFound.selector, bytes32(0)));
    s_helper.removePolicy(bytes32(0));
  }

  function test__removePolicy_RevertWhen_PolicyInUse() public {
    // The policy stays in use while one of its two tokens is left.
    s_helper.moveToken(bytes32(0), POLICY);
    s_helper.moveToken(bytes32(0), POLICY);
    s_helper.moveToken(POLICY, bytes32(0));
    s_helper.removePolicyRecord(POLICY, CHAIN_SELECTOR);

    vm.expectRevert(abi.encodeWithSelector(PolicyConfig.PolicyInUse.selector, POLICY));
    s_helper.removePolicy(POLICY);
  }

  function test__removePolicy_RevertWhen_PolicyHasRecords() public {
    vm.expectRevert(abi.encodeWithSelector(PolicyConfig.PolicyHasRecords.selector, POLICY));
    s_helper.removePolicy(POLICY);
  }
}
