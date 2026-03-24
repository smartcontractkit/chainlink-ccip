// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {FinalityCodec} from "../../../libraries/FinalityCodec.sol";
import {FinalityCodecSetup} from "./FinalityCodecSetup.t.sol";

contract FinalityCodec__encodeBlockDepth is FinalityCodecSetup {
  function test__encodeBlockDepth() public view {
    assertEq(bytes2(uint16(0)), s_helper.encodeBlockDepth(0));
    assertEq(bytes2(0x0001), s_helper.encodeBlockDepth(1));
    assertEq(bytes2(0x0010), s_helper.encodeBlockDepth(16));
    assertEq(bytes2(0x0100), s_helper.encodeBlockDepth(256));
    assertEq(bytes2(0x03ff), s_helper.encodeBlockDepth(FinalityCodec.MAX_BLOCK_DEPTH));

    // same as binary 0000 0100 0000 0000
    assertEq(bytes2(0x0400), FinalityCodec.WAIT_FOR_SAFE_FLAG);
  }

  // Reverts

  function test__encodeBlockDepth_RevertWhen_InvalidBlockDepth_DepthExceedsMax() public {
    vm.expectRevert(
      abi.encodeWithSelector(FinalityCodec.InvalidBlockDepth.selector, uint16(1024), FinalityCodec.MAX_BLOCK_DEPTH)
    );
    s_helper.encodeBlockDepth(1024);
  }
}
