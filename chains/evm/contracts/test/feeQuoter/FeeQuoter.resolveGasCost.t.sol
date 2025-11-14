// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {FeeQuoter} from "../../FeeQuoter.sol";
import {FeeQuoterSetup} from "./FeeQuoterSetup.t.sol";

contract FeeQuoter_resolveGasCost is FeeQuoterSetup {
  function test_resolveGasCost_ZeroCalldata() public view {
    uint32 nonCalldataGas = 100_000;
    uint32 calldataSize = 0;

    (uint32 totalGas, uint256 gasCostInUsdCents) =
      s_feeQuoter.resolveGasCost(DEST_CHAIN_SELECTOR, nonCalldataGas, calldataSize);

    // With zero calldata, total gas should equal non-calldata gas.
    assertEq(nonCalldataGas, totalGas);

    uint256 expectedCost = (uint256(totalGas) * USD_PER_GAS + (1e16 - 1)) / 1e16;
    assertEq(expectedCost, gasCostInUsdCents);
  }

  function test_resolveGasCost_ZeroNonCalldataGas() public view {
    uint32 nonCalldataGas = 0;
    uint32 calldataSize = 1000;

    (uint32 totalGas, uint256 gasCostInUsdCents) =
      s_feeQuoter.resolveGasCost(DEST_CHAIN_SELECTOR, nonCalldataGas, calldataSize);

    // With zero non-calldata gas, total should be calldata cost only.
    uint32 expectedTotalGas = calldataSize * DEST_GAS_PER_PAYLOAD_BYTE_BASE;
    assertEq(expectedTotalGas, totalGas);

    uint256 expectedCost = (uint256(totalGas) * USD_PER_GAS + (1e16 - 1)) / 1e16;
    assertEq(expectedCost, gasCostInUsdCents);
  }

  function test_resolveGasCost_WithBothGasTypes() public view {
    uint32 nonCalldataGas = 200_000;
    uint32 calldataSize = 500;

    (uint32 totalGas, uint256 gasCostInUsdCents) =
      s_feeQuoter.resolveGasCost(DEST_CHAIN_SELECTOR, nonCalldataGas, calldataSize);

    // Total gas should be sum of non-calldata and calldata gas.
    uint32 expectedTotalGas = nonCalldataGas + (calldataSize * DEST_GAS_PER_PAYLOAD_BYTE_BASE);
    assertEq(expectedTotalGas, totalGas);

    uint256 expectedCost = (uint256(totalGas) * USD_PER_GAS + (1e16 - 1)) / 1e16;
    assertEq(expectedCost, gasCostInUsdCents);
  }

  // Reverts

  function test_resolveGasCost_RevertWhen_DestinationChainNotEnabled() public {
    uint64 disabledChainSelector = 999999;

    vm.expectRevert(abi.encodeWithSelector(FeeQuoter.DestinationChainNotEnabled.selector, disabledChainSelector));
    s_feeQuoter.resolveGasCost(disabledChainSelector, 0, 0);
  }

  function test_resolveGasCost_RevertWhen_NoGasPriceAvailable() public {
    uint64 chainWithoutGasPrice = DEST_CHAIN_SELECTOR + 1;

    // Enable the destination chain but don't set a gas price.
    FeeQuoter.DestChainConfigArgs[] memory destChainConfigArgs = new FeeQuoter.DestChainConfigArgs[](1);
    destChainConfigArgs[0] = _generateFeeQuoterDestChainConfigArgs()[0];
    destChainConfigArgs[0].destChainSelector = chainWithoutGasPrice;

    s_feeQuoter.applyDestChainConfigUpdates(destChainConfigArgs);

    vm.expectRevert(abi.encodeWithSelector(FeeQuoter.NoGasPriceAvailable.selector, chainWithoutGasPrice));
    s_feeQuoter.resolveGasCost(chainWithoutGasPrice, 0, 0);
  }
}
