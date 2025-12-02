// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";

import {Pool} from "../../libraries/Pool.sol";
import {TokenPool} from "../TokenPool.sol";

import {IERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/IERC20.sol";
import {IERC20Metadata} from "@openzeppelin/contracts@4.8.3/token/ERC20/extensions/IERC20Metadata.sol";
import {SafeERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/utils/SafeERC20.sol";

/// @notice Lombard CCIP token pool.
/// For v2 flows, token movement (burn/mint or bridging) is handled by the Lombard verifier,
/// the pool performs validation, rate limiting, accounting and event emission.
/// IPoolV2.lockOrBurn forwards tokens to the verifier with _lockOrBurn.
/// IPoolV2.releaseOrMint does not move tokens; it validates, consumes rate limits, and emits the event while the verifier delivers funds.
/// TODO: Add explicit V1 support/backwards compatibility.
contract LombardTokenPool is TokenPool, ITypeAndVersion {
  using SafeERC20 for IERC20;
  using SafeERC20 for IERC20Metadata;

  error ZeroVerifierNotAllowed();

  /// @notice CCIP contract type and version.
  string public constant override typeAndVersion = "LombardTokenPool 1.7.0-dev";

  /// @notice Lombard verifier that executes the cross-chain flow and handles token movement.
  address public s_verifier;

  constructor(
    IERC20Metadata token,
    address verifier,
    address advancedPoolHooks,
    address rmnProxy,
    address router,
    uint8 fallbackDecimals
  ) TokenPool(token, _getTokenDecimals(token, fallbackDecimals), advancedPoolHooks, rmnProxy, router) {
    _setVerifier(verifier);
  }

  // ================================================================
  // │                        Lock or Burn                          │
  // ================================================================

  /// @notice For IPoolV2.lockOrBurn call, this contract only overrides _lockOrBurn to forward tokens to the verifier
  /// after validation/fee handling in the base class.
  /// @dev Forward the net amount to the verifier; actual burn/bridge is done there.
  function _lockOrBurn(
    uint256 amount
  ) internal virtual override {
    i_token.safeTransfer(s_verifier, amount);
  }

  function lockOrBurn(
    Pool.LockOrBurnInV1 calldata
  ) public pure override(TokenPool) returns (Pool.LockOrBurnOutV1 memory lockOrBurnOut) {
    // TODO: Implement V1 path for backward compatability with old lanes.
    return lockOrBurnOut;
  }

  // ================================================================
  // │                      Release or Mint                         │
  // ================================================================

  function releaseOrMint(
    Pool.ReleaseOrMintInV1 calldata
  ) public pure override(TokenPool) returns (Pool.ReleaseOrMintOutV1 memory releaseOrMintOut) {
    // TODO: Implement V1 path for backward compatability with old lanes.
    return releaseOrMintOut;
  }

  // ================================================================
  // │                        Internal utils                        │
  // ================================================================

  function _getTokenDecimals(IERC20Metadata token, uint8 fallbackDecimals) internal view returns (uint8) {
    try token.decimals() returns (uint8 dec) {
      return dec;
    } catch {
      return fallbackDecimals;
    }
  }

  /// @notice Updates the verifier address used for forwarding tokens.
  /// @param verifier New verifier address.
  function setVerifier(
    address verifier
  ) external onlyOwner {
    _setVerifier(verifier);
  }

  function _setVerifier(
    address verifier
  ) internal {
    if (verifier == address(0)) {
      revert ZeroVerifierNotAllowed();
    }
    if (verifier == s_verifier) {
      return;
    }

    // Revoke old allowance if set.
    if (s_verifier != address(0)) {
      i_token.safeApprove(s_verifier, 0);
    }

    s_verifier = verifier;
    i_token.safeApprove(verifier, type(uint256).max);
  }
}
