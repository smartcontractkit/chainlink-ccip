// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Internal} from "../../../libraries/Internal.sol";
import {OnRampOverSuperchainInterop} from "../../../onRamp/OnRampOverSuperchainInterop.sol";
import {OnRampOverSuperchainInteropSetup} from "./OnRampOverSuperchainInteropSetup.t.sol";

contract OnRampOverSuperchainInterop_reemitInteropMessage is OnRampOverSuperchainInteropSetup {
  function test_ReemitBasicValidation() public {
    Internal.Any2EVMRampMessage memory message = _generateBasicAny2EVMMessage();

    // The function validates source chain and calls hashInteropMessage
    // Since we haven't stored any messages, it should revert with MessageDoesNotExist
    bytes32 messageHash = s_onRampOverSuperchainInterop.hashInteropMessage(message);

    vm.expectRevert(
      abi.encodeWithSelector(
        OnRampOverSuperchainInterop.MessageDoesNotExist.selector,
        message.header.destChainSelector,
        message.header.sequenceNumber,
        messageHash
      )
    );
    s_onRampOverSuperchainInterop.reemitInteropMessage(message);
  }

  // Reverts

  function test_RevertWhen_InvalidSourceChainSelector() public {
    Internal.Any2EVMRampMessage memory message = _generateBasicAny2EVMMessage();
    message.header.sourceChainSelector = SOURCE_CHAIN_SELECTOR + 1;

    vm.expectRevert(
      abi.encodeWithSelector(OnRampOverSuperchainInterop.InvalidSourceChainSelector.selector, SOURCE_CHAIN_SELECTOR + 1)
    );
    s_onRampOverSuperchainInterop.reemitInteropMessage(message);
  }

  function test_RevertWhen_MessageDoesNotExist() public {
    Internal.Any2EVMRampMessage memory message = _generateBasicAny2EVMMessage();
    message.header.sourceChainSelector = SOURCE_CHAIN_SELECTOR;
    bytes32 messageHash = s_onRampOverSuperchainInterop.hashInteropMessage(message);

    vm.expectRevert(
      abi.encodeWithSelector(
        OnRampOverSuperchainInterop.MessageDoesNotExist.selector,
        message.header.destChainSelector,
        message.header.sequenceNumber,
        messageHash
      )
    );
    s_onRampOverSuperchainInterop.reemitInteropMessage(message);
  }
}
