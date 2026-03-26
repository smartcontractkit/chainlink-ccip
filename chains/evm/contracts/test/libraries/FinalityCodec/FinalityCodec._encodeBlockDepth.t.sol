// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {FinalityCodec} from "../../../libraries/FinalityCodec.sol";
import {FinalityCodecSetup} from "./FinalityCodecSetup.t.sol";

contract FinalityCodec__encodeBlockDepth is FinalityCodecSetup {
  function test__encodeBlockDepth() public view {
    assertEq(bytes4(uint32(0)), s_helper.encodeBlockDepth(0));
    assertEq(bytes4(0x00000001), s_helper.encodeBlockDepth(1));
    assertEq(bytes4(0x00000010), s_helper.encodeBlockDepth(16));
    assertEq(bytes4(0x00000100), s_helper.encodeBlockDepth(256));
    assertEq(bytes4(uint32(FinalityCodec.MAX_BLOCK_DEPTH)), s_helper.encodeBlockDepth(FinalityCodec.MAX_BLOCK_DEPTH));

    assertEq(bytes4(0x00010000), FinalityCodec.WAIT_FOR_SAFE_FLAG);
  }
}
