// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {USDCTokenPoolProxy} from "../../../../pools/USDC/USDCTokenPoolProxy.sol";
import {USDCSetup} from "../USDCSetup.t.sol";

contract USDCTokenPoolProxy_constructor is USDCSetup {
  address internal s_legacyCctpV1Pool = makeAddr("legacyCctpV1Pool");
  address internal s_cctpV1Pool = makeAddr("cctpV1Pool");
  address internal s_cctpV2Pool = makeAddr("cctpV2Pool");
  address internal s_lockReleasePool = makeAddr("lockReleasePool");

  function test_constructor() public {
    // Arrange: Define test constants
    USDCTokenPoolProxy proxy = new USDCTokenPoolProxy(
      s_USDCToken,
      address(s_router),
      new address[](0),
      address(s_mockRMNRemote),
      USDCTokenPoolProxy.PoolAddresses({
        legacyCctpV1Pool: s_legacyCctpV1Pool,
        cctpV1Pool: s_cctpV1Pool,
        cctpV2Pool: s_cctpV2Pool,
        lockReleasePool: s_lockReleasePool
      })
    );

    USDCTokenPoolProxy.PoolAddresses memory pools = proxy.getPools();
    assertEq(pools.legacyCctpV1Pool, s_legacyCctpV1Pool);
    assertEq(pools.cctpV1Pool, s_cctpV1Pool);
    assertEq(pools.cctpV2Pool, s_cctpV2Pool);
    assertEq(pools.lockReleasePool, s_lockReleasePool);
  }

  // Reverts

  function test_constructor_RevertWhen_CCTPV2PoolIsZero() public {
    vm.expectRevert(USDCTokenPoolProxy.PoolAddressCannotBeZero.selector);
    new USDCTokenPoolProxy(
      s_USDCToken,
      address(s_router),
      new address[](0),
      address(s_mockRMNRemote),
      USDCTokenPoolProxy.PoolAddresses({
        legacyCctpV1Pool: s_legacyCctpV1Pool,
        cctpV1Pool: s_cctpV1Pool,
        cctpV2Pool: address(0),
        lockReleasePool: s_lockReleasePool
      })
    );
  }

  function test_constructor_RevertWhen_LockReleasePoolIsZero() public {
    vm.expectRevert(USDCTokenPoolProxy.PoolAddressCannotBeZero.selector);
    new USDCTokenPoolProxy(
      s_USDCToken,
      address(s_router),
      new address[](0),
      address(s_mockRMNRemote),
      USDCTokenPoolProxy.PoolAddresses({
        legacyCctpV1Pool: s_legacyCctpV1Pool,
        cctpV1Pool: s_cctpV1Pool,
        cctpV2Pool: s_cctpV2Pool,
        lockReleasePool: address(0)
      })
    );
  }
}
