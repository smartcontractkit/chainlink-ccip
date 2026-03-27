// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {FinalityCodec} from "../../../libraries/FinalityCodec.sol";
import {TokenPool} from "../../../pools/TokenPool.sol";

import {AdvancedPoolHooksSetup} from "../AdvancedPoolHooks/AdvancedPoolHooksSetup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract TokenPool_setAllowedFinalityConfig is AdvancedPoolHooksSetup {
  function test_setAllowedFinalityConfig() public {
    bytes4 newMinFinality = FinalityCodec._encodeBlockDepth(42);
    vm.expectEmit();
    emit TokenPool.FinalityConfigSet(newMinFinality);
    s_tokenPool.setAllowedFinalityConfig(newMinFinality);
    assertEq(s_tokenPool.getAllowedFinalityConfig(), newMinFinality);
  }

  // Reverts
  function test_setAllowedFinalityConfig_RevertWhen_OnlyCallableByOwner() public {
    bytes4 newMinFinality = FinalityCodec._encodeBlockDepth(42);
    changePrank(STRANGER);
    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_tokenPool.setAllowedFinalityConfig(newMinFinality);
  }
}
