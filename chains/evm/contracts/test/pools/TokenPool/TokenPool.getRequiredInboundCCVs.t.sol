// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {TokenPool} from "../../../pools/TokenPool.sol";
import {TokenPoolV2Setup} from "./TokenPoolV2Setup.t.sol";

contract TokenPoolV2_getRequiredInboundCCVs is TokenPoolV2Setup {
  function test_getRequiredInboundCCVs_BaseCCVs() public {
    address[] memory inboundCCVs = new address[](1);
    inboundCCVs[0] = makeAddr("inboundCCV1");

    TokenPool.CCVConfigArg[] memory configArgs = new TokenPool.CCVConfigArg[](1);
    configArgs[0] = TokenPool.CCVConfigArg({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      outboundCCVs: new address[](0),
      additionalOutboundCCVs: new address[](0),
      inboundCCVs: inboundCCVs,
      additionalInboundCCVs: new address[](0)
    });

    s_tokenPool.applyCCVConfigUpdates(configArgs);

    // Test with amount below threshold, should return only base CCVs.
    address[] memory storedInbound =
      s_tokenPool.getRequiredInboundCCVs(address(s_token), DEST_CHAIN_SELECTOR, 100, 0, "");

    assertEq(storedInbound.length, inboundCCVs.length);
    assertEq(storedInbound[0], inboundCCVs[0]);
  }

  function test_getRequiredInboundCCVs_WithAdditionalCCVs() public {
    address[] memory inboundCCVs = new address[](1);
    inboundCCVs[0] = makeAddr("inboundCCV1");

    address[] memory additionalInboundCCVs = new address[](1);
    additionalInboundCCVs[0] = makeAddr("additionalInboundCCV1");

    TokenPool.CCVConfigArg[] memory configArgs = new TokenPool.CCVConfigArg[](1);
    configArgs[0] = TokenPool.CCVConfigArg({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      outboundCCVs: new address[](0),
      additionalOutboundCCVs: new address[](0),
      inboundCCVs: inboundCCVs,
      additionalInboundCCVs: additionalInboundCCVs
    });

    s_tokenPool.applyCCVConfigUpdates(configArgs);

    // Set threshold amount.
    uint96 thresholdAmount = 1000;
    s_tokenPool.setDynamicConfig(address(s_sourceRouter), thresholdAmount);

    // Test with amount below threshold, should return only base CCVs.
    address[] memory storedInboundBelow =
      s_tokenPool.getRequiredInboundCCVs(address(s_token), DEST_CHAIN_SELECTOR, thresholdAmount - 500, 0, "");
    assertEq(storedInboundBelow.length, 1);
    assertEq(storedInboundBelow[0], inboundCCVs[0]);

    // Test with amount above threshold, should return base + additional CCVs.
    address[] memory storedInboundAbove =
      s_tokenPool.getRequiredInboundCCVs(address(s_token), DEST_CHAIN_SELECTOR, thresholdAmount + 500, 0, "");
    assertEq(storedInboundAbove.length, 2);
    assertEq(storedInboundAbove[0], inboundCCVs[0]);
    assertEq(storedInboundAbove[1], additionalInboundCCVs[0]);
  }

  function test_getRequiredInboundCCVs_NoAdditionalCCVs() public {
    address[] memory inboundCCVs = new address[](1);
    inboundCCVs[0] = makeAddr("inboundCCV1");

    TokenPool.CCVConfigArg[] memory configArgs = new TokenPool.CCVConfigArg[](1);
    configArgs[0] = TokenPool.CCVConfigArg({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      outboundCCVs: new address[](0),
      additionalOutboundCCVs: new address[](0),
      inboundCCVs: inboundCCVs,
      additionalInboundCCVs: new address[](0)
    });

    s_tokenPool.applyCCVConfigUpdates(configArgs);

    // Set threshold amount.
    uint96 thresholdAmount = 1000;
    s_tokenPool.setDynamicConfig(address(s_sourceRouter), thresholdAmount);

    // Test with amount above threshold but no additional CCVs, should return only base CCVs.
    address[] memory storedInbound =
      s_tokenPool.getRequiredInboundCCVs(address(s_token), DEST_CHAIN_SELECTOR, thresholdAmount + 500, 0, "");
    assertEq(storedInbound.length, 1);
    assertEq(storedInbound[0], inboundCCVs[0]);
  }
}
