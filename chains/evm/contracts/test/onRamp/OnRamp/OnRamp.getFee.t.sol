// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ICrossChainVerifierResolver} from "../../../interfaces/ICrossChainVerifierResolver.sol";
import {ICrossChainVerifierV1} from "../../../interfaces/ICrossChainVerifierV1.sol";
import {IExecutor} from "../../../interfaces/IExecutor.sol";

import {Client} from "../../../libraries/Client.sol";
import {ExtraArgsCodec} from "../../../libraries/ExtraArgsCodec.sol";
import {OnRamp} from "../../../onRamp/OnRamp.sol";
import {OnRampSetup} from "./OnRampSetup.t.sol";

contract GasTest is OnRampSetup {
  function setUp() public virtual override {}

  event LogBytes(bytes data);

  function test_gas_abi_encode() public {
    bytes memory extraArgs = bytes.concat(
      ExtraArgsCodec.GENERIC_EXTRA_ARGS_V3_TAG,
      abi.encode(
        ExtraArgsCodec.GenericExtraArgsV3({
          ccvs: new address[](2),
          ccvArgs: new bytes[](2),
          blockConfirmations: 34,
          gasLimit: 59499,
          executor: address(0x1234567890123456789012345678901234567890),
          executorArgs: "3282389428935872359872395885792839273525",
          tokenReceiver: "3282389428935872359872329385792837273525",
          tokenArgs: ""
        })
      )
    );

    vm.pauseGasMetering();
    emit LogBytes(extraArgs);
    vm.resumeGasMetering();
  }

  function test_gas_abi_packed() public {
    bytes memory extraArgs = ExtraArgsCodec._encodeGenericExtraArgsV3(
      ExtraArgsCodec.GenericExtraArgsV3({
        ccvs: new address[](2),
        ccvArgs: new bytes[](2),
        blockConfirmations: 34,
        gasLimit: 59499,
        executor: address(0x1234567890123456789012345678901234567890),
        executorArgs: "3282389428935872359872395885792839273525",
        tokenReceiver: "3282389428935872359872329385792837273525",
        tokenArgs: ""
      })
    );

    vm.pauseGasMetering();
    emit LogBytes(extraArgs);
    vm.resumeGasMetering();
  }

  function test_gas_decode_abi_packed() public {
    vm.pauseGasMetering();
    EncodeDecoder encoderDecoder = new EncodeDecoder();

    bytes memory extraArgs = ExtraArgsCodec._encodeGenericExtraArgsV3(
      ExtraArgsCodec.GenericExtraArgsV3({
        ccvs: new address[](2),
        ccvArgs: new bytes[](2),
        blockConfirmations: 34,
        gasLimit: 59499,
        executor: address(0x1234567890123456789012345678901234567890),
        executorArgs: "3282389428935872359872395885792839273525",
        tokenReceiver: "3282389428935872359872329385792837273525",
        tokenArgs: ""
      })
    );

    vm.resumeGasMetering();

    encoderDecoder.decodePacked(extraArgs);
  }

  function test_gas_decode_abi_encode() public {
    vm.pauseGasMetering();

    EncodeDecoder encoderDecoder = new EncodeDecoder();
    bytes memory extraArgs = bytes.concat(
      ExtraArgsCodec.GENERIC_EXTRA_ARGS_V3_TAG,
      abi.encode(
        ExtraArgsCodec.GenericExtraArgsV3({
          ccvs: new address[](2),
          ccvArgs: new bytes[](2),
          blockConfirmations: 34,
          gasLimit: 59499,
          executor: address(0x1234567890123456789012345678901234567890),
          executorArgs: "3282389428935872359872395885792839273525",
          tokenReceiver: "3282389428935872359872329385792837273525",
          tokenArgs: ""
        })
      )
    );

    vm.resumeGasMetering();

    encoderDecoder.decodeABI(extraArgs);
  }

  function test_gas_decode_empty_abi_encode() public {
    vm.pauseGasMetering();

    EncodeDecoder encoderDecoder = new EncodeDecoder();
    bytes memory extraArgs = bytes.concat(
      ExtraArgsCodec.GENERIC_EXTRA_ARGS_V3_TAG,
      abi.encode(
        ExtraArgsCodec.GenericExtraArgsV3({
          ccvs: new address[](0),
          ccvArgs: new bytes[](0),
          blockConfirmations: 34,
          gasLimit: 59499,
          executor: address(0),
          executorArgs: "",
          tokenReceiver: "",
          tokenArgs: ""
        })
      )
    );

    vm.resumeGasMetering();

    encoderDecoder.decodeABI(extraArgs);
  }

  function test_gas_decode_empty_packed() public {
    vm.pauseGasMetering();

    EncodeDecoder encoderDecoder = new EncodeDecoder();
    bytes memory extraArgs = ExtraArgsCodec._encodeGenericExtraArgsV3(
      ExtraArgsCodec.GenericExtraArgsV3({
        ccvs: new address[](0),
        ccvArgs: new bytes[](0),
        blockConfirmations: 34,
        gasLimit: 59499,
        executor: address(0),
        executorArgs: "",
        tokenReceiver: "",
        tokenArgs: ""
      })
    );

    vm.resumeGasMetering();

    encoderDecoder.decodePacked(extraArgs);
  }
}

