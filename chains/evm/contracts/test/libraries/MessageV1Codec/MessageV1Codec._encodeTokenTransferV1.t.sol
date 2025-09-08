// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {MessageV1Codec} from "../../../libraries/MessageV1Codec.sol";
import {MessageV1CodecSetup} from "./MessageV1CodecSetup.t.sol";

contract MessageV1Codec__encodeTokenTransferV1 is MessageV1CodecSetup {
  function test__encodeTokenTransferV1_checkRawBytes() public pure {
    MessageV1Codec.TokenTransferV1 memory tokenTransfer = MessageV1Codec.TokenTransferV1({
      amount: 100,
      sourcePoolAddress: hex"1234567890abcdef",
      sourceTokenAddress: hex"abcdef1234567890",
      destTokenAddress: hex"fedcba0987654321",
      extraData: hex"0c117e57da7a"
    });

    bytes memory encoded = MessageV1Codec._encodeTokenTransferV1(tokenTransfer);

    // Verify the encoding structure:
    // version (1 byte) + amount (32 bytes) + sourcePoolAddressLength (1 byte) + sourcePoolAddress +
    // sourceTokenAddressLength (1 byte) + sourceTokenAddress + destTokenAddressLength (1 byte) + destTokenAddress +
    // extraDataLength (2 bytes) + extraData
    uint256 expectedLength = 1 + 32 + 1 + tokenTransfer.sourcePoolAddress.length + 1
      + tokenTransfer.sourceTokenAddress.length + 1 + tokenTransfer.destTokenAddress.length + 2
      + tokenTransfer.extraData.length;
    assertEq(expectedLength, encoded.length);

    // Check version
    assertEq(1, uint8(encoded[0]));

    // Check amount
    bytes32 amountBytes;
    assembly {
      amountBytes := mload(add(encoded, 0x21)) // 0x20 (length prefix) + 0x01 (version byte)
    }
    assertEq(tokenTransfer.amount, uint256(amountBytes));

    // Check sourcePoolAddress length and content
    uint256 sourcePoolOffset = 33;
    assertEq(tokenTransfer.sourcePoolAddress.length, uint8(encoded[sourcePoolOffset]));
    bytes memory sourcePoolAddress = new bytes(tokenTransfer.sourcePoolAddress.length);
    for (uint256 i = 0; i < tokenTransfer.sourcePoolAddress.length; i++) {
      sourcePoolAddress[i] = encoded[sourcePoolOffset + 1 + i];
    }
    assertEq(keccak256(tokenTransfer.sourcePoolAddress), keccak256(sourcePoolAddress));

    // Check sourceTokenAddress length and content
    uint256 sourceTokenOffset = sourcePoolOffset + 1 + tokenTransfer.sourcePoolAddress.length;
    assertEq(tokenTransfer.sourceTokenAddress.length, uint8(encoded[sourceTokenOffset]));
    bytes memory sourceTokenAddress = new bytes(tokenTransfer.sourceTokenAddress.length);
    for (uint256 i = 0; i < tokenTransfer.sourceTokenAddress.length; i++) {
      sourceTokenAddress[i] = encoded[sourceTokenOffset + 1 + i];
    }
    assertEq(keccak256(tokenTransfer.sourceTokenAddress), keccak256(sourceTokenAddress));

    // Check destTokenAddress length and content
    uint256 destTokenOffset = sourceTokenOffset + 1 + tokenTransfer.sourceTokenAddress.length;
    assertEq(tokenTransfer.destTokenAddress.length, uint8(encoded[destTokenOffset]));
    bytes memory destTokenAddress = new bytes(tokenTransfer.destTokenAddress.length);
    for (uint256 i = 0; i < tokenTransfer.destTokenAddress.length; i++) {
      destTokenAddress[i] = encoded[destTokenOffset + 1 + i];
    }
    assertEq(keccak256(tokenTransfer.destTokenAddress), keccak256(destTokenAddress));

    // Check extraData length and content
    uint256 extraDataLengthOffset = destTokenOffset + 1 + tokenTransfer.destTokenAddress.length;
    assertEq(
      tokenTransfer.extraData.length,
      uint16(uint8(encoded[extraDataLengthOffset])) << 8 | uint16(uint8(encoded[extraDataLengthOffset + 1]))
    );
    bytes memory extraData = new bytes(tokenTransfer.extraData.length);
    for (uint256 i = 0; i < tokenTransfer.extraData.length; i++) {
      extraData[i] = encoded[extraDataLengthOffset + 2 + i];
    }
    assertEq(keccak256(tokenTransfer.extraData), keccak256(extraData));
  }

  function test__encodeTokenTransferV1_EmptyAddresses() public view {
    MessageV1Codec.TokenTransferV1 memory tokenTransfer = MessageV1Codec.TokenTransferV1({
      amount: 0,
      sourcePoolAddress: "",
      sourceTokenAddress: "",
      destTokenAddress: "",
      extraData: ""
    });

    bytes memory encoded = MessageV1Codec._encodeTokenTransferV1(tokenTransfer);

    // Minimum encoding: version (1) + amount (32) + 3 address lengths (3) + extraData length (2)
    assertEq(1 + 32 + 1 + 1 + 1 + 2, encoded.length);

    // Decode and verify the result matches the original.
    MessageV1Codec.TokenTransferV1 memory decoded = s_helper.decodeTokenTransferV1(encoded);

    assertEq(0, decoded.amount);
    assertEq(0, decoded.sourcePoolAddress.length);
    assertEq(0, decoded.sourceTokenAddress.length);
    assertEq(0, decoded.destTokenAddress.length);
    assertEq(0, decoded.extraData.length);
  }

  function test__encodeTokenTransferV1_EVMAddresses() public view {
    // Test with typical EVM addresses (20 bytes)
    MessageV1Codec.TokenTransferV1 memory tokenTransfer = MessageV1Codec.TokenTransferV1({
      amount: type(uint256).max,
      sourcePoolAddress: abi.encode(address(0x1234567890123456789012345678901234567890)),
      sourceTokenAddress: abi.encode(address(0xABcdEFABcdEFabcdEfAbCdefabcdeFABcDEFabCD)),
      destTokenAddress: abi.encode(address(0xfEdcBA9876543210FedCBa9876543210fEdCBa98)),
      extraData: hex"00112233445566778899aabbccddeeff"
    });

    bytes memory encoded = MessageV1Codec._encodeTokenTransferV1(tokenTransfer);

    // 1 + 32 + 1 + 32 + 1 + 32 + 1 + 32 + 2 + 16 = 150 bytes
    assertEq(150, encoded.length);

    // Decode and verify the result matches the original.
    MessageV1Codec.TokenTransferV1 memory decoded = s_helper.decodeTokenTransferV1(encoded);

    assertEq(type(uint256).max, decoded.amount);
    assertEq(keccak256(tokenTransfer.sourcePoolAddress), keccak256(decoded.sourcePoolAddress));
    assertEq(keccak256(tokenTransfer.sourceTokenAddress), keccak256(decoded.sourceTokenAddress));
    assertEq(keccak256(tokenTransfer.destTokenAddress), keccak256(decoded.destTokenAddress));
    assertEq(keccak256(tokenTransfer.extraData), keccak256(decoded.extraData));
  }

  function test__encodeTokenTransferV1_MaxLengthAddresses() public view {
    // Test with maximum uint8 length (255 bytes)
    bytes memory maxLengthAddress = new bytes(255);
    for (uint256 i = 0; i < 255; i++) {
      maxLengthAddress[i] = bytes1(uint8(i));
    }

    MessageV1Codec.TokenTransferV1 memory tokenTransfer = MessageV1Codec.TokenTransferV1({
      amount: 42,
      sourcePoolAddress: maxLengthAddress,
      sourceTokenAddress: maxLengthAddress,
      destTokenAddress: maxLengthAddress,
      extraData: hex""
    });

    bytes memory encoded = MessageV1Codec._encodeTokenTransferV1(tokenTransfer);

    // 1 + 32 + 1 + 1 + 1 + 1 + 1 + 1 + 2 = 38 static length
    assertEq(38 + 3 * maxLengthAddress.length, encoded.length);

    // Decode and verify the result matches the original.
    MessageV1Codec.TokenTransferV1 memory decoded = s_helper.decodeTokenTransferV1(encoded);

    assertEq(tokenTransfer.amount, decoded.amount);
    assertEq(keccak256(tokenTransfer.sourcePoolAddress), keccak256(decoded.sourcePoolAddress));
    assertEq(keccak256(tokenTransfer.sourceTokenAddress), keccak256(decoded.sourceTokenAddress));
    assertEq(keccak256(tokenTransfer.destTokenAddress), keccak256(decoded.destTokenAddress));
    assertEq(keccak256(tokenTransfer.extraData), keccak256(decoded.extraData));
  }

  function test__encodeTokenTransferV1_MaxLengthExtraData() public view {
    // Test with maximum uint16 length (65535 bytes)
    bytes memory maxLengthExtraData = new bytes(type(uint16).max);
    for (uint256 i = 0; i < 100; i++) {
      maxLengthExtraData[i] = bytes1(uint8(i));
    }

    MessageV1Codec.TokenTransferV1 memory tokenTransfer = MessageV1Codec.TokenTransferV1({
      amount: 999,
      sourcePoolAddress: hex"",
      sourceTokenAddress: hex"",
      destTokenAddress: hex"",
      extraData: maxLengthExtraData
    });

    bytes memory encoded = MessageV1Codec._encodeTokenTransferV1(tokenTransfer);

    // 1 + 32 + 1 + 1 + 1 + 1 + 1 + 1 + 2 = 38 static length
    assertEq(38 + maxLengthExtraData.length, encoded.length);

    // Decode and verify the result matches the original.
    MessageV1Codec.TokenTransferV1 memory decoded = s_helper.decodeTokenTransferV1(encoded);

    assertEq(tokenTransfer.amount, decoded.amount);
    assertEq(keccak256(tokenTransfer.sourcePoolAddress), keccak256(decoded.sourcePoolAddress));
    assertEq(keccak256(tokenTransfer.sourceTokenAddress), keccak256(decoded.sourceTokenAddress));
    assertEq(keccak256(tokenTransfer.destTokenAddress), keccak256(decoded.destTokenAddress));
    assertEq(type(uint16).max, decoded.extraData.length);
  }

  function testFuzz__encodeTokenTransferV1(
    uint256 amount,
    bytes memory sourcePoolAddress,
    bytes memory sourceTokenAddress,
    bytes memory destTokenAddress,
    bytes memory extraData
  ) public view {
    // Bound inputs to valid ranges
    vm.assume(sourcePoolAddress.length <= type(uint8).max);
    vm.assume(sourceTokenAddress.length <= type(uint8).max);
    vm.assume(destTokenAddress.length <= type(uint8).max);
    vm.assume(extraData.length <= type(uint16).max);

    MessageV1Codec.TokenTransferV1 memory tokenTransfer = MessageV1Codec.TokenTransferV1({
      amount: amount,
      sourcePoolAddress: sourcePoolAddress,
      sourceTokenAddress: sourceTokenAddress,
      destTokenAddress: destTokenAddress,
      extraData: extraData
    });

    bytes memory encoded = MessageV1Codec._encodeTokenTransferV1(tokenTransfer);

    // Verify encoding structure
    uint256 expectedLength = 1 + 32 + 1 + sourcePoolAddress.length + 1 + sourceTokenAddress.length + 1
      + destTokenAddress.length + 2 + extraData.length;
    assertEq(expectedLength, encoded.length);

    // Decode and verify the result matches the original.
    MessageV1Codec.TokenTransferV1 memory decoded = s_helper.decodeTokenTransferV1(encoded);

    assertEq(amount, decoded.amount);
    assertEq(keccak256(sourcePoolAddress), keccak256(decoded.sourcePoolAddress));
    assertEq(keccak256(sourceTokenAddress), keccak256(decoded.sourceTokenAddress));
    assertEq(keccak256(destTokenAddress), keccak256(decoded.destTokenAddress));
    assertEq(keccak256(extraData), keccak256(decoded.extraData));
  }

  // Reverts

  /// forge-config: default.allow_internal_expect_revert = true
  function test__encodeTokenTransferV1_RevertWhen_SourcePoolAddressTooLong() public {
    bytes memory tooLongAddress = new bytes(256);

    MessageV1Codec.TokenTransferV1 memory tokenTransfer = MessageV1Codec.TokenTransferV1({
      amount: 100,
      sourcePoolAddress: tooLongAddress,
      sourceTokenAddress: hex"abcd",
      destTokenAddress: hex"1234",
      extraData: hex""
    });

    vm.expectRevert(
      abi.encodeWithSelector(
        MessageV1Codec.InvalidDataLength.selector, MessageV1Codec.EncodingErrorLocation.ENCODE_TOKEN_SOURCE_POOL_LENGTH
      )
    );
    MessageV1Codec._encodeTokenTransferV1(tokenTransfer);
  }

  /// forge-config: default.allow_internal_expect_revert = true
  function test__encodeTokenTransferV1_RevertWhen_SourceTokenAddressTooLong() public {
    bytes memory tooLongAddress = new bytes(256);

    MessageV1Codec.TokenTransferV1 memory tokenTransfer = MessageV1Codec.TokenTransferV1({
      amount: 100,
      sourcePoolAddress: hex"abcd",
      sourceTokenAddress: tooLongAddress,
      destTokenAddress: hex"1234",
      extraData: hex""
    });

    vm.expectRevert(
      abi.encodeWithSelector(
        MessageV1Codec.InvalidDataLength.selector, MessageV1Codec.EncodingErrorLocation.ENCODE_TOKEN_SOURCE_TOKEN_LENGTH
      )
    );
    MessageV1Codec._encodeTokenTransferV1(tokenTransfer);
  }

  /// forge-config: default.allow_internal_expect_revert = true
  function test__encodeTokenTransferV1_RevertWhen_DestTokenAddressTooLong() public {
    bytes memory tooLongAddress = new bytes(256);

    MessageV1Codec.TokenTransferV1 memory tokenTransfer = MessageV1Codec.TokenTransferV1({
      amount: 100,
      sourcePoolAddress: hex"abcd",
      sourceTokenAddress: hex"1234",
      destTokenAddress: tooLongAddress,
      extraData: hex""
    });

    vm.expectRevert(
      abi.encodeWithSelector(
        MessageV1Codec.InvalidDataLength.selector, MessageV1Codec.EncodingErrorLocation.ENCODE_TOKEN_DEST_TOKEN_LENGTH
      )
    );
    MessageV1Codec._encodeTokenTransferV1(tokenTransfer);
  }

  /// forge-config: default.allow_internal_expect_revert = true
  function test__encodeTokenTransferV1_RevertWhen_ExtraDataTooLong() public {
    bytes memory tooLongExtraData = new bytes(65536); // uint16 max + 1

    MessageV1Codec.TokenTransferV1 memory tokenTransfer = MessageV1Codec.TokenTransferV1({
      amount: 100,
      sourcePoolAddress: hex"abcd",
      sourceTokenAddress: hex"1234",
      destTokenAddress: hex"5678",
      extraData: tooLongExtraData
    });

    vm.expectRevert(
      abi.encodeWithSelector(
        MessageV1Codec.InvalidDataLength.selector, MessageV1Codec.EncodingErrorLocation.ENCODE_TOKEN_EXTRA_DATA_LENGTH
      )
    );
    MessageV1Codec._encodeTokenTransferV1(tokenTransfer);
  }
}
