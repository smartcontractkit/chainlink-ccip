// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";

import {ICrossChainVerifierResolver} from "../../interfaces/ICrossChainVerifierResolver.sol";
import {Pool} from "../../libraries/Pool.sol";
import {TokenPool} from "../TokenPool.sol";

import {IERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/IERC20.sol";
import {IERC20Metadata} from "@openzeppelin/contracts@4.8.3/token/ERC20/extensions/IERC20Metadata.sol";
import {SafeERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/utils/SafeERC20.sol";

/// @notice Lombard CCIP token pool.
/// For v2 flows, token movement (burn/mint or bridging) is handled by the Lombard verifier,
/// the pool performs validation, rate limiting, accounting and event emission.
/// IPoolV2.lockOrBurn forwards tokens to the verifier with _lockOrBurn.
/// IPoolV2.releaseOrMint does not move tokens, _releaseOrMint is a no-op.
/// TODO: Add explicit V1 support/backwards compatibility.
contract LombardTokenPool is TokenPool, ITypeAndVersion {
  using SafeERC20 for IERC20;
  using SafeERC20 for IERC20Metadata;

  error ZeroVerifierNotAllowed();
  error OutboundImplementationNotFoundForVerifier();

  event LombardVerifierSet(address indexed verifier);

  /// @notice CCIP contract type and version.
  string public constant override typeAndVersion = "LombardTokenPool 1.7.0-dev";

  /// @notice Lombard verifier proxy / resolver address. lockOrBurn fetches the outbound implementation and forwards tokens to it.
  address private immutable i_lombardVerifierResolver;

  constructor(
    IERC20Metadata token,
    address verifier,
    address advancedPoolHooks,
    address rmnProxy,
    address router,
    uint8 fallbackDecimals
  ) TokenPool(token, _getTokenDecimals(token, fallbackDecimals), advancedPoolHooks, rmnProxy, router) {
    if (verifier == address(0)) {
      revert ZeroVerifierNotAllowed();
    }
    i_lombardVerifierResolver = verifier;
    emit LombardVerifierSet(verifier);
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
    uint64 remoteChainSelector;
    // Calldata layout for lockOrBurn((bytes,uint64,address,uint256,address),uint16,bytes):
    // - 0x00-0x03: function selector
    // - 0x04: offset to the tuple (expected 0x60)
    // Within the tuple, remoteChainSelector is the second element at tupleOffset + 0x20.
    assembly {
      let tuplePtr := add(0x04, calldataload(0x04))
      remoteChainSelector := calldataload(add(tuplePtr, 0x20))
    }
    address verifierImpl =
      ICrossChainVerifierResolver(i_lombardVerifierResolver).getOutboundImplementation(remoteChainSelector, "");
    if (verifierImpl == address(0)) {
      revert OutboundImplementationNotFoundForVerifier();
    }
    i_token.safeTransfer(verifierImpl, amount);
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

  /// @notice Returns the verifier resolver address.
  function getVerifierResolver() external view returns (address) {
    return i_lombardVerifierResolver;
  }
}
