// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {CCTPMessageTransmitterProxy} from "../../../../pools/USDC/CCTPMessageTransmitterProxy.sol";
import {ITokenMessenger} from "../../../../pools/USDC/interfaces/ITokenMessenger.sol";
import {BaseTest} from "../../../BaseTest.t.sol";

contract CCTPMessageTransmitterProxy_constructor is BaseTest {
  address internal s_tokenMessenger = makeAddr("TOKEN_MESSENGER");
  address internal s_cctpMessageTransmitter = makeAddr("CCTP_MT");

  function test_constructor() public {
    vm.mockCall(
      s_tokenMessenger,
      abi.encodeWithSelector(ITokenMessenger.localMessageTransmitter.selector),
      abi.encode(s_cctpMessageTransmitter)
    );

    CCTPMessageTransmitterProxy proxy = new CCTPMessageTransmitterProxy(ITokenMessenger(s_tokenMessenger));

    assertEq(address(proxy.i_cctpTransmitter()), s_cctpMessageTransmitter);
    assertEq(proxy.typeAndVersion(), "CCTPMessageTransmitterProxy 2.0.0-dev");
  }

  // Reverts

  function test_constructor_RevertWhen_TransmitterIsZero() public {
    vm.mockCall(
      s_tokenMessenger, abi.encodeWithSelector(ITokenMessenger.localMessageTransmitter.selector), abi.encode(address(0))
    );

    vm.expectRevert(CCTPMessageTransmitterProxy.TransmitterCannotBeZero.selector);
    new CCTPMessageTransmitterProxy(ITokenMessenger(s_tokenMessenger));
  }
}

