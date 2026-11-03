// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../../../interfaces/IPoolConfigManager.sol";

import {FeePolicyStore} from "../../../../pools/MultiTokenPool/policies/FeePolicyStore.sol";
import {PoolConfigManagerSetup} from "./PoolConfigManagerSetup.t.sol";

/// @notice Tokens A and B use the DEFAULT fee policy, which has a disabled policy default record. Token B has a lane to
/// the destination chain and no lane to the other chain.
contract PoolConfigManager_getInlineFeeConfigs is PoolConfigManagerSetup {
  /// @notice The view returns each inline fee record of the token in the shape of applyInlineFeeConfigUpdates, in the
  /// order of the sets. A record on a chain with a lane and a zero record on a chain without a lane both show.
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

  /// @dev An enabled record. The FTF values are twice the default finality values.
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
