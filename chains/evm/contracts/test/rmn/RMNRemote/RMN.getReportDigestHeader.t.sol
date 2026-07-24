// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {RMNRemoteSetup} from "./RMNSetup.t.sol";

contract RMNRemote_getReportDigestHeader is RMNRemoteSetup {
  function test_getReportDigestHeader() public view {
    assertEq(s_rmn.getReportDigestHeader(), keccak256("RMN_V1_6_ANY2EVM_REPORT"));
  }
}
