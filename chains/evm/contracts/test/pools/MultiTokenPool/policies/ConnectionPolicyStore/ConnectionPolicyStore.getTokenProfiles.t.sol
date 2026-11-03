// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../../../../interfaces/IPoolConfigManager.sol";

import {PolicyConfig} from "../../../../../libraries/PolicyConfig.sol";
import {ConnectionPolicyStore} from "../../../../../pools/MultiTokenPool/policies/ConnectionPolicyStore.sol";
import {ConnectionPolicyStoreSetup} from "./ConnectionPolicyStoreSetup.t.sol";

contract ConnectionPolicyStore_getTokenProfiles is ConnectionPolicyStoreSetup {
  function test_getTokenProfiles() public {
    PolicyConfig.TokenProfileArgs[] memory profiles = _getTokenProfileArgs(s_token);
    profiles[0].lockOrBurnFlavor = IPoolConfigManager.LockOrBurnFlavor.BurnToAddress;
    profiles[0].lockOrBurnAddress = makeAddr("burnAddress");
    profiles[0].allowedFinalityConfig = ALLOWED_FINALITY;
    profiles[0].additionalPoolDataProvider = s_additionalPoolDataProvider;
    s_connectionStore.applyTokenProfileUpdates(new address[](0), profiles);
    s_connectionStore.setTokenPause(_arrayOf(s_token), ConnectionPolicyStore.PauseDirection.Inbound, true);
    // Pause flags can be set before a token is registered.
    address unknownToken = makeAddr("unknownToken");
    s_connectionStore.setTokenPause(_arrayOf(unknownToken), ConnectionPolicyStore.PauseDirection.Outbound, true);
    address[] memory localTokens = new address[](2);
    localTokens[0] = s_token;
    localTokens[1] = unknownToken;

    ConnectionPolicyStore.TokenProfile[] memory storedProfiles = s_connectionStore.getTokenProfiles(localTokens);

    assertEq(2, storedProfiles.length);
    ConnectionPolicyStore.TokenProfile memory expectedProfile = ConnectionPolicyStore.TokenProfile({
      isRegistered: true,
      localDecimals: DEFAULT_TOKEN_DECIMALS,
      lockOrBurnFlavor: profiles[0].lockOrBurnFlavor,
      lockOrBurnAddress: profiles[0].lockOrBurnAddress,
      allowedFinalityConfig: profiles[0].allowedFinalityConfig,
      isOutboundPaused: false,
      isInboundPaused: true,
      additionalPoolDataProvider: profiles[0].additionalPoolDataProvider
    });
    assertEq(abi.encode(expectedProfile), abi.encode(storedProfiles[0]));
    ConnectionPolicyStore.TokenProfile memory unknownProfile;
    unknownProfile.isOutboundPaused = true;
    assertEq(abi.encode(unknownProfile), abi.encode(storedProfiles[1]));
  }

  function test_getTokenProfiles_DeregisteredToken() public {
    s_connectionStore.setTokenPause(_arrayOf(s_otherToken), ConnectionPolicyStore.PauseDirection.Inbound, true);
    s_connectionStore.applyTokenProfileUpdates(_arrayOf(s_otherToken), new PolicyConfig.TokenProfileArgs[](0));

    ConnectionPolicyStore.TokenProfile memory storedProfile =
      s_connectionStore.getTokenProfiles(_arrayOf(s_otherToken))[0];

    // Deregistration preserves the stored profile and pause flags.
    assertFalse(storedProfile.isRegistered);
    assertEq(USDC_TOKEN_DECIMALS, storedProfile.localDecimals);
    assertEq(uint8(IPoolConfigManager.LockOrBurnFlavor.Burn), uint8(storedProfile.lockOrBurnFlavor));
    assertTrue(storedProfile.isInboundPaused);
  }
}
