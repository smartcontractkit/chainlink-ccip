// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IAdvancedPoolHooks} from "../../../interfaces/IAdvancedPoolHooks.sol";
import {IPolicyEngine} from "@chainlink/ace/policy-management/interfaces/IPolicyEngine.sol";
import {Client} from "../../../libraries/Client.sol";
import {Router} from "../../../Router.sol";
import {AdvancedPoolHooksE2ESetup} from "./AdvancedPoolHooksE2ESetup.t.sol";

import {IERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/IERC20.sol";
import {VmSafe} from "forge-std/Vm.sol";

contract AdvancedPoolHooksE2E_SourceChain is AdvancedPoolHooksE2ESetup {
  function setUp() public virtual override {
    super.setUp();

    // Register Router -> OnRamp mapping for source chain
    Router.OnRamp[] memory onRampUpdates = new Router.OnRamp[](1);
    onRampUpdates[0] = Router.OnRamp({
      destChainSelector: DEST_CHAIN_SELECTOR,
      onRamp: address(s_onRamp)
    });
    s_sourceRouter.applyRampUpdates(
      onRampUpdates,
      new Router.OffRamp[](0),
      new Router.OffRamp[](0)
    );

    // Mint tokens to the test contract
    deal(address(s_aceToken), OWNER, type(uint256).max);

    // Mint fee tokens to the test contract
    deal(s_sourceFeeToken, OWNER, type(uint256).max);
  }

  /// @notice Validates the complete source chain flow: Router -> OnRamp -> BurnMintTokenPool -> AdvancedPoolHooks
  /// -> PolicyEngine -> AdvancedPoolHooksExtractor -> VolumePolicy.
  /// Verifies: transfer succeeds, tokens burned, and all 7 extracted parameters are correct.
  function test_sourceChain_ccipSend_policyAccepted() public {
    address receiver = makeAddr("receiver");
    Client.EVM2AnyMessage memory message = _buildCCIPMessage(receiver, VALID_AMOUNT);

    uint256 fee = s_sourceRouter.getFee(DEST_CHAIN_SELECTOR, message);
    IERC20(s_sourceFeeToken).approve(address(s_sourceRouter), fee);
    IERC20(address(s_aceToken)).approve(address(s_sourceRouter), VALID_AMOUNT);

    uint256 senderBalanceBefore = s_aceToken.balanceOf(OWNER);

    vm.recordLogs();

    s_sourceRouter.ccipSend(DEST_CHAIN_SELECTOR, message);

    // Verify tokens were burned from sender
    uint256 senderBalanceAfter = s_aceToken.balanceOf(OWNER);
    assertEq(VALID_AMOUNT, senderBalanceBefore - senderBalanceAfter, "Sender balance should decrease by transfer amount");

    // Find and decode the PolicyRunComplete event to verify the full ACE pipeline
    VmSafe.Log[] memory logs = vm.getRecordedLogs();
    bytes32 policyRunCompleteTopic = keccak256(
      "PolicyRunComplete(address,address,bytes4,(bytes32,bytes)[],bytes)"
    );

    bool foundEvent = false;
    for (uint256 i = 0; i < logs.length; i++) {
      if (logs[i].emitter == address(s_policyEngine)
          && logs[i].topics.length > 0
          && logs[i].topics[0] == policyRunCompleteTopic) {
        foundEvent = true;

        // Verify indexed fields
        assertEq(address(s_burnMintPool), address(uint160(uint256(logs[i].topics[1]))));
        assertEq(address(s_advancedPoolHooks), address(uint160(uint256(logs[i].topics[2]))));
        assertEq(IAdvancedPoolHooks.preflightCheck.selector, bytes4(logs[i].topics[3]));

        // Decode non-indexed data: extractedParameters and context
        (IPolicyEngine.Parameter[] memory params,) = abi.decode(
          logs[i].data,
          (IPolicyEngine.Parameter[], bytes)
        );

        assertEq(7, params.length, "Should have 7 preflight parameters");

        // PARAM_FROM = originalSender
        assertEq(s_extractor.PARAM_FROM(), params[0].name);
        assertEq(OWNER, abi.decode(params[0].value, (address)));

        // PARAM_TO = receiver
        assertEq(s_extractor.PARAM_TO(), params[1].name);
        assertEq(receiver, abi.decode(params[1].value, (address)));

        // PARAM_AMOUNT = lockOrBurnIn.amount
        assertEq(s_extractor.PARAM_AMOUNT(), params[2].name);
        assertEq(VALID_AMOUNT, abi.decode(params[2].value, (uint256)));

        // PARAM_AMOUNT_POST_FEE = amount after pool bps fee
        assertEq(s_extractor.PARAM_AMOUNT_POST_FEE(), params[3].name);
        uint256 amountPostFee = abi.decode(params[3].value, (uint256));
        assertLe(amountPostFee, VALID_AMOUNT);

        // PARAM_REMOTE_CHAIN_SELECTOR
        assertEq(s_extractor.PARAM_REMOTE_CHAIN_SELECTOR(), params[4].name);
        assertEq(DEST_CHAIN_SELECTOR, abi.decode(params[4].value, (uint64)));

        // PARAM_TOKEN = local token address
        assertEq(s_extractor.PARAM_TOKEN(), params[5].name);
        assertEq(address(s_aceToken), abi.decode(params[5].value, (address)));

        // PARAM_BLOCK_CONFIRMATION_REQUESTED
        assertEq(s_extractor.PARAM_BLOCK_CONFIRMATION_REQUESTED(), params[6].name);

        break;
      }
    }

    assertTrue(foundEvent, "PolicyRunComplete event should be emitted");
  }

  /// @notice Validates that a transfer amount exceeding VolumePolicy max is rejected.
  function test_sourceChain_ccipSend_policyRejected_amountTooHigh() public {
    address receiver = makeAddr("receiver");
    Client.EVM2AnyMessage memory message = _buildCCIPMessage(receiver, TOO_HIGH_AMOUNT);

    uint256 fee = s_sourceRouter.getFee(DEST_CHAIN_SELECTOR, message);
    IERC20(s_sourceFeeToken).approve(address(s_sourceRouter), fee);
    IERC20(address(s_aceToken)).approve(address(s_sourceRouter), TOO_HIGH_AMOUNT);

    vm.expectRevert();
    s_sourceRouter.ccipSend(DEST_CHAIN_SELECTOR, message);
  }

  /// @notice Validates that a transfer amount below VolumePolicy min is rejected.
  function test_sourceChain_ccipSend_policyRejected_amountTooLow() public {
    address receiver = makeAddr("receiver");
    Client.EVM2AnyMessage memory message = _buildCCIPMessage(receiver, TOO_LOW_AMOUNT);

    uint256 fee = s_sourceRouter.getFee(DEST_CHAIN_SELECTOR, message);
    IERC20(s_sourceFeeToken).approve(address(s_sourceRouter), fee);
    IERC20(address(s_aceToken)).approve(address(s_sourceRouter), TOO_LOW_AMOUNT);

    vm.expectRevert();
    s_sourceRouter.ccipSend(DEST_CHAIN_SELECTOR, message);
  }
}
