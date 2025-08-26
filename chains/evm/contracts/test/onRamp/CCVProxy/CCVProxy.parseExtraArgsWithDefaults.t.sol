// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Client} from "../../../libraries/Client.sol";
import {CCVProxy} from "../../../onRamp/CCVProxy.sol";

import {CCVProxySetup} from "./CCVProxySetup.t.sol";
import {CCVProxyTestWrapper} from "./CCVProxyTestWrapper.sol";

contract CCVProxy_parseExtraArgsWithDefaults is CCVProxySetup {
  CCVProxyTestWrapper internal s_ccvProxyWrapper;

  function setUp() public override {
    super.setUp();
    s_ccvProxyWrapper = new CCVProxyTestWrapper(
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
  }

  function test_parseExtraArgsWithDefaults_WithGenericExtraArgsV3Tag() public {
    CCVProxy.DestChainConfig memory destChainConfig = CCVProxy.DestChainConfig({
      router: s_sourceRouter,
      sequenceNumber: 1,
      requiredCCV: address(0x123),
      defaultCCV: address(0x456),
      defaultExecutor: address(0x789)
    });

    Client.CCV[] memory requiredCCVs = new Client.CCV[](1);
    requiredCCVs[0] = Client.CCV({ccvAddress: address(0xABC), args: "test args"});

    Client.CCV[] memory optionalCCVs = new Client.CCV[](2);
    optionalCCVs[0] = Client.CCV({ccvAddress: address(0xDEF0), args: "optional1"});
    optionalCCVs[1] = Client.CCV({ccvAddress: address(0x1234), args: "optional2"});

    Client.EVMExtraArgsV3 memory extraArgs = Client.EVMExtraArgsV3({
      requiredCCV: requiredCCVs,
      optionalCCV: optionalCCVs,
      optionalThreshold: 1,
      finalityConfig: 10,
      executor: address(0xABCD),
      executorArgs: "executor args",
      tokenArgs: "token args"
    });

    bytes memory encodedExtraArgs = abi.encodePacked(Client.GENERIC_EXTRA_ARGS_V3_TAG, abi.encode(extraArgs));

    Client.EVMExtraArgsV3 memory result =
      s_ccvProxyWrapper.parseExtraArgsWithDefaults(destChainConfig, encodedExtraArgs);

    assertEq(result.requiredCCV.length, 2); // 1 from config + 1 from extraArgs
    assertEq(result.requiredCCV[0].ccvAddress, address(0x123)); // Required CCV from config should be first
    assertEq(result.requiredCCV[0].args, ""); // Required CCV from config should have empty args
    assertEq(result.requiredCCV[1].ccvAddress, address(0xABC)); // Original required CCV should be second
    assertEq(result.requiredCCV[1].args, "test args");

    assertEq(result.optionalCCV.length, 2);
    assertEq(result.optionalCCV[0].ccvAddress, address(0xDEF0));
    assertEq(result.optionalCCV[0].args, "optional1");
    assertEq(result.optionalCCV[1].ccvAddress, address(0x1234));
    assertEq(result.optionalCCV[1].args, "optional2");

    assertEq(result.optionalThreshold, 1);
    assertEq(result.finalityConfig, 10);
    assertEq(result.executor, address(0xABCD));
    assertEq(result.executorArgs, "executor args");
    assertEq(result.tokenArgs, "token args");
  }

  function test_parseExtraArgsWithDefaults_WithGenericExtraArgsV3Tag_NoRequiredCCVInConfig() public {
    CCVProxy.DestChainConfig memory destChainConfig = CCVProxy.DestChainConfig({
      router: s_sourceRouter,
      sequenceNumber: 1,
      requiredCCV: address(0), // No required CCV in config
      defaultCCV: address(0x456),
      defaultExecutor: address(0x789)
    });

    Client.CCV[] memory requiredCCVs = new Client.CCV[](1);
    requiredCCVs[0] = Client.CCV({ccvAddress: address(0xABC), args: "test args"});

    Client.EVMExtraArgsV3 memory extraArgs = Client.EVMExtraArgsV3({
      requiredCCV: requiredCCVs,
      optionalCCV: new Client.CCV[](0),
      optionalThreshold: 0,
      finalityConfig: 10,
      executor: address(0xABCD),
      executorArgs: "executor args",
      tokenArgs: "token args"
    });

    bytes memory encodedExtraArgs = abi.encodePacked(Client.GENERIC_EXTRA_ARGS_V3_TAG, abi.encode(extraArgs));

    Client.EVMExtraArgsV3 memory result =
      s_ccvProxyWrapper.parseExtraArgsWithDefaults(destChainConfig, encodedExtraArgs);

    assertEq(result.requiredCCV.length, 1); // Only the original required CCV
    assertEq(result.requiredCCV[0].ccvAddress, address(0xABC));
    assertEq(result.requiredCCV[0].args, "test args");
  }

  function test_parseExtraArgsWithDefaults_WithGenericExtraArgsV3Tag_NoCCVsSpecified() public {
    CCVProxy.DestChainConfig memory destChainConfig = CCVProxy.DestChainConfig({
      router: s_sourceRouter,
      sequenceNumber: 1,
      requiredCCV: address(0x123),
      defaultCCV: address(0x456),
      defaultExecutor: address(0x789)
    });

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
      s_ccvProxyWrapper.parseExtraArgsWithDefaults(destChainConfig, encodedExtraArgs);

    assertEq(result.requiredCCV.length, 2); // 1 from config + 1 default
    assertEq(result.requiredCCV[0].ccvAddress, address(0x123)); // Required CCV from config
    assertEq(result.requiredCCV[1].ccvAddress, address(0x456)); // Default CCV
    assertEq(result.executor, address(0x789)); // Default executor
  }

  function test_parseExtraArgsWithDefaults_WithLegacyExtraArgs() public {
    CCVProxy.DestChainConfig memory destChainConfig = CCVProxy.DestChainConfig({
      router: s_sourceRouter,
      sequenceNumber: 1,
      requiredCCV: address(0x123),
      defaultCCV: address(0x456),
      defaultExecutor: address(0x789)
    });

    bytes memory legacyExtraArgs = "legacy args";

    Client.EVMExtraArgsV3 memory result = s_ccvProxyWrapper.parseExtraArgsWithDefaults(destChainConfig, legacyExtraArgs);

    assertEq(result.requiredCCV.length, 2); // 1 from config + 1 default
    assertEq(result.requiredCCV[0].ccvAddress, address(0x123)); // Required CCV from config
    assertEq(result.requiredCCV[1].ccvAddress, address(0x456)); // Default CCV with legacy args
    assertEq(result.requiredCCV[1].args, legacyExtraArgs);
    assertEq(result.executorArgs, legacyExtraArgs);
    assertEq(result.executor, address(0x789)); // Default executor
  }

  function test_parseExtraArgsWithDefaults_RevertWhen_InvalidOptionalCCVThreshold() public {
    CCVProxy.DestChainConfig memory destChainConfig = CCVProxy.DestChainConfig({
      router: s_sourceRouter,
      sequenceNumber: 1,
      requiredCCV: address(0),
      defaultCCV: address(0x456),
      defaultExecutor: address(0x789)
    });

    Client.CCV[] memory optionalCCVs = new Client.CCV[](2);
    optionalCCVs[0] = Client.CCV({ccvAddress: address(0xDEF0), args: "optional1"});
    optionalCCVs[1] = Client.CCV({ccvAddress: address(0x5678), args: "optional2"});

    Client.EVMExtraArgsV3 memory extraArgs = Client.EVMExtraArgsV3({
      requiredCCV: new Client.CCV[](0),
      optionalCCV: optionalCCVs,
      optionalThreshold: 2, // Invalid: threshold >= optionalCCV.length
      finalityConfig: 10,
      executor: address(0),
      executorArgs: "",
      tokenArgs: ""
    });

    bytes memory encodedExtraArgs = abi.encodePacked(Client.GENERIC_EXTRA_ARGS_V3_TAG, abi.encode(extraArgs));

    vm.expectRevert(CCVProxy.InvalidOptionalCCVThreshold.selector);
    s_ccvProxyWrapper.parseExtraArgsWithDefaults(destChainConfig, encodedExtraArgs);
  }

  function test_parseExtraArgsWithDefaults_RevertWhen_ZeroOptionalThreshold() public {
    CCVProxy.DestChainConfig memory destChainConfig = CCVProxy.DestChainConfig({
      router: s_sourceRouter,
      sequenceNumber: 1,
      requiredCCV: address(0),
      defaultCCV: address(0x456),
      defaultExecutor: address(0x789)
    });

    Client.CCV[] memory optionalCCVs = new Client.CCV[](2);
    optionalCCVs[0] = Client.CCV({ccvAddress: address(0xDEF0), args: "optional1"});
    optionalCCVs[1] = Client.CCV({ccvAddress: address(0x9ABC), args: "optional2"});

    Client.EVMExtraArgsV3 memory extraArgs = Client.EVMExtraArgsV3({
      requiredCCV: new Client.CCV[](0),
      optionalCCV: optionalCCVs,
      optionalThreshold: 0, // Invalid: zero threshold
      finalityConfig: 10,
      executor: address(0),
      executorArgs: "",
      tokenArgs: ""
    });

    bytes memory encodedExtraArgs = abi.encodePacked(Client.GENERIC_EXTRA_ARGS_V3_TAG, abi.encode(extraArgs));

    vm.expectRevert(CCVProxy.InvalidOptionalCCVThreshold.selector);
    s_ccvProxyWrapper.parseExtraArgsWithDefaults(destChainConfig, encodedExtraArgs);
  }
}
