// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {FeeQuoter} from "../../FeeQuoter.sol";
import {Client} from "../../libraries/Client.sol";

import {Internal} from "../../libraries/Internal.sol";
import {FeeQuoterSetup} from "./FeeQuoterSetup.t.sol";

contract FeeQuoter_parseSuiExtraArgsFromBytes is FeeQuoterSetup {
  FeeQuoter.DestChainConfig private s_destChainConfig;

  function setUp() public virtual override {
    super.setUp();
    s_destChainConfig = _generateFeeQuoterDestChainConfigArgs()[0].destChainConfig;
    s_destChainConfig.enforceOutOfOrder = true; // Enforcing out of order execution for messages to SUI
    s_destChainConfig.chainFamilySelector = Internal.CHAIN_FAMILY_SELECTOR_SUI;

    FeeQuoter.DestChainConfigArgs[] memory destChainConfigs = new FeeQuoter.DestChainConfigArgs[](1);
    destChainConfigs[0] =
      FeeQuoter.DestChainConfigArgs({destChainSelector: DEST_CHAIN_SELECTOR, destChainConfig: s_destChainConfig});
    s_feeQuoter.applyDestChainConfigUpdates(destChainConfigs);
  }

  function test_SuiExtraArgsV1TagSelector() public pure {
    assertEq(Client.SUI_EXTRA_ARGS_V1_TAG, bytes4(keccak256("CCIP SuiExtraArgsV1")));
  }

  function test_SuiExtraArgsV1() public view {
    Client.SuiExtraArgsV1 memory extraArgs = Client.SuiExtraArgsV1({
      gasLimit: GAS_LIMIT,
      allowOutOfOrderExecution: true,
      tokenReceiver: bytes32(uint256(123)),
      receiverObjectIds: new bytes32[](2)
    });

    bytes memory inputExtraArgs = Client._suiArgsToBytes(extraArgs);

    vm.assertEq(
      abi.encode(s_feeQuoter.parseSuiExtraArgsFromBytes(inputExtraArgs, s_destChainConfig)), abi.encode(extraArgs)
    );
  }

  // Reverts
  function test_RevertWhen_ExtraArgsAreEmpty() public {
    bytes memory inputExtraArgs = new bytes(0);
    vm.expectRevert(FeeQuoter.InvalidExtraArgsData.selector);
    s_feeQuoter.parseSuiExtraArgsFromBytes(inputExtraArgs, s_destChainConfig);
  }

  function test_RevertWhen_InvalidExtraArgsTag() public {
    bytes memory inputExtraArgs = new bytes(4);

    vm.expectRevert(FeeQuoter.InvalidExtraArgsTag.selector);
    s_feeQuoter.parseSuiExtraArgsFromBytes(inputExtraArgs, s_destChainConfig);
  }

  function test_RevertWhen_SuiMessageGasLimitTooHigh() public {
    Client.SuiExtraArgsV1 memory inputArgs = Client.SuiExtraArgsV1({
      gasLimit: s_destChainConfig.maxPerMsgGasLimit + 1,
      allowOutOfOrderExecution: true,
      tokenReceiver: bytes32(uint256(0)),
      receiverObjectIds: new bytes32[](2)
    });

    bytes memory inputExtraArgs = Client._suiArgsToBytes(inputArgs);

    vm.expectRevert(FeeQuoter.MessageGasLimitTooHigh.selector);
    s_feeQuoter.parseSuiExtraArgsFromBytes(inputExtraArgs, s_destChainConfig);
  }

  function test_RevertWhen_ExtraArgOutOfOrderExecutionIsFalse() public {
    bytes memory inputExtraArgs = abi.encodeWithSelector(
      Client.SUI_EXTRA_ARGS_V1_TAG,
      Client.SuiExtraArgsV1({
        gasLimit: GAS_LIMIT,
        allowOutOfOrderExecution: false,
        tokenReceiver: bytes32(uint256(0)),
        receiverObjectIds: new bytes32[](2)
      })
    );

    vm.expectRevert(FeeQuoter.ExtraArgOutOfOrderExecutionMustBeTrue.selector);
    s_feeQuoter.parseSuiExtraArgsFromBytes(inputExtraArgs, s_destChainConfig);
  }
}
