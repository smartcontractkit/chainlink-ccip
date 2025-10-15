// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Executor} from "../../../executor/Executor.sol";
import {Client} from "../../../libraries/Client.sol";
import {ExecutorSetup} from "./ExecutorSetup.t.sol";

contract Executor_getFee is ExecutorSetup {
  function test_getFee() public view {
    Client.EVM2AnyMessage memory message;
    Client.CCV[] memory requiredCCVs = new Client.CCV[](1);
    requiredCCVs[0] = Client.CCV({ccvAddress: INITIAL_CCV, args: ""});
    Client.CCV[] memory optionalCCVs = new Client.CCV[](0);

    uint256 fee = s_executor.getFee(INITIAL_DEST, message, requiredCCVs, optionalCCVs, "");
    assertEq(fee, 0);
  }

  function test_getFee_RevertWhen_InvalidDestChain() public {
    Client.EVM2AnyMessage memory message;
    Client.CCV[] memory requiredCCVs = new Client.CCV[](1);
    requiredCCVs[0] = Client.CCV({ccvAddress: INITIAL_CCV, args: ""});
    Client.CCV[] memory optionalCCVs = new Client.CCV[](0);

    vm.expectRevert(abi.encodeWithSelector(Executor.InvalidDestChain.selector, 999));
    s_executor.getFee(999, message, requiredCCVs, optionalCCVs, "");
  }

  function test_getFee_RevertWhen_UnsupportedRequiredCCV() public {
    address unsupportedCCV = makeAddr("unsupportedCCV");
    Client.EVM2AnyMessage memory message;
    Client.CCV[] memory requiredCCVs = new Client.CCV[](1);
    requiredCCVs[0] = Client.CCV({ccvAddress: unsupportedCCV, args: ""});
    Client.CCV[] memory optionalCCVs = new Client.CCV[](0);

    vm.expectRevert(abi.encodeWithSelector(Executor.InvalidCCV.selector, unsupportedCCV));
    s_executor.getFee(INITIAL_DEST, message, requiredCCVs, optionalCCVs, "");
  }

  function test_getFee_RevertWhen_UnsupportedOptionalCCV() public {
    address unsupportedCCV = makeAddr("unsupportedCCV");
    Client.EVM2AnyMessage memory message;
    Client.CCV[] memory requiredCCVs = new Client.CCV[](1);
    requiredCCVs[0] = Client.CCV({ccvAddress: INITIAL_CCV, args: ""});
    Client.CCV[] memory optionalCCVs = new Client.CCV[](1);
    optionalCCVs[0] = Client.CCV({ccvAddress: unsupportedCCV, args: ""});

    vm.expectRevert(abi.encodeWithSelector(Executor.InvalidCCV.selector, unsupportedCCV));
    s_executor.getFee(INITIAL_DEST, message, requiredCCVs, optionalCCVs, "");
  }

  function test_getFee_RevertWhen_ExceedsMaxCCVs() public {
    Client.EVM2AnyMessage memory message;
    Client.CCV[] memory requiredCCVs = new Client.CCV[](1);
    requiredCCVs[0] = Client.CCV({ccvAddress: INITIAL_CCV, args: ""});
    Client.CCV[] memory optionalCCVs = new Client.CCV[](1);
    optionalCCVs[0] = Client.CCV({ccvAddress: INITIAL_CCV, args: ""});

    vm.expectRevert(
      abi.encodeWithSelector(
        Executor.ExceedsMaxCCVs.selector, requiredCCVs.length + optionalCCVs.length, INITIAL_MAX_CCVS
      )
    );
    s_executor.getFee(INITIAL_DEST, message, requiredCCVs, optionalCCVs, "");
  }
}
