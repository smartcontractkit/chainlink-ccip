// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

import {HyperLiquidCompatibleERC20} from "../../../tokenAdminRegistry/TokenPoolFactory/HyperLiquidCompatibleERC20.sol";
import {HyperLiquidCompatibleERC20Setup} from "./HyperLiquidCompatibleERC20Setup.t.sol";

contract HyperLiquidCompatibleERC20_setRemoteToken is HyperLiquidCompatibleERC20Setup {
  function test_setRemoteToken_Success() public {
    address testRemoteToken = makeAddr("testRemoteToken");
    uint8 testDecimals = 6;

    vm.expectEmit();
    emit HyperLiquidCompatibleERC20.RemoteTokenSet(testRemoteToken, testDecimals);
    s_hyperLiquidToken.setRemoteToken(testRemoteToken, testDecimals);

    // Note: Since these are internal variables, we can't directly test them
    // but we can verify the function doesn't revert and emits the correct event
  }

  function test_setRemoteToken_Success_ZeroDecimals() public {
    address testRemoteToken = makeAddr("testRemoteToken");
    uint8 testDecimals = 0;

    vm.expectEmit();
    emit HyperLiquidCompatibleERC20.RemoteTokenSet(testRemoteToken, testDecimals);
    s_hyperLiquidToken.setRemoteToken(testRemoteToken, testDecimals);
  }

  function test_setRemoteToken_Success_MaxDecimals() public {
    address testRemoteToken = makeAddr("testRemoteToken");
    uint8 testDecimals = 255;

    vm.expectEmit();
    emit HyperLiquidCompatibleERC20.RemoteTokenSet(testRemoteToken, testDecimals);
    s_hyperLiquidToken.setRemoteToken(testRemoteToken, testDecimals);
  }

  // Reverts

  function test_setRemoteToken_RevertWhen_ZeroAddress() public {
    uint8 testDecimals = 18;

    vm.expectRevert(abi.encodeWithSelector(HyperLiquidCompatibleERC20.ZeroAddressNotAllowed.selector));
    s_hyperLiquidToken.setRemoteToken(address(0), testDecimals);
  }
}
