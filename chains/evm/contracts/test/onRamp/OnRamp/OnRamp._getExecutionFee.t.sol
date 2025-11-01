// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IExecutor} from "../../../interfaces/IExecutor.sol";

import {Client} from "../../../libraries/Client.sol";
import {MessageV1Codec} from "../../../libraries/MessageV1Codec.sol";
import {OnRamp} from "../../../onRamp/OnRamp.sol";
import {OnRampTestHelper} from "../../helpers/OnRampTestHelper.sol";
import {MockExecutor} from "../../mocks/MockExecutor.sol";
import {OnRampSetup} from "./OnRampSetup.t.sol";

contract OnRamp_getExecutionFee is OnRampSetup {
  OnRampTestHelper internal s_onRampTestHelper;
  address internal s_customExecutor;

  uint16 internal constant EXECUTOR_FEE_USD_CENTS = 123;

  function setUp() public virtual override {
    super.setUp();

    s_onRampTestHelper = new OnRampTestHelper(
      OnRamp.StaticConfig({
        chainSelector: SOURCE_CHAIN_SELECTOR,
        rmnRemote: s_mockRMNRemote,
        tokenAdminRegistry: address(s_tokenAdminRegistry)
      }),
      OnRamp.DynamicConfig({
        feeQuoter: address(s_feeQuoter),
        reentrancyGuardEntered: false,
        feeAggregator: FEE_AGGREGATOR
      })
    );

    s_customExecutor = address(new MockExecutor());

    address[] memory defaultCCVs = new address[](1);
    defaultCCVs[0] = s_defaultCCV;

    OnRamp.DestChainConfigArgs[] memory destChainConfigArgs = new OnRamp.DestChainConfigArgs[](1);
    destChainConfigArgs[0] = OnRamp.DestChainConfigArgs({
      destChainSelector: DEST_CHAIN_SELECTOR,
      router: s_sourceRouter,
      addressBytesLength: EVM_ADDRESS_LENGTH,
      baseExecutionGasCost: BASE_EXEC_GAS_COST,
      laneMandatedCCVs: new address[](0),
      defaultCCVs: defaultCCVs,
      defaultExecutor: s_defaultExecutor,
      offRamp: abi.encodePacked(address(s_offRampOnRemoteChain))
    });

    s_onRampTestHelper.applyDestChainConfigUpdates(destChainConfigArgs);

    // Mock executor fee
    vm.mockCall(s_customExecutor, abi.encodeWithSelector(IExecutor.getFee.selector), abi.encode(EXECUTOR_FEE_USD_CENTS));
  }

  function test_getExecutionFee_WithExecutor() public view {
    Client.GenericExtraArgsV3 memory extraArgs = Client.GenericExtraArgsV3({
      ccvs: new Client.CCV[](0),
      finalityConfig: 12,
      gasLimit: GAS_LIMIT,
      executor: s_customExecutor,
      executorArgs: "",
      tokenReceiver: "",
      tokenArgs: "args"
    });

    OnRamp.Receipt memory receipt = s_onRampTestHelper.getExecutionFee(DEST_CHAIN_SELECTOR, 0, 0, extraArgs);

    assertEq(receipt.issuer, s_customExecutor, "Issuer should be the executor");
    assertEq(receipt.destGasLimit, BASE_EXEC_GAS_COST + GAS_LIMIT, "Gas limit should include base cost");
    assertEq(receipt.feeTokenAmount, EXECUTOR_FEE_USD_CENTS, "Fee should match executor fee");
    assertEq(receipt.extraArgs, extraArgs.executorArgs, "Extra args should match");
  }

  function test_getExecutionFee_NoExecutor() public view {
    Client.GenericExtraArgsV3 memory extraArgs = Client.GenericExtraArgsV3({
      ccvs: new Client.CCV[](0),
      finalityConfig: 12,
      gasLimit: GAS_LIMIT,
      executor: Client.NO_EXECUTION_ADDRESS,
      executorArgs: "",
      tokenReceiver: "",
      tokenArgs: "args with no executor"
    });

    OnRamp.Receipt memory receipt = s_onRampTestHelper.getExecutionFee(DEST_CHAIN_SELECTOR, 0, 0, extraArgs);

    assertEq(receipt.issuer, Client.NO_EXECUTION_ADDRESS, "Issuer should be NO_EXECUTION_ADDRESS");
    assertEq(receipt.destGasLimit, BASE_EXEC_GAS_COST + GAS_LIMIT, "Gas limit should still include base cost");
    assertEq(receipt.feeTokenAmount, 0, "Fee should be zero for NO_EXECUTION_ADDRESS");
    assertEq(receipt.extraArgs, extraArgs.executorArgs, "Extra args should match");
  }

  function test_getExecutionFee_CalculatesDestBytesOverhead_WithTokens() public view {
    uint256 dataLength = 500;
    uint256 numberOfTokens = 1;
    Client.GenericExtraArgsV3 memory extraArgs = Client.GenericExtraArgsV3({
      ccvs: new Client.CCV[](0),
      finalityConfig: 12,
      gasLimit: GAS_LIMIT,
      executor: s_customExecutor,
      executorArgs: "",
      tokenReceiver: "",
      tokenArgs: ""
    });

    OnRamp.Receipt memory receipt =
      s_onRampTestHelper.getExecutionFee(DEST_CHAIN_SELECTOR, dataLength, numberOfTokens, extraArgs);

    uint32 expectedOverhead = uint32(
      MessageV1Codec.MESSAGE_V1_EVM_SOURCE_BASE_SIZE + dataLength
        + (MessageV1Codec.MESSAGE_V1_REMOTE_CHAIN_ADDRESSES * EVM_ADDRESS_LENGTH)
        + (numberOfTokens * (MessageV1Codec.TOKEN_TRANSFER_V1_EVM_SOURCE_BASE_SIZE + EVM_ADDRESS_LENGTH))
    );

    assertEq(receipt.destBytesOverhead, expectedOverhead, "Bytes overhead should include token overhead");
  }

  function test_getExecutionFee_CalculatesDestBytesOverhead_WithExecutorArgs() public view {
    uint256 dataLength = 500;
    bytes memory executorArgs = new bytes(200);
    Client.GenericExtraArgsV3 memory extraArgs = Client.GenericExtraArgsV3({
      ccvs: new Client.CCV[](0),
      finalityConfig: 12,
      gasLimit: GAS_LIMIT,
      executor: s_customExecutor,
      executorArgs: executorArgs,
      tokenReceiver: "",
      tokenArgs: ""
    });

    OnRamp.Receipt memory receipt = s_onRampTestHelper.getExecutionFee(DEST_CHAIN_SELECTOR, dataLength, 0, extraArgs);

    uint32 expectedOverhead = uint32(
      MessageV1Codec.MESSAGE_V1_EVM_SOURCE_BASE_SIZE + dataLength + executorArgs.length
        + (MessageV1Codec.MESSAGE_V1_REMOTE_CHAIN_ADDRESSES * EVM_ADDRESS_LENGTH)
    );

    assertEq(receipt.destBytesOverhead, expectedOverhead, "Bytes overhead should include executor args length");
  }

  function test_getExecutionFee_ZeroGasLimit() public view {
    Client.GenericExtraArgsV3 memory extraArgs = Client.GenericExtraArgsV3({
      ccvs: new Client.CCV[](0),
      finalityConfig: 12,
      gasLimit: 0,
      executor: s_customExecutor,
      executorArgs: "",
      tokenReceiver: "",
      tokenArgs: ""
    });

    OnRamp.Receipt memory receipt = s_onRampTestHelper.getExecutionFee(DEST_CHAIN_SELECTOR, 100, 0, extraArgs);

    assertEq(receipt.destGasLimit, BASE_EXEC_GAS_COST, "Gas limit should only be base cost when user gas limit is 0");
  }
}
