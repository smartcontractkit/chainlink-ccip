// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IAdvancedPoolHooks} from "../../../../interfaces/IAdvancedPoolHooks.sol";
import {IPolicyEngine} from "@chainlink/policy-management/interfaces/IPolicyEngine.sol";

import {FinalityCodec} from "../../../../libraries/FinalityCodec.sol";
import {MultiTokenPoolHooks} from "../../../../pools/MultiTokenPool/MultiTokenPoolHooks.sol";
import {MockPolicyEngine} from "../../../mocks/MockPolicyEngine.sol";
import {MultiTokenPoolHooksSetup} from "./MultiTokenPoolHooksSetup.t.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

contract MultiTokenPoolHooks_preflightCheck is MultiTokenPoolHooksSetup {
  function test_preflightCheck() public {
    s_hooks.preflightCheck(s_lockOrBurnIn, FinalityCodec.WAIT_FOR_FINALITY_FLAG, "", 1e18);
  }

  function test_preflightCheck_AllowedSender() public {
    MultiTokenPoolHooks hooks = _deployHooksWithAllowList();

    hooks.preflightCheck(s_lockOrBurnIn, FinalityCodec.WAIT_FOR_FINALITY_FLAG, "", 1e18);
  }

  function test_preflightCheck_WithPolicyEngine() public {
    s_hooks.setPolicyEngine(s_policyEngine);
    bytes4 requestedFinalityConfig = FinalityCodec._encodeBlockDepth(5);
    bytes memory tokenArgs = abi.encode("custom token args");
    uint256 amountPostFee = 9e17;

    s_hooks.preflightCheck(s_lockOrBurnIn, requestedFinalityConfig, tokenArgs, amountPostFee);

    IPolicyEngine.Payload memory payload = s_policyEngine.getLastPayload();
    // The selector matches IAdvancedPoolHooks, so one policy engine extractor serves both kinds of hooks.
    assertEq(IAdvancedPoolHooks.preflightCheck.selector, payload.selector);
    assertEq(OWNER, payload.sender);
    assertEq(abi.encode(s_lockOrBurnIn, requestedFinalityConfig, tokenArgs, amountPostFee), payload.data);
    assertEq(tokenArgs, payload.context);
  }

  // Reverts

  function test_preflightCheck_RevertWhen_UnauthorizedCaller() public {
    vm.startPrank(STRANGER);

    vm.expectRevert(abi.encodeWithSelector(AuthorizedCallers.UnauthorizedCaller.selector, STRANGER));
    s_hooks.preflightCheck(s_lockOrBurnIn, FinalityCodec.WAIT_FOR_FINALITY_FLAG, "", 1e18);
  }

  function test_preflightCheck_RevertWhen_SenderNotAllowed() public {
    MultiTokenPoolHooks hooks = _deployHooksWithAllowList();
    s_lockOrBurnIn.originalSender = STRANGER;

    vm.expectRevert(abi.encodeWithSelector(MultiTokenPoolHooks.SenderNotAllowed.selector, STRANGER));
    hooks.preflightCheck(s_lockOrBurnIn, FinalityCodec.WAIT_FOR_FINALITY_FLAG, "", 1e18);
  }

  function test_preflightCheck_RevertWhen_PolicyEngineRejects() public {
    s_hooks.setPolicyEngine(s_policyEngine);
    s_policyEngine.setShouldRevert(true, "policy rejected");

    vm.expectRevert(abi.encodeWithSelector(MockPolicyEngine.MockPolicyEngineRejection.selector, "policy rejected"));
    s_hooks.preflightCheck(s_lockOrBurnIn, FinalityCodec.WAIT_FOR_FINALITY_FLAG, "", 1e18);
  }
}
