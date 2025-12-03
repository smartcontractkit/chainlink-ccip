// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {USDCSourcePoolDataCodec} from "../../../libraries/USDCSourcePoolDataCodec.sol";
import {Test} from "forge-std/Test.sol";

contract USDCSourcePoolDataCodec__calculateDepositHash is Test {
  function test__calculateDepositHash() public {
    uint32 sourceDomain = 1553252;
    uint256 amount = 1e6;
    uint32 destinationDomain = 1;
    bytes32 mintRecipient = bytes32(abi.encode(makeAddr("mintRecipient")));
    bytes32 burnToken = bytes32(abi.encode(makeAddr("burnToken")));
    bytes32 destinationCaller = bytes32(abi.encode(makeAddr("destinationCaller")));
    uint256 maxFee = 0;
    uint32 minFinalityThreshold = 2000;

    bytes32 expectedHash = keccak256(
      abi.encode(
        sourceDomain,
        amount,
        destinationDomain,
        mintRecipient,
        burnToken,
        destinationCaller,
        maxFee,
        minFinalityThreshold
      )
    );

    bytes32 actualHash = USDCSourcePoolDataCodec._calculateDepositHash(
      sourceDomain, amount, destinationDomain, mintRecipient, burnToken, destinationCaller, maxFee, minFinalityThreshold
    );

    assertEq(expectedHash, actualHash, "Deposit hash calculation mismatch");
  }
}
