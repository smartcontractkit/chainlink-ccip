// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Executor} from "../../../executor/Executor.sol";
import {Client} from "../../../libraries/Client.sol";
import {ExecutorSetup} from "./ExecutorSetup.t.sol";

contract Executor_getFee is ExecutorSetup {
  function test_getFee_EmptyMessage() public view {
    Client.CCV[] memory ccvs = new Client.CCV[](1);
    ccvs[0] = Client.CCV({ccvAddress: INITIAL_CCV, args: ""});

    uint16 fee = s_executor.getFee(DEST_CHAIN_SELECTOR, 0, ccvs, "");

    assertEq(DEFAULT_EXEC_FEE_USD_CENTS, fee);
  }

  function test_getFee_RevertWhen_Executor__RequestedBlockDepthTooLow() public {
    uint16 requestedBlockDepth = MIN_BLOCK_CONFIRMATIONS - 1;

    vm.expectRevert(
      abi.encodeWithSelector(
        Executor.Executor__RequestedBlockDepthTooLow.selector, requestedBlockDepth, MIN_BLOCK_CONFIRMATIONS
      )
    );
    s_executor.getFee(DEST_CHAIN_SELECTOR, requestedBlockDepth, new Client.CCV[](1), "");
  }

  function test_getFee_RevertWhen_InvalidDestChain() public {
    Client.CCV[] memory ccvs = new Client.CCV[](1);
    ccvs[0] = Client.CCV({ccvAddress: INITIAL_CCV, args: ""});

    vm.expectRevert(abi.encodeWithSelector(Executor.InvalidDestChain.selector, DEST_CHAIN_SELECTOR + 1));
    s_executor.getFee(DEST_CHAIN_SELECTOR + 1, 0, ccvs, "");
  }

  function test_getFee_RevertWhen_UnsupportedRequiredCCV() public {
    address unsupportedCCV = makeAddr("unsupportedCCV");
    Client.CCV[] memory ccvs = new Client.CCV[](1);
    ccvs[0] = Client.CCV({ccvAddress: unsupportedCCV, args: ""});

    vm.expectRevert(abi.encodeWithSelector(Executor.InvalidCCV.selector, unsupportedCCV));
    s_executor.getFee(DEST_CHAIN_SELECTOR, 0, ccvs, "");
  }

  function test_getFee_RevertWhen_ExceedsMaxCCVs() public {
    Client.CCV[] memory ccvs = new Client.CCV[](2);
    ccvs[0] = Client.CCV({ccvAddress: INITIAL_CCV, args: ""});
    ccvs[1] = Client.CCV({ccvAddress: INITIAL_CCV, args: ""});

    vm.expectRevert(abi.encodeWithSelector(Executor.ExceedsMaxCCVs.selector, ccvs.length, INITIAL_MAX_CCVS));
    s_executor.getFee(DEST_CHAIN_SELECTOR, 0, ccvs, "");
  }
}
