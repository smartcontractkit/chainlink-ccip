// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {USDCTokenPoolProxy} from "../../../../pools/USDC/USDCTokenPoolProxy.sol";
import {USDCTokenPoolProxySetup} from "./USDCTokenPoolProxySetup.t.sol";

contract USDCTokenPoolProxy_updatePoolAddresses is USDCTokenPoolProxySetup {
  address internal s_newCctpV1Pool = makeAddr("newCctpV1Pool");
  address internal s_newCctpV2Pool = makeAddr("newCctpV2Pool");
  address internal s_newLockReleasePool = makeAddr("newLockReleasePool");

  // Test successful pool address updates by owner
  function test_updatePoolAddresses() public {
    // Arrange: Define test constants
    USDCTokenPoolProxy.PoolAddresses memory newPools = USDCTokenPoolProxy.PoolAddresses({
      cctpV1Pool: s_newCctpV1Pool,
      cctpV2Pool: s_newCctpV2Pool,
      lockReleasePool: s_newLockReleasePool
    });

    // Act: Update pool addresses as owner
    changePrank(OWNER);
    s_usdcTokenPoolProxy.updatePoolAddresses(newPools);

    // Assert: Verify all pool addresses were updated correctly
    USDCTokenPoolProxy.PoolAddresses memory updatedPools = s_usdcTokenPoolProxy.getPools();
    assertEq(updatedPools.cctpV1Pool, s_newCctpV1Pool);
    assertEq(updatedPools.cctpV2Pool, s_newCctpV2Pool);
    assertEq(updatedPools.lockReleasePool, s_newLockReleasePool);
  }

  // Test that non-owner cannot update pool addresses
  function test_updatePoolAddresses_RevertWhen_NotOwner() public {
    // Arrange: Define test constants
    USDCTokenPoolProxy.PoolAddresses memory newPools = USDCTokenPoolProxy.PoolAddresses({
      cctpV1Pool: s_newCctpV1Pool,
      cctpV2Pool: s_newCctpV2Pool,
      lockReleasePool: s_newLockReleasePool
    });

    // Act & Assert: Non-owner should not be able to update addresses
    changePrank(makeAddr("notOwner"));
    vm.expectRevert();
    s_usdcTokenPoolProxy.updatePoolAddresses(newPools);
  }

  // Test that zero address for CCTP V1 pool is rejected
  function test_updatePoolAddresses_RevertWhen_CCTPV1PoolIsZero() public {
    // Arrange: Define test constants
    USDCTokenPoolProxy.PoolAddresses memory newPools = USDCTokenPoolProxy.PoolAddresses({
      cctpV1Pool: address(0), // Zero address
      cctpV2Pool: s_newCctpV2Pool,
      lockReleasePool: s_newLockReleasePool
    });

    // Act & Assert: Should revert with PoolAddressCannotBeZero error
    changePrank(OWNER);
    vm.expectRevert(USDCTokenPoolProxy.PoolAddressCannotBeZero.selector);
    s_usdcTokenPoolProxy.updatePoolAddresses(newPools);
  }

  // Test that zero address for CCTP V2 pool is rejected
  function test_updatePoolAddresses_RevertWhen_CCTPV2PoolIsZero() public {
    // Arrange: Define test constants
    USDCTokenPoolProxy.PoolAddresses memory newPools = USDCTokenPoolProxy.PoolAddresses({
      cctpV1Pool: s_newCctpV1Pool,
      cctpV2Pool: address(0), // Zero address
      lockReleasePool: s_newLockReleasePool
    });

    // Act & Assert: Should revert with PoolAddressCannotBeZero error
    changePrank(OWNER);
    vm.expectRevert(USDCTokenPoolProxy.PoolAddressCannotBeZero.selector);
    s_usdcTokenPoolProxy.updatePoolAddresses(newPools);
  }

  // Test that zero address for lock release pool is rejected
  function test_updatePoolAddresses_RevertWhen_LockReleasePoolIsZero() public {
    // Arrange: Define test constants
    USDCTokenPoolProxy.PoolAddresses memory newPools = USDCTokenPoolProxy.PoolAddresses({
      cctpV1Pool: s_newCctpV1Pool,
      cctpV2Pool: s_newCctpV2Pool,
      lockReleasePool: address(0) // Zero address
    });

    // Act & Assert: Should revert with PoolAddressCannotBeZero error
    changePrank(OWNER);
    vm.expectRevert(USDCTokenPoolProxy.PoolAddressCannotBeZero.selector);
    s_usdcTokenPoolProxy.updatePoolAddresses(newPools);
  }

  // Test that pool address update emits correct event
  function test_updatePoolAddresses_EmitsEvent() public {
    // Arrange: Define test constants
    USDCTokenPoolProxy.PoolAddresses memory newPools = USDCTokenPoolProxy.PoolAddresses({
      cctpV1Pool: s_newCctpV1Pool,
      cctpV2Pool: s_newCctpV2Pool,
      lockReleasePool: s_newLockReleasePool
    });

    // Act & Assert: Verify PoolAddressesUpdated event is emitted with correct data
    changePrank(OWNER);
    vm.expectEmit(true, true, true, true);
    emit USDCTokenPoolProxy.PoolAddressesUpdated(newPools);
    s_usdcTokenPoolProxy.updatePoolAddresses(newPools);
  }
}
