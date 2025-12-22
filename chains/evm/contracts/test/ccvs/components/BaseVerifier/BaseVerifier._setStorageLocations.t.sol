// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {BaseVerifier} from "../../../../ccvs/components/BaseVerifier.sol";
import {BaseVerifierSetup} from "./BaseVerifierSetup.t.sol";

contract BaseVerifier_setStorageLocations is BaseVerifierSetup {
  function test_setStorageLocation_SomeToNone() external {
    string[] memory oldLocations = s_baseVerifier.getStorageLocations();
    string[] memory none = _empty();

    _expectStorageLocationEvent(oldLocations, none);
    s_baseVerifier.setStorageLocations(none);

    _assertEq(s_baseVerifier.getStorageLocations(), none);
  }

  function test_setStorageLocation_NoneToSome() external {
    // Move to none first (no expectations)
    s_baseVerifier.setStorageLocations(_empty());

    string[] memory none = _empty();
    string[] memory some = _one("loc-1");

    _expectStorageLocationEvent(none, some);
    s_baseVerifier.setStorageLocations(some);

    _assertEq(s_baseVerifier.getStorageLocations(), some);
  }

  function test_setStorageLocation_SomeToMany() external {
    string[] memory oldLocations = s_baseVerifier.getStorageLocations(); // some (1 entry)
    string[] memory many = _three("loc-1", "loc-2", "loc-3");

    _expectStorageLocationEvent(oldLocations, many);
    s_baseVerifier.setStorageLocations(many);

    _assertEq(s_baseVerifier.getStorageLocations(), many);
  }

  function test_setStorageLocation_ManyToSome() external {
    string[] memory many = _three("loc-1", "loc-2", "loc-3");
    s_baseVerifier.setStorageLocations(many);

    string[] memory some = _one("loc-4");

    _expectStorageLocationEvent(many, some);
    s_baseVerifier.setStorageLocations(some);

    _assertEq(s_baseVerifier.getStorageLocations(), some);
  }

  function test_setStorageLocation_SomeToSome() external {
    string[] memory someOld = s_baseVerifier.getStorageLocations(); // initial some
    string[] memory someNew = _one("loc-2");

    _expectStorageLocationEvent(someOld, someNew);
    s_baseVerifier.setStorageLocations(someNew);

    _assertEq(s_baseVerifier.getStorageLocations(), someNew);
  }

  function test_setStorageLocation_ManyToMany() external {
    string[] memory manyOld = _three("loc-1", "loc-2", "loc-3");
    s_baseVerifier.setStorageLocations(manyOld);

    string[] memory manyNew = _three("loc-a", "loc-b", "loc-c");

    _expectStorageLocationEvent(manyOld, manyNew);
    s_baseVerifier.setStorageLocations(manyNew);

    _assertEq(s_baseVerifier.getStorageLocations(), manyNew);
  }

  function test_setStorageLocation_Noop() external {
    string[] memory same = _three("loc-1", "loc-2", "loc-3");
    s_baseVerifier.setStorageLocations(same);

    _expectStorageLocationEvent(same, same);
    s_baseVerifier.setStorageLocations(same);

    _assertEq(s_baseVerifier.getStorageLocations(), same);
  }

  // Helpers

  function _expectStorageLocationEvent(string[] memory oldLocations, string[] memory newLocations) internal {
    vm.expectEmit(address(s_baseVerifier));
    emit BaseVerifier.StorageLocationsUpdated(oldLocations, newLocations);
  }

  function _empty() internal pure returns (string[] memory arr) {
    arr = new string[](0);
    return arr;
  }

  function _one(
    string memory a
  ) internal pure returns (string[] memory arr) {
    arr = new string[](1);
    arr[0] = a;
    return arr;
  }

  function _three(string memory a, string memory b, string memory c) internal pure returns (string[] memory arr) {
    arr = new string[](3);
    arr[0] = a;
    arr[1] = b;
    arr[2] = c;
    return arr;
  }

  function _assertEq(string[] memory a, string[] memory b) internal pure {
    assertEq(a.length, b.length);
    for (uint256 i; i < a.length; ++i) {
      assertEq(keccak256(bytes(a[i])), keccak256(bytes(b[i])));
    }
  }
}
