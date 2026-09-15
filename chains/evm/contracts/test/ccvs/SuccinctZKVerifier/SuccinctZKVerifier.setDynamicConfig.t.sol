// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {SuccinctZKVerifier} from "../../../ccvs/SuccinctZKVerifier.sol";
import {SuccinctZKVerifierSetup} from "./SuccinctZKVerifierSetup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract SuccinctZKVerifier_setDynamicConfig is SuccinctZKVerifierSetup {
  function test_setDynamicConfig() public {
    SuccinctZKVerifier.DynamicConfig memory dynamicConfig =
      SuccinctZKVerifier.DynamicConfig({feeAggregator: makeAddr("newFeeAggregator")});

    vm.expectEmit();
    emit SuccinctZKVerifier.DynamicConfigSet(dynamicConfig);

    s_zkVerifier.setDynamicConfig(dynamicConfig);

    assertEq(dynamicConfig.feeAggregator, s_zkVerifier.getDynamicConfig().feeAggregator);
  }

  // Reverts

  function test_setDynamicConfig_RevertWhen_OnlyCallableByOwner() public {
    vm.stopPrank();
    vm.startPrank(STRANGER);

    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_zkVerifier.setDynamicConfig(SuccinctZKVerifier.DynamicConfig({feeAggregator: FEE_AGGREGATOR}));
  }
}
