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

  /// @dev Remote pool policies have no default records. Add an explicit record so lanes on the other chain pass
  /// validation.
  function setUp() public virtual override {
    super.setUp();
    _setRemotePoolPolicyConfig(DEFAULT_POLICY, OTHER_CHAIN_SELECTOR, _getRemotePoolAddresses(abi.encode(s_remotePool)));
  }

  /// @notice Creates lanes for multiple tokens in one call. Each token has the other records required for its new
  /// lane.
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

  /// @notice Removing an inline connection does not check lane completeness. Token A can remove its connection
  /// despite having no fee record.
  function test_applyInlineConnectionConfigUpdates_RemovalsAreNotChecked() public {
    _setTokenFeePolicy(s_tokenA, bytes32(0));

    vm.expectEmit(address(s_connectionStore));
    emit ConnectionPolicyStore.InlineConnectionConfigRemoved(s_tokenA, DEST_CHAIN_SELECTOR);
    s_poolConfigManager.applyInlineConnectionConfigUpdates(
      _lanesOf(s_tokenA, DEST_CHAIN_SELECTOR), new ConnectionPolicyStore.InlineConnectionConfigArgs[](0)
    );

    assertFalse(_getLaneConnection(s_tokenA, DEST_CHAIN_SELECTOR).hasConnection);
  }

  /// @notice Each allowed role authorizes the call on its own.
  function test_applyInlineConnectionConfigUpdates_AllowedRoles() public {
    _assertRolesPass(
      abi.encodeCall(
        PoolConfigManager.applyInlineConnectionConfigUpdates,
        (new IPoolConfigManager.LaneKey[](0), new ConnectionPolicyStore.InlineConnectionConfigArgs[](0))
      ),
      _getTokenWriteRoles(s_poolConfigManager.CONNECTION_ROLE())
    );
  }

  // Reverts

  /// @notice The connection store rejects an unregistered token before the manager checks for a remote pool record.
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

  /// @notice Lane validation reports the first missing record. The remote pool record is checked first.
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

  /// @notice Every inline connection is checked. Token A has all required records, but token B has no fee record.
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

  /// @notice Calls without an allowed role revert.
  function test_applyInlineConnectionConfigUpdates_RevertWhen_MissingRole() public {
    _assertRolesRevert(
      abi.encodeCall(
        PoolConfigManager.applyInlineConnectionConfigUpdates,
        (new IPoolConfigManager.LaneKey[](0), new ConnectionPolicyStore.InlineConnectionConfigArgs[](0))
      ),
      _getTokenWriteRoles(s_poolConfigManager.CONNECTION_ROLE())
    );
  }

  /// @dev Assigns policies in lane validation order. The remote pool policy has an explicit chain record; the other
  /// policies use default records. A zero policy name leaves that record type missing.
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
