// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IGetPoolConfigManager} from "../../../interfaces/IGetPoolConfigManager.sol";
import {IMultiTokenPool} from "../../../interfaces/IMultiTokenPool.sol";
import {IPoolV1} from "../../../interfaces/IPool.sol";
import {IPoolV2} from "../../../interfaces/IPoolV2.sol";

import {Pool} from "../../../libraries/Pool.sol";
import {MultiTokenPoolSetup} from "./MultiTokenPoolSetup.t.sol";

import {IERC165} from "@openzeppelin/contracts@5.3.0/utils/introspection/IERC165.sol";

contract MultiTokenPool_supportsInterface is MultiTokenPoolSetup {
  function test_supportsInterface() public view {
    assertTrue(s_pool.supportsInterface(Pool.CCIP_POOL_V1));
    assertTrue(s_pool.supportsInterface(type(IPoolV1).interfaceId));
    assertTrue(s_pool.supportsInterface(type(IPoolV2).interfaceId));
    assertTrue(s_pool.supportsInterface(type(IMultiTokenPool).interfaceId));
    assertTrue(s_pool.supportsInterface(type(IGetPoolConfigManager).interfaceId));
    assertTrue(s_pool.supportsInterface(type(IERC165).interfaceId));
  }

  function test_supportsInterface_FalseForUnknown() public view {
    assertFalse(s_pool.supportsInterface(bytes4(0xffffffff)));
  }
}
