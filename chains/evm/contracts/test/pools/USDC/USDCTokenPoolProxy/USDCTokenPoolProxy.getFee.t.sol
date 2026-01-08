// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolV1} from "../../../../interfaces/IPool.sol";
import {IPoolV2} from "../../../../interfaces/IPoolV2.sol";
import {USDCTokenPoolProxy} from "../../../../pools/USDC/USDCTokenPoolProxy.sol";
import {USDCTokenPoolProxySetup} from "./USDCTokenPoolProxySetup.t.sol";

contract USDCTokenPoolProxy_getFee is USDCTokenPoolProxySetup {
  uint256 internal constant FEE_USD_CENTS = 1e18;
  uint32 internal constant DEST_GAS_OVERHEAD = 200_000;
  uint32 internal constant DEST_BYTES_OVERHEAD = 32;
  uint16 internal constant TOKEN_FEE_BPS = 5;
  bool internal constant IS_ENABLED = true;

  function test_getFee() public {
    vm.mockCall(
      address(s_cctpThroughCCVTokenPool),
      abi.encodeWithSelector(IPoolV2.getFee.selector, address(0), s_remoteCCTPChainSelector, 0, address(0), 0, ""),
      abi.encode(FEE_USD_CENTS, DEST_GAS_OVERHEAD, DEST_BYTES_OVERHEAD, TOKEN_FEE_BPS, IS_ENABLED)
    );

    (uint256 usdFeeCents, uint32 destGasOverhead, uint32 destBytesOverhead, uint16 tokenFeeBps, bool isEnabled) =
      s_usdcTokenPoolProxy.getFee(address(0), s_remoteCCTPChainSelector, 0, address(0), 0, "");

    assertEq(usdFeeCents, FEE_USD_CENTS);
    assertEq(destGasOverhead, DEST_GAS_OVERHEAD);
    assertEq(destBytesOverhead, DEST_BYTES_OVERHEAD);
    assertEq(tokenFeeBps, TOKEN_FEE_BPS);
    assertEq(isEnabled, IS_ENABLED);
  }

  function test_getFee_RevertWhen_MustSetPoolForMechanism() public {
    _enableERC165InterfaceChecks(s_cctpV2Pool, type(IPoolV1).interfaceId);
    _enableERC165InterfaceChecks(s_cctpV1Pool, type(IPoolV1).interfaceId);
    s_usdcTokenPoolProxy.updatePoolAddresses(
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
    s_usdcTokenPoolProxy.getFee(address(0), s_remoteCCTPChainSelector, 0, address(0), 0, "");
  }

  function test_getFee_SiloedPool() public {
    // Configure the siloed lock/release pool
    _enableERC165InterfaceChecks(s_lockReleasePool, type(IPoolV1).interfaceId);
    s_usdcTokenPoolProxy.updatePoolAddresses(
      USDCTokenPoolProxy.PoolAddresses({
        cctpV1Pool: s_cctpV1Pool,
        cctpV2Pool: s_cctpV2Pool,
        cctpV2PoolWithCCV: s_cctpThroughCCVTokenPool,
        siloedLockReleasePool: s_lockReleasePool
      })
    );

    // Configure LOCK_RELEASE mechanism for the remote chain
    uint64[] memory remoteChainSelectors = new uint64[](1);
    remoteChainSelectors[0] = s_remoteLockReleaseChainSelector;

    USDCTokenPoolProxy.LockOrBurnMechanism[] memory mechanisms = new USDCTokenPoolProxy.LockOrBurnMechanism[](1);
    mechanisms[0] = USDCTokenPoolProxy.LockOrBurnMechanism.LOCK_RELEASE;
    s_usdcTokenPoolProxy.updateLockOrBurnMechanisms(remoteChainSelectors, mechanisms);

    // Mock the siloed pool's getFee to return specific values
    vm.mockCall(
      address(s_lockReleasePool),
      abi.encodeWithSelector(
        IPoolV2.getFee.selector, address(0), s_remoteLockReleaseChainSelector, 0, address(0), 0, ""
      ),
      abi.encode(FEE_USD_CENTS, DEST_GAS_OVERHEAD, DEST_BYTES_OVERHEAD, TOKEN_FEE_BPS, IS_ENABLED)
    );

    (uint256 usdFeeCents, uint32 destGasOverhead, uint32 destBytesOverhead, uint16 tokenFeeBps, bool isEnabled) =
      s_usdcTokenPoolProxy.getFee(address(0), s_remoteLockReleaseChainSelector, 0, address(0), 0, "");

    assertEq(usdFeeCents, FEE_USD_CENTS);
    assertEq(destGasOverhead, DEST_GAS_OVERHEAD);
    assertEq(destBytesOverhead, DEST_BYTES_OVERHEAD);
    assertEq(tokenFeeBps, TOKEN_FEE_BPS);
    assertEq(isEnabled, IS_ENABLED);
  }

  function test_getFee_RevertWhen_InvalidLockOrBurnMechanism_NoMechanismSet() public {
    // Use a chain selector that has no mechanism configured
    uint64 unconfiguredChainSelector = 99999;

    vm.expectRevert(
      abi.encodeWithSelector(
        USDCTokenPoolProxy.InvalidLockOrBurnMechanism.selector, USDCTokenPoolProxy.LockOrBurnMechanism.INVALID_MECHANISM
      )
    );
    s_usdcTokenPoolProxy.getFee(address(0), unconfiguredChainSelector, 0, address(0), 0, "");
  }

  function test_getFee_RevertWhen_InvalidLockOrBurnMechanism_OldMechanism() public {
    // Configure an old mechanism (CCTP_V1) that is not supported by getFee
    uint64[] memory remoteChainSelectors = new uint64[](1);
    remoteChainSelectors[0] = s_remoteLockReleaseChainSelector;

    USDCTokenPoolProxy.LockOrBurnMechanism[] memory mechanisms = new USDCTokenPoolProxy.LockOrBurnMechanism[](1);
    mechanisms[0] = USDCTokenPoolProxy.LockOrBurnMechanism.CCTP_V1;
    s_usdcTokenPoolProxy.updateLockOrBurnMechanisms(remoteChainSelectors, mechanisms);

    vm.expectRevert(
      abi.encodeWithSelector(
        USDCTokenPoolProxy.InvalidLockOrBurnMechanism.selector, USDCTokenPoolProxy.LockOrBurnMechanism.CCTP_V1
      )
    );
    s_usdcTokenPoolProxy.getFee(address(0), s_remoteLockReleaseChainSelector, 0, address(0), 0, "");
  }
}
