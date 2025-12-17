// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IAny2EVMMessageReceiver} from "../../../interfaces/IAny2EVMMessageReceiver.sol";

import {OffRampSetup} from "./OffRampSetup.t.sol";

import {IERC165} from "@openzeppelin/contracts@5.3.0/utils/introspection/IERC165.sol";

contract OffRamp_isTokenOnlyTransfer is OffRampSetup {
  address internal s_receiver = makeAddr("receiver");

  function test_isTokenOnlyTransfer_TrueWhen_EmptyDataAndZeroGasLimit() public view {
    // data.length == 0 && ccipReceiveGasLimit == 0 -> true
    assertTrue(s_offRamp.checkIsTokenOnlyTransfer(0, 0, s_receiver));
  }

  function test_isTokenOnlyTransfer_TrueWhen_ReceiverHasNoCode() public view {
    // receiver.code.length == 0 -> true, regardless of other params
    assertTrue(s_offRamp.checkIsTokenOnlyTransfer(100, 100_000, s_receiver));
  }

  function test_isTokenOnlyTransfer_TrueWhen_ReceiverDoesNotSupportInterface() public {
    // Etch bytecode to receiver
    vm.etch(s_receiver, hex"60006000f3");

    // Mock supportsInterface to return true for IERC165 but false for IAny2EVMMessageReceiver
    vm.mockCall(s_receiver, abi.encodeCall(IERC165.supportsInterface, (type(IERC165).interfaceId)), abi.encode(true));
    vm.mockCall(
      s_receiver,
      abi.encodeCall(IERC165.supportsInterface, (type(IAny2EVMMessageReceiver).interfaceId)),
      abi.encode(false)
    );

    // Even with data and gas limit, if receiver doesn't support interface -> true
    assertTrue(s_offRamp.checkIsTokenOnlyTransfer(100, 100_000, s_receiver));
  }

  function test_isTokenOnlyTransfer_FalseWhen_ReceiverSupportsInterface_WithDataAndGas() public {
    // Etch bytecode to receiver
    vm.etch(s_receiver, hex"60006000f3");

    // Mock supportsInterface to return true for both IERC165 and IAny2EVMMessageReceiver
    vm.mockCall(s_receiver, abi.encodeCall(IERC165.supportsInterface, (type(IERC165).interfaceId)), abi.encode(true));
    vm.mockCall(
      s_receiver,
      abi.encodeCall(IERC165.supportsInterface, (type(IAny2EVMMessageReceiver).interfaceId)),
      abi.encode(true)
    );

    // With data, gas limit, and receiver supporting interface -> false (not token-only)
    assertFalse(s_offRamp.checkIsTokenOnlyTransfer(100, 100_000, s_receiver));
  }

  function test_isTokenOnlyTransfer_TrueWhen_ReceiverSupportsInterface_ButNoDataAndNoGas() public {
    // Etch bytecode to receiver
    vm.etch(s_receiver, hex"60006000f3");

    // Mock supportsInterface to return true for both IERC165 and IAny2EVMMessageReceiver
    vm.mockCall(s_receiver, abi.encodeCall(IERC165.supportsInterface, (type(IERC165).interfaceId)), abi.encode(true));
    vm.mockCall(
      s_receiver,
      abi.encodeCall(IERC165.supportsInterface, (type(IAny2EVMMessageReceiver).interfaceId)),
      abi.encode(true)
    );

    // Even though receiver supports interface, no data and no gas -> true (token-only)
    assertTrue(s_offRamp.checkIsTokenOnlyTransfer(0, 0, s_receiver));
  }

  function test_isTokenOnlyTransfer_FalseWhen_HasDataOnly() public {
    // Etch bytecode to receiver
    vm.etch(s_receiver, hex"60006000f3");

    // Mock supportsInterface to return true for both interfaces
    vm.mockCall(s_receiver, abi.encodeCall(IERC165.supportsInterface, (type(IERC165).interfaceId)), abi.encode(true));
    vm.mockCall(
      s_receiver,
      abi.encodeCall(IERC165.supportsInterface, (type(IAny2EVMMessageReceiver).interfaceId)),
      abi.encode(true)
    );

    // Has data but no gas limit -> still not token-only because data is present
    assertFalse(s_offRamp.checkIsTokenOnlyTransfer(100, 0, s_receiver));
  }

  function test_isTokenOnlyTransfer_FalseWhen_HasGasLimitOnly() public {
    // Etch bytecode to receiver
    vm.etch(s_receiver, hex"60006000f3");

    // Mock supportsInterface to return true for both interfaces
    vm.mockCall(s_receiver, abi.encodeCall(IERC165.supportsInterface, (type(IERC165).interfaceId)), abi.encode(true));
    vm.mockCall(
      s_receiver,
      abi.encodeCall(IERC165.supportsInterface, (type(IAny2EVMMessageReceiver).interfaceId)),
      abi.encode(true)
    );

    // Has gas limit but no data -> still not token-only because gas limit indicates receiver call
    assertFalse(s_offRamp.checkIsTokenOnlyTransfer(0, 100_000, s_receiver));
  }
}
