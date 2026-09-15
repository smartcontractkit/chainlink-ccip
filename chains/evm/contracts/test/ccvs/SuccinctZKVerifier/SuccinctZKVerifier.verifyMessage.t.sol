// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IExecutionRootOracle} from "../../../interfaces/IExecutionRootOracle.sol";

import {SuccinctZKVerifier} from "../../../ccvs/SuccinctZKVerifier.sol";
import {MessageV1Codec} from "../../../libraries/MessageV1Codec.sol";
import {BaseVerifierSetup} from "../components/BaseVerifier/BaseVerifierSetup.t.sol";
import {SuccinctZKVerifierFixture} from "./SuccinctZKVerifierFixture.sol";

/// @notice Root oracle stub with explicitly anchored blocks.
contract StubExecutionRootOracle is IExecutionRootOracle {
  mapping(uint256 blockNumber => bytes32 blockHash) internal s_blockHashes;
  mapping(uint256 blockNumber => bytes32 receiptsRoot) internal s_receiptsRoots;

  function anchor(
    uint256 blockNumber,
    bytes32 blockHash,
    bytes32 receiptsRoot
  ) external {
    s_blockHashes[blockNumber] = blockHash;
    s_receiptsRoots[blockNumber] = receiptsRoot;
  }

  function executionBlockHashes(
    uint256 blockNumber
  ) external view returns (bytes32) {
    return s_blockHashes[blockNumber];
  }

  function executionReceiptsRoots(
    uint256 blockNumber
  ) external view returns (bytes32) {
    return s_receiptsRoots[blockNumber];
  }
}

