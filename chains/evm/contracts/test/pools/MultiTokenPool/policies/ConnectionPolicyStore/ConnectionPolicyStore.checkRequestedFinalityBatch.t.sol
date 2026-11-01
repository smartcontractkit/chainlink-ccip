// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../../../../interfaces/IPoolConfigManager.sol";

import {FinalityCodec} from "../../../../../libraries/FinalityCodec.sol";
import {PolicyConfig} from "../../../../../libraries/PolicyConfig.sol";
import {ConnectionPolicyStoreSetup} from "./ConnectionPolicyStoreSetup.t.sol";

contract ConnectionPolicyStore_checkRequestedFinalityBatch is ConnectionPolicyStoreSetup {
  function test_checkRequestedFinalityBatch() public view {
    // Default finality is always allowed, also for a token that is not registered.
    IPoolConfigManager.LaneKey[] memory lanes = new IPoolConfigManager.LaneKey[](3);
    lanes[0] = IPoolConfigManager.LaneKey({localToken: s_token, remoteChainSelector: DEST_CHAIN_SELECTOR});
    lanes[1] = IPoolConfigManager.LaneKey({localToken: s_otherToken, remoteChainSelector: DEST_CHAIN_SELECTOR});
    lanes[2] = IPoolConfigManager.LaneKey({localToken: STRANGER, remoteChainSelector: DEST_CHAIN_SELECTOR});

    s_connectionStore.checkRequestedFinalityBatch(lanes, FinalityCodec.WAIT_FOR_FINALITY_FLAG);
  }

  function test_checkRequestedFinalityBatch_FastFinality() public {
    _setAllowedFinality(s_token, ALLOWED_FINALITY);
    _setAllowedFinality(s_otherToken, ALLOWED_FINALITY);
    // s_otherToken has no lane, and its allowed finality still applies.
    IPoolConfigManager.LaneKey[] memory lanes = new IPoolConfigManager.LaneKey[](2);
    lanes[0] = IPoolConfigManager.LaneKey({localToken: s_token, remoteChainSelector: DEST_CHAIN_SELECTOR});
    lanes[1] = IPoolConfigManager.LaneKey({localToken: s_otherToken, remoteChainSelector: SOURCE_CHAIN_SELECTOR});

    s_connectionStore.checkRequestedFinalityBatch(lanes, REQUESTED_FINALITY);
  }

  // Reverts

  function test_checkRequestedFinalityBatch_RevertWhen_InvalidRequestedFinality() public {
    // s_token allows the request, and s_otherToken has no lane and allows default finality only. So the batch reverts.
    _setAllowedFinality(s_token, ALLOWED_FINALITY);
    IPoolConfigManager.LaneKey[] memory lanes = new IPoolConfigManager.LaneKey[](2);
    lanes[0] = IPoolConfigManager.LaneKey({localToken: s_token, remoteChainSelector: DEST_CHAIN_SELECTOR});
    lanes[1] = IPoolConfigManager.LaneKey({localToken: s_otherToken, remoteChainSelector: DEST_CHAIN_SELECTOR});

    vm.expectRevert(
      abi.encodeWithSelector(FinalityCodec.InvalidRequestedFinality.selector, REQUESTED_FINALITY, bytes4(0))
    );
    s_connectionStore.checkRequestedFinalityBatch(lanes, REQUESTED_FINALITY);
  }

  function test_checkRequestedFinalityBatch_RevertWhen_InvalidRequestedFinality_TokenNotRegistered() public {
    _setAllowedFinality(s_otherToken, ALLOWED_FINALITY);
    IPoolConfigManager.LaneKey[] memory lanes = _getLaneKeys(s_otherToken, DEST_CHAIN_SELECTOR);
    s_connectionStore.checkRequestedFinalityBatch(lanes, REQUESTED_FINALITY);

    // A deregistered token keeps its profile, but it allows default finality only.
    s_connectionStore.applyTokenProfileUpdates(_arrayOf(s_otherToken), new PolicyConfig.TokenProfileArgs[](0));

    vm.expectRevert(
      abi.encodeWithSelector(FinalityCodec.InvalidRequestedFinality.selector, REQUESTED_FINALITY, bytes4(0))
    );
    s_connectionStore.checkRequestedFinalityBatch(lanes, REQUESTED_FINALITY);
  }

  function test_checkRequestedFinalityBatch_RevertWhen_RequestedFinalityCanOnlyHaveOneMode() public {
    _setAllowedFinality(s_token, ALLOWED_FINALITY);
    bytes4 requestedFinality = FinalityCodec._encodeBlockDepthAndSafeFlag(10);

    vm.expectRevert(
      abi.encodeWithSelector(FinalityCodec.RequestedFinalityCanOnlyHaveOneMode.selector, requestedFinality)
    );
    s_connectionStore.checkRequestedFinalityBatch(_getLaneKeys(s_token, DEST_CHAIN_SELECTOR), requestedFinality);
  }

  function _setAllowedFinality(
    address localToken,
    bytes4 allowedFinalityConfig
  ) private {
    PolicyConfig.TokenProfileArgs[] memory profiles = _getTokenProfileArgs(localToken);
    profiles[0].allowedFinalityConfig = allowedFinalityConfig;
    s_connectionStore.applyTokenProfileUpdates(new address[](0), profiles);
  }
}
