// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Pool} from "../../../libraries/Pool.sol";
import {TokenPool} from "../../../poolsV2/TokenPool.sol";
import {TokenPoolV2Setup} from "./TokenPoolV2Setup.t.sol";

contract TokenPoolV2_lockOrBurn is TokenPoolV2Setup {
  uint256 constant BPS_DEVIDER = 10_000;

  function test_lockOrBurn() public {
    uint256 amount = 1000e18;
    s_token.transfer(address(s_tokenPool), amount);

    Pool.LockOrBurnInV1 memory lockOrBurnIn = Pool.LockOrBurnInV1({
      originalSender: OWNER,
      receiver: abi.encode(OWNER),
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
    TokenPool.FastFinalityConfig memory config = TokenPool.FastFinalityConfig({
      finalityThreshold: 8,
      fastTransferFeeBps: 500, // 5%
      maxAmountPerRequest: 1000e18
    });
    s_tokenPool.applyFinalityConfigUpdates(config);

    uint256 amount = 1000e18;
    s_token.transfer(address(s_tokenPool), amount);

    Pool.LockOrBurnInV1 memory lockOrBurnIn = Pool.LockOrBurnInV1({
      originalSender: OWNER,
      receiver: abi.encode(OWNER),
      amount: amount,
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      localToken: address(s_token)
    });

    vm.startPrank(s_allowedOnRamp);
    (Pool.LockOrBurnOutV1 memory lockOrBurnOut, uint256 destTokenAmount) =
      s_tokenPool.lockOrBurn(lockOrBurnIn, config.finalityThreshold, "");

    // Should deduct 5% fee
    uint256 expectedDestAmount = amount - (amount * config.fastTransferFeeBps / BPS_DEVIDER);
    assertEq(destTokenAmount, expectedDestAmount);
    assertEq(lockOrBurnOut.destTokenAddress, abi.encode(s_initialRemoteToken));
  }

  function test_lockOrBurn_WithFinalityConfig_User() public {
    // Set up finality config
    TokenPool.FastFinalityConfig memory config = TokenPool.FastFinalityConfig({
      finalityThreshold: 8,
      fastTransferFeeBps: 500, // 5%
      maxAmountPerRequest: 1000e18
    });
    s_tokenPool.applyFinalityConfigUpdates(config);

    uint256 amount = 1000e18;
    s_token.transfer(address(s_tokenPool), amount);

    Pool.LockOrBurnInV1 memory lockOrBurnIn = Pool.LockOrBurnInV1({
      originalSender: OWNER,
      receiver: abi.encode(OWNER),
      amount: amount,
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      localToken: address(s_token)
    });

    vm.startPrank(s_allowedOnRamp);
    (Pool.LockOrBurnOutV1 memory lockOrBurnOut, uint256 destTokenAmount) =
      s_tokenPool.lockOrBurn(lockOrBurnIn, config.finalityThreshold, "");

    // Should deduct 5% fee
    uint256 expectedDestAmount = amount - (amount * config.fastTransferFeeBps / BPS_DEVIDER);
    assertEq(destTokenAmount, expectedDestAmount);
    assertEq(lockOrBurnOut.destTokenAddress, abi.encode(s_initialRemoteToken));
  }

  // Reverts

  function test_lockOrBurn_RevertWhen_InvalidFinality() public {
    // Set up finality config
    TokenPool.FastFinalityConfig memory config =
      TokenPool.FastFinalityConfig({finalityThreshold: 5, fastTransferFeeBps: 500, maxAmountPerRequest: 1000e18});
    s_tokenPool.applyFinalityConfigUpdates(config);

    uint256 amount = 1000e18;
    s_token.transfer(address(s_tokenPool), amount);

    Pool.LockOrBurnInV1 memory lockOrBurnIn = Pool.LockOrBurnInV1({
      originalSender: OWNER,
      receiver: abi.encode(OWNER),
      amount: amount,
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      localToken: address(s_token)
    });

    // Finality below threshold should revert
    vm.expectRevert(abi.encodeWithSelector(TokenPool.InvalidFinality.selector, 1, 5));
    vm.startPrank(s_allowedOnRamp);
    s_tokenPool.lockOrBurn(lockOrBurnIn, 1, "");
  }

  function test_lockOrBurn_RevertWhen_AmountExceedsMaxPerRequest() public {
    // Set up finality config with low max amount
    TokenPool.FastFinalityConfig memory config = TokenPool.FastFinalityConfig({
      finalityThreshold: 8,
      fastTransferFeeBps: 500,
      maxAmountPerRequest: 500e18 // Lower than our test amount
    });
    s_tokenPool.applyFinalityConfigUpdates(config);

    uint256 amount = 1000e18; // Exceeds max
    s_token.transfer(address(s_tokenPool), amount);

    Pool.LockOrBurnInV1 memory lockOrBurnIn = Pool.LockOrBurnInV1({
      originalSender: OWNER,
      receiver: abi.encode(OWNER),
      amount: amount,
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      localToken: address(s_token)
    });

    vm.expectRevert(
      abi.encodeWithSelector(TokenPool.AmountExceedsMaxPerRequest.selector, amount, config.maxAmountPerRequest)
    );
    vm.startPrank(s_allowedOnRamp);
    s_tokenPool.lockOrBurn(lockOrBurnIn, 8, "");
  }
}
