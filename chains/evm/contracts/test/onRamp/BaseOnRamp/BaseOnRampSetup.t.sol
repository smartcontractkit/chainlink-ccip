// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {BaseOnRamp} from "../../../onRamp/BaseOnRamp.sol";
import {FeeQuoterSetup} from "../../feeQuoter/FeeQuoterSetup.t.sol";
import {BaseOnRampTestHelper} from "../../helpers/BaseOnRampTestHelper.sol";
import {BurnMintERC20} from "@chainlink/contracts/src/v0.8/shared/token/ERC20/BurnMintERC20.sol";

contract BaseOnRampSetup is FeeQuoterSetup {
  address internal constant FEE_AGGREGATOR = 0xa33CDB32eAEce34F6affEfF4899cef45744EDea3;
  address internal constant ALLOWLIST_ADMIN = 0x1234567890123456789012345678901234567890;

  BaseOnRampTestHelper internal s_baseOnRamp;

  address internal s_ccvProxy;

  function setUp() public virtual override {
    super.setUp();

    s_ccvProxy = makeAddr("CCVProxy");
    s_sourceFeeToken = address(new BurnMintERC20("Chainlink Token", "LINK", 18, 0, 0));

    s_baseOnRamp = new BaseOnRampTestHelper(address(s_mockRMNRemote));

    // Set up initial destination chain config.
    BaseOnRamp.DestChainConfigArgs[] memory destChainConfigs = new BaseOnRamp.DestChainConfigArgs[](1);
    destChainConfigs[0] = BaseOnRamp.DestChainConfigArgs({
      ccvProxy: s_ccvProxy,
      destChainSelector: DEST_CHAIN_SELECTOR,
      allowlistEnabled: false
    });

    s_baseOnRamp.applyDestChainConfigUpdates(destChainConfigs);

    vm.startPrank(OWNER);
  }

  /// @notice Helper to create a destination chain config.
  function _getDestChainConfig(
    address ccvProxy,
    uint64 destChainSelector,
    bool allowlistEnabled
  ) internal pure returns (BaseOnRamp.DestChainConfigArgs memory) {
    return BaseOnRamp.DestChainConfigArgs({
      ccvProxy: ccvProxy,
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
