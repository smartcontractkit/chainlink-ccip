// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Client} from "../../../libraries/Client.sol";
import {CCVProxy} from "../../../onRamp/CCVProxy.sol";

import {CCVProxyTestHelper} from "../../helpers/CCVProxyTestHelper.sol";
import {CCVProxySetup} from "./CCVProxySetup.t.sol";

contract CCVProxy_parseExtraArgsWithDefaults is CCVProxySetup {
  CCVProxyTestHelper internal s_ccvProxyTestHelper;
  address internal constant REQUIRED_CCV = address(0x1111);
  address internal constant DEFAULT_CCV = address(0x2222);
  address internal constant DEFAULT_EXECUTOR = address(0x3333);

  address internal s_userRequiredCCV1;
  address internal s_userExecutor;
  address internal s_optionalCCV1;
  address internal s_optionalCCV2;

  CCVProxy.DestChainConfig internal s_defaultDestChainConfig;

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

    // Initialize common test addresses.
    s_userRequiredCCV1 = makeAddr("userRequiredCCV1");
    s_userExecutor = makeAddr("userExecutor");
    s_optionalCCV1 = makeAddr("optionalCCV1");
    s_optionalCCV2 = makeAddr("optionalCCV2");

    // Initialize common test configuration.
    s_defaultDestChainConfig = CCVProxy.DestChainConfig({
      router: s_sourceRouter,
      sequenceNumber: 1,
      requiredCCV: REQUIRED_CCV,
      defaultCCV: DEFAULT_CCV,
      defaultExecutor: DEFAULT_EXECUTOR
    });
  }

  function test_parseExtraArgsWithDefaults_WithGenericExtraArgsV3Tag() public view {
    CCVProxy.DestChainConfig memory destChainConfig = s_defaultDestChainConfig;

    Client.CCV[] memory requiredCCVs = new Client.CCV[](1);
    requiredCCVs[0] = Client.CCV({ccvAddress: s_userRequiredCCV1, args: "test args"});

    Client.CCV[] memory optionalCCVs = new Client.CCV[](2);
    optionalCCVs[0] = Client.CCV({ccvAddress: s_optionalCCV1, args: "optional1"});
    optionalCCVs[1] = Client.CCV({ccvAddress: s_optionalCCV2, args: "optional2"});

    Client.EVMExtraArgsV3 memory extraArgs = Client.EVMExtraArgsV3({
      requiredCCV: requiredCCVs,
      optionalCCV: optionalCCVs,
      optionalThreshold: 1,
      finalityConfig: 10,
      executor: s_userExecutor,
      executorArgs: "executor args",
      tokenArgs: "token args"
    });

    bytes memory encodedExtraArgs = abi.encodePacked(Client.GENERIC_EXTRA_ARGS_V3_TAG, abi.encode(extraArgs));

    Client.EVMExtraArgsV3 memory result =
      s_ccvProxyTestHelper.parseExtraArgsWithDefaults(destChainConfig, encodedExtraArgs);

    assertEq(result.requiredCCV.length, 1 + extraArgs.requiredCCV.length); // 1 from DestChainConfig.requiredCCV + 1 from extraArgs.
    assertEq(result.requiredCCV[0].ccvAddress, destChainConfig.requiredCCV); // Required CCV from config should be first.
    assertEq(result.requiredCCV[0].args, ""); // Required CCV from config should have empty args.
    assertEq(result.requiredCCV[1].ccvAddress, s_userRequiredCCV1); // Original required CCV should be second.
    assertEq(result.requiredCCV[1].args, "test args");

    assertEq(result.optionalCCV.length, extraArgs.optionalCCV.length);
    assertEq(result.optionalCCV[0].ccvAddress, s_optionalCCV1);
    assertEq(result.optionalCCV[0].args, "optional1");
    assertEq(result.optionalCCV[1].ccvAddress, s_optionalCCV2);
    assertEq(result.optionalCCV[1].args, "optional2");

    assertEq(result.optionalThreshold, extraArgs.optionalThreshold);
    assertEq(result.finalityConfig, extraArgs.finalityConfig);
    assertEq(result.executor, extraArgs.executor);
    assertEq(result.executorArgs, extraArgs.executorArgs);
    assertEq(result.tokenArgs, extraArgs.tokenArgs);
  }

  function test_parseExtraArgsWithDefaults_WithGenericExtraArgsV3Tag_NoRequiredCCVInConfig() public view {
    CCVProxy.DestChainConfig memory destChainConfig = s_defaultDestChainConfig;
    destChainConfig.requiredCCV = address(0); // Override: No required CCV in config.

    Client.CCV[] memory requiredCCVs = new Client.CCV[](1);
    requiredCCVs[0] = Client.CCV({ccvAddress: s_userRequiredCCV1, args: "test args"});

    Client.EVMExtraArgsV3 memory extraArgs = Client.EVMExtraArgsV3({
      requiredCCV: requiredCCVs,
      optionalCCV: new Client.CCV[](0),
      optionalThreshold: 0,
      finalityConfig: 10,
      executor: s_userExecutor,
      executorArgs: "executor args",
      tokenArgs: "token args"
    });

    bytes memory encodedExtraArgs = abi.encodePacked(Client.GENERIC_EXTRA_ARGS_V3_TAG, abi.encode(extraArgs));

    Client.EVMExtraArgsV3 memory result =
      s_ccvProxyTestHelper.parseExtraArgsWithDefaults(destChainConfig, encodedExtraArgs);

    assertEq(result.requiredCCV.length, extraArgs.requiredCCV.length);
    assertEq(result.requiredCCV[0].ccvAddress, s_userRequiredCCV1);
    assertEq(result.requiredCCV[0].args, "test args");
  }

  function test_parseExtraArgsWithDefaults_WithGenericExtraArgsV3Tag_NoCCVsSpecified() public view {
    CCVProxy.DestChainConfig memory destChainConfig = s_defaultDestChainConfig;

    Client.EVMExtraArgsV3 memory extraArgs = Client.EVMExtraArgsV3({
      requiredCCV: new Client.CCV[](0),
      optionalCCV: new Client.CCV[](0),
      optionalThreshold: 0,
      finalityConfig: 10,
      executor: address(0),
      executorArgs: "",
      tokenArgs: ""
    });

    bytes memory encodedExtraArgs = abi.encodePacked(Client.GENERIC_EXTRA_ARGS_V3_TAG, abi.encode(extraArgs));

    Client.EVMExtraArgsV3 memory result =
      s_ccvProxyTestHelper.parseExtraArgsWithDefaults(destChainConfig, encodedExtraArgs);

    assertEq(result.requiredCCV.length, 2); // 1 from DestChainConfig.requiredCCV + 1 from DestChainConfig.defaultCCV.
    assertEq(result.requiredCCV[0].ccvAddress, destChainConfig.requiredCCV);
    assertEq(result.requiredCCV[1].ccvAddress, destChainConfig.defaultCCV);
    assertEq(result.executor, destChainConfig.defaultExecutor);
  }

  function test_parseExtraArgsWithDefaults_WithLegacyExtraArgs() public view {
    CCVProxy.DestChainConfig memory destChainConfig = s_defaultDestChainConfig;

    bytes memory oldExtraArgs = Client._argsToBytes(Client.EVMExtraArgsV1({gasLimit: 1000}));

    Client.EVMExtraArgsV3 memory result = s_ccvProxyTestHelper.parseExtraArgsWithDefaults(destChainConfig, oldExtraArgs);

    assertEq(result.requiredCCV.length, 2); // 1 from DestChainConfig.requiredCCV + 1 from DestChainConfig.defaultCCV.
    assertEq(result.requiredCCV[0].ccvAddress, destChainConfig.requiredCCV);
    assertEq(result.requiredCCV[1].ccvAddress, destChainConfig.defaultCCV);
    assertEq(result.requiredCCV[1].args, oldExtraArgs);
    assertEq(result.executorArgs, oldExtraArgs);
    assertEq(result.executor, destChainConfig.defaultExecutor); // Default executor.
  }

  function test_parseExtraArgsWithDefaults_RevertWhen_InvalidOptionalCCVThreshold() public {
    CCVProxy.DestChainConfig memory destChainConfig = s_defaultDestChainConfig;
    destChainConfig.requiredCCV = address(0); // Override: No required CCV for this test.

    Client.CCV[] memory optionalCCVs = new Client.CCV[](2);
    optionalCCVs[0] = Client.CCV({ccvAddress: s_optionalCCV1, args: "optional1"});
    optionalCCVs[1] = Client.CCV({ccvAddress: s_optionalCCV2, args: "optional2"});

    Client.EVMExtraArgsV3 memory extraArgs = Client.EVMExtraArgsV3({
      requiredCCV: new Client.CCV[](0),
      optionalCCV: optionalCCVs,
      optionalThreshold: uint8(optionalCCVs.length), // Invalid: threshold >= optionalCCV.length.
      finalityConfig: 10,
      executor: address(0),
      executorArgs: "",
      tokenArgs: ""
    });

    bytes memory encodedExtraArgs = abi.encodePacked(Client.GENERIC_EXTRA_ARGS_V3_TAG, abi.encode(extraArgs));

    vm.expectRevert(CCVProxy.InvalidOptionalCCVThreshold.selector);
    s_ccvProxyTestHelper.parseExtraArgsWithDefaults(destChainConfig, encodedExtraArgs);
  }

  function test_parseExtraArgsWithDefaults_RevertWhen_ZeroOptionalThreshold() public {
    CCVProxy.DestChainConfig memory destChainConfig = s_defaultDestChainConfig;
    destChainConfig.requiredCCV = address(0); // Override: No required CCV for this test.

    Client.CCV[] memory optionalCCVs = new Client.CCV[](1);
    optionalCCVs[0] = Client.CCV({ccvAddress: s_optionalCCV1, args: "optional1"});

    Client.EVMExtraArgsV3 memory extraArgs = Client.EVMExtraArgsV3({
      requiredCCV: new Client.CCV[](0),
      optionalCCV: optionalCCVs,
      optionalThreshold: 0, // Invalid: zero threshold.
      finalityConfig: 10,
      executor: address(0),
      executorArgs: "",
      tokenArgs: ""
    });

    bytes memory encodedExtraArgs = abi.encodePacked(Client.GENERIC_EXTRA_ARGS_V3_TAG, abi.encode(extraArgs));

    vm.expectRevert(CCVProxy.InvalidOptionalCCVThreshold.selector);
    s_ccvProxyTestHelper.parseExtraArgsWithDefaults(destChainConfig, encodedExtraArgs);
  }
}
