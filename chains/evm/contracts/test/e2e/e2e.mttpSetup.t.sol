// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolConfigManager} from "../../interfaces/IPoolConfigManager.sol";
import {IPoolV2} from "../../interfaces/IPoolV2.sol";
import {IRemoteTokenAddressProvider} from "../../interfaces/IRemoteTokenAddressProvider.sol";
import {IRouter} from "../../interfaces/IRouter.sol";
import {ITokenAdminRegistry} from "../../interfaces/ITokenAdminRegistry.sol";

import {IBurnMintERC20} from "../../interfaces/IBurnMintERC20.sol";
import {FinalityCodec} from "../../libraries/FinalityCodec.sol";
import {PolicyConfig} from "../../libraries/PolicyConfig.sol";
import {Pool} from "../../libraries/Pool.sol";
import {RateLimiter} from "../../libraries/RateLimiter.sol";
import {BurnMintTokenPool} from "../../pools/BurnMintTokenPool.sol";
import {BurnMintMultiTokenPool} from "../../pools/MultiTokenPool/BurnMintMultiTokenPool.sol";
import {PoolConfigManager} from "../../pools/MultiTokenPool/PoolConfigManager.sol";
import {ConnectionPolicyStore} from "../../pools/MultiTokenPool/policies/ConnectionPolicyStore.sol";
import {FeePolicyStore} from "../../pools/MultiTokenPool/policies/FeePolicyStore.sol";
import {SecurityPolicyStore} from "../../pools/MultiTokenPool/policies/SecurityPolicyStore.sol";
import {TokenPool} from "../../pools/TokenPool.sol";
import {TokenAdminRegistry} from "../../tokenAdminRegistry/TokenAdminRegistry.sol";
import {TokenFixture} from "../TokenFixture.t.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

