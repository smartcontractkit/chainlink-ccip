// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IOwner} from "../../../interfaces/IOwner.sol";

import {Router} from "../../../Router.sol";
import {RateLimiter} from "../../../libraries/RateLimiter.sol";
import {BurnFromMintTokenPool} from "../../../pools/BurnFromMintTokenPool.sol";
import {BurnMintTokenPool} from "../../../pools/BurnMintTokenPool.sol";
import {ERC20LockBox} from "../../../pools/ERC20LockBox.sol";
import {LockReleaseTokenPool} from "../../../pools/LockReleaseTokenPool.sol";
import {TokenPool} from "../../../pools/TokenPool.sol";
import {BaseERC20} from "../../../tmp/BaseERC20.sol";
import {CrossChainToken} from "../../../tmp/CrossChainToken.sol";
import {RegistryModuleOwnerCustom} from "../../../tokenAdminRegistry/RegistryModuleOwnerCustom.sol";
import {TokenAdminRegistry} from "../../../tokenAdminRegistry/TokenAdminRegistry.sol";
import {TokenPoolFactory} from "../../../tokenAdminRegistry/TokenPoolFactory/TokenPoolFactory.sol";
import {TokenPoolFactorySetup} from "./TokenPoolFactorySetup.t.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

import {IERC20Metadata} from "@openzeppelin/contracts@5.3.0/token/ERC20/extensions/IERC20Metadata.sol";
import {Create2} from "@openzeppelin/contracts@5.3.0/utils/Create2.sol";

