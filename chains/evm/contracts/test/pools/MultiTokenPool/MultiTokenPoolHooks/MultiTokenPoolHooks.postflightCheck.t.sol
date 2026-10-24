// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IAdvancedPoolHooks} from "../../../../interfaces/IAdvancedPoolHooks.sol";
import {IPolicyEngine} from "@chainlink/policy-management/interfaces/IPolicyEngine.sol";

import {FinalityCodec} from "../../../../libraries/FinalityCodec.sol";
import {MultiTokenPoolHooks} from "../../../../pools/MultiTokenPool/MultiTokenPoolHooks.sol";
import {MockPolicyEngine} from "../../../mocks/MockPolicyEngine.sol";
import {MultiTokenPoolHooksSetup} from "./MultiTokenPoolHooksSetup.t.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

contract MultiTokenPoolHooks_postflightCheck is MultiTokenPoolHooksSetup {
  function test_postflightCheck() public {
    s_hooks.postflightCheck(s_releaseOrMintIn, 1e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG);
  }

  function test_postflightCheck_AllowListDoesNotApply() public {
    MultiTokenPoolHooks hooks = _deployHooksWithAllowList();
    s_releaseOrMintIn.originalSender = abi.encode(STRANGER);
    s_releaseOrMintIn.receiver = STRANGER;

    hooks.postflightCheck(s_releaseOrMintIn, 1e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG);
  }

  function test_postflightCheck_WithPolicyEngine() public {
    s_hooks.setPolicyEngine(s_policyEngine);
    s_releaseOrMintIn.offchainTokenData = abi.encode("custom offchain token data");
    uint256 localAmount = 1e6;
    bytes4 requestedFinalityConfig = FinalityCodec._encodeBlockDepth(5);

    s_hooks.postflightCheck(s_releaseOrMintIn, localAmount, requestedFinalityConfig);

    IPolicyEngine.Payload memory payload = s_policyEngine.getLastPayload();
    // The selector matches IAdvancedPoolHooks, so one policy engine extractor serves both kinds of hooks.
    assertEq(IAdvancedPoolHooks.postflightCheck.selector, payload.selector);
    assertEq(OWNER, payload.sender);
    assertEq(abi.encode(s_releaseOrMintIn, localAmount, requestedFinalityConfig), payload.data);
    assertEq(s_releaseOrMintIn.offchainTokenData, payload.context);
  }

  // Reverts

  function test_postflightCheck_RevertWhen_UnauthorizedCaller() public {
    vm.startPrank(STRANGER);

    vm.expectRevert(abi.encodeWithSelector(AuthorizedCallers.UnauthorizedCaller.selector, STRANGER));
    s_hooks.postflightCheck(s_releaseOrMintIn, 1e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG);
  }

  function test_postflightCheck_RevertWhen_PolicyEngineRejects() public {
    s_hooks.setPolicyEngine(s_policyEngine);
    s_policyEngine.setShouldRevert(true, "policy rejected");

    vm.expectRevert(abi.encodeWithSelector(MockPolicyEngine.MockPolicyEngineRejection.selector, "policy rejected"));
    s_hooks.postflightCheck(s_releaseOrMintIn, 1e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG);
  }
}
