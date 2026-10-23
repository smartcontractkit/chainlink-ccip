// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolV2} from "../../../interfaces/IPoolV2.sol";

import {MultiTokenPoolSetup} from "./MultiTokenPoolSetup.t.sol";

contract MultiTokenPool_getRequiredCCVs is MultiTokenPoolSetup {
  function test_getRequiredCCVs() public {
    address[] memory outboundCCVs = new address[](2);
    outboundCCVs[0] = makeAddr("outboundCCV1");
    outboundCCVs[1] = makeAddr("outboundCCV2");
    address[] memory inboundCCVs = new address[](1);
    inboundCCVs[0] = makeAddr("inboundCCV");
    s_poolConfigManager.setRequiredCCVs(outboundCCVs, inboundCCVs);

    assertEq(
      s_pool.getRequiredCCVs(
        address(s_token), DEST_CHAIN_SELECTOR, 1e18, bytes4(0), "", IPoolV2.MessageDirection.Outbound
      ),
      outboundCCVs
    );
    assertEq(
      s_pool.getRequiredCCVs(
        address(s_token), DEST_CHAIN_SELECTOR, 1e18, bytes4(0), "", IPoolV2.MessageDirection.Inbound
      ),
      inboundCCVs
    );
  }
}
