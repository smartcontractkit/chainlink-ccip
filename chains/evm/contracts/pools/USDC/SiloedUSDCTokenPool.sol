// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ITokenMessenger} from "./interfaces/ITokenMessenger.sol";

import {Pool} from "../../libraries/Pool.sol";
import {ERC20LockBox} from "../ERC20LockBox.sol";
import {TokenPool} from "../TokenPool.sol";
import {CCTPMessageTransmitterProxy} from "../USDC/CCTPMessageTransmitterProxy.sol";
import {USDCTokenPoolCCTPV2} from "../USDC/USDCTokenPoolCCTPV2.sol";

import {SiloedLockReleaseTokenPool} from "../SiloedLockReleaseTokenPool.sol";
import {USDCBridgeMigrator} from "./USDCBridgeMigrator.sol";

import {IERC20} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/utils/SafeERC20.sol";
import {EnumerableSet} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v5.0.2/contracts/utils/structs/EnumerableSet.sol";

/// @notice A token pool for USDC which uses CCTP for supported chains and Lock/Release for all others
/// @dev The functionality from LockReleaseTokenPool.sol has been duplicated due to lack of compiler support for shared
/// constructors between parents
/// @dev The primary token mechanism in this pool is Burn/Mint with CCTP, with Lock/Release as the
/// secondary, opt in mechanism for chains not currently supporting CCTP.
contract SiloedUSDCTokenPool is SiloedLockReleaseTokenPool, USDCBridgeMigrator {
  using EnumerableSet for EnumerableSet.AddressSet;

  event AllowedTokenPoolProxyAdded(address tokenPoolProxy);
  event AllowedTokenPoolProxyRemoved(address tokenPoolProxy);

  error TokenPoolProxyAlreadyAllowed(address tokenPoolProxy);
  error TokenPoolProxyNotAllowed(address tokenPoolProxy);

  EnumerableSet.AddressSet internal s_allowedTokenPoolProxies;

  constructor(
    IERC20 token,
    uint8 localTokenDecimals,
    address[] memory allowlist,
    address rmnProxy,
    address router,
    address lockBox
  )
    SiloedLockReleaseTokenPool(token, localTokenDecimals, allowlist, rmnProxy, router)
    USDCBridgeMigrator(address(token), lockBox)
  {}

  /// @notice Lock tokens for a specific chain selector.
  /// @dev This function can only be called by an address specified by the owner to be controlled by circle
  /// @dev proposeCCTPMigration must be called first on an approved lane to execute properly.
  /// @dev This function signature should NEVER be overwritten, otherwise it will be unable to be called by
  /// circle to properly migrate USDC over to CCTP.
  function lockOrBurn(
    Pool.LockOrBurnInV1 calldata lockOrBurnIn
  ) public virtual override returns (Pool.LockOrBurnOutV1 memory lockOrBurnOut) {

    // The super call is made first to take advantage of the validation logic in the parent contract
    lockOrBurnOut = super.lockOrBurn(lockOrBurnIn);

    // If the chain is Siloed then add it to the accounting for the bridge migrator code  
    if (s_chainConfigs[lockOrBurnIn.remoteChainSelector].isSiloed) {
      s_lockedTokensByChainSelector[lockOrBurnIn.remoteChainSelector] += lockOrBurnIn.amount;
    }

    return lockOrBurnOut;
  }

  /// @notice Release tokens for a specific chain selector.
  /// @dev This function can only be called by an address specified by the owner to be controlled by circle
  /// @dev proposeCCTPMigration must be called first on an approved lane to execute properly.
  /// @dev This function signature should NEVER be overwritten, otherwise it will be unable to be called by
  /// circle to properly migrate USDC over to CCTP.
  function releaseOrMint(
    Pool.ReleaseOrMintInV1 calldata releaseOrMintIn
  ) public virtual override returns (Pool.ReleaseOrMintOutV1 memory releaseOrMintOut) {

    // The super call is made first to take advantage of the validation logic in the parent contract
    releaseOrMintOut = super.releaseOrMint(releaseOrMintIn);

    // 6 Decimals is ok to hard code since USDC is always 6 decimals
    uint256 localAmount = _calculateLocalAmount(releaseOrMintIn.sourceDenominatedAmount, 6);

    if (localAmount > s_lockedTokensByChainSelector[releaseOrMintIn.remoteChainSelector]) {
      revert InsufficientLiquidity(localAmount, s_lockedTokensByChainSelector[releaseOrMintIn.remoteChainSelector]);
    }

    // If the chain is Siloed then subtract the amount from the accounting for the bridge migrator code
    if (s_chainConfigs[releaseOrMintIn.remoteChainSelector].isSiloed) {
      s_lockedTokensByChainSelector[releaseOrMintIn.remoteChainSelector] -= localAmount;
    } 

    return releaseOrMintOut;
  }

  /// @notice This function is overridden to hard code the decimals to 6 since USDC is always 6 decimals
  function _parseRemoteDecimals(
    bytes memory sourcePoolData
  ) internal pure override returns (uint8) {
    // The parent contract attempts to parse the decimals from the source pool data. However, since the source pool data
    // is always equal to the LOCK_RELEASE_FLAG, we need to hard code the decimals to 6 otherwise the function will revert.

    // Since it is an invariant that a remote USDC token has 6 decimals, we can hard code the decimals to 6.
    // If this invariant is violated, the mint amount will be incorrect and the recipient will receive the wrong amount.
    return 6;
  }

    /// @notice Checks whether remote chain selector is configured on this contract, and if the msg.sender
  /// is a permissioned onRamp for the given chain on the Router.
  function _onlyOnRamp(
    uint64 remoteChainSelector
  ) internal view override {
    if (!isSupportedChain(remoteChainSelector)) revert ChainNotAllowed(remoteChainSelector);
    if (!s_allowedTokenPoolProxies.contains(msg.sender)) revert CallerIsNotARampOnRouter(msg.sender);
  }

  /// @notice Checks whether remote chain selector is configured on this contract, and if the msg.sender
  /// is a permissioned offRamp for the given chain on the Router.
  function _onlyOffRamp(
    uint64 remoteChainSelector
  ) internal view override {
    if (!isSupportedChain(remoteChainSelector)) revert ChainNotAllowed(remoteChainSelector);

    if (!s_allowedTokenPoolProxies.contains(msg.sender)) revert CallerIsNotARampOnRouter(msg.sender);
  }

  /// @notice This function is overridden to update the locked tokens accounting for the bridge migrator code
  function _provideLiquidity(uint64 remoteChainSelector, uint256 amount) internal override {
    super._provideLiquidity(remoteChainSelector, amount);

    // If the chain is Siloed then add it to the accounting for the bridge 
    if (s_chainConfigs[remoteChainSelector].isSiloed) {
      s_lockedTokensByChainSelector[remoteChainSelector] += amount;
    }
  }

  function setAllowedTokenPoolProxies(
    address[] calldata tokenPoolProxies, // The token pool proxies to set the allowed status for
    bool[] calldata allowed
  ) external onlyOwner {
    for (uint256 i = 0; i < tokenPoolProxies.length; ++i) {
      if (allowed[i]) {
        if (!s_allowedTokenPoolProxies.add(tokenPoolProxies[i])) {
          revert TokenPoolProxyAlreadyAllowed(tokenPoolProxies[i]);
        }

        emit AllowedTokenPoolProxyAdded(tokenPoolProxies[i]);
      } else {
        if (!s_allowedTokenPoolProxies.remove(tokenPoolProxies[i])) {
          revert TokenPoolProxyNotAllowed(tokenPoolProxies[i]);
        }

        emit AllowedTokenPoolProxyRemoved(tokenPoolProxies[i]);
      }
    }
  }
}
