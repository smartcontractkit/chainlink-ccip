// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Pool} from "../../libraries/Pool.sol";
import {USDCTokenPoolProxy} from "../../pools/USDC/USDCTokenPoolProxy.sol";

import {IERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/IERC20.sol";

contract USDCTokenPoolProxyHelper is USDCTokenPoolProxy {
  constructor(
    IERC20 token,
    USDCTokenPoolProxy.PoolAddresses memory pools,
    address router,
    address cctpVerifier
  ) USDCTokenPoolProxy(token, pools, router, cctpVerifier) {}

  function generateNewReleaseOrMintIn(
    Pool.ReleaseOrMintInV1 calldata releaseOrMintIn
  ) public pure returns (Pool.ReleaseOrMintInV1 memory newReleaseOrMintIn) {
    return _generateNewReleaseOrMintIn(releaseOrMintIn);
  }
}
