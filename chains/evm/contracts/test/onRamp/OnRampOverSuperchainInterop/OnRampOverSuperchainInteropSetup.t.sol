// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Internal} from "../../../libraries/Internal.sol";
import {OnRampOverSuperchainInterop} from "../../../onRamp/OnRampOverSuperchainInterop.sol";
import {OnRampSetup} from "../OnRamp/OnRampSetup.t.sol";

contract OnRampOverSuperchainInteropSetup is OnRampSetup {
  OnRampOverSuperchainInterop internal s_onRampOverSuperchainInterop;

  function setUp() public virtual override {
    super.setUp();

    s_onRampOverSuperchainInterop = new OnRampOverSuperchainInterop(
      s_onRamp.getStaticConfig(), s_onRamp.getDynamicConfig(), _generateDestChainConfigArgs(s_sourceRouter)
    );
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
  ) internal view returns (Internal.Any2EVMRampMessage memory) {
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
}
