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
    s_poolConfigManager.setRequiredCCVs(outboundCCVs, inboundCCVs);

    // The pool passes the requested finality to the pool config manager. Each direction uses a different value.
    bytes4 outboundFinality = FinalityCodec._encodeBlockDepth(5);
    _expectLaneVerificationCall(outboundFinality);
    assertEq(
      s_pool.getRequiredCCVs(
        address(s_token), DEST_CHAIN_SELECTOR, 1e18, outboundFinality, "", IPoolV2.MessageDirection.Outbound
      ),
      outboundCCVs
    );
    bytes4 inboundFinality = FinalityCodec.WAIT_FOR_SAFE_FLAG;
    _expectLaneVerificationCall(inboundFinality);
    assertEq(
      s_pool.getRequiredCCVs(
        address(s_token), DEST_CHAIN_SELECTOR, 1e18, inboundFinality, "", IPoolV2.MessageDirection.Inbound
      ),
      inboundCCVs
    );
  }

  function _expectLaneVerificationCall(
    bytes4 requestedFinalityConfig
  ) internal {
    IPoolConfigManager.LaneKey[] memory lanes = new IPoolConfigManager.LaneKey[](1);
    lanes[0] = IPoolConfigManager.LaneKey({localToken: address(s_token), remoteChainSelector: DEST_CHAIN_SELECTOR});
    vm.expectCall(
      address(s_poolConfigManager),
      abi.encodeCall(IPoolConfigManager.getLaneVerificationBatch, (lanes, requestedFinalityConfig)),
      1
    );
  }
}
