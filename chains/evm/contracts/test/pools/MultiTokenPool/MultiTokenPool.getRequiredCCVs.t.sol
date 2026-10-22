// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolV2} from "../../../interfaces/IPoolV2.sol";

import {MultiTokenPoolSetup} from "./MultiTokenPoolSetup.t.sol";

contract MultiTokenPool_getRequiredCCVs is MultiTokenPoolSetup {
  function test_getRequiredCCVs() public {
    address[] memory requiredCCVs = new address[](2);
    requiredCCVs[0] = makeAddr("ccv1");
    requiredCCVs[1] = makeAddr("ccv2");
    s_poolConfigManager.setRequiredCCVs(requiredCCVs);

    address[] memory returned = s_pool.getRequiredCCVs(
      address(s_token), DEST_CHAIN_SELECTOR, 1e18, bytes4(0), "", IPoolV2.MessageDirection.Outbound
    );

    assertEq(returned.length, 2);
    assertEq(returned[0], requiredCCVs[0]);
    assertEq(returned[1], requiredCCVs[1]);
  }
}
