// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ExtraArgsCodec} from "../../libraries/ExtraArgsCodec.sol";

contract ExtraArgsCodecHelper {
  function _decodeGenericExtraArgsV3(
    bytes calldata data
  ) external pure returns (ExtraArgsCodec.GenericExtraArgsV3 memory) {
    return ExtraArgsCodec._decodeGenericExtraArgsV3(data);
  }

  function _decodeSVMExecutorArgsV1(
    bytes calldata data
  ) external pure returns (ExtraArgsCodec.SVMExecutorArgsV1 memory) {
    return ExtraArgsCodec._decodeSVMExecutorArgsV1(data);
  }

  function _decodeSuiExecutorArgsV1(
    bytes calldata data
  ) external pure returns (ExtraArgsCodec.SuiExecutorArgsV1 memory) {
    return ExtraArgsCodec._decodeSuiExecutorArgsV1(data);
  }
}
