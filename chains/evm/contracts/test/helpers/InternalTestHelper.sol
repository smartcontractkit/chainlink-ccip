// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {Internal} from "../../libraries/Internal.sol";

/// @notice This helper contract exposes the internal functions of the Internal library
contract InternalTestHelper {
  function validateEVMAddress(
    bytes memory encodedAddress
  ) public pure {
    Internal._validateEVMAddress(encodedAddress);
  }

  function validate32ByteAddress(bytes memory encodedAddress, uint256 minValue) public pure {
    Internal._validate32ByteAddress(encodedAddress, minValue);
  }

  function validateTVMAddress(
    bytes memory encodedAddress
  ) public pure {
    Internal._validateTVMAddress(encodedAddress);
  }
}
