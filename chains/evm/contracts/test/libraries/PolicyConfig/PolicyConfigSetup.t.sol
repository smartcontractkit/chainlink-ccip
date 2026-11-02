// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {PolicyConfigHelper} from "../../helpers/PolicyConfigHelper.sol";
import {Test} from "forge-std/Test.sol";

contract PolicyConfigSetup is Test {
  bytes32 internal constant POLICY = "POLICY";
  bytes32 internal constant OTHER_POLICY = "OTHER_POLICY";
  uint64 internal constant CHAIN_SELECTOR = 5009297550715157269;
  uint64 internal constant OTHER_CHAIN_SELECTOR = 4949039107694359620;

  PolicyConfigHelper internal s_helper;
  address internal s_token;

  function setUp() public virtual {
    s_helper = new PolicyConfigHelper();
    s_token = makeAddr("token");
    s_helper.registerToken(s_token);
  }
}
