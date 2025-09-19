// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

contract MockRampImplementation {
  error Failed();

  uint8 private s_value;

  constructor(
    uint8 value
  ) {
    s_value = value;
  }

  function getValue() external view returns (uint8) {
    return s_value;
  }

  function revertWithError() external pure {
    revert Failed();
  }
}
