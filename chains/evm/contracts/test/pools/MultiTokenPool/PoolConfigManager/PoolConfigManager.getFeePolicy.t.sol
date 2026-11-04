// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {PolicyConfig} from "../../../../libraries/PolicyConfig.sol";
import {FeePolicyStore} from "../../../../pools/MultiTokenPool/policies/FeePolicyStore.sol";
import {PoolConfigManagerSetup} from "./PoolConfigManagerSetup.t.sol";

/// @notice Tokens A and B use the DEFAULT fee policy, which has a policy default record.
contract PoolConfigManager_getFeePolicy is PoolConfigManagerSetup {
  /// @notice Returns the token count and records for each policy. Chain selector zero identifies the policy default
  /// record.
  function test_getFeePolicy() public {
    FeePolicyStore.FeeConfig memory feeConfig = _getFeeConfig(100, 10);
    _setFeePolicyConfig(DEFAULT_POLICY, DEST_CHAIN_SELECTOR, feeConfig);
    _setFeePolicyConfig(OTHER_POLICY, OTHER_CHAIN_SELECTOR, feeConfig);

    (uint32 tokenCount, FeePolicyStore.FeePolicyConfigArgs[] memory policyConfigs) =
      s_poolConfigManager.getFeePolicy(DEFAULT_POLICY);
    assertEq(2, tokenCount);
    assertEq(2, policyConfigs.length);
    assertEq(PolicyConfig.DEFAULT_CHAIN_SELECTOR, policyConfigs[0].remoteChainSelector);
    assertEq(
      abi.encode(
        FeePolicyStore.FeePolicyConfigArgs({
          policyName: DEFAULT_POLICY, remoteChainSelector: DEST_CHAIN_SELECTOR, feeConfig: feeConfig
        })
      ),
      abi.encode(policyConfigs[1])
    );

    (tokenCount, policyConfigs) = s_poolConfigManager.getFeePolicy(OTHER_POLICY);
    assertEq(0, tokenCount);
    assertEq(1, policyConfigs.length);
    assertEq(OTHER_CHAIN_SELECTOR, policyConfigs[0].remoteChainSelector);
  }
}
