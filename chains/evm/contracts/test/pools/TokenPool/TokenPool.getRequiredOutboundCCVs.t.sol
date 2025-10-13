// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {TokenPool} from "../../../pools/TokenPool.sol";
import {TokenPoolV2Setup} from "./TokenPoolV2Setup.t.sol";

contract TokenPoolV2_getRequiredOutboundCCVs is TokenPoolV2Setup {
  function test_getRequiredOutboundCCVs_BaseCCVs() public {
    address[] memory outboundCCVs = new address[](1);
    outboundCCVs[0] = makeAddr("outboundCCV1");

    TokenPool.CCVConfigArg[] memory configArgs = new TokenPool.CCVConfigArg[](1);
    configArgs[0] = TokenPool.CCVConfigArg({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      outboundCCVs: outboundCCVs,
      additionalOutboundCCVs: new address[](0),
      inboundCCVs: new address[](0),
      additionalInboundCCVs: new address[](0)
    });

    s_tokenPool.applyCCVConfigUpdates(configArgs);

    // Test with amount below threshold, should return only base CCVs.
    address[] memory storedOutbound =
      s_tokenPool.getRequiredOutboundCCVs(address(s_token), DEST_CHAIN_SELECTOR, 100, 0, "");

    assertEq(storedOutbound.length, outboundCCVs.length);
    assertEq(storedOutbound[0], outboundCCVs[0]);
  }

  function test_getRequiredOutboundCCVs_WithAdditionalCCVs() public {
    address[] memory outboundCCVs = new address[](1);
    outboundCCVs[0] = makeAddr("outboundCCV1");

    address[] memory additionalOutboundCCVs = new address[](1);
    additionalOutboundCCVs[0] = makeAddr("additionalOutboundCCV1");

    TokenPool.CCVConfigArg[] memory configArgs = new TokenPool.CCVConfigArg[](1);
    configArgs[0] = TokenPool.CCVConfigArg({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      outboundCCVs: outboundCCVs,
      additionalOutboundCCVs: additionalOutboundCCVs,
      inboundCCVs: new address[](0),
      additionalInboundCCVs: new address[](0)
    });

    s_tokenPool.applyCCVConfigUpdates(configArgs);

    // Set threshold amount.
    uint256 thresholdAmount = 1000;
    s_tokenPool.setDynamicConfig(address(s_sourceRouter), thresholdAmount);

    // Test with amount below threshold, should return only base CCVs.
    address[] memory storedOutboundBelow =
      s_tokenPool.getRequiredOutboundCCVs(address(s_token), DEST_CHAIN_SELECTOR, thresholdAmount - 500, 0, "");
    assertEq(storedOutboundBelow.length, 1);
    assertEq(storedOutboundBelow[0], outboundCCVs[0]);

    // Test with amount above threshold, should return base + additional CCVs.
    address[] memory storedOutboundAbove =
      s_tokenPool.getRequiredOutboundCCVs(address(s_token), DEST_CHAIN_SELECTOR, thresholdAmount + 500, 0, "");
    assertEq(storedOutboundAbove.length, 2);
    assertEq(storedOutboundAbove[0], outboundCCVs[0]);
    assertEq(storedOutboundAbove[1], additionalOutboundCCVs[0]);
  }

  function test_getRequiredOutboundCCVs_NoAdditionalCCVs() public {
    address[] memory outboundCCVs = new address[](1);
    outboundCCVs[0] = makeAddr("outboundCCV1");

    TokenPool.CCVConfigArg[] memory configArgs = new TokenPool.CCVConfigArg[](1);
    configArgs[0] = TokenPool.CCVConfigArg({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      outboundCCVs: outboundCCVs,
      additionalOutboundCCVs: new address[](0),
      inboundCCVs: new address[](0),
      additionalInboundCCVs: new address[](0)
    });

    s_tokenPool.applyCCVConfigUpdates(configArgs);

    // Set threshold amount.
    uint256 thresholdAmount = 1000;
    s_tokenPool.setDynamicConfig(address(s_sourceRouter), thresholdAmount);

    // Test with amount above threshold but no additional CCVs, should return only base CCVs.
    address[] memory storedOutbound =
      s_tokenPool.getRequiredOutboundCCVs(address(s_token), DEST_CHAIN_SELECTOR, thresholdAmount + 500, 0, "");
    assertEq(storedOutbound.length, 1);
    assertEq(storedOutbound[0], outboundCCVs[0]);
  }
}