import {IERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/utils/SafeERC20.sol";

/// @notice Remote token address provider backed by explicit mappings. The spec wants a provider for policy
/// connections. SameAddressRemoteTokenAddressProvider only works when the local and remote token addresses match, and
/// they don't in these fixtures, so each token is mapped by hand.
contract MappingRemoteTokenAddressProvider is IRemoteTokenAddressProvider {
  mapping(address localToken => mapping(uint64 remoteChainSelector => bytes remoteToken)) private s_remoteTokens;

  function setRemoteToken(
    address localToken,
    uint64 remoteChainSelector,
    bytes calldata remoteToken
  ) external {
    s_remoteTokens[localToken][remoteChainSelector] = remoteToken;
  }

  function getRemoteTokenAddress(
    address localToken,
    uint64 remoteChainSelector
  ) external view returns (bytes memory) {
    return s_remoteTokens[localToken][remoteChainSelector];
  }
}

/// @notice Two MTTP stacks deployed on SOURCE_CHAIN_SELECTOR (A) and DEST_CHAIN_SELECTOR (B).
/// @dev Each chain has T1/T2/T3 (6 tokens total), one BurnMintMultiTokenPool, one PoolConfigManager and the three
/// stores. The pool config manager is an authorized caller on its stores, and the pool is registered in the token
/// admin registry for every token. The baseline is a working long-format lane in both directions: no fees, rate limits disabled, minimum destination overhead, and no hook or
/// CCVs.
abstract contract e2e_mttpSetup is TokenFixture {
  // One policy name per type so a swapped field fails the assertions.
  bytes32 internal constant CONNECTION_POLICY = "CONNECTION";
  bytes32 internal constant REMOTE_POOL_POLICY = "REMOTE_POOL";
  bytes32 internal constant RATE_LIMIT_POLICY = "RATE_LIMIT";
  bytes32 internal constant FEE_POLICY = "FEE";
  bytes32 internal constant DEST_OVERHEAD_POLICY = "DEST_OVERHEAD";
  bytes32 internal constant SECURITY_POLICY = "SECURITY";
  /// @dev A third chain selector. Only used to check that a chain-specific record doesn't leak into another lane's
  /// default.
  uint64 internal constant THIRD_CHAIN_SELECTOR = 4949039107694359620; // Arbitrum One.

  /// @dev The minimum destination byte overhead the store accepts. It is also enough for the 192 byte long pool data.
  uint32 internal constant MIN_DEST_BYTES_OVERHEAD = 192;
  /// @dev abi.encode(decimals, abi.encode(token), "") is exactly 192 bytes, the length of the long pool data format.
  uint256 internal constant LONG_POOL_DATA_LENGTH = 192;
  /// @dev Legacy pool data is only the token decimals.
  uint256 internal constant LEGACY_POOL_DATA_LENGTH = 32;

  uint8 internal constant T1_DECIMALS = 18;
  uint8 internal constant T2_DECIMALS = 18;
  uint8 internal constant T3_DECIMALS = 18;

  address internal s_issuer = makeAddr("issuer");
  address internal s_pauser = makeAddr("pauser");
  address internal s_unpauser = makeAddr("unpauser");
  address internal s_feeCollector = makeAddr("feeCollector");
  address internal s_allowedSender = makeAddr("allowedSender");
  address internal s_disallowedSender = makeAddr("disallowedSender");
  address internal s_receiver = makeAddr("receiver");

  struct MttpChain {
    uint64 localSelector;
    uint64 remoteSelector;
    IRouter router;
    address onRamp;
    address offRamp;
    TokenAdminRegistry tar;
    ConnectionPolicyStore connectionStore;
    FeePolicyStore feeStore;
    SecurityPolicyStore securityStore;
    PoolConfigManager pcm;
    BurnMintMultiTokenPool pool;
    MttpRouterHarness harness;
    MappingRemoteTokenAddressProvider remoteTokenProvider;
    address[3] tokens; // T1, T2, T3 on this chain.
  }

  MttpChain internal s_a;
  MttpChain internal s_b;

  function setUp() public virtual override {
    super.setUp();

    _deployChain(s_a, SOURCE_CHAIN_SELECTOR, DEST_CHAIN_SELECTOR, s_sourceRouter, "A");
    _deployChain(s_b, DEST_CHAIN_SELECTOR, SOURCE_CHAIN_SELECTOR, s_destRouter, "B");

    // Tokens and the TAR are per chain. Register the pool for every token and grant it mint/burn.
    for (uint256 i = 0; i < 3; ++i) {
      _registerToken(s_a, s_a.tokens[i], _decimalsOf(i));
      _registerToken(s_b, s_b.tokens[i], _decimalsOf(i));
    }

    _configureBaseline(s_a, s_b);
    _configureBaseline(s_b, s_a);

    vm.stopPrank();
  }

  // ================================================================
  // │                        Deployment                            │
  // ================================================================

  function _decimalsOf(
    uint256 tokenIndex
  ) internal pure returns (uint8) {
    if (tokenIndex == 0) return T1_DECIMALS;
    if (tokenIndex == 1) return T2_DECIMALS;
    return T3_DECIMALS;
  }

  function _deployChain(
    MttpChain storage c,
    uint64 localSelector,
    uint64 remoteSelector,
    IRouter router,
    string memory label
  ) internal {
    c.localSelector = localSelector;
    c.remoteSelector = remoteSelector;
    c.router = router;
    // The Router funds the pool and calls lockOrBurn in one transaction. The harness does the same, so a rejected
    // send rolls the funding back with it.
    c.harness = new MttpRouterHarness();
    c.onRamp = address(c.harness);
    c.offRamp = makeAddr(string.concat("offRamp", label));
    c.tar = new TokenAdminRegistry();

    c.connectionStore = new ConnectionPolicyStore(new address[](0));
    c.feeStore = new FeePolicyStore(new address[](0));
    c.securityStore = new SecurityPolicyStore(new address[](0));
    c.pcm =
      new PoolConfigManager(ITokenAdminRegistry(address(c.tar)), c.connectionStore, c.feeStore, c.securityStore, OWNER);

    AuthorizedCallers.AuthorizedCallerArgs memory callerArgs = AuthorizedCallers.AuthorizedCallerArgs({
      addedCallers: _arrayOf(address(c.pcm)), removedCallers: new address[](0)
    });
    c.connectionStore.applyAuthorizedCallerUpdates(callerArgs);
    c.feeStore.applyAuthorizedCallerUpdates(callerArgs);
    c.securityStore.applyAuthorizedCallerUpdates(callerArgs);
    c.pcm.grantRole(c.pcm.CONFIG_ROLE(), OWNER);
    c.pcm.grantRole(c.pcm.FEE_COLLECTOR_ROLE(), OWNER);

    c.pool = new BurnMintMultiTokenPool(address(s_mockRMNRemote), address(router), address(c.pcm));
    c.remoteTokenProvider = new MappingRemoteTokenAddressProvider();

    for (uint256 i = 0; i < 3; ++i) {
      c.tokens[i] = address(_deployCrossChainToken(string.concat("T", vm.toString(i + 1), label), _decimalsOf(i)));
      ICrossChainTokenLike(c.tokens[i]).grantMintAndBurnRoles(address(c.pool));
      ICrossChainTokenLike(c.tokens[i]).grantMintAndBurnRoles(OWNER);
      ICrossChainTokenLike(c.tokens[i]).mint(s_allowedSender, 1e30);
      ICrossChainTokenLike(c.tokens[i]).mint(s_disallowedSender, 1e30);
    }

    // Same as Router.applyRampUpdates: the pool accepts the local onRamp and offRamp.
    _setMockRouterOnRamp(address(router), remoteSelector, c.onRamp);
    _setMockRouterOffRamp(address(router), remoteSelector, c.offRamp, true);
  }

  /// @dev Registers the pool for a token and sets its profile. The profile has to come first, because the registry's
  /// setPool calls isSupportedToken, which reads the pool config manager.
  function _registerToken(
    MttpChain storage c,
    address token,
    uint8 decimals
  ) internal {
    _setProfile(c.pcm, token, decimals, IPoolConfigManager.LockOrBurnFlavor.Burn);
    c.tar.proposeAdministrator(token, OWNER);
    c.tar.acceptAdminRole(token);
    c.tar.setPool(token, address(c.pool));
  }

  /// @dev Deploys an extra token with custom decimals, outside the three baseline assets, and gives the pool mint and
  /// burn roles. Used by the decimal compatibility test. Call as OWNER.
  function _deployExtraToken(
    MttpChain storage c,
    string memory name,
    uint8 decimals
  ) internal returns (address token) {
    token = address(_deployCrossChainToken(name, decimals));
    ICrossChainTokenLike(token).grantMintAndBurnRoles(address(c.pool));
    return token;
  }

  /// @dev Registers an extra token with the full policy set and a remote token mapping, the same way the baseline
  /// tokens are set up. Call as OWNER.
  function _registerExtraToken(
    MttpChain storage c,
    address token,
    uint8 decimals,
    address remoteToken
  ) internal {
    _setProfile(c.pcm, token, decimals, IPoolConfigManager.LockOrBurnFlavor.Burn);
    _assignPolicy(c.pcm, token, CONNECTION_POLICY, REMOTE_POOL_POLICY);

    // Rate limit policies fix their decimals on the first record, so a token with other decimals needs its own.
    bytes32 rateLimitPolicy = RATE_LIMIT_POLICY;
    if (decimals != T1_DECIMALS) {
      rateLimitPolicy = bytes32(abi.encodePacked("RATE_LIMIT_", decimals));
      RateLimiter.Config memory disabled;
      _setRateLimitPolicyConfigWithDecimals(
        c.pcm, rateLimitPolicy, PolicyConfig.DEFAULT_CHAIN_SELECTOR, disabled, decimals
      );
    }
    _assignPolicy(c.pcm, token, rateLimitPolicy, FEE_POLICY, DEST_OVERHEAD_POLICY, SECURITY_POLICY);

    c.tar.proposeAdministrator(token, OWNER);
    c.tar.acceptAdminRole(token);
    c.tar.setPool(token, address(c.pool));
    c.remoteTokenProvider.setRemoteToken(token, c.remoteSelector, abi.encode(remoteToken));
  }

  // ================================================================
  // │                     Baseline configuration                   │
  // ================================================================

  /// @dev Creates one policy of each type and assigns every token to them. The connection and remote pool policies
  /// need explicit chain records because they have no defaults. Fee, rate limit, destination overhead and security
  /// use default records. The remote token provider maps each local token to its remote token.
  function _configureBaseline(
    MttpChain storage c,
    MttpChain storage remote
  ) internal {
    _setConnectionPolicyConfig(c.pcm, CONNECTION_POLICY, c.remoteSelector, false);
    _setRemoteTokenProvider(c.pcm, CONNECTION_POLICY, address(c.remoteTokenProvider));

    bytes[] memory remotePools = new bytes[](1);
    remotePools[0] = abi.encode(address(remote.pool));
    _setRemotePoolPolicyConfig(c.pcm, REMOTE_POOL_POLICY, c.remoteSelector, remotePools);

    RateLimiter.Config memory disabled;
    _setRateLimitPolicyConfig(c.pcm, RATE_LIMIT_POLICY, PolicyConfig.DEFAULT_CHAIN_SELECTOR, disabled);
    FeePolicyStore.FeeConfig memory zeroFee;
    _setFeePolicyConfig(c.pcm, FEE_POLICY, PolicyConfig.DEFAULT_CHAIN_SELECTOR, zeroFee);
    _setDestOverheadPolicyConfig(
      c.pcm,
      DEST_OVERHEAD_POLICY,
      PolicyConfig.DEFAULT_CHAIN_SELECTOR,
      DEFAULT_TOKEN_DEST_GAS_OVERHEAD,
      MIN_DEST_BYTES_OVERHEAD
    );
    _setSecurityPolicyConfig(
      c.pcm, SECURITY_POLICY, PolicyConfig.DEFAULT_CHAIN_SELECTOR, new address[](0), new address[](0), address(0)
    );

    for (uint256 i = 0; i < 3; ++i) {
      address token = c.tokens[i];
      _assignPolicy(c.pcm, token, CONNECTION_POLICY, REMOTE_POOL_POLICY);
      _assignPolicy(c.pcm, token, RATE_LIMIT_POLICY, FEE_POLICY, DEST_OVERHEAD_POLICY, SECURITY_POLICY);
      c.remoteTokenProvider.setRemoteToken(token, c.remoteSelector, abi.encode(remote.tokens[i]));
    }
  }

  // ================================================================
  // │                       Transfer helpers                       │
  // ================================================================

  /// @dev Approves the harness, then atomically pulls the sender's tokens and burns the post-fee amount.
  function _send(
    MttpChain storage src,
    MttpChain storage dst,
    uint256 tokenIndex,
    uint256 amount,
    bytes4 finality,
    address originalSender
  ) internal returns (bytes memory destPoolData, uint256 destAmount, bytes memory destTokenAddress) {
    _approvePool(src, src.tokens[tokenIndex], amount, originalSender);
    return _lockOrBurn(src, dst, tokenIndex, amount, finality, originalSender);
  }

  /// @dev Calls the router harness, which moves the amount from the sender into the pool and calls lockOrBurn in one
  /// transaction. Split out from _send so a test can put expectRevert directly on that call.
  function _lockOrBurn(
    MttpChain storage src,
    MttpChain storage dst,
    uint256 tokenIndex,
    uint256 amount,
    bytes4 finality,
    address originalSender
  ) internal returns (bytes memory destPoolData, uint256 destAmount, bytes memory destTokenAddress) {
    Pool.LockOrBurnInV1 memory lockOrBurnIn = Pool.LockOrBurnInV1({
      receiver: abi.encode(s_receiver),
      remoteChainSelector: dst.localSelector,
      originalSender: originalSender,
      amount: amount,
      localToken: src.tokens[tokenIndex]
    });

    return src.harness.forward(src.tokens[tokenIndex], address(src.pool), lockOrBurnIn, finality, "");
  }

  /// @dev Sends and delivers in one call, and also returns the fee that stayed on the source pool.
  function _sendAndDeliver(
    MttpChain storage src,
    MttpChain storage dst,
    uint256 tokenIndex,
    uint256 amount
  ) internal returns (uint256 received, uint256 fee) {
    uint256 poolBefore = _balance(src.tokens[tokenIndex], address(src.pool));
    (bytes memory poolData, uint256 postFeeAmount, bytes memory destTokenAddress) =
      _send(src, dst, tokenIndex, amount, FinalityCodec.WAIT_FOR_FINALITY_FLAG, s_allowedSender);
    received = _deliver(
      dst, src, abi.decode(destTokenAddress, (address)), postFeeAmount, poolData, FinalityCodec.WAIT_FOR_FINALITY_FLAG
    );
    fee = _balance(src.tokens[tokenIndex], address(src.pool)) - poolBefore;
    return (received, fee);
  }

  function _feeConfig(
    uint16 transferFeeBps
  ) internal pure returns (FeePolicyStore.FeeConfig memory) {
    return FeePolicyStore.FeeConfig({
      finalityFeeUSDCents: 0,
      fastFinalityFeeUSDCents: 0,
      finalityTransferFeeBps: transferFeeBps,
      fastFinalityTransferFeeBps: transferFeeBps * 2
    });
  }

  /// @dev An enabled rate limit with almost no refill, so a test can drain the bucket within a single block.
  function _rateLimit(
    uint128 capacity
  ) internal pure returns (RateLimiter.Config memory) {
    return RateLimiter.Config({isEnabled: true, capacity: capacity, rate: 1});
  }

  function _feeConfigWithFlat(
    uint16 transferFeeBps,
    uint32 flatFeeUSDCents
  ) internal pure returns (FeePolicyStore.FeeConfig memory) {
    return FeePolicyStore.FeeConfig({
      finalityFeeUSDCents: flatFeeUSDCents,
      fastFinalityFeeUSDCents: flatFeeUSDCents,
      finalityTransferFeeBps: transferFeeBps,
      fastFinalityTransferFeeBps: transferFeeBps
    });
  }

  function _removeInlineFee(
    PoolConfigManager pcm,
    address token,
    uint64 remoteChainSelector
  ) internal {
    pcm.applyInlineFeeConfigUpdates(_lane(token, remoteChainSelector), new FeePolicyStore.InlineFeeConfigArgs[](0));
  }

  function _removeFeePolicyRecord(
    PoolConfigManager pcm,
    bytes32 policyName,
    uint64 remoteChainSelector
  ) internal {
    pcm.applyFeePolicyConfigUpdates(
      _policyRecordKeysOf(policyName, remoteChainSelector),
      new bytes32[](0),
      new FeePolicyStore.FeePolicyConfigArgs[](0)
    );
  }

  function _policyRecordKeysOf(
    bytes32 policyName,
    uint64 remoteChainSelector
  ) internal pure returns (PolicyConfig.PolicyRecordKey[] memory keys) {
    keys = new PolicyConfig.PolicyRecordKey[](1);
    keys[0] = PolicyConfig.PolicyRecordKey({policyName: policyName, remoteChainSelector: remoteChainSelector});
    return keys;
  }

  function _twoAddresses(
    address a,
    address b
  ) internal pure returns (address[] memory arr) {
    arr = new address[](2);
    arr[0] = a;
    arr[1] = b;
    return arr;
  }

  /// @dev Sets a token profile with an explicit lock or burn address. The BurnToAddress flavor needs a nonzero
  /// address, and the store rejects a zero one for that flavor.
  function _setProfileFull(
    PoolConfigManager pcm,
    address token,
    uint8 decimals,
    IPoolConfigManager.LockOrBurnFlavor flavor,
    address lockOrBurnAddress
  ) internal {
    PolicyConfig.TokenProfileArgs[] memory profiles = new PolicyConfig.TokenProfileArgs[](1);
    profiles[0] = PolicyConfig.TokenProfileArgs({
      localToken: token,
      localDecimals: decimals,
      lockOrBurnFlavor: flavor,
      allowedFinalityConfig: FinalityCodec.WAIT_FOR_FINALITY_FLAG,
      lockOrBurnAddress: lockOrBurnAddress,
      additionalPoolDataProvider: address(0)
    });
    pcm.applyTokenProfileUpdates(new address[](0), profiles);
  }

  /// @dev Delivers source pool data to the destination pool. The local token is a parameter so tests can deliver to
  /// the wrong token on purpose.
  function _deliver(
    MttpChain storage dst,
    MttpChain storage src,
    address localToken,
    uint256 sourceAmount,
    bytes memory sourcePoolData,
    bytes4 finality
  ) internal returns (uint256 destAmount) {
    return
      _deliverFromPool(dst, src.localSelector, address(src.pool), localToken, sourceAmount, sourcePoolData, finality);
  }

  /// @dev Same as _deliver, but with an explicit source pool. Used when the source is a legacy single-token pool
  /// instead of an MTTP.
  function _deliverFromPool(
    MttpChain storage dst,
    uint64 sourceChainSelector,
    address sourcePool,
    address localToken,
    uint256 sourceAmount,
    bytes memory sourcePoolData,
    bytes4 finality
  ) internal returns (uint256 destAmount) {
    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = Pool.ReleaseOrMintInV1({
      originalSender: abi.encode(s_allowedSender),
      remoteChainSelector: sourceChainSelector,
      receiver: s_receiver,
      sourceDenominatedAmount: sourceAmount,
      localToken: localToken,
      sourcePoolAddress: abi.encode(sourcePool),
      sourcePoolData: sourcePoolData,
      offchainTokenData: ""
    });

    vm.prank(dst.offRamp);
    Pool.ReleaseOrMintOutV1 memory out = dst.pool.releaseOrMint(releaseOrMintIn, finality);
    return out.destinationAmount;
  }

  /// @dev A full transfer: send on the source and deliver to the matching token on the destination.
  function _transfer(
    MttpChain storage src,
    MttpChain storage dst,
    uint256 tokenIndex,
    uint256 amount
  ) internal returns (uint256 destAmount) {
    (bytes memory poolData, uint256 postFeeAmount, bytes memory destTokenAddress) =
      _send(src, dst, tokenIndex, amount, FinalityCodec.WAIT_FOR_FINALITY_FLAG, s_allowedSender);
    // Deliver to the token the source pool chose, so a wrong remote-token mapping fails here.
    return _deliver(
      dst, src, abi.decode(destTokenAddress, (address)), postFeeAmount, poolData, FinalityCodec.WAIT_FOR_FINALITY_FLAG
    );
  }

  /// @dev Approval does not move funds. The harness pulls tokens inside the same call as lockOrBurn, so a
  /// rejection restores the sender's balance, the pool balance and the allowance.
  function _approvePool(
    MttpChain storage c,
    address token,
    uint256 amount
  ) internal {
    _approvePool(c, token, amount, s_allowedSender);
  }

  function _approvePool(
    MttpChain storage c,
    address token,
    uint256 amount,
    address sender
  ) internal {
    vm.prank(sender);
    IERC20(token).approve(address(c.harness), amount);
  }

  function _fundAddress(
    address token,
    address recipient,
    uint256 amount
  ) internal {
    vm.prank(s_allowedSender);
    IERC20(token).transfer(recipient, amount);
  }

  function _balance(
    address token,
    address account
  ) internal view returns (uint256) {
    return IERC20(token).balanceOf(account);
  }

  // ================================================================
  // │                    Configuration helpers                     │
  // ================================================================

  function _setProfile(
    PoolConfigManager pcm,
    address token,
    uint8 decimals,
    IPoolConfigManager.LockOrBurnFlavor flavor
  ) internal {
    PolicyConfig.TokenProfileArgs[] memory profiles = new PolicyConfig.TokenProfileArgs[](1);
    profiles[0] = PolicyConfig.TokenProfileArgs({
      localToken: token,
      localDecimals: decimals,
      lockOrBurnFlavor: flavor,
      allowedFinalityConfig: FinalityCodec.WAIT_FOR_FINALITY_FLAG,
      lockOrBurnAddress: address(0),
      additionalPoolDataProvider: address(0)
    });
    pcm.applyTokenProfileUpdates(new address[](0), profiles);
  }

  function _assignPolicy(
    PoolConfigManager pcm,
    address token,
    bytes32 connectionPolicy,
    bytes32 remotePoolPolicy
  ) internal {
    if (connectionPolicy != bytes32(0)) {
      pcm.applyTokenConnectionPolicyUpdates(_tokenPolicyOf(token, connectionPolicy));
    }
    if (remotePoolPolicy != bytes32(0)) {
      pcm.applyTokenRemotePoolPolicyUpdates(_tokenPolicyOf(token, remotePoolPolicy));
    }
  }

  function _assignPolicy(
    PoolConfigManager pcm,
    address token,
    bytes32 rateLimitPolicy,
    bytes32 feePolicy,
    bytes32 destOverheadPolicy,
    bytes32 securityPolicy
  ) internal {
    if (rateLimitPolicy != bytes32(0)) pcm.applyTokenRateLimitPolicyUpdates(_tokenPolicyOf(token, rateLimitPolicy));
    if (feePolicy != bytes32(0)) pcm.applyTokenFeePolicyUpdates(_tokenPolicyOf(token, feePolicy));
    if (destOverheadPolicy != bytes32(0)) {
      pcm.applyTokenDestOverheadPolicyUpdates(_tokenPolicyOf(token, destOverheadPolicy));
    }
    if (securityPolicy != bytes32(0)) pcm.applyTokenSecurityPolicyUpdates(_tokenPolicyOf(token, securityPolicy));
  }

  function _tokenPolicyOf(
    address token,
    bytes32 policyName
  ) internal pure returns (PolicyConfig.TokenPolicyArgs[] memory tokenPolicies) {
    tokenPolicies = new PolicyConfig.TokenPolicyArgs[](1);
    tokenPolicies[0] = PolicyConfig.TokenPolicyArgs({localToken: token, policyName: policyName});
    return tokenPolicies;
  }

  function _setConnectionPolicyConfig(
    PoolConfigManager pcm,
    bytes32 policyName,
    uint64 remoteChainSelector,
    bool onlySubsetTokens
  ) internal {
    ConnectionPolicyStore.ConnectionPolicyConfigArgs[] memory configs =
      new ConnectionPolicyStore.ConnectionPolicyConfigArgs[](1);
    configs[0] = ConnectionPolicyStore.ConnectionPolicyConfigArgs({
      policyName: policyName, remoteChainSelector: remoteChainSelector, onlySubsetTokens: onlySubsetTokens
    });
    pcm.applyConnectionPolicyConfigUpdates(new PolicyConfig.PolicyRecordKey[](0), new bytes32[](0), configs);
  }

  function _setRemoteTokenProvider(
    PoolConfigManager pcm,
    bytes32 policyName,
    address provider
  ) internal {
    ConnectionPolicyStore.RemoteTokenAddressProviderArgs[] memory providers =
      new ConnectionPolicyStore.RemoteTokenAddressProviderArgs[](1);
    providers[0] = ConnectionPolicyStore.RemoteTokenAddressProviderArgs({
      policyName: policyName, remoteTokenAddressProvider: provider
    });
    pcm.applyRemoteTokenAddressProviderUpdates(providers);
  }

  function _setInlineConnectionConfig(
    PoolConfigManager pcm,
    address token,
    uint64 remoteChainSelector,
    bytes memory remoteToken
  ) internal {
    ConnectionPolicyStore.InlineConnectionConfigArgs[] memory configs =
      new ConnectionPolicyStore.InlineConnectionConfigArgs[](1);
    configs[0] = ConnectionPolicyStore.InlineConnectionConfigArgs({
      localToken: token, remoteChainSelector: remoteChainSelector, remoteToken: remoteToken
    });
    pcm.applyInlineConnectionConfigUpdates(new IPoolConfigManager.LaneKey[](0), configs);
  }

  function _setRemotePoolPolicyConfig(
    PoolConfigManager pcm,
    bytes32 policyName,
    uint64 remoteChainSelector,
    bytes[] memory remotePoolAddresses
  ) internal {
    ConnectionPolicyStore.RemotePoolPolicyConfigArgs[] memory configs =
      new ConnectionPolicyStore.RemotePoolPolicyConfigArgs[](1);
    configs[0] = ConnectionPolicyStore.RemotePoolPolicyConfigArgs({
      policyName: policyName, remoteChainSelector: remoteChainSelector, remotePoolAddresses: remotePoolAddresses
    });
    pcm.applyRemotePoolPolicyConfigUpdates(new PolicyConfig.PolicyRecordKey[](0), new bytes32[](0), configs);
  }

  function _setInlineRemotePoolConfig(
    PoolConfigManager pcm,
    address token,
    uint64 remoteChainSelector,
    bool trustPolicyRemotePools,
    ConnectionPolicyStore.RemotePoolConfig memory config
  ) internal {
    ConnectionPolicyStore.InlineRemotePoolConfigArgs[] memory configs =
      new ConnectionPolicyStore.InlineRemotePoolConfigArgs[](1);
    configs[0] = ConnectionPolicyStore.InlineRemotePoolConfigArgs({
      localToken: token,
      remoteChainSelector: remoteChainSelector,
      trustPolicyRemotePools: trustPolicyRemotePools,
      remotePoolConfig: config
    });
    pcm.applyInlineRemotePoolConfigUpdates(new IPoolConfigManager.LaneKey[](0), configs);
  }

  function _setRateLimitPolicyConfig(
    PoolConfigManager pcm,
    bytes32 policyName,
    uint64 remoteChainSelector,
    RateLimiter.Config memory config
  ) internal {
    _setRateLimitPolicyConfigWithDecimals(pcm, policyName, remoteChainSelector, config, T1_DECIMALS);
  }

  function _setRateLimitPolicyConfigWithDecimals(
    PoolConfigManager pcm,
    bytes32 policyName,
    uint64 remoteChainSelector,
    RateLimiter.Config memory config,
    uint8 localDecimals
  ) internal {
    SecurityPolicyStore.RateLimitPolicyConfigArgs[] memory configs =
      new SecurityPolicyStore.RateLimitPolicyConfigArgs[](1);
    configs[0] = SecurityPolicyStore.RateLimitPolicyConfigArgs({
      policyName: policyName,
      remoteChainSelector: remoteChainSelector,
      localDecimals: localDecimals,
      rateLimitConfig: SecurityPolicyStore.RateLimitConfig({
        outbound: config, inbound: config, fastFinalityOutbound: config, fastFinalityInbound: config
      })
    });
    pcm.applyRateLimitPolicyConfigUpdates(new PolicyConfig.PolicyRecordKey[](0), new bytes32[](0), configs);
  }

  function _setInlineRateLimitConfig(
    PoolConfigManager pcm,
    address token,
    uint64 remoteChainSelector,
    RateLimiter.Config memory config
  ) internal {
    SecurityPolicyStore.InlineRateLimitConfigArgs[] memory configs =
      new SecurityPolicyStore.InlineRateLimitConfigArgs[](1);
    configs[0] = SecurityPolicyStore.InlineRateLimitConfigArgs({
      localToken: token,
      remoteChainSelector: remoteChainSelector,
      rateLimitConfig: SecurityPolicyStore.RateLimitConfig({
        outbound: config, inbound: config, fastFinalityOutbound: config, fastFinalityInbound: config
      })
    });
    pcm.applyInlineRateLimitConfigUpdates(new IPoolConfigManager.LaneKey[](0), configs);
  }

  function _setFeePolicyConfig(
    PoolConfigManager pcm,
    bytes32 policyName,
    uint64 remoteChainSelector,
    FeePolicyStore.FeeConfig memory feeConfig
  ) internal {
    FeePolicyStore.FeePolicyConfigArgs[] memory configs = new FeePolicyStore.FeePolicyConfigArgs[](1);
    configs[0] = FeePolicyStore.FeePolicyConfigArgs({
      policyName: policyName, remoteChainSelector: remoteChainSelector, feeConfig: feeConfig
    });
    pcm.applyFeePolicyConfigUpdates(new PolicyConfig.PolicyRecordKey[](0), new bytes32[](0), configs);
  }

  function _setInlineFeeConfig(
    PoolConfigManager pcm,
    address token,
    uint64 remoteChainSelector,
    FeePolicyStore.FeeConfig memory feeConfig
  ) internal {
    FeePolicyStore.InlineFeeConfigArgs[] memory configs = new FeePolicyStore.InlineFeeConfigArgs[](1);
    configs[0] = FeePolicyStore.InlineFeeConfigArgs({
      localToken: token, remoteChainSelector: remoteChainSelector, feeConfig: feeConfig
    });
    pcm.applyInlineFeeConfigUpdates(new IPoolConfigManager.LaneKey[](0), configs);
  }

  function _setDestOverheadPolicyConfig(
    PoolConfigManager pcm,
    bytes32 policyName,
    uint64 remoteChainSelector,
    uint32 destGasOverhead,
    uint32 destBytesOverhead
  ) internal {
    FeePolicyStore.DestOverheadPolicyConfigArgs[] memory configs = new FeePolicyStore.DestOverheadPolicyConfigArgs[](1);
    configs[0] = FeePolicyStore.DestOverheadPolicyConfigArgs({
      policyName: policyName,
      remoteChainSelector: remoteChainSelector,
      destOverheadConfig: FeePolicyStore.DestOverheadConfig({
        destGasOverhead: destGasOverhead, destBytesOverhead: destBytesOverhead
      })
    });
    pcm.applyDestOverheadPolicyConfigUpdates(new PolicyConfig.PolicyRecordKey[](0), new bytes32[](0), configs);
  }

  function _setInlineDestOverheadConfig(
    PoolConfigManager pcm,
    address token,
    uint64 remoteChainSelector,
    uint32 destGasOverhead,
    uint32 destBytesOverhead
  ) internal {
    FeePolicyStore.InlineDestOverheadConfigArgs[] memory configs = new FeePolicyStore.InlineDestOverheadConfigArgs[](1);
    configs[0] = FeePolicyStore.InlineDestOverheadConfigArgs({
      localToken: token,
      remoteChainSelector: remoteChainSelector,
      destOverheadConfig: FeePolicyStore.DestOverheadConfig({
        destGasOverhead: destGasOverhead, destBytesOverhead: destBytesOverhead
      })
    });
    pcm.applyInlineDestOverheadConfigUpdates(new IPoolConfigManager.LaneKey[](0), configs);
  }

  function _setSecurityPolicyConfig(
    PoolConfigManager pcm,
    bytes32 policyName,
    uint64 remoteChainSelector,
    address[] memory outboundCCVs,
    address[] memory inboundCCVs,
    address hook
  ) internal {
    SecurityPolicyStore.SecurityPolicyConfigArgs[] memory configs =
      new SecurityPolicyStore.SecurityPolicyConfigArgs[](1);
    configs[0] = SecurityPolicyStore.SecurityPolicyConfigArgs({
      policyName: policyName,
      remoteChainSelector: remoteChainSelector,
      securityConfig: SecurityPolicyStore.SecurityConfig({
        outboundCCVs: outboundCCVs, inboundCCVs: inboundCCVs, hook: hook
      })
    });
    pcm.applySecurityPolicyConfigUpdates(new PolicyConfig.PolicyRecordKey[](0), new bytes32[](0), configs);
  }

  function _setInlineSecurityConfig(
    PoolConfigManager pcm,
    address token,
    uint64 remoteChainSelector,
    address[] memory outboundCCVs,
    address[] memory inboundCCVs,
    address hook
  ) internal {
    SecurityPolicyStore.InlineSecurityConfigArgs[] memory configs =
      new SecurityPolicyStore.InlineSecurityConfigArgs[](1);
    configs[0] = SecurityPolicyStore.InlineSecurityConfigArgs({
      localToken: token,
      remoteChainSelector: remoteChainSelector,
      securityConfig: SecurityPolicyStore.SecurityConfig({
        outboundCCVs: outboundCCVs, inboundCCVs: inboundCCVs, hook: hook
      })
    });
    pcm.applyInlineSecurityConfigUpdates(new IPoolConfigManager.LaneKey[](0), configs);
  }

  function _addSubsetTokens(
    PoolConfigManager pcm,
    bytes32 policyName,
    uint64 remoteChainSelector,
    address[] memory tokens
  ) internal {
    ConnectionPolicyStore.SubsetTokenArgs[] memory adds = new ConnectionPolicyStore.SubsetTokenArgs[](tokens.length);
    for (uint256 i = 0; i < tokens.length; ++i) {
      adds[i] = ConnectionPolicyStore.SubsetTokenArgs({
        policyName: policyName, remoteChainSelector: remoteChainSelector, localToken: tokens[i]
      });
    }
    pcm.applySubsetTokenUpdates(new ConnectionPolicyStore.SubsetTokenArgs[](0), adds);
  }

  // ================================================================
  // │                    Migration helpers                         │
  // ================================================================

  /// @dev Deploys a legacy BurnMintTokenPool for one token on `c`, registers it in the token admin registry, and
  /// connects it to `remote`'s MTTP. Call as OWNER.
  function _deployLegacyPoolOn(
    MttpChain storage c,
    MttpChain storage remote,
    uint256 tokenIndex
  ) internal returns (address legacyPool) {
    BurnMintTokenPool pool = new BurnMintTokenPool(
      IBurnMintERC20(c.tokens[tokenIndex]),
      _decimalsOf(tokenIndex),
      address(0),
      address(s_mockRMNRemote),
      address(c.router)
    );
    legacyPool = address(pool);
    ICrossChainTokenLike(c.tokens[tokenIndex]).grantMintAndBurnRoles(legacyPool);
    c.tar.setPool(c.tokens[tokenIndex], legacyPool);

    RateLimiter.Config memory disabled;
    bytes[] memory remotePools = new bytes[](1);
    remotePools[0] = abi.encode(address(remote.pool));
    TokenPool.ChainUpdate[] memory updates = new TokenPool.ChainUpdate[](1);
    updates[0] = TokenPool.ChainUpdate({
      remoteChainSelector: remote.localSelector,
      remotePoolAddresses: remotePools,
      remoteTokenAddress: abi.encode(remote.tokens[tokenIndex]),
      outboundRateLimiterConfig: disabled,
      inboundRateLimiterConfig: disabled
    });
    pool.applyChainUpdates(new uint64[](0), updates);
    return legacyPool;
  }

  /// @dev Points the token admin registry at a different pool for a token. Call as the token admin.
  function _switchPool(
    MttpChain storage c,
    uint256 tokenIndex,
    address newPool
  ) internal {
    c.tar.setPool(c.tokens[tokenIndex], newPool);
  }

  /// @dev Sets an inline remote pool record with explicit legacy flags. This is how the migration tests control the
  /// pool data format in each direction.
  function _setMigrationRemotePool(
    MttpChain storage c,
    uint256 tokenIndex,
    uint64 remoteChainSelector,
    address remotePool,
    bool useLegacyOnSend,
    bool allowLegacyOnReceive
  ) internal {
    IPoolConfigManager.RemotePool[] memory pools = new IPoolConfigManager.RemotePool[](1);
    pools[0] = IPoolConfigManager.RemotePool({
      remotePoolAddress: abi.encode(remotePool), allowLegacyPoolDataOnReceive: allowLegacyOnReceive
    });
    _setMigrationRemotePools(c, tokenIndex, remoteChainSelector, useLegacyOnSend, pools);
  }

  /// @dev Sets an inline remote pool record from an explicit pool list, so a migration can keep the old pool trusted
  /// while it adds the new one.
  function _setMigrationRemotePools(
    MttpChain storage c,
    uint256 tokenIndex,
    uint64 remoteChainSelector,
    bool useLegacyOnSend,
    IPoolConfigManager.RemotePool[] memory pools
  ) internal {
    ConnectionPolicyStore.RemotePoolConfig memory config =
      ConnectionPolicyStore.RemotePoolConfig({useLegacyPoolDataOnSend: useLegacyOnSend, remotePools: pools});
    _setInlineRemotePoolConfig(c.pcm, c.tokens[tokenIndex], remoteChainSelector, false, config);
  }

  /// @dev A remote-pool list that keeps the legacy pool trusted for the short-format cohort and adds the MTTP for
  /// long-format sends.
  function _legacyThenMttpPools(
    address legacyPool,
    address mttp
  ) internal pure returns (IPoolConfigManager.RemotePool[] memory pools) {
    pools = new IPoolConfigManager.RemotePool[](2);
    pools[0] =
      IPoolConfigManager.RemotePool({remotePoolAddress: abi.encode(legacyPool), allowLegacyPoolDataOnReceive: true});
    pools[1] = IPoolConfigManager.RemotePool({remotePoolAddress: abi.encode(mttp), allowLegacyPoolDataOnReceive: false});
    return pools;
  }

  /// @dev Revokes a retired pool's mint and burn roles so it can no longer serve its token. Call as OWNER.
  function _revokeMintAndBurnRoles(
    address token,
    address pool
  ) internal {
    ICrossChainTokenLike(token).revokeRole(ICrossChainTokenLike(token).MINTER_ROLE(), pool);
    ICrossChainTokenLike(token).revokeRole(ICrossChainTokenLike(token).BURNER_ROLE(), pool);
  }

  /// @dev Calls lockOrBurn on a legacy pool and returns the 32-byte legacy pool data it produces.
  function _legacyLockOrBurn(
    address legacyPool,
    MttpChain storage c,
    MttpChain storage remote,
    uint256 tokenIndex,
    uint256 amount
  ) internal returns (bytes memory poolData, uint256 destAmount) {
    _fundAddress(c.tokens[tokenIndex], legacyPool, amount);
    Pool.LockOrBurnInV1 memory lockOrBurnIn = Pool.LockOrBurnInV1({
      receiver: abi.encode(s_receiver),
      remoteChainSelector: remote.localSelector,
      originalSender: s_allowedSender,
      amount: amount,
      localToken: c.tokens[tokenIndex]
    });
    vm.prank(c.onRamp);
    (Pool.LockOrBurnOutV1 memory out, uint256 postFeeAmount) =
      BurnMintTokenPool(legacyPool).lockOrBurn(lockOrBurnIn, FinalityCodec.WAIT_FOR_FINALITY_FLAG, "");
    return (out.destPoolData, postFeeAmount);
  }

  /// @dev Calls releaseOrMint on a legacy pool.
  function _legacyReleaseOrMint(
    address legacyPool,
    MttpChain storage c,
    uint64 sourceChainSelector,
    address sourcePool,
    uint256 tokenIndex,
    uint256 sourceAmount,
    bytes memory sourcePoolData
  ) internal returns (uint256 destAmount) {
    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = Pool.ReleaseOrMintInV1({
      originalSender: abi.encode(s_allowedSender),
      remoteChainSelector: sourceChainSelector,
      receiver: s_receiver,
      sourceDenominatedAmount: sourceAmount,
      localToken: c.tokens[tokenIndex],
      sourcePoolAddress: abi.encode(sourcePool),
      sourcePoolData: sourcePoolData,
      offchainTokenData: ""
    });
    vm.prank(c.offRamp);
    Pool.ReleaseOrMintOutV1 memory out =
      BurnMintTokenPool(legacyPool).releaseOrMint(releaseOrMintIn, FinalityCodec.WAIT_FOR_FINALITY_FLAG);
    return out.destinationAmount;
  }

  // ================================================================
  // │                         View helpers                         │
  // ================================================================

  function _lane(
    address token,
    uint64 remoteChainSelector
  ) internal pure returns (IPoolConfigManager.LaneKey[] memory lanes) {
    lanes = new IPoolConfigManager.LaneKey[](1);
    lanes[0] = IPoolConfigManager.LaneKey({localToken: token, remoteChainSelector: remoteChainSelector});
    return lanes;
  }

  function _fee(
    PoolConfigManager pcm,
    address token,
    uint64 remoteChainSelector
  ) internal view returns (IPoolConfigManager.LaneFee memory) {
    return pcm.getFeeBatch(_lane(token, remoteChainSelector), FinalityCodec.WAIT_FOR_FINALITY_FLAG)[0];
  }

  function _getRateLimiterState(
    PoolConfigManager pcm,
    address token,
    uint64 remoteChainSelector
  ) internal view returns (IPoolConfigManager.LaneRateLimiterState memory) {
    return pcm.getCurrentRateLimiterStateBatch(_lane(token, remoteChainSelector))[0];
  }

  /// @dev Delivers with a fixed gas limit and returns whether it succeeded instead of reverting, so a test can model
  /// a low-gas execution attempt without failing the whole test.
  function _tryDeliver(
    MttpChain storage dst,
    uint64 sourceChainSelector,
    address sourcePool,
    address localToken,
    uint256 sourceAmount,
    bytes memory sourcePoolData,
    bytes4 finality,
    uint256 gasLimit
  ) internal returns (bool success) {
    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = Pool.ReleaseOrMintInV1({
      originalSender: abi.encode(s_allowedSender),
      remoteChainSelector: sourceChainSelector,
      receiver: s_receiver,
      sourceDenominatedAmount: sourceAmount,
      localToken: localToken,
      sourcePoolAddress: abi.encode(sourcePool),
      sourcePoolData: sourcePoolData,
      offchainTokenData: ""
    });

    vm.prank(dst.offRamp);
    (success,) =
      address(dst.pool).call{gas: gasLimit}(abi.encodeCall(IPoolV2.releaseOrMint, (releaseOrMintIn, finality)));
    return success;
  }

  function _removeInlineRateLimit(
    PoolConfigManager pcm,
    address token,
    uint64 remoteChainSelector
  ) internal {
    pcm.applyInlineRateLimitConfigUpdates(
      _lane(token, remoteChainSelector), new SecurityPolicyStore.InlineRateLimitConfigArgs[](0)
    );
  }
}

/// @notice Stands in for the Router: moves the amount into the pool and calls lockOrBurn in one transaction, so a
/// rejected send rolls the funding back with it.
contract MttpRouterHarness {
  using SafeERC20 for IERC20;

  function forward(
    address token,
    address pool,
    Pool.LockOrBurnInV1 calldata lockOrBurnIn,
    bytes4 finality,
    bytes calldata tokenArgs
  ) external returns (bytes memory destPoolData, uint256 destAmount, bytes memory destTokenAddress) {
    IERC20(token).safeTransferFrom(lockOrBurnIn.originalSender, pool, lockOrBurnIn.amount);
    (Pool.LockOrBurnOutV1 memory out, uint256 postFeeAmount) =
      BurnMintMultiTokenPool(pool).lockOrBurn(lockOrBurnIn, finality, tokenArgs);
    return (out.destPoolData, postFeeAmount, out.destTokenAddress);
  }
}

/// @dev The CrossChainToken functions these tests call. The full token interface pulls in a lot of extra imports.
interface ICrossChainTokenLike {
  function grantMintAndBurnRoles(
    address burnMintPool
  ) external;

  function revokeRole(
    bytes32 role,
    address account
  ) external;

  function MINTER_ROLE() external view returns (bytes32);

  function BURNER_ROLE() external view returns (bytes32);

  function mint(
    address account,
    uint256 amount
  ) external;
}
