// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {CCIPPolicyEnginePayloads} from "../../../libraries/CCIPPolicyEnginePayloads.sol";
import {AdvancedPoolHooksSetup} from "./AdvancedPoolHooksSetup.t.sol";

contract AdvancedPoolHooks_policyEngineEncoding is AdvancedPoolHooksSetup {
  function test_PoolHookOutboundPolicyDataV1TagSelector() public pure {
    assertEq(
      bytes4(keccak256("PoolHookOutboundPolicyDataV1")), CCIPPolicyEnginePayloads.POOL_HOOK_OUTBOUND_POLICY_DATA_V1_TAG
    );
  }

  function test_PoolHookInboundPolicyDataV1TagSelector() public pure {
    assertEq(
      bytes4(keccak256("PoolHookInboundPolicyDataV1")), CCIPPolicyEnginePayloads.POOL_HOOK_INBOUND_POLICY_DATA_V1_TAG
    );
  }

  function test_PoolHookOutboundPolicyDataV1_RoundTrip() public pure {
    CCIPPolicyEnginePayloads.PoolHookOutboundPolicyDataV1 memory original = CCIPPolicyEnginePayloads
      .PoolHookOutboundPolicyDataV1({
        originalSender: address(0x456),
        blockConfirmationRequested: 5,
        remoteChainSelector: 1,
        receiver: abi.encode(address(0x123)),
        amount: 100e18,
        localToken: address(0x789),
        tokenArgs: abi.encode("test")
      });

    bytes memory encoded =
      abi.encodeWithSelector(CCIPPolicyEnginePayloads.POOL_HOOK_OUTBOUND_POLICY_DATA_V1_TAG, original);

    assertEq(CCIPPolicyEnginePayloads.POOL_HOOK_OUTBOUND_POLICY_DATA_V1_TAG, bytes4(encoded));

    CCIPPolicyEnginePayloads.PoolHookOutboundPolicyDataV1 memory decoded =
      abi.decode(_sliceBytes(encoded, 4), (CCIPPolicyEnginePayloads.PoolHookOutboundPolicyDataV1));

    assertEq(original.originalSender, decoded.originalSender);
    assertEq(original.blockConfirmationRequested, decoded.blockConfirmationRequested);
    assertEq(original.remoteChainSelector, decoded.remoteChainSelector);
    assertEq(original.receiver, decoded.receiver);
    assertEq(original.amount, decoded.amount);
    assertEq(original.localToken, decoded.localToken);
    assertEq(original.tokenArgs, decoded.tokenArgs);
  }

  function test_PoolHookInboundPolicyDataV1_RoundTrip() public pure {
    CCIPPolicyEnginePayloads.PoolHookInboundPolicyDataV1 memory original = CCIPPolicyEnginePayloads
      .PoolHookInboundPolicyDataV1({
        originalSender: abi.encode(address(0x123)),
        blockConfirmationRequested: 10,
        remoteChainSelector: 1,
        receiver: address(0x456),
        amount: 100e18,
        localToken: address(0x789),
        sourcePoolAddress: abi.encode(address(0xabc)),
        sourcePoolData: abi.encode("poolData"),
        offchainTokenData: abi.encode("offchain"),
        localAmount: 99e18
      });

    bytes memory encoded =
      abi.encodeWithSelector(CCIPPolicyEnginePayloads.POOL_HOOK_INBOUND_POLICY_DATA_V1_TAG, original);

    assertEq(CCIPPolicyEnginePayloads.POOL_HOOK_INBOUND_POLICY_DATA_V1_TAG, bytes4(encoded));

    CCIPPolicyEnginePayloads.PoolHookInboundPolicyDataV1 memory decoded =
      abi.decode(_sliceBytes(encoded, 4), (CCIPPolicyEnginePayloads.PoolHookInboundPolicyDataV1));

    assertEq(original.originalSender, decoded.originalSender);
    assertEq(original.blockConfirmationRequested, decoded.blockConfirmationRequested);
    assertEq(original.remoteChainSelector, decoded.remoteChainSelector);
    assertEq(original.receiver, decoded.receiver);
    assertEq(original.amount, decoded.amount);
    assertEq(original.localToken, decoded.localToken);
    assertEq(original.sourcePoolAddress, decoded.sourcePoolAddress);
    assertEq(original.sourcePoolData, decoded.sourcePoolData);
    assertEq(original.offchainTokenData, decoded.offchainTokenData);
    assertEq(original.localAmount, decoded.localAmount);
  }
}
