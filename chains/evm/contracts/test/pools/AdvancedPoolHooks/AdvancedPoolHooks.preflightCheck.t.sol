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
  // bytes4(keccak256("OutboundPolicyDataV1"))
  bytes4 internal constant OUTBOUND_POLICY_DATA_V1_TAG = 0x73bb902c;

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

    // Verify policy engine was called with correct payload
    IPolicyEngine.Payload memory lastPayload = s_mockPolicyEngine.getLastPayload();
    assertEq(lastPayload.selector, IAdvancedPoolHooks.preflightCheck.selector);
    assertEq(lastPayload.sender, OWNER);
    assertEq(lastPayload.context, "");

    // Verify tag prefix
    bytes4 tag = bytes4(lastPayload.data);
    assertEq(tag, OUTBOUND_POLICY_DATA_V1_TAG);

    // Decode and verify the payload data
    CCIPPolicyEnginePayloads.OutboundPolicyDataV1 memory decoded =
      abi.decode(this._sliceBytes(lastPayload.data, 4), (CCIPPolicyEnginePayloads.OutboundPolicyDataV1));

    assertEq(decoded.receiver, lockOrBurnIn.receiver);
    assertEq(decoded.remoteChainSelector, lockOrBurnIn.remoteChainSelector);
    assertEq(decoded.originalSender, lockOrBurnIn.originalSender);
    assertEq(decoded.amount, lockOrBurnIn.amount);
    assertEq(decoded.localToken, lockOrBurnIn.localToken);
    assertEq(decoded.blockConfirmationRequested, blockConfirmationRequested);
    assertEq(decoded.tokenArgs, tokenArgs);
  }

  function test_preflightCheck_WithoutPolicyEngine() public {
    assertEq(s_advancedPoolHooks.getPolicyEngine(), address(0));

    Pool.LockOrBurnInV1 memory lockOrBurnIn = _createLockOrBurnIn(OWNER);

    // Should not revert when policy engine is not set
    s_advancedPoolHooks.preflightCheck(lockOrBurnIn, 5, "");
  }

  function test_preflightCheck_AllowListAndPolicyEngine() public {
    address[] memory allowedSenders = new address[](1);
    allowedSenders[0] = OWNER;
    AdvancedPoolHooks hooksWithBoth = new AdvancedPoolHooks(allowedSenders, 0, address(s_mockPolicyEngine), new address[](0), true);

    Pool.LockOrBurnInV1 memory lockOrBurnIn = _createLockOrBurnIn(OWNER);

    hooksWithBoth.preflightCheck(lockOrBurnIn, 5, "");

    // Verify policy engine was called with correct tag
    IPolicyEngine.Payload memory lastPayload = s_mockPolicyEngine.getLastPayload();
    assertEq(lastPayload.selector, IAdvancedPoolHooks.preflightCheck.selector);
    assertEq(bytes4(lastPayload.data), OUTBOUND_POLICY_DATA_V1_TAG);
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
    AdvancedPoolHooks hooksWithAllowList = new AdvancedPoolHooks(allowedSenders, 0, address(0), new address[](0), true);

    Pool.LockOrBurnInV1 memory lockOrBurnIn = _createLockOrBurnIn(STRANGER);

    vm.expectRevert(abi.encodeWithSelector(AdvancedPoolHooks.SenderNotAllowed.selector, STRANGER));
    hooksWithAllowList.preflightCheck(lockOrBurnIn, 5, "");
  }

  // Helper to slice bytes, exposed as external for use with this.
  function _sliceBytes(bytes memory data, uint256 start) external pure returns (bytes memory) {
    bytes memory result = new bytes(data.length - start);
    for (uint256 i = 0; i < result.length; ++i) {
      result[i] = data[start + i];
    }
    return result;
  }
}
