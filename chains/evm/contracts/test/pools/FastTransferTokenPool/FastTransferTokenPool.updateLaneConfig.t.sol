// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Internal} from "../../../libraries/Internal.sol";
import {FastTransferTokenPoolAbstract} from "../../../pools/FastTransferTokenPoolAbstract.sol";
import {FastTransferTokenPoolSetup} from "./FastTransferTokenPoolSetup.t.sol";

contract FastTransferTokenPool_updateDestChainConfig is FastTransferTokenPoolSetup {
  uint64 internal constant NEW_CHAIN_SELECTOR = 12345;
  bytes internal constant NEW_DESTINATION_POOL = abi.encode(address(0x5678));
  uint16 internal constant NEW_FAST_FEE_BPS = 200; // 2%
  uint256 internal constant NEW_FILL_AMOUNT_MAX = 2000 ether;

  function test_UpdateDestChainConfig() public {
    address[] memory addFillers = new address[](2);
    addFillers[0] = makeAddr("newFiller1");
    addFillers[1] = makeAddr("newFiller2");

    FastTransferTokenPoolAbstract.DestChainConfigUpdateArgs memory laneConfigArgs = FastTransferTokenPoolAbstract
      .DestChainConfigUpdateArgs({
      remoteChainSelector: NEW_CHAIN_SELECTOR,
      fastTransferBpsFee: NEW_FAST_FEE_BPS,
      fillerAllowlistEnabled: false,
      destinationPool: NEW_DESTINATION_POOL,
      maxFillAmountPerRequest: NEW_FILL_AMOUNT_MAX,
      settlementOverheadGas: 200_000,
      chainFamilySelector: Internal.CHAIN_FAMILY_SELECTOR_EVM,
      customExtraArgs: ""
    });

    vm.expectEmit();
    emit FastTransferTokenPoolAbstract.DestChainConfigUpdated(
      NEW_CHAIN_SELECTOR,
      NEW_FAST_FEE_BPS,
      NEW_FILL_AMOUNT_MAX,
      NEW_DESTINATION_POOL,
      Internal.CHAIN_FAMILY_SELECTOR_EVM,
      200_000,
      false
    );

    s_pool.updateDestChainConfig(_singleConfigToList(laneConfigArgs));
    s_pool.updateFillerAllowList(NEW_CHAIN_SELECTOR, addFillers, new address[](0));

    (FastTransferTokenPoolAbstract.DestChainConfig memory config,) = s_pool.getDestChainConfig(NEW_CHAIN_SELECTOR);
    assertEq(config.fastTransferBpsFee, NEW_FAST_FEE_BPS);
    assertFalse(config.fillerAllowlistEnabled);
    assertEq(config.destinationPool, NEW_DESTINATION_POOL);
    assertEq(config.maxFillAmountPerRequest, NEW_FILL_AMOUNT_MAX);
  }

  function test_UpdateDestChainConfig_ModifyExisting() public {
    // Modify existing lane config
    FastTransferTokenPoolAbstract.DestChainConfigUpdateArgs memory laneConfigArgs = FastTransferTokenPoolAbstract
      .DestChainConfigUpdateArgs({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      fastTransferBpsFee: NEW_FAST_FEE_BPS,
      fillerAllowlistEnabled: false, // disable whitelist
      destinationPool: NEW_DESTINATION_POOL,
      maxFillAmountPerRequest: NEW_FILL_AMOUNT_MAX,
      settlementOverheadGas: 200_000,
      chainFamilySelector: Internal.CHAIN_FAMILY_SELECTOR_EVM,
      customExtraArgs: ""
    });

    s_pool.updateDestChainConfig(_singleConfigToList(laneConfigArgs));

    (FastTransferTokenPoolAbstract.DestChainConfig memory config,) = s_pool.getDestChainConfig(DEST_CHAIN_SELECTOR);
    assertEq(config.fastTransferBpsFee, NEW_FAST_FEE_BPS);
    assertFalse(config.fillerAllowlistEnabled);
    assertEq(config.destinationPool, NEW_DESTINATION_POOL);
    assertEq(config.maxFillAmountPerRequest, NEW_FILL_AMOUNT_MAX);
  }

  function test_RevertWhen_InvalidFastFeeBps() public {
    FastTransferTokenPoolAbstract.DestChainConfigUpdateArgs memory laneConfigArgs = FastTransferTokenPoolAbstract
      .DestChainConfigUpdateArgs({
      remoteChainSelector: NEW_CHAIN_SELECTOR,
      fastTransferBpsFee: 10_001, // > 10_000 (100%)
      fillerAllowlistEnabled: true,
      destinationPool: NEW_DESTINATION_POOL,
      maxFillAmountPerRequest: NEW_FILL_AMOUNT_MAX,
      settlementOverheadGas: 200_000,
      chainFamilySelector: Internal.CHAIN_FAMILY_SELECTOR_EVM,
      customExtraArgs: ""
    });

    vm.expectRevert(FastTransferTokenPoolAbstract.InvalidDestChainConfig.selector);
    s_pool.updateDestChainConfig(_singleConfigToList(laneConfigArgs));
  }

  function test_RevertWhen_NotOwners() public {
    FastTransferTokenPoolAbstract.DestChainConfigUpdateArgs memory laneConfigArgs;

    vm.stopPrank();

    vm.expectRevert(); // TODO specify revert reason
    s_pool.updateDestChainConfig(_singleConfigToList(laneConfigArgs));
  }

  function test_UpdateDestChainConfig_ZeroFastFee() public {
    FastTransferTokenPoolAbstract.DestChainConfigUpdateArgs memory laneConfigArgs = FastTransferTokenPoolAbstract
      .DestChainConfigUpdateArgs({
      remoteChainSelector: NEW_CHAIN_SELECTOR,
      fastTransferBpsFee: 0, // No fast fee
      fillerAllowlistEnabled: true,
      destinationPool: NEW_DESTINATION_POOL,
      maxFillAmountPerRequest: NEW_FILL_AMOUNT_MAX,
      settlementOverheadGas: 200_000,
      chainFamilySelector: Internal.CHAIN_FAMILY_SELECTOR_EVM,
      customExtraArgs: ""
    });

    s_pool.updateDestChainConfig(_singleConfigToList(laneConfigArgs));

    (FastTransferTokenPoolAbstract.DestChainConfig memory config,) = s_pool.getDestChainConfig(NEW_CHAIN_SELECTOR);
    assertEq(config.fastTransferBpsFee, 0);
  }

  function test_UpdateDestChainConfig_MaxFastFee() public {
    FastTransferTokenPoolAbstract.DestChainConfigUpdateArgs memory laneConfigArgs = FastTransferTokenPoolAbstract
      .DestChainConfigUpdateArgs({
      remoteChainSelector: NEW_CHAIN_SELECTOR,
      fastTransferBpsFee: 10_000, // 100% fee
      fillerAllowlistEnabled: true,
      destinationPool: NEW_DESTINATION_POOL,
      maxFillAmountPerRequest: NEW_FILL_AMOUNT_MAX,
      settlementOverheadGas: 200_000,
      chainFamilySelector: Internal.CHAIN_FAMILY_SELECTOR_EVM,
      customExtraArgs: ""
    });

    s_pool.updateDestChainConfig(_singleConfigToList(laneConfigArgs));

    (FastTransferTokenPoolAbstract.DestChainConfig memory config,) = s_pool.getDestChainConfig(NEW_CHAIN_SELECTOR);
    assertEq(config.fastTransferBpsFee, 10_000);
  }
}
