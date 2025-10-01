pragma solidity ^0.8.25;

import {FactoryBurnMintERC20} from "./FactoryBurnMintERC20.sol";

contract HyperLiquidCompatibleERC20 is FactoryBurnMintERC20 {
  error LinkerAddressCannotBeZero();
  error HyperEVMTransferFailed();
  error InsufficientSpotBalance();
  error OverflowDetected(uint8 remoteDecimals, uint8 localDecimals, uint256 remoteAmount);
  error ZeroAddressNotAllowed();
  error RemoteTokenAlreadySet();

  event HyperEVMLinkerSet(address indexed hyperEVMLinker);
  event RemoteTokenSet(uint64 indexed remoteTokenId, address indexed remoteToken, uint8 indexed remoteTokenDecimals);

  // In order to bridge to HyperCore, factory-deployed contracts must store the address of a finalizer linker at
  // storage slot keccak256("HyperCore deployer")
  bytes32 internal constant HYPER_EVM_LINKER_SLOT = 0x8c306a6a12fff1951878e8621be6674add1102cd359dd968efbbe797629ef84f;

  address internal constant SPOT_BALANCE_PRECOMPILE_ADDRESS = 0x0000000000000000000000000000000000000801;

  /// @dev The address of the HyperCore token and its associated number of decimals may not be known
  /// at deployment time, and must be set later after the HyperCore token is deployed.
  uint64 internal s_hypercoreTokenSpotId;
  address internal s_hypercoreTokenSystemAddress;
  uint8 internal s_hypercoreTokenDecimals;

  struct SpotBalance {
    uint64 total;
    uint64 hold;
    uint64 entryNtl;
  }

  constructor(
    string memory name,
    string memory symbol,
    uint8 decimals,
    uint256 maxSupply,
    uint256 preMint,
    address newOwner
  ) FactoryBurnMintERC20(name, symbol, decimals, maxSupply, preMint, newOwner) {}

  /// @notice Using a function because constant state variables cannot be overridden by child contracts.
  function typeAndVersion() external pure virtual override returns (string memory) {
    return "HyperLiquidCompatibleERC20 1.6.2";
  }

  /// @notice Sets the hyperEVMLinker address.
  /// @param newLinker The address of the hyperEVMLinker.
  function setHyperEVMLinker(
    address newLinker
  ) external onlyOwner {
    if (newLinker == address(0)) {
      revert LinkerAddressCannotBeZero();
    }

    assembly {
      sstore(HYPER_EVM_LINKER_SLOT, newLinker)
    }

    emit HyperEVMLinkerSet(newLinker);
  }

  /// @notice Sets the remote token and decimals.
  /// @param hipTokenId The id of the remote token.
  /// @param remoteTokenDecimals The decimals of the remote token.
  /// @dev While the zero address is not allowed, it is allowed for the remote token to have zero decimals.
  function setRemoteToken(uint64 hipTokenId, uint8 remoteTokenDecimals) external onlyOwner {
    if (s_hypercoreTokenSystemAddress != address(0)) {
      revert RemoteTokenAlreadySet();
    }

    if (hipTokenId == 0) {
      revert ZeroAddressNotAllowed();
    }

    s_hypercoreTokenSpotId = hipTokenId;
    s_hypercoreTokenDecimals = remoteTokenDecimals;

    s_hypercoreTokenSystemAddress = address((uint160(0x20) << 152) | uint160(hipTokenId));

    emit RemoteTokenSet(hipTokenId, s_hypercoreTokenSystemAddress, remoteTokenDecimals);
  }

  /// @notice Gets the hyperEVMLinker address.
  /// @return hyperEVMLinker The address of the hyperEVMLinker.
  function getHyperEVMLinker() public view returns (address) {
    address hyperEVMLinker;
    assembly {
      hyperEVMLinker := sload(HYPER_EVM_LINKER_SLOT)
    }
    return hyperEVMLinker;
  }

  /// @notice Overrides the standard ERC20 transfer hook to add a balance check before bridging tokens to HyperCore.
  /// @dev This internal hook intercepts transfers to the `hypercoreTokenSystemAddress`. Before allowing the transfer,
  /// it performs a `staticcall` to the `SPOT_BALANCE_PRECOMPILE_ADDRESS` to fetch the system address's current
  /// spot balance on HyperCore. It then compares this remote balance (normalized to local decimals) with the transfer
  /// amount. This check prevents users from losing funds by ensuring the bridge destination on HyperCore has
  /// sufficient liquidity before the HyperEVM-side transfer occurs. The function reverts if the transfer amount
  /// exceeds the available spot balance or if the precompile call fails.
  /// @param to The recipient address of the token transfer.
  /// @param amount The amount of tokens being transferred.
  function _beforeTokenTransfer(address, address to, uint256 amount) internal virtual override {
    if (to == s_hypercoreTokenSystemAddress) {
      (bool success, bytes memory result) =
        SPOT_BALANCE_PRECOMPILE_ADDRESS.staticcall(abi.encode(to, s_hypercoreTokenSpotId));

      if (!success) {
        revert HyperEVMTransferFailed();
      }

      SpotBalance memory spotBalance = abi.decode(result, (SpotBalance));

      uint256 remoteAmountNormalizedLocalDecimals = _calculateLocalAmount(spotBalance.total, s_hypercoreTokenDecimals);

      if (amount > remoteAmountNormalizedLocalDecimals) {
        revert InsufficientSpotBalance();
      }
    }
  }

  /// @notice Calculates the local amount based on the remote amount and decimals.
  /// @param remoteAmount The amount on the remote chain.
  /// @param remoteDecimals The decimals of the token on the remote chain.
  /// @return The local amount.
  /// @dev This function protects against overflows. If there is a transaction that hits the overflow check, it is
  /// probably incorrect as that means the amount cannot be represented on this chain. If the local decimals have been
  /// wrongly configured, the token issuer could redeploy the pool with the correct decimals and manually re-execute the
  /// CCIP tx to fix the issue.
  function _calculateLocalAmount(uint256 remoteAmount, uint8 remoteDecimals) internal view virtual returns (uint256) {
    uint8 localDecimals = decimals();

    if (remoteDecimals == localDecimals) {
      return remoteAmount;
    }
    if (remoteDecimals > localDecimals) {
      uint8 decimalsDiff = remoteDecimals - localDecimals;
      if (decimalsDiff > 77) {
        // This is a safety check to prevent overflow in the next calculation.
        revert OverflowDetected(remoteDecimals, localDecimals, remoteAmount);
      }
      // Solidity rounds down so there is no risk of minting more tokens than the remote chain sent.
      return remoteAmount / (10 ** decimalsDiff);
    }

    // This is a safety check to prevent overflow in the next calculation.
    // More than 77 would never fit in a uint256 and would cause an overflow. We also check if the resulting amount
    // would overflow.
    uint8 diffDecimals = localDecimals - remoteDecimals;
    if (diffDecimals > 77 || remoteAmount > type(uint256).max / (10 ** diffDecimals)) {
      revert OverflowDetected(remoteDecimals, localDecimals, remoteAmount);
    }

    return remoteAmount * (10 ** diffDecimals);
  }
}
