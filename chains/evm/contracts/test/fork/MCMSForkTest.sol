// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Test} from "forge-std/Test.sol";
import {IOwner} from "../../interfaces/IOwner.sol";

contract MCMSForkTest is Test {
    struct Call {
        address target;
        uint256 value;
        bytes data;
    }

    error TransactionReverted();

    function applyPayload(bytes memory payload) internal {
        (MCMSForkTest.Call[] memory calls, , , ) = abi.decode(payload, (MCMSForkTest.Call[], bytes32, bytes32, uint256));
        for (uint256 i = 0; i < calls.length; ++i) {
            MCMSForkTest.Call memory call = calls[i];
            vm.startPrank(IOwner(call.target).owner());
            (bool success, ) = call.target.call{value: call.value}(call.data);
            if (!success) revert TransactionReverted();
            vm.stopPrank();
        }
    }
}
