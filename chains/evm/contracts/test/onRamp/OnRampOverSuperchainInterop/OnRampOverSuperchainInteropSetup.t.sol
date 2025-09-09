// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Router} from "../../../Router.sol";
import {Client} from "../../../libraries/Client.sol";
import {Internal} from "../../../libraries/Internal.sol";

import {OnRampOverSuperchainInteropHelper} from "../../helpers/OnRampOverSuperchainInteropHelper.sol";
import {OnRampSetup} from "../OnRamp/OnRampSetup.t.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

contract OnRampOverSuperchainInteropSetup is OnRampSetup {
  OnRampOverSuperchainInteropHelper internal s_onRampOverSuperchainInterop;

  function setUp() public virtual override {
    super.setUp();

    s_onRampOverSuperchainInterop = new OnRampOverSuperchainInteropHelper(
      s_onRamp.getStaticConfig(), s_onRamp.getDynamicConfig(), _generateDestChainConfigArgs(s_sourceRouter)
    );

    // Authorize the SuperchainInterop Onramp to call the NonceManager
    address[] memory authorizedCallers = new address[](1);
    authorizedCallers[0] = address(s_onRampOverSuperchainInterop);
    s_outboundNonceManager.applyAuthorizedCallerUpdates(
      AuthorizedCallers.AuthorizedCallerArgs({addedCallers: authorizedCallers, removedCallers: new address[](0)})
    );

    // Authorize the SuperchainInterop Onramp to call TokenPools
    Router.OnRamp[] memory onRampUpdates = new Router.OnRamp[](1);
    onRampUpdates[0] =
      Router.OnRamp({destChainSelector: DEST_CHAIN_SELECTOR, onRamp: address(s_onRampOverSuperchainInterop)});

    s_sourceRouter.applyRampUpdates(onRampUpdates, new Router.OffRamp[](0), new Router.OffRamp[](0));

    assertEq("OnRampOverSuperchainInterop 1.6.2-dev", s_onRampOverSuperchainInterop.typeAndVersion());
  }

  function _EVM2AnyRampMessageToAny2EVMRampMessage(
    Internal.EVM2AnyRampMessage memory message
  ) internal pure returns (Internal.Any2EVMRampMessage memory) {
    Internal.Any2EVMTokenTransfer[] memory tokenTransfers =
      new Internal.Any2EVMTokenTransfer[](message.tokenAmounts.length);

    for (uint256 i = 0; i < message.tokenAmounts.length; ++i) {
      tokenTransfers[i] = Internal.Any2EVMTokenTransfer({
        sourcePoolAddress: abi.encode(message.tokenAmounts[i].sourcePoolAddress),
        destTokenAddress: abi.decode(message.tokenAmounts[i].destTokenAddress, (address)),
        destGasAmount: abi.decode(message.tokenAmounts[i].destExecData, (uint32)),
        extraData: message.tokenAmounts[i].extraData,
        amount: message.tokenAmounts[i].amount
      });
    }

    bytes memory gasLimitBytes = new bytes(32);
    for (uint256 i = 0; i < 32; ++i) {
      gasLimitBytes[i] = message.extraArgs[4 + i];
    }
    uint256 gasLimit = abi.decode(gasLimitBytes, (uint256));

    return Internal.Any2EVMRampMessage({
      header: Internal.RampMessageHeader({
        messageId: message.header.messageId,
        sourceChainSelector: message.header.sourceChainSelector,
        destChainSelector: message.header.destChainSelector,
        sequenceNumber: message.header.sequenceNumber,
        nonce: message.header.nonce
      }),
      sender: abi.encode(message.sender),
      data: message.data,
      receiver: abi.decode(message.receiver, (address)),
      gasLimit: gasLimit,
      tokenAmounts: tokenTransfers
    });
  }

  function _generateBasicAny2EVMMessage() internal returns (Internal.Any2EVMRampMessage memory) {
    Internal.Any2EVMTokenTransfer[] memory tokenTransfers = new Internal.Any2EVMTokenTransfer[](0);

    return Internal.Any2EVMRampMessage({
      header: Internal.RampMessageHeader({
        messageId: keccak256("test-message"),
        sourceChainSelector: SOURCE_CHAIN_SELECTOR,
        destChainSelector: DEST_CHAIN_SELECTOR,
        sequenceNumber: 1,
        nonce: 1
      }),
      sender: abi.encode(OWNER),
      data: abi.encode("test data"),
      receiver: makeAddr("receiver"),
      gasLimit: 200000,
      tokenAmounts: tokenTransfers
    });
  }

  function _generateAny2EVMMessageWithTokens() internal returns (Internal.Any2EVMRampMessage memory) {
    Internal.Any2EVMTokenTransfer[] memory tokenTransfers = new Internal.Any2EVMTokenTransfer[](1);
    tokenTransfers[0] = Internal.Any2EVMTokenTransfer({
      sourcePoolAddress: abi.encode(makeAddr("sourcePool")),
      destTokenAddress: makeAddr("destToken"),
      destGasAmount: 50000,
      extraData: "",
      amount: 1000e18
    });

    return Internal.Any2EVMRampMessage({
      header: Internal.RampMessageHeader({
        messageId: keccak256("test-message-with-tokens"),
        sourceChainSelector: SOURCE_CHAIN_SELECTOR,
        destChainSelector: DEST_CHAIN_SELECTOR,
        sequenceNumber: 2,
        nonce: 2
      }),
      sender: abi.encode(OWNER),
      data: abi.encode("test data with tokens"),
      receiver: makeAddr("receiver"),
      gasLimit: 300000,
      tokenAmounts: tokenTransfers
    });
  }

  function _generateAny2EVMMessageWithCustomFields(
    bytes memory sender,
    bytes memory data,
    address receiver,
    uint256 gasLimit,
    uint64 sequenceNumber,
    uint64 nonce
  ) internal pure returns (Internal.Any2EVMRampMessage memory) {
    Internal.Any2EVMTokenTransfer[] memory tokenTransfers = new Internal.Any2EVMTokenTransfer[](0);

    return Internal.Any2EVMRampMessage({
      header: Internal.RampMessageHeader({
        messageId: keccak256("custom-message"),
        sourceChainSelector: SOURCE_CHAIN_SELECTOR,
        destChainSelector: DEST_CHAIN_SELECTOR,
        sequenceNumber: sequenceNumber,
        nonce: nonce
      }),
      sender: sender,
      data: data,
      receiver: receiver,
      gasLimit: gasLimit,
      tokenAmounts: tokenTransfers
    });
  }

  function _getOffRampMetadataHash() internal view returns (bytes32) {
    return keccak256(
      abi.encode(
        Internal.ANY_2_EVM_MESSAGE_HASH,
        SOURCE_CHAIN_SELECTOR,
        DEST_CHAIN_SELECTOR,
        keccak256(abi.encode(address(s_onRampOverSuperchainInterop)))
      )
    );
  }

  function _getOnRampMetadataHash(
    uint64 destChainSelector
  ) internal view returns (bytes32) {
    return keccak256(
      abi.encode(
        Internal.EVM_2_ANY_MESSAGE_HASH,
        SOURCE_CHAIN_SELECTOR,
        destChainSelector,
        address(s_onRampOverSuperchainInterop)
      )
    );
  }

  function _generateInitialSourceDestMessages(
    uint64 destChainSelector,
    Client.EVM2AnyMessage memory message,
    uint256 feeAmount
  ) internal view returns (Internal.EVM2AnyRampMessage memory, Internal.Any2EVMRampMessage memory) {
    // Need to pass in custom metadata hash because s_onRampOverSuperchainInterop is not the same address as s_onRamp
    Internal.EVM2AnyRampMessage memory evm2AnyMessage = _evmMessageToEvent(
      message,
      SOURCE_CHAIN_SELECTOR,
      1,
      1,
      feeAmount,
      feeAmount,
      OWNER,
      _getOnRampMetadataHash(destChainSelector),
      s_tokenAdminRegistry
    );

    return (evm2AnyMessage, _EVM2AnyRampMessageToAny2EVMRampMessage(evm2AnyMessage));
  }
}
