// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

/// @notice Test helper library for bytes manipulation.
library BytesTestHelper {
  /// @notice Slice bytes, removes prefix, return bytes array starting from start index.
  /// @param data The bytes to slice.
  /// @param start The index starting from which to return the sub array.
  /// @return Bytes sub array starting from start index.
  function _slice(
    bytes memory data,
    uint256 start
  ) internal pure returns (bytes memory) {
    bytes memory result = new bytes(data.length - start);
    for (uint256 i = 0; i < result.length; ++i) {
      result[i] = data[start + i];
    }
    return result;
  }
}
