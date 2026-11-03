// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../../../../interfaces/IPoolConfigManager.sol";

import {PolicyConfig} from "../../../../../libraries/PolicyConfig.sol";
import {ConnectionPolicyStoreSetup} from "./ConnectionPolicyStoreSetup.t.sol";

contract ConnectionPolicyStore_getTokenConfigBatch is ConnectionPolicyStoreSetup {
  /// @dev OP Mainnet, configured only through an inline connection.
  uint64 internal constant THIRD_CHAIN_SELECTOR = 3734403246176062136;

  function test_getTokenConfigBatch() public {
    PolicyConfig.TokenProfileArgs[] memory profiles = _getTokenProfileArgs(s_token);
    profiles[0].lockOrBurnFlavor = IPoolConfigManager.LockOrBurnFlavor.BurnToAddress;
    profiles[0].lockOrBurnAddress = makeAddr("burnAddress");
    profiles[0].allowedFinalityConfig = ALLOWED_FINALITY;
    profiles[0].additionalPoolDataProvider = s_additionalPoolDataProvider;
    s_connectionStore.applyTokenProfileUpdates(new address[](0), profiles);
    // Overlap inline and policy chains to check deduplication, and exclude OTHER through its subset restriction.
    _setConnectionPolicyConfig(_getConnectionPolicyConfigArgs(CONNECTION_POLICY, SOURCE_CHAIN_SELECTOR, false));
    _setConnectionPolicyConfig(_getConnectionPolicyConfigArgs(CONNECTION_POLICY, OTHER_CHAIN_SELECTOR, true));
    _setInlineConnectionConfig(_getInlineConnectionConfigArgs(s_token, SOURCE_CHAIN_SELECTOR, abi.encode(s_token)));
    _setInlineConnectionConfig(_getInlineConnectionConfigArgs(s_token, THIRD_CHAIN_SELECTOR, abi.encode(s_token)));
    address[] memory localTokens = new address[](3);
    localTokens[0] = s_token;
    localTokens[1] = makeAddr("unknownToken");
    localTokens[2] = s_otherToken;

    IPoolConfigManager.TokenConfig[] memory configs = s_connectionStore.getTokenConfigBatch(localTokens);

    assertEq(3, configs.length);
    IPoolConfigManager.TokenConfig memory expectedConfig;
    expectedConfig.isRegistered = true;
    expectedConfig.localDecimals = DEFAULT_TOKEN_DECIMALS;
    expectedConfig.lockOrBurnFlavor = profiles[0].lockOrBurnFlavor;
    expectedConfig.allowedFinalityConfig = profiles[0].allowedFinalityConfig;
    expectedConfig.lockOrBurnAddress = profiles[0].lockOrBurnAddress;
    expectedConfig.additionalPoolDataProvider = profiles[0].additionalPoolDataProvider;
    // Inline chains precede policy chains, with no duplicates.
    expectedConfig.supportedChains = new uint64[](3);
    expectedConfig.supportedChains[0] = SOURCE_CHAIN_SELECTOR;
    expectedConfig.supportedChains[1] = THIRD_CHAIN_SELECTOR;
    expectedConfig.supportedChains[2] = DEST_CHAIN_SELECTOR;
    expectedConfig.connectionPolicyName = CONNECTION_POLICY;
    expectedConfig.remotePoolPolicyName = REMOTE_POOL_POLICY;
    assertEq(abi.encode(expectedConfig), abi.encode(configs[0]));

    IPoolConfigManager.TokenConfig memory emptyConfig;
    assertEq(abi.encode(emptyConfig), abi.encode(configs[1]));

    IPoolConfigManager.TokenConfig memory otherConfig;
    otherConfig.isRegistered = true;
    otherConfig.localDecimals = USDC_TOKEN_DECIMALS;
    assertEq(abi.encode(otherConfig), abi.encode(configs[2]));
  }

  function test_getTokenConfigBatch_SubsetToken() public {
    _setConnectionPolicyConfig(_getConnectionPolicyConfigArgs(CONNECTION_POLICY, OTHER_CHAIN_SELECTOR, true));
    _addSubsetToken(CONNECTION_POLICY, OTHER_CHAIN_SELECTOR, s_token);

    uint64[] memory supportedChains = s_connectionStore.getTokenConfigBatch(_arrayOf(s_token))[0].supportedChains;

    assertEq(2, supportedChains.length);
    assertEq(DEST_CHAIN_SELECTOR, supportedChains[0]);
    assertEq(OTHER_CHAIN_SELECTOR, supportedChains[1]);
  }

  function test_getTokenConfigBatch_DeregisteredToken() public {
    s_connectionStore.applyTokenProfileUpdates(_arrayOf(s_otherToken), new PolicyConfig.TokenProfileArgs[](0));

    IPoolConfigManager.TokenConfig[] memory configs = s_connectionStore.getTokenConfigBatch(_arrayOf(s_otherToken));

    // Unlike getTokenProfiles, this view returns zero values for deregistered tokens.
    IPoolConfigManager.TokenConfig memory emptyConfig;
    assertEq(abi.encode(emptyConfig), abi.encode(configs[0]));
  }
}
