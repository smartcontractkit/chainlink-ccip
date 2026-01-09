// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolV1} from "../../../../interfaces/IPool.sol";
import {IPoolV2} from "../../../../interfaces/IPoolV2.sol";
import {USDCTokenPoolProxy} from "../../../../pools/USDC/USDCTokenPoolProxy.sol";
import {USDCTokenPoolProxySetup} from "./USDCTokenPoolProxySetup.t.sol";

contract USDCTokenPoolProxy_getTokenTransferFeeConfig is USDCTokenPoolProxySetup {
  function test_getTokenTransferFeeConfig_CCV() public {
    IPoolV2.TokenTransferFeeConfig memory expectedFeeConfig = IPoolV2.TokenTransferFeeConfig({
      destGasOverhead: 50_000,
      destBytesOverhead: 32,
      defaultBlockConfirmationFeeUSDCents: 100,
      customBlockConfirmationFeeUSDCents: 150,
      defaultBlockConfirmationTransferFeeBps: 123,
      customBlockConfirmationTransferFeeBps: 456,
      isEnabled: true
    });

    vm.mockCall(
      address(s_cctpThroughCCVTokenPool),
      abi.encodeWithSelector(
        IPoolV2.getTokenTransferFeeConfig.selector, address(0), s_remoteCCTPChainSelector, uint16(0), ""
      ),
      abi.encode(expectedFeeConfig)
    );

    IPoolV2.TokenTransferFeeConfig memory feeConfig =
      s_usdcTokenPoolProxy.getTokenTransferFeeConfig(address(0), s_remoteCCTPChainSelector, 0, "");

    _assertSameFeeConfig(feeConfig, expectedFeeConfig);
  }

  function test_getTokenTransferFeeConfig_LockRelease() public {
    // Set up siloed lock release pool
    address siloedPool = makeAddr("siloedLockReleasePool");
    _enableERC165InterfaceChecks(siloedPool, type(IPoolV1).interfaceId);

    s_usdcTokenPoolProxy.updatePoolAddresses(
      USDCTokenPoolProxy.PoolAddresses({
        cctpV1Pool: s_cctpV1Pool,
        cctpV2Pool: s_cctpV2Pool,
        cctpV2PoolWithCCV: s_cctpThroughCCVTokenPool,
        siloedLockReleasePool: siloedPool
      })
    );

    // Configure LOCK_RELEASE mechanism for a chain
    uint64[] memory remoteChainSelectors = new uint64[](1);
    remoteChainSelectors[0] = s_remoteLockReleaseChainSelector;

    USDCTokenPoolProxy.LockOrBurnMechanism[] memory mechanisms = new USDCTokenPoolProxy.LockOrBurnMechanism[](1);
    mechanisms[0] = USDCTokenPoolProxy.LockOrBurnMechanism.LOCK_RELEASE;
    s_usdcTokenPoolProxy.updateLockOrBurnMechanisms(remoteChainSelectors, mechanisms);

    IPoolV2.TokenTransferFeeConfig memory expectedFeeConfig = IPoolV2.TokenTransferFeeConfig({
      destGasOverhead: 60_000,
      destBytesOverhead: 64,
      defaultBlockConfirmationFeeUSDCents: 200,
      customBlockConfirmationFeeUSDCents: 250,
      defaultBlockConfirmationTransferFeeBps: 789,
      customBlockConfirmationTransferFeeBps: 999,
      isEnabled: true
    });

    vm.mockCall(
      siloedPool,
      abi.encodeWithSelector(
        IPoolV2.getTokenTransferFeeConfig.selector, address(0), s_remoteLockReleaseChainSelector, uint16(0), ""
      ),
      abi.encode(expectedFeeConfig)
    );

    IPoolV2.TokenTransferFeeConfig memory feeConfig =
      s_usdcTokenPoolProxy.getTokenTransferFeeConfig(address(0), s_remoteLockReleaseChainSelector, 0, "");

    _assertSameFeeConfig(feeConfig, expectedFeeConfig);
  }

  function test_getTokenTransferFeeConfig_CCTPV1_ReturnsEmptyConfig() public {
    // Configure CCTP_V1 mechanism
    uint64 cctpV1ChainSelector = 11111;
    uint64[] memory remoteChainSelectors = new uint64[](1);
    remoteChainSelectors[0] = cctpV1ChainSelector;

    USDCTokenPoolProxy.LockOrBurnMechanism[] memory mechanisms = new USDCTokenPoolProxy.LockOrBurnMechanism[](1);
    mechanisms[0] = USDCTokenPoolProxy.LockOrBurnMechanism.CCTP_V1;
    s_usdcTokenPoolProxy.updateLockOrBurnMechanisms(remoteChainSelectors, mechanisms);

    IPoolV2.TokenTransferFeeConfig memory feeConfig =
      s_usdcTokenPoolProxy.getTokenTransferFeeConfig(address(0), cctpV1ChainSelector, 0, "");

    // V1 pools return empty config
    IPoolV2.TokenTransferFeeConfig memory empty;
    _assertSameFeeConfig(feeConfig, empty);
  }

  function test_getTokenTransferFeeConfig_CCTPV2_ReturnsEmptyConfig() public {
    // Configure CCTP_V2 mechanism
    uint64 cctpV2ChainSelector = 22222;
    uint64[] memory remoteChainSelectors = new uint64[](1);
    remoteChainSelectors[0] = cctpV2ChainSelector;

    USDCTokenPoolProxy.LockOrBurnMechanism[] memory mechanisms = new USDCTokenPoolProxy.LockOrBurnMechanism[](1);
    mechanisms[0] = USDCTokenPoolProxy.LockOrBurnMechanism.CCTP_V2;
    s_usdcTokenPoolProxy.updateLockOrBurnMechanisms(remoteChainSelectors, mechanisms);

    IPoolV2.TokenTransferFeeConfig memory feeConfig =
      s_usdcTokenPoolProxy.getTokenTransferFeeConfig(address(0), cctpV2ChainSelector, 0, "");

    // V1 pools return empty config
    IPoolV2.TokenTransferFeeConfig memory empty;
    _assertSameFeeConfig(feeConfig, empty);
  }

  function test_getTokenTransferFeeConfig_RevertWhen_MustSetPoolForMechanism_CCV() public {
    // Remove the CCV pool
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
    s_usdcTokenPoolProxy.getTokenTransferFeeConfig(address(0), s_remoteCCTPChainSelector, 0, "");
  }

  function test_getTokenTransferFeeConfig_RevertWhen_MustSetPoolForMechanism_LockRelease() public {
    // First set the pool so we can configure the mechanism
    address siloedPool = makeAddr("siloedLockReleasePool");
    _enableERC165InterfaceChecks(siloedPool, type(IPoolV1).interfaceId);

    s_usdcTokenPoolProxy.updatePoolAddresses(
      USDCTokenPoolProxy.PoolAddresses({
        cctpV1Pool: s_cctpV1Pool,
        cctpV2Pool: s_cctpV2Pool,
        cctpV2PoolWithCCV: s_cctpThroughCCVTokenPool,
        siloedLockReleasePool: siloedPool
      })
    );

    // Configure LOCK_RELEASE mechanism
    uint64[] memory remoteChainSelectors = new uint64[](1);
    remoteChainSelectors[0] = s_remoteLockReleaseChainSelector;

    USDCTokenPoolProxy.LockOrBurnMechanism[] memory mechanisms = new USDCTokenPoolProxy.LockOrBurnMechanism[](1);
    mechanisms[0] = USDCTokenPoolProxy.LockOrBurnMechanism.LOCK_RELEASE;
    s_usdcTokenPoolProxy.updateLockOrBurnMechanisms(remoteChainSelectors, mechanisms);

    // Now remove the pool
    s_usdcTokenPoolProxy.updatePoolAddresses(
      USDCTokenPoolProxy.PoolAddresses({
        cctpV1Pool: s_cctpV1Pool,
        cctpV2Pool: s_cctpV2Pool,
        cctpV2PoolWithCCV: s_cctpThroughCCVTokenPool,
        siloedLockReleasePool: address(0)
      })
    );

    vm.expectRevert(
      abi.encodeWithSelector(
        USDCTokenPoolProxy.MustSetPoolForMechanism.selector,
        s_remoteLockReleaseChainSelector,
        USDCTokenPoolProxy.LockOrBurnMechanism.LOCK_RELEASE
      )
    );
    s_usdcTokenPoolProxy.getTokenTransferFeeConfig(address(0), s_remoteLockReleaseChainSelector, 0, "");
  }

  function test_getTokenTransferFeeConfig_RevertWhen_InvalidLockOrBurnMechanism_InvalidMechanism() public {
    // Use a chain selector with no mechanism configured (defaults to INVALID_MECHANISM)
    uint64 unconfiguredChainSelector = 99999;

    vm.expectRevert(
      abi.encodeWithSelector(
        USDCTokenPoolProxy.InvalidLockOrBurnMechanism.selector, USDCTokenPoolProxy.LockOrBurnMechanism.INVALID_MECHANISM
      )
    );
    s_usdcTokenPoolProxy.getTokenTransferFeeConfig(address(0), unconfiguredChainSelector, 0, "");
  }

  function _assertSameFeeConfig(
    IPoolV2.TokenTransferFeeConfig memory configA,
    IPoolV2.TokenTransferFeeConfig memory configB
  ) internal pure {
    assertEq(configA.destGasOverhead, configB.destGasOverhead);
    assertEq(configA.destBytesOverhead, configB.destBytesOverhead);
    assertEq(configA.defaultBlockConfirmationFeeUSDCents, configB.defaultBlockConfirmationFeeUSDCents);
    assertEq(configA.customBlockConfirmationFeeUSDCents, configB.customBlockConfirmationFeeUSDCents);
    assertEq(configA.defaultBlockConfirmationTransferFeeBps, configB.defaultBlockConfirmationTransferFeeBps);
    assertEq(configA.customBlockConfirmationTransferFeeBps, configB.customBlockConfirmationTransferFeeBps);
    assertEq(configA.isEnabled, configB.isEnabled);
  }
}
