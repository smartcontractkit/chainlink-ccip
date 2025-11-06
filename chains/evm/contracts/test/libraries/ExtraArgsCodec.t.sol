// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Client} from "../../libraries/Client.sol";
import {ExtraArgsCodec} from "../../libraries/ExtraArgsCodec.sol";
import {ExtraArgsCodecUnoptimized} from "../helpers/ExtraArgsCodecUnoptimized.sol";
import {Test} from "forge-std/Test.sol";

contract ExtraArgsCodecHelper {
  function decode(
    bytes calldata encoded
  ) external pure returns (ExtraArgsCodec.GenericExtraArgsV3 memory) {
    return ExtraArgsCodecUnoptimized._decodeGenericExtraArgsV3(encoded);
  }
}

contract ExtraArgsCodec_Test is Test {
  ExtraArgsCodecHelper internal helper;

  function setUp() public {
    helper = new ExtraArgsCodecHelper();
  }

  function test_encodeDecodeExecutorZeroAddress() public view {
    ExtraArgsCodec.GenericExtraArgsV3 memory args = ExtraArgsCodec.GenericExtraArgsV3({
      ccvs: new address[](0),
      ccvArgs: new bytes[](0),
      finalityConfig: 12,
      gasLimit: 200_000,
      executor: address(0),
      executorArgs: "",
      tokenReceiver: "",
      tokenArgs: ""
    });

    bytes memory encoded = ExtraArgsCodecUnoptimized._encodeGenericExtraArgsV3(args);
    ExtraArgsCodec.GenericExtraArgsV3 memory decoded = helper.decode(encoded);

    assertEq(decoded.executor, address(0), "Executor should be address(0)");
    assertEq(decoded.finalityConfig, 12, "FinalityConfig should match");
    assertEq(decoded.gasLimit, 200_000, "GasLimit should match");
  }

  function test_encodeDecodeExecutorNonZeroAddress() public view {
    address executor = address(0x1234567890123456789012345678901234567890);
    ExtraArgsCodec.GenericExtraArgsV3 memory args = ExtraArgsCodec.GenericExtraArgsV3({
      ccvs: new address[](0),
      ccvArgs: new bytes[](0),
      finalityConfig: 12,
      gasLimit: 200_000,
      executor: executor,
      executorArgs: "",
      tokenReceiver: "",
      tokenArgs: ""
    });

    bytes memory encoded = ExtraArgsCodecUnoptimized._encodeGenericExtraArgsV3(args);
    ExtraArgsCodec.GenericExtraArgsV3 memory decoded = helper.decode(encoded);

    assertEq(decoded.executor, executor, "Executor should match");
    assertEq(decoded.finalityConfig, 12, "FinalityConfig should match");
    assertEq(decoded.gasLimit, 200_000, "GasLimit should match");
  }

  function test_encodeExecutorZeroAddress_ChecksLength() public pure {
    ExtraArgsCodec.GenericExtraArgsV3 memory args = ExtraArgsCodec.GenericExtraArgsV3({
      ccvs: new address[](0),
      ccvArgs: new bytes[](0),
      finalityConfig: 12,
      gasLimit: 200_000,
      executor: address(0),
      executorArgs: "",
      tokenReceiver: "",
      tokenArgs: ""
    });

    bytes memory encoded = ExtraArgsCodecUnoptimized._encodeGenericExtraArgsV3(args);

    // Check the executor length field is 0
    // Format: 4 (tag) + 4 (gasLimit) + 2 (finality) + 1 (ccvs length) = 11 bytes offset
    // Then 1 byte for executor length
    uint8 executorLength = uint8(encoded[11]);
    assertEq(executorLength, 0, "Executor length should be 0 for address(0)");
  }

  function test_encodeExecutorNonZeroAddress_ChecksLength() public pure {
    address executor = address(0x1234567890123456789012345678901234567890);
    ExtraArgsCodec.GenericExtraArgsV3 memory args = ExtraArgsCodec.GenericExtraArgsV3({
      ccvs: new address[](0),
      ccvArgs: new bytes[](0),
      finalityConfig: 12,
      gasLimit: 200_000,
      executor: executor,
      executorArgs: "",
      tokenReceiver: "",
      tokenArgs: ""
    });

    bytes memory encoded = ExtraArgsCodecUnoptimized._encodeGenericExtraArgsV3(args);

    // Check the executor length field is 20
    // Format: 4 (tag) + 4 (gasLimit) + 2 (finality) + 1 (ccvs length) = 11 bytes offset
    // Then 1 byte for executor length
    uint8 executorLength = uint8(encoded[11]);
    assertEq(executorLength, 20, "Executor length should be 20 for non-zero address");
  }

  function test_decodeExecutorInvalidLength_Reverts() public {
    // Manually craft an encoded payload with invalid executor length (10 bytes)
    bytes memory invalidEncoded = abi.encodePacked(
      ExtraArgsCodec.GENERIC_EXTRA_ARGS_V3_TAG,
      uint32(200_000), // gasLimit
      uint16(12), // finalityConfig
      uint8(0), // ccvs length
      uint8(10), // executor length - INVALID (must be 0 or 20)
      bytes10(0x12345678901234567890), // 10 bytes of executor data
      uint16(0), // executorArgs length
      uint16(0), // tokenReceiver length
      uint16(0) // tokenArgs length
    );

    vm.expectRevert(abi.encodeWithSelector(ExtraArgsCodec.InvalidExecutorLength.selector, 10));
    helper.decode(invalidEncoded);
  }

  function test_decodeExecutorLength32_Reverts() public {
    // Manually craft an encoded payload with invalid executor length (32 bytes)
    bytes memory invalidEncoded = abi.encodePacked(
      ExtraArgsCodec.GENERIC_EXTRA_ARGS_V3_TAG,
      uint32(200_000), // gasLimit
      uint16(12), // finalityConfig
      uint8(0), // ccvs length
      uint8(32), // executor length - INVALID (must be 0 or 20)
      bytes32(0x1234567890123456789012345678901234567890123456789012345678901234), // 32 bytes
      uint16(0), // executorArgs length
      uint16(0), // tokenReceiver length
      uint16(0) // tokenArgs length
    );

    vm.expectRevert(abi.encodeWithSelector(ExtraArgsCodec.InvalidExecutorLength.selector, 32));
    helper.decode(invalidEncoded);
  }

  function test_encodeDecodeWithCCVs() public view {
    address executor = address(0x1234567890123456789012345678901234567890);
    address[] memory ccvAddresses = new address[](2);
    ccvAddresses[0] = address(0x1111);
    ccvAddresses[1] = address(0x2222);
    bytes[] memory ccvArgs = new bytes[](2);
    ccvArgs[0] = "args1";
    ccvArgs[1] = "args2";

    ExtraArgsCodec.GenericExtraArgsV3 memory args = ExtraArgsCodec.GenericExtraArgsV3({
      ccvs: ccvAddresses,
      ccvArgs: ccvArgs,
      finalityConfig: 12,
      gasLimit: 200_000,
      executor: executor,
      executorArgs: "execArgs",
      tokenReceiver: abi.encodePacked(address(0x3333)),
      tokenArgs: "tokenArgs"
    });

    bytes memory encoded = ExtraArgsCodecUnoptimized._encodeGenericExtraArgsV3(args);
    ExtraArgsCodec.GenericExtraArgsV3 memory decoded = helper.decode(encoded);

    assertEq(decoded.executor, executor, "Executor should match");
    assertEq(decoded.ccvs.length, 2, "CCVs length should match");
    assertEq(decoded.ccvs[0], address(0x1111), "CCV 0 address should match");
    assertEq(decoded.ccvArgs[0], "args1", "CCV 0 args should match");
    assertEq(decoded.ccvs[1], address(0x2222), "CCV 1 address should match");
    assertEq(decoded.ccvArgs[1], "args2", "CCV 1 args should match");
    assertEq(decoded.executorArgs, "execArgs", "ExecutorArgs should match");
    assertEq(decoded.tokenReceiver, abi.encodePacked(address(0x3333)), "TokenReceiver should match");
    assertEq(decoded.tokenArgs, "tokenArgs", "TokenArgs should match");
  }
}