contract TokenPoolFactory_deployTokenAndTokenPool is TokenPoolFactorySetup {
  using Create2 for bytes32;

  uint8 private constant LOCAL_TOKEN_DECIMALS = 18;
  uint8 private constant REMOTE_TOKEN_DECIMALS = 6;
  bytes private constant LOCK_RELEASE_INIT_CODE = type(LockReleaseTokenPool).creationCode;

  bytes32 internal constant DYNAMIC_SALT = keccak256(abi.encodePacked(FAKE_SALT, OWNER));

  address internal s_burnMintOffRamp = makeAddr("burn_mint_offRamp");

  function setUp() public override {
    super.setUp();

    Router.OffRamp[] memory offRampUpdates = new Router.OffRamp[](1);
    offRampUpdates[0] = Router.OffRamp({sourceChainSelector: DEST_CHAIN_SELECTOR, offRamp: s_burnMintOffRamp});
    s_sourceRouter.applyRampUpdates(new Router.OnRamp[](0), new Router.OffRamp[](0), offRampUpdates);

    vm.startPrank(OWNER);
  }

  function test_deployTokenAndTokenPool_WithNoExistingTokenOnRemoteChain() public {
    address predictedTokenAddress =
      Create2.computeAddress(DYNAMIC_SALT, keccak256(s_tokenInitCode), address(s_tokenPoolFactory));

    bytes memory poolCreationParams =
      abi.encode(predictedTokenAddress, LOCAL_TOKEN_DECIMALS, address(0), s_rmnProxy, s_sourceRouter);

    bytes memory predictedPoolInitCode = abi.encodePacked(POOL_INIT_CODE, poolCreationParams);

    address predictedPoolAddress =
      DYNAMIC_SALT.computeAddress(keccak256(predictedPoolInitCode), address(s_tokenPoolFactory));

    (address tokenAddress, address poolAddress) = s_tokenPoolFactory.deployTokenAndTokenPool(
      new TokenPoolFactory.RemoteTokenPoolInfo[](0),
      LOCAL_TOKEN_DECIMALS,
      TokenPoolFactory.PoolType.BURN_MINT,
      s_tokenInitCode,
      POOL_INIT_CODE,
      address(0),
      FAKE_SALT
    );

    assertNotEq(address(0), tokenAddress, "Token Address should not be 0");
    assertNotEq(address(0), poolAddress, "Pool Address should not be 0");

    assertEq(predictedTokenAddress, tokenAddress, "Token Address should have been predicted");
    assertEq(predictedPoolAddress, poolAddress, "Pool Address should have been predicted");

    s_tokenAdminRegistry.acceptAdminRole(tokenAddress);
    Ownable2Step(poolAddress).acceptOwnership();

    assertEq(poolAddress, s_tokenAdminRegistry.getPool(tokenAddress), "Token Pool should be set");
    assertEq(CrossChainToken(tokenAddress).owner(), OWNER, "Token should be owned by the owner");
    assertEq(IOwner(poolAddress).owner(), OWNER, "Pool should be owned by the owner");

    assertTrue(
      CrossChainToken(tokenAddress).hasRole(CrossChainToken(tokenAddress).MINTER_ROLE(), poolAddress),
      "pool should be minter"
    );
    assertTrue(
      CrossChainToken(tokenAddress).hasRole(CrossChainToken(tokenAddress).BURNER_ROLE(), poolAddress),
      "pool should be burner"
    );
  }

  function test_deployTokenAndTokenPool_WithNoExistingRemoteContracts_Predict() public {
    TokenAdminRegistry newTokenAdminRegistry = new TokenAdminRegistry();
    RegistryModuleOwnerCustom newRegistryModule = new RegistryModuleOwnerCustom(address(newTokenAdminRegistry));

    TokenPoolFactory newTokenPoolFactory =
      new TokenPoolFactory(newTokenAdminRegistry, newRegistryModule, s_rmnProxy, address(s_destRouter));

    newTokenAdminRegistry.addRegistryModule(address(newRegistryModule));

    bytes memory remoteTokenInitCode = _buildTokenInitCode(address(newTokenPoolFactory));

    TokenPoolFactory.RemoteChainConfig memory remoteChainConfig = TokenPoolFactory.RemoteChainConfig({
      remotePoolFactory: address(newTokenPoolFactory),
      remoteRouter: address(s_destRouter),
      remoteRMNProxy: address(s_rmnProxy),
      remoteLockBox: address(0),
      remoteTokenDecimals: LOCAL_TOKEN_DECIMALS
    });

    TokenPoolFactory.RemoteTokenPoolInfo[] memory remoteTokenPools = new TokenPoolFactory.RemoteTokenPoolInfo[](1);

    remoteTokenPools[0] = TokenPoolFactory.RemoteTokenPoolInfo(
      DEST_CHAIN_SELECTOR, // remoteChainSelector
      "", // remotePoolAddress
      type(BurnMintTokenPool).creationCode, // remotePoolInitCode
      remoteChainConfig, // remoteChainConfig
      TokenPoolFactory.PoolType.BURN_MINT, // poolType
      "", // remoteTokenAddress
      remoteTokenInitCode, // remoteTokenInitCode
      RateLimiter.Config(false, 0, 0)
    );

    address predictedTokenAddress =
      DYNAMIC_SALT.computeAddress(keccak256(remoteTokenInitCode), address(newTokenPoolFactory));

    (address tokenAddress, address poolAddress) = s_tokenPoolFactory.deployTokenAndTokenPool(
      remoteTokenPools,
      LOCAL_TOKEN_DECIMALS,
      TokenPoolFactory.PoolType.BURN_MINT,
      s_tokenInitCode,
      POOL_INIT_CODE,
      address(0),
      FAKE_SALT
    );

    assertEq(
      abi.encode(predictedTokenAddress),
      TokenPool(poolAddress).getRemoteToken(DEST_CHAIN_SELECTOR),
      "Token Address should have been predicted"
    );

    {
      bytes memory predictedPoolCreationParams =
        abi.encode(predictedTokenAddress, LOCAL_TOKEN_DECIMALS, address(0), s_rmnProxy, address(s_destRouter));

      bytes memory predictedPoolInitCode = abi.encodePacked(POOL_INIT_CODE, predictedPoolCreationParams);

      address predictedPoolAddress =
        DYNAMIC_SALT.computeAddress(keccak256(predictedPoolInitCode), address(newTokenPoolFactory));

      assertEq(
        abi.encode(predictedPoolAddress),
        TokenPool(poolAddress).getRemotePools(DEST_CHAIN_SELECTOR)[0],
        "Pool Address should have been predicted"
      );
    }

    (address newTokenAddress, address newPoolAddress) = newTokenPoolFactory.deployTokenAndTokenPool(
      new TokenPoolFactory.RemoteTokenPoolInfo[](0),
      LOCAL_TOKEN_DECIMALS,
      TokenPoolFactory.PoolType.BURN_MINT,
      remoteTokenInitCode,
      POOL_INIT_CODE,
      address(0),
      FAKE_SALT
    );

    assertEq(
      TokenPool(poolAddress).getRemotePools(DEST_CHAIN_SELECTOR)[0],
      abi.encode(newPoolAddress),
      "New Pool Address should have been deployed correctly"
    );

    assertEq(
      TokenPool(poolAddress).getRemoteToken(DEST_CHAIN_SELECTOR),
      abi.encode(newTokenAddress),
      "New Token Address should have been deployed correctly"
    );

    assertTrue(
      CrossChainToken(tokenAddress).hasRole(CrossChainToken(tokenAddress).MINTER_ROLE(), poolAddress),
      "pool should be minter"
    );
    assertTrue(
      CrossChainToken(tokenAddress).hasRole(CrossChainToken(tokenAddress).BURNER_ROLE(), poolAddress),
      "pool should be burner"
    );

    assertEq(s_tokenAdminRegistry.getPool(tokenAddress), poolAddress, "Token Pool should be set");

    TokenAdminRegistry.TokenConfig memory tokenConfig = s_tokenAdminRegistry.getTokenConfig(tokenAddress);
    assertEq(tokenConfig.administrator, address(s_tokenPoolFactory), "Administrator should be set");
    assertEq(tokenConfig.pendingAdministrator, OWNER, "Pending Administrator should be OWNER");
    assertEq(tokenConfig.tokenPool, poolAddress, "Pool Address should be set");

    vm.startPrank(OWNER);
    s_tokenAdminRegistry.acceptAdminRole(tokenAddress);
    assertEq(s_tokenAdminRegistry.getTokenConfig(tokenAddress).administrator, OWNER, "Administrator should be set");
    assertEq(
      s_tokenAdminRegistry.getTokenConfig(tokenAddress).pendingAdministrator, address(0), "Pending should be cleared"
    );

    Ownable2Step(poolAddress).acceptOwnership();

    assertEq(CrossChainToken(tokenAddress).owner(), OWNER, "Token should be controlled by the OWNER");
    assertEq(IOwner(poolAddress).owner(), OWNER, "Pool should be controlled by the OWNER");
  }

  function test_deployTokenPoolWithExistingToken_ExistingRemoteToken_AndPredictPool() public {
    CrossChainToken newRemoteToken = new CrossChainToken(
      BaseERC20.ConstructorParams({
        name: "TestToken",
        symbol: "TT",
        decimals: LOCAL_TOKEN_DECIMALS,
        maxSupply: type(uint256).max,
        preMint: PREMINT_AMOUNT,
        ccipAdmin: OWNER
      }),
      address(0),
      OWNER
    );

    TokenAdminRegistry newTokenAdminRegistry = new TokenAdminRegistry();
    RegistryModuleOwnerCustom newRegistryModule = new RegistryModuleOwnerCustom(address(newTokenAdminRegistry));

    TokenPoolFactory newTokenPoolFactory =
      new TokenPoolFactory(newTokenAdminRegistry, newRegistryModule, s_rmnProxy, address(s_destRouter));

    newTokenAdminRegistry.addRegistryModule(address(newRegistryModule));

    TokenPoolFactory.RemoteChainConfig memory remoteChainConfig = TokenPoolFactory.RemoteChainConfig({
      remotePoolFactory: address(newTokenPoolFactory),
      remoteRouter: address(s_destRouter),
      remoteRMNProxy: address(s_rmnProxy),
      remoteLockBox: address(0),
      remoteTokenDecimals: LOCAL_TOKEN_DECIMALS
    });

    TokenPoolFactory.RemoteTokenPoolInfo[] memory remoteTokenPools = new TokenPoolFactory.RemoteTokenPoolInfo[](1);

    remoteTokenPools[0] = TokenPoolFactory.RemoteTokenPoolInfo(
      DEST_CHAIN_SELECTOR, // remoteChainSelector
      "", // remotePoolAddress
      type(BurnMintTokenPool).creationCode, // remotePoolInitCode
      remoteChainConfig, // remoteChainConfig
      TokenPoolFactory.PoolType.BURN_MINT, // poolType
      abi.encode(address(newRemoteToken)), // remoteTokenAddress
      s_tokenInitCode, // remoteTokenInitCode
      RateLimiter.Config(false, 0, 0) // rateLimiterConfig
    );

    (address tokenAddress, address poolAddress) = s_tokenPoolFactory.deployTokenAndTokenPool(
      remoteTokenPools,
      LOCAL_TOKEN_DECIMALS,
      TokenPoolFactory.PoolType.BURN_MINT,
      s_tokenInitCode,
      POOL_INIT_CODE,
      address(0),
      FAKE_SALT
    );

    assertEq(address(TokenPool(poolAddress).getToken()), tokenAddress, "Token Address should have been set locally");

    assertEq(
      abi.encode(address(newRemoteToken)),
      TokenPool(poolAddress).getRemoteToken(DEST_CHAIN_SELECTOR),
      "Token Address should have been predicted"
    );

    bytes memory predictedPoolCreationParams =
      abi.encode(address(newRemoteToken), LOCAL_TOKEN_DECIMALS, address(0), s_rmnProxy, address(s_destRouter));

    bytes memory predictedPoolInitCode = abi.encodePacked(POOL_INIT_CODE, predictedPoolCreationParams);

    address predictedPoolAddress =
      DYNAMIC_SALT.computeAddress(keccak256(predictedPoolInitCode), address(newTokenPoolFactory));

    assertEq(
      abi.encode(predictedPoolAddress),
      TokenPool(poolAddress).getRemotePools(DEST_CHAIN_SELECTOR)[0],
      "Pool Address should have been predicted"
    );

    address newPoolAddress = newTokenPoolFactory.deployTokenPoolWithExistingToken(
      address(newRemoteToken),
      LOCAL_TOKEN_DECIMALS,
      TokenPoolFactory.PoolType.BURN_MINT,
      new TokenPoolFactory.RemoteTokenPoolInfo[](0),
      POOL_INIT_CODE,
      address(0),
      FAKE_SALT
    );

    assertEq(
      abi.encode(newRemoteToken),
      TokenPool(poolAddress).getRemoteToken(DEST_CHAIN_SELECTOR),
      "Remote Token Address should have been set correctly"
    );

    assertEq(
      TokenPool(poolAddress).getRemotePools(DEST_CHAIN_SELECTOR)[0],
      abi.encode(newPoolAddress),
      "New Pool Address should have been deployed correctly"
    );
  }

  function test_deployTokenAndTokenPool_WithRemoteTokenAndRemotePool() public {
    bytes memory RANDOM_TOKEN_ADDRESS = abi.encode(makeAddr("RANDOM_TOKEN"));
    bytes memory RANDOM_POOL_ADDRESS = abi.encode(makeAddr("RANDOM_POOL"));

    TokenPoolFactory.RemoteTokenPoolInfo[] memory remoteTokenPools = new TokenPoolFactory.RemoteTokenPoolInfo[](1);

    remoteTokenPools[0] = TokenPoolFactory.RemoteTokenPoolInfo(
      DEST_CHAIN_SELECTOR, // remoteChainSelector
      RANDOM_POOL_ADDRESS, // remotePoolAddress
      type(BurnMintTokenPool).creationCode, // remotePoolInitCode
      TokenPoolFactory.RemoteChainConfig({
        remotePoolFactory: address(0),
        remoteRouter: address(0),
        remoteRMNProxy: address(0),
        remoteLockBox: address(0),
        remoteTokenDecimals: 0
      }), // remoteChainConfig
      TokenPoolFactory.PoolType.BURN_MINT, // poolType
      RANDOM_TOKEN_ADDRESS, // remoteTokenAddress
      "", // remoteTokenInitCode
      RateLimiter.Config(false, 0, 0) // rateLimiterConfig
    );

    (address tokenAddress, address poolAddress) = s_tokenPoolFactory.deployTokenAndTokenPool(
      remoteTokenPools,
      LOCAL_TOKEN_DECIMALS,
      TokenPoolFactory.PoolType.BURN_MINT,
      s_tokenInitCode,
      POOL_INIT_CODE,
      address(0),
      FAKE_SALT
    );

    assertNotEq(address(0), tokenAddress, "Token Address should not be 0");
    assertNotEq(address(0), poolAddress, "Pool Address should not be 0");

    s_tokenAdminRegistry.acceptAdminRole(tokenAddress);
    Ownable2Step(poolAddress).acceptOwnership();

    assertEq(
      TokenPool(poolAddress).getRemoteToken(DEST_CHAIN_SELECTOR),
      RANDOM_TOKEN_ADDRESS,
      "Remote Token Address should have been set"
    );

    assertEq(
      TokenPool(poolAddress).getRemotePools(DEST_CHAIN_SELECTOR)[0],
      RANDOM_POOL_ADDRESS,
      "Remote Pool Address should have been set"
    );

    assertEq(poolAddress, s_tokenAdminRegistry.getPool(tokenAddress), "Token Pool should be set");

    assertEq(CrossChainToken(tokenAddress).owner(), OWNER, "Token should be owned by the owner");

    assertEq(IOwner(poolAddress).owner(), OWNER, "Pool should be owned by the owner");
  }

  function test_deployTokenPoolWithExistingToken_LockRelease_UserLockBoxOwnershipPreserved() public {
    CrossChainToken token = new CrossChainToken(
      BaseERC20.ConstructorParams({
        name: "TestToken",
        symbol: "TEST",
        decimals: LOCAL_TOKEN_DECIMALS,
        maxSupply: type(uint256).max,
        preMint: PREMINT_AMOUNT,
        ccipAdmin: OWNER
      }),
      address(0),
      OWNER
    );
    ERC20LockBox userLockBox = new ERC20LockBox(address(token));

    address poolAddress = s_tokenPoolFactory.deployTokenPoolWithExistingToken(
      address(token),
      LOCAL_TOKEN_DECIMALS,
      TokenPoolFactory.PoolType.LOCK_RELEASE,
      new TokenPoolFactory.RemoteTokenPoolInfo[](0),
      LOCK_RELEASE_INIT_CODE,
      address(userLockBox),
      FAKE_SALT
    );

    assertEq(userLockBox.owner(), OWNER, "lockbox owner should remain user");

    Ownable2Step(poolAddress).acceptOwnership();

    assertEq(Ownable2Step(poolAddress).owner(), OWNER, "pool should be owned by owner");
  }

  function test_deployTokenAndTokenPool_LockRelease_AuthorizesPoolForLockBox() public {
    (address tokenAddress, address poolAddress) = s_tokenPoolFactory.deployTokenAndTokenPool(
      new TokenPoolFactory.RemoteTokenPoolInfo[](0),
      LOCAL_TOKEN_DECIMALS,
      TokenPoolFactory.PoolType.LOCK_RELEASE,
      s_tokenInitCode,
      LOCK_RELEASE_INIT_CODE,
      address(0),
      FAKE_SALT
    );

    bytes memory lockBoxCreationCode = abi.encodePacked(type(ERC20LockBox).creationCode, abi.encode(tokenAddress));
    address predictedLockBox = DYNAMIC_SALT.computeAddress(keccak256(lockBoxCreationCode), address(s_tokenPoolFactory));

    Ownable2Step(poolAddress).acceptOwnership();

    assertEq(
      AuthorizedCallers(predictedLockBox).getAllAuthorizedCallers()[0],
      poolAddress,
      "pool should be authorized caller on lockbox"
    );
  }

  function test_deployTokenPoolWithExistingToken_LockRelease_ExistingToken_Predict() public {
    TokenAdminRegistry newTokenAdminRegistry = new TokenAdminRegistry();
    RegistryModuleOwnerCustom newRegistryModule = new RegistryModuleOwnerCustom(address(newTokenAdminRegistry));

    TokenPoolFactory newTokenPoolFactory =
      new TokenPoolFactory(newTokenAdminRegistry, newRegistryModule, s_rmnProxy, address(s_destRouter));

    newTokenAdminRegistry.addRegistryModule(address(newRegistryModule));

    CrossChainToken newLocalToken = new CrossChainToken(
      BaseERC20.ConstructorParams({
        name: "TestToken",
        symbol: "TEST",
        decimals: LOCAL_TOKEN_DECIMALS,
        maxSupply: type(uint256).max,
        preMint: PREMINT_AMOUNT,
        ccipAdmin: OWNER
      }),
      address(0),
      OWNER
    );

    CrossChainToken newRemoteToken = new CrossChainToken(
      BaseERC20.ConstructorParams({
        name: "TestToken",
        symbol: "TEST",
        decimals: LOCAL_TOKEN_DECIMALS,
        maxSupply: type(uint256).max,
        preMint: PREMINT_AMOUNT,
        ccipAdmin: OWNER
      }),
      address(0),
      OWNER
    );

    ERC20LockBox localLockBox = new ERC20LockBox(address(newLocalToken));
    ERC20LockBox remoteLockBox = new ERC20LockBox(address(newRemoteToken));

    TokenPoolFactory.RemoteChainConfig memory remoteChainConfig = TokenPoolFactory.RemoteChainConfig({
      remotePoolFactory: address(newTokenPoolFactory),
      remoteRouter: address(s_destRouter),
      remoteRMNProxy: address(s_rmnProxy),
      remoteLockBox: address(remoteLockBox),
      remoteTokenDecimals: LOCAL_TOKEN_DECIMALS
    });

    TokenPoolFactory.RemoteTokenPoolInfo[] memory remoteTokenPools = new TokenPoolFactory.RemoteTokenPoolInfo[](1);

    remoteTokenPools[0] = TokenPoolFactory.RemoteTokenPoolInfo(
      DEST_CHAIN_SELECTOR, // remoteChainSelector
      "", // remotePoolAddress
      type(LockReleaseTokenPool).creationCode, // remotePoolInitCode
      remoteChainConfig, // remoteChainConfig
      TokenPoolFactory.PoolType.LOCK_RELEASE, // poolType
      abi.encode(address(newRemoteToken)), // remoteTokenAddress
      s_tokenInitCode, // remoteTokenInitCode
      RateLimiter.Config(false, 0, 0)
    );

    address poolAddress = s_tokenPoolFactory.deployTokenPoolWithExistingToken(
      address(newLocalToken),
      LOCAL_TOKEN_DECIMALS,
      TokenPoolFactory.PoolType.LOCK_RELEASE,
      remoteTokenPools,
      type(LockReleaseTokenPool).creationCode,
      address(localLockBox),
      FAKE_SALT
    );

    Ownable2Step(poolAddress).acceptOwnership();

    assertEq(
      address(LockReleaseTokenPool(poolAddress).getToken()),
      address(newLocalToken),
      "Token Address should have been set"
    );

    address newPoolAddress = newTokenPoolFactory.deployTokenPoolWithExistingToken(
      address(newRemoteToken),
      LOCAL_TOKEN_DECIMALS,
      TokenPoolFactory.PoolType.LOCK_RELEASE,
      new TokenPoolFactory.RemoteTokenPoolInfo[](0),
      type(LockReleaseTokenPool).creationCode,
      address(remoteLockBox),
      FAKE_SALT
    );

    address[] memory allowedCallers = new address[](1);
    allowedCallers[0] = poolAddress;
    localLockBox.applyAuthorizedCallerUpdates(
      AuthorizedCallers.AuthorizedCallerArgs({addedCallers: allowedCallers, removedCallers: new address[](0)})
    );

    allowedCallers[0] = newPoolAddress;
    remoteLockBox.applyAuthorizedCallerUpdates(
      AuthorizedCallers.AuthorizedCallerArgs({addedCallers: allowedCallers, removedCallers: new address[](0)})
    );

    assertEq(
      LockReleaseTokenPool(poolAddress).getRemotePools(DEST_CHAIN_SELECTOR)[0],
      abi.encode(newPoolAddress),
      "New Pool Address should have been deployed correctly"
    );

    assertEq(
      LockReleaseTokenPool(poolAddress).getRemoteToken(DEST_CHAIN_SELECTOR),
      abi.encode(address(newRemoteToken)),
      "New Token Address should have been deployed correctly"
    );

    assertEq(
      address(LockReleaseTokenPool(newPoolAddress).getToken()),
      address(newRemoteToken),
      "New Remote Token should be set correctly"
    );
  }

  function test_deployTokenAndTokenPool_BurnFromMintTokenPool() public {
    bytes memory RANDOM_TOKEN_ADDRESS = abi.encode(makeAddr("RANDOM_TOKEN"));
    bytes memory RANDOM_POOL_ADDRESS = abi.encode(makeAddr("RANDOM_POOL"));

    TokenPoolFactory.RemoteTokenPoolInfo[] memory remoteTokenPools = new TokenPoolFactory.RemoteTokenPoolInfo[](1);

    remoteTokenPools[0] = TokenPoolFactory.RemoteTokenPoolInfo(
      DEST_CHAIN_SELECTOR, // remoteChainSelector
      RANDOM_POOL_ADDRESS, // remotePoolAddress
      type(BurnFromMintTokenPool).creationCode, // remotePoolInitCode
      TokenPoolFactory.RemoteChainConfig({
        remotePoolFactory: address(0),
        remoteRouter: address(0),
        remoteRMNProxy: address(0),
        remoteLockBox: address(0),
        remoteTokenDecimals: 0
      }), // remoteChainConfig
      TokenPoolFactory.PoolType.BURN_MINT, // poolType
      RANDOM_TOKEN_ADDRESS, // remoteTokenAddress
      "", // remoteTokenInitCode
      RateLimiter.Config(false, 0, 0) // rateLimiterConfig
    );

    (address tokenAddress, address poolAddress) = s_tokenPoolFactory.deployTokenAndTokenPool(
      remoteTokenPools,
      LOCAL_TOKEN_DECIMALS,
      TokenPoolFactory.PoolType.BURN_MINT,
      s_tokenInitCode,
      POOL_INIT_CODE,
      address(0),
      FAKE_SALT
    );

    assertNotEq(address(0), tokenAddress, "Token Address should not be 0");
    assertNotEq(address(0), poolAddress, "Pool Address should not be 0");

    s_tokenAdminRegistry.acceptAdminRole(tokenAddress);
    Ownable2Step(poolAddress).acceptOwnership();

    assertEq(
      TokenPool(poolAddress).getRemoteToken(DEST_CHAIN_SELECTOR),
      RANDOM_TOKEN_ADDRESS,
      "Remote Token Address should have been set"
    );

    assertEq(
      TokenPool(poolAddress).getRemotePools(DEST_CHAIN_SELECTOR)[0],
      RANDOM_POOL_ADDRESS,
      "Remote Pool Address should have been set"
    );

    assertEq(poolAddress, s_tokenAdminRegistry.getPool(tokenAddress), "Token Pool should be set");

    assertEq(CrossChainToken(tokenAddress).owner(), OWNER, "Token should be owned by the owner");

    assertEq(IOwner(poolAddress).owner(), OWNER, "Pool should be owned by the owner");
  }

  function test_deployTokenAndTokenPool_RemoteTokenHasDifferentDecimals() public {
    CrossChainToken newRemoteToken = new CrossChainToken(
      BaseERC20.ConstructorParams({
        name: "TestToken",
        symbol: "TT",
        decimals: 6,
        maxSupply: type(uint256).max,
        preMint: PREMINT_AMOUNT,
        ccipAdmin: OWNER
      }),
      address(0),
      OWNER
    );

    TokenAdminRegistry newTokenAdminRegistry = new TokenAdminRegistry();
    RegistryModuleOwnerCustom newRegistryModule = new RegistryModuleOwnerCustom(address(newTokenAdminRegistry));

    TokenPoolFactory newTokenPoolFactory =
      new TokenPoolFactory(newTokenAdminRegistry, newRegistryModule, s_rmnProxy, address(s_destRouter));

    newTokenAdminRegistry.addRegistryModule(address(newRegistryModule));

    TokenPoolFactory.RemoteChainConfig memory remoteChainConfig = TokenPoolFactory.RemoteChainConfig({
      remotePoolFactory: address(newTokenPoolFactory),
      remoteRouter: address(s_destRouter),
      remoteRMNProxy: address(s_rmnProxy),
      remoteLockBox: address(0),
      remoteTokenDecimals: REMOTE_TOKEN_DECIMALS
    });

    TokenPoolFactory.RemoteTokenPoolInfo[] memory remoteTokenPools = new TokenPoolFactory.RemoteTokenPoolInfo[](1);

    remoteTokenPools[0] = TokenPoolFactory.RemoteTokenPoolInfo(
      DEST_CHAIN_SELECTOR, // remoteChainSelector
      "", // remotePoolAddress
      type(BurnMintTokenPool).creationCode, // remotePoolInitCode
      remoteChainConfig, // remoteChainConfig
      TokenPoolFactory.PoolType.BURN_MINT, // poolType
      abi.encode(address(newRemoteToken)), // remoteTokenAddress
      s_tokenInitCode, // remoteTokenInitCode
      RateLimiter.Config(false, 0, 0) // rateLimiterConfig
    );

    (address tokenAddress, address poolAddress) = s_tokenPoolFactory.deployTokenAndTokenPool(
      remoteTokenPools,
      LOCAL_TOKEN_DECIMALS,
      TokenPoolFactory.PoolType.BURN_MINT,
      s_tokenInitCode,
      POOL_INIT_CODE,
      address(0),
      FAKE_SALT
    );

    assertEq(address(TokenPool(poolAddress).getToken()), tokenAddress, "Token Address should have been set locally");

    assertEq(
      abi.encode(address(newRemoteToken)),
      TokenPool(poolAddress).getRemoteToken(DEST_CHAIN_SELECTOR),
      "Token Address should have been predicted"
    );

    bytes memory predictedPoolCreationParams =
      abi.encode(address(newRemoteToken), REMOTE_TOKEN_DECIMALS, address(0), s_rmnProxy, address(s_destRouter));

    bytes memory predictedPoolInitCode = abi.encodePacked(POOL_INIT_CODE, predictedPoolCreationParams);

    address predictedPoolAddress =
      DYNAMIC_SALT.computeAddress(keccak256(predictedPoolInitCode), address(newTokenPoolFactory));

    assertEq(
      abi.encode(predictedPoolAddress),
      TokenPool(poolAddress).getRemotePools(DEST_CHAIN_SELECTOR)[0],
      "Pool Address should have been predicted"
    );

    address newPoolAddress = newTokenPoolFactory.deployTokenPoolWithExistingToken(
      address(newRemoteToken),
      REMOTE_TOKEN_DECIMALS,
      TokenPoolFactory.PoolType.BURN_MINT,
      new TokenPoolFactory.RemoteTokenPoolInfo[](0),
      POOL_INIT_CODE,
      address(0),
      FAKE_SALT
    );

    assertEq(
      abi.encode(newRemoteToken),
      TokenPool(poolAddress).getRemoteToken(DEST_CHAIN_SELECTOR),
      "Remote Token Address should have been set correctly"
    );

    assertEq(
      TokenPool(poolAddress).getRemotePools(DEST_CHAIN_SELECTOR)[0],
      abi.encode(newPoolAddress),
      "New Pool Address should have been deployed correctly"
    );

    assertEq(TokenPool(poolAddress).getTokenDecimals(), LOCAL_TOKEN_DECIMALS, "Local token pool should use 18 decimals");

    assertEq(IERC20Metadata(tokenAddress).decimals(), LOCAL_TOKEN_DECIMALS, "Token Decimals should be 18");

    assertEq(TokenPool(newPoolAddress).getTokenDecimals(), REMOTE_TOKEN_DECIMALS, "Token Decimals should be 6");
    assertEq(address(TokenPool(newPoolAddress).getToken()), address(newRemoteToken), "Token Address should be set");
    assertEq(IERC20Metadata(address(newRemoteToken)).decimals(), REMOTE_TOKEN_DECIMALS, "Token Decimals should be 6");
  }

  function test_deployTokenPoolWithExistingToken_RevertWhen_InvalidLockBoxToken() public {
    CrossChainToken token = new CrossChainToken(
      BaseERC20.ConstructorParams({
        name: "TestToken",
        symbol: "TT",
        decimals: LOCAL_TOKEN_DECIMALS,
        maxSupply: type(uint256).max,
        preMint: PREMINT_AMOUNT,
        ccipAdmin: OWNER
      }),
      address(0),
      OWNER
    );
    CrossChainToken otherToken = new CrossChainToken(
      BaseERC20.ConstructorParams({
        name: "TestToken",
        symbol: "TT",
        decimals: LOCAL_TOKEN_DECIMALS,
        maxSupply: type(uint256).max,
        preMint: PREMINT_AMOUNT,
        ccipAdmin: OWNER
      }),
      address(0),
      OWNER
    );
    ERC20LockBox lockBox = new ERC20LockBox(address(otherToken));

    TokenPoolFactory.RemoteTokenPoolInfo[] memory remotes = new TokenPoolFactory.RemoteTokenPoolInfo[](0);
    vm.expectRevert(abi.encodeWithSelector(TokenPoolFactory.InvalidLockBoxToken.selector, address(token)));
    s_tokenPoolFactory.deployTokenPoolWithExistingToken(
      address(token),
      LOCAL_TOKEN_DECIMALS,
      TokenPoolFactory.PoolType.LOCK_RELEASE,
      remotes,
      LOCK_RELEASE_INIT_CODE,
      address(lockBox),
      FAKE_SALT
    );
  }
}
