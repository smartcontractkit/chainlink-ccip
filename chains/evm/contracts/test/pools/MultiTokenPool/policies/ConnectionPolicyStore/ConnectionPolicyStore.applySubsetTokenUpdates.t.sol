// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../../../../interfaces/IPoolConfigManager.sol";

import {FinalityCodec} from "../../../../../libraries/FinalityCodec.sol";
import {PolicyConfig} from "../../../../../libraries/PolicyConfig.sol";
import {ConnectionPolicyStore} from "../../../../../pools/MultiTokenPool/policies/ConnectionPolicyStore.sol";
import {ConnectionPolicyStoreSetup} from "./ConnectionPolicyStoreSetup.t.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

contract ConnectionPolicyStore_applySubsetTokenUpdates is ConnectionPolicyStoreSetup {
  /// @dev Restrict SOURCE_CHAIN_SELECTOR to subset tokens and configure its remote pools.
  function setUp() public override {
    super.setUp();
    _setConnectionPolicyConfig(_getConnectionPolicyConfigArgs(CONNECTION_POLICY, SOURCE_CHAIN_SELECTOR, true));
    _setRemotePoolPolicyConfig(
      _getRemotePoolPolicyConfigArgs(REMOTE_POOL_POLICY, SOURCE_CHAIN_SELECTOR, _getRemotePoolAddresses(s_remotePool))
    );
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
      s_connectionStore.getSubsetTokens(CONNECTION_POLICY, SOURCE_CHAIN_SELECTOR, 0, type(uint64).max);
    assertTrue(onlySubsetTokens);
    assertEq(2, subsetTokens.length);
    assertEq(s_token, subsetTokens[0]);
    assertEq(s_otherToken, subsetTokens[1]);
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

    (, address[] memory subsetTokens) =
      s_connectionStore.getSubsetTokens(CONNECTION_POLICY, SOURCE_CHAIN_SELECTOR, 0, type(uint64).max);
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

    (, address[] memory subsetTokens) =
      s_connectionStore.getSubsetTokens(CONNECTION_POLICY, SOURCE_CHAIN_SELECTOR, 0, type(uint64).max);
    assertEq(_arrayOf(s_token), subsetTokens);
  }

  function test_applySubsetTokenUpdates_TokenLeavesPolicy() public {
    _addSubsetToken(CONNECTION_POLICY, SOURCE_CHAIN_SELECTOR, s_token);

    vm.expectRevert(abi.encodeWithSelector(ConnectionPolicyStore.TokenHasConfig.selector, s_token));
    _setTokenConnectionPolicy(s_token, bytes32(0));

    // Removing the subset entry allows the token to leave the policy.
    s_connectionStore.applySubsetTokenUpdates(
      _getSubsetTokenArgs(CONNECTION_POLICY, SOURCE_CHAIN_SELECTOR, s_token),
      new ConnectionPolicyStore.SubsetTokenArgs[](0)
    );
    _setTokenConnectionPolicy(s_token, bytes32(0));
    (, address[] memory subsetTokens) =
      s_connectionStore.getSubsetTokens(CONNECTION_POLICY, SOURCE_CHAIN_SELECTOR, 0, type(uint64).max);
    assertEq(0, subsetTokens.length);

    // Rejoining must not restore the old subset membership.
    _setTokenConnectionPolicy(s_token, CONNECTION_POLICY);
    assertFalse(s_connectionStore.hasLaneBatch(_getLaneKeys(s_token, SOURCE_CHAIN_SELECTOR))[0]);
  }

  /// @dev The subset can be filled while onlySubsetTokens is false, so switching to subset closes no listed lane.
  function test_applySubsetTokenUpdates_BeforeOnlySubsetTokensIsSet() public {
    _setTokenConnectionPolicy(s_otherToken, CONNECTION_POLICY);
    _setTokenRemotePoolPolicy(s_otherToken, REMOTE_POOL_POLICY);
    IPoolConfigManager.LaneKey[] memory lanes = new IPoolConfigManager.LaneKey[](2);
    lanes[0] = _getLaneKeys(s_token, DEST_CHAIN_SELECTOR)[0];
    lanes[1] = _getLaneKeys(s_otherToken, DEST_CHAIN_SELECTOR)[0];

    vm.expectEmit();
    emit ConnectionPolicyStore.SubsetTokenAdded(CONNECTION_POLICY, DEST_CHAIN_SELECTOR, s_token);
    s_connectionStore.applySubsetTokenUpdates(
      new ConnectionPolicyStore.SubsetTokenArgs[](0),
      _getSubsetTokenArgs(CONNECTION_POLICY, DEST_CHAIN_SELECTOR, s_token)
    );

    // The subset list has no effect until onlySubsetTokens is set.
    (bool onlySubsetTokens, address[] memory subsetTokens) =
      s_connectionStore.getSubsetTokens(CONNECTION_POLICY, DEST_CHAIN_SELECTOR, 0, type(uint64).max);
    assertFalse(onlySubsetTokens);
    assertEq(_arrayOf(s_token), subsetTokens);
    bool[] memory hasLane = s_connectionStore.hasLaneBatch(lanes);
    assertTrue(hasLane[0]);
    assertTrue(hasLane[1]);

    // Switching keeps the listed token's lane open and closes the unlisted one.
    _setConnectionPolicyConfig(_getConnectionPolicyConfigArgs(CONNECTION_POLICY, DEST_CHAIN_SELECTOR, true));
    hasLane = s_connectionStore.hasLaneBatch(lanes);
    assertTrue(hasLane[0]);
    assertFalse(hasLane[1]);
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
    // Connection policies have no default record at chain selector zero.
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
