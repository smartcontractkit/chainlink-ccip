// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../../../../interfaces/IPoolConfigManager.sol";
import {IRemoteTokenAddressProvider} from "../../../../../interfaces/IRemoteTokenAddressProvider.sol";

import {FinalityCodec} from "../../../../../libraries/FinalityCodec.sol";
import {PolicyConfig} from "../../../../../libraries/PolicyConfig.sol";
import {ConnectionPolicyStore} from "../../../../../pools/MultiTokenPool/policies/ConnectionPolicyStore.sol";
import {ConnectionPolicyStoreSetup} from "./ConnectionPolicyStoreSetup.t.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

contract ConnectionPolicyStore_applyInlineConnectionConfigUpdates is ConnectionPolicyStoreSetup {
  bytes internal s_inlineRemoteToken = abi.encode(makeAddr("inlineRemoteToken"));

  function test_applyInlineConnectionConfigUpdates() public {
    ConnectionPolicyStore.InlineConnectionConfigArgs[] memory inlineConfigs =
      new ConnectionPolicyStore.InlineConnectionConfigArgs[](2);
    inlineConfigs[0] = _getInlineConnectionConfigArgs(s_otherToken, DEST_CHAIN_SELECTOR, s_inlineRemoteToken)[0];
    // Remote token addresses need not use the EVM address length.
    inlineConfigs[1] = _getInlineConnectionConfigArgs(s_otherToken, SOURCE_CHAIN_SELECTOR, hex"01")[0];

    vm.expectEmit();
    emit ConnectionPolicyStore.InlineConnectionConfigSet(s_otherToken, DEST_CHAIN_SELECTOR, s_inlineRemoteToken);
    vm.expectEmit();
    emit ConnectionPolicyStore.InlineConnectionConfigSet(s_otherToken, SOURCE_CHAIN_SELECTOR, hex"01");
    s_connectionStore.applyInlineConnectionConfigUpdates(new IPoolConfigManager.LaneKey[](0), inlineConfigs);

    assertEq(abi.encode(inlineConfigs), abi.encode(s_connectionStore.getInlineConnectionConfigs(s_otherToken)));
    assertTrue(s_connectionStore.hasLaneBatch(_getLaneKeys(s_otherToken, DEST_CHAIN_SELECTOR))[0]);
    assertTrue(s_connectionStore.hasLaneBatch(_getLaneKeys(s_otherToken, SOURCE_CHAIN_SELECTOR))[0]);
  }

  function test_applyInlineConnectionConfigUpdates_InlineRecordOverridesPolicy() public {
    _setInlineConnectionConfig(_getInlineConnectionConfigArgs(s_token, DEST_CHAIN_SELECTOR, s_inlineRemoteToken));

    assertEq(
      s_inlineRemoteToken,
      s_connectionStore.resolveOutboundLane(s_token, DEST_CHAIN_SELECTOR, FinalityCodec.WAIT_FOR_FINALITY_FLAG, "")
      .remoteToken
    );
  }

  function test_applyInlineConnectionConfigUpdates_ReplaceRecord() public {
    _setInlineConnectionConfig(_getInlineConnectionConfigArgs(s_token, DEST_CHAIN_SELECTOR, s_inlineRemoteToken));
    ConnectionPolicyStore.InlineConnectionConfigArgs[] memory inlineConfigs =
      _getInlineConnectionConfigArgs(s_token, DEST_CHAIN_SELECTOR, abi.encode(makeAddr("newRemoteToken")));

    // Updating an inline address does not require it to match the policy's address.
    _setInlineConnectionConfig(inlineConfigs);

    assertEq(abi.encode(inlineConfigs), abi.encode(s_connectionStore.getInlineConnectionConfigs(s_token)));
  }

  function test_applyInlineConnectionConfigUpdates_RemoveInlineRecord() public {
    _setInlineConnectionConfig(_getInlineConnectionConfigArgs(s_otherToken, DEST_CHAIN_SELECTOR, s_inlineRemoteToken));

    vm.expectEmit();
    emit ConnectionPolicyStore.InlineConnectionConfigRemoved(s_otherToken, DEST_CHAIN_SELECTOR);
    s_connectionStore.applyInlineConnectionConfigUpdates(
      _getLaneKeys(s_otherToken, DEST_CHAIN_SELECTOR), new ConnectionPolicyStore.InlineConnectionConfigArgs[](0)
    );

    assertEq(0, s_connectionStore.getInlineConnectionConfigs(s_otherToken).length);
    // The token has no policy, so the lane closes.
    assertFalse(s_connectionStore.hasLaneBatch(_getLaneKeys(s_otherToken, DEST_CHAIN_SELECTOR))[0]);
  }

  function test_applyInlineConnectionConfigUpdates_RemoveWhenPolicyGivesSameRemoteToken() public {
    _setInlineConnectionConfig(_getInlineConnectionConfigArgs(s_token, DEST_CHAIN_SELECTOR, abi.encode(s_token)));

    _removeInlineConnectionConfig(s_token, DEST_CHAIN_SELECTOR);

    IPoolConfigManager.LaneConnection memory connection = _getLaneConnection(s_token, DEST_CHAIN_SELECTOR);
    assertTrue(connection.hasConnection);
    assertEq(abi.encode(s_token), connection.remoteToken);
  }

  function test_applyInlineConnectionConfigUpdates_RemoveWhenPolicyDoesNotServeToken() public {
    // The policy has no record for SOURCE_CHAIN_SELECTOR.
    _setInlineConnectionConfig(_getInlineConnectionConfigArgs(s_token, SOURCE_CHAIN_SELECTOR, s_inlineRemoteToken));
    _removeInlineConnectionConfig(s_token, SOURCE_CHAIN_SELECTOR);

    // The subset restriction excludes s_token from this policy connection.
    _setConnectionPolicyConfig(_getConnectionPolicyConfigArgs(CONNECTION_POLICY, DEST_CHAIN_SELECTOR, true));
    _setInlineConnectionConfig(_getInlineConnectionConfigArgs(s_token, DEST_CHAIN_SELECTOR, s_inlineRemoteToken));
    _removeInlineConnectionConfig(s_token, DEST_CHAIN_SELECTOR);

    assertEq(0, s_connectionStore.getInlineConnectionConfigs(s_token).length);
    assertFalse(s_connectionStore.hasLaneBatch(_getLaneKeys(s_token, DEST_CHAIN_SELECTOR))[0]);
  }

  /// @notice Removal is allowed when the policy has no usable remote token because outbound transfers would revert.
  function test_applyInlineConnectionConfigUpdates_RemoveWhenPolicyGivesNoRemoteToken() public {
    address remoteTokenAddressProvider = address(s_remoteTokenAddressProvider);
    bytes memory getRemoteTokenAddressCall =
      abi.encodeCall(IRemoteTokenAddressProvider.getRemoteTokenAddress, (s_token, DEST_CHAIN_SELECTOR));

    vm.mockCall(remoteTokenAddressProvider, getRemoteTokenAddressCall, abi.encode(bytes("")));
    _setInlineConnectionConfig(_getInlineConnectionConfigArgs(s_token, DEST_CHAIN_SELECTOR, s_inlineRemoteToken));
    _removeInlineConnectionConfig(s_token, DEST_CHAIN_SELECTOR);

    vm.mockCall(remoteTokenAddressProvider, getRemoteTokenAddressCall, abi.encode(abi.encode(address(0))));
    _setInlineConnectionConfig(_getInlineConnectionConfigArgs(s_token, DEST_CHAIN_SELECTOR, s_inlineRemoteToken));
    _removeInlineConnectionConfig(s_token, DEST_CHAIN_SELECTOR);

    vm.mockCallRevert(remoteTokenAddressProvider, getRemoteTokenAddressCall, "");
    _setInlineConnectionConfig(_getInlineConnectionConfigArgs(s_token, DEST_CHAIN_SELECTOR, s_inlineRemoteToken));
    _removeInlineConnectionConfig(s_token, DEST_CHAIN_SELECTOR);

    // The provider had code when configured, but no longer does.
    vm.clearMockedCalls();
    vm.etch(remoteTokenAddressProvider, "");
    _setInlineConnectionConfig(_getInlineConnectionConfigArgs(s_token, DEST_CHAIN_SELECTOR, s_inlineRemoteToken));
    _removeInlineConnectionConfig(s_token, DEST_CHAIN_SELECTOR);

    _setRemoteTokenAddressProvider(CONNECTION_POLICY, address(0));
    _setInlineConnectionConfig(_getInlineConnectionConfigArgs(s_token, DEST_CHAIN_SELECTOR, s_inlineRemoteToken));
    _removeInlineConnectionConfig(s_token, DEST_CHAIN_SELECTOR);

    assertEq(0, s_connectionStore.getInlineConnectionConfigs(s_token).length);
  }

  function test_applyInlineConnectionConfigUpdates_RemovalsRunFirst() public {
    _setInlineConnectionConfig(_getInlineConnectionConfigArgs(s_otherToken, DEST_CHAIN_SELECTOR, s_inlineRemoteToken));
    ConnectionPolicyStore.InlineConnectionConfigArgs[] memory inlineConfigs =
      _getInlineConnectionConfigArgs(s_otherToken, DEST_CHAIN_SELECTOR, abi.encode(makeAddr("newRemoteToken")));

    IPoolConfigManager.LaneKey[] memory inlineConfigsToRemove = _getLaneKeys(s_otherToken, DEST_CHAIN_SELECTOR);
    s_connectionStore.applyInlineConnectionConfigUpdates(inlineConfigsToRemove, inlineConfigs);

    assertEq(abi.encode(inlineConfigs), abi.encode(s_connectionStore.getInlineConnectionConfigs(s_otherToken)));
  }

  // Reverts

  function test_applyInlineConnectionConfigUpdates_RevertWhen_PolicyRemoteTokenMismatch() public {
    _setInlineConnectionConfig(_getInlineConnectionConfigArgs(s_token, DEST_CHAIN_SELECTOR, s_inlineRemoteToken));

    vm.expectRevert(
      abi.encodeWithSelector(ConnectionPolicyStore.PolicyRemoteTokenMismatch.selector, s_token, DEST_CHAIN_SELECTOR)
    );
    s_connectionStore.applyInlineConnectionConfigUpdates(
      _getLaneKeys(s_token, DEST_CHAIN_SELECTOR), new ConnectionPolicyStore.InlineConnectionConfigArgs[](0)
    );

    // Re-adding the inline record in the same call does not bypass the removal check.
    vm.expectRevert(
      abi.encodeWithSelector(ConnectionPolicyStore.PolicyRemoteTokenMismatch.selector, s_token, DEST_CHAIN_SELECTOR)
    );
    s_connectionStore.applyInlineConnectionConfigUpdates(
      _getLaneKeys(s_token, DEST_CHAIN_SELECTOR),
      _getInlineConnectionConfigArgs(s_token, DEST_CHAIN_SELECTOR, s_inlineRemoteToken)
    );
  }

  function test_applyInlineConnectionConfigUpdates_RevertWhen_ProviderAnswerDoesNotDecode() public {
    _setInlineConnectionConfig(_getInlineConnectionConfigArgs(s_token, DEST_CHAIN_SELECTOR, s_inlineRemoteToken));
    vm.mockCall(
      address(s_remoteTokenAddressProvider),
      abi.encodeCall(IRemoteTokenAddressProvider.getRemoteTokenAddress, (s_token, DEST_CHAIN_SELECTOR)),
      hex"01"
    );

    // ABI decoding errors are not caught by try/catch.
    vm.expectRevert(bytes(""));
    s_connectionStore.applyInlineConnectionConfigUpdates(
      _getLaneKeys(s_token, DEST_CHAIN_SELECTOR), new ConnectionPolicyStore.InlineConnectionConfigArgs[](0)
    );
  }

  function test_applyInlineConnectionConfigUpdates_RevertWhen_InvalidRemoteToken() public {
    bytes[] memory remoteTokens = new bytes[](3);
    remoteTokens[0] = "";
    remoteTokens[1] = abi.encode(address(0));
    remoteTokens[2] = hex"00";

    for (uint256 i = 0; i < remoteTokens.length; ++i) {
      ConnectionPolicyStore.InlineConnectionConfigArgs[] memory inlineConfigs =
        _getInlineConnectionConfigArgs(s_otherToken, DEST_CHAIN_SELECTOR, remoteTokens[i]);

      vm.expectRevert(
        abi.encodeWithSelector(ConnectionPolicyStore.InvalidRemoteToken.selector, s_otherToken, DEST_CHAIN_SELECTOR)
      );
      s_connectionStore.applyInlineConnectionConfigUpdates(new IPoolConfigManager.LaneKey[](0), inlineConfigs);
    }
  }

  function test_applyInlineConnectionConfigUpdates_RevertWhen_ZeroChainSelectorNotAllowed() public {
    ConnectionPolicyStore.InlineConnectionConfigArgs[] memory inlineConfigs =
      _getInlineConnectionConfigArgs(s_otherToken, PolicyConfig.DEFAULT_CHAIN_SELECTOR, s_inlineRemoteToken);

    vm.expectRevert(PolicyConfig.ZeroChainSelectorNotAllowed.selector);
    s_connectionStore.applyInlineConnectionConfigUpdates(new IPoolConfigManager.LaneKey[](0), inlineConfigs);
  }

  function test_applyInlineConnectionConfigUpdates_RevertWhen_LocalTokenNotRegistered() public {
    address unknownToken = makeAddr("unknownToken");
    ConnectionPolicyStore.InlineConnectionConfigArgs[] memory inlineConfigs =
      _getInlineConnectionConfigArgs(unknownToken, DEST_CHAIN_SELECTOR, s_inlineRemoteToken);

    vm.expectRevert(abi.encodeWithSelector(PolicyConfig.LocalTokenNotRegistered.selector, unknownToken));
    s_connectionStore.applyInlineConnectionConfigUpdates(new IPoolConfigManager.LaneKey[](0), inlineConfigs);
  }

  function test_applyInlineConnectionConfigUpdates_RevertWhen_InlineRecordNotFound() public {
    // Report the missing inline record before the policy's mismatched remote token.
    vm.expectRevert(abi.encodeWithSelector(PolicyConfig.InlineRecordNotFound.selector, s_token, DEST_CHAIN_SELECTOR));
    s_connectionStore.applyInlineConnectionConfigUpdates(
      _getLaneKeys(s_token, DEST_CHAIN_SELECTOR), new ConnectionPolicyStore.InlineConnectionConfigArgs[](0)
    );
  }

  function test_applyInlineConnectionConfigUpdates_RevertWhen_UnauthorizedCaller() public {
    ConnectionPolicyStore.InlineConnectionConfigArgs[] memory inlineConfigs =
      _getInlineConnectionConfigArgs(s_otherToken, DEST_CHAIN_SELECTOR, s_inlineRemoteToken);
    vm.startPrank(STRANGER);

    vm.expectRevert(abi.encodeWithSelector(AuthorizedCallers.UnauthorizedCaller.selector, STRANGER));
    s_connectionStore.applyInlineConnectionConfigUpdates(new IPoolConfigManager.LaneKey[](0), inlineConfigs);
  }

  function _removeInlineConnectionConfig(
    address localToken,
    uint64 remoteChainSelector
  ) private {
    s_connectionStore.applyInlineConnectionConfigUpdates(
      _getLaneKeys(localToken, remoteChainSelector), new ConnectionPolicyStore.InlineConnectionConfigArgs[](0)
    );
  }
}
