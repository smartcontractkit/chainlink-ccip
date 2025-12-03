// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {BaseVerifier} from "./components/BaseVerifier.sol";
import {ICrossChainVerifierV1} from "../interfaces/ICrossChainVerifierV1.sol";
import {MessageV1Codec} from "../libraries/MessageV1Codec.sol";
import {Client} from "../libraries/Client.sol";

import {Ownable2StepMsgSender} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2StepMsgSender.sol";

contract SimpleVerifier is Ownable2StepMsgSender, ICrossChainVerifierV1, BaseVerifier {
    struct Proof {
        bytes32 messageID;
        bool finalized;
    }


    string public constant override typeAndVersion = "SimpleVerifier 1.7.0-dev";
    /// @dev The preimage is bytes4(keccak256("SimpleVerifier 1.7.0")).
    bytes4 internal constant VERSION_TAG_V1_7_0 = 0x510b355b;

    constructor(string memory storageLoc)
        BaseVerifier(storageLoc)
    {}

    function forwardToVerifier(
        MessageV1Codec.MessageV1 calldata message,
        bytes32, // messageId
        address, // feeToken
        uint256, // feeTokenAmount
        bytes calldata // verifierArgs
    ) external view returns (bytes memory verifierReturnData) {
        // For EVM, sender is expected to be 20 bytes.
        address senderAddress = address(bytes20(message.sender));
        _assertSenderIsAllowed(message.destChainSelector, senderAddress);

        // TODO: Process msg & return verifier data
        return abi.encodePacked(VERSION_TAG_V1_7_0);
    }

    /// @inheritdoc ICrossChainVerifierV1
    // simple verification
    function verifyMessage(
        MessageV1Codec.MessageV1 calldata,
        bytes32 messageHash,
        bytes calldata verifierResults
    ) external pure override {
        Proof memory p = abi.decode(verifierResults, (Proof));
        require(p.messageID == messageHash, "SimpleVerifier: bad messageID");
        require(p.finalized, "SimpleVerifier: not finalized");
    }


    /// @notice Updates the storage location identifier.
    /// @param newLocation The new storage location identifier.
    function updateStorageLocation(
        string memory newLocation
    ) external onlyOwner {
        _setStorageLocation(newLocation);
    } 
}
