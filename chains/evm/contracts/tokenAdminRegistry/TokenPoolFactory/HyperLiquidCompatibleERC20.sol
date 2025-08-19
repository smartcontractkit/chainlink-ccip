pragma solidity ^0.8.25;

import {FactoryBurnMintERC20} from "./FactoryBurnMintERC20.sol";

contract HyperLiquidCompatibleERC20 is FactoryBurnMintERC20 {
  error LinkerAddressCannotBeZero();
  error HyperEVMTransferFailed();
  error InsufficientSpotBalance();
  error OverflowDetected(uint8 remoteDecimals, uint8 localDecimals, uint256 remoteAmount);
  error ZeroAddressNotAllowed();

  event HyperEVMLinkerSet(address indexed hyperEVMLinker);
  event RemoteTokenSet(address indexed remoteToken, uint8 indexed remoteTokenDecimals);

  // In order to bridge to HyperCore, factory-deployed contracts must store the address of a finalizer linker at
  // storage slot keccak256("HyperCore deployer")
  bytes32 internal constant HYPER_EVM_LINKER_SLOT = 0x8c306a6a12fff1951878e8621be6674add1102cd359dd968efbbe797629ef84f;

  address public constant SPOT_BALANCE_PRECOMPILE = 0x0000000000000000000000000000000000000800;

  /// @dev The address of the remote token and its associated number of decimals may not be known
  /// at deployment time, and must be set later after the remote token is deployed.
  address internal s_remoteToken;
  uint8 internal s_remoteTokenDecimals;

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
  /// @param remoteToken The address of the remote token.
  /// @param remoteTokenDecimals The decimals of the remote token.
  /// @dev While the zero address is not allowed, it is allowed for the remote token to have zero decimals.
  function setRemoteToken(address remoteToken, uint8 remoteTokenDecimals) external onlyOwner {
    if (remoteToken == address(0)) {
      revert ZeroAddressNotAllowed();
    }

    s_remoteToken = remoteToken;
    s_remoteTokenDecimals = remoteTokenDecimals;

    emit RemoteTokenSet(remoteToken, remoteTokenDecimals);
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

  function _beforeTokenTransfer(address, address to, uint256 amount) internal virtual override {
    if (to == SPOT_BALANCE_PRECOMPILE) {
      (bool success, bytes memory result) =
        SPOT_BALANCE_PRECOMPILE.staticcall(abi.encode(getHyperEVMLinker(), s_remoteToken));

      if (!success) {
        revert HyperEVMTransferFailed();
      }

      SpotBalance memory spotBalance = abi.decode(result, (SpotBalance));

      uint256 remoteAmountNormalizedLocalDecimals = _calculateLocalAmount(spotBalance.total, s_remoteTokenDecimals);

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
