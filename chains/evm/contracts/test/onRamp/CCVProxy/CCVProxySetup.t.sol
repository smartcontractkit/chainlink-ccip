// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Client} from "../../../libraries/Client.sol";
import {Internal} from "../../../libraries/Internal.sol";
import {CCVProxy} from "../../../onRamp/CCVProxy.sol";
import {FeeQuoterFeeSetup} from "../../feeQuoter/FeeQuoterSetup.t.sol";
import {MockCCVOnRamp} from "../../mocks/MockCCVOnRamp.sol";
import {MockExecutor} from "../../mocks/MockExecutor.sol";

contract CCVProxySetup is FeeQuoterFeeSetup {
  address internal constant FEE_AGGREGATOR = 0xa33CDB32eAEce34F6affEfF4899cef45744EDea3;
  bytes32 internal s_metadataHash;

  CCVProxy internal s_ccvProxy;
  MockCCVOnRamp internal s_mockCCVOne;
  MockExecutor internal s_mockExecutor;

  function setUp() public virtual override {
    super.setUp();

    s_ccvProxy = new CCVProxy(
      CCVProxy.StaticConfig({
        chainSelector: SOURCE_CHAIN_SELECTOR,
        rmnRemote: s_mockRMNRemote,
        tokenAdminRegistry: address(s_tokenAdminRegistry)
      }),
      CCVProxy.DynamicConfig({
        feeQuoter: address(s_feeQuoter),
        reentrancyGuardEntered: false,
        feeAggregator: FEE_AGGREGATOR
      })
    );
    s_mockCCVOne = new MockCCVOnRamp();
    s_mockExecutor = new MockExecutor();

    CCVProxy.DestChainConfigArgs[] memory destChainConfigs = new CCVProxy.DestChainConfigArgs[](1);
    destChainConfigs[0] = CCVProxy.DestChainConfigArgs({
      destChainSelector: DEST_CHAIN_SELECTOR,
      router: s_sourceRouter,
      defaultCCV: address(s_mockCCVOne),
      requiredCCV: address(0),
      defaultExecutor: address(s_mockExecutor)
    });

    s_ccvProxy.applyDestChainConfigUpdates(destChainConfigs);
  }

  function _evmMessageToEvent(
    Client.EVM2AnyMessage memory message,
    uint64 destChainSelector,
    uint64 seqNum,
    uint256 feeTokenAmount,
    uint256 feeValueJuels,
    address originalSender,
    bytes32 metadataHash
  ) internal pure returns (Internal.EVM2AnyVerifierMessage memory) {
    // TODO
    //    if (message.tokenAmounts.length > 0) {
    //      tokenTransfer =
    //        Internal.EVM2AnyTokenTransfer({token: message.tokenAmounts[0].token, amount: message.tokenAmounts[0].amount});
    //    }

    Internal.EVM2AnyVerifierMessage memory messageEvent = Internal.EVM2AnyVerifierMessage({
      header: Internal.Header({
        messageId: "",
        sourceChainSelector: SOURCE_CHAIN_SELECTOR,
        destChainSelector: destChainSelector,
        sequenceNumber: seqNum
      }),
      sender: originalSender,
      data: message.data,
      receiver: message.receiver,
      feeToken: message.feeToken,
      feeTokenAmount: feeTokenAmount,
      feeValueJuels: feeValueJuels,
      tokenTransfer: new Internal.EVMTokenTransfer[](message.tokenAmounts.length),
      verifierReceipts: new Internal.Receipt[](0),
      executorReceipt: Internal.Receipt({
        issuer: address(0),
        feeTokenAmount: 0,
        destGasLimit: 0,
        destBytesOverhead: 0,
        extraArgs: ""
      })
    });

    messageEvent.header.messageId = Internal._hash(messageEvent, metadataHash);
    return messageEvent;
  }
}
