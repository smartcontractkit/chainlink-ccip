// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {SuccinctZKVerifier} from "../../../ccvs/SuccinctZKVerifier.sol";
import {MockSP1Helios} from "../../mocks/MockSP1Helios.sol";
import {SuccinctZKVerifierSetup} from "./SuccinctZKVerifierSetup.t.sol";

contract SuccinctZKVerifier_proveBlockHash is SuccinctZKVerifierSetup {
  uint256 internal constant HEADER_COUNT = 3;

  bytes[] internal s_headers;

  function setUp() public override {
    super.setUp();

    (, bytes32 messageId) = _messageWithId();
    s_headers = _buildWitness(_encodeMessageSentReceipt(s_sourceOnRamp, messageId), HEADER_COUNT).headers;
  }

  function test_proveBlockHash() public {
    uint256 provenBlockNumber = ANCHOR_BLOCK_NUMBER - HEADER_COUNT;
    bytes32 provenBlockHash = keccak256("parent");

    vm.expectEmit();
    emit SuccinctZKVerifier.BlockHashProven(SOURCE_CHAIN_SELECTOR, provenBlockNumber, provenBlockHash);

    s_zkVerifier.proveBlockHash(SOURCE_CHAIN_SELECTOR, ANCHOR_BLOCK_NUMBER, s_headers);

    assertEq(provenBlockHash, s_zkVerifier.getProvenBlockHash(SOURCE_CHAIN_SELECTOR, provenBlockNumber));
  }

  function test_proveBlockHash_FromProvenBlock() public {
    bytes[] memory headers = new bytes[](1);
    headers[0] = s_headers[0];
    s_zkVerifier.proveBlockHash(SOURCE_CHAIN_SELECTOR, ANCHOR_BLOCK_NUMBER, headers);

    headers[0] = s_headers[1];
    vm.expectEmit();
    emit SuccinctZKVerifier.BlockHashProven(SOURCE_CHAIN_SELECTOR, ANCHOR_BLOCK_NUMBER - 2, keccak256(s_headers[2]));

    s_zkVerifier.proveBlockHash(SOURCE_CHAIN_SELECTOR, ANCHOR_BLOCK_NUMBER - 1, headers);
  }

  function test_getProvenBlockHash_AfterLightClientReplaced() public {
    s_zkVerifier.proveBlockHash(SOURCE_CHAIN_SELECTOR, ANCHOR_BLOCK_NUMBER, s_headers);

    MockSP1Helios otherHelios = new MockSP1Helios();
    otherHelios.setVkeys(LIGHT_CLIENT_VKEY, EXECUTION_HEADER_VKEY);
    _setSourceChainConfig(otherHelios);

    assertEq(bytes32(0), s_zkVerifier.getProvenBlockHash(SOURCE_CHAIN_SELECTOR, ANCHOR_BLOCK_NUMBER - HEADER_COUNT));
  }

  // Reverts

  function test_proveBlockHash_RevertWhen_SourceChainNotSupported() public {
    _setSourceChainConfig(MockSP1Helios(address(0)));

    vm.expectRevert(abi.encodeWithSelector(SuccinctZKVerifier.SourceChainNotSupported.selector, SOURCE_CHAIN_SELECTOR));
    s_zkVerifier.proveBlockHash(SOURCE_CHAIN_SELECTOR, ANCHOR_BLOCK_NUMBER, s_headers);
  }

  function test_proveBlockHash_RevertWhen_BlockNotAnchored() public {
    vm.expectRevert(
      abi.encodeWithSelector(
        SuccinctZKVerifier.BlockNotAnchored.selector, SOURCE_CHAIN_SELECTOR, ANCHOR_BLOCK_NUMBER - 1
      )
    );
    s_zkVerifier.proveBlockHash(SOURCE_CHAIN_SELECTOR, ANCHOR_BLOCK_NUMBER - 1, s_headers);
  }

  function test_proveBlockHash_RevertWhen_InvalidHeaderHash() public {
    bytes32 parentHash = keccak256(s_headers[1]);
    s_headers[1][40] ^= 0x01;

    vm.expectRevert(
      abi.encodeWithSelector(SuccinctZKVerifier.InvalidHeaderHash.selector, 1, parentHash, keccak256(s_headers[1]))
    );
    s_zkVerifier.proveBlockHash(SOURCE_CHAIN_SELECTOR, ANCHOR_BLOCK_NUMBER, s_headers);
  }
}
