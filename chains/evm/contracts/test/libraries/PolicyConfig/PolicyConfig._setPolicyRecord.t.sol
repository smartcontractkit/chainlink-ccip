// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {PolicyConfig} from "../../../libraries/PolicyConfig.sol";
import {PolicyConfigSetup} from "./PolicyConfigSetup.t.sol";

contract PolicyConfig__setPolicyRecord is PolicyConfigSetup {
  function test__setPolicyRecord() public {
    s_helper.setPolicyRecord(POLICY, CHAIN_SELECTOR);

    // The first record that is set under a new name creates the policy.
    bytes32[] memory policyNames = s_helper.getPolicyNames();
    assertEq(1, policyNames.length);
    assertEq(POLICY, policyNames[0]);
    uint256[] memory policyChainSelectors = s_helper.getPolicyChainSelectors(POLICY);
    assertEq(1, policyChainSelectors.length);
    assertEq(CHAIN_SELECTOR, policyChainSelectors[0]);
    assertEq(0, s_helper.getTokenCount(POLICY));
  }

  function test__setPolicyRecord_ExistingRecord() public {
    s_helper.setPolicyRecord(POLICY, CHAIN_SELECTOR);
    s_helper.setPolicyRecord(POLICY, CHAIN_SELECTOR);

    assertEq(1, s_helper.getPolicyNames().length);
    assertEq(1, s_helper.getPolicyChainSelectors(POLICY).length);
  }

  function test__setPolicyRecord_DefaultRecord() public {
    s_helper.setPolicyRecord(POLICY, CHAIN_SELECTOR);
    s_helper.setPolicyRecord(POLICY, PolicyConfig.DEFAULT_CHAIN_SELECTOR);

    uint256[] memory policyChainSelectors = s_helper.getPolicyChainSelectors(POLICY);
    assertEq(2, policyChainSelectors.length);
    assertEq(CHAIN_SELECTOR, policyChainSelectors[0]);
    assertEq(PolicyConfig.DEFAULT_CHAIN_SELECTOR, policyChainSelectors[1]);
  }

  // Reverts

  function test__setPolicyRecord_RevertWhen_ZeroPolicyNameNotAllowed() public {
    vm.expectRevert(PolicyConfig.ZeroPolicyNameNotAllowed.selector);
    s_helper.setPolicyRecord(bytes32(0), CHAIN_SELECTOR);
  }
}
