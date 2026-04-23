// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {BaseVerifier} from "../../../../ccvs/components/BaseVerifier.sol";
import {FinalityCodec} from "../../../../libraries/FinalityCodec.sol";
import {BaseVerifierSetup} from "./BaseVerifierSetup.t.sol";

contract BaseVerifier_setAllowedFinalityConfig is BaseVerifierSetup {
  function test_setAllowedFinalityConfig() public {
    bytes4 newFinality = FinalityCodec._encodeBlockDepth(42);

    vm.expectEmit();
    emit BaseVerifier.FinalityConfigSet(newFinality);

    s_baseVerifier.setAllowedFinalityConfig(newFinality);
    assertEq(s_baseVerifier.getAllowedFinalityConfig(), newFinality);
  }

  function test_setAllowedFinalityConfig_WaitForSafe() public {
    s_baseVerifier.setAllowedFinalityConfig(FinalityCodec.WAIT_FOR_SAFE_FLAG);
    assertEq(s_baseVerifier.getAllowedFinalityConfig(), FinalityCodec.WAIT_FOR_SAFE_FLAG);
  }

  function test_setAllowedFinalityConfig_WaitForFinality() public {
    s_baseVerifier.setAllowedFinalityConfig(FinalityCodec._encodeBlockDepth(10));
    s_baseVerifier.setAllowedFinalityConfig(FinalityCodec.WAIT_FOR_FINALITY_FLAG);
    assertEq(s_baseVerifier.getAllowedFinalityConfig(), FinalityCodec.WAIT_FOR_FINALITY_FLAG);
  }

  function test_getAllowedFinalityConfig_DefaultIsZero() public view {
    assertEq(s_baseVerifier.getAllowedFinalityConfig(), bytes4(0));
  }
}
