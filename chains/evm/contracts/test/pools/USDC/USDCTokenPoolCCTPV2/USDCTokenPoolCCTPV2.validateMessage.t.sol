// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {USDCSourcePoolDataCodec} from "../../../../libraries/USDCSourcePoolDataCodec.sol";
import {USDCTokenPool} from "../../../../pools/USDC/USDCTokenPool.sol";
import {USDCTokenPoolCCTPV2} from "../../../../pools/USDC/USDCTokenPoolCCTPV2.sol";
import {USDCTokenPoolCCTPV2Setup} from "./USDCTokenPoolCCTPV2Setup.t.sol";

contract USDCTokenPoolCCTPV2_validateMessage is USDCTokenPoolCCTPV2Setup {
  USDCMessageCCTPV2 internal s_validUsdcMessage;
  USDCSourcePoolDataCodec.SourceTokenDataPayloadV2 internal s_validSourceTokenData;

  function setUp() public virtual override {
    super.setUp();

    // Create a valid USDC message for testing - this serves as the base message
    // that individual tests can copy and modify to test specific validation failures
    s_validUsdcMessage = USDCMessageCCTPV2({
      version: 1,
      sourceDomain: 1553252,
      destinationDomain: DEST_DOMAIN_IDENTIFIER,
      nonce: keccak256("0xC11"),
      sender: SOURCE_CHAIN_TOKEN_SENDER,
      recipient: bytes32(uint256(92398429395823)),
      destinationCaller: bytes32(uint256(uint160(address(s_usdcTokenPool)))),
      minFinalityThreshold: s_usdcTokenPool.FINALITY_THRESHOLD(),
      finalityThresholdExecuted: s_usdcTokenPool.FINALITY_THRESHOLD(),
      messageBody: abi.encodePacked(
        uint32(1), // version
        bytes32(uint256(uint160(address(s_USDCToken)))), // burnToken
        bytes32(uint256(uint160(299999))), // mintRecipient
        uint256(1e6), // amount
        bytes32(SOURCE_CHAIN_TOKEN_SENDER) // messageSender
      )
    });

    // Create valid source token data payload that matches the valid USDC message
    // This is used by tests that need to test payload validation failures
    s_validSourceTokenData = USDCSourcePoolDataCodec.SourceTokenDataPayloadV2({
      sourceDomain: s_validUsdcMessage.sourceDomain,
      depositHash: bytes32(0)
    });
  }

  function testFuzz_validateMessage_Success(
    uint32 sourceDomain,
    bytes32 nonce,
    uint256 amount,
    bytes32 mintRecipient
  ) public {
    vm.assume(amount != 0);
    vm.assume(mintRecipient != bytes32(0));

    vm.pauseGasMetering();
    USDCMessageCCTPV2 memory usdcMessage = USDCMessageCCTPV2({
      version: 1,
      sourceDomain: sourceDomain,
      destinationDomain: DEST_DOMAIN_IDENTIFIER,
      nonce: nonce,
      sender: SOURCE_CHAIN_TOKEN_SENDER,
      recipient: mintRecipient,
      destinationCaller: bytes32(uint256(uint160(address(s_usdcTokenPool)))),
      minFinalityThreshold: s_usdcTokenPool.FINALITY_THRESHOLD(),
      finalityThresholdExecuted: s_usdcTokenPool.FINALITY_THRESHOLD(),
      messageBody: abi.encodePacked(
        uint32(1), // version
        bytes32(uint256(uint160(address(s_USDCToken)))), // burnToken
        mintRecipient, // mintRecipient
        amount, // amount
        bytes32(SOURCE_CHAIN_TOKEN_SENDER) // messageSender
      )
    });

    bytes memory encodedUsdcMessage = _generateUSDCMessageCCTPV2(usdcMessage);

    bytes32 depositHash = USDCSourcePoolDataCodec._calculateDepositHash(
      sourceDomain,
      amount,
      DEST_DOMAIN_IDENTIFIER,
      mintRecipient,
      bytes32(uint256(uint160(address(s_USDCToken)))),
      bytes32(uint256(uint160(address(s_usdcTokenPool)))),
      s_usdcTokenPool.MAX_FEE(),
      s_usdcTokenPool.FINALITY_THRESHOLD()
    );

    vm.resumeGasMetering();
    s_usdcTokenPool.validateMessage(
      encodedUsdcMessage,
      USDCSourcePoolDataCodec.SourceTokenDataPayloadV2({sourceDomain: sourceDomain, depositHash: depositHash})
    );
  }

  // Reverts

  function test_validateMessage_RevertWhen_InvalidSourceDomain() public {
    uint32 expectedSourceDomain = s_validUsdcMessage.sourceDomain + 1;

    vm.expectRevert(
      abi.encodeWithSelector(
        USDCTokenPool.InvalidSourceDomain.selector, expectedSourceDomain, s_validUsdcMessage.sourceDomain
      )
    );
    s_usdcTokenPool.validateMessage(
      _generateUSDCMessageCCTPV2(s_validUsdcMessage),
      USDCSourcePoolDataCodec.SourceTokenDataPayloadV2({sourceDomain: expectedSourceDomain, depositHash: bytes32(0)})
    );
  }

  function test_validateMessage_RevertWhen_InvalidDestinationDomain() public {
    // Create a message with invalid destination domain
    USDCMessageCCTPV2 memory invalidMessage = s_validUsdcMessage;
    invalidMessage.destinationDomain = DEST_DOMAIN_IDENTIFIER + 1;

    vm.expectRevert(
      abi.encodeWithSelector(
        USDCTokenPool.InvalidDestinationDomain.selector, DEST_DOMAIN_IDENTIFIER, invalidMessage.destinationDomain
      )
    );

    s_usdcTokenPool.validateMessage(_generateUSDCMessageCCTPV2(invalidMessage), s_validSourceTokenData);
  }

  function test_validateMessage_RevertWhen_InvalidMessageVersion() public {
    // Create a message with invalid version
    USDCMessageCCTPV2 memory invalidMessage = s_validUsdcMessage;
    invalidMessage.version = 2;

    vm.expectRevert(abi.encodeWithSelector(USDCTokenPool.InvalidMessageVersion.selector, 1, 2));
    s_usdcTokenPool.validateMessage(_generateUSDCMessageCCTPV2(invalidMessage), s_validSourceTokenData);
  }

  function test_validateMessage_RevertWhen_InvalidMessageLength() public {
    bytes memory shortMessage = new bytes(100);
    vm.expectRevert(abi.encodeWithSelector(USDCTokenPool.InvalidMessageLength.selector, 100));
    s_usdcTokenPool.validateMessage(shortMessage, s_validSourceTokenData);
  }

  function test_validateMessage_RevertWhen_InvalidMinFinalityThreshold() public {
    // Create a message with invalid min finality threshold
    USDCMessageCCTPV2 memory invalidMessage = s_validUsdcMessage;
    invalidMessage.minFinalityThreshold = s_usdcTokenPool.FINALITY_THRESHOLD() + 1;

    vm.expectRevert(
      abi.encodeWithSelector(
        USDCTokenPoolCCTPV2.InvalidMinFinalityThreshold.selector,
        s_usdcTokenPool.FINALITY_THRESHOLD(),
        invalidMessage.minFinalityThreshold
      )
    );
    s_usdcTokenPool.validateMessage(_generateUSDCMessageCCTPV2(invalidMessage), s_validSourceTokenData);
  }

  function test_validateMessage_RevertWhen_InvalidExecutionFinalityThreshold() public {
    // Create a message with invalid execution finality threshold
    USDCMessageCCTPV2 memory invalidMessage = s_validUsdcMessage;
    invalidMessage.finalityThresholdExecuted = s_usdcTokenPool.FINALITY_THRESHOLD() + 1;

    vm.expectRevert(
      abi.encodeWithSelector(
        USDCTokenPoolCCTPV2.InvalidExecutionFinalityThreshold.selector,
        s_usdcTokenPool.FINALITY_THRESHOLD(),
        invalidMessage.finalityThresholdExecuted
      )
    );
    s_usdcTokenPool.validateMessage(_generateUSDCMessageCCTPV2(invalidMessage), s_validSourceTokenData);
  }

  function testFuzz_validateMessage_RevertWhen_InvalidDepositHash(
    uint256 seed
  ) public {
    // Generate 280 pseudo-random bytes using the seed which will become the message body
    bytes memory randomBytes = new bytes(260);
    for (uint256 i = 0; i < 260; i++) {
      randomBytes[i] = bytes1(uint8(uint256(keccak256(abi.encodePacked(seed, i))) % 256));
    }

    // Set the message body to the pseudo-random bytes but preserving the rest of the
    // message header to pass the validation checks
    s_validUsdcMessage.messageBody = randomBytes;
    bytes memory usdcMessage = _generateUSDCMessageCCTPV2(s_validUsdcMessage);

    // Define the fields from the message header and body so that we can calculate the invalid deposit hash based on the message
    uint32 messageSourceDomain;
    uint32 destinationDomain;
    uint32 minFinalityThreshold;
    uint32 finalityThresholdExecuted;
    bytes32 destinationCaller;
    uint256 amount;
    bytes32 burnToken;
    bytes32 mintRecipient;

    // solhint-disable-next-line no-inline-assembly
    assembly {
      // Parse the message header and body into the fields
      messageSourceDomain := mload(add(usdcMessage, 8)) // 4 + 4 = 8
      destinationDomain := mload(add(usdcMessage, 12)) // 8 + 4 = 12
      destinationCaller := mload(add(usdcMessage, 140)) // 32 + 108 = 140
      minFinalityThreshold := mload(add(usdcMessage, 144)) // 140 + 4 = 144
      finalityThresholdExecuted := mload(add(usdcMessage, 148)) // 144 + 4 = 148

      // The message body starts at index 148 and because it is a dynamic byte array, contains a 32-byte
      // length field prefixing the data.
      burnToken := mload(add(usdcMessage, 184)) // 148 + 32 + 4 = 184
      mintRecipient := mload(add(usdcMessage, 216)) // 148 + 32 + 36 = 216
      amount := mload(add(usdcMessage, 248)) // 148 + 32 + 68 = 248
    }

    // Calculate the invalid deposit hash based on the message contents
    bytes32 invalidDepositHash = USDCSourcePoolDataCodec._calculateDepositHash(
      messageSourceDomain,
      amount,
      destinationDomain,
      mintRecipient,
      burnToken,
      destinationCaller,
      s_usdcTokenPool.MAX_FEE(),
      minFinalityThreshold
    );

    // Expect the revert with the invalid deposit hash because the message data does not match
    // the deposit hash provided by the source pool
    vm.expectRevert(
      abi.encodeWithSelector(USDCTokenPoolCCTPV2.InvalidDepositHash.selector, bytes32(0), invalidDepositHash)
    );

    s_usdcTokenPool.validateMessage(usdcMessage, s_validSourceTokenData);
  }
}
