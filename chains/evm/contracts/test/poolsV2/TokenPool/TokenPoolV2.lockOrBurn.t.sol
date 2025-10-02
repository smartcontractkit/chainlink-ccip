// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Pool} from "../../../libraries/Pool.sol";
import {TokenPool} from "../../../poolsV2/TokenPool.sol";
import {TokenPoolV2Setup} from "./TokenPoolV2Setup.t.sol";

contract TokenPoolV2_lockOrBurn is TokenPoolV2Setup {
  function test_lockOrBurn() public {
    uint256 amount = 1000e18;
    s_token.transfer(address(s_tokenPool), amount);

    Pool.LockOrBurnInV1 memory lockOrBurnIn = Pool.LockOrBurnInV1({
      originalSender: s_sender,
      receiver: s_receiver,
      amount: amount,
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      localToken: address(s_token)
    });

    vm.startPrank(s_allowedOnRamp);
    (Pool.LockOrBurnOutV1 memory lockOrBurnOut, uint256 destTokenAmount) = s_tokenPool.lockOrBurn(lockOrBurnIn, 0, "");

    assertEq(destTokenAmount, amount);
    assertEq(lockOrBurnOut.destTokenAddress, abi.encode(s_initialRemoteToken));
    assertEq(lockOrBurnOut.destPoolData, abi.encode(s_token.decimals()));
  }

  function test_lockOrBurn_WithFinalityConfig() public {
    // Set up finality config
    uint16 finalityThreshold = 8;
    uint16 fastTransferFeeBps = 500; // 5%
    uint256 maxAmountPerRequest = 1000e18;
    s_tokenPool.applyFinalityConfigUpdates(
      finalityThreshold, fastTransferFeeBps, maxAmountPerRequest, new TokenPool.FastTransferRateLimitConfigArgs[](0)
    );
    uint256 amount = 1000e18;
    s_token.transfer(address(s_tokenPool), amount);

    Pool.LockOrBurnInV1 memory lockOrBurnIn = Pool.LockOrBurnInV1({
      originalSender: s_sender,
      receiver: s_receiver,
      amount: amount,
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      localToken: address(s_token)
    });

    vm.startPrank(s_allowedOnRamp);
    (Pool.LockOrBurnOutV1 memory lockOrBurnOut, uint256 destTokenAmount) =
      s_tokenPool.lockOrBurn(lockOrBurnIn, finalityThreshold, "");

    // Should deduct 5% fee
    uint256 expectedDestAmount = amount - (amount * fastTransferFeeBps / BPS_DEVIDER);
    assertEq(destTokenAmount, expectedDestAmount);
    assertEq(lockOrBurnOut.destTokenAddress, abi.encode(s_initialRemoteToken));
  }

  // Reverts

  function test_lockOrBurn_RevertWhen_InvalidFinality() public {
    // Set up finality config
    uint16 finalityThreshold = 5;
    uint16 fastTransferFeeBps = 500;
    uint256 maxAmountPerRequest = 1000e18;

    s_tokenPool.applyFinalityConfigUpdates(
      finalityThreshold, fastTransferFeeBps, maxAmountPerRequest, new TokenPool.FastTransferRateLimitConfigArgs[](0)
    );

    uint256 amount = 1000e18;
    s_token.transfer(address(s_tokenPool), amount);

    Pool.LockOrBurnInV1 memory lockOrBurnIn = Pool.LockOrBurnInV1({
      originalSender: s_sender,
      receiver: s_receiver,
      amount: amount,
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      localToken: address(s_token)
    });

    // Finality below threshold should revert
    vm.expectRevert(abi.encodeWithSelector(TokenPool.InvalidFinality.selector, finalityThreshold - 1, 5));
    vm.startPrank(s_allowedOnRamp);
    s_tokenPool.lockOrBurn(lockOrBurnIn, finalityThreshold - 1, "");
  }

  function test_lockOrBurn_RevertWhen_AmountExceedsMaxPerRequest() public {
    // Set up finality config with low max amount
    uint16 finalityThreshold = 8;
    uint16 fastTransferFeeBps = 500;
    uint256 maxAmountPerRequest = 500e18; // Lower than our test amount.
    s_tokenPool.applyFinalityConfigUpdates(
      finalityThreshold, fastTransferFeeBps, maxAmountPerRequest, new TokenPool.FastTransferRateLimitConfigArgs[](0)
    );
    uint256 amount = maxAmountPerRequest + 1; // Exceeds max.
    s_token.transfer(address(s_tokenPool), amount);

    Pool.LockOrBurnInV1 memory lockOrBurnIn = Pool.LockOrBurnInV1({
      originalSender: s_sender,
      receiver: s_receiver,
      amount: amount,
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      localToken: address(s_token)
    });

    vm.expectRevert(abi.encodeWithSelector(TokenPool.AmountExceedsMaxPerRequest.selector, amount, maxAmountPerRequest));
    vm.startPrank(s_allowedOnRamp);
    s_tokenPool.lockOrBurn(lockOrBurnIn, finalityThreshold, "");
  }
}
