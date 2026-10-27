// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IAdvancedPoolHooks} from "../../../../interfaces/IAdvancedPoolHooks.sol";
import {IPolicyEngine} from "@chainlink/policy-management/interfaces/IPolicyEngine.sol";

import {FinalityCodec} from "../../../../libraries/FinalityCodec.sol";
import {MockPolicyEngine} from "../../../mocks/MockPolicyEngine.sol";
import {MultiTokenPoolHooksSetup} from "./MultiTokenPoolHooksSetup.t.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

contract MultiTokenPoolHooks_postflightCheck is MultiTokenPoolHooksSetup {
  function test_postflightCheck() public {
    vm.startPrank(s_pool);

    s_hooks.postflightCheck(s_releaseOrMintIn, 1e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG);
  }

  function test_postflightCheck_AllowListDoesNotApply() public {
    s_releaseOrMintIn.originalSender = abi.encode(STRANGER);
    s_releaseOrMintIn.receiver = STRANGER;
    vm.startPrank(s_pool);

    s_hooksWithAllowList.postflightCheck(s_releaseOrMintIn, 1e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG);
  }

  function test_postflightCheck_WithPolicyEngine() public {
    s_hooks.setPolicyEngine(s_policyEngine);
    s_releaseOrMintIn.offchainTokenData = abi.encode("custom offchain token data");
    uint256 localAmount = 1e6;
    bytes4 requestedFinalityConfig = FinalityCodec._encodeBlockDepth(5);
    vm.startPrank(s_pool);

    s_hooks.postflightCheck(s_releaseOrMintIn, localAmount, requestedFinalityConfig);

    IPolicyEngine.Payload memory payload = s_policyEngine.getLastPayload();
    // Same selector as IAdvancedPoolHooks, so AdvancedPoolHooksExtractor can decode the payload.
    assertEq(IAdvancedPoolHooks.postflightCheck.selector, payload.selector);
    assertEq(s_pool, payload.sender);
    assertEq(abi.encode(s_releaseOrMintIn, localAmount, requestedFinalityConfig), payload.data);
    assertEq(s_releaseOrMintIn.offchainTokenData, payload.context);
  }

  // Reverts

  function test_postflightCheck_RevertWhen_UnauthorizedCaller() public {
    // Owning the hooks does not make the owner an authorized caller.
    vm.expectRevert(abi.encodeWithSelector(AuthorizedCallers.UnauthorizedCaller.selector, OWNER));
    s_hooks.postflightCheck(s_releaseOrMintIn, 1e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG);
  }

  function test_postflightCheck_RevertWhen_PolicyEngineRejects() public {
    s_hooks.setPolicyEngine(s_policyEngine);
    s_policyEngine.setShouldRevert(true, "policy rejected");
    vm.startPrank(s_pool);

    vm.expectRevert(abi.encodeWithSelector(MockPolicyEngine.MockPolicyEngineRejection.selector, "policy rejected"));
    s_hooks.postflightCheck(s_releaseOrMintIn, 1e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG);
  }
}
