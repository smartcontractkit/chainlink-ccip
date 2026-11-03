// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../../../interfaces/IPoolConfigManager.sol";

import {FeePolicyStore} from "../../../../pools/MultiTokenPool/policies/FeePolicyStore.sol";
import {PoolConfigManagerSetup} from "./PoolConfigManagerSetup.t.sol";

/// @notice Tokens A and B use the DEFAULT destination overhead policy, which has a policy default record. Token B has a
/// lane to the destination chain and no lane to the other chain.
contract PoolConfigManager_getInlineDestOverheadConfigs is PoolConfigManagerSetup {
  /// @notice The view returns each inline destination overhead record of the token in the shape of
  /// applyInlineDestOverheadConfigUpdates, in the order of the sets. A record on a chain with a lane and a record on a
  /// chain without a lane both show.
  function test_getInlineDestOverheadConfigs() public {
    FeePolicyStore.InlineDestOverheadConfigArgs[] memory expectedConfigs =
      new FeePolicyStore.InlineDestOverheadConfigArgs[](2);
    expectedConfigs[0] = FeePolicyStore.InlineDestOverheadConfigArgs({
      localToken: s_tokenB,
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      destOverheadConfig: FeePolicyStore.DestOverheadConfig({destGasOverhead: 300_000, destBytesOverhead: 320})
    });
    expectedConfigs[1] = FeePolicyStore.InlineDestOverheadConfigArgs({
      localToken: s_tokenB,
      remoteChainSelector: OTHER_CHAIN_SELECTOR,
      destOverheadConfig: FeePolicyStore.DestOverheadConfig({destGasOverhead: 400_000, destBytesOverhead: 352})
    });
    s_poolConfigManager.applyInlineDestOverheadConfigUpdates(new IPoolConfigManager.LaneKey[](0), expectedConfigs);

    assertEq(abi.encode(expectedConfigs), abi.encode(s_poolConfigManager.getInlineDestOverheadConfigs(s_tokenB)));
  }
}
