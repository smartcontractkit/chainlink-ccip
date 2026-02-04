// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IAdvancedPoolHooks} from "../../../interfaces/IAdvancedPoolHooks.sol";
import {IPolicyEngine} from "../../../interfaces/IPolicyEngine.sol";

import {Pool} from "../../../libraries/Pool.sol";
import {AdvancedPoolHooks} from "../../../pools/AdvancedPoolHooks.sol";
import {MockPolicyEngine} from "../../mocks/MockPolicyEngine.sol";
import {AdvancedPoolHooksSetup} from "./AdvancedPoolHooksSetup.t.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

contract AdvancedPoolHooks_preflightCheck is AdvancedPoolHooksSetup {
  MockPolicyEngine internal s_mockPolicyEngine;

  address internal s_authorizedCaller = makeAddr("authorizedCaller");
  address internal s_unauthorizedCaller = makeAddr("unauthorizedCaller");
  AdvancedPoolHooks internal s_hooksWithAuthorizedCallers;

  function setUp() public virtual override {
    super.setUp();
    s_mockPolicyEngine = new MockPolicyEngine();

    address[] memory authorizedCallers = new address[](1);
    authorizedCallers[0] = s_authorizedCaller;
    s_hooksWithAuthorizedCallers = new AdvancedPoolHooks(new address[](0), 0, address(0), authorizedCallers);
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

  function testFuzz_preflightCheck_WithPolicyEngine(
    bytes memory tokenArgs
  ) public {
    s_advancedPoolHooks.setPolicyEngine(address(s_mockPolicyEngine));

    Pool.LockOrBurnInV1 memory lockOrBurnIn = _createLockOrBurnIn(OWNER);
    uint16 blockConfirmationRequested = 5;

    s_advancedPoolHooks.preflightCheck(lockOrBurnIn, blockConfirmationRequested, tokenArgs);

    IPolicyEngine.Payload memory lastPayload = s_mockPolicyEngine.getLastPayload();
    assertEq(IAdvancedPoolHooks.preflightCheck.selector, lastPayload.selector);
    assertEq(OWNER, lastPayload.sender);
    assertEq(tokenArgs, lastPayload.context);
    assertEq(abi.encode(lockOrBurnIn, blockConfirmationRequested, tokenArgs), lastPayload.data);
  }

  function test_preflightCheck_WithoutPolicyEngine() public {
    assertEq(address(0), s_advancedPoolHooks.getPolicyEngine());

    Pool.LockOrBurnInV1 memory lockOrBurnIn = _createLockOrBurnIn(OWNER);

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
    assertEq(abi.encode(lockOrBurnIn, uint16(5), bytes("")), lastPayload.data);
    assertEq(bytes(""), lastPayload.context);
  }

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

  function test_preflightCheck_AnyoneCanInvoke_WhenAuthorizedCallersDisabled() public {
    vm.stopPrank();
    assertFalse(s_advancedPoolHooks.getAuthorizedCallersEnabled());

    Pool.LockOrBurnInV1 memory lockOrBurnIn = _createLockOrBurnIn(OWNER);

    vm.prank(s_unauthorizedCaller);
    s_advancedPoolHooks.preflightCheck(lockOrBurnIn, 5, "");
  }

  function test_preflightCheck_OnlyAuthorizedCallersCanInvoke() public {
    vm.stopPrank();
    assertTrue(s_hooksWithAuthorizedCallers.getAuthorizedCallersEnabled());

    Pool.LockOrBurnInV1 memory lockOrBurnIn = _createLockOrBurnIn(OWNER);

    vm.prank(s_authorizedCaller);
    s_hooksWithAuthorizedCallers.preflightCheck(lockOrBurnIn, 5, "");
  }

  function test_preflightCheck_RevertWhen_UnauthorizedCaller() public {
    vm.stopPrank();
    assertTrue(s_hooksWithAuthorizedCallers.getAuthorizedCallersEnabled());

    Pool.LockOrBurnInV1 memory lockOrBurnIn = _createLockOrBurnIn(OWNER);

    vm.prank(s_unauthorizedCaller);
    vm.expectRevert(abi.encodeWithSelector(AuthorizedCallers.UnauthorizedCaller.selector, s_unauthorizedCaller));
    s_hooksWithAuthorizedCallers.preflightCheck(lockOrBurnIn, 5, "");
  }
}
