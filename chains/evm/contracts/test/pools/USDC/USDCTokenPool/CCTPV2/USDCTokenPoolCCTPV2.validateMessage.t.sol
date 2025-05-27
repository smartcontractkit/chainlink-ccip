// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {USDCTokenPoolCCTPV2} from "../../../../../pools/USDC/CCTPV2/USDCTokenPoolCCTPV2.sol";
import {USDCTokenPool} from "../../../../../pools/USDC/USDCTokenPool.sol";
import {USDCTokenPoolCCTPV2Setup} from "./USDCTokenPoolCCTPV2Setup.t.sol";

contract USDCTokenPoolCCTPV2__validateMessage is USDCTokenPoolCCTPV2Setup {
  function testFuzz_ValidateMessage_Success(uint32 sourceDomain, bytes32 nonce) public {
    vm.pauseGasMetering();
    USDCMessageCCTPV2 memory usdcMessage = USDCMessageCCTPV2({
      version: 0,
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

    vm.resumeGasMetering();
    s_usdcTokenPool.validateMessage(encodedUsdcMessage, sourceDomain);
  }

  // Reverts

  function test_RevertWhen_ValidateInvalidMessage() public {
    USDCMessageCCTPV2 memory usdcMessage = USDCMessageCCTPV2({
      version: 0,
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

    s_usdcTokenPool.validateMessage(encodedUsdcMessage, usdcMessage.sourceDomain);

    uint32 expectedSourceDomain = usdcMessage.sourceDomain + 1;

    vm.expectRevert(
      abi.encodeWithSelector(USDCTokenPool.InvalidSourceDomain.selector, expectedSourceDomain, usdcMessage.sourceDomain)
    );
    s_usdcTokenPool.validateMessage(encodedUsdcMessage, expectedSourceDomain);

    usdcMessage.destinationDomain = DEST_DOMAIN_IDENTIFIER + 1;
    vm.expectRevert(
      abi.encodeWithSelector(
        USDCTokenPool.InvalidDestinationDomain.selector, DEST_DOMAIN_IDENTIFIER, usdcMessage.destinationDomain
      )
    );

    s_usdcTokenPool.validateMessage(_generateUSDCMessageCCTPV2(usdcMessage), usdcMessage.sourceDomain);
    usdcMessage.destinationDomain = DEST_DOMAIN_IDENTIFIER;

    uint32 wrongVersion = usdcMessage.version + 1;

    usdcMessage.version = wrongVersion;
    encodedUsdcMessage = _generateUSDCMessageCCTPV2(usdcMessage);

    vm.expectRevert(abi.encodeWithSelector(USDCTokenPool.InvalidMessageVersion.selector, wrongVersion));
    s_usdcTokenPool.validateMessage(encodedUsdcMessage, usdcMessage.sourceDomain);
    usdcMessage.version = 0;

    // Change Finality threshold and finalityThresholdExecuted to 1000 to intentionally revert
    usdcMessage.minFinalityThreshold = 1000;
    encodedUsdcMessage = _generateUSDCMessageCCTPV2(usdcMessage);

    vm.expectRevert(abi.encodeWithSelector(USDCTokenPoolCCTPV2.InvalidMinFinalityThreshold.selector, 2000, 1000));

    s_usdcTokenPool.validateMessage(encodedUsdcMessage, usdcMessage.sourceDomain);

    // Change the min threshold back to 2k and the finality threshold to 1k to trigger
    // the other short half of the short-circuit.
    usdcMessage.minFinalityThreshold = 2000;
    usdcMessage.finalityThresholdExecuted = 1000;
    encodedUsdcMessage = _generateUSDCMessageCCTPV2(usdcMessage);

    vm.expectRevert(abi.encodeWithSelector(USDCTokenPoolCCTPV2.InvalidExecutionFinalityThreshold.selector, 2000, 1000));
    s_usdcTokenPool.validateMessage(encodedUsdcMessage, usdcMessage.sourceDomain);
  }
}
