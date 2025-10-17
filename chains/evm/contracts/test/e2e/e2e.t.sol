// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Router} from "../../Router.sol";
import {CommitteeVerifier} from "../../ccvs/CommitteeVerifier.sol";
import {VerifierProxy} from "../../ccvs/VerifierProxy.sol";
import {BaseVerifier} from "../../ccvs/components/BaseVerifier.sol";
import {Client} from "../../libraries/Client.sol";
import {Internal} from "../../libraries/Internal.sol";
import {OffRamp} from "../../offRamp/OffRamp.sol";
import {OnRamp} from "../../onRamp/OnRamp.sol";
import {OffRampHelper} from "../helpers/OffRampHelper.sol";
import {MockVerifier} from "../mocks/MockVerifier.sol";
import {OnRampSetup} from "../onRamp/OnRamp/OnRampSetup.t.sol";

import {IERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/IERC20.sol";

contract e2e is OnRampSetup {
  OffRampHelper internal s_offRamp;

  address internal s_destVerifier;
  address internal s_userSpecifiedCCV;

  function setUp() public virtual override {
    super.setUp();

    Router.OnRamp[] memory onRampUpdates = new Router.OnRamp[](1);
    onRampUpdates[0] = Router.OnRamp({destChainSelector: DEST_CHAIN_SELECTOR, onRamp: address(s_onRamp)});
    s_sourceRouter.applyRampUpdates(onRampUpdates, new Router.OffRamp[](0), new Router.OffRamp[](0));

    CommitteeVerifier committeeVerifier = new CommitteeVerifier(
      CommitteeVerifier.DynamicConfig({
        feeQuoter: address(s_feeQuoter),
        feeAggregator: address(1),
        allowlistAdmin: address(0)
      }),
      ""
    );

    BaseVerifier.DestChainConfigArgs[] memory destChainConfigs = new BaseVerifier.DestChainConfigArgs[](1);
    destChainConfigs[0] = BaseVerifier.DestChainConfigArgs({
      router: s_sourceRouter,
      destChainSelector: DEST_CHAIN_SELECTOR,
      allowlistEnabled: false,
      feeUSDCents: DEFAULT_CCV_FEE_USD_CENTS,
      gasForVerification: DEFAULT_CCV_GAS_LIMIT,
      payloadSizeBytes: DEFAULT_CCV_PAYLOAD_SIZE
    });
    committeeVerifier.applyDestChainConfigUpdates(destChainConfigs);

    s_userSpecifiedCCV = address(new VerifierProxy(address(committeeVerifier)));

    // OffRamp side
    s_offRamp = new OffRampHelper(
      OffRamp.StaticConfig({
        localChainSelector: DEST_CHAIN_SELECTOR,
        gasForCallExactCheck: GAS_FOR_CALL_EXACT_CHECK,
        rmnRemote: s_mockRMNRemote,
        tokenAdminRegistry: address(s_tokenAdminRegistry)
      })
    );

    s_destVerifier = address(new MockVerifier(""));

    address[] memory defaultCCVs = new address[](1);
    defaultCCVs[0] = s_destVerifier;

    OffRamp.SourceChainConfigArgs[] memory updates = new OffRamp.SourceChainConfigArgs[](1);
    updates[0] = OffRamp.SourceChainConfigArgs({
      router: s_destRouter,
      sourceChainSelector: SOURCE_CHAIN_SELECTOR,
      isEnabled: true,
      onRamp: abi.encodePacked(s_onRamp),
      defaultCCV: defaultCCVs,
      laneMandatedCCVs: new address[](0)
    });
    s_offRamp.applySourceChainConfigUpdates(updates);

    Router.OffRamp[] memory offRampUpdates = new Router.OffRamp[](1);
    offRampUpdates[0] = Router.OffRamp({sourceChainSelector: SOURCE_CHAIN_SELECTOR, offRamp: address(s_offRamp)});
    s_destRouter.applyRampUpdates(new Router.OnRamp[](0), new Router.OffRamp[](0), offRampUpdates);
  }

  function test_e2e() public {
    uint64 expectedSeqNum = s_onRamp.getDestChainConfig(DEST_CHAIN_SELECTOR).sequenceNumber + 1;

    IERC20(s_sourceFeeToken).approve(address(s_sourceRouter), type(uint256).max);

    Client.CCV[] memory userCCVs = new Client.CCV[](1);
    userCCVs[0] = Client.CCV({ccvAddress: s_userSpecifiedCCV, args: "1"});

    Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
      receiver: abi.encode(OWNER),
      data: "e2e test data",
      tokenAmounts: new Client.EVMTokenAmount[](1),
      feeToken: s_sourceFeeToken,
      extraArgs: Client._argsToBytes(
        Client.EVMExtraArgsV3({ccvs: userCCVs, finalityConfig: 0, executor: address(0), executorArgs: "", tokenArgs: ""})
      )
    });
    message.tokenAmounts[0] = Client.EVMTokenAmount({token: s_sourceFeeToken, amount: 1e18});

    (bytes32 messageId, bytes memory encodedMessage, OnRamp.Receipt[] memory receipts, bytes[] memory verifierBlobs) =
    _evmMessageToEvent({
      message: message,
      destChainSelector: DEST_CHAIN_SELECTOR,
      seqNum: expectedSeqNum,
      originalSender: OWNER
    });

    vm.expectEmit();
    emit OnRamp.CCIPMessageSent({
      destChainSelector: DEST_CHAIN_SELECTOR,
      sequenceNumber: expectedSeqNum,
      messageId: messageId,
      encodedMessage: encodedMessage,
      receipts: receipts,
      verifierBlobs: verifierBlobs
    });

    s_sourceRouter.ccipSend(DEST_CHAIN_SELECTOR, message);

    assertEq(s_onRamp.getDestChainConfig(DEST_CHAIN_SELECTOR).sequenceNumber, expectedSeqNum);

    address[] memory ccvAddresses = new address[](1);
    ccvAddresses[0] = s_destVerifier;

    vm.expectEmit();
    emit OffRamp.ExecutionStateChanged({
      sourceChainSelector: SOURCE_CHAIN_SELECTOR,
      sequenceNumber: expectedSeqNum,
      messageId: messageId,
      state: Internal.MessageExecutionState.SUCCESS,
      returnData: ""
    });

    s_offRamp.execute(encodedMessage, ccvAddresses, new bytes[](1));
  }
}
