// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolV1} from "../../../../interfaces/IPool.sol";
import {USDCTokenPoolProxy} from "../../../../pools/USDC/USDCTokenPoolProxy.sol";
import {USDCTokenPoolProxySetup} from "./USDCTokenPoolProxySetup.t.sol";

import {IERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/IERC20.sol";

/// @notice Helper contract to expose internal _getPoolForMechanism function for testing.
contract USDCTokenPoolProxyHelper is USDCTokenPoolProxy {
  constructor(
    IERC20 token,
    PoolAddresses memory poolAddresses,
    address router,
    address cctpVerifier
  ) USDCTokenPoolProxy(token, poolAddresses, router, cctpVerifier) {}

  function getPoolForMechanism(
    uint64 remoteChainSelector
  ) external view returns (address pool, bool isPoolV2) {
    return _getPoolForMechanism(remoteChainSelector);
  }
}

contract USDCTokenPoolProxy__getPoolForMechanism is USDCTokenPoolProxySetup {
  USDCTokenPoolProxyHelper internal s_helper;

  function setUp() public override {
    super.setUp();

    // Deploy the helper with the same config as the proxy.
    s_helper = new USDCTokenPoolProxyHelper(
      s_USDCToken,
      USDCTokenPoolProxy.PoolAddresses({
        cctpV1Pool: s_cctpV1Pool,
        cctpV2Pool: s_cctpV2Pool,
        cctpV2PoolWithCCV: s_cctpThroughCCVTokenPool,
        siloedLockReleasePool: address(0)
      }),
      address(s_router),
      address(s_cctpVerifier)
    );

    // Configure the CCV mechanism for the remote chain.
    uint64[] memory remoteChainSelectors = new uint64[](1);
    remoteChainSelectors[0] = s_remoteCCTPChainSelector;

    USDCTokenPoolProxy.LockOrBurnMechanism[] memory mechanisms = new USDCTokenPoolProxy.LockOrBurnMechanism[](1);
    mechanisms[0] = USDCTokenPoolProxy.LockOrBurnMechanism.CCV;
    s_helper.updateLockOrBurnMechanisms(remoteChainSelectors, mechanisms);
  }

  function test_getPoolForMechanism_CCV() public view {
    (address pool, bool isPoolV2) = s_helper.getPoolForMechanism(s_remoteCCTPChainSelector);

    assertEq(pool, s_cctpThroughCCVTokenPool);
    assertTrue(isPoolV2);
  }

  function test_getPoolForMechanism_CCTPV1() public {
    uint64 cctpV1ChainSelector = 11111;

    uint64[] memory remoteChainSelectors = new uint64[](1);
    remoteChainSelectors[0] = cctpV1ChainSelector;

    USDCTokenPoolProxy.LockOrBurnMechanism[] memory mechanisms = new USDCTokenPoolProxy.LockOrBurnMechanism[](1);
    mechanisms[0] = USDCTokenPoolProxy.LockOrBurnMechanism.CCTP_V1;
    s_helper.updateLockOrBurnMechanisms(remoteChainSelectors, mechanisms);

    (address pool, bool isPoolV2) = s_helper.getPoolForMechanism(cctpV1ChainSelector);

    assertEq(pool, s_cctpV1Pool);
    assertFalse(isPoolV2);
  }

  function test_getPoolForMechanism_CCTPV2() public {
    uint64 cctpV2ChainSelector = 22222;

    uint64[] memory remoteChainSelectors = new uint64[](1);
    remoteChainSelectors[0] = cctpV2ChainSelector;

    USDCTokenPoolProxy.LockOrBurnMechanism[] memory mechanisms = new USDCTokenPoolProxy.LockOrBurnMechanism[](1);
    mechanisms[0] = USDCTokenPoolProxy.LockOrBurnMechanism.CCTP_V2;
    s_helper.updateLockOrBurnMechanisms(remoteChainSelectors, mechanisms);

    (address pool, bool isPoolV2) = s_helper.getPoolForMechanism(cctpV2ChainSelector);

    assertEq(pool, s_cctpV2Pool);
    assertFalse(isPoolV2);
  }

  function test_getPoolForMechanism_LockRelease() public {
    address siloedPool = makeAddr("siloedLockReleasePool");
    _enableERC165InterfaceChecks(siloedPool, type(IPoolV1).interfaceId);

    s_helper.updatePoolAddresses(
      USDCTokenPoolProxy.PoolAddresses({
        cctpV1Pool: s_cctpV1Pool,
        cctpV2Pool: s_cctpV2Pool,
        cctpV2PoolWithCCV: s_cctpThroughCCVTokenPool,
        siloedLockReleasePool: siloedPool
      })
    );

    uint64[] memory remoteChainSelectors = new uint64[](1);
    remoteChainSelectors[0] = s_remoteLockReleaseChainSelector;

    USDCTokenPoolProxy.LockOrBurnMechanism[] memory mechanisms = new USDCTokenPoolProxy.LockOrBurnMechanism[](1);
    mechanisms[0] = USDCTokenPoolProxy.LockOrBurnMechanism.LOCK_RELEASE;
    s_helper.updateLockOrBurnMechanisms(remoteChainSelectors, mechanisms);

    (address pool, bool isPoolV2) = s_helper.getPoolForMechanism(s_remoteLockReleaseChainSelector);

    assertEq(pool, siloedPool);
    assertTrue(isPoolV2);
  }

  function test_getPoolForMechanism_RevertWhen_InvalidLockOrBurnMechanism() public {
    uint64 unconfiguredChainSelector = 99999;

    vm.expectRevert(
      abi.encodeWithSelector(
        USDCTokenPoolProxy.InvalidLockOrBurnMechanism.selector, USDCTokenPoolProxy.LockOrBurnMechanism.INVALID_MECHANISM
      )
    );
    s_helper.getPoolForMechanism(unconfiguredChainSelector);
  }

  function test_getPoolForMechanism_RevertWhen_MustSetPoolForMechanism_CCV() public {
    // Remove the CCV pool
    s_helper.updatePoolAddresses(
      USDCTokenPoolProxy.PoolAddresses({
        cctpV1Pool: s_cctpV1Pool,
        cctpV2Pool: s_cctpV2Pool,
        cctpV2PoolWithCCV: address(0),
        siloedLockReleasePool: address(0)
      })
    );

    vm.expectRevert(
      abi.encodeWithSelector(
        USDCTokenPoolProxy.MustSetPoolForMechanism.selector,
        s_remoteCCTPChainSelector,
        USDCTokenPoolProxy.LockOrBurnMechanism.CCV
      )
    );
    s_helper.getPoolForMechanism(s_remoteCCTPChainSelector);
  }

  function test_getPoolForMechanism_RevertWhen_MustSetPoolForMechanism_CCTPV1() public {
    uint64 cctpV1ChainSelector = 11111;

    // Configure mechanism but set pool to address(0).
    s_helper.updatePoolAddresses(
      USDCTokenPoolProxy.PoolAddresses({
        cctpV1Pool: address(0),
        cctpV2Pool: s_cctpV2Pool,
        cctpV2PoolWithCCV: s_cctpThroughCCVTokenPool,
        siloedLockReleasePool: address(0)
      })
    );

    // First set the pool so we can configure the mechanism
    _enableERC165InterfaceChecks(s_cctpV1Pool, type(IPoolV1).interfaceId);
    s_helper.updatePoolAddresses(
      USDCTokenPoolProxy.PoolAddresses({
        cctpV1Pool: s_cctpV1Pool,
        cctpV2Pool: s_cctpV2Pool,
        cctpV2PoolWithCCV: s_cctpThroughCCVTokenPool,
        siloedLockReleasePool: address(0)
      })
    );

    uint64[] memory remoteChainSelectors = new uint64[](1);
    remoteChainSelectors[0] = cctpV1ChainSelector;

    USDCTokenPoolProxy.LockOrBurnMechanism[] memory mechanisms = new USDCTokenPoolProxy.LockOrBurnMechanism[](1);
    mechanisms[0] = USDCTokenPoolProxy.LockOrBurnMechanism.CCTP_V1;
    s_helper.updateLockOrBurnMechanisms(remoteChainSelectors, mechanisms);

    // Now remove the pool
    s_helper.updatePoolAddresses(
      USDCTokenPoolProxy.PoolAddresses({
        cctpV1Pool: address(0),
        cctpV2Pool: s_cctpV2Pool,
        cctpV2PoolWithCCV: s_cctpThroughCCVTokenPool,
        siloedLockReleasePool: address(0)
      })
    );

    vm.expectRevert(
      abi.encodeWithSelector(
        USDCTokenPoolProxy.MustSetPoolForMechanism.selector,
        cctpV1ChainSelector,
        USDCTokenPoolProxy.LockOrBurnMechanism.CCTP_V1
      )
    );
    s_helper.getPoolForMechanism(cctpV1ChainSelector);
  }
}
