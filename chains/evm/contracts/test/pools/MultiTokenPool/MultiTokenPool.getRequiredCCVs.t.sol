// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../../interfaces/IPoolConfigManager.sol";
import {IPoolV2} from "../../../interfaces/IPoolV2.sol";

import {FinalityCodec} from "../../../libraries/FinalityCodec.sol";
import {MultiTokenPoolSetup} from "./MultiTokenPoolSetup.t.sol";

contract MultiTokenPool_getRequiredCCVs is MultiTokenPoolSetup {
  function test_getRequiredCCVs() public {
    address[] memory outboundCCVs = new address[](2);
    outboundCCVs[0] = makeAddr("outboundCCV1");
    outboundCCVs[1] = makeAddr("outboundCCV2");
    address[] memory inboundCCVs = new address[](1);
    inboundCCVs[0] = makeAddr("inboundCCV");
    // Each direction uses a different finality, and each mock matches only the finality that the pool must forward.
    bytes4 outboundFinality = FinalityCodec._encodeBlockDepth(5);
    bytes4 inboundFinality = FinalityCodec.WAIT_FOR_SAFE_FLAG;
    _mockLaneVerification(outboundFinality, outboundCCVs, inboundCCVs);
    _mockLaneVerification(inboundFinality, outboundCCVs, inboundCCVs);

    assertEq(
      s_pool.getRequiredCCVs(
        address(s_token), DEST_CHAIN_SELECTOR, 1e18, outboundFinality, "", IPoolV2.MessageDirection.Outbound
      ),
      outboundCCVs
    );
    assertEq(
      s_pool.getRequiredCCVs(
        address(s_token), DEST_CHAIN_SELECTOR, 1e18, inboundFinality, "", IPoolV2.MessageDirection.Inbound
      ),
      inboundCCVs
    );
  }

  function _mockLaneVerification(
    bytes4 requestedFinalityConfig,
    address[] memory outboundCCVs,
    address[] memory inboundCCVs
  ) internal {
    address[][] memory outboundCCVLists = new address[][](1);
    outboundCCVLists[0] = outboundCCVs;
    address[][] memory inboundCCVLists = new address[][](1);
    inboundCCVLists[0] = inboundCCVs;
    vm.mockCall(
      s_poolConfigManager,
      abi.encodeCall(IPoolConfigManager.getLaneVerificationBatch, (_laneKeys(), requestedFinalityConfig)),
      abi.encode(new address[](1), outboundCCVLists, inboundCCVLists)
    );
  }
}
