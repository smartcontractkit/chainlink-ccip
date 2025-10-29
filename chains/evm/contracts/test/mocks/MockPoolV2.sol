// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolV2} from "../../interfaces/IPoolV2.sol";

import {IERC165} from "@openzeppelin/contracts@5.0.2/utils/introspection/IERC165.sol";

contract MockPoolV2 {
  address[] internal s_requiredCCVs;

  constructor(
    address[] memory requiredCCVs
  ) {
    s_requiredCCVs = requiredCCVs;
  }

  function getRequiredCCVs(
    address,
    uint64,
    uint256,
    uint16,
    bytes memory,
    IPoolV2.MessageDirection
  ) external view returns (address[] memory) {
    return s_requiredCCVs;
  }

  function supportsInterface(
    bytes4 interfaceId
  ) external pure returns (bool) {
    return interfaceId == type(IPoolV2).interfaceId || interfaceId == type(IERC165).interfaceId;
  }
}
