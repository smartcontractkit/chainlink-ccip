// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Client} from "../../../libraries/Client.sol";
import {Internal} from "../../../libraries/Internal.sol";
import {SuperchainInterop} from "../../../libraries/SuperchainInterop.sol";
import {OnRampOverSuperchainInteropSetup} from "./OnRampOverSuperchainInteropSetup.t.sol";

contract OnRampOverSuperchainInterop_postProcessMessage is OnRampOverSuperchainInteropSetup {
  function test_postProcessMessage() public {
    Client.EVM2AnyMessage memory message = _generateSingleTokenMessage(s_sourceFeeToken, 100);
    message.data = "test message data";
    uint256 feeAmount = 1234567890;

    (Internal.EVM2AnyRampMessage memory evm2AnyMessage, Internal.Any2EVMRampMessage memory any2EvmMessage) =
      _generateInitialSourceDestMessages(DEST_CHAIN_SELECTOR, message, feeAmount);
    evm2AnyMessage.header.messageId = "";

    bytes memory originalMessage = abi.encode(evm2AnyMessage);

    vm.expectEmit();
    emit SuperchainInterop.CCIPSuperchainMessageSent(DEST_CHAIN_SELECTOR, 1, any2EvmMessage);

    Internal.EVM2AnyRampMessage memory hookResult = s_onRampOverSuperchainInterop.postProcessMessage(evm2AnyMessage);
    assertEq(originalMessage, abi.encode(hookResult));
  }

  // Reverts

  function test_postProcessMessage_RevertWhen_InvalidDestTokenAddress() public {
    Client.EVM2AnyMessage memory message = _generateSingleTokenMessage(s_sourceFeeToken, 100);
    uint256 feeAmount = 1234567890;

    (Internal.EVM2AnyRampMessage memory evm2AnyMessage,) =
      _generateInitialSourceDestMessages(DEST_CHAIN_SELECTOR, message, feeAmount);
    evm2AnyMessage.header.messageId = "";

    bytes memory encodePackedAddress = abi.encodePacked(address(234));
    bytes memory longAddress = abi.encode(abi.encode(address(0x1234567890123456789012345678901234567890)));
    bytes memory highAddress = abi.encode(uint256(type(uint160).max) + 1);
    bytes memory precompileAddress = abi.encode(address(1));

    evm2AnyMessage.tokenAmounts[0].destTokenAddress = encodePackedAddress;
    vm.expectRevert(abi.encodeWithSelector(Internal.InvalidEVMAddress.selector, encodePackedAddress));
    s_onRampOverSuperchainInterop.postProcessMessage(evm2AnyMessage);

    evm2AnyMessage.tokenAmounts[0].destTokenAddress = longAddress;
    vm.expectRevert(abi.encodeWithSelector(Internal.InvalidEVMAddress.selector, longAddress));
    s_onRampOverSuperchainInterop.postProcessMessage(evm2AnyMessage);

    evm2AnyMessage.tokenAmounts[0].destTokenAddress = highAddress;
    vm.expectRevert(abi.encodeWithSelector(Internal.InvalidEVMAddress.selector, highAddress));
    s_onRampOverSuperchainInterop.postProcessMessage(evm2AnyMessage);

    evm2AnyMessage.tokenAmounts[0].destTokenAddress = precompileAddress;
    vm.expectRevert(abi.encodeWithSelector(Internal.InvalidEVMAddress.selector, precompileAddress));
    s_onRampOverSuperchainInterop.postProcessMessage(evm2AnyMessage);
  }
}
