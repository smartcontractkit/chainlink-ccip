// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {FactoryBurnMintERC20} from "../../../tokenAdminRegistry/TokenPoolFactory/FactoryBurnMintERC20.sol";

import {HyperEVMLinker} from "../../../tokenAdminRegistry/TokenPoolFactory/HyperEVMLinker.sol";
import {BurnMintERC20Setup} from "./BurnMintERC20Setup.t.sol";

contract FactoryBurnMintERC20_setHyperEVMLinker is BurnMintERC20Setup {
  FactoryBurnMintERC20 public s_factoryBurnMintERC20;

  function setUp() public override {
    super.setUp();

    s_factoryBurnMintERC20 = new FactoryBurnMintERC20("MOCK TOKEN", "MOCK", 18, 1e27, 0, OWNER);
  }

  function test_setHyperEVMLinker() public {
    address testHyperEVMLinker = makeAddr("testHyperEVMLinker");

    vm.expectEmit(true, false, false, false);
    emit HyperEVMLinker.HyperEVMLinkerSet(testHyperEVMLinker);
    s_factoryBurnMintERC20.setHyperEVMLinker(testHyperEVMLinker);

    assertEq(s_factoryBurnMintERC20.getHyperEVMLinker(), testHyperEVMLinker);
  }

  function test_setHyperEVMLinker_RevertWhen_ZeroAddress() public {
    vm.expectRevert(abi.encodeWithSelector(HyperEVMLinker.LinkerAddressCannotBeZero.selector));
    s_factoryBurnMintERC20.setHyperEVMLinker(address(0));
  }
}
