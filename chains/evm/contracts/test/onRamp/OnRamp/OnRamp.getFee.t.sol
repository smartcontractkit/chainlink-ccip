// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ICrossChainVerifierV1} from "../../../interfaces/ICrossChainVerifierV1.sol";
import {IExecutor} from "../../../interfaces/IExecutor.sol";

import {Client} from "../../../libraries/Client.sol";
import {OnRamp} from "../../../onRamp/OnRamp.sol";
import {OnRampSetup} from "./OnRampSetup.t.sol";

contract OnRamp_getFee is OnRampSetup {
  uint256 internal constant MOCKED_DEFAULT_CCV_FEE = 1e18;
  uint256 internal constant MOCKED_DEFAULT_EXECUTOR_FEE = 2e20;

  function setUp() public virtual override {
    super.setUp();

    _mockVerifierFee(s_defaultCCV, MOCKED_DEFAULT_CCV_FEE, DEFAULT_CCV_GAS_LIMIT, DEFAULT_CCV_PAYLOAD_SIZE);
    _mockExecutorFee(s_defaultExecutor, MOCKED_DEFAULT_EXECUTOR_FEE);
  }

  function _mockVerifierFee(
    address verifier,
    uint256 feeUSDCents,
    uint32 gasForVerification,
    uint32 payloadSizeBytes
  ) internal {
    vm.mockCall(
      verifier,
      abi.encodeWithSelector(ICrossChainVerifierV1.getFee.selector),
      abi.encode(feeUSDCents, gasForVerification, payloadSizeBytes)
    );
  }

  function _mockExecutorFee(address executor, uint256 fee) internal {
    vm.mockCall(executor, abi.encodeWithSelector(IExecutor.getFee.selector), abi.encode(fee));
  }

  function test_getFee_WithV3ExtraArgs_EmptyCCVs_UsesDefaults() public view {
    // When no CCVs are provided in V3 extra args, default CCVs should be used.

    Client.CCV[] memory ccvs = new Client.CCV[](0);
    Client.EVMExtraArgsV3 memory extraArgsV3 = _createV3ExtraArgs(ccvs);

    Client.EVM2AnyMessage memory message = _generateEmptyMessage();
    message.extraArgs = abi.encodePacked(Client.GENERIC_EXTRA_ARGS_V3_TAG, abi.encode(extraArgsV3));

    uint256 feeAmount = s_onRamp.getFee(DEST_CHAIN_SELECTOR, message);

    // Should use default CCV + executor.
    assertEq(MOCKED_DEFAULT_CCV_FEE + MOCKED_DEFAULT_EXECUTOR_FEE, feeAmount);
  }

  function test_getFee_WithV3ExtraArgs_CustomCCV_SkipsDefaults() public {
    address newVerifier = makeAddr("custom_verifier");
    uint256 differentFee = 12345;
    _mockVerifierFee(newVerifier, differentFee, DEFAULT_CCV_GAS_LIMIT, DEFAULT_CCV_PAYLOAD_SIZE);

    Client.CCV[] memory ccvs = new Client.CCV[](1);
    ccvs[0] = Client.CCV({ccvAddress: newVerifier, args: ""});

    Client.EVM2AnyMessage memory message = _generateEmptyMessage();
    message.extraArgs = abi.encodePacked(Client.GENERIC_EXTRA_ARGS_V3_TAG, abi.encode(_createV3ExtraArgs(ccvs)));

    uint256 feeAmount = s_onRamp.getFee(DEST_CHAIN_SELECTOR, message);

    assertEq(differentFee + MOCKED_DEFAULT_EXECUTOR_FEE, feeAmount);
  }

  function test_getFee_WithLaneMandatedCCVs() public {
    address mandatedVerifier = makeAddr("mandated_verifier");
    uint256 mandatedFee = 150;

    _mockVerifierFee(mandatedVerifier, mandatedFee, 0, 0);

    address[] memory laneMandatedCCVs = new address[](1);
    laneMandatedCCVs[0] = mandatedVerifier;

    address[] memory defaultCCVs = new address[](1);
    defaultCCVs[0] = s_defaultCCV;

    OnRamp.DestChainConfigArgs[] memory destChainConfigArgs = new OnRamp.DestChainConfigArgs[](1);
    destChainConfigArgs[0] = OnRamp.DestChainConfigArgs({
      destChainSelector: DEST_CHAIN_SELECTOR,
      router: s_sourceRouter,
      laneMandatedCCVs: laneMandatedCCVs,
      defaultCCVs: defaultCCVs,
      defaultExecutor: s_defaultExecutor,
      offRamp: abi.encodePacked(address(s_offRampOnRemoteChain))
    });

    s_onRamp.applyDestChainConfigUpdates(destChainConfigArgs);

    Client.EVM2AnyMessage memory message = _generateEmptyMessage();
    uint256 feeAmount = s_onRamp.getFee(DEST_CHAIN_SELECTOR, message);

    assertEq(MOCKED_DEFAULT_CCV_FEE + MOCKED_DEFAULT_EXECUTOR_FEE + mandatedFee, feeAmount);
  }

  function test_getFee_WithCustomExecutorAndCCVs() public {
    address customExecutor = makeAddr("custom_executor_2");
    address verifier = makeAddr("verifier_with_executor");

    uint256 differentExecutorFee = 300;
    uint256 differentVerifierFee = 200;

    _mockExecutorFee(customExecutor, differentExecutorFee);
    _mockVerifierFee(verifier, differentVerifierFee, 0, 0);

    Client.CCV[] memory ccvs = new Client.CCV[](1);
    ccvs[0] = Client.CCV({ccvAddress: verifier, args: ""});

    Client.EVMExtraArgsV3 memory extraArgsV3 =
      Client.EVMExtraArgsV3({ccvs: ccvs, finalityConfig: 12, executor: customExecutor, executorArgs: "", tokenArgs: ""});

    Client.EVM2AnyMessage memory message = _generateEmptyMessage();
    message.extraArgs = abi.encodePacked(Client.GENERIC_EXTRA_ARGS_V3_TAG, abi.encode(extraArgsV3));

    uint256 feeAmount = s_onRamp.getFee(DEST_CHAIN_SELECTOR, message);

    assertEq(differentExecutorFee + differentVerifierFee, feeAmount);
  }

  // Reverts

  function test_getFee_RevertWhen_InvalidDestChainSelector() public {
    uint64 invalidChainSelector = 999999;

    vm.expectRevert(abi.encodeWithSelector(OnRamp.DestinationChainNotSupported.selector, invalidChainSelector));
    s_onRamp.getFee(invalidChainSelector, _generateEmptyMessage());
  }
}
