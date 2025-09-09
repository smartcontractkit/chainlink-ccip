// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {Internal} from "../libraries/Internal.sol";
import {MessageV1Codec} from "../libraries/MessageV1Codec.sol";

import {IERC165} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v5.0.2/contracts/utils/introspection/IERC165.sol";

interface ICCVOffRampV1 is IERC165 {
  /// @notice Message execution
  function verifyMessage(
    MessageV1Codec.MessageV1 memory message,
    bytes32 messageHash,
    bytes memory ccvData,
    Internal.MessageExecutionState originalState
  ) external;
}
