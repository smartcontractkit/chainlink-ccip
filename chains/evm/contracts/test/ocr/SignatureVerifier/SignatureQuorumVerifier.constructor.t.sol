// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {SignatureQuorumVerifier} from "../../../offRamp/components/SignatureQuorumVerifier.sol";
import {SignatureQuorumVerifierHelper} from "../../helpers/SignatureQuorumVerifierHelper.sol";
import {SignatureVerifierSetup} from "./SignatureVerifierSetup.t.sol";

contract SignatureQuorumVerifier_constructor is SignatureVerifierSetup {
  function test_constructor_ForkedChainDetection() public {
    SignatureQuorumVerifierHelper verifier = new SignatureQuorumVerifierHelper();

    address[] memory testSigners = new address[](1);
    testSigners[0] = vm.addr(PRIVATE_KEY_0);
    verifier.setSignatureConfig(testSigners, 1);

    uint256 expectedChainId = block.chainid;
    uint256 newChainId = expectedChainId + 100000;

    // Simulate chain fork by changing chain ID.
    vm.chainId(newChainId);

    bytes32 testHash = keccak256("test");
    (bytes32 r, bytes32 s) = _signWithV27(PRIVATE_KEY_0, testHash);
    bytes memory signature = abi.encodePacked(r, s);

    vm.expectRevert(abi.encodeWithSelector(SignatureQuorumVerifier.ForkedChain.selector, expectedChainId, newChainId));
    verifier.validateSignatures(testHash, signature);
  }
}
