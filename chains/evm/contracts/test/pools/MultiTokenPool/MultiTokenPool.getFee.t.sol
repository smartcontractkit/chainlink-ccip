// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../../interfaces/IPoolConfigManager.sol";

import {MultiTokenPoolSetup} from "./MultiTokenPoolSetup.t.sol";

contract MultiTokenPool_getFee is MultiTokenPoolSetup {
  function test_getFee() public {
    s_poolConfigManager.setLaneFee(
      IPoolConfigManager.LaneFee({
        feeUSDCents: 100, destGasOverhead: 200_000, destBytesOverhead: 256, tokenFeeBps: 10, isEnabled: true
      })
    );

    (uint256 feeUSDCents, uint32 destGasOverhead, uint32 destBytesOverhead, uint16 tokenFeeBps, bool isEnabled) =
      s_pool.getFee(address(s_token), DEST_CHAIN_SELECTOR, 1e18, address(0), bytes4(0), "");

    assertEq(feeUSDCents, 100);
    assertEq(destGasOverhead, 200_000);
    assertEq(destBytesOverhead, 256);
    assertEq(tokenFeeBps, 10);
    assertTrue(isEnabled);
  }
}
