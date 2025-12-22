// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {CCVConfigValidation} from "../../../libraries/CCVConfigValidation.sol";
import {Client} from "../../../libraries/Client.sol";
import {ExtraArgsCodec} from "../../../libraries/ExtraArgsCodec.sol";
import {OnRamp} from "../../../onRamp/OnRamp.sol";
import {OnRampHelper} from "../../helpers/OnRampHelper.sol";
import {OnRampSetup} from "./OnRampSetup.t.sol";

contract OnRamp_parseExtraArgsWithDefaults is OnRampSetup {
  OnRampHelper internal s_OnRampHelper;

  address[] internal s_defaultCCVs;
  address[] internal s_laneMandatedCCVs;
  OnRamp.DestChainConfig internal s_destChainConfig;

  function setUp() public override {
    super.setUp();
    s_OnRampHelper = new OnRampHelper(
      OnRamp.StaticConfig({
        chainSelector: SOURCE_CHAIN_SELECTOR,
        rmnRemote: s_mockRMNRemote,
        tokenAdminRegistry: address(s_tokenAdminRegistry)
      }),
      OnRamp.DynamicConfig({
        feeQuoter: address(s_feeQuoter), reentrancyGuardEntered: false, feeAggregator: FEE_AGGREGATOR
      })
    );

    // Initialize default test configuration
    s_defaultCCVs = new address[](2);
    s_defaultCCVs[0] = makeAddr("defaultCCV1");
    s_defaultCCVs[1] = makeAddr("defaultCCV2");

    s_laneMandatedCCVs = new address[](1);
    s_laneMandatedCCVs[0] = makeAddr("mandatedCCV1");

    s_destChainConfig = OnRamp.DestChainConfig({
      router: s_sourceRouter,
      messageNumber: 0,
      addressBytesLength: EVM_ADDRESS_LENGTH,
      networkFeeUSDCents: NETWORK_FEE_USD_CENTS,
      tokenReceiverAllowed: true,
      baseExecutionGasCost: BASE_EXEC_GAS_COST,
      defaultExecutor: s_defaultExecutor,
      laneMandatedCCVs: s_laneMandatedCCVs,
      defaultCCVs: s_defaultCCVs,
      offRamp: abi.encodePacked(address(s_offRampOnRemoteChain))
    });
  }

  function test_parseExtraArgsWithDefaults_V3WithUserProvidedCCVs() public {
    address[] memory userCCVAddresses = new address[](1);
    userCCVAddresses[0] = makeAddr("userCCV1");
    bytes[] memory userCCVArgs = new bytes[](1);
    userCCVArgs[0] = "userArgs";

    ExtraArgsCodec.GenericExtraArgsV3 memory inputArgs = ExtraArgsCodec.GenericExtraArgsV3({
      ccvs: userCCVAddresses,
      ccvArgs: userCCVArgs,
      blockConfirmations: 0,
      gasLimit: GAS_LIMIT,
      executor: makeAddr("userExecutor"),
      executorArgs: "execArgs",
      tokenReceiver: abi.encodePacked(makeAddr("tokenReceiver")),
      tokenArgs: "tokenArgs"
    });

    bytes memory extraArgs = ExtraArgsCodec._encodeGenericExtraArgsV3(inputArgs);

    ExtraArgsCodec.GenericExtraArgsV3 memory result =
      s_OnRampHelper.parseExtraArgsWithDefaults(DEST_CHAIN_SELECTOR, s_destChainConfig, extraArgs, false);

    // User-provided CCVs should be used (no lane mandated CCVs added in parseExtraArgsWithDefaults anymore)
    assertEq(userCCVAddresses.length, result.ccvs.length);
    assertEq(userCCVAddresses[0], result.ccvs[0]);
    assertEq(userCCVArgs[0], result.ccvArgs[0]);
    assertEq(inputArgs.executor, result.executor);
  }

  function test_parseExtraArgsWithDefaults_V3WithEmptyRequiredCCVs() public view {
    ExtraArgsCodec.GenericExtraArgsV3 memory inputArgs = _createV3ExtraArgs(new address[](0), new bytes[](0));

    bytes memory extraArgs = ExtraArgsCodec._encodeGenericExtraArgsV3(inputArgs);

    ExtraArgsCodec.GenericExtraArgsV3 memory result =
      s_OnRampHelper.parseExtraArgsWithDefaults(DEST_CHAIN_SELECTOR, s_destChainConfig, extraArgs, false);

    // Default CCVs should be applied (no lane mandated CCVs added in parseExtraArgsWithDefaults anymore)
    assertEq(s_defaultCCVs.length, result.ccvs.length);
    assertEq(s_defaultCCVs[0], result.ccvs[0]);
    assertEq("", result.ccvArgs[0]); // Default CCVs have empty args.
    assertEq(s_defaultCCVs[1], result.ccvs[1]);
    assertEq("", result.ccvArgs[1]); // Default CCVs have empty args.

    // Default executor should be applied.
    assertEq(s_defaultExecutor, result.executor);
  }

  // TODO Sui/SVM
  function test_parseExtraArgsWithDefaults_OldExtraArgs() public view {
    // Use GenericExtraArgsV2 format.
    uint256 gasLimit = 300_000;
    Client.GenericExtraArgsV2 memory v2Args =
      Client.GenericExtraArgsV2({gasLimit: gasLimit, allowOutOfOrderExecution: true});

    bytes memory legacyExtraArgs = Client._argsToBytes(v2Args);

    ExtraArgsCodec.GenericExtraArgsV3 memory result =
      s_OnRampHelper.parseExtraArgsWithDefaults(DEST_CHAIN_SELECTOR, s_destChainConfig, legacyExtraArgs, false);

    assertEq(s_defaultCCVs.length, result.ccvs.length);
    assertEq(result.ccvs[0], s_defaultCCVs[0]);
    assertEq(result.ccvArgs[0], "", "ccv args 0");
    assertEq(result.ccvs[1], s_defaultCCVs[1]);
    assertEq(result.ccvArgs[1], "", "ccv args 1");

    assertEq(result.executorArgs, "");
    assertEq(result.executor, s_defaultExecutor);
  }

  // Additional test for defaults when no user CCVs provided
  function test_parseExtraArgsWithDefaults_DefaultCCVsAlwaysPresent() public view {
    // Ensure defaultCCVs.length > 0 (invariant)
    assertTrue(s_destChainConfig.defaultCCVs.length > 0, "defaultCCVs must not be empty");

    // Test with empty user input
    ExtraArgsCodec.GenericExtraArgsV3 memory inputArgs = _createV3ExtraArgs(new address[](0), new bytes[](0));

    bytes memory extraArgs = ExtraArgsCodec._encodeGenericExtraArgsV3(inputArgs);
    ExtraArgsCodec.GenericExtraArgsV3 memory result =
      s_OnRampHelper.parseExtraArgsWithDefaults(DEST_CHAIN_SELECTOR, s_destChainConfig, extraArgs, false);

    // Should have default CCVs
    assertEq(s_defaultCCVs.length, result.ccvs.length);
  }

  function test_parseExtraArgsWithDefaults_V3DoesNotAddDefaults_IsTokenTransferWithoutDataAndGasLimitZero()
    public
    view
  {
    ExtraArgsCodec.GenericExtraArgsV3 memory inputArgs = _createV3ExtraArgs(new address[](0), new bytes[](0));
    inputArgs.gasLimit = 0;

    bytes memory extraArgs = ExtraArgsCodec._encodeGenericExtraArgsV3(inputArgs);
    ExtraArgsCodec.GenericExtraArgsV3 memory result =
      s_OnRampHelper.parseExtraArgsWithDefaults(DEST_CHAIN_SELECTOR, s_destChainConfig, extraArgs, true);

    assertEq(result.ccvs.length, 0, "Should not inject default CCVs for token-only transfer");
    assertEq(result.ccvArgs.length, 0, "Should not inject default CCV args for token-only transfer");
  }

  function test_parseExtraArgsWithDefaults_LegacyDoesNotAddDefaults_IsTokenTransferWithoutDataAndGasLimitZero()
    public
    view
  {
    // Legacy GenericExtraArgsV2 can explicitly set gasLimit=0 (unlike empty extraArgs which uses defaults).
    bytes memory legacyExtraArgs =
      Client._argsToBytes(Client.GenericExtraArgsV2({gasLimit: 0, allowOutOfOrderExecution: false}));

    ExtraArgsCodec.GenericExtraArgsV3 memory result =
      s_OnRampHelper.parseExtraArgsWithDefaults(DEST_CHAIN_SELECTOR, s_destChainConfig, legacyExtraArgs, true);

    assertEq(result.ccvs.length, 0, "Should not inject default CCVs for token-only transfer");
    assertEq(result.ccvArgs.length, 0, "Should not inject default CCV args for token-only transfer");
  }

  function test_parseExtraArgsWithDefaults_RevertWhen_TokenReceiverNotAllowed() public {
    OnRamp.DestChainConfig memory cfg = s_destChainConfig;
    cfg.tokenReceiverAllowed = false;

    ExtraArgsCodec.GenericExtraArgsV3 memory inputArgs = _createV3ExtraArgs(new address[](0), new bytes[](0));
    inputArgs.tokenReceiver = abi.encodePacked(makeAddr("tokenReceiver"));
    bytes memory extraArgs = ExtraArgsCodec._encodeGenericExtraArgsV3(inputArgs);

    vm.expectRevert(abi.encodeWithSelector(OnRamp.TokenReceiverNotAllowed.selector, DEST_CHAIN_SELECTOR));
    s_OnRampHelper.parseExtraArgsWithDefaults(DEST_CHAIN_SELECTOR, cfg, extraArgs, false);
  }

  // Reverts

  function test_parseExtraArgsWithDefaults_RevertWhen_NoDuplicatesAllowed_WithinRequiredCCVs() public {
    // Create user-provided CCVs with duplicates in required list
    address duplicateCCV = makeAddr("duplicateCCV");
    address[] memory userCCVAddresses = new address[](2);
    userCCVAddresses[0] = duplicateCCV;
    userCCVAddresses[1] = duplicateCCV; // Duplicate
    bytes[] memory userCCVArgs = new bytes[](2);
    userCCVArgs[0] = "args1";
    userCCVArgs[1] = "args2";

    ExtraArgsCodec.GenericExtraArgsV3 memory inputArgs = _createV3ExtraArgs(userCCVAddresses, userCCVArgs);

    bytes memory extraArgs = ExtraArgsCodec._encodeGenericExtraArgsV3(inputArgs);

    // Should revert due to duplicate CCVs
    vm.expectRevert(abi.encodeWithSelector(CCVConfigValidation.DuplicateCCVNotAllowed.selector, duplicateCCV));
    s_OnRampHelper.parseExtraArgsWithDefaults(DEST_CHAIN_SELECTOR, s_destChainConfig, extraArgs, false);
  }
}
