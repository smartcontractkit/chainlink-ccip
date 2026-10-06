// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Test} from "forge-std/Test.sol";
import {console} from "forge-std/console.sol";

contract MCMSForkTest is Test {
  struct Call {
    address target;
    uint256 value;
    bytes data;
  }

  error TransactionReverted();

  /// @dev ABI tuple `(Call[], bytes32, bytes32, uint256)`: 4-word head (128) + at least one word for `Call[]` tail (empty array).
  uint256 private constant MIN_MCMS_PAYLOAD_BYTES = 160;

  error PayloadTooShortForMCMSEnvelope(uint256 length, uint256 minLength);

  function _applyPayload(address sender, bytes memory payload) internal {
    console.log("MCMSForkTest._applyPayload: prank sender");
    console.logAddress(sender);
    console.log("MCMSForkTest._applyPayload: payload length (bytes)");
    console.logUint(payload.length);

    if (payload.length < MIN_MCMS_PAYLOAD_BYTES) {
      console.log("MCMSForkTest._applyPayload: payload too short; expected full abi.encode(Call[],bytes32,bytes32,uint256) bytes");
      console.log("MCMSForkTest._applyPayload: min length (bytes)");
      console.logUint(MIN_MCMS_PAYLOAD_BYTES);
      revert PayloadTooShortForMCMSEnvelope(payload.length, MIN_MCMS_PAYLOAD_BYTES);
    }

    MCMSForkTest.Call[] memory calls;
    {
      bytes32 preHash;
      bytes32 postHash;
      uint256 chainId;
      (calls, preHash, postHash, chainId) = abi.decode(payload, (MCMSForkTest.Call[], bytes32, bytes32, uint256));
      console.log("MCMSForkTest._applyPayload: decoded call count");
      console.logUint(calls.length);
      console.log("MCMSForkTest._applyPayload: pre/post hash, chainId");
      console.logBytes32(preHash);
      console.logBytes32(postHash);
      console.logUint(chainId);
    }

    for (uint256 i = 0; i < calls.length; ++i) {
      MCMSForkTest.Call memory call = calls[i];
      console.log("MCMSForkTest._applyPayload: --- call index ---");
      console.logUint(i);
      console.log("MCMSForkTest._applyPayload: target");
      console.logAddress(call.target);
      console.log("MCMSForkTest._applyPayload: value (wei)");
      console.logUint(call.value);
      console.log("MCMSForkTest._applyPayload: calldata length");
      console.logUint(call.data.length);
      if (call.data.length >= 4) {
        console.log("MCMSForkTest._applyPayload: selector");
        console.logBytes4(bytes4(call.data));
      }

      vm.startPrank(sender);
      (bool success, bytes memory returndata) = call.target.call{value: call.value}(call.data);
      vm.stopPrank();

      if (!success) {
        console.log("MCMSForkTest._applyPayload: CALL FAILED at index");
        console.logUint(i);
        console.log("MCMSForkTest._applyPayload: failing calldata");
        console.logBytes(call.data);
        console.log("MCMSForkTest._applyPayload: revert returndata (empty => bare revert)");
        console.logBytes(returndata);
        if (returndata.length > 0) {
          assembly ("memory-safe") {
            revert(add(returndata, 0x20), mload(returndata))
          }
        }
        revert TransactionReverted();
      }
    }
    console.log("MCMSForkTest._applyPayload: all calls succeeded");
  }
}