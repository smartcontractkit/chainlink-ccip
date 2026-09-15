// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IRouterClient} from "../../interfaces/IRouterClient.sol";

import {Client} from "../../libraries/Client.sol";
import {ExtraArgsCodec} from "../../libraries/ExtraArgsCodec.sol";

import {Script, console2} from "forge-std/Script.sol";

/// @notice Sends one data message from Sepolia to Arbitrum Sepolia verified only by the ZK CCV, with manual execution.
/// @dev Reads the source resolver from ZK_RESOLVER and the destination receiver from RECEIVER.
contract Send is Script {
  address internal constant ROUTER = 0x784d49a71BB4C48eB7dA4cD7e6Ecb424f9b5EAB1;
  uint64 internal constant ARBITRUM_SEPOLIA_SELECTOR = 3478487238524512106;

  function run() external {
    address zkResolver = vm.envAddress("ZK_RESOLVER");
    address receiver = vm.envAddress("RECEIVER");

    address[] memory ccvs = new address[](1);
    ccvs[0] = zkResolver;
    ExtraArgsCodec.GenericExtraArgsV3 memory extraArgs = ExtraArgsCodec.GenericExtraArgsV3({
      gasLimit: 50_000,
      requestedFinalityConfig: bytes4(0),
      ccvs: ccvs,
      ccvArgs: new bytes[](1),
      executor: Client.NO_EXECUTION_ADDRESS,
      executorArgs: "",
      tokenReceiver: "",
      tokenArgs: ""
    });

    Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
      receiver: abi.encode(receiver),
      data: "zk ccv poc",
      tokenAmounts: new Client.EVMTokenAmount[](0),
      feeToken: address(0),
      extraArgs: ExtraArgsCodec._encodeGenericExtraArgsV3(extraArgs)
    });

    uint256 fee = IRouterClient(ROUTER).getFee(ARBITRUM_SEPOLIA_SELECTOR, message);
    console2.log("fee wei", fee);

    vm.startBroadcast();
    bytes32 messageId = IRouterClient(ROUTER).ccipSend{value: fee}(ARBITRUM_SEPOLIA_SELECTOR, message);
    vm.stopBroadcast();

    console2.log("messageId");
    console2.logBytes32(messageId);
  }
}
