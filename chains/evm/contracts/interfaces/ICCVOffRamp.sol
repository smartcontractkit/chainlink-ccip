// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {Internal} from "../libraries/Internal.sol";

interface ICCVOffRamp {
  /// @notice Message execution
  function validateReport(
    bytes memory rawMessage,
    bytes32 messageHash,
    bytes memory ccvBlob,
    bytes memory proof,
    Internal.MessageExecutionState originalState
  ) external;
}
