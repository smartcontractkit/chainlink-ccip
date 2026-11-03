// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ITokenAdminRegistry} from "../../../../interfaces/ITokenAdminRegistry.sol";

import {PoolConfigManager} from "../../../../pools/MultiTokenPool/PoolConfigManager.sol";
import {ConnectionPolicyStore} from "../../../../pools/MultiTokenPool/policies/ConnectionPolicyStore.sol";
import {FeePolicyStore} from "../../../../pools/MultiTokenPool/policies/FeePolicyStore.sol";
import {SecurityPolicyStore} from "../../../../pools/MultiTokenPool/policies/SecurityPolicyStore.sol";
import {PoolConfigManagerSetup} from "./PoolConfigManagerSetup.t.sol";

contract PoolConfigManager_constructor is PoolConfigManagerSetup {
  function test_constructor() public {
    vm.expectEmit();
    emit PoolConfigManager.StaticConfigSet(
      ITokenAdminRegistry(address(s_tokenAdminRegistry)), s_connectionStore, s_feeStore, s_securityStore
    );
    PoolConfigManager poolConfigManager = new PoolConfigManager(
      ITokenAdminRegistry(address(s_tokenAdminRegistry)), s_connectionStore, s_feeStore, s_securityStore, OWNER
    );

    assertEq(address(s_tokenAdminRegistry), poolConfigManager.getTokenAdminRegistry());
    assertEq(OWNER, poolConfigManager.defaultAdmin());
    assertEq("PoolConfigManager 2.1.0-dev", poolConfigManager.typeAndVersion());
  }

  // Reverts

  function test_constructor_RevertWhen_ZeroAddressInvalid() public {
    vm.expectRevert(PoolConfigManager.ZeroAddressInvalid.selector);
    new PoolConfigManager(ITokenAdminRegistry(address(0)), s_connectionStore, s_feeStore, s_securityStore, OWNER);

    vm.expectRevert(PoolConfigManager.ZeroAddressInvalid.selector);
    new PoolConfigManager(
      ITokenAdminRegistry(address(s_tokenAdminRegistry)),
      ConnectionPolicyStore(address(0)),
      s_feeStore,
      s_securityStore,
      OWNER
    );

    vm.expectRevert(PoolConfigManager.ZeroAddressInvalid.selector);
    new PoolConfigManager(
      ITokenAdminRegistry(address(s_tokenAdminRegistry)),
      s_connectionStore,
      FeePolicyStore(address(0)),
      s_securityStore,
      OWNER
    );

    vm.expectRevert(PoolConfigManager.ZeroAddressInvalid.selector);
    new PoolConfigManager(
      ITokenAdminRegistry(address(s_tokenAdminRegistry)),
      s_connectionStore,
      s_feeStore,
      SecurityPolicyStore(address(0)),
      OWNER
    );
  }
}
