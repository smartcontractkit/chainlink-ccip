// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {CCTPTokenMessengerProxy} from "../../../../pools/USDC/CCTPTokenMessengerProxy.sol";
import {CCTPTokenMessengerProxySetup} from "./CCTPTokenMessengerProxySetup.t.sol";

contract CCTPTokenMessengerProxy_constructor is CCTPTokenMessengerProxySetup {
  function test_Constructor() public {
    address[] memory expectedAuthorizedCallers = new address[](1);
    expectedAuthorizedCallers[0] = s_authorizedCaller;
    CCTPTokenMessengerProxy proxy =
      new CCTPTokenMessengerProxy(s_tokenMessenger, s_USDCToken, expectedAuthorizedCallers);

    assertEq(proxy.getTokenMessenger(), address(s_tokenMessenger));
    assertEq(proxy.messageBodyVersion(), s_tokenMessenger.messageBodyVersion());
    assertEq(proxy.localMessageTransmitter(), s_messageTransmitter);

    address[] memory actualAuthorizedCallers = proxy.getAllAuthorizedCallers();
    assertEq(actualAuthorizedCallers.length, 1);
    assertEq(actualAuthorizedCallers[0], s_authorizedCaller);
  }
}
