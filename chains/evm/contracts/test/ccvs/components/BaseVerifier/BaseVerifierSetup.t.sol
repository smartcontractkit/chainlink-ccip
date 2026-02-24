// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IRouter} from "../../../../interfaces/IRouter.sol";

import {BaseVerifier} from "../../../../ccvs/components/BaseVerifier.sol";
import {MessageV1Codec} from "../../../../libraries/MessageV1Codec.sol";
import {BaseERC20} from "../../../../tmp/BaseERC20.sol";
import {CrossChainToken} from "../../../../tmp/CrossChainToken.sol";
import {FeeQuoterSetup} from "../../../feeQuoter/FeeQuoterSetup.t.sol";
import {BaseVerifierTestHelper} from "../../../helpers/BaseVerifierTestHelper.sol";

contract BaseVerifierSetup is FeeQuoterSetup {
  address internal constant FEE_AGGREGATOR = 0xa33CDB32eAEce34F6affEfF4899cef45744EDea3;
  address internal constant ALLOWLIST_ADMIN = 0x1234567890123456789012345678901234567890;

  BaseVerifierTestHelper internal s_baseVerifier;

  IRouter internal s_router;
  address internal s_onRamp;
  address internal s_offRamp;

  string[] internal s_storageLocations;

  function setUp() public virtual override {
    super.setUp();

    s_storageLocations.push("testStorageLocation");

    s_router = IRouter(makeAddr("Router"));
    s_onRamp = makeAddr("OnRamp");
    vm.mockCall(
      address(s_router), abi.encodeWithSelector(IRouter.getOnRamp.selector, DEST_CHAIN_SELECTOR), abi.encode(s_onRamp)
    );
    s_offRamp = makeAddr("OffRamp");
    s_sourceFeeToken = address(
      new CrossChainToken(
        BaseERC20.ConstructorParams({
          name: "Chainlink Token", symbol: "LINK", decimals: 18, maxSupply: 0, preMint: 0, ccipAdmin: OWNER
        }),
        OWNER,
        OWNER
      )
    );

    s_baseVerifier = new BaseVerifierTestHelper(s_storageLocations, address(s_mockRMNRemote));

    // Set up initial destination chain config.
    BaseVerifier.RemoteChainConfigArgs[] memory remoteChainConfigs = new BaseVerifier.RemoteChainConfigArgs[](1);
    remoteChainConfigs[0] = BaseVerifier.RemoteChainConfigArgs({
      router: s_router,
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      allowlistEnabled: false,
      feeUSDCents: DEFAULT_CCV_FEE_USD_CENTS,
      gasForVerification: DEFAULT_CCV_GAS_LIMIT,
      payloadSizeBytes: DEFAULT_CCV_PAYLOAD_SIZE
    });

    s_baseVerifier.applyRemoteChainConfigUpdates(remoteChainConfigs);

    vm.startPrank(OWNER);
  }

  function _getRemoteChainConfig(
    IRouter router,
    uint64 remoteChainSelector,
    bool allowlistEnabled
  ) internal pure returns (BaseVerifier.RemoteChainConfigArgs memory) {
    return BaseVerifier.RemoteChainConfigArgs({
      router: router,
      remoteChainSelector: remoteChainSelector,
      allowlistEnabled: allowlistEnabled,
      feeUSDCents: DEFAULT_CCV_FEE_USD_CENTS,
      gasForVerification: DEFAULT_CCV_GAS_LIMIT,
      payloadSizeBytes: DEFAULT_CCV_PAYLOAD_SIZE
    });
  }

  function _getAllowlistConfig(
    uint64 destChainSelector,
    bool allowlistEnabled,
    address[] memory addedSenders,
    address[] memory removedSenders
  ) internal pure returns (BaseVerifier.AllowlistConfigArgs memory) {
    return BaseVerifier.AllowlistConfigArgs({
      destChainSelector: destChainSelector,
      allowlistEnabled: allowlistEnabled,
      addedAllowlistedSenders: addedSenders,
      removedAllowlistedSenders: removedSenders
    });
  }

  /// @notice Creates a basic MessageV1 without token transfers.
  function _createBasicMessageV1(
    uint64 sourceChainSelector
  ) internal pure returns (MessageV1Codec.MessageV1 memory) {
    return MessageV1Codec.MessageV1({
      sourceChainSelector: sourceChainSelector,
      destChainSelector: DEST_CHAIN_SELECTOR,
      messageNumber: 1,
      executionGasLimit: GAS_LIMIT * 2,
      ccipReceiveGasLimit: GAS_LIMIT,
      finality: 0,
      ccvAndExecutorHash: bytes32(0),
      onRampAddress: abi.encode(address(0x1111111111111111111111111111111111111111)),
      offRampAddress: abi.encodePacked(address(0x2222222222222222222222222222222222222222)),
      sender: abi.encode(address(0x3333333333333333333333333333333333333333)),
      receiver: abi.encodePacked(address(0x4444444444444444444444444444444444444444)),
      destBlob: "",
      tokenTransfer: new MessageV1Codec.TokenTransferV1[](0),
      data: ""
    });
  }

  /// @notice Creates a MessageV1 with a single token transfer.
  function _createMessageV1WithTokenTransfer(
    uint64 sourceChainSelector,
    uint64 destChainSelector,
    uint16 finality,
    address sourceTokenAddress,
    uint256 amount,
    bytes memory tokenReceiver
  ) internal returns (MessageV1Codec.MessageV1 memory, bytes32 messageId) {
    MessageV1Codec.TokenTransferV1[] memory tokenTransfer = new MessageV1Codec.TokenTransferV1[](1);
    tokenTransfer[0] = MessageV1Codec.TokenTransferV1({
      amount: amount,
      sourcePoolAddress: abi.encode(makeAddr("sourcePool")),
      sourceTokenAddress: abi.encode(sourceTokenAddress),
      destTokenAddress: abi.encodePacked(makeAddr("destToken")),
      tokenReceiver: tokenReceiver,
      extraData: ""
    });

    MessageV1Codec.MessageV1 memory message = MessageV1Codec.MessageV1({
      sourceChainSelector: sourceChainSelector,
      destChainSelector: destChainSelector,
      messageNumber: 1,
      executionGasLimit: 400_000,
      ccipReceiveGasLimit: 200_000,
      finality: finality,
      ccvAndExecutorHash: bytes32(0),
      onRampAddress: abi.encode(address(0x1111111111111111111111111111111111111111)),
      offRampAddress: abi.encodePacked(address(0x2222222222222222222222222222222222222222)),
      sender: abi.encode(address(0x3333333333333333333333333333333333333333)),
      receiver: abi.encodePacked(address(0x4444444444444444444444444444444444444444)),
      destBlob: "",
      tokenTransfer: tokenTransfer,
      data: ""
    });

    messageId = keccak256(MessageV1Codec._encodeMessageV1(message));
    return (message, messageId);
  }
}
