// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {NonceManager} from "../../../NonceManager.sol";
import {Internal} from "../../../libraries/Internal.sol";
import {BaseOnRamp} from "../../../onRamp/BaseOnRamp.sol";
import {CommitOnRamp} from "../../../onRamp/CommitOnRamp.sol";
import {FeeQuoterFeeSetup} from "../../feeQuoter/FeeQuoterSetup.t.sol";

contract CommitOnRampSetup is FeeQuoterFeeSetup {
  address internal constant FEE_AGGREGATOR = 0xa33CDB32eAEce34F6affEfF4899cef45744EDea3;
  address internal constant ALLOWLIST_ADMIN = 0x1234567890123456789012345678901234567890;
  address internal s_ccvProxy;
  CommitOnRamp internal s_commitOnRamp;
  NonceManager internal s_nonceManager;

  function setUp() public virtual override {
    super.setUp();
    s_ccvProxy = makeAddr("CCVProxy");
    s_nonceManager = NonceManager(makeAddr("NonceManager"));

    s_commitOnRamp = new CommitOnRamp(
      address(s_mockRMNRemote),
      address(s_nonceManager),
      CommitOnRamp.DynamicConfig({
        feeQuoter: address(s_feeQuoter),
        feeAggregator: FEE_AGGREGATOR,
        allowlistAdmin: ALLOWLIST_ADMIN
      })
    );

    BaseOnRamp.DestChainConfigArgs[] memory destChainConfigs = new BaseOnRamp.DestChainConfigArgs[](1);
    destChainConfigs[0] = BaseOnRamp.DestChainConfigArgs({
      ccvProxy: s_ccvProxy,
      destChainSelector: DEST_CHAIN_SELECTOR,
      allowlistEnabled: false
    });

    s_commitOnRamp.applyDestChainConfigUpdates(destChainConfigs);
  }

  function _createEVM2AnyVerifierMessage(
    uint64 destChainSelector,
    address sender,
    bytes memory data,
    address receiver,
    address feeToken,
    uint256 feeTokenAmount
  ) internal pure returns (Internal.EVM2AnyVerifierMessage memory) {
    Internal.EVM2AnyVerifierMessage memory message = Internal.EVM2AnyVerifierMessage({
      header: Internal.Header({
        messageId: "",
        sourceChainSelector: SOURCE_CHAIN_SELECTOR,
        destChainSelector: destChainSelector,
        sequenceNumber: 1
      }),
      sender: sender,
      data: data,
      receiver: abi.encode(receiver),
      feeToken: feeToken,
      feeTokenAmount: feeTokenAmount,
      feeValueJuels: 0,
      tokenTransfer: new Internal.EVMTokenTransfer[](0),
      verifierReceipts: new Internal.Receipt[](1),
      executorReceipt: Internal.Receipt({
        issuer: address(0),
        feeTokenAmount: 0,
        destGasLimit: 0,
        destBytesOverhead: 0,
        extraArgs: ""
      })
    });

    message.verifierReceipts[0] =
      Internal.Receipt({issuer: address(0), feeTokenAmount: 0, destGasLimit: 0, destBytesOverhead: 0, extraArgs: ""});

    return message;
  }

  /// @notice Helper function to mock fee quoter response for forwardToVerifier tests
  /// @param isOutOfOrderExecution Whether the message should be processed out of order
  /// @param destChainSelector The destination chain selector
  /// @param feeToken The fee token address
  /// @param feeTokenAmount The fee token amount
  /// @param extraArgs The extra arguments
  /// @param receiver The receiver address
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

  /// @notice Helper function to mock nonce manager response for forwardToVerifier tests
  /// @param destChainSelector The destination chain selector
  /// @param sender The sender address
  /// @param nonce The nonce to return
  function _mockNonceManagerResponse(uint64 destChainSelector, address sender, uint64 nonce) internal {
    vm.mockCall(
      address(s_nonceManager),
      abi.encodeWithSelector(s_nonceManager.getIncrementedOutboundNonce.selector, destChainSelector, sender),
      abi.encode(nonce)
    );
  }

  /// @notice Helper function to set up common mocks for forwardToVerifier tests
  /// @param isOutOfOrderExecution Whether the message should be processed out of order
  /// @param destChainSelector The destination chain selector
  /// @param feeToken The fee token address
  /// @param feeTokenAmount The fee token amount
  /// @param receiver The receiver address
  /// @param nonce The nonce to return (only used for ordered execution)
  function _setupForwardToVerifierMocks(
    bool isOutOfOrderExecution,
    uint64 destChainSelector,
    address feeToken,
    uint256 feeTokenAmount,
    address receiver,
    uint64 nonce
  ) internal {
    _mockFeeQuoterResponse(isOutOfOrderExecution, destChainSelector, feeToken, feeTokenAmount, "", receiver);

    if (!isOutOfOrderExecution) {
      _mockNonceManagerResponse(destChainSelector, receiver, nonce);
    }
  }
}
