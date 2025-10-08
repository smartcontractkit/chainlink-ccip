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

    vm.expectEmit();
    emit SignatureQuorumValidator.SignatureConfigSet(newSigners, newThreshold);

    s_sigQuorumVerifier.setSignatureConfig(newSigners, newThreshold);

    (address[] memory actualSigners, uint8 actualThreshold) = s_sigQuorumVerifier.getSignatureConfig();

    assertEq(actualSigners.length, newSigners.length);
    assertEq(actualThreshold, newThreshold);

    for (uint256 i = 0; i < newSigners.length; ++i) {
      assertEq(actualSigners[i], newSigners[i]);
    }
  }

  function test_setSignatureConfig_UpdatesExistingConfig() public {
    // First set initial config.
    address[] memory initialSigners = new address[](2);
    initialSigners[0] = makeAddr("initial1");
    initialSigners[1] = makeAddr("initial2");
    s_sigQuorumVerifier.setSignatureConfig(initialSigners, 1);

    // Update with new config.
    address[] memory newSigners = new address[](3);
    newSigners[0] = makeAddr("new1");
    newSigners[1] = makeAddr("new2");
    newSigners[2] = makeAddr("new3");
    uint8 newThreshold = 2;

    vm.expectEmit();
    emit SignatureQuorumValidator.SignatureConfigSet(newSigners, newThreshold);

    s_sigQuorumVerifier.setSignatureConfig(newSigners, newThreshold);

    (address[] memory actualSigners, uint8 actualThreshold) = s_sigQuorumVerifier.getSignatureConfig();

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

    s_sigQuorumVerifier.setSignatureConfig(signers, threshold);

    (address[] memory actualSigners, uint8 actualThreshold) = s_sigQuorumVerifier.getSignatureConfig();

    assertEq(actualSigners.length, 1);
    assertEq(actualThreshold, 1);
    assertEq(actualSigners[0], signers[0]);
  }

  // Reverts

  function test_setSignatureConfig_RevertWhen_CallerNotOwner() public {
    vm.startPrank(makeAddr("notOwner"));

    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_sigQuorumVerifier.setSignatureConfig(new address[](0), 1);
  }

  function test_setSignatureConfig_RevertWhen_ThresholdZero() public {
    vm.expectRevert(abi.encodeWithSelector(SignatureQuorumValidator.InvalidSignatureConfig.selector));
    s_sigQuorumVerifier.setSignatureConfig(new address[](0), 0);
  }

  function test_setSignatureConfig_RevertWhen_ThresholdExceedsSignerCount() public {
    vm.expectRevert(abi.encodeWithSelector(SignatureQuorumValidator.InvalidSignatureConfig.selector));
    s_sigQuorumVerifier.setSignatureConfig(new address[](0), 3); // threshold > signers.length
  }

  function test_setSignatureConfig_RevertWhen_ZeroAddressSigner() public {
    vm.expectRevert(abi.encodeWithSelector(SignatureQuorumValidator.OracleCannotBeZeroAddress.selector));
    s_sigQuorumVerifier.setSignatureConfig(new address[](1), 1);
  }

  function test_setSignatureConfig_RevertWhen_DuplicateSigners() public {
    address duplicateAddress = makeAddr("duplicate");
    address[] memory signers = new address[](3);
    signers[0] = makeAddr("signer1");
    signers[1] = duplicateAddress;
    signers[2] = duplicateAddress; // Duplicate.

    vm.expectRevert(abi.encodeWithSelector(SignatureQuorumValidator.InvalidSignatureConfig.selector));
    s_sigQuorumVerifier.setSignatureConfig(signers, 2);
  }

  function test_setSignatureConfig_RevertWhen_EmptySignerArray() public {
    address[] memory signers = new address[](0);

    vm.expectRevert(abi.encodeWithSelector(SignatureQuorumValidator.InvalidSignatureConfig.selector));
    s_sigQuorumVerifier.setSignatureConfig(signers, 1);
  }
}
