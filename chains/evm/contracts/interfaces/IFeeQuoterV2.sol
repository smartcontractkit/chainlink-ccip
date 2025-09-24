// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

import {IFeeQuoter} from "./IFeeQuoter.sol";

interface IFeeQuoterV2 is IFeeQuoter {
  function resolveTokenReceiver(
    bytes calldata extraArgs
  ) external view returns (bytes memory tokenReceiver);
}
