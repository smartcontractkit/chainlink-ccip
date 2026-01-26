// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {CCIPPolicyEnginePayloads} from "../../../libraries/CCIPPolicyEnginePayloads.sol";
import {AdvancedPoolHooksSetup} from "./AdvancedPoolHooksSetup.t.sol";

contract AdvancedPoolHooks_policyEngineEncoding is AdvancedPoolHooksSetup {
  // bytes4(keccak256("PoolHookOutboundPolicyDataV1"))
  bytes4 internal constant POOL_HOOK_OUTBOUND_POLICY_DATA_V1_TAG = 0x12bebcb8;

  // bytes4(keccak256("PoolHookInboundPolicyDataV1"))
  bytes4 internal constant POOL_HOOK_INBOUND_POLICY_DATA_V1_TAG = 0x44d1de78;

  function test_PoolHookOutboundPolicyDataV1TagSelector() public pure {
    assertEq(POOL_HOOK_OUTBOUND_POLICY_DATA_V1_TAG, bytes4(keccak256("PoolHookOutboundPolicyDataV1")));
  }

  function test_PoolHookInboundPolicyDataV1TagSelector() public pure {
    assertEq(POOL_HOOK_INBOUND_POLICY_DATA_V1_TAG, bytes4(keccak256("PoolHookInboundPolicyDataV1")));
  }

  function test_PoolHookOutboundPolicyDataV1_RoundTrip() public pure {
    CCIPPolicyEnginePayloads.PoolHookOutboundPolicyDataV1 memory original = CCIPPolicyEnginePayloads.PoolHookOutboundPolicyDataV1({
      originalSender: address(0x456),
      blockConfirmationRequested: 5,
      remoteChainSelector: 1,
      receiver: abi.encode(address(0x123)),
      amount: 100e18,
      localToken: address(0x789),
      tokenArgs: abi.encode("test")
    });

    // Encode with tag
    bytes memory encoded = abi.encodeWithSelector(POOL_HOOK_OUTBOUND_POLICY_DATA_V1_TAG, original);

    // Verify tag
    assertEq(bytes4(encoded), POOL_HOOK_OUTBOUND_POLICY_DATA_V1_TAG);

    // Decode and verify
    bytes memory dataWithoutTag = _sliceBytes(encoded, 4);
    CCIPPolicyEnginePayloads.PoolHookOutboundPolicyDataV1 memory decoded =
      abi.decode(dataWithoutTag, (CCIPPolicyEnginePayloads.PoolHookOutboundPolicyDataV1));

    assertEq(decoded.receiver, original.receiver);
    assertEq(decoded.remoteChainSelector, original.remoteChainSelector);
    assertEq(decoded.originalSender, original.originalSender);
    assertEq(decoded.amount, original.amount);
    assertEq(decoded.localToken, original.localToken);
    assertEq(decoded.blockConfirmationRequested, original.blockConfirmationRequested);
    assertEq(decoded.tokenArgs, original.tokenArgs);
  }

  function test_PoolHookInboundPolicyDataV1_RoundTrip() public pure {
    CCIPPolicyEnginePayloads.PoolHookInboundPolicyDataV1 memory original = CCIPPolicyEnginePayloads.PoolHookInboundPolicyDataV1({
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

    // Encode with tag
    bytes memory encoded = abi.encodeWithSelector(POOL_HOOK_INBOUND_POLICY_DATA_V1_TAG, original);

    // Verify tag
    assertEq(bytes4(encoded), POOL_HOOK_INBOUND_POLICY_DATA_V1_TAG);

    // Decode and verify
    bytes memory dataWithoutTag = _sliceBytes(encoded, 4);
    CCIPPolicyEnginePayloads.PoolHookInboundPolicyDataV1 memory decoded =
      abi.decode(dataWithoutTag, (CCIPPolicyEnginePayloads.PoolHookInboundPolicyDataV1));

    assertEq(decoded.originalSender, original.originalSender);
    assertEq(decoded.remoteChainSelector, original.remoteChainSelector);
    assertEq(decoded.receiver, original.receiver);
    assertEq(decoded.amount, original.amount);
    assertEq(decoded.localToken, original.localToken);
    assertEq(decoded.sourcePoolAddress, original.sourcePoolAddress);
    assertEq(decoded.sourcePoolData, original.sourcePoolData);
    assertEq(decoded.offchainTokenData, original.offchainTokenData);
    assertEq(decoded.localAmount, original.localAmount);
    assertEq(decoded.blockConfirmationRequested, original.blockConfirmationRequested);
  }

  function _sliceBytes(bytes memory data, uint256 start) internal pure returns (bytes memory) {
    bytes memory result = new bytes(data.length - start);
    for (uint256 i = 0; i < result.length; ++i) {
      result[i] = data[start + i];
    }
    return result;
  }
}
