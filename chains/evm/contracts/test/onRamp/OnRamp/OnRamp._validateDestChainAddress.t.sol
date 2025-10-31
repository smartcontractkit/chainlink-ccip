// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {OnRamp} from "../../../onRamp/OnRamp.sol";
import {OnRampTestHelper} from "../../helpers/OnRampTestHelper.sol";
import {OnRampSetup} from "./OnRampSetup.t.sol";

contract OnRamp_validateDestChainAddress is OnRampSetup {
  OnRampTestHelper internal s_onRampTestHelper;

  function setUp() public override {
    super.setUp();
    s_onRampTestHelper = new OnRampTestHelper(
      OnRamp.StaticConfig({
        chainSelector: SOURCE_CHAIN_SELECTOR,
        rmnRemote: s_mockRMNRemote,
        tokenAdminRegistry: address(s_tokenAdminRegistry)
      }),
      OnRamp.DynamicConfig({
        feeQuoter: address(s_feeQuoter),
        reentrancyGuardEntered: false,
        feeAggregator: FEE_AGGREGATOR
      })
    );
  }

  function test_validateDestChainAddress_20ByteExactMatch() public {
    bytes memory rawAddress = abi.encodePacked(makeAddr("receiver"));
    uint8 addressBytesLength = 20;

    bytes memory validated = s_onRampTestHelper.validateDestChainAddress(rawAddress, addressBytesLength);

    assertEq(rawAddress, validated);
  }

  function testFuzz_validateDestChainAddress(
    bytes calldata rawAddress
  ) public view {
    vm.assume(rawAddress.length > 0);

    uint8 addressBytesLength = uint8(rawAddress.length);

    bytes memory validated = s_onRampTestHelper.validateDestChainAddress(rawAddress, addressBytesLength);

    assertEq(rawAddress, validated);
  }

  function test_validateDestChainAddress_20ByteAbiEncoded() public {
    address addr = makeAddr("receiver");
    bytes memory abiEncodedAddress = abi.encode(addr);
    uint8 addressBytesLength = 20;

    bytes memory validated = s_onRampTestHelper.validateDestChainAddress(abiEncodedAddress, addressBytesLength);

    assertEq(abi.encodePacked(addr), validated);
    assertEq(20, validated.length);
  }

  function testFuzz_validateDestChainAddress_AbiEncoded(
    uint8 addressBytesLength
  ) public view {
    addressBytesLength = uint8(bound(addressBytesLength, 1, 31));

    bytes memory actualAddress = new bytes(addressBytesLength);
    for (uint256 i = 0; i < addressBytesLength; i++) {
      actualAddress[i] = bytes1(uint8(i + 1));
    }

    bytes memory paddedAddress = new bytes(32);
    uint256 paddingLength = 32 - addressBytesLength;
    for (uint256 i = 0; i < addressBytesLength; i++) {
      paddedAddress[paddingLength + i] = actualAddress[i];
    }

    bytes memory validated = s_onRampTestHelper.validateDestChainAddress(paddedAddress, addressBytesLength);

    assertEq(actualAddress, validated);
    assertEq(addressBytesLength, validated.length);
  }

  // Reverts

  function test_validateDestChainAddress_RevertWhen_NonZeroPadding() public {
    uint8 addressBytesLength = 20;
    bytes memory paddedAddress = new bytes(32);
    // Set first byte to non-zero (invalid padding)
    paddedAddress[0] = 0x01;
    // Fill the address part with valid data
    for (uint256 i = 12; i < 32; i++) {
      paddedAddress[i] = bytes1(uint8(i));
    }

    vm.expectRevert(abi.encodeWithSelector(OnRamp.InvalidDestChainAddress.selector, paddedAddress));
    s_onRampTestHelper.validateDestChainAddress(paddedAddress, addressBytesLength);
  }

  function test_validateDestChainAddress_RevertWhen_InvalidDestChainAddress_LengthTooShort() public {
    bytes memory rawAddress = new bytes(31);
    uint8 addressBytesLength = 32;

    vm.expectRevert(abi.encodeWithSelector(OnRamp.InvalidDestChainAddress.selector, rawAddress));
    s_onRampTestHelper.validateDestChainAddress(rawAddress, addressBytesLength);
  }

  function test_validateDestChainAddress_RevertWhen_InvalidDestChainAddress_LengthTooLong() public {
    bytes memory rawAddress = new bytes(33);
    uint8 addressBytesLength = 32;

    vm.expectRevert(abi.encodeWithSelector(OnRamp.InvalidDestChainAddress.selector, rawAddress));
    s_onRampTestHelper.validateDestChainAddress(rawAddress, addressBytesLength);
  }

  function test_validateDestChainAddress_RevertWhen_21BytesFor20ByteAddress() public {
    bytes memory rawAddress = new bytes(21);
    uint8 addressBytesLength = 20;

    vm.expectRevert(abi.encodeWithSelector(OnRamp.InvalidDestChainAddress.selector, rawAddress));
    s_onRampTestHelper.validateDestChainAddress(rawAddress, addressBytesLength);
  }
}
