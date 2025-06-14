// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface IVerifier {
  function validateReport(bytes memory rawReport, bytes memory ocrProof, uint256 verifierIndex) external;
}
