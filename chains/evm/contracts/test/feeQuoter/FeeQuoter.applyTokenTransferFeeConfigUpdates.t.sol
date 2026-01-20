// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {FeeQuoter} from "../../FeeQuoter.sol";
import {Pool} from "../../libraries/Pool.sol";
import {FeeQuoterSetup} from "./FeeQuoterSetup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract FeeQuoter_applyTokenTransferFeeConfigUpdates is FeeQuoterSetup {
  function testFuzz_applyTokenTransferFeeConfigUpdates(
    FeeQuoter.TokenTransferFeeConfig[2] memory tokenTransferFeeConfigs
  ) public {
    // To prevent Invalid Fee Range error from the fuzzer, bound the results to a valid range that
    // where minFee < maxFee
    tokenTransferFeeConfigs[0].feeUSDCents = uint32(bound(tokenTransferFeeConfigs[0].feeUSDCents, 0, type(uint8).max));
    tokenTransferFeeConfigs[1].feeUSDCents = uint32(bound(tokenTransferFeeConfigs[1].feeUSDCents, 0, type(uint8).max));

    FeeQuoter.TokenTransferFeeConfigArgs[] memory tokenTransferFeeConfigArgs = _generateTokenTransferFeeConfigArgs(2, 2);
    tokenTransferFeeConfigArgs[0].destChainSelector = DEST_CHAIN_SELECTOR;
    tokenTransferFeeConfigArgs[1].destChainSelector = DEST_CHAIN_SELECTOR + 1;

    for (uint256 i = 0; i < tokenTransferFeeConfigArgs.length; ++i) {
      for (uint256 j = 0; j < tokenTransferFeeConfigs.length; ++j) {
        tokenTransferFeeConfigs[j].destBytesOverhead = uint32(
          bound(tokenTransferFeeConfigs[j].destBytesOverhead, Pool.CCIP_LOCK_OR_BURN_V1_RET_BYTES, type(uint32).max)
        );
        address feeToken = s_sourceTokens[j];
        tokenTransferFeeConfigArgs[i].tokenTransferFeeConfigs[j].token = feeToken;
        tokenTransferFeeConfigArgs[i].tokenTransferFeeConfigs[j].tokenTransferFeeConfig = tokenTransferFeeConfigs[j];

        vm.expectEmit();
        emit FeeQuoter.TokenTransferFeeConfigUpdated(
          tokenTransferFeeConfigArgs[i].destChainSelector, feeToken, tokenTransferFeeConfigs[j]
        );
      }
    }

    s_feeQuoter.applyTokenTransferFeeConfigUpdates(
      tokenTransferFeeConfigArgs, new FeeQuoter.TokenTransferFeeConfigRemoveArgs[](0)
    );

    for (uint256 i = 0; i < tokenTransferFeeConfigs.length; ++i) {
      _assertTokenTransferFeeConfigEqual(
        tokenTransferFeeConfigs[i],
        s_feeQuoter.getTokenTransferFeeConfig(
          tokenTransferFeeConfigArgs[0].destChainSelector,
          tokenTransferFeeConfigArgs[0].tokenTransferFeeConfigs[i].token
        )
      );
    }
  }

  function test_applyTokenTransferFeeConfigUpdates() public {
    FeeQuoter.TokenTransferFeeConfigArgs[] memory tokenTransferFeeConfigArgs = _generateTokenTransferFeeConfigArgs(1, 2);
    tokenTransferFeeConfigArgs[0].destChainSelector = DEST_CHAIN_SELECTOR;
    tokenTransferFeeConfigArgs[0].tokenTransferFeeConfigs[0].token = address(5);
    tokenTransferFeeConfigArgs[0].tokenTransferFeeConfigs[0].tokenTransferFeeConfig =
      FeeQuoter.TokenTransferFeeConfig({feeUSDCents: 6, destGasOverhead: 9, destBytesOverhead: 312, isEnabled: true});
    tokenTransferFeeConfigArgs[0].tokenTransferFeeConfigs[1].token = address(11);
    tokenTransferFeeConfigArgs[0].tokenTransferFeeConfigs[1].tokenTransferFeeConfig =
      FeeQuoter.TokenTransferFeeConfig({feeUSDCents: 12, destGasOverhead: 15, destBytesOverhead: 394, isEnabled: true});

    vm.expectEmit();
    emit FeeQuoter.TokenTransferFeeConfigUpdated(
      tokenTransferFeeConfigArgs[0].destChainSelector,
      tokenTransferFeeConfigArgs[0].tokenTransferFeeConfigs[0].token,
      tokenTransferFeeConfigArgs[0].tokenTransferFeeConfigs[0].tokenTransferFeeConfig
    );
    vm.expectEmit();
    emit FeeQuoter.TokenTransferFeeConfigUpdated(
      tokenTransferFeeConfigArgs[0].destChainSelector,
      tokenTransferFeeConfigArgs[0].tokenTransferFeeConfigs[1].token,
      tokenTransferFeeConfigArgs[0].tokenTransferFeeConfigs[1].tokenTransferFeeConfig
    );

    FeeQuoter.TokenTransferFeeConfigRemoveArgs[] memory tokensToRemove =
      new FeeQuoter.TokenTransferFeeConfigRemoveArgs[](0);
    s_feeQuoter.applyTokenTransferFeeConfigUpdates(tokenTransferFeeConfigArgs, tokensToRemove);

    FeeQuoter.TokenTransferFeeConfig memory config0 = s_feeQuoter.getTokenTransferFeeConfig(
      tokenTransferFeeConfigArgs[0].destChainSelector, tokenTransferFeeConfigArgs[0].tokenTransferFeeConfigs[0].token
    );
    FeeQuoter.TokenTransferFeeConfig memory config1 = s_feeQuoter.getTokenTransferFeeConfig(
      tokenTransferFeeConfigArgs[0].destChainSelector, tokenTransferFeeConfigArgs[0].tokenTransferFeeConfigs[1].token
    );

    _assertTokenTransferFeeConfigEqual(
      tokenTransferFeeConfigArgs[0].tokenTransferFeeConfigs[0].tokenTransferFeeConfig, config0
    );
    _assertTokenTransferFeeConfigEqual(
      tokenTransferFeeConfigArgs[0].tokenTransferFeeConfigs[1].tokenTransferFeeConfig, config1
    );

    // Remove only the first token and validate only the first token is removed
    tokensToRemove = new FeeQuoter.TokenTransferFeeConfigRemoveArgs[](1);
    tokensToRemove[0] = FeeQuoter.TokenTransferFeeConfigRemoveArgs({
      destChainSelector: tokenTransferFeeConfigArgs[0].destChainSelector,
      token: tokenTransferFeeConfigArgs[0].tokenTransferFeeConfigs[0].token
    });

    vm.expectEmit();
    emit FeeQuoter.TokenTransferFeeConfigDeleted(
      tokenTransferFeeConfigArgs[0].destChainSelector, tokenTransferFeeConfigArgs[0].tokenTransferFeeConfigs[0].token
    );

    s_feeQuoter.applyTokenTransferFeeConfigUpdates(new FeeQuoter.TokenTransferFeeConfigArgs[](0), tokensToRemove);

    config0 = s_feeQuoter.getTokenTransferFeeConfig(
      tokenTransferFeeConfigArgs[0].destChainSelector, tokenTransferFeeConfigArgs[0].tokenTransferFeeConfigs[0].token
    );
    config1 = s_feeQuoter.getTokenTransferFeeConfig(
      tokenTransferFeeConfigArgs[0].destChainSelector, tokenTransferFeeConfigArgs[0].tokenTransferFeeConfigs[1].token
    );

    FeeQuoter.TokenTransferFeeConfig memory emptyConfig;

    _assertTokenTransferFeeConfigEqual(emptyConfig, config0);
    _assertTokenTransferFeeConfigEqual(
      tokenTransferFeeConfigArgs[0].tokenTransferFeeConfigs[1].tokenTransferFeeConfig, config1
    );
  }

  function test_getAllTokenTransferFeeConfigs() public {
    // Set up token transfer fee configs for a chain selector
    uint64 testChainSelector = DEST_CHAIN_SELECTOR + 20;
    FeeQuoter.TokenTransferFeeConfigArgs[] memory tokenConfigArgs = new FeeQuoter.TokenTransferFeeConfigArgs[](1);
    tokenConfigArgs[0].destChainSelector = testChainSelector;
    tokenConfigArgs[0].tokenTransferFeeConfigs = new FeeQuoter.TokenTransferFeeConfigSingleTokenArgs[](2);

    // Use source tokens as transfer tokens
    address transferToken0 = makeAddr("transferToken0");
    address transferToken1 = makeAddr("transferToken1");
    tokenConfigArgs[0].tokenTransferFeeConfigs[0].token = transferToken0;
    tokenConfigArgs[0].tokenTransferFeeConfigs[0].tokenTransferFeeConfig = FeeQuoter.TokenTransferFeeConfig({
      feeUSDCents: 100, destGasOverhead: 50000, destBytesOverhead: 64, isEnabled: true
    });

    tokenConfigArgs[0].tokenTransferFeeConfigs[1].token = transferToken1;
    tokenConfigArgs[0].tokenTransferFeeConfigs[1].tokenTransferFeeConfig = FeeQuoter.TokenTransferFeeConfig({
      feeUSDCents: 200, destGasOverhead: 60000, destBytesOverhead: 96, isEnabled: true
    });

    s_feeQuoter.applyTokenTransferFeeConfigUpdates(tokenConfigArgs, new FeeQuoter.TokenTransferFeeConfigRemoveArgs[](0));

    // Get all token transfer fee configs for all chain selectors
    (
      uint64[] memory destChainSelectors,
      address[][] memory transferTokens,
      FeeQuoter.TokenTransferFeeConfig[][] memory tokenTransferFeeConfigs
    ) = s_feeQuoter.getAllTokenTransferFeeConfigs();

    // Find the index of our test chain selector
    uint256 chainIndex = type(uint256).max;
    for (uint256 i = 0; i < destChainSelectors.length; ++i) {
      if (destChainSelectors[i] == testChainSelector) {
        chainIndex = i;
        break;
      }
    }
    require(chainIndex != type(uint256).max, "Chain selector not found");

    address[] memory tokens = transferTokens[chainIndex];
    FeeQuoter.TokenTransferFeeConfig[] memory configs = tokenTransferFeeConfigs[chainIndex];

    // Verify we got the transfer tokens we configured
    assertGe(tokens.length, 2, "Should return at least the configured transfer tokens");

    // Find our configured tokens and verify their configs
    bool foundToken0 = false;
    bool foundToken1 = false;
    for (uint256 i = 0; i < tokens.length; ++i) {
      if (tokens[i] == transferToken0) {
        _assertTokenTransferFeeConfigEqual(
          tokenConfigArgs[0].tokenTransferFeeConfigs[0].tokenTransferFeeConfig, configs[i]
        );
        foundToken0 = true;
      }
      if (tokens[i] == transferToken1) {
        _assertTokenTransferFeeConfigEqual(
          tokenConfigArgs[0].tokenTransferFeeConfigs[1].tokenTransferFeeConfig, configs[i]
        );
        foundToken1 = true;
      }
    }
    assertTrue(foundToken0, "Should find first configured transfer token");
    assertTrue(foundToken1, "Should find second configured transfer token");
    assertEq(tokens.length, configs.length, "Arrays should have same length");
  }

  function test_getAllTokenTransferFeeConfigs_ReturnsEmptyStructsForUnconfiguredTokens() public view {
    uint64 testChainSelector = DEST_CHAIN_SELECTOR + 30;

    // Don't set any transfer token configs, just get all transfer tokens
    (
      uint64[] memory destChainSelectors,
      address[][] memory transferTokens,
      FeeQuoter.TokenTransferFeeConfig[][] memory tokenTransferFeeConfigs
    ) = s_feeQuoter.getAllTokenTransferFeeConfigs();

    // Find the index of our test chain selector
    uint256 chainIndex = type(uint256).max;
    for (uint256 i = 0; i < destChainSelectors.length; ++i) {
      if (destChainSelectors[i] == testChainSelector) {
        chainIndex = i;
        break;
      }
    }
    // If chain selector not found, it means no transfer token configs exist for it, which is expected
    if (chainIndex == type(uint256).max) {
      return;
    }

    address[] memory tokens = transferTokens[chainIndex];
    FeeQuoter.TokenTransferFeeConfig[] memory configs = tokenTransferFeeConfigs[chainIndex];

    // Should return transfer tokens - if any exist for this chain selector
    assertEq(tokens.length, configs.length, "Arrays should have same length");

    // Verify all configs are empty (isEnabled should be false)
    for (uint256 i = 0; i < configs.length; ++i) {
      assertFalse(configs[i].isEnabled, "Unconfigured transfer tokens should have isEnabled = false");
      assertEq(configs[i].feeUSDCents, 0, "Unconfigured transfer tokens should have zero feeUSDCents");
      assertEq(configs[i].destGasOverhead, 0, "Unconfigured transfer tokens should have zero destGasOverhead");
      assertEq(configs[i].destBytesOverhead, 0, "Unconfigured transfer tokens should have zero destBytesOverhead");
    }
  }

  // Reverts

  function test_applyTokenTransferFeeConfigUpdates_RevertWhen_OnlyCallableByOwner() public {
    vm.startPrank(STRANGER);
    FeeQuoter.TokenTransferFeeConfigArgs[] memory tokenTransferFeeConfigArgs;

    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);

    s_feeQuoter.applyTokenTransferFeeConfigUpdates(
      tokenTransferFeeConfigArgs, new FeeQuoter.TokenTransferFeeConfigRemoveArgs[](0)
    );
  }

  function test_applyTokenTransferFeeConfigUpdates_RevertWhen_InvalidDestBytesOverhead() public {
    FeeQuoter.TokenTransferFeeConfigArgs[] memory tokenTransferFeeConfigArgs = _generateTokenTransferFeeConfigArgs(1, 1);
    tokenTransferFeeConfigArgs[0].destChainSelector = DEST_CHAIN_SELECTOR;
    tokenTransferFeeConfigArgs[0].tokenTransferFeeConfigs[0].token = address(5);
    tokenTransferFeeConfigArgs[0].tokenTransferFeeConfigs[0].tokenTransferFeeConfig = FeeQuoter.TokenTransferFeeConfig({
      feeUSDCents: 6,
      destGasOverhead: 9,
      destBytesOverhead: uint32(Pool.CCIP_LOCK_OR_BURN_V1_RET_BYTES - 1),
      isEnabled: true
    });

    vm.expectRevert(
      abi.encodeWithSelector(
        FeeQuoter.InvalidDestBytesOverhead.selector,
        tokenTransferFeeConfigArgs[0].tokenTransferFeeConfigs[0].token,
        tokenTransferFeeConfigArgs[0].tokenTransferFeeConfigs[0].tokenTransferFeeConfig.destBytesOverhead
      )
    );

    s_feeQuoter.applyTokenTransferFeeConfigUpdates(
      tokenTransferFeeConfigArgs, new FeeQuoter.TokenTransferFeeConfigRemoveArgs[](0)
    );
  }
}
