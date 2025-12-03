// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolV1} from "../../../../interfaces/IPool.sol";
import {IPoolV2} from "../../../../interfaces/IPoolV2.sol";
import {USDCTokenPoolProxy} from "../../../../pools/USDC/USDCTokenPoolProxy.sol";
import {USDCTokenPoolProxySetup} from "./USDCTokenPoolProxySetup.t.sol";

contract USDCTokenPoolProxy_getFee is USDCTokenPoolProxySetup {
  function test_getFee() public {
    vm.mockCall(
      address(s_cctpV2PoolWithCCV),
      abi.encodeWithSelector(IPoolV2.getFee.selector, address(0), 0, 0, address(0), 0, ""),
      abi.encode(100, 1000, 1000, 100, true)
    );

    (uint256 usdFeeCents, uint32 destGasOverhead, uint32 destBytesOverhead, uint16 tokenFeeBps, bool isEnabled) =
      s_usdcTokenPoolProxy.getFee(address(0), 0, 0, address(0), 0, "");

    assertEq(usdFeeCents, 100);
    assertEq(destGasOverhead, 1000);
    assertEq(destBytesOverhead, 1000);
    assertEq(tokenFeeBps, 100);
    assertEq(isEnabled, true);
  }

  function test_getFee_RevertWhen_NoCCVCompatiblePoolSet() public {
    USDCTokenPoolProxy.PoolAddresses memory pools = s_usdcTokenPoolProxy.getPools();
    pools.cctpV2PoolWithCCV = address(0);
    _enableERC165InterfaceChecks(pools.cctpV2PoolWithCCV, type(IPoolV2).interfaceId);
    _enableERC165InterfaceChecks(pools.cctpV2Pool, type(IPoolV1).interfaceId);
    _enableERC165InterfaceChecks(pools.cctpV1Pool, type(IPoolV1).interfaceId);
    _enableERC165InterfaceChecks(pools.legacyCctpV1Pool, type(IPoolV1).interfaceId);
    s_usdcTokenPoolProxy.updatePoolAddresses(pools);

    vm.expectRevert(abi.encodeWithSelector(USDCTokenPoolProxy.CCVCompatiblePoolNotSet.selector));
    s_usdcTokenPoolProxy.getFee(address(0), 0, 0, address(0), 0, "");
  }
}
