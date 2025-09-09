// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {BaseOnRamp} from "../../../onRamp/BaseOnRamp.sol";
import {CommitOnRampSetup} from "./CommitOnRampSetup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract CommitOnRamp_applyDestChainConfigUpdates is CommitOnRampSetup {
  uint64 internal constant NEW_DEST_SELECTOR = uint64(uint256(keccak256("COMMIT_ONRAMP_NEW_DEST_SELECTOR")));

  function test_applyDestChainConfigUpdates() public {
    address newCCVProxy = makeAddr("newCCVProxy");

    BaseOnRamp.DestChainConfigArgs[] memory args = new BaseOnRamp.DestChainConfigArgs[](1);
    args[0] = BaseOnRamp.DestChainConfigArgs({
      ccvProxy: newCCVProxy,
      destChainSelector: NEW_DEST_SELECTOR,
      allowlistEnabled: true
    });

    vm.expectEmit();
    emit BaseOnRamp.DestChainConfigSet(NEW_DEST_SELECTOR, newCCVProxy, true);

    s_commitOnRamp.applyDestChainConfigUpdates(args);

    (bool allowlistEnabled, address ccvProxy, address[] memory allowedSenders) =
      s_commitOnRamp.getDestChainConfig(NEW_DEST_SELECTOR);

    assertEq(allowlistEnabled, true);
    assertEq(ccvProxy, newCCVProxy);
    assertEq(allowedSenders.length, 0);
  }

  function test_applyDestChainConfigUpdates_RevertWhen_OnlyOwner() public {
    BaseOnRamp.DestChainConfigArgs[] memory args = new BaseOnRamp.DestChainConfigArgs[](1);
    args[0] = BaseOnRamp.DestChainConfigArgs({
      ccvProxy: makeAddr("someCCV"),
      destChainSelector: NEW_DEST_SELECTOR,
      allowlistEnabled: false
    });

    vm.startPrank(STRANGER);
    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_commitOnRamp.applyDestChainConfigUpdates(args);
    vm.stopPrank();
  }
}
