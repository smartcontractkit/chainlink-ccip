// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Client} from "../../../libraries/Client.sol";
import {OnRampOverSuperchainInterop} from "../../../onRamp/OnRampOverSuperchainInterop.sol";
import {OnRampOverSuperchainInteropSetup} from "./OnRampOverSuperchainInteropSetup.t.sol";

contract OnRampOverSuperchainInterop_extractGasLimit is OnRampOverSuperchainInteropSetup {
  function test_extractGasLimit() public view {
    uint256 expectedGasLimit = 200000;
    bytes memory extraArgs = Client._argsToBytes(Client.EVMExtraArgsV1({gasLimit: expectedGasLimit}));

    assertEq(expectedGasLimit, s_onRampOverSuperchainInterop.extractGasLimit(extraArgs));

    extraArgs =
      Client._argsToBytes(Client.GenericExtraArgsV2({gasLimit: expectedGasLimit, allowOutOfOrderExecution: false}));
    assertEq(expectedGasLimit, s_onRampOverSuperchainInterop.extractGasLimit(extraArgs));
  }

  function testFuzz_extractGasLimit_VariousGasLimitAndDataSize(uint256 gasLimit, uint256 length) public view {
    length = bound(length, 0, 10_000);

    assertEq(
      gasLimit,
      s_onRampOverSuperchainInterop.extractGasLimit(
        abi.encodeWithSelector(bytes4(keccak256("extraArgsV3")), gasLimit, new bytes(length))
      )
    );
  }

  // Reverts

  function test_extractGasLimit_RevertWhen_ExtraArgsTooShort() public {
    vm.expectRevert(abi.encodeWithSelector(OnRampOverSuperchainInterop.ExtraArgsTooShort.selector, 0));
    s_onRampOverSuperchainInterop.extractGasLimit(new bytes(0));

    vm.expectRevert(abi.encodeWithSelector(OnRampOverSuperchainInterop.ExtraArgsTooShort.selector, 35));
    s_onRampOverSuperchainInterop.extractGasLimit(new bytes(35));
  }
}
