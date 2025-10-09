// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolV2} from "../../../interfaces/IPoolV2.sol";

import {Pool} from "../../../libraries/Pool.sol";
import {TokenPool} from "../../../pools/TokenPool.sol";
import {TokenPoolV2Setup} from "./TokenPoolV2Setup.t.sol";

contract TokenPoolV2_applyFee is TokenPoolV2Setup {
  function test_applyFee_DefaultFinality() public {
    uint16 finalityThreshold = 5;
    uint16 defaultFinalityTransferFeeBps = 200;
    uint256 maxAmountPerRequest = 1000e18;
    uint256 amount = 1000e18;
    vm.startPrank(OWNER);
    s_tokenPool.applyFinalityConfigUpdates(finalityThreshold, new TokenPool.CustomFinalityRateLimitConfigArgs[](0));
    // Set a fee config with default finality fee
    _applyTokenTransferFeeConfigUpdates(0, defaultFinalityTransferFeeBps);

    Pool.LockOrBurnInV1 memory lockOrBurnIn = Pool.LockOrBurnInV1({
      originalSender: s_sender,
      receiver: s_receiver,
      amount: amount,
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      localToken: address(s_token)
    });

    uint256 amountAfterFee = s_tokenPool.applyFee(lockOrBurnIn, 0);
    assertEq(amountAfterFee, amount - ((amount * defaultFinalityTransferFeeBps) / BPS_DIVIDER));
  }

  function test_applyFee_CustomFinality() public {
    uint16 finalityThreshold = 5;
    uint16 customFinalityTransferFeeBps = 500;
    uint256 amount = 1000e18;
    vm.startPrank(OWNER);
    s_tokenPool.applyFinalityConfigUpdates(finalityThreshold, new TokenPool.CustomFinalityRateLimitConfigArgs[](0));
    // Set a fee config with custom finality fee
    _applyTokenTransferFeeConfigUpdates(customFinalityTransferFeeBps, 0);

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

  function _applyTokenTransferFeeConfigUpdates(
    uint16 customFinalityTransferFeeBps,
    uint16 defaultFinalityTransferFeeBps
  ) internal {
    IPoolV2.TokenTransferFeeConfig memory feeConfig = IPoolV2.TokenTransferFeeConfig({
      destGasOverhead: 50000,
      destBytesOverhead: 32,
      feeUSDCents: 0,
      customFinalityTransferFeeBps: customFinalityTransferFeeBps, // 0.50%
      defaultFinalityTransferFeeBps: defaultFinalityTransferFeeBps, // 0.20%
      isEnabled: true
    });

    TokenPool.TokenTransferFeeConfigArgs[] memory feeConfigArgs = new TokenPool.TokenTransferFeeConfigArgs[](1);
    feeConfigArgs[0] =
      TokenPool.TokenTransferFeeConfigArgs({destChainSelector: DEST_CHAIN_SELECTOR, tokenTransferFeeConfig: feeConfig});

    s_tokenPool.applyTokenTransferFeeConfigUpdates(feeConfigArgs, new uint64[](0));
  }
}
