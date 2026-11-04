// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../../../../interfaces/IPoolConfigManager.sol";

import {FinalityCodec} from "../../../../../libraries/FinalityCodec.sol";
import {PolicyConfig} from "../../../../../libraries/PolicyConfig.sol";
import {ConnectionPolicyStore} from "../../../../../pools/MultiTokenPool/policies/ConnectionPolicyStore.sol";
import {ConnectionPolicyStoreSetup} from "./ConnectionPolicyStoreSetup.t.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

contract ConnectionPolicyStore_applyTokenRemotePoolPolicyUpdates is ConnectionPolicyStoreSetup {
  /// @dev OTHER_POLICY trusts s_otherRemotePool on DEST_CHAIN_SELECTOR only.
  function setUp() public override {
    super.setUp();
    _setRemotePoolPolicyConfig(
      _getRemotePoolPolicyConfigArgs(OTHER_POLICY, DEST_CHAIN_SELECTOR, _getRemotePoolAddresses(s_otherRemotePool))
    );
  }

  function test_applyTokenRemotePoolPolicyUpdates() public {
    PolicyConfig.TokenPolicyArgs[] memory tokenPolicies = new PolicyConfig.TokenPolicyArgs[](2);
    tokenPolicies[0] = PolicyConfig.TokenPolicyArgs({localToken: s_token, policyName: OTHER_POLICY});
    tokenPolicies[1] = PolicyConfig.TokenPolicyArgs({localToken: s_otherToken, policyName: REMOTE_POOL_POLICY});

    vm.expectEmit();
    emit ConnectionPolicyStore.TokenRemotePoolPolicySet(s_token, REMOTE_POOL_POLICY, OTHER_POLICY);
    vm.expectEmit();
    emit ConnectionPolicyStore.TokenRemotePoolPolicySet(s_otherToken, bytes32(0), REMOTE_POOL_POLICY);
    s_connectionStore.applyTokenRemotePoolPolicyUpdates(tokenPolicies);

    (,, bytes32 remotePoolPolicyName) = s_connectionStore.getTokenPolicies(s_token);
    assertEq(OTHER_POLICY, remotePoolPolicyName);
    (,, remotePoolPolicyName) = s_connectionStore.getTokenPolicies(s_otherToken);
    assertEq(REMOTE_POOL_POLICY, remotePoolPolicyName);
    uint32 tokenCount = _getRemotePoolPolicyTokenCount(REMOTE_POOL_POLICY);
    assertEq(1, tokenCount);
    tokenCount = _getRemotePoolPolicyTokenCount(OTHER_POLICY);
    assertEq(1, tokenCount);
    IPoolConfigManager.LaneConnection[] memory connections =
      s_connectionStore.getLaneConnectionBatch(_getLaneKeys(s_token, DEST_CHAIN_SELECTOR));
    assertEq(1, connections[0].remotePools.length);
    assertEq(s_otherRemotePool, connections[0].remotePools[0].remotePoolAddress);
  }

  function test_applyTokenRemotePoolPolicyUpdates_InlineRecordTrustingPolicy() public {
    bytes memory legacyPool = abi.encode(makeAddr("legacyPool"));
    _setInlineRemotePoolConfigTrustingPolicy(s_token, legacyPool, true);
    bytes memory sourcePoolData = _getSourcePoolData(DEFAULT_TOKEN_DECIMALS, s_token, "");
    bytes memory legacyPoolData = abi.encode(uint256(DEFAULT_TOKEN_DECIMALS));

    // Changing policy replaces the policy pools while preserving the inline pool.
    _setTokenRemotePoolPolicy(s_token, OTHER_POLICY);
    _resolveInbound(s_otherRemotePool, sourcePoolData);
    _resolveInbound(legacyPool, legacyPoolData);
    vm.expectRevert(
      abi.encodeWithSelector(
        ConnectionPolicyStore.InvalidSourcePoolAddress.selector, s_token, DEST_CHAIN_SELECTOR, s_remotePool
      )
    );
    _resolveInbound(s_remotePool, sourcePoolData);

    // Leaving the policy keeps only the inline pool.
    _setTokenRemotePoolPolicy(s_token, bytes32(0));
    _resolveInbound(legacyPool, legacyPoolData);
    vm.expectRevert(
      abi.encodeWithSelector(
        ConnectionPolicyStore.InvalidSourcePoolAddress.selector, s_token, DEST_CHAIN_SELECTOR, s_otherRemotePool
      )
    );
    _resolveInbound(s_otherRemotePool, sourcePoolData);
  }

  function test_applyTokenRemotePoolPolicyUpdates_ZeroPolicyNameClearsTokenPolicy() public {
    vm.expectEmit();
    emit ConnectionPolicyStore.TokenRemotePoolPolicySet(s_token, REMOTE_POOL_POLICY, bytes32(0));
    s_connectionStore.applyTokenRemotePoolPolicyUpdates(_getTokenPolicyArgs(s_token, bytes32(0)));

    (,, bytes32 remotePoolPolicyName) = s_connectionStore.getTokenPolicies(s_token);
    assertEq(bytes32(0), remotePoolPolicyName);
    uint32 tokenCount = _getRemotePoolPolicyTokenCount(REMOTE_POOL_POLICY);
    assertEq(0, tokenCount);
    // The connection still exists, but transfers require a remote pool record.
    (, PolicyConfig.RecordLevel[] memory remotePoolLevels) =
      s_connectionStore.getLaneRecordLevelBatch(_getLaneKeys(s_token, DEST_CHAIN_SELECTOR));
    assertEq(uint8(PolicyConfig.RecordLevel.None), uint8(remotePoolLevels[0]));
    vm.expectRevert(
      abi.encodeWithSelector(ConnectionPolicyStore.RemotePoolRecordNotFound.selector, s_token, DEST_CHAIN_SELECTOR)
    );
    s_connectionStore.resolveOutboundLane(s_token, DEST_CHAIN_SELECTOR, FinalityCodec.WAIT_FOR_FINALITY_FLAG, "");
  }

  function test_applyTokenRemotePoolPolicyUpdates_SamePolicy() public {
    vm.expectEmit();
    emit ConnectionPolicyStore.TokenRemotePoolPolicySet(s_token, REMOTE_POOL_POLICY, REMOTE_POOL_POLICY);
    s_connectionStore.applyTokenRemotePoolPolicyUpdates(_getTokenPolicyArgs(s_token, REMOTE_POOL_POLICY));

    uint32 tokenCount = _getRemotePoolPolicyTokenCount(REMOTE_POOL_POLICY);
    assertEq(1, tokenCount);
  }

  function test_applyTokenRemotePoolPolicyUpdates_SameTokenTwice() public {
    // Sequential assignments for the same token must update each policy's token count correctly.
    PolicyConfig.TokenPolicyArgs[] memory tokenPolicies = new PolicyConfig.TokenPolicyArgs[](2);
    tokenPolicies[0] = PolicyConfig.TokenPolicyArgs({localToken: s_token, policyName: OTHER_POLICY});
    tokenPolicies[1] = PolicyConfig.TokenPolicyArgs({localToken: s_token, policyName: bytes32(0)});

    vm.expectEmit();
    emit ConnectionPolicyStore.TokenRemotePoolPolicySet(s_token, REMOTE_POOL_POLICY, OTHER_POLICY);
    vm.expectEmit();
    emit ConnectionPolicyStore.TokenRemotePoolPolicySet(s_token, OTHER_POLICY, bytes32(0));
    s_connectionStore.applyTokenRemotePoolPolicyUpdates(tokenPolicies);

    (,, bytes32 remotePoolPolicyName) = s_connectionStore.getTokenPolicies(s_token);
    assertEq(bytes32(0), remotePoolPolicyName);
    uint32 tokenCount = _getRemotePoolPolicyTokenCount(REMOTE_POOL_POLICY);
    assertEq(0, tokenCount);
    tokenCount = _getRemotePoolPolicyTokenCount(OTHER_POLICY);
    assertEq(0, tokenCount);
  }

  function test_applyTokenRemotePoolPolicyUpdates_InlineRecordStays() public {
    _setInlineRemotePoolConfig(_getInlineRemotePoolConfigArgs(s_token, DEST_CHAIN_SELECTOR, _getRemotePoolConfig()));

    _setTokenRemotePoolPolicy(s_token, OTHER_POLICY);

    assertEq(1, s_connectionStore.getInlineRemotePoolConfigs(s_token).length);
    IPoolConfigManager.LaneConnection[] memory connections =
      s_connectionStore.getLaneConnectionBatch(_getLaneKeys(s_token, DEST_CHAIN_SELECTOR));
    assertEq(s_remotePool, connections[0].remotePools[0].remotePoolAddress);
  }

  // Reverts

  function test_applyTokenRemotePoolPolicyUpdates_RevertWhen_LocalTokenNotRegistered() public {
    address unknownToken = makeAddr("unknownToken");

    vm.expectRevert(abi.encodeWithSelector(PolicyConfig.LocalTokenNotRegistered.selector, unknownToken));
    s_connectionStore.applyTokenRemotePoolPolicyUpdates(_getTokenPolicyArgs(unknownToken, REMOTE_POOL_POLICY));
  }

  function test_applyTokenRemotePoolPolicyUpdates_RevertWhen_PolicyNotFound() public {
    // CONNECTION_POLICY exists only as a connection policy.
    vm.expectRevert(abi.encodeWithSelector(PolicyConfig.PolicyNotFound.selector, CONNECTION_POLICY));
    s_connectionStore.applyTokenRemotePoolPolicyUpdates(_getTokenPolicyArgs(s_token, CONNECTION_POLICY));
  }

  function test_applyTokenRemotePoolPolicyUpdates_RevertWhen_UnauthorizedCaller() public {
    vm.startPrank(STRANGER);

    vm.expectRevert(abi.encodeWithSelector(AuthorizedCallers.UnauthorizedCaller.selector, STRANGER));
    s_connectionStore.applyTokenRemotePoolPolicyUpdates(_getTokenPolicyArgs(s_token, OTHER_POLICY));
  }

  function _resolveInbound(
    bytes memory sourcePoolAddress,
    bytes memory sourcePoolData
  ) private view {
    s_connectionStore.resolveInboundLane(
      s_token, DEST_CHAIN_SELECTOR, sourcePoolAddress, sourcePoolData, 1e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG
    );
  }
}
