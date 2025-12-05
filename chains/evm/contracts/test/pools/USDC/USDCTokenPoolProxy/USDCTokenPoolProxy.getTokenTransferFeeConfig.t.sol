// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolV1} from "../../../../interfaces/IPool.sol";
import {IPoolV2} from "../../../../interfaces/IPoolV2.sol";
import {USDCTokenPoolProxy} from "../../../../pools/USDC/USDCTokenPoolProxy.sol";
import {USDCTokenPoolProxySetup} from "./USDCTokenPoolProxySetup.t.sol";

contract USDCTokenPoolProxy_getTokenTransferFeeConfig is USDCTokenPoolProxySetup {
  function test_getTokenTransferFeeConfig() public {
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
      address(s_cctpV2PoolWithCCV),
      abi.encodeWithSelector(IPoolV2.getTokenTransferFeeConfig.selector, address(0), uint64(1), uint16(0), ""),
      abi.encode(expectedFeeConfig)
    );

    IPoolV2.TokenTransferFeeConfig memory feeConfig =
      s_usdcTokenPoolProxy.getTokenTransferFeeConfig(address(0), 1, 0, "");

    assertEq(feeConfig.destGasOverhead, expectedFeeConfig.destGasOverhead);
    assertEq(feeConfig.destBytesOverhead, expectedFeeConfig.destBytesOverhead);
    assertEq(feeConfig.defaultBlockConfirmationFeeUSDCents, expectedFeeConfig.defaultBlockConfirmationFeeUSDCents);
    assertEq(feeConfig.customBlockConfirmationFeeUSDCents, expectedFeeConfig.customBlockConfirmationFeeUSDCents);
    assertEq(feeConfig.defaultBlockConfirmationTransferFeeBps, expectedFeeConfig.defaultBlockConfirmationTransferFeeBps);
    assertEq(feeConfig.customBlockConfirmationTransferFeeBps, expectedFeeConfig.customBlockConfirmationTransferFeeBps);
    assertEq(feeConfig.isEnabled, expectedFeeConfig.isEnabled);
  }

  function test_getTokenTransferFeeConfig_RevertWhen_NoCCVCompatiblePoolSet() public {
    USDCTokenPoolProxy.PoolAddresses memory pools = s_usdcTokenPoolProxy.getPools();
    pools.cctpV2PoolWithCCV = address(0);
    _enableERC165InterfaceChecks(pools.cctpV2PoolWithCCV, type(IPoolV2).interfaceId);
    _enableERC165InterfaceChecks(pools.cctpV2Pool, type(IPoolV1).interfaceId);
    _enableERC165InterfaceChecks(pools.cctpV1Pool, type(IPoolV1).interfaceId);
    _enableERC165InterfaceChecks(pools.legacyCctpV1Pool, type(IPoolV1).interfaceId);
    s_usdcTokenPoolProxy.updatePoolAddresses(pools);

    vm.expectRevert(abi.encodeWithSelector(USDCTokenPoolProxy.CCVCompatiblePoolNotSet.selector));
    s_usdcTokenPoolProxy.getTokenTransferFeeConfig(address(0), 1, 0, "");
  }
}
