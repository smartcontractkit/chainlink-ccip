// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";
import {ICrossChainVerifierResolver} from "../../interfaces/ICrossChainVerifierResolver.sol";

import {CCTPVerifier} from "../../ccvs/CCTPVerifier.sol";
import {TokenPool} from "./TokenPool.sol";
import {Pool} from "../../libraries/Pool.sol";
import {USDCSourcePoolDataCodec} from "../../libraries/USDCSourcePoolDataCodec.sol";

import {IERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/IERC20.sol";
import {ERC165Checker} from "@openzeppelin/contracts@5.0.2/utils/introspection/ERC165Checker.sol";

/// @notice CCTP token pool that delegates minting and burning responsibilities of USDC to the CCTPVerifier contract.
/// @dev This pool does not mutate the token state. It does not actually burn USDC via TokenMessenger on source or mint via MessageTransmitter on destination.
/// It remains responsible for rate limiting and other validations while outsourcing token management to the CCTPVerfier contract.
/// This token pool should never have a balance of USDC at any point during a transaction, otherwise funds will be lost.
/// The caller of lockOrBurn is responsible for sending USDC to the CCTPVerifier contract instead.
contract CCTPTokenPool is TokenPool, ITypeAndVersion {
  using ERC165Checker for address;

  error InboundImplementationNotFoundForVerifier(uint64 remoteChainSelector);
  error InvalidCCTPVerifier(address cctpVerifier);

  string public constant override typeAndVersion = "CCTPTokenPool 1.7.0-dev";

  /// @notice The CCTP verifier contract.
  /// @dev Immutable because the address should correspond to a static proxy contract.
  address internal immutable i_cctpVerifier;

  constructor(
    IERC20 token,
    uint8 localTokenDecimals,
    address[] memory allowlist,
    address rmnProxy,
    address router,
    address cctpVerifier
  ) TokenPool(token, localTokenDecimals, allowlist, rmnProxy, router) {
    if (!cctpVerifier.supportsInterface(type(ICrossChainVerifierResolver).interfaceId)) {
      revert InvalidCCTPVerifier(cctpVerifier);
    }
    
    i_cctpVerifier = cctpVerifier;
  }

  /// @inheritdoc IPoolV2
  /// @dev The _validateLockOrBurn check is an essential security check.
  /// @dev The _applyFee function deducts the fee from the amount and returns the amount after fee deduction.
  /// @dev The call to _lockOrBurn is omitted because this pool is not responsible for token management.
  /// @dev LockedOrBurned is still emitted for consumers that expect it.
  function lockOrBurn(
    Pool.LockOrBurnInV1 calldata lockOrBurnIn,
    uint16 blockConfirmationRequested,
    bytes memory // tokenArgs
  ) public virtual returns (Pool.LockOrBurnOutV1 memory, uint256 destTokenAmount) {
    _validateLockOrBurn(lockOrBurnIn, blockConfirmationRequested);
    destTokenAmount = _applyFee(lockOrBurnIn, blockConfirmationRequested);

    emit LockedOrBurned({
      remoteChainSelector: lockOrBurnIn.remoteChainSelector,
      token: address(i_token),
      sender: msg.sender,
      amount: destTokenAmount
    });

    return (
      Pool.LockOrBurnOutV1({
        destTokenAddress: getRemoteToken(lockOrBurnIn.remoteChainSelector),
        destPoolData: abi.encodePacked(USDCSourcePoolDataCodec.CCTP_VERSION_2_CCV_TAG)
      }),
      destTokenAmount
    );
  }

  /// @inheritdoc IPoolV1
  /// @dev The _validateLockOrBurn check is an essential security check.
  /// @dev _applyFee is not called in this legacy method, so the full amount is locked or burned.
  /// @dev The call to _lockOrBurn is omitted because this pool is not responsible for token management.
  /// @dev LockedOrBurned is still emitted for consumers that expect it.
  function lockOrBurn(
    Pool.LockOrBurnInV1 calldata lockOrBurnIn
  ) public virtual returns (Pool.LockOrBurnOutV1 memory lockOrBurnOutV1) {
    _validateLockOrBurn(lockOrBurnIn, WAIT_FOR_FINALITY);

    emit LockedOrBurned({
      remoteChainSelector: lockOrBurnIn.remoteChainSelector,
      token: address(i_token),
      sender: msg.sender,
      amount: lockOrBurnIn.amount
    });

    return Pool.LockOrBurnOutV1({
      destTokenAddress: getRemoteToken(lockOrBurnIn.remoteChainSelector),
      destPoolData: abi.encodePacked(USDCSourcePoolDataCodec.CCTP_VERSION_2_CCV_TAG)
    });
  }

  /// @inheritdoc IPoolV2
  /// @dev The _validateReleaseOrMint check is an essential security check.
  /// @dev The call to _releaseOrMint is omitted because this pool is not responsible for token management.
  /// @dev ReleasedOrMinted is still emitted for consumers that expect it.
  function releaseOrMint(
    Pool.ReleaseOrMintInV1 calldata releaseOrMintIn,
    uint16 blockConfirmationRequested
  ) public virtual override(IPoolV2) returns (Pool.ReleaseOrMintOutV1 memory) {
    uint256 localAmount = _calculateLocalAmount(
      releaseOrMintIn.sourceDenominatedAmount, _parseRemoteDecimals(releaseOrMintIn.sourcePoolData)
    );

    _validateReleaseOrMint(releaseOrMintIn, localAmount, blockConfirmationRequested);

    emit ReleasedOrMinted({
      remoteChainSelector: releaseOrMintIn.remoteChainSelector,
      token: address(i_token),
      sender: msg.sender,
      recipient: releaseOrMintIn.receiver,
      amount: localAmount
    });

    return Pool.ReleaseOrMintOutV1({destinationAmount: localAmount});
  }

  /// @inheritdoc IPoolV1
  /// @dev calls IPoolV2.releaseOrMint with default finality.
  /// @dev The call to _releaseOrMint is omitted because this pool is not responsible for token management.
  /// @dev ReleasedOrMinted is still emitted for consumers that expect it.
  function releaseOrMint(
    Pool.ReleaseOrMintInV1 calldata releaseOrMintIn
  ) public virtual override returns (Pool.ReleaseOrMintOutV1 memory) {
    return releaseOrMint(releaseOrMintIn, WAIT_FOR_FINALITY);
  }

  /// @notice Properly quotes the amount received on destination based on the requested finality.
  /// @dev Uses the CCTPVerifier contract as the source of truth for fast finality bps.
  /// @param lockOrBurnIn The original lock or burn request.
  /// @param blockConfirmationRequested The minimum block confirmation requested by the message.
  /// A value of zero (WAIT_FOR_FINALITY) applies default finality fees.
  /// @return destAmount The amount received on destination after fee deduction.
  /// CCTP deducts fees on destination rather than on source.
  function _applyFee(
    Pool.LockOrBurnInV1 calldata lockOrBurnIn,
    uint16 blockConfirmationRequested
  ) internal view virtual returns (uint256 destAmount) {
    address verifierImpl = ICrossChainVerifierResolver(i_cctpVerifier).getOutboundImplementation(lockOrBurnIn.remoteChainSelector);

    if (verifierImpl == address(0)) {
      revert InboundImplementationNotFoundForVerifier(lockOrBurnIn.remoteChainSelector);
    }

    // Standard finality transfers are not subject to fees.
    // Therefore, the token amount received on destination equals the amount sent on source.
    if (blockConfirmationRequested == WAIT_FOR_FINALITY) {
      return lockOrBurnIn.amount;
    }
    
    // Otherwise, we use the verifier contract as the source of truth for fast finality bps.
    return lockOrBurnIn.amount - (lockOrBurnIn.amount * CCTPVerifier(verifierImpl).getDynamicConfig().fastFinalityBps) / BPS_DIVIDER;
  }

  /// @notice Returns the CCTP verifier contract.
  /// @return cctpVerifier The CCTP verifier contract.
  function getCCTPVerifier() external view returns (address) {
    return i_cctpVerifier;
  }
}
