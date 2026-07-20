// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/// @notice This interface contains the only RMN-related functions that might be used on-chain by other CCIP contracts.
interface IRMN {
  /// @notice gets the current set of cursed subjects.
  /// @return subjects the list of cursed subjects.
  function getCursedSubjects() external view returns (bytes16[] memory subjects);

  /// @notice Iff there is an active global or legacy curse, this function returns true.
  /// @return bool true if there is an active global curse.
  function isCursed() external view returns (bool);

  /// @notice Iff there is an active global curse, or an active curse for `subject`, this function returns true.
  /// @param subject To check whether a particular chain is cursed, set to bytes16(uint128(chainSelector)).
  /// @return bool true if the provided subject is cursed *or* if there is an active global curse.
  function isCursed(
    bytes16 subject
  ) external view returns (bool);

  /// @notice Legacy struct: used for blessings, kept to maintain compatibility with the IRMN interface.
  struct TaggedRoot {
    address commitStore;
    bytes32 root;
  }

  /// @notice Legacy CCIP needs this function but blessings have been removed so it always returns true.
  function isBlessed(
    TaggedRoot calldata taggedRoot
  ) external view returns (bool);
}
