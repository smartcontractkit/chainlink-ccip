// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {FinalityCodec} from "../../../libraries/FinalityCodec.sol";
import {FinalityCodecSetup} from "./FinalityCodecSetup.t.sol";

contract FinalityCodec__ensureRequestedFinalityAllowed is FinalityCodecSetup {
  function test__ensureRequestedFinalityAllowed_FinalityAlwaysAllowed() public view {
    s_helper.ensureRequestedFinalityAllowed(FinalityCodec.WAIT_FOR_FINALITY_FLAG, FinalityCodec.WAIT_FOR_FINALITY_FLAG);
    s_helper.ensureRequestedFinalityAllowed(FinalityCodec.WAIT_FOR_FINALITY_FLAG, FinalityCodec._encodeBlockDepth(1));
    s_helper.ensureRequestedFinalityAllowed(FinalityCodec.WAIT_FOR_FINALITY_FLAG, FinalityCodec.WAIT_FOR_SAFE_FLAG);
  }

  function test__ensureRequestedFinalityAllowed_AllowedWhen_UpperFlagBitsOverlap() public view {
    bytes4 requested = FinalityCodec.WAIT_FOR_SAFE_FLAG;
    bytes4 allowed = bytes4(uint32(uint32(FinalityCodec.WAIT_FOR_SAFE_FLAG) | 500));
    s_helper.ensureRequestedFinalityAllowed(requested, allowed);
  }

  function test__ensureRequestedFinalityAllowed_AllowedWhen_BlockDepthMeetsMinimum() public view {
    uint16 requestedDepth = 100;
    // Exact match — requesting exactly the minimum is allowed.
    s_helper.ensureRequestedFinalityAllowed(
      FinalityCodec._encodeBlockDepth(requestedDepth), FinalityCodec._encodeBlockDepth(requestedDepth)
    );
    // Requesting more confirmations than the minimum is also allowed.
    s_helper.ensureRequestedFinalityAllowed(
      FinalityCodec._encodeBlockDepth(requestedDepth * 2), FinalityCodec._encodeBlockDepth(requestedDepth)
    );
  }

  function test__ensureRequestedFinalityAllowed_RevertWhen_InvalidBlockDepth_RequestedSafeButAllowedIsDepthOnly()
    public
  {
    vm.expectRevert(
      abi.encodeWithSelector(
        FinalityCodec.InvalidRequestedFinality.selector,
        FinalityCodec.WAIT_FOR_SAFE_FLAG,
        FinalityCodec._encodeBlockDepth(200)
      )
    );
    s_helper.ensureRequestedFinalityAllowed(FinalityCodec.WAIT_FOR_SAFE_FLAG, FinalityCodec._encodeBlockDepth(200));
  }

  function test__ensureRequestedFinalityAllowed_RevertWhen_InvalidBlockDepth_RequestedDepthBelowMinimum() public {
    uint16 requested = 99;
    uint16 allowed = requested + 1;

    vm.expectRevert(
      abi.encodeWithSelector(
        FinalityCodec.InvalidRequestedFinality.selector,
        FinalityCodec._encodeBlockDepth(requested),
        FinalityCodec._encodeBlockDepth(allowed)
      )
    );
    s_helper.ensureRequestedFinalityAllowed(
      FinalityCodec._encodeBlockDepth(requested), FinalityCodec._encodeBlockDepth(allowed)
    );
  }

  function test__ensureRequestedFinalityAllowed_RevertWhen_InvalidBlockDepth_NoMatchingFlagAndRequestedDepthExceedsAllowed()
    public
  {
    uint16 requested = 1;
    bytes4 allowed = FinalityCodec.WAIT_FOR_SAFE_FLAG;

    vm.expectRevert(
      abi.encodeWithSelector(
        FinalityCodec.InvalidRequestedFinality.selector, FinalityCodec._encodeBlockDepth(requested), allowed
      )
    );
    s_helper.ensureRequestedFinalityAllowed(FinalityCodec._encodeBlockDepth(requested), allowed);
  }

  function test__ensureRequestedFinalityAllowed_RevertWhen_MalformedRequest_FlagAndDepthCombined() public {
    // requested has both WAIT_FOR_SAFE_FLAG (bit 16) and non-zero depth bits — two active modes.
    // _validateRequestedFinality (called internally) catches this before the allowance check.
    bytes4 requested = bytes4(uint32(uint32(FinalityCodec.WAIT_FOR_SAFE_FLAG) | 7));
    bytes4 allowed = bytes4(uint32(uint32(FinalityCodec.WAIT_FOR_SAFE_FLAG) | 500));
    vm.expectRevert(abi.encodeWithSelector(FinalityCodec.RequestedFinalityCanOnlyHaveOneMode.selector, requested));
    s_helper.ensureRequestedFinalityAllowed(requested, allowed);
  }

  function test__ensureRequestedFinalityAllowed_RevertWhen_InvalidRequestedFinality_FinalityRequired() public {
    uint16 requested = 50;
    vm.expectRevert(
      abi.encodeWithSelector(
        FinalityCodec.InvalidRequestedFinality.selector,
        FinalityCodec._encodeBlockDepth(requested),
        FinalityCodec.WAIT_FOR_FINALITY_FLAG
      )
    );
    s_helper.ensureRequestedFinalityAllowed(
      FinalityCodec._encodeBlockDepth(requested), FinalityCodec.WAIT_FOR_FINALITY_FLAG
    );
  }

  function test__ensureRequestedFinalityAllowed_AllowedWhen_BlockDepthAtMaxVsMinAllowed() public view {
    s_helper.ensureRequestedFinalityAllowed(
      FinalityCodec._encodeBlockDepth(FinalityCodec.MAX_BLOCK_DEPTH), FinalityCodec._encodeBlockDepth(1)
    );
  }

  function test__ensureRequestedFinalityAllowed_AllowedWhen_BlockDepthMinimumExact() public view {
    s_helper.ensureRequestedFinalityAllowed(FinalityCodec._encodeBlockDepth(1), FinalityCodec._encodeBlockDepth(1));
  }

  function test__ensureRequestedFinalityAllowed_AllowedWhen_ReservedFlagOverlap() public view {
    bytes4 reservedFlag = bytes4(uint32(1 << 17));
    s_helper.ensureRequestedFinalityAllowed(reservedFlag, reservedFlag);
  }

  function test__ensureRequestedFinalityAllowed_RevertWhen_InvalidRequestedFinality_ReservedFlagNotInAllowed() public {
    bytes4 requestedReserved = bytes4(uint32(1 << 17));
    bytes4 allowedSafe = FinalityCodec.WAIT_FOR_SAFE_FLAG; // bit 16 only, no bit 17
    vm.expectRevert(
      abi.encodeWithSelector(FinalityCodec.InvalidRequestedFinality.selector, requestedReserved, allowedSafe)
    );
    s_helper.ensureRequestedFinalityAllowed(requestedReserved, allowedSafe);
  }
}
