// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Router} from "../../../Router.sol";
import {PingPongDemo} from "../../../applications/PingPongDemo.sol";

import {PingPongDappSetup} from "./PingPongDappSetup.t.sol";

contract PingPong_startPingPong is PingPongDappSetup {
  uint256 internal s_pingPongNumber = 1;

  function test_StartPingPong_With_Sequenced_Ordered() public {
    _assertPingPongSuccess();
  }

  function test_StartPingPong_With_OOO() public {
    s_pingPong.setOutOfOrderExecution(true);

    _assertPingPongSuccess();
  }

  function _assertPingPongSuccess() internal {
    vm.expectEmit();
    emit PingPongDemo.Ping(s_pingPongNumber);

    vm.expectCall(address(s_sourceRouter), abi.encodeWithSelector(Router.ccipSend.selector));

    s_pingPong.startPingPong();
  }
}
