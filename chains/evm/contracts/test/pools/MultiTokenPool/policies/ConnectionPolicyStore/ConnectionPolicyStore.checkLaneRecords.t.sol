// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {PolicyConfig} from "../../../../../libraries/PolicyConfig.sol";
import {ConnectionPolicyStore} from "../../../../../pools/MultiTokenPool/policies/ConnectionPolicyStore.sol";
import {ConnectionPolicyStoreSetup} from "./ConnectionPolicyStoreSetup.t.sol";

contract ConnectionPolicyStore_checkLaneRecords is ConnectionPolicyStoreSetup {
  function test_checkLaneRecords() public view {
    s_connectionStore.checkLaneRecords(s_token, DEST_CHAIN_SELECTOR);
  }

  function test_checkLaneRecords_InlineRecordWithoutLane() public {
    // Remote pool records can be checked before a connection is configured, regardless of pause state.
    _setInlineRemotePoolConfig(
      _getInlineRemotePoolConfigArgs(s_otherToken, DEST_CHAIN_SELECTOR, _getRemotePoolConfig())
    );
    s_connectionStore.setGlobalPause(true);

    s_connectionStore.checkLaneRecords(s_otherToken, DEST_CHAIN_SELECTOR);
  }

  // Reverts

  function test_checkLaneRecords_RevertWhen_RemotePoolRecordNotFound() public {
    vm.expectRevert(
      abi.encodeWithSelector(ConnectionPolicyStore.RemotePoolRecordNotFound.selector, s_otherToken, DEST_CHAIN_SELECTOR)
    );
    s_connectionStore.checkLaneRecords(s_otherToken, DEST_CHAIN_SELECTOR);

    // Remote pool policies have no default record at chain selector zero.
    vm.expectRevert(
      abi.encodeWithSelector(
        ConnectionPolicyStore.RemotePoolRecordNotFound.selector, s_token, PolicyConfig.DEFAULT_CHAIN_SELECTOR
      )
    );
    s_connectionStore.checkLaneRecords(s_token, PolicyConfig.DEFAULT_CHAIN_SELECTOR);
  }
}
