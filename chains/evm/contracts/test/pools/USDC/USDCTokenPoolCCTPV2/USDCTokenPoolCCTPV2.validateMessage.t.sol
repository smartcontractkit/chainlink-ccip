// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {USDCTokenPool} from "../../../../pools/USDC/USDCTokenPool.sol";
import {USDCTokenPoolCCTPV2} from "../../../../pools/USDC/USDCTokenPoolCCTPV2.sol";
import {USDCTokenPoolCCTPV2Setup} from "./USDCTokenPoolCCTPV2Setup.t.sol";

contract USDCTokenPoolCCTPV2_validateMessage is USDCTokenPoolCCTPV2Setup {
  USDCMessageCCTPV2 internal s_validUsdcMessage;
  USDCTokenPool.SourceTokenDataPayload internal s_validSourceTokenData;

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
      messageBody: bytes("")
    });

    // Create valid source token data payload that matches the valid USDC message
    // This is used by tests that need to test payload validation failures
    s_validSourceTokenData = USDCTokenPool.SourceTokenDataPayload({
      nonce: 0,
      sourceDomain: s_validUsdcMessage.sourceDomain,
      cctpVersion: USDCTokenPool.CCTPVersion.CCTP_V2,
      amount: 0,
      destinationDomain: DEST_DOMAIN_IDENTIFIER,
      mintRecipient: bytes32(0),
      burnToken: address(0),
      destinationCaller: bytes32(0),
      maxFee: 0,
      minFinalityThreshold: 0
    });
  }

  function testFuzz_validateMessage_Success(uint32 sourceDomain, bytes32 nonce) public {
    vm.pauseGasMetering();
    USDCMessageCCTPV2 memory usdcMessage = USDCMessageCCTPV2({
      version: 1,
      sourceDomain: sourceDomain,
      destinationDomain: DEST_DOMAIN_IDENTIFIER,
      nonce: nonce,
      sender: SOURCE_CHAIN_TOKEN_SENDER,
      recipient: bytes32(uint256(299999)),
      destinationCaller: bytes32(uint256(uint160(address(s_usdcTokenPool)))),
      minFinalityThreshold: s_usdcTokenPool.FINALITY_THRESHOLD(),
      finalityThresholdExecuted: s_usdcTokenPool.FINALITY_THRESHOLD(),
      messageBody: bytes("")
    });
    bytes memory encodedUsdcMessage = _generateUSDCMessageCCTPV2(usdcMessage);

    vm.resumeGasMetering();
    s_usdcTokenPool.validateMessage(
      encodedUsdcMessage,
      USDCTokenPool.SourceTokenDataPayload({
        nonce: 0,
        sourceDomain: sourceDomain,
        cctpVersion: USDCTokenPool.CCTPVersion.CCTP_V2,
        amount: 0,
        destinationDomain: DEST_DOMAIN_IDENTIFIER,
        mintRecipient: bytes32(0),
        burnToken: address(0),
        destinationCaller: bytes32(0),
        maxFee: 0,
        minFinalityThreshold: 0
      })
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
      USDCTokenPool.SourceTokenDataPayload({
        nonce: 0,
        sourceDomain: expectedSourceDomain,
        cctpVersion: USDCTokenPool.CCTPVersion.CCTP_V2,
        amount: 0,
        destinationDomain: DEST_DOMAIN_IDENTIFIER,
        mintRecipient: bytes32(0),
        burnToken: address(0),
        destinationCaller: bytes32(0),
        maxFee: 0,
        minFinalityThreshold: 0
      })
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

  function test_validateMessage_RevertWhen_InvalidCCTPVersion() public {
    // Create source token data with invalid CCTP version
    USDCTokenPool.SourceTokenDataPayload memory invalidCCTPVersionData = USDCTokenPool.SourceTokenDataPayload({
      nonce: 0,
      sourceDomain: s_validUsdcMessage.sourceDomain,
      cctpVersion: USDCTokenPool.CCTPVersion.CCTP_V1,
      amount: 0,
      destinationDomain: DEST_DOMAIN_IDENTIFIER,
      mintRecipient: bytes32(0),
      burnToken: address(0),
      destinationCaller: bytes32(0),
      maxFee: 0,
      minFinalityThreshold: 0
    });

    vm.expectRevert(
      abi.encodeWithSelector(
        USDCTokenPool.InvalidCCTPVersion.selector, USDCTokenPool.CCTPVersion.CCTP_V2, USDCTokenPool.CCTPVersion.CCTP_V1
      )
    );
    s_usdcTokenPool.validateMessage(_generateUSDCMessageCCTPV2(s_validUsdcMessage), invalidCCTPVersionData);
  }
}
