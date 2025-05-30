  // SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IFastTransferPool} from "../../../interfaces/IFastTransferPool.sol";
import {IRouterClient} from "../../../interfaces/IRouterClient.sol";
import {Client} from "../../../libraries/Client.sol";
import {Internal} from "../../../libraries/Internal.sol";
  import {TokenPool} from "../../../pools/TokenPool.sol";
import {FastTransferTokenPoolAbstract} from "../../../pools/FastTransferTokenPoolAbstract.sol";
import {FastTransferTokenPoolHelperSetup} from "./FastTransferTokenPoolHelperSetup.t.sol";

contract FastTransferTokenPoolHelper_ccipSendToken_Test is FastTransferTokenPoolHelperSetup {
  function setUp() public override {
    super.setUp();
  }

  function test_CcipSendToken_NativeFee() public {
    uint256 amount = 100 ether;
    uint256 balanceBefore = s_token.balanceOf(OWNER);
    bytes memory receiver = abi.encodePacked(address(0x5));
    bytes memory extraArgs = "";
    bytes32 mockMessageId = keccak256("mockMessageId");
    uint256 fastFeeExpected = amount * FAST_FEE_BPS / 10000;
    vm.mockCall(address(s_sourceRouter), abi.encodeWithSelector(IRouterClient.getFee.selector), abi.encode(1 ether));
    vm.mockCall(
      address(s_sourceRouter), abi.encodeWithSelector(IRouterClient.ccipSend.selector), abi.encode(mockMessageId)
    );
    Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
      receiver: destPoolAddress,
      data: abi.encode(
        FastTransferTokenPoolAbstract.MintMessage({
          sourceAmountToTransfer: amount,
          sourceDecimals: 18,
          fastTransferFee: fastFeeExpected,
          receiver: receiver
        })
      ),
      feeToken: address(0),
      tokenAmounts: new Client.EVMTokenAmount[](0),
      extraArgs: Client._argsToBytes(
        Client.GenericExtraArgsV2({gasLimit: SETTLEMENT_GAS_OVERHEAD, allowOutOfOrderExecution: true})
      )
    });
    vm.expectCall(
      address(s_sourceRouter), abi.encodeWithSelector(IRouterClient.ccipSend.selector, DEST_CHAIN_SELECTOR, message)
    );
    vm.expectEmit();
    emit IFastTransferPool.FastTransferRequested(mockMessageId, DEST_CHAIN_SELECTOR, amount, fastFeeExpected, receiver);
    bytes32 fillRequestId =
      s_tokenPool.ccipSendToken{value: 1 ether}(address(0), DEST_CHAIN_SELECTOR, amount, receiver, extraArgs);

    // Verify fillRequestId is non-zero
    assertEq(fillRequestId, mockMessageId);

    // Verify token balances
    assertEq(s_token.balanceOf(OWNER), balanceBefore - amount - fastFeeExpected);
    assertEq(s_token.balanceOf(address(s_tokenPool)), amount + fastFeeExpected);
  }

  function test_CcipSendToken_NativeFee_ToSVM() public {
    uint256 amount = 100 ether;
    uint256 balanceBefore = s_token.balanceOf(OWNER);
    bytes memory receiver = abi.encodePacked(address(0x5));
    bytes memory extraArgs = "";
    bytes32 mockMessageId = keccak256("mockMessageId");
    uint256 fastFeeExpected = amount * FAST_FEE_BPS / 10000;
    vm.mockCall(address(s_sourceRouter), abi.encodeWithSelector(IRouterClient.getFee.selector), abi.encode(1 ether));
    vm.mockCall(
      address(s_sourceRouter), abi.encodeWithSelector(IRouterClient.ccipSend.selector), abi.encode(mockMessageId)
    );
    Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
      receiver: destPoolAddress,
      data: abi.encode(
        FastTransferTokenPoolAbstract.MintMessage({
          sourceAmountToTransfer: amount,
          sourceDecimals: 18,
          fastTransferFee: fastFeeExpected,
          receiver: receiver
        })
      ),
      feeToken: address(0),
      tokenAmounts: new Client.EVMTokenAmount[](0),
      extraArgs: svmExtraArgsBytesEncoded
    });
    vm.expectCall(
      address(s_sourceRouter), abi.encodeWithSelector(IRouterClient.ccipSend.selector, SVM_CHAIN_SELECTOR, message)
    );
    vm.expectEmit();
    emit IFastTransferPool.FastTransferRequested(mockMessageId, SVM_CHAIN_SELECTOR, amount, fastFeeExpected, receiver);
    bytes32 fillRequestId =
              s_tokenPool.ccipSendToken{value: 1 ether}(address(0), SVM_CHAIN_SELECTOR, amount, receiver, extraArgs);

    // Verify fillRequestId is non-zero
    assertEq(fillRequestId, mockMessageId);

    // Verify token balances
    assertEq(s_token.balanceOf(OWNER), balanceBefore - amount - fastFeeExpected);
    assertEq(s_token.balanceOf(address(s_tokenPool)), amount + fastFeeExpected);
  }

  function test_CcipSendToken_EVM_WithZeroSettlementGas() public {
    uint64 testChainSelector = uint64(uint256(keccak256("EVM_ZERO_GAS")));
    uint256 amount = 100 ether;
    bytes memory receiver = abi.encodePacked(address(0x5));
    bytes32 mockMessageId = keccak256("mockMessageId");
    uint256 fastFeeExpected = amount * FAST_FEE_BPS / 10000;
    
    // Setup EVM chain config with zero settlement gas
    bytes memory customExtraArgs = abi.encode("customExtraArgs");
    address[] memory addFillers = new address[](1);
    addFillers[0] = s_filler;
    
    FastTransferTokenPoolAbstract.DestChainConfigUpdateArgs memory laneConfigArgs = FastTransferTokenPoolAbstract
      .DestChainConfigUpdateArgs({
      remoteChainSelector: testChainSelector,
      fastTransferBpsFee: FAST_FEE_BPS,
      fillerAllowlistEnabled: true,
      destinationPool: destPoolAddress,
      maxFillAmountPerRequest: MAX_FILL_AMOUNT_PER_REQUEST,
      settlementOverheadGas: 0, // Zero settlement gas
      chainFamilySelector: Internal.CHAIN_FAMILY_SELECTOR_EVM,
      evmToAnyMessageExtraArgsBytes: customExtraArgs,
      addFillers: addFillers,
      removeFillers: new address[](0)
    });
    s_tokenPool.updateDestChainConfig(laneConfigArgs);

    // Add chain update for rate limiter
    bytes[] memory remotePoolAddresses = new bytes[](1);
    remotePoolAddresses[0] = destPoolAddress;
    TokenPool.ChainUpdate[] memory chainUpdate = new TokenPool.ChainUpdate[](1);
    chainUpdate[0] = TokenPool.ChainUpdate({
      remoteChainSelector: testChainSelector,
      remotePoolAddresses: remotePoolAddresses,
      remoteTokenAddress: abi.encode(address(2)),
      outboundRateLimiterConfig: _getOutboundRateLimiterConfig(),
      inboundRateLimiterConfig: _getInboundRateLimiterConfig()
    });
    s_tokenPool.applyChainUpdates(new uint64[](0), chainUpdate);

    vm.mockCall(address(s_sourceRouter), abi.encodeWithSelector(IRouterClient.getFee.selector), abi.encode(1 ether));
    vm.mockCall(
      address(s_sourceRouter), abi.encodeWithSelector(IRouterClient.ccipSend.selector), abi.encode(mockMessageId)
    );
    
    Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
      receiver: destPoolAddress,
      data: abi.encode(
        FastTransferTokenPoolAbstract.MintMessage({
          sourceAmountToTransfer: amount,
          sourceDecimals: 18,
          fastTransferFee: fastFeeExpected,
          receiver: receiver
        })
      ),
      feeToken: address(0),
      tokenAmounts: new Client.EVMTokenAmount[](0),
      extraArgs: customExtraArgs // Should use custom extra args when settlementOverheadGas == 0
    });
    vm.expectCall(
      address(s_sourceRouter), abi.encodeWithSelector(IRouterClient.ccipSend.selector, testChainSelector, message)
    );
    vm.expectEmit();
    emit IFastTransferPool.FastTransferRequested(mockMessageId, testChainSelector, amount, fastFeeExpected, receiver);

    bytes32 fillRequestId =
      s_tokenPool.ccipSendToken{value: 1 ether}(address(0), testChainSelector, amount, receiver, "");

    assertEq(fillRequestId, mockMessageId);
  }

  function test_CcipSendToken_APTOS_WithSettlementGas() public {
    uint64 testChainSelector = uint64(uint256(keccak256("APTOS_WITH_GAS")));
    uint256 amount = 100 ether;
    bytes memory receiver = abi.encodePacked(address(0x5));
    bytes memory extraArgs = "";
    bytes32 mockMessageId = keccak256("mockMessageId");
    uint256 fastFeeExpected = amount * FAST_FEE_BPS / 10000;
    
    // Setup APTOS chain config with settlement gas
    address[] memory addFillers = new address[](1);
    addFillers[0] = s_filler;
    
    FastTransferTokenPoolAbstract.DestChainConfigUpdateArgs memory laneConfigArgs = FastTransferTokenPoolAbstract
      .DestChainConfigUpdateArgs({
      remoteChainSelector: testChainSelector,
      fastTransferBpsFee: FAST_FEE_BPS,
      fillerAllowlistEnabled: true,
      destinationPool: destPoolAddress,
      maxFillAmountPerRequest: MAX_FILL_AMOUNT_PER_REQUEST,
      settlementOverheadGas: SETTLEMENT_GAS_OVERHEAD,
      chainFamilySelector: Internal.CHAIN_FAMILY_SELECTOR_APTOS,
      evmToAnyMessageExtraArgsBytes: "",
      addFillers: addFillers,
      removeFillers: new address[](0)
    });
    s_tokenPool.updateDestChainConfig(laneConfigArgs);

    // Add chain update for rate limiter
    bytes[] memory remotePoolAddresses = new bytes[](1);
    remotePoolAddresses[0] = destPoolAddress;
    TokenPool.ChainUpdate[] memory chainUpdate = new TokenPool.ChainUpdate[](1);
    chainUpdate[0] = TokenPool.ChainUpdate({
      remoteChainSelector: testChainSelector,
      remotePoolAddresses: remotePoolAddresses,
      remoteTokenAddress: abi.encode(address(2)),
      outboundRateLimiterConfig: _getOutboundRateLimiterConfig(),
      inboundRateLimiterConfig: _getInboundRateLimiterConfig()
    });
    s_tokenPool.applyChainUpdates(new uint64[](0), chainUpdate);

    vm.mockCall(address(s_sourceRouter), abi.encodeWithSelector(IRouterClient.getFee.selector), abi.encode(1 ether));
    vm.mockCall(
      address(s_sourceRouter), abi.encodeWithSelector(IRouterClient.ccipSend.selector), abi.encode(mockMessageId)
    );
    
    Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
      receiver: destPoolAddress,
      data: abi.encode(
        FastTransferTokenPoolAbstract.MintMessage({
          sourceAmountToTransfer: amount,
          sourceDecimals: 18,
          fastTransferFee: fastFeeExpected,
          receiver: receiver
        })
      ),
      feeToken: address(0),
      tokenAmounts: new Client.EVMTokenAmount[](0),
      extraArgs: Client._argsToBytes(
        Client.GenericExtraArgsV2({gasLimit: SETTLEMENT_GAS_OVERHEAD, allowOutOfOrderExecution: true})
      )
    });
    vm.expectCall(
      address(s_sourceRouter), abi.encodeWithSelector(IRouterClient.ccipSend.selector, testChainSelector, message)
    );
    vm.expectEmit();
    emit IFastTransferPool.FastTransferRequested(mockMessageId, testChainSelector, amount, fastFeeExpected, receiver);

    bytes32 fillRequestId =
      s_tokenPool.ccipSendToken{value: 1 ether}(address(0), testChainSelector, amount, receiver, extraArgs);

    assertEq(fillRequestId, mockMessageId);
  }

  function test_CcipSendToken_APTOS_WithZeroSettlementGas() public {
    uint64 testChainSelector = uint64(uint256(keccak256("APTOS_ZERO_GAS")));
    uint256 amount = 100 ether;
    bytes memory receiver = abi.encodePacked(address(0x5));
    bytes32 mockMessageId = keccak256("mockMessageId");
    uint256 fastFeeExpected = amount * FAST_FEE_BPS / 10000;
    
    // Setup APTOS chain config with zero settlement gas
    bytes memory customExtraArgs = abi.encode("aptosCustomExtraArgs");
    address[] memory addFillers = new address[](1);
    addFillers[0] = s_filler;
    
    FastTransferTokenPoolAbstract.DestChainConfigUpdateArgs memory laneConfigArgs = FastTransferTokenPoolAbstract
      .DestChainConfigUpdateArgs({
      remoteChainSelector: testChainSelector,
      fastTransferBpsFee: FAST_FEE_BPS,
      fillerAllowlistEnabled: true,
      destinationPool: destPoolAddress,
      maxFillAmountPerRequest: MAX_FILL_AMOUNT_PER_REQUEST,
      settlementOverheadGas: 0, // Zero settlement gas
      chainFamilySelector: Internal.CHAIN_FAMILY_SELECTOR_APTOS,
      evmToAnyMessageExtraArgsBytes: customExtraArgs,
      addFillers: addFillers,
      removeFillers: new address[](0)
    });
    s_tokenPool.updateDestChainConfig(laneConfigArgs);

    // Add chain update for rate limiter
    bytes[] memory remotePoolAddresses = new bytes[](1);
    remotePoolAddresses[0] = destPoolAddress;
    TokenPool.ChainUpdate[] memory chainUpdate = new TokenPool.ChainUpdate[](1);
    chainUpdate[0] = TokenPool.ChainUpdate({
      remoteChainSelector: testChainSelector,
      remotePoolAddresses: remotePoolAddresses,
      remoteTokenAddress: abi.encode(address(2)),
      outboundRateLimiterConfig: _getOutboundRateLimiterConfig(),
      inboundRateLimiterConfig: _getInboundRateLimiterConfig()
    });
    s_tokenPool.applyChainUpdates(new uint64[](0), chainUpdate);

    vm.mockCall(address(s_sourceRouter), abi.encodeWithSelector(IRouterClient.getFee.selector), abi.encode(1 ether));
    vm.mockCall(
      address(s_sourceRouter), abi.encodeWithSelector(IRouterClient.ccipSend.selector), abi.encode(mockMessageId)
    );
    
    Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
      receiver: destPoolAddress,
      data: abi.encode(
        FastTransferTokenPoolAbstract.MintMessage({
          sourceAmountToTransfer: amount,
          sourceDecimals: 18,
          fastTransferFee: fastFeeExpected,
          receiver: receiver
        })
      ),
      feeToken: address(0),
      tokenAmounts: new Client.EVMTokenAmount[](0),
      extraArgs: customExtraArgs // Should use custom extra args when settlementOverheadGas == 0
    });
    vm.expectCall(
      address(s_sourceRouter), abi.encodeWithSelector(IRouterClient.ccipSend.selector, testChainSelector, message)
    );
    vm.expectEmit();
    emit IFastTransferPool.FastTransferRequested(mockMessageId, testChainSelector, amount, fastFeeExpected, receiver);

    bytes32 fillRequestId =
      s_tokenPool.ccipSendToken{value: 1 ether}(address(0), testChainSelector, amount, receiver, "");

    assertEq(fillRequestId, mockMessageId);
  }

  function test_CcipSendToken_SUI_WithSettlementGas() public {
    uint64 testChainSelector = uint64(uint256(keccak256("SUI_WITH_GAS")));
    uint256 amount = 100 ether;
    bytes memory receiver = abi.encodePacked(address(0x5));
    bytes memory extraArgs = "";
    bytes32 mockMessageId = keccak256("mockMessageId");
    uint256 fastFeeExpected = amount * FAST_FEE_BPS / 10000;
    
    // Setup SUI chain config with settlement gas
    address[] memory addFillers = new address[](1);
    addFillers[0] = s_filler;
    
    FastTransferTokenPoolAbstract.DestChainConfigUpdateArgs memory laneConfigArgs = FastTransferTokenPoolAbstract
      .DestChainConfigUpdateArgs({
      remoteChainSelector: testChainSelector,
      fastTransferBpsFee: FAST_FEE_BPS,
      fillerAllowlistEnabled: true,
      destinationPool: destPoolAddress,
      maxFillAmountPerRequest: MAX_FILL_AMOUNT_PER_REQUEST,
      settlementOverheadGas: SETTLEMENT_GAS_OVERHEAD,
      chainFamilySelector: Internal.CHAIN_FAMILY_SELECTOR_SUI,
      evmToAnyMessageExtraArgsBytes: "",
      addFillers: addFillers,
      removeFillers: new address[](0)
    });
    s_tokenPool.updateDestChainConfig(laneConfigArgs);

    // Add chain update for rate limiter
    bytes[] memory remotePoolAddresses = new bytes[](1);
    remotePoolAddresses[0] = destPoolAddress;
    TokenPool.ChainUpdate[] memory chainUpdate = new TokenPool.ChainUpdate[](1);
    chainUpdate[0] = TokenPool.ChainUpdate({
      remoteChainSelector: testChainSelector,
      remotePoolAddresses: remotePoolAddresses,
      remoteTokenAddress: abi.encode(address(2)),
      outboundRateLimiterConfig: _getOutboundRateLimiterConfig(),
      inboundRateLimiterConfig: _getInboundRateLimiterConfig()
    });
    s_tokenPool.applyChainUpdates(new uint64[](0), chainUpdate);

    vm.mockCall(address(s_sourceRouter), abi.encodeWithSelector(IRouterClient.getFee.selector), abi.encode(1 ether));
    vm.mockCall(
      address(s_sourceRouter), abi.encodeWithSelector(IRouterClient.ccipSend.selector), abi.encode(mockMessageId)
    );
    
    Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
      receiver: destPoolAddress,
      data: abi.encode(
        FastTransferTokenPoolAbstract.MintMessage({
          sourceAmountToTransfer: amount,
          sourceDecimals: 18,
          fastTransferFee: fastFeeExpected,
          receiver: receiver
        })
      ),
      feeToken: address(0),
      tokenAmounts: new Client.EVMTokenAmount[](0),
      extraArgs: Client._argsToBytes(
        Client.GenericExtraArgsV2({gasLimit: SETTLEMENT_GAS_OVERHEAD, allowOutOfOrderExecution: true})
      )
    });
    vm.expectCall(
      address(s_sourceRouter), abi.encodeWithSelector(IRouterClient.ccipSend.selector, testChainSelector, message)
    );
    vm.expectEmit();
    emit IFastTransferPool.FastTransferRequested(mockMessageId, testChainSelector, amount, fastFeeExpected, receiver);

    bytes32 fillRequestId =
      s_tokenPool.ccipSendToken{value: 1 ether}(address(0), testChainSelector, amount, receiver, extraArgs);

    assertEq(fillRequestId, mockMessageId);
  }

  function test_CcipSendToken_SUI_WithZeroSettlementGas() public {
    uint64 testChainSelector = uint64(uint256(keccak256("SUI_ZERO_GAS")));
    uint256 amount = 100 ether;
    bytes memory receiver = abi.encodePacked(address(0x5));
    bytes32 mockMessageId = keccak256("mockMessageId");
    uint256 fastFeeExpected = amount * FAST_FEE_BPS / 10000;
    
    // Setup SUI chain config with zero settlement gas
    bytes memory customExtraArgs = abi.encode("suiCustomExtraArgs");
    address[] memory addFillers = new address[](1);
    addFillers[0] = s_filler;
    
    FastTransferTokenPoolAbstract.DestChainConfigUpdateArgs memory laneConfigArgs = FastTransferTokenPoolAbstract
      .DestChainConfigUpdateArgs({
      remoteChainSelector: testChainSelector,
      fastTransferBpsFee: FAST_FEE_BPS,
      fillerAllowlistEnabled: true,
      destinationPool: destPoolAddress,
      maxFillAmountPerRequest: MAX_FILL_AMOUNT_PER_REQUEST,
      settlementOverheadGas: 0, // Zero settlement gas
      chainFamilySelector: Internal.CHAIN_FAMILY_SELECTOR_SUI,
      evmToAnyMessageExtraArgsBytes: customExtraArgs,
      addFillers: addFillers,
      removeFillers: new address[](0)
    });
    s_tokenPool.updateDestChainConfig(laneConfigArgs);

    // Add chain update for rate limiter
    bytes[] memory remotePoolAddresses = new bytes[](1);
    remotePoolAddresses[0] = destPoolAddress;
    TokenPool.ChainUpdate[] memory chainUpdate = new TokenPool.ChainUpdate[](1);
    chainUpdate[0] = TokenPool.ChainUpdate({
      remoteChainSelector: testChainSelector,
      remotePoolAddresses: remotePoolAddresses,
      remoteTokenAddress: abi.encode(address(2)),
      outboundRateLimiterConfig: _getOutboundRateLimiterConfig(),
      inboundRateLimiterConfig: _getInboundRateLimiterConfig()
    });
    s_tokenPool.applyChainUpdates(new uint64[](0), chainUpdate);

    vm.mockCall(address(s_sourceRouter), abi.encodeWithSelector(IRouterClient.getFee.selector), abi.encode(1 ether));
    vm.mockCall(
      address(s_sourceRouter), abi.encodeWithSelector(IRouterClient.ccipSend.selector), abi.encode(mockMessageId)
    );
    
    Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
      receiver: destPoolAddress,
      data: abi.encode(
        FastTransferTokenPoolAbstract.MintMessage({
          sourceAmountToTransfer: amount,
          sourceDecimals: 18,
          fastTransferFee: fastFeeExpected,
          receiver: receiver
        })
      ),
      feeToken: address(0),
      tokenAmounts: new Client.EVMTokenAmount[](0),
      extraArgs: customExtraArgs // Should use custom extra args when settlementOverheadGas == 0
    });
    vm.expectCall(
      address(s_sourceRouter), abi.encodeWithSelector(IRouterClient.ccipSend.selector, testChainSelector, message)
    );
    vm.expectEmit();
    emit IFastTransferPool.FastTransferRequested(mockMessageId, testChainSelector, amount, fastFeeExpected, receiver);

    bytes32 fillRequestId =
      s_tokenPool.ccipSendToken{value: 1 ether}(address(0), testChainSelector, amount, receiver, "");

    assertEq(fillRequestId, mockMessageId);
  }

  function test_CcipSendToken_RevertWhen_InvalidChainFamilySelector() public {
    uint64 testChainSelector = uint64(uint256(keccak256("INVALID_CHAIN")));
    uint256 amount = 100 ether;
    bytes memory receiver = abi.encodePacked(address(0x5));

    // Setup chain config with invalid chain family selector
    address[] memory addFillers = new address[](1);
    addFillers[0] = s_filler;
    
    FastTransferTokenPoolAbstract.DestChainConfigUpdateArgs memory laneConfigArgs = FastTransferTokenPoolAbstract
      .DestChainConfigUpdateArgs({
      remoteChainSelector: testChainSelector,
      fastTransferBpsFee: FAST_FEE_BPS,
      fillerAllowlistEnabled: true,
      destinationPool: destPoolAddress,
      maxFillAmountPerRequest: MAX_FILL_AMOUNT_PER_REQUEST,
      settlementOverheadGas: SETTLEMENT_GAS_OVERHEAD,
      chainFamilySelector: bytes4(0xdeadbeef), // Invalid chain family selector
      evmToAnyMessageExtraArgsBytes: "",
      addFillers: addFillers,
      removeFillers: new address[](0)
    });
    s_tokenPool.updateDestChainConfig(laneConfigArgs);

    // Add chain update for rate limiter
    bytes[] memory remotePoolAddresses = new bytes[](1);
    remotePoolAddresses[0] = destPoolAddress;
    TokenPool.ChainUpdate[] memory chainUpdate = new TokenPool.ChainUpdate[](1);
    chainUpdate[0] = TokenPool.ChainUpdate({
      remoteChainSelector: testChainSelector,
      remotePoolAddresses: remotePoolAddresses,
      remoteTokenAddress: abi.encode(address(2)),
      outboundRateLimiterConfig: _getOutboundRateLimiterConfig(),
      inboundRateLimiterConfig: _getInboundRateLimiterConfig()
    });
    s_tokenPool.applyChainUpdates(new uint64[](0), chainUpdate);

    vm.expectRevert(FastTransferTokenPoolAbstract.InvalidDestChainConfig.selector);
    s_tokenPool.ccipSendToken{value: 1 ether}(address(0), testChainSelector, amount, receiver, "");
  }
}
