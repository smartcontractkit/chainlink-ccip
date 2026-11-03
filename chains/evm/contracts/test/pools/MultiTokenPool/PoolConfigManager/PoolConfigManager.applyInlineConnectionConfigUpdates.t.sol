// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../../../interfaces/IPoolConfigManager.sol";

import {PolicyConfig} from "../../../../libraries/PolicyConfig.sol";
import {PoolConfigManager} from "../../../../pools/MultiTokenPool/PoolConfigManager.sol";
import {ConnectionPolicyStore} from "../../../../pools/MultiTokenPool/policies/ConnectionPolicyStore.sol";
import {FeePolicyStore} from "../../../../pools/MultiTokenPool/policies/FeePolicyStore.sol";
import {SecurityPolicyStore} from "../../../../pools/MultiTokenPool/policies/SecurityPolicyStore.sol";
import {PoolConfigManagerRolesSetup} from "./PoolConfigManagerRolesSetup.t.sol";

contract PoolConfigManager_applyInlineConnectionConfigUpdates is PoolConfigManagerRolesSetup {
  address internal s_remoteTokenB = makeAddr("remoteTokenB");

  /// @dev A remote pool policy has no default record, so the DEFAULT policy needs a record on the other chain for the
  /// lane check of the inline records that the tests set there.
  function setUp() public virtual override {
    super.setUp();
    _setRemotePoolPolicyConfig(DEFAULT_POLICY, OTHER_CHAIN_SELECTOR, _getRemotePoolAddresses(abi.encode(s_remotePool)));
  }

  /// @notice CONFIG_ROLE sets inline records of several tokens in one call. Each inline record opens a lane. Each token
  /// has a record of every other type on the other chain, so the lane check passes for each record.
  function test_applyInlineConnectionConfigUpdates() public {
    ConnectionPolicyStore.InlineConnectionConfigArgs[] memory inlineConfigs =
      new ConnectionPolicyStore.InlineConnectionConfigArgs[](2);
    inlineConfigs[0] = ConnectionPolicyStore.InlineConnectionConfigArgs({
      localToken: s_tokenA, remoteChainSelector: OTHER_CHAIN_SELECTOR, remoteToken: abi.encode(s_remoteTokenA)
    });
    inlineConfigs[1] = ConnectionPolicyStore.InlineConnectionConfigArgs({
      localToken: s_tokenB, remoteChainSelector: OTHER_CHAIN_SELECTOR, remoteToken: abi.encode(s_remoteTokenB)
    });

    vm.expectEmit(address(s_connectionStore));
    emit ConnectionPolicyStore.InlineConnectionConfigSet(s_tokenA, OTHER_CHAIN_SELECTOR, abi.encode(s_remoteTokenA));
    vm.expectEmit(address(s_connectionStore));
    emit ConnectionPolicyStore.InlineConnectionConfigSet(s_tokenB, OTHER_CHAIN_SELECTOR, abi.encode(s_remoteTokenB));
    s_poolConfigManager.applyInlineConnectionConfigUpdates(new IPoolConfigManager.LaneKey[](0), inlineConfigs);

    assertEq(abi.encode(s_remoteTokenA), _getLaneConnection(s_tokenA, OTHER_CHAIN_SELECTOR).remoteToken);
    assertEq(abi.encode(s_remoteTokenB), _getLaneConnection(s_tokenB, OTHER_CHAIN_SELECTOR).remoteToken);
  }

  /// @notice A removal opens no lane, so the lane check does not run for it. Token A has no fee record when its inline
  /// record is removed.
  function test_applyInlineConnectionConfigUpdates_RemovalsAreNotChecked() public {
    _setTokenFeePolicy(s_tokenA, bytes32(0));

    vm.expectEmit(address(s_connectionStore));
    emit ConnectionPolicyStore.InlineConnectionConfigRemoved(s_tokenA, DEST_CHAIN_SELECTOR);
    s_poolConfigManager.applyInlineConnectionConfigUpdates(
      _lanesOf(s_tokenA, DEST_CHAIN_SELECTOR), new ConnectionPolicyStore.InlineConnectionConfigArgs[](0)
    );

    assertFalse(_getLaneConnection(s_tokenA, DEST_CHAIN_SELECTOR).hasConnection);
  }

  /// @notice CONFIG_ROLE and TOKEN_MANAGER_ROLE each pass alone.
  function test_applyInlineConnectionConfigUpdates_AllowedRoles() public {
    _assertRolesPass(
      abi.encodeCall(
        PoolConfigManager.applyInlineConnectionConfigUpdates,
        (new IPoolConfigManager.LaneKey[](0), new ConnectionPolicyStore.InlineConnectionConfigArgs[](0))
      ),
      _getTokenWriteRoles()
    );
  }

  // Reverts

  /// @notice The lane check runs after the forward. An unknown token has no remote pool record, so a lane check before
  /// the forward would revert with RemotePoolRecordNotFound.
  function test_applyInlineConnectionConfigUpdates_RevertWhen_StoreRevertsBeforeLaneCheck() public {
    address unknownToken = makeAddr("unknownToken");
    ConnectionPolicyStore.InlineConnectionConfigArgs[] memory inlineConfigs =
      new ConnectionPolicyStore.InlineConnectionConfigArgs[](1);
    inlineConfigs[0] = ConnectionPolicyStore.InlineConnectionConfigArgs({
      localToken: unknownToken, remoteChainSelector: OTHER_CHAIN_SELECTOR, remoteToken: abi.encode(s_remoteTokenB)
    });

    vm.expectRevert(abi.encodeWithSelector(PolicyConfig.LocalTokenNotRegistered.selector, unknownToken));
    s_poolConfigManager.applyInlineConnectionConfigUpdates(new IPoolConfigManager.LaneKey[](0), inlineConfigs);
  }

  /// @notice The lane check reverts for the first type that the token misses. The remote pool type is checked first.
  function test_applyInlineConnectionConfigUpdates_RevertWhen_RemotePoolRecordNotFound() public {
    _setRecordPolicies(s_tokenB, bytes32(0), bytes32(0), bytes32(0), bytes32(0), bytes32(0));

    vm.expectRevert(
      abi.encodeWithSelector(ConnectionPolicyStore.RemotePoolRecordNotFound.selector, s_tokenB, OTHER_CHAIN_SELECTOR)
    );
    s_poolConfigManager.applyInlineConnectionConfigUpdates(
      new IPoolConfigManager.LaneKey[](0), _getInlineConfigsOfTokenB()
    );
  }

  /// @notice The fee type is checked after the remote pool type.
  function test_applyInlineConnectionConfigUpdates_RevertWhen_FeeRecordNotFound() public {
    _setRecordPolicies(s_tokenB, DEFAULT_POLICY, bytes32(0), bytes32(0), bytes32(0), bytes32(0));

    vm.expectRevert(abi.encodeWithSelector(FeePolicyStore.FeeRecordNotFound.selector, s_tokenB, OTHER_CHAIN_SELECTOR));
    s_poolConfigManager.applyInlineConnectionConfigUpdates(
      new IPoolConfigManager.LaneKey[](0), _getInlineConfigsOfTokenB()
    );
  }

  /// @notice The destination overhead type is checked after the fee type.
  function test_applyInlineConnectionConfigUpdates_RevertWhen_DestOverheadRecordNotFound() public {
    _setRecordPolicies(s_tokenB, DEFAULT_POLICY, DEFAULT_POLICY, bytes32(0), bytes32(0), bytes32(0));

    vm.expectRevert(
      abi.encodeWithSelector(FeePolicyStore.DestOverheadRecordNotFound.selector, s_tokenB, OTHER_CHAIN_SELECTOR)
    );
    s_poolConfigManager.applyInlineConnectionConfigUpdates(
      new IPoolConfigManager.LaneKey[](0), _getInlineConfigsOfTokenB()
    );
  }

  /// @notice The rate limit type is checked after the destination overhead type.
  function test_applyInlineConnectionConfigUpdates_RevertWhen_RateLimitRecordNotFound() public {
    _setRecordPolicies(s_tokenB, DEFAULT_POLICY, DEFAULT_POLICY, DEFAULT_POLICY, bytes32(0), bytes32(0));

    vm.expectRevert(
      abi.encodeWithSelector(SecurityPolicyStore.RateLimitRecordNotFound.selector, s_tokenB, OTHER_CHAIN_SELECTOR)
    );
    s_poolConfigManager.applyInlineConnectionConfigUpdates(
      new IPoolConfigManager.LaneKey[](0), _getInlineConfigsOfTokenB()
    );
  }

  /// @notice The security type is checked last.
  function test_applyInlineConnectionConfigUpdates_RevertWhen_SecurityRecordNotFound() public {
    _setRecordPolicies(s_tokenB, DEFAULT_POLICY, DEFAULT_POLICY, DEFAULT_POLICY, DEFAULT_POLICY, bytes32(0));

    vm.expectRevert(
      abi.encodeWithSelector(SecurityPolicyStore.SecurityRecordNotFound.selector, s_tokenB, OTHER_CHAIN_SELECTOR)
    );
    s_poolConfigManager.applyInlineConnectionConfigUpdates(
      new IPoolConfigManager.LaneKey[](0), _getInlineConfigsOfTokenB()
    );
  }

  /// @notice The lane check runs for each inline record of the call. The record of token A passes, and token B misses
  /// a fee record.
  function test_applyInlineConnectionConfigUpdates_RevertWhen_LaterInlineConfigMissesRecord() public {
    _setRecordPolicies(s_tokenB, DEFAULT_POLICY, bytes32(0), DEFAULT_POLICY, DEFAULT_POLICY, DEFAULT_POLICY);
    ConnectionPolicyStore.InlineConnectionConfigArgs[] memory inlineConfigs =
      new ConnectionPolicyStore.InlineConnectionConfigArgs[](2);
    inlineConfigs[0] = ConnectionPolicyStore.InlineConnectionConfigArgs({
      localToken: s_tokenA, remoteChainSelector: OTHER_CHAIN_SELECTOR, remoteToken: abi.encode(s_remoteTokenA)
    });
    inlineConfigs[1] = _getInlineConfigsOfTokenB()[0];

    vm.expectRevert(abi.encodeWithSelector(FeePolicyStore.FeeRecordNotFound.selector, s_tokenB, OTHER_CHAIN_SELECTOR));
    s_poolConfigManager.applyInlineConnectionConfigUpdates(new IPoolConfigManager.LaneKey[](0), inlineConfigs);
  }

  /// @notice Each other PCM role and a caller with no role revert. The error names the second role of the check.
  function test_applyInlineConnectionConfigUpdates_RevertWhen_MissingRole() public {
    _assertRolesRevert(
      abi.encodeCall(
        PoolConfigManager.applyInlineConnectionConfigUpdates,
        (new IPoolConfigManager.LaneKey[](0), new ConnectionPolicyStore.InlineConnectionConfigArgs[](0))
      ),
      _getTokenWriteRoles()
    );
  }

  /// @dev Sets the policies of the five types that the lane check reads, in the order of the check. Each DEFAULT policy
  /// has a policy default record, so it gives the token a record on every chain. Zero gives the token no record of that
  /// type.
  function _setRecordPolicies(
    address localToken,
    bytes32 remotePoolPolicyName,
    bytes32 feePolicyName,
    bytes32 destOverheadPolicyName,
    bytes32 rateLimitPolicyName,
    bytes32 securityPolicyName
  ) internal {
    _setTokenRemotePoolPolicy(localToken, remotePoolPolicyName);
    _setTokenFeePolicy(localToken, feePolicyName);
    _setTokenDestOverheadPolicy(localToken, destOverheadPolicyName);
    _setTokenRateLimitPolicy(localToken, rateLimitPolicyName);
    _setTokenSecurityPolicy(localToken, securityPolicyName);
  }

  /// @dev An inline connection record of token B to the other chain.
  function _getInlineConfigsOfTokenB()
    internal
    view
    returns (ConnectionPolicyStore.InlineConnectionConfigArgs[] memory inlineConfigs)
  {
    inlineConfigs = new ConnectionPolicyStore.InlineConnectionConfigArgs[](1);
    inlineConfigs[0] = ConnectionPolicyStore.InlineConnectionConfigArgs({
      localToken: s_tokenB, remoteChainSelector: OTHER_CHAIN_SELECTOR, remoteToken: abi.encode(s_remoteTokenB)
    });
    return inlineConfigs;
  }
}
