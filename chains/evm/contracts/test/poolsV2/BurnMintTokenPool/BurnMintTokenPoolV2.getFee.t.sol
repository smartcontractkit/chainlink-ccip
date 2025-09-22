// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Client} from "../../../libraries/Client.sol";
import {BurnMintSetup} from "./BurnMintSetup.t.sol";

contract BurnMintTokenPoolV2_releaseOrMint is BurnMintSetup {
  function test_getFee() public {
    Client.EVM2AnyMessage memory message;
    uint256 fee = s_pool.getFee(DEST_CHAIN_SELECTOR, message);

    assertEq(fee, 0);
  }

  // TODO, complete tests once billing is finalised.
}
