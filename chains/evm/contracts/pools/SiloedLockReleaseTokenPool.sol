// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";

import {Pool} from "../libraries/Pool.sol";
import {ERC20LockBox} from "./ERC20LockBox.sol";
import {TokenPool} from "./TokenPool.sol";

import {IERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/utils/SafeERC20.sol";

/// @notice A variation on Lock Release token pools where liquidity is shared among some chains, and stored independently
/// for others. Chains which do not share liquidity are known as siloed chains.
/// @dev This pool defines the ERC20LockBox liquidity domain id as bytes32(remoteChainSelector).
contract SiloedLockReleaseTokenPool is TokenPool, ITypeAndVersion {
  using SafeERC20 for IERC20;

  error InsufficientLiquidity(uint256 availableLiquidity, uint256 requestedAmount);
  error ChainNotSiloed(uint64 remoteChainSelector);
  error InvalidChainSelector(uint64 remoteChainSelector);
  error LiquidityAmountCannotBeZero();
  error InvalidLockBoxLiquidityDomain(bytes32 liquidityDomainId);
  error LockBoxNotConfigured(bytes32 liquidityDomainId);

  event LiquidityAdded(uint64 remoteChainSelector, address indexed provider, uint256 amount);
  event LiquidityRemoved(uint64 remoteChainSelector, address indexed remover, uint256 amount);
  event ChainUnsiloed(uint64 remoteChainSelector, uint256 amountUnsiloed);
  event ChainSiloed(uint64 remoteChainSelector, address rebalancer);
  event SiloRebalancerSet(uint64 indexed remoteChainSelector, address oldRebalancer, address newRebalancer);
  event UnsiloedRebalancerSet(address oldRebalancer, address newRebalancer);

  struct LockBoxConfig {
    uint64 remoteChainSelector;
    address lockBox;
  }

  struct SiloConfigUpdate {
    uint64 remoteChainSelector;
    address rebalancer;
  }

  struct SiloConfig {
    address rebalancer; // ─╮ The address allowed to add liquidity for the given siloed chain.
    bool isSiloed; // ──────╯ Whether funds should be isolated from all other chains or shared amongst all non-siloed chains.
  }

  /// @notice The rebalancer for unsiloed chains, which can add liquidity to the shared pool.
  address internal s_rebalancer;

  /// @notice Lock boxes keyed by chain selector.
  mapping(uint64 remoteChainSelector => ERC20LockBox lockBox) internal s_lockBoxes;

  /// @notice The configuration for each chain that is siloed, or not. By default chains are not siloed.
  mapping(uint64 remoteChainSelector => SiloConfig) internal s_chainConfigs;

  constructor(
    IERC20 token,
    uint8 localTokenDecimals,
    address advancedPoolHooks,
    address rmnProxy,
    address router,
    address lockBox
  ) TokenPool(token, localTokenDecimals, advancedPoolHooks, rmnProxy, router) {
    if (lockBox == address(0)) revert ZeroAddressInvalid();

    ERC20LockBox erc20LockBox = ERC20LockBox(lockBox);
    if (address(erc20LockBox.getToken()) != address(token)) revert InvalidToken(address(erc20LockBox.getToken()));
    // We should have a lockbox for unsiloed liquidity.
    if (erc20LockBox.getLiquidityDomainId() != 0) {
      revert InvalidLockBoxLiquidityDomain(erc20LockBox.getLiquidityDomainId());
    }

    token.safeApprove(lockBox, type(uint256).max);
    s_lockBoxes[0] = erc20LockBox;
  }

  /// @notice Using a function because constant state variables cannot be overridden by child contracts.
  function typeAndVersion() external pure virtual override returns (string memory) {
    return "SiloedLockReleaseTokenPool 1.7.0-dev";
  }

  /// @notice Locks the token in the pool
  /// @param lockOrBurnIn The lock or burn input parameters.
  function lockOrBurn(
    Pool.LockOrBurnInV1 calldata lockOrBurnIn
  ) public virtual override returns (Pool.LockOrBurnOutV1 memory out) {
    // super.lockOrBurn will validate the lockOrBurnIn and revert if invalid.
    out = super.lockOrBurn(lockOrBurnIn);

    // lockBoxSelector chooses which lockbox holds liquidity for this transfer; unsiloed chains use selector 0.
    uint64 lockBoxSelector =
      s_chainConfigs[lockOrBurnIn.remoteChainSelector].isSiloed ? lockOrBurnIn.remoteChainSelector : 0;

    // Transfer the tokens to the appropriate lock box.
    _getLockBox(lockBoxSelector)
      .deposit(
        lockOrBurnIn.remoteChainSelector, address(i_token), bytes32(uint256(lockBoxSelector)), lockOrBurnIn.amount
      );

    return out;
  }

  /// @notice Release tokens from the pool to the recipient
  /// @dev The _validateReleaseOrMint check is an essential security check
  /// @dev If the releaseOrMintIn amount is greater than available liquidity, the function will revert as a security
  /// measure to prevent funds from a Silo being released by another chain.
  /// @param releaseOrMintIn The release or mint input parameters.
  function releaseOrMint(
    Pool.ReleaseOrMintInV1 calldata releaseOrMintIn
  ) public virtual override returns (Pool.ReleaseOrMintOutV1 memory) {
    // Calculate the local amount
    uint256 localAmount = _calculateLocalAmount(
      releaseOrMintIn.sourceDenominatedAmount, _parseRemoteDecimals(releaseOrMintIn.sourcePoolData)
    );

    _validateReleaseOrMint(releaseOrMintIn, localAmount, WAIT_FOR_FINALITY);

    // lockBoxSelector chooses which lockbox holds liquidity for this transfer; unsiloed chains use selector 0.
    uint64 lockBoxSelector =
      s_chainConfigs[releaseOrMintIn.remoteChainSelector].isSiloed ? releaseOrMintIn.remoteChainSelector : 0;

    uint256 availableLiquidity = i_token.balanceOf(address(_getLockBox(lockBoxSelector)));
    if (localAmount > availableLiquidity) revert InsufficientLiquidity(availableLiquidity, localAmount);

    // Release to the recipient
    _getLockBox(lockBoxSelector)
      .withdraw(
        releaseOrMintIn.remoteChainSelector,
        address(i_token),
        bytes32(uint256(lockBoxSelector)),
        localAmount,
        releaseOrMintIn.receiver
      );

    emit ReleasedOrMinted({
      remoteChainSelector: releaseOrMintIn.remoteChainSelector,
      token: address(i_token),
      sender: msg.sender,
      recipient: releaseOrMintIn.receiver,
      amount: localAmount
    });

    return Pool.ReleaseOrMintOutV1({destinationAmount: localAmount});
  }

  /// @notice Returns the amount of tokens in the token pool that were siloed for a specific remote chain selector.
  /// @param remoteChainSelector the CCIP specific selector for the remote chain being interacted with.
  /// @return lockedTokens The tokens locked into this token pool for the given selector. If the chain is not siloed,
  /// the amount will be the amount of liquidity shared among all unsiloed chains.
  function getAvailableTokens(
    uint64 remoteChainSelector
  ) external view returns (uint256 lockedTokens) {
    if (!isSupportedChain(remoteChainSelector)) revert InvalidChainSelector(remoteChainSelector);

    uint64 lockBoxSelector = s_chainConfigs[remoteChainSelector].isSiloed ? remoteChainSelector : 0;
    return i_token.balanceOf(address(_getLockBox(lockBoxSelector)));
  }

  /// @notice Returns the amount of tokens in the token pool that are shared among all unsiloed chains.
  /// @return unsiloedTokens amount of tokens available to all unsiloed chains.
  function getUnsiloedLiquidity() external view returns (uint256) {
    return i_token.balanceOf(address(_getLockBox(0)));
  }

  // ================================================================
  // │                      Chain Management                        │
  // ================================================================

  /// @notice Returns whether the tokens locked for a given remote chain should be siloed independently
  /// from all other remote chains.
  /// @param remoteChainSelector the CCIP specific selector for the remote chain being interacted with.
  /// @return isSiloed Whether the funds should be isolated from all the others.
  function isSiloed(
    uint64 remoteChainSelector
  ) external view returns (bool) {
    return s_chainConfigs[remoteChainSelector].isSiloed;
  }

  /// @notice Updates designations for chains on whether to mark funds as siloed or not.
  /// @param removes A list of chain selectors to disable siloing. Their funds will be moved into the unsiloed lockbox.
  /// If a chain is not siloed, and attempted to be removed, the function will revert.
  /// @param adds A list of chain selectors to enable siloing.
  function updateSiloDesignations(
    uint64[] calldata removes,
    SiloConfigUpdate[] calldata adds
  ) external onlyOwner {
    for (uint256 i = 0; i < removes.length; ++i) {
      if (!s_chainConfigs[removes[i]].isSiloed) revert ChainNotSiloed(removes[i]);

      ERC20LockBox chainLockBox = _getLockBox(removes[i]);
      uint256 amountUnsiloed = i_token.balanceOf(address(chainLockBox));

      if (amountUnsiloed > 0) {
        chainLockBox.withdraw(removes[i], address(i_token), bytes32(uint256(removes[i])), amountUnsiloed, address(this));
        _getLockBox(0).deposit(0, address(i_token), bytes32(0), amountUnsiloed);
      }

      delete s_chainConfigs[removes[i]];

      // Emit a removal event which includes the amount of funds moved to the shared lockbox.
      emit ChainUnsiloed(removes[i], amountUnsiloed);
    }

    for (uint256 i = 0; i < adds.length; ++i) {
      // Since the zero chain selector is used to designate unsiloed chains, it should never be used for siloed chains.
      if (
        adds[i].remoteChainSelector == 0 || s_chainConfigs[adds[i].remoteChainSelector].isSiloed
          || !isSupportedChain(adds[i].remoteChainSelector)
      ) {
        revert InvalidChainSelector(adds[i].remoteChainSelector);
      }

      if (adds[i].rebalancer == address(0)) revert ZeroAddressInvalid();

      if (address(s_lockBoxes[adds[i].remoteChainSelector]) == address(0)) {
        revert LockBoxNotConfigured(bytes32(uint256(adds[i].remoteChainSelector)));
      }
      s_chainConfigs[adds[i].remoteChainSelector] = SiloConfig({rebalancer: adds[i].rebalancer, isSiloed: true});

      emit ChainSiloed(adds[i].remoteChainSelector, adds[i].rebalancer);
    }
  }

  /// @notice Gets the rebalancer able to provide liquidity for a remote chain selector
  /// @param remoteChainSelector The CCIP specific selector for the remote chain being interacted with.
  /// @return The current liquidity manager for the given siloed chain, or the unsiloed rebalancer if the chain is not siloed.
  function getChainRebalancer(
    uint64 remoteChainSelector
  ) public view returns (address) {
    SiloConfig storage remoteConfig = s_chainConfigs[remoteChainSelector];
    if (remoteConfig.isSiloed) {
      return remoteConfig.rebalancer;
    }

    return s_rebalancer;
  }

  /// @notice Gets the rebalancer for the unsiloed chains.
  /// @return The current liquidity manager for the unsiloed chains.
  function getRebalancer() external view returns (address) {
    return s_rebalancer;
  }

  /// @notice Sets the Rebalancer address for a given remoteChainSelector.
  /// @dev Only callable by the owner.
  /// @param remoteChainSelector the remote chain to set.
  /// @param newRebalancer the address allowed to add liquidity for the given siloed chain.
  function setSiloRebalancer(
    uint64 remoteChainSelector,
    address newRebalancer
  ) external onlyOwner {
    SiloConfig storage remoteConfig = s_chainConfigs[remoteChainSelector];

    if (!remoteConfig.isSiloed) revert ChainNotSiloed(remoteChainSelector);

    address oldRebalancer = remoteConfig.rebalancer;

    remoteConfig.rebalancer = newRebalancer;

    emit SiloRebalancerSet(remoteChainSelector, oldRebalancer, newRebalancer);
  }

  /// @notice Sets the Rebalancer address for unsiloed chains.
  /// @dev Only callable by the owner.
  /// @param newRebalancer the address allowed to add liquidity for the given siloed chain.
  function setRebalancer(
    address newRebalancer
  ) external onlyOwner {
    address oldRebalancer = s_rebalancer;

    s_rebalancer = newRebalancer;

    emit UnsiloedRebalancerSet(oldRebalancer, newRebalancer);
  }

  // ================================================================
  // │                    Provide Liquidity                         │
  // ================================================================

  /// @notice Adds liquidity to the pool. The tokens should be approved first.
  /// @param remoteChainSelector the remote chain to set. If the chain is not siloed, the liquidity will be shared among all
  /// non-siloed chains.
  /// @param amount The amount of liquidity to provide.
  /// @dev Only the rebalancer for the chain can add liquidity
  function provideSiloedLiquidity(
    uint64 remoteChainSelector,
    uint256 amount
  ) external virtual {
    if (!s_chainConfigs[remoteChainSelector].isSiloed || remoteChainSelector == 0) {
      revert ChainNotSiloed(remoteChainSelector);
    }
    _provideLiquidity(remoteChainSelector, amount);
  }

  /// @notice Adds liquidity to the pool for unsiloed chains. Function is used to support legacy liquidity operations
  /// by using a function selector available to previous L/R pools.
  /// @dev Since the remoteChainSelector 0 should never be applied to a real chain, it is used to designate unsiloed chains.
  /// @param amount The amount of liquidity to provide.
  function provideLiquidity(
    uint256 amount
  ) external virtual {
    // The zero chain selector is used to designate unsiloed chains, so hard coding it in allows for a more efficient
    // implementation where both liquidity functions can use the same internal function but with different external
    // functions for liquidity providers.
    _provideLiquidity(0, amount);
  }

  function _provideLiquidity(
    uint64 remoteChainSelector,
    uint256 amount
  ) internal virtual {
    if (amount == 0) revert LiquidityAmountCannotBeZero();
    if (msg.sender != getChainRebalancer(remoteChainSelector)) revert Unauthorized(msg.sender);

    uint64 lockBoxSelector = s_chainConfigs[remoteChainSelector].isSiloed ? remoteChainSelector : 0;

    i_token.safeTransferFrom(msg.sender, address(this), amount);
    _getLockBox(lockBoxSelector)
      .deposit(remoteChainSelector, address(i_token), bytes32(uint256(lockBoxSelector)), amount);

    emit LiquidityAdded(remoteChainSelector, msg.sender, amount);
  }

  // ================================================================
  // │                    Withdraw Liquidity                        │
  // ================================================================

  /// @notice Removes liquidity from the pool for unsiloed chains. Function is used to support legacy liquidity operations
  /// by using a function selector available to previous L/R pools.
  /// @dev Since the remoteChainSelector 0 should never be applied to a real chain, it is used to designate unsiloed chains.
  /// @param amount The amount of liquidity to remove.
  function withdrawLiquidity(
    uint256 amount
  ) external {
    // The zero chain selector is used to designate unsiloed chains, so hard coding it in allows for a more efficient
    // implementation where both liquidity functions can use the same internal function but with different external
    // functions for liquidity providers.
    _withdrawLiquidity(0, amount);
  }

  /// @notice Removed liquidity to the pool. The tokens will be sent to msg.sender.
  /// @dev Only the rebalancer can remove liquidity from the contract, for both siloed and unsiloed chains.
  /// @param remoteChainSelector the remote chain to set. If the chain is not siloed, then no accounting will be updated,
  /// which can be considered the liquidity for all non-siloed chains sharing liquidity.
  /// @param amount The amount of liquidity to remove.
  function withdrawSiloedLiquidity(
    uint64 remoteChainSelector,
    uint256 amount
  ) external {
    // The zero chain selector is used to designate unsiloed chains, and should never be used for siloed chains,
    // so we revert instead of proceeding.
    if (!s_chainConfigs[remoteChainSelector].isSiloed || remoteChainSelector == 0) {
      revert ChainNotSiloed(remoteChainSelector);
    }

    _withdrawLiquidity(remoteChainSelector, amount);
  }

  function _withdrawLiquidity(
    uint64 remoteChainSelector,
    uint256 amount
  ) internal {
    if (amount == 0) revert LiquidityAmountCannotBeZero();
    if (msg.sender != getChainRebalancer(remoteChainSelector)) revert Unauthorized(msg.sender);

    uint64 lockBoxSelector = s_chainConfigs[remoteChainSelector].isSiloed ? remoteChainSelector : 0;

    uint256 availableLiquidity = i_token.balanceOf(address(_getLockBox(lockBoxSelector)));
    if (amount > availableLiquidity) revert InsufficientLiquidity(availableLiquidity, amount);

    // Withdraw the tokens directly from the lockbox to the rebalancer. This saves gas by avoiding the need to transfer
    // the tokens to the contract first.
    _getLockBox(lockBoxSelector)
      .withdraw(remoteChainSelector, address(i_token), bytes32(uint256(lockBoxSelector)), amount, msg.sender);

    emit LiquidityRemoved(remoteChainSelector, msg.sender, amount);
  }

  /// @notice No-op override to purge the unused code path from the contract.
  function _postFlightCheck(
    Pool.ReleaseOrMintInV1 calldata,
    uint256,
    uint16
  ) internal pure virtual override {}

  /// @notice No-op override to purge the unused code path from the contract.
  function _preFlightCheck(
    Pool.LockOrBurnInV1 calldata,
    uint16,
    bytes memory
  ) internal pure virtual override {}

  /// @notice Configure lockboxes for liquidity domains using a struct to avoid array misconfigs.
  function configureLockBoxes(
    LockBoxConfig[] calldata lockBoxConfigs
  ) external onlyOwner {
    for (uint256 i = 0; i < lockBoxConfigs.length; ++i) {
      address lockBox = lockBoxConfigs[i].lockBox;
      bytes32 liquidityDomainId = bytes32(uint256(lockBoxConfigs[i].remoteChainSelector));
      if (lockBox == address(0)) revert ZeroAddressInvalid();
      ERC20LockBox erc20LockBox = ERC20LockBox(lockBox);
      if (address(erc20LockBox.getToken()) != address(i_token)) {
        revert InvalidToken(address(erc20LockBox.getToken()));
      }
      if (erc20LockBox.getLiquidityDomainId() != liquidityDomainId) {
        revert InvalidLockBoxLiquidityDomain(erc20LockBox.getLiquidityDomainId());
      }
      s_lockBoxes[lockBoxConfigs[i].remoteChainSelector] = erc20LockBox;
      // Reset then set to max to avoid sticky non-zero allowance issues.
      i_token.safeApprove(lockBox, 0);
      i_token.safeApprove(lockBox, type(uint256).max);
    }
  }

  function _getLockBox(
    uint64 remoteChainSelector
  ) internal view returns (ERC20LockBox) {
    ERC20LockBox lockBox = s_lockBoxes[remoteChainSelector];
    if (address(lockBox) == address(0)) revert LockBoxNotConfigured(bytes32(uint256(remoteChainSelector)));
    return lockBox;
  }
}
