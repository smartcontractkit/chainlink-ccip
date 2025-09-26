// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {CommitteeVerifier} from "../../../ccvs/CommitteeVerifier.sol";
import {BaseVerifier} from "../../../ccvs/components/BaseVerifier.sol";
import {MessageV1Codec} from "../../../libraries/MessageV1Codec.sol";

import {BaseVerifierSetup} from "../components/BaseVerifier/BaseVerifierSetup.t.sol";

contract CommitteeVerifierSetup is BaseVerifierSetup {
  CommitteeVerifier internal s_committeeVerifier;

  uint256 internal constant PRIVATE_KEY_0 = 0x60b919c82f0b4791a5b7c6a7275970ace1748759ebdaa8a5c3a4b2f5a8b1e8d1;
  address internal constant MOCK_SENDER = 0x3333333333333333333333333333333333333333;
  address internal constant MOCK_RECEIVER = 0x4444444444444444444444444444444444444444;

  function setUp() public virtual override {
    super.setUp();

    s_committeeVerifier = new CommitteeVerifier(_createBasicDynamicConfigArgs(), "testStorageLocation");

    address[] memory validSigner = new address[](1);
    validSigner[0] = vm.addr(PRIVATE_KEY_0);

    s_committeeVerifier.setSignatureConfig(validSigner, 1);

    BaseVerifier.DestChainConfigArgs[] memory destChainConfigs = new BaseVerifier.DestChainConfigArgs[](1);
    destChainConfigs[0] = BaseVerifier.DestChainConfigArgs({
      router: s_router,
      destChainSelector: DEST_CHAIN_SELECTOR,
      allowlistEnabled: false
    });

    s_committeeVerifier.applyDestChainConfigUpdates(destChainConfigs);
  }

  /// @notice Helper to create a minimal dynamic config.
  function _createBasicDynamicConfigArgs() internal view returns (CommitteeVerifier.DynamicConfig memory) {
    return CommitteeVerifier.DynamicConfig({
      feeQuoter: address(s_feeQuoter),
      feeAggregator: FEE_AGGREGATOR,
      allowlistAdmin: ALLOWLIST_ADMIN
    });
  }

  /// @notice Helper to create a dynamic config with custom addresses.
  function _createDynamicConfigArgs(
    address feeQuoter,
    address feeAggregator,
    address allowlistAdmin
  ) internal pure returns (CommitteeVerifier.DynamicConfig memory) {
    return CommitteeVerifier.DynamicConfig({
      feeQuoter: feeQuoter,
      feeAggregator: feeAggregator,
      allowlistAdmin: allowlistAdmin
    });
  }

  function _generateBasicMessageV1() internal pure returns (MessageV1Codec.MessageV1 memory, bytes32 messageId) {
    MessageV1Codec.MessageV1 memory message = MessageV1Codec.MessageV1({
      sourceChainSelector: 1,
      destChainSelector: 2,
      sequenceNumber: 1,
      onRampAddress: abi.encodePacked(address(0x1111111111111111111111111111111111111111)),
      offRampAddress: abi.encodePacked(address(0x2222222222222222222222222222222222222222)),
      finality: 100,
      sender: abi.encodePacked(address(0x3333333333333333333333333333333333333333)),
      receiver: abi.encodePacked(address(0x4444444444444444444444444444444444444444)),
      destBlob: "",
      tokenTransfer: new MessageV1Codec.TokenTransferV1[](0),
      data: ""
    });

    return (message, keccak256(MessageV1Codec._encodeMessageV1(message)));
  }

  /// @notice Helper function to mock fee quoter response for forwardToVerifier tests.
  /// @param isOutOfOrderExecution Whether the message should be processed out of order.
  /// @param destChainSelector The destination chain selector.
  /// @param feeToken The fee token address.
  /// @param feeTokenAmount The fee token amount.
  /// @param extraArgs The extra arguments.
  /// @param receiver The receiver address.
  function _mockFeeQuoterResponse(
    bool isOutOfOrderExecution,
    uint64 destChainSelector,
    address feeToken,
    uint256 feeTokenAmount,
    bytes memory extraArgs,
    address receiver
  ) internal {
    vm.mockCall(
      address(s_feeQuoter),
      abi.encodeWithSelector(
        s_feeQuoter.processMessageArgs.selector,
        destChainSelector,
        feeToken,
        feeTokenAmount,
        extraArgs,
        abi.encode(receiver)
      ),
      abi.encode(0, isOutOfOrderExecution, "", "")
    );
  }

  /// @notice Helper to create a signature with v=27 (required by SignatureQuorumValidator).
  /// @param privateKey The private key to sign with.
  /// @param hash The hash to sign.
  /// @return r The r component of the signature.
  /// @return s The s component of the signature (adjusted for v=27).
  function _signWithV27(uint256 privateKey, bytes32 hash) internal pure returns (bytes32 r, bytes32 s) {
    (uint8 v, bytes32 _r, bytes32 _s) = vm.sign(privateKey, hash);

    // SignatureQuorumValidator only supports sigs with v=27, so adjust if necessary.
    // Any valid ECDSA sig (r, s, v) can be "flipped" into (r, s*, v*) without knowing the private key.
    // https://github.com/kadenzipfel/smart-contract-vulnerabilities/blob/master/vulnerabilities/signature-malleability.md
    if (v == 28) {
      uint256 N = 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEBAAEDCE6AF48A03BBFD25E8CD0364141;
      _s = bytes32(N - uint256(_s));
    }

    return (_r, _s);
  }

  function _generateMessageHash(
    MessageV1Codec.MessageV1 memory message
  ) internal pure returns (bytes32) {
    return keccak256(MessageV1Codec._encodeMessageV1(message));
  }
}
