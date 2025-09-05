// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

import {MockHyperLiquidCompatibleERC20} from "../../mocks/MockHyperLiquidCompatibleERC20.sol";
import {HyperLiquidCompatibleERC20Setup} from "./HyperLiquidCompatibleERC20Setup.t.sol";

contract HyperLiquidCompatibleERC20_getHyperEVMLinker is HyperLiquidCompatibleERC20Setup {
  MockHyperLiquidCompatibleERC20 public s_testToken;

  function setUp() public override {
    super.setUp();

    s_testToken = new MockHyperLiquidCompatibleERC20("TEST TOKEN", "TEST", 18, 1e27, 0, OWNER);
  }

  function test_getHyperEVMLinker_ReturnsZeroAddress_WhenNotSet() public view {
    address hyperEVMLinker = s_testToken.getHyperEVMLinker();
    assertEq(hyperEVMLinker, address(0));
  }

  function test_getHyperEVMLinker_ReturnsCorrectAddress_WhenSet() public {
    address testHyperEVMLinker = makeAddr("testHyperEVMLinker");
    s_testToken.setHyperEVMLinker(testHyperEVMLinker);

    address hyperEVMLinker = s_testToken.getHyperEVMLinker();
    assertEq(hyperEVMLinker, testHyperEVMLinker);
  }

  function test_getHyperEVMLinker_ReturnsUpdatedAddress_WhenChanged() public {
    address firstLinker = makeAddr("firstLinker");
    address secondLinker = makeAddr("secondLinker");

    s_testToken.setHyperEVMLinker(firstLinker);
    assertEq(s_testToken.getHyperEVMLinker(), firstLinker);

    s_testToken.setHyperEVMLinker(secondLinker);
    assertEq(s_testToken.getHyperEVMLinker(), secondLinker);
  }
}
