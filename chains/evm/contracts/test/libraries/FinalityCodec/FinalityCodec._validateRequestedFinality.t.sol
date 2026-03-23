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
    s_helper.validateRequestedFinality(bytes2(uint16(1)));
    s_helper.validateRequestedFinality(bytes2(FinalityCodec.MAX_BLOCK_DEPTH));
  }

  function test__validateRequestedFinality_PureBlockDepth_MidRange() public view {
    s_helper.validateRequestedFinality(bytes2(uint16(500)));
  }

  // Reverts

  function test__validateRequestedFinality_RevertWhen_InvalidRequestedFinality_FlagWithNonZeroDepth() public {
    bytes2 invalid = bytes2(uint16(uint16(FinalityCodec.WAIT_FOR_SAFE_FLAG) | 1));
    vm.expectRevert(abi.encodeWithSelector(FinalityCodec.InvalidRequestedFinality.selector, invalid));
    s_helper.validateRequestedFinality(invalid);
  }

  function test__validateRequestedFinality_RevertWhen_InvalidRequestedFinality_MultipleFlagBits() public {
    bytes2 invalid = bytes2(uint16((1 << 10) | (1 << 11)));
    vm.expectRevert(abi.encodeWithSelector(FinalityCodec.InvalidRequestedFinality.selector, invalid));
    s_helper.validateRequestedFinality(invalid);
  }
}
