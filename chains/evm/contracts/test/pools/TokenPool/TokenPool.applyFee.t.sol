// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolV2} from "../../../interfaces/IPoolV2.sol";

import {Pool} from "../../../libraries/Pool.sol";
import {TokenPool} from "../../../pools/TokenPool.sol";
import {TokenPoolV2Setup} from "./TokenPoolV2Setup.t.sol";

contract TokenPoolV2_applyFee is TokenPoolV2Setup {
  function test_applyFee_CustomFinality() public {
    uint16 minBlockConfirmation = 5;
    uint16 defaultBlockConfirmationTransferFeeBps = 100;
    uint16 customBlockConfirmationTransferFeeBps = 500;
    uint256 amount = 1_000e18;
    vm.startPrank(OWNER);
    s_tokenPool.applyCustomBlockConfirmationConfigUpdates(
      minBlockConfirmation, new TokenPool.CustomBlockConfirmationRateLimitConfigArgs[](0)
    );
    TokenPool.TokenTransferFeeConfigArgs[] memory feeConfigArgs = new TokenPool.TokenTransferFeeConfigArgs[](1);
    feeConfigArgs[0] = TokenPool.TokenTransferFeeConfigArgs({
      destChainSelector: DEST_CHAIN_SELECTOR,
      tokenTransferFeeConfig: IPoolV2.TokenTransferFeeConfig({
        destGasOverhead: 50_000,
        destBytesOverhead: Pool.CCIP_LOCK_OR_BURN_V1_RET_BYTES,
        defaultBlockConfirmationFeeUSDCents: 0,
        customBlockConfirmationFeeUSDCents: 0,
        defaultBlockConfirmationTransferFeeBps: defaultBlockConfirmationTransferFeeBps,
        customBlockConfirmationTransferFeeBps: customBlockConfirmationTransferFeeBps,
        isEnabled: true
      })
    });
    s_tokenPool.applyTokenTransferFeeConfigUpdates(feeConfigArgs, new uint64[](0));

    Pool.LockOrBurnInV1 memory lockOrBurnIn = Pool.LockOrBurnInV1({
      originalSender: s_sender,
      receiver: s_receiver,
      amount: amount,
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      localToken: address(s_token)
    });

    uint256 amountAfterFee = s_tokenPool.applyFee(lockOrBurnIn, minBlockConfirmation);
    assertEq(amountAfterFee, amount - ((amount * customBlockConfirmationTransferFeeBps) / BPS_DIVIDER));
  }

  function test_applyFee_DefaultFinality() public {
    uint16 defaultBlockConfirmationTransferFeeBps = 250; // 2.5%
    uint256 amount = 1_000e18;

    vm.startPrank(OWNER);
    TokenPool.TokenTransferFeeConfigArgs[] memory feeConfigArgs = new TokenPool.TokenTransferFeeConfigArgs[](1);
    feeConfigArgs[0] = TokenPool.TokenTransferFeeConfigArgs({
      destChainSelector: DEST_CHAIN_SELECTOR,
      tokenTransferFeeConfig: IPoolV2.TokenTransferFeeConfig({
        destGasOverhead: 50_000,
        destBytesOverhead: Pool.CCIP_LOCK_OR_BURN_V1_RET_BYTES,
        defaultBlockConfirmationFeeUSDCents: 0,
        customBlockConfirmationFeeUSDCents: 0,
        defaultBlockConfirmationTransferFeeBps: defaultBlockConfirmationTransferFeeBps,
        customBlockConfirmationTransferFeeBps: 0,
        isEnabled: true
      })
    });
    s_tokenPool.applyTokenTransferFeeConfigUpdates(feeConfigArgs, new uint64[](0));

    Pool.LockOrBurnInV1 memory lockOrBurnIn = Pool.LockOrBurnInV1({
      originalSender: s_sender,
      receiver: s_receiver,
      amount: amount,
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      localToken: address(s_token)
    });

    uint256 amountAfterFee = s_tokenPool.applyFee(lockOrBurnIn, 0);
    assertEq(amountAfterFee, amount - ((amount * defaultBlockConfirmationTransferFeeBps) / BPS_DIVIDER));
  }
}