contract SuccinctZKVerifier_verifyMessage is BaseVerifierSetup {
  bytes4 internal constant VERSION_TAG = bytes4(keccak256("SuccinctZKVerifier 2.0.0"));

  SuccinctZKVerifier internal s_verifier;
  StubExecutionRootOracle internal s_oracle;

  function setUp() public virtual override {
    super.setUp();

    s_oracle = new StubExecutionRootOracle();
    s_oracle.anchor(
      SuccinctZKVerifierFixture.ANCHOR_BLOCK_NUMBER, SuccinctZKVerifierFixture.ANCHOR_BLOCK_HASH, bytes32(0)
    );

    s_verifier = new SuccinctZKVerifier(s_storageLocations, address(s_mockRMNRemote), VERSION_TAG);

    SuccinctZKVerifier.SourceChainConfigArgs[] memory configs = new SuccinctZKVerifier.SourceChainConfigArgs[](1);
    configs[0] = SuccinctZKVerifier.SourceChainConfigArgs({
      sourceChainSelector: SuccinctZKVerifierFixture.SOURCE_CHAIN_SELECTOR,
      rootOracle: s_oracle,
      onRamp: SuccinctZKVerifierFixture.ON_RAMP,
      maxHeaderChainLength: 64
    });
    s_verifier.applySourceChainConfigUpdates(configs);
  }

  /// @dev Exposes the codec, which needs calldata, to setUp.
  function decodeMessage(
    bytes calldata encodedMessage
  ) external pure returns (MessageV1Codec.MessageV1 memory) {
    return MessageV1Codec._decodeMessageV1(encodedMessage);
  }

  function _message() internal view returns (MessageV1Codec.MessageV1 memory) {
    return this.decodeMessage(SuccinctZKVerifierFixture.encodedMessage());
  }

  function test_verifyMessage_HeaderChain() public view {
    s_verifier.verifyMessage(_message(), SuccinctZKVerifierFixture.MESSAGE_ID, SuccinctZKVerifierFixture.envelope());
  }

  function test_verifyMessage_DirectAnchor() public {
    s_oracle.anchor(SuccinctZKVerifierFixture.MESSAGE_BLOCK_NUMBER, bytes32(0), SuccinctZKVerifierFixture.RECEIPTS_ROOT);

    s_verifier.verifyMessage(
      _message(), SuccinctZKVerifierFixture.MESSAGE_ID, _envelopeWithoutHeaders(SuccinctZKVerifierFixture.envelope())
    );
  }

  function test_verifyMessage_RevertWhen_MessageIdMismatch() public {
    MessageV1Codec.MessageV1 memory message = _message();
    bytes32 wrongMessageId = keccak256("wrong");
    vm.expectRevert(
      abi.encodeWithSelector(
        SuccinctZKVerifier.InvalidMessageId.selector, wrongMessageId, SuccinctZKVerifierFixture.MESSAGE_ID
      )
    );
    s_verifier.verifyMessage(message, wrongMessageId, SuccinctZKVerifierFixture.envelope());
  }

  function test_verifyMessage_RevertWhen_AnchorNotAvailable() public {
    MessageV1Codec.MessageV1 memory message = _message();
    s_oracle.anchor(SuccinctZKVerifierFixture.ANCHOR_BLOCK_NUMBER, bytes32(0), bytes32(0));
    vm.expectRevert(
      abi.encodeWithSelector(
        SuccinctZKVerifier.AnchorNotAvailable.selector,
        SuccinctZKVerifierFixture.SOURCE_CHAIN_SELECTOR,
        SuccinctZKVerifierFixture.ANCHOR_BLOCK_NUMBER
      )
    );
    s_verifier.verifyMessage(message, SuccinctZKVerifierFixture.MESSAGE_ID, SuccinctZKVerifierFixture.envelope());
  }

  function test_verifyMessage_RevertWhen_HeaderTampered() public {
    MessageV1Codec.MessageV1 memory message = _message();
    bytes memory envelope = SuccinctZKVerifierFixture.envelope();
    // Flip one byte inside the first header.
    envelope[4 + 8 + 2 + 2 + 40] ^= 0x01;
    vm.expectPartialRevert(SuccinctZKVerifier.HeaderHashMismatch.selector);
    s_verifier.verifyMessage(message, SuccinctZKVerifierFixture.MESSAGE_ID, envelope);
  }

  function test_verifyMessage_RevertWhen_WrongVersion() public {
    MessageV1Codec.MessageV1 memory message = _message();
    bytes memory envelope = SuccinctZKVerifierFixture.envelope();
    envelope[0] ^= 0xff;
    vm.expectRevert(abi.encodeWithSelector(SuccinctZKVerifier.InvalidCCVVersion.selector, bytes4(envelope)));
    s_verifier.verifyMessage(message, SuccinctZKVerifierFixture.MESSAGE_ID, envelope);
  }

  function test_verifyMessage_RevertWhen_WrongOnRamp() public {
    MessageV1Codec.MessageV1 memory message = _message();
    SuccinctZKVerifier.SourceChainConfigArgs[] memory configs = new SuccinctZKVerifier.SourceChainConfigArgs[](1);
    configs[0] = SuccinctZKVerifier.SourceChainConfigArgs({
      sourceChainSelector: SuccinctZKVerifierFixture.SOURCE_CHAIN_SELECTOR,
      rootOracle: s_oracle,
      onRamp: makeAddr("otherOnRamp"),
      maxHeaderChainLength: 64
    });
    s_verifier.applySourceChainConfigUpdates(configs);

    vm.expectRevert(
      abi.encodeWithSelector(
        SuccinctZKVerifier.UnexpectedLogEmitter.selector, SuccinctZKVerifierFixture.ON_RAMP, configs[0].onRamp
      )
    );
    s_verifier.verifyMessage(message, SuccinctZKVerifierFixture.MESSAGE_ID, SuccinctZKVerifierFixture.envelope());
  }

  /// @dev Rewrites the envelope so the anchor is the message block and the header chain is empty.
  function _envelopeWithoutHeaders(
    bytes memory envelope
  ) internal pure returns (bytes memory) {
    uint256 offset = 4 + 8;
    uint256 headerCount = uint16(bytes2(_slice(envelope, offset, 2)));
    offset += 2;
    for (uint256 i = 0; i < headerCount; ++i) {
      uint256 length = uint16(bytes2(_slice(envelope, offset, 2)));
      offset += 2 + length;
    }
    return bytes.concat(
      _slice(envelope, 0, 4),
      bytes8(SuccinctZKVerifierFixture.MESSAGE_BLOCK_NUMBER),
      bytes2(0),
      _slice(envelope, offset, envelope.length - offset)
    );
  }

  function _slice(
    bytes memory data,
    uint256 start,
    uint256 length
  ) internal pure returns (bytes memory out) {
    out = new bytes(length);
    for (uint256 i = 0; i < length; ++i) {
      out[i] = data[start + i];
    }
  }
}
