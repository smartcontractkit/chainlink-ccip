// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ICrossChainVerifierResolver} from "../../interfaces/ICrossChainVerifierResolver.sol";
import {IPoolV1} from "../../interfaces/IPool.sol";
import {IPoolV2} from "../../interfaces/IPoolV2.sol";
import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";

import {CCTPVerifier} from "../../ccvs/CCTPVerifier.sol";
import {Pool} from "../../libraries/Pool.sol";
import {USDCSourcePoolDataCodec} from "../../libraries/USDCSourcePoolDataCodec.sol";
import {TokenPool} from "../TokenPool.sol";

import {IERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/IERC20.sol";
import {ERC165Checker} from "@openzeppelin/contracts@5.3.0/utils/introspection/ERC165Checker.sol";

/// @notice CCTP token pool that delegates minting and burning responsibilities of USDC to the CCTPVerifier contract.
/// @dev This pool does not mutate the token state. It does not actually burn USDC via TokenMessenger on source or mint via MessageTransmitter on destination.
/// It remains responsible for rate limiting and other validations while outsourcing token management to the CCTPVerfier contract.
/// This token pool should never have a balance of USDC at any point during a transaction, otherwise funds will be lost.
/// The caller of lockOrBurn is responsible for sending USDC to the CCTPVerifier contract instead, which this contract points to.
contract CCTPTokenPool is TokenPool, ITypeAndVersion {
  using ERC165Checker for address;

  error InvalidCCTPVerifier(address cctpVerifier);
  error InboundImplementationNotFoundForVerifier(bytes4 verifierVersion);
  error OutboundImplementationNotFoundForVerifier(uint64 remoteChainSelector);

  string public constant override typeAndVersion = "CCTPTokenPool 1.7.0-dev";

  /// @notice The CCTP verifier contract.
  /// @dev Immutable because the address should correspond to a static proxy contract.
  address internal immutable i_cctpVerifier;

  constructor(
    IERC20 token,
    uint8 localTokenDecimals,
    address advancedPoolHooks,
    address rmnProxy,
    address router,
    address cctpVerifier
  ) TokenPool(token, localTokenDecimals, advancedPoolHooks, rmnProxy, router) {
    if (!cctpVerifier.supportsInterface(type(ICrossChainVerifierResolver).interfaceId)) {
      revert InvalidCCTPVerifier(cctpVerifier);
    }

    i_cctpVerifier = cctpVerifier;
  }

  /// @inheritdoc IPoolV2
  /// @dev The _validateLockOrBurn check is an essential security check.
  /// @dev The _applyFee function deducts the fee from the amount and returns the amount after fee deduction.
  /// @dev The call to _lockOrBurn is omitted because this pool is not responsible for token management.
  /// LockedOrBurned is still emitted for consumers that expect it.
  function lockOrBurn(
    Pool.LockOrBurnInV1 calldata lockOrBurnIn,
    uint16 blockConfirmationRequested,
    bytes calldata tokenArgs
  ) public virtual override returns (Pool.LockOrBurnOutV1 memory, uint256 destTokenAmount) {
    _validateLockOrBurn(lockOrBurnIn, blockConfirmationRequested, tokenArgs);
    destTokenAmount = _applyFee(lockOrBurnIn, blockConfirmationRequested);

    emit LockedOrBurned({
      remoteChainSelector: lockOrBurnIn.remoteChainSelector,
      token: address(i_token),
      sender: msg.sender,
      amount: destTokenAmount
    });

    address verifierImpl =
      ICrossChainVerifierResolver(i_cctpVerifier).getOutboundImplementation(lockOrBurnIn.remoteChainSelector, "");
    if (verifierImpl == address(0)) {
      revert OutboundImplementationNotFoundForVerifier(lockOrBurnIn.remoteChainSelector);
    }

    return (
      Pool.LockOrBurnOutV1({
        destTokenAddress: getRemoteToken(lockOrBurnIn.remoteChainSelector),
        destPoolData: USDCSourcePoolDataCodec._encodeSourceTokenDataPayloadV2WithCCV(
          CCTPVerifier(verifierImpl).versionTag()
        )
      }),
      destTokenAmount
    );
  }

  /// @inheritdoc IPoolV1
  /// @dev The _validateLockOrBurn check is an essential security check.
  /// @dev _applyFee is not called in this legacy method, so the full amount is locked or burned.
  /// @dev The call to _lockOrBurn is omitted because this pool is not responsible for token management.
  /// LockedOrBurned is still emitted for consumers that expect it.
  function lockOrBurn(
    Pool.LockOrBurnInV1 calldata lockOrBurnIn
  ) public virtual override returns (Pool.LockOrBurnOutV1 memory lockOrBurnOutV1) {
    _validateLockOrBurn(lockOrBurnIn, WAIT_FOR_FINALITY, "");

    emit LockedOrBurned({
      remoteChainSelector: lockOrBurnIn.remoteChainSelector,
      token: address(i_token),
      sender: msg.sender,
      amount: lockOrBurnIn.amount
    });

    address verifierImpl =
      ICrossChainVerifierResolver(i_cctpVerifier).getOutboundImplementation(lockOrBurnIn.remoteChainSelector, "");
    if (verifierImpl == address(0)) {
      revert OutboundImplementationNotFoundForVerifier(lockOrBurnIn.remoteChainSelector);
    }

    return Pool.LockOrBurnOutV1({
      destTokenAddress: getRemoteToken(lockOrBurnIn.remoteChainSelector),
      destPoolData: USDCSourcePoolDataCodec._encodeSourceTokenDataPayloadV2WithCCV(
        CCTPVerifier(verifierImpl).versionTag()
      )
    });
  }

  /// @inheritdoc IPoolV2
  /// @dev The _validateReleaseOrMint check is an essential security check.
  /// @dev The call to _releaseOrMint is omitted because this pool is not responsible for token management.
  /// ReleasedOrMinted is still emitted for consumers that expect it.
  function releaseOrMint(
    Pool.ReleaseOrMintInV1 calldata releaseOrMintIn,
    uint16 blockConfirmationRequested
  ) public virtual override returns (Pool.ReleaseOrMintOutV1 memory) {
    _validateReleaseOrMint(releaseOrMintIn, releaseOrMintIn.sourceDenominatedAmount, blockConfirmationRequested);

    emit ReleasedOrMinted({
      remoteChainSelector: releaseOrMintIn.remoteChainSelector,
      token: address(i_token),
      sender: msg.sender,
      recipient: releaseOrMintIn.receiver,
      amount: releaseOrMintIn.sourceDenominatedAmount
    });

    // We return 0 because this pool doesn't actually mint tokens, the CCTPVerifier contract does.
    // Returning a value here would cause the OffRamp to revert with ReleaseOrMintBalanceMismatch.
    // The CCTPVerifier contract is responsible for validating that receiver balance has increased by the correct amount.
    return Pool.ReleaseOrMintOutV1({destinationAmount: 0});
  }

  /// @inheritdoc IPoolV1
  /// @dev calls IPoolV2.releaseOrMint with default finality.
  /// @dev The call to _releaseOrMint is omitted because this pool is not responsible for token management.
  /// ReleasedOrMinted is still emitted for consumers that expect it.
  function releaseOrMint(
    Pool.ReleaseOrMintInV1 calldata releaseOrMintIn
  ) public virtual override returns (Pool.ReleaseOrMintOutV1 memory) {
    return releaseOrMint(releaseOrMintIn, WAIT_FOR_FINALITY);
  }

  /// @notice Returns the CCTP verifier contract.
  /// @return cctpVerifier The CCTP verifier contract.
  function getCCTPVerifier() external view returns (address) {
    return i_cctpVerifier;
  }
}
