// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {FinalityCodec} from "../../../../../libraries/FinalityCodec.sol";
import {PolicyConfig} from "../../../../../libraries/PolicyConfig.sol";
import {ConnectionPolicyStore} from "../../../../../pools/MultiTokenPool/policies/ConnectionPolicyStore.sol";
import {ConnectionPolicyStoreSetup} from "./ConnectionPolicyStoreSetup.t.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

contract ConnectionPolicyStore_applyConnectionPolicyConfigUpdates is ConnectionPolicyStoreSetup {
  function test_applyConnectionPolicyConfigUpdates() public {
    ConnectionPolicyStore.ConnectionPolicyConfigArgs[] memory policyConfigs =
      new ConnectionPolicyStore.ConnectionPolicyConfigArgs[](2);
    policyConfigs[0] = _getConnectionPolicyConfigArgs(OTHER_POLICY, DEST_CHAIN_SELECTOR, false)[0];
    policyConfigs[1] = _getConnectionPolicyConfigArgs(OTHER_POLICY, SOURCE_CHAIN_SELECTOR, true)[0];

    vm.expectEmit();
    emit ConnectionPolicyStore.ConnectionPolicyConfigSet(OTHER_POLICY, DEST_CHAIN_SELECTOR, false);
    vm.expectEmit();
    emit ConnectionPolicyStore.ConnectionPolicyConfigSet(OTHER_POLICY, SOURCE_CHAIN_SELECTOR, true);
    s_connectionStore.applyConnectionPolicyConfigUpdates(
      new PolicyConfig.PolicyRecordKey[](0), new bytes32[](0), policyConfigs
    );

    // Adding the first record creates the policy without a provider.
    (
      uint32 tokenCount,
      address remoteTokenAddressProvider,
      ConnectionPolicyStore.ConnectionPolicyConfigArgs[] memory storedConfigs
    ) = s_connectionStore.getConnectionPolicy(OTHER_POLICY);
    assertEq(0, tokenCount);
    assertEq(address(0), remoteTokenAddressProvider);
    assertEq(abi.encode(policyConfigs), abi.encode(storedConfigs));
    (bytes32[] memory connectionPolicyNames,) = s_connectionStore.getPolicyNames();
    assertEq(2, connectionPolicyNames.length);
    assertEq(CONNECTION_POLICY, connectionPolicyNames[0]);
    assertEq(OTHER_POLICY, connectionPolicyNames[1]);
  }

  function test_applyConnectionPolicyConfigUpdates_ReplaceRecord() public {
    // Enabling the subset restriction excludes s_token because the list is empty.
    vm.expectEmit();
    emit ConnectionPolicyStore.ConnectionPolicyConfigSet(CONNECTION_POLICY, DEST_CHAIN_SELECTOR, true);
    _setConnectionPolicyConfig(_getConnectionPolicyConfigArgs(CONNECTION_POLICY, DEST_CHAIN_SELECTOR, true));

    (, ConnectionPolicyStore.ConnectionPolicyConfigArgs[] memory storedConfigs) = _getConnectionPolicyRecords();
    assertEq(1, storedConfigs.length);
    assertTrue(storedConfigs[0].onlySubsetTokens);
    assertFalse(s_connectionStore.hasLaneBatch(_getLaneKeys(s_token, DEST_CHAIN_SELECTOR))[0]);
  }

  function test_applyConnectionPolicyConfigUpdates_SameRecordTwice() public {
    ConnectionPolicyStore.ConnectionPolicyConfigArgs[] memory policyConfigs =
      new ConnectionPolicyStore.ConnectionPolicyConfigArgs[](2);
    policyConfigs[0] = _getConnectionPolicyConfigArgs(CONNECTION_POLICY, DEST_CHAIN_SELECTOR, true)[0];
    policyConfigs[1] = _getConnectionPolicyConfigArgs(CONNECTION_POLICY, DEST_CHAIN_SELECTOR, false)[0];

    _setConnectionPolicyConfig(policyConfigs);

    (, ConnectionPolicyStore.ConnectionPolicyConfigArgs[] memory storedConfigs) = _getConnectionPolicyRecords();
    assertEq(1, storedConfigs.length);
    assertEq(abi.encode(policyConfigs[1]), abi.encode(storedConfigs[0]));
  }

  function test_applyConnectionPolicyConfigUpdates_RecordSetKeepsSubsetTokens() public {
    _setConnectionPolicyConfig(_getConnectionPolicyConfigArgs(CONNECTION_POLICY, DEST_CHAIN_SELECTOR, true));
    _addSubsetToken(CONNECTION_POLICY, DEST_CHAIN_SELECTOR, s_token);
    _setConnectionPolicyConfig(_getConnectionPolicyConfigArgs(CONNECTION_POLICY, DEST_CHAIN_SELECTOR, false));

    // Disabling and re-enabling the restriction preserves the subset list.
    _setConnectionPolicyConfig(_getConnectionPolicyConfigArgs(CONNECTION_POLICY, DEST_CHAIN_SELECTOR, true));

    (bool onlySubsetTokens, address[] memory subsetTokens) =
      s_connectionStore.getSubsetTokens(CONNECTION_POLICY, DEST_CHAIN_SELECTOR, 0, type(uint64).max);
    assertTrue(onlySubsetTokens);
    assertEq(_arrayOf(s_token), subsetTokens);
    assertTrue(s_connectionStore.hasLaneBatch(_getLaneKeys(s_token, DEST_CHAIN_SELECTOR))[0]);
  }

  function test_applyConnectionPolicyConfigUpdates_RemoveRecord() public {
    vm.expectEmit();
    emit ConnectionPolicyStore.ConnectionPolicyConfigRemoved(CONNECTION_POLICY, DEST_CHAIN_SELECTOR);
    s_connectionStore.applyConnectionPolicyConfigUpdates(
      _getPolicyRecordKeys(CONNECTION_POLICY, DEST_CHAIN_SELECTOR),
      new bytes32[](0),
      new ConnectionPolicyStore.ConnectionPolicyConfigArgs[](0)
    );

    // Removing a chain record preserves the policy, assigned tokens, and provider.
    (
      uint32 tokenCount,
      address remoteTokenAddressProvider,
      ConnectionPolicyStore.ConnectionPolicyConfigArgs[] memory storedConfigs
    ) = s_connectionStore.getConnectionPolicy(CONNECTION_POLICY);
    assertEq(1, tokenCount);
    assertEq(address(s_remoteTokenAddressProvider), remoteTokenAddressProvider);
    assertEq(0, storedConfigs.length);
    assertFalse(s_connectionStore.hasLaneBatch(_getLaneKeys(s_token, DEST_CHAIN_SELECTOR))[0]);
    vm.expectRevert(
      abi.encodeWithSelector(ConnectionPolicyStore.ConnectionRecordNotFound.selector, s_token, DEST_CHAIN_SELECTOR)
    );
    s_connectionStore.resolveOutboundLane(s_token, DEST_CHAIN_SELECTOR, FinalityCodec.WAIT_FOR_FINALITY_FLAG, "");
  }

  function test_applyConnectionPolicyConfigUpdates_RemoveRecordDeletesSubsetTokens() public {
    _setTokenConnectionPolicy(s_otherToken, CONNECTION_POLICY);
    _setConnectionPolicyConfig(_getConnectionPolicyConfigArgs(CONNECTION_POLICY, DEST_CHAIN_SELECTOR, true));
    _addSubsetToken(CONNECTION_POLICY, DEST_CHAIN_SELECTOR, s_token);
    _addSubsetToken(CONNECTION_POLICY, DEST_CHAIN_SELECTOR, s_otherToken);

    s_connectionStore.applyConnectionPolicyConfigUpdates(
      _getPolicyRecordKeys(CONNECTION_POLICY, DEST_CHAIN_SELECTOR),
      new bytes32[](0),
      new ConnectionPolicyStore.ConnectionPolicyConfigArgs[](0)
    );

    (bool onlySubsetTokens, address[] memory subsetTokens) =
      s_connectionStore.getSubsetTokens(CONNECTION_POLICY, DEST_CHAIN_SELECTOR, 0, type(uint64).max);
    assertFalse(onlySubsetTokens);
    assertEq(0, subsetTokens.length);

    // Recreating the record must not restore its old subset list.
    _setConnectionPolicyConfig(_getConnectionPolicyConfigArgs(CONNECTION_POLICY, DEST_CHAIN_SELECTOR, true));
    (, subsetTokens) = s_connectionStore.getSubsetTokens(CONNECTION_POLICY, DEST_CHAIN_SELECTOR, 0, type(uint64).max);
    assertEq(0, subsetTokens.length);
    assertFalse(s_connectionStore.hasLaneBatch(_getLaneKeys(s_token, DEST_CHAIN_SELECTOR))[0]);
  }

  function test_applyConnectionPolicyConfigUpdates_RemovePolicy() public {
    _setTokenConnectionPolicy(s_token, bytes32(0));
    vm.expectEmit();
    emit ConnectionPolicyStore.ConnectionPolicyConfigRemoved(CONNECTION_POLICY, DEST_CHAIN_SELECTOR);
    vm.expectEmit();
    emit ConnectionPolicyStore.ConnectionPolicyRemoved(CONNECTION_POLICY);
    s_connectionStore.applyConnectionPolicyConfigUpdates(
      _getPolicyRecordKeys(CONNECTION_POLICY, DEST_CHAIN_SELECTOR),
      _getPolicyNames(CONNECTION_POLICY),
      new ConnectionPolicyStore.ConnectionPolicyConfigArgs[](0)
    );

    (bytes32[] memory connectionPolicyNames,) = s_connectionStore.getPolicyNames();
    assertEq(0, connectionPolicyNames.length);
    // Recreating the policy must not restore its old provider.
    _setConnectionPolicyConfig(_getConnectionPolicyConfigArgs(CONNECTION_POLICY, DEST_CHAIN_SELECTOR, false));
    (, address remoteTokenAddressProvider,) = s_connectionStore.getConnectionPolicy(CONNECTION_POLICY);
    assertEq(address(0), remoteTokenAddressProvider);
  }

  function test_applyConnectionPolicyConfigUpdates_RemovalsRunFirst() public {
    _setConnectionPolicyConfig(_getConnectionPolicyConfigArgs(CONNECTION_POLICY, DEST_CHAIN_SELECTOR, true));
    _addSubsetToken(CONNECTION_POLICY, DEST_CHAIN_SELECTOR, s_token);
    ConnectionPolicyStore.ConnectionPolicyConfigArgs[] memory policyConfigs =
      _getConnectionPolicyConfigArgs(CONNECTION_POLICY, DEST_CHAIN_SELECTOR, true);

    vm.expectEmit();
    emit ConnectionPolicyStore.ConnectionPolicyConfigRemoved(CONNECTION_POLICY, DEST_CHAIN_SELECTOR);
    vm.expectEmit();
    emit ConnectionPolicyStore.ConnectionPolicyConfigSet(CONNECTION_POLICY, DEST_CHAIN_SELECTOR, true);
    s_connectionStore.applyConnectionPolicyConfigUpdates(
      _getPolicyRecordKeys(CONNECTION_POLICY, DEST_CHAIN_SELECTOR), new bytes32[](0), policyConfigs
    );

    (, ConnectionPolicyStore.ConnectionPolicyConfigArgs[] memory storedConfigs) = _getConnectionPolicyRecords();
    assertEq(abi.encode(policyConfigs), abi.encode(storedConfigs));
    (, address[] memory subsetTokens) =
      s_connectionStore.getSubsetTokens(CONNECTION_POLICY, DEST_CHAIN_SELECTOR, 0, type(uint64).max);
    assertEq(0, subsetTokens.length);
  }

  // Reverts

  function test_applyConnectionPolicyConfigUpdates_RevertWhen_ZeroChainSelectorNotAllowed() public {
    ConnectionPolicyStore.ConnectionPolicyConfigArgs[] memory policyConfigs =
      _getConnectionPolicyConfigArgs(CONNECTION_POLICY, PolicyConfig.DEFAULT_CHAIN_SELECTOR, false);

    vm.expectRevert(PolicyConfig.ZeroChainSelectorNotAllowed.selector);
    s_connectionStore.applyConnectionPolicyConfigUpdates(
      new PolicyConfig.PolicyRecordKey[](0), new bytes32[](0), policyConfigs
    );
  }

  function test_applyConnectionPolicyConfigUpdates_RevertWhen_ZeroPolicyNameNotAllowed() public {
    ConnectionPolicyStore.ConnectionPolicyConfigArgs[] memory policyConfigs =
      _getConnectionPolicyConfigArgs(bytes32(0), DEST_CHAIN_SELECTOR, false);

    vm.expectRevert(PolicyConfig.ZeroPolicyNameNotAllowed.selector);
    s_connectionStore.applyConnectionPolicyConfigUpdates(
      new PolicyConfig.PolicyRecordKey[](0), new bytes32[](0), policyConfigs
    );
  }

  function test_applyConnectionPolicyConfigUpdates_RevertWhen_PolicyRecordNotFound() public {
    vm.expectRevert(
      abi.encodeWithSelector(PolicyConfig.PolicyRecordNotFound.selector, CONNECTION_POLICY, SOURCE_CHAIN_SELECTOR)
    );
    s_connectionStore.applyConnectionPolicyConfigUpdates(
      _getPolicyRecordKeys(CONNECTION_POLICY, SOURCE_CHAIN_SELECTOR),
      new bytes32[](0),
      new ConnectionPolicyStore.ConnectionPolicyConfigArgs[](0)
    );
  }

  function test_applyConnectionPolicyConfigUpdates_RevertWhen_PolicyNotFound() public {
    // REMOTE_POOL_POLICY exists only as a remote pool policy.
    vm.expectRevert(abi.encodeWithSelector(PolicyConfig.PolicyNotFound.selector, REMOTE_POOL_POLICY));
    s_connectionStore.applyConnectionPolicyConfigUpdates(
      new PolicyConfig.PolicyRecordKey[](0),
      _getPolicyNames(REMOTE_POOL_POLICY),
      new ConnectionPolicyStore.ConnectionPolicyConfigArgs[](0)
    );
  }

  function test_applyConnectionPolicyConfigUpdates_RevertWhen_PolicyInUse() public {
    vm.expectRevert(abi.encodeWithSelector(PolicyConfig.PolicyInUse.selector, CONNECTION_POLICY));
    s_connectionStore.applyConnectionPolicyConfigUpdates(
      _getPolicyRecordKeys(CONNECTION_POLICY, DEST_CHAIN_SELECTOR),
      _getPolicyNames(CONNECTION_POLICY),
      new ConnectionPolicyStore.ConnectionPolicyConfigArgs[](0)
    );
  }

  function test_applyConnectionPolicyConfigUpdates_RevertWhen_PolicyHasRecords() public {
    _setTokenConnectionPolicy(s_token, bytes32(0));

    vm.expectRevert(abi.encodeWithSelector(PolicyConfig.PolicyHasRecords.selector, CONNECTION_POLICY));
    s_connectionStore.applyConnectionPolicyConfigUpdates(
      new PolicyConfig.PolicyRecordKey[](0),
      _getPolicyNames(CONNECTION_POLICY),
      new ConnectionPolicyStore.ConnectionPolicyConfigArgs[](0)
    );
  }

  function test_applyConnectionPolicyConfigUpdates_RevertWhen_UnauthorizedCaller() public {
    ConnectionPolicyStore.ConnectionPolicyConfigArgs[] memory policyConfigs =
      _getConnectionPolicyConfigArgs(OTHER_POLICY, DEST_CHAIN_SELECTOR, false);
    vm.startPrank(STRANGER);

    vm.expectRevert(abi.encodeWithSelector(AuthorizedCallers.UnauthorizedCaller.selector, STRANGER));
    s_connectionStore.applyConnectionPolicyConfigUpdates(
      new PolicyConfig.PolicyRecordKey[](0), new bytes32[](0), policyConfigs
    );
  }

  function _getConnectionPolicyRecords()
    private
    view
    returns (uint32 tokenCount, ConnectionPolicyStore.ConnectionPolicyConfigArgs[] memory policyConfigs)
  {
    (tokenCount,, policyConfigs) = s_connectionStore.getConnectionPolicy(CONNECTION_POLICY);
    return (tokenCount, policyConfigs);
  }
}
