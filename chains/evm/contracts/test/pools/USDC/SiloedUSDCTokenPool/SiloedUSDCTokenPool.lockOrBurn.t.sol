// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolV2} from "../../../../interfaces/IPoolV2.sol";

import {Pool} from "../../../../libraries/Pool.sol";
import {TokenPool} from "../../../../pools/TokenPool.sol";
import {SiloedUSDCTokenPoolSetup} from "./SiloedUSDCTokenPoolSetup.sol";

import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

contract SiloedUSDCTokenPool_lockOrBurn is SiloedUSDCTokenPoolSetup {
  address public s_sender = makeAddr("sender");
  bytes public s_receiver = abi.encode(makeAddr("receiver"));
  uint256 public s_amount = 1000e6;

  function setUp() public virtual override {
    super.setUp();

    // Deposit 1e12 USDC into the pool so that it can be transferred to the lock box
    deal(address(s_USDCToken), address(s_usdcTokenPool), 1e18);

    vm.startPrank(s_routerAllowedOnRamp);
  }

  function test_lockOrBurn_Success() public {
    Pool.LockOrBurnInV1 memory lockOrBurnIn = Pool.LockOrBurnInV1({
      receiver: s_receiver,
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      originalSender: s_sender,
      amount: s_amount,
      localToken: address(s_USDCToken)
    });

    vm.expectEmit();
    emit TokenPool.LockedOrBurned(DEST_CHAIN_SELECTOR, address(s_USDCToken), s_routerAllowedOnRamp, s_amount);

    Pool.LockOrBurnOutV1 memory result = s_usdcTokenPool.lockOrBurn(lockOrBurnIn);

    assertEq(s_usdcTokenPool.getAvailableTokens(DEST_CHAIN_SELECTOR), s_amount);

    // destPoolData is the local token decimals abi-encoded to 32 bytes
    assertEq(result.destPoolData.length, 32);
  }

  function test_lockOrBurn_UpdatesLockedTokensAccounting() public {
    Pool.LockOrBurnInV1 memory lockOrBurnIn = Pool.LockOrBurnInV1({
      receiver: s_receiver,
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      originalSender: s_sender,
      amount: s_amount,
      localToken: address(s_USDCToken)
    });

    Pool.LockOrBurnOutV1 memory result = s_usdcTokenPool.lockOrBurn(lockOrBurnIn);

    // Assert: Verify the locked tokens accounting is updated
    assertEq(s_usdcTokenPool.getAvailableTokens(DEST_CHAIN_SELECTOR), s_amount);

    // destPoolData is the local token decimals abi-encoded to 32 bytes
    assertEq(result.destPoolData.length, 32);
  }

  function test_lockOrBurn_MultipleLocks() public {
    address sender2 = makeAddr("sender2");
    uint256 amount2 = 300e6;
    bytes memory receiver2 = abi.encode(makeAddr("receiver2"));

    Pool.LockOrBurnInV1 memory lockOrBurnIn1 = Pool.LockOrBurnInV1({
      receiver: s_receiver,
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      originalSender: s_sender,
      amount: s_amount,
      localToken: address(s_USDCToken)
    });

    Pool.LockOrBurnOutV1 memory result1 = s_usdcTokenPool.lockOrBurn(lockOrBurnIn1);

    Pool.LockOrBurnInV1 memory lockOrBurnIn2 = Pool.LockOrBurnInV1({
      receiver: receiver2,
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      originalSender: sender2,
      amount: amount2,
      localToken: address(s_USDCToken)
    });

    Pool.LockOrBurnOutV1 memory result2 = s_usdcTokenPool.lockOrBurn(lockOrBurnIn2);

    assertEq(s_usdcTokenPool.getAvailableTokens(DEST_CHAIN_SELECTOR), s_amount + amount2);

    // destPoolData is the local token decimals abi-encoded to 32 bytes
    assertEq(result1.destPoolData.length, 32);
    assertEq(result2.destPoolData.length, 32);
  }

  function test_lockOrBurn_UpdatesTokensAccounting() public {
    Pool.LockOrBurnInV1 memory lockOrBurnIn = Pool.LockOrBurnInV1({
      receiver: s_receiver,
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      originalSender: s_sender,
      amount: s_amount,
      localToken: address(s_USDCToken)
    });

    Pool.LockOrBurnOutV1 memory result = s_usdcTokenPool.lockOrBurn(lockOrBurnIn);

    assertEq(s_usdcTokenPool.getAvailableTokens(DEST_CHAIN_SELECTOR), s_amount);

    assertEq(result.destPoolData.length, 32);
  }

  function test_lockOrBurnV2_WithFee() public {
    uint16 defaultFeeBps = 100;
    uint256 expectedFee = (s_amount * defaultFeeBps) / 10_000;
    uint256 expectedDestAmount = s_amount - expectedFee;

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
    changePrank(OWNER);
    s_usdcTokenPool.applyTokenTransferFeeConfigUpdates(feeConfigArgs, new uint64[](0));
    changePrank(s_routerAllowedOnRamp);

    uint256 lockBoxBalanceBefore = s_USDCToken.balanceOf(address(s_destLockBox));
    deal(address(s_USDCToken), address(s_usdcTokenPool), s_amount);

    Pool.LockOrBurnInV1 memory lockOrBurnIn = Pool.LockOrBurnInV1({
      receiver: s_receiver,
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      originalSender: s_sender,
      amount: s_amount,
      localToken: address(s_USDCToken)
    });

    (, uint256 destTokenAmount) = s_usdcTokenPool.lockOrBurn(lockOrBurnIn, 0, "");

    assertEq(destTokenAmount, expectedDestAmount);
    assertEq(s_USDCToken.balanceOf(address(s_destLockBox)), lockBoxBalanceBefore + expectedDestAmount);
    assertEq(s_USDCToken.balanceOf(address(s_usdcTokenPool)), expectedFee);
  }

  function test_lockOrBurn_RevertWhen_NotAllowedOnRamp() public {
    address unauthorizedCaller = makeAddr("unauthorized");

    changePrank(unauthorizedCaller);

    Pool.LockOrBurnInV1 memory lockOrBurnIn = Pool.LockOrBurnInV1({
      receiver: s_receiver,
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      originalSender: s_sender,
      amount: s_amount,
      localToken: address(s_USDCToken)
    });

    vm.expectRevert(abi.encodeWithSelector(AuthorizedCallers.UnauthorizedCaller.selector, unauthorizedCaller));
    s_usdcTokenPool.lockOrBurn(lockOrBurnIn);
  }

  function test_lockOrBurn_RevertWhen_ChainNotSupported() public {
    uint64 unsupportedChain = 999999999; // Chain that's not configured

    Pool.LockOrBurnInV1 memory lockOrBurnIn = Pool.LockOrBurnInV1({
      receiver: s_receiver,
      remoteChainSelector: unsupportedChain,
      originalSender: s_sender,
      amount: s_amount,
      localToken: address(s_USDCToken)
    });

    vm.expectRevert(abi.encodeWithSelector(TokenPool.ChainNotAllowed.selector, unsupportedChain));
    s_usdcTokenPool.lockOrBurn(lockOrBurnIn);
  }

  function test_lockOrBurn_RevertWhen_NotAllowedTokenPoolProxy() public {
    address unauthorizedProxy = makeAddr("unauthorizedProxy");

    changePrank(unauthorizedProxy);

    Pool.LockOrBurnInV1 memory lockOrBurnIn = Pool.LockOrBurnInV1({
      receiver: s_receiver,
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      originalSender: s_sender,
      amount: s_amount,
      localToken: address(s_USDCToken)
    });

    vm.expectRevert(abi.encodeWithSelector(AuthorizedCallers.UnauthorizedCaller.selector, unauthorizedProxy));
    s_usdcTokenPool.lockOrBurn(lockOrBurnIn);
  }
}
