// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {SignatureQuorumVerifier} from "../../../ocr/SignatureQuorumVerifier.sol";
import {SignatureVerifierSetup} from "./SignatureVerifierSetup.t.sol";

contract SignatureQuorumVerifier_setSignatureConfigs is SignatureVerifierSetup {
  function setUp() public virtual override {
    super.setUp();
  }
}
