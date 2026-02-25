// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolV1} from "../../../../interfaces/IPool.sol";
import {USDCTokenPoolProxy} from "../../../../pools/USDC/USDCTokenPoolProxy.sol";
import {USDCTokenPoolProxySetup} from "./USDCTokenPoolProxySetup.t.sol";

contract USDCTokenPoolProxy_getRemotePools is USDCTokenPoolProxySetup {
  function test_getRemotePools() public {
    bytes[] memory expectedRemotePools = new bytes[](2);
    expectedRemotePools[0] = abi.encode(address(0x1234));
    expectedRemotePools[1] = abi.encode(address(0x5678));

    vm.mockCall(
      address(s_cctpThroughCCVTokenPool),
      abi.encodeWithSignature("getRemotePools(uint64)", s_remoteCCTPChainSelector),
      abi.encode(expectedRemotePools)
    );

    bytes[] memory remotePools = s_usdcTokenPoolProxy.getRemotePools(s_remoteCCTPChainSelector);

    assertEq(remotePools.length, expectedRemotePools.length);
    assertEq(remotePools[0], expectedRemotePools[0]);
    assertEq(remotePools[1], expectedRemotePools[1]);
  }

  function test_getRemotePools_RevertWhen_MustSetPoolForMechanism() public {
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
    s_usdcTokenPoolProxy.getRemotePools(s_remoteCCTPChainSelector);
  }
}
