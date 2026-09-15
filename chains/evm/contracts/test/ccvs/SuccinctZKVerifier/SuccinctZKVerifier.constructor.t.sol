// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {SuccinctZKVerifier} from "../../../ccvs/SuccinctZKVerifier.sol";
import {BaseVerifier} from "../../../ccvs/components/BaseVerifier.sol";
import {SuccinctZKVerifierSetup} from "./SuccinctZKVerifierSetup.t.sol";

contract SuccinctZKVerifier_constructor is SuccinctZKVerifierSetup {
  function test_constructor() public {
    SuccinctZKVerifier.DynamicConfig memory dynamicConfig =
      SuccinctZKVerifier.DynamicConfig({feeAggregator: FEE_AGGREGATOR});

    vm.expectEmit();
    emit SuccinctZKVerifier.DynamicConfigSet(dynamicConfig);

    SuccinctZKVerifier zkVerifier =
      new SuccinctZKVerifier(dynamicConfig, s_storageLocations, address(s_mockRMNRemote), VERSION_TAG_V0_0_1);

    assertEq("SuccinctZKVerifier 0.0.1-dev", zkVerifier.typeAndVersion());
    assertEq(VERSION_TAG_V0_0_1, zkVerifier.versionTag());
    assertEq(FEE_AGGREGATOR, zkVerifier.getDynamicConfig().feeAggregator);
    assertEq(OWNER, zkVerifier.owner());
    assertEq(s_storageLocations, zkVerifier.getStorageLocations());
  }

  // Reverts

  function test_constructor_RevertWhen_ZeroAddressNotAllowed() public {
    vm.expectRevert(BaseVerifier.ZeroAddressNotAllowed.selector);
    new SuccinctZKVerifier(
      SuccinctZKVerifier.DynamicConfig({feeAggregator: FEE_AGGREGATOR}),
      s_storageLocations,
      address(0),
      VERSION_TAG_V0_0_1
    );
  }

  function test_constructor_RevertWhen_VersionTagCannotBeZero() public {
    vm.expectRevert(BaseVerifier.VersionTagCannotBeZero.selector);
    new SuccinctZKVerifier(
      SuccinctZKVerifier.DynamicConfig({feeAggregator: FEE_AGGREGATOR}),
      s_storageLocations,
      address(s_mockRMNRemote),
      bytes4(0)
    );
  }
}
