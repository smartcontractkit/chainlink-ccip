// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IRouter} from "../../../interfaces/IRouter.sol";

import {NonceManager} from "../../../NonceManager.sol";
import {Router} from "../../../Router.sol";
import {Client} from "../../../libraries/Client.sol";
import {Internal} from "../../../libraries/Internal.sol";
import {OnRamp} from "../../../onRamp/OnRamp.sol";
import {TokenAdminRegistry} from "../../../tokenAdminRegistry/TokenAdminRegistry.sol";
import {FeeQuoterSetup} from "../../feeQuoter/FeeQuoterSetup.t.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

import {IERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/IERC20.sol";

contract OnRampSetup is FeeQuoterSetup {
  address internal constant FEE_AGGREGATOR = 0xa33CDB32eAEce34F6affEfF4899cef45744EDea3;

  OnRamp internal s_onRamp;
  NonceManager internal s_outboundNonceManager;

  function setUp() public virtual override {
    super.setUp();

    s_outboundNonceManager = new NonceManager(new address[](0));
    s_onRamp = new OnRamp(
      OnRamp.StaticConfig({
        chainSelector: SOURCE_CHAIN_SELECTOR,
        rmnRemote: s_mockRMNRemote,
        nonceManager: address(s_outboundNonceManager),
        tokenAdminRegistry: address(s_tokenAdminRegistry)
      }),
      OnRamp.DynamicConfig({
        feeQuoter: address(s_feeQuoter),
        reentrancyGuardEntered: false,
        messageInterceptor: address(0),
        feeAggregator: FEE_AGGREGATOR,
        allowlistAdmin: address(0)
      }),
      _generateDestChainConfigArgs(s_sourceRouter)
    );

    address[] memory authorizedCallers = new address[](1);
    authorizedCallers[0] = address(s_onRamp);

    s_outboundNonceManager.applyAuthorizedCallerUpdates(
      AuthorizedCallers.AuthorizedCallerArgs({addedCallers: authorizedCallers, removedCallers: new address[](0)})
    );

    Router.OnRamp[] memory onRampUpdates = new Router.OnRamp[](1);
    onRampUpdates[0] = Router.OnRamp({destChainSelector: DEST_CHAIN_SELECTOR, onRamp: address(s_onRamp)});

    Router.OffRamp[] memory offRampUpdates = new Router.OffRamp[](2);
    offRampUpdates[0] = Router.OffRamp({sourceChainSelector: SOURCE_CHAIN_SELECTOR, offRamp: makeAddr("offRamp0")});
    offRampUpdates[1] = Router.OffRamp({sourceChainSelector: SOURCE_CHAIN_SELECTOR, offRamp: makeAddr("offRamp1")});
    s_sourceRouter.applyRampUpdates(onRampUpdates, new Router.OffRamp[](0), offRampUpdates);

    // Pre approve the first token so the gas estimates of the tests
    // only cover actual gas usage from the ramps
    IERC20(s_sourceTokens[0]).approve(address(s_sourceRouter), 2 ** 128);
    IERC20(s_sourceTokens[1]).approve(address(s_sourceRouter), 2 ** 128);
  }

  function _generateDestChainConfigArgs(
    IRouter router
  ) internal pure returns (OnRamp.DestChainConfigArgs[] memory) {
    OnRamp.DestChainConfigArgs[] memory destChainConfigs = new OnRamp.DestChainConfigArgs[](1);
    destChainConfigs[0] =
      OnRamp.DestChainConfigArgs({destChainSelector: DEST_CHAIN_SELECTOR, router: router, allowlistEnabled: false});
    return destChainConfigs;
  }
}
