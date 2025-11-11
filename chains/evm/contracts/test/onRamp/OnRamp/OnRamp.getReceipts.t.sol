// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ICrossChainVerifierResolver} from "../../../interfaces/ICrossChainVerifierResolver.sol";
import {ICrossChainVerifierV1} from "../../../interfaces/ICrossChainVerifierV1.sol";
import {IFeeQuoter} from "../../../interfaces/IFeeQuoter.sol";
import {IPoolV2} from "../../../interfaces/IPoolV2.sol";

import {Client} from "../../../libraries/Client.sol";
import {ExtraArgsCodec} from "../../../libraries/ExtraArgsCodec.sol";
import {Pool} from "../../../libraries/Pool.sol";
import {OnRamp} from "../../../onRamp/OnRamp.sol";
import {OnRampTestHelper} from "../../helpers/OnRampTestHelper.sol";
import {OnRampSetup} from "./OnRampSetup.t.sol";

import {IERC165} from "@openzeppelin/contracts@5.0.2/utils/introspection/IERC165.sol";

contract OnRamp_getReceipts is OnRampSetup {
  OnRampTestHelper internal s_onRampHelper;
  address internal s_sourceToken;
  address internal s_pool;
  address internal s_verifier1;
  address internal s_verifier2;

  uint32 internal constant POOL_FEE_USD_CENTS = 100; // $1.00
  uint32 internal constant POOL_GAS_OVERHEAD = 50000;
  uint32 internal constant POOL_BYTES_OVERHEAD = 128;

  uint32 internal constant FEE_QUOTER_FEE_USD_CENTS = 50; // $0.50
  uint32 internal constant FEE_QUOTER_GAS_OVERHEAD = 30000;
  uint32 internal constant FEE_QUOTER_BYTES_OVERHEAD = 64;

  uint32 internal constant VERIFIER_FEE_USD_CENTS = 200; // $2.00
  uint32 internal constant VERIFIER_GAS = 100000;
  uint32 internal constant VERIFIER_BYTES = 256;

  function setUp() public override {
    super.setUp();

    s_onRampHelper = new OnRampTestHelper(
      OnRamp.StaticConfig({
        chainSelector: SOURCE_CHAIN_SELECTOR,
        rmnRemote: s_mockRMNRemote,
        tokenAdminRegistry: address(s_tokenAdminRegistry)
      }),
      OnRamp.DynamicConfig({
        feeQuoter: address(s_feeQuoter),
        reentrancyGuardEntered: false,
        feeAggregator: FEE_AGGREGATOR
      })
    );

    s_verifier1 = makeAddr("verifier1");
    s_verifier2 = makeAddr("verifier2");
    address[] memory defaultCCVs = new address[](2);
    defaultCCVs[0] = s_verifier1;
    defaultCCVs[1] = s_verifier2;

    OnRamp.DestChainConfigArgs[] memory args = new OnRamp.DestChainConfigArgs[](1);
    args[0] = OnRamp.DestChainConfigArgs({
      destChainSelector: DEST_CHAIN_SELECTOR,
      router: s_sourceRouter,
      addressBytesLength: EVM_ADDRESS_LENGTH,
      baseExecutionGasCost: BASE_EXEC_GAS_COST,
      laneMandatedCCVs: new address[](0),
      defaultCCVs: defaultCCVs,
      defaultExecutor: s_defaultExecutor,
      offRamp: abi.encodePacked(address(s_offRampOnRemoteChain))
    });
    s_onRampHelper.applyDestChainConfigUpdates(args);

    s_pool = makeAddr("sourcePool");
    s_sourceToken = makeAddr("sourceToken");

    vm.mockCall(s_pool, abi.encodeCall(IERC165(s_pool).supportsInterface, (Pool.CCIP_POOL_V1)), abi.encode(true));
    vm.mockCall(
      address(s_tokenAdminRegistry), abi.encodeCall(s_tokenAdminRegistry.getPool, (s_sourceToken)), abi.encode(s_pool)
    );
  }

  function _createMessage(
    uint256 tokenAmount
  ) internal returns (Client.EVM2AnyMessage memory) {
    Client.EVMTokenAmount[] memory tokenAmounts = new Client.EVMTokenAmount[](tokenAmount > 0 ? 1 : 0);
    if (tokenAmount > 0) {
      tokenAmounts[0] = Client.EVMTokenAmount({token: s_sourceToken, amount: tokenAmount});
    }

    return Client.EVM2AnyMessage({
      receiver: abi.encodePacked(makeAddr("receiver")),
      data: "test data",
      tokenAmounts: tokenAmounts,
      feeToken: s_sourceFeeToken,
      extraArgs: ""
    });
  }

  function _createExtraArgs(
    address[] memory ccvAddresses
  ) internal view returns (ExtraArgsCodec.GenericExtraArgsV3 memory) {
    address[] memory ccvs = new address[](ccvAddresses.length);
    bytes[] memory execArgs = new bytes[](ccvAddresses.length);
    for (uint256 i = 0; i < ccvAddresses.length; i++) {
      ccvs[i] = ccvAddresses[i];
    }

    return ExtraArgsCodec.GenericExtraArgsV3({
      ccvs: ccvs,
      ccvArgs: execArgs,
      blockConfirmations: 0,
      gasLimit: GAS_LIMIT,
      executor: s_defaultExecutor,
      executorArgs: "",
      tokenReceiver: "",
      tokenArgs: ""
    });
  }

  function _mockVerifier(address verifierAddress, address implAddress) internal {
    vm.mockCall(
      verifierAddress,
      abi.encodeCall(ICrossChainVerifierResolver.getOutboundImplementation, (DEST_CHAIN_SELECTOR, "")),
      abi.encode(implAddress)
    );
  }

  function _mockVerifierFee(
    address implAddress
  ) internal {
    vm.mockCall(
      implAddress,
      abi.encodeWithSelector(ICrossChainVerifierV1.getFee.selector),
      abi.encode(VERIFIER_FEE_USD_CENTS, VERIFIER_GAS, VERIFIER_BYTES)
    );
  }

  function test_getReceipts_WithTokens_PoolV2Fee_Success() public {
    // Mock pool to support IPoolV2
    vm.mockCall(s_pool, abi.encodeCall(IERC165.supportsInterface, (type(IPoolV2).interfaceId)), abi.encode(true));

    // Mock pool's getFee to return custom fees
    vm.mockCall(
      s_pool,
      abi.encodeWithSelector(IPoolV2.getFee.selector),
      abi.encode(POOL_FEE_USD_CENTS, POOL_GAS_OVERHEAD, POOL_BYTES_OVERHEAD, uint16(0))
    );

    // Mock verifiers
    address verifier1Impl = makeAddr("verifier1Impl");
    address verifier2Impl = makeAddr("verifier2Impl");
    _mockVerifier(s_verifier1, verifier1Impl);
    _mockVerifier(s_verifier2, verifier2Impl);
    _mockVerifierFee(verifier1Impl);
    _mockVerifierFee(verifier2Impl);

    Client.EVM2AnyMessage memory message = _createMessage(100 ether);
    address[] memory ccvs = new address[](2);
    ccvs[0] = s_verifier1;
    ccvs[1] = s_verifier2;
    ExtraArgsCodec.GenericExtraArgsV3 memory extraArgs = _createExtraArgs(ccvs);

    OnRamp.Receipt[] memory receipts = s_onRampHelper.getReceipts(DEST_CHAIN_SELECTOR, message, extraArgs);

    // Should have: 2 verifiers + 1 token + 1 executor = 4 receipts
    assertEq(receipts.length, 4, "Should have 4 receipts");

    // Check verifier receipts (first 2)
    assertEq(receipts[0].issuer, s_verifier1, "First receipt should be from verifier1");
    assertEq(receipts[0].feeTokenAmount, VERIFIER_FEE_USD_CENTS, "Verifier1 fee should match");
    assertEq(receipts[0].destGasLimit, VERIFIER_GAS, "Verifier1 gas should match");
    assertEq(receipts[0].destBytesOverhead, VERIFIER_BYTES, "Verifier1 bytes should match");

    assertEq(receipts[1].issuer, s_verifier2, "Second receipt should be from verifier2");
    assertEq(receipts[1].feeTokenAmount, VERIFIER_FEE_USD_CENTS, "Verifier2 fee should match");

    // Check executor receipt (last)
    assertEq(receipts[3].issuer, s_defaultExecutor, "Last receipt should be from executor");

    // Check token pool receipt (second to last)
    assertEq(receipts[2].issuer, s_sourceToken, "Second to last receipt should be from token");
    assertEq(receipts[2].feeTokenAmount, POOL_FEE_USD_CENTS, "Pool fee should match");
    assertEq(receipts[2].destGasLimit, POOL_GAS_OVERHEAD, "Pool gas overhead should match");
    assertEq(receipts[2].destBytesOverhead, POOL_BYTES_OVERHEAD, "Pool bytes overhead should match");
    assertEq(receipts[2].extraArgs, extraArgs.tokenArgs, "Pool extra args should match");
  }

  function test_getReceipts_WithTokens_FeeQuoterFallback_Success() public {
    // Mock pool to NOT support IPoolV2
    vm.mockCall(s_pool, abi.encodeCall(IERC165.supportsInterface, (type(IPoolV2).interfaceId)), abi.encode(false));

    // Mock FeeQuoter's getTokenTransferFee
    vm.mockCall(
      address(s_feeQuoter),
      abi.encodeCall(IFeeQuoter.getTokenTransferFee, (DEST_CHAIN_SELECTOR, s_sourceToken)),
      abi.encode(FEE_QUOTER_FEE_USD_CENTS, FEE_QUOTER_GAS_OVERHEAD, FEE_QUOTER_BYTES_OVERHEAD)
    );

    // Mock verifier
    address verifier1Impl = makeAddr("verifier1Impl");
    _mockVerifier(s_verifier1, verifier1Impl);
    _mockVerifierFee(verifier1Impl);

    Client.EVM2AnyMessage memory message = _createMessage(100 ether);
    address[] memory ccvs = new address[](1);
    ccvs[0] = s_verifier1;
    ExtraArgsCodec.GenericExtraArgsV3 memory extraArgs = _createExtraArgs(ccvs);

    OnRamp.Receipt[] memory receipts = s_onRampHelper.getReceipts(DEST_CHAIN_SELECTOR, message, extraArgs);

    // Should have: 1 verifier + 1 token + 1 executor = 3 receipts
    assertEq(receipts.length, 3, "Should have 3 receipts");

    // Check token pool receipt uses FeeQuoter values
    assertEq(receipts[1].issuer, s_sourceToken, "Token receipt should be second to last");
    assertEq(receipts[1].feeTokenAmount, FEE_QUOTER_FEE_USD_CENTS, "Should use FeeQuoter fee");
    assertEq(receipts[1].destGasLimit, FEE_QUOTER_GAS_OVERHEAD, "Should use FeeQuoter gas overhead");
    assertEq(receipts[1].destBytesOverhead, FEE_QUOTER_BYTES_OVERHEAD, "Should use FeeQuoter bytes overhead");
  }

  function test_getReceipts_WithTokens_PoolV2ReturnsZero_FallsBackToFeeQuoter() public {
    // Mock pool to support IPoolV2
    vm.mockCall(s_pool, abi.encodeCall(IERC165.supportsInterface, (type(IPoolV2).interfaceId)), abi.encode(true));

    // Mock pool's getFee to return zeros (indicating it wants to use FeeQuoter)
    vm.mockCall(s_pool, abi.encodeWithSelector(IPoolV2.getFee.selector), abi.encode(0, 0, 0, uint16(0)));

    // Mock FeeQuoter's getTokenTransferFee
    vm.mockCall(
      address(s_feeQuoter),
      abi.encodeCall(IFeeQuoter.getTokenTransferFee, (DEST_CHAIN_SELECTOR, s_sourceToken)),
      abi.encode(FEE_QUOTER_FEE_USD_CENTS, FEE_QUOTER_GAS_OVERHEAD, FEE_QUOTER_BYTES_OVERHEAD)
    );

    // Mock verifier
    address verifier1Impl = makeAddr("verifier1Impl");
    _mockVerifier(s_verifier1, verifier1Impl);
    _mockVerifierFee(verifier1Impl);

    Client.EVM2AnyMessage memory message = _createMessage(100 ether);
    address[] memory ccvs = new address[](1);
    ccvs[0] = s_verifier1;
    ExtraArgsCodec.GenericExtraArgsV3 memory extraArgs = _createExtraArgs(ccvs);

    OnRamp.Receipt[] memory receipts = s_onRampHelper.getReceipts(DEST_CHAIN_SELECTOR, message, extraArgs);

    // Check token receipt falls back to FeeQuoter values
    assertEq(receipts[1].issuer, s_sourceToken, "Token receipt should be present");
    assertEq(receipts[1].feeTokenAmount, FEE_QUOTER_FEE_USD_CENTS, "Should fall back to FeeQuoter fee");
    assertEq(receipts[1].destGasLimit, FEE_QUOTER_GAS_OVERHEAD, "Should fall back to FeeQuoter gas");
    assertEq(receipts[1].destBytesOverhead, FEE_QUOTER_BYTES_OVERHEAD, "Should fall back to FeeQuoter bytes");
  }

  function test_getReceipts_NoTokens_Success() public {
    // Mock verifiers
    address verifier1Impl = makeAddr("verifier1Impl");
    address verifier2Impl = makeAddr("verifier2Impl");
    _mockVerifier(s_verifier1, verifier1Impl);
    _mockVerifier(s_verifier2, verifier2Impl);
    _mockVerifierFee(verifier1Impl);
    _mockVerifierFee(verifier2Impl);

    Client.EVM2AnyMessage memory message = _createMessage(0); // No tokens
    address[] memory ccvs = new address[](2);
    ccvs[0] = s_verifier1;
    ccvs[1] = s_verifier2;
    ExtraArgsCodec.GenericExtraArgsV3 memory extraArgs = _createExtraArgs(ccvs);

    OnRamp.Receipt[] memory receipts = s_onRampHelper.getReceipts(DEST_CHAIN_SELECTOR, message, extraArgs);

    // Should have: 2 verifiers + 1 executor = 3 receipts (no token receipt)
    assertEq(receipts.length, 3, "Should have 3 receipts without tokens");

    // Verify all are verifiers and executor
    assertEq(receipts[0].issuer, s_verifier1, "First should be verifier1");
    assertEq(receipts[1].issuer, s_verifier2, "Second should be verifier2");
    assertEq(receipts[2].issuer, s_defaultExecutor, "Last should be executor");
  }

  function test_getReceipts_NoVerifiers_WithTokens_Success() public {
    // Mock pool to support IPoolV2
    vm.mockCall(s_pool, abi.encodeCall(IERC165.supportsInterface, (type(IPoolV2).interfaceId)), abi.encode(true));
    vm.mockCall(
      s_pool,
      abi.encodeWithSelector(IPoolV2.getFee.selector),
      abi.encode(POOL_FEE_USD_CENTS, POOL_GAS_OVERHEAD, POOL_BYTES_OVERHEAD, uint16(0))
    );

    Client.EVM2AnyMessage memory message = _createMessage(100 ether);
    address[] memory ccvs = new address[](0); // No verifiers
    ExtraArgsCodec.GenericExtraArgsV3 memory extraArgs = _createExtraArgs(ccvs);

    OnRamp.Receipt[] memory receipts = s_onRampHelper.getReceipts(DEST_CHAIN_SELECTOR, message, extraArgs);

    // Should have: 0 verifiers + 1 token + 1 executor = 2 receipts
    assertEq(receipts.length, 2, "Should have 2 receipts");

    // Check receipts order
    assertEq(receipts[0].issuer, s_sourceToken, "First should be token");
    assertEq(receipts[0].feeTokenAmount, POOL_FEE_USD_CENTS, "Token fee should match");
    assertEq(receipts[1].issuer, s_defaultExecutor, "Last should be executor");
  }

  function test_getReceipts_MultipleVerifiers_WithTokens_OrderIsCorrect() public {
    // Mock pool
    vm.mockCall(s_pool, abi.encodeCall(IERC165.supportsInterface, (type(IPoolV2).interfaceId)), abi.encode(true));
    vm.mockCall(
      s_pool,
      abi.encodeWithSelector(IPoolV2.getFee.selector),
      abi.encode(POOL_FEE_USD_CENTS, POOL_GAS_OVERHEAD, POOL_BYTES_OVERHEAD, uint16(0))
    );

    // Mock 3 verifiers
    address verifier3 = makeAddr("verifier3");
    address verifier1Impl = makeAddr("verifier1Impl");
    address verifier2Impl = makeAddr("verifier2Impl");
    address verifier3Impl = makeAddr("verifier3Impl");

    _mockVerifier(s_verifier1, verifier1Impl);
    _mockVerifier(s_verifier2, verifier2Impl);
    _mockVerifier(verifier3, verifier3Impl);
    _mockVerifierFee(verifier1Impl);
    _mockVerifierFee(verifier2Impl);
    _mockVerifierFee(verifier3Impl);

    Client.EVM2AnyMessage memory message = _createMessage(100 ether);
    address[] memory ccvs = new address[](3);
    ccvs[0] = s_verifier1;
    ccvs[1] = s_verifier2;
    ccvs[2] = verifier3;
    ExtraArgsCodec.GenericExtraArgsV3 memory extraArgs = _createExtraArgs(ccvs);

    OnRamp.Receipt[] memory receipts = s_onRampHelper.getReceipts(DEST_CHAIN_SELECTOR, message, extraArgs);

    // Should have: 3 verifiers + 1 token + 1 executor = 5 receipts
    assertEq(receipts.length, 5, "Should have 5 receipts");

    // Verify order: verifiers (0-2), token (3), executor (4)
    assertEq(receipts[0].issuer, s_verifier1, "Receipt 0: verifier1");
    assertEq(receipts[1].issuer, s_verifier2, "Receipt 1: verifier2");
    assertEq(receipts[2].issuer, verifier3, "Receipt 2: verifier3");
    assertEq(receipts[3].issuer, s_sourceToken, "Receipt 3: token (second to last)");
    assertEq(receipts[4].issuer, s_defaultExecutor, "Receipt 4: executor (last)");
  }

  function test_getReceipts_TokenArgsPassedToPool() public {
    bytes memory customTokenArgs = abi.encode("custom token args");

    // Mock pool to support IPoolV2
    vm.mockCall(s_pool, abi.encodeCall(IERC165.supportsInterface, (type(IPoolV2).interfaceId)), abi.encode(true));

    // Expect pool's getFee to be called with correct tokenArgs
    vm.mockCall(
      s_pool,
      abi.encodeWithSelector(IPoolV2.getFee.selector),
      abi.encode(POOL_FEE_USD_CENTS, POOL_GAS_OVERHEAD, POOL_BYTES_OVERHEAD, uint16(0))
    );

    // Mock verifier
    address verifier1Impl = makeAddr("verifier1Impl");
    _mockVerifier(s_verifier1, verifier1Impl);
    _mockVerifierFee(verifier1Impl);

    Client.EVM2AnyMessage memory message = _createMessage(100 ether);
    address[] memory ccvs = new address[](1);
    ccvs[0] = s_verifier1;
    ExtraArgsCodec.GenericExtraArgsV3 memory extraArgs = _createExtraArgs(ccvs);
    extraArgs.tokenArgs = customTokenArgs;

    OnRamp.Receipt[] memory receipts = s_onRampHelper.getReceipts(DEST_CHAIN_SELECTOR, message, extraArgs);

    // Verify token receipt has correct tokenArgs
    assertEq(receipts[1].extraArgs, customTokenArgs, "Token receipt should have custom token args");
  }
}
