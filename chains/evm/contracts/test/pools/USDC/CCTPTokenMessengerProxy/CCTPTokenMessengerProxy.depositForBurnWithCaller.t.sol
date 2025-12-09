// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ITokenMessenger} from "../../../../pools/USDC/interfaces/ITokenMessenger.sol";

import {CCTPTokenMessengerProxySetup} from "./CCTPTokenMessengerProxySetup.t.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

contract CCTPTokenMessengerProxy_depositForBurnWithCaller is CCTPTokenMessengerProxySetup {
  function test_depositForBurnWithCaller() public {
    vm.startPrank(s_authorizedCaller);
    vm.expectEmit();
    emit ITokenMessenger.DepositForBurn(
      s_tokenMessenger.s_nonce(),
      address(s_USDCToken),
      BURN_AMOUNT,
      address(s_cctpTokenMessengerProxy),
      s_mintRecipient,
      DESTINATION_DOMAIN,
      s_tokenMessenger.DESTINATION_TOKEN_MESSENGER(),
      s_destinationCaller
    );

    uint64 nonce = s_cctpTokenMessengerProxy.depositForBurnWithCaller(
      BURN_AMOUNT, DESTINATION_DOMAIN, s_mintRecipient, address(s_USDCToken), s_destinationCaller
    );

    assertEq(nonce, s_tokenMessenger.s_nonce() - 1);
  }

  function test_depositForBurnWithCaller_RevertWhen_UnauthorizedCaller() public {
    address unauthorizedCaller = makeAddr("UNAUTHORIZED");

    vm.startPrank(unauthorizedCaller);
    vm.expectRevert(abi.encodeWithSelector(AuthorizedCallers.UnauthorizedCaller.selector, unauthorizedCaller));
    s_cctpTokenMessengerProxy.depositForBurnWithCaller(0, 0, bytes32(0), address(0), bytes32(0));
  }
}
