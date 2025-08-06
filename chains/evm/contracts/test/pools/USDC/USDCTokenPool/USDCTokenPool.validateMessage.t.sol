// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {USDCTokenPool} from "../../../../pools/USDC/USDCTokenPool.sol";
import {USDCTokenPoolSetup} from "./USDCTokenPoolSetup.t.sol";

contract USDCTokenPool__validateMessage is USDCTokenPoolSetup {
  function testFuzz_ValidateMessage_Success(uint32 sourceDomain, uint64 nonce) public {
    vm.pauseGasMetering();
    USDCMessage memory usdcMessage = USDCMessage({
      version: 0,
      sourceDomain: sourceDomain,
      destinationDomain: DEST_DOMAIN_IDENTIFIER,
      nonce: nonce,
      sender: SOURCE_CHAIN_TOKEN_SENDER,
      recipient: bytes32(uint256(299999)),
      destinationCaller: bytes32(uint256(uint160(address(s_usdcTokenPool)))),
      messageBody: bytes("")
    });

    bytes memory encodedUsdcMessage = _generateUSDCMessage(usdcMessage);

    vm.resumeGasMetering();
    s_usdcTokenPool.validateMessage(
      encodedUsdcMessage,
      USDCTokenPool.SourceTokenDataPayload({
        nonce: nonce,
        sourceDomain: sourceDomain,
        cctpVersion: USDCTokenPool.CCTPVersion.CCTP_V1,
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
    USDCMessage memory usdcMessage = USDCMessage({
      version: 0,
      sourceDomain: 1553252,
      destinationDomain: DEST_DOMAIN_IDENTIFIER,
      nonce: 387289284924,
      sender: SOURCE_CHAIN_TOKEN_SENDER,
      recipient: bytes32(uint256(92398429395823)),
      destinationCaller: bytes32(uint256(uint160(address(s_usdcTokenPool)))),
      messageBody: bytes("")
    });

    USDCTokenPool.SourceTokenDataPayload memory sourceTokenData = USDCTokenPool.SourceTokenDataPayload({
      nonce: usdcMessage.nonce,
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

    bytes memory encodedUsdcMessage = _generateUSDCMessage(usdcMessage);

    s_usdcTokenPool.validateMessage(encodedUsdcMessage, sourceTokenData);

    uint32 expectedSourceDomain = usdcMessage.sourceDomain + 1;

    vm.expectRevert(
      abi.encodeWithSelector(USDCTokenPool.InvalidSourceDomain.selector, expectedSourceDomain, usdcMessage.sourceDomain)
    );
    s_usdcTokenPool.validateMessage(
      encodedUsdcMessage,
      USDCTokenPool.SourceTokenDataPayload({
        nonce: usdcMessage.nonce,
        sourceDomain: expectedSourceDomain,
        cctpVersion: USDCTokenPool.CCTPVersion.CCTP_V1,
        amount: 0,
        destinationDomain: DEST_DOMAIN_IDENTIFIER,
        mintRecipient: bytes32(0),
        burnToken: address(0),
        destinationCaller: bytes32(0),
        maxFee: 0,
        minFinalityThreshold: 0
      })
    );

    uint64 expectedNonce = usdcMessage.nonce + 1;

    vm.expectRevert(abi.encodeWithSelector(USDCTokenPool.InvalidNonce.selector, expectedNonce, usdcMessage.nonce));
    s_usdcTokenPool.validateMessage(
      encodedUsdcMessage,
      USDCTokenPool.SourceTokenDataPayload({
        nonce: expectedNonce,
        sourceDomain: usdcMessage.sourceDomain,
        cctpVersion: USDCTokenPool.CCTPVersion.CCTP_V1,
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
      _generateUSDCMessage(usdcMessage),
      USDCTokenPool.SourceTokenDataPayload({
        nonce: usdcMessage.nonce,
        sourceDomain: usdcMessage.sourceDomain,
        cctpVersion: USDCTokenPool.CCTPVersion.CCTP_V1,
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
    encodedUsdcMessage = _generateUSDCMessage(usdcMessage);

    vm.expectRevert(abi.encodeWithSelector(USDCTokenPool.InvalidMessageVersion.selector, wrongVersion, 0));
    s_usdcTokenPool.validateMessage(encodedUsdcMessage, sourceTokenData);

    // Create a byte string of length less than 116 (e.g., 100)
    bytes memory shortMessage = new bytes(100);
    vm.expectRevert(abi.encodeWithSelector(USDCTokenPool.InvalidMessageLength.selector, 100));
    s_usdcTokenPool.validateMessage(shortMessage, sourceTokenData);

    // Undo the wrong message version and re-encode the message with the correct version
    usdcMessage.version--;
    encodedUsdcMessage = _generateUSDCMessage(usdcMessage);

    vm.expectRevert(
      abi.encodeWithSelector(
        USDCTokenPool.InvalidCCTPVersion.selector, USDCTokenPool.CCTPVersion.CCTP_V1, USDCTokenPool.CCTPVersion.CCTP_V2
      )
    );
    s_usdcTokenPool.validateMessage(
      encodedUsdcMessage,
      USDCTokenPool.SourceTokenDataPayload({
        nonce: usdcMessage.nonce,
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
}