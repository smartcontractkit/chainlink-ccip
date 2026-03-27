// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {FinalityCodec} from "../../../libraries/FinalityCodec.sol";
import {FinalityCodecSetup} from "./FinalityCodecSetup.t.sol";

contract FinalityCodec__validateRequestedFinality is FinalityCodecSetup {
  function test__validateRequestedFinality_WaitForFinality() public view {
    s_helper.validateRequestedFinality(FinalityCodec.WAIT_FOR_FINALITY_FLAG);
  }

  function test__validateRequestedFinality_WaitForSafe() public view {
    s_helper.validateRequestedFinality(FinalityCodec.WAIT_FOR_SAFE_FLAG);
  }

  function test__validateRequestedFinality_PureBlockDepth_Boundaries() public view {
    s_helper.validateRequestedFinality(FinalityCodec._encodeBlockDepth(1));
    s_helper.validateRequestedFinality(FinalityCodec._encodeBlockDepth(10));
    s_helper.validateRequestedFinality(FinalityCodec._encodeBlockDepth(FinalityCodec.MAX_BLOCK_DEPTH));
  }

  // Reverts

  function test__validateRequestedFinality_RevertWhen_InvalidRequestedFinality_FlagWithNonZeroDepth() public {
    bytes4 invalid = FinalityCodec.WAIT_FOR_SAFE_FLAG | FinalityCodec._encodeBlockDepth(1);
    vm.expectRevert(abi.encodeWithSelector(FinalityCodec.RequestedFinalityCanOnlyHaveOneMode.selector, invalid));
    s_helper.validateRequestedFinality(invalid);
  }

  function test__validateRequestedFinality_RevertWhen_InvalidRequestedFinality_MultipleFlagBits() public {
    bytes4 invalid = bytes4(uint32((1 << 16) | (1 << 17)));
    vm.expectRevert(abi.encodeWithSelector(FinalityCodec.RequestedFinalityCanOnlyHaveOneMode.selector, invalid));
    s_helper.validateRequestedFinality(invalid);
  }
}
