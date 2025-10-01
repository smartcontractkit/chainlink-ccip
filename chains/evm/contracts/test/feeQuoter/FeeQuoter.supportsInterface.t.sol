// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IFeeQuoter} from "../../interfaces/IFeeQuoter.sol";
import {FeeQuoterSetup} from "./FeeQuoterSetup.t.sol";

import {IERC165} from "@openzeppelin/contracts@5.0.2/interfaces/IERC165.sol";

contract FeeQuoter_supportsInterface is FeeQuoterSetup {
  function test_SupportsInterface() public view {
    assertTrue(s_feeQuoter.supportsInterface(type(IFeeQuoter).interfaceId));
    assertTrue(s_feeQuoter.supportsInterface(type(IERC165).interfaceId));
  }
}
