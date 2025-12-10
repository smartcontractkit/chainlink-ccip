// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Pool} from "../../../libraries/Pool.sol";
import {TokenPool} from "../../../pools/TokenPool.sol";
import {SiloedLockReleaseTokenPoolSetup} from "./SiloedLockReleaseTokenPoolSetup.t.sol";

contract SiloedLockReleaseTokenPool_lockOrBurn is SiloedLockReleaseTokenPoolSetup {
  uint256 public constant AMOUNT = 10e18;

  function test_lockOrBurn_SiloedFunds() public {
    assertTrue(s_siloedLockReleaseTokenPool.isSiloed(SILOED_CHAIN_SELECTOR));
    deal(address(s_token), address(s_siloedLockReleaseTokenPool), AMOUNT);

    vm.startPrank(s_allowedOnRamp);

    vm.expectEmit();
    emit TokenPool.OutboundRateLimitConsumed({
      remoteChainSelector: SILOED_CHAIN_SELECTOR,
      token: address(s_token),
      amount: AMOUNT
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

    assertEq(s_siloedLockReleaseTokenPool.getAvailableTokens(SILOED_CHAIN_SELECTOR), AMOUNT);
  }

  function test_lockOrBurn_UnsiloedFunds() public {
    vm.startPrank(s_allowedOnRamp);
    deal(address(s_token), address(s_siloedLockReleaseTokenPool), AMOUNT);

    assertFalse(s_siloedLockReleaseTokenPool.isSiloed(DEST_CHAIN_SELECTOR));

    vm.expectEmit();
    emit TokenPool.OutboundRateLimitConsumed({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      token: address(s_token),
      amount: AMOUNT
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

    assertEq(s_siloedLockReleaseTokenPool.getAvailableTokens(DEST_CHAIN_SELECTOR), AMOUNT);
  }

  function test_lockOrBurn_V2_UsesNetAmountForLiquidity() public {
    uint16 feeBps = 1_000;
    uint256 expectedLockedAmount = AMOUNT - (AMOUNT * feeBps) / 10_000;

    _setTokenTransferFee(DEST_CHAIN_SELECTOR, feeBps);
    deal(address(s_token), address(s_siloedLockReleaseTokenPool), AMOUNT);

    vm.startPrank(s_allowedOnRamp);

    vm.expectEmit();
    emit TokenPool.OutboundRateLimitConsumed({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      token: address(s_token),
      amount: AMOUNT
    });

    vm.expectEmit();
    emit TokenPool.LockedOrBurned({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      token: address(s_token),
      sender: address(s_allowedOnRamp),
      amount: expectedLockedAmount
    });

    (Pool.LockOrBurnOutV1 memory lockOrBurnOut, uint256 destTokenAmount) = s_siloedLockReleaseTokenPool.lockOrBurn(
      Pool.LockOrBurnInV1({
        originalSender: STRANGER,
        receiver: bytes(""),
        amount: AMOUNT,
        remoteChainSelector: DEST_CHAIN_SELECTOR,
        localToken: address(s_token)
      }),
      0,
      ""
    );

    assertEq(destTokenAmount, expectedLockedAmount);
    assertEq(lockOrBurnOut.destTokenAddress, abi.encode(address(2)));
    assertEq(s_siloedLockReleaseTokenPool.getAvailableTokens(DEST_CHAIN_SELECTOR), expectedLockedAmount);
    assertEq(s_token.balanceOf(address(s_lockBox)), expectedLockedAmount);
    assertEq(s_token.balanceOf(address(s_siloedLockReleaseTokenPool)), AMOUNT - expectedLockedAmount);
  }
}
