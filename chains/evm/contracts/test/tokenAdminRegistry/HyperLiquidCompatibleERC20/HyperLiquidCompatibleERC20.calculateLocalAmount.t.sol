// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

import {HyperLiquidCompatibleERC20} from "../../../tokenAdminRegistry/TokenPoolFactory/HyperLiquidCompatibleERC20.sol";

import {MockHyperLiquidCompatibleERC20} from "../../mocks/MockHyperLiquidCompatibleERC20.sol";
import {HyperLiquidCompatibleERC20Setup} from "./HyperLiquidCompatibleERC20Setup.t.sol";

contract HyperLiquidCompatibleERC20_calculateLocalAmount is HyperLiquidCompatibleERC20Setup {
  function test_calculateLocalAmount() public view {
    uint8 localDecimals = s_hyperLiquidToken.decimals();
    uint256 remoteAmount = 123e18;

    // Zero decimals should return amount * 10^localDecimals
    assertEq(s_hyperLiquidToken.calculateLocalAmount(remoteAmount, 0), remoteAmount * 10 ** localDecimals);

    // Equal decimals should return the same amount
    assertEq(s_hyperLiquidToken.calculateLocalAmount(remoteAmount, localDecimals), remoteAmount);

    // Remote amount with more decimals should return less local amount
    uint256 expectedAmount = remoteAmount;
    for (uint8 remoteDecimals = localDecimals + 1; remoteDecimals < 36; ++remoteDecimals) {
      expectedAmount /= 10;
      assertEq(s_hyperLiquidToken.calculateLocalAmount(remoteAmount, remoteDecimals), expectedAmount);
    }

    // Remote amount with less decimals should return more local amount
    expectedAmount = remoteAmount;
    for (uint8 remoteDecimals = localDecimals - 1; remoteDecimals > 0; --remoteDecimals) {
      expectedAmount *= 10;
      assertEq(s_hyperLiquidToken.calculateLocalAmount(remoteAmount, remoteDecimals), expectedAmount);
    }
  }

  // Reverts

  function test_RevertWhen_calculateLocalAmountWhen_LowRemoteDecimalsOverflows() public {
    uint8 remoteDecimals = 0;
    uint8 localDecimals = 78;
    uint256 remoteAmount = 1;

    s_hyperLiquidToken = new MockHyperLiquidCompatibleERC20(
      "HyperLiquid Token", "HLT", localDecimals, type(uint256).max, type(uint256).max, OWNER
    );

    vm.expectRevert(
      abi.encodeWithSelector(
        HyperLiquidCompatibleERC20.OverflowDetected.selector, remoteDecimals, localDecimals, remoteAmount
      )
    );

    s_hyperLiquidToken.calculateLocalAmount(remoteAmount, remoteDecimals);
  }

  function test_RevertWhen_calculateLocalAmountWhen_HighLocalDecimalsOverflows() public {
    uint8 remoteDecimals = 18;
    uint8 localDecimals = 18 + 78;
    uint256 remoteAmount = 1;

    s_hyperLiquidToken = new MockHyperLiquidCompatibleERC20(
      "HyperLiquid Token", "HLT", localDecimals, type(uint256).max, type(uint256).max, OWNER
    );

    vm.expectRevert(
      abi.encodeWithSelector(
        HyperLiquidCompatibleERC20.OverflowDetected.selector, remoteDecimals, localDecimals, remoteAmount
      )
    );

    s_hyperLiquidToken.calculateLocalAmount(remoteAmount, remoteDecimals);
  }

  function test_RevertWhen_calculateLocalAmountWhen_HighRemoteDecimalsOverflows() public {
    uint8 remoteDecimals = 18 + 78;
    uint8 localDecimals = 18;
    uint256 remoteAmount = 1;

    s_hyperLiquidToken = new MockHyperLiquidCompatibleERC20(
      "HyperLiquid Token", "HLT", localDecimals, type(uint256).max, type(uint256).max, OWNER
    );

    vm.expectRevert(
      abi.encodeWithSelector(
        HyperLiquidCompatibleERC20.OverflowDetected.selector, remoteDecimals, localDecimals, remoteAmount
      )
    );

    s_hyperLiquidToken.calculateLocalAmount(remoteAmount, remoteDecimals);
  }

  function test_RevertWhen_calculateLocalAmountWhen_HighAmountOverflows() public {
    uint8 remoteDecimals = 18;
    uint8 localDecimals = 18 + 28;
    uint256 remoteAmount = 1e50;

    s_hyperLiquidToken = new MockHyperLiquidCompatibleERC20(
      "HyperLiquid Token", "HLT", localDecimals, type(uint256).max, type(uint256).max, OWNER
    );

    vm.expectRevert(
      abi.encodeWithSelector(
        HyperLiquidCompatibleERC20.OverflowDetected.selector, remoteDecimals, localDecimals, remoteAmount
      )
    );

    s_hyperLiquidToken.calculateLocalAmount(remoteAmount, remoteDecimals);
  }
}
