// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {SignatureQuorumValidator} from "../../../../ccvs/components/SignatureQuorumValidator.sol";
import {SignatureValidatorSetup} from "./SignatureValidatorSetup.t.sol";

contract SignatureQuorumValidator_getAllSignatureConfigs is SignatureValidatorSetup {
  function test_getAllSignatureConfigs_ReturnsDefaultConfiguration() public view {
    (uint64[] memory selectors, address[][] memory signerSets, uint8[] memory thresholds) =
      s_sigQuorumVerifier.getAllSignatureConfigs();

    assertEq(selectors.length, 1);
    assertEq(signerSets.length, 1);
    assertEq(thresholds.length, 1);

    (address[] memory defaultSigners, uint8 defaultThreshold) =
      s_sigQuorumVerifier.getSignatureConfig(SOURCE_CHAIN_SELECTOR);

    _assertConfigPresent(selectors, signerSets, thresholds, SOURCE_CHAIN_SELECTOR, defaultSigners, defaultThreshold);
  }

  function test_getAllSignatureConfigs_MultipleSelectorsReturned() public {
    uint64 selectorB = 800;
    uint64 selectorC = 801;

    address[] memory signersB = new address[](2);
    signersB[0] = makeAddr("selectorB-1");
    signersB[1] = makeAddr("selectorB-2");

    address[] memory signersC = new address[](3);
    signersC[0] = makeAddr("selectorC-1");
    signersC[1] = makeAddr("selectorC-2");
    signersC[2] = makeAddr("selectorC-3");

    SignatureQuorumValidator.SignersUpdate[] memory updates = new SignatureQuorumValidator.SignersUpdate[](2);
    updates[0] =
      SignatureQuorumValidator.SignersUpdate({sourceChainSelector: selectorB, signers: signersB, threshold: 2});
    updates[1] =
      SignatureQuorumValidator.SignersUpdate({sourceChainSelector: selectorC, signers: signersC, threshold: 3});

    s_sigQuorumVerifier.applySignersUpdates(new uint64[](0), updates);

    (uint64[] memory selectors, address[][] memory signerSets, uint8[] memory thresholds) =
      s_sigQuorumVerifier.getAllSignatureConfigs();

    assertEq(selectors.length, 3);
    assertEq(signerSets.length, 3);
    assertEq(thresholds.length, 3);

    (address[] memory defaultSigners, uint8 defaultThreshold) =
      s_sigQuorumVerifier.getSignatureConfig(SOURCE_CHAIN_SELECTOR);

    _assertConfigPresent(selectors, signerSets, thresholds, SOURCE_CHAIN_SELECTOR, defaultSigners, defaultThreshold);
    _assertConfigPresent(selectors, signerSets, thresholds, selectorB, signersB, 2);
    _assertConfigPresent(selectors, signerSets, thresholds, selectorC, signersC, 3);
  }

  function test_getAllSignatureConfigs_SourceChainSelectorsReflectRemovals() public {
    uint64 selectorB = 900;
    address[] memory signersB = new address[](1);
    signersB[0] = makeAddr("selectorB");

    SignatureQuorumValidator.SignersUpdate[] memory updates = new SignatureQuorumValidator.SignersUpdate[](1);
    updates[0] =
      SignatureQuorumValidator.SignersUpdate({sourceChainSelector: selectorB, signers: signersB, threshold: 1});
    s_sigQuorumVerifier.applySignersUpdates(new uint64[](0), updates);

    uint64[] memory removals = new uint64[](1);
    removals[0] = selectorB;
    s_sigQuorumVerifier.applySignersUpdates(removals, new SignatureQuorumValidator.SignersUpdate[](0));

    (uint64[] memory selectors,,) = s_sigQuorumVerifier.getAllSignatureConfigs();

    bool selectorFound;
    for (uint256 i; i < selectors.length; ++i) {
      if (selectors[i] == selectorB) {
        selectorFound = true;
        break;
      }
    }
    assertFalse(selectorFound, "selectorB should have been removed");
  }

  function test_getAllSignatureConfigs_EmptyAfterRemovingEverySelector() public {
    uint64[] memory removals = new uint64[](1);
    removals[0] = SOURCE_CHAIN_SELECTOR;
    s_sigQuorumVerifier.applySignersUpdates(removals, new SignatureQuorumValidator.SignersUpdate[](0));

    (uint64[] memory selectors, address[][] memory signerSets, uint8[] memory thresholds) =
      s_sigQuorumVerifier.getAllSignatureConfigs();

    assertEq(selectors.length, 0);
    assertEq(signerSets.length, 0);
    assertEq(thresholds.length, 0);
  }
}
