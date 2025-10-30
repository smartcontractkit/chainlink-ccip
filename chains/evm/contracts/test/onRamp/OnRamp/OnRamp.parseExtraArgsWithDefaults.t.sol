// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {CCVConfigValidation} from "../../../libraries/CCVConfigValidation.sol";
import {Client} from "../../../libraries/Client.sol";
import {OnRamp} from "../../../onRamp/OnRamp.sol";
import {OnRampTestHelper} from "../../helpers/OnRampTestHelper.sol";
import {OnRampSetup} from "./OnRampSetup.t.sol";

contract OnRamp_parseExtraArgsWithDefaults is OnRampSetup {
  OnRampTestHelper internal s_onRampTestHelper;

  address[] internal s_defaultCCVs;
  address[] internal s_laneMandatedCCVs;
  OnRamp.DestChainConfig internal s_destChainConfig;

  function setUp() public override {
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

    // Initialize default test configuration
    s_defaultCCVs = new address[](2);
    s_defaultCCVs[0] = makeAddr("defaultCCV1");
    s_defaultCCVs[1] = makeAddr("defaultCCV2");

    s_laneMandatedCCVs = new address[](1);
    s_laneMandatedCCVs[0] = makeAddr("mandatedCCV1");

    s_destChainConfig = OnRamp.DestChainConfig({
      router: s_sourceRouter,
      sequenceNumber: 0,
      addressBytesLength: EVM_ADDRESS_LENGTH,
      defaultExecutor: s_defaultExecutor,
      laneMandatedCCVs: s_laneMandatedCCVs,
      defaultCCVs: s_defaultCCVs,
      offRamp: abi.encodePacked(address(s_offRampOnRemoteChain))
    });
  }

  function test_parseExtraArgsWithDefaults_V3WithUserProvidedCCVs() public {
    Client.CCV[] memory userRequiredCCVs = new Client.CCV[](1);
    userRequiredCCVs[0] = Client.CCV({ccvAddress: makeAddr("userCCV1"), args: "userArgs"});

    Client.GenericExtraArgsV3 memory inputArgs = Client.GenericExtraArgsV3({
      ccvs: userRequiredCCVs,
      finalityConfig: 0,
      gasLimit: GAS_LIMIT,
      executor: makeAddr("userExecutor"),
      executorArgs: "execArgs",
      tokenReceiver: abi.encodePacked(makeAddr("tokenReceiver")),
      tokenArgs: "tokenArgs"
    });

    bytes memory extraArgs = abi.encodePacked(Client.GENERIC_EXTRA_ARGS_V3_TAG, abi.encode(inputArgs));

    Client.GenericExtraArgsV3 memory result =
      s_onRampTestHelper.parseExtraArgsWithDefaults(DEST_CHAIN_SELECTOR, s_destChainConfig, extraArgs);

    // User-provided CCVs should be used (no lane mandated CCVs added in parseExtraArgsWithDefaults anymore)
    assertEq(userRequiredCCVs.length, result.ccvs.length);
    assertEq(userRequiredCCVs[0].ccvAddress, result.ccvs[0].ccvAddress);
    assertEq(userRequiredCCVs[0].args, result.ccvs[0].args);
    assertEq(inputArgs.executor, result.executor);
  }

  function test_parseExtraArgsWithDefaults_V3WithEmptyRequiredCCVs() public view {
    Client.GenericExtraArgsV3 memory inputArgs = _createV3ExtraArgs(new Client.CCV[](0));

    bytes memory extraArgs = abi.encodePacked(Client.GENERIC_EXTRA_ARGS_V3_TAG, abi.encode(inputArgs));

    Client.GenericExtraArgsV3 memory result =
      s_onRampTestHelper.parseExtraArgsWithDefaults(DEST_CHAIN_SELECTOR, s_destChainConfig, extraArgs);

    // Default CCVs should be applied (no lane mandated CCVs added in parseExtraArgsWithDefaults anymore)
    assertEq(s_defaultCCVs.length, result.ccvs.length);
    assertEq(s_defaultCCVs[0], result.ccvs[0].ccvAddress);
    assertEq("", result.ccvs[0].args); // Default CCVs have empty args.
    assertEq(s_defaultCCVs[1], result.ccvs[1].ccvAddress);
    assertEq("", result.ccvs[1].args); // Default CCVs have empty args.

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

    Client.GenericExtraArgsV3 memory result =
      s_onRampTestHelper.parseExtraArgsWithDefaults(DEST_CHAIN_SELECTOR, s_destChainConfig, legacyExtraArgs);

    assertEq(s_defaultCCVs.length, result.ccvs.length);
    assertEq(result.ccvs[0].ccvAddress, s_defaultCCVs[0]);
    assertEq(result.ccvs[0].args, "", "ccv args 0");
    assertEq(result.ccvs[1].ccvAddress, s_defaultCCVs[1]);
    assertEq(result.ccvs[1].args, "", "ccv args 1");

    assertEq(result.executorArgs, "");
    assertEq(result.executor, s_defaultExecutor);
  }

  // Additional test for defaults when no user CCVs provided
  function test_parseExtraArgsWithDefaults_DefaultCCVsAlwaysPresent() public view {
    // Ensure defaultCCVs.length > 0 (invariant)
    assertTrue(s_destChainConfig.defaultCCVs.length > 0, "defaultCCVs must not be empty");

    // Test with empty user input
    Client.GenericExtraArgsV3 memory inputArgs = _createV3ExtraArgs(new Client.CCV[](0));

    bytes memory extraArgs = abi.encodePacked(Client.GENERIC_EXTRA_ARGS_V3_TAG, abi.encode(inputArgs));
    Client.GenericExtraArgsV3 memory result =
      s_onRampTestHelper.parseExtraArgsWithDefaults(DEST_CHAIN_SELECTOR, s_destChainConfig, extraArgs);

    // Should have default CCVs
    assertEq(s_defaultCCVs.length, result.ccvs.length);
  }

  // Reverts

  function test_parseExtraArgsWithDefaults_RevertWhen_NoDuplicatesAllowed_WithinRequiredCCVs() public {
    // Create user-provided CCVs with duplicates in required list
    address duplicateCCV = makeAddr("duplicateCCV");
    Client.CCV[] memory userRequiredCCVs = new Client.CCV[](2);
    userRequiredCCVs[0] = Client.CCV({ccvAddress: duplicateCCV, args: "args1"});
    userRequiredCCVs[1] = Client.CCV({ccvAddress: duplicateCCV, args: "args2"}); // Duplicate

    Client.GenericExtraArgsV3 memory inputArgs = _createV3ExtraArgs(userRequiredCCVs);

    bytes memory extraArgs = abi.encodePacked(Client.GENERIC_EXTRA_ARGS_V3_TAG, abi.encode(inputArgs));

    // Should revert due to duplicate CCVs
    vm.expectRevert(abi.encodeWithSelector(CCVConfigValidation.DuplicateCCVNotAllowed.selector, duplicateCCV));
    s_onRampTestHelper.parseExtraArgsWithDefaults(DEST_CHAIN_SELECTOR, s_destChainConfig, extraArgs);
  }
}
