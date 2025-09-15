// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IRouter} from "../../../interfaces/IRouter.sol";
import {BaseOnRamp} from "../../../onRamp/BaseOnRamp.sol";
import {CommitOnRampSetup} from "./CommitOnRampSetup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract CommitOnRamp_applyDestChainConfigUpdates is CommitOnRampSetup {
  uint64 internal constant NEW_DEST_SELECTOR = uint64(uint256(keccak256("COMMIT_ONRAMP_NEW_DEST_SELECTOR")));

  function test_applyDestChainConfigUpdates() public {
    address router = makeAddr("newRouter");

    BaseOnRamp.DestChainConfigArgs[] memory args = new BaseOnRamp.DestChainConfigArgs[](1);
    args[0] = BaseOnRamp.DestChainConfigArgs({
      router: IRouter(router),
      destChainSelector: NEW_DEST_SELECTOR,
      allowlistEnabled: true
    });

    vm.expectEmit();
    emit BaseOnRamp.DestChainConfigSet(NEW_DEST_SELECTOR, router, true);

    s_commitOnRamp.applyDestChainConfigUpdates(args);

    (bool allowlistEnabled, address newRouter, address[] memory allowedSenders) =
      s_commitOnRamp.getDestChainConfig(NEW_DEST_SELECTOR);

    assertEq(allowlistEnabled, true);
    assertEq(newRouter, router);
    assertEq(allowedSenders.length, 0);
  }

  function test_applyDestChainConfigUpdates_RevertWhen_OnlyOwner() public {
    BaseOnRamp.DestChainConfigArgs[] memory args = new BaseOnRamp.DestChainConfigArgs[](1);
    args[0] =
      BaseOnRamp.DestChainConfigArgs({router: s_router, destChainSelector: NEW_DEST_SELECTOR, allowlistEnabled: false});

    vm.stopPrank();
    vm.prank(STRANGER);
    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_commitOnRamp.applyDestChainConfigUpdates(args);
  }
}
