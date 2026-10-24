// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {MultiTokenPoolSetup} from "../MultiTokenPool/MultiTokenPoolSetup.t.sol";

contract BurnMintMultiTokenPool_typeAndVersion is MultiTokenPoolSetup {
  function test_typeAndVersion() public view {
    assertEq("BurnMintMultiTokenPool 2.1.0-dev", s_pool.typeAndVersion());
  }
}
