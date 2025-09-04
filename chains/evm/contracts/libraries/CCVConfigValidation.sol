// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

/// @notice CCV config validation helper.
library CCVConfigValidation {
  error MustSpecifyDefaultOrRequiredCCVs();
  error DuplicateCCVNotAllowed(address ccvAddress);
  error ZeroAddressNotAllowed();

  /// @notice Validates default and mandated CCV arrays for non-empty, non-zero, and deduped addresses.
  /// Reverts with detailed custom errors when invalid.
  function _validateDefaultAndMandatedCCVs(
    address[] memory defaultCCV,
    address[] memory laneMandatedCCVs
  ) internal pure {
    uint256 defaultLength = defaultCCV.length;
    uint256 mandatedLength = laneMandatedCCVs.length;
    uint256 totalLength = defaultLength + mandatedLength;

    // There must always be at least one default or mandated CCV. This ensures that any receiver who does not specify
    // CCVs will always have at least one CCV to validate the message.
    if (totalLength == 0) revert MustSpecifyDefaultOrRequiredCCVs();

    // We check for duplicates and zero addresses in the default and mandated CCVs. We need to check for duplicates
    // between the two sets of CCVs as well as within each set. Doing these checks here means we can assume there are
    // no duplicates or zero addresses in the rest of the code.
    for (uint256 combinedIndex = 0; combinedIndex < totalLength; ++combinedIndex) {
      address currentCCVAddress =
        combinedIndex < defaultLength ? defaultCCV[combinedIndex] : laneMandatedCCVs[combinedIndex - defaultLength];
      if (currentCCVAddress == address(0)) revert ZeroAddressNotAllowed();

      for (uint256 nextIndex = combinedIndex + 1; nextIndex < totalLength; ++nextIndex) {
        address compareCCVAddress =
          nextIndex < defaultLength ? defaultCCV[nextIndex] : laneMandatedCCVs[nextIndex - defaultLength];
        if (currentCCVAddress == compareCCVAddress) revert DuplicateCCVNotAllowed(currentCCVAddress);
      }
    }
  }
}
