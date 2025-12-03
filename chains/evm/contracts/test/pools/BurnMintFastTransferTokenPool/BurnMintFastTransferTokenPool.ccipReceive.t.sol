// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IFastTransferPool} from "../../../interfaces/IFastTransferPool.sol";

import {Client} from "../../../libraries/Client.sol";
import {FastTransferTokenPoolAbstract} from "../../../pools/FastTransferTokenPoolAbstract.sol";
import {TokenPool} from "../../../pools/TokenPool.sol";
import {BurnMintFastTransferTokenPoolSetup} from "./BurnMintFastTransferTokenPoolSetup.t.sol";

import {IERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/IERC20.sol";

contract BurnMintFastTransferTokenPool_ccipReceive is BurnMintFastTransferTokenPoolSetup {
  function test_ccipReceive_SlowFilled() public {
    vm.stopPrank();
    vm.startPrank(address(s_sourceRouter));

    uint256 receiverBalanceBefore = s_token.balanceOf(RECEIVER);
    uint256 expectedMintAmount = TRANSFER_AMOUNT;

    Client.Any2EVMMessage memory message = _createCcipMessage();

    uint256 fastTransferFee = (TRANSFER_AMOUNT * FAST_FEE_FILLER_BPS) / 10_000;
    bytes32 fillId = s_pool.computeFillId(
      message.messageId,
      message.sourceChainSelector,
      TRANSFER_AMOUNT - fastTransferFee,
      SOURCE_DECIMALS,
      abi.encode(RECEIVER)
    );

    // Expect inbound rate limit consumption (mint amount consumed)
    vm.expectEmit();
    emit TokenPool.InboundRateLimitConsumed(DEST_CHAIN_SELECTOR, address(s_token), TRANSFER_AMOUNT);

    // A transfer from address(0) meant the tokens were minted
    vm.expectEmit();
    emit IERC20.Transfer(address(0), RECEIVER, expectedMintAmount);

    vm.expectEmit();
    emit IFastTransferPool.FastTransferSettled(fillId, message.messageId, 0, 0, IFastTransferPool.FillState.NOT_FILLED);

    s_pool.ccipReceive(message);

    assertEq(s_token.balanceOf(RECEIVER), receiverBalanceBefore + expectedMintAmount);
  }

  function test_ccipReceive_RevertWhen_InvalidSourcePoolAddress() public {
    vm.stopPrank();
    vm.startPrank(address(s_sourceRouter));

    Client.Any2EVMMessage memory message = _createCcipMessage();
    message.sender = abi.encode(makeAddr("invalidPool"));

    vm.expectRevert(abi.encodeWithSelector(TokenPool.InvalidSourcePoolAddress.selector, message.sender));
    s_pool.ccipReceive(message);
  }

  function _createCcipMessage() internal view returns (Client.Any2EVMMessage memory) {
    return Client.Any2EVMMessage({
      messageId: keccak256("settlementId"),
      sourceChainSelector: DEST_CHAIN_SELECTOR,
      sender: abi.encode(s_remoteBurnMintPool),
      data: abi.encode(
        FastTransferTokenPoolAbstract.MintMessage({
          sourceAmount: TRANSFER_AMOUNT,
          sourceDecimals: SOURCE_DECIMALS,
          fastTransferFillerFeeBps: FAST_FEE_FILLER_BPS,
          fastTransferPoolFeeBps: 0,
          receiver: abi.encode(RECEIVER)
        })
      ),
      destTokenAmounts: new Client.EVMTokenAmount[](0)
    });
  }
}
