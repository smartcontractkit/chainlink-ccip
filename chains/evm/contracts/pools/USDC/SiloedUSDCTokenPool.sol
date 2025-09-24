// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Pool} from "../../libraries/Pool.sol";
import {SiloedLockReleaseTokenPool} from "../SiloedLockReleaseTokenPool.sol";

import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";
import {IBurnMintERC20} from "@chainlink/contracts/src/v0.8/shared/token/ERC20/IBurnMintERC20.sol";
import {IERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/IERC20.sol";
import {EnumerableSet} from "@openzeppelin/contracts@5.0.2/utils/structs/EnumerableSet.sol";

/// @dev The flag used to indicate that the source pool data is coming from a chain that does not have CCTP Support,
/// and so the lock release pool should be used. The BurnMintWithLockReleaseTokenPool uses this flag as its source pool
/// data to indicate that the tokens should be released from the lock release pool rather than attempting to be minted
/// through CCTP.
/// @dev The preimage is bytes4(keccak256("NO_CCTP_USE_LOCK_RELEASE")).
/// Note: This will be removed in a following PR but is included here to prevent breaking changes to the
/// BurnMintWithLockReleaseTokenPool.
bytes4 constant LOCK_RELEASE_FLAG = 0xfa7c07de;

/// @notice A token pool for USDC which inherits the Siloed token functionality while adding the CCTP migration functionality.
/// @dev The CCTP migration functions have been previously audited. The code has been moved from its own contract
/// to this, to maximize simplicity. The only difference is that custom balance tracking
/// has been removed and instead is now inherited from the SiloedLockReleaseTokenPool.
/// @dev While this technically supports unsiloed chains, as inherited from the parent contract,
/// it is not recommended to use them. All chains should be siloed, otherwise the chain will not be
/// able to migrate to CCTP in the future, due to the inability to manage the token
/// balances under CCTP accounting rules defined at:
/// https://github.com/circlefin/stablecoin-evm/blob/master/doc/bridged_USDC_standard.md
contract SiloedUSDCTokenPool is SiloedLockReleaseTokenPool, AuthorizedCallers {
  using EnumerableSet for EnumerableSet.UintSet;

  event CCTPMigrationProposed(uint64 remoteChainSelector);
  event CCTPMigrationExecuted(uint64 remoteChainSelector, uint256 USDCBurned);
  event CCTPMigrationCancelled(uint64 existingProposalSelector);
  event CircleMigratorAddressSet(address migratorAddress);
  event TokensExcludedFromBurn(
    uint64 indexed remoteChainSelector, uint256 amount, uint256 burnableAmountAfterExclusion
  );

  error OnlyCircle();
  error ExistingMigrationProposal();
  error NoMigrationProposalPending();
  error ChainAlreadyMigrated(uint64 remoteChainSelector);
  error TokenLockingNotAllowedAfterMigration(uint64 remoteChainSelector);

  /// @notice The address of the circle-controlled wallet which will execute a CCTP lane migration
  address internal s_circleUSDCMigrator;
  uint64 internal s_proposedUSDCMigrationChain;

  /// @notice The tokens excluded from being burned in a CCTP-migration.
  mapping(uint64 remoteChainSelector => uint256 excludedTokens) internal s_tokensExcludedFromBurn;

  /// @notice The chains that have been migrated to CCTP.
  EnumerableSet.UintSet internal s_migratedChains;

  /// @dev The authorized callers are set as empty since the USDCTokenPoolProxy is the only authorized caller,
  /// but cannot be deployed until after this contract. The allowed callers will be set after deployment.
  constructor(
    IERC20 token,
    uint8 localTokenDecimals,
    address[] memory allowlist,
    address rmnProxy,
    address router,
    address lockBox
  )
    SiloedLockReleaseTokenPool(token, localTokenDecimals, allowlist, rmnProxy, router, lockBox)
    AuthorizedCallers(new address[](0))
  {}

  /// @notice Using a function because constant state variables cannot be overridden by child contracts.
  function typeAndVersion() external pure virtual override returns (string memory) {
    return "SiloedUSDCTokenPool 1.6.3-dev";
  }

  /// @inheritdoc SiloedLockReleaseTokenPool
  function releaseOrMint(
    Pool.ReleaseOrMintInV1 calldata releaseOrMintIn
  ) public virtual override returns (Pool.ReleaseOrMintOutV1 memory) {
    // Calculate the local amount. Since USDC is always 6 decimals, we can hard code the decimals to 6.
    uint256 localAmount = _calculateLocalAmount(releaseOrMintIn.sourceDenominatedAmount, 6);

    _validateReleaseOrMint(releaseOrMintIn, localAmount);

    // Save gas by using storage instead of memory as a value may need to be updated.
    SiloConfig storage remoteConfig = s_chainConfigs[releaseOrMintIn.remoteChainSelector];

    uint256 excludedTokens = s_tokensExcludedFromBurn[releaseOrMintIn.remoteChainSelector];

    // No excluded tokens is the common path, as it means no migration has occured yet, and any released
    // tokens should come from the stored token balance of previously deposited tokens.
    if (excludedTokens == 0) {
      if (localAmount > remoteConfig.tokenBalance) revert InsufficientLiquidity(remoteConfig.tokenBalance, localAmount);

      remoteConfig.tokenBalance -= localAmount;

      // The existence of excluded tokens indicates a migration has occured on the chain, and that any tokens
      // being released should come from those excluded tokens reserved for processing inflight messages.
    } else {
      if (localAmount > excludedTokens) revert InsufficientLiquidity(excludedTokens, localAmount);
      s_tokensExcludedFromBurn[releaseOrMintIn.remoteChainSelector] -= localAmount;
    }

    // Release to the recipient
    i_lockBox.withdraw(address(i_token), localAmount, releaseOrMintIn.receiver);

    emit ReleasedOrMinted({
      remoteChainSelector: releaseOrMintIn.remoteChainSelector,
      token: address(i_token),
      sender: msg.sender,
      recipient: releaseOrMintIn.receiver,
      amount: localAmount
    });

    return Pool.ReleaseOrMintOutV1({destinationAmount: localAmount});
  }

  /// @dev This function is overridden to prevent providing liquidity to a chain that has already been migrated, and thus should use CCTP-proper instead of a Lock/Release mechanism.
  function _provideLiquidity(uint64 remoteChainSelector, uint256 amount) internal override {
    if (s_migratedChains.contains(remoteChainSelector)) {
      revert TokenLockingNotAllowedAfterMigration(remoteChainSelector);
    }

    super._provideLiquidity(remoteChainSelector, amount);
  }

  // ================================================================
  // │                           Access                             │
  // ================================================================

  /// @dev This function is overridden to remove the On-Ramp check, as this pool does not receive calls
  /// from the ramps directly, instead receiving them from a proxy contract first.
  function _onlyOnRamp(
    uint64 remoteChainSelector
  ) internal view virtual override {
    if (!isSupportedChain(remoteChainSelector)) revert ChainNotAllowed(remoteChainSelector);

    // Validate logic is inherited from AuthorizedCallers, and is used to validate that the caller is the authorized USDC proxy contract rather than the ramp.
    _validateCaller();
  }

  /// @dev This function is overridden to remove the Off-Ramp check, as this pool does not receive calls
  /// from the ramps directly, instead receiving them from a proxy contract first.
  function _onlyOffRamp(
    uint64 remoteChainSelector
  ) internal view virtual override {
    if (!isSupportedChain(remoteChainSelector)) revert ChainNotAllowed(remoteChainSelector);

    // Validate logic is inherited from AuthorizedCallers, and is used to validate that the caller is the authorized USDC proxy contract rather than the ramp.
    _validateCaller();
  }

  // ================================================================
  // │                  CCTP Migration functions                    │
  // ================================================================

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
    if (remoteChainSelector == 0) revert SiloedLockReleaseTokenPool.InvalidChainSelector(0);

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
    if (msg.sender != s_circleUSDCMigrator) revert OnlyCircle();

    uint64 burnChainSelector = s_proposedUSDCMigrationChain;
    if (burnChainSelector == 0) revert NoMigrationProposalPending();

    // Burnable tokens is the total locked minus the amount excluded from burn
    uint256 tokensToBurn = s_chainConfigs[burnChainSelector].tokenBalance - s_tokensExcludedFromBurn[burnChainSelector];

    // The CCTP burn function will attempt to burn out of the contract that calls it, so we need to withdraw the tokens
    // from the lock box first otherwise the burn will revert.
    i_lockBox.withdraw(address(i_token), tokensToBurn, address(this));

    // Even though USDC is a trusted call, ensure CEI by updating state first
    delete s_chainConfigs[burnChainSelector].tokenBalance;
    delete s_proposedUSDCMigrationChain;

    // This should only be called after this contract has been granted a "zero allowance minter role" on USDC by Circle,
    // otherwise the call will revert. Executing this burn will functionally convert all USDC on the destination chain
    // to canonical USDC by removing the canonical USDC backing it from circulation.
    IBurnMintERC20(address(i_token)).burn(tokensToBurn);

    s_migratedChains.add(burnChainSelector);

    emit CCTPMigrationExecuted(burnChainSelector, tokensToBurn);
  }
}
