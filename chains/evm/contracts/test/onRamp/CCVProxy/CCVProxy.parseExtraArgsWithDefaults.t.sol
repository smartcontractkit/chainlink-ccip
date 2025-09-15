// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Client} from "../../../libraries/Client.sol";
import {CCVProxy} from "../../../onRamp/CCVProxy.sol";
import {CCVProxyTestHelper} from "../../helpers/CCVProxyTestHelper.sol";
import {CCVProxySetup} from "./CCVProxySetup.t.sol";

contract CCVProxy_parseExtraArgsWithDefaults is CCVProxySetup {
  CCVProxyTestHelper internal s_ccvProxyTestHelper;

  address[] internal s_defaultCCVs;
  address[] internal s_laneMandatedCCVs;
  address internal s_defaultExecutor;
  CCVProxy.DestChainConfig internal s_destChainConfig;

  function setUp() public override {
    super.setUp();
    s_ccvProxyTestHelper = new CCVProxyTestHelper(
      CCVProxy.StaticConfig({
        chainSelector: SOURCE_CHAIN_SELECTOR,
        rmnRemote: s_mockRMNRemote,
        tokenAdminRegistry: address(s_tokenAdminRegistry)
      }),
      CCVProxy.DynamicConfig({
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

    s_defaultExecutor = makeAddr("defaultExecutor");

    s_destChainConfig = CCVProxy.DestChainConfig({
      router: s_sourceRouter,
      sequenceNumber: 0,
      defaultExecutor: s_defaultExecutor,
      laneMandatedCCVs: s_laneMandatedCCVs,
      defaultCCVs: s_defaultCCVs,
      ccvAggregator: abi.encodePacked(address(s_ccvAggregatorRemote))
    });
  }

  function test_parseExtraArgsWithDefaults_V3WithUserProvidedCCVs() public {
    Client.CCV[] memory userRequiredCCVs = new Client.CCV[](1);
    userRequiredCCVs[0] = Client.CCV({ccvAddress: makeAddr("userCCV1"), args: "userArgs"});

    Client.EVMExtraArgsV3 memory inputArgs = Client.EVMExtraArgsV3({
      requiredCCV: userRequiredCCVs,
      optionalCCV: new Client.CCV[](0),
      optionalThreshold: 0,
      finalityConfig: 0,
      executor: makeAddr("userExecutor"),
      executorArgs: "execArgs",
      tokenArgs: "tokenArgs"
    });

    bytes memory extraArgs = abi.encodePacked(Client.GENERIC_EXTRA_ARGS_V3_TAG, abi.encode(inputArgs));

    Client.EVMExtraArgsV3 memory result = s_ccvProxyTestHelper.parseExtraArgsWithDefaults(s_destChainConfig, extraArgs);

    // User-provided CCVs should be used (no lane mandated CCVs added in parseExtraArgsWithDefaults anymore)
    assertEq(userRequiredCCVs.length, result.requiredCCV.length);
    assertEq(userRequiredCCVs[0].ccvAddress, result.requiredCCV[0].ccvAddress);
    assertEq(userRequiredCCVs[0].args, result.requiredCCV[0].args);
    assertEq(inputArgs.executor, result.executor);
  }

  function test_parseExtraArgsWithDefaults_V3WithEmptyRequiredCCVs() public view {
    Client.EVMExtraArgsV3 memory inputArgs = _createV3ExtraArgs(
      new Client.CCV[](0), // Empty required CCVs.
      new Client.CCV[](0),
      0
    );

    bytes memory extraArgs = abi.encodePacked(Client.GENERIC_EXTRA_ARGS_V3_TAG, abi.encode(inputArgs));

    Client.EVMExtraArgsV3 memory result = s_ccvProxyTestHelper.parseExtraArgsWithDefaults(s_destChainConfig, extraArgs);

    // Default CCVs should be applied (no lane mandated CCVs added in parseExtraArgsWithDefaults anymore)
    assertEq(s_defaultCCVs.length, result.requiredCCV.length);
    assertEq(s_defaultCCVs[0], result.requiredCCV[0].ccvAddress);
    assertEq("", result.requiredCCV[0].args); // Default CCVs have empty args.
    assertEq(s_defaultCCVs[1], result.requiredCCV[1].ccvAddress);
    assertEq("", result.requiredCCV[1].args); // Default CCVs have empty args.

    // Default executor should be applied.
    assertEq(s_defaultExecutor, result.executor);
  }

  function test_parseExtraArgsWithDefaults_OldExtraArgs() public view {
    // Use GenericExtraArgsV2 format.
    uint256 gasLimit = 300_000;
    Client.GenericExtraArgsV2 memory v2Args =
      Client.GenericExtraArgsV2({gasLimit: gasLimit, allowOutOfOrderExecution: true});

    bytes memory legacyExtraArgs = Client._argsToBytes(v2Args);

    Client.EVMExtraArgsV3 memory result =
      s_ccvProxyTestHelper.parseExtraArgsWithDefaults(s_destChainConfig, legacyExtraArgs);

    // Default CCVs should be used with V2 args passed to each CCV (no lane mandated CCVs added in parseExtraArgsWithDefaults anymore)
    assertEq(s_defaultCCVs.length, result.requiredCCV.length);
    assertEq(s_defaultCCVs[0], result.requiredCCV[0].ccvAddress);
    assertEq(legacyExtraArgs, result.requiredCCV[0].args);
    assertEq(s_defaultCCVs[1], result.requiredCCV[1].ccvAddress);
    assertEq(legacyExtraArgs, result.requiredCCV[1].args);

    // V2 args should be set as executor args.
    assertEq(legacyExtraArgs, result.executorArgs);
    assertEq(s_defaultExecutor, result.executor);
  }

  function test_parseExtraArgsWithDefaults_NoDuplicatesAllowed_WithinRequiredCCVs() public {
    // Create user-provided CCVs with duplicates in required list
    address duplicateCCV = makeAddr("duplicateCCV");
    Client.CCV[] memory userRequiredCCVs = new Client.CCV[](2);
    userRequiredCCVs[0] = Client.CCV({ccvAddress: duplicateCCV, args: "args1"});
    userRequiredCCVs[1] = Client.CCV({ccvAddress: duplicateCCV, args: "args2"}); // Duplicate

    Client.EVMExtraArgsV3 memory inputArgs = _createV3ExtraArgs(userRequiredCCVs, new Client.CCV[](0), 0);

    bytes memory extraArgs = abi.encodePacked(Client.GENERIC_EXTRA_ARGS_V3_TAG, abi.encode(inputArgs));

    // Should revert due to duplicate CCVs
    vm.expectRevert(abi.encodeWithSelector(CCVProxy.DuplicateCCVInUserInput.selector, duplicateCCV));
    s_ccvProxyTestHelper.parseExtraArgsWithDefaults(s_destChainConfig, extraArgs);
  }

  function test_parseExtraArgsWithDefaults_NoDuplicatesAllowed_WithinOptionalCCVs() public {
    // Create user-provided CCVs with duplicates in optional list
    address duplicateCCV = makeAddr("duplicateCCV");
    Client.CCV[] memory optionalCCVs = new Client.CCV[](2);
    optionalCCVs[0] = Client.CCV({ccvAddress: duplicateCCV, args: "opt1"});
    optionalCCVs[1] = Client.CCV({ccvAddress: duplicateCCV, args: "opt2"}); // Duplicate

    Client.EVMExtraArgsV3 memory inputArgs = _createV3ExtraArgs(new Client.CCV[](0), optionalCCVs, 1);

    bytes memory extraArgs = abi.encodePacked(Client.GENERIC_EXTRA_ARGS_V3_TAG, abi.encode(inputArgs));

    // Should revert due to duplicate CCVs
    vm.expectRevert(abi.encodeWithSelector(CCVProxy.DuplicateCCVInUserInput.selector, duplicateCCV));
    s_ccvProxyTestHelper.parseExtraArgsWithDefaults(s_destChainConfig, extraArgs);
  }

  function test_parseExtraArgsWithDefaults_NoDuplicatesAllowed_BetweenRequiredAndOptional() public {
    // Create user-provided CCVs with same CCV in both required and optional
    address duplicateCCV = makeAddr("duplicateCCV");
    Client.CCV[] memory requiredCCVs = new Client.CCV[](1);
    requiredCCVs[0] = Client.CCV({ccvAddress: duplicateCCV, args: "required"});

    Client.CCV[] memory optionalCCVs = new Client.CCV[](2);
    optionalCCVs[0] = Client.CCV({ccvAddress: duplicateCCV, args: "optional"}); // Same as required
    optionalCCVs[1] = Client.CCV({ccvAddress: makeAddr("optionalCCV2"), args: "optional"}); // Need one more here to set threshold 1

    Client.EVMExtraArgsV3 memory inputArgs = _createV3ExtraArgs(requiredCCVs, optionalCCVs, 1);

    bytes memory extraArgs = abi.encodePacked(Client.GENERIC_EXTRA_ARGS_V3_TAG, abi.encode(inputArgs));

    // Should revert due to duplicate CCVs between required and optional
    vm.expectRevert(abi.encodeWithSelector(CCVProxy.DuplicateCCVInUserInput.selector, duplicateCCV));
    s_ccvProxyTestHelper.parseExtraArgsWithDefaults(s_destChainConfig, extraArgs);
  }

  function test_parseExtraArgsWithDefaults_WithOptionalCCVs() public {
    // Use non-empty defaultCCVs as it's an invariant
    // The test should work with defaults but user provides their own CCVs

    Client.CCV[] memory requiredCCVs = new Client.CCV[](1);
    requiredCCVs[0] = Client.CCV({ccvAddress: makeAddr("requiredCCV"), args: ""});

    uint256 optionalCCVCount = 3;
    Client.CCV[] memory optionalCCVs = new Client.CCV[](optionalCCVCount);
    optionalCCVs[0] = Client.CCV({ccvAddress: makeAddr("optionalCCV1"), args: "opt1"});
    optionalCCVs[1] = Client.CCV({ccvAddress: makeAddr("optionalCCV2"), args: "opt2"});
    optionalCCVs[2] = Client.CCV({ccvAddress: makeAddr("optionalCCV3"), args: "opt3"});

    uint8 optionalThreshold = 2;
    Client.EVMExtraArgsV3 memory inputArgs = _createV3ExtraArgs(
      requiredCCVs,
      optionalCCVs,
      optionalThreshold // 2 out of 3 required
    );

    bytes memory extraArgs = abi.encodePacked(Client.GENERIC_EXTRA_ARGS_V3_TAG, abi.encode(inputArgs));

    Client.EVMExtraArgsV3 memory result = s_ccvProxyTestHelper.parseExtraArgsWithDefaults(s_destChainConfig, extraArgs);

    // User provided required CCVs override defaults
    assertEq(requiredCCVs.length, result.requiredCCV.length);
    assertEq(requiredCCVs[0].ccvAddress, result.requiredCCV[0].ccvAddress);
    assertEq(optionalCCVCount, result.optionalCCV.length);
    assertEq(optionalThreshold, result.optionalThreshold);
  }

  function test_parseExtraArgsWithDefaults_EmptyRequiredWithOptionalCCVs_NoDefaultsApplied() public {
    // Test case: empty required CCVs but non-empty optional CCVs
    Client.CCV[] memory optionalCCVs = new Client.CCV[](2);
    optionalCCVs[0] = Client.CCV({ccvAddress: makeAddr("optionalCCV1"), args: "opt1"});
    optionalCCVs[1] = Client.CCV({ccvAddress: makeAddr("optionalCCV2"), args: "opt2"});

    uint8 optionalThreshold = 1;
    Client.EVMExtraArgsV3 memory inputArgs = _createV3ExtraArgs(
      new Client.CCV[](0), // Empty required CCVs
      optionalCCVs,
      optionalThreshold
    );

    bytes memory extraArgs = abi.encodePacked(Client.GENERIC_EXTRA_ARGS_V3_TAG, abi.encode(inputArgs));

    Client.EVMExtraArgsV3 memory result = s_ccvProxyTestHelper.parseExtraArgsWithDefaults(s_destChainConfig, extraArgs);

    // Required CCVs should remain empty (no defaults applied)
    assertEq(0, result.requiredCCV.length);
    // Optional CCVs should be preserved as provided
    _assertCCVArraysEqual(optionalCCVs, result.optionalCCV);
    assertEq(optionalThreshold, result.optionalThreshold);
    // Default executor should still be applied
    assertEq(s_defaultExecutor, result.executor);
  }

  // Additional test for defaults when no user CCVs provided
  function test_parseExtraArgsWithDefaults_DefaultCCVsAlwaysPresent() public view {
    // Ensure defaultCCVs.length > 0 (invariant)
    assertTrue(s_destChainConfig.defaultCCVs.length > 0, "defaultCCVs must not be empty");

    // Test with empty user input
    Client.EVMExtraArgsV3 memory inputArgs = _createV3ExtraArgs(new Client.CCV[](0), new Client.CCV[](0), 0);

    bytes memory extraArgs = abi.encodePacked(Client.GENERIC_EXTRA_ARGS_V3_TAG, abi.encode(inputArgs));
    Client.EVMExtraArgsV3 memory result = s_ccvProxyTestHelper.parseExtraArgsWithDefaults(s_destChainConfig, extraArgs);

    // Should have default CCVs
    assertEq(s_defaultCCVs.length, result.requiredCCV.length);
  }

  // Reverts

  function test_parseExtraArgsWithDefaults_RevertWhen_InvalidOptionalThreshold_TooHigh() public {
    uint256 invalidCCVCount = 2;
    Client.CCV[] memory optionalCCVs = new Client.CCV[](invalidCCVCount);
    optionalCCVs[0] = Client.CCV({ccvAddress: makeAddr("optionalCCV1"), args: ""});
    optionalCCVs[1] = Client.CCV({ccvAddress: makeAddr("optionalCCV2"), args: ""});

    uint8 invalidThreshold = uint8(optionalCCVs.length + 1); // Threshold > array length
    Client.EVMExtraArgsV3 memory inputArgs = _createV3ExtraArgs(
      new Client.CCV[](0),
      optionalCCVs,
      invalidThreshold // Threshold > array length
    );

    bytes memory extraArgs = abi.encodePacked(Client.GENERIC_EXTRA_ARGS_V3_TAG, abi.encode(inputArgs));

    vm.expectRevert(CCVProxy.InvalidOptionalCCVThreshold.selector);
    s_ccvProxyTestHelper.parseExtraArgsWithDefaults(s_destChainConfig, extraArgs);
  }

  function test_parseExtraArgsWithDefaults_RevertWhen_InvalidOptionalThreshold_Zero() public {
    Client.CCV[] memory optionalCCVs = new Client.CCV[](1);
    optionalCCVs[0] = Client.CCV({ccvAddress: makeAddr("optionalCCV1"), args: ""});

    Client.EVMExtraArgsV3 memory inputArgs = _createV3ExtraArgs(
      new Client.CCV[](0),
      optionalCCVs,
      0 // Zero threshold with optional CCVs (invalid).
    );

    bytes memory extraArgs = abi.encodePacked(Client.GENERIC_EXTRA_ARGS_V3_TAG, abi.encode(inputArgs));

    vm.expectRevert(CCVProxy.InvalidOptionalCCVThreshold.selector);
    s_ccvProxyTestHelper.parseExtraArgsWithDefaults(s_destChainConfig, extraArgs);
  }

  function test_parseExtraArgsWithDefaults_RevertWhen_InvalidOptionalThreshold_EqualLength() public {
    Client.CCV[] memory optionalCCVs = new Client.CCV[](2);
    optionalCCVs[0] = Client.CCV({ccvAddress: makeAddr("optionalCCV1"), args: ""});
    optionalCCVs[1] = Client.CCV({ccvAddress: makeAddr("optionalCCV2"), args: ""});

    Client.EVMExtraArgsV3 memory inputArgs = _createV3ExtraArgs(
      new Client.CCV[](0),
      optionalCCVs,
      uint8(optionalCCVs.length) // Threshold == array length (all required, defeats purpose).
    );

    bytes memory extraArgs = abi.encodePacked(Client.GENERIC_EXTRA_ARGS_V3_TAG, abi.encode(inputArgs));

    vm.expectRevert(CCVProxy.InvalidOptionalCCVThreshold.selector);
    s_ccvProxyTestHelper.parseExtraArgsWithDefaults(s_destChainConfig, extraArgs);
  }
}
