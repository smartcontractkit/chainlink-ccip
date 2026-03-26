// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {FinalityCodec} from "../../libraries/FinalityCodec.sol";

/// @notice Exposes `FinalityCodec` internal functions for unit tests.
contract FinalityCodecHelper {
  function encodeBlockDepth(
    uint16 blockDepth
  ) external pure returns (bytes4) {
    return FinalityCodec._encodeBlockDepth(blockDepth);
  }

  function encodeBlockDepthAndSafeFlag(
    uint16 blockDepth
  ) external pure returns (bytes4) {
    return FinalityCodec._encodeBlockDepthAndSafeFlag(blockDepth);
  }

  function validateRequestedFinality(
    bytes4 encodedFinality
  ) external pure {
    FinalityCodec._validateRequestedFinality(encodedFinality);
  }

  function ensureRequestedFinalityAllowed(
    bytes4 requestedFinality,
    bytes4 allowedFinality
  ) external pure {
    FinalityCodec._ensureRequestedFinalityAllowed(requestedFinality, allowedFinality);
  }
}
