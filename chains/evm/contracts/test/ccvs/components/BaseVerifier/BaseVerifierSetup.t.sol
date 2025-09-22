// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IRouter} from "../../../../interfaces/IRouter.sol";

import {BaseVerifier} from "../../../../ccvs/components/BaseVerifier.sol";
import {FeeQuoterSetup} from "../../../feeQuoter/FeeQuoterSetup.t.sol";
import {BaseVerifierTestHelper} from "../../../helpers/BaseVerifierTestHelper.sol";
import {BurnMintERC20} from "@chainlink/contracts/src/v0.8/shared/token/ERC20/BurnMintERC20.sol";

contract BaseVerifierSetup is FeeQuoterSetup {
  address internal constant FEE_AGGREGATOR = 0xa33CDB32eAEce34F6affEfF4899cef45744EDea3;
  address internal constant ALLOWLIST_ADMIN = 0x1234567890123456789012345678901234567890;

  BaseVerifierTestHelper internal s_baseVerifier;

  IRouter internal s_router;
  address internal s_ccvProxy;
  address internal s_ccvAggregatorRemote;

  string internal constant STORAGE_LOCATION = "testStorageLocation";

  function setUp() public virtual override {
    super.setUp();

    s_router = IRouter(makeAddr("Router"));
    s_ccvProxy = makeAddr("CCVProxy");
    vm.mockCall(
      address(s_router), abi.encodeWithSelector(IRouter.getOnRamp.selector, DEST_CHAIN_SELECTOR), abi.encode(s_ccvProxy)
    );
    s_ccvAggregatorRemote = makeAddr("CCVAggregatorRemote");
    s_sourceFeeToken = address(new BurnMintERC20("Chainlink Token", "LINK", 18, 0, 0));

    s_baseVerifier = new BaseVerifierTestHelper(STORAGE_LOCATION);

    // Set up initial destination chain config.
    BaseVerifier.DestChainConfigArgs[] memory destChainConfigs = new BaseVerifier.DestChainConfigArgs[](1);
    destChainConfigs[0] = BaseVerifier.DestChainConfigArgs({
      router: s_router,
      destChainSelector: DEST_CHAIN_SELECTOR,
      allowlistEnabled: false
    });

    s_baseVerifier.applyDestChainConfigUpdates(destChainConfigs);

    vm.startPrank(OWNER);
  }

  /// @notice Helper to create a destination chain config.
  function _getDestChainConfig(
    IRouter router,
    uint64 destChainSelector,
    bool allowlistEnabled
  ) internal pure returns (BaseVerifier.DestChainConfigArgs memory) {
    return BaseVerifier.DestChainConfigArgs({
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
  ) internal pure returns (BaseVerifier.AllowlistConfigArgs memory) {
    return BaseVerifier.AllowlistConfigArgs({
      destChainSelector: destChainSelector,
      allowlistEnabled: allowlistEnabled,
      addedAllowlistedSenders: addedSenders,
      removedAllowlistedSenders: removedSenders
    });
  }
}
