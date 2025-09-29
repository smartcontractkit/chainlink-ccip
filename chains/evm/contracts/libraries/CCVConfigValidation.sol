// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

/// @notice CCV config validation helpers.
library CCVConfigValidation {
  error MustSpecifyDefaultOrRequiredCCVs();
  error DuplicateCCVNotAllowed(address ccvAddress);
  error ZeroAddressNotAllowed();

  /// @notice Ensures at least one CCV combined, no zero addresses, no duplicates within or across both sets.
  /// @param defaultCCV The default CCVs.
  /// @param laneMandatedCCVs The mandated CCVs.
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

  function _assertNoDuplicates(
    address[] memory addresses
  ) internal pure {
    uint256 length = addresses.length;
    for (uint256 i = 0; i < length; ++i) {
      for (uint256 j = i + 1; j < length; ++j) {
        if (addresses[i] == addresses[j]) revert DuplicateCCVNotAllowed(addresses[i]);
      }
    }
  }
}
