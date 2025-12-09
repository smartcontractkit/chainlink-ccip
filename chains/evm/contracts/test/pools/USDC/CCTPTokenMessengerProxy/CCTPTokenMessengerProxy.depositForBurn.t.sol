// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ITokenMessenger} from "../../../../pools/USDC/interfaces/ITokenMessenger.sol";

import {CCTPTokenMessengerProxySetup} from "./CCTPTokenMessengerProxySetup.t.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

contract CCTPTokenMessengerProxy_depositForBurn is CCTPTokenMessengerProxySetup {
  function test_depositForBurn() public {
    vm.startPrank(s_authorizedCaller);
    vm.expectEmit();
    emit ITokenMessenger.DepositForBurn(
      address(s_USDCToken),
      BURN_AMOUNT,
      address(s_cctpTokenMessengerProxy),
      s_mintRecipient,
      DESTINATION_DOMAIN,
      s_tokenMessenger.DESTINATION_TOKEN_MESSENGER(),
      s_destinationCaller,
      MAX_FEE,
      MIN_FINALITY_THRESHOLD,
      ""
    );

    s_cctpTokenMessengerProxy.depositForBurn(
      BURN_AMOUNT,
      DESTINATION_DOMAIN,
      s_mintRecipient,
      address(s_USDCToken),
      s_destinationCaller,
      MAX_FEE,
      MIN_FINALITY_THRESHOLD
    );
  }

  function test_depositForBurn_RevertWhen_UnauthorizedCaller() public {
    address unauthorizedCaller = makeAddr("UNAUTHORIZED");

    vm.startPrank(unauthorizedCaller);
    vm.expectRevert(abi.encodeWithSelector(AuthorizedCallers.UnauthorizedCaller.selector, unauthorizedCaller));
    s_cctpTokenMessengerProxy.depositForBurn(0, 0, bytes32(0), address(0), bytes32(0), 0, 0);
  }
}
