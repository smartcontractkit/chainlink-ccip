// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolV1} from "../../interfaces/IPool.sol";
import {IPoolV2} from "../../interfaces/IPoolV2.sol";
import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";

import {CCTPVerifier} from "../../ccvs/CCTPVerifier.sol";
import {ICrossChainVerifierResolver} from "../../interfaces/ICrossChainVerifierResolver.sol";
import {Pool} from "../../libraries/Pool.sol";
import {USDCSourcePoolDataCodec} from "../../libraries/USDCSourcePoolDataCodec.sol";
import {TokenPool} from "../TokenPool.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

import {IERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/IERC20.sol";

/// @notice CCTP token pool that delegates minting and burning responsibilities of USDC to the CCTPVerifier contract.
/// @dev This pool does not mutate the token state. It does not actually burn USDC via TokenMessenger on source or mint via MessageTransmitter on destination.
/// It remains responsible for rate limiting and other validations while outsourcing token management to the CCTPVerifier contract.
/// This token pool should never have a balance of USDC at any point during a transaction, otherwise funds will be lost.
/// The caller of lockOrBurn is responsible for sending USDC to the CCTPVerifier contract instead.
contract CCTPTokenPool is TokenPool, ITypeAndVersion, AuthorizedCallers {
  string public constant override typeAndVersion = "CCTPTokenPool 1.7.0-dev";

  error IPoolV1NotSupported();

  /// @notice The CCTP verifier.
  ICrossChainVerifierResolver internal immutable i_cctpVerifier;

  constructor(
    IERC20 token,
    uint8 localTokenDecimals,
    address rmnProxy,
    address router,
    address cctpVerifier,
    address[] memory allowedCallers,
    address feeAggregator
  )
    TokenPool(token, localTokenDecimals, address(0), rmnProxy, router, feeAggregator)
    AuthorizedCallers(allowedCallers)
  {
    i_cctpVerifier = ICrossChainVerifierResolver(cctpVerifier);
  }

  /// @inheritdoc IPoolV2
  /// @dev The _validateLockOrBurn check is an essential security check.
  /// @dev The call to _lockOrBurn(amount) is omitted because this pool is not responsible for token management.
  /// LockedOrBurned is still emitted for consumers that expect it.
  /// @param lockOrBurnIn Encoded data fields for the processing of tokens on the source chain.
  /// @param blockConfirmationRequested Requested block confirmation.
  /// @param tokenArgs Additional token arguments.
  function lockOrBurn(
    Pool.LockOrBurnInV1 calldata lockOrBurnIn,
    uint16 blockConfirmationRequested,
    bytes calldata tokenArgs
  ) public virtual override returns (Pool.LockOrBurnOutV1 memory, uint256 destTokenAmount) {
    uint256 feeAmount = _getFee(lockOrBurnIn, blockConfirmationRequested);
    _validateLockOrBurn(lockOrBurnIn, blockConfirmationRequested, tokenArgs, feeAmount);

    emit LockedOrBurned({
      remoteChainSelector: lockOrBurnIn.remoteChainSelector,
      token: address(i_token),
      sender: msg.sender,
      amount: lockOrBurnIn.amount
    });

    return (
      Pool.LockOrBurnOutV1({
        destTokenAddress: getRemoteToken(lockOrBurnIn.remoteChainSelector),
        destPoolData: abi.encodePacked(USDCSourcePoolDataCodec.CCTP_VERSION_2_CCV_TAG)
      }),
      lockOrBurnIn.amount
    );
  }

  /// @inheritdoc IPoolV2
  /// @dev The _validateReleaseOrMint check is an essential security check.
  /// @dev The call to _releaseOrMint is omitted because this pool is not responsible for token management.
  /// ReleasedOrMinted is still emitted for consumers that expect it.
  /// @param releaseOrMintIn Encoded data fields for the processing of tokens on the destination chain.
  /// @param blockConfirmationRequested Requested block confirmation.
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
  function lockOrBurn(
    Pool.LockOrBurnInV1 calldata
  ) public virtual override returns (Pool.LockOrBurnOutV1 memory) {
    revert IPoolV1NotSupported();
  }

  /// @inheritdoc IPoolV1
  function releaseOrMint(
    Pool.ReleaseOrMintInV1 calldata
  ) public virtual override returns (Pool.ReleaseOrMintOutV1 memory) {
    revert IPoolV1NotSupported();
  }

  /// @inheritdoc IPoolV2
  /// @dev Uses the CCTPVerifier to determine the bps charged by CCTP on destination.
  /// getFee will not actually account for these bps. Otherwise users would be doubly charged, on source and destination.
  /// @param destChainSelector The chain selector of the destination chain.
  function getTokenTransferFeeConfig(
    address, // localToken
    uint64 destChainSelector,
    uint16, // blockConfirmationRequested,
    bytes calldata // tokenArgs
  ) external view override returns (TokenTransferFeeConfig memory feeConfig) {
    TokenTransferFeeConfig memory transferFeeConfig = s_tokenTransferFeeConfig[destChainSelector];

    address verifierImpl = i_cctpVerifier.getOutboundImplementation(destChainSelector, "");
    CCTPVerifier.DynamicConfig memory dynamicConfig = CCTPVerifier(verifierImpl).getDynamicConfig();
    transferFeeConfig.customBlockConfirmationTransferFeeBps = dynamicConfig.fastFinalityBps;

    return transferFeeConfig;
  }

  /// @notice Validates the caller of lockOrBurn against a set of allowed callers.
  /// @dev Overrides the default behavior of _onlyOnRamp because this contract may be invoked by a proxy contract.
  /// @param remoteChainSelector The remote chain selector to validate the caller against.
  function _onlyOnRamp(
    uint64 remoteChainSelector
  ) internal view virtual override {
    if (!isSupportedChain(remoteChainSelector)) revert ChainNotAllowed(remoteChainSelector);
    _validateCaller();
  }

  /// @notice Validates the caller of releaseOrMint against a set of allowed callers.
  /// @dev Overrides the default behavior of _onlyOffRamp because this contract may be invoked by a proxy contract.
  /// @param remoteChainSelector The remote chain selector to validate the caller against.
  function _onlyOffRamp(
    uint64 remoteChainSelector
  ) internal view virtual override {
    if (!isSupportedChain(remoteChainSelector)) revert ChainNotAllowed(remoteChainSelector);
    _validateCaller();
  }

  /// @notice Returns the CCTP verifier.
  /// @return cctpVerifier The CCTP verifier.
  function getCCTPVerifier() external view returns (address) {
    return address(i_cctpVerifier);
  }
}
