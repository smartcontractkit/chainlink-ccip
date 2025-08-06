// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {FactoryBurnMintERC20} from "../../../tokenAdminRegistry/TokenPoolFactory/FactoryBurnMintERC20.sol";

import {HyperEVMLinker} from "../../../tokenAdminRegistry/TokenPoolFactory/HyperEVMLinker.sol";
import {BurnMintERC20Setup} from "./BurnMintERC20Setup.t.sol";

contract FactoryBurnMintERC20_setHyperEVMLinker is BurnMintERC20Setup {
  bytes32 public constant HYPER_EVM_LINKER_SLOT = 0x8c306a6a12fff1951878e8621be6674add1102cd359dd968efbbe797629ef84f;

  FactoryBurnMintERC20 public s_factoryBurnMintERC20;

  function setUp() public override {
    super.setUp();

    s_factoryBurnMintERC20 = new FactoryBurnMintERC20("MOCK TOKEN", "MOCK", 18, 1e27, 0, OWNER);
  }

  function test_setHyperEVMLinker_Success() public {
    address testHyperEVMLinker = makeAddr("testHyperEVMLinker");

    vm.expectEmit();
    emit HyperEVMLinker.HyperEVMLinkerSet(testHyperEVMLinker);
    s_factoryBurnMintERC20.setHyperEVMLinker(testHyperEVMLinker);

    assertEq(s_factoryBurnMintERC20.getHyperEVMLinker(), testHyperEVMLinker);

    // Manually inspect the storage at the HYPER_EVM_LINKER_SLOT to ensure the address is stored correctly
    bytes32 storedValue = vm.load(address(s_factoryBurnMintERC20), HYPER_EVM_LINKER_SLOT);
    assertEq(address(uint160(uint256(storedValue))), testHyperEVMLinker);
  }

  // Reverts

  function test_setHyperEVMLinker_RevertWhen_ZeroAddress() public {
    vm.expectRevert(abi.encodeWithSelector(HyperEVMLinker.LinkerAddressCannotBeZero.selector));
    s_factoryBurnMintERC20.setHyperEVMLinker(address(0));
  }
}
