// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IAdditionalPoolDataProvider} from "../../../../../interfaces/IAdditionalPoolDataProvider.sol";
import {IPoolConfigManager} from "../../../../../interfaces/IPoolConfigManager.sol";
import {IRemoteTokenAddressProvider} from "../../../../../interfaces/IRemoteTokenAddressProvider.sol";

import {FinalityCodec} from "../../../../../libraries/FinalityCodec.sol";
import {PolicyConfig} from "../../../../../libraries/PolicyConfig.sol";
import {ConnectionPolicyStore} from "../../../../../pools/MultiTokenPool/policies/ConnectionPolicyStore.sol";
import {ConnectionPolicyStoreSetup} from "./ConnectionPolicyStoreSetup.t.sol";

contract ConnectionPolicyStore_resolveOutboundLane is ConnectionPolicyStoreSetup {
  bytes4 internal constant WAIT_FOR_FINALITY = FinalityCodec.WAIT_FOR_FINALITY_FLAG;

  function test_resolveOutboundLane() public view {
    IPoolConfigManager.OutboundLane memory lane =
      s_connectionStore.resolveOutboundLane(s_token, DEST_CHAIN_SELECTOR, WAIT_FOR_FINALITY, "");

    IPoolConfigManager.OutboundLane memory expectedLane;
    expectedLane.remoteToken = abi.encode(s_token);
    expectedLane.destPoolData = _getSourcePoolData(DEFAULT_TOKEN_DECIMALS, s_token, "");
    expectedLane.localDecimals = DEFAULT_TOKEN_DECIMALS;
    expectedLane.lockOrBurnFlavor = IPoolConfigManager.LockOrBurnFlavor.Burn;
    assertEq(abi.encode(expectedLane), abi.encode(lane));
  }

  function test_resolveOutboundLane_InlineConnectionRecord() public {
    bytes memory remoteToken = abi.encode(makeAddr("remoteToken"));
    _setInlineConnectionConfig(_getInlineConnectionConfigArgs(s_token, DEST_CHAIN_SELECTOR, remoteToken));

    // The inline remote token takes precedence without calling the policy's provider.
    vm.expectCall(
      address(s_remoteTokenAddressProvider),
      abi.encodeCall(IRemoteTokenAddressProvider.getRemoteTokenAddress, (s_token, DEST_CHAIN_SELECTOR)),
      0
    );
    IPoolConfigManager.OutboundLane memory lane =
      s_connectionStore.resolveOutboundLane(s_token, DEST_CHAIN_SELECTOR, WAIT_FOR_FINALITY, "");

    assertEq(remoteToken, lane.remoteToken);
    assertEq(_getSourcePoolData(DEFAULT_TOKEN_DECIMALS, s_token, ""), lane.destPoolData);
  }

  function test_resolveOutboundLane_SubsetToken() public {
    _setConnectionPolicyConfig(_getConnectionPolicyConfigArgs(CONNECTION_POLICY, SOURCE_CHAIN_SELECTOR, true));
    _addSubsetToken(CONNECTION_POLICY, SOURCE_CHAIN_SELECTOR, s_token);
    _setRemotePoolPolicyConfig(
      _getRemotePoolPolicyConfigArgs(REMOTE_POOL_POLICY, SOURCE_CHAIN_SELECTOR, _getRemotePoolAddresses(s_remotePool))
    );

    IPoolConfigManager.OutboundLane memory lane =
      s_connectionStore.resolveOutboundLane(s_token, SOURCE_CHAIN_SELECTOR, WAIT_FOR_FINALITY, "");

    assertEq(abi.encode(s_token), lane.remoteToken);
  }

  function test_resolveOutboundLane_AdditionalPoolData() public {
    // Use nonzero finality to detect a missing argument when calling the provider.
    PolicyConfig.TokenProfileArgs[] memory profiles = _getTokenProfileArgs(s_token);
    profiles[0].allowedFinalityConfig = ALLOWED_FINALITY;
    profiles[0].additionalPoolDataProvider = s_additionalPoolDataProvider;
    s_connectionStore.applyTokenProfileUpdates(new address[](0), profiles);
    bytes memory tokenArgs = "tokenArgs";
    bytes memory additionalPoolData = abi.encode(keccak256("uid"));
    bytes memory callData = abi.encodeCall(
      IAdditionalPoolDataProvider.getAdditionalPoolData, (s_token, DEST_CHAIN_SELECTOR, REQUESTED_FINALITY, tokenArgs)
    );
    vm.mockCall(s_additionalPoolDataProvider, callData, abi.encode(additionalPoolData));

    vm.expectCall(s_additionalPoolDataProvider, callData);
    IPoolConfigManager.OutboundLane memory lane =
      s_connectionStore.resolveOutboundLane(s_token, DEST_CHAIN_SELECTOR, REQUESTED_FINALITY, tokenArgs);

    assertEq(_getSourcePoolData(DEFAULT_TOKEN_DECIMALS, s_token, additionalPoolData), lane.destPoolData);
  }

  function test_resolveOutboundLane_LegacyPoolData() public {
    _setAdditionalPoolDataProvider(s_token);
    ConnectionPolicyStore.RemotePoolConfig memory remotePoolConfig = _getRemotePoolConfig();
    remotePoolConfig.useLegacyPoolDataOnSend = true;
    _setInlineRemotePoolConfig(_getInlineRemotePoolConfigArgs(s_token, DEST_CHAIN_SELECTOR, remotePoolConfig));

    // Legacy destinations expect only token decimals, so the additional data provider must not be called.
    vm.expectCall(
      s_additionalPoolDataProvider,
      abi.encodeWithSelector(IAdditionalPoolDataProvider.getAdditionalPoolData.selector),
      0
    );
    IPoolConfigManager.OutboundLane memory lane =
      s_connectionStore.resolveOutboundLane(s_token, DEST_CHAIN_SELECTOR, WAIT_FOR_FINALITY, "");

    assertEq(abi.encode(uint256(DEFAULT_TOKEN_DECIMALS)), lane.destPoolData);
    // Omitting the source token from pool data does not change the destination token address.
    assertEq(abi.encode(s_token), lane.remoteToken);
  }

  function test_resolveOutboundLane_TrustPolicyRemotePools() public {
    // Trusting policy pools must not override the inline outbound data format.
    ConnectionPolicyStore.RemotePoolConfig memory remotePoolConfig = _getRemotePoolConfig();
    remotePoolConfig.remotePools[0].remotePoolAddress = s_otherRemotePool;
    remotePoolConfig.useLegacyPoolDataOnSend = true;
    ConnectionPolicyStore.InlineRemotePoolConfigArgs[] memory inlineConfigs =
      _getInlineRemotePoolConfigArgs(s_token, DEST_CHAIN_SELECTOR, remotePoolConfig);
    inlineConfigs[0].trustPolicyRemotePools = true;
    _setInlineRemotePoolConfig(inlineConfigs);

    IPoolConfigManager.OutboundLane memory lane =
      s_connectionStore.resolveOutboundLane(s_token, DEST_CHAIN_SELECTOR, WAIT_FOR_FINALITY, "");

    assertEq(abi.encode(uint256(DEFAULT_TOKEN_DECIMALS)), lane.destPoolData);
  }

  function test_resolveOutboundLane_FastFinality() public {
    PolicyConfig.TokenProfileArgs[] memory profiles = _getTokenProfileArgs(s_token);
    profiles[0].allowedFinalityConfig = ALLOWED_FINALITY;
    s_connectionStore.applyTokenProfileUpdates(new address[](0), profiles);

    IPoolConfigManager.OutboundLane memory lane =
      s_connectionStore.resolveOutboundLane(s_token, DEST_CHAIN_SELECTOR, REQUESTED_FINALITY, "");

    assertEq(abi.encode(s_token), lane.remoteToken);
  }

  function test_resolveOutboundLane_TokenProfile() public {
    address burnAddress = makeAddr("burnAddress");
    PolicyConfig.TokenProfileArgs[] memory profiles = _getTokenProfileArgs(s_token);
    profiles[0].lockOrBurnFlavor = IPoolConfigManager.LockOrBurnFlavor.BurnToAddress;
    profiles[0].lockOrBurnAddress = burnAddress;
    s_connectionStore.applyTokenProfileUpdates(new address[](0), profiles);

    IPoolConfigManager.OutboundLane memory lane =
      s_connectionStore.resolveOutboundLane(s_token, DEST_CHAIN_SELECTOR, WAIT_FOR_FINALITY, "");

    assertEq(uint8(IPoolConfigManager.LockOrBurnFlavor.BurnToAddress), uint8(lane.lockOrBurnFlavor));
    assertEq(burnAddress, lane.lockOrBurnAddress);
  }

  function test_resolveOutboundLane_InboundPauseIgnored() public {
    s_connectionStore.setTokenPause(_arrayOf(s_token), ConnectionPolicyStore.PauseDirection.Inbound, true);
    uint64[] memory remoteChainSelectors = _getRemoteChainSelectors(DEST_CHAIN_SELECTOR);
    s_connectionStore.setChainPause(remoteChainSelectors, ConnectionPolicyStore.PauseDirection.Inbound, true);

    IPoolConfigManager.OutboundLane memory lane =
      s_connectionStore.resolveOutboundLane(s_token, DEST_CHAIN_SELECTOR, WAIT_FOR_FINALITY, "");

    assertEq(abi.encode(s_token), lane.remoteToken);
  }

  // Reverts

  function test_resolveOutboundLane_RevertWhen_GloballyPaused() public {
    s_connectionStore.setGlobalPause(true);

    vm.expectRevert(ConnectionPolicyStore.GloballyPaused.selector);
    s_connectionStore.resolveOutboundLane(s_token, DEST_CHAIN_SELECTOR, WAIT_FOR_FINALITY, "");
  }

  function test_resolveOutboundLane_RevertWhen_TokenIsPaused() public {
    s_connectionStore.setTokenPause(_arrayOf(s_token), ConnectionPolicyStore.PauseDirection.Outbound, true);

    vm.expectRevert(
      abi.encodeWithSelector(
        ConnectionPolicyStore.TokenIsPaused.selector, s_token, ConnectionPolicyStore.PauseDirection.Outbound
      )
    );
    s_connectionStore.resolveOutboundLane(s_token, DEST_CHAIN_SELECTOR, WAIT_FOR_FINALITY, "");
  }

  function test_resolveOutboundLane_RevertWhen_ChainIsPaused() public {
    uint64[] memory remoteChainSelectors = _getRemoteChainSelectors(DEST_CHAIN_SELECTOR);
    s_connectionStore.setChainPause(remoteChainSelectors, ConnectionPolicyStore.PauseDirection.Outbound, true);

    vm.expectRevert(
      abi.encodeWithSelector(
        ConnectionPolicyStore.ChainIsPaused.selector, DEST_CHAIN_SELECTOR, ConnectionPolicyStore.PauseDirection.Outbound
      )
    );
    s_connectionStore.resolveOutboundLane(s_token, DEST_CHAIN_SELECTOR, WAIT_FOR_FINALITY, "");
  }

  function test_resolveOutboundLane_RevertWhen_ConnectionRecordNotFound() public {
    IPoolConfigManager.LaneKey[] memory lanes = new IPoolConfigManager.LaneKey[](3);
    lanes[0] = IPoolConfigManager.LaneKey({localToken: s_otherToken, remoteChainSelector: DEST_CHAIN_SELECTOR});
    lanes[1] = IPoolConfigManager.LaneKey({localToken: s_token, remoteChainSelector: OTHER_CHAIN_SELECTOR});
    lanes[2] =
      IPoolConfigManager.LaneKey({localToken: s_token, remoteChainSelector: PolicyConfig.DEFAULT_CHAIN_SELECTOR});

    for (uint256 i = 0; i < lanes.length; ++i) {
      vm.expectRevert(
        abi.encodeWithSelector(
          ConnectionPolicyStore.ConnectionRecordNotFound.selector, lanes[i].localToken, lanes[i].remoteChainSelector
        )
      );
      s_connectionStore.resolveOutboundLane(lanes[i].localToken, lanes[i].remoteChainSelector, WAIT_FOR_FINALITY, "");
    }
  }

  function test_resolveOutboundLane_RevertWhen_SubsetTokenNotFound() public {
    _setConnectionPolicyConfig(_getConnectionPolicyConfigArgs(CONNECTION_POLICY, SOURCE_CHAIN_SELECTOR, true));

    vm.expectRevert(
      abi.encodeWithSelector(
        ConnectionPolicyStore.SubsetTokenNotFound.selector, CONNECTION_POLICY, SOURCE_CHAIN_SELECTOR, s_token
      )
    );
    s_connectionStore.resolveOutboundLane(s_token, SOURCE_CHAIN_SELECTOR, WAIT_FOR_FINALITY, "");
  }

  function test_resolveOutboundLane_RevertWhen_RemotePoolRecordNotFound() public {
    _setTokenRemotePoolPolicy(s_token, bytes32(0));

    vm.expectRevert(
      abi.encodeWithSelector(ConnectionPolicyStore.RemotePoolRecordNotFound.selector, s_token, DEST_CHAIN_SELECTOR)
    );
    s_connectionStore.resolveOutboundLane(s_token, DEST_CHAIN_SELECTOR, WAIT_FOR_FINALITY, "");
  }

  function test_resolveOutboundLane_RevertWhen_InvalidRequestedFinality() public {
    bytes4 requestedFinality = REQUESTED_FINALITY;

    vm.expectRevert(
      abi.encodeWithSelector(FinalityCodec.InvalidRequestedFinality.selector, requestedFinality, bytes4(0))
    );
    s_connectionStore.resolveOutboundLane(s_token, DEST_CHAIN_SELECTOR, requestedFinality, "");
  }

  function test_resolveOutboundLane_RevertWhen_InvalidRemoteToken() public {
    bytes[] memory providerAnswers = new bytes[](3);
    providerAnswers[0] = "";
    providerAnswers[1] = abi.encode(address(0));
    providerAnswers[2] = hex"00";

    for (uint256 i = 0; i < providerAnswers.length; ++i) {
      vm.mockCall(
        address(s_remoteTokenAddressProvider),
        abi.encodeCall(IRemoteTokenAddressProvider.getRemoteTokenAddress, (s_token, DEST_CHAIN_SELECTOR)),
        abi.encode(providerAnswers[i])
      );

      vm.expectRevert(
        abi.encodeWithSelector(ConnectionPolicyStore.InvalidRemoteToken.selector, s_token, DEST_CHAIN_SELECTOR)
      );
      s_connectionStore.resolveOutboundLane(s_token, DEST_CHAIN_SELECTOR, WAIT_FOR_FINALITY, "");
    }

    _setRemoteTokenAddressProvider(CONNECTION_POLICY, address(0));

    vm.expectRevert(
      abi.encodeWithSelector(ConnectionPolicyStore.InvalidRemoteToken.selector, s_token, DEST_CHAIN_SELECTOR)
    );
    s_connectionStore.resolveOutboundLane(s_token, DEST_CHAIN_SELECTOR, WAIT_FOR_FINALITY, "");
  }

  function test_resolveOutboundLane_RevertWhen_RemoteTokenAddressProviderReverts() public {
    bytes memory revertData = abi.encodeWithSignature("Error(string)", "unknown token");
    vm.mockCallRevert(
      address(s_remoteTokenAddressProvider),
      abi.encodeCall(IRemoteTokenAddressProvider.getRemoteTokenAddress, (s_token, DEST_CHAIN_SELECTOR)),
      revertData
    );

    vm.expectRevert(revertData);
    s_connectionStore.resolveOutboundLane(s_token, DEST_CHAIN_SELECTOR, WAIT_FOR_FINALITY, "");
  }

  function test_resolveOutboundLane_RevertWhen_RemoteTokenAddressProviderHasNoCode() public {
    vm.etch(address(s_remoteTokenAddressProvider), "");

    vm.expectRevert();
    s_connectionStore.resolveOutboundLane(s_token, DEST_CHAIN_SELECTOR, WAIT_FOR_FINALITY, "");
  }

  function test_resolveOutboundLane_RevertWhen_AdditionalPoolDataProviderReverts() public {
    _setAdditionalPoolDataProvider(s_token);
    bytes memory revertData = abi.encodeWithSignature("Error(string)", "unknown uid");
    vm.mockCallRevert(
      s_additionalPoolDataProvider,
      abi.encodeWithSelector(IAdditionalPoolDataProvider.getAdditionalPoolData.selector),
      revertData
    );

    vm.expectRevert(revertData);
    s_connectionStore.resolveOutboundLane(s_token, DEST_CHAIN_SELECTOR, WAIT_FOR_FINALITY, "");
  }
}
