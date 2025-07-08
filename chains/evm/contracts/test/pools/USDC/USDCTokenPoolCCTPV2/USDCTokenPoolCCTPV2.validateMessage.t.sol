// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {USDCTokenPool} from "../../../../pools/USDC/USDCTokenPool.sol";
import {USDCTokenPoolCCTPV2} from "../../../../pools/USDC/USDCTokenPoolCCTPV2.sol";
import {USDCTokenPoolCCTPV2Setup} from "./USDCTokenPoolCCTPV2Setup.t.sol";

contract USDCTokenPoolCCTPV2__validateMessage is USDCTokenPoolCCTPV2Setup {
  function testFuzz_ValidateMessage_Success(uint32 sourceDomain, bytes32 nonce) public {
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

  function test_RevertWhen_ValidateInvalidMessage() public {
    USDCMessageCCTPV2 memory usdcMessage = USDCMessageCCTPV2({
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

    USDCTokenPool.SourceTokenDataPayload memory sourceTokenData = USDCTokenPool.SourceTokenDataPayload({
      nonce: 0,
      sourceDomain: usdcMessage.sourceDomain,
      cctpVersion: USDCTokenPool.CCTPVersion.CCTP_V2,
      amount: 0,
      destinationDomain: DEST_DOMAIN_IDENTIFIER,
      mintRecipient: bytes32(0),
      burnToken: address(0),
      destinationCaller: bytes32(0),
      maxFee: 0,
      minFinalityThreshold: 0
    });

    bytes memory encodedUsdcMessage = _generateUSDCMessageCCTPV2(usdcMessage);

    s_usdcTokenPool.validateMessage(encodedUsdcMessage, sourceTokenData);

    uint32 expectedSourceDomain = usdcMessage.sourceDomain + 1;

    vm.expectRevert(
      abi.encodeWithSelector(USDCTokenPool.InvalidSourceDomain.selector, expectedSourceDomain, usdcMessage.sourceDomain)
    );
    s_usdcTokenPool.validateMessage(
      encodedUsdcMessage,
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

    usdcMessage.destinationDomain = DEST_DOMAIN_IDENTIFIER + 1;
    vm.expectRevert(
      abi.encodeWithSelector(
        USDCTokenPool.InvalidDestinationDomain.selector, DEST_DOMAIN_IDENTIFIER, usdcMessage.destinationDomain
      )
    );

    s_usdcTokenPool.validateMessage(
      _generateUSDCMessageCCTPV2(usdcMessage),
      USDCTokenPool.SourceTokenDataPayload({
        nonce: 0,
        sourceDomain: usdcMessage.sourceDomain,
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
    usdcMessage.destinationDomain = DEST_DOMAIN_IDENTIFIER;

    uint32 wrongVersion = usdcMessage.version + 1;

    usdcMessage.version = wrongVersion;
    encodedUsdcMessage = _generateUSDCMessageCCTPV2(usdcMessage);

    vm.expectRevert(abi.encodeWithSelector(USDCTokenPool.InvalidMessageVersion.selector, wrongVersion));
    s_usdcTokenPool.validateMessage(encodedUsdcMessage, sourceTokenData);

    // Create a byte string of length less than 116 (e.g., 100)
    bytes memory shortMessage = new bytes(100);
    vm.expectRevert(abi.encodeWithSelector(USDCTokenPool.InvalidMessageLength.selector, 100));
    s_usdcTokenPool.validateMessage(shortMessage, sourceTokenData);

    // Test for InvalidMinFinalityThreshold
    usdcMessage.version = 1;
    usdcMessage.destinationDomain = DEST_DOMAIN_IDENTIFIER;
    usdcMessage.minFinalityThreshold = s_usdcTokenPool.FINALITY_THRESHOLD() + 1;
    bytes memory invalidMinFinalityMsg = _generateUSDCMessageCCTPV2(usdcMessage);
    vm.expectRevert(
      abi.encodeWithSelector(
        USDCTokenPoolCCTPV2.InvalidMinFinalityThreshold.selector,
        s_usdcTokenPool.FINALITY_THRESHOLD(),
        usdcMessage.minFinalityThreshold
      )
    );
    s_usdcTokenPool.validateMessage(invalidMinFinalityMsg, sourceTokenData);

    // Test for InvalidExecutionFinalityThreshold
    usdcMessage.minFinalityThreshold = s_usdcTokenPool.FINALITY_THRESHOLD();
    usdcMessage.finalityThresholdExecuted = s_usdcTokenPool.FINALITY_THRESHOLD() + 1;
    bytes memory invalidExecFinalityMsg = _generateUSDCMessageCCTPV2(usdcMessage);
    vm.expectRevert(
      abi.encodeWithSelector(
        USDCTokenPoolCCTPV2.InvalidExecutionFinalityThreshold.selector,
        s_usdcTokenPool.FINALITY_THRESHOLD(),
        usdcMessage.finalityThresholdExecuted
      )
    );
    s_usdcTokenPool.validateMessage(invalidExecFinalityMsg, sourceTokenData);

    // Test for InvalidCCTPVersion
    usdcMessage.minFinalityThreshold = s_usdcTokenPool.FINALITY_THRESHOLD();
    usdcMessage.finalityThresholdExecuted = s_usdcTokenPool.FINALITY_THRESHOLD();
    bytes memory validMessage = _generateUSDCMessageCCTPV2(usdcMessage);
    USDCTokenPool.SourceTokenDataPayload memory invalidCCTPVersionData = USDCTokenPool.SourceTokenDataPayload({
      nonce: 0,
      sourceDomain: usdcMessage.sourceDomain,
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
    s_usdcTokenPool.validateMessage(validMessage, invalidCCTPVersionData);
  }
}
