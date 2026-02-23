// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

import {IGetCCIPAdmin} from "../interfaces/IGetCCIPAdmin.sol";
import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";

import {ERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/ERC20.sol";
import {IERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/IERC20.sol";
import {IERC165} from "@openzeppelin/contracts@5.3.0/utils/introspection/IERC165.sol";

/// @notice A basic ERC20 compatible token contract with burn and minting roles.
/// @dev The total supply can be limited during deployment.
contract BaseERC20 is IGetCCIPAdmin, ERC20, ITypeAndVersion, IERC165 {
  function typeAndVersion() external pure virtual override returns (string memory) {
    return "BaseERC20 2.0.0-dev";
  }

  error InvalidRecipient(address recipient);
  error OnlyCCIPAdmin();

  event CCIPAdminTransferred(address indexed previousAdmin, address indexed newAdmin);

  /// @param name The name of the token
  /// @param symbol The symbol of the token
  /// @param decimals_ The number of decimals the token uses
  /// @param maxSupply_ The maximum supply of the token, 0 if unlimited
  /// @param preMint The amount of tokens to mint to the deployer upon construction. NOTE: the base version of this
  /// contract does not support minting additional tokens after deployment, so this should be set to the full supply.
  struct ConstructorParams {
    string name;
    string symbol;
    uint256 maxSupply;
    uint256 preMint;
    uint8 decimals;
    address ccipAdmin;
  }

  /// @dev The number of decimals for the token
  uint8 internal immutable i_decimals;

  /// @dev The maximum supply of the token, 0 if unlimited
  uint256 internal immutable i_maxSupply;

  /// @dev the CCIPAdmin can be used to register with the CCIP token admin registry, but has no other special powers,
  /// and can only be transferred by the owner.
  address internal s_ccipAdmin;

  constructor(
    ConstructorParams memory args
  ) ERC20(args.name, args.symbol) {
    i_decimals = args.decimals;
    i_maxSupply = args.maxSupply;

    address ccipAdmin = args.ccipAdmin == address(0) ? msg.sender : args.ccipAdmin;

    // Mint the initial supply to the new Owner, saving gas by not calling if the mint amount is zero.
    if (args.preMint != 0) _mint(ccipAdmin, args.preMint);

    _setCCIPAdmin(ccipAdmin);
  }

  /// @inheritdoc IERC165
  function supportsInterface(
    bytes4 interfaceId
  ) public view virtual returns (bool) {
    return interfaceId == type(IERC20).interfaceId || interfaceId == type(IGetCCIPAdmin).interfaceId;
  }

  // ================================================================
  // │                            ERC20                             │
  // ================================================================

  /// @dev Returns the number of decimals used in its user representation.
  function decimals() public view virtual override returns (uint8) {
    return i_decimals;
  }

  /// @dev Returns the max supply of the token, 0 if unlimited.
  function maxSupply() public view virtual returns (uint256) {
    return i_maxSupply;
  }

  /// @dev Uses OZ ERC20 _approve to disallow approving for address(0).
  /// @dev Disallows approving for address(this).
  function _approve(
    address owner,
    address spender,
    uint256 value,
    bool emitEvent
  ) internal virtual override {
    if (spender == address(this)) revert InvalidRecipient(spender);

    super._approve(owner, spender, value, emitEvent);
  }

  /// @dev This check applies to transfer, minting, and burning.
  /// @dev Disallows approving for address(this).
  function _update(
    address from,
    address to,
    uint256 value
  ) internal virtual override {
    if (to == address(this)) revert InvalidRecipient(to);

    super._update(from, to, value);
  }

  // ================================================================
  // │                            Roles                             │
  // ================================================================

  /// @notice Returns the current CCIPAdmin.
  function getCCIPAdmin() external view virtual returns (address) {
    return s_ccipAdmin;
  }

  /// @notice Transfers the CCIPAdmin role to a new address
  function setCCIPAdmin(
    address newAdmin
  ) external virtual {
    if (msg.sender != s_ccipAdmin) {
      revert OnlyCCIPAdmin();
    }
    _setCCIPAdmin(newAdmin);
  }

  function _setCCIPAdmin(
    address newAdmin
  ) internal virtual {
    address currentAdmin = s_ccipAdmin;

    s_ccipAdmin = newAdmin;

    emit CCIPAdminTransferred(currentAdmin, newAdmin);
  }
}
