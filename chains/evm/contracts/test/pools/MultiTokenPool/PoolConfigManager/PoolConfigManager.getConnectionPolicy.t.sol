// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ConnectionPolicyStore} from "../../../../pools/MultiTokenPool/policies/ConnectionPolicyStore.sol";
import {PoolConfigManagerSetup} from "./PoolConfigManagerSetup.t.sol";

/// @notice Token B uses the DEFAULT connection policy. Token A has no connection policy.
contract PoolConfigManager_getConnectionPolicy is PoolConfigManagerSetup {
  /// @notice Returns the token count, remote token address provider, and records for each policy. A missing
  /// provider returns the zero address.
  function test_getConnectionPolicy() public {
    address remoteTokenAddressProvider = makeAddr("remoteTokenAddressProvider");
    // A provider needs code.
    vm.etch(remoteTokenAddressProvider, hex"00");
    _setConnectionPolicyConfig(DEFAULT_POLICY, OTHER_CHAIN_SELECTOR, true);
    _setConnectionPolicyConfig(OTHER_POLICY, OTHER_CHAIN_SELECTOR, false);
    _setRemoteTokenAddressProvider(DEFAULT_POLICY, remoteTokenAddressProvider);
    ConnectionPolicyStore.ConnectionPolicyConfigArgs[] memory expectedConfigs =
      new ConnectionPolicyStore.ConnectionPolicyConfigArgs[](2);
    expectedConfigs[0] = ConnectionPolicyStore.ConnectionPolicyConfigArgs({
      policyName: DEFAULT_POLICY, remoteChainSelector: DEST_CHAIN_SELECTOR, onlySubsetTokens: false
    });
    expectedConfigs[1] = ConnectionPolicyStore.ConnectionPolicyConfigArgs({
      policyName: DEFAULT_POLICY, remoteChainSelector: OTHER_CHAIN_SELECTOR, onlySubsetTokens: true
    });

    (
      uint32 tokenCount,
      address storedProvider,
      ConnectionPolicyStore.ConnectionPolicyConfigArgs[] memory policyConfigs
    ) = s_poolConfigManager.getConnectionPolicy(DEFAULT_POLICY);
    assertEq(1, tokenCount);
    assertEq(remoteTokenAddressProvider, storedProvider);
    assertEq(abi.encode(expectedConfigs), abi.encode(policyConfigs));

    (tokenCount, storedProvider, policyConfigs) = s_poolConfigManager.getConnectionPolicy(OTHER_POLICY);
    assertEq(0, tokenCount);
    assertEq(address(0), storedProvider);
    assertEq(1, policyConfigs.length);
    assertEq(OTHER_POLICY, policyConfigs[0].policyName);
    assertEq(OTHER_CHAIN_SELECTOR, policyConfigs[0].remoteChainSelector);
    assertFalse(policyConfigs[0].onlySubsetTokens);
  }
}
