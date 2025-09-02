// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Internal} from "../../libraries/Internal.sol";
import {OffRampOverSuperchainInterop} from "../../offRamp/OffRampOverSuperchainInterop.sol";
import {Identifier} from "../../vendor/optimism/interop-lib/v0/src/interfaces/IIdentifier.sol";

contract OffRampOverSuperchainInteropHelper is OffRampOverSuperchainInterop {
  constructor(
    StaticConfig memory staticConfig,
    DynamicConfig memory dynamicConfig,
    SourceChainConfigArgs[] memory sourceChainConfigs,
    address crossL2Inbox,
    ChainSelectorToChainIdConfigArgs[] memory chainSelectorToChainIdConfigArgs
  )
    OffRampOverSuperchainInterop(
      staticConfig,
      dynamicConfig,
      sourceChainConfigs,
      crossL2Inbox,
      chainSelectorToChainIdConfigArgs
    )
  {}

  /// @notice Exposes the internal _constructProofs function as public for testing
  function constructProofs(
    Internal.Any2EVMRampMessage memory message,
    bytes32[] memory proofs
  ) public pure returns (Identifier memory identifier, bytes32 logHash) {
    return _constructProofs(message, proofs);
  }

  /// @notice Exposes the internal _verifyReport function as public for testing
  function verifyMessage(
    uint64 sourceChainSelector,
    Internal.ExecutionReport memory report
  ) public returns (uint256 timestampCommitted, bytes32[] memory hashedLeaves) {
    return _verifyReport(sourceChainSelector, report);
  }

  /// @notice Exposes the internal _executeSingleReport function as public for testing
  function executeSingleReport(
    Internal.ExecutionReport memory rep,
    GasLimitOverride[] memory manualExecGasExecOverrides
  ) public {
    _executeSingleReport(rep, manualExecGasExecOverrides);
  }
}
