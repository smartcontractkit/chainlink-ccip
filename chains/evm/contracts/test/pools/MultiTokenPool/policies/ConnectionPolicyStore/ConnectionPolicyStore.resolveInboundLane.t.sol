// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IAdditionalPoolDataProvider} from "../../../../../interfaces/IAdditionalPoolDataProvider.sol";
import {IPoolConfigManager} from "../../../../../interfaces/IPoolConfigManager.sol";
import {IRemoteTokenAddressProvider} from "../../../../../interfaces/IRemoteTokenAddressProvider.sol";

import {FinalityCodec} from "../../../../../libraries/FinalityCodec.sol";
import {ConnectionPolicyStore} from "../../../../../pools/MultiTokenPool/policies/ConnectionPolicyStore.sol";
import {ConnectionPolicyStoreSetup} from "./ConnectionPolicyStoreSetup.t.sol";

contract ConnectionPolicyStore_resolveInboundLane is ConnectionPolicyStoreSetup {
  bytes4 internal constant WAIT_FOR_FINALITY = FinalityCodec.WAIT_FOR_FINALITY_FLAG;
  uint256 internal constant AMOUNT = 1e18;

  function test_resolveInboundLane() public view {
    IPoolConfigManager.ResolvedInboundTransfer memory transfer =
      _resolveInboundLane(s_token, s_remotePool, _getSourcePoolData(DEFAULT_TOKEN_DECIMALS, s_token, ""), AMOUNT);

    IPoolConfigManager.ResolvedInboundTransfer memory expectedTransfer;
    expectedTransfer.localAmount = AMOUNT;
    expectedTransfer.lockOrBurnFlavor = IPoolConfigManager.LockOrBurnFlavor.Burn;
    assertEq(abi.encode(expectedTransfer), abi.encode(transfer));
  }

  function test_resolveInboundLane_ConvertsDecimals() public view {
    // From fewer remote decimals, the amount scales up.
    IPoolConfigManager.ResolvedInboundTransfer memory transfer =
      _resolveInboundLane(s_token, s_remotePool, _getSourcePoolData(6, s_token, ""), 1e6);
    assertEq(1e18, transfer.localAmount);

    // From more remote decimals, the amount rounds down, so the pool never mints more than the remote chain sent.
    transfer = _resolveInboundLane(s_token, s_remotePool, _getSourcePoolData(20, s_token, ""), 1e20 + 99);
    assertEq(1e18, transfer.localAmount);
  }

  function test_resolveInboundLane_InlineConnectionRecord() public {
    address remoteToken = makeAddr("remoteToken");
    _setInlineConnectionConfig(_getInlineConnectionConfigArgs(s_token, DEST_CHAIN_SELECTOR, abi.encode(remoteToken)));

    // The source token must equal the inline remote token, so the provider of the policy is not asked.
    vm.expectCall(
      address(s_remoteTokenAddressProvider),
      abi.encodeWithSelector(IRemoteTokenAddressProvider.getRemoteTokenAddress.selector),
      0
    );
    IPoolConfigManager.ResolvedInboundTransfer memory transfer =
      _resolveInboundLane(s_token, s_remotePool, _getSourcePoolData(DEFAULT_TOKEN_DECIMALS, remoteToken, ""), AMOUNT);

    assertEq(AMOUNT, transfer.localAmount);
  }

  function test_resolveInboundLane_SourceTokenAddressMissing() public {
    _allowSourceTokenAddressMissingOnReceive();

    // A source without the source token does not depend on the remote token address provider.
    vm.expectCall(
      address(s_remoteTokenAddressProvider),
      abi.encodeWithSelector(IRemoteTokenAddressProvider.getRemoteTokenAddress.selector),
      0
    );
    // Pool data with the decimals only, as a legacy pool sends it.
    IPoolConfigManager.ResolvedInboundTransfer memory transfer =
      _resolveInboundLane(s_token, s_remotePool, abi.encode(uint256(6)), 1e6);
    assertEq(1e18, transfer.localAmount);

    // Empty pool data, as a pool older than 1.5.1 sends it, uses the local decimals.
    transfer = _resolveInboundLane(s_token, s_remotePool, "", AMOUNT);
    assertEq(AMOUNT, transfer.localAmount);
  }

  function test_resolveInboundLane_AdditionalPoolData() public {
    _setAdditionalPoolDataProvider(s_token);
    bytes memory additionalPoolData = abi.encode(keccak256("uid"));
    bytes4 requestedFinality = ALLOWED_FINALITY;
    bytes memory callData = abi.encodeCall(
      IAdditionalPoolDataProvider.isAdditionalPoolDataValid,
      (s_token, DEST_CHAIN_SELECTOR, requestedFinality, additionalPoolData)
    );
    vm.mockCall(s_additionalPoolDataProvider, callData, abi.encode(true));

    // The inbound side does not check the finality, so a token with default finality only still accepts it.
    vm.expectCall(s_additionalPoolDataProvider, callData);
    IPoolConfigManager.ResolvedInboundTransfer memory transfer = s_connectionStore.resolveInboundLane(
      s_token,
      DEST_CHAIN_SELECTOR,
      s_remotePool,
      _getSourcePoolData(DEFAULT_TOKEN_DECIMALS, s_token, additionalPoolData),
      AMOUNT,
      requestedFinality
    );

    assertEq(AMOUNT, transfer.localAmount);
  }

  function test_resolveInboundLane_AdditionalPoolDataProviderGetsEmptyData() public {
    _setAdditionalPoolDataProvider(s_token);
    _allowSourceTokenAddressMissingOnReceive();
    bytes memory callData = abi.encodeCall(
      IAdditionalPoolDataProvider.isAdditionalPoolDataValid, (s_token, DEST_CHAIN_SELECTOR, WAIT_FOR_FINALITY, "")
    );
    vm.mockCall(s_additionalPoolDataProvider, callData, abi.encode(true));

    // The provider decides whether a source without additional pool data is valid.
    vm.expectCall(s_additionalPoolDataProvider, callData);
    _resolveInboundLane(s_token, s_remotePool, abi.encode(uint256(DEFAULT_TOKEN_DECIMALS)), AMOUNT);
  }

  function test_resolveInboundLane_OutboundPauseIgnored() public {
    s_connectionStore.setTokenPause(_arrayOf(s_token), ConnectionPolicyStore.PauseDirection.Outbound, true);
    uint64[] memory remoteChainSelectors = _getRemoteChainSelectors(DEST_CHAIN_SELECTOR);
    s_connectionStore.setChainPause(remoteChainSelectors, ConnectionPolicyStore.PauseDirection.Outbound, true);

    IPoolConfigManager.ResolvedInboundTransfer memory transfer =
      _resolveInboundLane(s_token, s_remotePool, _getSourcePoolData(DEFAULT_TOKEN_DECIMALS, s_token, ""), AMOUNT);

    assertEq(AMOUNT, transfer.localAmount);
  }

  /// @notice The pool data that a lockOrBurn builds passes the checks of a releaseOrMint for the same token.
  function testFuzz_resolveInboundLane_PoolDataFromResolveOutboundLane(
    bytes memory additionalPoolData,
    bytes memory tokenArgs,
    uint256 amount
  ) public {
    _setAdditionalPoolDataProvider(s_token);
    vm.mockCall(
      s_additionalPoolDataProvider,
      abi.encodeCall(
        IAdditionalPoolDataProvider.getAdditionalPoolData, (s_token, DEST_CHAIN_SELECTOR, WAIT_FOR_FINALITY, tokenArgs)
      ),
      abi.encode(additionalPoolData)
    );
    vm.mockCall(
      s_additionalPoolDataProvider,
      abi.encodeCall(
        IAdditionalPoolDataProvider.isAdditionalPoolDataValid,
        (s_token, DEST_CHAIN_SELECTOR, WAIT_FOR_FINALITY, additionalPoolData)
      ),
      abi.encode(true)
    );
    bytes memory sourcePoolData =
      s_connectionStore.resolveOutboundLane(s_token, DEST_CHAIN_SELECTOR, WAIT_FOR_FINALITY, tokenArgs).destPoolData;

    IPoolConfigManager.ResolvedInboundTransfer memory transfer =
      _resolveInboundLane(s_token, s_remotePool, sourcePoolData, amount);

    assertEq(amount, transfer.localAmount);
  }

  function testFuzz_resolveInboundLane_ConvertsDecimals(
    uint8 remoteDecimals,
    uint256 amount
  ) public view {
    // Beyond a difference of 77 decimals, the conversion reverts.
    remoteDecimals = uint8(bound(remoteDecimals, 0, DEFAULT_TOKEN_DECIMALS + 77));

    uint256 expectedAmount;
    if (remoteDecimals > DEFAULT_TOKEN_DECIMALS) {
      expectedAmount = amount / 10 ** (remoteDecimals - DEFAULT_TOKEN_DECIMALS);
    } else {
      uint256 factor = 10 ** (DEFAULT_TOKEN_DECIMALS - remoteDecimals);
      amount = bound(amount, 0, type(uint256).max / factor);
      expectedAmount = amount * factor;
    }

    IPoolConfigManager.ResolvedInboundTransfer memory transfer =
      _resolveInboundLane(s_token, s_remotePool, _getSourcePoolData(remoteDecimals, s_token, ""), amount);

    assertEq(expectedAmount, transfer.localAmount);
  }

  // Reverts

  function test_resolveInboundLane_RevertWhen_GloballyPaused() public {
    s_connectionStore.setGlobalPause(true);
    bytes memory sourcePoolData = _getSourcePoolData(DEFAULT_TOKEN_DECIMALS, s_token, "");

    vm.expectRevert(ConnectionPolicyStore.GloballyPaused.selector);
    _resolveInboundLane(s_token, s_remotePool, sourcePoolData, AMOUNT);
  }

  function test_resolveInboundLane_RevertWhen_TokenIsPaused() public {
    s_connectionStore.setTokenPause(_arrayOf(s_token), ConnectionPolicyStore.PauseDirection.Inbound, true);
    bytes memory sourcePoolData = _getSourcePoolData(DEFAULT_TOKEN_DECIMALS, s_token, "");

    vm.expectRevert(
      abi.encodeWithSelector(
        ConnectionPolicyStore.TokenIsPaused.selector, s_token, ConnectionPolicyStore.PauseDirection.Inbound
      )
    );
    _resolveInboundLane(s_token, s_remotePool, sourcePoolData, AMOUNT);
  }

  function test_resolveInboundLane_RevertWhen_ChainIsPaused() public {
    uint64[] memory remoteChainSelectors = _getRemoteChainSelectors(DEST_CHAIN_SELECTOR);
    s_connectionStore.setChainPause(remoteChainSelectors, ConnectionPolicyStore.PauseDirection.Inbound, true);
    bytes memory sourcePoolData = _getSourcePoolData(DEFAULT_TOKEN_DECIMALS, s_token, "");

    vm.expectRevert(
      abi.encodeWithSelector(
        ConnectionPolicyStore.ChainIsPaused.selector, DEST_CHAIN_SELECTOR, ConnectionPolicyStore.PauseDirection.Inbound
      )
    );
    _resolveInboundLane(s_token, s_remotePool, sourcePoolData, AMOUNT);
  }

  function test_resolveInboundLane_RevertWhen_ConnectionRecordNotFound() public {
    bytes memory sourcePoolData = _getSourcePoolData(USDC_TOKEN_DECIMALS, s_otherToken, "");

    vm.expectRevert(
      abi.encodeWithSelector(ConnectionPolicyStore.ConnectionRecordNotFound.selector, s_otherToken, DEST_CHAIN_SELECTOR)
    );
    _resolveInboundLane(s_otherToken, s_remotePool, sourcePoolData, AMOUNT);
  }

  function test_resolveInboundLane_RevertWhen_RemotePoolRecordNotFound() public {
    _setTokenRemotePoolPolicy(s_token, bytes32(0));
    bytes memory sourcePoolData = _getSourcePoolData(DEFAULT_TOKEN_DECIMALS, s_token, "");

    vm.expectRevert(
      abi.encodeWithSelector(ConnectionPolicyStore.RemotePoolRecordNotFound.selector, s_token, DEST_CHAIN_SELECTOR)
    );
    _resolveInboundLane(s_token, s_remotePool, sourcePoolData, AMOUNT);
  }

  function test_resolveInboundLane_RevertWhen_InvalidSourcePoolAddress() public {
    bytes memory sourcePoolData = _getSourcePoolData(DEFAULT_TOKEN_DECIMALS, s_token, "");
    bytes memory unknownPool = abi.encode(makeAddr("unknownPool"));

    vm.expectRevert(
      abi.encodeWithSelector(
        ConnectionPolicyStore.InvalidSourcePoolAddress.selector, s_token, DEST_CHAIN_SELECTOR, unknownPool
      )
    );
    _resolveInboundLane(s_token, unknownPool, sourcePoolData, AMOUNT);

    // An inline record without remote pools works outbound only.
    ConnectionPolicyStore.RemotePoolConfig memory remotePoolConfig = _getRemotePoolConfig();
    remotePoolConfig.remotePoolAddresses = new bytes[](0);
    _setInlineRemotePoolConfig(_getInlineRemotePoolConfigArgs(s_token, DEST_CHAIN_SELECTOR, remotePoolConfig));

    vm.expectRevert(
      abi.encodeWithSelector(
        ConnectionPolicyStore.InvalidSourcePoolAddress.selector, s_token, DEST_CHAIN_SELECTOR, s_remotePool
      )
    );
    _resolveInboundLane(s_token, s_remotePool, sourcePoolData, AMOUNT);
  }

  function test_resolveInboundLane_RevertWhen_SourceTokenAddressMissing() public {
    bytes[] memory sourcePoolDatas = new bytes[](2);
    sourcePoolDatas[0] = abi.encode(uint256(DEFAULT_TOKEN_DECIMALS));
    sourcePoolDatas[1] = "";

    for (uint256 i = 0; i < sourcePoolDatas.length; ++i) {
      vm.expectRevert(
        abi.encodeWithSelector(ConnectionPolicyStore.SourceTokenAddressMissing.selector, s_token, DEST_CHAIN_SELECTOR)
      );
      _resolveInboundLane(s_token, s_remotePool, sourcePoolDatas[i], AMOUNT);
    }
  }

  function test_resolveInboundLane_RevertWhen_SourceTokenMismatch() public {
    // A trusted remote pool burned another token. A source token in the pool data is checked also while the record
    // accepts pool data without it.
    _allowSourceTokenAddressMissingOnReceive();
    address otherSourceToken = makeAddr("otherSourceToken");
    bytes memory sourcePoolData = _getSourcePoolData(DEFAULT_TOKEN_DECIMALS, otherSourceToken, "");

    vm.expectRevert(
      abi.encodeWithSelector(
        ConnectionPolicyStore.SourceTokenMismatch.selector,
        s_token,
        DEST_CHAIN_SELECTOR,
        abi.encode(s_token),
        abi.encode(otherSourceToken)
      )
    );
    _resolveInboundLane(s_token, s_remotePool, sourcePoolData, AMOUNT);
  }

  function test_resolveInboundLane_RevertWhen_InvalidSourcePoolData() public {
    bytes memory validPoolData = _getSourcePoolData(DEFAULT_TOKEN_DECIMALS, s_token, "");
    bytes[] memory sourcePoolDatas = new bytes[](5);
    // Decimals above 255, in both formats.
    sourcePoolDatas[0] = abi.encode(uint256(256));
    sourcePoolDatas[1] = abi.encode(uint256(256), abi.encode(s_token), "");
    // The long format with an empty source token.
    sourcePoolDatas[2] = abi.encode(uint256(DEFAULT_TOKEN_DECIMALS), "", "");
    // One byte more than abi.encode makes.
    sourcePoolDatas[3] = abi.encodePacked(validPoolData, bytes1(0));
    // Both offsets point to the source token. abi.decode accepts it, but abi.encode never makes it.
    sourcePoolDatas[4] =
      abi.encode(uint256(DEFAULT_TOKEN_DECIMALS), uint256(0x60), uint256(0x60), uint256(0x20), s_token);

    for (uint256 i = 0; i < sourcePoolDatas.length; ++i) {
      vm.expectRevert(abi.encodeWithSelector(ConnectionPolicyStore.InvalidSourcePoolData.selector, sourcePoolDatas[i]));
      _resolveInboundLane(s_token, s_remotePool, sourcePoolDatas[i], AMOUNT);
    }
  }

  function test_resolveInboundLane_RevertWhen_PoolDataDoesNotDecode() public {
    // Too short for the long format, so abi.decode reverts.
    bytes memory sourcePoolData = abi.encode(uint256(DEFAULT_TOKEN_DECIMALS), uint256(0x40));

    vm.expectRevert(bytes(""));
    _resolveInboundLane(s_token, s_remotePool, sourcePoolData, AMOUNT);
  }

  function test_resolveInboundLane_RevertWhen_InvalidRemoteToken() public {
    bytes memory sourcePoolData = _getSourcePoolData(DEFAULT_TOKEN_DECIMALS, s_token, "");

    // The provider gives an all-zero answer.
    vm.mockCall(
      address(s_remoteTokenAddressProvider),
      abi.encodeCall(IRemoteTokenAddressProvider.getRemoteTokenAddress, (s_token, DEST_CHAIN_SELECTOR)),
      abi.encode(abi.encode(address(0)))
    );

    vm.expectRevert(
      abi.encodeWithSelector(ConnectionPolicyStore.InvalidRemoteToken.selector, s_token, DEST_CHAIN_SELECTOR)
    );
    _resolveInboundLane(s_token, s_remotePool, sourcePoolData, AMOUNT);

    // The policy has no provider.
    _setRemoteTokenAddressProvider(CONNECTION_POLICY, address(0));

    vm.expectRevert(
      abi.encodeWithSelector(ConnectionPolicyStore.InvalidRemoteToken.selector, s_token, DEST_CHAIN_SELECTOR)
    );
    _resolveInboundLane(s_token, s_remotePool, sourcePoolData, AMOUNT);
  }

  function test_resolveInboundLane_RevertWhen_RemoteTokenAddressProviderReverts() public {
    bytes memory revertData = abi.encodeWithSignature("Error(string)", "unknown token");
    vm.mockCallRevert(
      address(s_remoteTokenAddressProvider),
      abi.encodeCall(IRemoteTokenAddressProvider.getRemoteTokenAddress, (s_token, DEST_CHAIN_SELECTOR)),
      revertData
    );
    bytes memory sourcePoolData = _getSourcePoolData(DEFAULT_TOKEN_DECIMALS, s_token, "");

    vm.expectRevert(revertData);
    _resolveInboundLane(s_token, s_remotePool, sourcePoolData, AMOUNT);
  }

  function test_resolveInboundLane_RevertWhen_RemoteTokenAddressProviderHasNoCode() public {
    _setRemoteTokenAddressProvider(CONNECTION_POLICY, makeAddr("providerWithoutCode"));
    bytes memory sourcePoolData = _getSourcePoolData(DEFAULT_TOKEN_DECIMALS, s_token, "");

    // The empty answer fails the decode.
    vm.expectRevert(bytes(""));
    _resolveInboundLane(s_token, s_remotePool, sourcePoolData, AMOUNT);
  }

  function test_resolveInboundLane_RevertWhen_InvalidAdditionalPoolData() public {
    _setAdditionalPoolDataProvider(s_token);
    bytes memory additionalPoolData = abi.encode(keccak256("otherUid"));
    vm.mockCall(
      s_additionalPoolDataProvider,
      abi.encodeWithSelector(IAdditionalPoolDataProvider.isAdditionalPoolDataValid.selector),
      abi.encode(false)
    );
    bytes memory sourcePoolData = _getSourcePoolData(DEFAULT_TOKEN_DECIMALS, s_token, additionalPoolData);

    vm.expectRevert(
      abi.encodeWithSelector(
        ConnectionPolicyStore.InvalidAdditionalPoolData.selector, s_token, DEST_CHAIN_SELECTOR, additionalPoolData
      )
    );
    _resolveInboundLane(s_token, s_remotePool, sourcePoolData, AMOUNT);
  }

  function test_resolveInboundLane_RevertWhen_OverflowDetected() public {
    // More remote decimals: a difference above 77.
    uint8 remoteDecimals = DEFAULT_TOKEN_DECIMALS + 78;
    bytes memory sourcePoolData = _getSourcePoolData(remoteDecimals, s_token, "");

    vm.expectRevert(
      abi.encodeWithSelector(
        ConnectionPolicyStore.OverflowDetected.selector, remoteDecimals, DEFAULT_TOKEN_DECIMALS, AMOUNT
      )
    );
    _resolveInboundLane(s_token, s_remotePool, sourcePoolData, AMOUNT);

    // Fewer remote decimals: the product overflows.
    uint256 amount = type(uint256).max / 1e18 + 1;
    sourcePoolData = _getSourcePoolData(0, s_token, "");

    vm.expectRevert(
      abi.encodeWithSelector(ConnectionPolicyStore.OverflowDetected.selector, 0, DEFAULT_TOKEN_DECIMALS, amount)
    );
    _resolveInboundLane(s_token, s_remotePool, sourcePoolData, amount);

    // Fewer remote decimals: a difference above 77.
    address largeDecimalsToken = address(_deployToken("LARGE", 78));
    _registerToken(largeDecimalsToken);
    _setInlineConnectionConfig(
      _getInlineConnectionConfigArgs(largeDecimalsToken, DEST_CHAIN_SELECTOR, abi.encode(largeDecimalsToken))
    );
    _setInlineRemotePoolConfig(
      _getInlineRemotePoolConfigArgs(largeDecimalsToken, DEST_CHAIN_SELECTOR, _getRemotePoolConfig())
    );
    sourcePoolData = _getSourcePoolData(0, largeDecimalsToken, "");

    vm.expectRevert(abi.encodeWithSelector(ConnectionPolicyStore.OverflowDetected.selector, 0, 78, 1));
    _resolveInboundLane(largeDecimalsToken, s_remotePool, sourcePoolData, 1);
  }

  function _resolveInboundLane(
    address localToken,
    bytes memory sourcePoolAddress,
    bytes memory sourcePoolData,
    uint256 sourceDenominatedAmount
  ) private view returns (IPoolConfigManager.ResolvedInboundTransfer memory) {
    return s_connectionStore.resolveInboundLane(
      localToken, DEST_CHAIN_SELECTOR, sourcePoolAddress, sourcePoolData, sourceDenominatedAmount, WAIT_FOR_FINALITY
    );
  }

  function _allowSourceTokenAddressMissingOnReceive() private {
    ConnectionPolicyStore.RemotePoolConfig memory remotePoolConfig = _getRemotePoolConfig();
    remotePoolConfig.allowSourceTokenAddressMissingOnReceive = true;
    _setInlineRemotePoolConfig(_getInlineRemotePoolConfigArgs(s_token, DEST_CHAIN_SELECTOR, remotePoolConfig));
  }
}
