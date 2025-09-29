// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Client} from "../../../libraries/Client.sol";
import {TokenPoolV2Setup} from "./TokenPoolV2Setup.t.sol";

contract TokenPoolV2_getFee is TokenPoolV2Setup {
  function test_getFee() public view {
    Client.EVM2AnyMessage memory message;
    assertEq(0, s_tokenPool.getFee(DEST_CHAIN_SELECTOR, message, 0, ""));
  }
}
