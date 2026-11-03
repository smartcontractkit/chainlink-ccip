// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {PolicyConfig} from "../../../../libraries/PolicyConfig.sol";
import {FeePolicyStore} from "../../../../pools/MultiTokenPool/policies/FeePolicyStore.sol";
import {PoolConfigManagerSetup} from "./PoolConfigManagerSetup.t.sol";

/// @notice Tokens A and B use the DEFAULT destination overhead policy, which has a policy default record.
contract PoolConfigManager_getDestOverheadPolicy is PoolConfigManagerSetup {
  /// @notice Each policy returns its token count and its records in the shape of applyDestOverheadPolicyConfigUpdates.
  /// Chain selector zero is the policy default record.
  function test_getDestOverheadPolicy() public {
    _setDestOverheadPolicyConfig(DEFAULT_POLICY, DEST_CHAIN_SELECTOR, 300_000, 320);
    _setDestOverheadPolicyConfig(OTHER_POLICY, OTHER_CHAIN_SELECTOR, 400_000, 352);

    (uint32 tokenCount, FeePolicyStore.DestOverheadPolicyConfigArgs[] memory policyConfigs) =
      s_poolConfigManager.getDestOverheadPolicy(DEFAULT_POLICY);
    assertEq(2, tokenCount);
    assertEq(2, policyConfigs.length);
    assertEq(
      abi.encode(
        FeePolicyStore.DestOverheadPolicyConfigArgs({
          policyName: DEFAULT_POLICY,
          remoteChainSelector: PolicyConfig.DEFAULT_CHAIN_SELECTOR,
          destOverheadConfig: FeePolicyStore.DestOverheadConfig({
            destGasOverhead: DEFAULT_TOKEN_DEST_GAS_OVERHEAD, destBytesOverhead: MIN_DEST_BYTES_OVERHEAD
          })
        })
      ),
      abi.encode(policyConfigs[0])
    );
    assertEq(
      abi.encode(
        FeePolicyStore.DestOverheadPolicyConfigArgs({
          policyName: DEFAULT_POLICY,
          remoteChainSelector: DEST_CHAIN_SELECTOR,
          destOverheadConfig: FeePolicyStore.DestOverheadConfig({destGasOverhead: 300_000, destBytesOverhead: 320})
        })
      ),
      abi.encode(policyConfigs[1])
    );

    (tokenCount, policyConfigs) = s_poolConfigManager.getDestOverheadPolicy(OTHER_POLICY);
    assertEq(0, tokenCount);
    assertEq(1, policyConfigs.length);
    assertEq(OTHER_CHAIN_SELECTOR, policyConfigs[0].remoteChainSelector);
  }
}
