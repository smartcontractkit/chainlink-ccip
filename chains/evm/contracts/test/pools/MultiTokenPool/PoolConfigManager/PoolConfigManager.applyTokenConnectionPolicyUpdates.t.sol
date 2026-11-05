// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {PolicyConfig} from "../../../../libraries/PolicyConfig.sol";
import {PoolConfigManager} from "../../../../pools/MultiTokenPool/PoolConfigManager.sol";
import {ConnectionPolicyStore} from "../../../../pools/MultiTokenPool/policies/ConnectionPolicyStore.sol";
import {PoolConfigManagerRolesSetup} from "./PoolConfigManagerRolesSetup.t.sol";

contract PoolConfigManager_applyTokenConnectionPolicyUpdates is PoolConfigManagerRolesSetup {
  address internal s_tokenC;

  function setUp() public virtual override {
    super.setUp();
    s_tokenC = _deployToken("C");
  }

  /// @notice A batch assigns token A to OTHER_POLICY and removes token B's policy assignment. The connection store
  /// emits both changes.
  function test_applyTokenConnectionPolicyUpdates() public {
    _setConnectionPolicyConfig(OTHER_POLICY, OTHER_CHAIN_SELECTOR, false);
    PolicyConfig.TokenPolicyArgs[] memory tokenPolicies = new PolicyConfig.TokenPolicyArgs[](2);
    tokenPolicies[0] = PolicyConfig.TokenPolicyArgs({localToken: s_tokenA, policyName: OTHER_POLICY});
    tokenPolicies[1] = PolicyConfig.TokenPolicyArgs({localToken: s_tokenB, policyName: bytes32(0)});

    vm.expectEmit(address(s_connectionStore));
    emit ConnectionPolicyStore.TokenConnectionPolicySet(s_tokenA, bytes32(0), OTHER_POLICY);
    vm.expectEmit(address(s_connectionStore));
    emit ConnectionPolicyStore.TokenConnectionPolicySet(s_tokenB, DEFAULT_POLICY, bytes32(0));
    s_poolConfigManager.applyTokenConnectionPolicyUpdates(tokenPolicies);

    assertEq(OTHER_POLICY, _getTokenConfig(s_tokenA).connectionPolicyName);
    assertEq(bytes32(0), _getTokenConfig(s_tokenB).connectionPolicyName);
  }

  /// @notice Assigning a connection policy does not check lane completeness. Token C gains a lane despite having no
  /// remote pool record.
  function test_applyTokenConnectionPolicyUpdates_LanesAreNotChecked() public {
    _setTokenProfile(_getTokenProfileArgs(s_tokenC));

    _setTokenConnectionPolicy(s_tokenC, DEFAULT_POLICY);

    assertTrue(_getLaneConnection(s_tokenC, DEST_CHAIN_SELECTOR).hasConnection);
    (, PolicyConfig.RecordLevel[] memory remotePoolLevels,,,,) =
      s_poolConfigManager.getLaneRecordLevelBatch(_lanesOf(s_tokenC, DEST_CHAIN_SELECTOR));
    assertEq(uint8(PolicyConfig.RecordLevel.None), uint8(remotePoolLevels[0]));
  }

  /// @notice Each allowed role authorizes the call on its own.
  function test_applyTokenConnectionPolicyUpdates_AllowedRoles() public {
    _assertRolesPass(
      abi.encodeCall(PoolConfigManager.applyTokenConnectionPolicyUpdates, (new PolicyConfig.TokenPolicyArgs[](0))),
      _getTokenWriteRoles(s_poolConfigManager.CONNECTION_ROLE())
    );
  }

  // Reverts

  /// @notice Calls without an allowed role revert.
  function test_applyTokenConnectionPolicyUpdates_RevertWhen_MissingRole() public {
    _assertRolesRevert(
      abi.encodeCall(PoolConfigManager.applyTokenConnectionPolicyUpdates, (new PolicyConfig.TokenPolicyArgs[](0))),
      _getTokenWriteRoles(s_poolConfigManager.CONNECTION_ROLE())
    );
  }
}
