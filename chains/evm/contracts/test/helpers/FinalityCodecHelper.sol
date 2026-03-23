// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {FinalityCodec} from "../../libraries/FinalityCodec.sol";

/// @notice Exposes `FinalityCodec` internal functions for unit tests.
contract FinalityCodecHelper {
  function encodeBlockDepth(
    uint16 blockDepth
  ) external pure returns (bytes2) {
    return FinalityCodec._encodeBlockDepth(blockDepth);
  }

  function encodeBlockDepthAndSafeFlag(
    uint16 blockDepth
  ) external pure returns (bytes2) {
    return FinalityCodec._encodeBlockDepthAndSafeFlag(blockDepth);
  }

  function validateRequestedFinality(
    bytes2 encodedFinality
  ) external pure {
    FinalityCodec._validateRequestedFinality(encodedFinality);
  }

  function ensureRequestedFinalityAllowed(
    bytes2 requestedFinality,
    bytes2 allowedFinality
  ) external pure {
    FinalityCodec._ensureRequestedFinalityAllowed(requestedFinality, allowedFinality);
  }
}
