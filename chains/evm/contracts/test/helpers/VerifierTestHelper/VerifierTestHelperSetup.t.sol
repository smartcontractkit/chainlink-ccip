// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IRouter} from "../../../interfaces/IRouter.sol";

import {BaseVerifier} from "../../../ccvs/components/BaseVerifier.sol";
import {MessageV1Codec} from "../../../libraries/MessageV1Codec.sol";
import {BaseTest} from "../../BaseTest.t.sol";
import {IOffRamp, VerifierTestHelper} from "../VerifierTestHelper.sol";

contract VerifierTestHelperSetup is BaseTest {
  VerifierTestHelper internal s_verifier;

  address internal s_testRouter;
  address internal s_testToken;
  address internal s_onRamp;
  address internal s_offRamp;
  address internal s_sender;
  address internal s_receiver;

  string[] internal s_storageLocations;

  bytes4 internal constant VERSION_TAG = bytes4(keccak256("VerifierTestHelper test"));

  function setUp() public virtual override {
    super.setUp();

    s_testRouter = makeAddr("testRouter");
    s_testToken = makeAddr("testToken");
    s_onRamp = makeAddr("onRamp");
    s_offRamp = makeAddr("offRamp");
    s_sender = makeAddr("sender");
    s_receiver = makeAddr("receiver");

    vm.etch(s_testRouter, bytes("fake bytecode"));
    vm.etch(s_offRamp, bytes("fake bytecode"));
    vm.mockCall(s_testRouter, abi.encodeCall(IRouter.getOnRamp, (DEST_CHAIN_SELECTOR)), abi.encode(s_onRamp));
    vm.mockCall(s_testRouter, abi.encodeCall(IRouter.isOffRamp, (SOURCE_CHAIN_SELECTOR, s_offRamp)), abi.encode(true));

    s_storageLocations.push("testStorageLocation");
    s_verifier =
      new VerifierTestHelper(s_testRouter, s_testToken, s_storageLocations, address(s_mockRMNRemote), VERSION_TAG);

    vm.stopPrank();
  }

  function _configureRemoteChain(
    uint64 remoteChainSelector
  ) internal {
    BaseVerifier.RemoteChainConfigArgs[] memory configs = new BaseVerifier.RemoteChainConfigArgs[](1);
    configs[0] = BaseVerifier.RemoteChainConfigArgs({
      router: IRouter(s_verifier.getTestRouter()),
      remoteChainSelector: remoteChainSelector,
      allowlistEnabled: true,
      feeUSDCents: 1,
      gasForVerification: 1,
      payloadSizeBytes: 1
    });

    vm.prank(OWNER);
    s_verifier.applyRemoteChainConfigUpdates(configs);
  }

  function _allowSender(
    uint64 destChainSelector,
    address sender
  ) internal {
    address[] memory senders = new address[](1);
    senders[0] = sender;
    BaseVerifier.AllowlistConfigArgs[] memory configs = new BaseVerifier.AllowlistConfigArgs[](1);
    configs[0] = BaseVerifier.AllowlistConfigArgs({
      destChainSelector: destChainSelector,
      allowlistEnabled: true,
      addedAllowlistedSenders: senders,
      removedAllowlistedSenders: new address[](0)
    });

    vm.prank(OWNER);
    s_verifier.applyAllowlistUpdates(configs);
  }

  function _mockOffRampSourceChainConfig(
    uint64 sourceChainSelector,
    address router
  ) internal {
    IOffRamp.SourceChainConfig memory config = IOffRamp.SourceChainConfig({
      router: IRouter(router),
      isEnabled: true,
      onRamps: new bytes[](0),
      defaultCCVs: new address[](0),
      laneMandatedCCVs: new address[](0)
    });
    vm.mockCall(s_offRamp, abi.encodeCall(IOffRamp.getSourceChainConfig, (sourceChainSelector)), abi.encode(config));
  }

  function _createMessage(
    uint64 sourceChainSelector,
    uint64 destChainSelector,
    address sourceToken,
    address destToken,
    address sender,
    address receiver
  ) internal pure returns (MessageV1Codec.MessageV1 memory message) {
    MessageV1Codec.TokenTransferV1[] memory tokenTransfers = new MessageV1Codec.TokenTransferV1[](1);
    tokenTransfers[0] = MessageV1Codec.TokenTransferV1({
      amount: 1,
      sourcePoolAddress: abi.encode(address(1)),
      sourceTokenAddress: abi.encode(sourceToken),
      destTokenAddress: abi.encodePacked(destToken),
      tokenReceiver: abi.encodePacked(receiver),
      extraData: ""
    });

    message = MessageV1Codec.MessageV1({
      sourceChainSelector: sourceChainSelector,
      destChainSelector: destChainSelector,
      messageNumber: 1,
      executionGasLimit: 1,
      ccipReceiveGasLimit: 0,
      finality: bytes4(0),
      ccvAndExecutorHash: bytes32(0),
      onRampAddress: abi.encode(address(1)),
      offRampAddress: abi.encodePacked(address(2)),
      sender: abi.encode(sender),
      receiver: abi.encodePacked(receiver),
      destBlob: "",
      tokenTransfer: tokenTransfers,
      data: ""
    });
  }
}
