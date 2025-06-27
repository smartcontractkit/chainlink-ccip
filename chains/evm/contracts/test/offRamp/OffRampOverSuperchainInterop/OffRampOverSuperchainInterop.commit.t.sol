// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {OffRampOverSuperchainInterop} from "../../../offRamp/OffRampOverSuperchainInterop.sol";
import {OffRampOverSuperchainInteropSetup} from "./OffRampOverSuperchainInteropSetup.t.sol";

contract OffRampOverSuperchainInterop_commit is OffRampOverSuperchainInteropSetup {
  function test_commit_RevertWhen_Always() public {
    bytes32[2] memory reportContext = [bytes32(0), bytes32(0)];
    bytes memory report = abi.encode("test report");
    bytes32[] memory rs = new bytes32[](1);
    bytes32[] memory ss = new bytes32[](1);
    bytes32 rawVs = keccak256("raw_vs");

    vm.expectRevert(OffRampOverSuperchainInterop.OperationNotSupportedbyThisOffRampType.selector);
    s_offRampOverSuperchainInterop.commit(reportContext, report, rs, ss, rawVs);
  }
}
