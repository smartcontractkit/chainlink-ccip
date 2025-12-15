// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IRouter} from "../../../interfaces/IRouter.sol";

import {BaseVerifier} from "../../../ccvs/components/BaseVerifier.sol";
import {CCTPVerifierSetup} from "./CCTPVerifierSetup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract CCTPVerifier_applyRemoteChainConfigUpdates is CCTPVerifierSetup {
  function test_applyRemoteChainConfigUpdates() public {
    address router = makeAddr("newRouter");
    uint64 newChainSelector = uint64(uint256(keccak256("CCTP_VERIFIER_NEW_SELECTOR")));

    BaseVerifier.RemoteChainConfigArgs[] memory args = new BaseVerifier.RemoteChainConfigArgs[](1);
    args[0] = BaseVerifier.RemoteChainConfigArgs({
      router: IRouter(router),
      remoteChainSelector: newChainSelector,
      allowlistEnabled: true,
      feeUSDCents: DEFAULT_CCV_FEE_USD_CENTS,
      gasForVerification: DEFAULT_CCV_GAS_LIMIT,
      payloadSizeBytes: DEFAULT_CCV_PAYLOAD_SIZE
    });

    vm.expectEmit();
    emit BaseVerifier.RemoteChainConfigSet(newChainSelector, router, true);

    s_cctpVerifier.applyRemoteChainConfigUpdates(args);

    (bool allowlistEnabled, address newRouter, address[] memory allowedSenders) =
      s_cctpVerifier.getRemoteChainConfig(newChainSelector);

    assertEq(allowlistEnabled, true);
    assertEq(newRouter, router);
    assertEq(allowedSenders.length, 0);
  }

  function test_applyRemoteChainConfigUpdates_RevertWhen_OnlyCallableByOwner() public {
    vm.stopPrank();

    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_cctpVerifier.applyRemoteChainConfigUpdates(new BaseVerifier.RemoteChainConfigArgs[](1));
  }
}
