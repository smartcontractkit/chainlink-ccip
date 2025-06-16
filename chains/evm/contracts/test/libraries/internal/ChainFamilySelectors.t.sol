// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Internal} from "../../../libraries/Internal.sol";
import {Test} from "forge-std/Test.sol";

contract Internal_ChainFamilySelectors is Test {
  function test_ChainFamilySelectors_CorrectValues() public pure {
    assertEq(Internal.CHAIN_FAMILY_SELECTOR_EVM, bytes4(keccak256("CCIP ChainFamilySelector EVM")));
    assertEq(Internal.CHAIN_FAMILY_SELECTOR_SVM, bytes4(keccak256("CCIP ChainFamilySelector SVM")));
    assertEq(Internal.CHAIN_FAMILY_SELECTOR_APTOS, bytes4(keccak256("CCIP ChainFamilySelector APTOS")));
    assertEq(Internal.CHAIN_FAMILY_SELECTOR_SUI, bytes4(keccak256("CCIP ChainFamilySelector SUI")));
    assertEq(Internal.CHAIN_FAMILY_SELECTOR_TVM, bytes4(keccak256("CCIP ChainFamilySelector TVM")));
  }
}
