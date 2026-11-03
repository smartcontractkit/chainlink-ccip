// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../../../../interfaces/IPoolConfigManager.sol";

import {PolicyConfig} from "../../../../../libraries/PolicyConfig.sol";
import {
  SameAddressRemoteTokenAddressProvider
} from "../../../../../pools/MultiTokenPool/SameAddressRemoteTokenAddressProvider.sol";
import {ConnectionPolicyStore} from "../../../../../pools/MultiTokenPool/policies/ConnectionPolicyStore.sol";
import {BaseERC20} from "../../../../../tokens/BaseERC20.sol";
import {CrossChainToken} from "../../../../../tokens/CrossChainToken.sol";
import {BaseTest} from "../../../../BaseTest.t.sol";

/// @notice Register two tokens, configure a policy lane for s_token on DEST_CHAIN_SELECTOR, and authorize
/// s_poolConfigManager. The lane uses s_remoteTokenAddressProvider and trusts s_remotePool; s_otherToken has no policy.
contract ConnectionPolicyStoreSetup is BaseTest {
  bytes32 internal constant CONNECTION_POLICY = "CONNECTION_POLICY";
  bytes32 internal constant REMOTE_POOL_POLICY = "REMOTE_POOL_POLICY";
  bytes32 internal constant OTHER_POLICY = "OTHER_POLICY";
  /// @dev A remote chain with no configuration in setUp.
  uint64 internal constant OTHER_CHAIN_SELECTOR = 4949039107694359620;
  /// @dev Minimum allowed fast finality of 5 blocks.
  bytes4 internal constant ALLOWED_FINALITY = bytes4(uint32(5));
  /// @dev A valid fast finality request of 10 blocks.
  bytes4 internal constant REQUESTED_FINALITY = bytes4(uint32(10));

  ConnectionPolicyStore internal s_connectionStore;
  SameAddressRemoteTokenAddressProvider internal s_remoteTokenAddressProvider;
  address internal s_poolConfigManager = makeAddr("poolConfigManager");
  address internal s_token;
  address internal s_otherToken;
  bytes internal s_remotePool = abi.encode(makeAddr("remotePool"));
  bytes internal s_otherRemotePool = abi.encode(makeAddr("otherRemotePool"));
  address internal s_additionalPoolDataProvider = makeAddr("additionalPoolDataProvider");

  function setUp() public virtual override {
    super.setUp();
    s_token = address(_deployToken("TOKEN", DEFAULT_TOKEN_DECIMALS));
    s_otherToken = address(_deployToken("OTHER", USDC_TOKEN_DECIMALS));
    s_remoteTokenAddressProvider = new SameAddressRemoteTokenAddressProvider();
    s_connectionStore = new ConnectionPolicyStore(_arrayOf(s_poolConfigManager));
    // Configuration requires provider code even though the tests mock its responses.
    vm.etch(s_additionalPoolDataProvider, hex"00");
    vm.startPrank(s_poolConfigManager);

    _registerToken(s_token);
    _registerToken(s_otherToken);
    _setConnectionPolicyConfig(_getConnectionPolicyConfigArgs(CONNECTION_POLICY, DEST_CHAIN_SELECTOR, false));
    _setRemoteTokenAddressProvider(CONNECTION_POLICY, address(s_remoteTokenAddressProvider));
    _setRemotePoolPolicyConfig(
      _getRemotePoolPolicyConfigArgs(REMOTE_POOL_POLICY, DEST_CHAIN_SELECTOR, _getRemotePoolAddresses(s_remotePool))
    );
    _setTokenConnectionPolicy(s_token, CONNECTION_POLICY);
    _setTokenRemotePoolPolicy(s_token, REMOTE_POOL_POLICY);
  }

  function _deployToken(
    string memory name,
    uint8 decimals
  ) internal returns (CrossChainToken) {
    return new CrossChainToken(
      BaseERC20.ConstructorParams({
        name: name,
        symbol: name,
        decimals: decimals,
        maxSupply: 0,
        preMint: 0,
        preMintRecipient: address(0),
        ccipAdmin: OWNER
      }),
      OWNER,
      OWNER
    );
  }

  // ================================================================
  // │                          Arguments                           │
  // ================================================================

  function _getTokenProfileArgs(
    address localToken
  ) internal view returns (PolicyConfig.TokenProfileArgs[] memory profiles) {
    profiles = new PolicyConfig.TokenProfileArgs[](1);
    profiles[0] = PolicyConfig.TokenProfileArgs({
      localToken: localToken,
      localDecimals: BaseERC20(localToken).decimals(),
      lockOrBurnFlavor: IPoolConfigManager.LockOrBurnFlavor.Burn,
      allowedFinalityConfig: bytes4(0),
      lockOrBurnAddress: address(0),
      additionalPoolDataProvider: address(0)
    });
    return profiles;
  }

  function _getTokenPolicyArgs(
    address localToken,
    bytes32 policyName
  ) internal pure returns (PolicyConfig.TokenPolicyArgs[] memory tokenPolicies) {
    tokenPolicies = new PolicyConfig.TokenPolicyArgs[](1);
    tokenPolicies[0] = PolicyConfig.TokenPolicyArgs({localToken: localToken, policyName: policyName});
    return tokenPolicies;
  }

  function _getConnectionPolicyConfigArgs(
    bytes32 policyName,
    uint64 remoteChainSelector,
    bool onlySubsetTokens
  ) internal pure returns (ConnectionPolicyStore.ConnectionPolicyConfigArgs[] memory policyConfigs) {
    policyConfigs = new ConnectionPolicyStore.ConnectionPolicyConfigArgs[](1);
    policyConfigs[0] = ConnectionPolicyStore.ConnectionPolicyConfigArgs({
      policyName: policyName, remoteChainSelector: remoteChainSelector, onlySubsetTokens: onlySubsetTokens
    });
    return policyConfigs;
  }

  function _getInlineConnectionConfigArgs(
    address localToken,
    uint64 remoteChainSelector,
    bytes memory remoteToken
  ) internal pure returns (ConnectionPolicyStore.InlineConnectionConfigArgs[] memory inlineConfigs) {
    inlineConfigs = new ConnectionPolicyStore.InlineConnectionConfigArgs[](1);
    inlineConfigs[0] = ConnectionPolicyStore.InlineConnectionConfigArgs({
      localToken: localToken, remoteChainSelector: remoteChainSelector, remoteToken: remoteToken
    });
    return inlineConfigs;
  }

  function _getSubsetTokenArgs(
    bytes32 policyName,
    uint64 remoteChainSelector,
    address localToken
  ) internal pure returns (ConnectionPolicyStore.SubsetTokenArgs[] memory subsetTokens) {
    subsetTokens = new ConnectionPolicyStore.SubsetTokenArgs[](1);
    subsetTokens[0] = ConnectionPolicyStore.SubsetTokenArgs({
      policyName: policyName, remoteChainSelector: remoteChainSelector, localToken: localToken
    });
    return subsetTokens;
  }

  function _getRemoteTokenAddressProviderArgs(
    bytes32 policyName,
    address remoteTokenAddressProvider
  ) internal pure returns (ConnectionPolicyStore.RemoteTokenAddressProviderArgs[] memory providers) {
    providers = new ConnectionPolicyStore.RemoteTokenAddressProviderArgs[](1);
    providers[0] = ConnectionPolicyStore.RemoteTokenAddressProviderArgs({
      policyName: policyName, remoteTokenAddressProvider: remoteTokenAddressProvider
    });
    return providers;
  }

  function _getRemotePoolConfig() internal view returns (ConnectionPolicyStore.RemotePoolConfig memory) {
    IPoolConfigManager.RemotePool[] memory remotePools = new IPoolConfigManager.RemotePool[](1);
    remotePools[0].remotePoolAddress = s_remotePool;
    return ConnectionPolicyStore.RemotePoolConfig({useLegacyPoolDataOnSend: false, remotePools: remotePools});
  }

  function _getRemotePoolAddresses(
    bytes memory remotePool
  ) internal pure returns (bytes[] memory remotePoolAddresses) {
    remotePoolAddresses = new bytes[](1);
    remotePoolAddresses[0] = remotePool;
    return remotePoolAddresses;
  }

  function _getRemotePoolPolicyConfigArgs(
    bytes32 policyName,
    uint64 remoteChainSelector,
    bytes[] memory remotePoolAddresses
  ) internal pure returns (ConnectionPolicyStore.RemotePoolPolicyConfigArgs[] memory policyConfigs) {
    policyConfigs = new ConnectionPolicyStore.RemotePoolPolicyConfigArgs[](1);
    policyConfigs[0] = ConnectionPolicyStore.RemotePoolPolicyConfigArgs({
      policyName: policyName, remoteChainSelector: remoteChainSelector, remotePoolAddresses: remotePoolAddresses
    });
    return policyConfigs;
  }

  function _getInlineRemotePoolConfigArgs(
    address localToken,
    uint64 remoteChainSelector,
    ConnectionPolicyStore.RemotePoolConfig memory remotePoolConfig
  ) internal pure returns (ConnectionPolicyStore.InlineRemotePoolConfigArgs[] memory inlineConfigs) {
    inlineConfigs = new ConnectionPolicyStore.InlineRemotePoolConfigArgs[](1);
    inlineConfigs[0] = ConnectionPolicyStore.InlineRemotePoolConfigArgs({
      localToken: localToken,
      remoteChainSelector: remoteChainSelector,
      trustPolicyRemotePools: false,
      remotePoolConfig: remotePoolConfig
    });
    return inlineConfigs;
  }

  function _getPolicyRecordKeys(
    bytes32 policyName,
    uint64 remoteChainSelector
  ) internal pure returns (PolicyConfig.PolicyRecordKey[] memory recordKeys) {
    recordKeys = new PolicyConfig.PolicyRecordKey[](1);
    recordKeys[0] = PolicyConfig.PolicyRecordKey({policyName: policyName, remoteChainSelector: remoteChainSelector});
    return recordKeys;
  }

  function _getPolicyNames(
    bytes32 policyName
  ) internal pure returns (bytes32[] memory policyNames) {
    policyNames = new bytes32[](1);
    policyNames[0] = policyName;
    return policyNames;
  }

  function _getRemoteChainSelectors(
    uint64 remoteChainSelector
  ) internal pure returns (uint64[] memory remoteChainSelectors) {
    remoteChainSelectors = new uint64[](1);
    remoteChainSelectors[0] = remoteChainSelector;
    return remoteChainSelectors;
  }

  function _getLaneKeys(
    address localToken,
    uint64 remoteChainSelector
  ) internal pure returns (IPoolConfigManager.LaneKey[] memory lanes) {
    lanes = new IPoolConfigManager.LaneKey[](1);
    lanes[0] = IPoolConfigManager.LaneKey({localToken: localToken, remoteChainSelector: remoteChainSelector});
    return lanes;
  }

  function _getLaneConnection(
    address localToken,
    uint64 remoteChainSelector
  ) internal view returns (IPoolConfigManager.LaneConnection memory) {
    return s_connectionStore.getLaneConnectionBatch(_getLaneKeys(localToken, remoteChainSelector))[0];
  }

  /// @dev Decode only the token count to avoid a stack-too-deep error in the test profile without via-ir.
  function _getRemotePoolPolicyTokenCount(
    bytes32 policyName
  ) internal view returns (uint32) {
    (bool success, bytes memory returnData) =
      address(s_connectionStore).staticcall(abi.encodeCall(ConnectionPolicyStore.getRemotePoolPolicy, (policyName)));
    assertTrue(success);
    return abi.decode(returnData, (uint32));
  }

  /// @dev Compare encoded results to avoid a stack-too-deep error in the test profile without via-ir.
  function _assertRemotePoolPolicy(
    bytes32 policyName,
    uint32 tokenCount,
    ConnectionPolicyStore.RemotePoolPolicyConfigArgs[] memory policyConfigs
  ) internal view {
    (bool success, bytes memory returnData) =
      address(s_connectionStore).staticcall(abi.encodeCall(ConnectionPolicyStore.getRemotePoolPolicy, (policyName)));
    assertTrue(success);
    assertEq(abi.encode(tokenCount, policyConfigs), returnData);
  }

  function _getSourcePoolData(
    uint8 sourceDecimals,
    address sourceToken,
    bytes memory additionalPoolData
  ) internal pure returns (bytes memory) {
    return abi.encode(uint256(sourceDecimals), abi.encode(sourceToken), additionalPoolData);
  }

  // ================================================================
  // │                         Config writes                        │
  // ================================================================

  function _setAdditionalPoolDataProvider(
    address localToken
  ) internal {
    PolicyConfig.TokenProfileArgs[] memory profiles = _getTokenProfileArgs(localToken);
    profiles[0].additionalPoolDataProvider = s_additionalPoolDataProvider;
    s_connectionStore.applyTokenProfileUpdates(new address[](0), profiles);
  }

  function _registerToken(
    address localToken
  ) internal {
    s_connectionStore.applyTokenProfileUpdates(new address[](0), _getTokenProfileArgs(localToken));
  }

  function _setTokenConnectionPolicy(
    address localToken,
    bytes32 policyName
  ) internal {
    s_connectionStore.applyTokenConnectionPolicyUpdates(_getTokenPolicyArgs(localToken, policyName));
  }

  function _setTokenRemotePoolPolicy(
    address localToken,
    bytes32 policyName
  ) internal {
    s_connectionStore.applyTokenRemotePoolPolicyUpdates(_getTokenPolicyArgs(localToken, policyName));
  }

  function _setConnectionPolicyConfig(
    ConnectionPolicyStore.ConnectionPolicyConfigArgs[] memory policyConfigs
  ) internal {
    s_connectionStore.applyConnectionPolicyConfigUpdates(
      new PolicyConfig.PolicyRecordKey[](0), new bytes32[](0), policyConfigs
    );
  }

  function _setInlineConnectionConfig(
    ConnectionPolicyStore.InlineConnectionConfigArgs[] memory inlineConfigs
  ) internal {
    s_connectionStore.applyInlineConnectionConfigUpdates(new IPoolConfigManager.LaneKey[](0), inlineConfigs);
  }

  function _addSubsetToken(
    bytes32 policyName,
    uint64 remoteChainSelector,
    address localToken
  ) internal {
    s_connectionStore.applySubsetTokenUpdates(
      new ConnectionPolicyStore.SubsetTokenArgs[](0), _getSubsetTokenArgs(policyName, remoteChainSelector, localToken)
    );
  }

  function _setRemoteTokenAddressProvider(
    bytes32 policyName,
    address remoteTokenAddressProvider
  ) internal {
    s_connectionStore.applyRemoteTokenAddressProviderUpdates(
      _getRemoteTokenAddressProviderArgs(policyName, remoteTokenAddressProvider)
    );
  }

  function _setRemotePoolPolicyConfig(
    ConnectionPolicyStore.RemotePoolPolicyConfigArgs[] memory policyConfigs
  ) internal {
    s_connectionStore.applyRemotePoolPolicyConfigUpdates(
      new PolicyConfig.PolicyRecordKey[](0), new bytes32[](0), policyConfigs
    );
  }

  function _setInlineRemotePoolConfig(
    ConnectionPolicyStore.InlineRemotePoolConfigArgs[] memory inlineConfigs
  ) internal {
    s_connectionStore.applyInlineRemotePoolConfigUpdates(new IPoolConfigManager.LaneKey[](0), inlineConfigs);
  }

  function _setInlineRemotePoolConfigTrustingPolicy(
    address localToken,
    bytes memory remotePool,
    bool allowLegacyPoolDataOnReceive
  ) internal {
    ConnectionPolicyStore.RemotePoolConfig memory remotePoolConfig = _getRemotePoolConfig();
    remotePoolConfig.remotePools[0].remotePoolAddress = remotePool;
    remotePoolConfig.remotePools[0].allowLegacyPoolDataOnReceive = allowLegacyPoolDataOnReceive;
    ConnectionPolicyStore.InlineRemotePoolConfigArgs[] memory inlineConfigs =
      _getInlineRemotePoolConfigArgs(localToken, DEST_CHAIN_SELECTOR, remotePoolConfig);
    inlineConfigs[0].trustPolicyRemotePools = true;
    _setInlineRemotePoolConfig(inlineConfigs);
  }
}
