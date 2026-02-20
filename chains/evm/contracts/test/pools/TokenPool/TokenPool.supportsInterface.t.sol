// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolV1} from "../../../interfaces/IPool.sol";
import {IPoolV2} from "../../../interfaces/IPoolV2.sol";
import {Pool} from "../../../libraries/Pool.sol";
import {AdvancedPoolHooksSetup} from "../AdvancedPoolHooks/AdvancedPoolHooksSetup.t.sol";

import {IERC165} from "@openzeppelin/contracts@5.3.0/utils/introspection/IERC165.sol";

contract TokenPool_supportsInterface is AdvancedPoolHooksSetup {
  function test_supportsInterface() public view {
    assertTrue(s_tokenPool.supportsInterface(type(IERC165).interfaceId));
    assertTrue(s_tokenPool.supportsInterface(type(IPoolV2).interfaceId));
    assertTrue(s_tokenPool.supportsInterface(type(IPoolV1).interfaceId));
    assertTrue(s_tokenPool.supportsInterface(Pool.CCIP_POOL_V1));
  }
}
