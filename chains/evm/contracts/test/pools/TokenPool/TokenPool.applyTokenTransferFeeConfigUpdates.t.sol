// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolV2} from "../../../interfaces/IPoolV2.sol";

import {Pool} from "../../../libraries/Pool.sol";
import {TokenPool} from "../../../pools/TokenPool.sol";
import {TokenPoolV2Setup} from "./TokenPoolV2Setup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract TokenPoolV2_applyTokenTransferFeeConfigUpdates is TokenPoolV2Setup {
  function test_applyTokenTransferFeeConfigUpdates() public {
    IPoolV2.TokenTransferFeeConfig memory feeConfig = IPoolV2.TokenTransferFeeConfig({
      destGasOverhead: 50_000,
      destBytesOverhead: Pool.CCIP_LOCK_OR_BURN_V1_RET_BYTES,
      defaultBlockConfirmationFeeUSDCents: 100, // $1.00
      customBlockConfirmationFeeUSDCents: 150, // $1.50
      defaultBlockConfirmationTransferFeeBps: 100, // 1%
      customBlockConfirmationTransferFeeBps: 200 // 2%
    });

    TokenPool.TokenTransferFeeConfigArgs[] memory feeConfigArgs = new TokenPool.TokenTransferFeeConfigArgs[](1);
    feeConfigArgs[0] =
      TokenPool.TokenTransferFeeConfigArgs({destChainSelector: DEST_CHAIN_SELECTOR, tokenTransferFeeConfig: feeConfig});

    uint64[] memory destToUseDefaultFeeConfigs = new uint64[](0);

    vm.expectEmit();
    emit TokenPool.TokenTransferFeeConfigUpdated(DEST_CHAIN_SELECTOR, feeConfig);
    s_tokenPool.applyTokenTransferFeeConfigUpdates(feeConfigArgs, destToUseDefaultFeeConfigs);
  }

  function test_applyTokenTransferFeeConfigUpdates_DeleteConfig() public {
    // First set a config
    IPoolV2.TokenTransferFeeConfig memory feeConfig = IPoolV2.TokenTransferFeeConfig({
      destGasOverhead: 50_000,
      destBytesOverhead: Pool.CCIP_LOCK_OR_BURN_V1_RET_BYTES,
      defaultBlockConfirmationFeeUSDCents: 100,
      customBlockConfirmationFeeUSDCents: 150,
      defaultBlockConfirmationTransferFeeBps: 100,
      customBlockConfirmationTransferFeeBps: 200
    });

    TokenPool.TokenTransferFeeConfigArgs[] memory feeConfigArgs = new TokenPool.TokenTransferFeeConfigArgs[](1);
    feeConfigArgs[0] =
      TokenPool.TokenTransferFeeConfigArgs({destChainSelector: DEST_CHAIN_SELECTOR, tokenTransferFeeConfig: feeConfig});

    s_tokenPool.applyTokenTransferFeeConfigUpdates(feeConfigArgs, new uint64[](0));

    // Now delete it
    uint64[] memory destToUseDefaultFeeConfigs = new uint64[](1);
    destToUseDefaultFeeConfigs[0] = DEST_CHAIN_SELECTOR;

    vm.expectEmit();
    emit TokenPool.TokenTransferFeeConfigDeleted(DEST_CHAIN_SELECTOR);
    s_tokenPool.applyTokenTransferFeeConfigUpdates(
      new TokenPool.TokenTransferFeeConfigArgs[](0), destToUseDefaultFeeConfigs
    );
  }

  // Reverts

  function test_applyTokenTransferFeeConfigUpdates_RevertWhen_OnlyCallableByOwner() public {
    vm.stopPrank();
    vm.prank(STRANGER);

    TokenPool.TokenTransferFeeConfigArgs[] memory feeConfigArgs = new TokenPool.TokenTransferFeeConfigArgs[](0);
    uint64[] memory destToUseDefaultFeeConfigs = new uint64[](0);

    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_tokenPool.applyTokenTransferFeeConfigUpdates(feeConfigArgs, destToUseDefaultFeeConfigs);
  }

  function test_applyTokenTransferFeeConfigUpdates_RevertWhen_DefaultBpsTooHigh() public {
    IPoolV2.TokenTransferFeeConfig memory feeConfig = IPoolV2.TokenTransferFeeConfig({
      destGasOverhead: 50_000,
      destBytesOverhead: Pool.CCIP_LOCK_OR_BURN_V1_RET_BYTES,
      defaultBlockConfirmationFeeUSDCents: 0,
      customBlockConfirmationFeeUSDCents: 0,
      defaultBlockConfirmationTransferFeeBps: uint16(BPS_DIVIDER),
      customBlockConfirmationTransferFeeBps: 0
    });

    TokenPool.TokenTransferFeeConfigArgs[] memory feeConfigArgs = new TokenPool.TokenTransferFeeConfigArgs[](1);
    feeConfigArgs[0] =
      TokenPool.TokenTransferFeeConfigArgs({destChainSelector: DEST_CHAIN_SELECTOR, tokenTransferFeeConfig: feeConfig});

    vm.expectRevert(abi.encodeWithSelector(TokenPool.InvalidTransferFeeBps.selector, BPS_DIVIDER));
    s_tokenPool.applyTokenTransferFeeConfigUpdates(feeConfigArgs, new uint64[](0));
  }

  function test_applyTokenTransferFeeConfigUpdates_RevertWhen_CustomBpsTooHigh() public {
    IPoolV2.TokenTransferFeeConfig memory feeConfig = IPoolV2.TokenTransferFeeConfig({
      destGasOverhead: 50_000,
      destBytesOverhead: Pool.CCIP_LOCK_OR_BURN_V1_RET_BYTES,
      defaultBlockConfirmationFeeUSDCents: 0,
      customBlockConfirmationFeeUSDCents: 0,
      defaultBlockConfirmationTransferFeeBps: 0,
      customBlockConfirmationTransferFeeBps: uint16(BPS_DIVIDER)
    });

    TokenPool.TokenTransferFeeConfigArgs[] memory feeConfigArgs = new TokenPool.TokenTransferFeeConfigArgs[](1);
    feeConfigArgs[0] =
      TokenPool.TokenTransferFeeConfigArgs({destChainSelector: DEST_CHAIN_SELECTOR, tokenTransferFeeConfig: feeConfig});

    vm.expectRevert(abi.encodeWithSelector(TokenPool.InvalidTransferFeeBps.selector, BPS_DIVIDER));
    s_tokenPool.applyTokenTransferFeeConfigUpdates(feeConfigArgs, new uint64[](0));
  }
}
