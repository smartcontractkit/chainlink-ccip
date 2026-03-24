// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {FinalityCodec} from "../../../libraries/FinalityCodec.sol";
import {FinalityCodecSetup} from "./FinalityCodecSetup.t.sol";

contract FinalityCodec__ensureRequestedFinalityAllowed is FinalityCodecSetup {
  function test__ensureRequestedFinalityAllowed_FinalityAlwaysAllowed() public view {
    s_helper.ensureRequestedFinalityAllowed(bytes2(0), bytes2(0));
    s_helper.ensureRequestedFinalityAllowed(bytes2(0), bytes2(uint16(1)));
    s_helper.ensureRequestedFinalityAllowed(bytes2(0), FinalityCodec.WAIT_FOR_SAFE_FLAG);
  }

  function test__ensureRequestedFinalityAllowed_AllowedWhen_UpperFlagBitsOverlap() public view {
    bytes2 requested = FinalityCodec.WAIT_FOR_SAFE_FLAG;
    bytes2 allowed = bytes2(uint16(FinalityCodec.WAIT_FOR_SAFE_FLAG) | 500);
    s_helper.ensureRequestedFinalityAllowed(requested, allowed);
  }

  function test__ensureRequestedFinalityAllowed_AllowedWhen_BlockDepthMeetsMinimum() public view {
    uint16 requestedDepth = 100;
    // Exact match — requesting exactly the minimum is allowed.
    s_helper.ensureRequestedFinalityAllowed(bytes2(requestedDepth), bytes2(requestedDepth));
    // Requesting more confirmations than the minimum is also allowed.
    s_helper.ensureRequestedFinalityAllowed(bytes2(requestedDepth * 2), bytes2(requestedDepth));
  }

  function test__ensureRequestedFinalityAllowed_RevertWhen_InvalidBlockDepth_RequestedSafeButAllowedIsDepthOnly()
    public
  {
    vm.expectRevert(
      abi.encodeWithSelector(
        FinalityCodec.InvalidRequestedFinality.selector, FinalityCodec.WAIT_FOR_SAFE_FLAG, bytes2(uint16(200))
      )
    );
    s_helper.ensureRequestedFinalityAllowed(FinalityCodec.WAIT_FOR_SAFE_FLAG, bytes2(uint16(200)));
  }

  function test__ensureRequestedFinalityAllowed_RevertWhen_InvalidBlockDepth_RequestedDepthBelowMinimum() public {
    uint16 requested = 99;
    uint16 allowed = requested + 1;

    vm.expectRevert(abi.encodeWithSelector(FinalityCodec.InvalidRequestedFinality.selector, bytes2(requested), bytes2(allowed)));
    s_helper.ensureRequestedFinalityAllowed(bytes2(requested), bytes2(allowed));
  }

  function test__ensureRequestedFinalityAllowed_RevertWhen_InvalidBlockDepth_NoMatchingFlagAndRequestedDepthExceedsAllowed()
    public
  {
    uint16 requested = 1;
    bytes2 allowed = FinalityCodec.WAIT_FOR_SAFE_FLAG;

    vm.expectRevert(abi.encodeWithSelector(FinalityCodec.InvalidRequestedFinality.selector, bytes2(requested), allowed));
    s_helper.ensureRequestedFinalityAllowed(bytes2(requested), allowed);
  }

  function test__ensureRequestedFinalityAllowed_RevertWhen_InvalidBlockDepth_FlagOverlapWithNonZeroDepth() public {
    bytes2 requested = bytes2(uint16(uint16(FinalityCodec.WAIT_FOR_SAFE_FLAG) | 7));
    bytes2 allowed = bytes2(uint16(uint16(FinalityCodec.WAIT_FOR_SAFE_FLAG) | 500));
    vm.expectRevert(abi.encodeWithSelector(FinalityCodec.InvalidRequestedFinality.selector, requested, allowed));
    s_helper.ensureRequestedFinalityAllowed(requested, allowed);
  }
}
