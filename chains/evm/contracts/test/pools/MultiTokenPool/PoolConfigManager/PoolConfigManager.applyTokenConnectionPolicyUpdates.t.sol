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

  /// @notice CONFIG_ROLE sets the connection policy of several tokens in one call. Token A joins the OTHER policy, and
  /// token B leaves its policy. The connection store applies the call and emits the events.
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

  /// @notice The PCM runs no lane check when a token gets a connection policy, the same as for a connection policy
  /// record. Token C has no policy of another type, so the new lane has no remote pool record.
  function test_applyTokenConnectionPolicyUpdates_LanesAreNotChecked() public {
    _setTokenProfile(_getTokenProfileArgs(s_tokenC));

    _setTokenConnectionPolicy(s_tokenC, DEFAULT_POLICY);

    assertTrue(_getLaneConnection(s_tokenC, DEST_CHAIN_SELECTOR).hasConnection);
    (, PolicyConfig.RecordLevel[] memory remotePoolLevels,,,,) =
      s_poolConfigManager.getLaneRecordLevelBatch(_lanesOf(s_tokenC, DEST_CHAIN_SELECTOR));
    assertEq(uint8(PolicyConfig.RecordLevel.None), uint8(remotePoolLevels[0]));
  }

  /// @notice CONFIG_ROLE and TOKEN_MANAGER_ROLE each pass alone.
  function test_applyTokenConnectionPolicyUpdates_AllowedRoles() public {
    _assertRolesPass(
      abi.encodeCall(PoolConfigManager.applyTokenConnectionPolicyUpdates, (new PolicyConfig.TokenPolicyArgs[](0))),
      _getTokenWriteRoles()
    );
  }

  // Reverts

  /// @notice Each other PCM role and a caller with no role revert. The error names the second role of the check.
  function test_applyTokenConnectionPolicyUpdates_RevertWhen_MissingRole() public {
    _assertRolesRevert(
      abi.encodeCall(PoolConfigManager.applyTokenConnectionPolicyUpdates, (new PolicyConfig.TokenPolicyArgs[](0))),
      _getTokenWriteRoles()
    );
  }
}
