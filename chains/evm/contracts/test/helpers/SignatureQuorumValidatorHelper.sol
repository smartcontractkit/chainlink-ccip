// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {SignatureQuorumValidator} from "../../ccvs/components/SignatureQuorumValidator.sol";

contract SignatureQuorumValidatorHelper is SignatureQuorumValidator {
  function validateSignatures(bytes32 reportHash, bytes calldata signatures) external view {
    _validateSignatures(reportHash, signatures);
  }
}
