// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IExecutor} from "../../interfaces/IExecutor.sol";

import {FinalityCodec} from "../../libraries/FinalityCodec.sol";

contract MockExecutor is IExecutor {
  function getAllowedFinalityConfig() external view virtual returns (bytes4) {
    return FinalityCodec.WAIT_FOR_FINALITY_FLAG;
  }

  function getFee(
    uint64, // destChainSelector,
    bytes4, // finalityConfig
    address[] memory, // ccvs,
    bytes memory, // extraArgs
    address // feeToken
  ) external pure returns (uint16 usdCents) {
    return 0;
  }
}
