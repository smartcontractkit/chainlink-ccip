// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {TokenPool} from "../../../poolsV2/TokenPool.sol";
import {TokenPoolV2Setup} from "./TokenPoolV2Setup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract TokenPoolV2_applyFinalityConfigUpdates is TokenPoolV2Setup {
  function test_applyFinalityConfigUpdates() public {
    TokenPool.FastFinalityConfig memory config = TokenPool.FastFinalityConfig({
      finalityThreshold: 100,
      fastTransferFeeBps: 500, // 5%
      maxAmountPerRequest: 1000e18
    });

    vm.expectEmit();
    emit TokenPool.FinalityConfigUpdated(
      config.finalityThreshold, config.fastTransferFeeBps, config.maxAmountPerRequest
    );
    s_tokenPool.applyFinalityConfigUpdates(config);
  }

  // Reverts
  function test_applyFinalityConfigUpdates_RevertWhen_OnlyCallableByOwner() public {
    vm.stopPrank();
    vm.prank(STRANGER);

    TokenPool.FastFinalityConfig memory config =
      TokenPool.FastFinalityConfig({finalityThreshold: 100, fastTransferFeeBps: 500, maxAmountPerRequest: 1000e18});

    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_tokenPool.applyFinalityConfigUpdates(config);
  }
}
