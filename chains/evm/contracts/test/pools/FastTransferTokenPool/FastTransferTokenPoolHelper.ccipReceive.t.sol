// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.10;

import {IFastTransferPool} from "../../../interfaces/IFastTransferPool.sol";
import {Client} from "../../../libraries/Client.sol";

import {FastTransferTokenPoolAbstract} from "../../../pools/FastTransferTokenPoolAbstract.sol";

import {FastTransferTokenPoolHelperSetup} from "./FastTransferTokenPoolHelperSetup.t.sol";

contract FastTransferTokenPoolHelper_ccipReceive_Test is FastTransferTokenPoolHelperSetup {
  bytes32 public messageId;
  address public sourcePool;
  uint64 public sourceChainSelector;
  uint256 public srcAmount;
  uint8 public srcDecimals;
  uint256 public fastTransferFee;
  address public receiver;

  function setUp() public override {
    super.setUp();
    vm.stopPrank();
    messageId = bytes32("messageId");
    sourcePool = address(0x123);
    sourceChainSelector = 1;
    srcAmount = 100 ether;
    srcDecimals = 18;
    fastTransferFee = srcAmount * FAST_FEE_BPS / 10000; // 1% fast fee
    receiver = address(0x5);
  }

  function test_CcipReceive() public {
    // Prepare CCIP message
    bytes memory data = abi.encode(
      FastTransferTokenPoolAbstract.MintMessage({
        receiver: abi.encode(receiver),
        srcAmountToTransfer: srcAmount,
        srcDecimals: srcDecimals,
        fastTransferFee: fastTransferFee
      })
    );
    Client.Any2EVMMessage memory message = Client.Any2EVMMessage({
      messageId: messageId,
      sourceChainSelector: sourceChainSelector,
      sender: abi.encode(sourcePool),
      data: data,
      destTokenAmounts: new Client.EVMTokenAmount[](0)
    });

    // Mock router call
    vm.expectEmit(true, false, false, true);
    emit IFastTransferPool.FastFillSettled(messageId);
    vm.prank(address(s_sourceRouter));
    s_tokenPool.ccipReceive(message);
  }

  function test_CcipReceive_RevertWhen_InvalidData() public {
    // Prepare CCIP message with invalid data
    bytes memory invalidData = abi.encode(srcAmount, srcDecimals); // Missing fastTransferFee and receiver
    Client.Any2EVMMessage memory message = Client.Any2EVMMessage({
      messageId: messageId,
      sourceChainSelector: sourceChainSelector,
      sender: abi.encode(sourcePool),
      data: invalidData,
      destTokenAmounts: new Client.EVMTokenAmount[](0)
    });

    // Mock router call
    vm.prank(address(s_sourceRouter));
    vm.expectRevert(); // Should revert due to invalid data format
    s_tokenPool.ccipReceive(message);
  }
}
