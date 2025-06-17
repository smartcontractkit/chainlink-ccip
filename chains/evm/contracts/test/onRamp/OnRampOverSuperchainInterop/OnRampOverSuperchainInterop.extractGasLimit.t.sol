// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Client} from "../../../libraries/Client.sol";
import {OnRampOverSuperchainInterop} from "../../../onRamp/OnRampOverSuperchainInterop.sol";
import {OnRampOverSuperchainInteropSetup} from "./OnRampOverSuperchainInteropSetup.t.sol";

contract OnRampOverSuperchainInterop_extractGasLimit is OnRampOverSuperchainInteropSetup {
  function _generateEVMExtraArgsV1(
    uint256 gasLimit
  ) internal pure returns (bytes memory) {
    Client.EVMExtraArgsV1 memory extraArgs = Client.EVMExtraArgsV1({gasLimit: gasLimit});
    return Client._argsToBytes(extraArgs);
  }

  function _generateGenericExtraArgsV2(
    uint256 gasLimit,
    bool allowOutOfOrderExecution
  ) internal pure returns (bytes memory) {
    Client.GenericExtraArgsV2 memory extraArgs =
      Client.GenericExtraArgsV2({gasLimit: gasLimit, allowOutOfOrderExecution: allowOutOfOrderExecution});
    return Client._argsToBytes(extraArgs);
  }

  function _generateExtraArgsWithAdditionalData(
    uint256 gasLimit,
    bytes memory additionalData
  ) internal pure returns (bytes memory) {
    bytes4 selector = bytes4(keccak256("extraArgs"));
    return abi.encodeWithSelector(selector, gasLimit, additionalData);
  }

  function _generateShortExtraArgs(
    uint256 length
  ) internal pure returns (bytes memory) {
    bytes memory shortArgs = new bytes(length);
    for (uint256 i = 0; i < length; i++) {
      shortArgs[i] = bytes1(uint8(i % 256));
    }
    return shortArgs;
  }

  function test_BasicGasLimitExtraction_EVMExtraArgsV1() public view {
    uint256 expectedGasLimit = 200000;
    bytes memory extraArgs = _generateEVMExtraArgsV1(expectedGasLimit);

    uint256 actualGasLimit = s_onRampOverSuperchainInterop.extractGasLimit(extraArgs);

    assertEq(actualGasLimit, expectedGasLimit);
  }

  function test_BasicGasLimitExtraction_GenericExtraArgsV2() public view {
    uint256 expectedGasLimit = 200000;
    bytes memory extraArgs = _generateGenericExtraArgsV2(expectedGasLimit, false);

    uint256 actualGasLimit = s_onRampOverSuperchainInterop.extractGasLimit(extraArgs);

    assertEq(actualGasLimit, expectedGasLimit);
  }

  function test_ZeroGasLimit_EVMExtraArgsV1() public view {
    uint256 expectedGasLimit = 0;
    bytes memory extraArgs = _generateEVMExtraArgsV1(expectedGasLimit);

    uint256 actualGasLimit = s_onRampOverSuperchainInterop.extractGasLimit(extraArgs);

    assertEq(actualGasLimit, expectedGasLimit);
  }

  function test_MaximumGasLimit_GenericExtraArgsV2() public view {
    uint256 expectedGasLimit = type(uint256).max;
    bytes memory extraArgs = _generateGenericExtraArgsV2(expectedGasLimit, true);

    uint256 actualGasLimit = s_onRampOverSuperchainInterop.extractGasLimit(extraArgs);

    assertEq(actualGasLimit, expectedGasLimit);
  }

  function test_ExtraArgsWithAdditionalData() public view {
    uint256 expectedGasLimit = 500000;
    bytes memory additionalData = abi.encode("extra", "data");
    bytes memory extraArgs = _generateExtraArgsWithAdditionalData(expectedGasLimit, additionalData);

    uint256 actualGasLimit = s_onRampOverSuperchainInterop.extractGasLimit(extraArgs);

    assertEq(actualGasLimit, expectedGasLimit);
  }

  function testFuzz_VariousGasLimits_Success(
    uint256 gasLimit
  ) public view {
    assertEq(s_onRampOverSuperchainInterop.extractGasLimit(_generateEVMExtraArgsV1(gasLimit)), gasLimit);
    assertEq(s_onRampOverSuperchainInterop.extractGasLimit(_generateGenericExtraArgsV2(gasLimit, true)), gasLimit);
    assertEq(
      s_onRampOverSuperchainInterop.extractGasLimit(_generateExtraArgsWithAdditionalData(gasLimit, new bytes(100))),
      gasLimit
    );
  }

  // Reverts

  function test_RevertWhen_EmptyExtraArgs() public {
    bytes memory emptyArgs = new bytes(0);

    vm.expectRevert(abi.encodeWithSelector(OnRampOverSuperchainInterop.ExtraArgsTooShort.selector, 0));
    s_onRampOverSuperchainInterop.extractGasLimit(emptyArgs);
  }

  function testFuzz_RevertWhen_ExtraArgsTooShort(
    uint8 length
  ) public {
    vm.assume(length < 36);
    bytes memory shortArgs = _generateShortExtraArgs(length);

    vm.expectRevert(abi.encodeWithSelector(OnRampOverSuperchainInterop.ExtraArgsTooShort.selector, length));
    s_onRampOverSuperchainInterop.extractGasLimit(shortArgs);
  }
}
