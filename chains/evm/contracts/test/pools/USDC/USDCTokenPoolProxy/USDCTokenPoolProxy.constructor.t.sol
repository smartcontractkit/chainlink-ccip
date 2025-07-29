// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {USDCTokenPoolProxy} from "../../../../pools/USDC/USDCTokenPoolProxy.sol";

import {LockReleaseTokenPoolHelper} from "../../../helpers/LockReleaseTokenPoolHelper.sol";
import {USDCTokenPoolCCTPV2Helper} from "../../../helpers/USDCTokenPoolCCTPV2Helper.sol";
import {USDCTokenPoolHelper} from "../../../helpers/USDCTokenPoolHelper.sol";

import {USDCSetup} from "../USDCSetup.t.sol";

contract USDCTokenPoolProxy_constructor is USDCSetup {
  address internal s_cctpV1Pool;
  address internal s_cctpV2Pool;
  address internal s_lockReleasePool;

  function setUp() public virtual override {
    super.setUp();

    // Create pool addresses using makeAddr
    s_cctpV1Pool = makeAddr("cctpV1Pool");
    s_cctpV2Pool = makeAddr("cctpV2Pool");
    s_lockReleasePool = makeAddr("lockReleasePool");
  }

  function test_constructor() public {
    // Arrange: Define test constants
    address[] memory emptyAllowlist = new address[](0);
    address routerAddress = address(s_router);
    address rmnRemoteAddress = address(s_mockRMNRemote);
    address cctpV1PoolAddress = address(s_cctpV1Pool);
    address cctpV2PoolAddress = address(s_cctpV2Pool);
    address lockReleasePoolAddress = address(s_lockReleasePool);

    USDCTokenPoolProxy proxy = new USDCTokenPoolProxy(
      s_USDCToken,
      routerAddress,
      emptyAllowlist,
      rmnRemoteAddress,
      cctpV1PoolAddress,
      cctpV2PoolAddress,
      lockReleasePoolAddress
    );

    USDCTokenPoolProxy.PoolAddresses memory pools = proxy.getPools();
    assertEq(pools.cctpV1Pool, cctpV1PoolAddress);
    assertEq(pools.cctpV2Pool, cctpV2PoolAddress);
    assertEq(pools.lockReleasePool, lockReleasePoolAddress);
  }

  function test_constructor_RevertWhen_CCTPV1PoolIsZero() public {
    // Arrange: Define test constants
    address[] memory emptyAllowlist = new address[](0);
    address routerAddress = address(s_router);
    address rmnRemoteAddress = address(s_mockRMNRemote);
    address cctpV2PoolAddress = address(s_cctpV2Pool);
    address lockReleasePoolAddress = address(s_lockReleasePool);

    vm.expectRevert(USDCTokenPoolProxy.InvalidPoolAddresses.selector);
    new USDCTokenPoolProxy(
      s_USDCToken,
      routerAddress,
      emptyAllowlist,
      rmnRemoteAddress,
      address(0), // cctpV1Pool is zero
      cctpV2PoolAddress,
      lockReleasePoolAddress
    );
  }

  function test_constructor_RevertWhen_CCTPV2PoolIsZero() public {
    // Arrange: Define test constants
    address[] memory emptyAllowlist = new address[](0);
    address routerAddress = address(s_router);
    address rmnRemoteAddress = address(s_mockRMNRemote);
    address cctpV1PoolAddress = address(s_cctpV1Pool);
    address lockReleasePoolAddress = address(s_lockReleasePool);

    vm.expectRevert(USDCTokenPoolProxy.InvalidPoolAddresses.selector);
    new USDCTokenPoolProxy(
      s_USDCToken,
      routerAddress,
      emptyAllowlist,
      rmnRemoteAddress,
      cctpV1PoolAddress,
      address(0), // cctpV2Pool is zero
      lockReleasePoolAddress
    );
  }

  function test_constructor_RevertWhen_LockReleasePoolIsZero() public {
    // Arrange: Define test constants
    address[] memory emptyAllowlist = new address[](0);
    address routerAddress = address(s_router);
    address rmnRemoteAddress = address(s_mockRMNRemote);
    address cctpV1PoolAddress = address(s_cctpV1Pool);
    address cctpV2PoolAddress = address(s_cctpV2Pool);

    vm.expectRevert(USDCTokenPoolProxy.InvalidPoolAddresses.selector);
    new USDCTokenPoolProxy(
      s_USDCToken,
      routerAddress,
      emptyAllowlist,
      rmnRemoteAddress,
      cctpV1PoolAddress,
      cctpV2PoolAddress,
      address(0) // lockReleasePool is zero
    );
  }
}
