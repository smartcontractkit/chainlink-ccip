// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {BaseVerifier} from "../../../ccvs/components/BaseVerifier.sol";
import {IRouter} from "../../../interfaces/IRouter.sol";
import {CommitteeVerifierSetup} from "./CommitteeVerifierSetup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract CommitteeVerifier_applyDestChainConfigUpdates is CommitteeVerifierSetup {
  uint64 internal constant NEW_DEST_SELECTOR = uint64(uint256(keccak256("COMMITTEE_RAMP_NEW_DEST_SELECTOR")));

  function test_applyDestChainConfigUpdates() public {
    address router = makeAddr("newRouter");

    BaseVerifier.DestChainConfigArgs[] memory args = new BaseVerifier.DestChainConfigArgs[](1);
    args[0] = BaseVerifier.DestChainConfigArgs({
      router: IRouter(router),
      destChainSelector: NEW_DEST_SELECTOR,
      allowlistEnabled: true
    });

    vm.expectEmit();
    emit BaseVerifier.DestChainConfigSet(NEW_DEST_SELECTOR, router, true);

    s_committeeVerifier.applyDestChainConfigUpdates(args);

    (bool allowlistEnabled, address newRouter, address[] memory allowedSenders) =
      s_committeeVerifier.getDestChainConfig(NEW_DEST_SELECTOR);

    assertEq(allowlistEnabled, true);
    assertEq(newRouter, router);
    assertEq(allowedSenders.length, 0);
  }

  function test_applyDestChainConfigUpdates_RevertWhen_OnlyOwner() public {
    BaseVerifier.DestChainConfigArgs[] memory args = new BaseVerifier.DestChainConfigArgs[](1);
    args[0] = BaseVerifier.DestChainConfigArgs({
      router: s_router,
      destChainSelector: NEW_DEST_SELECTOR,
      allowlistEnabled: false
    });

    vm.stopPrank();
    vm.prank(STRANGER);
    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_committeeVerifier.applyDestChainConfigUpdates(args);
  }
}
