// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolV1} from "../../../../interfaces/IPool.sol";
import {IPoolV2} from "../../../../interfaces/IPoolV2.sol";
import {USDCTokenPoolProxy} from "../../../../pools/USDC/USDCTokenPoolProxy.sol";
import {USDCTokenPoolProxySetup} from "./USDCTokenPoolProxySetup.t.sol";

contract USDCTokenPoolProxy_getRemoteToken is USDCTokenPoolProxySetup {
  function test_getRemoteToken() public {
    bytes memory expectedRemoteToken = abi.encode(address(0x1234));
    vm.mockCall(
      address(s_cctpV2PoolWithCCV),
      abi.encodeWithSelector(IPoolV2.getRemoteToken.selector, uint64(1)),
      abi.encode(expectedRemoteToken)
    );

    bytes memory remoteToken = s_usdcTokenPoolProxy.getRemoteToken(1);

    assertEq(remoteToken, expectedRemoteToken);
  }

  function test_getRemoteToken_RevertWhen_NoCCVCompatiblePoolSet() public {
    USDCTokenPoolProxy.PoolAddresses memory pools = s_usdcTokenPoolProxy.getPools();
    pools.cctpV2PoolWithCCV = address(0);
    _enableERC165InterfaceChecks(pools.cctpV2PoolWithCCV, type(IPoolV2).interfaceId);
    _enableERC165InterfaceChecks(pools.cctpV2Pool, type(IPoolV1).interfaceId);
    _enableERC165InterfaceChecks(pools.cctpV1Pool, type(IPoolV1).interfaceId);
    _enableERC165InterfaceChecks(pools.legacyCctpV1Pool, type(IPoolV1).interfaceId);
    s_usdcTokenPoolProxy.updatePoolAddresses(pools);

    vm.expectRevert(abi.encodeWithSelector(USDCTokenPoolProxy.CCVCompatiblePoolNotSet.selector));
    s_usdcTokenPoolProxy.getRemoteToken(1);
  }
}
