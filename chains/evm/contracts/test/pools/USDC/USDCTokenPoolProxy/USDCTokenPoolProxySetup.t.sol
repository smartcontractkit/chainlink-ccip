// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {CCTPMessageTransmitterProxy} from "../../../../pools/USDC/CCTPMessageTransmitterProxy.sol";
import {USDCTokenPool} from "../../../../pools/USDC/USDCTokenPool.sol";
import {USDCTokenPoolCCTPV2} from "../../../../pools/USDC/USDCTokenPoolCCTPV2.sol";
import {USDCTokenPoolProxy} from "../../../../pools/USDC/USDCTokenPoolProxy.sol";

import {LockReleaseTokenPoolHelper} from "../../../helpers/LockReleaseTokenPoolHelper.sol";
import {USDCTokenPoolCCTPV2Helper} from "../../../helpers/USDCTokenPoolCCTPV2Helper.sol";
import {USDCTokenPoolHelper} from "../../../helpers/USDCTokenPoolHelper.sol";

import {USDCSetup} from "../USDCSetup.t.sol";
import {BurnMintERC677} from "@chainlink/contracts/src/v0.8/shared/token/ERC677/BurnMintERC677.sol";

contract USDCTokenPoolProxySetup is USDCSetup {
  address internal s_cctpV1Pool = makeAddr("cctpV1Pool");
  address internal s_cctpV2Pool = makeAddr("cctpV2Pool");
  address internal s_lockReleasePool = makeAddr("lockReleasePool");
  USDCTokenPoolProxy internal s_usdcTokenPoolProxy;

  function setUp() public virtual override {
    super.setUp();

    // Mock lockOrBurn and releaseOrMint calls for the pools

    // Mock lockOrBurn for s_cctpV1Pool
    vm.mockCall(address(s_cctpV1Pool), abi.encodeWithSelector(USDCTokenPool.lockOrBurn.selector), abi.encode());
    // Mock releaseOrMint for s_cctpV1Pool
    vm.mockCall(address(s_cctpV1Pool), abi.encodeWithSelector(USDCTokenPool.releaseOrMint.selector), abi.encode());

    // Mock lockOrBurn for s_cctpV2Pool
    vm.mockCall(address(s_cctpV2Pool), abi.encodeWithSelector(USDCTokenPool.lockOrBurn.selector), abi.encode());
    // Mock releaseOrMint for s_cctpV2Pool
    vm.mockCall(address(s_cctpV2Pool), abi.encodeWithSelector(USDCTokenPool.releaseOrMint.selector), abi.encode());

    // Mock lockOrBurn for s_lockReleasePool
    vm.mockCall(address(s_lockReleasePool), abi.encodeWithSelector(USDCTokenPool.lockOrBurn.selector), abi.encode());
    // Mock releaseOrMint for s_lockReleasePool
    vm.mockCall(address(s_lockReleasePool), abi.encodeWithSelector(USDCTokenPool.releaseOrMint.selector), abi.encode());

    // Deploy the proxy
    s_usdcTokenPoolProxy = new USDCTokenPoolProxy(
      s_USDCToken,
      address(s_router),
      new address[](0),
      address(s_mockRMNRemote),
      address(s_cctpV1Pool),
      address(s_cctpV2Pool),
      address(s_lockReleasePool)
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
