// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Client} from "../../../libraries/Client.sol";
import {Internal} from "../../../libraries/Internal.sol";
import {SuperchainInterop} from "../../../libraries/SuperchainInterop.sol";
import {OnRampOverSuperchainInterop} from "../../../onRamp/OnRampOverSuperchainInterop.sol";
import {OnRampOverSuperchainInteropSetup} from "./OnRampOverSuperchainInteropSetup.t.sol";

import {IERC20} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";

contract OnRampOverSuperchainInterop_reemitInteropMessage is OnRampOverSuperchainInteropSetup {
  function test_ReemitBasicMessage() public {
    // First send a basic message to populate the storage
    vm.startPrank(address(s_sourceRouter));

    Client.EVM2AnyMessage memory message = _generateEmptyMessage();
    uint256 feeAmount = 1234567890;
    IERC20(s_sourceFeeToken).transferFrom(OWNER, address(s_onRampOverSuperchainInterop), feeAmount);

    s_onRampOverSuperchainInterop.forwardFromRouter(DEST_CHAIN_SELECTOR, message, feeAmount, OWNER);
    vm.stopPrank();

    // Generate the correct Any2EVM message that matches what was sent
    (, Internal.Any2EVMRampMessage memory any2EvmMessage) =
      _generateInitialSourceDestMessages(DEST_CHAIN_SELECTOR, message, feeAmount);

    // Expect the CCIPSuperchainMessageSent event to be re-emitted
    vm.expectEmit();
    emit SuperchainInterop.CCIPSuperchainMessageSent(DEST_CHAIN_SELECTOR, 1, any2EvmMessage);

    s_onRampOverSuperchainInterop.reemitInteropMessage(any2EvmMessage);
  }

  function test_ReemitTokenMessage() public {
    // First send a message with tokens to populate the storage
    vm.startPrank(address(s_sourceRouter));

    Client.EVM2AnyMessage memory message = _generateSingleTokenMessage(s_sourceFeeToken, 100);
    uint256 feeAmount = 1234567890;
    IERC20(s_sourceFeeToken).transferFrom(OWNER, address(s_onRampOverSuperchainInterop), feeAmount);

    s_onRampOverSuperchainInterop.forwardFromRouter(DEST_CHAIN_SELECTOR, message, feeAmount, OWNER);
    vm.stopPrank();

    // Generate the correct Any2EVM message that matches what was sent
    (, Internal.Any2EVMRampMessage memory any2EvmMessage) =
      _generateInitialSourceDestMessages(DEST_CHAIN_SELECTOR, message, feeAmount);

    // Expect the CCIPSuperchainMessageSent event to be re-emitted
    vm.expectEmit();
    emit SuperchainInterop.CCIPSuperchainMessageSent(DEST_CHAIN_SELECTOR, 1, any2EvmMessage);

    s_onRampOverSuperchainInterop.reemitInteropMessage(any2EvmMessage);
  }

  function test_RepeatReemitPTTMessage() public {
    // First send a message with tokens and data to populate the storage
    vm.startPrank(address(s_sourceRouter));

    Client.EVM2AnyMessage memory message = _generateSingleTokenMessage(s_sourceFeeToken, 1000);
    message.data = abi.encode("custom test data");
    message.extraArgs =
      Client._argsToBytes(Client.GenericExtraArgsV2({gasLimit: 500_000, allowOutOfOrderExecution: false}));

    uint256 feeAmount = 2000000000;
    IERC20(s_sourceFeeToken).transferFrom(OWNER, address(s_onRampOverSuperchainInterop), feeAmount);

    s_onRampOverSuperchainInterop.forwardFromRouter(DEST_CHAIN_SELECTOR, message, feeAmount, OWNER);
    vm.stopPrank();

    // Generate the correct Any2EVM message that matches what was sent
    (, Internal.Any2EVMRampMessage memory any2EvmMessage) =
      _generateInitialSourceDestMessages(DEST_CHAIN_SELECTOR, message, feeAmount);

    // Same message can be re-emitted multiple times, by any address
    for (uint256 i = 0; i < 3; i++) {
      vm.expectEmit();
      emit SuperchainInterop.CCIPSuperchainMessageSent(DEST_CHAIN_SELECTOR, 1, any2EvmMessage);

      // Mock calling from a different account each time to validate reemitInteropMessage is permissionless
      vm.prank(makeAddr(string(abi.encode("caller", i))));
      s_onRampOverSuperchainInterop.reemitInteropMessage(any2EvmMessage);
    }
  }

  // Reverts

  function test_RevertWhen_InvalidSourceChainSelector() public {
    Internal.Any2EVMRampMessage memory any2EvmMessage = _generateBasicAny2EVMMessage();
    // Change source chain selector to an invalid one
    any2EvmMessage.header.sourceChainSelector = SOURCE_CHAIN_SELECTOR + 1; // Wrong chain

    vm.expectRevert(
      abi.encodeWithSelector(OnRampOverSuperchainInterop.InvalidSourceChainSelector.selector, SOURCE_CHAIN_SELECTOR + 1)
    );

    s_onRampOverSuperchainInterop.reemitInteropMessage(any2EvmMessage);
  }

  function test_RevertWhen_MessageDoesNotExist() public {
    // Try to re-emit a message that was never sent
    Internal.Any2EVMRampMessage memory any2EvmMessage = _generateBasicAny2EVMMessage();

    bytes32 expectedHash = SuperchainInterop._hashInteropMessage(any2EvmMessage);

    vm.expectRevert(
      abi.encodeWithSelector(
        OnRampOverSuperchainInterop.MessageDoesNotExist.selector, DEST_CHAIN_SELECTOR, 1, expectedHash
      )
    );

    s_onRampOverSuperchainInterop.reemitInteropMessage(any2EvmMessage);
  }

  function test_RevertWhen_WrongMessageHash() public {
    // First send a message to populate the storage
    vm.startPrank(address(s_sourceRouter));

    Client.EVM2AnyMessage memory message = _generateEmptyMessage();
    uint256 feeAmount = 1234567890;
    IERC20(s_sourceFeeToken).transferFrom(OWNER, address(s_onRampOverSuperchainInterop), feeAmount);

    s_onRampOverSuperchainInterop.forwardFromRouter(DEST_CHAIN_SELECTOR, message, feeAmount, OWNER);
    vm.stopPrank();

    Internal.Any2EVMRampMessage[] memory any2EvmMessages = new Internal.Any2EVMRampMessage[](3);
    for (uint256 i = 0; i < 3; i++) {
      any2EvmMessages[i] = _generateBasicAny2EVMMessage();
    }

    // Modify the receiver to make the hash different
    any2EvmMessages[0].receiver = makeAddr("wrongReceiver");
    // Wrong sequence number
    any2EvmMessages[1].header.sequenceNumber = 100;
    // Wrong destination chain selector
    any2EvmMessages[1].header.destChainSelector = DEST_CHAIN_SELECTOR + 1;

    for (uint256 i = 0; i < 3; i++) {
      bytes32 expectedHash = SuperchainInterop._hashInteropMessage(any2EvmMessages[i]);

      vm.expectRevert(
        abi.encodeWithSelector(
          OnRampOverSuperchainInterop.MessageDoesNotExist.selector,
          any2EvmMessages[i].header.destChainSelector,
          any2EvmMessages[i].header.sequenceNumber,
          expectedHash
        )
      );

      s_onRampOverSuperchainInterop.reemitInteropMessage(any2EvmMessages[i]);
    }
  }
}
