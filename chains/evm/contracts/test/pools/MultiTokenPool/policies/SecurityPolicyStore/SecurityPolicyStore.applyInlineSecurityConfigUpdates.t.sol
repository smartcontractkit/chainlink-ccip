// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../../../../interfaces/IPoolConfigManager.sol";

import {CCVConfigValidation} from "../../../../../libraries/CCVConfigValidation.sol";
import {PolicyConfig} from "../../../../../libraries/PolicyConfig.sol";
import {SecurityPolicyStore} from "../../../../../pools/MultiTokenPool/policies/SecurityPolicyStore.sol";
import {SecurityPolicyStoreSetup} from "./SecurityPolicyStoreSetup.t.sol";

contract SecurityPolicyStore_applyInlineSecurityConfigUpdates is SecurityPolicyStoreSetup {
  /// @notice One call sets an inline security record. The two CCV lists differ, so a swap of the two is caught on the
  /// read-back.
  function test_applyInlineSecurityConfigUpdates_SetsRecord() public {
    SecurityPolicyStore.InlineSecurityConfigArgs[] memory configs =
      new SecurityPolicyStore.InlineSecurityConfigArgs[](1);
    configs[0] = SecurityPolicyStore.InlineSecurityConfigArgs({
      localToken: s_tokenA,
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      outboundCCVs: _ccvsOf(s_ccv1, s_ccv2),
      inboundCCVs: _arrayOf(s_ccv2),
      hook: s_hook
    });

    vm.expectEmit(address(s_store));
    emit SecurityPolicyStore.InlineSecurityConfigSet(
      s_tokenA, DEST_CHAIN_SELECTOR, configs[0].outboundCCVs, configs[0].inboundCCVs, s_hook
    );
    _setInlineSecurityConfig(s_tokenA, DEST_CHAIN_SELECTOR, configs[0].outboundCCVs, configs[0].inboundCCVs, s_hook);

    SecurityPolicyStore.InlineSecurityConfigArgs[] memory records = s_store.getInlineSecurityConfigs(s_tokenA);
    assertEq(1, records.length);
    assertEq(abi.encode(configs[0]), abi.encode(records[0]));
  }

  /// @notice A set replaces an existing record, and a removal deletes it.
  function test_applyInlineSecurityConfigUpdates_ReplacesAndRemovesRecord() public {
    _setInlineSecurityConfig(s_tokenA, DEST_CHAIN_SELECTOR, _arrayOf(s_ccv1), _arrayOf(s_ccv2), s_hook);
    _setInlineSecurityConfig(s_tokenA, DEST_CHAIN_SELECTOR, _arrayOf(s_ccv2), _arrayOf(s_ccv1), address(0));

    SecurityPolicyStore.InlineSecurityConfigArgs[] memory records = s_store.getInlineSecurityConfigs(s_tokenA);
    assertEq(address(0), records[0].hook);
    assertEq(abi.encode(_arrayOf(s_ccv2)), abi.encode(records[0].outboundCCVs));

    vm.expectEmit(address(s_store));
    emit SecurityPolicyStore.InlineSecurityConfigRemoved(s_tokenA, DEST_CHAIN_SELECTOR);
    vm.prank(s_poolConfigManager);
    s_store.applyInlineSecurityConfigUpdates(
      _lanesOf(s_tokenA, DEST_CHAIN_SELECTOR), new SecurityPolicyStore.InlineSecurityConfigArgs[](0)
    );
    assertEq(0, s_store.getInlineSecurityConfigs(s_tokenA).length);
  }

  // Reverts

  /// @notice A CCV list has no duplicate.
  function test_applyInlineSecurityConfigUpdates_RevertWhen_DuplicateCCVNotAllowed() public {
    vm.expectRevert(abi.encodeWithSelector(CCVConfigValidation.DuplicateCCVNotAllowed.selector, s_ccv1));
    _setInlineSecurityConfig(s_tokenA, DEST_CHAIN_SELECTOR, _ccvsOf(s_ccv1, s_ccv1), _arrayOf(s_ccv2), s_hook);
  }

  /// @notice A token has no default record, so an inline record needs a chain.
  function test_applyInlineSecurityConfigUpdates_RevertWhen_ZeroChainSelectorNotAllowed() public {
    vm.expectRevert(PolicyConfig.ZeroChainSelectorNotAllowed.selector);
    _setInlineSecurityConfig(s_tokenA, PolicyConfig.DEFAULT_CHAIN_SELECTOR, _arrayOf(s_ccv1), _arrayOf(s_ccv2), s_hook);
  }

  /// @notice An inline record needs a registered token.
  function test_applyInlineSecurityConfigUpdates_RevertWhen_LocalTokenNotRegistered() public {
    address unknownToken = makeAddr("unknownToken");
    vm.expectRevert(abi.encodeWithSelector(PolicyConfig.LocalTokenNotRegistered.selector, unknownToken));
    _setInlineSecurityConfig(unknownToken, DEST_CHAIN_SELECTOR, _arrayOf(s_ccv1), _arrayOf(s_ccv2), s_hook);
  }

  /// @notice A removal names an inline record that exists.
  function test_applyInlineSecurityConfigUpdates_RevertWhen_InlineRecordNotFound() public {
    vm.expectRevert(abi.encodeWithSelector(PolicyConfig.InlineRecordNotFound.selector, s_tokenA, DEST_CHAIN_SELECTOR));
    vm.prank(s_poolConfigManager);
    s_store.applyInlineSecurityConfigUpdates(
      _lanesOf(s_tokenA, DEST_CHAIN_SELECTOR), new SecurityPolicyStore.InlineSecurityConfigArgs[](0)
    );
  }
}
