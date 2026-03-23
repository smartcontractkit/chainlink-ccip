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
    bytes2 allowed = bytes2(uint16(uint16(FinalityCodec.WAIT_FOR_SAFE_FLAG) | 500));
    s_helper.ensureRequestedFinalityAllowed(requested, allowed);
  }

  function test__ensureRequestedFinalityAllowed_AllowedWhen_BlockDepthWithinAllowance() public view {
    s_helper.ensureRequestedFinalityAllowed(bytes2(uint16(50)), bytes2(uint16(100)));
    s_helper.ensureRequestedFinalityAllowed(bytes2(uint16(100)), bytes2(uint16(100)));
  }

  function test__ensureRequestedFinalityAllowed_AllowedWhen_RequestedSafeAndAllowedIsDepthOnly() public view {
    s_helper.ensureRequestedFinalityAllowed(FinalityCodec.WAIT_FOR_SAFE_FLAG, bytes2(uint16(200)));
  }

  /// forge-config: default.fuzz.runs = 1024
  /// forge-config: ccip.fuzz.runs = 1024
  function testFuzz__ensureRequestedFinalityAllowed_FinalityAlwaysAllowed(
    bytes2 allowed
  ) public view {
    s_helper.ensureRequestedFinalityAllowed(bytes2(0), allowed);
  }

  /// forge-config: default.fuzz.runs = 1024
  /// forge-config: ccip.fuzz.runs = 1024
  function testFuzz__ensureRequestedFinalityAllowed_BlockDepthAllowedWhen_LessOrEqual(
    uint16 allowedDepth,
    uint16 requestedDepth
  ) public view {
    allowedDepth = uint16(bound(uint256(allowedDepth), 0, uint256(FinalityCodec.MAX_BLOCK_DEPTH)));
    requestedDepth = uint16(bound(uint256(requestedDepth), 0, uint256(allowedDepth)));
    s_helper.ensureRequestedFinalityAllowed(bytes2(requestedDepth), bytes2(allowedDepth));
  }

  // Reverts

  function test__ensureRequestedFinalityAllowed_RevertWhen_InvalidBlockDepth_RequestedDepthExceedsAllowed() public {
    bytes2 requested = bytes2(uint16(101));
    bytes2 allowed = bytes2(uint16(100));
    vm.expectRevert(abi.encodeWithSelector(FinalityCodec.InvalidBlockDepth.selector, uint16(101), uint16(100)));
    s_helper.ensureRequestedFinalityAllowed(requested, allowed);
  }

  function test__ensureRequestedFinalityAllowed_RevertWhen_InvalidBlockDepth_NoMatchingFlagAndRequestedDepthExceedsAllowed()
    public
  {
    bytes2 requested = bytes2(uint16(1));
    bytes2 allowed = FinalityCodec.WAIT_FOR_SAFE_FLAG;
    vm.expectRevert(abi.encodeWithSelector(FinalityCodec.InvalidBlockDepth.selector, uint16(1), uint16(0)));
    s_helper.ensureRequestedFinalityAllowed(requested, allowed);
  }

  /// @dev Documents `_ensureRequestedFinalityAllowed` when flag bits overlap but lower bits are non-zero (ill-formed
  /// for a validated request; callers should run `_validateRequestedFinality` first).
  function test__ensureRequestedFinalityAllowed_RevertWhen_InvalidRequestedFinality_FlagOverlapWithNonZeroDepth()
    public
  {
    bytes2 requested = bytes2(uint16(uint16(FinalityCodec.WAIT_FOR_SAFE_FLAG) | 7));
    bytes2 allowed = bytes2(uint16(uint16(FinalityCodec.WAIT_FOR_SAFE_FLAG) | 500));
    vm.expectRevert(abi.encodeWithSelector(FinalityCodec.InvalidRequestedFinality.selector, requested));
    s_helper.ensureRequestedFinalityAllowed(requested, allowed);
  }

  /// forge-config: default.fuzz.runs = 1024
  /// forge-config: ccip.fuzz.runs = 1024
  function testFuzz__ensureRequestedFinalityAllowed_RevertWhen_InvalidBlockDepth_RequestedDepthGreaterThanAllowed(
    uint16 allowedDepth,
    uint16 requestedDepth
  ) public {
    allowedDepth = uint16(bound(uint256(allowedDepth), 0, uint256(FinalityCodec.MAX_BLOCK_DEPTH - 1)));
    requestedDepth =
      uint16(bound(uint256(requestedDepth), uint256(allowedDepth) + 1, uint256(FinalityCodec.MAX_BLOCK_DEPTH)));
    vm.expectRevert(abi.encodeWithSelector(FinalityCodec.InvalidBlockDepth.selector, requestedDepth, allowedDepth));
    s_helper.ensureRequestedFinalityAllowed(bytes2(requestedDepth), bytes2(allowedDepth));
  }
}
