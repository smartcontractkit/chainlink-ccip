// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {PoolConfigManager} from "../../../../pools/MultiTokenPool/PoolConfigManager.sol";
import {ConnectionPolicyStore} from "../../../../pools/MultiTokenPool/policies/ConnectionPolicyStore.sol";
import {PoolConfigManagerSetup} from "./PoolConfigManagerSetup.t.sol";

/// @notice Tokens A and B use the DEFAULT remote pool policy, which has a record on the destination chain. The tests
/// compare the whole answer of the view as bytes, and read the token count from its first word. The compiler without
/// via-IR runs out of stack when this test contract decodes RemotePoolPolicyConfigArgs[].
contract PoolConfigManager_getRemotePoolPolicy is PoolConfigManagerSetup {
  /// @notice Each policy returns its token count and its records in the shape of applyRemotePoolPolicyConfigUpdates.
  /// @dev The merged store has no remote pool policy default record, so every record is a chain record.
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

  /// @dev The answer of the view as bytes. The call must succeed.
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
