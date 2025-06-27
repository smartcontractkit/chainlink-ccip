// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface IVerifierRegistry {
  function getVerifier(
    bytes32 verifierId
  ) external view returns (address);
}
