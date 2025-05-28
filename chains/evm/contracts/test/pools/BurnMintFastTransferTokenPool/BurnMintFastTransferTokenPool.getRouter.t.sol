// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {BurnMintFastTransferTokenPoolSetup} from "./BurnMintFastTransferTokenPoolSetup.t.sol";

contract BurnMintFastTransferTokenPool_getRouter is BurnMintFastTransferTokenPoolSetup {
  function test_GetRouter() public view {
    assertEq(s_pool.getRouter(), address(s_sourceRouter));
  }

  function test_TypeAndVersion() public view {
    assertEq(s_pool.typeAndVersion(), "BurnMintFastTransferTokenPool 1.6.1");
  }
}
