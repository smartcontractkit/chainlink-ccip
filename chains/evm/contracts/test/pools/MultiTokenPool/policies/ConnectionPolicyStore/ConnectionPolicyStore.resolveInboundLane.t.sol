// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IAdditionalPoolDataProvider} from "../../../../../interfaces/IAdditionalPoolDataProvider.sol";
import {IPoolConfigManager} from "../../../../../interfaces/IPoolConfigManager.sol";
import {IRemoteTokenAddressProvider} from "../../../../../interfaces/IRemoteTokenAddressProvider.sol";

import {FinalityCodec} from "../../../../../libraries/FinalityCodec.sol";
import {PolicyConfig} from "../../../../../libraries/PolicyConfig.sol";
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

  function test_resolveInboundLane_TokenProfile() public {
    address burnAddress = makeAddr("burnAddress");
    PolicyConfig.TokenProfileArgs[] memory profiles = _getTokenProfileArgs(s_token);
    profiles[0].lockOrBurnFlavor = IPoolConfigManager.LockOrBurnFlavor.BurnToAddress;
    profiles[0].lockOrBurnAddress = burnAddress;
    s_connectionStore.applyTokenProfileUpdates(new address[](0), profiles);

    IPoolConfigManager.ResolvedInboundTransfer memory transfer =
      _resolveInboundLane(s_token, s_remotePool, _getSourcePoolData(DEFAULT_TOKEN_DECIMALS, s_token, ""), AMOUNT);

    assertEq(uint8(IPoolConfigManager.LockOrBurnFlavor.BurnToAddress), uint8(transfer.lockOrBurnFlavor));
    assertEq(burnAddress, transfer.lockOrBurnAddress);
  }

  function test_resolveInboundLane_ConvertsDecimals() public view {
    IPoolConfigManager.ResolvedInboundTransfer memory transfer =
      _resolveInboundLane(s_token, s_remotePool, _getSourcePoolData(6, s_token, ""), 1e6);
    assertEq(1e18, transfer.localAmount);

    // Converting to fewer decimals must round down.
    transfer = _resolveInboundLane(s_token, s_remotePool, _getSourcePoolData(20, s_token, ""), 1e20 + 99);
    assertEq(1e18, transfer.localAmount);
  }

  function test_resolveInboundLane_InlineConnectionRecord() public {
    address remoteToken = makeAddr("remoteToken");
    _setInlineConnectionConfig(_getInlineConnectionConfigArgs(s_token, DEST_CHAIN_SELECTOR, abi.encode(remoteToken)));

    // The inline remote token takes precedence without calling the policy's provider.
    vm.expectCall(
      address(s_remoteTokenAddressProvider),
      abi.encodeWithSelector(IRemoteTokenAddressProvider.getRemoteTokenAddress.selector),
      0
    );
    IPoolConfigManager.ResolvedInboundTransfer memory transfer =
      _resolveInboundLane(s_token, s_remotePool, _getSourcePoolData(DEFAULT_TOKEN_DECIMALS, remoteToken, ""), AMOUNT);

    assertEq(AMOUNT, transfer.localAmount);
  }

  function test_resolveInboundLane_LegacyPoolData() public {
    _allowLegacyPoolDataOnReceive();

    // Legacy data has no source token to compare, so the remote token provider is not needed.
    vm.expectCall(
      address(s_remoteTokenAddressProvider),
      abi.encodeWithSelector(IRemoteTokenAddressProvider.getRemoteTokenAddress.selector),
      0
    );
    // Legacy pool data with token decimals still requires amount conversion.
    IPoolConfigManager.ResolvedInboundTransfer memory transfer =
      _resolveInboundLane(s_token, s_remotePool, abi.encode(uint256(6)), 1e6);
    assertEq(1e18, transfer.localAmount);

    // Older pools send no decimals, so the local decimals are used.
    transfer = _resolveInboundLane(s_token, s_remotePool, "", AMOUNT);
    assertEq(AMOUNT, transfer.localAmount);
  }

  function test_resolveInboundLane_TrustPolicyRemotePools() public {
    // Accept legacy data from the inline pool while also trusting the policy pool.
    _setInlineRemotePoolConfigTrustingPolicy(s_token, s_otherRemotePool, true);
    bytes memory legacyPoolData = abi.encode(uint256(DEFAULT_TOKEN_DECIMALS));

    _resolveInboundLane(s_token, s_remotePool, _getSourcePoolData(DEFAULT_TOKEN_DECIMALS, s_token, ""), AMOUNT);
    _resolveInboundLane(s_token, s_otherRemotePool, legacyPoolData, AMOUNT);

    // Trusting the policy pool does not allow it to send legacy data.
    vm.expectRevert(
      abi.encodeWithSelector(ConnectionPolicyStore.LegacyPoolDataNotAllowed.selector, s_token, DEST_CHAIN_SELECTOR)
    );
    _resolveInboundLane(s_token, s_remotePool, legacyPoolData, AMOUNT);
  }

  function test_resolveInboundLane_TrustPolicyRemotePools_InlineFlagWins() public {
    bytes memory legacyPoolData = abi.encode(uint256(DEFAULT_TOKEN_DECIMALS));
    // Both records list this pool; only the inline record allows legacy data.
    _setInlineRemotePoolConfigTrustingPolicy(s_token, s_remotePool, true);
    _resolveInboundLane(s_token, s_remotePool, legacyPoolData, AMOUNT);

    // The inline legacy permission must not affect another token assigned to the policy.
    _setTokenRemotePoolPolicy(s_otherToken, REMOTE_POOL_POLICY);
    _setInlineConnectionConfig(
      _getInlineConnectionConfigArgs(s_otherToken, DEST_CHAIN_SELECTOR, abi.encode(s_otherToken))
    );
    vm.expectRevert(
      abi.encodeWithSelector(ConnectionPolicyStore.LegacyPoolDataNotAllowed.selector, s_otherToken, DEST_CHAIN_SELECTOR)
    );
    _resolveInboundLane(s_otherToken, s_remotePool, legacyPoolData, AMOUNT);
  }

  function test_resolveInboundLane_TrustPolicyRemotePools_NoInlineRemotePools() public {
    ConnectionPolicyStore.RemotePoolConfig memory remotePoolConfig = _getRemotePoolConfig();
    remotePoolConfig.remotePools = new IPoolConfigManager.RemotePool[](0);
    ConnectionPolicyStore.InlineRemotePoolConfigArgs[] memory inlineConfigs =
      _getInlineRemotePoolConfigArgs(s_token, DEST_CHAIN_SELECTOR, remotePoolConfig);
    inlineConfigs[0].trustPolicyRemotePools = true;
    _setInlineRemotePoolConfig(inlineConfigs);

    IPoolConfigManager.ResolvedInboundTransfer memory transfer =
      _resolveInboundLane(s_token, s_remotePool, _getSourcePoolData(DEFAULT_TOKEN_DECIMALS, s_token, ""), AMOUNT);
    assertEq(AMOUNT, transfer.localAmount);
  }

  function test_resolveInboundLane_TrustPolicyRemotePools_PolicyRecordOfTheChain() public {
    // Configure different pools for the two chains to detect a lookup on the wrong chain.
    _setRemotePoolPolicyConfig(
      _getRemotePoolPolicyConfigArgs(
        REMOTE_POOL_POLICY, SOURCE_CHAIN_SELECTOR, _getRemotePoolAddresses(s_otherRemotePool)
      )
    );
    bytes memory legacyPool = abi.encode(makeAddr("legacyPool"));
    _setInlineRemotePoolConfigTrustingPolicy(s_token, legacyPool, true);
    bytes memory sourcePoolData = _getSourcePoolData(DEFAULT_TOKEN_DECIMALS, s_token, "");

    // Only the policy pools for this chain are trusted.
    _resolveInboundLane(s_token, s_remotePool, sourcePoolData, AMOUNT);
    vm.expectRevert(
      abi.encodeWithSelector(
        ConnectionPolicyStore.InvalidSourcePoolAddress.selector, s_token, DEST_CHAIN_SELECTOR, s_otherRemotePool
      )
    );
    _resolveInboundLane(s_token, s_otherRemotePool, sourcePoolData, AMOUNT);

    // Removing the chain's policy record leaves only the inline pool trusted.
    s_connectionStore.applyRemotePoolPolicyConfigUpdates(
      _getPolicyRecordKeys(REMOTE_POOL_POLICY, DEST_CHAIN_SELECTOR),
      new bytes32[](0),
      new ConnectionPolicyStore.RemotePoolPolicyConfigArgs[](0)
    );
    vm.expectRevert(
      abi.encodeWithSelector(
        ConnectionPolicyStore.InvalidSourcePoolAddress.selector, s_token, DEST_CHAIN_SELECTOR, s_remotePool
      )
    );
    _resolveInboundLane(s_token, s_remotePool, sourcePoolData, AMOUNT);
    _resolveInboundLane(s_token, legacyPool, abi.encode(uint256(DEFAULT_TOKEN_DECIMALS)), AMOUNT);
  }

  function test_resolveInboundLane_TrustPolicyRemotePools_InlineRemotePoolStays() public {
    _setInlineRemotePoolConfigTrustingPolicy(s_token, s_remotePool, false);
    _setRemotePoolPolicyConfig(
      _getRemotePoolPolicyConfigArgs(
        REMOTE_POOL_POLICY, DEST_CHAIN_SELECTOR, _getRemotePoolAddresses(s_otherRemotePool)
      )
    );

    // Removing a pool from the policy does not revoke trust granted by an inline record.
    _resolveInboundLane(s_token, s_remotePool, _getSourcePoolData(DEFAULT_TOKEN_DECIMALS, s_token, ""), AMOUNT);
    _resolveInboundLane(s_token, s_otherRemotePool, _getSourcePoolData(DEFAULT_TOKEN_DECIMALS, s_token, ""), AMOUNT);
    _setTokenRemotePoolPolicy(s_otherToken, REMOTE_POOL_POLICY);
    _setInlineConnectionConfig(
      _getInlineConnectionConfigArgs(s_otherToken, DEST_CHAIN_SELECTOR, abi.encode(s_otherToken))
    );
    vm.expectRevert(
      abi.encodeWithSelector(
        ConnectionPolicyStore.InvalidSourcePoolAddress.selector, s_otherToken, DEST_CHAIN_SELECTOR, s_remotePool
      )
    );
    _resolveInboundLane(s_otherToken, s_remotePool, _getSourcePoolData(USDC_TOKEN_DECIMALS, s_otherToken, ""), AMOUNT);
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

    // The token's outbound finality restriction does not apply to inbound transfers.
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
    _allowLegacyPoolDataOnReceive();
    bytes memory callData = abi.encodeCall(
      IAdditionalPoolDataProvider.isAdditionalPoolDataValid, (s_token, DEST_CHAIN_SELECTOR, WAIT_FOR_FINALITY, "")
    );
    vm.mockCall(s_additionalPoolDataProvider, callData, abi.encode(true));

    // The provider must validate empty additional data in all three pool data formats.
    vm.expectCall(s_additionalPoolDataProvider, callData, 3);
    // Legacy pool data of 32 and 0 bytes.
    _resolveInboundLane(s_token, s_remotePool, abi.encode(uint256(DEFAULT_TOKEN_DECIMALS)), AMOUNT);
    _resolveInboundLane(s_token, s_remotePool, "", AMOUNT);
    // Pool data containing a source token does not require legacy receive permission.
    _resolveInboundLane(s_token, s_otherRemotePool, _getSourcePoolData(DEFAULT_TOKEN_DECIMALS, s_token, ""), AMOUNT);
  }

  function test_resolveInboundLane_OutboundPauseIgnored() public {
    s_connectionStore.setTokenPause(_arrayOf(s_token), ConnectionPolicyStore.PauseDirection.Outbound, true);
    uint64[] memory remoteChainSelectors = _getRemoteChainSelectors(DEST_CHAIN_SELECTOR);
    s_connectionStore.setChainPause(remoteChainSelectors, ConnectionPolicyStore.PauseDirection.Outbound, true);

    IPoolConfigManager.ResolvedInboundTransfer memory transfer =
      _resolveInboundLane(s_token, s_remotePool, _getSourcePoolData(DEFAULT_TOKEN_DECIMALS, s_token, ""), AMOUNT);

    assertEq(AMOUNT, transfer.localAmount);
  }

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
    // Keep the decimal difference within uint256 exponentiation limits.
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

  function test_resolveInboundLane_RevertWhen_SubsetTokenNotFound() public {
    _setConnectionPolicyConfig(_getConnectionPolicyConfigArgs(CONNECTION_POLICY, DEST_CHAIN_SELECTOR, true));

    vm.expectRevert(
      abi.encodeWithSelector(
        ConnectionPolicyStore.SubsetTokenNotFound.selector, CONNECTION_POLICY, DEST_CHAIN_SELECTOR, s_token
      )
    );
    _resolveInboundLane(s_token, s_remotePool, _getSourcePoolData(DEFAULT_TOKEN_DECIMALS, s_token, ""), AMOUNT);
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

    // An empty inline pool list rejects all source pools when policy trust is disabled.
    ConnectionPolicyStore.RemotePoolConfig memory remotePoolConfig = _getRemotePoolConfig();
    remotePoolConfig.remotePools = new IPoolConfigManager.RemotePool[](0);
    _setInlineRemotePoolConfig(_getInlineRemotePoolConfigArgs(s_token, DEST_CHAIN_SELECTOR, remotePoolConfig));

    vm.expectRevert(
      abi.encodeWithSelector(
        ConnectionPolicyStore.InvalidSourcePoolAddress.selector, s_token, DEST_CHAIN_SELECTOR, s_remotePool
      )
    );
    _resolveInboundLane(s_token, s_remotePool, sourcePoolData, AMOUNT);

    // Opting into policy pools must not allow a pool absent from both records.
    _setInlineRemotePoolConfigTrustingPolicy(s_token, s_otherRemotePool, false);
    vm.expectRevert(
      abi.encodeWithSelector(
        ConnectionPolicyStore.InvalidSourcePoolAddress.selector, s_token, DEST_CHAIN_SELECTOR, unknownPool
      )
    );
    _resolveInboundLane(s_token, unknownPool, sourcePoolData, AMOUNT);

    // Without a policy record for this chain, only inline pools are trusted.
    _setTokenRemotePoolPolicy(s_token, bytes32(0));
    vm.expectRevert(
      abi.encodeWithSelector(
        ConnectionPolicyStore.InvalidSourcePoolAddress.selector, s_token, DEST_CHAIN_SELECTOR, s_remotePool
      )
    );
    _resolveInboundLane(s_token, s_remotePool, sourcePoolData, AMOUNT);
  }

  function test_resolveInboundLane_RevertWhen_LegacyPoolDataNotAllowed() public {
    bytes[] memory sourcePoolDatas = new bytes[](2);
    sourcePoolDatas[0] = abi.encode(uint256(DEFAULT_TOKEN_DECIMALS));
    sourcePoolDatas[1] = "";

    for (uint256 i = 0; i < sourcePoolDatas.length; ++i) {
      vm.expectRevert(
        abi.encodeWithSelector(ConnectionPolicyStore.LegacyPoolDataNotAllowed.selector, s_token, DEST_CHAIN_SELECTOR)
      );
      _resolveInboundLane(s_token, s_remotePool, sourcePoolDatas[i], AMOUNT);
    }

    // Legacy receive permission applies only to the pool it is configured for.
    _allowLegacyPoolDataOnReceive();
    for (uint256 i = 0; i < sourcePoolDatas.length; ++i) {
      vm.expectRevert(
        abi.encodeWithSelector(ConnectionPolicyStore.LegacyPoolDataNotAllowed.selector, s_token, DEST_CHAIN_SELECTOR)
      );
      _resolveInboundLane(s_token, s_otherRemotePool, sourcePoolDatas[i], AMOUNT);
    }

    // An inline legacy permission must not affect other tokens using the same remote pool policy.
    _setTokenRemotePoolPolicy(s_otherToken, REMOTE_POOL_POLICY);
    _setInlineConnectionConfig(
      _getInlineConnectionConfigArgs(s_otherToken, DEST_CHAIN_SELECTOR, abi.encode(s_otherToken))
    );
    vm.expectRevert(
      abi.encodeWithSelector(ConnectionPolicyStore.LegacyPoolDataNotAllowed.selector, s_otherToken, DEST_CHAIN_SELECTOR)
    );
    _resolveInboundLane(s_otherToken, s_remotePool, sourcePoolDatas[0], AMOUNT);
  }

  function test_resolveInboundLane_RevertWhen_LegacyPoolDataNotAllowed_PolicyRecord() public {
    _assertPolicyPoolRejectsLegacyPoolData(s_remotePool);
  }

  function test_resolveInboundLane_RevertWhen_SourceTokenMismatch() public {
    // The source token is present, so it must match the configured remote token. Legacy permission does not skip this check.
    _allowLegacyPoolDataOnReceive();
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

    // Source token validation also applies to pools trusted through the policy opt-in.
    _setInlineRemotePoolConfigTrustingPolicy(s_token, s_otherRemotePool, true);
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
    bytes[] memory sourcePoolDatas = new bytes[](6);
    // Decimals above 255, in both formats.
    sourcePoolDatas[0] = abi.encode(uint256(256));
    sourcePoolDatas[1] = abi.encode(uint256(256), abi.encode(s_token), "");
    // Multi-token pool data with an empty source token.
    sourcePoolDatas[2] = abi.encode(uint256(DEFAULT_TOKEN_DECIMALS), "", "");
    // Trailing data after a valid ABI encoding.
    sourcePoolDatas[3] = abi.encodePacked(validPoolData, bytes1(0));
    // Overlapping offsets are decodable but not a canonical ABI encoding.
    sourcePoolDatas[4] =
      abi.encode(uint256(DEFAULT_TOKEN_DECIMALS), uint256(0x60), uint256(0x60), uint256(0x20), s_token);
    // Truncated data must return InvalidSourcePoolData instead of an empty decoding error.
    sourcePoolDatas[5] = abi.encode(uint256(DEFAULT_TOKEN_DECIMALS), uint256(0x40));

    for (uint256 i = 0; i < sourcePoolDatas.length; ++i) {
      vm.expectRevert(abi.encodeWithSelector(ConnectionPolicyStore.InvalidSourcePoolData.selector, sourcePoolDatas[i]));
      _resolveInboundLane(s_token, s_remotePool, sourcePoolDatas[i], AMOUNT);
    }
  }

  function test_resolveInboundLane_RevertWhen_PoolDataDoesNotDecode() public {
    // The length check passes, but an offset beyond the end must fail ABI decoding.
    bytes memory sourcePoolData = abi.encode(
      uint256(DEFAULT_TOKEN_DECIMALS), uint256(0x1000), uint256(0x80), uint256(0x20), bytes32(0), uint256(0)
    );

    vm.expectRevert(bytes(""));
    _resolveInboundLane(s_token, s_remotePool, sourcePoolData, AMOUNT);
  }

  function test_resolveInboundLane_RevertWhen_InvalidRemoteToken() public {
    bytes memory sourcePoolData = _getSourcePoolData(DEFAULT_TOKEN_DECIMALS, s_token, "");

    vm.mockCall(
      address(s_remoteTokenAddressProvider),
      abi.encodeCall(IRemoteTokenAddressProvider.getRemoteTokenAddress, (s_token, DEST_CHAIN_SELECTOR)),
      abi.encode(abi.encode(address(0)))
    );

    vm.expectRevert(
      abi.encodeWithSelector(ConnectionPolicyStore.InvalidRemoteToken.selector, s_token, DEST_CHAIN_SELECTOR)
    );
    _resolveInboundLane(s_token, s_remotePool, sourcePoolData, AMOUNT);

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
    vm.etch(address(s_remoteTokenAddressProvider), "");
    bytes memory sourcePoolData = _getSourcePoolData(DEFAULT_TOKEN_DECIMALS, s_token, "");

    vm.expectRevert();
    _resolveInboundLane(s_token, s_remotePool, sourcePoolData, AMOUNT);
  }

  function test_resolveInboundLane_RevertWhen_InvalidAdditionalPoolData() public {
    _setAdditionalPoolDataProvider(s_token);
    vm.mockCall(
      s_additionalPoolDataProvider,
      abi.encodeWithSelector(IAdditionalPoolDataProvider.isAdditionalPoolDataValid.selector),
      abi.encode(false)
    );
    bytes[] memory additionalPoolDatas = new bytes[](2);
    additionalPoolDatas[0] = abi.encode(keccak256("otherUid"));
    // Legacy permission must not bypass additional pool data validation.
    additionalPoolDatas[1] = "";

    for (uint256 i = 0; i < additionalPoolDatas.length; ++i) {
      bytes memory sourcePoolData = _getSourcePoolData(DEFAULT_TOKEN_DECIMALS, s_token, additionalPoolDatas[i]);
      vm.expectRevert(
        abi.encodeWithSelector(
          ConnectionPolicyStore.InvalidAdditionalPoolData.selector, s_token, DEST_CHAIN_SELECTOR, additionalPoolDatas[i]
        )
      );
      _resolveInboundLane(s_token, s_remotePool, sourcePoolData, AMOUNT);
    }
  }

  function test_resolveInboundLane_RevertWhen_OverflowDetected() public {
    // Scaling down by more than 77 decimals overflows the divisor.
    uint8 remoteDecimals = DEFAULT_TOKEN_DECIMALS + 78;
    bytes memory sourcePoolData = _getSourcePoolData(remoteDecimals, s_token, "");

    vm.expectRevert(
      abi.encodeWithSelector(
        ConnectionPolicyStore.OverflowDetected.selector, remoteDecimals, DEFAULT_TOKEN_DECIMALS, AMOUNT
      )
    );
    _resolveInboundLane(s_token, s_remotePool, sourcePoolData, AMOUNT);

    // Scaling up can overflow even when the decimal difference is within range.
    uint256 amount = type(uint256).max / 1e18 + 1;
    sourcePoolData = _getSourcePoolData(0, s_token, "");

    vm.expectRevert(
      abi.encodeWithSelector(ConnectionPolicyStore.OverflowDetected.selector, 0, DEFAULT_TOKEN_DECIMALS, amount)
    );
    _resolveInboundLane(s_token, s_remotePool, sourcePoolData, amount);

    // Scaling up by more than 77 decimals overflows the multiplier.
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

  /// @dev Policy pools require a source token, whether trusted directly or through trustPolicyRemotePools.
  function _assertPolicyPoolRejectsLegacyPoolData(
    bytes memory policyRemotePool
  ) private {
    bytes[2] memory legacyPoolDatas = [abi.encode(uint256(DEFAULT_TOKEN_DECIMALS)), bytes("")];
    for (uint256 i = 0; i < legacyPoolDatas.length; ++i) {
      vm.expectRevert(
        abi.encodeWithSelector(ConnectionPolicyStore.LegacyPoolDataNotAllowed.selector, s_token, DEST_CHAIN_SELECTOR)
      );
      _resolveInboundLane(s_token, policyRemotePool, legacyPoolDatas[i], AMOUNT);
    }

    _setInlineRemotePoolConfigTrustingPolicy(s_token, abi.encode(makeAddr("legacyPool")), true);
    for (uint256 i = 0; i < legacyPoolDatas.length; ++i) {
      vm.expectRevert(
        abi.encodeWithSelector(ConnectionPolicyStore.LegacyPoolDataNotAllowed.selector, s_token, DEST_CHAIN_SELECTOR)
      );
      _resolveInboundLane(s_token, policyRemotePool, legacyPoolDatas[i], AMOUNT);
    }
    _resolveInboundLane(s_token, policyRemotePool, _getSourcePoolData(DEFAULT_TOKEN_DECIMALS, s_token, ""), AMOUNT);
  }

  function _allowLegacyPoolDataOnReceive() private {
    _setInlineRemotePoolConfig(
      _getInlineRemotePoolConfigArgs(s_token, DEST_CHAIN_SELECTOR, _getRemotePoolConfigWithFlag())
    );
  }

  /// @dev Allow legacy data from s_remotePool only, with s_otherRemotePool as a control.
  function _getRemotePoolConfigWithFlag()
    private
    view
    returns (ConnectionPolicyStore.RemotePoolConfig memory remotePoolConfig)
  {
    remotePoolConfig = _getRemotePoolConfig();
    remotePoolConfig.remotePools = new IPoolConfigManager.RemotePool[](2);
    remotePoolConfig.remotePools[0].remotePoolAddress = s_remotePool;
    remotePoolConfig.remotePools[0].allowLegacyPoolDataOnReceive = true;
    remotePoolConfig.remotePools[1].remotePoolAddress = s_otherRemotePool;
    return remotePoolConfig;
  }
}
