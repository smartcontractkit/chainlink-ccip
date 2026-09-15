// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {BaseVerifier} from "../../../ccvs/components/BaseVerifier.sol";
import {SuccinctZKVerifierSetup} from "./SuccinctZKVerifierSetup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract SuccinctZKVerifier_applyAllowlistUpdates is SuccinctZKVerifierSetup {
  function test_applyAllowlistUpdates() public {
    address[] memory senders = new address[](1);
    senders[0] = makeAddr("sender1");
    BaseVerifier.AllowlistConfigArgs[] memory allowlistConfigs = new BaseVerifier.AllowlistConfigArgs[](1);
    allowlistConfigs[0] = _getAllowlistConfig(DEST_CHAIN_SELECTOR, true, senders, new address[](0));

    vm.expectEmit();
    emit BaseVerifier.AllowListStateChanged(DEST_CHAIN_SELECTOR, true);
    vm.expectEmit();
    emit BaseVerifier.AllowListSendersAdded(DEST_CHAIN_SELECTOR, senders[0]);

    s_zkVerifier.applyAllowlistUpdates(allowlistConfigs);

    (BaseVerifier.RemoteChainConfigArgs memory config, address[] memory allowedSenders) =
      s_zkVerifier.getRemoteChainConfig(DEST_CHAIN_SELECTOR);
    assertTrue(config.allowlistEnabled);
    assertEq(senders, allowedSenders);
  }

  // Reverts

  function test_applyAllowlistUpdates_RevertWhen_OnlyCallableByOwner() public {
    vm.stopPrank();
    vm.startPrank(STRANGER);

    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_zkVerifier.applyAllowlistUpdates(new BaseVerifier.AllowlistConfigArgs[](0));
  }
}
