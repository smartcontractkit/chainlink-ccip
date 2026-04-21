// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

/// @notice This interface contains the only RMN-related functions that might be used on-chain by other CCIP contracts.
interface IRMNRemote {
  /// @notice gets the current set of cursed subjects.
  /// @return subjects the list of cursed subjects.
  function getCursedSubjects() external view returns (bytes16[] memory subjects);

  /// @notice If there is an active global or legacy curse, this function returns true.
  /// @return bool true if there is an active global curse.
  function isCursed() external view returns (bool);

  /// @notice If there is an active global curse, or an active curse for `subject`, this function returns true.
  /// @param subject To check whether a particular chain is cursed, set to bytes16(uint128(chainSelector)).
  /// @return bool true if the provided subject is cured *or* if there is an active global curse.
  function isCursed(
    bytes16 subject
  ) external view returns (bool);
}
