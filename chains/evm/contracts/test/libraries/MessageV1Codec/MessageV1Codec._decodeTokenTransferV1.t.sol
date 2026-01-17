// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {MessageV1Codec} from "../../../libraries/MessageV1Codec.sol";
import {MessageV1CodecSetup} from "./MessageV1CodecSetup.t.sol";

contract MessageV1Codec__decodeTokenTransferV1 is MessageV1CodecSetup {
  function test__decodeTokenTransferV1_BasicDecode() public view {
    MessageV1Codec.TokenTransferV1 memory originalTransfer = _createBasicTokenTransfer();
    bytes memory encoded = MessageV1Codec._encodeTokenTransferV1(originalTransfer);

    MessageV1Codec.TokenTransferV1 memory decoded = s_helper.decodeTokenTransferV1(encoded);

    assertEq(originalTransfer.amount, decoded.amount);
    assertEq(keccak256(originalTransfer.sourcePoolAddress), keccak256(decoded.sourcePoolAddress));
    assertEq(keccak256(originalTransfer.sourceTokenAddress), keccak256(decoded.sourceTokenAddress));
    assertEq(keccak256(originalTransfer.destTokenAddress), keccak256(decoded.destTokenAddress));
    assertEq(keccak256(originalTransfer.extraData), keccak256(decoded.extraData));
  }

  function test__decodeTokenTransferV1_EmptyFields() public view {
    MessageV1Codec.TokenTransferV1 memory originalTransfer = MessageV1Codec.TokenTransferV1({
      amount: 0,
      sourcePoolAddress: "",
      sourceTokenAddress: "",
      destTokenAddress: "",
      extraData: ""
    });

    bytes memory encoded = MessageV1Codec._encodeTokenTransferV1(originalTransfer);
    MessageV1Codec.TokenTransferV1 memory decoded = s_helper.decodeTokenTransferV1(encoded);

    assertEq(0, decoded.amount);
    assertEq(0, decoded.sourcePoolAddress.length);
    assertEq(0, decoded.sourceTokenAddress.length);
    assertEq(0, decoded.destTokenAddress.length);
    assertEq(0, decoded.extraData.length);
  }

  function test__decodeTokenTransferV1_MaxLengthFields() public view {
    bytes memory maxAddressLength = new bytes(type(uint8).max);
    bytes memory maxExtraDataLength = new bytes(type(uint16).max);

    // Fill with test patterns
    for (uint256 i = 0; i < maxAddressLength.length; i++) {
      maxAddressLength[i] = bytes1(uint8(i % 256));
    }
    for (uint256 i = 0; i < 100; i++) {
      maxExtraDataLength[i] = bytes1(uint8(i % 256));
    }

    MessageV1Codec.TokenTransferV1 memory originalTransfer = MessageV1Codec.TokenTransferV1({
      amount: type(uint256).max,
      sourcePoolAddress: maxAddressLength,
      sourceTokenAddress: maxAddressLength,
      destTokenAddress: maxAddressLength,
      extraData: maxExtraDataLength
    });

    bytes memory encoded = MessageV1Codec._encodeTokenTransferV1(originalTransfer);
    MessageV1Codec.TokenTransferV1 memory decoded = s_helper.decodeTokenTransferV1(encoded);

    assertEq(type(uint256).max, decoded.amount);
    assertEq(type(uint8).max, decoded.sourcePoolAddress.length);
    assertEq(type(uint8).max, decoded.sourceTokenAddress.length);
    assertEq(type(uint8).max, decoded.destTokenAddress.length);
    assertEq(type(uint16).max, decoded.extraData.length);
    assertEq(keccak256(originalTransfer.sourcePoolAddress), keccak256(decoded.sourcePoolAddress));
    assertEq(keccak256(originalTransfer.sourceTokenAddress), keccak256(decoded.sourceTokenAddress));
    assertEq(keccak256(originalTransfer.destTokenAddress), keccak256(decoded.destTokenAddress));
  }

  function testFuzz__decodeTokenTransferV1_RoundTrip(
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

    MessageV1Codec.TokenTransferV1 memory originalTransfer = MessageV1Codec.TokenTransferV1({
      amount: amount,
      sourcePoolAddress: sourcePoolAddress,
      sourceTokenAddress: sourceTokenAddress,
      destTokenAddress: destTokenAddress,
      extraData: extraData
    });

    bytes memory encoded = MessageV1Codec._encodeTokenTransferV1(originalTransfer);
    MessageV1Codec.TokenTransferV1 memory decoded = s_helper.decodeTokenTransferV1(encoded);

    assertEq(amount, decoded.amount);
    assertEq(keccak256(sourcePoolAddress), keccak256(decoded.sourcePoolAddress));
    assertEq(keccak256(sourceTokenAddress), keccak256(decoded.sourceTokenAddress));
    assertEq(keccak256(destTokenAddress), keccak256(decoded.destTokenAddress));
    assertEq(keccak256(extraData), keccak256(decoded.extraData));
  }

  // Reverts

  function test__decodeTokenTransferV1_RevertWhen_InvalidVersion() public {
    MessageV1Codec.TokenTransferV1 memory tokenTransfer = _createBasicTokenTransfer();
    bytes memory encoded = MessageV1Codec._encodeTokenTransferV1(tokenTransfer);

    // Corrupt the version byte (first byte should be 1)
    bytes memory corruptedEncoded = encoded;
    corruptedEncoded[0] = 0x02; // Invalid version

    vm.expectRevert(abi.encodeWithSelector(MessageV1Codec.InvalidEncodingVersion.selector, 2));
    s_helper.decodeTokenTransferV1(corruptedEncoded);
  }

  function test__decodeTokenTransferV1_RevertWhen_TruncatedAtVersion() public {
    bytes memory truncated = new bytes(0); // Empty data

    vm.expectRevert(
      abi.encodeWithSelector(
        MessageV1Codec.InvalidDataLength.selector, MessageV1Codec.EncodingErrorLocation.TOKEN_TRANSFER_VERSION
      )
    );
    s_helper.decodeTokenTransferV1(truncated);
  }

  function test__decodeTokenTransferV1_RevertWhen_TruncatedAtAmount() public {
    bytes memory partialData = new bytes(1 + 16); // Version + partial amount (32 bytes needed)
    partialData[0] = 0x01; // Valid version

    vm.expectRevert(
      abi.encodeWithSelector(
        MessageV1Codec.InvalidDataLength.selector, MessageV1Codec.EncodingErrorLocation.TOKEN_TRANSFER_AMOUNT
      )
    );
    s_helper.decodeTokenTransferV1(partialData);
  }

  function test__decodeTokenTransferV1_RevertWhen_TruncatedAtSourcePoolLength() public {
    bytes memory partialData = new bytes(1 + 32); // Version + amount, missing source pool length
    partialData[0] = 0x01; // Valid version

    vm.expectRevert(
      abi.encodeWithSelector(
        MessageV1Codec.InvalidDataLength.selector,
        MessageV1Codec.EncodingErrorLocation.TOKEN_TRANSFER_SOURCE_POOL_LENGTH
      )
    );
    s_helper.decodeTokenTransferV1(partialData);
  }

  function test__decodeTokenTransferV1_RevertWhen_TruncatedAtSourcePoolContent() public {
    bytes memory partialData = new bytes(1 + 32 + 1 + 5); // Version + amount + length(10) + partial content(5)
    partialData[0] = 0x01; // Valid version
    partialData[33] = 0x0A; // Source pool length = 10, but only 5 bytes provided

    vm.expectRevert(
      abi.encodeWithSelector(
        MessageV1Codec.InvalidDataLength.selector,
        MessageV1Codec.EncodingErrorLocation.TOKEN_TRANSFER_SOURCE_POOL_CONTENT
      )
    );
    s_helper.decodeTokenTransferV1(partialData);
  }

  function test__decodeTokenTransferV1_RevertWhen_TruncatedAtSourceTokenLength() public {
    bytes memory partialData = new bytes(1 + 32 + 1 + 8); // Version + amount + sourcePoolLength(8) + sourcePool(8), missing sourceTokenLength
    partialData[0] = 0x01; // Valid version
    partialData[33] = 0x08; // Source pool length = 8

    vm.expectRevert(
      abi.encodeWithSelector(
        MessageV1Codec.InvalidDataLength.selector,
        MessageV1Codec.EncodingErrorLocation.TOKEN_TRANSFER_SOURCE_TOKEN_LENGTH
      )
    );
    s_helper.decodeTokenTransferV1(partialData);
  }

  function test__decodeTokenTransferV1_RevertWhen_TruncatedAtSourceTokenContent() public {
    bytes memory partialData = new bytes(1 + 32 + 1 + 8 + 1 + 3); // Missing sourceToken content
    partialData[0] = 0x01; // Valid version
    partialData[33] = 0x08; // Source pool length = 8
    partialData[42] = 0x06; // Source token length = 6, but only 3 bytes provided

    vm.expectRevert(
      abi.encodeWithSelector(
        MessageV1Codec.InvalidDataLength.selector,
        MessageV1Codec.EncodingErrorLocation.TOKEN_TRANSFER_SOURCE_TOKEN_CONTENT
      )
    );
    s_helper.decodeTokenTransferV1(partialData);
  }

  function test__decodeTokenTransferV1_RevertWhen_TruncatedAtDestTokenLength() public {
    bytes memory partialData = new bytes(1 + 32 + 1 + 8 + 1 + 8); // Missing destTokenLength
    partialData[0] = 0x01; // Valid version
    partialData[33] = 0x08; // Source pool length = 8
    partialData[42] = 0x08; // Source token length = 8

    vm.expectRevert(
      abi.encodeWithSelector(
        MessageV1Codec.InvalidDataLength.selector, MessageV1Codec.EncodingErrorLocation.TOKEN_TRANSFER_DEST_TOKEN_LENGTH
      )
    );
    s_helper.decodeTokenTransferV1(partialData);
  }

  function test__decodeTokenTransferV1_RevertWhen_TruncatedAtDestTokenContent() public {
    bytes memory partialData = new bytes(1 + 32 + 1 + 8 + 1 + 8 + 1 + 4); // Missing destToken content
    partialData[0] = 0x01; // Valid version
    partialData[33] = 0x08; // Source pool length = 8
    partialData[42] = 0x08; // Source token length = 8
    partialData[51] = 0x08; // Dest token length = 8, but only 4 bytes provided

    vm.expectRevert(
      abi.encodeWithSelector(
        MessageV1Codec.InvalidDataLength.selector,
        MessageV1Codec.EncodingErrorLocation.TOKEN_TRANSFER_DEST_TOKEN_CONTENT
      )
    );
    s_helper.decodeTokenTransferV1(partialData);
  }

  function test__decodeTokenTransferV1_RevertWhen_TruncatedAtExtraDataLength() public {
    bytes memory partialData = new bytes(1 + 32 + 1 + 8 + 1 + 8 + 1 + 8 + 1); // Missing extraData length (2 bytes needed)
    partialData[0] = 0x01; // Valid version
    partialData[33] = 0x08; // Source pool length = 8
    partialData[42] = 0x08; // Source token length = 8
    partialData[51] = 0x08; // Dest token length = 8

    vm.expectRevert(
      abi.encodeWithSelector(
        MessageV1Codec.InvalidDataLength.selector, MessageV1Codec.EncodingErrorLocation.TOKEN_TRANSFER_EXTRA_DATA_LENGTH
      )
    );
    s_helper.decodeTokenTransferV1(partialData);
  }

  function test__decodeTokenTransferV1_RevertWhen_TruncatedAtExtraDataContent() public {
    bytes memory partialData = new bytes(1 + 32 + 1 + 8 + 1 + 8 + 1 + 8 + 2 + 5); // Missing extraData content
    partialData[0] = 0x01; // Valid version
    partialData[33] = 0x08; // Source pool length = 8
    partialData[42] = 0x08; // Source token length = 8
    partialData[51] = 0x08; // Dest token length = 8
    partialData[60] = 0x00; // Extra data length high byte = 0
    partialData[61] = 0x0A; // Extra data length low byte = 10, but only 5 bytes provided

    vm.expectRevert(
      abi.encodeWithSelector(
        MessageV1Codec.InvalidDataLength.selector,
        MessageV1Codec.EncodingErrorLocation.TOKEN_TRANSFER_EXTRA_DATA_CONTENT
      )
    );
    s_helper.decodeTokenTransferV1(partialData);
  }
}
