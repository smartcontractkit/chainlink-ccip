pragma solidity ^0.8.24;

import {Router} from "../Router.sol";
import {Internal} from "../libraries/Internal.sol";
import {OffRamp} from "../offRamp/OffRamp.sol";
import {Script} from "forge-std/Script.sol";
// solhint-disable-next-line no-console
import {console2 as console} from "forge-std/console2.sol";

// solhint-disable no-console
/// @title CCIPManualExecutionScript
/// @notice A foundry script for manually executing undelivered messages on a destination chain in CCIP.
/// @dev This script has NOT been audited and is NOT intended for production usage. It is intended only for
/// local debugging and testing with existing deployed contracts.
/// @dev Usage: "forge script scripts/CCIPManualExecutionScript.s.sol:CCIPManualExecutionScript -vvvv"
contract CCIPManualExecutionScript is Script {
  error ManualExecutionFailed();
  error ManualExecutionNotAllowed();
  error InvalidRouter();
  error InvalidSourceChainSelector();
  error InvalidSequenceNumber();
  error InvalidManualExecutionData();

  string public rpcUrl;
  address public router;

  uint64 public sourceChainSelector; // The CCIP chain selector the message originated from

  // Define the sequence number of the message to be manually executed. The sequence number can be found
  // on the CCIP-Explorer page for the given message.
  // Note: The OffRamp will use the sequencer number and NOT the messageId to acquire message status, so if this
  // value is set incorrectly, unexpected behavior may result by this script.
  uint64 public sequenceNumber;

  // Define the manual execution data to be used. Given that manual execution data can be hard to derive manually,
  // due to the existence of offChain data and various proof flags, the data for manual execution should be acquired
  // from the CCIP-Explorer webpage or the ccip-tools-ts repository on Github (https://github.com/smartcontractkit/ccip-tools-ts)
  // Manual Execution Tutorial: https://docs.chain.link/ccip/tutorials/manual-execution#trigger-manual-execution
  // Note: The manual execution data should be formatted as calldata to be sent to the OffRamp and invoke the
  // function "manuallyExecute"
  bytes public manualExecutionData;

  function run() public {
    if (router == address(0)) revert InvalidRouter();
    if (sourceChainSelector == 0) revert InvalidSourceChainSelector();
    if (sequenceNumber == 0) revert InvalidSequenceNumber();
    if (manualExecutionData.length == 0) revert InvalidManualExecutionData();

    vm.createSelectFork(rpcUrl);


    // Acquire the private key from the .env file and derive address
    uint256 privateKey = vm.envUint("PRIVATE_KEY");
    address sender = vm.rememberKey(privateKey);

    console.log("Sender: %s", sender);
    console.log("Starting Script...");

    address offRamp;
    Router.OffRamp[] memory offRamps = Router(router).getOffRamps();

    // Perform a linear search for the offRamp contract using the known source chain selector.
    // Note: Given that this operation is being performed off-chain, and thus gas is not a consideration,
    // a linear search is an acceptable time-complexity.
    for (uint256 i = 0; i < offRamps.length; ++i) {
      if (offRamps[i].sourceChainSelector == sourceChainSelector) {
        offRamp = offRamps[i].offRamp;
        break;
      }
    }

    Internal.MessageExecutionState executionState =
      OffRamp(offRamp).getExecutionState(sourceChainSelector, sequenceNumber);
    if (
      executionState != Internal.MessageExecutionState.FAILURE
        && executionState != Internal.MessageExecutionState.UNTOUCHED
    ) revert ManualExecutionNotAllowed();

    // Attempt the call to the offRamp to manually execute the message.
    vm.startBroadcast(privateKey);
    (bool success,) = address(offRamp).call(manualExecutionData);
    vm.stopBroadcast();

    // Revert if the execution was not successful.
    if (!success) revert ManualExecutionFailed();
    executionState = OffRamp(offRamp).getExecutionState(sourceChainSelector, sequenceNumber);
    if (executionState != Internal.MessageExecutionState.SUCCESS) revert ManualExecutionFailed();
  }
}
