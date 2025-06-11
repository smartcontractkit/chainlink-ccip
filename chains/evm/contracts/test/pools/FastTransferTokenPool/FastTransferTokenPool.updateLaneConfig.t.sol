// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Internal} from "../../../libraries/Internal.sol";
import {FastTransferTokenPoolAbstract} from "../../../pools/FastTransferTokenPoolAbstract.sol";
import {FastTransferTokenPoolSetup} from "./FastTransferTokenPoolSetup.t.sol";

contract FastTransferTokenPool_updateDestChainConfig is FastTransferTokenPoolSetup {
  uint64 internal constant NEW_CHAIN_SELECTOR = 12345;
  bytes internal constant NEW_DESTINATION_POOL = abi.encode(address(0x5678));
  uint16 internal constant NEW_FAST_FEE_FILLER_BPS = 200; // 2%
  uint256 internal constant NEW_FILL_AMOUNT_MAX = 2000 ether;
  uint32 internal constant NEW_SETTLEMENT_GAS_OVERHEAD = SETTLEMENT_GAS_OVERHEAD + 100_000; // Increase by 100k

  function test_updateDestChainConfig() public {
    FastTransferTokenPoolAbstract.DestChainConfigUpdateArgs memory laneConfigArgs = FastTransferTokenPoolAbstract
      .DestChainConfigUpdateArgs({
      remoteChainSelector: NEW_CHAIN_SELECTOR,
      fastTransferFillerFeeBps: NEW_FAST_FEE_FILLER_BPS,
      fastTransferPoolFeeBps: 0, // No pool fee for this test
      fillerAllowlistEnabled: false,
      destinationPool: NEW_DESTINATION_POOL,
      maxFillAmountPerRequest: NEW_FILL_AMOUNT_MAX,
      settlementOverheadGas: NEW_SETTLEMENT_GAS_OVERHEAD,
      chainFamilySelector: Internal.CHAIN_FAMILY_SELECTOR_EVM,
      customExtraArgs: ""
    });

    vm.expectEmit();
    emit FastTransferTokenPoolAbstract.DestChainConfigUpdated(
      NEW_CHAIN_SELECTOR,
      NEW_FAST_FEE_FILLER_BPS,
      0, // pool fee BPS
      NEW_FILL_AMOUNT_MAX,
      NEW_DESTINATION_POOL,
      Internal.CHAIN_FAMILY_SELECTOR_EVM,
      NEW_SETTLEMENT_GAS_OVERHEAD,
      false
    );

    s_pool.updateDestChainConfig(_singleConfigToList(laneConfigArgs));

    (FastTransferTokenPoolAbstract.DestChainConfig memory config,) = s_pool.getDestChainConfig(NEW_CHAIN_SELECTOR);
    assertEq(config.fastTransferFillerFeeBps, NEW_FAST_FEE_FILLER_BPS);
    assertFalse(config.fillerAllowlistEnabled);
    assertEq(config.destinationPool, NEW_DESTINATION_POOL);
    assertEq(config.maxFillAmountPerRequest, NEW_FILL_AMOUNT_MAX);
  }

  function test_updateDestChainConfig_ModifyExisting() public {
    FastTransferTokenPoolAbstract.DestChainConfigUpdateArgs memory laneConfigArgs = FastTransferTokenPoolAbstract
      .DestChainConfigUpdateArgs({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      fastTransferFillerFeeBps: NEW_FAST_FEE_FILLER_BPS,
      fastTransferPoolFeeBps: 0, // No pool fee for this test
      fillerAllowlistEnabled: false, // disable allowlist
      destinationPool: NEW_DESTINATION_POOL,
      maxFillAmountPerRequest: NEW_FILL_AMOUNT_MAX,
      settlementOverheadGas: NEW_SETTLEMENT_GAS_OVERHEAD,
      chainFamilySelector: Internal.CHAIN_FAMILY_SELECTOR_EVM,
      customExtraArgs: ""
    });

    s_pool.updateDestChainConfig(_singleConfigToList(laneConfigArgs));

    (FastTransferTokenPoolAbstract.DestChainConfig memory config,) = s_pool.getDestChainConfig(DEST_CHAIN_SELECTOR);
    assertEq(config.fastTransferFillerFeeBps, NEW_FAST_FEE_FILLER_BPS);
    assertFalse(config.fillerAllowlistEnabled);
    assertEq(config.destinationPool, NEW_DESTINATION_POOL);
    assertEq(config.maxFillAmountPerRequest, NEW_FILL_AMOUNT_MAX);
  }

  function test_RevertWhen_InvalidFastFeeBps() public {
    FastTransferTokenPoolAbstract.DestChainConfigUpdateArgs memory laneConfigArgs = FastTransferTokenPoolAbstract
      .DestChainConfigUpdateArgs({
      remoteChainSelector: NEW_CHAIN_SELECTOR,
      fastTransferFillerFeeBps: 10_001, // > 10_000 (100%)
      fastTransferPoolFeeBps: 0, // No pool fee for this test
      fillerAllowlistEnabled: true,
      destinationPool: NEW_DESTINATION_POOL,
      maxFillAmountPerRequest: NEW_FILL_AMOUNT_MAX,
      settlementOverheadGas: NEW_SETTLEMENT_GAS_OVERHEAD,
      chainFamilySelector: Internal.CHAIN_FAMILY_SELECTOR_EVM,
      customExtraArgs: ""
    });

    vm.expectRevert(FastTransferTokenPoolAbstract.InvalidDestChainConfig.selector);
    s_pool.updateDestChainConfig(_singleConfigToList(laneConfigArgs));
  }

  function test_RevertWhen_InvalidPoolFeeBps() public {
    FastTransferTokenPoolAbstract.DestChainConfigUpdateArgs memory laneConfigArgs = FastTransferTokenPoolAbstract
      .DestChainConfigUpdateArgs({
      remoteChainSelector: NEW_CHAIN_SELECTOR,
      fastTransferFillerFeeBps: 100,
      fastTransferPoolFeeBps: 10_001, // > 10_000 (100%)
      fillerAllowlistEnabled: true,
      destinationPool: NEW_DESTINATION_POOL,
      maxFillAmountPerRequest: NEW_FILL_AMOUNT_MAX,
      settlementOverheadGas: NEW_SETTLEMENT_GAS_OVERHEAD,
      chainFamilySelector: Internal.CHAIN_FAMILY_SELECTOR_EVM,
      customExtraArgs: ""
    });

    vm.expectRevert(FastTransferTokenPoolAbstract.InvalidDestChainConfig.selector);
    s_pool.updateDestChainConfig(_singleConfigToList(laneConfigArgs));
  }

  function test_RevertWhen_NotOwners() public {
    FastTransferTokenPoolAbstract.DestChainConfigUpdateArgs memory laneConfigArgs;

    vm.stopPrank();

    vm.expectRevert();
    s_pool.updateDestChainConfig(_singleConfigToList(laneConfigArgs));
  }

  function test_updateDestChainConfig_ZeroFastFee() public {
    FastTransferTokenPoolAbstract.DestChainConfigUpdateArgs memory laneConfigArgs = FastTransferTokenPoolAbstract
      .DestChainConfigUpdateArgs({
      remoteChainSelector: NEW_CHAIN_SELECTOR,
      fastTransferFillerFeeBps: 0, // No fast fee
      fastTransferPoolFeeBps: 0, // No pool fee for this test
      fillerAllowlistEnabled: true,
      destinationPool: NEW_DESTINATION_POOL,
      maxFillAmountPerRequest: NEW_FILL_AMOUNT_MAX,
      settlementOverheadGas: NEW_SETTLEMENT_GAS_OVERHEAD,
      chainFamilySelector: Internal.CHAIN_FAMILY_SELECTOR_EVM,
      customExtraArgs: ""
    });

    s_pool.updateDestChainConfig(_singleConfigToList(laneConfigArgs));

    (FastTransferTokenPoolAbstract.DestChainConfig memory config,) = s_pool.getDestChainConfig(NEW_CHAIN_SELECTOR);
    assertEq(config.fastTransferFillerFeeBps, 0);
  }

  function test_updateDestChainConfig_MaxFastFee() public {
    FastTransferTokenPoolAbstract.DestChainConfigUpdateArgs memory laneConfigArgs = FastTransferTokenPoolAbstract
      .DestChainConfigUpdateArgs({
      remoteChainSelector: NEW_CHAIN_SELECTOR,
      fastTransferFillerFeeBps: 9_999, // 99.99% fee
      fastTransferPoolFeeBps: 0, // No pool fee for this test
      fillerAllowlistEnabled: true,
      destinationPool: NEW_DESTINATION_POOL,
      maxFillAmountPerRequest: NEW_FILL_AMOUNT_MAX,
      settlementOverheadGas: NEW_SETTLEMENT_GAS_OVERHEAD,
      chainFamilySelector: Internal.CHAIN_FAMILY_SELECTOR_EVM,
      customExtraArgs: ""
    });

    s_pool.updateDestChainConfig(_singleConfigToList(laneConfigArgs));

    (FastTransferTokenPoolAbstract.DestChainConfig memory config,) = s_pool.getDestChainConfig(NEW_CHAIN_SELECTOR);
    assertEq(config.fastTransferFillerFeeBps, 9_999);
  }

  function test_updateDestChainConfig_WithPoolFee() public {
    uint16 poolFeeBps = 150; // 1.5%

    FastTransferTokenPoolAbstract.DestChainConfigUpdateArgs memory laneConfigArgs = FastTransferTokenPoolAbstract
      .DestChainConfigUpdateArgs({
      remoteChainSelector: NEW_CHAIN_SELECTOR,
      fastTransferFillerFeeBps: NEW_FAST_FEE_FILLER_BPS,
      fastTransferPoolFeeBps: poolFeeBps,
      fillerAllowlistEnabled: false,
      destinationPool: NEW_DESTINATION_POOL,
      maxFillAmountPerRequest: NEW_FILL_AMOUNT_MAX,
      settlementOverheadGas: NEW_SETTLEMENT_GAS_OVERHEAD,
      chainFamilySelector: Internal.CHAIN_FAMILY_SELECTOR_EVM,
      customExtraArgs: ""
    });

    vm.expectEmit();
    emit FastTransferTokenPoolAbstract.DestChainConfigUpdated(
      NEW_CHAIN_SELECTOR,
      NEW_FAST_FEE_FILLER_BPS,
      poolFeeBps, // pool fee BPS
      NEW_FILL_AMOUNT_MAX,
      NEW_DESTINATION_POOL,
      Internal.CHAIN_FAMILY_SELECTOR_EVM,
      NEW_SETTLEMENT_GAS_OVERHEAD,
      false
    );

    s_pool.updateDestChainConfig(_singleConfigToList(laneConfigArgs));

    (FastTransferTokenPoolAbstract.DestChainConfig memory config,) = s_pool.getDestChainConfig(NEW_CHAIN_SELECTOR);
    assertEq(config.fastTransferFillerFeeBps, NEW_FAST_FEE_FILLER_BPS);
    assertEq(config.fastTransferPoolFeeBps, poolFeeBps);
    assertFalse(config.fillerAllowlistEnabled);
    assertEq(config.destinationPool, NEW_DESTINATION_POOL);
    assertEq(config.maxFillAmountPerRequest, NEW_FILL_AMOUNT_MAX);
  }

  function test_updateDestChainConfig_RevertWhen_TotalFeesExactly100Percent() public {
    uint16 fillerFee = 3_000; // 30%
    uint16 poolFee = 7_000; // 70% -> Total exactly 100%

    FastTransferTokenPoolAbstract.DestChainConfigUpdateArgs memory laneConfigArgs = FastTransferTokenPoolAbstract
      .DestChainConfigUpdateArgs({
      remoteChainSelector: NEW_CHAIN_SELECTOR,
      fastTransferFillerFeeBps: fillerFee,
      fastTransferPoolFeeBps: poolFee,
      fillerAllowlistEnabled: true,
      destinationPool: NEW_DESTINATION_POOL,
      maxFillAmountPerRequest: NEW_FILL_AMOUNT_MAX,
      settlementOverheadGas: NEW_SETTLEMENT_GAS_OVERHEAD,
      chainFamilySelector: Internal.CHAIN_FAMILY_SELECTOR_EVM,
      customExtraArgs: ""
    });

    vm.expectRevert(FastTransferTokenPoolAbstract.InvalidDestChainConfig.selector);
    s_pool.updateDestChainConfig(_singleConfigToList(laneConfigArgs));
  }

  function test_RevertWhen_TotalFeesExceed100Percent() public {
    FastTransferTokenPoolAbstract.DestChainConfigUpdateArgs memory laneConfigArgs = FastTransferTokenPoolAbstract
      .DestChainConfigUpdateArgs({
      remoteChainSelector: NEW_CHAIN_SELECTOR,
      fastTransferFillerFeeBps: 6_000, // 60%
      fastTransferPoolFeeBps: 5_000, // 50% -> Total 110%
      fillerAllowlistEnabled: true,
      destinationPool: NEW_DESTINATION_POOL,
      maxFillAmountPerRequest: NEW_FILL_AMOUNT_MAX,
      settlementOverheadGas: NEW_SETTLEMENT_GAS_OVERHEAD,
      chainFamilySelector: Internal.CHAIN_FAMILY_SELECTOR_EVM,
      customExtraArgs: ""
    });

    vm.expectRevert(FastTransferTokenPoolAbstract.InvalidDestChainConfig.selector);
    s_pool.updateDestChainConfig(_singleConfigToList(laneConfigArgs));
  }
}
