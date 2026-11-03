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

  function test_resolveInboundLane_LegacyPoolData() public {
    // s_remotePool has the flag on, and s_otherRemotePool has it off.
    _allowLegacyPoolDataOnReceive();

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

  function test_resolveInboundLane_TrustPolicyRemotePools() public {
    // The inline record lists only s_otherRemotePool, as a legacy pool. The policy record gives s_remotePool.
    _setInlineRemotePoolConfigTrustingPolicy(s_token, s_otherRemotePool, true);
    bytes memory legacyPoolData = abi.encode(uint256(DEFAULT_TOKEN_DECIMALS));

    _resolveInboundLane(s_token, s_remotePool, _getSourcePoolData(DEFAULT_TOKEN_DECIMALS, s_token, ""), AMOUNT);
    _resolveInboundLane(s_token, s_otherRemotePool, legacyPoolData, AMOUNT);

    // The remote pool of the policy record keeps the flag of the policy record.
    vm.expectRevert(
      abi.encodeWithSelector(ConnectionPolicyStore.LegacyPoolDataNotAllowed.selector, s_token, DEST_CHAIN_SELECTOR)
    );
    _resolveInboundLane(s_token, s_remotePool, legacyPoolData, AMOUNT);
  }

  function test_resolveInboundLane_TrustPolicyRemotePools_InlineFlagWins() public {
    bytes memory legacyPoolData = abi.encode(uint256(DEFAULT_TOKEN_DECIMALS));
    // Both records list s_remotePool. The inline record has the flag on, and the policy record has it off.
    _setInlineRemotePoolConfigTrustingPolicy(s_token, s_remotePool, true);
    _resolveInboundLane(s_token, s_remotePool, legacyPoolData, AMOUNT);

    // Another token of the policy keeps the flag of the policy record.
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
    // An inline record without remote pools still works inbound when it trusts the remote pools of the policy record.
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
    // The policy record of DEST_CHAIN_SELECTOR trusts s_remotePool, and the one of SOURCE_CHAIN_SELECTOR trusts
    // s_otherRemotePool.
    _setRemotePoolPolicyConfig(
      _getRemotePoolPolicyConfigArgs(
        REMOTE_POOL_POLICY, SOURCE_CHAIN_SELECTOR, _getRemotePoolAddresses(s_otherRemotePool)
      )
    );
    bytes memory legacyPool = abi.encode(makeAddr("legacyPool"));
    _setInlineRemotePoolConfigTrustingPolicy(s_token, legacyPool, true);
    bytes memory sourcePoolData = _getSourcePoolData(DEFAULT_TOKEN_DECIMALS, s_token, "");

    // The inline record adds only the policy record of its own chain.
    _resolveInboundLane(s_token, s_remotePool, sourcePoolData, AMOUNT);
    vm.expectRevert(
      abi.encodeWithSelector(
        ConnectionPolicyStore.InvalidSourcePoolAddress.selector, s_token, DEST_CHAIN_SELECTOR, s_otherRemotePool
      )
    );
    _resolveInboundLane(s_token, s_otherRemotePool, sourcePoolData, AMOUNT);

    // After the removal of the policy record, no other record of the policy serves the chain. The token trusts only
    // the remote pools of its inline record.
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
    // The inline record lists s_remotePool. Then the policy record replaces s_remotePool with s_otherRemotePool.
    _setInlineRemotePoolConfigTrustingPolicy(s_token, s_remotePool, false);
    _setRemotePoolPolicyConfig(
      _getRemotePoolPolicyConfigArgs(
        REMOTE_POOL_POLICY, DEST_CHAIN_SELECTOR, _getRemotePoolAddresses(s_otherRemotePool)
      )
    );

    // A change of the policy record does not reach a remote pool that the inline record lists. The token trusts the new
    // remote pool of the policy record at once.
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
    // s_remotePool has the flag on, and s_otherRemotePool has it off.
    _allowLegacyPoolDataOnReceive();
    bytes memory callData = abi.encodeCall(
      IAdditionalPoolDataProvider.isAdditionalPoolDataValid, (s_token, DEST_CHAIN_SELECTOR, WAIT_FOR_FINALITY, "")
    );
    vm.mockCall(s_additionalPoolDataProvider, callData, abi.encode(true));

    // The provider decides whether a source without additional pool data is valid.
    vm.expectCall(s_additionalPoolDataProvider, callData, 3);
    // Legacy pool data of 32 and 0 bytes.
    _resolveInboundLane(s_token, s_remotePool, abi.encode(uint256(DEFAULT_TOKEN_DECIMALS)), AMOUNT);
    _resolveInboundLane(s_token, s_remotePool, "", AMOUNT);
    // Pool data with the source token, from a source token without a provider. The flag of the pool does not matter.
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

  function test_resolveInboundLane_RevertWhen_SubsetTokenNotFound() public {
    // A record with onlySubsetTokens on, without s_token in its subset.
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

    // An inline record without remote pools works outbound only.
    ConnectionPolicyStore.RemotePoolConfig memory remotePoolConfig = _getRemotePoolConfig();
    remotePoolConfig.remotePools = new IPoolConfigManager.RemotePool[](0);
    _setInlineRemotePoolConfig(_getInlineRemotePoolConfigArgs(s_token, DEST_CHAIN_SELECTOR, remotePoolConfig));

    vm.expectRevert(
      abi.encodeWithSelector(
        ConnectionPolicyStore.InvalidSourcePoolAddress.selector, s_token, DEST_CHAIN_SELECTOR, s_remotePool
      )
    );
    _resolveInboundLane(s_token, s_remotePool, sourcePoolData, AMOUNT);

    // An inline record that trusts the policy record still rejects a remote pool that neither record lists.
    _setInlineRemotePoolConfigTrustingPolicy(s_token, s_otherRemotePool, false);
    vm.expectRevert(
      abi.encodeWithSelector(
        ConnectionPolicyStore.InvalidSourcePoolAddress.selector, s_token, DEST_CHAIN_SELECTOR, unknownPool
      )
    );
    _resolveInboundLane(s_token, unknownPool, sourcePoolData, AMOUNT);

    // A token without a policy record for the chain trusts only the remote pools of its inline record.
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

    // The flag of s_remotePool does not apply to s_otherRemotePool in the same record.
    _allowLegacyPoolDataOnReceive();
    for (uint256 i = 0; i < sourcePoolDatas.length; ++i) {
      vm.expectRevert(
        abi.encodeWithSelector(ConnectionPolicyStore.LegacyPoolDataNotAllowed.selector, s_token, DEST_CHAIN_SELECTOR)
      );
      _resolveInboundLane(s_token, s_otherRemotePool, sourcePoolDatas[i], AMOUNT);
    }

    // The flag of the inline record of s_token does not apply to the policy record that s_otherToken uses.
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
    // A trusted remote pool burned another token. A source token in the pool data is checked also from a remote pool
    // that can send pool data without it.
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

    // The same check runs for a remote pool of the policy record that the inline record trusts.
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
    // The long format with an empty source token.
    sourcePoolDatas[2] = abi.encode(uint256(DEFAULT_TOKEN_DECIMALS), "", "");
    // One byte more than abi.encode makes.
    sourcePoolDatas[3] = abi.encodePacked(validPoolData, bytes1(0));
    // Both offsets point to the source token. abi.decode accepts it, but abi.encode never makes it.
    sourcePoolDatas[4] =
      abi.encode(uint256(DEFAULT_TOKEN_DECIMALS), uint256(0x60), uint256(0x60), uint256(0x20), s_token);
    // Too short to hold a source token. abi.decode alone reverts on it with no error data.
    sourcePoolDatas[5] = abi.encode(uint256(DEFAULT_TOKEN_DECIMALS), uint256(0x40));

    for (uint256 i = 0; i < sourcePoolDatas.length; ++i) {
      vm.expectRevert(abi.encodeWithSelector(ConnectionPolicyStore.InvalidSourcePoolData.selector, sourcePoolDatas[i]));
      _resolveInboundLane(s_token, s_remotePool, sourcePoolDatas[i], AMOUNT);
    }
  }

  function test_resolveInboundLane_RevertWhen_PoolDataDoesNotDecode() public {
    // Long enough for a source token, but the first offset points past the end, so abi.decode reverts.
    bytes memory sourcePoolData = abi.encode(
      uint256(DEFAULT_TOKEN_DECIMALS), uint256(0x1000), uint256(0x80), uint256(0x20), bytes32(0), uint256(0)
    );

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
    // The provider loses its code after the set.
    vm.etch(address(s_remoteTokenAddressProvider), "");
    bytes memory sourcePoolData = _getSourcePoolData(DEFAULT_TOKEN_DECIMALS, s_token, "");

    // The empty answer fails the decode.
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
    // A provider can reject a source without additional pool data.
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

  /// @dev A remote pool of a policy record never accepts legacy pool data: not for a token that uses the policy record,
  /// and not for a token whose inline record trusts it. Long pool data from it runs.
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

  /// @dev Sets _getRemotePoolConfigWithFlag as an inline record of s_token.
  function _allowLegacyPoolDataOnReceive() private {
    _setInlineRemotePoolConfig(
      _getInlineRemotePoolConfigArgs(s_token, DEST_CHAIN_SELECTOR, _getRemotePoolConfigWithFlag())
    );
  }

  /// @dev s_remotePool can send pool data without the source token, and s_otherRemotePool cannot.
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
