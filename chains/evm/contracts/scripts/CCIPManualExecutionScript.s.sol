pragma solidity ^0.8.24;

import {Internal} from "../libraries/Internal.sol";
import {OffRamp} from "../offRamp/OffRamp.sol";

import {IERC20} from "@vendor/openzeppelin-solidity/v5.0.2/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "@vendor/openzeppelin-solidity/v5.0.2/contracts/token/ERC20/utils/SafeERC20.sol";

import {Script} from "forge-std/Script.sol";
import {Test} from "forge-std/Test.sol";

// solhint-disable-next-line no-console
import {console2 as console} from "forge-std/console2.sol";

/* solhint-disable no-console */
contract CCIPSendTestScript is Script, Test {
  using SafeERC20 for IERC20;

  // Ex: "ETHEREUM_RPC_URL" as defined in .env
  string public RPC_IDENTIFIER;

  OffRamp public s_offRamp;

  bytes32 public s_messageId;

  uint64 public s_sourceChainSelector;
  uint64 public s_sequenceNumber;
  bytes public s_manualExecutionData;

  function run() public {
    vm.createSelectFork(RPC_IDENTIFIER);

    uint256 privateKey = vm.envUint("PRIVATE_KEY");

    address sender = vm.rememberKey(privateKey);

    vm.startBroadcast(privateKey);

    console.log("Sender: %s", sender);
    console.log("Starting Script...");

    // Check that the messageId is not empty
    Internal.MessageExecutionState executionState = s_offRamp.getExecutionState(s_sourceChainSelector, s_sequenceNumber);
    assertTrue(
      executionState == Internal.MessageExecutionState.FAILURE
        || executionState == Internal.MessageExecutionState.UNTOUCHED,
      "Message is not ready for Manual Execution"
    );

    // Manual Execution data can be invoked from a different tool or front-end to avoid having to
    // gather execution report data manually
    (bool success,) = address(s_offRamp).call(s_manualExecutionData);
    assertTrue(success, "Manual execution call reverted");

    executionState = s_offRamp.getExecutionState(s_sourceChainSelector, s_sequenceNumber);
    assertTrue(executionState == Internal.MessageExecutionState.SUCCESS, "Message was not executed successfully");

    console.log("Script completed...");
  }
}
/* solhint-enable no-console */
