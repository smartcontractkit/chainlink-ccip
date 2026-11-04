// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../../../interfaces/IPoolConfigManager.sol";

import {FeePolicyStore} from "../../../../pools/MultiTokenPool/policies/FeePolicyStore.sol";
import {PoolConfigManagerSetup} from "./PoolConfigManagerSetup.t.sol";

/// @notice Tokens A and B use the DEFAULT fee policy, which has a default record with zero fees. Token B has a lane
/// to the destination chain and no lane to the other chain.
contract PoolConfigManager_getInlineFeeConfigs is PoolConfigManagerSetup {
  /// @notice Returns both inline fee records, including the record with zero fees on a chain without a lane.
  function test_getInlineFeeConfigs() public {
    FeePolicyStore.InlineFeeConfigArgs[] memory expectedConfigs = new FeePolicyStore.InlineFeeConfigArgs[](2);
    expectedConfigs[0] = _getInlineFeeConfigArgs(s_tokenB, DEST_CHAIN_SELECTOR, 100, 10);
    expectedConfigs[1] = FeePolicyStore.InlineFeeConfigArgs({
      localToken: s_tokenB,
      remoteChainSelector: OTHER_CHAIN_SELECTOR,
      feeConfig: FeePolicyStore.FeeConfig({
        finalityFeeUSDCents: 0, fastFinalityFeeUSDCents: 0, finalityTransferFeeBps: 0, fastFinalityTransferFeeBps: 0
      })
    });
    _applyInlineFeeConfigs(expectedConfigs);

    assertEq(abi.encode(expectedConfigs), abi.encode(s_poolConfigManager.getInlineFeeConfigs(s_tokenB)));
  }

  /// @dev FTF fees are twice the default finality fees.
  function _getInlineFeeConfigArgs(
    address localToken,
    uint64 remoteChainSelector,
    uint32 finalityFeeUSDCents,
    uint16 finalityTransferFeeBps
  ) internal pure returns (FeePolicyStore.InlineFeeConfigArgs memory) {
    FeePolicyStore.FeeConfig memory feeConfig = _getFeeConfig(finalityFeeUSDCents, finalityTransferFeeBps);
    return FeePolicyStore.InlineFeeConfigArgs({
      localToken: localToken, remoteChainSelector: remoteChainSelector, feeConfig: feeConfig
    });
  }

  function _applyInlineFeeConfigs(
    FeePolicyStore.InlineFeeConfigArgs[] memory inlineConfigs
  ) internal {
    s_poolConfigManager.applyInlineFeeConfigUpdates(new IPoolConfigManager.LaneKey[](0), inlineConfigs);
  }
}
