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
      if (logs[i].topics.length != 0 && logs[i].topics[0] == CCIP_MESSAGE_SENT_TOPIC) {
        (, encodedMessage,,) = abi.decode(logs[i].data, (address, bytes, OnRamp.Receipt[], bytes[]));
        break;
      }
    }
    require(encodedMessage.length != 0, "encoded message not found");

    return this.decode(encodedMessage);
  }

  function _setDestChainAddressLength(
    uint8 addressBytesLength
  ) internal {
    address[] memory defaultCCVs = new address[](1);
    defaultCCVs[0] = s_defaultCCV;

    OnRamp.DestChainConfigArgs[] memory args = new OnRamp.DestChainConfigArgs[](1);
    args[0] = OnRamp.DestChainConfigArgs({
      destChainSelector: DEST_CHAIN_SELECTOR,
      router: s_sourceRouter,
      addressBytesLength: addressBytesLength,
      networkFeeUSDCents: NETWORK_FEE_USD_CENTS,
      baseExecutionGasCost: BASE_EXEC_GAS_COST,
      laneMandatedCCVs: new address[](0),
      defaultCCVs: defaultCCVs,
      defaultExecutor: s_defaultExecutor,
      offRamp: _encodeAddressToLength(address(s_offRampOnRemoteChain), addressBytesLength)
    });

    s_onRamp.applyDestChainConfigUpdates(args);
  }

  function _expectTrimmed(bytes memory actual, address addr, uint8 addressBytesLength) internal pure {
    assertEq(actual.length, addressBytesLength);
    if (addressBytesLength <= 20) {
      assertEq(actual, abi.encodePacked(addr));
    } else if (addressBytesLength == 32) {
      // For 32 bytes we expect abi.encode(addr)
      assertEq(actual, abi.encode(addr));
    } else {
      assertEq(actual, _encodeAddressToLength(addr, addressBytesLength));
    }
  }

  function _encodeAddressToLength(address addr, uint8 addressBytesLength) internal pure returns (bytes memory) {
    if (addressBytesLength == 20) return abi.encodePacked(addr);
    if (addressBytesLength == 32) return abi.encode(addr);

    // addressBytesLength must be between 21 and 31 here. Left-pad zeros and place the address in the last 20 bytes.
    bytes memory out = new bytes(addressBytesLength);
    bytes20 addrBytes = bytes20(addr);
    uint256 start = addressBytesLength - 20;
    for (uint256 i = 0; i < 20; ++i) {
      out[start + i] = addrBytes[i];
    }
    return out;
  }

  function test_forwardFromRouter_SenderAbiEncodedForEvmDest() public {
    _setDestChainAddressLength(EVM_ADDRESS_LENGTH);
    address originalSender = makeAddr("originalSender");
    Client.EVM2AnyMessage memory message = _generateEmptyMessage();

    MessageV1Codec.MessageV1 memory decoded = _forwardAndDecode(message, originalSender);

    assertEq(decoded.sender.length, 32, "sender should be 32 bytes");
    assertEq(decoded.sender, abi.encode(originalSender), "sender should be abi.encode(address)");
  }

  function test_forwardFromRouter_ReceiverAbiEncodedTrimsToDestLength() public {
    _setDestChainAddressLength(EVM_ADDRESS_LENGTH);
    address receiver = makeAddr("receiver");
    Client.EVM2AnyMessage memory message = _generateEmptyMessage();
    message.receiver = abi.encode(receiver); // 32 bytes

    MessageV1Codec.MessageV1 memory decoded = _forwardAndDecode(message, STRANGER);
    _expectTrimmed(decoded.receiver, receiver, EVM_ADDRESS_LENGTH);
  }

  function test_forwardFromRouter_ReceiverAcceptsExactDestLength() public {
    _setDestChainAddressLength(EVM_ADDRESS_LENGTH);
    address receiver = makeAddr("receiverPacked");
    Client.EVM2AnyMessage memory message = _generateEmptyMessage();
    message.receiver = abi.encodePacked(receiver); // exact length

    MessageV1Codec.MessageV1 memory decoded = _forwardAndDecode(message, STRANGER);
    _expectTrimmed(decoded.receiver, receiver, EVM_ADDRESS_LENGTH);
  }

  function test_forwardFromRouter_TokenReceiverAbiEncodedTrimsToDestLength() public {
    _setDestChainAddressLength(EVM_ADDRESS_LENGTH);
    address tokenReceiver = makeAddr("tokenReceiver");

    ExtraArgsCodec.GenericExtraArgsV3 memory extraArgs = ExtraArgsCodec.GenericExtraArgsV3({
      ccvs: new address[](0),
      ccvArgs: new bytes[](0),
      blockConfirmations: 0,
      gasLimit: GAS_LIMIT,
      executor: address(0),
      executorArgs: "",
      tokenReceiver: abi.encode(tokenReceiver), // 32 bytes
      tokenArgs: ""
    });

    Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
      receiver: abi.encode(makeAddr("receiver")),
      data: "",
      tokenAmounts: new Client.EVMTokenAmount[](0),
      feeToken: s_sourceFeeToken,
      extraArgs: ExtraArgsCodec._encodeGenericExtraArgsV3(extraArgs)
    });

    MessageV1Codec.MessageV1 memory decoded = _forwardAndDecode(message, STRANGER);
    // token receiver is part of the message only when tokens are present; here we just ensure no revert and receiver normalized
    _expectTrimmed(decoded.receiver, makeAddr("receiver"), EVM_ADDRESS_LENGTH);
  }

  function test_forwardFromRouter_DestTokenTrimsToDestLengthWhenPoolReturnsAbiEncoded() public {
    _setDestChainAddressLength(EVM_ADDRESS_LENGTH);
    address token = s_sourceFeeToken;
    address pool = makeAddr("mockPool");
    address destToken = makeAddr("destToken");
    address receiver = makeAddr("receiver");
    address tokenReceiver = makeAddr("tokenReceiver");

    vm.mockCall(address(s_tokenAdminRegistry), abi.encodeCall(s_tokenAdminRegistry.getPool, (token)), abi.encode(pool));
    vm.mockCall(pool, abi.encodeCall(IERC165(pool).supportsInterface, (Pool.CCIP_POOL_V1)), abi.encode(true));
    vm.mockCall(pool, abi.encodeCall(IERC165(pool).supportsInterface, (type(IPoolV2).interfaceId)), abi.encode(false));

    Pool.LockOrBurnOutV1 memory poolReturnData = Pool.LockOrBurnOutV1({
      destTokenAddress: abi.encode(destToken), // 32 bytes
      destPoolData: abi.encode("poolData")
    });

    Pool.LockOrBurnInV1 memory expectedInput = Pool.LockOrBurnInV1({
      receiver: abi.encodePacked(tokenReceiver),
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      originalSender: STRANGER,
      amount: 1e18,
      localToken: token
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
      tokenReceiver: abi.encode(tokenReceiver),
      tokenArgs: ""
    });

    Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
      receiver: abi.encodePacked(receiver),
      data: "payload",
      tokenAmounts: tokenAmounts,
      feeToken: s_sourceFeeToken,
      extraArgs: ExtraArgsCodec._encodeGenericExtraArgsV3(extraArgs)
    });

    MessageV1Codec.MessageV1 memory decoded = _forwardAndDecode(message, STRANGER);

    assertEq(decoded.tokenTransfer.length, 1);
    _expectTrimmed(decoded.tokenTransfer[0].destTokenAddress, destToken, EVM_ADDRESS_LENGTH);
  }

  function test_forwardFromRouter_ReceiverAcceptsCustomNonEvmLength() public {
    uint8 customLen = 24;
    _setDestChainAddressLength(customLen);

    bytes memory receiver = new bytes(customLen);
    for (uint256 i = 0; i < customLen; ++i) {
      receiver[i] = bytes1(uint8(i + 1));
    }

    Client.EVM2AnyMessage memory message = _generateEmptyMessage();
    message.receiver = receiver;

    MessageV1Codec.MessageV1 memory decoded = _forwardAndDecode(message, STRANGER);
    assertEq(decoded.receiver, receiver);
    assertEq(decoded.receiver.length, customLen);
  }

  function test_forwardFromRouter_TokenReceiverAcceptsCustomNonEvmLength() public {
    uint8 customLen = 24;
    _setDestChainAddressLength(customLen);

    bytes memory tokenReceiver = new bytes(customLen);
    for (uint256 i = 0; i < customLen; ++i) {
      tokenReceiver[i] = bytes1(uint8(0xAA + i));
    }

    ExtraArgsCodec.GenericExtraArgsV3 memory extraArgs = ExtraArgsCodec.GenericExtraArgsV3({
      ccvs: new address[](0),
      ccvArgs: new bytes[](0),
      blockConfirmations: 0,
      gasLimit: GAS_LIMIT,
      executor: address(0),
      executorArgs: "",
      tokenReceiver: tokenReceiver,
      tokenArgs: ""
    });

    Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
      receiver: tokenReceiver, // same length, should pass
      data: "",
      tokenAmounts: new Client.EVMTokenAmount[](0),
      feeToken: s_sourceFeeToken,
      extraArgs: ExtraArgsCodec._encodeGenericExtraArgsV3(extraArgs)
    });

    MessageV1Codec.MessageV1 memory decoded = _forwardAndDecode(message, STRANGER);
    assertEq(decoded.receiver, tokenReceiver);
    assertEq(decoded.receiver.length, customLen);
  }

  // ================================================================
  // │                          Reverts                             │
  // ================================================================

  function test_forwardFromRouter_RevertWhen_ReceiverPaddingNonZero() public {
    _setDestChainAddressLength(EVM_ADDRESS_LENGTH);
    bytes memory bad = abi.encode(bytes32(type(uint256).max)); // high bits non-zero
    Client.EVM2AnyMessage memory message = _generateEmptyMessage();
    message.receiver = bad;

    vm.startPrank(address(s_sourceRouter));
    vm.expectRevert(abi.encodeWithSelector(OnRamp.InvalidDestChainAddress.selector, bad));
    s_onRamp.forwardFromRouter(DEST_CHAIN_SELECTOR, message, 0, STRANGER);
    vm.stopPrank();
  }

  function test_forwardFromRouter_RevertWhen_ReceiverLengthNotDestOr32() public {
    _setDestChainAddressLength(EVM_ADDRESS_LENGTH);
    Client.EVM2AnyMessage memory message = _generateEmptyMessage();
    address receiver = makeAddr("receiver");
    message.receiver = abi.encodePacked(receiver, uint8(0));

    vm.startPrank(address(s_sourceRouter));
    vm.expectRevert(abi.encodeWithSelector(OnRamp.InvalidDestChainAddress.selector, message.receiver));
    s_onRamp.forwardFromRouter(DEST_CHAIN_SELECTOR, message, 0, STRANGER);
    vm.stopPrank();
  }

  function test_forwardFromRouter_RevertWhen_ReceiverShortFor32ByteChain() public {
    _setDestChainAddressLength(NON_EVM_ADDRESS_LENGTH);
    address receiver = makeAddr("receiverShort");
    Client.EVM2AnyMessage memory message = _generateEmptyMessage();
    message.receiver = abi.encodePacked(receiver); // 20 bytes, but dest expects 32

    vm.startPrank(address(s_sourceRouter));
    vm.expectRevert(abi.encodeWithSelector(OnRamp.InvalidDestChainAddress.selector, message.receiver));
    s_onRamp.forwardFromRouter(DEST_CHAIN_SELECTOR, message, 0, STRANGER);
    vm.stopPrank();
  }

  function test_forwardFromRouter_RevertWhen_TokenReceiverPaddingNonZero() public {
    _setDestChainAddressLength(EVM_ADDRESS_LENGTH);
    bytes memory bad = abi.encode(bytes32(type(uint256).max));

    ExtraArgsCodec.GenericExtraArgsV3 memory extraArgs = ExtraArgsCodec.GenericExtraArgsV3({
      ccvs: new address[](0),
      ccvArgs: new bytes[](0),
      blockConfirmations: 0,
      gasLimit: GAS_LIMIT,
      executor: address(0),
      executorArgs: "",
      tokenReceiver: bad,
      tokenArgs: ""
    });

    Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
      receiver: abi.encode(makeAddr("receiver")),
      data: "",
      tokenAmounts: new Client.EVMTokenAmount[](0),
      feeToken: s_sourceFeeToken,
      extraArgs: ExtraArgsCodec._encodeGenericExtraArgsV3(extraArgs)
    });

    vm.startPrank(address(s_sourceRouter));
    vm.expectRevert(abi.encodeWithSelector(OnRamp.InvalidDestChainAddress.selector, bad));
    s_onRamp.forwardFromRouter(DEST_CHAIN_SELECTOR, message, 0, STRANGER);
    vm.stopPrank();
  }

  function test_forwardFromRouter_RevertWhen_TokenReceiverLengthNotDestOr32() public {
    _setDestChainAddressLength(EVM_ADDRESS_LENGTH);
    bytes memory wrongLen = bytes("wronglengthhere!"); // 17 bytes

    ExtraArgsCodec.GenericExtraArgsV3 memory extraArgs = ExtraArgsCodec.GenericExtraArgsV3({
      ccvs: new address[](0),
      ccvArgs: new bytes[](0),
      blockConfirmations: 0,
      gasLimit: GAS_LIMIT,
      executor: address(0),
      executorArgs: "",
      tokenReceiver: wrongLen,
      tokenArgs: ""
    });

    Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
      receiver: abi.encode(makeAddr("receiver")),
      data: "",
      tokenAmounts: new Client.EVMTokenAmount[](0),
      feeToken: s_sourceFeeToken,
      extraArgs: ExtraArgsCodec._encodeGenericExtraArgsV3(extraArgs)
    });

    vm.startPrank(address(s_sourceRouter));
    vm.expectRevert(abi.encodeWithSelector(OnRamp.InvalidDestChainAddress.selector, wrongLen));
    s_onRamp.forwardFromRouter(DEST_CHAIN_SELECTOR, message, 0, STRANGER);
    vm.stopPrank();
  }

  function test_applyDestChainConfigUpdates_RevertWhen_OffRampLengthMismatch() public {
    address[] memory defaultCCVs = new address[](1);
    defaultCCVs[0] = s_defaultCCV;

    bytes memory offRamp = abi.encode(address(s_offRampOnRemoteChain)); // 32 bytes, expected 20 for EVM

    OnRamp.DestChainConfigArgs[] memory args = new OnRamp.DestChainConfigArgs[](1);
    args[0] = OnRamp.DestChainConfigArgs({
      destChainSelector: DEST_CHAIN_SELECTOR,
      router: s_sourceRouter,
      addressBytesLength: EVM_ADDRESS_LENGTH,
      networkFeeUSDCents: NETWORK_FEE_USD_CENTS,
      baseExecutionGasCost: BASE_EXEC_GAS_COST,
      laneMandatedCCVs: new address[](0),
      defaultCCVs: defaultCCVs,
      defaultExecutor: s_defaultExecutor,
      offRamp: offRamp
    });

    vm.expectRevert(abi.encodeWithSelector(OnRamp.InvalidDestChainAddress.selector, offRamp));
    s_onRamp.applyDestChainConfigUpdates(args);
  }
}
