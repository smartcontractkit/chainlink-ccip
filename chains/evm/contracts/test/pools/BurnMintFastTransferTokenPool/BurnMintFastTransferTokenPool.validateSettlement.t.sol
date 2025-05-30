// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Client} from "../../../libraries/Client.sol";
import {FastTransferTokenPoolAbstract} from "../../../pools/FastTransferTokenPoolAbstract.sol";
import {TokenPool} from "../../../pools/TokenPool.sol";
import {BurnMintFastTransferTokenPoolSetup} from "./BurnMintFastTransferTokenPoolSetup.t.sol";

contract BurnMintFastTransferTokenPool_validateSettlement is BurnMintFastTransferTokenPoolSetup {
  uint256 internal constant TRANSFER_AMOUNT = 100 ether;
  address internal constant RECEIVER = address(0x1234);
  bytes32 internal constant FILL_REQUEST_ID = keccak256("fillRequestId");
  uint8 internal constant SRC_DECIMALS = 18;

  function setUp() public virtual override {
    super.setUp();
    vm.stopPrank();
  }

  function test_ValidateSettlement_Success() public {
    // Create a valid CCIP message that should pass validation
    Client.Any2EVMMessage memory message = _createCcipMessage();

    // This should not revert - validation passes
    vm.prank(address(s_sourceRouter));
    s_pool.ccipReceive(message);

    // Verify the message was processed successfully
    assertEq(s_burnMintERC20.balanceOf(RECEIVER), TRANSFER_AMOUNT);
  }

  function test_ValidateSettlement_RevertWhen_CursedByRMN() public {
    // Mock RMN to return cursed status
    vm.mockCall(address(s_mockRMNRemote), abi.encodeWithSignature("isCursed(bytes16)"), abi.encode(true));

    Client.Any2EVMMessage memory message = _createCcipMessage();

    // Should revert with CursedByRMN error
    vm.expectRevert(TokenPool.CursedByRMN.selector);
    vm.prank(address(s_sourceRouter));
    s_pool.ccipReceive(message);
  }

  function test_ValidateSettlement_RevertWhen_InvalidSourcePoolAddress() public {
    Client.Any2EVMMessage memory message = _createCcipMessage();
    // Set an invalid source pool address
    message.sender = abi.encode(makeAddr("invalidPool"));

    // Should revert with InvalidSourcePoolAddress error
    vm.expectRevert(abi.encodeWithSelector(TokenPool.InvalidSourcePoolAddress.selector, message.sender));
    vm.prank(address(s_sourceRouter));
    s_pool.ccipReceive(message);
  }

  function test_ValidateSettlement_RevertWhen_EmptySourcePoolAddress() public {
    Client.Any2EVMMessage memory message = _createCcipMessage();
    // Set empty source pool address
    message.sender = "";

    // Should revert with InvalidSourcePoolAddress error
    vm.expectRevert(abi.encodeWithSelector(TokenPool.InvalidSourcePoolAddress.selector, message.sender));
    vm.prank(address(s_sourceRouter));
    s_pool.ccipReceive(message);
  }

  function test_ValidateSettlement_RevertWhen_WrongSourceChainSelector() public {
    // Create message with a chain selector that doesn't have remote pools configured
    uint64 wrongChainSelector = 999999;

    Client.Any2EVMMessage memory message = Client.Any2EVMMessage({
      messageId: FILL_REQUEST_ID,
      sourceChainSelector: wrongChainSelector,
      sender: abi.encode(s_remoteBurnMintPool),
      data: abi.encode(
        FastTransferTokenPoolAbstract.MintMessage({
          sourceAmountToTransfer: TRANSFER_AMOUNT,
          sourceDecimals: SRC_DECIMALS,
          fastTransferFee: TRANSFER_AMOUNT * FAST_FEE_BPS / 10_000,
          receiver: abi.encode(RECEIVER)
        })
      ),
      destTokenAmounts: new Client.EVMTokenAmount[](0)
    });

    // Should revert with InvalidSourcePoolAddress error since the chain is not configured
    vm.expectRevert(abi.encodeWithSelector(TokenPool.InvalidSourcePoolAddress.selector, message.sender));
    vm.prank(address(s_sourceRouter));
    s_pool.ccipReceive(message);
  }

  function _createCcipMessage() internal view returns (Client.Any2EVMMessage memory) {
    return Client.Any2EVMMessage({
      messageId: FILL_REQUEST_ID,
      sourceChainSelector: DEST_CHAIN_SELECTOR,
      sender: abi.encode(s_remoteBurnMintPool),
      data: abi.encode(
        FastTransferTokenPoolAbstract.MintMessage({
          sourceAmountToTransfer: TRANSFER_AMOUNT,
          sourceDecimals: SRC_DECIMALS,
          fastTransferFee: TRANSFER_AMOUNT * FAST_FEE_BPS / 10_000,
          receiver: abi.encode(RECEIVER)
        })
      ),
      destTokenAmounts: new Client.EVMTokenAmount[](0)
    });
  }
}
