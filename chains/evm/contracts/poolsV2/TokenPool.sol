// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolV2} from "../interfaces/IPoolV2.sol";

import {Client} from "../libraries/Client.sol";
import {Pool} from "../libraries/Pool.sol";
import {TokenPool as TokenPoolV1} from "../pools/TokenPool.sol";

import {IERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/IERC20.sol";
import {IERC165} from "@openzeppelin/contracts@5.0.2/utils/introspection/IERC165.sol";

import {SafeERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/utils/SafeERC20.sol";

abstract contract TokenPool is IPoolV2, TokenPoolV1 {
  using SafeERC20 for IERC20;

  error DuplicateCCV(address ccv);
  error InvalidDestBytesOverhead(uint32 destBytesOverhead);
  error InvalidFinality(uint16 userSuppliedFinality, uint16 tokenConfigFinality);
  error AmountExceedsMaxPerRequest(uint256 userSuppliedAmount, uint256 max);

  event CCVConfigUpdated(uint64 indexed remoteChainSelector, address[] outboundCCVs, address[] inboundCCVs);
  event FinalityConfigUpdated(uint16 finalityConfig, uint16 fastTransferFeeBps, uint256 maxAmountPerRequest);
  event TokenTransferFeeConfigUpdated(uint64 indexed destChainSelector, TokenTransferFeeConfig tokenTransferFeeConfig);
  event TokenTransferFeeConfigDeleted(uint64 indexed destChainSelector);
  /// @notice Emitted when pool fees are withdrawn.
  event PoolFeeWithdrawn(address indexed recipient, uint256 amount);

  struct FastFinalityConfig {
    uint16 finalityThreshold; // ──╮ Maximum block depth required for token transfers.
    uint16 fastTransferFeeBps; // ─╯ Fee in basis points for fast transfers.
    uint256 maxAmountPerRequest; // Maximum amount allowed per transfer request.
  }

  struct CCVConfig {
    address[] outboundCCVs; // CCVs required for outgoing messages to the remote chain.
    address[] inboundCCVs; // CCVs required for incoming messages from the remote chain.
  }

  struct CCVConfigArg {
    uint64 remoteChainSelector;
    address[] outboundCCVs;
    address[] inboundCCVs;
  }

  struct TokenTransferFeeConfig {
    uint32 destGasOverhead; // ──╮ Gas charged to execute the token transfer on the destination chain.
    uint32 destBytesOverhead; // │ Data availability bytes.
    uint32 feeUSDCents; //       │ Fee to charge per token transfer, multiples of 0.01 USD.
    bool isEnabled; // ──────────╯ Whether this token has custom transfer fees.
  }

  /// @dev Struct with args for setting the token transfer fee configurations for a destination chain and a set of tokens.
  struct TokenTransferFeeConfigArgs {
    uint64 destChainSelector; // Destination chain selector.
    TokenTransferFeeConfig tokenTransferFeeConfig; // Token transfer fee configuration.
  }

  /// @notice The division factor for basis points (BPS). This also represents the maximum BPS fee for fast transfer.
  uint256 internal constant BPS_DIVIDER = 10_000;

  FastFinalityConfig internal s_finalityConfig;
  mapping(uint64 remoteChainSelector => CCVConfig ccvConfig) internal s_verifierConfig;
  mapping(uint64 destChainSelector => TokenTransferFeeConfig tokenTransferFeeConfig) internal s_tokenTransferFeeConfig;

  constructor(
    IERC20 token,
    uint8 localTokenDecimals,
    address[] memory allowlist,
    address rmnProxy,
    address router
  ) TokenPoolV1(token, localTokenDecimals, allowlist, rmnProxy, router) {}

  // ================================================================
  // │                        Lock or Burn                          │
  // ================================================================

  function lockOrBurn(
    Pool.LockOrBurnInV1 memory lockOrBurnIn,
    uint16 finality,
    bytes calldata // tokenArgs
  ) public virtual override returns (Pool.LockOrBurnOutV1 memory, uint256 destTokenAmount) {
    FastFinalityConfig memory finalityConfig = s_finalityConfig;
    if (finalityConfig.finalityThreshold != 0) {
      if (finality != 0 && finality < finalityConfig.finalityThreshold) {
        revert InvalidFinality(finality, finalityConfig.finalityThreshold);
      }
      if (lockOrBurnIn.amount > finalityConfig.maxAmountPerRequest) {
        revert AmountExceedsMaxPerRequest(lockOrBurnIn.amount, finalityConfig.maxAmountPerRequest);
      }
      // deduct fast transfer fee
      lockOrBurnIn.amount -= (lockOrBurnIn.amount * finalityConfig.fastTransferFeeBps) / BPS_DIVIDER;
    }

    _validateLockOrBurn(lockOrBurnIn);

    _lockOrBurn(lockOrBurnIn.amount);

    emit LockedOrBurned({
      remoteChainSelector: lockOrBurnIn.remoteChainSelector,
      token: address(i_token),
      sender: msg.sender,
      amount: lockOrBurnIn.amount
    });

    return (
      Pool.LockOrBurnOutV1({
        destTokenAddress: getRemoteToken(lockOrBurnIn.remoteChainSelector),
        destPoolData: _encodeLocalDecimals()
      }),
      lockOrBurnIn.amount
    );
  }

  // ================================================================
  // │                          Finality                             │
  // ================================================================
  /// @notice Updates the finality configuration for token transfers.
  function applyFinalityConfigUpdates(
    FastFinalityConfig calldata finalityConfig
  ) external virtual onlyOwner {
    s_finalityConfig = finalityConfig;
    emit FinalityConfigUpdated(
      finalityConfig.finalityThreshold, finalityConfig.fastTransferFeeBps, finalityConfig.maxAmountPerRequest
    );
  }

  // ================================================================
  // │                          CCV                                  │
  // ================================================================

  /// @notice Updates the CCV configuration for specified remote chains.
  /// If the array includes address(0), it indicates that the default CCV should be used alongside any other specified CCVs.
  function applyCCVConfigUpdates(
    CCVConfigArg[] calldata ccvConfigArgs
  ) external virtual onlyOwner {
    for (uint256 i = 0; i < ccvConfigArgs.length; ++i) {
      uint64 remoteChainSelector = ccvConfigArgs[i].remoteChainSelector;
      address[] calldata outboundCCVs = ccvConfigArgs[i].outboundCCVs;
      address[] calldata inboundCCVs = ccvConfigArgs[i].inboundCCVs;

      // check for duplicates in outbound CCVs.
      _checkNoDuplicateAddresses(outboundCCVs);

      // check for duplicates in inbound CCVs.
      _checkNoDuplicateAddresses(inboundCCVs);

      CCVConfig memory ccvConfig = CCVConfig({outboundCCVs: outboundCCVs, inboundCCVs: inboundCCVs});
      emit CCVConfigUpdated(remoteChainSelector, outboundCCVs, inboundCCVs);
      s_verifierConfig[remoteChainSelector] = ccvConfig;
    }
  }

  /// @notice Returns the set of required CCVs for incoming messages from a source chain.
  /// @param sourceChainSelector The source chain selector for incoming messages.
  /// This implementation assumes the same set of CCVs are used for all transfers on a lane.
  /// Implementers can override this function to define custom logic based on these params.
  /// @return requiredCCVs Set of required CCV addresses.
  function getRequiredInboundCCVs(
    address, // localToken
    uint64 sourceChainSelector,
    uint256, // amount,
    uint16, // finality
    bytes calldata // sourcePoolData
  ) external view virtual returns (address[] memory requiredCCVs) {
    return s_verifierConfig[sourceChainSelector].inboundCCVs;
  }

  /// @notice Returns the set of required CCVs for outgoing messages to a destination chain.
  /// @param destChainSelector The destination chain selector for outgoing messages.
  /// This implementation assumes the same set of CCVs are used for all transfers on a lane.
  /// Implementers can override this function to define custom logic based on these params.
  /// @return requiredCCVs Set of required CCV addresses.
  function getRequiredOutboundCCVs(
    address, // localToken
    uint64 destChainSelector,
    uint256, // amount
    uint16, // finality
    bytes calldata // tokenArgs
  ) external view virtual returns (address[] memory requiredCCVs) {
    return s_verifierConfig[destChainSelector].outboundCCVs;
  }

  /// @notice Checks a CCV address array for duplicate entries.
  /// @param ccvs The array of CCV addresses to check for duplicates.
  function _checkNoDuplicateAddresses(
    address[] calldata ccvs
  ) private pure {
    for (uint256 i = 0; i < ccvs.length; ++i) {
      for (uint256 j = i + 1; j < ccvs.length; ++j) {
        if (ccvs[i] == ccvs[j]) {
          revert DuplicateCCV(ccvs[i]);
        }
      }
    }
  }

  // ================================================================
  // │                          Fee                                  │
  // ================================================================
  function applyTokenTransferFeeConfigUpdates(
    TokenTransferFeeConfigArgs[] memory tokenTransferFeeConfigArgs,
    uint64[] calldata destToUseDefaultFeeConfigs
  ) external virtual onlyOwner {
    for (uint256 i = 0; i < tokenTransferFeeConfigArgs.length; ++i) {
      uint64 destChainSelector = tokenTransferFeeConfigArgs[i].destChainSelector;
      TokenTransferFeeConfig memory tokenTransferFeeConfig = tokenTransferFeeConfigArgs[i].tokenTransferFeeConfig;

      if (tokenTransferFeeConfig.destBytesOverhead < Pool.CCIP_LOCK_OR_BURN_V1_RET_BYTES) {
        revert InvalidDestBytesOverhead(tokenTransferFeeConfig.destBytesOverhead);
      }

      s_tokenTransferFeeConfig[destChainSelector] = tokenTransferFeeConfig;
      emit TokenTransferFeeConfigUpdated(destChainSelector, tokenTransferFeeConfig);
    }

    for (uint256 i = 0; i < destToUseDefaultFeeConfigs.length; ++i) {
      uint64 destChainSelector = destToUseDefaultFeeConfigs[i];
      delete s_tokenTransferFeeConfig[destChainSelector];
      emit TokenTransferFeeConfigDeleted(destChainSelector);
    }
  }

  function getTokenTransferFeeConfig(
    address, // localToken
    uint64 destChainSelector,
    Client.EVM2AnyMessage calldata, // message
    uint16, // finality
    bytes calldata // tokenArgs
  )
    external
    view
    virtual
    returns (bool isEnabled, uint32 destGasOverhead, uint32 destBytesOverhead, uint32 feeUSDCents)
  {
    TokenTransferFeeConfig memory feeConfig = s_tokenTransferFeeConfig[destChainSelector];
    return (feeConfig.isEnabled, feeConfig.destGasOverhead, feeConfig.destBytesOverhead, feeConfig.feeUSDCents);
  }

  // @inheritdoc IPoolV2
  function withdrawPoolFees(
    address recipient
  ) external virtual onlyOwner {
    uint256 amount = getAccumulatedPoolFees();
    if (amount > 0) {
      getToken().safeTransfer(recipient, amount);
      emit PoolFeeWithdrawn(recipient, amount);
    }
  }

  // @inheritdoc IPoolV2
  function getAccumulatedPoolFees() public view virtual returns (uint256) {
    return getToken().balanceOf(address(this));
  }

  /// @notice Signals which version of the pool interface is supported.
  function supportsInterface(
    bytes4 interfaceId
  ) public pure virtual override(TokenPoolV1, IERC165) returns (bool) {
    return interfaceId == Pool.CCIP_POOL_V2 || interfaceId == type(IPoolV2).interfaceId
      || super.supportsInterface(interfaceId);
  }
}
