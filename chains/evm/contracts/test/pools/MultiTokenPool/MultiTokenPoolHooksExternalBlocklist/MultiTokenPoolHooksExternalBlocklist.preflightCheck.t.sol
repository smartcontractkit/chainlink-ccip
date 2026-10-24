// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IAddressBlocklist} from "../../../../interfaces/IAddressBlocklist.sol";
import {IMultiTokenPoolHooks} from "../../../../interfaces/IMultiTokenPoolHooks.sol";
import {IPolicyEngine} from "@chainlink/policy-management/interfaces/IPolicyEngine.sol";

import {FinalityCodec} from "../../../../libraries/FinalityCodec.sol";
import {MultiTokenPoolHooks} from "../../../../pools/MultiTokenPool/MultiTokenPoolHooks.sol";
import {
  MultiTokenPoolHooksExternalBlocklist
} from "../../../../pools/MultiTokenPool/MultiTokenPoolHooksExternalBlocklist.sol";
import {MockPolicyEngine} from "../../../mocks/MockPolicyEngine.sol";
import {MultiTokenPoolHooksExternalBlocklistSetup} from "./MultiTokenPoolHooksExternalBlocklistSetup.t.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

contract MultiTokenPoolHooksExternalBlocklist_preflightCheck is MultiTokenPoolHooksExternalBlocklistSetup {
  function test_preflightCheck() public {
    vm.expectCall(address(s_blocklist), abi.encodeCall(IAddressBlocklist.isBlocked, (s_sender)), 1);
    vm.expectCall(address(s_blocklist), abi.encodeCall(IAddressBlocklist.isBlocked, (s_receiver)), 1);
    s_hooks.preflightCheck(s_lockOrBurnIn, FinalityCodec.WAIT_FOR_FINALITY_FLAG, "", 1e18);
  }

  function test_preflightCheck_EVMReceiverOf20Bytes() public {
    s_lockOrBurnIn.receiver = abi.encodePacked(s_receiver);

    vm.expectCall(address(s_blocklist), abi.encodeCall(IAddressBlocklist.isBlocked, (s_receiver)), 1);
    s_hooks.preflightCheck(s_lockOrBurnIn, FinalityCodec.WAIT_FOR_FINALITY_FLAG, "", 1e18);
  }

  function test_preflightCheck_SolanaReceiver() public {
    s_lockOrBurnIn.receiver = abi.encodePacked(SOLANA_ACCOUNT);

    vm.expectCall(address(s_blocklist), abi.encodeCall(IAddressBlocklist.isBlocked, (SOLANA_ACCOUNT_AS_ADDRESS)), 1);
    s_hooks.preflightCheck(s_lockOrBurnIn, FinalityCodec.WAIT_FOR_FINALITY_FLAG, "", 1e18);
  }

  function test_preflightCheck_ReceiverWithoutEVMForm() public {
    vm.mockCall(address(s_blocklist), abi.encodeWithSelector(IAddressBlocklist.isBlocked.selector), abi.encode(true));
    _setBlocked(s_sender, false);
    bytes[] memory receivers = _remoteAddressesWithoutEVMForm();

    for (uint256 i = 0; i < receivers.length; ++i) {
      s_lockOrBurnIn.receiver = receivers[i];
      // Every address but the sender is blocked, so a call for the receiver would revert.
      s_hooks.preflightCheck(s_lockOrBurnIn, FinalityCodec.WAIT_FOR_FINALITY_FLAG, "", 1e18);
    }
  }

  function test_preflightCheck_NoBlocklist() public {
    MultiTokenPoolHooksExternalBlocklist hooks = new MultiTokenPoolHooksExternalBlocklist(
      new address[](0), IPolicyEngine(address(0)), IAddressBlocklist(address(0)), _arrayOf(OWNER)
    );
    vm.mockCall(address(s_blocklist), abi.encodeWithSelector(IAddressBlocklist.isBlocked.selector), abi.encode(true));

    vm.expectCall(address(s_blocklist), abi.encodeWithSelector(IAddressBlocklist.isBlocked.selector), 0);
    hooks.preflightCheck(s_lockOrBurnIn, FinalityCodec.WAIT_FOR_FINALITY_FLAG, "", 1e18);
  }

  function test_preflightCheck_WithAllowListAndPolicyEngine() public {
    MockPolicyEngine policyEngine = new MockPolicyEngine();
    MultiTokenPoolHooksExternalBlocklist hooks =
      new MultiTokenPoolHooksExternalBlocklist(_arrayOf(s_sender), policyEngine, s_blocklist, _arrayOf(OWNER));
    bytes memory tokenArgs = abi.encode("custom token args");

    vm.expectCall(address(s_blocklist), abi.encodeCall(IAddressBlocklist.isBlocked, (s_receiver)), 1);
    hooks.preflightCheck(s_lockOrBurnIn, FinalityCodec.WAIT_FOR_FINALITY_FLAG, tokenArgs, 1e18);

    IPolicyEngine.Payload memory payload = policyEngine.getLastPayload();
    assertEq(IMultiTokenPoolHooks.preflightCheck.selector, payload.selector);
    assertEq(abi.encode(s_lockOrBurnIn, FinalityCodec.WAIT_FOR_FINALITY_FLAG, tokenArgs, 1e18), payload.data);
    assertEq(tokenArgs, payload.context);
  }

  function testFuzz_preflightCheck_ReceiverOf32Bytes(
    bytes32 receiver
  ) public {
    s_lockOrBurnIn.receiver = abi.encodePacked(receiver);
    // The low 20 bytes, moved to the top so that bytes20 takes them.
    address expectedAddress = address(bytes20(receiver << 96));

    vm.expectCall(address(s_blocklist), abi.encodeCall(IAddressBlocklist.isBlocked, (expectedAddress)));
    s_hooks.preflightCheck(s_lockOrBurnIn, FinalityCodec.WAIT_FOR_FINALITY_FLAG, "", 1e18);
  }

  // Reverts

  function test_preflightCheck_RevertWhen_UnauthorizedCaller() public {
    vm.startPrank(STRANGER);

    vm.expectCall(address(s_blocklist), abi.encodeWithSelector(IAddressBlocklist.isBlocked.selector), 0);
    vm.expectRevert(abi.encodeWithSelector(AuthorizedCallers.UnauthorizedCaller.selector, STRANGER));
    s_hooks.preflightCheck(s_lockOrBurnIn, FinalityCodec.WAIT_FOR_FINALITY_FLAG, "", 1e18);
  }

  function test_preflightCheck_RevertWhen_SenderNotAllowed() public {
    MultiTokenPoolHooksExternalBlocklist hooks = new MultiTokenPoolHooksExternalBlocklist(
      _arrayOf(STRANGER), IPolicyEngine(address(0)), s_blocklist, _arrayOf(OWNER)
    );

    vm.expectRevert(abi.encodeWithSelector(MultiTokenPoolHooks.SenderNotAllowed.selector, s_sender));
    hooks.preflightCheck(s_lockOrBurnIn, FinalityCodec.WAIT_FOR_FINALITY_FLAG, "", 1e18);
  }

  function test_preflightCheck_RevertWhen_LocalSenderBlocked() public {
    _setBlocked(s_sender, true);

    vm.expectRevert(
      abi.encodeWithSelector(MultiTokenPoolHooksExternalBlocklist.AddressBlocked.selector, s_blocklist, s_sender)
    );
    s_hooks.preflightCheck(s_lockOrBurnIn, FinalityCodec.WAIT_FOR_FINALITY_FLAG, "", 1e18);
  }

  function test_preflightCheck_RevertWhen_EVMReceiverBlocked() public {
    _setBlocked(s_receiver, true);
    bytes[] memory receivers = new bytes[](2);
    receivers[0] = abi.encode(s_receiver);
    receivers[1] = abi.encodePacked(s_receiver);

    for (uint256 i = 0; i < receivers.length; ++i) {
      s_lockOrBurnIn.receiver = receivers[i];
      vm.expectRevert(
        abi.encodeWithSelector(MultiTokenPoolHooksExternalBlocklist.AddressBlocked.selector, s_blocklist, s_receiver)
      );
      s_hooks.preflightCheck(s_lockOrBurnIn, FinalityCodec.WAIT_FOR_FINALITY_FLAG, "", 1e18);
    }
  }

  function test_preflightCheck_RevertWhen_SolanaReceiverBlocked() public {
    _setBlocked(SOLANA_ACCOUNT_AS_ADDRESS, true);
    s_lockOrBurnIn.receiver = abi.encodePacked(SOLANA_ACCOUNT);

    vm.expectRevert(
      abi.encodeWithSelector(
        MultiTokenPoolHooksExternalBlocklist.AddressBlocked.selector, s_blocklist, SOLANA_ACCOUNT_AS_ADDRESS
      )
    );
    s_hooks.preflightCheck(s_lockOrBurnIn, FinalityCodec.WAIT_FOR_FINALITY_FLAG, "", 1e18);
  }

  function test_preflightCheck_RevertWhen_BlocklistReverts() public {
    bytes memory revertData = hex"deadbeef";
    vm.mockCallRevert(address(s_blocklist), abi.encodeCall(IAddressBlocklist.isBlocked, (s_sender)), revertData);

    vm.expectRevert(revertData);
    s_hooks.preflightCheck(s_lockOrBurnIn, FinalityCodec.WAIT_FOR_FINALITY_FLAG, "", 1e18);
  }

  function test_preflightCheck_RevertWhen_BlocklistReturnsInvalidData() public {
    vm.mockCall(address(s_blocklist), abi.encodeCall(IAddressBlocklist.isBlocked, (s_sender)), hex"01");

    // The answer is shorter than one ABI word, so decoding it reverts without data.
    vm.expectRevert(bytes(""));
    s_hooks.preflightCheck(s_lockOrBurnIn, FinalityCodec.WAIT_FOR_FINALITY_FLAG, "", 1e18);
  }
}
