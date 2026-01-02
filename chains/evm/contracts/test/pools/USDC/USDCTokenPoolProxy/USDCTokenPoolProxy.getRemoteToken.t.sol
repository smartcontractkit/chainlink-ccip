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
      address(s_cctpThroughCCVTokenPool),
      abi.encodeWithSelector(IPoolV2.getRemoteToken.selector, uint64(1)),
      abi.encode(expectedRemoteToken)
    );

    bytes memory remoteToken = s_usdcTokenPoolProxy.getRemoteToken(1);

    assertEq(remoteToken, expectedRemoteToken);
  }

  function test_getRemoteToken_RevertWhen_NoCCVCompatiblePoolSet() public {
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

    vm.expectRevert(abi.encodeWithSelector(USDCTokenPoolProxy.CCVCompatiblePoolNotSet.selector));
    s_usdcTokenPoolProxy.getRemoteToken(1);
  }
}
