// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

import {HyperLiquidCompatibleERC20} from "../../../tokenAdminRegistry/TokenPoolFactory/HyperLiquidCompatibleERC20.sol";
import {HyperLiquidCompatibleERC20Setup} from "./HyperLiquidCompatibleERC20Setup.t.sol";

contract HyperLiquidCompatibleERC20_setRemoteToken is HyperLiquidCompatibleERC20Setup {
  function testFuzz_setRemoteToken_Success(
    uint64 remoteTokenId
  ) public {
    vm.assume(remoteTokenId != 0);

    uint8 testDecimals = 6;

    address systemAddressInt = address((uint160(0x20) << 152) | uint160(remoteTokenId));

    vm.expectEmit();
    emit HyperLiquidCompatibleERC20.RemoteTokenSet(remoteTokenId, systemAddressInt, testDecimals);

    s_hyperLiquidToken.setRemoteToken(remoteTokenId, testDecimals);

    // Note: Since these are internal variables, we can't directly test them
    // but we can verify the function doesn't revert and emits the correct event
  }

  function test_setRemoteToken_Success_ZeroDecimals() public {
    uint8 testDecimals = 0;
    uint64 remoteTokenId = 1;
    address systemAddressInt = address((uint160(0x20) << 152) | uint160(remoteTokenId));

    vm.expectEmit();
    emit HyperLiquidCompatibleERC20.RemoteTokenSet(remoteTokenId, systemAddressInt, testDecimals);
    s_hyperLiquidToken.setRemoteToken(remoteTokenId, testDecimals);
  }

  function test_setRemoteToken_Success_MaxDecimals() public {
    uint8 testDecimals = 255;
    uint64 remoteTokenId = 1;
    address systemAddressInt = address((uint160(0x20) << 152) | uint160(remoteTokenId));

    vm.expectEmit();
    emit HyperLiquidCompatibleERC20.RemoteTokenSet(remoteTokenId, systemAddressInt, testDecimals);
    s_hyperLiquidToken.setRemoteToken(remoteTokenId, testDecimals);
  }

  // Reverts

  function test_setRemoteToken_RevertWhen_ZeroAddress() public {
    uint8 testDecimals = 18;

    vm.expectRevert(abi.encodeWithSelector(HyperLiquidCompatibleERC20.ZeroAddressNotAllowed.selector));
    s_hyperLiquidToken.setRemoteToken(0, testDecimals);
  }

  function test_setRemoteToken_RevertWhen_RemoteTokenAlreadySet() public {
    // Set the remote token first so that the second call reverts
    s_hyperLiquidToken.setRemoteToken(s_remoteTokenId, 18);

    vm.expectRevert(abi.encodeWithSelector(HyperLiquidCompatibleERC20.RemoteTokenAlreadySet.selector));
    s_hyperLiquidToken.setRemoteToken(s_remoteTokenId, 18);
  }
}
