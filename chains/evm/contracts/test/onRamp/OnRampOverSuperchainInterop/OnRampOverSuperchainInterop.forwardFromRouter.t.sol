// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Client} from "../../../libraries/Client.sol";
import {Internal} from "../../../libraries/Internal.sol";
import {SuperchainInterop} from "../../../libraries/SuperchainInterop.sol";
import {OnRamp} from "../../../onRamp/OnRamp.sol";
import {OnRampOverSuperchainInteropSetup} from "./OnRampOverSuperchainInteropSetup.t.sol";

import {IERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/IERC20.sol";

contract OnRampOverSuperchainInterop_forwardFromRouter is OnRampOverSuperchainInteropSetup {
  function setUp() public virtual override {
    super.setUp();
    vm.startPrank(address(s_sourceRouter));
  }

  function test_forwardFromRouter_SingleBasicMessage() public {
    Client.EVM2AnyMessage memory message = _generateSingleTokenMessage(s_sourceFeeToken, 100);
    uint256 gasLimit = 300000;
    message.extraArgs =
      Client._argsToBytes(Client.GenericExtraArgsV2({gasLimit: gasLimit, allowOutOfOrderExecution: true}));

    uint256 feeAmount = 1234567890;
    IERC20(s_sourceFeeToken).transferFrom(OWNER, address(s_onRampOverSuperchainInterop), feeAmount);

    (Internal.EVM2AnyRampMessage memory evm2AnyMessage, Internal.Any2EVMRampMessage memory any2EvmMessage) =
      _generateInitialSourceDestMessages(DEST_CHAIN_SELECTOR, message, feeAmount);

    vm.expectEmit();
    emit SuperchainInterop.CCIPSuperchainMessageSent(DEST_CHAIN_SELECTOR, 1, any2EvmMessage);

    vm.expectEmit();
    emit OnRamp.CCIPMessageSent(DEST_CHAIN_SELECTOR, 1, evm2AnyMessage);

    s_onRampOverSuperchainInterop.forwardFromRouter(DEST_CHAIN_SELECTOR, message, feeAmount, OWNER);
  }

  function test_forwardFromRouter_ThreeConsecutiveMessages() public {
    Client.EVM2AnyMessage memory message = _generateEmptyMessage();

    uint256 feeAmount = 1234567890;
    IERC20(s_sourceFeeToken).transferFrom(OWNER, address(s_onRampOverSuperchainInterop), feeAmount);

    (Internal.EVM2AnyRampMessage memory evm2AnyMessage, Internal.Any2EVMRampMessage memory any2EvmMessage) =
      _generateInitialSourceDestMessages(DEST_CHAIN_SELECTOR, message, feeAmount);

    for (uint64 seqNum = 1; seqNum <= 3; ++seqNum) {
      // Hashing is always done with empty messageId
      evm2AnyMessage.header.messageId = "";
      evm2AnyMessage.header.nonce = seqNum;
      evm2AnyMessage.header.sequenceNumber = seqNum;
      evm2AnyMessage.header.messageId = Internal._hash(evm2AnyMessage, _getOnRampMetadataHash(DEST_CHAIN_SELECTOR));

      any2EvmMessage = _EVM2AnyRampMessageToAny2EVMRampMessage(evm2AnyMessage);

      vm.expectEmit();
      emit SuperchainInterop.CCIPSuperchainMessageSent(DEST_CHAIN_SELECTOR, seqNum, any2EvmMessage);

      vm.expectEmit();
      emit OnRamp.CCIPMessageSent(DEST_CHAIN_SELECTOR, seqNum, evm2AnyMessage);

      s_onRampOverSuperchainInterop.forwardFromRouter(DEST_CHAIN_SELECTOR, message, feeAmount, OWNER);
    }
  }
}
