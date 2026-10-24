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
import {MultiTokenPoolHooksExternalBlocklistSetup} from "./MultiTokenPoolHooksExternalBlocklistSetup.t.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

contract MultiTokenPoolHooksExternalBlocklist_preflightCheck is MultiTokenPoolHooksExternalBlocklistSetup {
  function test_preflightCheck() public {
    vm.startPrank(s_pool);

    vm.expectCall(address(s_blocklist), abi.encodeCall(IAddressBlocklist.isBlocked, (s_sender)), 1);
    vm.expectCall(address(s_blocklist), abi.encodeCall(IAddressBlocklist.isBlocked, (s_receiver)), 1);
    s_hooks.preflightCheck(s_lockOrBurnIn, FinalityCodec.WAIT_FOR_FINALITY_FLAG, "", 1e18);
  }

  function test_preflightCheck_EVMReceiverOf20Bytes() public {
    s_lockOrBurnIn.receiver = abi.encodePacked(s_receiver);
    vm.startPrank(s_pool);

    vm.expectCall(address(s_blocklist), abi.encodeCall(IAddressBlocklist.isBlocked, (s_receiver)), 1);
    s_hooks.preflightCheck(s_lockOrBurnIn, FinalityCodec.WAIT_FOR_FINALITY_FLAG, "", 1e18);
  }

  function test_preflightCheck_SolanaReceiver() public {
    s_lockOrBurnIn.receiver = abi.encodePacked(SOLANA_ACCOUNT);
    vm.startPrank(s_pool);

    vm.expectCall(address(s_blocklist), abi.encodeCall(IAddressBlocklist.isBlocked, (SOLANA_ACCOUNT_AS_ADDRESS)), 1);
    s_hooks.preflightCheck(s_lockOrBurnIn, FinalityCodec.WAIT_FOR_FINALITY_FLAG, "", 1e18);
  }

  function test_preflightCheck_ReceiverWithoutEVMForm() public {
    // Every address except the sender is blocked, so checking the receiver would revert.
    vm.mockCall(address(s_blocklist), abi.encodeWithSelector(IAddressBlocklist.isBlocked.selector), abi.encode(true));
    _setBlocked(s_sender, false);
    bytes[] memory receivers = _remoteAddressesWithoutEVMForm();
    vm.startPrank(s_pool);

    for (uint256 i = 0; i < receivers.length; ++i) {
      s_lockOrBurnIn.receiver = receivers[i];
      s_hooks.preflightCheck(s_lockOrBurnIn, FinalityCodec.WAIT_FOR_FINALITY_FLAG, "", 1e18);
    }
  }

  function test_preflightCheck_NoBlocklist() public {
    s_hooks.setBlocklist(IAddressBlocklist(address(0)));
    vm.mockCall(address(s_blocklist), abi.encodeWithSelector(IAddressBlocklist.isBlocked.selector), abi.encode(true));
    vm.startPrank(s_pool);

    vm.expectCall(address(s_blocklist), abi.encodeWithSelector(IAddressBlocklist.isBlocked.selector), 0);
    s_hooks.preflightCheck(s_lockOrBurnIn, FinalityCodec.WAIT_FOR_FINALITY_FLAG, "", 1e18);
  }

  function test_preflightCheck_WithPolicyEngine() public {
    s_hooks.setPolicyEngine(s_policyEngine);
    bytes memory tokenArgs = abi.encode("custom token args");
    vm.startPrank(s_pool);

    vm.expectCall(address(s_blocklist), abi.encodeCall(IAddressBlocklist.isBlocked, (s_receiver)), 1);
    s_hooks.preflightCheck(s_lockOrBurnIn, FinalityCodec.WAIT_FOR_FINALITY_FLAG, tokenArgs, 1e18);

    IPolicyEngine.Payload memory payload = s_policyEngine.getLastPayload();
    assertEq(IMultiTokenPoolHooks.preflightCheck.selector, payload.selector);
    assertEq(s_pool, payload.sender);
    assertEq(abi.encode(s_lockOrBurnIn, FinalityCodec.WAIT_FOR_FINALITY_FLAG, tokenArgs, 1e18), payload.data);
    assertEq(tokenArgs, payload.context);
  }

  function testFuzz_preflightCheck_ReceiverOf32Bytes(
    bytes32 receiver
  ) public {
    s_lockOrBurnIn.receiver = abi.encodePacked(receiver);
    // bytes20 keeps the leading bytes, so shift the low 20 bytes to the front first.
    address expectedAddress = address(bytes20(receiver << 96));
    vm.startPrank(s_pool);

    vm.expectCall(address(s_blocklist), abi.encodeCall(IAddressBlocklist.isBlocked, (expectedAddress)));
    s_hooks.preflightCheck(s_lockOrBurnIn, FinalityCodec.WAIT_FOR_FINALITY_FLAG, "", 1e18);
  }

  // Reverts

  function test_preflightCheck_RevertWhen_UnauthorizedCaller() public {
    // Owning the hooks does not make the owner an authorized caller.
    vm.expectCall(address(s_blocklist), abi.encodeWithSelector(IAddressBlocklist.isBlocked.selector), 0);
    vm.expectRevert(abi.encodeWithSelector(AuthorizedCallers.UnauthorizedCaller.selector, OWNER));
    s_hooks.preflightCheck(s_lockOrBurnIn, FinalityCodec.WAIT_FOR_FINALITY_FLAG, "", 1e18);
  }

  function test_preflightCheck_RevertWhen_SenderNotAllowed() public {
    MultiTokenPoolHooksExternalBlocklist hooks = new MultiTokenPoolHooksExternalBlocklist(
      _arrayOf(STRANGER), IPolicyEngine(address(0)), s_blocklist, _arrayOf(s_pool)
    );
    vm.startPrank(s_pool);

    vm.expectRevert(abi.encodeWithSelector(MultiTokenPoolHooks.SenderNotAllowed.selector, s_sender));
    hooks.preflightCheck(s_lockOrBurnIn, FinalityCodec.WAIT_FOR_FINALITY_FLAG, "", 1e18);
  }

  function test_preflightCheck_RevertWhen_LocalSenderBlocked() public {
    _setBlocked(s_sender, true);
    vm.startPrank(s_pool);

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
    vm.startPrank(s_pool);

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
    vm.startPrank(s_pool);

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
    vm.startPrank(s_pool);

    vm.expectRevert(revertData);
    s_hooks.preflightCheck(s_lockOrBurnIn, FinalityCodec.WAIT_FOR_FINALITY_FLAG, "", 1e18);
  }

  function test_preflightCheck_RevertWhen_BlocklistReturnsInvalidData() public {
    vm.mockCall(address(s_blocklist), abi.encodeCall(IAddressBlocklist.isBlocked, (s_sender)), hex"01");
    vm.startPrank(s_pool);

    // A single byte is too short to decode as a bool, so the call reverts without a reason.
    vm.expectRevert(bytes(""));
    s_hooks.preflightCheck(s_lockOrBurnIn, FinalityCodec.WAIT_FOR_FINALITY_FLAG, "", 1e18);
  }
}
