// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolV2} from "../../../interfaces/IPoolV2.sol";
import {Pool} from "../../../libraries/Pool.sol";
import {TokenPool} from "../../../pools/TokenPool.sol";
import {LockReleaseTokenPoolSetup} from "./LockReleaseTokenPoolSetup.t.sol";

contract LockReleaseTokenPool_lockOrBurn is LockReleaseTokenPoolSetup {
  function testFuzz_lockOrBurn_LockOrBurnNoAllowList(
    uint256 amount
  ) public {
    amount = bound(amount, 1, _getOutboundRateLimiterConfig().capacity);

    // Transfer tokens to the pool (simulating Router behavior).
    deal(address(s_token), address(s_lockReleaseTokenPool), amount);
    vm.startPrank(s_allowedOnRamp);

    vm.expectEmit();
    emit TokenPool.OutboundRateLimitConsumed({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      token: address(s_token),
      amount: amount
    });

    vm.expectEmit();
    emit TokenPool.LockedOrBurned({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      token: address(s_token),
      sender: address(s_allowedOnRamp),
      amount: amount
    });

    s_lockReleaseTokenPool.lockOrBurn(
      Pool.LockOrBurnInV1({
        originalSender: STRANGER,
        receiver: bytes(""),
        amount: amount,
        remoteChainSelector: DEST_CHAIN_SELECTOR,
        localToken: address(s_token)
      })
    );
  }

  function test_lockOrBurn_LockOrBurnWithAllowList() public {
    uint256 amount = 100;

    // Transfer tokens to the pool (simulating Router behavior).
    deal(address(s_token), address(s_lockReleaseTokenPoolWithAllowList), amount * 2);
    vm.startPrank(s_allowedOnRamp);

    vm.expectEmit();
    emit TokenPool.OutboundRateLimitConsumed({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      token: address(s_token),
      amount: amount
    });

    vm.expectEmit();
    emit TokenPool.LockedOrBurned({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      token: address(s_token),
      sender: address(s_allowedOnRamp),
      amount: amount
    });

    s_lockReleaseTokenPoolWithAllowList.lockOrBurn(
      Pool.LockOrBurnInV1({
        originalSender: s_allowedList[0],
        receiver: bytes(""),
        amount: amount,
        remoteChainSelector: DEST_CHAIN_SELECTOR,
        localToken: address(s_token)
      })
    );

    vm.expectEmit();
    emit TokenPool.LockedOrBurned({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      token: address(s_token),
      sender: address(s_allowedOnRamp),
      amount: amount
    });

    s_lockReleaseTokenPoolWithAllowList.lockOrBurn(
      Pool.LockOrBurnInV1({
        originalSender: s_allowedList[1],
        receiver: bytes(""),
        amount: amount,
        remoteChainSelector: DEST_CHAIN_SELECTOR,
        localToken: address(s_token)
      })
    );
  }

  function test_lockOrBurn_RevertWhen_SenderNotAllowed_LockOrBurnWithAllowList() public {
    vm.startPrank(s_allowedOnRamp);

    vm.expectRevert(abi.encodeWithSelector(TokenPool.SenderNotAllowed.selector, STRANGER));

    s_lockReleaseTokenPoolWithAllowList.lockOrBurn(
      Pool.LockOrBurnInV1({
        originalSender: STRANGER,
        receiver: bytes(""),
        amount: 100,
        remoteChainSelector: DEST_CHAIN_SELECTOR,
        localToken: address(s_token)
      })
    );
  }

  function test_lockOrBurn_RevertWhen_CursedByRMN_PoolBurnRevertNotHealthy() public {
    // Should not burn tokens if cursed.
    vm.mockCall(address(s_mockRMNRemote), abi.encodeWithSignature("isCursed(bytes16)"), abi.encode(true));
    uint256 before = s_token.balanceOf(address(s_lockBox));

    vm.startPrank(s_allowedOnRamp);
    vm.expectRevert(TokenPool.CursedByRMN.selector);

    s_lockReleaseTokenPoolWithAllowList.lockOrBurn(
      Pool.LockOrBurnInV1({
        originalSender: s_allowedList[0],
        receiver: bytes(""),
        amount: 1e5,
        remoteChainSelector: DEST_CHAIN_SELECTOR,
        localToken: address(s_token)
      })
    );

    assertEq(s_token.balanceOf(address(s_lockBox)), before);
  }

  function test_lockOrBurnV2_WithFee() public {
    uint256 amount = 1000e18;
    uint16 defaultFeeBps = 100; // 1%
    uint256 expectedFee = (amount * defaultFeeBps) / 10_000;
    uint256 expectedDestAmount = amount - expectedFee;

    // Configure fee.
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
    feeConfigArgs[0] =
      TokenPool.TokenTransferFeeConfigArgs({destChainSelector: DEST_CHAIN_SELECTOR, tokenTransferFeeConfig: feeConfig});

    s_lockReleaseTokenPool.applyTokenTransferFeeConfigUpdates(feeConfigArgs, new uint64[](0));

    // Setup tokens - transfer to pool (simulating Router behavior).
    uint256 lockBoxBalanceBefore = s_token.balanceOf(address(s_lockBox));
    deal(address(s_token), address(s_lockReleaseTokenPool), amount);
    vm.startPrank(s_allowedOnRamp);

    // Call V2 lockOrBurn with default finality (0).
    (, uint256 destTokenAmount) = s_lockReleaseTokenPool.lockOrBurn(
      Pool.LockOrBurnInV1({
        originalSender: STRANGER,
        receiver: bytes(""),
        amount: amount,
        remoteChainSelector: DEST_CHAIN_SELECTOR,
        localToken: address(s_token)
      }),
      0, // default finality.
      ""
    );

    // Verify destTokenAmount is correct (amount minus fee).
    assertEq(destTokenAmount, expectedDestAmount);

    // Verify only destTokenAmount went to lockbox (bridge liquidity).
    assertEq(s_token.balanceOf(address(s_lockBox)), lockBoxBalanceBefore + expectedDestAmount);

    // Verify fees remained on the pool contract.
    assertEq(s_token.balanceOf(address(s_lockReleaseTokenPool)), expectedFee);
  }
}
