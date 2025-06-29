// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {OffRamp} from "../../../offRamp/OffRamp.sol";
import {OffRampOverSuperchainInterop} from "../../../offRamp/OffRampOverSuperchainInterop.sol";
import {OffRampOverSuperchainInteropSetup} from "./OffRampOverSuperchainInteropSetup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract OffRampOverSuperchainInterop_applyChainSelectorToChainIdConfigUpdates is OffRampOverSuperchainInteropSetup {
  function test_applyChainSelectorToChainIdConfigUpdates() public {
    OffRampOverSuperchainInterop.ChainSelectorToChainIdConfigArgs[] memory chainSelectorsUpdates =
      new OffRampOverSuperchainInterop.ChainSelectorToChainIdConfigArgs[](3);
    chainSelectorsUpdates[0] = OffRampOverSuperchainInterop.ChainSelectorToChainIdConfigArgs({
      chainSelector: SOURCE_CHAIN_SELECTOR_1,
      chainId: CHAIN_ID_1
    });
    chainSelectorsUpdates[1] = OffRampOverSuperchainInterop.ChainSelectorToChainIdConfigArgs({
      chainSelector: SOURCE_CHAIN_SELECTOR_2,
      chainId: CHAIN_ID_2
    });
    chainSelectorsUpdates[2] = OffRampOverSuperchainInterop.ChainSelectorToChainIdConfigArgs({
      chainSelector: SOURCE_CHAIN_SELECTOR_3,
      chainId: CHAIN_ID_3
    });

    vm.expectEmit();
    emit OffRampOverSuperchainInterop.ChainSelectorToChainIdConfigUpdated(SOURCE_CHAIN_SELECTOR_1, CHAIN_ID_1);
    vm.expectEmit();
    emit OffRampOverSuperchainInterop.ChainSelectorToChainIdConfigUpdated(SOURCE_CHAIN_SELECTOR_2, CHAIN_ID_2);
    vm.expectEmit();
    emit OffRampOverSuperchainInterop.ChainSelectorToChainIdConfigUpdated(SOURCE_CHAIN_SELECTOR_3, CHAIN_ID_3);

    s_offRampOverSuperchainInterop.applyChainSelectorToChainIdConfigUpdates(new uint64[](0), chainSelectorsUpdates);

    // Verify the config added successfully
    assertEq(CHAIN_ID_1, s_offRampOverSuperchainInterop.getChainId(SOURCE_CHAIN_SELECTOR_1));
    assertEq(CHAIN_ID_2, s_offRampOverSuperchainInterop.getChainId(SOURCE_CHAIN_SELECTOR_2));
    assertEq(CHAIN_ID_3, s_offRampOverSuperchainInterop.getChainId(SOURCE_CHAIN_SELECTOR_3));
  }

  function test_applyChainSelectorToChainIdConfigUpdates_AddUpdateRemove() public {
    // First add another chain selector
    OffRampOverSuperchainInterop.ChainSelectorToChainIdConfigArgs[] memory initialConfig =
      new OffRampOverSuperchainInterop.ChainSelectorToChainIdConfigArgs[](1);
    initialConfig[0] = OffRampOverSuperchainInterop.ChainSelectorToChainIdConfigArgs({
      chainSelector: SOURCE_CHAIN_SELECTOR_2,
      chainId: CHAIN_ID_2
    });
    s_offRampOverSuperchainInterop.applyChainSelectorToChainIdConfigUpdates(new uint64[](0), initialConfig);

    assertEq(CHAIN_ID_2, s_offRampOverSuperchainInterop.getChainId(SOURCE_CHAIN_SELECTOR_2));

    // Remove SOURCE_CHAIN_SELECTOR_1, update SOURCE_CHAIN_SELECTOR_2, add SOURCE_CHAIN_SELECTOR_3
    uint64[] memory chainSelectorsToRemove = new uint64[](1);
    chainSelectorsToRemove[0] = SOURCE_CHAIN_SELECTOR_1;

    OffRampOverSuperchainInterop.ChainSelectorToChainIdConfigArgs[] memory chainSelectorsUpdates =
      new OffRampOverSuperchainInterop.ChainSelectorToChainIdConfigArgs[](2);
    chainSelectorsUpdates[0] = OffRampOverSuperchainInterop.ChainSelectorToChainIdConfigArgs({
      chainSelector: SOURCE_CHAIN_SELECTOR_2,
      chainId: CHAIN_ID_2 + 1
    });
    chainSelectorsUpdates[1] = OffRampOverSuperchainInterop.ChainSelectorToChainIdConfigArgs({
      chainSelector: SOURCE_CHAIN_SELECTOR_3,
      chainId: CHAIN_ID_3
    });

    vm.expectEmit();
    emit OffRampOverSuperchainInterop.ChainSelectorToChainIdConfigRemoved(SOURCE_CHAIN_SELECTOR_1, CHAIN_ID_1);
    vm.expectEmit();
    emit OffRampOverSuperchainInterop.ChainSelectorToChainIdConfigUpdated(SOURCE_CHAIN_SELECTOR_2, CHAIN_ID_2 + 1);
    vm.expectEmit();
    emit OffRampOverSuperchainInterop.ChainSelectorToChainIdConfigUpdated(SOURCE_CHAIN_SELECTOR_3, CHAIN_ID_3);
    s_offRampOverSuperchainInterop.applyChainSelectorToChainIdConfigUpdates(
      chainSelectorsToRemove, chainSelectorsUpdates
    );

    assertEq(0, s_offRampOverSuperchainInterop.getChainId(SOURCE_CHAIN_SELECTOR_1));
    assertEq(CHAIN_ID_2 + 1, s_offRampOverSuperchainInterop.getChainId(SOURCE_CHAIN_SELECTOR_2));
    assertEq(CHAIN_ID_3, s_offRampOverSuperchainInterop.getChainId(SOURCE_CHAIN_SELECTOR_3));
  }

  function test_applyChainSelectorToChainIdConfigUpdates_RemoveNonExistentConfig() public {
    uint64 nonExistentChainSelector = 99999;
    uint64[] memory chainSelectorsToRemove = new uint64[](1);
    chainSelectorsToRemove[0] = nonExistentChainSelector;

    vm.recordLogs();

    s_offRampOverSuperchainInterop.applyChainSelectorToChainIdConfigUpdates(
      chainSelectorsToRemove, new OffRampOverSuperchainInterop.ChainSelectorToChainIdConfigArgs[](0)
    );

    // No events should be emitted
    assertEq(vm.getRecordedLogs().length, 0);
  }

  function test_applyChainSelectorToChainIdConfigUpdates_EmptyInput() public {
    vm.recordLogs();

    s_offRampOverSuperchainInterop.applyChainSelectorToChainIdConfigUpdates(
      new uint64[](0), new OffRampOverSuperchainInterop.ChainSelectorToChainIdConfigArgs[](0)
    );

    // No events should be emitted
    assertEq(vm.getRecordedLogs().length, 0);
  }

  // Reverts

  function test_applyChainSelectorToChainIdConfigUpdates_RevertWhen_NotOwner() public {
    vm.startPrank(STRANGER);

    OffRampOverSuperchainInterop.ChainSelectorToChainIdConfigArgs[] memory chainSelectorsUpdate =
      new OffRampOverSuperchainInterop.ChainSelectorToChainIdConfigArgs[](1);
    chainSelectorsUpdate[0] =
      OffRampOverSuperchainInterop.ChainSelectorToChainIdConfigArgs({chainSelector: 10, chainId: 1000});

    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_offRampOverSuperchainInterop.applyChainSelectorToChainIdConfigUpdates(new uint64[](0), chainSelectorsUpdate);
  }

  function test_applyChainSelectorToChainIdConfigUpdates_RevertWhen_ChainIdOrSelectorZero() public {
    OffRampOverSuperchainInterop.ChainSelectorToChainIdConfigArgs[] memory chainSelectorsUpdate =
      new OffRampOverSuperchainInterop.ChainSelectorToChainIdConfigArgs[](1);
    chainSelectorsUpdate[0] =
      OffRampOverSuperchainInterop.ChainSelectorToChainIdConfigArgs({chainSelector: 10, chainId: 0});

    vm.expectRevert(OffRampOverSuperchainInterop.ZeroChainIdNotAllowed.selector);
    s_offRampOverSuperchainInterop.applyChainSelectorToChainIdConfigUpdates(new uint64[](0), chainSelectorsUpdate);

    chainSelectorsUpdate[0] =
      OffRampOverSuperchainInterop.ChainSelectorToChainIdConfigArgs({chainSelector: 0, chainId: 1000});

    vm.expectRevert(OffRamp.ZeroChainSelectorNotAllowed.selector);
    s_offRampOverSuperchainInterop.applyChainSelectorToChainIdConfigUpdates(new uint64[](0), chainSelectorsUpdate);
  }
}
