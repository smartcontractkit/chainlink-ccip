// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {CCTPMessageTransmitterProxy} from "../../../../pools/USDC/CCTPMessageTransmitterProxy.sol";
import {USDCTokenPool} from "../../../../pools/USDC/USDCTokenPool.sol";
import {USDCTokenPoolProxy} from "../../../../pools/USDC/USDCTokenPoolProxy.sol";
import {USDCSetup} from "../USDCSetup.t.sol";

contract USDCTokenPoolProxySetup is USDCSetup {
  address internal s_cctpV1Pool = makeAddr("cctpV1Pool");
  address internal s_cctpV2Pool = makeAddr("cctpV2Pool");
  address internal s_lockReleasePool = makeAddr("lockReleasePool");
  address internal s_mockTransmitterProxy = makeAddr("mockTransmitterProxy");

  USDCTokenPoolProxy internal s_usdcTokenPoolProxy;

  function setUp() public virtual override {
    super.setUp();

    // Mock the transmitter proxy's receiveMessage function to return true
    vm.mockCall(address(s_cctpV1Pool), abi.encodeWithSelector(bytes4(keccak256("i_messageTransmitterProxy()"))), abi.encode(s_mockTransmitterProxy));

    // Deploy the proxy
    s_usdcTokenPoolProxy = new USDCTokenPoolProxy(
      s_USDCToken,
      address(s_router),
      new address[](0),
      address(s_mockRMNRemote),
      USDCTokenPoolProxy.PoolAddresses({
        cctpV1Pool: s_cctpV1Pool,
        cctpV2Pool: s_cctpV2Pool,
        lockReleasePool: s_lockReleasePool
      })
    );

    // Deal some tokens to the proxy to test the transfer to the destination pool
    deal(address(s_USDCToken), address(s_usdcTokenPoolProxy), 1000e6);

    // Configure the pools
    _poolApplyChainUpdates(address(s_usdcTokenPoolProxy));

    // Configure allowed callers for the CCTP message transmitter proxy
    CCTPMessageTransmitterProxy.AllowedCallerConfigArgs[] memory allowedCallerParams =
      new CCTPMessageTransmitterProxy.AllowedCallerConfigArgs[](3);
    allowedCallerParams[0] =
      CCTPMessageTransmitterProxy.AllowedCallerConfigArgs({caller: address(s_cctpV1Pool), allowed: true});
    allowedCallerParams[1] =
      CCTPMessageTransmitterProxy.AllowedCallerConfigArgs({caller: address(s_cctpV2Pool), allowed: true});
    allowedCallerParams[2] =
      CCTPMessageTransmitterProxy.AllowedCallerConfigArgs({caller: address(s_lockReleasePool), allowed: true});
    s_cctpMessageTransmitterProxy.configureAllowedCallers(allowedCallerParams);
  }
}
