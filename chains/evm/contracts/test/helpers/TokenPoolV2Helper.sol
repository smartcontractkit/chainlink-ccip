// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Client} from "../../libraries/Client.sol";
import {TokenPool} from "../../poolsV2/TokenPool.sol";

import {IERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/IERC20.sol";

/// @notice Helper contract for testing TokenPool V2 functionality.
contract TokenPoolV2Helper is TokenPool {
  constructor(
    IERC20 token,
    uint8 localTokenDecimals,
    address[] memory allowlist,
    address rmnProxy,
    address router
  ) TokenPool(token, localTokenDecimals, allowlist, rmnProxy, router) {}

  function getFee(
    uint64, // destChainSelector
    Client.EVM2AnyMessage calldata // message
  ) external view returns (uint256 feeTokenAmount) {
    return 0;
  }
}
