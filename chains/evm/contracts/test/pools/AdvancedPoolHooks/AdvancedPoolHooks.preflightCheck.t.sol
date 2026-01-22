// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IAdvancedPoolHooks} from "../../../interfaces/IAdvancedPoolHooks.sol";
import {IPolicyEngine} from "../../../interfaces/IPolicyEngine.sol";

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
    // Set up policy engine
    s_advancedPoolHooks.setPolicyEngine(address(s_mockPolicyEngine));

    Pool.LockOrBurnInV1 memory lockOrBurnIn = _createLockOrBurnIn(OWNER);
    uint16 blockConfirmationRequested = 5;
    bytes memory tokenArgs = abi.encode("test");

    // Call preflightCheck
    s_advancedPoolHooks.preflightCheck(lockOrBurnIn, blockConfirmationRequested, tokenArgs);

    // Verify policy engine was called with correct payload
    IPolicyEngine.Payload memory lastPayload = s_mockPolicyEngine.getLastPayload();
    assertEq(lastPayload.selector, IAdvancedPoolHooks.preflightCheck.selector);
    assertEq(lastPayload.sender, OWNER); // msg.sender is OWNER from vm.startPrank in setup

    // Verify the encoded data contains the lockOrBurnIn, blockConfirmationRequested, and tokenArgs
    bytes memory expectedData = abi.encode(lockOrBurnIn, blockConfirmationRequested, tokenArgs);
    assertEq(lastPayload.data, expectedData);
    assertEq(lastPayload.context, "");
  }

  function test_preflightCheck_WithoutPolicyEngine() public {
    // Ensure no policy engine is set
    assertEq(s_advancedPoolHooks.getPolicyEngine(), address(0));

    Pool.LockOrBurnInV1 memory lockOrBurnIn = _createLockOrBurnIn(OWNER);

    // Should not revert when policy engine is not set
    s_advancedPoolHooks.preflightCheck(lockOrBurnIn, 5, "");
  }

  function test_preflightCheck_RevertWhen_PolicyEngineRejects() public {
    // Set up policy engine to reject
    s_advancedPoolHooks.setPolicyEngine(address(s_mockPolicyEngine));
    s_mockPolicyEngine.setShouldRevert(true, "Policy rejected");

    Pool.LockOrBurnInV1 memory lockOrBurnIn = _createLockOrBurnIn(OWNER);

    vm.expectRevert(abi.encodeWithSelector(MockPolicyEngine.MockPolicyEngineRejection.selector, "Policy rejected"));
    s_advancedPoolHooks.preflightCheck(lockOrBurnIn, 5, "");
  }

  function test_preflightCheck_RevertWhen_SenderNotAllowed() public {
    // Create hooks with allowlist enabled
    address[] memory allowedSenders = new address[](1);
    allowedSenders[0] = OWNER;
    AdvancedPoolHooks hooksWithAllowList = new AdvancedPoolHooks(allowedSenders, 0, address(0));

    Pool.LockOrBurnInV1 memory lockOrBurnIn = _createLockOrBurnIn(STRANGER);

    vm.expectRevert(abi.encodeWithSelector(AdvancedPoolHooks.SenderNotAllowed.selector, STRANGER));
    hooksWithAllowList.preflightCheck(lockOrBurnIn, 5, "");
  }

  function test_preflightCheck_AllowListAndPolicyEngine() public {
    // Create hooks with both allowlist and policy engine
    address[] memory allowedSenders = new address[](1);
    allowedSenders[0] = OWNER;
    AdvancedPoolHooks hooksWithBoth = new AdvancedPoolHooks(allowedSenders, 0, address(s_mockPolicyEngine));

    Pool.LockOrBurnInV1 memory lockOrBurnIn = _createLockOrBurnIn(OWNER);

    // Should succeed when sender is allowed and policy engine passes
    hooksWithBoth.preflightCheck(lockOrBurnIn, 5, "");

    // Verify policy engine was called
    IPolicyEngine.Payload memory lastPayload = s_mockPolicyEngine.getLastPayload();
    assertEq(lastPayload.selector, IAdvancedPoolHooks.preflightCheck.selector);
  }
}
