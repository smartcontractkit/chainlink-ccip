// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {FinalityCodec} from "../../../libraries/FinalityCodec.sol";
import {FinalityCodecSetup} from "./FinalityCodecSetup.t.sol";

contract FinalityCodec__encodeBlockDepthAndSafeFlag is FinalityCodecSetup {
  function test__encodeBlockDepthAndSafeFlag_CombinesFlagAndDepth() public view {
    assertEq(
      FinalityCodec.WAIT_FOR_SAFE_FLAG,
      s_helper.encodeBlockDepthAndSafeFlag(0),
      "depth 0 leaves only the safe flag in the lower bits as 0"
    );
    bytes2 expected = bytes2(uint16(uint16(FinalityCodec.WAIT_FOR_SAFE_FLAG) | 42));
    assertEq(expected, s_helper.encodeBlockDepthAndSafeFlag(42));
  }

  // Reverts

  function test__encodeBlockDepthAndSafeFlag_RevertWhen_InvalidBlockDepth_DepthExceedsMax() public {
    vm.expectRevert(
      abi.encodeWithSelector(FinalityCodec.InvalidBlockDepth.selector, uint16(1024), FinalityCodec.MAX_BLOCK_DEPTH)
    );
    s_helper.encodeBlockDepthAndSafeFlag(1024);
  }
}
