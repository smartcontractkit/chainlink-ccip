pragma solidity ^0.8.24;

import {Router} from "../Router.sol";
import {Client} from "../libraries/Client.sol";

import {IERC20} from "@chainlink/vendor/openzeppelin-solidity/v5.0.2/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "@chainlink/vendor/openzeppelin-solidity/v5.0.2/contracts/token/ERC20/utils/SafeERC20.sol";

import {Script} from "forge-std/Script.sol";

/* solhint-disable-next-line no-console */
import {console2 as console} from "forge-std/console2.sol";

/* solhint-disable no-console */
contract CCIPSendTestScript is Script {
  using SafeERC20 for IERC20;

  address public ROUTER;
  address public FEE_TOKEN;

  address public TOKEN0;
  uint256 public TOKEN0_AMOUNT;

  uint64 public DESTINATION_CHAIN_SELECTOR;
  uint64 public SOURCE_CHAIN_SELECTOR;

  // Ex: "ETHEREUM_RPC_URL" as defined in .env
  string public RPC_DESCRIPTOR;

  bytes public s_extraArgs;
  bytes public s_recipient;
  bytes public s_data;

  function run() public {
    vm.createSelectFork(RPC_DESCRIPTOR);

    uint256 privateKey = vm.envUint("PRIVATE_KEY");

    address sender = vm.rememberKey(privateKey);
    s_recipient = abi.encode(sender);

    vm.startBroadcast(privateKey);

    console.log("Sender: %s", sender);
    console.log("Starting Script...");

    // Create the EVMTokenAmount array and populate with the first token
    Client.EVMTokenAmount[] memory tokens;
    if (TOKEN0 != address(0)) {
      tokens = new Client.EVMTokenAmount[](1);
      tokens[0] = Client.EVMTokenAmount({token: TOKEN0, amount: TOKEN0_AMOUNT});
    }

    // Construct the EVM2AnyMessage
    Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
      receiver: abi.encode(sender),
      data: s_data,
      tokenAmounts: tokens,
      feeToken: address(0),
      extraArgs: s_extraArgs
    });

    uint256 fee = Router(ROUTER).getFee(DESTINATION_CHAIN_SELECTOR, message);

    console.log("Fee in WEI: %s", fee);

    console.log("1. Approving Send Tokens...");

    for (uint256 i = 0; i < tokens.length; i++) {
      // Since sender may be an EOA with an existing approval, the allowance should be checked first
      uint256 allowance = IERC20(tokens[i].token).allowance(sender, ROUTER);

      // Approving Tokens for Router if allowance is currently insufficient.
      if (allowance < tokens[i].amount) {
        console.log("Approving %i tokens to Router for %s", tokens[i].amount, tokens[i].token);
        IERC20(tokens[i].token).safeIncreaseAllowance(ROUTER, tokens[i].amount);
      }
    }

    console.log("--- Tokens Approved ---");

    console.log("2. Approving Fee Token");
    if (FEE_TOKEN != address(0)) {
      console.log("Approving Fee Token %s to Router", FEE_TOKEN);
      IERC20(FEE_TOKEN).safeIncreaseAllowance(ROUTER, fee);
    }
    // --- Fee Token Approved ---

    console.log("3. Sending message from: %i to %i", SOURCE_CHAIN_SELECTOR, DESTINATION_CHAIN_SELECTOR);

    // Send the message, forwarding native tokens if necessary to pay the fee
    bytes32 messageId =
      Router(ROUTER).ccipSend{value: FEE_TOKEN == address(0) ? fee : 0}(DESTINATION_CHAIN_SELECTOR, message);

    console.log("--- Message sent: MessageId ---");

    console.logBytes32(messageId);
    vm.stopBroadcast();

    console.log("Script completed...");
  }
}
/* solhint-enable no-console */