contract EncodeDecoder {
  function decodePacked(
    bytes calldata extraArgs
  ) external pure {
    ExtraArgsCodec._decodeGenericExtraArgsV3(extraArgs);
  }

  function decodeABI(
    bytes calldata extraArgs
  ) external pure {
    abi.decode(extraArgs[4:], (ExtraArgsCodec.GenericExtraArgsV3));
  }
}

contract OnRamp_getFee is OnRampSetup {
  uint16 internal constant MOCKED_DEFAULT_CCV_FEE_USD_CENTS = 5_00;
  uint16 internal constant MOCKED_DEFAULT_EXECUTOR_FEE_USD_CENTS = 4_25;

  function setUp() public virtual override {
    super.setUp();

    _mockVerifierFee(s_defaultCCV, MOCKED_DEFAULT_CCV_FEE_USD_CENTS, DEFAULT_CCV_GAS_LIMIT, DEFAULT_CCV_PAYLOAD_SIZE);
    _mockExecutorFee(s_defaultExecutor, MOCKED_DEFAULT_EXECUTOR_FEE_USD_CENTS, 0, 0);
  }

  function _mockVerifierFee(
    address verifier,
    uint16 feeUSDCents,
    uint64 gasForVerification,
    uint32 payloadSizeBytes
  ) internal {
    vm.mockCall(
      verifier,
      abi.encodeWithSelector(ICrossChainVerifierResolver.getOutboundImplementation.selector, DEST_CHAIN_SELECTOR),
      abi.encode(verifier)
    );
    vm.mockCall(
      verifier,
      abi.encodeWithSelector(ICrossChainVerifierV1.getFee.selector),
      abi.encode(feeUSDCents, gasForVerification, payloadSizeBytes)
    );
  }

  function _mockExecutorFee(
    address executor,
    uint16 feeUSDCents,
    uint64 gasForVerification,
    uint32 payloadSizeBytes
  ) internal {
    vm.mockCall(
      executor,
      abi.encodeWithSelector(IExecutor.getFee.selector),
      abi.encode(feeUSDCents, gasForVerification, payloadSizeBytes)
    );
  }

  function test_getFee_WithV3ExtraArgs_CustomCCV_SkipsDefaults() public {
    address newVerifier = makeAddr("custom_verifier");
    uint16 differentFee = 123_45;
    _mockVerifierFee(newVerifier, differentFee, DEFAULT_CCV_GAS_LIMIT, DEFAULT_CCV_PAYLOAD_SIZE);

    address[] memory ccvAddresses = new address[](1);
    ccvAddresses[0] = newVerifier;
    bytes[] memory ccvArgs = new bytes[](1);
    ccvArgs[0] = "";

    Client.EVM2AnyMessage memory message = _generateEmptyMessage();
    message.extraArgs = ExtraArgsCodec._encodeGenericExtraArgsV3(_createV3ExtraArgs(ccvAddresses, ccvArgs));

    uint256 feeAmount = s_onRamp.getFee(DEST_CHAIN_SELECTOR, message);

    assertEq(differentFee + MOCKED_DEFAULT_EXECUTOR_FEE_USD_CENTS, feeAmount);
  }

  function test_getFee_WithLaneMandatedCCVs() public {
    address mandatedVerifier = makeAddr("mandated_verifier");
    uint16 mandatedFee = 1_50;

    _mockVerifierFee(mandatedVerifier, mandatedFee, 0, 0);

    address[] memory laneMandatedCCVs = new address[](1);
    laneMandatedCCVs[0] = mandatedVerifier;

    address[] memory defaultCCVs = new address[](1);
    defaultCCVs[0] = s_defaultCCV;

    OnRamp.DestChainConfigArgs[] memory destChainConfigArgs = new OnRamp.DestChainConfigArgs[](1);
    destChainConfigArgs[0] = OnRamp.DestChainConfigArgs({
      destChainSelector: DEST_CHAIN_SELECTOR,
      router: s_sourceRouter,
      addressBytesLength: EVM_ADDRESS_LENGTH,
      baseExecutionGasCost: BASE_EXEC_GAS_COST,
      laneMandatedCCVs: laneMandatedCCVs,
      defaultCCVs: defaultCCVs,
      defaultExecutor: s_defaultExecutor,
      offRamp: abi.encodePacked(address(s_offRampOnRemoteChain))
    });

    s_onRamp.applyDestChainConfigUpdates(destChainConfigArgs);

    Client.EVM2AnyMessage memory message = _generateEmptyMessage();
    uint256 feeAmount = s_onRamp.getFee(DEST_CHAIN_SELECTOR, message);

    assertEq(MOCKED_DEFAULT_CCV_FEE_USD_CENTS + MOCKED_DEFAULT_EXECUTOR_FEE_USD_CENTS + mandatedFee, feeAmount);
  }

  function test_getFee_WithCustomExecutorAndCCVs() public {
    address customExecutor = makeAddr("custom_executor_2");
    address verifier = makeAddr("verifier_with_executor");

    uint16 differentExecutorFee = 300;
    uint16 differentVerifierFee = 200;

    _mockExecutorFee(customExecutor, differentExecutorFee, 0, 0);
    _mockVerifierFee(verifier, differentVerifierFee, 0, 0);

    address[] memory ccvAddresses = new address[](1);
    ccvAddresses[0] = verifier;
    bytes[] memory ccvArgs = new bytes[](1);
    ccvArgs[0] = "";

    ExtraArgsCodec.GenericExtraArgsV3 memory extraArgsV3 = ExtraArgsCodec.GenericExtraArgsV3({
      ccvs: ccvAddresses,
      ccvArgs: ccvArgs,
      blockConfirmations: 12,
      gasLimit: GAS_LIMIT,
      executor: customExecutor,
      executorArgs: "",
      tokenReceiver: "",
      tokenArgs: ""
    });

    Client.EVM2AnyMessage memory message = _generateEmptyMessage();
    message.extraArgs = ExtraArgsCodec._encodeGenericExtraArgsV3(extraArgsV3);

    uint256 feeAmount = s_onRamp.getFee(DEST_CHAIN_SELECTOR, message);

    assertEq(differentExecutorFee + differentVerifierFee, feeAmount);
  }

  // Reverts

  function test_getFee_RevertWhen_InvalidDestChainSelector() public {
    uint64 invalidChainSelector = 999999;

    vm.expectRevert(abi.encodeWithSelector(OnRamp.DestinationChainNotSupported.selector, invalidChainSelector));
    s_onRamp.getFee(invalidChainSelector, _generateEmptyMessage());
  }

  function test_getFee_RevertWhen_DestinationChainNotSupportedByCCV() public {
    address verifier = makeAddr("verifier_no_support");

    _mockVerifierFee(verifier, 100, 0, 0);
    // Mock to return address(0) to simulate no support.
    vm.mockCall(
      verifier,
      abi.encodeWithSelector(ICrossChainVerifierResolver.getOutboundImplementation.selector, DEST_CHAIN_SELECTOR),
      abi.encode(address(0))
    );

    address[] memory ccvs = new address[](1);
    ccvs[0] = verifier;

    Client.EVM2AnyMessage memory message = _generateEmptyMessage();
    message.extraArgs = ExtraArgsCodec._encodeGenericExtraArgsV3(_createV3ExtraArgs(ccvs, new bytes[](1)));

    vm.expectRevert(
      abi.encodeWithSelector(OnRamp.DestinationChainNotSupportedByCCV.selector, verifier, DEST_CHAIN_SELECTOR)
    );
    s_onRamp.getFee(DEST_CHAIN_SELECTOR, message);
  }
}
