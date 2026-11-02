// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {SecurityPolicyStore} from "../../../../../pools/MultiTokenPool/policies/SecurityPolicyStore.sol";
import {SecurityPolicyStoreSetup} from "./SecurityPolicyStoreSetup.t.sol";

contract SecurityPolicyStore_checkLaneRecords is SecurityPolicyStoreSetup {
  /// @notice checkLaneRecords accepts a token with a rate limit record and a security record for the chain.
  function test_checkLaneRecords() public {
    _setInlineRateLimitConfig(s_tokenA, DEST_CHAIN_SELECTOR, _getRateLimitConfig(CAPACITY));
    _setInlineSecurityConfig(s_tokenA, DEST_CHAIN_SELECTOR, _arrayOf(s_ccv1), _arrayOf(s_ccv2), s_hook);
    s_store.checkLaneRecords(s_tokenA, DEST_CHAIN_SELECTOR);
  }

  // Reverts

  /// @notice checkLaneRecords reverts when the token has no rate limit record for the chain.
  function test_checkLaneRecords_RevertWhen_RateLimitRecordNotFound() public {
    _setInlineSecurityConfig(s_tokenA, DEST_CHAIN_SELECTOR, _arrayOf(s_ccv1), _arrayOf(s_ccv2), s_hook);

    vm.expectRevert(
      abi.encodeWithSelector(SecurityPolicyStore.RateLimitRecordNotFound.selector, s_tokenA, DEST_CHAIN_SELECTOR)
    );
    s_store.checkLaneRecords(s_tokenA, DEST_CHAIN_SELECTOR);
  }

  /// @notice checkLaneRecords reverts when the token has a rate limit record but no security record for the chain.
  function test_checkLaneRecords_RevertWhen_SecurityRecordNotFound() public {
    _setInlineRateLimitConfig(s_tokenA, DEST_CHAIN_SELECTOR, _getRateLimitConfig(CAPACITY));

    vm.expectRevert(
      abi.encodeWithSelector(SecurityPolicyStore.SecurityRecordNotFound.selector, s_tokenA, DEST_CHAIN_SELECTOR)
    );
    s_store.checkLaneRecords(s_tokenA, DEST_CHAIN_SELECTOR);
  }
}
