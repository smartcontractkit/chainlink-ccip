// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {OffRampOverSuperchainInterop} from "../../../offRamp/OffRampOverSuperchainInterop.sol";
import {OffRampOverSuperchainInteropSetup} from "./OffRampOverSuperchainInteropSetup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract OffRampOverSuperchainInterop_applyChainSelectorToChainIdConfigUpdates is OffRampOverSuperchainInteropSetup {
  function test_applyChainSelectorToChainIdConfigUpdates_SetMappings() public {
    OffRampOverSuperchainInterop.ChainSelectorToChainIdConfigArgs[] memory configs = 
      new OffRampOverSuperchainInterop.ChainSelectorToChainIdConfigArgs[](2);
    
    configs[0] = OffRampOverSuperchainInterop.ChainSelectorToChainIdConfigArgs({
      chainSelector: SOURCE_CHAIN_SELECTOR_2,
      chainId: SOURCE_CHAIN_ID_2
    });
    
    configs[1] = OffRampOverSuperchainInterop.ChainSelectorToChainIdConfigArgs({
      chainSelector: SOURCE_CHAIN_SELECTOR_3,
      chainId: SOURCE_CHAIN_ID_3
    });
    
    vm.expectEmit();
    emit OffRampOverSuperchainInterop.ChainSelectorToChainIdConfigAdded(SOURCE_CHAIN_SELECTOR_2, SOURCE_CHAIN_ID_2);
    vm.expectEmit();
    emit OffRampOverSuperchainInterop.ChainSelectorToChainIdConfigAdded(SOURCE_CHAIN_SELECTOR_3, SOURCE_CHAIN_ID_3);
    
    s_offRamp.applyChainSelectorToChainIdConfigUpdates(new uint64[](0), configs);
    
    // Verify mappings
    assertEq(s_offRamp.getChainIdBySourceChainSelector(SOURCE_CHAIN_SELECTOR_2), SOURCE_CHAIN_ID_2);
    assertEq(s_offRamp.getChainIdBySourceChainSelector(SOURCE_CHAIN_SELECTOR_3), SOURCE_CHAIN_ID_3);
  }

  function test_applyChainSelectorToChainIdConfigUpdates_UnsetMappings() public {
    // First set a mapping
    OffRampOverSuperchainInterop.ChainSelectorToChainIdConfigArgs[] memory configs = 
      new OffRampOverSuperchainInterop.ChainSelectorToChainIdConfigArgs[](1);
    configs[0] = OffRampOverSuperchainInterop.ChainSelectorToChainIdConfigArgs({
      chainSelector: SOURCE_CHAIN_SELECTOR_2,
      chainId: SOURCE_CHAIN_ID_2
    });
    s_offRamp.applyChainSelectorToChainIdConfigUpdates(new uint64[](0), configs);
    
    // Now unset it
    uint64[] memory chainSelectorsToUnset = new uint64[](1);
    chainSelectorsToUnset[0] = SOURCE_CHAIN_SELECTOR_2;
    
    vm.expectEmit();
    emit OffRampOverSuperchainInterop.ChainSelectorToChainIdConfigRemoved(SOURCE_CHAIN_SELECTOR_2, SOURCE_CHAIN_ID_2);
    
    s_offRamp.applyChainSelectorToChainIdConfigUpdates(chainSelectorsToUnset, new OffRampOverSuperchainInterop.ChainSelectorToChainIdConfigArgs[](0));
    
    // Verify mapping removed
    assertEq(s_offRamp.getChainIdBySourceChainSelector(SOURCE_CHAIN_SELECTOR_2), 0);
  }

  function test_applyChainSelectorToChainIdConfigUpdates_UpdateExisting() public {
    // Update the existing SOURCE_CHAIN_SELECTOR_1 mapping
    OffRampOverSuperchainInterop.ChainSelectorToChainIdConfigArgs[] memory configs = 
      new OffRampOverSuperchainInterop.ChainSelectorToChainIdConfigArgs[](1);
    
    uint256 newChainId = 999;
    configs[0] = OffRampOverSuperchainInterop.ChainSelectorToChainIdConfigArgs({
      chainSelector: SOURCE_CHAIN_SELECTOR_1,
      chainId: newChainId
    });
    
    vm.expectEmit();
    emit OffRampOverSuperchainInterop.ChainSelectorToChainIdConfigAdded(SOURCE_CHAIN_SELECTOR_1, newChainId);
    
    s_offRamp.applyChainSelectorToChainIdConfigUpdates(new uint64[](0), configs);
    
    // Verify mapping updated
    assertEq(s_offRamp.getChainIdBySourceChainSelector(SOURCE_CHAIN_SELECTOR_1), newChainId);
  }

  function test_applyChainSelectorToChainIdConfigUpdates_RevertWhen_ZeroChainId() public {
    OffRampOverSuperchainInterop.ChainSelectorToChainIdConfigArgs[] memory configs = 
      new OffRampOverSuperchainInterop.ChainSelectorToChainIdConfigArgs[](1);
    
    configs[0] = OffRampOverSuperchainInterop.ChainSelectorToChainIdConfigArgs({
      chainSelector: SOURCE_CHAIN_SELECTOR_2,
      chainId: 0
    });
    
    vm.expectRevert(OffRampOverSuperchainInterop.ZeroChainIdNotAllowed.selector);
    s_offRamp.applyChainSelectorToChainIdConfigUpdates(new uint64[](0), configs);
  }

  function test_applyChainSelectorToChainIdConfigUpdates_RevertWhen_ZeroChainSelector() public {
    OffRampOverSuperchainInterop.ChainSelectorToChainIdConfigArgs[] memory configs = 
      new OffRampOverSuperchainInterop.ChainSelectorToChainIdConfigArgs[](1);
    
    configs[0] = OffRampOverSuperchainInterop.ChainSelectorToChainIdConfigArgs({
      chainSelector: 0,
      chainId: SOURCE_CHAIN_ID_2
    });
    
    vm.expectRevert(OffRampOverSuperchainInterop.ZeroChainSelectorNotAllowed.selector);
    s_offRamp.applyChainSelectorToChainIdConfigUpdates(new uint64[](0), configs);
  }

  function test_applyChainSelectorToChainIdConfigUpdates_RevertWhen_NotOwner() public {
    OffRampOverSuperchainInterop.ChainSelectorToChainIdConfigArgs[] memory configs = 
      new OffRampOverSuperchainInterop.ChainSelectorToChainIdConfigArgs[](1);
    
    configs[0] = OffRampOverSuperchainInterop.ChainSelectorToChainIdConfigArgs({
      chainSelector: SOURCE_CHAIN_SELECTOR_2,
      chainId: SOURCE_CHAIN_ID_2
    });
    
    vm.stopPrank();
    vm.prank(STRANGER);
    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_offRamp.applyChainSelectorToChainIdConfigUpdates(new uint64[](0), configs);
  }
}