// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IAny2EVMMessageReceiverV2} from "../../../interfaces/IAny2EVMMessageReceiverV2.sol";
import {IPoolV2} from "../../../interfaces/IPoolV2.sol";
import {ITokenAdminRegistry} from "../../../interfaces/ITokenAdminRegistry.sol";

import {ERC165CheckerReverting} from "../../../libraries/ERC165CheckerReverting.sol";
import {MessageV1Codec} from "../../../libraries/MessageV1Codec.sol";
import {CCVAggregator} from "../../../offRamp/CCVAggregator.sol";
import {MockReceiverV2} from "../../mocks/MockReceiverV2.sol";
import {CCVAggregatorSetup} from "./CCVAggregatorSetup.t.sol";

import {IERC165} from "@openzeppelin/contracts@4.8.3/utils/introspection/IERC165.sol";

using ERC165CheckerReverting for address;

contract CCVAggregator_ensureCCVQuorumIsReached is CCVAggregatorSetup {
  address internal s_receiver;
  address internal s_requiredCCV;
  address internal s_optionalCCV1;
  address internal s_optionalCCV2;
  address internal s_laneMandatedCCV;
  address internal s_poolRequiredCCV;
  address internal s_destToken;
  address internal s_destTokenPool;

  uint16 internal constant FINALITY = 0;

  function setUp() public override {
    super.setUp();

    s_receiver = address(new MockReceiverV2(new address[](0), new address[](0), 0));

    s_requiredCCV = makeAddr("requiredCCV");
    s_optionalCCV1 = makeAddr("optionalCCV1");
    s_optionalCCV2 = makeAddr("optionalCCV2");
    s_laneMandatedCCV = makeAddr("laneMandatedCCV");
    s_poolRequiredCCV = makeAddr("poolRequiredCCV");
    s_destToken = makeAddr("destToken");
    s_destTokenPool = makeAddr("destTokenPool");

    // Configure source chain with lane mandated CCVs.
    CCVAggregator.SourceChainConfigArgs[] memory configs = new CCVAggregator.SourceChainConfigArgs[](1);
    configs[0] = CCVAggregator.SourceChainConfigArgs({
      router: s_sourceRouter,
      sourceChainSelector: SOURCE_CHAIN_SELECTOR,
      isEnabled: true,
      onRamp: abi.encode(makeAddr("onRamp")),
      defaultCCV: new address[](1),
      laneMandatedCCVs: new address[](1)
    });
    configs[0].laneMandatedCCVs[0] = s_laneMandatedCCV;
    configs[0].defaultCCV[0] = s_defaultCCV;

    s_agg.applySourceChainConfigUpdates(configs);

    // Mock the token admin registry to return our pool.
    vm.mockCall(
      s_tokenAdminRegistry,
      abi.encodeWithSelector(ITokenAdminRegistry.getPool.selector, s_destToken),
      abi.encode(s_destTokenPool)
    );

    // Mock the pool to support IERC165, but not IPoolV2 by default - tests can override this.
    vm.mockCall(
      s_destTokenPool,
      abi.encodeWithSelector(IERC165.supportsInterface.selector, type(IERC165).interfaceId),
      abi.encode(true)
    );
  }

  function _createTokenTransferWithPool() internal view returns (MessageV1Codec.TokenTransferV1[] memory) {
    MessageV1Codec.TokenTransferV1[] memory tokenTransfers = new MessageV1Codec.TokenTransferV1[](1);
    tokenTransfers[0] = MessageV1Codec.TokenTransferV1({
      amount: 100,
      sourcePoolAddress: abi.encodePacked(s_destTokenPool),
      sourceTokenAddress: abi.encodePacked(address(0x1111111111111111111111111111111111111111)),
      destTokenAddress: abi.encodePacked(s_destToken),
      extraData: ""
    });
    return tokenTransfers;
  }

  function test_ensureCCVQuorumIsReached_AllCCVsFound() public {
    address[] memory ccvs = new address[](5);
    ccvs[0] = s_requiredCCV;
    ccvs[1] = s_optionalCCV1;
    ccvs[2] = s_laneMandatedCCV;
    ccvs[3] = s_poolRequiredCCV;
    ccvs[4] = s_defaultCCV;

    // Mock receiver to return no required CCVs (so it falls back to defaults).
    vm.mockCall(
      s_receiver,
      abi.encodeWithSelector(IAny2EVMMessageReceiverV2.getCCVs.selector, SOURCE_CHAIN_SELECTOR),
      abi.encode(
        new address[](0), // no required CCVs - will fall back to defaults.
        new address[](0), // no optional CCVs.
        uint8(0) // no optional threshold.
      )
    );

    // Create token transfers that will trigger pool requirements.
    MessageV1Codec.TokenTransferV1[] memory tokenTransfers = _createTokenTransferWithPool();

    // Mock the pool to support IPoolV2 and return required CCVs.
    vm.mockCall(
      s_destTokenPool,
      abi.encodeWithSelector(IERC165.supportsInterface.selector, type(IPoolV2).interfaceId),
      abi.encode(true)
    );
    address[] memory poolRequiredCCVs = new address[](1);
    poolRequiredCCVs[0] = s_poolRequiredCCV;
    vm.mockCall(
      s_destTokenPool, abi.encodeWithSelector(IPoolV2.getRequiredInboundCCVs.selector), abi.encode(poolRequiredCCVs)
    );

    (address[] memory ccvsToQuery, uint256[] memory dataIndexes) =
      s_agg.ensureCCVQuorumIsReached(SOURCE_CHAIN_SELECTOR, s_receiver, tokenTransfers, FINALITY, ccvs);

    // Since we have 1 default, 1 lane mandated, and 1 pool required.
    assertEq(ccvsToQuery.length, 3);
    assertEq(ccvsToQuery[0], s_defaultCCV);
    assertEq(ccvsToQuery[1], s_poolRequiredCCV);
    assertEq(ccvsToQuery[2], s_laneMandatedCCV);

    assertEq(dataIndexes.length, 3, "right number of data indexes");
    assertEq(dataIndexes[0], 4);
    assertEq(dataIndexes[1], 3);
    assertEq(dataIndexes[2], 2);
  }

  function test_ensureCCVQuorumIsReached_OptionalIsAlsoRequired() public {
    address[] memory ccvs = new address[](2);
    ccvs[0] = s_optionalCCV1;
    ccvs[1] = s_laneMandatedCCV;

    MessageV1Codec.TokenTransferV1[] memory tokenTransfers = new MessageV1Codec.TokenTransferV1[](0);

    address[] memory receiverOptional = new address[](3);
    receiverOptional[0] = s_laneMandatedCCV;
    receiverOptional[1] = s_optionalCCV1;
    receiverOptional[2] = s_optionalCCV2;

    vm.mockCall(
      s_receiver,
      abi.encodeWithSelector(IAny2EVMMessageReceiverV2.getCCVs.selector, SOURCE_CHAIN_SELECTOR),
      abi.encode(new address[](0), receiverOptional, uint8(2))
    );

    (address[] memory ccvsToQuery, uint256[] memory dataIndexes) =
      s_agg.ensureCCVQuorumIsReached(SOURCE_CHAIN_SELECTOR, s_receiver, tokenTransfers, FINALITY, ccvs);

    assertEq(ccvsToQuery.length, 2);
    assertEq(ccvsToQuery[0], s_laneMandatedCCV);
    assertEq(ccvsToQuery[1], s_optionalCCV1);

    assertEq(dataIndexes.length, 2);
    assertEq(dataIndexes[0], 1);
    assertEq(dataIndexes[1], 0);
  }

  function test_ensureCCVQuorumIsReached_RevertWhen_RequiredCCVMissing_Receiver() public {
    address[] memory ccvs = new address[](2);
    ccvs[0] = s_optionalCCV1;
    ccvs[1] = s_laneMandatedCCV;

    MessageV1Codec.TokenTransferV1[] memory tokenTransfers = new MessageV1Codec.TokenTransferV1[](0);

    // Mock receiver to return a required CCV that's not in the provided CCVs.
    address[] memory receiverRequired = new address[](1);
    receiverRequired[0] = s_requiredCCV;
    vm.mockCall(
      s_receiver,
      abi.encodeWithSelector(IAny2EVMMessageReceiverV2.getCCVs.selector, SOURCE_CHAIN_SELECTOR),
      abi.encode(receiverRequired, new address[](0), uint8(0))
    );

    vm.expectRevert(abi.encodeWithSelector(CCVAggregator.RequiredCCVMissing.selector, s_requiredCCV));
    s_agg.ensureCCVQuorumIsReached(SOURCE_CHAIN_SELECTOR, s_receiver, tokenTransfers, FINALITY, ccvs);
  }

  function test_ensureCCVQuorumIsReached_RevertWhen_RequiredCCVMissing_Pool() public {
    address[] memory ccvs = new address[](3);
    ccvs[0] = s_requiredCCV;
    ccvs[1] = s_laneMandatedCCV;
    ccvs[2] = s_defaultCCV;

    // Create token transfers that will require pool CCVs but with missing CCV
    MessageV1Codec.TokenTransferV1[] memory tokenTransfers = _createTokenTransferWithPool();

    // Mock the pool to support IPoolV2 and return required CCVs (but CCV is missing from ccvs array).
    vm.mockCall(
      s_destTokenPool,
      abi.encodeWithSelector(IERC165.supportsInterface.selector, type(IPoolV2).interfaceId),
      abi.encode(true)
    );
    address[] memory poolRequiredCCVs = new address[](1);
    poolRequiredCCVs[0] = s_poolRequiredCCV;
    vm.mockCall(
      s_destTokenPool, abi.encodeWithSelector(IPoolV2.getRequiredInboundCCVs.selector), abi.encode(poolRequiredCCVs)
    );

    // Mock receiver to return required CCVs that are found.
    address[] memory receiverRequired = new address[](1);
    receiverRequired[0] = s_requiredCCV;
    vm.mockCall(
      s_receiver,
      abi.encodeWithSelector(IAny2EVMMessageReceiverV2.getCCVs.selector, SOURCE_CHAIN_SELECTOR),
      abi.encode(receiverRequired, new address[](0), uint8(0))
    );

    vm.expectRevert(abi.encodeWithSelector(CCVAggregator.RequiredCCVMissing.selector, s_poolRequiredCCV));
    s_agg.ensureCCVQuorumIsReached(SOURCE_CHAIN_SELECTOR, s_receiver, tokenTransfers, FINALITY, ccvs);
  }

  function test_ensureCCVQuorumIsReached_RevertWhen_RequiredCCVMissing_LaneMandated() public {
    address[] memory ccvs = new address[](3);
    ccvs[0] = s_requiredCCV;
    ccvs[1] = s_optionalCCV1;
    ccvs[2] = s_defaultCCV;
    MessageV1Codec.TokenTransferV1[] memory tokenTransfers = new MessageV1Codec.TokenTransferV1[](0);

    vm.expectRevert(abi.encodeWithSelector(CCVAggregator.RequiredCCVMissing.selector, s_laneMandatedCCV));
    s_agg.ensureCCVQuorumIsReached(SOURCE_CHAIN_SELECTOR, s_receiver, tokenTransfers, FINALITY, ccvs);
  }

  function test_ensureCCVQuorumIsReached_RevertWhen_OptionalCCVQuorumNotReached() public {
    address[] memory ccvs = new address[](4);
    ccvs[0] = s_requiredCCV;
    ccvs[1] = s_optionalCCV1;
    ccvs[2] = s_laneMandatedCCV;

    // Mock receiver to return optional CCVs with threshold 2.
    address[] memory receiverOptional = new address[](2);
    receiverOptional[0] = s_optionalCCV1;
    receiverOptional[1] = s_optionalCCV2;

    vm.mockCall(
      s_receiver,
      abi.encodeWithSelector(IAny2EVMMessageReceiverV2.getCCVs.selector, SOURCE_CHAIN_SELECTOR),
      abi.encode(new address[](0), receiverOptional, uint8(2))
    );

    vm.expectRevert(
      abi.encodeWithSelector(CCVAggregator.OptionalCCVQuorumNotReached.selector, receiverOptional.length, 1)
    );
    s_agg.ensureCCVQuorumIsReached(
      SOURCE_CHAIN_SELECTOR, s_receiver, new MessageV1Codec.TokenTransferV1[](0), FINALITY, ccvs
    );
  }

  function test_ensureCCVQuorumIsReached_Success_OptionalCCVsFound() public {
    address[] memory ccvs = new address[](2);
    ccvs[0] = s_optionalCCV1;
    ccvs[1] = s_laneMandatedCCV;

    MessageV1Codec.TokenTransferV1[] memory tokenTransfers = new MessageV1Codec.TokenTransferV1[](0);

    // Mock receiver to return optional CCVs with threshold 1.
    address[] memory receiverOptional = new address[](2);
    receiverOptional[0] = s_optionalCCV1;
    receiverOptional[1] = s_optionalCCV2;

    vm.mockCall(
      s_receiver,
      abi.encodeWithSelector(IAny2EVMMessageReceiverV2.getCCVs.selector, SOURCE_CHAIN_SELECTOR),
      abi.encode(new address[](0), receiverOptional, uint8(1))
    );

    (address[] memory ccvsToQuery, uint256[] memory dataIndexes) =
      s_agg.ensureCCVQuorumIsReached(SOURCE_CHAIN_SELECTOR, s_receiver, tokenTransfers, FINALITY, ccvs);

    assertEq(ccvsToQuery.length, 2);
    assertEq(ccvsToQuery[0], s_laneMandatedCCV);
    assertEq(ccvsToQuery[1], s_optionalCCV1);

    assertEq(dataIndexes.length, 2);
    assertEq(dataIndexes[0], 1);
    assertEq(dataIndexes[1], 0);
  }
}
