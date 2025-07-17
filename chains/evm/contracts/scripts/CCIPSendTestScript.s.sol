pragma solidity ^0.8.24;

import {Router} from "../Router.sol";
import {Client} from "../libraries/Client.sol";

import {IERC20} from "@chainlink/vendor/openzeppelin-solidity/v5.0.2/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "@chainlink/vendor/openzeppelin-solidity/v5.0.2/contracts/token/ERC20/utils/SafeERC20.sol";
import {Script} from "forge-std/Script.sol";

// solhint-disable-next-line no-console
import {console2 as console} from "forge-std/console2.sol";

/* solhint-disable no-console */
/// @title CCIPSendTestScript
/// @notice This is a foundry script for sending messages through CCIP.
/// @dev This script has NOT been audited, and is NOT intended for use in production. It is intended to aid in
/// local debugging and testing with existing deployed contracts.
/// @dev Usage: "forge script scripts/CCIPSendTestScript.s.sol:CCIPSendTestScript"
contract CCIPSendTestScript is Script {
  using SafeERC20 for IERC20;

  error ChainNotSupported(uint64 destChainSelector);

  // For the script to run successfully, please define the following constants below
  // [REQUIRED]
  string public chainIdentifier; // The Chain to use (Ex: "ETHEREUM" as defined in the .env)
  uint64 public destChainSelector; // The CCIP-Specific chain selector to send the message to
  uint256 public numTokens; // The number of tokens to be sent

  bytes public recipient; // The recipient to receive both the tokens and the arbitrary data.
  bytes public data; // Define the message data to be passed to the recipient if it is NOT an EOA
  bytes public extraArgs; // If any extraArgs are needed, define below. Due to different chains families having different extraArgs formats, they should be passed as raw bytes, and encoded separately.

  address public feeToken; // The token to pay CCIP Message Fees in, address(0) for native.

  function run() public {
    vm.createSelectFork(string.concat(chainIdentifier, "_RPC_URL"));

    // Acquire the private key from the .env file and derive address
    uint256 privateKey = vm.envUint("PRIVATE_KEY");
    address sender = vm.rememberKey(privateKey);

    vm.startBroadcast(privateKey);

    console.log("Sender: %s", sender);
    console.log("Starting Script...");

    Client.EVM2AnyMessage memory message;

    // Get the router using the chain identifier Ex: "ETHEREUM" -> "ETHEREUM_ROUTER"
    address router = vm.envAddress(string.concat(chainIdentifier, "_ROUTER"));

    // Scoping to prevent a "stack-too-deep" error.
    {
      bool isSupported = Router(router).isChainSupported(destChainSelector);
      if (!isSupported) revert ChainNotSupported(destChainSelector);

      address[] memory tokenAddresses = new address[](numTokens);
      uint256[] memory tokenAmounts = new uint256[](numTokens);

      // Since solidity does not support array literals being defined in storage or constant, manually define the addresses and amounts of each token that should be sent. They will automatically
      // be converted into a Client.EVMTokenAmount format for the CCIP-Message.
      // Ex: tokenAddresses[0] = address(1);
      // Ex: tokenAmounts[0] = 1e18;
      // [INSERT HERE]

      console.log("Approving Send Tokens...");

      Client.EVMTokenAmount[] memory tokens = new Client.EVMTokenAmount[](numTokens);
      for (uint256 i = 0; i < tokens.length; ++i) {
        if (tokenAddresses[i] != address(0)) {
          // Since the sender may be an EOA with an existing approval, the allowance is checked first.
          uint256 allowance = IERC20(tokens[i].token).allowance(sender, router);

          // If the existing allowance is insufficient, increase it to allow sending through CCIP.
          if (allowance < tokens[i].amount) {
            console.log("Approving %i tokens to Router for %s", tokenAmounts[i], tokenAddresses[i]);
            IERC20(tokens[i].token).safeIncreaseAllowance(router, tokenAmounts[i]);
          }

          // Once approval is granted, copy into the EVM Token Amount Array to be included in the message-proper.
          tokens[i] = Client.EVMTokenAmount({token: tokenAddresses[i], amount: tokenAmounts[i]});
        }
      }
      console.log("--- Tokens Approved ---");

      // Construct the EVM2AnyMessage using the fields defined above.
      message = Client.EVM2AnyMessage({
        receiver: recipient,
        data: data,
        tokenAmounts: tokens,
        feeToken: feeToken,
        extraArgs: extraArgs
      });
    }

    // Calculate the fee in WEI for the message and approve the router if necessary.
    // Note: Even if the token is not native, it will still be provided in WEI.
    uint256 fee = Router(router).getFee(destChainSelector, message);
    console.log("Fee in WEI: %s", fee);
    if (feeToken != address(0)) {
      console.log("Approving Fee Token %s to Router", feeToken);
      IERC20(feeToken).safeIncreaseAllowance(router, fee);
    }

    // Send the message, forwarding native tokens if necessary to pay the fee.
    console.log("Sending message to %i", destChainSelector);
    bytes32 messageId = Router(router).ccipSend{value: feeToken == address(0) ? fee : 0}(destChainSelector, message);

    // Foundry's console library does not support including bytes32 as a parameter so it is printed separately.
    console.log("--- Message sent: MessageId ---");
    console.logBytes32(messageId);

    vm.stopBroadcast();
  }
}
