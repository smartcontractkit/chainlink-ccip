// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {SignatureQuorumValidatorHelper} from "../../../helpers/SignatureQuorumValidatorHelper.sol";
import {SignatureValidatorSetup} from "./SignatureValidatorSetup.t.sol";

contract SignatureQuorumValidator_getSignatureConfig is SignatureValidatorSetup {
  function test_getSignatureConfig_InitialState() public view {
    (address[] memory signers, uint8 threshold) = s_sigQuorumVerifier.getSignatureConfig();

    assertEq(signers.length, 4);
    assertEq(threshold, 1);

    for (uint256 i = 0; i < s_validSigners.length; ++i) {
      assertEq(signers[i], s_validSigners[i]);
    }
  }

  function test_getSignatureConfig_AfterUpdate() public {
    address[] memory newSigners = new address[](3);
    newSigners[0] = makeAddr("newSigner1");
    newSigners[1] = makeAddr("newSigner2");
    newSigners[2] = makeAddr("newSigner3");
    uint8 newThreshold = 2;

    s_sigQuorumVerifier.setSignatureConfig(newSigners, newThreshold);

    (address[] memory actualSigners, uint8 actualThreshold) = s_sigQuorumVerifier.getSignatureConfig();

    assertEq(actualSigners.length, 3);
    assertEq(actualThreshold, 2);

    for (uint256 i = 0; i < newSigners.length; ++i) {
      assertEq(actualSigners[i], newSigners[i]);
    }
  }

  function test_getSignatureConfig_EmptyConfiguration() public {
    // Deploy new verifier with no initial setup.
    SignatureQuorumValidatorHelper newVerifier = new SignatureQuorumValidatorHelper();

    (address[] memory signers, uint8 threshold) = newVerifier.getSignatureConfig();

    assertEq(signers.length, 0);
    assertEq(threshold, 0);
  }

  function test_getSignatureConfig_SingleSigner() public {
    address[] memory singleSigner = new address[](1);
    singleSigner[0] = makeAddr("soloSigner");

    s_sigQuorumVerifier.setSignatureConfig(singleSigner, 1);

    (address[] memory signers, uint8 threshold) = s_sigQuorumVerifier.getSignatureConfig();

    assertEq(signers.length, 1);
    assertEq(threshold, 1);
    assertEq(signers[0], singleSigner[0]);
  }

  function test_getSignatureConfig_LargeSignerSet() public {
    address[] memory largeSignerSet = new address[](10);
    for (uint256 i = 0; i < 10; ++i) {
      largeSignerSet[i] = makeAddr(string(abi.encodePacked("signer", i)));
    }

    s_sigQuorumVerifier.setSignatureConfig(largeSignerSet, 7);

    (address[] memory signers, uint8 threshold) = s_sigQuorumVerifier.getSignatureConfig();

    assertEq(signers.length, 10);
    assertEq(threshold, 7);

    for (uint256 i = 0; i < largeSignerSet.length; ++i) {
      assertEq(signers[i], largeSignerSet[i]);
    }
  }

  function test_getSignatureConfig_MultipleUpdates() public {
    // First update.
    address[] memory firstSigners = new address[](2);
    firstSigners[0] = makeAddr("first1");
    firstSigners[1] = makeAddr("first2");
    s_sigQuorumVerifier.setSignatureConfig(firstSigners, 1);

    (address[] memory signers1, uint8 threshold1) = s_sigQuorumVerifier.getSignatureConfig();
    assertEq(signers1.length, 2);
    assertEq(threshold1, 1);

    // Second update.
    address[] memory secondSigners = new address[](3);
    secondSigners[0] = makeAddr("second1");
    secondSigners[1] = makeAddr("second2");
    secondSigners[2] = makeAddr("second3");
    s_sigQuorumVerifier.setSignatureConfig(secondSigners, 3);

    (address[] memory signers2, uint8 threshold2) = s_sigQuorumVerifier.getSignatureConfig();
    assertEq(signers2.length, 3);
    assertEq(threshold2, 3);

    // Verify final state matches latest update.
    for (uint256 i = 0; i < secondSigners.length; ++i) {
      assertEq(signers2[i], secondSigners[i]);
    }
  }
}
