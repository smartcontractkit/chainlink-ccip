// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {SignatureQuorumVerifier} from "../../offRamp/components/SignatureQuorumVerifier.sol";

contract SignatureQuorumVerifierHelper is SignatureQuorumVerifier {
  function validateSignatures(bytes32 reportHash, bytes32[] calldata rs, bytes32[] calldata ss) external view {
    _validateSignatures(reportHash, rs, ss);
  }
}
