// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

import {HyperLiquidCompatibleERC20} from "../../../tokenAdminRegistry/TokenPoolFactory/HyperLiquidCompatibleERC20.sol";
import {HyperLiquidCompatibleERC20Setup} from "./HyperLiquidCompatibleERC20Setup.t.sol";

contract HyperLiquidCompatibleERC20_setHyperEVMLinker is HyperLiquidCompatibleERC20Setup {
  bytes32 public constant HYPER_EVM_LINKER_SLOT = 0x8c306a6a12fff1951878e8621be6674add1102cd359dd968efbbe797629ef84f;

  function test_setHyperEVMLinker_Success() public {
    address testHyperEVMLinker = makeAddr("testHyperEVMLinker");

    vm.expectEmit();
    emit HyperLiquidCompatibleERC20.HyperEVMLinkerSet(testHyperEVMLinker);
    s_hyperLiquidToken.setHyperEVMLinker(testHyperEVMLinker);

    assertEq(s_hyperLiquidToken.getHyperEVMLinker(), testHyperEVMLinker);

    // Manually inspect the storage at the HYPER_EVM_LINKER_SLOT to ensure the address is stored correctly
    bytes32 storedValue = vm.load(address(s_hyperLiquidToken), HYPER_EVM_LINKER_SLOT);
    assertEq(address(uint160(uint256(storedValue))), testHyperEVMLinker);
  }

  // Reverts

  function test_setHyperEVMLinker_RevertWhen_ZeroAddress() public {
    vm.expectRevert(abi.encodeWithSelector(HyperLiquidCompatibleERC20.LinkerAddressCannotBeZero.selector));
    s_hyperLiquidToken.setHyperEVMLinker(address(0));
  }
}
