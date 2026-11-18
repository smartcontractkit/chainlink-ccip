// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {FinalityCodec} from "../../libraries/FinalityCodec.sol";
import {MultiTokenPoolHooks} from "../../pools/MultiTokenPool/MultiTokenPoolHooks.sol";
import {ConnectionPolicyStore} from "../../pools/MultiTokenPool/policies/ConnectionPolicyStore.sol";
import {e2e_mttpSetup} from "./e2e.mttpSetup.t.sol";
import {IPolicyEngine} from "@chainlink/policy-management/interfaces/IPolicyEngine.sol";

/// @notice Security behavior: source token identity, the allowlist state across a hook migration, and ACE hook
/// replacement.
contract e2e_mttpSecurity is e2e_mttpSetup {
  /// @dev A wrong source token cannot mint another asset. Long-format pool data carries the source token, so
  /// delivering T1's data to T2, or T2's data to T1, is rejected. The emitted data names the source token, so the
  /// mismatch cannot be repaired by changing which token the source pool points at.
  function test_mttp_wrongSourceTokenCannotMintAnotherAsset() public {
    // The correct pairing works.
    assertEq(_transfer(s_a, s_b, 0, 10e18), 10e18, "correct T1 pairing");

    // T1's long-format pool data delivered to T2 on B is rejected.
    (bytes memory poolData, uint256 postFee,) =
      _send(s_a, s_b, 0, 10e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG, s_allowedSender);
    uint256 t2ReceiverBefore = _balance(s_b.tokens[1], s_receiver);
    vm.expectRevert(
      abi.encodeWithSelector(
        ConnectionPolicyStore.SourceTokenMismatch.selector,
        s_b.tokens[1],
        s_a.localSelector,
        abi.encode(s_a.tokens[1]),
        abi.encode(s_a.tokens[0])
      )
    );
    _deliver(s_b, s_a, s_b.tokens[1], postFee, poolData, FinalityCodec.WAIT_FOR_FINALITY_FLAG);
    assertEq(_balance(s_b.tokens[1], s_receiver), t2ReceiverBefore, "no T2 mint on mismatch");

    // The reverse misroute fails too: T2's data delivered to T1 on B.
    (bytes memory t2Data, uint256 t2PostFee,) =
      _send(s_a, s_b, 1, 10e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG, s_allowedSender);
    uint256 t1ReceiverBefore = _balance(s_b.tokens[0], s_receiver);
    vm.expectRevert(
      abi.encodeWithSelector(
        ConnectionPolicyStore.SourceTokenMismatch.selector,
        s_b.tokens[0],
        s_a.localSelector,
        abi.encode(s_a.tokens[0]),
        abi.encode(s_a.tokens[1])
      )
    );
    _deliver(s_b, s_a, s_b.tokens[0], t2PostFee, t2Data, FinalityCodec.WAIT_FOR_FINALITY_FLAG);
    assertEq(_balance(s_b.tokens[0], s_receiver), t1ReceiverBefore, "no T1 mint on reverse mismatch");

    // The same immutable T1 message delivered to its matching token succeeds.
    uint256 received = _deliver(s_b, s_a, s_b.tokens[0], postFee, poolData, FinalityCodec.WAIT_FOR_FINALITY_FLAG);
    assertEq(received, 10e18, "legitimate pairing recovers");
  }

  /// @dev Allowlist migration preserves deny-all. A disabled allowlist stays disabled, an enabled-populated one
  /// keeps its members, and an enabled-empty one, built with [address(0)], stays deny-all.
  function test_mttp_allowlistMigrationPreservesDenyAll() public {
    vm.startPrank(OWNER);
    MultiTokenPoolHooks disabled =
      new MultiTokenPoolHooks(new address[](0), IPolicyEngine(address(0)), _arrayOf(address(s_a.pool)));
    MultiTokenPoolHooks populated =
      new MultiTokenPoolHooks(_arrayOf(s_allowedSender), IPolicyEngine(address(0)), _arrayOf(address(s_a.pool)));
    // The current constructor treats [address(0)] as enabled-empty: enabled, but no allowed sender.
    MultiTokenPoolHooks enabledEmpty =
      new MultiTokenPoolHooks(_arrayOf(address(0)), IPolicyEngine(address(0)), _arrayOf(address(s_a.pool)));
    vm.stopPrank();

    assertFalse(disabled.getAllowListEnabled(), "disabled stays disabled");
    assertEq(disabled.getAllowList().length, 0, "disabled list");
    assertTrue(populated.getAllowListEnabled(), "populated enabled");
    assertEq(populated.getAllowList().length, 1, "populated membership preserved");
    assertEq(populated.getAllowList()[0], s_allowedSender, "populated member");
    assertTrue(enabledEmpty.getAllowListEnabled(), "enabled-empty enabled");
    assertEq(enabledEmpty.getAllowList().length, 0, "enabled-empty is deny-all");

    // Sender behavior matches the state.
    populated.checkAllowList(s_allowedSender);
    vm.expectRevert(abi.encodeWithSelector(MultiTokenPoolHooks.SenderNotAllowed.selector, s_disallowedSender));
    populated.checkAllowList(s_disallowedSender);
    vm.expectRevert(abi.encodeWithSelector(MultiTokenPoolHooks.SenderNotAllowed.selector, s_allowedSender));
    enabledEmpty.checkAllowList(s_allowedSender);

    // Exercise senders before and after switching to a replacement hook with the same membership.
    vm.prank(OWNER);
    _setInlineSecurityConfig(
      s_a.pcm, s_a.tokens[0], s_a.remoteSelector, new address[](0), new address[](0), address(populated)
    );
    assertEq(_transfer(s_a, s_b, 0, 10e18), 10e18, "allowed before switch");
    _approvePool(s_a, s_a.tokens[0], 10e18, s_disallowedSender);
    vm.expectRevert(abi.encodeWithSelector(MultiTokenPoolHooks.SenderNotAllowed.selector, s_disallowedSender));
    _lockOrBurn(s_a, s_b, 0, 10e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG, s_disallowedSender);

    vm.startPrank(OWNER);
    MultiTokenPoolHooks replacement =
      new MultiTokenPoolHooks(_arrayOf(s_allowedSender), IPolicyEngine(address(0)), _arrayOf(address(s_a.pool)));
    _setInlineSecurityConfig(
      s_a.pcm, s_a.tokens[0], s_a.remoteSelector, new address[](0), new address[](0), address(replacement)
    );
    vm.stopPrank();
    assertEq(_transfer(s_a, s_b, 0, 10e18), 10e18, "allowed after switch");
    _approvePool(s_a, s_a.tokens[0], 10e18, s_disallowedSender);
    vm.expectRevert(abi.encodeWithSelector(MultiTokenPoolHooks.SenderNotAllowed.selector, s_disallowedSender));
    _lockOrBurn(s_a, s_b, 0, 10e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG, s_disallowedSender);

    // An enabled-empty hook stays deny-all across a replacement. Rebuilding the list from getAllowList() returns
    // [] and would silently disable the allowlist, so the replacement is built from [address(0)].
    vm.prank(OWNER);
    _setInlineSecurityConfig(
      s_a.pcm, s_a.tokens[0], s_a.remoteSelector, new address[](0), new address[](0), address(enabledEmpty)
    );
    _approvePool(s_a, s_a.tokens[0], 10e18);
    vm.expectRevert(abi.encodeWithSelector(MultiTokenPoolHooks.SenderNotAllowed.selector, s_allowedSender));
    _lockOrBurn(s_a, s_b, 0, 10e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG, s_allowedSender);

    vm.startPrank(OWNER);
    MultiTokenPoolHooks enabledEmptyReplacement =
      new MultiTokenPoolHooks(_arrayOf(address(0)), IPolicyEngine(address(0)), _arrayOf(address(s_a.pool)));
    _setInlineSecurityConfig(
      s_a.pcm, s_a.tokens[0], s_a.remoteSelector, new address[](0), new address[](0), address(enabledEmptyReplacement)
    );
    vm.stopPrank();
    assertTrue(enabledEmptyReplacement.getAllowListEnabled(), "replacement keeps the allowlist enabled");
    _approvePool(s_a, s_a.tokens[0], 10e18);
    vm.expectRevert(abi.encodeWithSelector(MultiTokenPoolHooks.SenderNotAllowed.selector, s_allowedSender));
    _lockOrBurn(s_a, s_b, 0, 10e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG, s_allowedSender);
  }

  /// @dev Replacing the ACE hook preserves default inheritance. A fresh hook with no explicit target default
  /// follows later changes to the global default, while an explicit target default keeps its override.
  function test_mttp_aceHookReplacementPreservesDefaultInheritance() public {
    DefaultInheritingPolicyEngine engine = new DefaultInheritingPolicyEngine(true);

    vm.startPrank(OWNER);
    MultiTokenPoolHooks hook1 =
      new MultiTokenPoolHooks(new address[](0), IPolicyEngine(address(engine)), _arrayOf(address(s_a.pool)));
    _setInlineSecurityConfig(
      s_a.pcm, s_a.tokens[0], s_a.remoteSelector, new address[](0), new address[](0), address(hook1)
    );
    vm.stopPrank();

    // Allowed under the inherited global allow.
    assertEq(_transfer(s_a, s_b, 0, 10e18), 10e18, "allowed under inherited global allow");

    // Replace the hook. Fresh hook inherits the default; we don't set a target default.
    vm.startPrank(OWNER);
    MultiTokenPoolHooks hook2 =
      new MultiTokenPoolHooks(new address[](0), IPolicyEngine(address(engine)), _arrayOf(address(s_a.pool)));
    _setInlineSecurityConfig(
      s_a.pcm, s_a.tokens[0], s_a.remoteSelector, new address[](0), new address[](0), address(hook2)
    );
    vm.stopPrank();

    assertFalse(engine.targetHasDefault(address(hook2)), "fresh hook did not silently get an explicit default");

    // A later global deny applies to the inherited target.
    engine.setDefaultPolicyAllow(false);
    _approvePool(s_a, s_a.tokens[0], 10e18);
    vm.expectRevert();
    _lockOrBurn(s_a, s_b, 0, 10e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG, s_allowedSender);

    // An explicit target default overrides the global default.
    engine.setTargetDefaultPolicyAllow(address(hook2), true);
    assertEq(_transfer(s_a, s_b, 0, 10e18), 10e18, "explicit allow overrides global deny");

    // An explicit deny survives a global allow.
    engine.setTargetDefaultPolicyAllow(address(hook2), false);
    engine.setDefaultPolicyAllow(true);
    _approvePool(s_a, s_a.tokens[0], 10e18);
    vm.expectRevert();
    _lockOrBurn(s_a, s_b, 0, 10e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG, s_allowedSender);
  }
}

