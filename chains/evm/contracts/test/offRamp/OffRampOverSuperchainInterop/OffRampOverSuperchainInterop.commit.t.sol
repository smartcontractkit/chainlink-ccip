// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {OffRampOverSuperchainInterop} from "../../../offRamp/OffRampOverSuperchainInterop.sol";
import {OffRampOverSuperchainInteropSetup} from "./OffRampOverSuperchainInteropSetup.t.sol";

contract OffRampOverSuperchainInterop_commit is OffRampOverSuperchainInteropSetup {
  function test_commit_RevertWhen_Always() public {
    bytes32[2] memory reportContext;
    bytes memory report;
    bytes32[] memory rs;
    bytes32[] memory ss;
    bytes32 rawVs;

    vm.expectRevert(OffRampOverSuperchainInterop.OperationNotSupportedByThisOffRampType.selector);
    s_offRampOverSuperchainInterop.commit(reportContext, report, rs, ss, rawVs);
  }
}
