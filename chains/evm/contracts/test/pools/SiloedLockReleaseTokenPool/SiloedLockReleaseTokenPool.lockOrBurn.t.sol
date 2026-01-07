// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolV2} from "../../../interfaces/IPoolV2.sol";

import {Pool} from "../../../libraries/Pool.sol";
import {TokenPool} from "../../../pools/TokenPool.sol";
import {SiloedLockReleaseTokenPoolSetup} from "./SiloedLockReleaseTokenPoolSetup.t.sol";

contract SiloedLockReleaseTokenPool_lockOrBurn is SiloedLockReleaseTokenPoolSetup {
  uint256 public constant AMOUNT = 10e18;

  function test_lockOrBurn_SiloedFunds() public {
    deal(address(s_token), address(s_siloedLockReleaseTokenPool), AMOUNT);

    vm.startPrank(s_allowedOnRamp);

    vm.expectEmit();
    emit TokenPool.OutboundRateLimitConsumed({
      remoteChainSelector: SILOED_CHAIN_SELECTOR, token: address(s_token), amount: AMOUNT
    });

    vm.expectEmit();
    emit TokenPool.LockedOrBurned({
      remoteChainSelector: SILOED_CHAIN_SELECTOR,
      token: address(s_token),
      sender: address(s_allowedOnRamp),
      amount: AMOUNT
    });

    s_siloedLockReleaseTokenPool.lockOrBurn(
      Pool.LockOrBurnInV1({
        originalSender: STRANGER,
        receiver: bytes(""),
        amount: AMOUNT,
        remoteChainSelector: SILOED_CHAIN_SELECTOR,
        localToken: address(s_token)
      })
    );

    assertEq(s_token.balanceOf(address(s_siloedLockReleaseTokenPool.getLockBox(SILOED_CHAIN_SELECTOR))), AMOUNT);
  }

  function test_lockOrBurn_UnsiloedFunds() public {
    vm.startPrank(s_allowedOnRamp);
    deal(address(s_token), address(s_siloedLockReleaseTokenPool), AMOUNT);

    vm.expectEmit();
    emit TokenPool.OutboundRateLimitConsumed({
      remoteChainSelector: DEST_CHAIN_SELECTOR, token: address(s_token), amount: AMOUNT
    });

    vm.expectEmit();
    emit TokenPool.LockedOrBurned({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      token: address(s_token),
      sender: address(s_allowedOnRamp),
      amount: AMOUNT
    });

    s_siloedLockReleaseTokenPool.lockOrBurn(
      Pool.LockOrBurnInV1({
        originalSender: STRANGER,
        receiver: bytes(""),
        amount: AMOUNT,
        remoteChainSelector: DEST_CHAIN_SELECTOR,
        localToken: address(s_token)
      })
    );

    assertEq(s_token.balanceOf(address(s_lockBox)), AMOUNT);
  }

  function test_lockOrBurnV2_SiloedFundsWithFee() public {
    uint256 amount = 1000e18;
    uint16 defaultFeeBps = 100;
    uint256 expectedFee = (amount * defaultFeeBps) / 10_000;
    uint256 expectedDestAmount = amount - expectedFee;

    IPoolV2.TokenTransferFeeConfig memory feeConfig = IPoolV2.TokenTransferFeeConfig({
      destGasOverhead: 50_000,
      destBytesOverhead: 32,
      defaultBlockConfirmationFeeUSDCents: 0,
      customBlockConfirmationFeeUSDCents: 0,
      defaultBlockConfirmationTransferFeeBps: defaultFeeBps,
      customBlockConfirmationTransferFeeBps: 0,
      isEnabled: true
    });

    TokenPool.TokenTransferFeeConfigArgs[] memory feeConfigArgs = new TokenPool.TokenTransferFeeConfigArgs[](1);
    feeConfigArgs[0] = TokenPool.TokenTransferFeeConfigArgs({
      destChainSelector: SILOED_CHAIN_SELECTOR, tokenTransferFeeConfig: feeConfig
    });

    s_siloedLockReleaseTokenPool.applyTokenTransferFeeConfigUpdates(feeConfigArgs, new uint64[](0));

    uint256 lockBoxBalanceBefore = s_token.balanceOf(address(s_siloLockBox));
    deal(address(s_token), address(s_siloedLockReleaseTokenPool), amount);

    vm.startPrank(s_allowedOnRamp);

    (, uint256 destTokenAmount) = s_siloedLockReleaseTokenPool.lockOrBurn(
      Pool.LockOrBurnInV1({
        originalSender: STRANGER,
        receiver: bytes(""),
        amount: amount,
        remoteChainSelector: SILOED_CHAIN_SELECTOR,
        localToken: address(s_token)
      }),
      0,
      ""
    );

    assertEq(destTokenAmount, expectedDestAmount);
    assertEq(s_token.balanceOf(address(s_siloLockBox)), lockBoxBalanceBefore + expectedDestAmount);
    assertEq(s_token.balanceOf(address(s_siloedLockReleaseTokenPool)), expectedFee);
  }

  // Reverts
}
