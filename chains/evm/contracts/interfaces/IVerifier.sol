// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {Internal} from "../libraries/Internal.sol";

interface IVerifier {
  /// @notice Message execution
  function validateReport(
    bytes memory rawReport,
    bytes memory ocrProof,
    uint256 verifierIndex,
    Internal.MessageExecutionState originalState
  ) external;
}

interface IVerifierSender {
  /// @notice Message sending
  // TODO versioning?
  function forwardToVerifier(bytes memory rawMessage, uint256 verifierIndex) external returns (bytes memory);
}
