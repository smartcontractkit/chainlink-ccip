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
      customBlockConfirmationTransferFeeBps: 200, // 2%
      isEnabled: true
    });

    TokenPool.TokenTransferFeeConfigArgs[] memory feeConfigArgs = new TokenPool.TokenTransferFeeConfigArgs[](1);
    feeConfigArgs[0] =
      TokenPool.TokenTransferFeeConfigArgs({destChainSelector: DEST_CHAIN_SELECTOR, tokenTransferFeeConfig: feeConfig});

    uint64[] memory disableTokenTransferFeeConfigs = new uint64[](0);

    vm.expectEmit();
    emit TokenPool.TokenTransferFeeConfigUpdated(DEST_CHAIN_SELECTOR, feeConfig);
    s_tokenPool.applyTokenTransferFeeConfigUpdates(feeConfigArgs, disableTokenTransferFeeConfigs);
  }

  function test_applyTokenTransferFeeConfigUpdates_DeleteConfig() public {
    // First set a config
    IPoolV2.TokenTransferFeeConfig memory feeConfig = IPoolV2.TokenTransferFeeConfig({
      destGasOverhead: 50_000,
      destBytesOverhead: Pool.CCIP_LOCK_OR_BURN_V1_RET_BYTES,
      defaultBlockConfirmationFeeUSDCents: 100,
      customBlockConfirmationFeeUSDCents: 150,
      defaultBlockConfirmationTransferFeeBps: 100,
      customBlockConfirmationTransferFeeBps: 200,
      isEnabled: true
    });

    TokenPool.TokenTransferFeeConfigArgs[] memory feeConfigArgs = new TokenPool.TokenTransferFeeConfigArgs[](1);
    feeConfigArgs[0] =
      TokenPool.TokenTransferFeeConfigArgs({destChainSelector: DEST_CHAIN_SELECTOR, tokenTransferFeeConfig: feeConfig});

    s_tokenPool.applyTokenTransferFeeConfigUpdates(feeConfigArgs, new uint64[](0));

    // Now delete it
    uint64[] memory disableTokenTransferFeeConfigs = new uint64[](1);
    disableTokenTransferFeeConfigs[0] = DEST_CHAIN_SELECTOR;

    vm.expectEmit();
    emit TokenPool.TokenTransferFeeConfigDeleted(DEST_CHAIN_SELECTOR);
    s_tokenPool.applyTokenTransferFeeConfigUpdates(
      new TokenPool.TokenTransferFeeConfigArgs[](0), disableTokenTransferFeeConfigs
    );
  }

  // Reverts

  function test_applyTokenTransferFeeConfigUpdates_RevertWhen_OnlyCallableByOwner() public {
    vm.stopPrank();
    vm.prank(STRANGER);

    TokenPool.TokenTransferFeeConfigArgs[] memory feeConfigArgs = new TokenPool.TokenTransferFeeConfigArgs[](0);
    uint64[] memory disableTokenTransferFeeConfigs = new uint64[](0);

    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_tokenPool.applyTokenTransferFeeConfigUpdates(feeConfigArgs, disableTokenTransferFeeConfigs);
  }

  function test_applyTokenTransferFeeConfigUpdates_RevertWhen_DefaultBpsTooHigh() public {
    IPoolV2.TokenTransferFeeConfig memory feeConfig = IPoolV2.TokenTransferFeeConfig({
      destGasOverhead: 50_000,
      destBytesOverhead: Pool.CCIP_LOCK_OR_BURN_V1_RET_BYTES,
      defaultBlockConfirmationFeeUSDCents: 0,
      customBlockConfirmationFeeUSDCents: 0,
      defaultBlockConfirmationTransferFeeBps: uint16(BPS_DIVIDER),
      customBlockConfirmationTransferFeeBps: 0,
      isEnabled: true
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
      customBlockConfirmationTransferFeeBps: uint16(BPS_DIVIDER),
      isEnabled: true
    });

    TokenPool.TokenTransferFeeConfigArgs[] memory feeConfigArgs = new TokenPool.TokenTransferFeeConfigArgs[](1);
    feeConfigArgs[0] =
      TokenPool.TokenTransferFeeConfigArgs({destChainSelector: DEST_CHAIN_SELECTOR, tokenTransferFeeConfig: feeConfig});

    vm.expectRevert(abi.encodeWithSelector(TokenPool.InvalidTransferFeeBps.selector, BPS_DIVIDER));
    s_tokenPool.applyTokenTransferFeeConfigUpdates(feeConfigArgs, new uint64[](0));
  }

  function test_applyTokenTransferFeeConfigUpdates_RevertWhen_InvalidTokenTransferFeeConfig_EnabledWithZeroGasOverhead() public {
    IPoolV2.TokenTransferFeeConfig memory feeConfig = IPoolV2.TokenTransferFeeConfig({
      destGasOverhead: 0, // Zero gas overhead
      destBytesOverhead: Pool.CCIP_LOCK_OR_BURN_V1_RET_BYTES,
      defaultBlockConfirmationFeeUSDCents: 100,
      customBlockConfirmationFeeUSDCents: 150,
      defaultBlockConfirmationTransferFeeBps: 100,
      customBlockConfirmationTransferFeeBps: 200,
      isEnabled: true // Enabled with zero gas
    });

    TokenPool.TokenTransferFeeConfigArgs[] memory feeConfigArgs = new TokenPool.TokenTransferFeeConfigArgs[](1);
    feeConfigArgs[0] =
      TokenPool.TokenTransferFeeConfigArgs({destChainSelector: DEST_CHAIN_SELECTOR, tokenTransferFeeConfig: feeConfig});

    vm.expectRevert(abi.encodeWithSelector(TokenPool.InvalidTokenTransferFeeConfig.selector, DEST_CHAIN_SELECTOR));
    s_tokenPool.applyTokenTransferFeeConfigUpdates(feeConfigArgs, new uint64[](0));
  }

  function test_applyTokenTransferFeeConfigUpdates_RevertWhen_InvalidTokenTransferFeeConfig_EnabledWithZeroBytesOverhead() public {
    IPoolV2.TokenTransferFeeConfig memory feeConfig = IPoolV2.TokenTransferFeeConfig({
      destGasOverhead: 50_000,
      destBytesOverhead: 0, // Zero bytes overhead
      defaultBlockConfirmationFeeUSDCents: 100,
      customBlockConfirmationFeeUSDCents: 150,
      defaultBlockConfirmationTransferFeeBps: 100,
      customBlockConfirmationTransferFeeBps: 200,
      isEnabled: true // Enabled with zero bytes
    });

    TokenPool.TokenTransferFeeConfigArgs[] memory feeConfigArgs = new TokenPool.TokenTransferFeeConfigArgs[](1);
    feeConfigArgs[0] =
      TokenPool.TokenTransferFeeConfigArgs({destChainSelector: DEST_CHAIN_SELECTOR, tokenTransferFeeConfig: feeConfig});

    vm.expectRevert(abi.encodeWithSelector(TokenPool.InvalidTokenTransferFeeConfig.selector, DEST_CHAIN_SELECTOR));
    s_tokenPool.applyTokenTransferFeeConfigUpdates(feeConfigArgs, new uint64[](0));
  }

  function test_applyTokenTransferFeeConfigUpdates_RevertWhen_InvalidTokenTransferFeeConfig_IsEnabledFalse() public {
    IPoolV2.TokenTransferFeeConfig memory feeConfig = IPoolV2.TokenTransferFeeConfig({
      destGasOverhead: 50_000,
      destBytesOverhead: Pool.CCIP_LOCK_OR_BURN_V1_RET_BYTES,
      defaultBlockConfirmationFeeUSDCents: 100,
      customBlockConfirmationFeeUSDCents: 150,
      defaultBlockConfirmationTransferFeeBps: 100,
      customBlockConfirmationTransferFeeBps: 200,
      isEnabled: false // Cannot set isEnabled: false directly
    });

    TokenPool.TokenTransferFeeConfigArgs[] memory feeConfigArgs = new TokenPool.TokenTransferFeeConfigArgs[](1);
    feeConfigArgs[0] =
      TokenPool.TokenTransferFeeConfigArgs({destChainSelector: DEST_CHAIN_SELECTOR, tokenTransferFeeConfig: feeConfig});

    vm.expectRevert(abi.encodeWithSelector(TokenPool.InvalidTokenTransferFeeConfig.selector, DEST_CHAIN_SELECTOR));
    s_tokenPool.applyTokenTransferFeeConfigUpdates(feeConfigArgs, new uint64[](0));
  }
}
