// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IFastTransferPool} from "../../../interfaces/IFastTransferPool.sol";
import {Client} from "../../../libraries/Client.sol";

import {FastTransferTokenPoolAbstract} from "../../../pools/FastTransferTokenPoolAbstract.sol";
import {TokenPool} from "../../../pools/TokenPool.sol";
import {BurnMintFastTransferTokenPoolSetup} from "./BurnMintFastTransferTokenPoolSetup.t.sol";
import {IERC20} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";

contract BurnMintFastTransferTokenPool_ccipReceive is BurnMintFastTransferTokenPoolSetup {
  bytes32 internal constant FILL_REQUEST_ID = keccak256("fillRequestId");

  function setUp() public virtual override {
    super.setUp();
    // Give filler tokens to fill with
    deal(address(s_token), s_filler, 2 * TRANSFER_AMOUNT);
    vm.stopPrank();
    vm.prank(s_filler);
    s_token.approve(address(s_pool), type(uint256).max);
  }

  function test_CcipReceive_NotFastFilled() public {
    uint256 receiverBalanceBefore = s_token.balanceOf(RECEIVER);
    uint256 expectedMintAmount = TRANSFER_AMOUNT;

    Client.Any2EVMMessage memory message = _createCcipMessage();

    vm.expectEmit();
    emit IERC20.Transfer(address(0), RECEIVER, expectedMintAmount);

    vm.expectEmit();
    emit IFastTransferPool.FastTransferSettled(FILL_REQUEST_ID);

    vm.prank(address(s_sourceRouter));
    s_pool.ccipReceive(message);

    assertEq(s_token.balanceOf(RECEIVER), receiverBalanceBefore + expectedMintAmount);
  }

  function test_CcipReceive_FastFilled() public {
    // First, fast fill the request
    uint256 amountToFill = TRANSFER_AMOUNT - (TRANSFER_AMOUNT * FAST_FEE_BPS / 10_000);
    bytes32 fillId = s_pool.computeFillId(FILL_REQUEST_ID, amountToFill, SOURCE_DECIMALS, RECEIVER);

    vm.prank(s_filler);
    s_pool.fastFill(FILL_REQUEST_ID, fillId, DEST_CHAIN_SELECTOR, amountToFill, SOURCE_DECIMALS, RECEIVER);

    uint256 fillerBalanceBefore = s_token.balanceOf(s_filler);
    uint256 receiverBalanceBefore = s_token.balanceOf(RECEIVER);
    uint256 expectedReimbursement = TRANSFER_AMOUNT;

    Client.Any2EVMMessage memory message = _createCcipMessage();

    vm.expectEmit();
    emit IFastTransferPool.FastTransferSettled(FILL_REQUEST_ID);

    vm.prank(address(s_sourceRouter));
    s_pool.ccipReceive(message);

    // Filler should be reimbursed the transfer amount + fee
    assertEq(s_token.balanceOf(s_filler), fillerBalanceBefore + expectedReimbursement);
    // Receiver balance should not change (already received tokens from fast fill)
    assertEq(s_token.balanceOf(RECEIVER), receiverBalanceBefore);
  }

  function test_RevertWhen_AlreadySettled() public {
    // First settlement
    Client.Any2EVMMessage memory message = _createCcipMessage();
    vm.startPrank(address(s_sourceRouter));
    s_pool.ccipReceive(message);

    // Try to settle again
    vm.expectRevert(abi.encodeWithSelector(IFastTransferPool.AlreadySettled.selector, FILL_REQUEST_ID));
    s_pool.ccipReceive(message);
  }

  function test_RevertWhen_CursedByRMN() public {
    vm.mockCall(address(s_mockRMNRemote), abi.encodeWithSignature("isCursed(bytes16)"), abi.encode(true));

    Client.Any2EVMMessage memory message = _createCcipMessage();

    vm.expectRevert(TokenPool.CursedByRMN.selector);
    vm.prank(address(s_sourceRouter));
    s_pool.ccipReceive(message);
  }

  function test_RevertWhen_InvalidSourcePoolAddress() public {
    Client.Any2EVMMessage memory message = _createCcipMessage();
    message.sender = abi.encode(makeAddr("invalidPool")); // Invalid source pool

    vm.expectRevert(abi.encodeWithSelector(TokenPool.InvalidSourcePoolAddress.selector, message.sender));
    vm.prank(address(s_sourceRouter));
    s_pool.ccipReceive(message);
  }

  function test_CcipReceive_WithDifferentDecimals() public {
    uint256 receiverBalanceBefore = s_token.balanceOf(RECEIVER);

    uint8 sourceDecimals = 6; // USDC-like decimals
    uint256 srcAmount = 100e6;
    uint256 expectedLocalAmount = 100 ether; // Should be scaled to 18 decimals

    Client.Any2EVMMessage memory message = _createCcipMessage();
    message.data = abi.encode(
      FastTransferTokenPoolAbstract.MintMessage({
        sourceAmount: srcAmount,
        sourceDecimals: sourceDecimals,
        fastTransferFeeBps: FAST_FEE_BPS,
        receiver: abi.encode(RECEIVER)
      })
    );

    vm.prank(address(s_sourceRouter));
    s_pool.ccipReceive(message);

    assertEq(s_token.balanceOf(RECEIVER), receiverBalanceBefore + expectedLocalAmount);
  }

  function test_CcipReceive_OnlyRouter() public {
    vm.prank(makeAddr("notRouter"));

    Client.Any2EVMMessage memory message = _createCcipMessage();

    vm.expectRevert(); // TODO define specific revert reason
    s_pool.ccipReceive(message);
  }

  function _createCcipMessage() internal view returns (Client.Any2EVMMessage memory) {
    return Client.Any2EVMMessage({
      messageId: FILL_REQUEST_ID,
      sourceChainSelector: DEST_CHAIN_SELECTOR,
      sender: abi.encode(s_remoteBurnMintPool),
      data: abi.encode(
        FastTransferTokenPoolAbstract.MintMessage({
          sourceAmount: TRANSFER_AMOUNT,
          sourceDecimals: SOURCE_DECIMALS,
          fastTransferFeeBps: FAST_FEE_BPS,
          receiver: abi.encode(RECEIVER)
        })
      ),
      destTokenAmounts: new Client.EVMTokenAmount[](0)
    });
  }
}
