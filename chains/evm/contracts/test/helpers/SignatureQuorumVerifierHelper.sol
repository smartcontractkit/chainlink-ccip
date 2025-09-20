// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {SignatureQuorumVerifier} from "../../ccvs/components/SignatureQuorumVerifier.sol";

contract SignatureQuorumVerifierHelper is SignatureQuorumVerifier {
  function validateSignatures(bytes32 reportHash, bytes calldata signatures) external view {
    _validateSignatures(reportHash, signatures);
  }
}
