// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {PolicyConfig} from "../../../libraries/PolicyConfig.sol";
import {PolicyConfigSetup} from "./PolicyConfigSetup.t.sol";

contract PolicyConfig__resolvePolicyRecordLevel is PolicyConfigSetup {
  function setUp() public override {
    super.setUp();
    s_helper.setPolicyRecord(POLICY, CHAIN_SELECTOR);
    s_helper.setPolicyRecord(POLICY, PolicyConfig.DEFAULT_CHAIN_SELECTOR);
  }

  function test__resolvePolicyRecordLevel_PolicyChain() public view {
    _assertLevel(PolicyConfig.RecordLevel.PolicyChain, POLICY, CHAIN_SELECTOR);
  }

  function test__resolvePolicyRecordLevel_PolicyDefault() public view {
    _assertLevel(PolicyConfig.RecordLevel.PolicyDefault, POLICY, OTHER_CHAIN_SELECTOR);
  }

  function test__resolvePolicyRecordLevel_NoDefaultRecord() public {
    s_helper.removePolicyRecord(POLICY, PolicyConfig.DEFAULT_CHAIN_SELECTOR);

    _assertLevel(PolicyConfig.RecordLevel.PolicyChain, POLICY, CHAIN_SELECTOR);
    _assertLevel(PolicyConfig.RecordLevel.None, POLICY, OTHER_CHAIN_SELECTOR);
  }

  function test__resolvePolicyRecordLevel_ZeroPolicyName() public view {
    // A token without a policy has the zero name, which never has records.
    _assertLevel(PolicyConfig.RecordLevel.None, bytes32(0), CHAIN_SELECTOR);
  }

  function test__resolvePolicyRecordLevel_ZeroChainSelector() public view {
    // Chain selector zero resolves to no record, even with a policy default record.
    _assertLevel(PolicyConfig.RecordLevel.None, POLICY, PolicyConfig.DEFAULT_CHAIN_SELECTOR);
  }

  function testFuzz__resolvePolicyRecordLevel_DefaultRecordServesOtherChains(
    uint64 remoteChainSelector
  ) public view {
    vm.assume(remoteChainSelector != CHAIN_SELECTOR && remoteChainSelector != PolicyConfig.DEFAULT_CHAIN_SELECTOR);

    _assertLevel(PolicyConfig.RecordLevel.PolicyDefault, POLICY, remoteChainSelector);
  }

  function _assertLevel(
    PolicyConfig.RecordLevel expected,
    bytes32 policyName,
    uint64 remoteChainSelector
  ) internal view {
    assertEq(uint8(expected), uint8(s_helper.resolvePolicyRecordLevel(policyName, remoteChainSelector)));
  }
}
