// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {CCIPPolicyEnginePayloads} from "../../../libraries/CCIPPolicyEnginePayloads.sol";
import {AdvancedPoolHooksSetup} from "./AdvancedPoolHooksSetup.t.sol";

contract AdvancedPoolHooks_policyEngineEncoding is AdvancedPoolHooksSetup {
  // bytes4(keccak256("OutboundPolicyDataV1"))
  bytes4 internal constant OUTBOUND_POLICY_DATA_V1_TAG = 0x73bb902c;

  // bytes4(keccak256("InboundPolicyDataV1"))
  bytes4 internal constant INBOUND_POLICY_DATA_V1_TAG = 0xe8deab79;

  function test_OutboundPolicyDataV1TagSelector() public pure {
    assertEq(OUTBOUND_POLICY_DATA_V1_TAG, bytes4(keccak256("OutboundPolicyDataV1")));
  }

  function test_InboundPolicyDataV1TagSelector() public pure {
    assertEq(INBOUND_POLICY_DATA_V1_TAG, bytes4(keccak256("InboundPolicyDataV1")));
  }

  function test_OutboundPolicyDataV1_RoundTrip() public pure {
    CCIPPolicyEnginePayloads.OutboundPolicyDataV1 memory original = CCIPPolicyEnginePayloads.OutboundPolicyDataV1({
      receiver: abi.encode(address(0x123)),
      remoteChainSelector: 1,
      originalSender: address(0x456),
      amount: 100e18,
      localToken: address(0x789),
      blockConfirmationRequested: 5,
      tokenArgs: abi.encode("test")
    });

    // Encode with tag
    bytes memory encoded = abi.encodeWithSelector(OUTBOUND_POLICY_DATA_V1_TAG, original);

    // Verify tag
    assertEq(bytes4(encoded), OUTBOUND_POLICY_DATA_V1_TAG);

    // Decode and verify
    bytes memory dataWithoutTag = _sliceBytes(encoded, 4);
    CCIPPolicyEnginePayloads.OutboundPolicyDataV1 memory decoded =
      abi.decode(dataWithoutTag, (CCIPPolicyEnginePayloads.OutboundPolicyDataV1));

    assertEq(decoded.receiver, original.receiver);
    assertEq(decoded.remoteChainSelector, original.remoteChainSelector);
    assertEq(decoded.originalSender, original.originalSender);
    assertEq(decoded.amount, original.amount);
    assertEq(decoded.localToken, original.localToken);
    assertEq(decoded.blockConfirmationRequested, original.blockConfirmationRequested);
    assertEq(decoded.tokenArgs, original.tokenArgs);
  }

  function test_InboundPolicyDataV1_RoundTrip() public pure {
    CCIPPolicyEnginePayloads.InboundPolicyDataV1 memory original = CCIPPolicyEnginePayloads.InboundPolicyDataV1({
      originalSender: abi.encode(address(0x123)),
      remoteChainSelector: 1,
      receiver: address(0x456),
      amount: 100e18,
      localToken: address(0x789),
      sourcePoolAddress: abi.encode(address(0xabc)),
      sourcePoolData: abi.encode("poolData"),
      offchainTokenData: abi.encode("offchain"),
      localAmount: 99e18,
      blockConfirmationRequested: 10
    });

    // Encode with tag
    bytes memory encoded = abi.encodeWithSelector(INBOUND_POLICY_DATA_V1_TAG, original);

    // Verify tag
    assertEq(bytes4(encoded), INBOUND_POLICY_DATA_V1_TAG);

    // Decode and verify
    bytes memory dataWithoutTag = _sliceBytes(encoded, 4);
    CCIPPolicyEnginePayloads.InboundPolicyDataV1 memory decoded =
      abi.decode(dataWithoutTag, (CCIPPolicyEnginePayloads.InboundPolicyDataV1));

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
