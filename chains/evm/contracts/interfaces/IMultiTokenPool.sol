// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {IGetPoolConfigManager} from "./IGetPoolConfigManager.sol";
import {IPoolV1V2} from "./IPoolV1V2.sol";

/// @notice Shared interface for CCIP pools that serve multiple tokens.
/// @dev The inherited isSupportedChain(uint64) and getRemoteToken(uint64) views revert with LocalTokenRequired.
/// Query per-token state from the pool config manager returned by getPoolConfigManager.
interface IMultiTokenPool is IPoolV1V2, IGetPoolConfigManager {
  error LocalTokenRequired();
}
