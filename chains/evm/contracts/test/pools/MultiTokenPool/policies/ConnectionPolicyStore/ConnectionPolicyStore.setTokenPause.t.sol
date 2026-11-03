// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {FinalityCodec} from "../../../../../libraries/FinalityCodec.sol";
import {PolicyConfig} from "../../../../../libraries/PolicyConfig.sol";
import {ConnectionPolicyStore} from "../../../../../pools/MultiTokenPool/policies/ConnectionPolicyStore.sol";
import {ConnectionPolicyStoreSetup} from "./ConnectionPolicyStoreSetup.t.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

import {Vm} from "forge-std/Vm.sol";

contract ConnectionPolicyStore_setTokenPause is ConnectionPolicyStoreSetup {
  function test_setTokenPause() public {
    vm.expectEmit();
    emit ConnectionPolicyStore.TokenPauseSet(s_token, true, false);
    s_connectionStore.setTokenPause(_arrayOf(s_token), ConnectionPolicyStore.PauseDirection.Outbound, true);

    _assertTokenPause(s_token, true, false);
    vm.expectRevert(
      abi.encodeWithSelector(
        ConnectionPolicyStore.TokenIsPaused.selector, s_token, ConnectionPolicyStore.PauseDirection.Outbound
      )
    );
    s_connectionStore.resolveOutboundLane(s_token, DEST_CHAIN_SELECTOR, FinalityCodec.WAIT_FOR_FINALITY_FLAG, "");
  }

  function test_setTokenPause_Inbound() public {
    vm.expectEmit();
    emit ConnectionPolicyStore.TokenPauseSet(s_token, false, true);
    s_connectionStore.setTokenPause(_arrayOf(s_token), ConnectionPolicyStore.PauseDirection.Inbound, true);

    _assertTokenPause(s_token, false, true);
  }

  function test_setTokenPause_Both() public {
    vm.expectEmit();
    emit ConnectionPolicyStore.TokenPauseSet(s_token, true, true);
    s_connectionStore.setTokenPause(_arrayOf(s_token), ConnectionPolicyStore.PauseDirection.Both, true);
    _assertTokenPause(s_token, true, true);

    vm.expectEmit();
    emit ConnectionPolicyStore.TokenPauseSet(s_token, false, false);
    s_connectionStore.setTokenPause(_arrayOf(s_token), ConnectionPolicyStore.PauseDirection.Both, false);
    _assertTokenPause(s_token, false, false);
  }

  function test_setTokenPause_OtherDirectionStays() public {
    s_connectionStore.setTokenPause(_arrayOf(s_token), ConnectionPolicyStore.PauseDirection.Outbound, true);

    vm.expectEmit();
    emit ConnectionPolicyStore.TokenPauseSet(s_token, true, true);
    s_connectionStore.setTokenPause(_arrayOf(s_token), ConnectionPolicyStore.PauseDirection.Inbound, true);

    vm.expectEmit();
    emit ConnectionPolicyStore.TokenPauseSet(s_token, false, true);
    s_connectionStore.setTokenPause(_arrayOf(s_token), ConnectionPolicyStore.PauseDirection.Outbound, false);

    _assertTokenPause(s_token, false, true);
  }

  function test_setTokenPause_NoChange() public {
    s_connectionStore.setTokenPause(_arrayOf(s_token), ConnectionPolicyStore.PauseDirection.Outbound, true);
    address[] memory localTokens = new address[](2);
    localTokens[0] = s_token;
    localTokens[1] = s_otherToken;

    // Only the token whose pause state changes should emit an event.
    vm.recordLogs();
    s_connectionStore.setTokenPause(localTokens, ConnectionPolicyStore.PauseDirection.Outbound, true);

    Vm.Log[] memory logs = vm.getRecordedLogs();
    assertEq(1, logs.length);
    assertEq(ConnectionPolicyStore.TokenPauseSet.selector, logs[0].topics[0]);
    assertEq(bytes32(uint256(uint160(s_otherToken))), logs[0].topics[1]);

    vm.recordLogs();
    s_connectionStore.setTokenPause(_arrayOf(s_token), ConnectionPolicyStore.PauseDirection.Inbound, false);
    assertEq(0, vm.getRecordedLogs().length);
  }

  function test_setTokenPause_BeforeRegistration() public {
    address newToken = address(_deployToken("NEW", DEFAULT_TOKEN_DECIMALS));

    vm.expectEmit();
    emit ConnectionPolicyStore.TokenPauseSet(newToken, true, true);
    s_connectionStore.setTokenPause(_arrayOf(newToken), ConnectionPolicyStore.PauseDirection.Both, true);

    // Registration and deregistration must preserve pause flags.
    _registerToken(newToken);
    _assertTokenPause(newToken, true, true);
    s_connectionStore.applyTokenProfileUpdates(_arrayOf(newToken), new PolicyConfig.TokenProfileArgs[](0));
    _assertTokenPause(newToken, true, true);
  }

  // Reverts

  function test_setTokenPause_RevertWhen_UnauthorizedCaller() public {
    vm.startPrank(STRANGER);

    vm.expectRevert(abi.encodeWithSelector(AuthorizedCallers.UnauthorizedCaller.selector, STRANGER));
    s_connectionStore.setTokenPause(_arrayOf(s_token), ConnectionPolicyStore.PauseDirection.Both, true);
  }

  function _assertTokenPause(
    address localToken,
    bool isOutboundPaused,
    bool isInboundPaused
  ) private view {
    ConnectionPolicyStore.TokenProfile memory profile = s_connectionStore.getTokenProfiles(_arrayOf(localToken))[0];
    assertEq(isOutboundPaused, profile.isOutboundPaused);
    assertEq(isInboundPaused, profile.isInboundPaused);
  }
}
