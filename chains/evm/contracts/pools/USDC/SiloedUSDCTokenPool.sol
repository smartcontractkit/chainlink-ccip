// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Pool} from "../../libraries/Pool.sol";
import {SiloedLockReleaseTokenPool} from "../SiloedLockReleaseTokenPool.sol";

import {Ownable2StepMsgSender} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2StepMsgSender.sol";
import {IBurnMintERC20} from "@chainlink/contracts/src/v0.8/shared/token/ERC20/IBurnMintERC20.sol";
import {IERC20} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";
import {EnumerableSet} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v5.0.2/contracts/utils/structs/EnumerableSet.sol";

/// @notice A token pool for USDC which inherits the Siloed token functionality while adding the CCTP migration functionality
contract SiloedUSDCTokenPool is SiloedLockReleaseTokenPool {
  using EnumerableSet for EnumerableSet.AddressSet;
  using EnumerableSet for EnumerableSet.UintSet;

  event AllowedTokenPoolProxyAdded(address tokenPoolProxy);
  event AllowedTokenPoolProxyRemoved(address tokenPoolProxy);
  event CCTPMigrationProposed(uint64 remoteChainSelector);
  event CCTPMigrationExecuted(uint64 remoteChainSelector, uint256 USDCBurned);
  event CCTPMigrationCancelled(uint64 existingProposalSelector);
  event CircleMigratorAddressSet(address migratorAddress);
  event TokensExcludedFromBurn(
    uint64 indexed remoteChainSelector, uint256 amount, uint256 burnableAmountAfterExclusion
  );

  error TokenPoolProxyAlreadyAllowed(address tokenPoolProxy);
  error TokenPoolProxyNotAllowed(address tokenPoolProxy);
  error onlyCircle();
  error ExistingMigrationProposal();
  error NoMigrationProposalPending();
  error ChainAlreadyMigrated(uint64 remoteChainSelector);
  error TokenLockingNotAllowedAfterMigration(uint64 remoteChainSelector);

  EnumerableSet.AddressSet internal s_allowedTokenPoolProxies;
  IBurnMintERC20 private immutable i_USDC;

  address internal s_circleUSDCMigrator;
  uint64 internal s_proposedUSDCMigrationChain;

  mapping(uint64 remoteChainSelector => uint256 excludedTokens) internal s_tokensExcludedFromBurn;

  EnumerableSet.UintSet internal s_migratedChains;

  constructor(
    IERC20 token,
    uint8 localTokenDecimals,
    address[] memory allowlist,
    address rmnProxy,
    address router,
    address lockBox
  ) SiloedLockReleaseTokenPool(token, localTokenDecimals, allowlist, rmnProxy, router, lockBox) {
    i_USDC = IBurnMintERC20(address(token));
  }

  /// @notice Release tokens for a specific chain selector.
  /// @dev This function can only be called by an address specified by the owner to be controlled by circle
  /// @dev proposeCCTPMigration must be called first on an approved lane to execute properly.
  /// @dev This function signature should NEVER be overwritten, otherwise it will be unable to be called by
  /// circle to properly migrate USDC over to CCTP.
  function releaseOrMint(
    Pool.ReleaseOrMintInV1 calldata releaseOrMintIn
  ) public virtual override returns (Pool.ReleaseOrMintOutV1 memory) {
    // Calculate the local amount. Since USDC is always 6 decimals, we can hard code the decimals to 6.
    uint256 localAmount = _calculateLocalAmount(releaseOrMintIn.sourceDenominatedAmount, 6);

    _validateReleaseOrMint(releaseOrMintIn, localAmount);

    // Save gas by using storage instead of memory as a value may need to be updated.
    SiloConfig storage remoteConfig = s_chainConfigs[releaseOrMintIn.remoteChainSelector];

    // Since remoteConfig.isSiloed is used more than once, caching in memory saves gas instead of multiple SLOADs.
    bool chainIsSiloed = remoteConfig.isSiloed;

    // Additional security check to prevent underflow by explicitly ensuring that enough funds are available to release
    uint256 availableLiquidity = chainIsSiloed ? remoteConfig.tokenBalance : s_unsiloedTokenBalance;
    if (localAmount > availableLiquidity) revert InsufficientLiquidity(availableLiquidity, localAmount);

    // If the chain is Siloed then subtract the amount from the accounting for the silo
    if (remoteConfig.isSiloed) {
      // If the chain is Siloed and has no locked tokens, that means a migration has already occurred
      // and the tokens should be excluded from burn, so we need to subtract the amount from the excluded tokens
      // instead of the locked tokens
      if (remoteConfig.tokenBalance == 0) {
        s_tokensExcludedFromBurn[releaseOrMintIn.remoteChainSelector] -= localAmount;
      } else {
        remoteConfig.tokenBalance -= localAmount;
      }
    }

    // Release to the recipient
    i_lockBox.withdraw(address(i_token), localAmount, releaseOrMintIn.receiver, releaseOrMintIn.remoteChainSelector);

    emit ReleasedOrMinted({
      remoteChainSelector: releaseOrMintIn.remoteChainSelector,
      token: address(i_token),
      sender: msg.sender,
      recipient: releaseOrMintIn.receiver,
      amount: localAmount
    });

    return Pool.ReleaseOrMintOutV1({destinationAmount: localAmount});
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
    if (s_migratedChains.contains(remoteChainSelector)) {
      revert TokenLockingNotAllowedAfterMigration(remoteChainSelector);
    }

    super._provideLiquidity(remoteChainSelector, amount);
  }

  /// @notice Set the allowed token pool proxies for the pool.
  /// @dev This function can only be called by the owner
  /// @param tokenPoolProxies The token pool proxies to set the allowed status for
  /// @param allowed The allowed status for the token pool proxies
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

  /// @notice Get the allowed token pool proxies for the pool.
  /// @return address[] The allowed token pool proxies
  function getAllowedTokenPoolProxies() public view returns (address[] memory) {
    return s_allowedTokenPoolProxies.values();
  }

  /// @notice Propose a destination chain to migrate from lock/release mechanism to CCTP enabled burn/mint
  /// through a Circle controlled burn.
  /// @param remoteChainSelector the CCIP specific selector for the remote chain currently using a
  /// non-canonical form of USDC which they wish to update to canonical. Function will revert if an existing migration
  /// proposal is already in progress.
  /// @dev This function can only be called by the owner
  function proposeCCTPMigration(
    uint64 remoteChainSelector
  ) external onlyOwner {
    // Prevent overwriting existing migration proposals until the current one is finished
    if (s_proposedUSDCMigrationChain != 0) revert ExistingMigrationProposal();
    if (s_migratedChains.contains(remoteChainSelector)) revert ChainAlreadyMigrated(remoteChainSelector);

    s_proposedUSDCMigrationChain = remoteChainSelector;

    emit CCTPMigrationProposed(remoteChainSelector);
  }

  /// @notice Cancel an existing proposal to migrate a lane to CCTP.
  /// @notice This function will revert if no proposal is currently in progress.
  function cancelExistingCCTPMigrationProposal() external onlyOwner {
    if (s_proposedUSDCMigrationChain == 0) revert NoMigrationProposalPending();

    uint64 currentProposalChainSelector = s_proposedUSDCMigrationChain;
    delete s_proposedUSDCMigrationChain;

    // If a migration is cancelled, the tokens excluded from burn should be reset, and must be manually
    // re-excluded if the proposal is re-proposed in the future
    delete s_tokensExcludedFromBurn[currentProposalChainSelector];

    emit CCTPMigrationCancelled(currentProposalChainSelector);
  }

  /// @notice retrieve the chain selector for an ongoing CCTP migration in progress.
  /// @return uint64 the chain selector of the lane to be migrated. Will be zero if no proposal currently
  /// exists
  function getCurrentProposedCCTPChainMigration() public view returns (uint64) {
    return s_proposedUSDCMigrationChain;
  }

  /// @notice Set the address of the circle-controlled wallet which will execute a CCTP lane migration
  /// @dev The function should only be invoked once the address has been confirmed by Circle prior to
  /// chain expansion.
  function setCircleMigratorAddress(
    address migrator
  ) external onlyOwner {
    s_circleUSDCMigrator = migrator;

    emit CircleMigratorAddressSet(migrator);
  }

  /// @notice Exclude tokens to be burned in a CCTP-migration because the amount are locked in an undelivered message.
  /// @dev When a message is sitting in manual execution from the L/R chain, those tokens need to be excluded from
  /// being burned in a CCTP-migration otherwise the message will never be able to be delivered due to it not having
  /// an attestation on the source-chain to mint. In that instance it should use provided liquidity that was designated
  /// @dev This function should ONLY be called on the home chain, where tokens are locked, NOT on the remote chain
  /// and strict scrutiny should be applied to ensure that the amount of tokens excluded is accurate.
  function excludeTokensFromBurn(uint64 remoteChainSelector, uint256 amount) external onlyOwner {
    if (s_proposedUSDCMigrationChain != remoteChainSelector) revert NoMigrationProposalPending();

    s_tokensExcludedFromBurn[remoteChainSelector] += amount;

    uint256 burnableAmountAfterExclusion =
      s_chainConfigs[remoteChainSelector].tokenBalance - s_tokensExcludedFromBurn[remoteChainSelector];

    emit TokensExcludedFromBurn(remoteChainSelector, amount, burnableAmountAfterExclusion);
  }

  /// @notice Get the amount of tokens excluded from being burned in a CCTP-migration
  /// @dev The sum of locked tokens and excluded tokens should equal the supply of the token on the remote chain
  /// @param remoteChainSelector The chain for which the excluded tokens are being queried
  /// @return uint256 amount of tokens excluded from being burned in a CCTP-migration
  function getExcludedTokensByChain(
    uint64 remoteChainSelector
  ) external view returns (uint256) {
    return s_tokensExcludedFromBurn[remoteChainSelector];
  }

  /// @notice Burn USDC locked for a specific lane so that destination USDC can be converted from
  /// non-canonical to canonical USDC.
  /// @dev This function can only be called by an address specified by the owner to be controlled by circle
  /// @dev This function signature should NEVER be overwritten, otherwise it will be unable to be called by
  /// circle to properly migrate USDC over to CCTP.
  function burnLockedUSDC() external {
    if (msg.sender != s_circleUSDCMigrator) revert onlyCircle();

    uint64 burnChainSelector = s_proposedUSDCMigrationChain;
    if (burnChainSelector == 0) revert NoMigrationProposalPending();

    // Burnable tokens is the total locked minus the amount excluded from burn
    uint256 tokensToBurn = s_chainConfigs[burnChainSelector].tokenBalance - s_tokensExcludedFromBurn[burnChainSelector];

    // The CCTP burn function will attempt to burn out of the contract that calls it, so we need to withdraw the tokens
    // from the lock box first otherwise the burn will revert.
    i_lockBox.withdraw(address(i_token), tokensToBurn, address(this), burnChainSelector);

    // Even though USDC is a trusted call, ensure CEI by updating state first
    delete s_chainConfigs[burnChainSelector].tokenBalance;
    delete s_proposedUSDCMigrationChain;

    // This should only be called after this contract has been granted a "zero allowance minter role" on USDC by Circle,
    // otherwise the call will revert. Executing this burn will functionally convert all USDC on the destination chain
    // to canonical USDC by removing the canonical USDC backing it from circulation.
    i_USDC.burn(tokensToBurn);

    s_migratedChains.add(burnChainSelector);

    emit CCTPMigrationExecuted(burnChainSelector, tokensToBurn);
  }
}