/// @notice A test double for the ACE policy engine. It models the documented default resolution: a target with an
/// explicit target default uses it, otherwise the target follows the global default. The ACE PolicyEngine depends
/// on openzeppelin contracts-upgradeable, which this repo does not install, so it cannot be imported here.
contract DefaultInheritingPolicyEngine is IPolicyEngine {
  error DefaultRejected(address target);

  bool internal s_defaultAllow;
  mapping(address target => bool hasDefault) internal s_targetHasDefault;
  mapping(address target => bool defaultAllow) internal s_targetDefaultAllow;

  constructor(
    bool defaultAllow
  ) {
    s_defaultAllow = defaultAllow;
  }

  function attach() external override {}

  function detach() external override {}

  function run(
    Payload calldata
  ) external view override {
    _checkDefault(msg.sender);
  }

  function check(
    Payload calldata
  ) external view override {
    _checkDefault(msg.sender);
  }

  function setDefaultPolicyAllow(
    bool defaultAllow
  ) external override {
    s_defaultAllow = defaultAllow;
  }

  function setTargetDefaultPolicyAllow(
    address target,
    bool defaultAllow
  ) external override {
    s_targetHasDefault[target] = true;
    s_targetDefaultAllow[target] = defaultAllow;
  }

  function targetHasDefault(
    address target
  ) external view returns (bool) {
    return s_targetHasDefault[target];
  }

  function _checkDefault(
    address target
  ) private view {
    bool allow = s_targetHasDefault[target] ? s_targetDefaultAllow[target] : s_defaultAllow;
    if (!allow) revert DefaultRejected(target);
  }

  // Unused interface surface.

  function typeAndVersion() external pure override returns (string memory) {
    return "DefaultInheritingPolicyEngine 1.0.0";
  }

  function setExtractor(
    bytes4,
    address
  ) external override {}

  function setExtractors(
    bytes4[] calldata,
    address
  ) external override {}

  function getExtractor(
    bytes4
  ) external pure override returns (address) {
    return address(0);
  }

  function setPolicyMapper(
    address,
    address
  ) external override {}

  function getPolicyMapper(
    address
  ) external pure override returns (address) {
    return address(0);
  }

  function addPolicy(
    address,
    bytes4,
    address,
    bytes32[] calldata
  ) external override {}

  function addPolicyAt(
    address,
    bytes4,
    address,
    bytes32[] calldata,
    uint256
  ) external override {}

  function removePolicy(
    address,
    bytes4,
    address
  ) external override {}

  function getPolicies(
    address,
    bytes4
  ) external pure override returns (address[] memory) {
    return new address[](0);
  }

  function setPolicyConfiguration(
    address,
    uint256,
    bytes4,
    bytes calldata
  ) external override {}

  function getPolicyConfigVersion(
    address
  ) external pure override returns (uint256) {
    return 0;
  }
}
