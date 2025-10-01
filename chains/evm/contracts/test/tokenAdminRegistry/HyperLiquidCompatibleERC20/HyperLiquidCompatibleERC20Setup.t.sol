// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

import {BaseTest} from "../../BaseTest.t.sol";
import {MockHyperLiquidCompatibleERC20} from "../../mocks/MockHyperLiquidCompatibleERC20.sol";

contract HyperLiquidCompatibleERC20Setup is BaseTest {
  MockHyperLiquidCompatibleERC20 internal s_hyperLiquidToken;

  address internal s_mockPool = makeAddr("s_mockPool");
  address internal s_hyperEVMLinker = makeAddr("s_hyperEVMLinker");
  uint64 internal s_remoteTokenId = 325;
  address internal s_hypercoreTokenSystemAddress = address((uint160(0x20) << 152) | uint160(s_remoteTokenId));

  function setUp() public virtual override {
    BaseTest.setUp();

    s_hyperLiquidToken =
      new MockHyperLiquidCompatibleERC20("HyperLiquid Token", "HLT", 18, type(uint256).max, type(uint256).max, OWNER);

    assertEq(s_hyperLiquidToken.balanceOf(OWNER), type(uint256).max);

    // Set up HyperEVM linker and remote token
    s_hyperLiquidToken.setHyperEVMLinker(s_hyperEVMLinker);

    assertEq(s_hyperLiquidToken.typeAndVersion(), "HyperLiquidCompatibleERC20 1.6.2");
  }
}
