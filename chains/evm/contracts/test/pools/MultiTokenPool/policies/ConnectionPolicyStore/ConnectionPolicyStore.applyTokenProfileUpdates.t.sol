// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ILockBox} from "../../../../../interfaces/ILockBox.sol";
import {IPoolConfigManager} from "../../../../../interfaces/IPoolConfigManager.sol";

import {PolicyConfig} from "../../../../../libraries/PolicyConfig.sol";
import {ConnectionPolicyStore} from "../../../../../pools/MultiTokenPool/policies/ConnectionPolicyStore.sol";
import {ConnectionPolicyStoreSetup} from "./ConnectionPolicyStoreSetup.t.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

import {IERC20Metadata} from "@openzeppelin/contracts@5.3.0/token/ERC20/extensions/IERC20Metadata.sol";

contract ConnectionPolicyStore_applyTokenProfileUpdates is ConnectionPolicyStoreSetup {
  uint8 internal constant NEW_TOKEN_DECIMALS = 8;

  address internal s_newToken;
  address internal s_burnAddress = makeAddr("burnAddress");
  /// @dev Has code. Each test mocks whether it supports s_newToken.
  address internal s_lockBox = makeAddr("lockBox");

  function setUp() public override {
    super.setUp();
    s_newToken = address(_deployToken("NEW", NEW_TOKEN_DECIMALS));
    vm.etch(s_lockBox, hex"00");
  }

  function test_applyTokenProfileUpdates() public {
    PolicyConfig.TokenProfileArgs[] memory profiles = _getTokenProfileArgs(s_newToken);
    profiles[0].lockOrBurnFlavor = IPoolConfigManager.LockOrBurnFlavor.BurnToAddress;
    profiles[0].lockOrBurnAddress = s_burnAddress;
    profiles[0].allowedFinalityConfig = ALLOWED_FINALITY;
    profiles[0].additionalPoolDataProvider = s_additionalPoolDataProvider;

    vm.expectEmit();
    emit ConnectionPolicyStore.TokenProfileSet(
      s_newToken,
      NEW_TOKEN_DECIMALS,
      IPoolConfigManager.LockOrBurnFlavor.BurnToAddress,
      s_burnAddress,
      ALLOWED_FINALITY,
      s_additionalPoolDataProvider
    );
    s_connectionStore.applyTokenProfileUpdates(new address[](0), profiles);

    assertEq(abi.encode(_getStoredProfile(profiles[0], true, false, false)), abi.encode(_getTokenProfile(s_newToken)));
    (bool isRegistered, bytes32 connectionPolicyName, bytes32 remotePoolPolicyName) =
      s_connectionStore.getTokenPolicies(s_newToken);
    assertTrue(isRegistered);
    assertEq(bytes32(0), connectionPolicyName);
    assertEq(bytes32(0), remotePoolPolicyName);
    assertEq(3, s_connectionStore.getAllConfiguredTokens(0, type(uint64).max).length);
    // Registering a token does not enable any connections.
    assertFalse(s_connectionStore.hasLaneBatch(_getLaneKeys(s_newToken, DEST_CHAIN_SELECTOR))[0]);
  }

  function test_applyTokenProfileUpdates_RegisteredTokenKeepsPolicies() public {
    PolicyConfig.TokenProfileArgs[] memory profiles = _getTokenProfileArgs(s_token);
    profiles[0].lockOrBurnFlavor = IPoolConfigManager.LockOrBurnFlavor.BurnFrom;
    profiles[0].allowedFinalityConfig = ALLOWED_FINALITY;
    profiles[0].additionalPoolDataProvider = s_additionalPoolDataProvider;

    vm.expectEmit();
    emit ConnectionPolicyStore.TokenProfileSet(
      s_token,
      DEFAULT_TOKEN_DECIMALS,
      IPoolConfigManager.LockOrBurnFlavor.BurnFrom,
      address(0),
      ALLOWED_FINALITY,
      s_additionalPoolDataProvider
    );
    s_connectionStore.applyTokenProfileUpdates(new address[](0), profiles);

    assertEq(abi.encode(_getStoredProfile(profiles[0], true, false, false)), abi.encode(_getTokenProfile(s_token)));
    (, bytes32 connectionPolicyName, bytes32 remotePoolPolicyName) = s_connectionStore.getTokenPolicies(s_token);
    assertEq(CONNECTION_POLICY, connectionPolicyName);
    assertEq(REMOTE_POOL_POLICY, remotePoolPolicyName);
    assertTrue(s_connectionStore.hasLaneBatch(_getLaneKeys(s_token, DEST_CHAIN_SELECTOR))[0]);
  }

  function test_applyTokenProfileUpdates_LockFlavor() public {
    PolicyConfig.TokenProfileArgs[] memory profiles = _getTokenProfileArgs(s_newToken);
    profiles[0].lockOrBurnFlavor = IPoolConfigManager.LockOrBurnFlavor.Lock;
    profiles[0].lockOrBurnAddress = s_lockBox;
    _mockLockBoxSupportsToken(true);

    s_connectionStore.applyTokenProfileUpdates(new address[](0), profiles);

    assertEq(abi.encode(_getStoredProfile(profiles[0], true, false, false)), abi.encode(_getTokenProfile(s_newToken)));
  }

  function test_applyTokenProfileUpdates_BurnToAddressAcceptsAnyAddress() public {
    PolicyConfig.TokenProfileArgs[] memory profiles = _getTokenProfileArgs(s_newToken);
    profiles[0].lockOrBurnFlavor = IPoolConfigManager.LockOrBurnFlavor.BurnToAddress;
    profiles[0].lockOrBurnAddress = s_burnAddress;

    s_connectionStore.applyTokenProfileUpdates(new address[](0), profiles);
    assertEq(s_burnAddress, _getTokenProfile(s_newToken).lockOrBurnAddress);

    // The target does not need code.
    address otherBurnAddress = makeAddr("otherBurnAddress");
    profiles[0].lockOrBurnAddress = otherBurnAddress;
    s_connectionStore.applyTokenProfileUpdates(new address[](0), profiles);
    assertEq(otherBurnAddress, _getTokenProfile(s_newToken).lockOrBurnAddress);

    // Some tokens burn on a transfer to address(0).
    profiles[0].lockOrBurnAddress = address(0);
    s_connectionStore.applyTokenProfileUpdates(new address[](0), profiles);
    assertEq(address(0), _getTokenProfile(s_newToken).lockOrBurnAddress);
  }

  function test_applyTokenProfileUpdates_TokenWithoutDecimals() public {
    // ERC20 decimals is optional, so use the supplied value if the call reverts.
    vm.mockCallRevert(s_newToken, abi.encodeCall(IERC20Metadata.decimals, ()), "");
    PolicyConfig.TokenProfileArgs[] memory profiles = new PolicyConfig.TokenProfileArgs[](1);
    profiles[0].localToken = s_newToken;
    profiles[0].localDecimals = USDC_TOKEN_DECIMALS;

    s_connectionStore.applyTokenProfileUpdates(new address[](0), profiles);

    assertEq(USDC_TOKEN_DECIMALS, _getTokenProfile(s_newToken).localDecimals);
  }

  function test_applyTokenProfileUpdates_RemoveToken() public {
    vm.expectEmit();
    emit ConnectionPolicyStore.TokenRemoved(s_otherToken);
    s_connectionStore.applyTokenProfileUpdates(_arrayOf(s_otherToken), new PolicyConfig.TokenProfileArgs[](0));

    (bool isRegistered,,) = s_connectionStore.getTokenPolicies(s_otherToken);
    assertFalse(isRegistered);
    assertEq(_arrayOf(s_token), s_connectionStore.getAllConfiguredTokens(0, type(uint64).max));
    assertFalse(s_connectionStore.getTokenConfigBatch(_arrayOf(s_otherToken))[0].isRegistered);
  }

  function test_applyTokenProfileUpdates_RemovedTokenKeepsProfileAndPauseFlags() public {
    PolicyConfig.TokenProfileArgs[] memory profiles = _getTokenProfileArgs(s_otherToken);
    s_connectionStore.setTokenPause(_arrayOf(s_otherToken), ConnectionPolicyStore.PauseDirection.Both, true);

    s_connectionStore.applyTokenProfileUpdates(_arrayOf(s_otherToken), new PolicyConfig.TokenProfileArgs[](0));

    assertEq(abi.encode(_getStoredProfile(profiles[0], false, true, true)), abi.encode(_getTokenProfile(s_otherToken)));

    // Re-registering a token must not clear its pause flags.
    s_connectionStore.applyTokenProfileUpdates(new address[](0), profiles);
    assertEq(abi.encode(_getStoredProfile(profiles[0], true, true, true)), abi.encode(_getTokenProfile(s_otherToken)));
  }

  function test_applyTokenProfileUpdates_RemovalsRunFirst() public {
    // Removing and re-registering a token allows new decimals in the same call.
    vm.mockCall(s_otherToken, abi.encodeCall(IERC20Metadata.decimals, ()), abi.encode(NEW_TOKEN_DECIMALS));
    PolicyConfig.TokenProfileArgs[] memory profiles = _getTokenProfileArgs(s_otherToken);

    vm.expectEmit();
    emit ConnectionPolicyStore.TokenRemoved(s_otherToken);
    vm.expectEmit();
    emit ConnectionPolicyStore.TokenProfileSet(
      s_otherToken, NEW_TOKEN_DECIMALS, IPoolConfigManager.LockOrBurnFlavor.Burn, address(0), bytes4(0), address(0)
    );
    s_connectionStore.applyTokenProfileUpdates(_arrayOf(s_otherToken), profiles);

    (bool isRegistered,,) = s_connectionStore.getTokenPolicies(s_otherToken);
    assertTrue(isRegistered);
    assertEq(NEW_TOKEN_DECIMALS, _getTokenProfile(s_otherToken).localDecimals);
  }

  // Reverts

  function test_applyTokenProfileUpdates_RevertWhen_ZeroAddressInvalid() public {
    PolicyConfig.TokenProfileArgs[] memory profiles = new PolicyConfig.TokenProfileArgs[](1);

    vm.expectRevert(ConnectionPolicyStore.ZeroAddressInvalid.selector);
    s_connectionStore.applyTokenProfileUpdates(new address[](0), profiles);
  }

  function test_applyTokenProfileUpdates_RevertWhen_InvalidDecimalArgs_NewToken() public {
    PolicyConfig.TokenProfileArgs[] memory profiles = _getTokenProfileArgs(s_newToken);
    profiles[0].localDecimals = USDC_TOKEN_DECIMALS;

    vm.expectRevert(
      abi.encodeWithSelector(
        ConnectionPolicyStore.InvalidDecimalArgs.selector, s_newToken, NEW_TOKEN_DECIMALS, USDC_TOKEN_DECIMALS
      )
    );
    s_connectionStore.applyTokenProfileUpdates(new address[](0), profiles);
  }

  function test_applyTokenProfileUpdates_RevertWhen_InvalidDecimalArgs_RegisteredToken() public {
    // Stored decimals cannot change while registered, even if the token reports a different value.
    vm.mockCall(s_token, abi.encodeCall(IERC20Metadata.decimals, ()), abi.encode(USDC_TOKEN_DECIMALS));
    PolicyConfig.TokenProfileArgs[] memory profiles = _getTokenProfileArgs(s_token);

    vm.expectRevert(
      abi.encodeWithSelector(
        ConnectionPolicyStore.InvalidDecimalArgs.selector, s_token, DEFAULT_TOKEN_DECIMALS, USDC_TOKEN_DECIMALS
      )
    );
    s_connectionStore.applyTokenProfileUpdates(new address[](0), profiles);
  }

  function test_applyTokenProfileUpdates_RevertWhen_InvalidLockOrBurnAddress() public {
    // Lock requires a compatible LockBox. Burn, BurnFrom, and BurnWithFrom require a zero target address.
    IPoolConfigManager.LockOrBurnFlavor[] memory flavors = new IPoolConfigManager.LockOrBurnFlavor[](5);
    address[] memory lockOrBurnAddresses = new address[](5);
    flavors[0] = IPoolConfigManager.LockOrBurnFlavor.Lock;
    flavors[1] = IPoolConfigManager.LockOrBurnFlavor.Lock;
    // A lock box that does not support the token.
    lockOrBurnAddresses[1] = s_lockBox;
    _mockLockBoxSupportsToken(false);
    flavors[2] = IPoolConfigManager.LockOrBurnFlavor.Burn;
    lockOrBurnAddresses[2] = s_burnAddress;
    flavors[3] = IPoolConfigManager.LockOrBurnFlavor.BurnFrom;
    lockOrBurnAddresses[3] = s_burnAddress;
    flavors[4] = IPoolConfigManager.LockOrBurnFlavor.BurnWithFrom;
    lockOrBurnAddresses[4] = s_burnAddress;

    for (uint256 i = 0; i < flavors.length; ++i) {
      PolicyConfig.TokenProfileArgs[] memory profiles = _getTokenProfileArgs(s_newToken);
      profiles[0].lockOrBurnFlavor = flavors[i];
      profiles[0].lockOrBurnAddress = lockOrBurnAddresses[i];

      vm.expectRevert(
        abi.encodeWithSelector(
          ConnectionPolicyStore.InvalidLockOrBurnAddress.selector, s_newToken, flavors[i], lockOrBurnAddresses[i]
        )
      );
      s_connectionStore.applyTokenProfileUpdates(new address[](0), profiles);
    }
  }

  function test_applyTokenProfileUpdates_RevertWhen_LocalTokenNotRegistered() public {
    vm.expectRevert(abi.encodeWithSelector(PolicyConfig.LocalTokenNotRegistered.selector, s_newToken));
    s_connectionStore.applyTokenProfileUpdates(_arrayOf(s_newToken), new PolicyConfig.TokenProfileArgs[](0));
  }

  function test_applyTokenProfileUpdates_RevertWhen_ProviderHasNoCode() public {
    address providerWithoutCode = makeAddr("providerWithoutCode");
    PolicyConfig.TokenProfileArgs[] memory profiles = _getTokenProfileArgs(s_newToken);
    profiles[0].additionalPoolDataProvider = providerWithoutCode;

    vm.expectRevert(abi.encodeWithSelector(ConnectionPolicyStore.ProviderHasNoCode.selector, providerWithoutCode));
    s_connectionStore.applyTokenProfileUpdates(new address[](0), profiles);
  }

  function test_applyTokenProfileUpdates_RevertWhen_TokenHasConfig() public {
    // Each policy assignment and inline record independently prevents deregistration.
    _setTokenConnectionPolicy(s_otherToken, CONNECTION_POLICY);
    _expectTokenHasConfigRevert();
    _setTokenConnectionPolicy(s_otherToken, bytes32(0));

    _setTokenRemotePoolPolicy(s_otherToken, REMOTE_POOL_POLICY);
    _expectTokenHasConfigRevert();
    _setTokenRemotePoolPolicy(s_otherToken, bytes32(0));

    _setInlineConnectionConfig(
      _getInlineConnectionConfigArgs(s_otherToken, DEST_CHAIN_SELECTOR, abi.encode(s_otherToken))
    );
    _expectTokenHasConfigRevert();
    s_connectionStore.applyInlineConnectionConfigUpdates(
      _getLaneKeys(s_otherToken, DEST_CHAIN_SELECTOR), new ConnectionPolicyStore.InlineConnectionConfigArgs[](0)
    );

    _setInlineRemotePoolConfig(
      _getInlineRemotePoolConfigArgs(s_otherToken, DEST_CHAIN_SELECTOR, _getRemotePoolConfig())
    );
    _expectTokenHasConfigRevert();
  }

  function test_applyTokenProfileUpdates_RevertWhen_UnauthorizedCaller() public {
    PolicyConfig.TokenProfileArgs[] memory profiles = _getTokenProfileArgs(s_newToken);
    vm.startPrank(STRANGER);

    vm.expectRevert(abi.encodeWithSelector(AuthorizedCallers.UnauthorizedCaller.selector, STRANGER));
    s_connectionStore.applyTokenProfileUpdates(new address[](0), profiles);
  }

  function _mockLockBoxSupportsToken(
    bool isSupported
  ) private {
    vm.mockCall(s_lockBox, abi.encodeCall(ILockBox.isTokenSupported, (s_newToken)), abi.encode(isSupported));
  }

  function _expectTokenHasConfigRevert() private {
    vm.expectRevert(abi.encodeWithSelector(ConnectionPolicyStore.TokenHasConfig.selector, s_otherToken));
    s_connectionStore.applyTokenProfileUpdates(_arrayOf(s_otherToken), new PolicyConfig.TokenProfileArgs[](0));
  }

  function _getTokenProfile(
    address localToken
  ) private view returns (ConnectionPolicyStore.TokenProfile memory) {
    return s_connectionStore.getTokenProfiles(_arrayOf(localToken))[0];
  }

  function _getStoredProfile(
    PolicyConfig.TokenProfileArgs memory profile,
    bool isRegistered,
    bool isOutboundPaused,
    bool isInboundPaused
  ) private pure returns (ConnectionPolicyStore.TokenProfile memory) {
    return ConnectionPolicyStore.TokenProfile({
      isRegistered: isRegistered,
      localDecimals: profile.localDecimals,
      lockOrBurnFlavor: profile.lockOrBurnFlavor,
      lockOrBurnAddress: profile.lockOrBurnAddress,
      allowedFinalityConfig: profile.allowedFinalityConfig,
      isOutboundPaused: isOutboundPaused,
      isInboundPaused: isInboundPaused,
      additionalPoolDataProvider: profile.additionalPoolDataProvider
    });
  }
}
