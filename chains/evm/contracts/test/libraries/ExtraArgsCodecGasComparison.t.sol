// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Test} from "forge-std/Test.sol";
import {ExtraArgsCodecUnoptimized} from "../../libraries/ExtraArgsCodecUnoptimized.sol";
import {ExtraArgsCodec} from "../../libraries/ExtraArgsCodec.sol";

/// @notice Helper contract to decode with original implementation
contract OriginalDecoder {
  function decode(bytes calldata data) external pure returns (ExtraArgsCodecUnoptimized.GenericExtraArgsV3 memory) {
    return ExtraArgsCodecUnoptimized._decodeGenericExtraArgsV3(data);
  }

  function decodeNoReturn(bytes calldata data) external pure {
    ExtraArgsCodecUnoptimized._decodeGenericExtraArgsV3(data);
  }
}

/// @notice Helper contract to decode with optimized implementation
contract OptimizedDecoder {
  function decode(bytes calldata data) external pure returns (ExtraArgsCodec.GenericExtraArgsV3 memory) {
    return ExtraArgsCodec._decodeGenericExtraArgsV3(data);
  }

  function decodeNoReturn(bytes calldata data) external pure {
    ExtraArgsCodec._decodeGenericExtraArgsV3(data);
  }
}

/// @notice Gas comparison test between original and optimized ExtraArgsCodec implementations.
/// @dev This test suite compares gas costs for encoding and decoding operations.
contract ExtraArgsCodecGasComparison is Test {
  OriginalDecoder originalDecoder;
  OptimizedDecoder optimizedDecoder;

  function setUp() public {
    originalDecoder = new OriginalDecoder();
    optimizedDecoder = new OptimizedDecoder();
  }
  /// @notice Test encoding with no CCVs - comparing gas costs
  function test_gas_encode_noCCVs() public {
    ExtraArgsCodecUnoptimized.GenericExtraArgsV3 memory args = ExtraArgsCodecUnoptimized.GenericExtraArgsV3({
      ccvs: new address[](0),
      ccvArgs: new bytes[](0),
      finalityConfig: 12,
      gasLimit: 200_000,
      executor: address(0x1234567890123456789012345678901234567890),
      executorArgs: "some executor args here",
      tokenReceiver: abi.encodePacked(address(0x9876543210987654321098765432109876543210)),
      tokenArgs: "token args data"
    });

    // Original implementation
    uint256 gasBefore = gasleft();
    bytes memory encoded1 = ExtraArgsCodecUnoptimized._encodeGenericExtraArgsV3(args);
    uint256 gasUsedOriginal = gasBefore - gasleft();

    // Optimized implementation
    ExtraArgsCodec.GenericExtraArgsV3 memory argsOpt = ExtraArgsCodec.GenericExtraArgsV3({
      ccvs: args.ccvs,
      ccvArgs: args.ccvArgs,
      finalityConfig: args.finalityConfig,
      gasLimit: args.gasLimit,
      executor: args.executor,
      executorArgs: args.executorArgs,
      tokenReceiver: args.tokenReceiver,
      tokenArgs: args.tokenArgs
    });

    gasBefore = gasleft();
    bytes memory encoded2 = ExtraArgsCodec._encodeGenericExtraArgsV3(argsOpt);
    uint256 gasUsedOptimized = gasBefore - gasleft();

    // Verify outputs are identical
    assertEq(encoded1, encoded2, "Encoded outputs should match");

    // Log gas comparison
    emit log_named_uint("Original gas (no CCVs)", gasUsedOriginal);
    emit log_named_uint("Optimized gas (no CCVs)", gasUsedOptimized);
    emit log_named_uint("Gas saved (no CCVs)", gasUsedOriginal - gasUsedOptimized);
    emit log_named_decimal_uint(
      "Gas savings % (no CCVs)", ((gasUsedOriginal - gasUsedOptimized) * 10000) / gasUsedOriginal, 2
    );
  }

  /// @notice Test encoding with 1 CCV
  function test_gas_encode_1CCV() public {
    address[] memory ccvs = new address[](1);
    ccvs[0] = address(0x1111111111111111111111111111111111111111);
    bytes[] memory ccvArgs = new bytes[](1);
    ccvArgs[0] = "ccv args 1";

    ExtraArgsCodecUnoptimized.GenericExtraArgsV3 memory args = ExtraArgsCodecUnoptimized.GenericExtraArgsV3({
      ccvs: ccvs,
      ccvArgs: ccvArgs,
      finalityConfig: 12,
      gasLimit: 200_000,
      executor: address(0x1234567890123456789012345678901234567890),
      executorArgs: "executor args",
      tokenReceiver: abi.encodePacked(address(0x9876543210987654321098765432109876543210)),
      tokenArgs: "token args"
    });

    // Original implementation
    uint256 gasBefore = gasleft();
    bytes memory encoded1 = ExtraArgsCodecUnoptimized._encodeGenericExtraArgsV3(args);
    uint256 gasUsedOriginal = gasBefore - gasleft();

    // Optimized implementation
    ExtraArgsCodec.GenericExtraArgsV3 memory argsOpt = ExtraArgsCodec.GenericExtraArgsV3({
      ccvs: ccvs,
      ccvArgs: ccvArgs,
      finalityConfig: args.finalityConfig,
      gasLimit: args.gasLimit,
      executor: args.executor,
      executorArgs: args.executorArgs,
      tokenReceiver: args.tokenReceiver,
      tokenArgs: args.tokenArgs
    });

    gasBefore = gasleft();
    bytes memory encoded2 = ExtraArgsCodec._encodeGenericExtraArgsV3(argsOpt);
    uint256 gasUsedOptimized = gasBefore - gasleft();

    // Verify outputs are identical
    assertEq(encoded1, encoded2, "Encoded outputs should match");

    // Log gas comparison
    emit log_named_uint("Original gas (1 CCV)", gasUsedOriginal);
    emit log_named_uint("Optimized gas (1 CCV)", gasUsedOptimized);
    emit log_named_uint("Gas saved (1 CCV)", gasUsedOriginal - gasUsedOptimized);
    emit log_named_decimal_uint(
      "Gas savings % (1 CCV)", ((gasUsedOriginal - gasUsedOptimized) * 10000) / gasUsedOriginal, 2
    );
  }

  /// @notice Test encoding with 3 CCVs
  function test_gas_encode_3CCVs() public {
    address[] memory ccvs = new address[](3);
    ccvs[0] = address(0x1111111111111111111111111111111111111111);
    ccvs[1] = address(0x2222222222222222222222222222222222222222);
    ccvs[2] = address(0x3333333333333333333333333333333333333333);
    
    bytes[] memory ccvArgs = new bytes[](3);
    ccvArgs[0] = "ccv args 1 with more data";
    ccvArgs[1] = "ccv args 2";
    ccvArgs[2] = "ccv args 3 longer arguments here";

    ExtraArgsCodecUnoptimized.GenericExtraArgsV3 memory args = ExtraArgsCodecUnoptimized.GenericExtraArgsV3({
      ccvs: ccvs,
      ccvArgs: ccvArgs,
      finalityConfig: 12,
      gasLimit: 200_000,
      executor: address(0x1234567890123456789012345678901234567890),
      executorArgs: "executor args with some data",
      tokenReceiver: abi.encodePacked(address(0x9876543210987654321098765432109876543210)),
      tokenArgs: "token args with data"
    });

    // Original implementation
    uint256 gasBefore = gasleft();
    bytes memory encoded1 = ExtraArgsCodecUnoptimized._encodeGenericExtraArgsV3(args);
    uint256 gasUsedOriginal = gasBefore - gasleft();

    // Optimized implementation
    ExtraArgsCodec.GenericExtraArgsV3 memory argsOpt = ExtraArgsCodec.GenericExtraArgsV3({
      ccvs: ccvs,
      ccvArgs: ccvArgs,
      finalityConfig: args.finalityConfig,
      gasLimit: args.gasLimit,
      executor: args.executor,
      executorArgs: args.executorArgs,
      tokenReceiver: args.tokenReceiver,
      tokenArgs: args.tokenArgs
    });

    gasBefore = gasleft();
    bytes memory encoded2 = ExtraArgsCodec._encodeGenericExtraArgsV3(argsOpt);
    uint256 gasUsedOptimized = gasBefore - gasleft();

    // Verify outputs are identical
    assertEq(encoded1, encoded2, "Encoded outputs should match");

    // Log gas comparison
    emit log_named_uint("Original gas (3 CCVs)", gasUsedOriginal);
    emit log_named_uint("Optimized gas (3 CCVs)", gasUsedOptimized);
    emit log_named_uint("Gas saved (3 CCVs)", gasUsedOriginal - gasUsedOptimized);
    emit log_named_decimal_uint(
      "Gas savings % (3 CCVs)", ((gasUsedOriginal - gasUsedOptimized) * 10000) / gasUsedOriginal, 2
    );
  }

  /// @notice Test decoding with no CCVs
  function test_gas_decode_noCCVs() public {
    ExtraArgsCodecUnoptimized.GenericExtraArgsV3 memory args = ExtraArgsCodecUnoptimized.GenericExtraArgsV3({
      ccvs: new address[](0),
      ccvArgs: new bytes[](0),
      finalityConfig: 12,
      gasLimit: 200_000,
      executor: address(0x1234567890123456789012345678901234567890),
      executorArgs: "some executor args here",
      tokenReceiver: abi.encodePacked(address(0x9876543210987654321098765432109876543210)),
      tokenArgs: "token args data"
    });

    bytes memory encoded = ExtraArgsCodecUnoptimized._encodeGenericExtraArgsV3(args);

    // Original implementation
    uint256 gasBefore = gasleft();
    ExtraArgsCodecUnoptimized.GenericExtraArgsV3 memory decoded1 = originalDecoder.decode(encoded);
    uint256 gasUsedOriginal = gasBefore - gasleft();

    // Optimized implementation
    gasBefore = gasleft();
    ExtraArgsCodec.GenericExtraArgsV3 memory decoded2 = optimizedDecoder.decode(encoded);
    uint256 gasUsedOptimized = gasBefore - gasleft();

    // Verify outputs match
    assertEq(decoded1.finalityConfig, decoded2.finalityConfig);
    assertEq(decoded1.gasLimit, decoded2.gasLimit);
    assertEq(decoded1.executor, decoded2.executor);
    assertEq(decoded1.executorArgs, decoded2.executorArgs);

    // Log gas comparison
    emit log_named_uint("Original gas decode (no CCVs)", gasUsedOriginal);
    emit log_named_uint("Optimized gas decode (no CCVs)", gasUsedOptimized);
    emit log_named_uint("Gas saved decode (no CCVs)", gasUsedOriginal - gasUsedOptimized);
    emit log_named_decimal_uint(
      "Gas savings % decode (no CCVs)", ((gasUsedOriginal - gasUsedOptimized) * 10000) / gasUsedOriginal, 2
    );
  }

  /// @notice Test decoding with 1 CCV
  function test_gas_decode_1CCV() public {
    address[] memory ccvs = new address[](1);
    ccvs[0] = address(0x1111111111111111111111111111111111111111);
    bytes[] memory ccvArgs = new bytes[](1);
    ccvArgs[0] = "ccv args 1";

    ExtraArgsCodecUnoptimized.GenericExtraArgsV3 memory args = ExtraArgsCodecUnoptimized.GenericExtraArgsV3({
      ccvs: ccvs,
      ccvArgs: ccvArgs,
      finalityConfig: 12,
      gasLimit: 200_000,
      executor: address(0x1234567890123456789012345678901234567890),
      executorArgs: "executor args",
      tokenReceiver: abi.encodePacked(address(0x9876543210987654321098765432109876543210)),
      tokenArgs: "token args"
    });

    bytes memory encoded = ExtraArgsCodecUnoptimized._encodeGenericExtraArgsV3(args);

    // Original implementation
    uint256 gasBefore = gasleft();
    ExtraArgsCodecUnoptimized.GenericExtraArgsV3 memory decoded1 = originalDecoder.decode(encoded);
    uint256 gasUsedOriginal = gasBefore - gasleft();

    // Optimized implementation
    gasBefore = gasleft();
    ExtraArgsCodec.GenericExtraArgsV3 memory decoded2 = optimizedDecoder.decode(encoded);
    uint256 gasUsedOptimized = gasBefore - gasleft();

    // Verify outputs match
    assertEq(decoded1.ccvs.length, decoded2.ccvs.length);
    assertEq(decoded1.ccvs[0], decoded2.ccvs[0]);
    assertEq(decoded1.ccvArgs[0], decoded2.ccvArgs[0]);

    // Log gas comparison
    emit log_named_uint("Original gas decode (1 CCV)", gasUsedOriginal);
    emit log_named_uint("Optimized gas decode (1 CCV)", gasUsedOptimized);
    emit log_named_uint("Gas saved decode (1 CCV)", gasUsedOriginal - gasUsedOptimized);
    emit log_named_decimal_uint(
      "Gas savings % decode (1 CCV)", ((gasUsedOriginal - gasUsedOptimized) * 10000) / gasUsedOriginal, 2
    );
  }

  /// @notice Test decoding with 3 CCVs
  function test_gas_decode_3CCVs() public {
    address[] memory ccvs = new address[](3);
    ccvs[0] = address(0x1111111111111111111111111111111111111111);
    ccvs[1] = address(0x2222222222222222222222222222222222222222);
    ccvs[2] = address(0x3333333333333333333333333333333333333333);
    
    bytes[] memory ccvArgs = new bytes[](3);
    ccvArgs[0] = "ccv args 1 with more data";
    ccvArgs[1] = "ccv args 2";
    ccvArgs[2] = "ccv args 3 longer arguments here";

    ExtraArgsCodecUnoptimized.GenericExtraArgsV3 memory args = ExtraArgsCodecUnoptimized.GenericExtraArgsV3({
      ccvs: ccvs,
      ccvArgs: ccvArgs,
      finalityConfig: 12,
      gasLimit: 200_000,
      executor: address(0x1234567890123456789012345678901234567890),
      executorArgs: "executor args with some data",
      tokenReceiver: abi.encodePacked(address(0x9876543210987654321098765432109876543210)),
      tokenArgs: "token args with data"
    });

    bytes memory encoded = ExtraArgsCodecUnoptimized._encodeGenericExtraArgsV3(args);

    // Original implementation
    uint256 gasBefore = gasleft();
    ExtraArgsCodecUnoptimized.GenericExtraArgsV3 memory decoded1 = originalDecoder.decode(encoded);
    uint256 gasUsedOriginal = gasBefore - gasleft();

    // Optimized implementation
    gasBefore = gasleft();
    ExtraArgsCodec.GenericExtraArgsV3 memory decoded2 = optimizedDecoder.decode(encoded);
    uint256 gasUsedOptimized = gasBefore - gasleft();

    // Verify outputs match
    assertEq(decoded1.ccvs.length, decoded2.ccvs.length);
    for (uint256 i = 0; i < decoded1.ccvs.length; i++) {
      assertEq(decoded1.ccvs[i], decoded2.ccvs[i]);
      assertEq(decoded1.ccvArgs[i], decoded2.ccvArgs[i]);
    }

    // Log gas comparison
    emit log_named_uint("Original gas decode (3 CCVs)", gasUsedOriginal);
    emit log_named_uint("Optimized gas decode (3 CCVs)", gasUsedOptimized);
    emit log_named_uint("Gas saved decode (3 CCVs)", gasUsedOriginal - gasUsedOptimized);
    emit log_named_decimal_uint(
      "Gas savings % decode (3 CCVs)", ((gasUsedOriginal - gasUsedOptimized) * 10000) / gasUsedOriginal, 2
    );
  }

  function test_gas_decode_empty() public {
    ExtraArgsCodecUnoptimized.GenericExtraArgsV3 memory args = ExtraArgsCodecUnoptimized.GenericExtraArgsV3({
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

    // Original implementation
    uint256 gasBefore = gasleft();
     originalDecoder.decodeNoReturn(encoded);
    uint256 gasUsedOriginal = gasBefore - gasleft();

    // Optimized implementation
    gasBefore = gasleft();
    optimizedDecoder.decodeNoReturn(encoded);
    uint256 gasUsedOptimized = gasBefore - gasleft();


    // Log gas comparison
    emit log_named_uint("Original gas decode empty", gasUsedOriginal);
    emit log_named_uint("Optimized gas decode empty", gasUsedOptimized);
    emit log_named_uint("Gas saved decode empty", gasUsedOriginal - gasUsedOptimized);
    emit log_named_decimal_uint(
      "Gas savings % decode empty", ((gasUsedOriginal - gasUsedOptimized) * 10000) / gasUsedOriginal, 2
    );
  }

  /// @notice Test encoding with zero executor
  function test_gas_encode_empty() public {
    ExtraArgsCodecUnoptimized.GenericExtraArgsV3 memory args = ExtraArgsCodecUnoptimized.GenericExtraArgsV3({
      ccvs: new address[](0),
      ccvArgs: new bytes[](0),
      finalityConfig: 12,
      gasLimit: 200_000,
      executor: address(0), // Zero executor
      executorArgs: "",
      tokenReceiver: "",
      tokenArgs: ""
    });

    // Original implementation
    uint256 gasBefore = gasleft();
    bytes memory encoded1 = ExtraArgsCodecUnoptimized._encodeGenericExtraArgsV3(args);
    uint256 gasUsedOriginal = gasBefore - gasleft();

    // Optimized implementation
    ExtraArgsCodec.GenericExtraArgsV3 memory argsOpt = ExtraArgsCodec.GenericExtraArgsV3({
      ccvs: new address[](0),
      ccvArgs: new bytes[](0),
      finalityConfig: args.finalityConfig,
      gasLimit: args.gasLimit,
      executor: args.executor,
      executorArgs: args.executorArgs,
      tokenReceiver: args.tokenReceiver,
      tokenArgs: args.tokenArgs
    });

    gasBefore = gasleft();
    bytes memory encoded2 = ExtraArgsCodec._encodeGenericExtraArgsV3(argsOpt);
    uint256 gasUsedOptimized = gasBefore - gasleft();

    // Verify outputs are identical
    assertEq(encoded1, encoded2, "Encoded outputs should match");

    // Log gas comparison
    emit log_named_uint("Original gas (zero executor)", gasUsedOriginal);
    emit log_named_uint("Optimized gas (zero executor)", gasUsedOptimized);
    emit log_named_uint("Gas saved (zero executor)", gasUsedOriginal - gasUsedOptimized);
    emit log_named_decimal_uint(
      "Gas savings % (zero executor)", ((gasUsedOriginal - gasUsedOptimized) * 10000) / gasUsedOriginal, 2
    );
  }
}

