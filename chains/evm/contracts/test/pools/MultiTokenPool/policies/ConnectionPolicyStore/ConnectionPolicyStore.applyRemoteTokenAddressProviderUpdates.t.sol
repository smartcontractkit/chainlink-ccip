// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../../../../interfaces/IPoolConfigManager.sol";
import {IRemoteTokenAddressProvider} from "../../../../../interfaces/IRemoteTokenAddressProvider.sol";

import {FinalityCodec} from "../../../../../libraries/FinalityCodec.sol";
import {PolicyConfig} from "../../../../../libraries/PolicyConfig.sol";
import {ConnectionPolicyStore} from "../../../../../pools/MultiTokenPool/policies/ConnectionPolicyStore.sol";
import {ConnectionPolicyStoreSetup} from "./ConnectionPolicyStoreSetup.t.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

contract ConnectionPolicyStore_applyRemoteTokenAddressProviderUpdates is ConnectionPolicyStoreSetup {
  /// @dev A provider needs code. The tests mock each call to it.
  address internal s_newProvider = makeAddr("newProvider");
  bytes internal s_newRemoteToken = abi.encode(makeAddr("newRemoteToken"));

  function setUp() public override {
    super.setUp();
    vm.etch(s_newProvider, hex"00");
    vm.mockCall(
      s_newProvider,
      abi.encodeCall(IRemoteTokenAddressProvider.getRemoteTokenAddress, (s_token, DEST_CHAIN_SELECTOR)),
      abi.encode(s_newRemoteToken)
    );
  }

  function test_applyRemoteTokenAddressProviderUpdates() public {
    vm.expectEmit();
    emit ConnectionPolicyStore.RemoteTokenAddressProviderSet(CONNECTION_POLICY, s_newProvider);
    s_connectionStore.applyRemoteTokenAddressProviderUpdates(
      _getRemoteTokenAddressProviderArgs(CONNECTION_POLICY, s_newProvider)
    );

    (, address remoteTokenAddressProvider,) = s_connectionStore.getConnectionPolicy(CONNECTION_POLICY);
    assertEq(s_newProvider, remoteTokenAddressProvider);
    // Both tokens resolve their remote address through the new provider.
    assertEq(s_newRemoteToken, _getLaneConnection(s_token, DEST_CHAIN_SELECTOR).remoteToken);
    IPoolConfigManager.OutboundLane memory lane =
      s_connectionStore.resolveOutboundLane(s_token, DEST_CHAIN_SELECTOR, FinalityCodec.WAIT_FOR_FINALITY_FLAG, "");
    assertEq(s_newRemoteToken, lane.remoteToken);
  }

  function test_applyRemoteTokenAddressProviderUpdates_ZeroProviderRemovesProvider() public {
    vm.expectEmit();
    emit ConnectionPolicyStore.RemoteTokenAddressProviderSet(CONNECTION_POLICY, address(0));
    s_connectionStore.applyRemoteTokenAddressProviderUpdates(
      _getRemoteTokenAddressProviderArgs(CONNECTION_POLICY, address(0))
    );

    (, address remoteTokenAddressProvider,) = s_connectionStore.getConnectionPolicy(CONNECTION_POLICY);
    assertEq(address(0), remoteTokenAddressProvider);
    // The connection still exists, but outbound transfers need a provider to resolve the remote token.
    IPoolConfigManager.LaneConnection memory connection = _getLaneConnection(s_token, DEST_CHAIN_SELECTOR);
    assertTrue(connection.hasConnection);
    assertEq(0, connection.remoteToken.length);
    vm.expectRevert(
      abi.encodeWithSelector(ConnectionPolicyStore.InvalidRemoteToken.selector, s_token, DEST_CHAIN_SELECTOR)
    );
    s_connectionStore.resolveOutboundLane(s_token, DEST_CHAIN_SELECTOR, FinalityCodec.WAIT_FOR_FINALITY_FLAG, "");
  }

  function test_applyRemoteTokenAddressProviderUpdates_AppliedInOrder() public {
    ConnectionPolicyStore.RemoteTokenAddressProviderArgs[] memory providers =
      new ConnectionPolicyStore.RemoteTokenAddressProviderArgs[](2);
    providers[0] = _getRemoteTokenAddressProviderArgs(CONNECTION_POLICY, s_newProvider)[0];
    providers[1] = _getRemoteTokenAddressProviderArgs(CONNECTION_POLICY, address(s_remoteTokenAddressProvider))[0];

    vm.expectEmit();
    emit ConnectionPolicyStore.RemoteTokenAddressProviderSet(CONNECTION_POLICY, s_newProvider);
    vm.expectEmit();
    emit ConnectionPolicyStore.RemoteTokenAddressProviderSet(CONNECTION_POLICY, address(s_remoteTokenAddressProvider));
    s_connectionStore.applyRemoteTokenAddressProviderUpdates(providers);

    (, address remoteTokenAddressProvider,) = s_connectionStore.getConnectionPolicy(CONNECTION_POLICY);
    assertEq(address(s_remoteTokenAddressProvider), remoteTokenAddressProvider);
    assertEq(abi.encode(s_token), _getLaneConnection(s_token, DEST_CHAIN_SELECTOR).remoteToken);
  }

  // Reverts

  function test_applyRemoteTokenAddressProviderUpdates_RevertWhen_ProviderHasNoCode() public {
    address providerWithoutCode = makeAddr("providerWithoutCode");

    vm.expectRevert(abi.encodeWithSelector(ConnectionPolicyStore.ProviderHasNoCode.selector, providerWithoutCode));
    s_connectionStore.applyRemoteTokenAddressProviderUpdates(
      _getRemoteTokenAddressProviderArgs(CONNECTION_POLICY, providerWithoutCode)
    );
  }

  function test_applyRemoteTokenAddressProviderUpdates_RevertWhen_PolicyNotFound() public {
    // REMOTE_POOL_POLICY exists only as a remote pool policy.
    vm.expectRevert(abi.encodeWithSelector(PolicyConfig.PolicyNotFound.selector, REMOTE_POOL_POLICY));
    s_connectionStore.applyRemoteTokenAddressProviderUpdates(
      _getRemoteTokenAddressProviderArgs(REMOTE_POOL_POLICY, s_newProvider)
    );
  }

  function test_applyRemoteTokenAddressProviderUpdates_RevertWhen_UnauthorizedCaller() public {
    vm.startPrank(STRANGER);

    vm.expectRevert(abi.encodeWithSelector(AuthorizedCallers.UnauthorizedCaller.selector, STRANGER));
    s_connectionStore.applyRemoteTokenAddressProviderUpdates(
      _getRemoteTokenAddressProviderArgs(CONNECTION_POLICY, s_newProvider)
    );
  }
}
