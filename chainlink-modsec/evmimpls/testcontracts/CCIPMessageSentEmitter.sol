// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

contract CCIPMessageSentEmitter {
    mapping(uint64 => mapping(address => uint64)) public s_nonces;
    mapping(uint64 => mapping(uint64 => bool)) public s_executed;

    struct Header {
        bytes32 messageId;
        uint64 sourceChainSelector;
        uint64 destChainSelector;
        uint64 sequenceNumber;
    }

    enum ReceiptType {
        Verifier,
        Executor
    }

    struct Receipt {
        ReceiptType receiptType;
        address issuer;
        uint256 feeTokenAmount;
        uint64 destGasLimit;
        uint32 destBytesOverhead;
        bytes extraArgs;
    }

    struct EVMTokenTransfer {
        address sourceTokenAddress;
        address sourcePoolAddress;
        bytes destTokenAddress;
        bytes extraData;
        uint256 amount;
        bytes destExecData;
        bytes32 requiredVerifierId;
    }

    struct EVM2AnyVerifierMessage {
        Header header;
        address sender;
        bytes data;
        bytes receiver;
        address feeToken;
        uint256 feeTokenAmount;
        uint256 feeValueJuels;
        EVMTokenTransfer tokenTransfer;
        Receipt[] receipts;
    }

    struct Any2EVMMultiProofTokenTransfer {
        bytes sourcePoolAddress;
        address destTokenAddress;
        bytes extraData;
        uint256 amount;
    }

    struct Any2EVMMultiProofMessage {
        Header header;
        bytes sender;
        bytes data;
        address receiver;
        uint32 gasLimit;
        Any2EVMMultiProofTokenTransfer[] tokenAmounts;
        bytes[] requiredVerifiers;
    }

    event CCIPMessageSent(
        uint64 indexed destChainSelector,
        uint64 indexed sequenceNumber,
        EVM2AnyVerifierMessage message
    );

    event Executed(
        Any2EVMMultiProofMessage message,
        bytes[] proofs
    );

    function emitCCIPMessageSent(EVM2AnyVerifierMessage memory message) public {
        emit CCIPMessageSent(message.header.destChainSelector, message.header.sequenceNumber, message);
    }

    function setNonce(uint64 sourceChainSelector, address account, uint64 nonce) public {
        s_nonces[sourceChainSelector][account] = nonce;
    }

    function getNonce(uint64 sourceChainSelector, address account) public view returns (uint64) {
        return s_nonces[sourceChainSelector][account];
    }

    function isExecuted(uint64 sourceChainSelector, uint64 sequenceNumber) public view returns (bool) {
        return s_executed[sourceChainSelector][sequenceNumber];
    }

    function setExecuted(uint64 sourceChainSelector, uint64 sequenceNumber, bool executed) public {
        s_executed[sourceChainSelector][sequenceNumber] = executed;
    }

    function execute(Any2EVMMultiProofMessage memory message, bytes[] memory proofs) public {
        emit Executed(message, proofs);
    }
}
