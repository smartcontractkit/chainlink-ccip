// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Executor} from "../../../executor/Executor.sol";
import {FinalityCodec} from "../../../libraries/FinalityCodec.sol";
import {ExecutorSetup} from "./ExecutorSetup.t.sol";

contract Executor_getFee is ExecutorSetup {
  function test_getFee_EmptyMessage() public view {
    address[] memory ccvAddresses = new address[](1);
    ccvAddresses[0] = INITIAL_CCV;

    uint16 fee = s_executor.getFee(DEST_CHAIN_SELECTOR, bytes2(0), ccvAddresses, "", s_sourceFeeToken);

    assertEq(DEFAULT_EXEC_FEE_USD_CENTS, fee);
  }

  function test_getFee_RevertWhen_InvalidRequestedFinality() public {
    bytes2 requestedFinality = bytes2(uint16(MIN_FINALITY_CONFIG) - 1);

    vm.expectRevert(
      abi.encodeWithSelector(FinalityCodec.InvalidRequestedFinality.selector, requestedFinality, MIN_FINALITY_CONFIG)
    );
    s_executor.getFee(DEST_CHAIN_SELECTOR, requestedFinality, new address[](1), "", s_sourceFeeToken);
  }

  function test_getFee_RevertWhen_InvalidDestChain() public {
    address[] memory ccvAddresses = new address[](1);
    ccvAddresses[0] = INITIAL_CCV;

    vm.expectRevert(abi.encodeWithSelector(Executor.InvalidDestChain.selector, DEST_CHAIN_SELECTOR + 1));
    s_executor.getFee(DEST_CHAIN_SELECTOR + 1, bytes2(0), ccvAddresses, "", s_sourceFeeToken);
  }

  function test_getFee_RevertWhen_UnsupportedRequiredCCV() public {
    address unsupportedCCV = makeAddr("unsupportedCCV");
    address[] memory ccvAddresses = new address[](1);
    ccvAddresses[0] = unsupportedCCV;

    vm.expectRevert(abi.encodeWithSelector(Executor.InvalidCCV.selector, unsupportedCCV));
    s_executor.getFee(DEST_CHAIN_SELECTOR, bytes2(0), ccvAddresses, "", s_sourceFeeToken);
  }

  function test_getFee_RevertWhen_ExceedsMaxCCVs() public {
    address[] memory ccvAddresses = new address[](2);
    ccvAddresses[0] = INITIAL_CCV;
    ccvAddresses[1] = INITIAL_CCV;

    vm.expectRevert(abi.encodeWithSelector(Executor.ExceedsMaxCCVs.selector, ccvAddresses.length, INITIAL_MAX_CCVS));
    s_executor.getFee(DEST_CHAIN_SELECTOR, bytes2(0), ccvAddresses, "", s_sourceFeeToken);
  }
}
