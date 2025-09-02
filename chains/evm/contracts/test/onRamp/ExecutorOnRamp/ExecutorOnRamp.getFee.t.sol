// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Client} from "../../../libraries/Client.sol";
import {ExecutorOnRamp} from "../../../onRamp/ExecutorOnRamp.sol";
import {ExecutorOnRampSetup} from "./ExecutorOnRampSetup.t.sol";

contract ExecutorOnRamp_getFee is ExecutorOnRampSetup {
  function test_getFee() public view {
    Client.EVM2AnyMessage memory message;
    Client.CCV[] memory requiredCCVs = new Client.CCV[](1);
    requiredCCVs[0] = Client.CCV({ccvAddress: INITIAL_CCV, args: ""});
    Client.CCV[] memory optionalCCVs = new Client.CCV[](0);

    uint256 fee = s_executorOnRamp.getFee(INITIAL_DEST, message, requiredCCVs, optionalCCVs, "");
    assertEq(fee, 0);
  }

  function test_getFee_RevertWhen_InvalidDestChain() public {
    Client.EVM2AnyMessage memory message;
    Client.CCV[] memory requiredCCVs = new Client.CCV[](1);
    requiredCCVs[0] = Client.CCV({ccvAddress: INITIAL_CCV, args: ""});
    Client.CCV[] memory optionalCCVs = new Client.CCV[](0);

    vm.expectRevert(abi.encodeWithSelector(ExecutorOnRamp.InvalidDestChain.selector, 999));
    s_executorOnRamp.getFee(999, message, requiredCCVs, optionalCCVs, "");
  }

  function test_getFee_RevertWhen_UnsupportedRequiredCCV() public {
    address unsupportedCCV = makeAddr("unsupportedCCV");
    Client.EVM2AnyMessage memory message;
    Client.CCV[] memory requiredCCVs = new Client.CCV[](1);
    requiredCCVs[0] = Client.CCV({ccvAddress: unsupportedCCV, args: ""});
    Client.CCV[] memory optionalCCVs = new Client.CCV[](0);

    vm.expectRevert(abi.encodeWithSelector(ExecutorOnRamp.InvalidCCV.selector, unsupportedCCV));
    s_executorOnRamp.getFee(INITIAL_DEST, message, requiredCCVs, optionalCCVs, "");
  }

  function test_getFee_RevertWhen_UnsupportedOptionalCCV() public {
    address unsupportedCCV = makeAddr("unsupportedCCV");
    Client.EVM2AnyMessage memory message;
    Client.CCV[] memory requiredCCVs = new Client.CCV[](1);
    requiredCCVs[0] = Client.CCV({ccvAddress: INITIAL_CCV, args: ""});
    Client.CCV[] memory optionalCCVs = new Client.CCV[](1);
    optionalCCVs[0] = Client.CCV({ccvAddress: unsupportedCCV, args: ""});

    vm.expectRevert(abi.encodeWithSelector(ExecutorOnRamp.InvalidCCV.selector, unsupportedCCV));
    s_executorOnRamp.getFee(INITIAL_DEST, message, requiredCCVs, optionalCCVs, "");
  }

  function test_getFee_RevertWhen_ExceedsMaxCCVs() public {
    Client.EVM2AnyMessage memory message;
    Client.CCV[] memory requiredCCVs = new Client.CCV[](1);
    requiredCCVs[0] = Client.CCV({ccvAddress: INITIAL_CCV, args: ""});
    Client.CCV[] memory optionalCCVs = new Client.CCV[](1);
    optionalCCVs[0] = Client.CCV({ccvAddress: INITIAL_CCV, args: ""});

    vm.expectRevert(
      abi.encodeWithSelector(
        ExecutorOnRamp.ExceedsMaxCCVs.selector, requiredCCVs.length + optionalCCVs.length, INITIAL_MAX_CCVS
      )
    );
    s_executorOnRamp.getFee(INITIAL_DEST, message, requiredCCVs, optionalCCVs, "");
  }
}
