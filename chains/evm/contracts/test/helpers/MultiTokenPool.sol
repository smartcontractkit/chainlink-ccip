// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolV1} from "../../interfaces/IPool.sol";

import {Pool} from "../../libraries/Pool.sol";
import {TokenPool} from "../../pools/TokenPool.sol";

import {IERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/IERC20.sol";
import {EnumerableSet} from "@openzeppelin/contracts@5.3.0/utils/structs/EnumerableSet.sol";

/// @notice This contract is a proof of concept and should NOT be used in production.
/// @dev All tokens share all config, including rate limits. The contract also assumes all remote
/// tokens have the same decimals. Decimals can be different per chain, but not per token.
/// @dev This contract has quite a few inherited functions that are focussed on a single token. This is not
/// recommended for production use, but since this is a PoC, it's acceptable.
abstract contract MultiTokenPool is TokenPool {
  using EnumerableSet for EnumerableSet.AddressSet;

  /// @dev The IERC20 token that this pool supports
  EnumerableSet.AddressSet internal s_tokens;
  mapping(address token => mapping(uint64 remoteChainSelector => bytes remoteToken)) internal s_remoteTokens;

  constructor(
    IERC20[] memory tokens,
    uint8 localTokenDecimals,
    address rmnProxy,
    address router
  ) TokenPool(tokens[0], localTokenDecimals, address(0), rmnProxy, router) {
    for (uint256 i = 0; i < tokens.length; ++i) {
      s_tokens.add(address(tokens[i]));
    }
  }

  function lockOrBurn(
    Pool.LockOrBurnInV1 calldata lockOrBurnIn,
    uint16 blockConfirmationRequested,
    bytes calldata tokenArgs
  ) public virtual override returns (Pool.LockOrBurnOutV1 memory out, uint256 destTokenAmount) {
    (out, destTokenAmount) = super.lockOrBurn(lockOrBurnIn, blockConfirmationRequested, tokenArgs);

    // The token-less `_lockOrBurn` is a no-op by default so calling it in the super call is safe.
    // Really lock/burn the tokens here.
    _lockOrBurn(lockOrBurnIn.localToken, lockOrBurnIn.remoteChainSelector, lockOrBurnIn.amount);

    // Override the dest token address as the base pool assumed a single token.
    out.destTokenAddress = getRemoteToken(lockOrBurnIn.localToken, lockOrBurnIn.remoteChainSelector);

    return (out, destTokenAmount);
  }

  /// @dev this would call into lockOrBurn if tokenArgs did not require calldata.
  function lockOrBurn(
    Pool.LockOrBurnInV1 calldata lockOrBurnIn
  ) public virtual override returns (Pool.LockOrBurnOutV1 memory out) {
    out = super.lockOrBurn(lockOrBurnIn);

    // The token-less `_lockOrBurn` is a no-op by default so calling it in the super call is safe.
    // Really lock/burn the tokens here.
    _lockOrBurn(lockOrBurnIn.localToken, lockOrBurnIn.remoteChainSelector, lockOrBurnIn.amount);

    // Override the dest token address as the base pool assumed a single token.
    out.destTokenAddress = getRemoteToken(lockOrBurnIn.localToken, lockOrBurnIn.remoteChainSelector);

    return out;
  }

  function _lockOrBurn(
    address token,
    uint64 remoteChainSelector,
    uint256 amount
  ) internal virtual {}

  function releaseOrMint(
    Pool.ReleaseOrMintInV1 calldata releaseOrMintIn,
    uint16 blockConfirmationRequested
  ) public virtual override returns (Pool.ReleaseOrMintOutV1 memory out) {
    out = super.releaseOrMint(releaseOrMintIn, blockConfirmationRequested);

    // The token-less `_releaseOrMint` is a no-op by default so calling it in the super call is safe.
    // Really release/mint the tokens here.
    _releaseOrMint(
      releaseOrMintIn.localToken,
      releaseOrMintIn.receiver,
      releaseOrMintIn.sourceDenominatedAmount,
      releaseOrMintIn.remoteChainSelector
    );

    return out;
  }

  function _releaseOrMint(
    address token,
    address receiver,
    uint256 amount,
    uint64 remoteChainSelector
  ) internal virtual {}

  // ================================================================
  // │                     Multi-token config                       │
  // ================================================================

  /// @inheritdoc IPoolV1
  function isSupportedToken(
    address token
  ) public view virtual override returns (bool) {
    return s_tokens.contains(token);
  }

  /// @notice Gets the IERC20 token that this pool can lock or burn.
  /// @return tokens The IERC20 token representation.
  function getTokens() public view returns (IERC20[] memory tokens) {
    tokens = new IERC20[](s_tokens.length());
    for (uint256 i = 0; i < s_tokens.length(); ++i) {
      tokens[i] = IERC20(s_tokens.at(i));
    }
    return tokens;
  }

  /// @notice Gets the token address on the remote chain.
  /// @param remoteChainSelector Remote chain selector.
  /// @dev To support non-evm chains, this value is encoded into bytes
  function getRemoteToken(
    address token,
    uint64 remoteChainSelector
  ) public view returns (bytes memory) {
    return s_remoteTokens[token][remoteChainSelector];
  }
}
