// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ITokenAdminRegistry} from "../../../../interfaces/ITokenAdminRegistry.sol";

import {ConnectionPolicyStore} from "../../../../pools/MultiTokenPool/policies/ConnectionPolicyStore.sol";
import {FeePolicyStore} from "../../../../pools/MultiTokenPool/policies/FeePolicyStore.sol";
import {SecurityPolicyStore} from "../../../../pools/MultiTokenPool/policies/SecurityPolicyStore.sol";
import {PoolConfigManagerSetup} from "./PoolConfigManagerSetup.t.sol";

contract PoolConfigManager_getStaticConfig is PoolConfigManagerSetup {
  /// @notice The static config returns the four addresses that the constructor set.
  function test_getStaticConfig() public view {
    (
      ITokenAdminRegistry tokenAdminRegistry,
      ConnectionPolicyStore connectionStore,
      FeePolicyStore feeStore,
      SecurityPolicyStore securityStore
    ) = s_poolConfigManager.getStaticConfig();

    assertEq(address(s_tokenAdminRegistry), address(tokenAdminRegistry));
    assertEq(address(s_connectionStore), address(connectionStore));
    assertEq(address(s_feeStore), address(feeStore));
    assertEq(address(s_securityStore), address(securityStore));
  }

  /// @notice The token admin registry view returns the same registry.
  function test_getTokenAdminRegistry() public view {
    assertEq(address(s_tokenAdminRegistry), s_poolConfigManager.getTokenAdminRegistry());
  }
}
