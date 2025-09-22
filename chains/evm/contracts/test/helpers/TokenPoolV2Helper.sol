// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Client} from "../../libraries/Client.sol";
import {Pool} from "../../libraries/Pool.sol";
import {TokenPoolV2} from "../../pools/TokenPoolV2.sol";

import {IERC20} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";

contract TokenPoolV2Helper is TokenPoolV2 {
  constructor(
    IERC20 token,
    uint8 localTokenDecimals,
    address[] memory allowlist,
    address rmnProxy,
    address router
  ) TokenPoolV2(token, localTokenDecimals, allowlist, rmnProxy, router) {}

  function getFee(
    uint64,
    address,
    address,
    Client.EVMTokenAmount[] calldata,
    bytes calldata
  ) external pure returns (Pool.Quote memory quote) {
    return quote;
  }
}
