// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Pool} from "../../../../libraries/Pool.sol";
import {SiloedLockReleaseTokenPool} from "../../../../pools/SiloedLockReleaseTokenPool.sol";
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

    // Set up silo designation for the test chain
    vm.startPrank(OWNER);
    uint64[] memory removes = new uint64[](0);
    SiloedLockReleaseTokenPool.SiloConfigUpdate[] memory adds = new SiloedLockReleaseTokenPool.SiloConfigUpdate[](1);
    adds[0] = SiloedLockReleaseTokenPool.SiloConfigUpdate({remoteChainSelector: DEST_CHAIN_SELECTOR, rebalancer: OWNER});
    s_usdcTokenPool.updateSiloDesignations(removes, adds);
    vm.stopPrank();
  }

  function test_lockOrBurn_Success() public {
    vm.startPrank(s_routerAllowedOnRamp);

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

    // Assert: Verify the result
    assertEq(s_usdcTokenPool.getAvailableTokens(DEST_CHAIN_SELECTOR), s_amount);

    // destPoolData is the local token decimals abi-encoded to 32 bytes
    assertEq(result.destPoolData.length, 32);
    vm.stopPrank();
  }

  function test_lockOrBurn_UpdatesLockedTokensAccounting() public {
    // Arrange: Define test constants

    vm.startPrank(s_routerAllowedOnRamp);

    // Act: Call lockOrBurn
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
    vm.stopPrank();
  }

  function test_lockOrBurn_MultipleLocks() public {
    address sender2 = makeAddr("sender2");
    uint256 amount2 = 300e6;
    bytes memory receiver2 = abi.encode(makeAddr("receiver2"));

    vm.startPrank(s_routerAllowedOnRamp);

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

    // Assert: Verify the locked tokens accounting is updated correctly
    assertEq(s_usdcTokenPool.getAvailableTokens(DEST_CHAIN_SELECTOR), s_amount + amount2);

    // destPoolData is the local token decimals abi-encoded to 32 bytes
    assertEq(result1.destPoolData.length, 32);
    assertEq(result2.destPoolData.length, 32);
    vm.stopPrank();
  }

  function test_lockOrBurn_UpdatesSiloedTokensAccounting() public {
    vm.startPrank(s_routerAllowedOnRamp);

    Pool.LockOrBurnInV1 memory lockOrBurnIn = Pool.LockOrBurnInV1({
      receiver: s_receiver,
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      originalSender: s_sender,
      amount: s_amount,
      localToken: address(s_USDCToken)
    });

    Pool.LockOrBurnOutV1 memory result = s_usdcTokenPool.lockOrBurn(lockOrBurnIn);

    assertEq(s_usdcTokenPool.getAvailableTokens(DEST_CHAIN_SELECTOR), s_amount);
    assertTrue(s_usdcTokenPool.isSiloed(DEST_CHAIN_SELECTOR));

    assertEq(result.destPoolData.length, 32);
    vm.stopPrank();
  }

  function test_lockOrBurn_RevertWhen_NotAllowedOnRamp() public {
    address unauthorizedCaller = makeAddr("unauthorized");

    vm.startPrank(unauthorizedCaller);

    Pool.LockOrBurnInV1 memory lockOrBurnIn = Pool.LockOrBurnInV1({
      receiver: s_receiver,
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      originalSender: s_sender,
      amount: s_amount,
      localToken: address(s_USDCToken)
    });

    vm.expectRevert(abi.encodeWithSelector(AuthorizedCallers.UnauthorizedCaller.selector, unauthorizedCaller));
    s_usdcTokenPool.lockOrBurn(lockOrBurnIn);

    vm.stopPrank();
  }

  function test_lockOrBurn_RevertWhen_ChainNotSupported() public {
    uint64 unsupportedChain = 999999999; // Chain that's not configured

    vm.startPrank(s_routerAllowedOnRamp);

    Pool.LockOrBurnInV1 memory lockOrBurnIn = Pool.LockOrBurnInV1({
      receiver: s_receiver,
      remoteChainSelector: unsupportedChain,
      originalSender: s_sender,
      amount: s_amount,
      localToken: address(s_USDCToken)
    });

    vm.expectRevert(abi.encodeWithSelector(TokenPool.ChainNotAllowed.selector, unsupportedChain));
    s_usdcTokenPool.lockOrBurn(lockOrBurnIn);

    vm.stopPrank();
  }

  function test_lockOrBurn_RevertWhen_NotAllowedTokenPoolProxy() public {
    address unauthorizedProxy = makeAddr("unauthorizedProxy");

    vm.startPrank(unauthorizedProxy);

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
