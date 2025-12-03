// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolV1} from "../../interfaces/IPool.sol";
import {IPoolV2} from "../../interfaces/IPoolV2.sol";
import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";

import {Pool} from "../../libraries/Pool.sol";
import {USDCSourcePoolDataCodec} from "../../libraries/USDCSourcePoolDataCodec.sol";
import {TokenPool} from "../TokenPool.sol";

import {IERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/IERC20.sol";

/// @notice CCTP token pool that delegates minting and burning responsibilities of USDC to the CCTPVerifier contract.
/// @dev This pool does not mutate the token state. It does not actually burn USDC via TokenMessenger on source or mint via MessageTransmitter on destination.
/// It remains responsible for rate limiting and other validations while outsourcing token management to the CCTPVerfier contract.
/// This token pool should never have a balance of USDC at any point during a transaction, otherwise funds will be lost.
/// The caller of lockOrBurn is responsible for sending USDC to the CCTPVerifier contract instead.
contract CCTPTokenPool is TokenPool, ITypeAndVersion {
  error InboundImplementationNotFoundForVerifier(bytes4 ccvVersionTag);
  error OutboundImplementationNotFoundForVerifier(uint64 remoteChainSelector);

  string public constant override typeAndVersion = "CCTPTokenPool 1.7.0-dev";

  constructor(
    IERC20 token,
    uint8 localTokenDecimals,
    address advancedPoolHooks,
    address rmnProxy,
    address router
  ) TokenPool(token, localTokenDecimals, advancedPoolHooks, rmnProxy, router) {}

  /// @inheritdoc IPoolV2
  function lockOrBurn(
    Pool.LockOrBurnInV1 calldata lockOrBurnIn,
    uint16 blockConfirmationRequested,
    bytes calldata tokenArgs
  ) public virtual override returns (Pool.LockOrBurnOutV1 memory, uint256 destTokenAmount) {
    return (_lockOrBurn(lockOrBurnIn, blockConfirmationRequested, tokenArgs), lockOrBurnIn.amount);
  }

  /// @inheritdoc IPoolV1
  function lockOrBurn(
    Pool.LockOrBurnInV1 calldata lockOrBurnIn
  ) public virtual override returns (Pool.LockOrBurnOutV1 memory lockOrBurnOutV1) {
    return _lockOrBurn(lockOrBurnIn, WAIT_FOR_FINALITY, "");
  }

  /// @dev The _validateLockOrBurn check is an essential security check.
  /// @dev The call to _lockOrBurn(amount) is omitted because this pool is not responsible for token management.
  /// LockedOrBurned is still emitted for consumers that expect it.
  function _lockOrBurn(
    Pool.LockOrBurnInV1 calldata lockOrBurnIn,
    uint16 blockConfirmationRequested,
    bytes memory tokenArgs
  ) internal virtual returns (Pool.LockOrBurnOutV1 memory lockOrBurnOutV1) {
    _validateLockOrBurn(lockOrBurnIn, blockConfirmationRequested, tokenArgs);

    emit LockedOrBurned({
      remoteChainSelector: lockOrBurnIn.remoteChainSelector,
      token: address(i_token),
      sender: msg.sender,
      amount: lockOrBurnIn.amount
    });

    return Pool.LockOrBurnOutV1({
      destTokenAddress: getRemoteToken(lockOrBurnIn.remoteChainSelector),
      destPoolData: USDCSourcePoolDataCodec._encodeSourceTokenDataPayloadV2WithCCV()
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
      // The CCTP verifier will mint some of the amount to the fee recipient and some to the receiver.
      // This event simply declares that the full amount has been minted, with no awareness of the split.
      // This token pool can't be aware of the split without requiring the CCTP verifier to write feeExecuted to storage every tx.
      // Additionally, the LockedOrBurned event on source includes the full amount, providing parity for offchain consumers.
      amount: releaseOrMintIn.sourceDenominatedAmount
    });

    return Pool.ReleaseOrMintOutV1({destinationAmount: releaseOrMintIn.sourceDenominatedAmount});
  }

  /// @inheritdoc IPoolV1
  /// @dev Calls IPoolV2.releaseOrMint with default finality.
  function releaseOrMint(
    Pool.ReleaseOrMintInV1 calldata releaseOrMintIn
  ) public virtual override returns (Pool.ReleaseOrMintOutV1 memory) {
    return releaseOrMint(releaseOrMintIn, WAIT_FOR_FINALITY);
  }
}
