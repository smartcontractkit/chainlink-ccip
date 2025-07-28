// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Internal} from "../../../../libraries/Internal.sol";
import {Pool} from "../../../../libraries/Pool.sol";
import {TokenPool} from "../../../../pools/TokenPool.sol";
import {LOCK_RELEASE_FLAG} from "../../../../pools/USDC/USDCTokenPoolProxy.sol";
import {USDCBridgeMigrator} from "../../../../pools/USDC/USDCBridgeMigrator.sol";
import {USDCTokenPool} from "../../../../pools/USDC/USDCTokenPool.sol";
import {SiloedLockReleaseTokenPool} from "../../../../pools/SiloedLockReleaseTokenPool.sol";
import {SiloedUSDCTokenPoolSetup} from "./SiloedUSDCTokenPoolSetup.sol";

contract SiloedUSDCTokenPool_lockOrBurn is SiloedUSDCTokenPoolSetup {

  function setUp() public virtual override {
    super.setUp();
    
    // Deposit 1000 USDC into the pool so that it can be transferred to the lock box
    deal(address(s_USDCToken), address(s_usdcTokenPool), 1000e6);

    // Set up silo designation for the test chain
    vm.startPrank(OWNER);
    uint64[] memory removes = new uint64[](0);
    SiloedLockReleaseTokenPool.SiloConfigUpdate[] memory adds = new SiloedLockReleaseTokenPool.SiloConfigUpdate[](1);
    adds[0] = SiloedLockReleaseTokenPool.SiloConfigUpdate({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      rebalancer: OWNER
    });
    s_usdcTokenPool.updateSiloDesignations(removes, adds);
    vm.stopPrank();
  }

  function test_lockOrBurn_Success() public {
    // Arrange: Define test constants
    address sender = makeAddr("sender");
    uint256 amount = 1000e6; // 1000 USDC (6 decimals)
    address localToken = address(s_USDCToken);
    bytes memory receiver = abi.encode(makeAddr("receiver"));

    vm.startPrank(s_routerAllowedOnRamp);

    // Act: Call lockOrBurn
    Pool.LockOrBurnInV1 memory lockOrBurnIn = Pool.LockOrBurnInV1({
      receiver: receiver,
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      originalSender: sender,
      amount: amount,
      localToken: localToken
    });

    vm.expectEmit();
    emit TokenPool.LockedOrBurned(DEST_CHAIN_SELECTOR, localToken, s_routerAllowedOnRamp, amount);

    Pool.LockOrBurnOutV1 memory result = s_usdcTokenPool.lockOrBurn(lockOrBurnIn);

    // Assert: Verify the result
    assertEq(s_usdcTokenPool.getLockedTokensForChain(DEST_CHAIN_SELECTOR), amount);

    // destPoolData is the local token decimals abi-encoded to 32 bytes
    assertEq(result.destPoolData.length, 32);
    vm.stopPrank();
  }


  function test_lockOrBurn_UpdatesLockedTokensAccounting() public {
    // Arrange: Define test constants
    address sender = makeAddr("sender");
    uint256 amount = 1000e6;
    address localToken = address(s_USDCToken);
    bytes memory receiver = abi.encode(makeAddr("receiver"));

    vm.startPrank(s_routerAllowedOnRamp);

    // Act: Call lockOrBurn
    Pool.LockOrBurnInV1 memory lockOrBurnIn = Pool.LockOrBurnInV1({
      receiver: receiver,
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      originalSender: sender,
      amount: amount,
      localToken: localToken
    });

    Pool.LockOrBurnOutV1 memory result = s_usdcTokenPool.lockOrBurn(lockOrBurnIn);

    // Assert: Verify the locked tokens accounting is updated
    assertEq(s_usdcTokenPool.getLockedTokensForChain(DEST_CHAIN_SELECTOR), amount);

    // destPoolData is the local token decimals abi-encoded to 32 bytes
    assertEq(result.destPoolData.length, 32);
    vm.stopPrank();
  }


  function test_lockOrBurn_MultipleLocks() public {
    // Arrange: Define test constants
    address sender1 = makeAddr("sender1");
    address sender2 = makeAddr("sender2");
    uint256 amount1 = 500e6;
    uint256 amount2 = 300e6;
    address localToken = address(s_USDCToken);
    bytes memory receiver1 = abi.encode(makeAddr("receiver1"));
    bytes memory receiver2 = abi.encode(makeAddr("receiver2"));

    vm.startPrank(s_routerAllowedOnRamp);

    // Act: Call lockOrBurn twice
    Pool.LockOrBurnInV1 memory lockOrBurnIn1 = Pool.LockOrBurnInV1({
      receiver: receiver1,
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      originalSender: sender1,
      amount: amount1,
      localToken: localToken
    });

    Pool.LockOrBurnOutV1 memory result1 = s_usdcTokenPool.lockOrBurn(lockOrBurnIn1);

    Pool.LockOrBurnInV1 memory lockOrBurnIn2 = Pool.LockOrBurnInV1({
      receiver: receiver2,
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      originalSender: sender2,
      amount: amount2,
      localToken: localToken
    });

    Pool.LockOrBurnOutV1 memory result2 = s_usdcTokenPool.lockOrBurn(lockOrBurnIn2);

    // Assert: Verify the locked tokens accounting is updated correctly
    assertEq(s_usdcTokenPool.getLockedTokensForChain(DEST_CHAIN_SELECTOR), amount1 + amount2);

    // destPoolData is the local token decimals abi-encoded to 32 bytes
    assertEq(result1.destPoolData.length, 32);
    assertEq(result2.destPoolData.length, 32);
    vm.stopPrank();
  }

  function test_lockOrBurn_UpdatesSiloedTokensAccounting() public {
    // Arrange: Define test constants
    address sender = makeAddr("sender");
    uint256 amount = 1000e6;
    address localToken = address(s_USDCToken);
    bytes memory receiver = abi.encode(makeAddr("receiver"));

    vm.startPrank(s_routerAllowedOnRamp);

    // Act: Call lockOrBurn
    Pool.LockOrBurnInV1 memory lockOrBurnIn = Pool.LockOrBurnInV1({
      receiver: receiver,
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      originalSender: sender,
      amount: amount,
      localToken: localToken
    });

    Pool.LockOrBurnOutV1 memory result = s_usdcTokenPool.lockOrBurn(lockOrBurnIn);

    // Assert: Verify the siloed tokens accounting is updated correctly
    assertEq(s_usdcTokenPool.getLockedTokensForChain(DEST_CHAIN_SELECTOR), amount);
    assertTrue(s_usdcTokenPool.isSiloed(DEST_CHAIN_SELECTOR));

    // destPoolData is the local token decimals abi-encoded to 32 bytes
    assertEq(result.destPoolData.length, 32);
    vm.stopPrank();
  }
}
