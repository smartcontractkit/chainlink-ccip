// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Executor} from "../../../executor/Executor.sol";
import {Client} from "../../../libraries/Client.sol";
import {ExecutorSetup} from "./ExecutorSetup.t.sol";

contract Executor_getFee is ExecutorSetup {
  function test_getFee() public view {
    Client.EVM2AnyMessage memory message;
    Client.CCV[] memory ccvs = new Client.CCV[](1);
    ccvs[0] = Client.CCV({ccvAddress: INITIAL_CCV, args: ""});

    uint256 fee = s_executor.getFee(INITIAL_DEST, message, ccvs, "");
    assertEq(fee, 0);
  }

  function test_getFee_RevertWhen_InvalidDestChain() public {
    Client.EVM2AnyMessage memory message;
    Client.CCV[] memory ccvs = new Client.CCV[](1);
    ccvs[0] = Client.CCV({ccvAddress: INITIAL_CCV, args: ""});

    vm.expectRevert(abi.encodeWithSelector(Executor.InvalidDestChain.selector, 999));
    s_executor.getFee(999, message, ccvs, "");
  }

  function test_getFee_RevertWhen_UnsupportedRequiredCCV() public {
    address unsupportedCCV = makeAddr("unsupportedCCV");
    Client.EVM2AnyMessage memory message;
    Client.CCV[] memory ccvs = new Client.CCV[](1);
    ccvs[0] = Client.CCV({ccvAddress: unsupportedCCV, args: ""});

    vm.expectRevert(abi.encodeWithSelector(Executor.InvalidCCV.selector, unsupportedCCV));
    s_executor.getFee(INITIAL_DEST, message, ccvs, "");
  }

  function test_getFee_RevertWhen_ExceedsMaxCCVs() public {
    Client.EVM2AnyMessage memory message;
    Client.CCV[] memory ccvs = new Client.CCV[](2);
    ccvs[0] = Client.CCV({ccvAddress: INITIAL_CCV, args: ""});
    ccvs[1] = Client.CCV({ccvAddress: INITIAL_CCV, args: ""});

    vm.expectRevert(abi.encodeWithSelector(Executor.ExceedsMaxCCVs.selector, ccvs.length, INITIAL_MAX_CCVS));
    s_executor.getFee(INITIAL_DEST, message, ccvs, "");
  }
}
