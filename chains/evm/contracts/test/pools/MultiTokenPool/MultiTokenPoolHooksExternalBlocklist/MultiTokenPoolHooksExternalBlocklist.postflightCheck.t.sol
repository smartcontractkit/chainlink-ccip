// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IAddressBlocklist} from "../../../../interfaces/IAddressBlocklist.sol";
import {IMultiTokenPoolHooks} from "../../../../interfaces/IMultiTokenPoolHooks.sol";
import {IPolicyEngine} from "@chainlink/policy-management/interfaces/IPolicyEngine.sol";

import {FinalityCodec} from "../../../../libraries/FinalityCodec.sol";
import {
  MultiTokenPoolHooksExternalBlocklist
} from "../../../../pools/MultiTokenPool/MultiTokenPoolHooksExternalBlocklist.sol";
import {MockPolicyEngine} from "../../../mocks/MockPolicyEngine.sol";
import {MultiTokenPoolHooksExternalBlocklistSetup} from "./MultiTokenPoolHooksExternalBlocklistSetup.t.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

contract MultiTokenPoolHooksExternalBlocklist_postflightCheck is MultiTokenPoolHooksExternalBlocklistSetup {
  function test_postflightCheck() public {
    vm.expectCall(address(s_blocklist), abi.encodeCall(IAddressBlocklist.isBlocked, (s_receiver)), 1);
    vm.expectCall(address(s_blocklist), abi.encodeCall(IAddressBlocklist.isBlocked, (s_sender)), 1);
    s_hooks.postflightCheck(s_releaseOrMintIn, 1e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG);
  }

  function test_postflightCheck_EVMSenderOf20Bytes() public {
    s_releaseOrMintIn.originalSender = abi.encodePacked(s_sender);

    vm.expectCall(address(s_blocklist), abi.encodeCall(IAddressBlocklist.isBlocked, (s_sender)), 1);
    s_hooks.postflightCheck(s_releaseOrMintIn, 1e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG);
  }

  function test_postflightCheck_SolanaSender() public {
    s_releaseOrMintIn.originalSender = abi.encodePacked(SOLANA_ACCOUNT);

    vm.expectCall(address(s_blocklist), abi.encodeCall(IAddressBlocklist.isBlocked, (SOLANA_ACCOUNT_AS_ADDRESS)), 1);
    s_hooks.postflightCheck(s_releaseOrMintIn, 1e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG);
  }

  function test_postflightCheck_SenderWithoutEVMForm() public {
    vm.mockCall(address(s_blocklist), abi.encodeWithSelector(IAddressBlocklist.isBlocked.selector), abi.encode(true));
    _setBlocked(s_receiver, false);
    bytes[] memory senders = _remoteAddressesWithoutEVMForm();

    for (uint256 i = 0; i < senders.length; ++i) {
      s_releaseOrMintIn.originalSender = senders[i];
      // Every address except the receiver is blocked, so checking the sender would revert.
      s_hooks.postflightCheck(s_releaseOrMintIn, 1e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG);
    }
  }

  function test_postflightCheck_NoBlocklist() public {
    MultiTokenPoolHooksExternalBlocklist hooks = new MultiTokenPoolHooksExternalBlocklist(
      new address[](0), IPolicyEngine(address(0)), IAddressBlocklist(address(0)), _arrayOf(OWNER)
    );
    vm.mockCall(address(s_blocklist), abi.encodeWithSelector(IAddressBlocklist.isBlocked.selector), abi.encode(true));

    vm.expectCall(address(s_blocklist), abi.encodeWithSelector(IAddressBlocklist.isBlocked.selector), 0);
    hooks.postflightCheck(s_releaseOrMintIn, 1e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG);
  }

  function test_postflightCheck_WithPolicyEngine() public {
    MockPolicyEngine policyEngine = new MockPolicyEngine();
    s_hooks.setPolicyEngine(policyEngine);
    s_releaseOrMintIn.offchainTokenData = abi.encode("custom offchain token data");

    vm.expectCall(address(s_blocklist), abi.encodeCall(IAddressBlocklist.isBlocked, (s_sender)), 1);
    s_hooks.postflightCheck(s_releaseOrMintIn, 1e6, FinalityCodec.WAIT_FOR_FINALITY_FLAG);

    IPolicyEngine.Payload memory payload = policyEngine.getLastPayload();
    assertEq(IMultiTokenPoolHooks.postflightCheck.selector, payload.selector);
    assertEq(abi.encode(s_releaseOrMintIn, 1e6, FinalityCodec.WAIT_FOR_FINALITY_FLAG), payload.data);
    assertEq(s_releaseOrMintIn.offchainTokenData, payload.context);
  }

  function testFuzz_postflightCheck_SenderOf32Bytes(
    bytes32 sender
  ) public {
    s_releaseOrMintIn.originalSender = abi.encodePacked(sender);
    // bytes20 keeps the leading bytes, so shift the low 20 bytes to the front first.
    address expectedAddress = address(bytes20(sender << 96));

    vm.expectCall(address(s_blocklist), abi.encodeCall(IAddressBlocklist.isBlocked, (expectedAddress)));
    s_hooks.postflightCheck(s_releaseOrMintIn, 1e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG);
  }

  // Reverts

  function test_postflightCheck_RevertWhen_UnauthorizedCaller() public {
    vm.startPrank(STRANGER);

    vm.expectCall(address(s_blocklist), abi.encodeWithSelector(IAddressBlocklist.isBlocked.selector), 0);
    vm.expectRevert(abi.encodeWithSelector(AuthorizedCallers.UnauthorizedCaller.selector, STRANGER));
    s_hooks.postflightCheck(s_releaseOrMintIn, 1e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG);
  }

  function test_postflightCheck_RevertWhen_LocalReceiverBlocked() public {
    _setBlocked(s_receiver, true);

    vm.expectRevert(
      abi.encodeWithSelector(MultiTokenPoolHooksExternalBlocklist.AddressBlocked.selector, s_blocklist, s_receiver)
    );
    s_hooks.postflightCheck(s_releaseOrMintIn, 1e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG);
  }

  function test_postflightCheck_RevertWhen_EVMSenderBlocked() public {
    _setBlocked(s_sender, true);
    bytes[] memory senders = new bytes[](2);
    senders[0] = abi.encode(s_sender);
    senders[1] = abi.encodePacked(s_sender);

    for (uint256 i = 0; i < senders.length; ++i) {
      s_releaseOrMintIn.originalSender = senders[i];
      vm.expectRevert(
        abi.encodeWithSelector(MultiTokenPoolHooksExternalBlocklist.AddressBlocked.selector, s_blocklist, s_sender)
      );
      s_hooks.postflightCheck(s_releaseOrMintIn, 1e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG);
    }
  }

  function test_postflightCheck_RevertWhen_SolanaSenderBlocked() public {
    _setBlocked(SOLANA_ACCOUNT_AS_ADDRESS, true);
    s_releaseOrMintIn.originalSender = abi.encodePacked(SOLANA_ACCOUNT);

    vm.expectRevert(
      abi.encodeWithSelector(
        MultiTokenPoolHooksExternalBlocklist.AddressBlocked.selector, s_blocklist, SOLANA_ACCOUNT_AS_ADDRESS
      )
    );
    s_hooks.postflightCheck(s_releaseOrMintIn, 1e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG);
  }

  function test_postflightCheck_RevertWhen_BlocklistReverts() public {
    bytes memory revertData = hex"deadbeef";
    vm.mockCallRevert(address(s_blocklist), abi.encodeCall(IAddressBlocklist.isBlocked, (s_sender)), revertData);

    vm.expectRevert(revertData);
    s_hooks.postflightCheck(s_releaseOrMintIn, 1e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG);
  }
}
