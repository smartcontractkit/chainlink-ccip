// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IAdvancedPoolHooks} from "../../../interfaces/IAdvancedPoolHooks.sol";
import {IPolicyEngine} from "@chainlink/ace/policy-management/interfaces/IPolicyEngine.sol";
import {Client} from "../../../libraries/Client.sol";
import {MessageV1Codec} from "../../../libraries/MessageV1Codec.sol";
import {OffRamp} from "../../../offRamp/OffRamp.sol";
import {Router} from "../../../Router.sol";
import {AdvancedPoolHooksE2ESetup} from "./AdvancedPoolHooksE2ESetup.t.sol";
import {OffRampHelper} from "../../helpers/OffRampHelper.sol";
import {MockVerifier} from "../../mocks/MockVerifier.sol";

import {VmSafe} from "forge-std/Vm.sol";

contract AdvancedPoolHooksE2E_DestChain is AdvancedPoolHooksE2ESetup {
  OffRampHelper internal s_offRamp;

  function setUp() public virtual override {
    super.setUp();

    // Deploy OffRamp
    s_offRamp = new OffRampHelper(
      OffRamp.StaticConfig({
        localChainSelector: DEST_CHAIN_SELECTOR,
        gasForCallExactCheck: GAS_FOR_CALL_EXACT_CHECK,
        rmnRemote: s_mockRMNRemote,
        tokenAdminRegistry: address(s_tokenAdminRegistry),
        maxGasBufferToUpdateState: DEFAULT_MAX_GAS_BUFFER_TO_UPDATE_STATE
      })
    );

    // Configure OffRamp source chain
    bytes[] memory onRamps = new bytes[](1);
    onRamps[0] = abi.encode(s_onRamp);
    OffRamp.SourceChainConfigArgs[] memory sourceChainConfigs = new OffRamp.SourceChainConfigArgs[](1);
    sourceChainConfigs[0] = OffRamp.SourceChainConfigArgs({
      router: s_destRouter,
      sourceChainSelector: SOURCE_CHAIN_SELECTOR,
      isEnabled: true,
      onRamps: onRamps,
      defaultCCVs: _defaultCCVs(),
      laneMandatedCCVs: new address[](0)
    });
    s_offRamp.applySourceChainConfigUpdates(sourceChainConfigs);

    // Register OffRamp in the destination Router so pool accepts it
    Router.OffRamp[] memory offRampUpdates = new Router.OffRamp[](1);
    offRampUpdates[0] = Router.OffRamp({
      sourceChainSelector: SOURCE_CHAIN_SELECTOR,
      offRamp: address(s_offRamp)
    });
    s_destRouter.applyRampUpdates(
      new Router.OnRamp[](0),
      new Router.OffRamp[](0),
      offRampUpdates
    );

    // Point the pool's router to the dest router so the offRamp is recognized
    s_burnMintPool.setDynamicConfig(address(s_destRouter), address(0), address(0));
  }

  function _defaultCCVs() internal returns (address[] memory) {
    address[] memory ccvs = new address[](1);
    ccvs[0] = address(new MockVerifier(""));
    return ccvs;
  }

  /// @notice Builds a MessageV1Codec.TokenTransferV1 for the dest chain test.
  function _buildTokenTransfer(
    address receiver,
    uint256 amount
  ) internal view returns (MessageV1Codec.TokenTransferV1 memory) {
    return MessageV1Codec.TokenTransferV1({
      amount: amount,
      sourcePoolAddress: abi.encode(address(s_burnMintPool)),
      sourceTokenAddress: abi.encode(address(s_aceToken)),
      destTokenAddress: abi.encodePacked(address(s_aceToken)),
      tokenReceiver: abi.encodePacked(receiver),
      extraData: abi.encode(uint256(18))
    });
  }

  /// @notice Validates the complete dest chain flow: OffRamp -> BurnMintTokenPool -> AdvancedPoolHooks
  /// -> PolicyEngine -> AdvancedPoolHooksExtractor -> VolumePolicy.
  /// Verifies: tokens minted, correct return values, and all 9 extracted parameters.
  function test_destChain_releaseOrMint_policyAccepted() public {
    address receiver = makeAddr("receiver");
    uint256 receiverBalanceBefore = s_aceToken.balanceOf(receiver);

    MessageV1Codec.TokenTransferV1 memory tokenTransfer = _buildTokenTransfer(receiver, VALID_AMOUNT);

    vm.recordLogs();

    (Client.EVMTokenAmount memory destTokenAmount,) = s_offRamp.releaseOrMintSingleToken(
      tokenTransfer,
      abi.encode(OWNER),
      SOURCE_CHAIN_SELECTOR,
      0
    );

    // Verify tokens were minted to receiver
    uint256 receiverBalanceAfter = s_aceToken.balanceOf(receiver);
    assertEq(VALID_AMOUNT, receiverBalanceAfter - receiverBalanceBefore, "Receiver should receive minted tokens");
    assertEq(VALID_AMOUNT, destTokenAmount.amount, "Returned amount should match");
    assertEq(address(s_aceToken), destTokenAmount.token, "Returned token should match");

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
        assertEq(IAdvancedPoolHooks.postflightCheck.selector, bytes4(logs[i].topics[3]));

        // Decode parameters
        (IPolicyEngine.Parameter[] memory params,) = abi.decode(
          logs[i].data,
          (IPolicyEngine.Parameter[], bytes)
        );

        assertEq(9, params.length, "Should have 9 postflight parameters");

        // PARAM_FROM = originalSender
        assertEq(s_extractor.PARAM_FROM(), params[0].name);
        assertEq(abi.encode(OWNER), params[0].value);

        // PARAM_TO = receiver
        assertEq(s_extractor.PARAM_TO(), params[1].name);
        assertEq(receiver, abi.decode(params[1].value, (address)));

        // PARAM_AMOUNT = localAmount
        assertEq(s_extractor.PARAM_AMOUNT(), params[2].name);
        assertEq(VALID_AMOUNT, abi.decode(params[2].value, (uint256)));

        // PARAM_REMOTE_CHAIN_SELECTOR = SOURCE_CHAIN_SELECTOR
        assertEq(s_extractor.PARAM_REMOTE_CHAIN_SELECTOR(), params[3].name);
        assertEq(SOURCE_CHAIN_SELECTOR, abi.decode(params[3].value, (uint64)));

        // PARAM_TOKEN = local token
        assertEq(s_extractor.PARAM_TOKEN(), params[4].name);
        assertEq(address(s_aceToken), abi.decode(params[4].value, (address)));

        // PARAM_BLOCK_CONFIRMATION_REQUESTED = 0
        assertEq(s_extractor.PARAM_BLOCK_CONFIRMATION_REQUESTED(), params[5].name);
        assertEq(uint16(0), abi.decode(params[5].value, (uint16)));

        // PARAM_SOURCE_POOL_ADDRESS
        assertEq(s_extractor.PARAM_SOURCE_POOL_ADDRESS(), params[6].name);
        assertEq(abi.encode(address(s_burnMintPool)), params[6].value);

        // PARAM_SOURCE_POOL_DATA = source decimals
        assertEq(s_extractor.PARAM_SOURCE_POOL_DATA(), params[7].name);
        assertEq(uint256(18), abi.decode(params[7].value, (uint256)));

        // PARAM_SOURCE_DENOMINATED_AMOUNT
        assertEq(s_extractor.PARAM_SOURCE_DENOMINATED_AMOUNT(), params[8].name);
        assertEq(VALID_AMOUNT, abi.decode(params[8].value, (uint256)));

        break;
      }
    }

    assertTrue(foundEvent, "PolicyRunComplete event should be emitted");
  }

  /// @notice Validates that the VolumePolicy rejects amounts above max on the dest chain.
  function test_destChain_releaseOrMint_policyRejected_amountTooHigh() public {
    address receiver = makeAddr("receiver");
    MessageV1Codec.TokenTransferV1 memory tokenTransfer = _buildTokenTransfer(receiver, TOO_HIGH_AMOUNT);

    vm.expectRevert();
    s_offRamp.releaseOrMintSingleToken(
      tokenTransfer,
      abi.encode(OWNER),
      SOURCE_CHAIN_SELECTOR,
      0
    );
  }

  /// @notice Validates that the VolumePolicy rejects amounts below min on the dest chain.
  function test_destChain_releaseOrMint_policyRejected_amountTooLow() public {
    address receiver = makeAddr("receiver");
    MessageV1Codec.TokenTransferV1 memory tokenTransfer = _buildTokenTransfer(receiver, TOO_LOW_AMOUNT);

    vm.expectRevert();
    s_offRamp.releaseOrMintSingleToken(
      tokenTransfer,
      abi.encode(OWNER),
      SOURCE_CHAIN_SELECTOR,
      0
    );
  }
}
