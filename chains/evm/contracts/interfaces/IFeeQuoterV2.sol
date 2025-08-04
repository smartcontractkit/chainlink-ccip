// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

import {Internal} from "../libraries/Internal.sol";
import {IFeeQuoter} from "./IFeeQuoter.sol";

interface IFeeQuoterV2 is IFeeQuoter {
  /// @notice Validates pool return data.
  /// @param destChainSelector Destination chain selector to which the token amounts are sent to.
  /// @param onRampTokenTransfers Token amounts with populated pool return data.
  /// @return destExecDataPerToken Destination chain execution data.
  function processPoolReturnDataNew(
    uint64 destChainSelector,
    Internal.EVMTokenTransfer[] calldata onRampTokenTransfers
  ) external view returns (bytes[] memory destExecDataPerToken);

  function resolveTokenReceiver(
    bytes calldata extraArgs
  ) external view returns (bytes memory tokenReceiver);
}
