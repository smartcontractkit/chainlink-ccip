// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {IVerifierSender} from "../../interfaces/verifiers/IVerifier.sol";

contract MockVerifier is IVerifierSender {
  function forwardToVerifier(bytes memory rawMessage, uint256 verifierIndex) external returns (bytes memory) {
    return "";
  }
}
