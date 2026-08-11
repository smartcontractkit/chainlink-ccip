// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {RateLimiter} from "../../../libraries/RateLimiter.sol";
import {RateLimiterHelper} from "../../helpers/RateLimiterHelper.sol";
import {Test} from "forge-std/Test.sol";

contract RateLimiterSetup is Test {
  RateLimiterHelper internal s_helper;
  RateLimiter.Config internal s_config;

  uint256 internal immutable i_blockTime = block.timestamp;

  function setUp() public virtual {
    s_config = RateLimiter.Config({isEnabled: true, rate: 5, capacity: 100});
    s_helper = new RateLimiterHelper(s_config);
  }
}
