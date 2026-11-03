// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ConnectionPolicyStore} from "../../../../../pools/MultiTokenPool/policies/ConnectionPolicyStore.sol";
import {ConnectionPolicyStoreSetup} from "./ConnectionPolicyStoreSetup.t.sol";

contract ConnectionPolicyStore_getConnectionPolicy is ConnectionPolicyStoreSetup {
  function test_getConnectionPolicy() public {
    _setConnectionPolicyConfig(_getConnectionPolicyConfigArgs(CONNECTION_POLICY, SOURCE_CHAIN_SELECTOR, true));
    ConnectionPolicyStore.ConnectionPolicyConfigArgs[] memory expectedConfigs =
      new ConnectionPolicyStore.ConnectionPolicyConfigArgs[](2);
    expectedConfigs[0] = _getConnectionPolicyConfigArgs(CONNECTION_POLICY, DEST_CHAIN_SELECTOR, false)[0];
    expectedConfigs[1] = _getConnectionPolicyConfigArgs(CONNECTION_POLICY, SOURCE_CHAIN_SELECTOR, true)[0];

    (
      uint32 tokenCount,
      address remoteTokenAddressProvider,
      ConnectionPolicyStore.ConnectionPolicyConfigArgs[] memory policyConfigs
    ) = s_connectionStore.getConnectionPolicy(CONNECTION_POLICY);

    assertEq(1, tokenCount);
    assertEq(address(s_remoteTokenAddressProvider), remoteTokenAddressProvider);
    assertEq(abi.encode(expectedConfigs), abi.encode(policyConfigs));
  }

  function test_getConnectionPolicy_UnknownPolicy() public view {
    (
      uint32 tokenCount,
      address remoteTokenAddressProvider,
      ConnectionPolicyStore.ConnectionPolicyConfigArgs[] memory policyConfigs
    ) = s_connectionStore.getConnectionPolicy(OTHER_POLICY);

    assertEq(0, tokenCount);
    assertEq(address(0), remoteTokenAddressProvider);
    assertEq(0, policyConfigs.length);
  }
}
