// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Client} from "../../../libraries/Client.sol";
import {Internal} from "../../../libraries/Internal.sol";
import {SuperchainInterop} from "../../../libraries/SuperchainInterop.sol";
import {OnRamp} from "../../../onRamp/OnRamp.sol";
import {OnRampOverSuperchainInterop} from "../../../onRamp/OnRampOverSuperchainInterop.sol";
import {OnRampOverSuperchainInteropSetup} from "./OnRampOverSuperchainInteropSetup.t.sol";

import {IERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/IERC20.sol";

contract OnRampOverSuperchainInterop_reemitInteropMessage is OnRampOverSuperchainInteropSetup {
  function _forwardFromRouter(Client.EVM2AnyMessage memory message, uint256 feeAmount) internal {
    // First send a basic message to populate the storage
    vm.stopPrank(); // Stop OWNER prank
    vm.startPrank(address(s_sourceRouter));

    IERC20(s_sourceFeeToken).transferFrom(OWNER, address(s_onRampOverSuperchainInterop), feeAmount);

    s_onRampOverSuperchainInterop.forwardFromRouter(DEST_CHAIN_SELECTOR, message, feeAmount, OWNER);
    vm.stopPrank();
  }

  function test_reemitInteropMessage_ReemitMessage() public {
    uint256 feeAmount = 1234567890;
    Client.EVM2AnyMessage memory message = _generateSingleTokenMessage(s_sourceFeeToken, 100);
    _forwardFromRouter(message, feeAmount);

    // Generate the correct Any2EVM message that matches what was sent
    (, Internal.Any2EVMRampMessage memory any2EvmMessage) =
      _generateInitialSourceDestMessages(DEST_CHAIN_SELECTOR, message, feeAmount);

    // Expect the CCIPSuperchainMessageSent event to be re-emitted
    vm.expectEmit();
    emit SuperchainInterop.CCIPSuperchainMessageSent(DEST_CHAIN_SELECTOR, 1, any2EvmMessage);

    s_onRampOverSuperchainInterop.reemitInteropMessage(any2EvmMessage);
  }

  function test_reemitInteropMessage_ReemitPTTMessageFromDifferentAccounts() public {
    uint256 feeAmount = 1234567890;
    Client.EVM2AnyMessage memory message = _generateSingleTokenMessage(s_sourceFeeToken, 1000);
    message.data = abi.encode("custom test data");
    message.extraArgs =
      Client._argsToBytes(Client.GenericExtraArgsV2({gasLimit: 500_000, allowOutOfOrderExecution: false}));

    _forwardFromRouter(message, feeAmount);

    // Generate the correct Any2EVM message that matches what was sent
    (, Internal.Any2EVMRampMessage memory any2EvmMessage) =
      _generateInitialSourceDestMessages(DEST_CHAIN_SELECTOR, message, feeAmount);

    vm.stopPrank();
    // Same message can be re-emitted multiple times, by different addresses at different blocks
    for (uint256 i = 0; i < 3; ++i) {
      vm.expectEmit();
      emit SuperchainInterop.CCIPSuperchainMessageSent(DEST_CHAIN_SELECTOR, 1, any2EvmMessage);

      vm.roll(block.number + 100);
      vm.warp(block.timestamp + 100);

      // Mock calling from a different account each time to validate reemitInteropMessage is permissionless
      vm.prank(makeAddr(string(abi.encode("reemitPTT", i))));
      s_onRampOverSuperchainInterop.reemitInteropMessage(any2EvmMessage);
    }
  }

  function test_reemitInteropMessage_ReemitAfterAllowlistChange() public {
    uint256 feeAmount = 1234567890;
    Client.EVM2AnyMessage memory message = _generateEmptyMessage();
    _forwardFromRouter(message, feeAmount);

    // Enable allowlist for the destination chain, since it is empty, no one should be able to send new messages.
    vm.stopPrank();
    vm.prank(OWNER);
    OnRamp.DestChainConfigArgs[] memory destChainConfigArgs = new OnRamp.DestChainConfigArgs[](1);
    destChainConfigArgs[0] = OnRamp.DestChainConfigArgs({
      destChainSelector: DEST_CHAIN_SELECTOR,
      router: s_sourceRouter,
      allowlistEnabled: true
    });
    s_onRampOverSuperchainInterop.applyDestChainConfigUpdates(destChainConfigArgs);

    // Generate the correct Any2EVM message that matches what was sent
    (, Internal.Any2EVMRampMessage memory any2EvmMessage) =
      _generateInitialSourceDestMessages(DEST_CHAIN_SELECTOR, message, feeAmount);

    // Expect a previous-sent message can be re-emitted even if the sender would be blocked by the current allowlist.
    vm.expectEmit();
    emit SuperchainInterop.CCIPSuperchainMessageSent(DEST_CHAIN_SELECTOR, 1, any2EvmMessage);

    s_onRampOverSuperchainInterop.reemitInteropMessage(any2EvmMessage);
  }

  // Reverts

  function test_reemitInteropMessage_RevertWhen_InvalidSourceChainSelector() public {
    Internal.Any2EVMRampMessage memory any2EvmMessage = _generateBasicAny2EVMMessage();
    // Change source chain selector to an invalid one
    any2EvmMessage.header.sourceChainSelector = SOURCE_CHAIN_SELECTOR + 1; // Wrong chain

    vm.expectRevert(
      abi.encodeWithSelector(OnRampOverSuperchainInterop.InvalidSourceChainSelector.selector, SOURCE_CHAIN_SELECTOR + 1)
    );

    s_onRampOverSuperchainInterop.reemitInteropMessage(any2EvmMessage);
  }

  function test_reemitInteropMessage_RevertWhen_MessageDoesNotExist() public {
    // Try to re-emit a message that was never sent
    Internal.Any2EVMRampMessage memory any2EvmMessage = _generateBasicAny2EVMMessage();

    // Calculate the hash from the OnRamp's perspective
    bytes32 expectedHash = SuperchainInterop._hashInteropMessage(any2EvmMessage, address(s_onRampOverSuperchainInterop));

    vm.expectRevert(
      abi.encodeWithSelector(
        OnRampOverSuperchainInterop.MessageDoesNotExist.selector, DEST_CHAIN_SELECTOR, 1, expectedHash
      )
    );

    s_onRampOverSuperchainInterop.reemitInteropMessage(any2EvmMessage);
  }

  function test_reemitInteropMessage_RevertWhen_WrongMessageHash() public {
    // First send a message to populate the storage
    vm.startPrank(address(s_sourceRouter));

    uint256 feeAmount = 1234567890;
    Client.EVM2AnyMessage memory message = _generateEmptyMessage();
    _forwardFromRouter(message, feeAmount);

    (, Internal.Any2EVMRampMessage memory any2EvmMessage) =
      _generateInitialSourceDestMessages(DEST_CHAIN_SELECTOR, message, feeAmount);

    // Modify the receiver to make the hash different
    any2EvmMessage.receiver = makeAddr("wrongReceiver");

    bytes32 expectedHash = SuperchainInterop._hashInteropMessage(any2EvmMessage, address(s_onRampOverSuperchainInterop));

    vm.expectRevert(
      abi.encodeWithSelector(
        OnRampOverSuperchainInterop.MessageDoesNotExist.selector,
        any2EvmMessage.header.destChainSelector,
        any2EvmMessage.header.sequenceNumber,
        expectedHash
      )
    );

    s_onRampOverSuperchainInterop.reemitInteropMessage(any2EvmMessage);
  }
}
