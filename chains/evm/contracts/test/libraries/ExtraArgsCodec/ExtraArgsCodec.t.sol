// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ExtraArgsCodec} from "../../../libraries/ExtraArgsCodec.sol";
import {BaseTest} from "../../BaseTest.t.sol";
import {ExtraArgsCodecHelper} from "../../helpers/ExtraArgsCodecHelpers.sol";
import {ExtraArgsCodecUnoptimized} from "../../helpers/ExtraArgsCodecUnoptimized.sol";

contract UnoptimizedDecodeHelper {
  function _decodeGenericExtraArgsV3(
    bytes calldata encoded
  ) external pure returns (ExtraArgsCodec.GenericExtraArgsV3 memory) {
    return ExtraArgsCodecUnoptimized._decodeGenericExtraArgsV3(encoded);
  }
}

contract ExtraArgsCodec_Test is BaseTest {
  ExtraArgsCodecHelper internal s_decoder;
  UnoptimizedDecodeHelper internal s_unoptimizedDecoder;

  function setUp() public override {
    super.setUp();
    s_decoder = new ExtraArgsCodecHelper();
    s_unoptimizedDecoder = new UnoptimizedDecodeHelper();
  }

  function test_EncodeGenericExtraArgsV3_AllDynamicArgsDefaultValues() public view {
    ExtraArgsCodec.GenericExtraArgsV3 memory args = ExtraArgsCodec.GenericExtraArgsV3({
      ccvs: new address[](0),
      ccvArgs: new bytes[](0),
      finalityConfig: 12,
      gasLimit: GAS_LIMIT,
      executor: address(0),
      executorArgs: "",
      tokenReceiver: "",
      tokenArgs: ""
    });

    bytes memory encoded = ExtraArgsCodecUnoptimized._encodeGenericExtraArgsV3(args);
    assertEq(encoded.length, ExtraArgsCodec.GENERIC_EXTRA_ARGS_V3_BASE_SIZE);

    _assertSameArgs(args, s_decoder._decodeGenericExtraArgsV3(encoded));
  }

  function test_EncodeGenericExtraArgsV3_ExecutorNonZeroAddress() public {
    address executor = makeAddr("executor");
    ExtraArgsCodec.GenericExtraArgsV3 memory args = ExtraArgsCodec.GenericExtraArgsV3({
      ccvs: new address[](0),
      ccvArgs: new bytes[](0),
      finalityConfig: 12,
      gasLimit: GAS_LIMIT,
      executor: executor,
      executorArgs: "",
      tokenReceiver: "",
      tokenArgs: ""
    });

    bytes memory encoded = ExtraArgsCodecUnoptimized._encodeGenericExtraArgsV3(args);
    // 20 bytes for the executor address since it's an EVM address.
    assertEq(encoded.length, ExtraArgsCodec.GENERIC_EXTRA_ARGS_V3_BASE_SIZE + 20);

    _assertSameArgs(args, s_decoder._decodeGenericExtraArgsV3(encoded));
  }

  function test_DecodeGenericExtraArgsV3_RevertWhen_InvalidExecutorLength() public {
    bytes memory invalidEncoded = abi.encodePacked(
      ExtraArgsCodec.GENERIC_EXTRA_ARGS_V3_TAG,
      GAS_LIMIT,
      uint16(12),
      uint8(0),
      uint8(10),
      bytes10(0x12345678901234567890),
      uint16(0),
      uint16(0),
      uint16(0)
    );

    vm.expectRevert(abi.encodeWithSelector(ExtraArgsCodec.InvalidAddressLength.selector, 10));
    s_decoder._decodeGenericExtraArgsV3(invalidEncoded);
  }

  function test_EncodeGenericExtraArgsV3_WithCCVs() public {
    address[] memory ccvs = new address[](2);
    ccvs[0] = makeAddr("ccv1");
    ccvs[1] = makeAddr("ccv2");
    bytes[] memory ccvArgs = new bytes[](2);
    ccvArgs[0] = "args1";
    ccvArgs[1] = "args2";

    ExtraArgsCodec.GenericExtraArgsV3 memory args = ExtraArgsCodec.GenericExtraArgsV3({
      ccvs: ccvs,
      ccvArgs: ccvArgs,
      finalityConfig: 12,
      gasLimit: GAS_LIMIT,
      executor: makeAddr("executor"),
      executorArgs: "execArgs",
      tokenReceiver: abi.encodePacked(makeAddr("tokenReceiver")),
      tokenArgs: "tokenArgs"
    });

    bytes memory encoded = ExtraArgsCodecUnoptimized._encodeGenericExtraArgsV3(args);

    _assertSameArgs(args, s_decoder._decodeGenericExtraArgsV3(encoded));
  }

  /// forge-config: default.fuzz.runs = 4096
  /// forge-config: ccip.fuzz.runs = 4096
  function testFuzz_EncodeGenericExtraArgsV3_Differential_Identical(
    uint32 gasLimit,
    uint16 finalityConfig,
    address[9] memory ccvs,
    bytes[9] memory ccvArgs,
    address executor,
    bytes memory executorArgs,
    bytes memory tokenReceiver,
    bytes memory tokenArgs
  ) public view {
    vm.assume(executorArgs.length <= type(uint16).max);
    vm.assume(tokenArgs.length <= type(uint16).max);
    vm.assume(tokenReceiver.length <= type(uint8).max);

    address[] memory ccvsDynamic = new address[](ccvs.length);
    bytes[] memory ccvArgsDynamic = new bytes[](ccvs.length);
    for (uint256 i = 0; i < ccvs.length; i++) {
      ccvsDynamic[i] = ccvs[i];
      ccvArgsDynamic[i] = ccvArgs[i];
      vm.assume(ccvArgs[i].length <= type(uint16).max);
    }

    ExtraArgsCodec.GenericExtraArgsV3 memory args = ExtraArgsCodec.GenericExtraArgsV3({
      gasLimit: gasLimit,
      finalityConfig: finalityConfig,
      ccvs: ccvsDynamic,
      ccvArgs: ccvArgsDynamic,
      executor: executor,
      executorArgs: executorArgs,
      tokenReceiver: tokenReceiver,
      tokenArgs: tokenArgs
    });

    bytes memory optimizedEncoded = ExtraArgsCodec._encodeGenericExtraArgsV3(args);
    bytes memory unoptimizedEncoded = ExtraArgsCodecUnoptimized._encodeGenericExtraArgsV3(args);

    assertEq(optimizedEncoded, unoptimizedEncoded);

    ExtraArgsCodec.GenericExtraArgsV3 memory optimizedDecoded = s_decoder._decodeGenericExtraArgsV3(optimizedEncoded);
    ExtraArgsCodec.GenericExtraArgsV3 memory unoptimizedDecoded =
      s_unoptimizedDecoder._decodeGenericExtraArgsV3(unoptimizedEncoded);

    // Assert that both decoders produce the same output and that it matches the original args.
    _assertSameArgs(optimizedDecoded, unoptimizedDecoded);
    _assertSameArgs(optimizedDecoded, args);
  }

  function _assertSameArgs(
    ExtraArgsCodec.GenericExtraArgsV3 memory args1,
    ExtraArgsCodec.GenericExtraArgsV3 memory args2
  ) internal pure {
    assertEq(keccak256(abi.encode(args1)), keccak256(abi.encode(args2)));
  }
}
