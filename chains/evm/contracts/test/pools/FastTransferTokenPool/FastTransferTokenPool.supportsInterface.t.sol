// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IAny2EVMMessageReceiver} from "../../../interfaces/IAny2EVMMessageReceiver.sol";
import {IFastTransferPool} from "../../../interfaces/IFastTransferPool.sol";

import {FastTransferTokenPoolSetup} from "./FastTransferTokenPoolSetup.t.sol";
import {IERC165} from "@openzeppelin/contracts@4.8.3/utils/introspection/IERC165.sol";

contract FastTransferTokenPool_supportsInterface is FastTransferTokenPoolSetup {
  function test_supportsInterface() public view {
    assertTrue(s_pool.supportsInterface(type(IFastTransferPool).interfaceId));
    assertTrue(s_pool.supportsInterface(type(IERC165).interfaceId));
    assertTrue(s_pool.supportsInterface(type(IAny2EVMMessageReceiver).interfaceId));
  }
}
