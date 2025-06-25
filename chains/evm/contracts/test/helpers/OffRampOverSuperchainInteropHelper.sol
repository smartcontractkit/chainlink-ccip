// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Internal} from "../../libraries/Internal.sol";
import {OffRampOverSuperchainInterop} from "../../offRamp/OffRampOverSuperchainInterop.sol";

contract OffRampOverSuperchainInteropHelper is OffRampOverSuperchainInterop {
  constructor(
    StaticConfig memory staticConfig,
    DynamicConfig memory dynamicConfig,
    SourceChainConfigArgs[] memory sourceChainConfigs,
    address[] memory allowedTransmitters,
    ChainSelectorToChainIdConfigArgs[] memory chainIdConfigs
  ) OffRampOverSuperchainInterop(
    staticConfig,
    dynamicConfig,
    sourceChainConfigs,
    allowedTransmitters,
    chainIdConfigs
  ) {}

  function decodeLogDataIntoMessageExposed(
    bytes calldata logData
  ) external view returns (Internal.Any2EVMRampMessage memory) {
    return decodeLogDataIntoMessage(logData);
  }
}