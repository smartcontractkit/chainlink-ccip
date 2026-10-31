// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {PolicyConfig} from "../../../../../libraries/PolicyConfig.sol";
import {FeePolicyStore} from "../../../../../pools/MultiTokenPool/policies/FeePolicyStore.sol";
import {FeePolicyStoreSetup} from "./FeePolicyStoreSetup.t.sol";

contract FeePolicyStore_checkLaneRecords is FeePolicyStoreSetup {
  function test_checkLaneRecords() public view {
    s_feeStore.checkLaneRecords(s_token, DEST_CHAIN_SELECTOR);
  }

  function test_checkLaneRecords_InlineRecords() public {
    _setInlineFeeConfig(_getInlineFeeConfigArgs(s_otherToken, DEST_CHAIN_SELECTOR, _getFeeConfig()));
    _setInlineDestOverheadConfig(
      _getInlineDestOverheadConfigArgs(s_otherToken, DEST_CHAIN_SELECTOR, _getDestOverheadConfig())
    );

    s_feeStore.checkLaneRecords(s_otherToken, DEST_CHAIN_SELECTOR);
  }

  // Reverts

  function test_checkLaneRecords_RevertWhen_FeeRecordNotFound() public {
    _setTokenFeePolicy(s_token, bytes32(0));

    vm.expectRevert(abi.encodeWithSelector(FeePolicyStore.FeeRecordNotFound.selector, s_token, DEST_CHAIN_SELECTOR));
    s_feeStore.checkLaneRecords(s_token, DEST_CHAIN_SELECTOR);

    // Chain selector zero names the policy default record, so it never resolves to a record.
    vm.expectRevert(
      abi.encodeWithSelector(FeePolicyStore.FeeRecordNotFound.selector, s_token, PolicyConfig.DEFAULT_CHAIN_SELECTOR)
    );
    s_feeStore.checkLaneRecords(s_token, PolicyConfig.DEFAULT_CHAIN_SELECTOR);
  }

  function test_checkLaneRecords_RevertWhen_DestOverheadRecordNotFound() public {
    _setTokenDestOverheadPolicy(s_token, bytes32(0));

    vm.expectRevert(
      abi.encodeWithSelector(FeePolicyStore.DestOverheadRecordNotFound.selector, s_token, DEST_CHAIN_SELECTOR)
    );
    s_feeStore.checkLaneRecords(s_token, DEST_CHAIN_SELECTOR);
  }
}
