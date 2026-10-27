// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IAddressBlocklist} from "../../../../interfaces/IAddressBlocklist.sol";
import {IMultiTokenPoolHooks} from "../../../../interfaces/IMultiTokenPoolHooks.sol";
import {IPolicyEngine} from "@chainlink/policy-management/interfaces/IPolicyEngine.sol";

import {FinalityCodec} from "../../../../libraries/FinalityCodec.sol";
import {
  MultiTokenPoolHooksExternalBlocklist
} from "../../../../pools/MultiTokenPool/MultiTokenPoolHooksExternalBlocklist.sol";
import {MultiTokenPoolHooksExternalBlocklistSetup} from "./MultiTokenPoolHooksExternalBlocklistSetup.t.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

contract MultiTokenPoolHooksExternalBlocklist_postflightCheck is MultiTokenPoolHooksExternalBlocklistSetup {
  function test_postflightCheck() public {
    vm.startPrank(s_pool);

    vm.expectCall(address(s_blocklist), abi.encodeCall(IAddressBlocklist.isBlocked, (s_receiver)), 1);
    vm.expectCall(address(s_blocklist), abi.encodeCall(IAddressBlocklist.isBlocked, (s_sender)), 1);
    s_hooks.postflightCheck(s_releaseOrMintIn, 1e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG);
  }

  function test_postflightCheck_EVMSenderOf20Bytes() public {
    s_releaseOrMintIn.originalSender = abi.encodePacked(s_sender);
    vm.startPrank(s_pool);

    vm.expectCall(address(s_blocklist), abi.encodeCall(IAddressBlocklist.isBlocked, (s_sender)), 1);
    s_hooks.postflightCheck(s_releaseOrMintIn, 1e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG);
  }

  function test_postflightCheck_SolanaSender() public {
    s_releaseOrMintIn.originalSender = abi.encodePacked(SOLANA_ACCOUNT);
    vm.startPrank(s_pool);

    vm.expectCall(address(s_blocklist), abi.encodeCall(IAddressBlocklist.isBlocked, (SOLANA_ACCOUNT_AS_ADDRESS)), 1);
    s_hooks.postflightCheck(s_releaseOrMintIn, 1e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG);
  }

  function test_postflightCheck_SenderWithoutEVMForm() public {
    // Every address except the receiver is blocked, so checking the sender would revert.
    vm.mockCall(address(s_blocklist), abi.encodeWithSelector(IAddressBlocklist.isBlocked.selector), abi.encode(true));
    _setBlocked(s_receiver, false);
    bytes[] memory senders = _remoteAddressesWithoutEVMForm();
    vm.startPrank(s_pool);

    for (uint256 i = 0; i < senders.length; ++i) {
      s_releaseOrMintIn.originalSender = senders[i];
      s_hooks.postflightCheck(s_releaseOrMintIn, 1e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG);
    }
  }

  function test_postflightCheck_NoBlocklist() public {
    s_hooks.setBlocklist(IAddressBlocklist(address(0)));
    vm.mockCall(address(s_blocklist), abi.encodeWithSelector(IAddressBlocklist.isBlocked.selector), abi.encode(true));
    vm.startPrank(s_pool);

    vm.expectCall(address(s_blocklist), abi.encodeWithSelector(IAddressBlocklist.isBlocked.selector), 0);
    s_hooks.postflightCheck(s_releaseOrMintIn, 1e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG);
  }

  function test_postflightCheck_WithPolicyEngine() public {
    s_hooks.setPolicyEngine(s_policyEngine);
    s_releaseOrMintIn.offchainTokenData = abi.encode("custom offchain token data");
    vm.startPrank(s_pool);

    vm.expectCall(address(s_blocklist), abi.encodeCall(IAddressBlocklist.isBlocked, (s_sender)), 1);
    s_hooks.postflightCheck(s_releaseOrMintIn, 1e6, FinalityCodec.WAIT_FOR_FINALITY_FLAG);

    IPolicyEngine.Payload memory payload = s_policyEngine.getLastPayload();
    assertEq(IMultiTokenPoolHooks.postflightCheck.selector, payload.selector);
    assertEq(s_pool, payload.sender);
    assertEq(abi.encode(s_releaseOrMintIn, 1e6, FinalityCodec.WAIT_FOR_FINALITY_FLAG), payload.data);
    assertEq(s_releaseOrMintIn.offchainTokenData, payload.context);
  }

  // Reverts

  function test_postflightCheck_RevertWhen_UnauthorizedCaller() public {
    // Owning the hooks does not make the owner an authorized caller.
    vm.expectCall(address(s_blocklist), abi.encodeWithSelector(IAddressBlocklist.isBlocked.selector), 0);
    vm.expectRevert(abi.encodeWithSelector(AuthorizedCallers.UnauthorizedCaller.selector, OWNER));
    s_hooks.postflightCheck(s_releaseOrMintIn, 1e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG);
  }

  function test_postflightCheck_RevertWhen_LocalReceiverBlocked() public {
    _setBlocked(s_receiver, true);
    vm.startPrank(s_pool);

    vm.expectRevert(
      abi.encodeWithSelector(MultiTokenPoolHooksExternalBlocklist.LocalAddressBlocked.selector, s_blocklist, s_receiver)
    );
    s_hooks.postflightCheck(s_releaseOrMintIn, 1e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG);
  }

  function test_postflightCheck_RevertWhen_EVMSenderBlocked() public {
    _setBlocked(s_sender, true);
    bytes[] memory senders = new bytes[](2);
    senders[0] = abi.encode(s_sender);
    senders[1] = abi.encodePacked(s_sender);
    vm.startPrank(s_pool);

    for (uint256 i = 0; i < senders.length; ++i) {
      s_releaseOrMintIn.originalSender = senders[i];
      vm.expectRevert(
        abi.encodeWithSelector(
          MultiTokenPoolHooksExternalBlocklist.RemoteAddressBlocked.selector, s_blocklist, senders[i]
        )
      );
      s_hooks.postflightCheck(s_releaseOrMintIn, 1e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG);
    }
  }

  function test_postflightCheck_RevertWhen_SolanaSenderBlocked() public {
    _setBlocked(SOLANA_ACCOUNT_AS_ADDRESS, true);
    s_releaseOrMintIn.originalSender = abi.encodePacked(SOLANA_ACCOUNT);
    vm.startPrank(s_pool);

    vm.expectRevert(
      abi.encodeWithSelector(
        MultiTokenPoolHooksExternalBlocklist.RemoteAddressBlocked.selector,
        s_blocklist,
        s_releaseOrMintIn.originalSender
      )
    );
    s_hooks.postflightCheck(s_releaseOrMintIn, 1e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG);
  }

  function test_postflightCheck_RevertWhen_BlocklistReverts() public {
    bytes memory revertData = hex"deadbeef";
    vm.mockCallRevert(address(s_blocklist), abi.encodeCall(IAddressBlocklist.isBlocked, (s_sender)), revertData);
    vm.startPrank(s_pool);

    vm.expectRevert(revertData);
    s_hooks.postflightCheck(s_releaseOrMintIn, 1e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG);
  }

  function test_postflightCheck_RevertWhen_BlocklistReturnsInvalidData() public {
    vm.mockCall(address(s_blocklist), abi.encodeCall(IAddressBlocklist.isBlocked, (s_sender)), hex"01");
    vm.startPrank(s_pool);

    // A single byte is too short to decode as a bool, so the call reverts without a reason.
    vm.expectRevert(bytes(""));
    s_hooks.postflightCheck(s_releaseOrMintIn, 1e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG);
  }
}
