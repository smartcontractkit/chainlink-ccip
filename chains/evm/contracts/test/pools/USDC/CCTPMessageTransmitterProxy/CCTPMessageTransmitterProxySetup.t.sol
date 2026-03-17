// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ITokenMessenger} from "../../../../pools/USDC/interfaces/ITokenMessenger.sol";

import {CCTPMessageTransmitterProxy} from "../../../../pools/USDC/CCTPMessageTransmitterProxy.sol";
import {BaseTest} from "../../../BaseTest.t.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

contract CCTPMessageTransmitterProxySetup is BaseTest {
  address internal s_tokenMessenger = makeAddr("TOKEN_MESSENGER");
  address internal s_cctpMessageTransmitter = makeAddr("CCTP_MT");
  address internal s_usdcTokenPool = makeAddr("USDC_TP");
  CCTPMessageTransmitterProxy internal s_cctpMessageTransmitterProxy;

  function setUp() public virtual override {
    super.setUp();
    vm.mockCall(
      s_tokenMessenger,
      abi.encodeWithSelector(ITokenMessenger.localMessageTransmitter.selector),
      abi.encode(s_cctpMessageTransmitter)
    );
    s_cctpMessageTransmitterProxy = new CCTPMessageTransmitterProxy(ITokenMessenger(s_tokenMessenger));

    address[] memory addedCallers = new address[](1);
    addedCallers[0] = s_usdcTokenPool;
    s_cctpMessageTransmitterProxy.applyAuthorizedCallerUpdates(
      AuthorizedCallers.AuthorizedCallerArgs({addedCallers: addedCallers, removedCallers: new address[](0)})
    );
  }
}
