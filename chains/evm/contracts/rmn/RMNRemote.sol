// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IRMN} from "../interfaces/IRMN.sol";
import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";

import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";
import {EnumerableSet} from "@chainlink/contracts/src/v0.8/shared/enumerable/EnumerableSetWithBytes16.sol";

/// @dev An active curse on this subject will cause isCursed() and isCursed(bytes16) to return true. Use this subject
/// for issues affecting all of CCIP chains, or pertaining to the chain that this contract is deployed on, instead of
/// using the local chain selector as a subject.
bytes16 constant GLOBAL_CURSE_SUBJECT = 0x01000000000000000000000000000001;

/// @notice This contract supports cursing and uncursing of chains.
contract RMNRemote is AuthorizedCallers, ITypeAndVersion, IRMN {
  using EnumerableSet for EnumerableSet.Bytes16Set;

  error AlreadyCursed(bytes16 subject);
  error NotCursed(bytes16 subject);

  event Cursed(bytes16[] subjects);
  event Uncursed(bytes16[] subjects);

  string public constant override typeAndVersion = "RMNRemote 2.1.0";

  EnumerableSet.Bytes16Set private s_cursedSubjects;

  /// @param curseAdmins initial set of addresses authorized to call curse.
  constructor(
    address[] memory curseAdmins
  ) AuthorizedCallers(curseAdmins) {}

  // ================================================================
  // │                           Cursing                            │
  // ================================================================

  /// @notice Curse a single subject.
  /// @param subject the subject to curse.
  function curse(
    bytes16 subject
  ) external {
    bytes16[] memory subjects = new bytes16[](1);
    subjects[0] = subject;
    curse(subjects);
  }

  /// @notice Curse an array of subjects.
  /// @param subjects the subjects to curse.
  /// @dev reverts if any of the subjects are already cursed or if there is a duplicate.
  function curse(
    bytes16[] memory subjects
  ) public {
    if (msg.sender != owner()) _validateCaller();
    for (uint256 i = 0; i < subjects.length; ++i) {
      if (!s_cursedSubjects.add(subjects[i])) {
        revert AlreadyCursed(subjects[i]);
      }
    }
    emit Cursed(subjects);
  }

  /// @notice Uncurse a single subject.
  /// @param subject the subject to uncurse.
  function uncurse(
    bytes16 subject
  ) external {
    bytes16[] memory subjects = new bytes16[](1);
    subjects[0] = subject;
    uncurse(subjects);
  }

  /// @notice Uncurse an array of subjects.
  /// @param subjects the subjects to uncurse.
  /// @dev reverts if any of the subjects are not cursed or if there is a duplicate.
  function uncurse(
    bytes16[] memory subjects
  ) public onlyOwner {
    for (uint256 i = 0; i < subjects.length; ++i) {
      if (!s_cursedSubjects.remove(subjects[i])) {
        revert NotCursed(subjects[i]);
      }
    }
    emit Uncursed(subjects);
  }

  /// @inheritdoc IRMN
  function getCursedSubjects() external view returns (bytes16[] memory subjects) {
    return s_cursedSubjects.values();
  }

  /// @inheritdoc IRMN
  function isCursed() external view override returns (bool) {
    // There are zero curses under normal circumstances, which means it's cheaper to check for the absence of curses.
    // than to check the subject list for the global curse subject.
    if (s_cursedSubjects.length() == 0) {
      return false;
    }
    return s_cursedSubjects.contains(GLOBAL_CURSE_SUBJECT);
  }

  /// @inheritdoc IRMN
  function isCursed(
    bytes16 subject
  ) external view override returns (bool) {
    // There are zero curses under normal circumstances, which means it's cheaper to check for the absence of curses.
    // than to check the subject list twice, as we have to check for both the given and global curse subjects.
    if (s_cursedSubjects.length() == 0) {
      return false;
    }
    return s_cursedSubjects.contains(subject) || s_cursedSubjects.contains(GLOBAL_CURSE_SUBJECT);
  }
}
