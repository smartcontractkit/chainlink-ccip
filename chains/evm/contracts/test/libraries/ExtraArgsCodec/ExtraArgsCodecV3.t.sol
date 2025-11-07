// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ExtraArgsCodec} from "../../../libraries/ExtraArgsCodec.sol";
import {BaseTest} from "../../BaseTest.t.sol";
import {ExtraArgsCodecHelper} from "../../helpers/ExtraArgsCodecHelpers.sol";

/// forge-config: default.allow_internal_expect_revert = true
contract ExtraArgsCodecV3_Test is BaseTest {
  ExtraArgsCodecHelper internal s_helper;

  function setUp() public override {
    super.setUp();
    s_helper = new ExtraArgsCodecHelper();
  }

  function test_GetBasicEncodedExtraArgsV3_ReturnsCorrectLength() public pure {
    bytes memory encoded = ExtraArgsCodec._getBasicEncodedExtraArgsV3(GAS_LIMIT, 12);
    assertEq(encoded.length, 18);
    bytes4 tag;
    assembly {
      tag := mload(add(encoded, 32))
    }
    assertEq(tag, ExtraArgsCodec.GENERIC_EXTRA_ARGS_V3_TAG);
  }

  function test_DecodeGenericExtraArgsV3_Basic() public view {
    ExtraArgsCodec.GenericExtraArgsV3 memory args = ExtraArgsCodec.GenericExtraArgsV3({
      gasLimit: GAS_LIMIT,
      finalityConfig: 12,
      ccvs: new address[](0),
      ccvArgs: new bytes[](0),
      executor: address(0),
      executorArgs: "",
      tokenReceiver: "",
      tokenArgs: ""
    });

    bytes memory encoded = ExtraArgsCodec._encodeGenericExtraArgsV3(args);
    ExtraArgsCodec.GenericExtraArgsV3 memory decoded = s_helper._decodeGenericExtraArgsV3(encoded);

    assertEq(decoded.gasLimit, GAS_LIMIT);
    assertEq(decoded.finalityConfig, 12);
    assertEq(decoded.ccvs.length, 0);
    assertEq(decoded.executor, address(0));
  }

  function test_DecodeGenericExtraArgsV3_WithExecutor() public {
    address executor = makeAddr("executor");
    ExtraArgsCodec.GenericExtraArgsV3 memory args = ExtraArgsCodec.GenericExtraArgsV3({
      gasLimit: GAS_LIMIT * 2,
      finalityConfig: 5,
      ccvs: new address[](0),
      ccvArgs: new bytes[](0),
      executor: executor,
      executorArgs: "data",
      tokenReceiver: "",
      tokenArgs: ""
    });

    bytes memory encoded = ExtraArgsCodec._encodeGenericExtraArgsV3(args);
    ExtraArgsCodec.GenericExtraArgsV3 memory decoded = s_helper._decodeGenericExtraArgsV3(encoded);

    assertEq(decoded.executor, executor);
    assertEq(decoded.executorArgs, "data");
  }

  function test_DecodeGenericExtraArgsV3_WithCCVs() public {
    address[] memory ccvs = new address[](2);
    ccvs[0] = makeAddr("ccv1");
    ccvs[1] = makeAddr("ccv2");
    bytes[] memory ccvArgs = new bytes[](2);
    ccvArgs[0] = "args1";
    ccvArgs[1] = "args2";

    ExtraArgsCodec.GenericExtraArgsV3 memory args = ExtraArgsCodec.GenericExtraArgsV3({
      gasLimit: GAS_LIMIT + 100_000,
      finalityConfig: 10,
      ccvs: ccvs,
      ccvArgs: ccvArgs,
      executor: address(0),
      executorArgs: "",
      tokenReceiver: "",
      tokenArgs: ""
    });

    bytes memory encoded = ExtraArgsCodec._encodeGenericExtraArgsV3(args);
    ExtraArgsCodec.GenericExtraArgsV3 memory decoded = s_helper._decodeGenericExtraArgsV3(encoded);

    assertEq(decoded.ccvs.length, 2);
    assertEq(decoded.ccvs[0], makeAddr("ccv1"));
    assertEq(decoded.ccvArgs[0], "args1");
  }

  function test_DecodeGenericExtraArgsV3_WithTokenArgs() public {
    address tokenReceiver = makeAddr("tokenReceiver");
    ExtraArgsCodec.GenericExtraArgsV3 memory args = ExtraArgsCodec.GenericExtraArgsV3({
      gasLimit: GAS_LIMIT / 2,
      finalityConfig: 1,
      ccvs: new address[](0),
      ccvArgs: new bytes[](0),
      executor: address(0),
      executorArgs: "",
      tokenReceiver: abi.encodePacked(tokenReceiver),
      tokenArgs: "token args"
    });

    bytes memory encoded = ExtraArgsCodec._encodeGenericExtraArgsV3(args);
    ExtraArgsCodec.GenericExtraArgsV3 memory decoded = s_helper._decodeGenericExtraArgsV3(encoded);

    assertEq(decoded.tokenReceiver, abi.encodePacked(tokenReceiver));
    assertEq(decoded.tokenArgs, "token args");
  }

  function test_DecodeGenericExtraArgsV3_ZeroValues() public view {
    ExtraArgsCodec.GenericExtraArgsV3 memory args = ExtraArgsCodec.GenericExtraArgsV3({
      gasLimit: 0,
      finalityConfig: 0,
      ccvs: new address[](0),
      ccvArgs: new bytes[](0),
      executor: address(0),
      executorArgs: "",
      tokenReceiver: "",
      tokenArgs: ""
    });

    bytes memory encoded = ExtraArgsCodec._encodeGenericExtraArgsV3(args);
    ExtraArgsCodec.GenericExtraArgsV3 memory decoded = s_helper._decodeGenericExtraArgsV3(encoded);

    assertEq(decoded.gasLimit, 0);
    assertEq(decoded.finalityConfig, 0);
  }

  function test_DecodeGenericExtraArgsV3_MaxValues() public view {
    ExtraArgsCodec.GenericExtraArgsV3 memory args = ExtraArgsCodec.GenericExtraArgsV3({
      gasLimit: type(uint32).max,
      finalityConfig: type(uint16).max,
      ccvs: new address[](0),
      ccvArgs: new bytes[](0),
      executor: address(0),
      executorArgs: "",
      tokenReceiver: "",
      tokenArgs: ""
    });

    bytes memory encoded = ExtraArgsCodec._encodeGenericExtraArgsV3(args);
    ExtraArgsCodec.GenericExtraArgsV3 memory decoded = s_helper._decodeGenericExtraArgsV3(encoded);

    assertEq(decoded.gasLimit, type(uint32).max);
    assertEq(decoded.finalityConfig, type(uint16).max);
  }

  function test_DecodeGenericExtraArgsV3_RevertWhen_DataTooShort() public {
    vm.expectRevert(
      abi.encodeWithSelector(
        ExtraArgsCodec.InvalidDataLength.selector, ExtraArgsCodec.EncodingErrorLocation.EXTRA_ARGS_STATIC_LENGTH_FIELDS
      )
    );
    s_helper._decodeGenericExtraArgsV3(new bytes(10));
  }

  function test_DecodeGenericExtraArgsV3_RevertWhen_InvalidExecutorLength() public {
    bytes memory invalidData = abi.encodePacked(
      ExtraArgsCodec.GENERIC_EXTRA_ARGS_V3_TAG,
      GAS_LIMIT / 2,
      uint16(1),
      uint8(0),
      uint8(10),
      bytes10(0x12345678901234567890),
      uint16(0),
      uint16(0),
      uint16(0)
    );

    vm.expectRevert(abi.encodeWithSelector(ExtraArgsCodec.InvalidExecutorLength.selector, 10));
    s_helper._decodeGenericExtraArgsV3(invalidData);
  }

  function test_DecodeGenericExtraArgsV3_RevertWhen_InvalidCCVAddressLength() public {
    bytes memory invalidData = abi.encodePacked(
      ExtraArgsCodec.GENERIC_EXTRA_ARGS_V3_TAG,
      GAS_LIMIT / 2,
      uint16(1),
      uint8(1),
      uint8(10),
      bytes10(0x12345678901234567890),
      uint16(0),
      uint8(0),
      uint16(0),
      uint16(0),
      uint16(0)
    );

    vm.expectRevert(abi.encodeWithSelector(ExtraArgsCodec.InvalidCCVAddressLength.selector, 10));
    s_helper._decodeGenericExtraArgsV3(invalidData);
  }

  function test_DecodeGenericExtraArgsV3_RevertWhen_ExtraBytes() public {
    ExtraArgsCodec.GenericExtraArgsV3 memory args = ExtraArgsCodec.GenericExtraArgsV3({
      gasLimit: GAS_LIMIT / 2,
      finalityConfig: 1,
      ccvs: new address[](0),
      ccvArgs: new bytes[](0),
      executor: address(0),
      executorArgs: "",
      tokenReceiver: "",
      tokenArgs: ""
    });

    bytes memory encoded = ExtraArgsCodec._encodeGenericExtraArgsV3(args);
    bytes memory withExtra = bytes.concat(encoded, bytes("extra"));

    vm.expectRevert(
      abi.encodeWithSelector(
        ExtraArgsCodec.InvalidDataLength.selector, ExtraArgsCodec.EncodingErrorLocation.EXTRA_ARGS_FINAL_OFFSET
      )
    );
    s_helper._decodeGenericExtraArgsV3(withExtra);
  }
}
