// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {TokenPoolV2} from "../../poolsV2/TokenPoolV2.sol";

import {IERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/IERC20.sol";

/// @notice Helper contract for testing TokenPool V2 functionality.
contract TokenPoolV2Helper is TokenPoolV2 {
  constructor(
    IERC20 token,
    uint8 localTokenDecimals,
    address[] memory allowlist,
    address rmnProxy,
    address router
  ) TokenPoolV2(token, localTokenDecimals, allowlist, rmnProxy, router) {}
}
