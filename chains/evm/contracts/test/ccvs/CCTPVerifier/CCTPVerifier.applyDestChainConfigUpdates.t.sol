// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {BaseVerifier} from "../../../ccvs/components/BaseVerifier.sol";
import {IRouter} from "../../../interfaces/IRouter.sol";
import {CCTPVerifierSetup} from "./CCTPVerifierSetup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract CCTPVerifier_applyDestChainConfigUpdates is CCTPVerifierSetup {
  uint64 internal constant NEW_DEST_SELECTOR = uint64(uint256(keccak256("CCTP_VERIFIER_NEW_DEST_SELECTOR")));

  function test_applyDestChainConfigUpdates() public {
    address router = makeAddr("newRouter");

    BaseVerifier.DestChainConfigArgs[] memory args = new BaseVerifier.DestChainConfigArgs[](1);
    args[0] = BaseVerifier.DestChainConfigArgs({
      router: IRouter(router),
      destChainSelector: NEW_DEST_SELECTOR,
      allowlistEnabled: true,
      feeUSDCents: DEFAULT_CCV_FEE_USD_CENTS,
      gasForVerification: DEFAULT_CCV_GAS_LIMIT,
      payloadSizeBytes: DEFAULT_CCV_PAYLOAD_SIZE
    });

    vm.expectEmit();
    emit BaseVerifier.DestChainConfigSet(NEW_DEST_SELECTOR, router, true);

    s_cctpVerifier.applyDestChainConfigUpdates(args);

    (bool allowlistEnabled, address newRouter, address[] memory allowedSenders) =
      s_cctpVerifier.getDestChainConfig(NEW_DEST_SELECTOR);

    assertEq(allowlistEnabled, true);
    assertEq(newRouter, router);
    assertEq(allowedSenders.length, 0);
  }

  function test_applyDestChainConfigUpdates_RevertWhen_OnlyCallableByOwner() public {
    BaseVerifier.DestChainConfigArgs[] memory args = new BaseVerifier.DestChainConfigArgs[](1);
    args[0] = BaseVerifier.DestChainConfigArgs({
      router: s_router,
      destChainSelector: NEW_DEST_SELECTOR,
      allowlistEnabled: false,
      feeUSDCents: DEFAULT_CCV_FEE_USD_CENTS,
      gasForVerification: DEFAULT_CCV_GAS_LIMIT,
      payloadSizeBytes: DEFAULT_CCV_PAYLOAD_SIZE
    });

    vm.stopPrank();
    vm.prank(STRANGER);
    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_cctpVerifier.applyDestChainConfigUpdates(args);
  }
}
