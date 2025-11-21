// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {SignatureQuorumValidator} from "../../../../ccvs/components/SignatureQuorumValidator.sol";
import {SignatureValidatorSetup} from "./SignatureValidatorSetup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract SignatureQuorumValidator_setSignatureConfigs is SignatureValidatorSetup {
  function test_setSignatureConfig() public {
    address[] memory newSigners = new address[](3);
    newSigners[0] = makeAddr("signer1");
    newSigners[1] = makeAddr("signer2");
    newSigners[2] = makeAddr("signer3");
    uint8 newThreshold = 2;
    uint64 sourceChainSelector = 1;

    vm.expectEmit();
    emit SignatureQuorumValidator.SignatureConfigSet(sourceChainSelector, newSigners, newThreshold);

    s_sigQuorumVerifier.setSignatureConfig(sourceChainSelector, newSigners, newThreshold);

    (address[] memory actualSigners, uint8 actualThreshold) =
      s_sigQuorumVerifier.getSignatureConfig(sourceChainSelector);

    assertEq(actualSigners.length, newSigners.length);
    assertEq(actualThreshold, newThreshold);

    for (uint256 i = 0; i < newSigners.length; ++i) {
      assertEq(actualSigners[i], newSigners[i]);
    }
  }

  function test_setSignatureConfigIsPerSource() public {
    uint64 sourceChainSelector0 = 77;
    address[] memory newSignersSource0 = new address[](3);
    newSignersSource0[0] = makeAddr("signer.77.1");
    newSignersSource0[1] = makeAddr("signer.77.2");
    newSignersSource0[2] = makeAddr("signer.77.3");
    uint8 newThreshold0 = 2;

    uint64 sourceChainSelector1 = 333;
    address[] memory newSignersSource1 = new address[](2);
    newSignersSource1[0] = makeAddr("signer.333.1");
    newSignersSource1[1] = makeAddr("signer.333.2");
    uint8 newThreshold1 = 1;

    vm.expectEmit();
    emit SignatureQuorumValidator.SignatureConfigSet(sourceChainSelector0, newSignersSource0, newThreshold0);

    s_sigQuorumVerifier.setSignatureConfig(sourceChainSelector0, newSignersSource0, newThreshold0);

    vm.expectEmit();
    emit SignatureQuorumValidator.SignatureConfigSet(sourceChainSelector1, newSignersSource1, newThreshold1);

    s_sigQuorumVerifier.setSignatureConfig(sourceChainSelector1, newSignersSource1, newThreshold1);

    // Check source 0 signers.
    (address[] memory actualSigners0, uint8 actualThreshold0) =
      s_sigQuorumVerifier.getSignatureConfig(sourceChainSelector0);

    assertEq(actualSigners0.length, newSignersSource0.length);
    assertEq(actualThreshold0, newThreshold0);

    for (uint256 i = 0; i < newSignersSource0.length; ++i) {
      assertEq(actualSigners0[i], newSignersSource0[i]);
    }

    // Check source 1 signers.
    (address[] memory actualSigners1, uint8 actualThreshold1) =
      s_sigQuorumVerifier.getSignatureConfig(sourceChainSelector1);

    assertEq(actualSigners1.length, newSignersSource1.length);
    assertEq(actualThreshold1, newThreshold1);

    for (uint256 i = 0; i < newSignersSource1.length; ++i) {
      assertEq(actualSigners1[i], newSignersSource1[i]);
    }

    // Check that there is no cross pollination between the signer sets.
    for (uint256 i = 0; i < newSignersSource0.length; ++i) {
      for (uint256 j = 0; j < newSignersSource1.length; ++j) {
        assertTrue(actualSigners0[i] != newSignersSource1[j]);
      }
    }
  }

  function test_setSignatureConfig_UpdatesExistingConfig() public {
    uint64 sourceChainSelector = 1;
    // First set initial config.
    address[] memory initialSigners = new address[](2);
    initialSigners[0] = makeAddr("initial1");
    initialSigners[1] = makeAddr("initial2");
    s_sigQuorumVerifier.setSignatureConfig(sourceChainSelector, initialSigners, 1);

    // Update with new config.
    address[] memory newSigners = new address[](3);
    newSigners[0] = makeAddr("new1");
    newSigners[1] = makeAddr("new2");
    newSigners[2] = makeAddr("new3");
    uint8 newThreshold = 2;

    vm.expectEmit();
    emit SignatureQuorumValidator.SignatureConfigSet(sourceChainSelector, newSigners, newThreshold);

    s_sigQuorumVerifier.setSignatureConfig(sourceChainSelector, newSigners, newThreshold);

    (address[] memory actualSigners, uint8 actualThreshold) =
      s_sigQuorumVerifier.getSignatureConfig(sourceChainSelector);

    assertEq(actualSigners.length, 3);
    assertEq(actualThreshold, 2);

    // Verify old signers are replaced.
    for (uint256 i = 0; i < newSigners.length; ++i) {
      assertEq(actualSigners[i], newSigners[i]);
    }
  }

  function test_setSignatureConfig_SingleSignerThresholdOne() public {
    address[] memory signers = new address[](1);
    signers[0] = makeAddr("soloSigner");
    uint8 threshold = 1;
    uint64 sourceChainSelector = 1;

    s_sigQuorumVerifier.setSignatureConfig(1, signers, threshold);

    (address[] memory actualSigners, uint8 actualThreshold) =
      s_sigQuorumVerifier.getSignatureConfig(sourceChainSelector);

    assertEq(actualSigners.length, 1);
    assertEq(actualThreshold, 1);
    assertEq(actualSigners[0], signers[0]);
  }

  // Reverts

  function test_setSignatureConfig_RevertWhen_CallerNotOwner() public {
    vm.startPrank(makeAddr("notOwner"));

    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_sigQuorumVerifier.setSignatureConfig(0, new address[](0), 1);
  }

  function test_setSignatureConfig_RevertWhen_ThresholdZero() public {
    vm.expectRevert(abi.encodeWithSelector(SignatureQuorumValidator.InvalidSignatureConfig.selector));
    s_sigQuorumVerifier.setSignatureConfig(0, new address[](0), 0);
  }

  function test_setSignatureConfig_RevertWhen_ThresholdExceedsSignerCount() public {
    vm.expectRevert(abi.encodeWithSelector(SignatureQuorumValidator.InvalidSignatureConfig.selector));
    s_sigQuorumVerifier.setSignatureConfig(0, new address[](0), 3); // threshold > signers.length
  }

  function test_setSignatureConfig_RevertWhen_ZeroAddressSigner() public {
    vm.expectRevert(abi.encodeWithSelector(SignatureQuorumValidator.OracleCannotBeZeroAddress.selector));
    s_sigQuorumVerifier.setSignatureConfig(0, new address[](1), 1);
  }

  function test_setSignatureConfig_RevertWhen_DuplicateSigners() public {
    address duplicateAddress = makeAddr("duplicate");
    address[] memory signers = new address[](3);
    signers[0] = makeAddr("signer1");
    signers[1] = duplicateAddress;
    signers[2] = duplicateAddress; // Duplicate.

    vm.expectRevert(abi.encodeWithSelector(SignatureQuorumValidator.InvalidSignatureConfig.selector));
    s_sigQuorumVerifier.setSignatureConfig(0, signers, 2);
  }

  function test_setSignatureConfig_RevertWhen_EmptySignerArray() public {
    address[] memory signers = new address[](0);

    vm.expectRevert(abi.encodeWithSelector(SignatureQuorumValidator.InvalidSignatureConfig.selector));
    s_sigQuorumVerifier.setSignatureConfig(0, signers, 1);
  }
}
