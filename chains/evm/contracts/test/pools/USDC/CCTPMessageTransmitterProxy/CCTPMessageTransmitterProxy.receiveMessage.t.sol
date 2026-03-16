// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IMessageTransmitter} from "../../../../pools/USDC/interfaces/IMessageTransmitter.sol";

import {CCTPMessageTransmitterProxySetup} from "./CCTPMessageTransmitterProxySetup.t.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

contract CCTPMessageTransmitterProxy_receiveMessage is CCTPMessageTransmitterProxySetup {
  function test_receiveMessage() public {
    bytes memory message = bytes("message");
    bytes memory attestation = bytes("attestation");

    // Mocking the call to the IMessageTransmitter to return true
    vm.mockCall(
      s_cctpMessageTransmitter,
      abi.encodeWithSelector(IMessageTransmitter.receiveMessage.selector, message, attestation),
      abi.encode(true)
    );

    changePrank(s_usdcTokenPool);
    assertTrue(s_cctpMessageTransmitterProxy.receiveMessage(message, attestation));

    // Mocking the call to the IMessageTransmitter to return false
    vm.mockCall(
      s_cctpMessageTransmitter,
      abi.encodeWithSelector(IMessageTransmitter.receiveMessage.selector, message, attestation),
      abi.encode(false)
    );

    changePrank(s_usdcTokenPool);
    assertFalse(s_cctpMessageTransmitterProxy.receiveMessage(message, attestation));
  }

  // Revert cases
  function test_receiveMessage_RevertWhen_UnauthorizedCaller() public {
    bytes memory message = bytes("message");
    bytes memory attestation = bytes("attestation");

    address random = makeAddr("RANDOM");
    changePrank(random);
    vm.expectRevert(abi.encodeWithSelector(AuthorizedCallers.UnauthorizedCaller.selector, random));
    s_cctpMessageTransmitterProxy.receiveMessage(message, attestation);
  }
}
