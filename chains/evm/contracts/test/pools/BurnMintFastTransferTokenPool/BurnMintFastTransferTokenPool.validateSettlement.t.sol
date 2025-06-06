// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Client} from "../../../libraries/Client.sol";
import {FastTransferTokenPoolAbstract} from "../../../pools/FastTransferTokenPoolAbstract.sol";
import {TokenPool} from "../../../pools/TokenPool.sol";
import {BurnMintFastTransferTokenPoolSetup} from "./BurnMintFastTransferTokenPoolSetup.t.sol";

contract BurnMintFastTransferTokenPool_validateSettlement is BurnMintFastTransferTokenPoolSetup {
  function setUp() public virtual override {
    super.setUp();
    vm.stopPrank();
    vm.startPrank(address(s_sourceRouter));
  }

  function test_validateSettlement_Success() public {
    // Create a valid CCIP message that should pass validation
    Client.Any2EVMMessage memory message = _createCcipMessage();

    // This should not revert - validation passes
    s_pool.ccipReceive(message);

    // Verify the message was processed successfully
    assertEq(s_token.balanceOf(RECEIVER), TRANSFER_AMOUNT);
  }

  function test_validateSettlement_RevertWhen_CursedByRMN() public {
    // Mock RMN to return cursed status
    vm.mockCall(address(s_mockRMNRemote), abi.encodeWithSignature("isCursed(bytes16)"), abi.encode(true));

    Client.Any2EVMMessage memory message = _createCcipMessage();

    // Should revert with CursedByRMN error
    vm.expectRevert(TokenPool.CursedByRMN.selector);
    s_pool.ccipReceive(message);
  }

  function test_validateSettlement_RevertWhen_InvalidSourcePoolAddress() public {
    Client.Any2EVMMessage memory message = _createCcipMessage();
    // Set an invalid source pool address
    message.sender = abi.encode(makeAddr("invalidPool"));

    // Should revert with InvalidSourcePoolAddress error
    vm.expectRevert(abi.encodeWithSelector(TokenPool.InvalidSourcePoolAddress.selector, message.sender));
    s_pool.ccipReceive(message);
  }

  function test_validateSettlement_RevertWhen_EmptySourcePoolAddress() public {
    Client.Any2EVMMessage memory message = _createCcipMessage();
    // Set empty source pool address
    message.sender = "";

    // Should revert with InvalidSourcePoolAddress error
    vm.expectRevert(abi.encodeWithSelector(TokenPool.InvalidSourcePoolAddress.selector, message.sender));
    s_pool.ccipReceive(message);
  }

  function test_validateSettlement_RevertWhen_WrongSourceChainSelector() public {
    // Create message with a chain selector that doesn't have remote pools configured
    uint64 wrongChainSelector = 999999;

    Client.Any2EVMMessage memory message = _createCcipMessage();
    message.sourceChainSelector = wrongChainSelector;

    // Should revert with InvalidSourcePoolAddress error since the chain is not configured
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
          fastTransferFillerFeeBps: FAST_FEE_FILLER_BPS,
          fastTransferPoolFeeBps: 0, // No pool fee for this test
          sourceDecimals: SOURCE_DECIMALS,
          receiver: abi.encode(RECEIVER)
        })
      ),
      destTokenAmounts: new Client.EVMTokenAmount[](0)
    });
  }
}
