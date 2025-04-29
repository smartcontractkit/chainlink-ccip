pragma solidity ^0.8.24;

import {Router} from "../Router.sol";
import {Internal} from "../libraries/Internal.sol";
import {OffRamp} from "../offRamp/OffRamp.sol";

import {IERC20} from "@chainlink/vendor/openzeppelin-solidity/v5.0.2/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "@chainlink/vendor/openzeppelin-solidity/v5.0.2/contracts/token/ERC20/utils/SafeERC20.sol";
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
  using SafeERC20 for IERC20;

  error ManualExecutionFailed();
  error ManualExecutionNotAllowed();

  function run() public {
    // 1. Define which chain you would like to use (Ex: "ETHEREUM" as defined in .env)
    // [INSERT BELOW]
    string memory chainIdentifier;

    // 2. Retrieve the RPC-URL based on the identifier and defined in .env
    // Ex: "ETHEREUM" -> "ETHEREUM_RPC_URL"
    vm.createSelectFork(string.concat(chainIdentifier, "_RPC_URL"));

    // 3. Acquire the private key from the .env file and derive address
    uint256 privateKey = vm.envUint("PRIVATE_KEY");
    address sender = vm.rememberKey(privateKey);

    console.log("Sender: %s", sender);
    console.log("Starting Script...");

    // 4. Define the selector of the source chain the message originated on.
    // [INSERT BELOW]
    uint64 sourceChainSelector;

    // 5. Define the sequence number of the message to be manually executed. The sequence number can be found
    // on the CCIP-Explorer page for the given message.
    // Note: The OffRamp will use the sequencer number and NOT the messageId to acquire message status, so if this
    // value is set incorrectly, unexpected behavior may result by this script.
    uint64 sequenceNumber;

    // 6. Acquire the address of the OffRamp to send the manual execution call to
    // [INSERT BELOW]
    address router = vm.envAddress(string.concat(chainIdentifier, "_ROUTER"));
    address offRamp;
    Router.OffRamp[] memory offRamps = Router(router).getOffRamps();
    // Perform a linear search for the offRamp contract using the known source chain selector.
    // Note: Given that this operation is performed off-chain, a linear search is an acceptable time-complexity.
    for (uint256 i = 0; i < offRamps.length; ++i) {
      if (offRamps[i].sourceChainSelector == sourceChainSelector) {
        offRamp = offRamps[i].offRamp;
        break;
      }
    }

    // 7. Check that the message status is appropriate for manual execution, revert otherwise.
    Internal.MessageExecutionState executionState =
      OffRamp(offRamp).getExecutionState(sourceChainSelector, sequenceNumber);
    if (
      executionState != Internal.MessageExecutionState.FAILURE
        && executionState != Internal.MessageExecutionState.UNTOUCHED
    ) revert ManualExecutionNotAllowed();

    // 8. Define the manual execution data to be used. Given that manual execution data can be hard to derive manually,
    // due to the existence of offChain data and various proof flags, the data for manual execution should be acquired
    // from the CCIP-Explorer webpage or the ccip-tools-ts repository on Github (https://github.com/smartcontractkit/ccip-tools-ts)
    // Manual Execution Tutorial: https://docs.chain.link/ccip/tutorials/manual-execution#trigger-manual-execution
    // Note: The manual execution data should be formatted as calldata to be sent to the OffRamp and invoke the
    // function "manuallyExecute"
    bytes memory manualExecutionData;

    // 9. Attempt the call to the offRamp to manually execute the message.
    vm.startBroadcast(privateKey);
    (bool success,) = address(offRamp).call(manualExecutionData);
    vm.stopBroadcast();

    // 10. Revert if the execution was not successful.
    if (!success) revert ManualExecutionFailed();
    executionState = OffRamp(offRamp).getExecutionState(sourceChainSelector, sequenceNumber);
    if (executionState != Internal.MessageExecutionState.SUCCESS) revert ManualExecutionFailed();
  }
}
