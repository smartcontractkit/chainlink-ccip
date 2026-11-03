// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../../../../interfaces/IPoolConfigManager.sol";

import {FinalityCodec} from "../../../../../libraries/FinalityCodec.sol";
import {PolicyConfig} from "../../../../../libraries/PolicyConfig.sol";
import {ConnectionPolicyStore} from "../../../../../pools/MultiTokenPool/policies/ConnectionPolicyStore.sol";
import {ConnectionPolicyStoreSetup} from "./ConnectionPolicyStoreSetup.t.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

contract ConnectionPolicyStore_applyTokenConnectionPolicyUpdates is ConnectionPolicyStoreSetup {
  /// @dev OTHER_POLICY enables SOURCE_CHAIN_SELECTOR for all assigned tokens.
  function setUp() public override {
    super.setUp();
    _setConnectionPolicyConfig(_getConnectionPolicyConfigArgs(OTHER_POLICY, SOURCE_CHAIN_SELECTOR, false));
    _setRemoteTokenAddressProvider(OTHER_POLICY, address(s_remoteTokenAddressProvider));
  }

  function test_applyTokenConnectionPolicyUpdates() public {
    PolicyConfig.TokenPolicyArgs[] memory tokenPolicies = new PolicyConfig.TokenPolicyArgs[](2);
    tokenPolicies[0] = PolicyConfig.TokenPolicyArgs({localToken: s_token, policyName: OTHER_POLICY});
    tokenPolicies[1] = PolicyConfig.TokenPolicyArgs({localToken: s_otherToken, policyName: CONNECTION_POLICY});

    vm.expectEmit();
    emit ConnectionPolicyStore.TokenConnectionPolicySet(s_token, CONNECTION_POLICY, OTHER_POLICY);
    vm.expectEmit();
    emit ConnectionPolicyStore.TokenConnectionPolicySet(s_otherToken, bytes32(0), CONNECTION_POLICY);
    s_connectionStore.applyTokenConnectionPolicyUpdates(tokenPolicies);

    (, bytes32 connectionPolicyName,) = s_connectionStore.getTokenPolicies(s_token);
    assertEq(OTHER_POLICY, connectionPolicyName);
    (, connectionPolicyName,) = s_connectionStore.getTokenPolicies(s_otherToken);
    assertEq(CONNECTION_POLICY, connectionPolicyName);
    (uint32 tokenCount,,) = s_connectionStore.getConnectionPolicy(CONNECTION_POLICY);
    assertEq(1, tokenCount);
    (tokenCount,,) = s_connectionStore.getConnectionPolicy(OTHER_POLICY);
    assertEq(1, tokenCount);
    bool[] memory hasLanes = s_connectionStore.hasLaneBatch(_getTokenLanes());
    assertFalse(hasLanes[0]);
    assertTrue(hasLanes[1]);
    assertTrue(hasLanes[2]);
    assertFalse(hasLanes[3]);
  }

  function test_applyTokenConnectionPolicyUpdates_NoLaneCheck() public {
    // Assigning a connection policy does not require a remote pool record.
    _setTokenConnectionPolicy(s_otherToken, CONNECTION_POLICY);

    assertTrue(s_connectionStore.hasLaneBatch(_getLaneKeys(s_otherToken, DEST_CHAIN_SELECTOR))[0]);
    vm.expectRevert(
      abi.encodeWithSelector(ConnectionPolicyStore.RemotePoolRecordNotFound.selector, s_otherToken, DEST_CHAIN_SELECTOR)
    );
    s_connectionStore.resolveOutboundLane(s_otherToken, DEST_CHAIN_SELECTOR, FinalityCodec.WAIT_FOR_FINALITY_FLAG, "");
  }

  function test_applyTokenConnectionPolicyUpdates_ZeroPolicyNameClearsTokenPolicy() public {
    vm.expectEmit();
    emit ConnectionPolicyStore.TokenConnectionPolicySet(s_token, CONNECTION_POLICY, bytes32(0));
    s_connectionStore.applyTokenConnectionPolicyUpdates(_getTokenPolicyArgs(s_token, bytes32(0)));

    (, bytes32 connectionPolicyName,) = s_connectionStore.getTokenPolicies(s_token);
    assertEq(bytes32(0), connectionPolicyName);
    (uint32 tokenCount,,) = s_connectionStore.getConnectionPolicy(CONNECTION_POLICY);
    assertEq(0, tokenCount);
    assertFalse(s_connectionStore.hasLaneBatch(_getLaneKeys(s_token, DEST_CHAIN_SELECTOR))[0]);
  }

  function test_applyTokenConnectionPolicyUpdates_SamePolicy() public {
    vm.expectEmit();
    emit ConnectionPolicyStore.TokenConnectionPolicySet(s_token, CONNECTION_POLICY, CONNECTION_POLICY);
    s_connectionStore.applyTokenConnectionPolicyUpdates(_getTokenPolicyArgs(s_token, CONNECTION_POLICY));

    (uint32 tokenCount,,) = s_connectionStore.getConnectionPolicy(CONNECTION_POLICY);
    assertEq(1, tokenCount);
  }

  function test_applyTokenConnectionPolicyUpdates_SameTokenTwice() public {
    // Sequential assignments for the same token must update each policy's token count correctly.
    PolicyConfig.TokenPolicyArgs[] memory tokenPolicies = new PolicyConfig.TokenPolicyArgs[](2);
    tokenPolicies[0] = PolicyConfig.TokenPolicyArgs({localToken: s_token, policyName: OTHER_POLICY});
    tokenPolicies[1] = PolicyConfig.TokenPolicyArgs({localToken: s_token, policyName: bytes32(0)});

    vm.expectEmit();
    emit ConnectionPolicyStore.TokenConnectionPolicySet(s_token, CONNECTION_POLICY, OTHER_POLICY);
    vm.expectEmit();
    emit ConnectionPolicyStore.TokenConnectionPolicySet(s_token, OTHER_POLICY, bytes32(0));
    s_connectionStore.applyTokenConnectionPolicyUpdates(tokenPolicies);

    (, bytes32 connectionPolicyName,) = s_connectionStore.getTokenPolicies(s_token);
    assertEq(bytes32(0), connectionPolicyName);
    (uint32 tokenCount,,) = s_connectionStore.getConnectionPolicy(CONNECTION_POLICY);
    assertEq(0, tokenCount);
    (tokenCount,,) = s_connectionStore.getConnectionPolicy(OTHER_POLICY);
    assertEq(0, tokenCount);
  }

  function test_applyTokenConnectionPolicyUpdates_InlineRecordStays() public {
    bytes memory inlineRemoteToken = abi.encode(makeAddr("inlineRemoteToken"));
    _setInlineConnectionConfig(_getInlineConnectionConfigArgs(s_token, SOURCE_CHAIN_SELECTOR, inlineRemoteToken));

    _setTokenConnectionPolicy(s_token, OTHER_POLICY);

    assertEq(1, s_connectionStore.getInlineConnectionConfigs(s_token).length);
    IPoolConfigManager.LaneConnection[] memory connections =
      s_connectionStore.getLaneConnectionBatch(_getLaneKeys(s_token, SOURCE_CHAIN_SELECTOR));
    assertEq(inlineRemoteToken, connections[0].remoteToken);
  }

  // Reverts

  function test_applyTokenConnectionPolicyUpdates_RevertWhen_TokenHasConfig() public {
    _setConnectionPolicyConfig(_getConnectionPolicyConfigArgs(CONNECTION_POLICY, SOURCE_CHAIN_SELECTOR, true));
    _addSubsetToken(CONNECTION_POLICY, SOURCE_CHAIN_SELECTOR, s_token);
    bytes32[2] memory newPolicyNames = [bytes32(0), OTHER_POLICY];

    for (uint256 i = 0; i < newPolicyNames.length; ++i) {
      vm.expectRevert(abi.encodeWithSelector(ConnectionPolicyStore.TokenHasConfig.selector, s_token));
      _setTokenConnectionPolicy(s_token, newPolicyNames[i]);
    }
    // Assigning the same policy is allowed because the token does not leave it.
    _setTokenConnectionPolicy(s_token, CONNECTION_POLICY);
  }

  function test_applyTokenConnectionPolicyUpdates_RevertWhen_LocalTokenNotRegistered() public {
    address unknownToken = makeAddr("unknownToken");

    vm.expectRevert(abi.encodeWithSelector(PolicyConfig.LocalTokenNotRegistered.selector, unknownToken));
    s_connectionStore.applyTokenConnectionPolicyUpdates(_getTokenPolicyArgs(unknownToken, CONNECTION_POLICY));
  }

  function test_applyTokenConnectionPolicyUpdates_RevertWhen_PolicyNotFound() public {
    // REMOTE_POOL_POLICY exists only as a remote pool policy.
    vm.expectRevert(abi.encodeWithSelector(PolicyConfig.PolicyNotFound.selector, REMOTE_POOL_POLICY));
    s_connectionStore.applyTokenConnectionPolicyUpdates(_getTokenPolicyArgs(s_token, REMOTE_POOL_POLICY));
  }

  function test_applyTokenConnectionPolicyUpdates_RevertWhen_UnauthorizedCaller() public {
    vm.startPrank(STRANGER);

    vm.expectRevert(abi.encodeWithSelector(AuthorizedCallers.UnauthorizedCaller.selector, STRANGER));
    s_connectionStore.applyTokenConnectionPolicyUpdates(_getTokenPolicyArgs(s_token, OTHER_POLICY));
  }

  /// @dev s_token and s_otherToken, each on DEST_CHAIN_SELECTOR and SOURCE_CHAIN_SELECTOR.
  function _getTokenLanes() private view returns (IPoolConfigManager.LaneKey[] memory lanes) {
    lanes = new IPoolConfigManager.LaneKey[](4);
    lanes[0] = IPoolConfigManager.LaneKey({localToken: s_token, remoteChainSelector: DEST_CHAIN_SELECTOR});
    lanes[1] = IPoolConfigManager.LaneKey({localToken: s_token, remoteChainSelector: SOURCE_CHAIN_SELECTOR});
    lanes[2] = IPoolConfigManager.LaneKey({localToken: s_otherToken, remoteChainSelector: DEST_CHAIN_SELECTOR});
    lanes[3] = IPoolConfigManager.LaneKey({localToken: s_otherToken, remoteChainSelector: SOURCE_CHAIN_SELECTOR});
    return lanes;
  }
}
