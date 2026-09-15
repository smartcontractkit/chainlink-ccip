// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IRouter} from "../../../interfaces/IRouter.sol";

import {SuccinctZKVerifier} from "../../../ccvs/SuccinctZKVerifier.sol";
import {BaseVerifier} from "../../../ccvs/components/BaseVerifier.sol";
import {MessageV1Codec} from "../../../libraries/MessageV1Codec.sol";
import {MockSP1Helios} from "../../mocks/MockSP1Helios.sol";
import {BaseVerifierSetup} from "../components/BaseVerifier/BaseVerifierSetup.t.sol";
import {SuccinctZKVerifierFixture} from "./SuccinctZKVerifierFixture.sol";

contract SuccinctZKVerifierSetup is BaseVerifierSetup {
  bytes4 internal constant VERSION_TAG_V0_0_1 = bytes4(keccak256("SuccinctZKVerifier 0.0.1-dev"));
  uint16 internal constant MAX_HEADER_CHAIN_LENGTH = 64;

  SuccinctZKVerifier internal s_zkVerifier;
  MockSP1Helios internal s_mockHelios;

  function setUp() public virtual override {
    super.setUp();

    s_mockHelios = new MockSP1Helios();
    s_mockHelios.setAnchor(
      SuccinctZKVerifierFixture.ANCHOR_BLOCK_NUMBER, SuccinctZKVerifierFixture.ANCHOR_BLOCK_HASH, bytes32(0)
    );

    s_zkVerifier = new SuccinctZKVerifier(
      SuccinctZKVerifier.DynamicConfig({feeAggregator: FEE_AGGREGATOR}),
      s_storageLocations,
      address(s_mockRMNRemote),
      VERSION_TAG_V0_0_1
    );

    BaseVerifier.RemoteChainConfigArgs[] memory remoteChainConfigs = new BaseVerifier.RemoteChainConfigArgs[](1);
    remoteChainConfigs[0] = _getRemoteChainConfig(s_router, DEST_CHAIN_SELECTOR, false);
    s_zkVerifier.applyRemoteChainConfigUpdates(remoteChainConfigs);

    _setSourceChainConfig(SuccinctZKVerifierFixture.ON_RAMP, MAX_HEADER_CHAIN_LENGTH);

    vm.mockCall(address(s_router), abi.encodeCall(IRouter.getOnRamp, (DEST_CHAIN_SELECTOR)), abi.encode(s_onRamp));
  }

  function _setSourceChainConfig(
    address onRamp,
    uint16 maxHeaderChainLength
  ) internal {
    SuccinctZKVerifier.SourceChainConfigArgs[] memory sourceChainConfigs =
      new SuccinctZKVerifier.SourceChainConfigArgs[](1);
    sourceChainConfigs[0] = SuccinctZKVerifier.SourceChainConfigArgs({
      helios: s_mockHelios,
      sourceChainSelector: SuccinctZKVerifierFixture.SOURCE_CHAIN_SELECTOR,
      maxHeaderChainLength: maxHeaderChainLength,
      onRamp: onRamp
    });
    s_zkVerifier.applySourceChainConfigUpdates(sourceChainConfigs);
  }

  /// @notice Exposes the codec, which needs calldata, to the tests.
  function decodeMessage(
    bytes calldata encodedMessage
  ) external pure returns (MessageV1Codec.MessageV1 memory) {
    return MessageV1Codec._decodeMessageV1(encodedMessage);
  }

  function _fixtureMessage() internal view returns (MessageV1Codec.MessageV1 memory) {
    return this.decodeMessage(SuccinctZKVerifierFixture._encodedMessage());
  }

  function _fixtureWitness() internal pure returns (SuccinctZKVerifier.Witness memory) {
    return abi.decode(SuccinctZKVerifierFixture._witness(), (SuccinctZKVerifier.Witness));
  }

  function _encodeVerifierResults(
    SuccinctZKVerifier.Witness memory witness
  ) internal pure returns (bytes memory) {
    return abi.encodePacked(VERSION_TAG_V0_0_1, abi.encode(witness));
  }
}
