// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {FeeQuoter} from "../../FeeQuoter.sol";
import {Client} from "../../libraries/Client.sol";
import {Internal} from "../../libraries/Internal.sol";
import {FeeQuoterSetup} from "./FeeQuoterSetup.t.sol";

contract FeeQuoter_resolveLegacyArgs is FeeQuoterSetup {
  function _setupDestChain(
    bytes4 chainFamilySelector
  ) internal returns (FeeQuoter.DestChainConfig memory config) {
    config = _generateFeeQuoterDestChainConfigArgs()[0].destChainConfig;
    config.chainFamilySelector = chainFamilySelector;

    FeeQuoter.DestChainConfigArgs[] memory destChainConfigs = new FeeQuoter.DestChainConfigArgs[](1);
    destChainConfigs[0] =
      FeeQuoter.DestChainConfigArgs({destChainSelector: DEST_CHAIN_SELECTOR, destChainConfig: config});
    s_feeQuoter.applyDestChainConfigUpdates(destChainConfigs);

    return config;
  }

  function _setupDestChainWithConfig(
    FeeQuoter.DestChainConfig memory config
  ) internal {
    FeeQuoter.DestChainConfigArgs[] memory destChainConfigs = new FeeQuoter.DestChainConfigArgs[](1);
    destChainConfigs[0] =
      FeeQuoter.DestChainConfigArgs({destChainSelector: DEST_CHAIN_SELECTOR, destChainConfig: config});
    s_feeQuoter.applyDestChainConfigUpdates(destChainConfigs);
  }

  function test_resolveLegacyArgs_EVM() public {
    _setupDestChain(Internal.CHAIN_FAMILY_SELECTOR_EVM);

    Client.EVMExtraArgsV1 memory evmArgs = Client.EVMExtraArgsV1({gasLimit: GAS_LIMIT});
    bytes memory extraArgs = Client._argsToBytes(evmArgs);

    (bytes memory tokenReceiver, uint32 gasLimit, bytes memory executorArgs) =
      s_feeQuoter.resolveLegacyArgs(DEST_CHAIN_SELECTOR, extraArgs);

    // For EVM, tokenReceiver and executorArgs should be empty.
    assertEq("", tokenReceiver);
    assertEq(GAS_LIMIT, gasLimit);
    assertEq("", executorArgs);
  }

  function test_resolveLegacyArgs_EVM_DefaultGasLimit() public {
    FeeQuoter.DestChainConfig memory config = _setupDestChain(Internal.CHAIN_FAMILY_SELECTOR_EVM);
    config.defaultTxGasLimit = 100_000;
    _setupDestChainWithConfig(config);

    bytes memory extraArgs = new bytes(0);

    (bytes memory tokenReceiver, uint32 gasLimit, bytes memory executorArgs) =
      s_feeQuoter.resolveLegacyArgs(DEST_CHAIN_SELECTOR, extraArgs);

    // Should use default gas limit.
    assertEq("", tokenReceiver);
    assertEq(config.defaultTxGasLimit, gasLimit);
    assertEq("", executorArgs);
  }

  function test_resolveLegacyArgs_Aptos() public {
    _setupDestChain(Internal.CHAIN_FAMILY_SELECTOR_APTOS);

    Client.GenericExtraArgsV2 memory genericArgs =
      Client.GenericExtraArgsV2({gasLimit: GAS_LIMIT, allowOutOfOrderExecution: false});
    bytes memory extraArgs = Client._argsToBytes(genericArgs);

    (bytes memory tokenReceiver, uint32 gasLimit, bytes memory executorArgs) =
      s_feeQuoter.resolveLegacyArgs(DEST_CHAIN_SELECTOR, extraArgs);

    // For Aptos, tokenReceiver and executorArgs should be empty.
    assertEq("", tokenReceiver);
    assertEq(GAS_LIMIT, gasLimit);
    assertEq("", executorArgs);
  }

  function test_resolveLegacyArgs_TVM() public {
    _setupDestChain(Internal.CHAIN_FAMILY_SELECTOR_TVM);

    Client.GenericExtraArgsV2 memory genericArgs =
      Client.GenericExtraArgsV2({gasLimit: GAS_LIMIT, allowOutOfOrderExecution: true});
    bytes memory extraArgs = Client._argsToBytes(genericArgs);

    (bytes memory tokenReceiver, uint32 gasLimit, bytes memory executorArgs) =
      s_feeQuoter.resolveLegacyArgs(DEST_CHAIN_SELECTOR, extraArgs);

    // For TVM, tokenReceiver and executorArgs should be empty.
    assertEq("", tokenReceiver);
    assertEq(GAS_LIMIT, gasLimit);
    assertEq("", executorArgs);
  }

  function test_resolveLegacyArgs_SVM() public {
    _setupDestChain(Internal.CHAIN_FAMILY_SELECTOR_SVM);

    bytes32 testTokenReceiver = bytes32(uint256(0xabcdef));
    bytes32[] memory accounts = new bytes32[](2);
    accounts[0] = bytes32(uint256(1));
    accounts[1] = bytes32(uint256(2));

    Client.SVMExtraArgsV1 memory svmArgs = Client.SVMExtraArgsV1({
      computeUnits: GAS_LIMIT,
      accountIsWritableBitmap: 3,
      allowOutOfOrderExecution: true,
      tokenReceiver: testTokenReceiver,
      accounts: accounts
    });
    bytes memory extraArgs = Client._svmArgsToBytes(svmArgs);

    (bytes memory tokenReceiver, uint32 gasLimit, bytes memory executorArgs) =
      s_feeQuoter.resolveLegacyArgs(DEST_CHAIN_SELECTOR, extraArgs);

    // For SVM, tokenReceiver should be encoded, executorArgs should be constructed.
    assertEq(abi.encode(testTokenReceiver), tokenReceiver);
    assertEq(GAS_LIMIT, gasLimit);
    // executorArgs should be 8 + 2 + (accounts.length * 32) = 74 bytes.
    assertEq(2 + 8 + 32 * accounts.length, executorArgs.length);
  }

  function test_resolveLegacyArgs_SVM_NoAccounts() public {
    _setupDestChain(Internal.CHAIN_FAMILY_SELECTOR_SVM);

    bytes32 testTokenReceiver = bytes32(uint256(0x123));

    Client.SVMExtraArgsV1 memory svmArgs = Client.SVMExtraArgsV1({
      computeUnits: GAS_LIMIT,
      accountIsWritableBitmap: 0,
      allowOutOfOrderExecution: false,
      tokenReceiver: testTokenReceiver,
      accounts: new bytes32[](0)
    });
    bytes memory extraArgs = Client._svmArgsToBytes(svmArgs);

    (bytes memory tokenReceiver, uint32 gasLimit, bytes memory executorArgs) =
      s_feeQuoter.resolveLegacyArgs(DEST_CHAIN_SELECTOR, extraArgs);

    // For SVM, tokenReceiver should be encoded.
    assertEq(abi.encode(testTokenReceiver), tokenReceiver);
    assertEq(GAS_LIMIT, gasLimit);
    // executorArgs should be 8 + 2 + 0 = 10 bytes.
    assertEq(10, executorArgs.length);
  }

  function test_resolveLegacyArgs_Sui() public {
    _setupDestChain(Internal.CHAIN_FAMILY_SELECTOR_SUI);

    bytes32 testTokenReceiver = bytes32(uint256(0xfedcba));
    bytes32[] memory objectIds = new bytes32[](3);
    objectIds[0] = bytes32(uint256(100));
    objectIds[1] = bytes32(uint256(200));
    objectIds[2] = bytes32(uint256(300));

    Client.SuiExtraArgsV1 memory suiArgs = Client.SuiExtraArgsV1({
      gasLimit: GAS_LIMIT,
      allowOutOfOrderExecution: false,
      tokenReceiver: testTokenReceiver,
      receiverObjectIds: objectIds
    });
    bytes memory extraArgs = Client._suiArgsToBytes(suiArgs);

    (bytes memory tokenReceiver, uint32 gasLimit, bytes memory executorArgs) =
      s_feeQuoter.resolveLegacyArgs(DEST_CHAIN_SELECTOR, extraArgs);

    // For SUI, tokenReceiver should be encoded, executorArgs should be constructed.
    assertEq(abi.encode(testTokenReceiver), tokenReceiver);
    assertEq(GAS_LIMIT, gasLimit);
    // executorArgs should be 2 + (receiverObjectIds.length * 32) = 98 bytes.
    assertEq(2 + 32 * objectIds.length, executorArgs.length);
  }

  function test_resolveLegacyArgs_Sui_NoObjectIds() public {
    _setupDestChain(Internal.CHAIN_FAMILY_SELECTOR_SUI);

    bytes32 testTokenReceiver = bytes32(uint256(0x456));

    Client.SuiExtraArgsV1 memory suiArgs = Client.SuiExtraArgsV1({
      gasLimit: GAS_LIMIT,
      allowOutOfOrderExecution: true,
      tokenReceiver: testTokenReceiver,
      receiverObjectIds: new bytes32[](0)
    });
    bytes memory extraArgs = Client._suiArgsToBytes(suiArgs);

    (bytes memory tokenReceiver, uint32 gasLimit, bytes memory executorArgs) =
      s_feeQuoter.resolveLegacyArgs(DEST_CHAIN_SELECTOR, extraArgs);

    // For SUI, tokenReceiver should be encoded.
    assertEq(abi.encode(testTokenReceiver), tokenReceiver);
    assertEq(GAS_LIMIT, gasLimit);
    // executorArgs should be 2 + 0 = 2 bytes.
    assertEq(2, executorArgs.length);
  }

  function test_resolveLegacyArgs_EVM_MaxGasLimit() public {
    FeeQuoter.DestChainConfig memory config = _setupDestChain(Internal.CHAIN_FAMILY_SELECTOR_EVM);
    config.maxPerMsgGasLimit = type(uint32).max;
    _setupDestChainWithConfig(config);

    Client.EVMExtraArgsV1 memory evmArgs = Client.EVMExtraArgsV1({gasLimit: uint256(type(uint32).max)});
    bytes memory extraArgs = Client._argsToBytes(evmArgs);

    (bytes memory tokenReceiver, uint32 gasLimit, bytes memory executorArgs) =
      s_feeQuoter.resolveLegacyArgs(DEST_CHAIN_SELECTOR, extraArgs);

    // Should accept maxPerMsgGasLimit.
    assertEq("", tokenReceiver);
    assertEq(type(uint32).max, gasLimit);
    assertEq("", executorArgs);
  }

  // Reverts

  function test_resolveLegacyArgs_RevertWhen_MessageGasLimitTooHigh_EVM() public {
    FeeQuoter.DestChainConfig memory config = _setupDestChain(Internal.CHAIN_FAMILY_SELECTOR_EVM);
    config.maxPerMsgGasLimit = 1_000_000;
    _setupDestChainWithConfig(config);

    Client.EVMExtraArgsV1 memory evmArgs = Client.EVMExtraArgsV1({gasLimit: config.maxPerMsgGasLimit + 1});
    bytes memory extraArgs = Client._argsToBytes(evmArgs);

    vm.expectRevert(FeeQuoter.MessageGasLimitTooHigh.selector);
    s_feeQuoter.resolveLegacyArgs(DEST_CHAIN_SELECTOR, extraArgs);
  }

  function test_resolveLegacyArgs_RevertWhen_InvalidExtraArgsData_SUI_EmptyExtraArgs() public {
    _setupDestChain(Internal.CHAIN_FAMILY_SELECTOR_SUI);

    bytes memory extraArgs = new bytes(0);

    vm.expectRevert(FeeQuoter.InvalidExtraArgsData.selector);
    s_feeQuoter.resolveLegacyArgs(DEST_CHAIN_SELECTOR, extraArgs);
  }

  function test_resolveLegacyArgs_RevertWhen_MessageGasLimitTooHigh_SUI() public {
    FeeQuoter.DestChainConfig memory config = _setupDestChain(Internal.CHAIN_FAMILY_SELECTOR_SUI);
    config.maxPerMsgGasLimit = 1_000_000;
    _setupDestChainWithConfig(config);

    Client.SuiExtraArgsV1 memory suiArgs = Client.SuiExtraArgsV1({
      gasLimit: config.maxPerMsgGasLimit + 1,
      allowOutOfOrderExecution: false,
      tokenReceiver: bytes32(uint256(1)),
      receiverObjectIds: new bytes32[](0)
    });
    bytes memory extraArgs = Client._suiArgsToBytes(suiArgs);

    vm.expectRevert(FeeQuoter.MessageGasLimitTooHigh.selector);
    s_feeQuoter.resolveLegacyArgs(DEST_CHAIN_SELECTOR, extraArgs);
  }

  function test_resolveLegacyArgs_RevertWhen_InvalidExtraArgsTag() public {
    _setupDestChain(Internal.CHAIN_FAMILY_SELECTOR_SUI);

    bytes memory extraArgs = abi.encodeWithSelector(bytes4(0x12345678));

    vm.expectRevert(FeeQuoter.InvalidExtraArgsTag.selector);
    s_feeQuoter.resolveLegacyArgs(DEST_CHAIN_SELECTOR, extraArgs);
  }
}
