// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ITokenAdminRegistry} from "../../interfaces/ITokenAdminRegistry.sol";
import {IOwnable} from "@chainlink/contracts/src/v0.8/shared/interfaces/IOwnable.sol";
import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";

import {RateLimiter} from "../../libraries/RateLimiter.sol";
import {ERC20LockBox} from "../../pools/ERC20LockBox.sol";
import {TokenPool} from "../../pools/TokenPool.sol";
import {RegistryModuleOwnerCustom} from "../RegistryModuleOwnerCustom.sol";
import {FactoryBurnMintERC20} from "./FactoryBurnMintERC20.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

import {Create2} from "@openzeppelin/contracts@5.3.0/utils/Create2.sol";

/// @notice A contract for deploying new tokens and token pools, and configuring them with the token admin registry.
/// @dev At the end of the transaction, the ownership transfer process will begin, but the user must accept the
/// ownership transfer in a separate transaction.
/// @dev The address prediction mechanism is only capable of deploying and predicting addresses for EVM based chains.
/// adding compatibility for other chains will require additional offchain computation.
contract TokenPoolFactory is ITypeAndVersion {
  using Create2 for bytes32;

  error InvalidZeroAddress();
  error InvalidLockBoxToken(address lockBoxToken, address poolToken);
  error InvalidLockBoxChainSelector(uint64 lockBoxSelector);

  /// @notice The type of pool to deploy. Types may be expanded in future versions.
  enum PoolType {
    BURN_MINT,
    LOCK_RELEASE
  }

  /// @dev This struct will only ever exist in memory and as calldata, and therefore does not need to be efficiently packed for storage. The struct is used to pass information to the create2 address generation function.
  struct RemoteTokenPoolInfo {
    uint64 remoteChainSelector; // The CCIP specific selector for the remote chain.
    bytes remotePoolAddress; // The address of the remote pool to either deploy or use as is. If empty, address
    // will be predicted.
    bytes remotePoolInitCode; // Remote pool creation code if it needs to be deployed, without constructor params
    // appended to the end.
    RemoteChainConfig remoteChainConfig; // The addresses of the remote RMNProxy, Router, factory, and token
    // decimals which are needed for determining the remote address.
    PoolType poolType; // The type of pool to deploy, either Burn/Mint or Lock/Release.
    bytes remoteTokenAddress; // EVM address for remote token. If empty, the address will be predicted.
    bytes remoteTokenInitCode; // The init code to be deployed on the remote chain and includes constructor params.
    RateLimiter.Config rateLimiterConfig; // Token Pool rate limit. Values will be applied on incoming an outgoing messages.
  }

  // solhint-disable-next-line gas-struct-packing
  struct RemoteChainConfig {
    address remotePoolFactory; // The factory contract on the remote chain which will make the deployment.
    address remoteRouter; // The router on the remote chain.
    address remoteRMNProxy; // The RMNProxy contract on the remote chain.
    address remoteLockBox; // The lockBox contract on the remote chain (for LOCK_RELEASE pools).
    uint8 remoteTokenDecimals; // The number of decimals for the token on the remote chain.
  }

  struct LocalPoolConfig {
    address token;
    uint8 localTokenDecimals;
    PoolType localPoolType;
    address lockBox;
    bytes32 salt;
  }

  string public constant typeAndVersion = "TokenPoolFactory 1.6.0-dev";
  bytes private constant LOCKBOX_INIT_CODE = type(ERC20LockBox).creationCode;
  address private immutable i_rmnProxy;

  ITokenAdminRegistry private immutable i_tokenAdminRegistry;
  RegistryModuleOwnerCustom private immutable i_registryModuleOwnerCustom;

  address private immutable i_ccipRouter;

  /// @notice Construct the TokenPoolFactory.
  /// @param tokenAdminRegistry The address of the token admin registry.
  /// @param tokenAdminModule The address of the token admin module which can register the token via ownership module.
  /// @param rmnProxy The address of the RMNProxy contract token pools will be deployed with.
  /// @param ccipRouter The address of the CCIPRouter contract token pools will be deployed with.
  constructor(
    ITokenAdminRegistry tokenAdminRegistry,
    RegistryModuleOwnerCustom tokenAdminModule,
    address rmnProxy,
    address ccipRouter
  ) {
    if (
      address(tokenAdminRegistry) == address(0) || address(tokenAdminModule) == address(0) || rmnProxy == address(0)
        || ccipRouter == address(0)
    ) revert InvalidZeroAddress();

    i_tokenAdminRegistry = ITokenAdminRegistry(tokenAdminRegistry);
    i_registryModuleOwnerCustom = RegistryModuleOwnerCustom(tokenAdminModule);
    i_rmnProxy = rmnProxy;
    i_ccipRouter = ccipRouter;
  }

  // ================================================================
  // │                   Top-Level Deployment                       │
  // ================================================================

  /// @notice Deploys a token and token pool with the given token information and configures it with remote token pools.
  /// @dev The token and token pool are deployed in the same transaction, and the token pool is configured with the
  /// remote token pools. The token pool is then set in the token admin registry. Ownership of the everything is transferred
  /// to the msg.sender, but must be accepted in a separate transaction due to 2-step ownership transfer.
  /// @param remoteTokenPools An array of remote token pools info to be used in the pool's applyChainUpdates function
  /// or to be predicted if the pool has not been deployed yet on the remote chain.
  /// @param localTokenDecimals The amount of decimals to be used in the new token. Since decimals() is not part of the
  /// the ERC20 standard, and thus cannot be certain to exist, the amount must be supplied via user input.
  /// @param localPoolType The type of pool to deploy locally (BURN_MINT or LOCK_RELEASE).
  /// @param tokenInitCode The creation code for the token, which includes the constructor parameters already appended.
  /// @param tokenPoolInitCode The creation code for the token pool, without the constructor parameters appended.
  /// @param lockBox The lockbox associated with the token, required for lock/release pools.
  /// @param salt The salt to be used in the create2 deployment of the token and token pool to ensure a unique address.
  /// @return token The address of the token that was deployed.
  /// @return pool The address of the token pool that was deployed.
  function deployTokenAndTokenPool(
    RemoteTokenPoolInfo[] calldata remoteTokenPools,
    uint8 localTokenDecimals,
    PoolType localPoolType,
    bytes memory tokenInitCode,
    bytes calldata tokenPoolInitCode,
    address lockBox,
    bytes32 salt
  ) external returns (address, address) {
    // Ensure a unique deployment between senders even if the same input parameter is used to prevent
    // DOS/front running attacks.
    salt = keccak256(abi.encodePacked(salt, msg.sender));

    // Deploy the token. The constructor parameters are already provided in the tokenInitCode.
    address token = Create2.deploy(0, salt, tokenInitCode);

    LocalPoolConfig memory localConfig = LocalPoolConfig({
      token: token, localTokenDecimals: localTokenDecimals, localPoolType: localPoolType, lockBox: lockBox, salt: salt
    });

    // Deploy the token pool.
    address pool = _createTokenPool(remoteTokenPools, tokenPoolInitCode, localConfig);

    // Grant the mint and burn roles to the pool for the token.
    FactoryBurnMintERC20(token).grantMintAndBurnRoles(pool);

    // Set the token pool for token in the token admin registry since this contract is the token and pool owner.
    _setTokenPoolInTokenAdminRegistry(token, pool);

    // Begin the 2 step ownership transfer of the newly deployed token to the msg.sender.
    IOwnable(token).transferOwnership(msg.sender);

    return (token, pool);
  }

  /// @notice Deploys a token pool with an existing ERC20 token.
  /// @dev Since the token already exists, this contract is not the owner and therefore cannot configure the
  /// token pool in the token admin registry in the same transaction. The user must invoke the calls to the
  /// tokenAdminRegistry manually.
  /// @dev since the token already exists, the owner must grant the mint and burn roles to the pool manually.
  /// @param token The address of the existing token to be used in the token pool.
  /// @param localTokenDecimals The amount of decimals used in the existing token. Since decimals() is not part of the
  /// the ERC20 standard, and thus cannot be certain to exist, the amount must be supplied via user input.
  /// @param localPoolType The type of pool to deploy locally (BURN_MINT or LOCK_RELEASE).
  /// @param remoteTokenPools An array of remote token pools info to be used in the pool's applyChainUpdates function.
  /// @param tokenPoolInitCode The creation code for the token pool.
  /// @param lockBox The lockbox associated with the token, required for lock/release pools.
  /// @param salt The salt to be used in the create2 deployment of the token pool.
  /// @return poolAddress The address of the token pool that was deployed.
  function deployTokenPoolWithExistingToken(
    address token,
    uint8 localTokenDecimals,
    PoolType localPoolType,
    RemoteTokenPoolInfo[] calldata remoteTokenPools,
    bytes calldata tokenPoolInitCode,
    address lockBox,
    bytes32 salt
  ) external returns (address poolAddress) {
    // Ensure a unique deployment between senders even if the same input parameter is used to prevent
    // DOS/front running attacks.
    salt = keccak256(abi.encodePacked(salt, msg.sender));

    LocalPoolConfig memory localConfig = LocalPoolConfig({
      token: token, localTokenDecimals: localTokenDecimals, localPoolType: localPoolType, lockBox: lockBox, salt: salt
    });

    // create the token pool and return the address.
    return _createTokenPool(remoteTokenPools, tokenPoolInitCode, localConfig);
  }

  // ================================================================
  // │                Pool Deployment/Configuration                 │
  // ================================================================

  /// @notice Deploys a token pool with the given token information and remote token pools.
  /// @param remoteTokenPools An array of remote token pools info to be used in the pool's applyChainUpdates function.
  /// @param tokenPoolInitCode The creation code for the token pool.
  /// @param localConfig Local deployment config including token, decimals, pool type, lockbox, and salt.
  /// @return poolAddress The address of the token pool that was deployed.
  function _createTokenPool(
    RemoteTokenPoolInfo[] calldata remoteTokenPools,
    bytes calldata tokenPoolInitCode,
    LocalPoolConfig memory localConfig
  ) private returns (address) {
    // Create an array of chain updates to apply to the token pool.
    TokenPool.ChainUpdate[] memory chainUpdates = new TokenPool.ChainUpdate[](remoteTokenPools.length);

    for (uint256 i = 0; i < remoteTokenPools.length; ++i) {
      chainUpdates[i] = _buildChainUpdate(remoteTokenPools[i], localConfig.salt);
    }

    // Construct the initArgs for the token pool using the immutable contracts for CCIP on the local chain.
    // LockRelease pools need lockBox, BurnMint pools don't.
    bytes memory tokenPoolInitArgs;
    address localLockBox;
    if (localConfig.localPoolType == PoolType.LOCK_RELEASE) {
      localLockBox = localConfig.lockBox;
      if (localLockBox == address(0)) {
        localLockBox = _deployLockBox(localConfig.token, localConfig.salt);
      } else {
        ERC20LockBox lockBoxContract = ERC20LockBox(localLockBox);
        if (address(lockBoxContract.getToken()) != localConfig.token) {
          revert InvalidLockBoxToken(address(lockBoxContract.getToken()), localConfig.token);
        }
      }
      tokenPoolInitArgs = abi.encode(
        localConfig.token, localConfig.localTokenDecimals, address(0), i_rmnProxy, i_ccipRouter, localLockBox
      );
    } else {
      tokenPoolInitArgs =
        abi.encode(localConfig.token, localConfig.localTokenDecimals, address(0), i_rmnProxy, i_ccipRouter);
    }

    // Construct the deployment code from the initCode and the initArgs and then deploy.
    address poolAddress = Create2.deploy(0, localConfig.salt, abi.encodePacked(tokenPoolInitCode, tokenPoolInitArgs));

    // Apply the chain updates to the token pool.
    TokenPool(poolAddress).applyChainUpdates(new uint64[](0), chainUpdates);

    // Authorize the new pool to interact with the local lockbox and transfer ownership to the caller for future admin.
    if (localConfig.localPoolType == PoolType.LOCK_RELEASE) {
      ERC20LockBox lockBoxContract = ERC20LockBox(localLockBox);
      // We check the owner as user supplied lockboxes can be different.
      if (lockBoxContract.owner() == address(this)) {
        _authorizePoolInLockBox(localLockBox, poolAddress);
        lockBoxContract.transferOwnership(msg.sender);
      }
    }

    // Begin the 2 step ownership transfer of the token pool to the msg.sender.
    IOwnable(poolAddress).transferOwnership(address(msg.sender));

    return poolAddress;
  }

  function _buildChainUpdate(
    RemoteTokenPoolInfo memory remoteTokenPool,
    bytes32 salt
  ) private pure returns (TokenPool.ChainUpdate memory) {
    address remoteLockBox = remoteTokenPool.remoteChainConfig.remoteLockBox;

    // If the user provides an empty byte string, indicated no token has already been deployed,
    // then the address of the token needs to be predicted. Otherwise the address provided will be used.
    if (remoteTokenPool.remoteTokenAddress.length == 0) {
      // The user must provide the initCode for the remote token, so its address can be predicted correctly. It's
      // provided in the remoteTokenInitCode field for the remoteTokenPool.
      remoteTokenPool.remoteTokenAddress = abi.encode(
        salt.computeAddress(
          keccak256(remoteTokenPool.remoteTokenInitCode), remoteTokenPool.remoteChainConfig.remotePoolFactory
        )
      );
    }

    // For lock/release pools, predict the remote lockbox if none is provided.
    if (remoteTokenPool.poolType == PoolType.LOCK_RELEASE && remoteLockBox == address(0)) {
      address decodedRemoteToken = abi.decode(remoteTokenPool.remoteTokenAddress, (address));
      (remoteLockBox,) =
        _computeLockBoxAddress(decodedRemoteToken, salt, remoteTokenPool.remoteChainConfig.remotePoolFactory);
      remoteTokenPool.remoteChainConfig.remoteLockBox = remoteLockBox;
    }

    // If the user provides an empty byte string parameter, indicating the pool has not been deployed yet,
    // the address of the pool should be predicted. Otherwise use the provided address.
    if (remoteTokenPool.remotePoolAddress.length == 0) {
      // Address is predicted based on the init code hash and the deployer, so the hash must first be computed
      // using the initCode and a concatenated set of constructor parameters.
      bytes32 remotePoolInitcodeHash = _generatePoolInitcodeHash(
        remoteTokenPool.remotePoolInitCode,
        remoteTokenPool.remoteChainConfig,
        abi.decode(remoteTokenPool.remoteTokenAddress, (address)),
        remoteTokenPool.poolType
      );

      // Abi encode the computed remote address so it can be used as bytes in the chain update.
      remoteTokenPool.remotePoolAddress =
        abi.encode(salt.computeAddress(remotePoolInitcodeHash, remoteTokenPool.remoteChainConfig.remotePoolFactory));
    }

    bytes[] memory remotePoolAddresses = new bytes[](1);
    remotePoolAddresses[0] = remoteTokenPool.remotePoolAddress;

    return TokenPool.ChainUpdate({
      remoteChainSelector: remoteTokenPool.remoteChainSelector,
      remotePoolAddresses: remotePoolAddresses,
      remoteTokenAddress: remoteTokenPool.remoteTokenAddress,
      outboundRateLimiterConfig: remoteTokenPool.rateLimiterConfig,
      inboundRateLimiterConfig: remoteTokenPool.rateLimiterConfig
    });
  }

  function _deployLockBox(
    address token,
    bytes32 salt
  ) private returns (address lockBox) {
    (address predicted, bytes memory creationCode) = _computeLockBoxAddress(token, salt, address(this));
    lockBox = Create2.deploy(0, salt, creationCode);
    // If deployment fails Create2 reverts; address mismatch is impossible since salt and init code are deterministic.
    if (lockBox != predicted) revert InvalidZeroAddress();
    return lockBox;
  }

  function _computeLockBoxAddress(
    address token,
    bytes32 salt,
    address deployer
  ) private pure returns (address predicted, bytes memory creationCode) {
    creationCode = abi.encodePacked(LOCKBOX_INIT_CODE, abi.encode(token, bytes32(0)));
    predicted = salt.computeAddress(keccak256(creationCode), deployer);
    return (predicted, creationCode);
  }

  function _authorizePoolInLockBox(
    address lockBox,
    address pool
  ) private {
    ERC20LockBox lockBoxContract = ERC20LockBox(lockBox);
    // Skip if this factory is not the owner; user-supplied lockboxes must already authorize the pool.
    if (lockBoxContract.owner() != address(this)) {
      return;
    }

    address[] memory added = new address[](1);
    added[0] = pool;
    lockBoxContract.applyAuthorizedCallerUpdates(
      AuthorizedCallers.AuthorizedCallerArgs({addedCallers: added, removedCallers: new address[](0)})
    );
  }

  /// @notice Generates the hash of the init code the pool will be deployed with.
  /// @dev The init code hash is used with Create2 to predict the address of the pool on the remote chain.
  /// @dev ABI-encoding limitations prevent arbitrary constructor parameters from being used, so pool type must be.
  /// restricted to those with known types in the constructor. This function should be updated if new pool types are needed.
  /// @param initCode The init code of the pool.
  /// @param remoteChainConfig The remote chain config for the pool.
  /// @param remoteTokenAddress The address of the remote token.
  /// @param poolType The type of pool being deployed.
  /// @return bytes32 hash of the init code to be used in the deterministic address calculation.
  function _generatePoolInitcodeHash(
    bytes memory initCode,
    RemoteChainConfig memory remoteChainConfig,
    address remoteTokenAddress,
    PoolType poolType
  ) private pure returns (bytes32) {
    bytes memory constructorParams;

    // LockRelease pools have an additional lockBox parameter.
    if (poolType == PoolType.LOCK_RELEASE) {
      // constructor(address token, uint8 localTokenDecimals, address advancedPoolHooks, address rmnProxy, address router, address lockBox).
      constructorParams = abi.encode(
        remoteTokenAddress,
        remoteChainConfig.remoteTokenDecimals,
        address(0),
        remoteChainConfig.remoteRMNProxy,
        remoteChainConfig.remoteRouter,
        remoteChainConfig.remoteLockBox
      );
    } else {
      // constructor(address token, uint8 localTokenDecimals, address advancedPoolHooks, address rmnProxy, address router).
      constructorParams = abi.encode(
        remoteTokenAddress,
        remoteChainConfig.remoteTokenDecimals,
        address(0),
        remoteChainConfig.remoteRMNProxy,
        remoteChainConfig.remoteRouter
      );
    }

    return keccak256(abi.encodePacked(initCode, constructorParams));
  }

  /// @notice Sets the token pool address in the token admin registry for a newly deployed token pool.
  /// @dev this function should only be called when the token is deployed by this contract as well, otherwise
  /// the token pool will not be able to be set in the token admin registry, and this function will revert.
  /// @param token The address of the token to set the pool for.
  /// @param pool The address of the pool to set in the token admin registry.
  function _setTokenPoolInTokenAdminRegistry(
    address token,
    address pool
  ) private {
    i_registryModuleOwnerCustom.registerAdminViaOwner(token);
    i_tokenAdminRegistry.acceptAdminRole(token);
    i_tokenAdminRegistry.setPool(token, pool);

    // Begin the 2 admin transfer process which must be accepted in a separate tx.
    i_tokenAdminRegistry.transferAdminRole(token, msg.sender);
  }
}
