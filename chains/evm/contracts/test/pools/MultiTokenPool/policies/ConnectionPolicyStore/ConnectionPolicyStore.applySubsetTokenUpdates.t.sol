// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../../../../interfaces/IPoolConfigManager.sol";

import {FinalityCodec} from "../../../../../libraries/FinalityCodec.sol";
import {PolicyConfig} from "../../../../../libraries/PolicyConfig.sol";
import {ConnectionPolicyStore} from "../../../../../pools/MultiTokenPool/policies/ConnectionPolicyStore.sol";
import {ConnectionPolicyStoreSetup} from "./ConnectionPolicyStoreSetup.t.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

contract ConnectionPolicyStore_applySubsetTokenUpdates is ConnectionPolicyStoreSetup {
  /// @dev CONNECTION_POLICY serves only its subset tokens on SOURCE_CHAIN_SELECTOR. The policy default record of
  /// REMOTE_POOL_POLICY covers that chain.
  function setUp() public override {
    super.setUp();
    _setConnectionPolicyConfig(_getConnectionPolicyConfigArgs(CONNECTION_POLICY, SOURCE_CHAIN_SELECTOR, true));
  }

  function test_applySubsetTokenUpdates() public {
    _setTokenConnectionPolicy(s_otherToken, CONNECTION_POLICY);
    ConnectionPolicyStore.SubsetTokenArgs[] memory tokensToAdd = new ConnectionPolicyStore.SubsetTokenArgs[](2);
    tokensToAdd[0] = _getSubsetTokenArgs(CONNECTION_POLICY, SOURCE_CHAIN_SELECTOR, s_token)[0];
    tokensToAdd[1] = _getSubsetTokenArgs(CONNECTION_POLICY, SOURCE_CHAIN_SELECTOR, s_otherToken)[0];

    vm.expectEmit();
    emit ConnectionPolicyStore.SubsetTokenAdded(CONNECTION_POLICY, SOURCE_CHAIN_SELECTOR, s_token);
    vm.expectEmit();
    emit ConnectionPolicyStore.SubsetTokenAdded(CONNECTION_POLICY, SOURCE_CHAIN_SELECTOR, s_otherToken);
    s_connectionStore.applySubsetTokenUpdates(new ConnectionPolicyStore.SubsetTokenArgs[](0), tokensToAdd);

    (bool onlySubsetTokens, address[] memory subsetTokens) =
      s_connectionStore.getSubsetTokens(CONNECTION_POLICY, SOURCE_CHAIN_SELECTOR);
    assertTrue(onlySubsetTokens);
    assertEq(2, subsetTokens.length);
    assertEq(s_token, subsetTokens[0]);
    assertEq(s_otherToken, subsetTokens[1]);
    // A subset token takes its remote token from the provider of the policy.
    IPoolConfigManager.OutboundLane memory lane =
      s_connectionStore.resolveOutboundLane(s_token, SOURCE_CHAIN_SELECTOR, FinalityCodec.WAIT_FOR_FINALITY_FLAG, "");
    assertEq(abi.encode(s_token), lane.remoteToken);
  }

  function test_applySubsetTokenUpdates_RemoveSubsetToken() public {
    _addSubsetToken(CONNECTION_POLICY, SOURCE_CHAIN_SELECTOR, s_token);

    vm.expectEmit();
    emit ConnectionPolicyStore.SubsetTokenRemoved(CONNECTION_POLICY, SOURCE_CHAIN_SELECTOR, s_token);
    s_connectionStore.applySubsetTokenUpdates(
      _getSubsetTokenArgs(CONNECTION_POLICY, SOURCE_CHAIN_SELECTOR, s_token),
      new ConnectionPolicyStore.SubsetTokenArgs[](0)
    );

    (, address[] memory subsetTokens) = s_connectionStore.getSubsetTokens(CONNECTION_POLICY, SOURCE_CHAIN_SELECTOR);
    assertEq(0, subsetTokens.length);
    assertFalse(s_connectionStore.hasLaneBatch(_getLaneKeys(s_token, SOURCE_CHAIN_SELECTOR))[0]);
  }

  function test_applySubsetTokenUpdates_RemovalsRunFirst() public {
    _addSubsetToken(CONNECTION_POLICY, SOURCE_CHAIN_SELECTOR, s_token);
    ConnectionPolicyStore.SubsetTokenArgs[] memory subsetTokenArgs =
      _getSubsetTokenArgs(CONNECTION_POLICY, SOURCE_CHAIN_SELECTOR, s_token);

    vm.expectEmit();
    emit ConnectionPolicyStore.SubsetTokenRemoved(CONNECTION_POLICY, SOURCE_CHAIN_SELECTOR, s_token);
    vm.expectEmit();
    emit ConnectionPolicyStore.SubsetTokenAdded(CONNECTION_POLICY, SOURCE_CHAIN_SELECTOR, s_token);
    s_connectionStore.applySubsetTokenUpdates(subsetTokenArgs, subsetTokenArgs);

    (, address[] memory subsetTokens) = s_connectionStore.getSubsetTokens(CONNECTION_POLICY, SOURCE_CHAIN_SELECTOR);
    assertEq(_arrayOf(s_token), subsetTokens);
  }

  function test_applySubsetTokenUpdates_TokenLeavesPolicy() public {
    _addSubsetToken(CONNECTION_POLICY, SOURCE_CHAIN_SELECTOR, s_token);

    // A subset token that leaves the policy stays in the set, but has no lane.
    _setTokenConnectionPolicy(s_token, bytes32(0));
    (, address[] memory subsetTokens) = s_connectionStore.getSubsetTokens(CONNECTION_POLICY, SOURCE_CHAIN_SELECTOR);
    assertEq(_arrayOf(s_token), subsetTokens);
    assertFalse(s_connectionStore.hasLaneBatch(_getLaneKeys(s_token, SOURCE_CHAIN_SELECTOR))[0]);

    // Its lane opens again at once when it joins the policy again.
    _setTokenConnectionPolicy(s_token, CONNECTION_POLICY);
    assertTrue(s_connectionStore.hasLaneBatch(_getLaneKeys(s_token, SOURCE_CHAIN_SELECTOR))[0]);
  }

  // Reverts

  function test_applySubsetTokenUpdates_RevertWhen_SubsetTokenNotFound() public {
    vm.expectRevert(
      abi.encodeWithSelector(
        ConnectionPolicyStore.SubsetTokenNotFound.selector, CONNECTION_POLICY, SOURCE_CHAIN_SELECTOR, s_token
      )
    );
    s_connectionStore.applySubsetTokenUpdates(
      _getSubsetTokenArgs(CONNECTION_POLICY, SOURCE_CHAIN_SELECTOR, s_token),
      new ConnectionPolicyStore.SubsetTokenArgs[](0)
    );
  }

  function test_applySubsetTokenUpdates_RevertWhen_PolicyRecordNotFound() public {
    // A connection policy has no policy default record, so chain selector zero never has a record.
    uint64[] memory remoteChainSelectors = new uint64[](2);
    remoteChainSelectors[0] = OTHER_CHAIN_SELECTOR;
    remoteChainSelectors[1] = PolicyConfig.DEFAULT_CHAIN_SELECTOR;

    for (uint256 i = 0; i < remoteChainSelectors.length; ++i) {
      vm.expectRevert(
        abi.encodeWithSelector(PolicyConfig.PolicyRecordNotFound.selector, CONNECTION_POLICY, remoteChainSelectors[i])
      );
      s_connectionStore.applySubsetTokenUpdates(
        new ConnectionPolicyStore.SubsetTokenArgs[](0),
        _getSubsetTokenArgs(CONNECTION_POLICY, remoteChainSelectors[i], s_token)
      );
    }
  }

  function test_applySubsetTokenUpdates_RevertWhen_OnlySubsetTokensNotSet() public {
    vm.expectRevert(
      abi.encodeWithSelector(
        ConnectionPolicyStore.OnlySubsetTokensNotSet.selector, CONNECTION_POLICY, DEST_CHAIN_SELECTOR
      )
    );
    s_connectionStore.applySubsetTokenUpdates(
      new ConnectionPolicyStore.SubsetTokenArgs[](0),
      _getSubsetTokenArgs(CONNECTION_POLICY, DEST_CHAIN_SELECTOR, s_token)
    );
  }

  function test_applySubsetTokenUpdates_RevertWhen_TokenNotInPolicy() public {
    vm.expectRevert(
      abi.encodeWithSelector(ConnectionPolicyStore.TokenNotInPolicy.selector, s_otherToken, CONNECTION_POLICY)
    );
    s_connectionStore.applySubsetTokenUpdates(
      new ConnectionPolicyStore.SubsetTokenArgs[](0),
      _getSubsetTokenArgs(CONNECTION_POLICY, SOURCE_CHAIN_SELECTOR, s_otherToken)
    );
  }

  function test_applySubsetTokenUpdates_RevertWhen_SubsetTokenAlreadyAdded() public {
    _addSubsetToken(CONNECTION_POLICY, SOURCE_CHAIN_SELECTOR, s_token);

    vm.expectRevert(
      abi.encodeWithSelector(
        ConnectionPolicyStore.SubsetTokenAlreadyAdded.selector, CONNECTION_POLICY, SOURCE_CHAIN_SELECTOR, s_token
      )
    );
    s_connectionStore.applySubsetTokenUpdates(
      new ConnectionPolicyStore.SubsetTokenArgs[](0),
      _getSubsetTokenArgs(CONNECTION_POLICY, SOURCE_CHAIN_SELECTOR, s_token)
    );
  }

  function test_applySubsetTokenUpdates_RevertWhen_UnauthorizedCaller() public {
    vm.startPrank(STRANGER);

    vm.expectRevert(abi.encodeWithSelector(AuthorizedCallers.UnauthorizedCaller.selector, STRANGER));
    s_connectionStore.applySubsetTokenUpdates(
      new ConnectionPolicyStore.SubsetTokenArgs[](0),
      _getSubsetTokenArgs(CONNECTION_POLICY, SOURCE_CHAIN_SELECTOR, s_token)
    );
  }
}
