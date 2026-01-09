// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ILockBox} from "../../../interfaces/ILockBox.sol";
import {IPoolV2} from "../../../interfaces/IPoolV2.sol";

import {Router} from "../../../Router.sol";
import {Pool} from "../../../libraries/Pool.sol";
import {SiloedLockReleaseTokenPool} from "../../../pools/SiloedLockReleaseTokenPool.sol";
import {TokenPool} from "../../../pools/TokenPool.sol";
import {SiloedLockReleaseTokenPoolSetup} from "./SiloedLockReleaseTokenPoolSetup.t.sol";

import {IERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/IERC20.sol";

/// @notice Mock lockbox that intentionally leaves dangling allowance for testing.
contract MockLockBoxWithDanglingAllowance is ILockBox {
  IERC20 internal immutable i_token;

  constructor(
    address token
  ) {
    i_token = IERC20(token);
  }

  function deposit(
    address,
    uint64,
    uint256 amount
  ) external {
    // Only transfer half the amount, leaving dangling allowance.
    i_token.transferFrom(msg.sender, address(this), amount / 2);
  }

  function withdraw(
    address,
    uint64,
    uint256,
    address
  ) external {}

  function isTokenSupported(
    address token
  ) external view returns (bool) {
    return token == address(i_token);
  }
}

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

  function test_lockOrBurn_ResetsAllowanceAfterDeposit() public {
    // Create a mock lockbox that leaves dangling allowance.
    MockLockBoxWithDanglingAllowance mockLockBox = new MockLockBoxWithDanglingAllowance(address(s_token));

    // Configure the mock lockbox for a new chain selector.
    uint64 testChainSelector = 99999;
    SiloedLockReleaseTokenPool.LockBoxConfig[] memory lockBoxConfigs = new SiloedLockReleaseTokenPool.LockBoxConfig[](1);
    lockBoxConfigs[0] =
      SiloedLockReleaseTokenPool.LockBoxConfig({remoteChainSelector: testChainSelector, lockBox: address(mockLockBox)});
    s_siloedLockReleaseTokenPool.configureLockBoxes(lockBoxConfigs);

    // Configure chain for the pool.
    bytes[] memory remotePoolAddresses = new bytes[](1);
    remotePoolAddresses[0] = abi.encode(address(1234));

    TokenPool.ChainUpdate[] memory chainUpdates = new TokenPool.ChainUpdate[](1);
    chainUpdates[0] = TokenPool.ChainUpdate({
      remoteChainSelector: testChainSelector,
      remotePoolAddresses: remotePoolAddresses,
      remoteTokenAddress: abi.encode(address(2)),
      outboundRateLimiterConfig: _getOutboundRateLimiterConfig(),
      inboundRateLimiterConfig: _getInboundRateLimiterConfig()
    });
    s_siloedLockReleaseTokenPool.applyChainUpdates(new uint64[](0), chainUpdates);

    // Add onRamp for the test chain.
    Router.OnRamp[] memory onRampUpdates = new Router.OnRamp[](1);
    onRampUpdates[0] = Router.OnRamp({destChainSelector: testChainSelector, onRamp: s_allowedOnRamp});
    s_sourceRouter.applyRampUpdates(onRampUpdates, new Router.OffRamp[](0), new Router.OffRamp[](0));

    // Fund the pool.
    deal(address(s_token), address(s_siloedLockReleaseTokenPool), AMOUNT);

    vm.stopPrank();
    vm.prank(s_allowedOnRamp);
    s_siloedLockReleaseTokenPool.lockOrBurn(
      Pool.LockOrBurnInV1({
        originalSender: STRANGER,
        receiver: bytes(""),
        amount: AMOUNT,
        remoteChainSelector: testChainSelector,
        localToken: address(s_token)
      })
    );

    // Verify the allowance is reset to 0 after the deposit.
    assertEq(s_token.allowance(address(s_siloedLockReleaseTokenPool), address(mockLockBox)), 0);
  }
}
