// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

/// @notice This library provides encoding and validation for finality parameters used in cross-chain transfers.
/// @dev this codec supports all the bit flags, even though some might not be assigned any meaning yet. This is
/// intentional to allow for future flexibility.
library FinalityCodec {
  error InvalidBlockDepth(uint16 requestedDepth, uint16 maxDepth);
  /// @notice Requested finality must be exactly one mode: any of the flag bits or a block depth with no upper flag bits.
  /// It cannot combine a flag with a block depth.
  error InvalidRequestedFinality(bytes2 encodedFinality);

  /// @notice The block depth is stored in the lower 10 bits, leaving the upper 6 bits for flags. This allows for a
  /// maximum block depth of 1023, which should be sufficient for most use cases. For more security, users should wait
  /// for finality instead (bytes2(0)).
  uint256 public constant BLOCK_DEPTH_BITS = 10;
  /// @notice The maximum block depth that can be encoded in the finality params: 1023.
  uint16 public constant MAX_BLOCK_DEPTH = uint16((1 << BLOCK_DEPTH_BITS) - 1);
  /// @notice The block depth mask to extract the block depth from the finality params.
  bytes2 public constant BLOCK_DEPTH_MASK = bytes2(MAX_BLOCK_DEPTH);

  /// @notice The finality flag for waiting for finality is 0, this is the safest option. Any block depth that's deeper
  /// than finality will fall back to finality, meaning a very deep block depth will not be more secure than finality.
  bytes2 public constant WAIT_FOR_FINALITY_FLAG = bytes2(0);
  /// @notice Signals to wait for the `safe` tag.
  bytes2 public constant WAIT_FOR_SAFE_FLAG = bytes2(uint16(1 << BLOCK_DEPTH_BITS));

  /// @notice Helper to encode block depth into the finality params. Will revert if the block depth is greater than the
  /// maximum block depth.
  /// @param blockDepth The block depth to encode into the finality params.
  /// @return The encoded finality params with the block depth.
  function _encodeBlockDepth(
    uint16 blockDepth
  ) internal pure returns (bytes2) {
    if (blockDepth > MAX_BLOCK_DEPTH) {
      revert InvalidBlockDepth(blockDepth, MAX_BLOCK_DEPTH);
    }
    return bytes2(blockDepth);
  }

  /// @notice Helper to encode the `safe` tag plus a block depth into the finality params.
  /// NOTE: this format is only allowed for allowed finality, not requested finality, as requested finality can only
  /// contain a single flag or block depth, but allowed finality can contain multiple.
  /// @param blockDepth The block depth to encode into the finality params.
  /// @return The encoded finality params with the `safe` tag and block depth.
  function _encodeBlockDepthAndSafeFlag(
    uint16 blockDepth
  ) internal pure returns (bytes2) {
    return _encodeBlockDepth(blockDepth) | WAIT_FOR_SAFE_FLAG;
  }

  /// @notice Validates requested finality: either `bytes2(0)`, exactly one set bit among the upper flag bits, or a pure
  /// block depth (no flag bits, depth in `1..MAX_BLOCK_DEPTH`). Never a flag combined with a non-zero depth. Unknown
  /// flags are accepted here for wire compatibility; pools/CCVs reject modes they do not implement.
  /// @param encodedFinality The encoded finality params to validate.
  function _validateRequestedFinality(
    bytes2 encodedFinality
  ) internal pure {
    // Waiting for finality is always valid.
    if (encodedFinality == WAIT_FOR_FINALITY_FLAG) {
      return;
    }
    uint16 finality = uint16(encodedFinality);
    bool hasBlockDepth = (finality & MAX_BLOCK_DEPTH) != 0;
    uint256 activeModes = hasBlockDepth ? 1 : 0; // If it has depth, it counts as one active mode.

    uint16 flags = finality >> BLOCK_DEPTH_BITS;
    if (flags != 0) {
      for (uint256 i = 0; i < 6; ++i) {
        if ((flags & (1 << i)) != 0) {
          activeModes += 1;
        }
      }
    }
    // There must be exactly one active mode: either a block depth or a single flag. Selecting multiple modes is only
    // allowed for `allowedFinality` set by Pools, CCVs, etc., but not for `requestedFinality` set by senders.
    if (activeModes != 1) {
      revert InvalidRequestedFinality(encodedFinality);
    }
  }

  /// @notice Ensures `requestedFinality` is permitted by `allowedFinality`. When matching on flags, the request must not
  /// carry a block depth (lower bits zero aside from the all-zero finality case, which is handled earlier).
  /// @param requestedFinality The requested finality params to check. This value must already be validated by
  /// `_validateRequestedFinality` to ensure it is well-formed.
  /// @param allowedFinality The allowed finality params to check against.
  function _ensureRequestedFinalityAllowed(
    bytes2 requestedFinality,
    bytes2 allowedFinality
  ) internal pure {
    // Finality is always allowed.
    if (requestedFinality == bytes2(0)) {
      return;
    }
    // If any of the flags match, the request is allowed only when it has no depth field (flag-only request).
    if (requestedFinality >> BLOCK_DEPTH_BITS & allowedFinality >> BLOCK_DEPTH_BITS != 0) {
      if (uint16(requestedFinality & BLOCK_DEPTH_MASK) != 0) {
        revert InvalidRequestedFinality(requestedFinality);
      }
      return;
    }
    // Otherwise, it must be block-depth based.
    uint16 requestedBlockDepth = uint16(requestedFinality & BLOCK_DEPTH_MASK);
    uint16 allowedBlockDepth = uint16(allowedFinality & BLOCK_DEPTH_MASK);
    if (allowedBlockDepth == 0 || requestedBlockDepth > allowedBlockDepth) {
      revert InvalidBlockDepth(requestedBlockDepth, allowedBlockDepth);
    }
  }
}
