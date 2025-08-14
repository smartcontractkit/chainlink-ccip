// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {Internal} from "../libraries/Internal.sol";

interface ICCVOffRamp {
  /// @notice Message execution
  function validateReport(
    bytes memory rawReport,
    bytes memory ccvBlob,
    bytes memory proof,
    uint256 verifierIndex,
    Internal.MessageExecutionState originalState
  ) external;
}
