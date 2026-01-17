// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IRouter} from "../../../interfaces/IRouter.sol";

import {BaseOnRamp} from "../../../onRamp/BaseOnRamp.sol";
import {FeeQuoterSetup} from "../../feeQuoter/FeeQuoterSetup.t.sol";
import {BaseOnRampTestHelper} from "../../helpers/BaseOnRampTestHelper.sol";
import {BurnMintERC20} from "@chainlink/contracts/src/v0.8/shared/token/ERC20/BurnMintERC20.sol";

contract BaseOnRampSetup is FeeQuoterSetup {
  address internal constant FEE_AGGREGATOR = 0xa33CDB32eAEce34F6affEfF4899cef45744EDea3;
  address internal constant ALLOWLIST_ADMIN = 0x1234567890123456789012345678901234567890;

  BaseOnRampTestHelper internal s_baseOnRamp;

  IRouter internal s_router;
  address internal s_ccvProxy;
  address internal s_ccvAggregatorRemote;

  function setUp() public virtual override {
    super.setUp();

    s_router = IRouter(makeAddr("Router"));
    s_ccvProxy = makeAddr("CCVProxy");
    vm.mockCall(
      address(s_router), abi.encodeWithSelector(IRouter.getOnRamp.selector, DEST_CHAIN_SELECTOR), abi.encode(s_ccvProxy)
    );
    s_ccvAggregatorRemote = makeAddr("CCVAggregatorRemote");
    s_sourceFeeToken = address(new BurnMintERC20("Chainlink Token", "LINK", 18, 0, 0));

    s_baseOnRamp = new BaseOnRampTestHelper();

    // Set up initial destination chain config.
    BaseOnRamp.DestChainConfigArgs[] memory destChainConfigs = new BaseOnRamp.DestChainConfigArgs[](1);
    destChainConfigs[0] = BaseOnRamp.DestChainConfigArgs({
      router: s_router,
      destChainSelector: DEST_CHAIN_SELECTOR,
      allowlistEnabled: false
    });

    s_baseOnRamp.applyDestChainConfigUpdates(destChainConfigs);

    vm.startPrank(OWNER);
  }

  /// @notice Helper to create a destination chain config.
  function _getDestChainConfig(
    IRouter router,
    uint64 destChainSelector,
    bool allowlistEnabled
  ) internal pure returns (BaseOnRamp.DestChainConfigArgs memory) {
    return BaseOnRamp.DestChainConfigArgs({
      router: router,
      destChainSelector: destChainSelector,
      allowlistEnabled: allowlistEnabled
    });
  }

  /// @notice Helper to create allowlist config args.
  function _getAllowlistConfig(
    uint64 destChainSelector,
    bool allowlistEnabled,
    address[] memory addedSenders,
    address[] memory removedSenders
  ) internal pure returns (BaseOnRamp.AllowlistConfigArgs memory) {
    return BaseOnRamp.AllowlistConfigArgs({
      destChainSelector: destChainSelector,
      allowlistEnabled: allowlistEnabled,
      addedAllowlistedSenders: addedSenders,
      removedAllowlistedSenders: removedSenders
    });
  }
}
