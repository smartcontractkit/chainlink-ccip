// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {BaseVerifier} from "../../../ccvs/components/BaseVerifier.sol";
import {IRouter} from "../../../interfaces/IRouter.sol";
import {CommitteeVerifierSetup} from "./CommitteeVerifierSetup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract CommitteeVerifier_applyRemoteChainConfigUpdates is CommitteeVerifierSetup {
  uint64 internal constant NEW_DEST_SELECTOR = uint64(uint256(keccak256("COMMITTEE_RAMP_NEW_DEST_SELECTOR")));

  function test_applyRemoteChainConfigUpdates() public {
    address router = makeAddr("newRouter");

    BaseVerifier.RemoteChainConfigArgs[] memory args = new BaseVerifier.RemoteChainConfigArgs[](1);
    args[0] = BaseVerifier.RemoteChainConfigArgs({
      router: IRouter(router),
      remoteChainSelector: NEW_DEST_SELECTOR,
      allowlistEnabled: true,
      feeUSDCents: DEFAULT_CCV_FEE_USD_CENTS,
      gasForVerification: DEFAULT_CCV_GAS_LIMIT,
      payloadSizeBytes: DEFAULT_CCV_PAYLOAD_SIZE
    });

    vm.expectEmit();
    emit BaseVerifier.RemoteChainConfigSet(NEW_DEST_SELECTOR, router, true);

    s_committeeVerifier.applyRemoteChainConfigUpdates(args);

    (bool allowlistEnabled, address newRouter, address[] memory allowedSenders) =
      s_committeeVerifier.getRemoteChainConfig(NEW_DEST_SELECTOR);

    assertEq(allowlistEnabled, true);
    assertEq(newRouter, router);
    assertEq(allowedSenders.length, 0);
  }

  function test_applyRemoteChainConfigUpdates_RevertWhen_OnlyOwner() public {
    vm.stopPrank();

    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_committeeVerifier.applyRemoteChainConfigUpdates(new BaseVerifier.RemoteChainConfigArgs[](1));
  }
}
