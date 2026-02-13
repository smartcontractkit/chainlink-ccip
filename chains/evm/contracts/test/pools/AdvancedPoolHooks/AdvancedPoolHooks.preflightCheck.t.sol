// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IAdvancedPoolHooks} from "../../../interfaces/IAdvancedPoolHooks.sol";
import {IPolicyEngine} from "@chainlink/ace/policy-management/interfaces/IPolicyEngine.sol";

import {Pool} from "../../../libraries/Pool.sol";
import {AdvancedPoolHooks} from "../../../pools/AdvancedPoolHooks.sol";
import {MockPolicyEngine} from "../../mocks/MockPolicyEngine.sol";
import {AdvancedPoolHooksSetup} from "./AdvancedPoolHooksSetup.t.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

contract AdvancedPoolHooks_preflightCheck is AdvancedPoolHooksSetup {
  MockPolicyEngine internal s_mockPolicyEngine;

  address internal s_unauthorizedCaller = makeAddr("unauthorizedCaller");

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
    bytes memory tokenArgs = abi.encode("custom token args");

    s_advancedPoolHooks.preflightCheck(lockOrBurnIn, blockConfirmationRequested, tokenArgs, lockOrBurnIn.amount);

    IPolicyEngine.Payload memory lastPayload = s_mockPolicyEngine.getLastPayload();
    assertEq(IAdvancedPoolHooks.preflightCheck.selector, lastPayload.selector);
    assertEq(OWNER, lastPayload.sender);
    assertEq(tokenArgs, lastPayload.context);
    assertEq(abi.encode(lockOrBurnIn, blockConfirmationRequested, tokenArgs, lockOrBurnIn.amount), lastPayload.data);
  }

  function test_preflightCheck_WithoutPolicyEngine() public {
    assertEq(address(0), s_advancedPoolHooks.getPolicyEngine());

    Pool.LockOrBurnInV1 memory lockOrBurnIn = _createLockOrBurnIn(OWNER);

    s_advancedPoolHooks.preflightCheck(lockOrBurnIn, 5, "", lockOrBurnIn.amount);
  }

  function test_preflightCheck_AllowListAndPolicyEngine() public {
    address[] memory allowedSenders = new address[](1);
    allowedSenders[0] = OWNER;
    address[] memory callers = new address[](1);
    callers[0] = OWNER;
    AdvancedPoolHooks hooksWithBoth = new AdvancedPoolHooks(allowedSenders, 0, address(s_mockPolicyEngine), callers);

    Pool.LockOrBurnInV1 memory lockOrBurnIn = _createLockOrBurnIn(OWNER);

    hooksWithBoth.preflightCheck(lockOrBurnIn, 5, "", lockOrBurnIn.amount);

    IPolicyEngine.Payload memory lastPayload = s_mockPolicyEngine.getLastPayload();
    assertEq(IAdvancedPoolHooks.preflightCheck.selector, lastPayload.selector);
    assertEq(abi.encode(lockOrBurnIn, uint16(5), bytes(""), lockOrBurnIn.amount), lastPayload.data);
    assertEq(bytes(""), lastPayload.context);
  }

  function test_preflightCheck_RevertWhen_PolicyEngineRejects() public {
    string memory expectedRevertReason = "policy rejected";
    s_advancedPoolHooks.setPolicyEngine(address(s_mockPolicyEngine));
    s_mockPolicyEngine.setShouldRevert(true, expectedRevertReason);

    Pool.LockOrBurnInV1 memory lockOrBurnIn = _createLockOrBurnIn(OWNER);

    vm.expectRevert(abi.encodeWithSelector(MockPolicyEngine.MockPolicyEngineRejection.selector, expectedRevertReason));
    s_advancedPoolHooks.preflightCheck(lockOrBurnIn, 5, "", lockOrBurnIn.amount);
  }

  function test_preflightCheck_RevertWhen_SenderNotAllowed() public {
    address[] memory allowedSenders = new address[](1);
    allowedSenders[0] = OWNER;
    address[] memory callers = new address[](1);
    callers[0] = OWNER;
    AdvancedPoolHooks hooksWithAllowList = new AdvancedPoolHooks(allowedSenders, 0, address(0), callers);

    Pool.LockOrBurnInV1 memory lockOrBurnIn = _createLockOrBurnIn(STRANGER);

    vm.expectRevert(abi.encodeWithSelector(AdvancedPoolHooks.SenderNotAllowed.selector, STRANGER));
    hooksWithAllowList.preflightCheck(lockOrBurnIn, 5, "", lockOrBurnIn.amount);
  }

  function test_preflightCheck_OnlyAuthorizedCallersCanInvoke() public {
    Pool.LockOrBurnInV1 memory lockOrBurnIn = _createLockOrBurnIn(OWNER);

    s_advancedPoolHooks.preflightCheck(lockOrBurnIn, 5, "", lockOrBurnIn.amount);
  }

  function test_preflightCheck_RevertWhen_UnauthorizedCaller() public {
    vm.stopPrank();

    Pool.LockOrBurnInV1 memory lockOrBurnIn = _createLockOrBurnIn(OWNER);

    vm.prank(s_unauthorizedCaller);
    vm.expectRevert(abi.encodeWithSelector(AuthorizedCallers.UnauthorizedCaller.selector, s_unauthorizedCaller));
    s_advancedPoolHooks.preflightCheck(lockOrBurnIn, 5, "", lockOrBurnIn.amount);
  }
}
