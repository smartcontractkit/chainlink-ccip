// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {PoolConfigManager} from "../../../../pools/MultiTokenPool/PoolConfigManager.sol";
import {ConnectionPolicyStore} from "../../../../pools/MultiTokenPool/policies/ConnectionPolicyStore.sol";
import {PoolConfigManagerRolesSetup} from "./PoolConfigManagerRolesSetup.t.sol";

contract PoolConfigManager_applyRemoteTokenAddressProviderUpdates is PoolConfigManagerRolesSetup {
  address internal s_provider = makeAddr("remoteTokenAddressProvider");

  /// @dev The connection store checks that a provider has code.
  function setUp() public virtual override {
    super.setUp();
    vm.etch(s_provider, hex"00");
  }

  /// @notice Updates providers for multiple policies in one call. The connection store emits the events.
  function test_applyRemoteTokenAddressProviderUpdates() public {
    _setConnectionPolicyConfig(OTHER_POLICY, OTHER_CHAIN_SELECTOR, false);
    ConnectionPolicyStore.RemoteTokenAddressProviderArgs[] memory providersToSet =
      new ConnectionPolicyStore.RemoteTokenAddressProviderArgs[](2);
    providersToSet[0] = _getProviderArgs(DEFAULT_POLICY);
    providersToSet[1] = _getProviderArgs(OTHER_POLICY);

    vm.expectEmit(address(s_connectionStore));
    emit ConnectionPolicyStore.RemoteTokenAddressProviderSet(DEFAULT_POLICY, s_provider);
    vm.expectEmit(address(s_connectionStore));
    emit ConnectionPolicyStore.RemoteTokenAddressProviderSet(OTHER_POLICY, s_provider);
    s_poolConfigManager.applyRemoteTokenAddressProviderUpdates(providersToSet);

    (, address storedProvider,) = s_poolConfigManager.getConnectionPolicy(DEFAULT_POLICY);
    assertEq(s_provider, storedProvider);
    (, storedProvider,) = s_poolConfigManager.getConnectionPolicy(OTHER_POLICY);
    assertEq(s_provider, storedProvider);
  }

  /// @notice Each allowed role authorizes the call on its own.
  function test_applyRemoteTokenAddressProviderUpdates_AllowedRoles() public {
    _assertRolesPass(
      abi.encodeCall(
        PoolConfigManager.applyRemoteTokenAddressProviderUpdates,
        (new ConnectionPolicyStore.RemoteTokenAddressProviderArgs[](0))
      ),
      _getPolicyWriteRoles(s_poolConfigManager.CONNECTION_ROLE())
    );
  }

  // Reverts

  /// @notice Calls without an allowed role revert.
  function test_applyRemoteTokenAddressProviderUpdates_RevertWhen_MissingRole() public {
    _assertRolesRevert(
      abi.encodeCall(
        PoolConfigManager.applyRemoteTokenAddressProviderUpdates,
        (new ConnectionPolicyStore.RemoteTokenAddressProviderArgs[](0))
      ),
      _getPolicyWriteRoles(s_poolConfigManager.CONNECTION_ROLE())
    );
  }

  function _getProviderArgs(
    bytes32 policyName
  ) internal view returns (ConnectionPolicyStore.RemoteTokenAddressProviderArgs memory) {
    return ConnectionPolicyStore.RemoteTokenAddressProviderArgs({
      policyName: policyName, remoteTokenAddressProvider: s_provider
    });
  }
}
