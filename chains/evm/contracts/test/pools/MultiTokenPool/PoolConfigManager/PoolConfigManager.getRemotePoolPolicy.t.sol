// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {PoolConfigManager} from "../../../../pools/MultiTokenPool/PoolConfigManager.sol";
import {ConnectionPolicyStore} from "../../../../pools/MultiTokenPool/policies/ConnectionPolicyStore.sol";
import {PoolConfigManagerSetup} from "./PoolConfigManagerSetup.t.sol";

/// @notice Tokens A and B use the DEFAULT remote pool policy, which has a record for the destination chain.
/// @dev Compare raw return data because decoding nested records exceeds the stack limit without viaIR. The first
/// word contains the token count.
contract PoolConfigManager_getRemotePoolPolicy is PoolConfigManagerSetup {
  /// @notice Returns the token count and records for each policy.
  /// @dev Remote pool policies have no default records, so each record is for a specific chain.
  function test_getRemotePoolPolicy() public {
    bytes[] memory policyPoolAddresses = _getRemotePoolAddresses(abi.encode(makeAddr("policyPool")));
    _setRemotePoolPolicyConfig(DEFAULT_POLICY, DEST_CHAIN_SELECTOR, policyPoolAddresses);
    _setRemotePoolPolicyConfig(OTHER_POLICY, OTHER_CHAIN_SELECTOR, _getRemotePoolAddresses(abi.encode(s_remotePool)));

    ConnectionPolicyStore.RemotePoolPolicyConfigArgs[] memory expectedConfigs =
      new ConnectionPolicyStore.RemotePoolPolicyConfigArgs[](1);
    expectedConfigs[0] = _getRemotePoolPolicyConfigArgs(DEFAULT_POLICY, DEST_CHAIN_SELECTOR, policyPoolAddresses);
    assertEq(abi.encode(uint32(2), expectedConfigs), _getRemotePoolPolicyAnswer(DEFAULT_POLICY));

    expectedConfigs = new ConnectionPolicyStore.RemotePoolPolicyConfigArgs[](1);
    expectedConfigs[0] = _getRemotePoolPolicyConfigArgs(
      OTHER_POLICY, OTHER_CHAIN_SELECTOR, _getRemotePoolAddresses(abi.encode(s_remotePool))
    );
    assertEq(abi.encode(uint32(0), expectedConfigs), _getRemotePoolPolicyAnswer(OTHER_POLICY));
  }

  /// @dev Returns the raw response from a successful view call.
  function _getRemotePoolPolicyAnswer(
    bytes32 policyName
  ) internal view returns (bytes memory answer) {
    bool success;
    (success, answer) =
      address(s_poolConfigManager).staticcall(abi.encodeCall(PoolConfigManager.getRemotePoolPolicy, (policyName)));
    assertTrue(success);
    return answer;
  }

  function _getRemotePoolPolicyConfigArgs(
    bytes32 policyName,
    uint64 remoteChainSelector,
    bytes[] memory remotePoolAddresses
  ) internal pure returns (ConnectionPolicyStore.RemotePoolPolicyConfigArgs memory) {
    return ConnectionPolicyStore.RemotePoolPolicyConfigArgs({
      policyName: policyName, remoteChainSelector: remoteChainSelector, remotePoolAddresses: remotePoolAddresses
    });
  }
}
