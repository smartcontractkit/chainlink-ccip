// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {BaseVerifier} from "../../../ccvs/components/BaseVerifier.sol";
import {FinalityCodec} from "../../../libraries/FinalityCodec.sol";
import {SuccinctZKVerifierSetup} from "./SuccinctZKVerifierSetup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract SuccinctZKVerifier_setAllowedFinalityConfig is SuccinctZKVerifierSetup {
  function test_setAllowedFinalityConfig() public {
    bytes4 allowedFinality = FinalityCodec.WAIT_FOR_FINALITY_FLAG;

    vm.expectEmit();
    emit BaseVerifier.FinalityConfigSet(allowedFinality);

    s_zkVerifier.setAllowedFinalityConfig(allowedFinality);

    assertEq(allowedFinality, s_zkVerifier.getAllowedFinalityConfig());
  }

  // Reverts

  function test_setAllowedFinalityConfig_RevertWhen_OnlyCallableByOwner() public {
    vm.stopPrank();
    vm.startPrank(STRANGER);

    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_zkVerifier.setAllowedFinalityConfig(FinalityCodec.WAIT_FOR_FINALITY_FLAG);
  }
}
