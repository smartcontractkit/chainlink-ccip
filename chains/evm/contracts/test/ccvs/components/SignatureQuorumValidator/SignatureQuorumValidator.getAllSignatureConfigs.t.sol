// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {SignatureQuorumValidator} from "../../../../ccvs/components/SignatureQuorumValidator.sol";
import {SignatureValidatorSetup} from "./SignatureValidatorSetup.t.sol";

contract SignatureQuorumValidator_getAllSignatureConfigs is SignatureValidatorSetup {
  function test_getAllSignatureConfigs_ReturnsDefaultConfiguration() public view {
    SignatureQuorumValidator.SignatureConfig[] memory configs = s_sigQuorumVerifier.getAllSignatureConfigs();

    assertEq(configs.length, 1);

    (address[] memory defaultSigners, uint8 defaultThreshold) =
      s_sigQuorumVerifier.getSignatureConfig(SOURCE_CHAIN_SELECTOR);

    _assertConfigPresent(configs, SOURCE_CHAIN_SELECTOR, defaultSigners, defaultThreshold);
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

    SignatureQuorumValidator.SignatureConfig[] memory updates = new SignatureQuorumValidator.SignatureConfig[](2);
    updates[0] =
      SignatureQuorumValidator.SignatureConfig({sourceChainSelector: selectorB, signers: signersB, threshold: 2});
    updates[1] =
      SignatureQuorumValidator.SignatureConfig({sourceChainSelector: selectorC, signers: signersC, threshold: 3});

    s_sigQuorumVerifier.applySignatureConfigs(new uint64[](0), updates);

    SignatureQuorumValidator.SignatureConfig[] memory configs = s_sigQuorumVerifier.getAllSignatureConfigs();

    assertEq(configs.length, 3);

    (address[] memory defaultSigners, uint8 defaultThreshold) =
      s_sigQuorumVerifier.getSignatureConfig(SOURCE_CHAIN_SELECTOR);

    _assertConfigPresent(configs, SOURCE_CHAIN_SELECTOR, defaultSigners, defaultThreshold);
    _assertConfigPresent(configs, selectorB, signersB, 2);
    _assertConfigPresent(configs, selectorC, signersC, 3);
  }

  function test_getAllSignatureConfigs_SourceChainSelectorsReflectRemovals() public {
    uint64 selectorB = 900;
    address[] memory signersB = new address[](1);
    signersB[0] = makeAddr("selectorB");

    SignatureQuorumValidator.SignatureConfig[] memory updates = new SignatureQuorumValidator.SignatureConfig[](1);
    updates[0] =
      SignatureQuorumValidator.SignatureConfig({sourceChainSelector: selectorB, signers: signersB, threshold: 1});
    s_sigQuorumVerifier.applySignatureConfigs(new uint64[](0), updates);

    uint64[] memory removals = new uint64[](1);
    removals[0] = selectorB;
    s_sigQuorumVerifier.applySignatureConfigs(removals, new SignatureQuorumValidator.SignatureConfig[](0));

    SignatureQuorumValidator.SignatureConfig[] memory configs = s_sigQuorumVerifier.getAllSignatureConfigs();

    bool selectorFound;
    for (uint256 i; i < configs.length; ++i) {
      if (configs[i].sourceChainSelector == selectorB) {
        selectorFound = true;
        break;
      }
    }
    assertFalse(selectorFound, "selectorB should have been removed");
  }

  function test_getAllSignatureConfigs_EmptyAfterRemovingEverySelector() public {
    uint64[] memory removals = new uint64[](1);
    removals[0] = SOURCE_CHAIN_SELECTOR;
    s_sigQuorumVerifier.applySignatureConfigs(removals, new SignatureQuorumValidator.SignatureConfig[](0));

    SignatureQuorumValidator.SignatureConfig[] memory configs = s_sigQuorumVerifier.getAllSignatureConfigs();

    assertEq(configs.length, 0);
  }
}
