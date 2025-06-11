// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Internal} from "../../libraries/Internal.sol";
import {InternalTestHelper} from "../helpers/InternalTestHelper.sol";
import {Test} from "forge-std/Test.sol";

contract InternalSetup is Test {
  address internal constant VALID_EVM_ADDRESS = 0x1234567890123456789012345678901234567890;
  address internal constant PRECOMPILE_ADDRESS = address(0x01);
  bytes internal constant VALID_32_BYTE_ADDRESS =
    abi.encode(uint256(0x1234567890123456789012345678901234567890123456789012345678901234));
  bytes internal constant VALID_TVM_ADDRESS =
    hex"11ff1234567890123456789012345678901234567890123456789012345678901234abcd";
  bytes internal constant INVALID_TVM_ADDRESS_ZERO_ACCOUNT =
    hex"11ff000000000000000000000000000000000000000000000000000000000000000012ab";

  InternalTestHelper internal s_helper;

  function setUp() public virtual {
    s_helper = new InternalTestHelper();
  }
}

contract Internal_validateEVMAddress is InternalSetup {
  function test_validateEVMAddress_succeeds_onValidAddress() public {
    bytes memory validAddress = abi.encode(VALID_EVM_ADDRESS);
    s_helper.validateEVMAddress(validAddress);
  }

  function test_validateEVMAddress_reverts_onInvalidLength() public {
    bytes memory invalidAddress = new bytes(31);
    vm.expectRevert(abi.encodeWithSelector(Internal.InvalidEVMAddress.selector, invalidAddress));
    s_helper.validateEVMAddress(invalidAddress);
  }

  function test_validateEVMAddress_reverts_onPrecompileAddress() public {
    bytes memory precompileAddress = abi.encode(PRECOMPILE_ADDRESS);
    vm.expectRevert(abi.encodeWithSelector(Internal.InvalidEVMAddress.selector, precompileAddress));
    s_helper.validateEVMAddress(precompileAddress);
  }

  function test_validateEVMAddress_reverts_onOversizedAddress() public {
    bytes memory invalidAddress = abi.encode(uint256(type(uint160).max) + 1);
    vm.expectRevert(abi.encodeWithSelector(Internal.InvalidEVMAddress.selector, invalidAddress));
    s_helper.validateEVMAddress(invalidAddress);
  }

  function test_validateEVMAddress_succeeds_onBoundaryAddresses() public {
    bytes memory lowerBoundary = abi.encode(uint256(Internal.EVM_PRECOMPILE_SPACE));
    s_helper.validateEVMAddress(lowerBoundary);

    bytes memory upperBoundary = abi.encode(uint256(type(uint160).max));
    s_helper.validateEVMAddress(upperBoundary);
  }
}

contract Internal_validate32ByteAddress is InternalSetup {
  function test_validate32ByteAddress_succeeds_onValidAddress() public {
    s_helper.validate32ByteAddress(VALID_32_BYTE_ADDRESS, 0);
  }

  function test_validate32ByteAddress_reverts_onInvalidLength() public {
    bytes memory invalidAddress = new bytes(31);
    vm.expectRevert(abi.encodeWithSelector(Internal.Invalid32ByteAddress.selector, invalidAddress));
    s_helper.validate32ByteAddress(invalidAddress, 0);
  }

  function test_validate32ByteAddress_reverts_onAddressBelowMinValue() public {
    bytes memory belowMinAddress = abi.encode(uint256(500));
    vm.expectRevert(abi.encodeWithSelector(Internal.Invalid32ByteAddress.selector, belowMinAddress));
    s_helper.validate32ByteAddress(belowMinAddress, 1000);
  }

  function test_validate32ByteAddress_succeeds_onBoundaryMinValue() public {
    uint256 minValue = 1000;
    bytes memory exactMinAddress = abi.encode(minValue);
    s_helper.validate32ByteAddress(exactMinAddress, minValue);
  }

  function test_validate32ByteAddress_reverts_onAptosPrecompileAddress() public {
    bytes memory precompileAddress = abi.encode(Internal.APTOS_PRECOMPILE_SPACE - 1);
    vm.expectRevert(abi.encodeWithSelector(Internal.Invalid32ByteAddress.selector, precompileAddress));
    s_helper.validate32ByteAddress(precompileAddress, Internal.APTOS_PRECOMPILE_SPACE);
  }
}

contract Internal_validateTVMAddress is InternalSetup {
  function test_validateTVMAddress_succeeds_onValidAddress() public {
    s_helper.validateTVMAddress(VALID_TVM_ADDRESS);
  }

  function test_validateTVMAddress_reverts_onShortLength() public {
    bytes memory shortAddress = new bytes(35);
    vm.expectRevert(abi.encodeWithSelector(Internal.InvalidTVMAddress.selector, shortAddress));
    s_helper.validateTVMAddress(shortAddress);
  }

  function test_validateTVMAddress_reverts_onLongLength() public {
    bytes memory longAddress = new bytes(37);
    vm.expectRevert(abi.encodeWithSelector(Internal.InvalidTVMAddress.selector, longAddress));
    s_helper.validateTVMAddress(longAddress);
  }

  function test_validateTVMAddress_reverts_onZeroAccountId() public {
    vm.expectRevert(abi.encodeWithSelector(Internal.InvalidTVMAddress.selector, INVALID_TVM_ADDRESS_ZERO_ACCOUNT));
    s_helper.validateTVMAddress(INVALID_TVM_ADDRESS_ZERO_ACCOUNT);
  }
}

contract Internal_ChainFamilySelectors is InternalSetup {
  function test_ChainFamilySelectors_areCorrect() public pure {
    assertEq(Internal.CHAIN_FAMILY_SELECTOR_EVM, bytes4(keccak256("CCIP ChainFamilySelector EVM")));
    assertEq(Internal.CHAIN_FAMILY_SELECTOR_SVM, bytes4(keccak256("CCIP ChainFamilySelector SVM")));
    assertEq(Internal.CHAIN_FAMILY_SELECTOR_APTOS, bytes4(keccak256("CCIP ChainFamilySelector APTOS")));
    assertEq(Internal.CHAIN_FAMILY_SELECTOR_SUI, bytes4(keccak256("CCIP ChainFamilySelector SUI")));
    assertEq(Internal.CHAIN_FAMILY_SELECTOR_TVM, bytes4(keccak256("CCIP ChainFamilySelector TVM")));
  }
}
