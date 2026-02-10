// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolV1} from "../../../../interfaces/IPool.sol";
import {IPoolV2} from "../../../../interfaces/IPoolV2.sol";

import {USDCTokenPoolProxy} from "../../../../pools/USDC/USDCTokenPoolProxy.sol";
import {USDCTokenPoolProxySetup} from "./USDCTokenPoolProxySetup.t.sol";

contract USDCTokenPoolProxy_updatePoolAddresses is USDCTokenPoolProxySetup {
  address internal s_newCctpV1Pool = makeAddr("newCctpV1Pool");
  address internal s_newCctpV2Pool = makeAddr("newCctpV2Pool");
  address internal s_newCctpV2PoolWithCCV = makeAddr("newCctpV2PoolWithCCV");
  address internal s_newLockReleasePool = makeAddr("newLockReleasePool");

  function test_updatePoolAddresses() public {
    _enableERC165InterfaceChecks(s_newCctpV1Pool, type(IPoolV1).interfaceId);
    _enableERC165InterfaceChecks(s_newCctpV2Pool, type(IPoolV1).interfaceId);
    _enableERC165InterfaceChecks(s_newCctpV2PoolWithCCV, type(IPoolV2).interfaceId);

    USDCTokenPoolProxy.PoolAddresses memory expectedEvent = USDCTokenPoolProxy.PoolAddresses({
      cctpV1Pool: s_newCctpV1Pool,
      cctpV2Pool: s_newCctpV2Pool,
      cctpV2PoolWithCCV: s_newCctpV2PoolWithCCV,
      siloedLockReleasePool: address(0)
    });

    vm.expectEmit();
    emit USDCTokenPoolProxy.PoolAddressesUpdated(expectedEvent);

    s_usdcTokenPoolProxy.updatePoolAddresses(expectedEvent);

    USDCTokenPoolProxy.PoolAddresses memory updatedPools = s_usdcTokenPoolProxy.getPools();
    assertEq(updatedPools.cctpV1Pool, s_newCctpV1Pool);
    assertEq(updatedPools.cctpV2Pool, s_newCctpV2Pool);
  }

  // Reverts

  function test_updatePoolAddresses_RevertWhen_TokenPoolUnsupported_CCTPV1PoolDoesNotSupportIPoolV1() public {
    vm.expectRevert(abi.encodeWithSelector(USDCTokenPoolProxy.TokenPoolUnsupported.selector, s_newCctpV1Pool));
    s_usdcTokenPoolProxy.updatePoolAddresses(
      USDCTokenPoolProxy.PoolAddresses({
        cctpV1Pool: s_newCctpV1Pool,
        cctpV2Pool: address(0),
        cctpV2PoolWithCCV: address(0),
        siloedLockReleasePool: address(0)
      })
    );

    _enableERC165InterfaceChecks(s_newCctpV1Pool, type(IPoolV1).interfaceId);

    s_usdcTokenPoolProxy.updatePoolAddresses(
      USDCTokenPoolProxy.PoolAddresses({
        cctpV1Pool: s_newCctpV1Pool,
        cctpV2Pool: address(0),
        cctpV2PoolWithCCV: address(0),
        siloedLockReleasePool: address(0)
      })
    );

    assertEq(s_usdcTokenPoolProxy.getPools().cctpV1Pool, s_newCctpV1Pool);
  }

  function test_updatePoolAddresses_RevertWhen_TokenPoolUnsupported_CCTPV2PoolDoesNotSupportIPoolV1() public {
    vm.expectRevert(abi.encodeWithSelector(USDCTokenPoolProxy.TokenPoolUnsupported.selector, s_newCctpV2Pool));
    s_usdcTokenPoolProxy.updatePoolAddresses(
      USDCTokenPoolProxy.PoolAddresses({
        cctpV1Pool: address(0),
        cctpV2Pool: s_newCctpV2Pool,
        cctpV2PoolWithCCV: address(0),
        siloedLockReleasePool: address(0)
      })
    );

    _enableERC165InterfaceChecks(s_newCctpV2Pool, type(IPoolV1).interfaceId);

    s_usdcTokenPoolProxy.updatePoolAddresses(
      USDCTokenPoolProxy.PoolAddresses({
        cctpV1Pool: address(0),
        cctpV2Pool: s_newCctpV2Pool,
        cctpV2PoolWithCCV: address(0),
        siloedLockReleasePool: address(0)
      })
    );

    assertEq(s_usdcTokenPoolProxy.getPools().cctpV2Pool, s_newCctpV2Pool);
  }

  function test_updatePoolAddresses_RevertWhen_TokenPoolUnsupported_CCTPV2PoolWithCCVDoesNotSupportIPoolV2() public {
    vm.expectRevert(abi.encodeWithSelector(USDCTokenPoolProxy.TokenPoolUnsupported.selector, s_newCctpV2PoolWithCCV));
    s_usdcTokenPoolProxy.updatePoolAddresses(
      USDCTokenPoolProxy.PoolAddresses({
        cctpV1Pool: address(0),
        cctpV2Pool: address(0),
        cctpV2PoolWithCCV: s_newCctpV2PoolWithCCV,
        siloedLockReleasePool: address(0)
      })
    );

    _enableERC165InterfaceChecks(s_newCctpV2PoolWithCCV, type(IPoolV2).interfaceId);

    s_usdcTokenPoolProxy.updatePoolAddresses(
      USDCTokenPoolProxy.PoolAddresses({
        cctpV1Pool: address(0),
        cctpV2Pool: address(0),
        cctpV2PoolWithCCV: s_newCctpV2PoolWithCCV,
        siloedLockReleasePool: address(0)
      })
    );
  }

  function test_updatePoolAddresses_RevertWhen_TokenPoolUnsupported_SiloedLockReleasePoolDoesNotSupportIPoolV1OrV2()
    public
  {
    // Test when siloedLockReleasePool doesn't support either IPoolV1 or IPoolV2.
    vm.expectRevert(abi.encodeWithSelector(USDCTokenPoolProxy.TokenPoolUnsupported.selector, s_newLockReleasePool));
    s_usdcTokenPoolProxy.updatePoolAddresses(
      USDCTokenPoolProxy.PoolAddresses({
        cctpV1Pool: address(0),
        cctpV2Pool: address(0),
        cctpV2PoolWithCCV: address(0),
        siloedLockReleasePool: s_newLockReleasePool
      })
    );

    // Now enable IPoolV1 support and verify it works.
    _enableERC165InterfaceChecks(s_newLockReleasePool, type(IPoolV1).interfaceId);

    s_usdcTokenPoolProxy.updatePoolAddresses(
      USDCTokenPoolProxy.PoolAddresses({
        cctpV1Pool: address(0),
        cctpV2Pool: address(0),
        cctpV2PoolWithCCV: address(0),
        siloedLockReleasePool: s_newLockReleasePool
      })
    );

    assertEq(s_usdcTokenPoolProxy.getPools().siloedLockReleasePool, s_newLockReleasePool);
  }
}
