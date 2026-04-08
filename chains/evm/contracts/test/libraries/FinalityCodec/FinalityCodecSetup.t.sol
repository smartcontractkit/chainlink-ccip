// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {FinalityCodecHelper} from "../../helpers/FinalityCodecHelper.sol";
import {Test} from "forge-std/Test.sol";

contract FinalityCodecSetup is Test {
  FinalityCodecHelper internal s_helper;

  function setUp() public virtual {
    s_helper = new FinalityCodecHelper();
  }
}
