// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {USDCTokenPool} from "../../../../pools/USDC/USDCTokenPool.sol";
import {USDCTokenPoolSetup} from "./USDCTokenPoolSetup.t.sol";

contract USDCTokenPool_validateMessage is USDCTokenPoolSetup {
  function testFuzz_validateMessage_Success(uint32 sourceDomain, uint64 nonce) public {
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

  function test_validateMessage_RevertWhen_InvalidSourceDomain() public {
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

    uint32 expectedSourceDomain = usdcMessage.sourceDomain + 1;

    vm.expectRevert(
      abi.encodeWithSelector(USDCTokenPool.InvalidSourceDomain.selector, expectedSourceDomain, usdcMessage.sourceDomain)
    );
    s_usdcTokenPool.validateMessage(
      _generateUSDCMessage(usdcMessage),
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
  }

  function test_validateMessage_RevertWhen_InvalidNonce() public {
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

    uint64 expectedNonce = usdcMessage.nonce + 1;

    vm.expectRevert(abi.encodeWithSelector(USDCTokenPool.InvalidNonce.selector, expectedNonce, usdcMessage.nonce));
    s_usdcTokenPool.validateMessage(
      _generateUSDCMessage(usdcMessage),
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
  }

  function test_validateMessage_RevertWhen_InvalidDestinationDomain() public {
    USDCMessage memory usdcMessage = USDCMessage({
      version: 0,
      sourceDomain: 1553252,
      destinationDomain: DEST_DOMAIN_IDENTIFIER + 1,
      nonce: 387289284924,
      sender: SOURCE_CHAIN_TOKEN_SENDER,
      recipient: bytes32(uint256(92398429395823)),
      destinationCaller: bytes32(uint256(uint160(address(s_usdcTokenPool)))),
      messageBody: bytes("")
    });

    vm.expectRevert(
      abi.encodeWithSelector(
        USDCTokenPool.InvalidDestinationDomain.selector, DEST_DOMAIN_IDENTIFIER, DEST_DOMAIN_IDENTIFIER + 1
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
  }

  function test_validateMessage_RevertWhen_InvalidMessageVersion() public {
    USDCMessage memory usdcMessage = USDCMessage({
      version: 1,
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

    // The right version is 0, so the invalid version is 1, since CCTP V1 uses a version number of 0
    vm.expectRevert(abi.encodeWithSelector(USDCTokenPool.InvalidMessageVersion.selector, 0, 1));
    s_usdcTokenPool.validateMessage(encodedUsdcMessage, sourceTokenData);
  }

  function test_validateMessage_RevertWhen_InvalidMessageLength() public {
    USDCTokenPool.SourceTokenDataPayload memory sourceTokenData = USDCTokenPool.SourceTokenDataPayload({
      nonce: 387289284924,
      sourceDomain: 1553252,
      cctpVersion: USDCTokenPool.CCTPVersion.CCTP_V1,
      amount: 0,
      destinationDomain: DEST_DOMAIN_IDENTIFIER,
      mintRecipient: bytes32(0),
      burnToken: address(0),
      destinationCaller: bytes32(0),
      maxFee: 0,
      minFinalityThreshold: 0
    });

    bytes memory shortMessage = new bytes(100);
    vm.expectRevert(abi.encodeWithSelector(USDCTokenPool.InvalidMessageLength.selector, 100));
    s_usdcTokenPool.validateMessage(shortMessage, sourceTokenData);
  }

  function test_validateMessage_RevertWhen_InvalidCCTPVersion() public {
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

    bytes memory encodedUsdcMessage = _generateUSDCMessage(usdcMessage);

    vm.expectRevert(
      abi.encodeWithSelector(
        USDCTokenPool.InvalidCCTPVersion.selector, USDCTokenPool.CCTPVersion.CCTP_V1, USDCTokenPool.CCTPVersion.CCTP_V2
      )
    );

    s_usdcTokenPool.validateMessage(
      encodedUsdcMessage,
      USDCTokenPool.SourceTokenDataPayload({
        nonce: usdcMessage.nonce,
        sourceDomain: usdcMessage.sourceDomain + 1, // Use different source domain to avoid other validation errors
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
