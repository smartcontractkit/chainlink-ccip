// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {TokenPool} from "../../../poolsV2/TokenPool.sol";
import {TokenPoolV2Setup} from "./TokenPoolV2Setup.t.sol";

contract TokenPoolV2_getRequiredInboundCCVs is TokenPoolV2Setup {
  function test_getRequiredInboundCCVs() public {
    address[] memory inboundCCVs = new address[](1);
    inboundCCVs[0] = makeAddr("inboundCCV1");

    TokenPool.CCVConfigArg[] memory configArgs = new TokenPool.CCVConfigArg[](1);
    configArgs[0] = TokenPool.CCVConfigArg({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      outboundCCVs: new address[](0),
      inboundCCVs: inboundCCVs
    });

    s_tokenPool.applyCCVConfigUpdates(configArgs);

    // Verify the configuration was stored correctly.
    address[] memory storedInbound = s_tokenPool.getRequiredInboundCCVs(address(s_token), DEST_CHAIN_SELECTOR, 0, 0, "");

    assertEq(storedInbound.length, inboundCCVs.length);
    assertEq(storedInbound[0], inboundCCVs[0]);
  }
}
