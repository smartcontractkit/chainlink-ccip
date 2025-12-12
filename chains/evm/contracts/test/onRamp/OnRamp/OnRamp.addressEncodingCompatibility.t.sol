// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolV1} from "../../../interfaces/IPool.sol";
import {IPoolV2} from "../../../interfaces/IPoolV2.sol";

import {Client} from "../../../libraries/Client.sol";
import {ExtraArgsCodec} from "../../../libraries/ExtraArgsCodec.sol";
import {MessageV1Codec} from "../../../libraries/MessageV1Codec.sol";
import {Pool} from "../../../libraries/Pool.sol";
import {OnRamp} from "../../../onRamp/OnRamp.sol";
import {OnRampSetup} from "./OnRampSetup.t.sol";

import {IERC165} from "@openzeppelin/contracts@5.3.0/utils/introspection/IERC165.sol";
import {VmSafe} from "forge-std/Vm.sol";

contract OnRamp_addressEncodingCompatibility is OnRampSetup {
  bytes32 internal constant CCIP_MESSAGE_SENT_TOPIC =
    keccak256("CCIPMessageSent(uint64,uint64,bytes32,address,bytes,(address,uint32,uint32,uint256,bytes)[],bytes[])");

  function decode(
    bytes calldata encodedMessage
  ) external pure returns (MessageV1Codec.MessageV1 memory) {
    return MessageV1Codec._decodeMessageV1(encodedMessage);
  }

  function _forwardAndDecode(
    Client.EVM2AnyMessage memory message,
    address originalSender
  ) internal returns (MessageV1Codec.MessageV1 memory) {
    uint256 fee = s_onRamp.getFee(DEST_CHAIN_SELECTOR, message);

    deal(s_sourceFeeToken, address(s_onRamp), type(uint96).max);
    vm.startPrank(address(s_sourceRouter));
    vm.recordLogs();
    s_onRamp.forwardFromRouter(DEST_CHAIN_SELECTOR, message, fee, originalSender);
    vm.stopPrank();

    VmSafe.Log[] memory logs = vm.getRecordedLogs();
    bytes memory encodedMessage;
    for (uint256 i = 0; i < logs.length; ++i) {
      if (
        logs[i].emitter == address(s_onRamp) && logs[i].topics.length != 0
          && logs[i].topics[0] == CCIP_MESSAGE_SENT_TOPIC
      ) {
        (, encodedMessage,,) = abi.decode(logs[i].data, (address, bytes, OnRamp.Receipt[], bytes[]));
        break;
      }
    }
    require(encodedMessage.length != 0, "encoded message not found");

    return this.decode(encodedMessage);
  }

  function test_forwardFromRouter_NormalizesReceiver_AbiEncoded() public {
    address receiver = makeAddr("receiverEncoded");
    Client.EVM2AnyMessage memory message = _generateEmptyMessage();
    message.receiver = abi.encode(receiver);

    MessageV1Codec.MessageV1 memory decoded = _forwardAndDecode(message, STRANGER);

    assertEq(decoded.receiver, abi.encodePacked(receiver));
    assertEq(decoded.receiver.length, EVM_ADDRESS_LENGTH);
  }

  function test_forwardFromRouter_NormalizesReceiver_AbiEncodePacked() public {
    address receiver = makeAddr("receiverPacked");
    Client.EVM2AnyMessage memory message = _generateEmptyMessage();
    message.receiver = abi.encodePacked(receiver);

    MessageV1Codec.MessageV1 memory decoded = _forwardAndDecode(message, STRANGER);

    assertEq(decoded.receiver, abi.encodePacked(receiver));
    assertEq(decoded.receiver.length, EVM_ADDRESS_LENGTH);
  }

  function _assertTokenTransferNormalization(bool tokenReceiverPackedInput, bool destTokenPackedInput) internal {
    address token = s_sourceFeeToken;
    address pool = makeAddr("mockPool");
    address tokenReceiver = makeAddr(tokenReceiverPackedInput ? "tokenReceiverPacked" : "tokenReceiverEncoded");
    address destToken = makeAddr(destTokenPackedInput ? "destTokenPacked" : "destTokenEncoded");
    address receiver = makeAddr("messageReceiver");

    vm.mockCall(address(s_tokenAdminRegistry), abi.encodeCall(s_tokenAdminRegistry.getPool, (token)), abi.encode(pool));
    vm.mockCall(pool, abi.encodeCall(IERC165(pool).supportsInterface, (Pool.CCIP_POOL_V1)), abi.encode(true));
    vm.mockCall(pool, abi.encodeCall(IERC165(pool).supportsInterface, (type(IPoolV2).interfaceId)), abi.encode(false));

    bytes memory expectedTokenReceiver = abi.encodePacked(tokenReceiver);
    Pool.LockOrBurnInV1 memory expectedInput = Pool.LockOrBurnInV1({
      receiver: expectedTokenReceiver,
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      originalSender: STRANGER,
      amount: 1e18,
      localToken: token
    });

    Pool.LockOrBurnOutV1 memory poolReturnData = Pool.LockOrBurnOutV1({
      destTokenAddress: destTokenPackedInput ? abi.encodePacked(destToken) : abi.encode(destToken),
      destPoolData: abi.encode("poolData")
    });

    bytes memory lockOrBurnCalldata = abi.encodeWithSelector(IPoolV1.lockOrBurn.selector, expectedInput);
    vm.expectCall(pool, lockOrBurnCalldata);
    vm.mockCall(pool, lockOrBurnCalldata, abi.encode(poolReturnData));

    Client.EVMTokenAmount[] memory tokenAmounts = new Client.EVMTokenAmount[](1);
    tokenAmounts[0] = Client.EVMTokenAmount({token: token, amount: expectedInput.amount});

    ExtraArgsCodec.GenericExtraArgsV3 memory extraArgs = ExtraArgsCodec.GenericExtraArgsV3({
      ccvs: new address[](0),
      ccvArgs: new bytes[](0),
      blockConfirmations: 0,
      gasLimit: GAS_LIMIT,
      executor: address(0),
      executorArgs: "",
      tokenReceiver: tokenReceiverPackedInput ? abi.encodePacked(tokenReceiver) : abi.encode(tokenReceiver),
      tokenArgs: ""
    });

    Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
      receiver: abi.encode(receiver),
      data: "payload",
      tokenAmounts: tokenAmounts,
      feeToken: s_sourceFeeToken,
      extraArgs: ExtraArgsCodec._encodeGenericExtraArgsV3(extraArgs)
    });

    MessageV1Codec.MessageV1 memory decoded = _forwardAndDecode(message, STRANGER);

    assertEq(decoded.receiver, abi.encodePacked(receiver));
    assertEq(decoded.tokenTransfer.length, 1);
    assertEq(decoded.tokenTransfer[0].tokenReceiver, expectedTokenReceiver);
    assertEq(decoded.tokenTransfer[0].destTokenAddress, abi.encodePacked(destToken));
  }

  function test_forwardFromRouter_NormalizesTokenReceiverAndDestToken_AbiEncoded() public {
    _assertTokenTransferNormalization(false, false);
  }

  function test_forwardFromRouter_NormalizesTokenReceiverAndDestToken_Packed() public {
    _assertTokenTransferNormalization(true, true);
  }

  function test_forwardFromRouter_NormalizesTokenReceiver_Packed_DestToken_AbiEncoded() public {
    _assertTokenTransferNormalization(true, false);
  }

  function test_forwardFromRouter_NormalizesTokenReceiver_AbiEncoded_DestToken_Packed() public {
    _assertTokenTransferNormalization(false, true);
  }
}
