// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IAdvancedPoolHooks} from "../../../interfaces/IAdvancedPoolHooks.sol";
import {IPolicyEngine} from "../../../interfaces/IPolicyEngine.sol";

import {CCIPPolicyEnginePayloads} from "../../../libraries/CCIPPolicyEnginePayloads.sol";
import {Pool} from "../../../libraries/Pool.sol";
import {AdvancedPoolHooks} from "../../../pools/AdvancedPoolHooks.sol";
import {MockPolicyEngine} from "../../mocks/MockPolicyEngine.sol";
import {AdvancedPoolHooksSetup} from "./AdvancedPoolHooksSetup.t.sol";

contract AdvancedPoolHooks_preflightCheck is AdvancedPoolHooksSetup {
  MockPolicyEngine internal s_mockPolicyEngine;

  function setUp() public virtual override {
    super.setUp();
    s_mockPolicyEngine = new MockPolicyEngine();
  }

  function _createLockOrBurnIn(
    address originalSender
  ) internal view returns (Pool.LockOrBurnInV1 memory) {
    return Pool.LockOrBurnInV1({
      receiver: s_receiver,
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      originalSender: originalSender,
      amount: 100e18,
      localToken: address(s_token)
    });
  }

  function test_preflightCheck_WithPolicyEngine() public {
    s_advancedPoolHooks.setPolicyEngine(address(s_mockPolicyEngine));

    Pool.LockOrBurnInV1 memory lockOrBurnIn = _createLockOrBurnIn(OWNER);
    uint16 blockConfirmationRequested = 5;
    bytes memory tokenArgs = abi.encode("test");

    s_advancedPoolHooks.preflightCheck(lockOrBurnIn, blockConfirmationRequested, tokenArgs);

    IPolicyEngine.Payload memory lastPayload = s_mockPolicyEngine.getLastPayload();
    assertEq(IAdvancedPoolHooks.preflightCheck.selector, lastPayload.selector);
    assertEq(OWNER, lastPayload.sender);
    assertEq("", lastPayload.context);
    assertEq(CCIPPolicyEnginePayloads.POOL_HOOK_OUTBOUND_POLICY_DATA_V1_TAG, bytes4(lastPayload.data));

    CCIPPolicyEnginePayloads.PoolHookOutboundPolicyDataV1 memory decoded =
      abi.decode(_sliceBytes(lastPayload.data, 4), (CCIPPolicyEnginePayloads.PoolHookOutboundPolicyDataV1));

    assertEq(lockOrBurnIn.originalSender, decoded.originalSender);
    assertEq(blockConfirmationRequested, decoded.blockConfirmationRequested);
    assertEq(lockOrBurnIn.remoteChainSelector, decoded.remoteChainSelector);
    assertEq(lockOrBurnIn.receiver, decoded.receiver);
    assertEq(lockOrBurnIn.amount, decoded.amount);
    assertEq(lockOrBurnIn.localToken, decoded.localToken);
    assertEq(tokenArgs, decoded.tokenArgs);
  }

  function test_preflightCheck_WithoutPolicyEngine() public {
    assertEq(address(0), s_advancedPoolHooks.getPolicyEngine());

    Pool.LockOrBurnInV1 memory lockOrBurnIn = _createLockOrBurnIn(OWNER);

    // Should not revert when policy engine is not set
    s_advancedPoolHooks.preflightCheck(lockOrBurnIn, 5, "");
  }

  function test_preflightCheck_AllowListAndPolicyEngine() public {
    address[] memory allowedSenders = new address[](1);
    allowedSenders[0] = OWNER;
    AdvancedPoolHooks hooksWithBoth =
      new AdvancedPoolHooks(allowedSenders, 0, address(s_mockPolicyEngine), new address[](0));

    Pool.LockOrBurnInV1 memory lockOrBurnIn = _createLockOrBurnIn(OWNER);

    hooksWithBoth.preflightCheck(lockOrBurnIn, 5, "");

    IPolicyEngine.Payload memory lastPayload = s_mockPolicyEngine.getLastPayload();
    assertEq(IAdvancedPoolHooks.preflightCheck.selector, lastPayload.selector);
    assertEq(CCIPPolicyEnginePayloads.POOL_HOOK_OUTBOUND_POLICY_DATA_V1_TAG, bytes4(lastPayload.data));
  }

  // Reverts

  function test_preflightCheck_RevertWhen_PolicyEngineRejects() public {
    s_advancedPoolHooks.setPolicyEngine(address(s_mockPolicyEngine));
    s_mockPolicyEngine.setShouldRevert(true, "Policy rejected");

    Pool.LockOrBurnInV1 memory lockOrBurnIn = _createLockOrBurnIn(OWNER);

    vm.expectRevert(abi.encodeWithSelector(MockPolicyEngine.MockPolicyEngineRejection.selector, "Policy rejected"));
    s_advancedPoolHooks.preflightCheck(lockOrBurnIn, 5, "");
  }

  function test_preflightCheck_RevertWhen_SenderNotAllowed() public {
    address[] memory allowedSenders = new address[](1);
    allowedSenders[0] = OWNER;
    AdvancedPoolHooks hooksWithAllowList = new AdvancedPoolHooks(allowedSenders, 0, address(0), new address[](0));

    Pool.LockOrBurnInV1 memory lockOrBurnIn = _createLockOrBurnIn(STRANGER);

    vm.expectRevert(abi.encodeWithSelector(AdvancedPoolHooks.SenderNotAllowed.selector, STRANGER));
    hooksWithAllowList.preflightCheck(lockOrBurnIn, 5, "");
  }
}
