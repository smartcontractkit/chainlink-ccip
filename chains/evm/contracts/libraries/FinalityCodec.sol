// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

/// @notice This library provides encoding and validation for finality parameters used in cross-chain transfers.
/// @dev this codec supports all the bit flags, even though some might not be assigned any meaning yet. This is
/// intentional to allow for future flexibility.
///
/// @dev Bit layout of the `bytes4` finality value (32 bits, MSB on the left):
///
///  Bit: 31  30  29  28  27  26  25  24  23  22  21  20  19  18  17  16 | 15  14  13  12  11  10   9   8   7   6   5   4   3   2   1   0
///      +---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
///      | R | R | R | R | R | R | R | R | R | R | R | R | R | R | R | S |                          block depth (16 bits)                |
///      +---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
///      \_______________________________  _____________________________/ \______________________________  _____________________________/
///                                      \/                                                              \/
///                                 flags (16 bits)                                                  depth (16 bits)
///                                                                                                max = 65535 (0xFFFF)
///
///  S  (bit 16) = WAIT_FOR_SAFE_FLAG  — wait for the `safe` tag.
///  R  (bits 17-31) = Reserved for future flags (currently unassigned; accepted on the wire).
///                    Reserved bits may be assigned in the future, read the docs for the latest bit definitions.
///
///  Special values:
///    0x00000000  WAIT_FOR_FINALITY_FLAG  — wait for full finality (safest, default).
///    0x00010000  WAIT_FOR_SAFE_FLAG      — wait for the `safe` head (bit 16 set, no depth).
///    0x00000001..0x0000FFFF              — wait for N blocks.
library FinalityCodec {
  error InvalidRequestedFinality(bytes4 requestedFinality, bytes4 allowedFinality);
  /// @notice Requested finality must be exactly one mode: any of the flag bits or a block depth with no upper flag bits.
  /// It cannot combine a flag with a block depth.
  error RequestedFinalityCanOnlyHaveOneMode(bytes4 encodedFinality);

  /// @notice The block depth is stored in the lower 16 bits, leaving the upper 16 bits for flags.
  /// For more security, users should wait for finality instead (bytes4(0)).
  uint256 public constant BLOCK_DEPTH_BITS = 16;
  /// @notice The maximum block depth that can be encoded in the finality params.
  uint16 public constant MAX_BLOCK_DEPTH = type(uint16).max;
  /// @notice The block depth mask to extract the block depth from the finality params.
  bytes4 public constant BLOCK_DEPTH_MASK = bytes4(uint32(MAX_BLOCK_DEPTH));

  /// @notice The finality flag for waiting for finality is 0, this is the safest option. Any block depth that's deeper
  /// than finality will fall back to finality, meaning a very deep block depth will not be more secure than finality.
  bytes4 public constant WAIT_FOR_FINALITY_FLAG = bytes4(0);
  /// @notice Signals to wait for the `safe` tag.
  bytes4 public constant WAIT_FOR_SAFE_FLAG = bytes4(uint32(1 << BLOCK_DEPTH_BITS));

  /// @notice Helper to encode block depth into the finality params. Returns WAIT_FOR_FINALITY_FLAG if the block depth
  /// is zero.
  /// @param blockDepth The block depth to encode into the finality params.
  /// @return The encoded finality params with the block depth.
  function _encodeBlockDepth(
    uint16 blockDepth
  ) internal pure returns (bytes4) {
    return bytes4(uint32(blockDepth));
  }

  /// @notice Helper to encode the `safe` tag plus a block depth into the finality params.
  /// NOTE: this format is only allowed for allowed finality, not requested finality, as requested finality can only
  /// contain a single flag or block depth, but allowed finality can contain multiple.
  /// @param blockDepth The block depth to encode into the finality params.
  /// @return The encoded finality params with the `safe` tag and block depth.
  function _encodeBlockDepthAndSafeFlag(
    uint16 blockDepth
  ) internal pure returns (bytes4) {
    return _encodeBlockDepth(blockDepth) | WAIT_FOR_SAFE_FLAG;
  }

  /// @notice Validates requested finality: either `bytes4(0)`, exactly one set bit among the upper flag bits, or a pure
  /// block depth (no flag bits, depth in `1..MAX_BLOCK_DEPTH`). Never a flag combined with a non-zero depth. Unknown
  /// flags are accepted here for wire compatibility; pools/CCVs reject modes they do not implement.
  /// @param encodedFinality The encoded finality params to validate.
  function _validateRequestedFinality(
    bytes4 encodedFinality
  ) internal pure {
    // Waiting for finality is always valid.
    if (encodedFinality == WAIT_FOR_FINALITY_FLAG) {
      return;
    }
    uint32 finality = uint32(encodedFinality);
    bool hasBlockDepth = (finality & uint32(MAX_BLOCK_DEPTH)) != 0;
    uint256 activeModes = hasBlockDepth ? 1 : 0; // If it has depth, it counts as one active mode.

    uint32 flags = finality >> BLOCK_DEPTH_BITS;
    if (flags != 0) {
      for (uint256 i = 0; i < 16; ++i) {
        if ((flags & (1 << i)) != 0) {
          activeModes += 1;
        }
      }
    }
    // There must be exactly one active mode: either a block depth or a single flag. Selecting multiple modes is only
    // allowed for `allowedFinality` set by Pools, CCVs, etc., but not for `requestedFinality` set by senders.
    if (activeModes != 1) {
      revert RequestedFinalityCanOnlyHaveOneMode(encodedFinality);
    }
  }

  /// @notice Validates that `requestedFinality` is well-formed and permitted by `allowedFinality`.
  /// @param requestedFinality The requested finality params to check.
  /// @param allowedFinality The allowed finality params to check against.
  function _ensureRequestedFinalityAllowed(
    bytes4 requestedFinality,
    bytes4 allowedFinality
  ) internal pure {
    // Finality is always allowed.
    if (requestedFinality == WAIT_FOR_FINALITY_FLAG) {
      return;
    }

    // Validate the structural shape of the requested finality, as it is only allowed to signal one mode.
    _validateRequestedFinality(requestedFinality);

    // If any of the flags match, the request is allowed only when it has no depth field (flag-only request).
    if (((requestedFinality >> BLOCK_DEPTH_BITS) & (allowedFinality >> BLOCK_DEPTH_BITS)) != 0) {
      if (uint32(requestedFinality & BLOCK_DEPTH_MASK) != 0) {
        revert InvalidRequestedFinality(requestedFinality, allowedFinality);
      }
      return;
    }
    // Otherwise, it must be block-depth based.
    uint32 requestedBlockDepth = uint32(requestedFinality & BLOCK_DEPTH_MASK);
    uint32 allowedBlockDepth = uint32(allowedFinality & BLOCK_DEPTH_MASK);
    if (allowedBlockDepth == 0 || requestedBlockDepth < allowedBlockDepth) {
      revert InvalidRequestedFinality(requestedFinality, allowedFinality);
    }
  }
}
