// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {FeePolicyStore} from "../../../../../pools/MultiTokenPool/policies/FeePolicyStore.sol";
import {FeePolicyStoreSetup} from "./FeePolicyStoreSetup.t.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

contract FeePolicyStore_constructor is FeePolicyStoreSetup {
  function test_constructor() public {
    vm.startPrank(OWNER);
    vm.expectEmit();
    emit AuthorizedCallers.AuthorizedCallerAdded(s_poolConfigManager);
    FeePolicyStore feeStore = new FeePolicyStore(_arrayOf(s_poolConfigManager));

    assertEq("FeePolicyStore 2.1.0-dev", feeStore.typeAndVersion());
    assertEq(OWNER, feeStore.owner());
    assertEq(_arrayOf(s_poolConfigManager), feeStore.getAllAuthorizedCallers());
  }
}
