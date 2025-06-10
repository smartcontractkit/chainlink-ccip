// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ITokenMessenger} from "../../../../pools/USDC/interfaces/ITokenMessenger.sol";

import {CCTPMessageTransmitterProxy} from "../../../../pools/USDC/CCTPMessageTransmitterProxy.sol";
import {HybridLockReleaseUSDCTokenPool} from "../../../../pools/USDC/HybridLockReleaseUSDCTokenPool.sol";
import {USDCTokenPool} from "../../../../pools/USDC/USDCTokenPool.sol";

import {MockE2EUSDCTransmitter} from "../../../mocks/MockE2EUSDCTransmitter.sol";
import {MockUSDCTokenMessenger} from "../../../mocks/MockUSDCTokenMessenger.sol";
import {USDCSetup} from "../USDCSetup.t.sol";

contract HybridLockReleaseUSDCTokenPool_constructor is USDCSetup {
  HybridLockReleaseUSDCTokenPool internal s_usdcTokenPool;
  HybridLockReleaseUSDCTokenPool internal s_usdcTokenPoolTransferLiquidity;
  CCTPMessageTransmitterProxy internal s_cctpMessageTransmitterProxyForTransferLiquidity;
  address[] internal s_allowedList;

  address internal s_previousPool = makeAddr("previousPool");

  // Reverts
  function test_RevertWhen_InvalidConfig() public {
    vm.expectRevert(USDCTokenPool.InvalidConfig.selector);

    s_usdcTokenPool = new HybridLockReleaseUSDCTokenPool(
      s_mockUSDC,
      ITokenMessenger(address(0)), // transmitter V2 that is checked for zero address
      s_cctpMessageTransmitterProxy,
      s_token,
      new address[](0),
      address(s_mockRMNRemote),
      address(s_router),
      address(1) // previousPool
    );

    vm.expectRevert(USDCTokenPool.InvalidConfig.selector);

    s_usdcTokenPool = new HybridLockReleaseUSDCTokenPool(
      s_mockUSDC,
      ITokenMessenger(address(0)), // transmitter V2 that is checked for zero address
      s_cctpMessageTransmitterProxy,
      s_token,
      new address[](0),
      address(s_mockRMNRemote),
      address(s_router),
      address(0) // previousPool
    );
  }

  function test_RevertWhen_InvalidMessageVersion() public {
    // Make a new fake transmitter with invalid message version
    s_mockUSDCTransmitterV2 = new MockE2EUSDCTransmitter(2, DEST_DOMAIN_IDENTIFIER, address(s_token));

    // Pass that to the token messenger which we will pass to the token pool
    MockUSDCTokenMessenger invalidTokenMessenger = new MockUSDCTokenMessenger(2, address(s_mockUSDCTransmitterV2));

    vm.expectRevert(abi.encodeWithSelector(USDCTokenPool.InvalidMessageVersion.selector, 2));

    s_usdcTokenPool = new HybridLockReleaseUSDCTokenPool(
      s_mockUSDC,
      invalidTokenMessenger, // transmitter V2 that is checked for zero address
      s_cctpMessageTransmitterProxy,
      s_token,
      new address[](0),
      address(s_mockRMNRemote),
      address(s_router),
      s_previousPool // previousPool
    );
  }

  function test_RevertWhen_InvalidTokenMessengerVersion() public {
    // Make a new fake token messenger with message version 2
    MockUSDCTokenMessenger invalidTokenMessenger = new MockUSDCTokenMessenger(2, address(s_mockUSDCTransmitterV2));

    vm.expectRevert(abi.encodeWithSelector(USDCTokenPool.InvalidTokenMessengerVersion.selector, 2));

    s_usdcTokenPool = new HybridLockReleaseUSDCTokenPool(
      s_mockUSDC,
      invalidTokenMessenger, // transmitter V2 that is checked for zero address
      s_cctpMessageTransmitterProxy,
      s_token,
      new address[](0),
      address(s_mockRMNRemote),
      address(s_router),
      s_previousPool // previousPool
    );
  }

  function test_RevertWhen_InvalidTransmitterInProxy() public {
    // Create a transmitter proxy where the V2 transmitter is the same as the V1 and so the hybrid contract
    // will point to the wrong address.
    CCTPMessageTransmitterProxy invalidProxy = new CCTPMessageTransmitterProxy(s_mockUSDC, s_mockUSDC);

    vm.expectRevert(abi.encodeWithSelector(USDCTokenPool.InvalidTransmitterInProxy.selector));
    // Since the V1 proxy is valid we can re-use it here since we only want to use a
    // proxy that points to any other messenger contract than the V2 one, which in this
    // case is the one for v1
    s_usdcTokenPool = new HybridLockReleaseUSDCTokenPool(
      s_mockUSDC,
      s_mockUSDCV2, // transmitter V2 that is checked for zero address
      invalidProxy,
      s_token,
      new address[](0),
      address(s_mockRMNRemote),
      address(s_router),
      s_previousPool // previousPool
    );
  }
}
