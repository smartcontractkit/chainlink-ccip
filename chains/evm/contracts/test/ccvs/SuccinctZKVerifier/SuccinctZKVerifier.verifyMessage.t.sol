// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ISP1Helios} from "../../../interfaces/succinct/ISP1Helios.sol";

import {SuccinctZKVerifier} from "../../../ccvs/SuccinctZKVerifier.sol";
import {BaseVerifier} from "../../../ccvs/components/BaseVerifier.sol";
import {MessageV1Codec} from "../../../libraries/MessageV1Codec.sol";
import {SuccinctZKVerifierFixture} from "./SuccinctZKVerifierFixture.sol";
import {SuccinctZKVerifierSetup} from "./SuccinctZKVerifierSetup.t.sol";

contract SuccinctZKVerifier_verifyMessage is SuccinctZKVerifierSetup {
  // The fixture receipt has an ERC20 Transfer log at index 1, which has three topics.
  address internal constant TRANSFER_LOG_EMITTER = 0x6846eF566e701136b2f77E3A7f21aAaDfc61B801;
  uint256 internal constant TRANSFER_LOG_INDEX = 1;

  function test_verifyMessage_HeaderChain() public view {
    s_zkVerifier.verifyMessage(
      _fixtureMessage(), SuccinctZKVerifierFixture.MESSAGE_ID, _encodeVerifierResults(_fixtureWitness())
    );
  }

  function test_verifyMessage_DirectAnchor() public {
    s_mockHelios.setAnchor(
      SuccinctZKVerifierFixture.MESSAGE_BLOCK_NUMBER, bytes32(0), SuccinctZKVerifierFixture.RECEIPTS_ROOT
    );
    SuccinctZKVerifier.Witness memory witness = _fixtureWitness();
    witness.anchorBlockNumber = SuccinctZKVerifierFixture.MESSAGE_BLOCK_NUMBER;
    witness.headers = new bytes[](0);

    s_zkVerifier.verifyMessage(_fixtureMessage(), SuccinctZKVerifierFixture.MESSAGE_ID, _encodeVerifierResults(witness));
  }

  function test_verifyMessage_HeaderChainAtMaxLength() public {
    SuccinctZKVerifier.Witness memory witness = _fixtureWitness();
    _setSourceChainConfig(SuccinctZKVerifierFixture.ON_RAMP, uint16(witness.headers.length));

    s_zkVerifier.verifyMessage(_fixtureMessage(), SuccinctZKVerifierFixture.MESSAGE_ID, _encodeVerifierResults(witness));
  }

  // Reverts

  function test_verifyMessage_RevertWhen_CursedByRMN() public {
    MessageV1Codec.MessageV1 memory message = _fixtureMessage();
    _setMockRMNChainCurse(message.sourceChainSelector, true);

    vm.expectRevert(abi.encodeWithSelector(BaseVerifier.CursedByRMN.selector, message.sourceChainSelector));
    s_zkVerifier.verifyMessage(message, SuccinctZKVerifierFixture.MESSAGE_ID, _encodeVerifierResults(_fixtureWitness()));
  }

  function test_verifyMessage_RevertWhen_SourceChainNotSupported() public {
    MessageV1Codec.MessageV1 memory message = _fixtureMessage();
    message.sourceChainSelector = SOURCE_CHAIN_SELECTOR;

    vm.expectRevert(abi.encodeWithSelector(SuccinctZKVerifier.SourceChainNotSupported.selector, SOURCE_CHAIN_SELECTOR));
    s_zkVerifier.verifyMessage(message, SuccinctZKVerifierFixture.MESSAGE_ID, _encodeVerifierResults(_fixtureWitness()));
  }

  function test_verifyMessage_RevertWhen_SourceChainNotSupported_Paused() public {
    MessageV1Codec.MessageV1 memory message = _fixtureMessage();
    SuccinctZKVerifier.SourceChainConfigArgs[] memory sourceChainConfigs =
      new SuccinctZKVerifier.SourceChainConfigArgs[](1);
    sourceChainConfigs[0] = SuccinctZKVerifier.SourceChainConfigArgs({
      helios: ISP1Helios(address(0)),
      sourceChainSelector: SuccinctZKVerifierFixture.SOURCE_CHAIN_SELECTOR,
      maxHeaderChainLength: MAX_HEADER_CHAIN_LENGTH,
      onRamp: SuccinctZKVerifierFixture.ON_RAMP
    });
    s_zkVerifier.applySourceChainConfigUpdates(sourceChainConfigs);

    vm.expectRevert(
      abi.encodeWithSelector(
        SuccinctZKVerifier.SourceChainNotSupported.selector, SuccinctZKVerifierFixture.SOURCE_CHAIN_SELECTOR
      )
    );
    s_zkVerifier.verifyMessage(message, SuccinctZKVerifierFixture.MESSAGE_ID, _encodeVerifierResults(_fixtureWitness()));
  }

  function test_verifyMessage_RevertWhen_InvalidVerifierResults_TooShort() public {
    MessageV1Codec.MessageV1 memory message = _fixtureMessage();

    vm.expectRevert(SuccinctZKVerifier.InvalidVerifierResults.selector);
    s_zkVerifier.verifyMessage(message, SuccinctZKVerifierFixture.MESSAGE_ID, hex"0102");
  }

  function test_verifyMessage_RevertWhen_InvalidCCVVersion() public {
    MessageV1Codec.MessageV1 memory message = _fixtureMessage();
    bytes memory verifierResults = _encodeVerifierResults(_fixtureWitness());
    verifierResults[0] ^= 0xff;

    vm.expectRevert(
      abi.encodeWithSelector(SuccinctZKVerifier.InvalidCCVVersion.selector, VERSION_TAG_V0_0_1, bytes4(verifierResults))
    );
    s_zkVerifier.verifyMessage(message, SuccinctZKVerifierFixture.MESSAGE_ID, verifierResults);
  }

  function test_verifyMessage_RevertWhen_BlockNotAnchored_HeaderChain() public {
    MessageV1Codec.MessageV1 memory message = _fixtureMessage();
    s_mockHelios.setAnchor(SuccinctZKVerifierFixture.ANCHOR_BLOCK_NUMBER, bytes32(0), bytes32(0));

    vm.expectRevert(
      abi.encodeWithSelector(
        SuccinctZKVerifier.BlockNotAnchored.selector,
        SuccinctZKVerifierFixture.SOURCE_CHAIN_SELECTOR,
        SuccinctZKVerifierFixture.ANCHOR_BLOCK_NUMBER
      )
    );
    s_zkVerifier.verifyMessage(message, SuccinctZKVerifierFixture.MESSAGE_ID, _encodeVerifierResults(_fixtureWitness()));
  }

  function test_verifyMessage_RevertWhen_BlockNotAnchored_DirectAnchor() public {
    MessageV1Codec.MessageV1 memory message = _fixtureMessage();
    SuccinctZKVerifier.Witness memory witness = _fixtureWitness();
    witness.anchorBlockNumber = SuccinctZKVerifierFixture.MESSAGE_BLOCK_NUMBER;
    witness.headers = new bytes[](0);

    vm.expectRevert(
      abi.encodeWithSelector(
        SuccinctZKVerifier.BlockNotAnchored.selector,
        SuccinctZKVerifierFixture.SOURCE_CHAIN_SELECTOR,
        SuccinctZKVerifierFixture.MESSAGE_BLOCK_NUMBER
      )
    );
    s_zkVerifier.verifyMessage(message, SuccinctZKVerifierFixture.MESSAGE_ID, _encodeVerifierResults(witness));
  }

  function test_verifyMessage_RevertWhen_HeaderChainTooLong() public {
    MessageV1Codec.MessageV1 memory message = _fixtureMessage();
    SuccinctZKVerifier.Witness memory witness = _fixtureWitness();
    uint16 maxHeaderChainLength = uint16(witness.headers.length - 1);
    _setSourceChainConfig(SuccinctZKVerifierFixture.ON_RAMP, maxHeaderChainLength);

    vm.expectRevert(
      abi.encodeWithSelector(
        SuccinctZKVerifier.HeaderChainTooLong.selector, witness.headers.length, maxHeaderChainLength
      )
    );
    s_zkVerifier.verifyMessage(message, SuccinctZKVerifierFixture.MESSAGE_ID, _encodeVerifierResults(witness));
  }

  function test_verifyMessage_RevertWhen_InvalidHeaderHash_FirstHeader() public {
    MessageV1Codec.MessageV1 memory message = _fixtureMessage();
    SuccinctZKVerifier.Witness memory witness = _fixtureWitness();
    witness.headers[0][40] ^= 0x01;

    vm.expectRevert(
      abi.encodeWithSelector(
        SuccinctZKVerifier.InvalidHeaderHash.selector,
        0,
        SuccinctZKVerifierFixture.ANCHOR_BLOCK_HASH,
        keccak256(witness.headers[0])
      )
    );
    s_zkVerifier.verifyMessage(message, SuccinctZKVerifierFixture.MESSAGE_ID, _encodeVerifierResults(witness));
  }

  function test_verifyMessage_RevertWhen_InvalidHeaderHash_LastHeader() public {
    MessageV1Codec.MessageV1 memory message = _fixtureMessage();
    SuccinctZKVerifier.Witness memory witness = _fixtureWitness();
    uint256 lastIndex = witness.headers.length - 1;
    witness.headers[lastIndex][40] ^= 0x01;

    vm.expectPartialRevert(SuccinctZKVerifier.InvalidHeaderHash.selector);
    s_zkVerifier.verifyMessage(message, SuccinctZKVerifierFixture.MESSAGE_ID, _encodeVerifierResults(witness));
  }

  function test_verifyMessage_RevertWhen_ProofNodeTampered() public {
    MessageV1Codec.MessageV1 memory message = _fixtureMessage();
    SuccinctZKVerifier.Witness memory witness = _fixtureWitness();
    witness.proofNodes[0][40] ^= 0x01;

    vm.expectRevert("MerkleTrie: invalid root hash");
    s_zkVerifier.verifyMessage(message, SuccinctZKVerifierFixture.MESSAGE_ID, _encodeVerifierResults(witness));
  }

  function test_verifyMessage_RevertWhen_WrongTxIndex() public {
    MessageV1Codec.MessageV1 memory message = _fixtureMessage();
    SuccinctZKVerifier.Witness memory witness = _fixtureWitness();
    witness.txIndex += 1;

    vm.expectRevert("MerkleTrie: invalid large internal hash");
    s_zkVerifier.verifyMessage(message, SuccinctZKVerifierFixture.MESSAGE_ID, _encodeVerifierResults(witness));
  }

  function test_verifyMessage_RevertWhen_LogIndexOutOfRange() public {
    MessageV1Codec.MessageV1 memory message = _fixtureMessage();
    SuccinctZKVerifier.Witness memory witness = _fixtureWitness();
    witness.logIndex = SuccinctZKVerifierFixture.LOG_COUNT;

    vm.expectRevert(
      abi.encodeWithSelector(
        SuccinctZKVerifier.LogIndexOutOfRange.selector, witness.logIndex, SuccinctZKVerifierFixture.LOG_COUNT
      )
    );
    s_zkVerifier.verifyMessage(message, SuccinctZKVerifierFixture.MESSAGE_ID, _encodeVerifierResults(witness));
  }

  function test_verifyMessage_RevertWhen_InvalidLogEmitter() public {
    MessageV1Codec.MessageV1 memory message = _fixtureMessage();
    address otherOnRamp = makeAddr("otherOnRamp");
    _setSourceChainConfig(otherOnRamp, MAX_HEADER_CHAIN_LENGTH);

    vm.expectRevert(
      abi.encodeWithSelector(
        SuccinctZKVerifier.InvalidLogEmitter.selector, otherOnRamp, abi.encodePacked(SuccinctZKVerifierFixture.ON_RAMP)
      )
    );
    s_zkVerifier.verifyMessage(message, SuccinctZKVerifierFixture.MESSAGE_ID, _encodeVerifierResults(_fixtureWitness()));
  }

  function test_verifyMessage_RevertWhen_InvalidTopicCount() public {
    MessageV1Codec.MessageV1 memory message = _fixtureMessage();
    SuccinctZKVerifier.Witness memory witness = _fixtureWitness();
    witness.logIndex = TRANSFER_LOG_INDEX;
    _setSourceChainConfig(TRANSFER_LOG_EMITTER, MAX_HEADER_CHAIN_LENGTH);

    vm.expectRevert(abi.encodeWithSelector(SuccinctZKVerifier.InvalidTopicCount.selector, 3));
    s_zkVerifier.verifyMessage(message, SuccinctZKVerifierFixture.MESSAGE_ID, _encodeVerifierResults(witness));
  }

  function test_verifyMessage_RevertWhen_InvalidMessageId() public {
    MessageV1Codec.MessageV1 memory message = _fixtureMessage();
    bytes32 wrongMessageId = keccak256("wrong");

    vm.expectRevert(
      abi.encodeWithSelector(
        SuccinctZKVerifier.InvalidMessageId.selector, wrongMessageId, SuccinctZKVerifierFixture.MESSAGE_ID
      )
    );
    s_zkVerifier.verifyMessage(message, wrongMessageId, _encodeVerifierResults(_fixtureWitness()));
  }
}
