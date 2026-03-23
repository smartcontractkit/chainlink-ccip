// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {FinalityCodec} from "../../../libraries/FinalityCodec.sol";
import {FinalityCodecSetup} from "./FinalityCodecSetup.t.sol";

contract FinalityCodec__encodeBlockDepth is FinalityCodecSetup {
  function test__encodeBlockDepth_ZeroAndMax() public view {
    assertEq(bytes2(uint16(0)), s_helper.encodeBlockDepth(0));
    assertEq(bytes2(uint16(1023)), s_helper.encodeBlockDepth(1023));
  }

  function test__encodeBlockDepth_ArbitraryDepth() public view {
    assertEq(bytes2(uint16(100)), s_helper.encodeBlockDepth(100));
  }

  // Reverts

  function test__encodeBlockDepth_RevertWhen_InvalidBlockDepth_DepthExceedsMax() public {
    vm.expectRevert(
      abi.encodeWithSelector(FinalityCodec.InvalidBlockDepth.selector, uint16(1024), FinalityCodec.MAX_BLOCK_DEPTH)
    );
    s_helper.encodeBlockDepth(1024);
  }
}
