// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IEVM2AnyOnRampClient} from "../../interfaces/IEVM2AnyOnRampClient.sol";

import {Router} from "../../Router.sol";
import {OnRamp} from "../../onRamp/OnRamp.sol";
import {FeeQuoterSetup} from "../feeQuoter/FeeQuoterSetup.t.sol";
import {MockExecutor} from "../mocks/MockExecutor.sol";
import {MockVerifier} from "../mocks/MockVerifier.sol";

import {IERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/IERC20.sol";

contract RouterSetup is FeeQuoterSetup {
  address internal constant FEE_AGGREGATOR = 0xa33CDB32eAEce34F6affEfF4899cef45744EDea3;

  OnRamp internal s_onRamp;
  address internal s_offRamp;

  function setUp() public virtual override {
    super.setUp();

    s_onRamp = new OnRamp(
      OnRamp.StaticConfig({
        chainSelector: SOURCE_CHAIN_SELECTOR,
        rmnRemote: s_mockRMNRemote,
        tokenAdminRegistry: address(s_tokenAdminRegistry)
      }),
      OnRamp.DynamicConfig({
        feeQuoter: address(s_feeQuoter),
        reentrancyGuardEntered: false,
        feeAggregator: FEE_AGGREGATOR
      })
    );

    address[] memory defaultCCVs = new address[](1);
    defaultCCVs[0] = address(new MockVerifier(""));

    OnRamp.DestChainConfigArgs[] memory destChainConfigs = new OnRamp.DestChainConfigArgs[](1);
    destChainConfigs[0] = OnRamp.DestChainConfigArgs({
      destChainSelector: DEST_CHAIN_SELECTOR,
      router: s_sourceRouter,
      laneMandatedCCVs: new address[](0),
      defaultCCVs: defaultCCVs,
      defaultExecutor: address(new MockExecutor()),
      offRamp: abi.encodePacked(s_offRamp)
    });
    s_onRamp.applyDestChainConfigUpdates(destChainConfigs);

    Router.OnRamp[] memory onRampUpdates = new Router.OnRamp[](1);
    onRampUpdates[0] = Router.OnRamp({destChainSelector: DEST_CHAIN_SELECTOR, onRamp: address(s_onRamp)});

    Router.OffRamp[] memory offRampUpdates = new Router.OffRamp[](2);
    offRampUpdates[0] = Router.OffRamp({sourceChainSelector: SOURCE_CHAIN_SELECTOR, offRamp: makeAddr("offRamp0")});
    offRampUpdates[1] = Router.OffRamp({sourceChainSelector: SOURCE_CHAIN_SELECTOR, offRamp: makeAddr("offRamp1")});
    s_sourceRouter.applyRampUpdates(onRampUpdates, new Router.OffRamp[](0), offRampUpdates);

    // Mock a fee for the onRamp so the router can calculate fees.
    vm.mockCall(address(s_onRamp), abi.encodeWithSelector(IEVM2AnyOnRampClient.getFee.selector), abi.encode(1e17));

    // Pre approve the first token so the gas estimates of the tests only cover actual gas usage from the ramps.
    IERC20(s_sourceTokens[0]).approve(address(s_sourceRouter), 2 ** 128);
    IERC20(s_sourceTokens[1]).approve(address(s_sourceRouter), 2 ** 128);
  }
}
