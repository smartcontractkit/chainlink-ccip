// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Pool} from "../../../libraries/Pool.sol";
import {TokenPool} from "../../../pools/TokenPool.sol";
import {TokenPoolV2Setup} from "./TokenPoolV2Setup.t.sol";

contract TokenPoolV2_applyFee is TokenPoolV2Setup {
  function test_applyFee_CustomFinality() public {
    uint16 finalityThreshold = 5;
    uint16 customFinalityTransferFeeBps = 500;
    uint256 amount = 1000e18;
    vm.startPrank(OWNER);
    s_tokenPool.applyFinalityConfigUpdates(
      finalityThreshold, customFinalityTransferFeeBps, new TokenPool.CustomFinalityRateLimitConfigArgs[](0)
    );

    Pool.LockOrBurnInV1 memory lockOrBurnIn = Pool.LockOrBurnInV1({
      originalSender: s_sender,
      receiver: s_receiver,
      amount: amount,
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      localToken: address(s_token)
    });

    uint256 amountAfterFee = s_tokenPool.applyFee(lockOrBurnIn, finalityThreshold);
    assertEq(amountAfterFee, amount - ((amount * customFinalityTransferFeeBps) / BPS_DIVIDER));
  }

  function test_applyFee_NoFee() public view {
    uint256 amount = 1000e18;
    Pool.LockOrBurnInV1 memory lockOrBurnIn = Pool.LockOrBurnInV1({
      originalSender: s_sender,
      receiver: s_receiver,
      amount: amount,
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      localToken: address(s_token)
    });

    uint256 amountAfterFee = s_tokenPool.applyFee(lockOrBurnIn, 0);
    assertEq(amountAfterFee, amount);
  }
}
