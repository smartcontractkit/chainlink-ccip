// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../../interfaces/IPoolConfigManager.sol";

import {FinalityCodec} from "../../../libraries/FinalityCodec.sol";
import {MultiTokenPool} from "../../../pools/MultiTokenPool/MultiTokenPool.sol";
import {MultiTokenPoolSetup} from "./MultiTokenPoolSetup.t.sol";

contract MultiTokenPool_getFee is MultiTokenPoolSetup {
  function test_getFee() public {
    // The mock matches only the lane and the requested finality that the pool must forward.
    bytes4 finality = FinalityCodec._encodeBlockDepth(5);
    _mockLaneFee(
      finality,
      IPoolConfigManager.LaneFee({
        feeUSDCents: 100, destGasOverhead: 200_000, destBytesOverhead: 256, tokenFeeBps: 10, isEnabled: true
      })
    );

    (uint256 feeUSDCents, uint32 destGasOverhead, uint32 destBytesOverhead, uint16 tokenFeeBps, bool isEnabled) =
      s_pool.getFee(address(s_token), DEST_CHAIN_SELECTOR, 1e18, address(0), finality, "");

    assertEq(feeUSDCents, 100);
    assertEq(destGasOverhead, 200_000);
    assertEq(destBytesOverhead, 256);
    assertEq(tokenFeeBps, 10);
    assertTrue(isEnabled);
  }

  // Reverts

  /// @notice Missing fee configuration must revert instead of requesting FeeQuoter's token transfer fees.
  function test_getFee_RevertWhen_LaneFeeNotFound() public {
    IPoolConfigManager.LaneFee memory noFee;
    _mockLaneFee(FinalityCodec.WAIT_FOR_FINALITY_FLAG, noFee);

    vm.expectRevert(
      abi.encodeWithSelector(MultiTokenPool.LaneFeeNotFound.selector, address(s_token), DEST_CHAIN_SELECTOR)
    );
    s_pool.getFee(address(s_token), DEST_CHAIN_SELECTOR, 1e18, address(0), FinalityCodec.WAIT_FOR_FINALITY_FLAG, "");
  }

  function _mockLaneFee(
    bytes4 requestedFinalityConfig,
    IPoolConfigManager.LaneFee memory fee
  ) internal {
    IPoolConfigManager.LaneFee[] memory fees = new IPoolConfigManager.LaneFee[](1);
    fees[0] = fee;
    vm.mockCall(
      s_poolConfigManager,
      abi.encodeCall(IPoolConfigManager.getFeeBatch, (_laneKeys(), requestedFinalityConfig)),
      abi.encode(fees)
    );
  }
}
