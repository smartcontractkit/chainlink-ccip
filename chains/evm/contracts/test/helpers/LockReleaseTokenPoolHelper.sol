// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {LockReleaseTokenPool} from "../../pools/LockReleaseTokenPool.sol";
import {IERC20} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";

contract LockReleaseTokenPoolHelper is LockReleaseTokenPool {
  constructor(
    IERC20 token,
    address[] memory allowlist,
    address rmnProxy,
    address router,
    address previousPool
  ) LockReleaseTokenPool(token, 6, allowlist, rmnProxy, router) {}
}
