// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {BaseVerifier} from "../../../ccvs/components/BaseVerifier.sol";
import {FinalityCodec} from "../../../libraries/FinalityCodec.sol";
import {CommitteeVerifierSetup} from "./CommitteeVerifierSetup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract CommitteeVerifier_setAllowedFinalityConfig is CommitteeVerifierSetup {
  function test_setAllowedFinalityConfig() public {
    bytes4 newFinality = FinalityCodec._encodeBlockDepth(42);

    vm.expectEmit();
    emit BaseVerifier.FinalityConfigSet(newFinality);

    s_committeeVerifier.setAllowedFinalityConfig(newFinality);
    assertEq(s_committeeVerifier.getAllowedFinalityConfig(), newFinality);
  }

  function test_setAllowedFinalityConfig_RevertWhen_OnlyCallableByOwner() public {
    vm.startPrank(STRANGER);

    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_committeeVerifier.setAllowedFinalityConfig(FinalityCodec._encodeBlockDepth(1));
  }
}
