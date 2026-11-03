// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {FinalityCodec} from "../../../../../libraries/FinalityCodec.sol";
import {PolicyConfig} from "../../../../../libraries/PolicyConfig.sol";
import {ConnectionPolicyStore} from "../../../../../pools/MultiTokenPool/policies/ConnectionPolicyStore.sol";
import {ConnectionPolicyStoreSetup} from "./ConnectionPolicyStoreSetup.t.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

contract ConnectionPolicyStore_applyRemotePoolPolicyConfigUpdates is ConnectionPolicyStoreSetup {
  function test_applyRemotePoolPolicyConfigUpdates() public {
    bytes[] memory chainRemotePoolAddresses = new bytes[](2);
    chainRemotePoolAddresses[0] = s_remotePool;
    chainRemotePoolAddresses[1] = s_otherRemotePool;
    ConnectionPolicyStore.RemotePoolPolicyConfigArgs[] memory policyConfigs =
      new ConnectionPolicyStore.RemotePoolPolicyConfigArgs[](2);
    policyConfigs[0] = _getRemotePoolPolicyConfigArgs(OTHER_POLICY, DEST_CHAIN_SELECTOR, chainRemotePoolAddresses)[0];
    policyConfigs[1] =
      _getRemotePoolPolicyConfigArgs(OTHER_POLICY, SOURCE_CHAIN_SELECTOR, _getRemotePoolAddresses(s_remotePool))[0];

    vm.expectEmit();
    emit ConnectionPolicyStore.RemotePoolPolicyConfigSet(OTHER_POLICY, DEST_CHAIN_SELECTOR, chainRemotePoolAddresses);
    vm.expectEmit();
    emit ConnectionPolicyStore.RemotePoolPolicyConfigSet(
      OTHER_POLICY, SOURCE_CHAIN_SELECTOR, _getRemotePoolAddresses(s_remotePool)
    );
    s_connectionStore.applyRemotePoolPolicyConfigUpdates(
      new PolicyConfig.PolicyRecordKey[](0), new bytes32[](0), policyConfigs
    );

    // The first record set under a new name creates the policy.
    _assertRemotePoolPolicy(OTHER_POLICY, 0, policyConfigs);
    (, bytes32[] memory remotePoolPolicyNames) = s_connectionStore.getPolicyNames();
    assertEq(2, remotePoolPolicyNames.length);
    assertEq(REMOTE_POOL_POLICY, remotePoolPolicyNames[0]);
    assertEq(OTHER_POLICY, remotePoolPolicyNames[1]);
  }

  function test_applyRemotePoolPolicyConfigUpdates_ReplaceRecord() public {
    ConnectionPolicyStore.RemotePoolPolicyConfigArgs[] memory policyConfigs = _getRemotePoolPolicyConfigArgs(
      REMOTE_POOL_POLICY, DEST_CHAIN_SELECTOR, _getRemotePoolAddresses(s_otherRemotePool)
    );
    bytes memory sourcePoolData = _getSourcePoolData(DEFAULT_TOKEN_DECIMALS, s_token, "");

    vm.expectEmit();
    emit ConnectionPolicyStore.RemotePoolPolicyConfigSet(
      REMOTE_POOL_POLICY, DEST_CHAIN_SELECTOR, _getRemotePoolAddresses(s_otherRemotePool)
    );
    s_connectionStore.applyRemotePoolPolicyConfigUpdates(
      new PolicyConfig.PolicyRecordKey[](0), new bytes32[](0), policyConfigs
    );

    _assertRemotePoolPolicy(REMOTE_POOL_POLICY, 1, policyConfigs);
    // The change applies at once to every token of the policy, and the old remote pool is no longer trusted.
    s_connectionStore.resolveInboundLane(
      s_token, DEST_CHAIN_SELECTOR, s_otherRemotePool, sourcePoolData, 1e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG
    );
    vm.expectRevert(
      abi.encodeWithSelector(
        ConnectionPolicyStore.InvalidSourcePoolAddress.selector, s_token, DEST_CHAIN_SELECTOR, s_remotePool
      )
    );
    s_connectionStore.resolveInboundLane(
      s_token, DEST_CHAIN_SELECTOR, s_remotePool, sourcePoolData, 1e18, FinalityCodec.WAIT_FOR_FINALITY_FLAG
    );
  }

  function test_applyRemotePoolPolicyConfigUpdates_RemoveRecord() public {
    _setRemotePoolPolicyConfig(
      _getRemotePoolPolicyConfigArgs(REMOTE_POOL_POLICY, SOURCE_CHAIN_SELECTOR, _getRemotePoolAddresses(s_remotePool))
    );
    PolicyConfig.PolicyRecordKey[] memory recordsToRemove =
      _getPolicyRecordKeys(REMOTE_POOL_POLICY, DEST_CHAIN_SELECTOR);

    vm.expectEmit();
    emit ConnectionPolicyStore.RemotePoolPolicyConfigRemoved(REMOTE_POOL_POLICY, DEST_CHAIN_SELECTOR);
    s_connectionStore.applyRemotePoolPolicyConfigUpdates(
      recordsToRemove, new bytes32[](0), new ConnectionPolicyStore.RemotePoolPolicyConfigArgs[](0)
    );

    _assertRemotePoolPolicy(
      REMOTE_POOL_POLICY,
      1,
      _getRemotePoolPolicyConfigArgs(REMOTE_POOL_POLICY, SOURCE_CHAIN_SELECTOR, _getRemotePoolAddresses(s_remotePool))
    );
    // The record of another chain does not serve the chain, so the token has no remote pool record there at once.
    vm.expectRevert(
      abi.encodeWithSelector(ConnectionPolicyStore.RemotePoolRecordNotFound.selector, s_token, DEST_CHAIN_SELECTOR)
    );
    s_connectionStore.checkLaneRecords(s_token, DEST_CHAIN_SELECTOR);
  }

  function test_applyRemotePoolPolicyConfigUpdates_RemoveLastRecord() public {
    PolicyConfig.PolicyRecordKey[] memory recordsToRemove =
      _getPolicyRecordKeys(REMOTE_POOL_POLICY, DEST_CHAIN_SELECTOR);

    vm.expectEmit();
    emit ConnectionPolicyStore.RemotePoolPolicyConfigRemoved(REMOTE_POOL_POLICY, DEST_CHAIN_SELECTOR);
    s_connectionStore.applyRemotePoolPolicyConfigUpdates(
      recordsToRemove, new bytes32[](0), new ConnectionPolicyStore.RemotePoolPolicyConfigArgs[](0)
    );

    // The policy stays with its token and no records, so the lane of the token has no remote pool record.
    _assertRemotePoolPolicy(REMOTE_POOL_POLICY, 1, new ConnectionPolicyStore.RemotePoolPolicyConfigArgs[](0));
    vm.expectRevert(
      abi.encodeWithSelector(ConnectionPolicyStore.RemotePoolRecordNotFound.selector, s_token, DEST_CHAIN_SELECTOR)
    );
    s_connectionStore.checkLaneRecords(s_token, DEST_CHAIN_SELECTOR);
  }

  function test_applyRemotePoolPolicyConfigUpdates_RemovePolicy() public {
    _setTokenRemotePoolPolicy(s_token, bytes32(0));
    // The record removals run before the policy removals, so one call can remove the records and the policy.
    PolicyConfig.PolicyRecordKey[] memory recordsToRemove =
      _getPolicyRecordKeys(REMOTE_POOL_POLICY, DEST_CHAIN_SELECTOR);
    bytes32[] memory policiesToRemove = _getPolicyNames(REMOTE_POOL_POLICY);

    vm.expectEmit();
    emit ConnectionPolicyStore.RemotePoolPolicyConfigRemoved(REMOTE_POOL_POLICY, DEST_CHAIN_SELECTOR);
    vm.expectEmit();
    emit ConnectionPolicyStore.RemotePoolPolicyRemoved(REMOTE_POOL_POLICY);
    s_connectionStore.applyRemotePoolPolicyConfigUpdates(
      recordsToRemove, policiesToRemove, new ConnectionPolicyStore.RemotePoolPolicyConfigArgs[](0)
    );

    (bytes32[] memory connectionPolicyNames, bytes32[] memory remotePoolPolicyNames) =
      s_connectionStore.getPolicyNames();
    assertEq(1, connectionPolicyNames.length);
    assertEq(0, remotePoolPolicyNames.length);
    _assertRemotePoolPolicy(REMOTE_POOL_POLICY, 0, new ConnectionPolicyStore.RemotePoolPolicyConfigArgs[](0));
  }

  function test_applyRemotePoolPolicyConfigUpdates_RemovalsRunFirst() public {
    _setTokenRemotePoolPolicy(s_token, bytes32(0));
    PolicyConfig.PolicyRecordKey[] memory recordsToRemove =
      _getPolicyRecordKeys(REMOTE_POOL_POLICY, DEST_CHAIN_SELECTOR);
    bytes32[] memory policiesToRemove = _getPolicyNames(REMOTE_POOL_POLICY);
    ConnectionPolicyStore.RemotePoolPolicyConfigArgs[] memory policyConfigs = _getRemotePoolPolicyConfigArgs(
      REMOTE_POOL_POLICY, DEST_CHAIN_SELECTOR, _getRemotePoolAddresses(s_otherRemotePool)
    );

    s_connectionStore.applyRemotePoolPolicyConfigUpdates(recordsToRemove, policiesToRemove, policyConfigs);

    // The set runs last, so it creates the policy again with only the new remote pool.
    _assertRemotePoolPolicy(REMOTE_POOL_POLICY, 0, policyConfigs);
  }

  // Reverts

  function test_applyRemotePoolPolicyConfigUpdates_RevertWhen_PolicyRecordHasNoPools() public {
    ConnectionPolicyStore.RemotePoolPolicyConfigArgs[] memory policyConfigs =
      _getRemotePoolPolicyConfigArgs(REMOTE_POOL_POLICY, DEST_CHAIN_SELECTOR, new bytes[](0));

    vm.expectRevert(
      abi.encodeWithSelector(
        ConnectionPolicyStore.PolicyRecordHasNoPools.selector, REMOTE_POOL_POLICY, DEST_CHAIN_SELECTOR
      )
    );
    s_connectionStore.applyRemotePoolPolicyConfigUpdates(
      new PolicyConfig.PolicyRecordKey[](0), new bytes32[](0), policyConfigs
    );
  }

  function test_applyRemotePoolPolicyConfigUpdates_RevertWhen_ZeroChainSelectorNotAllowed() public {
    // A remote pool policy has no policy default record.
    vm.expectRevert(PolicyConfig.ZeroChainSelectorNotAllowed.selector);
    _setRemotePoolPolicyConfig(
      _getRemotePoolPolicyConfigArgs(
        REMOTE_POOL_POLICY, PolicyConfig.DEFAULT_CHAIN_SELECTOR, _getRemotePoolAddresses(s_remotePool)
      )
    );
  }

  function test_applyRemotePoolPolicyConfigUpdates_RevertWhen_ZeroAddressInvalid() public {
    // An empty first remote pool reverts, and so does an empty later remote pool.
    bytes[] memory remotePoolAddresses = new bytes[](2);
    remotePoolAddresses[1] = s_remotePool;
    vm.expectRevert(ConnectionPolicyStore.ZeroAddressInvalid.selector);
    _setRemotePoolPolicyConfig(
      _getRemotePoolPolicyConfigArgs(REMOTE_POOL_POLICY, DEST_CHAIN_SELECTOR, remotePoolAddresses)
    );

    remotePoolAddresses[0] = s_remotePool;
    remotePoolAddresses[1] = "";
    vm.expectRevert(ConnectionPolicyStore.ZeroAddressInvalid.selector);
    _setRemotePoolPolicyConfig(
      _getRemotePoolPolicyConfigArgs(REMOTE_POOL_POLICY, DEST_CHAIN_SELECTOR, remotePoolAddresses)
    );
  }

  function test_applyRemotePoolPolicyConfigUpdates_RevertWhen_RemotePoolAlreadyAdded() public {
    bytes[] memory remotePoolAddresses = new bytes[](3);
    remotePoolAddresses[0] = s_otherRemotePool;
    remotePoolAddresses[1] = s_remotePool;
    remotePoolAddresses[2] = s_otherRemotePool;
    ConnectionPolicyStore.RemotePoolPolicyConfigArgs[] memory policyConfigs =
      _getRemotePoolPolicyConfigArgs(REMOTE_POOL_POLICY, DEST_CHAIN_SELECTOR, remotePoolAddresses);

    vm.expectRevert(
      abi.encodeWithSelector(
        ConnectionPolicyStore.RemotePoolAlreadyAdded.selector, DEST_CHAIN_SELECTOR, s_otherRemotePool
      )
    );
    s_connectionStore.applyRemotePoolPolicyConfigUpdates(
      new PolicyConfig.PolicyRecordKey[](0), new bytes32[](0), policyConfigs
    );
  }

  function test_applyRemotePoolPolicyConfigUpdates_RevertWhen_ZeroPolicyNameNotAllowed() public {
    ConnectionPolicyStore.RemotePoolPolicyConfigArgs[] memory policyConfigs =
      _getRemotePoolPolicyConfigArgs(bytes32(0), DEST_CHAIN_SELECTOR, _getRemotePoolAddresses(s_remotePool));

    vm.expectRevert(PolicyConfig.ZeroPolicyNameNotAllowed.selector);
    s_connectionStore.applyRemotePoolPolicyConfigUpdates(
      new PolicyConfig.PolicyRecordKey[](0), new bytes32[](0), policyConfigs
    );
  }

  function test_applyRemotePoolPolicyConfigUpdates_RevertWhen_PolicyRecordNotFound() public {
    PolicyConfig.PolicyRecordKey[] memory recordsToRemove =
      _getPolicyRecordKeys(REMOTE_POOL_POLICY, SOURCE_CHAIN_SELECTOR);

    vm.expectRevert(
      abi.encodeWithSelector(PolicyConfig.PolicyRecordNotFound.selector, REMOTE_POOL_POLICY, SOURCE_CHAIN_SELECTOR)
    );
    s_connectionStore.applyRemotePoolPolicyConfigUpdates(
      recordsToRemove, new bytes32[](0), new ConnectionPolicyStore.RemotePoolPolicyConfigArgs[](0)
    );
  }

  function test_applyRemotePoolPolicyConfigUpdates_RevertWhen_PolicyNotFound() public {
    // CONNECTION_POLICY exists only as a connection policy.
    bytes32[] memory policiesToRemove = _getPolicyNames(CONNECTION_POLICY);

    vm.expectRevert(abi.encodeWithSelector(PolicyConfig.PolicyNotFound.selector, CONNECTION_POLICY));
    s_connectionStore.applyRemotePoolPolicyConfigUpdates(
      new PolicyConfig.PolicyRecordKey[](0), policiesToRemove, new ConnectionPolicyStore.RemotePoolPolicyConfigArgs[](0)
    );
  }

  function test_applyRemotePoolPolicyConfigUpdates_RevertWhen_PolicyInUse() public {
    bytes32[] memory policiesToRemove = _getPolicyNames(REMOTE_POOL_POLICY);

    vm.expectRevert(abi.encodeWithSelector(PolicyConfig.PolicyInUse.selector, REMOTE_POOL_POLICY));
    s_connectionStore.applyRemotePoolPolicyConfigUpdates(
      new PolicyConfig.PolicyRecordKey[](0), policiesToRemove, new ConnectionPolicyStore.RemotePoolPolicyConfigArgs[](0)
    );
  }

  function test_applyRemotePoolPolicyConfigUpdates_RevertWhen_PolicyHasRecords() public {
    _setTokenRemotePoolPolicy(s_token, bytes32(0));
    bytes32[] memory policiesToRemove = _getPolicyNames(REMOTE_POOL_POLICY);

    vm.expectRevert(abi.encodeWithSelector(PolicyConfig.PolicyHasRecords.selector, REMOTE_POOL_POLICY));
    s_connectionStore.applyRemotePoolPolicyConfigUpdates(
      new PolicyConfig.PolicyRecordKey[](0), policiesToRemove, new ConnectionPolicyStore.RemotePoolPolicyConfigArgs[](0)
    );
  }

  function test_applyRemotePoolPolicyConfigUpdates_RevertWhen_UnauthorizedCaller() public {
    ConnectionPolicyStore.RemotePoolPolicyConfigArgs[] memory policyConfigs =
      _getRemotePoolPolicyConfigArgs(OTHER_POLICY, DEST_CHAIN_SELECTOR, _getRemotePoolAddresses(s_remotePool));
    vm.startPrank(STRANGER);

    vm.expectRevert(abi.encodeWithSelector(AuthorizedCallers.UnauthorizedCaller.selector, STRANGER));
    s_connectionStore.applyRemotePoolPolicyConfigUpdates(
      new PolicyConfig.PolicyRecordKey[](0), new bytes32[](0), policyConfigs
    );
  }
}
