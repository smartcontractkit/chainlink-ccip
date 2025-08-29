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
      defaultCCVs: s_defaultCCVs
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

    // User-provided CCVs should be used, plus lane mandated CCVs.
    assertEq(s_laneMandatedCCVs.length + userRequiredCCVs.length, result.requiredCCV.length);
    assertEq(s_laneMandatedCCVs[0], result.requiredCCV[0].ccvAddress);
    assertEq("", result.requiredCCV[0].args); // Lane mandated CCVs have empty args.
    assertEq(userRequiredCCVs[0].ccvAddress, result.requiredCCV[1].ccvAddress);
    assertEq(userRequiredCCVs[0].args, result.requiredCCV[1].args);
    assertEq(inputArgs.executor, result.executor);
  }

  function test_parseExtraArgsWithDefaults_V3WithEmptyRequiredCCVs() public view {
    Client.EVMExtraArgsV3 memory inputArgs = Client.EVMExtraArgsV3({
      requiredCCV: new Client.CCV[](0), // Empty required CCVs.
      optionalCCV: new Client.CCV[](0),
      optionalThreshold: 0,
      finalityConfig: 0,
      executor: address(0), // No executor specified.
      executorArgs: "",
      tokenArgs: ""
    });

    bytes memory extraArgs = abi.encodePacked(Client.GENERIC_EXTRA_ARGS_V3_TAG, abi.encode(inputArgs));

    Client.EVMExtraArgsV3 memory result = s_ccvProxyTestHelper.parseExtraArgsWithDefaults(s_destChainConfig, extraArgs);

    // Default CCVs should be applied plus mandated CCVs.
    assertEq(s_laneMandatedCCVs.length + s_defaultCCVs.length, result.requiredCCV.length); // mandated + defaults
    assertEq(s_laneMandatedCCVs[0], result.requiredCCV[0].ccvAddress);
    assertEq("", result.requiredCCV[0].args); // Lane mandated CCVs have empty args.
    assertEq(s_defaultCCVs[0], result.requiredCCV[1].ccvAddress);
    assertEq("", result.requiredCCV[1].args); // Default CCVs have empty args.
    assertEq(s_defaultCCVs[1], result.requiredCCV[2].ccvAddress);
    assertEq("", result.requiredCCV[2].args); // Default CCVs have empty args.

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

    // Default CCVs should be used with V2 args passed to each CCV.
    assertEq(s_laneMandatedCCVs.length + s_defaultCCVs.length, result.requiredCCV.length); // mandated + defaults
    assertEq(s_laneMandatedCCVs[0], result.requiredCCV[0].ccvAddress);
    assertEq("", result.requiredCCV[0].args); // Mandated CCVs always have empty args.
    assertEq(s_defaultCCVs[0], result.requiredCCV[1].ccvAddress);
    assertEq(legacyExtraArgs, result.requiredCCV[1].args);
    assertEq(s_defaultCCVs[1], result.requiredCCV[2].ccvAddress);
    assertEq(legacyExtraArgs, result.requiredCCV[2].args);

    // V2 args should be set as executor args.
    assertEq(legacyExtraArgs, result.executorArgs);
    assertEq(s_defaultExecutor, result.executor);
  }

  function test_parseExtraArgsWithDefaults_WithLaneMandatedCCVsDuplicates_PreserveUserExtraArgs() public {
    // Modify lane mandated CCVs to include a duplicate.
    uint256 duplicateTestSize = 2;
    s_laneMandatedCCVs = new address[](duplicateTestSize);
    s_laneMandatedCCVs[0] = makeAddr("mandatedCCV1");
    s_laneMandatedCCVs[1] = makeAddr("userCCV1"); // Will be duplicate with user CCV.

    s_destChainConfig.laneMandatedCCVs = s_laneMandatedCCVs;

    Client.CCV[] memory userRequiredCCVs = new Client.CCV[](1);
    userRequiredCCVs[0] = Client.CCV({ccvAddress: s_laneMandatedCCVs[1], args: "userArgs"});

    Client.EVMExtraArgsV3 memory inputArgs = Client.EVMExtraArgsV3({
      requiredCCV: userRequiredCCVs,
      optionalCCV: new Client.CCV[](0),
      optionalThreshold: 0,
      finalityConfig: 0,
      executor: address(0),
      executorArgs: "",
      tokenArgs: ""
    });

    bytes memory extraArgs = abi.encodePacked(Client.GENERIC_EXTRA_ARGS_V3_TAG, abi.encode(inputArgs));

    Client.EVMExtraArgsV3 memory result = s_ccvProxyTestHelper.parseExtraArgsWithDefaults(s_destChainConfig, extraArgs);

    // Should have 2 CCVs (mandatedCCV1 + userCCV1, no duplicate due to deduplication).
    assertEq(2, result.requiredCCV.length); // After deduplication: 1 unique mandated + 1 user.
    assertEq(s_laneMandatedCCVs[0], result.requiredCCV[0].ccvAddress);
    assertEq("", result.requiredCCV[0].args); // 1st Mandated CCVs has empty args.
    assertEq(s_laneMandatedCCVs[1], result.requiredCCV[1].ccvAddress);
    assertEq(userRequiredCCVs[0].args, result.requiredCCV[1].args); // User args preserved, even with duplicate mandated CCV.
  }

  function test_parseExtraArgsWithDefaults_WithOptionalCCVs() public {
    // Clear lane mandated CCVs and default CCVs for this test.
    s_destChainConfig.laneMandatedCCVs = new address[](0);
    s_destChainConfig.defaultCCVs = new address[](0);

    Client.CCV[] memory requiredCCVs = new Client.CCV[](1);
    requiredCCVs[0] = Client.CCV({ccvAddress: makeAddr("requiredCCV"), args: ""});

    uint256 optionalCCVCount = 3;
    Client.CCV[] memory optionalCCVs = new Client.CCV[](optionalCCVCount);
    optionalCCVs[0] = Client.CCV({ccvAddress: makeAddr("optionalCCV1"), args: "opt1"});
    optionalCCVs[1] = Client.CCV({ccvAddress: makeAddr("optionalCCV2"), args: "opt2"});
    optionalCCVs[2] = Client.CCV({ccvAddress: makeAddr("optionalCCV3"), args: "opt3"});

    uint8 optionalThreshold = 2;
    Client.EVMExtraArgsV3 memory inputArgs = Client.EVMExtraArgsV3({
      requiredCCV: requiredCCVs,
      optionalCCV: optionalCCVs,
      optionalThreshold: optionalThreshold, // 2 out of 3 required
      finalityConfig: 0,
      executor: address(0),
      executorArgs: "",
      tokenArgs: ""
    });

    bytes memory extraArgs = abi.encodePacked(Client.GENERIC_EXTRA_ARGS_V3_TAG, abi.encode(inputArgs));

    Client.EVMExtraArgsV3 memory result = s_ccvProxyTestHelper.parseExtraArgsWithDefaults(s_destChainConfig, extraArgs);

    assertEq(requiredCCVs.length, result.requiredCCV.length);
    assertEq(optionalCCVCount, result.optionalCCV.length);
    assertEq(optionalThreshold, result.optionalThreshold);
  }

  // Reverts

  function test_parseExtraArgsWithDefaults_RevertWhen_InvalidOptionalThreshold_TooHigh() public {
    uint256 invalidCCVCount = 2;
    Client.CCV[] memory optionalCCVs = new Client.CCV[](invalidCCVCount);
    optionalCCVs[0] = Client.CCV({ccvAddress: makeAddr("optionalCCV1"), args: ""});
    optionalCCVs[1] = Client.CCV({ccvAddress: makeAddr("optionalCCV2"), args: ""});

    uint8 invalidThreshold = uint8(optionalCCVs.length + 1); // Threshold > array length
    Client.EVMExtraArgsV3 memory inputArgs = Client.EVMExtraArgsV3({
      requiredCCV: new Client.CCV[](0),
      optionalCCV: optionalCCVs,
      optionalThreshold: invalidThreshold, // Threshold > array length
      finalityConfig: 0,
      executor: address(0),
      executorArgs: "",
      tokenArgs: ""
    });

    bytes memory extraArgs = abi.encodePacked(Client.GENERIC_EXTRA_ARGS_V3_TAG, abi.encode(inputArgs));

    vm.expectRevert(CCVProxy.InvalidOptionalCCVThreshold.selector);
    s_ccvProxyTestHelper.parseExtraArgsWithDefaults(s_destChainConfig, extraArgs);
  }

  function test_parseExtraArgsWithDefaults_RevertWhen_InvalidOptionalThreshold_Zero() public {
    Client.CCV[] memory optionalCCVs = new Client.CCV[](2);
    optionalCCVs[0] = Client.CCV({ccvAddress: makeAddr("optionalCCV1"), args: ""});
    optionalCCVs[1] = Client.CCV({ccvAddress: makeAddr("optionalCCV2"), args: ""});

    Client.EVMExtraArgsV3 memory inputArgs = Client.EVMExtraArgsV3({
      requiredCCV: new Client.CCV[](0),
      optionalCCV: optionalCCVs,
      optionalThreshold: 0, // Zero threshold with optional CCVs (invalid).
      finalityConfig: 0,
      executor: address(0),
      executorArgs: "",
      tokenArgs: ""
    });

    bytes memory extraArgs = abi.encodePacked(Client.GENERIC_EXTRA_ARGS_V3_TAG, abi.encode(inputArgs));

    vm.expectRevert(CCVProxy.InvalidOptionalCCVThreshold.selector);
    s_ccvProxyTestHelper.parseExtraArgsWithDefaults(s_destChainConfig, extraArgs);
  }

  function test_parseExtraArgsWithDefaults_RevertWhen_InvalidOptionalThreshold_EqualLength() public {
    Client.CCV[] memory optionalCCVs = new Client.CCV[](2);
    optionalCCVs[0] = Client.CCV({ccvAddress: makeAddr("optionalCCV1"), args: ""});
    optionalCCVs[1] = Client.CCV({ccvAddress: makeAddr("optionalCCV2"), args: ""});

    Client.EVMExtraArgsV3 memory inputArgs = Client.EVMExtraArgsV3({
      requiredCCV: new Client.CCV[](0),
      optionalCCV: optionalCCVs,
      optionalThreshold: uint8(optionalCCVs.length), // Threshold == array length (all required, defeats purpose).
      finalityConfig: 0,
      executor: address(0),
      executorArgs: "",
      tokenArgs: ""
    });

    bytes memory extraArgs = abi.encodePacked(Client.GENERIC_EXTRA_ARGS_V3_TAG, abi.encode(inputArgs));

    vm.expectRevert(CCVProxy.InvalidOptionalCCVThreshold.selector);
    s_ccvProxyTestHelper.parseExtraArgsWithDefaults(s_destChainConfig, extraArgs);
  }
}
