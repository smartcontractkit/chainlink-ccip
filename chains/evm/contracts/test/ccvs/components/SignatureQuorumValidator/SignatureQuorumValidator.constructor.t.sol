// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {SignatureQuorumValidator} from "../../../../ccvs/components/SignatureQuorumValidator.sol";
import {SignatureQuorumValidatorHelper} from "../../../helpers/SignatureQuorumValidatorHelper.sol";
import {SignatureValidatorSetup} from "./SignatureValidatorSetup.t.sol";

contract SignatureQuorumValidator_constructor is SignatureValidatorSetup {
  function test_constructor_ForkedChainDetection() public {
    SignatureQuorumValidatorHelper verifier = new SignatureQuorumValidatorHelper();

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

    vm.expectRevert(abi.encodeWithSelector(SignatureQuorumValidator.ForkedChain.selector, expectedChainId, newChainId));
    verifier.validateSignatures(testHash, signature);
  }
}
