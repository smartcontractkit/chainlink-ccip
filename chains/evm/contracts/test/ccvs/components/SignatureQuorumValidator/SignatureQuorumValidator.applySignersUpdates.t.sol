// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {SignatureQuorumValidator} from "../../../../ccvs/components/SignatureQuorumValidator.sol";
import {SignatureValidatorSetup} from "./SignatureValidatorSetup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract SignatureQuorumValidator_applySignersUpdates is SignatureValidatorSetup {
  function test_applySignersUpdates() public {
    address[] memory newSigners = new address[](3);
    newSigners[0] = makeAddr("signer1");
    newSigners[1] = makeAddr("signer2");
    newSigners[2] = makeAddr("signer3");
    uint8 newThreshold = 2;
    uint64 sourceChainSelector = 1;

    vm.expectEmit();
    emit SignatureQuorumValidator.SignatureConfigSet(sourceChainSelector, newSigners, newThreshold);

    SignatureQuorumValidator.SignersUpdate[] memory updates =
      _createUpdate(sourceChainSelector, newSigners, newThreshold);
    s_sigQuorumVerifier.applySignersUpdates(new uint64[](0), updates);

    (address[] memory actualSigners, uint8 actualThreshold) =
      s_sigQuorumVerifier.getSignatureConfig(sourceChainSelector);

    assertEq(actualSigners.length, newSigners.length);
    assertEq(actualThreshold, newThreshold);

    for (uint256 i = 0; i < newSigners.length; ++i) {
      assertEq(actualSigners[i], newSigners[i]);
    }
  }

  function test_applySignersUpdates_IsPerSource() public {
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

    SignatureQuorumValidator.SignersUpdate[] memory updates =
      _createUpdate(sourceChainSelector0, newSignersSource0, newThreshold0);
    s_sigQuorumVerifier.applySignersUpdates(new uint64[](0), updates);

    vm.expectEmit();
    emit SignatureQuorumValidator.SignatureConfigSet(sourceChainSelector1, newSignersSource1, newThreshold1);

    updates = _createUpdate(sourceChainSelector1, newSignersSource1, newThreshold1);
    s_sigQuorumVerifier.applySignersUpdates(new uint64[](0), updates);

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

  function test_applySignersUpdates_UpdatesExistingConfig() public {
    uint64 sourceChainSelector = 1;
    // First set initial config.
    address[] memory initialSigners = new address[](2);
    initialSigners[0] = makeAddr("initial1");
    initialSigners[1] = makeAddr("initial2");
    SignatureQuorumValidator.SignersUpdate[] memory updates = _createUpdate(sourceChainSelector, initialSigners, 1);
    s_sigQuorumVerifier.applySignersUpdates(new uint64[](0), updates);

    // Update with new config.
    address[] memory newSigners = new address[](3);
    newSigners[0] = makeAddr("new1");
    newSigners[1] = makeAddr("new2");
    newSigners[2] = makeAddr("new3");
    uint8 newThreshold = 2;

    vm.expectEmit();
    emit SignatureQuorumValidator.SignatureConfigSet(sourceChainSelector, newSigners, newThreshold);

    updates = _createUpdate(sourceChainSelector, newSigners, newThreshold);
    s_sigQuorumVerifier.applySignersUpdates(new uint64[](0), updates);

    (address[] memory actualSigners, uint8 actualThreshold) =
      s_sigQuorumVerifier.getSignatureConfig(sourceChainSelector);

    assertEq(actualSigners.length, 3);
    assertEq(actualThreshold, 2);

    // Verify old signers are replaced.
    for (uint256 i = 0; i < newSigners.length; ++i) {
      assertEq(actualSigners[i], newSigners[i]);
    }
  }

  function test_applySignersUpdates_SingleSignerThresholdOne() public {
    address[] memory signers = new address[](1);
    signers[0] = makeAddr("soloSigner");
    uint8 threshold = 1;
    uint64 sourceChainSelector = 1;

    SignatureQuorumValidator.SignersUpdate[] memory updates = _createUpdate(1, signers, threshold);
    s_sigQuorumVerifier.applySignersUpdates(new uint64[](0), updates);

    (address[] memory actualSigners, uint8 actualThreshold) =
      s_sigQuorumVerifier.getSignatureConfig(sourceChainSelector);

    assertEq(actualSigners.length, 1);
    assertEq(actualThreshold, 1);
    assertEq(actualSigners[0], signers[0]);
  }

  function test_applySignersUpdates_RemovalTakesPriorityWithinOneCall() public {
    uint64 selector = 555;
    address[] memory signers = new address[](2);
    signers[0] = makeAddr("selector555-signer1");
    signers[1] = makeAddr("selector555-signer2");

    SignatureQuorumValidator.SignersUpdate[] memory updates = _createUpdate(selector, signers, 2);

    uint64[] memory removals = new uint64[](1);
    removals[0] = selector;

    vm.expectEmit();
    emit SignatureQuorumValidator.SignatureConfigSet(selector, signers, 2);

    vm.expectEmit();
    emit SignatureQuorumValidator.SignatureConfigSet(selector, new address[](0), 0);

    s_sigQuorumVerifier.applySignersUpdates(removals, updates);

    (address[] memory actualSigners, uint8 threshold) = s_sigQuorumVerifier.getSignatureConfig(selector);
    assertEq(actualSigners.length, 0);
    assertEq(threshold, 0);
  }

  function test_applySignersUpdates_AllowsNoops() public {
    (address[] memory beforeSigners, uint8 beforeThreshold) =
      s_sigQuorumVerifier.getSignatureConfig(SOURCE_CHAIN_SELECTOR);

    // Pure noop: no updates and no removals.
    s_sigQuorumVerifier.applySignersUpdates(new uint64[](0), new SignatureQuorumValidator.SignersUpdate[](0));

    // Updating with the exact same config should also be a noop.
    SignatureQuorumValidator.SignersUpdate[] memory identicalUpdate =
      _createUpdate(SOURCE_CHAIN_SELECTOR, beforeSigners, beforeThreshold);
    s_sigQuorumVerifier.applySignersUpdates(new uint64[](0), identicalUpdate);

    // Removing a selector that was never configured should be a noop.
    uint64[] memory nonexistentRemoval = new uint64[](1);
    nonexistentRemoval[0] = 999_999;
    s_sigQuorumVerifier.applySignersUpdates(nonexistentRemoval, new SignatureQuorumValidator.SignersUpdate[](0));

    (address[] memory afterSigners, uint8 afterThreshold) =
      s_sigQuorumVerifier.getSignatureConfig(SOURCE_CHAIN_SELECTOR);

    assertEq(beforeSigners.length, afterSigners.length);
    assertEq(beforeThreshold, afterThreshold);
    for (uint256 i; i < beforeSigners.length; ++i) {
      assertEq(beforeSigners[i], afterSigners[i]);
    }
  }

  function test_applySignersUpdates_RemovalEmitsEventAndClearsConfig() public {
    uint64[] memory removals = new uint64[](1);
    removals[0] = SOURCE_CHAIN_SELECTOR;

    vm.expectEmit();
    emit SignatureQuorumValidator.SignatureConfigSet(SOURCE_CHAIN_SELECTOR, new address[](0), 0);

    s_sigQuorumVerifier.applySignersUpdates(removals, new SignatureQuorumValidator.SignersUpdate[](0));

    (address[] memory signers, uint8 threshold) = s_sigQuorumVerifier.getSignatureConfig(SOURCE_CHAIN_SELECTOR);
    assertEq(signers.length, 0);
    assertEq(threshold, 0);
  }

  function test_applySignersUpdates_MultipleChainsSingleCall_NoCrossPollination() public {
    uint64 selectorA = 700;
    uint64 selectorB = 900;

    address[] memory signersA = new address[](3);
    signersA[0] = makeAddr("selectorA-1");
    signersA[1] = makeAddr("selectorA-2");
    signersA[2] = makeAddr("selectorA-3");

    address[] memory signersB = new address[](2);
    signersB[0] = makeAddr("selectorB-1");
    signersB[1] = makeAddr("selectorB-2");

    SignatureQuorumValidator.SignersUpdate[] memory updates = new SignatureQuorumValidator.SignersUpdate[](2);
    updates[0] =
      SignatureQuorumValidator.SignersUpdate({sourceChainSelector: selectorA, signers: signersA, threshold: 3});
    updates[1] =
      SignatureQuorumValidator.SignersUpdate({sourceChainSelector: selectorB, signers: signersB, threshold: 1});

    s_sigQuorumVerifier.applySignersUpdates(new uint64[](0), updates);

    (address[] memory actualSignersA, uint8 thresholdA) = s_sigQuorumVerifier.getSignatureConfig(selectorA);
    (address[] memory actualSignersB, uint8 thresholdB) = s_sigQuorumVerifier.getSignatureConfig(selectorB);

    assertEq(thresholdA, 3);
    assertEq(thresholdB, 1);
    _assertAddressArraysEqual(signersA, actualSignersA);
    _assertAddressArraysEqual(signersB, actualSignersB);
  }

  function test_getAllSignatureConfigs_ReturnsExpectedConfigurations() public {
    (address[] memory defaultSigners, uint8 defaultThreshold) =
      s_sigQuorumVerifier.getSignatureConfig(SOURCE_CHAIN_SELECTOR);

    uint64 selectorB = 800;
    address[] memory signersB = new address[](2);
    signersB[0] = makeAddr("selectorB-1");
    signersB[1] = makeAddr("selectorB-2");

    uint64 selectorC = 801;
    address[] memory signersC = new address[](1);
    signersC[0] = makeAddr("selectorC-1");

    SignatureQuorumValidator.SignersUpdate[] memory updates = new SignatureQuorumValidator.SignersUpdate[](2);
    updates[0] =
      SignatureQuorumValidator.SignersUpdate({sourceChainSelector: selectorB, signers: signersB, threshold: 2});
    updates[1] =
      SignatureQuorumValidator.SignersUpdate({sourceChainSelector: selectorC, signers: signersC, threshold: 1});

    s_sigQuorumVerifier.applySignersUpdates(new uint64[](0), updates);

    (uint64[] memory selectors, address[][] memory signerSets, uint8[] memory thresholds) =
      s_sigQuorumVerifier.getAllSignatureConfigs();

    assertEq(selectors.length, 3);
    assertEq(signerSets.length, 3);
    assertEq(thresholds.length, 3);

    _assertConfigPresent(selectors, signerSets, thresholds, SOURCE_CHAIN_SELECTOR, defaultSigners, defaultThreshold);
    _assertConfigPresent(selectors, signerSets, thresholds, selectorB, signersB, 2);
    _assertConfigPresent(selectors, signerSets, thresholds, selectorC, signersC, 1);
  }

  function test_applySignersUpdates_LastUpdateWinsBeforeRemoval() public {
    uint64 selector = 8080;
    address[] memory firstSigners = new address[](1);
    firstSigners[0] = makeAddr("first");

    address[] memory secondSigners = new address[](2);
    secondSigners[0] = makeAddr("second1");
    secondSigners[1] = makeAddr("second2");

    SignatureQuorumValidator.SignersUpdate[] memory updates = new SignatureQuorumValidator.SignersUpdate[](2);
    updates[0] =
      SignatureQuorumValidator.SignersUpdate({sourceChainSelector: selector, signers: firstSigners, threshold: 1});
    updates[1] =
      SignatureQuorumValidator.SignersUpdate({sourceChainSelector: selector, signers: secondSigners, threshold: 2});

    s_sigQuorumVerifier.applySignersUpdates(new uint64[](0), updates);

    (address[] memory signers, uint8 threshold) = s_sigQuorumVerifier.getSignatureConfig(selector);
    _assertAddressArraysEqual(secondSigners, signers);
    assertEq(threshold, 2);
  }

  function test_getAllSignatureConfigs_EmptyAfterAllRemovals() public {
    // Configure two additional selectors.
    uint64 selectorB = 8800;
    uint64 selectorC = 9900;

    SignatureQuorumValidator.SignersUpdate[] memory updates = new SignatureQuorumValidator.SignersUpdate[](2);
    updates[0] = SignatureQuorumValidator.SignersUpdate({
      sourceChainSelector: selectorB,
      signers: _arrayOf(makeAddr("bSigner")),
      threshold: 1
    });
    updates[1] = SignatureQuorumValidator.SignersUpdate({
      sourceChainSelector: selectorC,
      signers: _arrayOf(makeAddr("cSigner")),
      threshold: 1
    });
    s_sigQuorumVerifier.applySignersUpdates(new uint64[](0), updates);

    // Remove every configured selector (default + the two new ones).
    uint64[] memory removals = new uint64[](3);
    removals[0] = SOURCE_CHAIN_SELECTOR;
    removals[1] = selectorB;
    removals[2] = selectorC;
    s_sigQuorumVerifier.applySignersUpdates(removals, new SignatureQuorumValidator.SignersUpdate[](0));

    (uint64[] memory selectors, address[][] memory signerSets, uint8[] memory thresholds) =
      s_sigQuorumVerifier.getAllSignatureConfigs();

    assertEq(selectors.length, 0);
    assertEq(signerSets.length, 0);
    assertEq(thresholds.length, 0);
  }

  // Reverts

  function test_applySignersUpdates_RevertWhen_CallerNotOwner() public {
    SignatureQuorumValidator.SignersUpdate[] memory updates = _createUpdate(0, new address[](0), 1);
    vm.startPrank(makeAddr("notOwner"));

    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_sigQuorumVerifier.applySignersUpdates(new uint64[](0), updates);
  }

  function test_applySignersUpdates_RevertWhen_ThresholdZero() public {
    SignatureQuorumValidator.SignersUpdate[] memory updates = _createUpdate(0, new address[](0), 0);
    vm.expectRevert(abi.encodeWithSelector(SignatureQuorumValidator.InvalidSignatureConfig.selector));
    s_sigQuorumVerifier.applySignersUpdates(new uint64[](0), updates);
  }

  function test_applySignersUpdates_RevertWhen_ThresholdExceedsSignerCount() public {
    SignatureQuorumValidator.SignersUpdate[] memory updates = _createUpdate(0, new address[](0), 3);
    vm.expectRevert(abi.encodeWithSelector(SignatureQuorumValidator.InvalidSignatureConfig.selector));
    s_sigQuorumVerifier.applySignersUpdates(new uint64[](0), updates); // threshold > signers.length
  }

  function test_applySignersUpdates_RevertWhen_ZeroAddressSigner() public {
    SignatureQuorumValidator.SignersUpdate[] memory updates = _createUpdate(0, new address[](1), 1);
    vm.expectRevert(abi.encodeWithSelector(SignatureQuorumValidator.SignerCannotBeZeroAddress.selector));
    s_sigQuorumVerifier.applySignersUpdates(new uint64[](0), updates);
  }

  function test_applySignersUpdates_RevertWhen_DuplicateSigners() public {
    address duplicateAddress = makeAddr("duplicate");
    address[] memory signers = new address[](3);
    signers[0] = makeAddr("signer1");
    signers[1] = duplicateAddress;
    signers[2] = duplicateAddress; // Duplicate.

    SignatureQuorumValidator.SignersUpdate[] memory updates = _createUpdate(0, signers, 2);

    vm.expectRevert(abi.encodeWithSelector(SignatureQuorumValidator.InvalidSignatureConfig.selector));
    s_sigQuorumVerifier.applySignersUpdates(new uint64[](0), updates);
  }

  function test_applySignersUpdates_RevertWhen_EmptySignerArray() public {
    address[] memory signers = new address[](0);
    SignatureQuorumValidator.SignersUpdate[] memory updates = _createUpdate(0, signers, 1);

    vm.expectRevert(abi.encodeWithSelector(SignatureQuorumValidator.InvalidSignatureConfig.selector));
    s_sigQuorumVerifier.applySignersUpdates(new uint64[](0), updates);
  }
}
