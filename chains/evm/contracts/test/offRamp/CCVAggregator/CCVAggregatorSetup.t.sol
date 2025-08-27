// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IRMNRemote} from "../../../interfaces/IRMNRemote.sol";
import {IRouter} from "../../../interfaces/IRouter.sol";
import {CCVAggregator} from "../../../offRamp/CCVAggregator.sol";
import {FeeQuoterSetup} from "../../feeQuoter/FeeQuoterSetup.t.sol";

contract CCVAggregatorHelper is CCVAggregator {
  constructor(
    CCVAggregator.StaticConfig memory staticConfig
  ) CCVAggregator(staticConfig) {}

  function getCCVsFromReceiver(
    uint64 sourceChainSelector,
    address receiver
  ) external view returns (address[] memory requiredCCV, address[] memory optionalCCVs, uint8 optionalThreshold) {
    return _getCCVsFromReceiver(sourceChainSelector, receiver);
  }
}

contract CCVAggregatorSetup is FeeQuoterSetup {
  CCVAggregatorHelper internal s_agg;

  function setUp() public virtual override {
    FeeQuoterSetup.setUp();

    s_agg = new CCVAggregatorHelper(
      CCVAggregator.StaticConfig({
        localChainSelector: DEST_CHAIN_SELECTOR,
        gasForCallExactCheck: GAS_FOR_CALL_EXACT_CHECK,
        rmnRemote: s_mockRMNRemote,
        tokenAdminRegistry: address(s_tokenAdminRegistry)
      })
    );
  }

  function _applySourceConfig(
    IRouter router,
    uint64 sourceChainSelector,
    bytes memory onRamp,
    bool isEnabled,
    address[] memory defaultCCVs,
    address[] memory laneMandatedCCVs
  ) internal {
    CCVAggregator.SourceChainConfigArgs[] memory updates = new CCVAggregator.SourceChainConfigArgs[](1);
    updates[0] = CCVAggregator.SourceChainConfigArgs({
      router: router,
      sourceChainSelector: sourceChainSelector,
      isEnabled: isEnabled,
      onRamp: onRamp,
      defaultCCV: defaultCCVs,
      laneMandatedCCVs: laneMandatedCCVs
    });
    s_agg.applySourceChainConfigUpdates(updates);
  }
}
