// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {USDCTokenPoolCCTPV2} from "../../../../../pools/USDC/CCTPV2/USDCTokenPoolCCTPV2.sol";
import {USDCTokenPool} from "../../../../../pools/USDC/USDCTokenPool.sol";
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
      minFinalityThreshold: 2000,
      finalityThresholdExecuted: 2000,
      messageBody: bytes("")
    });

    bytes memory encodedUsdcMessage = _generateUSDCMessageCCTPV2(usdcMessage);

    USDCTokenPool.SourceTokenDataPayload memory sourceTokenDataPayload = USDCTokenPool.SourceTokenDataPayload({
      nonce: uint64(0),
      sourceDomain: uint32(usdcMessage.sourceDomain),
      cctpVersion: USDCTokenPool.CCTPVersion.VERSION_2
    });

    vm.resumeGasMetering();

    s_usdcTokenPool.validateMessage(encodedUsdcMessage, sourceTokenDataPayload);
  }

  // Reverts

  function test_RevertWhen_ValidateInvalidMessage() public {
    USDCMessageCCTPV2 memory usdcMessage = USDCMessageCCTPV2({
      version: 1,
      sourceDomain: 1553252,
      destinationDomain: DEST_DOMAIN_IDENTIFIER,
      nonce: keccak256("0xCLL"),
      sender: SOURCE_CHAIN_TOKEN_SENDER,
      recipient: bytes32(uint256(92398429395823)),
      destinationCaller: bytes32(uint256(uint160(address(s_usdcTokenPool)))),
      minFinalityThreshold: 2000,
      finalityThresholdExecuted: 2000,
      messageBody: bytes("")
    });

    bytes memory encodedUsdcMessage = _generateUSDCMessageCCTPV2(usdcMessage);

    USDCTokenPool.SourceTokenDataPayload memory sourceTokenDataPayload = USDCTokenPool.SourceTokenDataPayload({
      nonce: uint64(0),
      sourceDomain: usdcMessage.sourceDomain + 1,
      cctpVersion: USDCTokenPool.CCTPVersion.VERSION_2
    });

    // The usdcMessage should have the source domain present in sourceTokenData payload but it doesn't so revert
    vm.expectRevert(
      abi.encodeWithSelector(
        USDCTokenPool.InvalidSourceDomain.selector, usdcMessage.sourceDomain + 1, usdcMessage.sourceDomain
      )
    );
    s_usdcTokenPool.validateMessage(encodedUsdcMessage, sourceTokenDataPayload);

    // Since we had the sourceDomain be incorrect in the previous call, we fix it here before proceeding so that the
    // correct revert error is invoked.
    sourceTokenDataPayload.sourceDomain = usdcMessage.sourceDomain;

    usdcMessage.destinationDomain = DEST_DOMAIN_IDENTIFIER + 1;
    vm.expectRevert(
      abi.encodeWithSelector(
        USDCTokenPool.InvalidDestinationDomain.selector, DEST_DOMAIN_IDENTIFIER, usdcMessage.destinationDomain
      )
    );

    s_usdcTokenPool.validateMessage(_generateUSDCMessageCCTPV2(usdcMessage), sourceTokenDataPayload);
    usdcMessage.destinationDomain = DEST_DOMAIN_IDENTIFIER;

    uint32 wrongVersion = usdcMessage.version + 1;

    usdcMessage.version = wrongVersion;
    encodedUsdcMessage = _generateUSDCMessageCCTPV2(usdcMessage);

    vm.expectRevert(abi.encodeWithSelector(USDCTokenPool.InvalidMessageVersion.selector, wrongVersion));
    s_usdcTokenPool.validateMessage(encodedUsdcMessage, sourceTokenDataPayload);
    usdcMessage.version = 1;

    // Change Finality threshold and finalityThresholdExecuted to 1000 to intentionally revert
    usdcMessage.minFinalityThreshold = 1000;
    encodedUsdcMessage = _generateUSDCMessageCCTPV2(usdcMessage);

    vm.expectRevert(abi.encodeWithSelector(USDCTokenPoolCCTPV2.InvalidMinFinalityThreshold.selector, 2000, 1000));

    s_usdcTokenPool.validateMessage(encodedUsdcMessage, sourceTokenDataPayload);

    // Change the min threshold back to 2k and the finality threshold to 1k to trigger
    // the other short half of the short-circuit.
    usdcMessage.minFinalityThreshold = 2000;
    usdcMessage.finalityThresholdExecuted = 1000;
    encodedUsdcMessage = _generateUSDCMessageCCTPV2(usdcMessage);

    vm.expectRevert(abi.encodeWithSelector(USDCTokenPoolCCTPV2.InvalidExecutionFinalityThreshold.selector, 2000, 1000));
    s_usdcTokenPool.validateMessage(encodedUsdcMessage, sourceTokenDataPayload);
  }
}
