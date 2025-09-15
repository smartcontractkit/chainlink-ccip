// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {MessageV1Codec} from "../../../libraries/MessageV1Codec.sol";
import {BaseOnRamp} from "../../../onRamp/BaseOnRamp.sol";
import {CommitOnRamp} from "../../../onRamp/CommitOnRamp.sol";
import {BaseOnRampSetup} from "../BaseOnRamp/BaseOnRampSetup.t.sol";

contract CommitOnRampSetup is BaseOnRampSetup {
  CommitOnRamp internal s_commitOnRamp;

  function setUp() public virtual override {
    super.setUp();

    s_commitOnRamp = new CommitOnRamp(_createBasicDynamicConfigArgs());

    BaseOnRamp.DestChainConfigArgs[] memory destChainConfigs = new BaseOnRamp.DestChainConfigArgs[](1);
    destChainConfigs[0] = BaseOnRamp.DestChainConfigArgs({
      router: s_router,
      destChainSelector: DEST_CHAIN_SELECTOR,
      allowlistEnabled: false
    });

    s_commitOnRamp.applyDestChainConfigUpdates(destChainConfigs);
  }

  /// @notice Helper to create a minimal dynamic config.
  function _createBasicDynamicConfigArgs() internal view returns (CommitOnRamp.DynamicConfig memory) {
    return CommitOnRamp.DynamicConfig({
      feeQuoter: address(s_feeQuoter),
      feeAggregator: FEE_AGGREGATOR,
      allowlistAdmin: ALLOWLIST_ADMIN
    });
  }

  /// @notice Helper to create a dynamic config with custom addresses.
  function _createDynamicConfigArgs(
    address feeQuoter,
    address feeAggregator,
    address allowlistAdmin
  ) internal pure returns (CommitOnRamp.DynamicConfig memory) {
    return
      CommitOnRamp.DynamicConfig({feeQuoter: feeQuoter, feeAggregator: feeAggregator, allowlistAdmin: allowlistAdmin});
  }

  function _createMessageV1(
    uint64 destChainSelector,
    address sender,
    bytes memory data,
    address receiver
  ) internal view returns (MessageV1Codec.MessageV1 memory, bytes32 messageId) {
    MessageV1Codec.MessageV1 memory message = MessageV1Codec.MessageV1({
      sourceChainSelector: SOURCE_CHAIN_SELECTOR,
      destChainSelector: destChainSelector,
      sequenceNumber: 1,
      onRampAddress: abi.encodePacked(address(s_ccvProxy)),
      offRampAddress: abi.encodePacked(address(s_ccvAggregatorRemote)),
      finality: 0,
      sender: abi.encodePacked(sender),
      receiver: abi.encodePacked(receiver),
      destBlob: "",
      tokenTransfer: new MessageV1Codec.TokenTransferV1[](0),
      data: data
    });

    return (message, keccak256(MessageV1Codec._encodeMessageV1(message)));
  }

  /// @notice Helper function to mock fee quoter response for forwardToVerifier tests.
  /// @param isOutOfOrderExecution Whether the message should be processed out of order.
  /// @param destChainSelector The destination chain selector.
  /// @param feeToken The fee token address.
  /// @param feeTokenAmount The fee token amount.
  /// @param extraArgs The extra arguments.
  /// @param receiver The receiver address.
  function _mockFeeQuoterResponse(
    bool isOutOfOrderExecution,
    uint64 destChainSelector,
    address feeToken,
    uint256 feeTokenAmount,
    bytes memory extraArgs,
    address receiver
  ) internal {
    vm.mockCall(
      address(s_feeQuoter),
      abi.encodeWithSelector(
        s_feeQuoter.processMessageArgs.selector,
        destChainSelector,
        feeToken,
        feeTokenAmount,
        extraArgs,
        abi.encode(receiver)
      ),
      abi.encode(0, isOutOfOrderExecution, "", "")
    );
  }
}
