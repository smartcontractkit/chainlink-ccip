// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Executor} from "../../../executor/Executor.sol";
import {Client} from "../../../libraries/Client.sol";
import {ExecutorSetup} from "./ExecutorSetup.t.sol";

contract Executor_getFee is ExecutorSetup {
  function test_getFee_EmptyMessage() public view {
    address[] memory ccvAddresses = new address[](1);
    ccvAddresses[0] = INITIAL_CCV;

    uint16 fee = s_executor.getFee(DEST_CHAIN_SELECTOR, 0, ccvAddresses, "");

    assertEq(DEFAULT_EXEC_FEE_USD_CENTS, fee);
  }

  function test_getFee_RevertWhen_Executor__RequestedBlockDepthTooLow() public {
    uint16 requestedBlockDepth = MIN_BLOCK_CONFIRMATIONS - 1;

    vm.expectRevert(
      abi.encodeWithSelector(
        Executor.Executor__RequestedBlockDepthTooLow.selector, requestedBlockDepth, MIN_BLOCK_CONFIRMATIONS
      )
    );
    s_executor.getFee(DEST_CHAIN_SELECTOR, requestedBlockDepth, new address[](1), "");
  }

  function test_getFee_RevertWhen_InvalidDestChain() public {
    address[] memory ccvAddresses = new address[](1);
    ccvAddresses[0] = INITIAL_CCV;

    vm.expectRevert(abi.encodeWithSelector(Executor.InvalidDestChain.selector, DEST_CHAIN_SELECTOR + 1));
    s_executor.getFee(DEST_CHAIN_SELECTOR + 1, 0, ccvAddresses, "");
  }

  function test_getFee_RevertWhen_UnsupportedRequiredCCV() public {
    address unsupportedCCV = makeAddr("unsupportedCCV");
    address[] memory ccvAddresses = new address[](1);
    ccvAddresses[0] = unsupportedCCV;

    vm.expectRevert(abi.encodeWithSelector(Executor.InvalidCCV.selector, unsupportedCCV));
    s_executor.getFee(DEST_CHAIN_SELECTOR, 0, ccvAddresses, "");
  }

  function test_getFee_RevertWhen_ExceedsMaxCCVs() public {
    address[] memory ccvAddresses = new address[](2);
    ccvAddresses[0] = INITIAL_CCV;
    ccvAddresses[1] = INITIAL_CCV;

    vm.expectRevert(abi.encodeWithSelector(Executor.ExceedsMaxCCVs.selector, ccvAddresses.length, INITIAL_MAX_CCVS));
    s_executor.getFee(DEST_CHAIN_SELECTOR, 0, ccvAddresses, "");
  }
}
