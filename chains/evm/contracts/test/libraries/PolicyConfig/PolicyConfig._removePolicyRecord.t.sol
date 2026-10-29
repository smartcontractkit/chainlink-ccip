// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {PolicyConfig} from "../../../libraries/PolicyConfig.sol";
import {PolicyConfigSetup} from "./PolicyConfigSetup.t.sol";

contract PolicyConfig__removePolicyRecord is PolicyConfigSetup {
  function setUp() public override {
    super.setUp();
    s_helper.setPolicyRecord(POLICY, CHAIN_SELECTOR);
  }

  function test__removePolicyRecord() public {
    s_helper.removePolicyRecord(POLICY, CHAIN_SELECTOR);

    assertEq(0, s_helper.getPolicyChains(POLICY).length);
    // The policy stays when it has no records left.
    bytes32[] memory policyNames = s_helper.getPolicyNames();
    assertEq(1, policyNames.length);
    assertEq(POLICY, policyNames[0]);
  }

  // Reverts

  function test__removePolicyRecord_RevertWhen_PolicyRecordNotFound() public {
    vm.expectRevert(abi.encodeWithSelector(PolicyConfig.PolicyRecordNotFound.selector, POLICY, OTHER_CHAIN_SELECTOR));
    s_helper.removePolicyRecord(POLICY, OTHER_CHAIN_SELECTOR);

    vm.expectRevert(abi.encodeWithSelector(PolicyConfig.PolicyRecordNotFound.selector, OTHER_POLICY, CHAIN_SELECTOR));
    s_helper.removePolicyRecord(OTHER_POLICY, CHAIN_SELECTOR);
  }
}
