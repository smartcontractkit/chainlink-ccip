// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {Internal} from "../libraries/Internal.sol";

interface IVerifier {
  function validateReport(
    bytes memory rawReport,
    bytes memory ocrProof,
    uint256 verifierIndex,
    Internal.MessageExecutionState originalState
  ) external;
}
