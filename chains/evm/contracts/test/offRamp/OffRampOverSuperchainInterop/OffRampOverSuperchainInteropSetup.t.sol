// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ICrossL2Inbox} from "../../../vendor/optimism/interop-lib/v0/src/interfaces/ICrossL2Inbox.sol";
import {Identifier} from "../../../vendor/optimism/interop-lib/v0/src/interfaces/IIdentifier.sol";

import {OffRampOverSuperchainInteropHelper} from "../../helpers/OffRampOverSuperchainInteropHelper.sol";
import {MockCrossL2Inbox} from "../../mocks/MockCrossL2Inbox.sol";

import {IRMNRemote} from "../../../interfaces/IRMNRemote.sol";

import {Internal} from "../../../libraries/Internal.sol";
import {SuperchainInterop} from "../../../libraries/SuperchainInterop.sol";
import {MultiOCR3Base} from "../../../ocr/MultiOCR3Base.sol";
import {OffRamp} from "../../../offRamp/OffRamp.sol";
import {OffRampOverSuperchainInterop} from "../../../offRamp/OffRampOverSuperchainInterop.sol";

import {OffRampHelper} from "../../helpers/OffRampHelper.sol";
import {OffRampSetup} from "../OffRamp/OffRampSetup.t.sol";

contract OffRampOverSuperchainInteropSetup is OffRampSetup {
  uint256 internal constant CHAIN_ID_1 = 100;
  uint256 internal constant CHAIN_ID_2 = 200;
  uint256 internal constant CHAIN_ID_3 = 300;

  MockCrossL2Inbox internal s_mockCrossL2Inbox;
  OffRampOverSuperchainInterop internal s_offRampOverSuperchainInterop;
  OffRampOverSuperchainInteropHelper internal s_offRampHelper;

  function setUp() public virtual override {
    super.setUp();

    // Deploy MockCrossL2Inbox
    s_mockCrossL2Inbox = new MockCrossL2Inbox();

    // Deploy OffRampOverSuperchainInterop with initial config
    _deployOffRampOverSuperchainInterop();
  }

  function _deployOffRampOverSuperchainInterop() internal {
    OffRamp.StaticConfig memory staticConfig = OffRamp.StaticConfig({
      chainSelector: DEST_CHAIN_SELECTOR,
      gasForCallExactCheck: GAS_FOR_CALL_EXACT_CHECK,
      rmnRemote: s_mockRMNRemote,
      tokenAdminRegistry: address(s_tokenAdminRegistry),
      nonceManager: address(s_inboundNonceManager)
    });

    OffRamp.DynamicConfig memory dynamicConfig = _generateDynamicOffRampConfig(address(s_feeQuoter));

    OffRamp.SourceChainConfigArgs[] memory sourceChainConfigs = new OffRamp.SourceChainConfigArgs[](1);
    sourceChainConfigs[0] = OffRamp.SourceChainConfigArgs({
      router: s_destRouter,
      sourceChainSelector: SOURCE_CHAIN_SELECTOR_1,
      isEnabled: true,
      isRMNVerificationDisabled: false,
      onRamp: abi.encode(ON_RAMP_ADDRESS)
    });

    OffRampOverSuperchainInterop.ChainSelectorToChainIdConfigArgs[] memory chainIdConfigs =
      new OffRampOverSuperchainInterop.ChainSelectorToChainIdConfigArgs[](1);
    chainIdConfigs[0] = OffRampOverSuperchainInterop.ChainSelectorToChainIdConfigArgs({
      chainSelector: SOURCE_CHAIN_SELECTOR_1,
      chainId: CHAIN_ID_1
    });

    s_offRampOverSuperchainInterop = new OffRampOverSuperchainInterop(
      staticConfig, dynamicConfig, sourceChainConfigs, address(s_mockCrossL2Inbox), chainIdConfigs
    );

    // Also deploy helper version for targeted testing
    s_offRampHelper = new OffRampOverSuperchainInteropHelper(
      staticConfig, dynamicConfig, sourceChainConfigs, address(s_mockCrossL2Inbox), chainIdConfigs
    );

    // Cast to OffRamp to access setOCR3Configs
    s_offRamp = OffRampHelper(address(s_offRampOverSuperchainInterop));

    // Setup OCR config to allow execution by transmitters
    s_configDigestExec = _getBasicConfigDigest(F, s_emptySigners, s_validTransmitters);

    MultiOCR3Base.OCRConfigArgs[] memory ocrConfigs = new MultiOCR3Base.OCRConfigArgs[](1);
    ocrConfigs[0] = MultiOCR3Base.OCRConfigArgs({
      ocrPluginType: uint8(Internal.OCRPluginType.Execution),
      configDigest: s_configDigestExec,
      F: F,
      isSignatureVerificationEnabled: false,
      signers: s_emptySigners,
      transmitters: s_validTransmitters
    });

    OffRamp(address(s_offRampOverSuperchainInterop)).setOCR3Configs(ocrConfigs);
  }

  function _generateValidMessage(
    uint64 sourceChainSelector,
    uint64 sequenceNumber
  ) internal view returns (Internal.Any2EVMRampMessage memory) {
    return Internal.Any2EVMRampMessage({
      header: Internal.RampMessageHeader({
        messageId: keccak256(abi.encodePacked("messageId", sequenceNumber)),
        sourceChainSelector: sourceChainSelector,
        destChainSelector: DEST_CHAIN_SELECTOR,
        sequenceNumber: sequenceNumber,
        nonce: 1
      }),
      sender: abi.encode(address(0x1234567890123456789012345678901234567890)),
      data: abi.encode("test data"),
      receiver: address(s_receiver),
      gasLimit: GAS_LIMIT,
      tokenAmounts: new Internal.Any2EVMTokenTransfer[](0)
    });
  }

  function _getOffRampMetadataHash() internal view returns (bytes32) {
    return keccak256(
      abi.encode(
        Internal.ANY_2_EVM_MESSAGE_HASH,
        SOURCE_CHAIN_SELECTOR,
        DEST_CHAIN_SELECTOR,
        keccak256(abi.encode(address(s_offRampHelper)))
      )
    );
  }
}
