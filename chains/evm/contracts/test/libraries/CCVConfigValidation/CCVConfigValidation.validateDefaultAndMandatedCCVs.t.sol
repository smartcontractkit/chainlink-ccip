// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {CCVConfigValidation} from "../../../libraries/CCVConfigValidation.sol";
import {Test} from "forge-std/Test.sol";

contract CCVConfigValidationTestHelper {
  function validateDefaultAndMandatedCCVs(address[] memory defaultCCV, address[] memory laneMandatedCCVs) external pure {
    CCVConfigValidation._validateDefaultAndMandatedCCVs(defaultCCV, laneMandatedCCVs);
  }
}

contract CCVConfigValidation_validateDefaultAndMandatedCCVs is Test {
  CCVConfigValidationTestHelper internal s_helper;

  function setUp() public {
    s_helper = new CCVConfigValidationTestHelper();
  }

  function test_validateDefaultAndMandatedCCVs_OnlyDefaultProvided() public view {
    address[] memory defaultCCV = new address[](1);
    defaultCCV[0] = address(0x1);
    s_helper.validateDefaultAndMandatedCCVs(defaultCCV, new address[](0));
  }

  function test_validateDefaultAndMandatedCCVs_OnlyMandatedProvided() public view {
    address[] memory mandated = new address[](1);
    mandated[0] = address(0x1);
    s_helper.validateDefaultAndMandatedCCVs(new address[](0), mandated);
  }

  function test_validateDefaultAndMandatedCCVs_RevertWhen_ZeroInDefault() public {
    address[] memory defaultCCV = new address[](1);
    defaultCCV[0] = address(0);
    address[] memory mandated = new address[](0);

    vm.expectRevert(CCVConfigValidation.ZeroAddressNotAllowed.selector);
    s_helper.validateDefaultAndMandatedCCVs(defaultCCV, mandated);
  }

  function test_validateDefaultAndMandatedCCVs_RevertWhen_ZeroInMandated() public {
    address[] memory defaultCCV = new address[](1);
    defaultCCV[0] = address(0x1);
    address[] memory mandated = new address[](2);
    mandated[0] = address(0x2);
    mandated[1] = address(0);

    vm.expectRevert(CCVConfigValidation.ZeroAddressNotAllowed.selector);
    s_helper.validateDefaultAndMandatedCCVs(defaultCCV, mandated);
  }

  function test_validateDefaultAndMandatedCCVs_RevertWhen_DuplicateWithinDefault() public {
    address dup = address(0xBEEF);
    address[] memory defaultCCV = new address[](2);
    defaultCCV[0] = dup;
    defaultCCV[1] = dup;
    address[] memory mandated = new address[](0);

    vm.expectRevert(abi.encodeWithSelector(CCVConfigValidation.DuplicateCCVNotAllowed.selector, dup));
    s_helper.validateDefaultAndMandatedCCVs(defaultCCV, mandated);
  }

  function test_validateDefaultAndMandatedCCVs_RevertWhen_DuplicateWithinMandated() public {
    address dup = address(0xBEEF);
    address[] memory defaultCCV = new address[](1);
    defaultCCV[0] = address(0x1);
    address[] memory mandated = new address[](2);
    mandated[0] = dup;
    mandated[1] = dup;

    vm.expectRevert(abi.encodeWithSelector(CCVConfigValidation.DuplicateCCVNotAllowed.selector, dup));
    s_helper.validateDefaultAndMandatedCCVs(defaultCCV, mandated);
  }

  function test_validateDefaultAndMandatedCCVs_RevertWhen_DuplicateAcrossDefaultAndMandated() public {
    address dup = address(0xBEEF);
    address[] memory defaultCCV = new address[](1);
    defaultCCV[0] = dup;
    address[] memory mandated = new address[](1);
    mandated[0] = dup;

    vm.expectRevert(abi.encodeWithSelector(CCVConfigValidation.DuplicateCCVNotAllowed.selector, dup));
    s_helper.validateDefaultAndMandatedCCVs(defaultCCV, mandated);
  }

  function test_validateDefaultAndMandatedCCVs_RevertWhen_BothEmpty() public {
    vm.expectRevert(CCVConfigValidation.MustSpecifyDefaultOrRequiredCCVs.selector);
    s_helper.validateDefaultAndMandatedCCVs(new address[](0), new address[](0));
  }
}
