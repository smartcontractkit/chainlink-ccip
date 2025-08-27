// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IRouter} from "../../../interfaces/IRouter.sol";

import {Client} from "../../../libraries/Client.sol";
import {Internal} from "../../../libraries/Internal.sol";
import {CCVAggregator} from "../../../offRamp/CCVAggregator.sol";

import {BaseTest} from "../../BaseTest.t.sol";
import {CCVAggregatorHelper} from "../../helpers/CCVAggregatorHelper.sol";

contract CCVAggregatorSetup is BaseTest {
  CCVAggregatorHelper internal s_agg;
  address internal s_defaultCCV;
  address internal s_tokenAdminRegistry;

  function setUp() public virtual override {
    BaseTest.setUp();

    s_defaultCCV = makeAddr("defaultCCV");
    s_tokenAdminRegistry = makeAddr("tokenAdminRegistry");

    s_agg = new CCVAggregatorHelper(
      CCVAggregator.StaticConfig({
        localChainSelector: DEST_CHAIN_SELECTOR,
        gasForCallExactCheck: GAS_FOR_CALL_EXACT_CHECK,
        rmnRemote: s_mockRMNRemote,
        tokenAdminRegistry: s_tokenAdminRegistry
      })
    );

    // Apply initial source chain configuration
    _applySourceConfig(
      s_sourceRouter, SOURCE_CHAIN_SELECTOR, abi.encode(makeAddr("onRamp")), true, new address[](1), new address[](0)
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
    defaultCCVs[0] = s_defaultCCV;

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
