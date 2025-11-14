// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ICrossChainVerifierResolver} from "../../../interfaces/ICrossChainVerifierResolver.sol";
import {ICrossChainVerifierV1} from "../../../interfaces/ICrossChainVerifierV1.sol";
import {IExecutor} from "../../../interfaces/IExecutor.sol";

import {Client} from "../../../libraries/Client.sol";
import {ExtraArgsCodec} from "../../../libraries/ExtraArgsCodec.sol";
import {OnRamp} from "../../../onRamp/OnRamp.sol";
import {OnRampSetup} from "./OnRampSetup.t.sol";

contract OnRamp_getReceipts is OnRampSetup {
  function setUp() public virtual override {
    super.setUp();

    // Mock CCV responses for the default CCV from OnRampSetup.
    vm.mockCall(
      s_defaultCCV,
      abi.encodeWithSelector(ICrossChainVerifierResolver.getOutboundImplementation.selector),
      abi.encode(s_defaultCCV)
    );

    vm.mockCall(
      s_defaultCCV,
      abi.encodeWithSelector(ICrossChainVerifierV1.getFee.selector),
      abi.encode(DEFAULT_CCV_FEE_USD_CENTS, DEFAULT_CCV_GAS_LIMIT, DEFAULT_CCV_PAYLOAD_SIZE)
    );

    vm.mockCall(
      s_defaultExecutor, abi.encodeWithSelector(IExecutor.getFee.selector), abi.encode(DEFAULT_EXEC_FEE_USD_CENTS)
    );
  }

  function test_getReceipts_IncludesExtraArgsGasLimitInGasLimitSum() public view {
    Client.EVM2AnyMessage memory message = _generateEmptyMessage();

    address[] memory ccvs = new address[](1);
    ccvs[0] = s_defaultCCV;

    bytes[] memory ccvArgs = new bytes[](1);
    ccvArgs[0] = "";

    ExtraArgsCodec.GenericExtraArgsV3 memory extraArgs = ExtraArgsCodec.GenericExtraArgsV3({
      ccvs: ccvs,
      ccvArgs: ccvArgs,
      blockConfirmations: 12,
      gasLimit: GAS_LIMIT,
      executor: s_defaultExecutor,
      executorArgs: "",
      tokenReceiver: "",
      tokenArgs: ""
    });

    (OnRamp.Receipt[] memory receipts, uint32 gasLimitSum, uint256 feeUSDCentsSum) =
      s_onRamp.getReceipts(DEST_CHAIN_SELECTOR, message, extraArgs);

    assertEq(receipts.length, 2, "Should have 2 receipts: 1 CCV + 1 executor");

    // Verify CCV receipt.
    assertEq(receipts[0].issuer, s_defaultCCV, "First receipt issuer should be CCV");
    assertEq(receipts[0].destGasLimit, DEFAULT_CCV_GAS_LIMIT, "CCV gas limit should match");

    // Verify executor receipt.
    assertEq(receipts[1].issuer, s_defaultExecutor, "Last receipt issuer should be executor");
    assertEq(receipts[1].destGasLimit, BASE_EXEC_GAS_COST + GAS_LIMIT);
    assertEq(receipts[1].feeTokenAmount, DEFAULT_EXEC_FEE_USD_CENTS);

    // Verify total gasLimitSum includes extraArgs.gasLimit.
    uint32 expectedGasLimitSum = DEFAULT_CCV_GAS_LIMIT + BASE_EXEC_GAS_COST + GAS_LIMIT;

    assertEq(gasLimitSum, expectedGasLimitSum);

    assertEq(feeUSDCentsSum, DEFAULT_CCV_FEE_USD_CENTS + DEFAULT_EXEC_FEE_USD_CENTS);
  }
}
