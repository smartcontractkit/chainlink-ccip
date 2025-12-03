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

  // Test successful pool address updates by owner
  function test_updatePoolAddresses() public {
    // Arrange: Define test constants
    USDCTokenPoolProxy.PoolAddresses memory newPools = USDCTokenPoolProxy.PoolAddresses({
      legacyCctpV1Pool: s_legacyCctpV1Pool,
      cctpV1Pool: s_newCctpV1Pool,
      cctpV2Pool: s_newCctpV2Pool,
      cctpV2PoolWithCCV: s_newCctpV2PoolWithCCV
    });

    _enableERC165InterfaceChecks(s_newCctpV1Pool, type(IPoolV1).interfaceId);
    _enableERC165InterfaceChecks(s_newCctpV2Pool, type(IPoolV1).interfaceId);
    _enableERC165InterfaceChecks(s_legacyCctpV1Pool, type(IPoolV1).interfaceId);
    _enableERC165InterfaceChecks(s_newCctpV2PoolWithCCV, type(IPoolV2).interfaceId);

    // Act: Update pool addresses as owner
    changePrank(OWNER);
    s_usdcTokenPoolProxy.updatePoolAddresses(newPools);

    // Assert: Verify all pool addresses were updated correctly
    USDCTokenPoolProxy.PoolAddresses memory updatedPools = s_usdcTokenPoolProxy.getPools();
    assertEq(updatedPools.legacyCctpV1Pool, s_legacyCctpV1Pool);
    assertEq(updatedPools.cctpV1Pool, s_newCctpV1Pool);
    assertEq(updatedPools.cctpV2Pool, s_newCctpV2Pool);
  }

  // Reverts

  function test_updatePoolAddresses_RevertWhen_CCTPV1PoolDoesNotSupportIPoolV1() public {
    USDCTokenPoolProxy.PoolAddresses memory newPools = USDCTokenPoolProxy.PoolAddresses({
      legacyCctpV1Pool: address(0),
      cctpV1Pool: s_newCctpV1Pool,
      cctpV2Pool: address(0),
      cctpV2PoolWithCCV: address(0)
    });

    changePrank(OWNER);
    vm.expectRevert(abi.encodeWithSelector(USDCTokenPoolProxy.TokenPoolUnsupported.selector, s_newCctpV1Pool));
    s_usdcTokenPoolProxy.updatePoolAddresses(newPools);

    _enableERC165InterfaceChecks(s_newCctpV1Pool, type(IPoolV1).interfaceId);

    changePrank(OWNER);
    s_usdcTokenPoolProxy.updatePoolAddresses(newPools);

    assertEq(s_usdcTokenPoolProxy.getPools().cctpV1Pool, s_newCctpV1Pool);
  }

  function test_updatePoolAddresses_RevertWhen_CCTPV2PoolDoesNotSupportIPoolV1() public {
    USDCTokenPoolProxy.PoolAddresses memory newPools = USDCTokenPoolProxy.PoolAddresses({
      legacyCctpV1Pool: address(0),
      cctpV1Pool: address(0),
      cctpV2Pool: s_newCctpV2Pool,
      cctpV2PoolWithCCV: address(0)
    });

    changePrank(OWNER);
    vm.expectRevert(abi.encodeWithSelector(USDCTokenPoolProxy.TokenPoolUnsupported.selector, s_newCctpV2Pool));
    s_usdcTokenPoolProxy.updatePoolAddresses(newPools);

    _enableERC165InterfaceChecks(s_newCctpV2Pool, type(IPoolV1).interfaceId);

    changePrank(OWNER);
    s_usdcTokenPoolProxy.updatePoolAddresses(newPools);

    assertEq(s_usdcTokenPoolProxy.getPools().cctpV2Pool, s_newCctpV2Pool);
  }

  function test_updatePoolAddresses_RevertWhen_CCTPV2PoolWithCCVDoesNotSupportIPoolV2() public {
    USDCTokenPoolProxy.PoolAddresses memory newPools = USDCTokenPoolProxy.PoolAddresses({
      legacyCctpV1Pool: address(0),
      cctpV1Pool: address(0),
      cctpV2Pool: address(0),
      cctpV2PoolWithCCV: s_newCctpV2PoolWithCCV
    });

    changePrank(OWNER);
    vm.expectRevert(abi.encodeWithSelector(USDCTokenPoolProxy.TokenPoolUnsupported.selector, s_newCctpV2PoolWithCCV));
    s_usdcTokenPoolProxy.updatePoolAddresses(newPools);

    _enableERC165InterfaceChecks(s_newCctpV2PoolWithCCV, type(IPoolV2).interfaceId);

    changePrank(OWNER);
    s_usdcTokenPoolProxy.updatePoolAddresses(newPools);
  }

  function test_updatePoolAddresses_RevertWhen_LegacyPoolDoesNotSupportIPoolV1() public {
    USDCTokenPoolProxy.PoolAddresses memory newPools = USDCTokenPoolProxy.PoolAddresses({
      legacyCctpV1Pool: s_legacyCctpV1Pool,
      cctpV1Pool: s_newCctpV1Pool,
      cctpV2Pool: s_newCctpV2Pool,
      cctpV2PoolWithCCV: s_newCctpV2PoolWithCCV
    });

    // Enable the V1 and V2 pools to support the IPoolV1 interface
    _enableERC165InterfaceChecks(s_newCctpV1Pool, type(IPoolV1).interfaceId);
    _enableERC165InterfaceChecks(s_newCctpV2Pool, type(IPoolV1).interfaceId);
    _enableERC165InterfaceChecks(s_newCctpV2PoolWithCCV, type(IPoolV2).interfaceId);

    // Should revert because the legacy pool does not support the IPoolV1 interface even though the V1 and V2 pools do
    changePrank(OWNER);
    vm.expectRevert(abi.encodeWithSelector(USDCTokenPoolProxy.TokenPoolUnsupported.selector, s_legacyCctpV1Pool));
    s_usdcTokenPoolProxy.updatePoolAddresses(newPools);

    // Now it should succeed because the legacy pool is not being used and thus the check is not performed
    newPools.legacyCctpV1Pool = address(0);
    s_usdcTokenPoolProxy.updatePoolAddresses(newPools);

    // Now we re-enable the legacy pool to support the IPoolV1 interface
    newPools.legacyCctpV1Pool = s_legacyCctpV1Pool;

    // enable the legacy pool to support the IPoolV1 interface
    _enableERC165InterfaceChecks(s_legacyCctpV1Pool, type(IPoolV1).interfaceId);

    // Now it should Succeed
    s_usdcTokenPoolProxy.updatePoolAddresses(newPools);

    assertEq(s_usdcTokenPoolProxy.getPools().legacyCctpV1Pool, s_legacyCctpV1Pool);
  }
}
