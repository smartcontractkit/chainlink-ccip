// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Executor} from "../../../executor/Executor.sol";
import {Client} from "../../../libraries/Client.sol";
import {MessageV1Codec} from "../../../libraries/MessageV1Codec.sol";
import {ExecutorSetup} from "./ExecutorSetup.t.sol";

contract Executor_getFee is ExecutorSetup {
  function test_getFee_EmptyMessage() public view {
    Client.CCV[] memory ccvs = new Client.CCV[](1);
    ccvs[0] = Client.CCV({ccvAddress: INITIAL_CCV, args: ""});

    (uint16 fee, uint32 gas, uint32 execBytes) = s_executor.getFee(DEST_CHAIN_SELECTOR, 0, 0, 0, ccvs, "");

    uint32 expectedBytes = uint32(MessageV1Codec.MESSAGE_V1_EVM_SOURCE_BASE_SIZE + 2 * EVM_ADDRESS_LENGTH);

    assertEq(DEFAULT_EXEC_FEE_USD_CENTS, fee);
    assertEq(DEFAULT_EXEC_GAS, gas);
    assertEq(execBytes, expectedBytes);
  }

  function test_getFee_WithData() public view {
    Client.CCV[] memory ccvs = new Client.CCV[](1);
    ccvs[0] = Client.CCV({ccvAddress: INITIAL_CCV, args: ""});
    uint32 dataLength = 100;

    (uint16 fee, uint32 gas, uint32 execBytes) = s_executor.getFee(DEST_CHAIN_SELECTOR, 0, dataLength, 0, ccvs, "");

    uint32 expectedBytes = uint32(MessageV1Codec.MESSAGE_V1_EVM_SOURCE_BASE_SIZE + 2 * EVM_ADDRESS_LENGTH + dataLength);

    assertEq(DEFAULT_EXEC_FEE_USD_CENTS, fee);
    assertEq(DEFAULT_EXEC_GAS, gas);
    assertEq(execBytes, expectedBytes);
  }

  function test_getFee_WithTokens() public view {
    Client.CCV[] memory ccvs = new Client.CCV[](1);
    ccvs[0] = Client.CCV({ccvAddress: INITIAL_CCV, args: ""});
    uint8 numberOfTokens = 1;

    (uint16 fee, uint32 gas, uint32 execBytes) = s_executor.getFee(DEST_CHAIN_SELECTOR, 0, 0, numberOfTokens, ccvs, "");

    uint32 expectedBytes = uint32(
      MessageV1Codec.MESSAGE_V1_EVM_SOURCE_BASE_SIZE + 3 * EVM_ADDRESS_LENGTH
        + MessageV1Codec.TOKEN_TRANSFER_V1_EVM_SOURCE_BASE_SIZE
    );

    assertEq(DEFAULT_EXEC_FEE_USD_CENTS, fee);
    assertEq(DEFAULT_EXEC_GAS, gas);
    assertEq(execBytes, expectedBytes);
  }

  function test_getFee_WithExtraArgs() public view {
    Client.CCV[] memory ccvs = new Client.CCV[](1);
    ccvs[0] = Client.CCV({ccvAddress: INITIAL_CCV, args: ""});
    bytes memory extraArgs = "extra_args_data";

    (uint16 fee, uint32 gas, uint32 execBytes) = s_executor.getFee(DEST_CHAIN_SELECTOR, 0, 0, 0, ccvs, extraArgs);

    uint32 expectedBytes =
      uint32(MessageV1Codec.MESSAGE_V1_EVM_SOURCE_BASE_SIZE + 2 * EVM_ADDRESS_LENGTH + extraArgs.length);

    assertEq(DEFAULT_EXEC_FEE_USD_CENTS, fee);
    assertEq(DEFAULT_EXEC_GAS, gas);
    assertEq(execBytes, expectedBytes);
  }

  function test_getFee_RevertWhen_Executor__RequestedBlockDepthTooLow() public {
    uint16 requestedBlockDepth = MIN_BLOCK_CONFIRMATIONS - 1;

    vm.expectRevert(
      abi.encodeWithSelector(
        Executor.Executor__RequestedBlockDepthTooLow.selector, requestedBlockDepth, MIN_BLOCK_CONFIRMATIONS
      )
    );
    s_executor.getFee(DEST_CHAIN_SELECTOR, requestedBlockDepth, 0, 0, new Client.CCV[](1), "");
  }

  function test_getFee_RevertWhen_InvalidDestChain() public {
    Client.CCV[] memory ccvs = new Client.CCV[](1);
    ccvs[0] = Client.CCV({ccvAddress: INITIAL_CCV, args: ""});

    vm.expectRevert(abi.encodeWithSelector(Executor.InvalidDestChain.selector, DEST_CHAIN_SELECTOR + 1));
    s_executor.getFee(DEST_CHAIN_SELECTOR + 1, 0, 0, 0, ccvs, "");
  }

  function test_getFee_RevertWhen_UnsupportedRequiredCCV() public {
    address unsupportedCCV = makeAddr("unsupportedCCV");
    Client.CCV[] memory ccvs = new Client.CCV[](1);
    ccvs[0] = Client.CCV({ccvAddress: unsupportedCCV, args: ""});

    vm.expectRevert(abi.encodeWithSelector(Executor.InvalidCCV.selector, unsupportedCCV));
    s_executor.getFee(DEST_CHAIN_SELECTOR, 0, 0, 0, ccvs, "");
  }

  function test_getFee_RevertWhen_ExceedsMaxCCVs() public {
    Client.CCV[] memory ccvs = new Client.CCV[](2);
    ccvs[0] = Client.CCV({ccvAddress: INITIAL_CCV, args: ""});
    ccvs[1] = Client.CCV({ccvAddress: INITIAL_CCV, args: ""});

    vm.expectRevert(abi.encodeWithSelector(Executor.ExceedsMaxCCVs.selector, ccvs.length, INITIAL_MAX_CCVS));
    s_executor.getFee(DEST_CHAIN_SELECTOR, 0, 0, 0, ccvs, "");
  }
}
