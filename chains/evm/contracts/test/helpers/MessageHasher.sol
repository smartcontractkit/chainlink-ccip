// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.4;

import {Client} from "../../libraries/Client.sol";
import {ExtraArgsCodec} from "../../libraries/ExtraArgsCodec.sol";

contract MessageHasher {
  function encodeEVMExtraArgsV1(
    Client.EVMExtraArgsV1 memory extraArgs
  ) public pure returns (bytes memory) {
    return Client._argsToBytes(extraArgs);
  }

  function encodeEVMExtraArgsV2(
    Client.GenericExtraArgsV2 memory extraArgs
  ) public pure returns (bytes memory) {
    return Client._argsToBytes(extraArgs);
  }

  function encodeGenericExtraArgsV2(
    Client.GenericExtraArgsV2 memory extraArgs
  ) public pure returns (bytes memory) {
    return Client._argsToBytes(extraArgs);
  }

  function encodeGenericExtraArgsV3(
    ExtraArgsCodec.GenericExtraArgsV3 memory extraArgs
  ) public pure returns (bytes memory) {
    return ExtraArgsCodec._encodeGenericExtraArgsV3(extraArgs);
  }

  function decodeEVMExtraArgsV1(
    uint256 gasLimit
  ) public pure returns (Client.EVMExtraArgsV1 memory) {
    return Client.EVMExtraArgsV1(gasLimit);
  }

  function decodeGenericExtraArgsV2(
    uint256 gasLimit,
    bool allowOutOfOrderExecution
  ) public pure returns (Client.GenericExtraArgsV2 memory) {
    return Client.GenericExtraArgsV2({gasLimit: gasLimit, allowOutOfOrderExecution: allowOutOfOrderExecution});
  }

  function decodeEVMExtraArgsV2(
    uint256 gasLimit,
    bool allowOutOfOrderExecution
  ) public pure returns (Client.GenericExtraArgsV2 memory) {
    return Client.GenericExtraArgsV2({gasLimit: gasLimit, allowOutOfOrderExecution: allowOutOfOrderExecution});
  }

  function encodeSVMExtraArgsV1(
    Client.SVMExtraArgsV1 memory extraArgs
  ) public pure returns (bytes memory) {
    return Client._svmArgsToBytes(extraArgs);
  }

  function encodeSUIExtraArgsV1(
    Client.SuiExtraArgsV1 memory extraArgs
  ) public pure returns (bytes memory) {
    return Client._suiArgsToBytes(extraArgs);
  }

  /// @notice used offchain to decode an encoded SVMExtraArgsV1 struct.
  /// @dev The unrolled version fails due to differences in encoding when the accounts[] array
  /// is empty or not.
  function decodeSVMExtraArgsStruct(
    Client.SVMExtraArgsV1 memory extraArgs
  )
    public
    pure
    returns (
      uint32 computeUnits,
      uint64 accountIsWritableBitmap,
      bool allowOutOfOrderExecution,
      bytes32 tokenReceiver,
      bytes32[] memory accounts
    )
  {
    return (
      extraArgs.computeUnits,
      extraArgs.accountIsWritableBitmap,
      extraArgs.allowOutOfOrderExecution,
      extraArgs.tokenReceiver,
      extraArgs.accounts
    );
  }

  /// @notice Used offchain to decode an encoded SuiExtraArgsV1 struct.
  function decodeSuiExtraArgsStruct(
    Client.SuiExtraArgsV1 memory extraArgs
  )
    public
    pure
    returns (uint256 gasLimit, bool allowOutOfOrderExecution, bytes32 tokenReceiver, bytes32[] memory receiverObjectIds)
  {
    return
      (extraArgs.gasLimit, extraArgs.allowOutOfOrderExecution, extraArgs.tokenReceiver, extraArgs.receiverObjectIds);
  }
}
