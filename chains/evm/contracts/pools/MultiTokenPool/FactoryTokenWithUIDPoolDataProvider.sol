// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IAdditionalPoolDataProvider} from "../../interfaces/IAdditionalPoolDataProvider.sol";
import {IHasUID} from "../../interfaces/IHasUID.sol";
import {IResolveTokenAddressFromUID} from "../../interfaces/IResolveTokenAddressFromUID.sol";
import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";

/// @notice Sends the UID of a factory token as the additional pool data, and checks the UID on the destination chain.
/// @dev The token factory on each chain deploys at most one token for each UID. The issuer gives the same asset the same
/// UID on every chain. So a token and its UID identify the asset, and the token addresses on two chains can differ.
/// - On send, the provider reads the UID from the local token. The factory must return the local token for that UID. So
///   a token cannot send the UID of another token.
/// - On receive, the factory must return the local token for the UID from the source chain. So a transfer can only
///   release or mint the token with the same UID as the token that was locked or burned.
/// - Empty data is not valid. So a source pool that does not send the UID cannot release or mint a token that uses this
///   provider.
/// The provider has no config other than the factory. So one instance can serve every token of the factory.
contract FactoryTokenWithUIDPoolDataProvider is IAdditionalPoolDataProvider, ITypeAndVersion {
  error InvalidTokenFactory(IResolveTokenAddressFromUID tokenFactory);
  error TokenNotFromFactory(address localToken, bytes32 uid);

  /// @dev The length of abi.encode(bytes32 uid).
  uint256 private constant UID_POOL_DATA_LENGTH = 32;

  /// @dev The factory that maps each UID to its token on this chain.
  IResolveTokenAddressFromUID private immutable i_tokenFactory;

  /// @param tokenFactory The factory that deploys the tokens on this chain.
  constructor(
    IResolveTokenAddressFromUID tokenFactory
  ) {
    // Calling tokenAddress on an address without code reverts, which would stop every transfer.
    if (address(tokenFactory).code.length == 0) revert InvalidTokenFactory(tokenFactory);
    i_tokenFactory = tokenFactory;
  }

  function typeAndVersion() external pure virtual override returns (string memory) {
    return "FactoryTokenWithUIDPoolDataProvider 2.1.0-dev";
  }

  /// @inheritdoc IAdditionalPoolDataProvider
  /// @dev Returns abi.encode(bytes32 uid). Reverts if the factory does not return the local token for its UID.
  function getAdditionalPoolData(
    address localToken,
    uint64 // remoteChainSelector
  ) public view virtual returns (bytes memory additionalPoolData) {
    bytes32 uid = IHasUID(localToken).uid();
    if (i_tokenFactory.tokenAddress(uid) != localToken) revert TokenNotFromFactory(localToken, uid);
    return abi.encode(uid);
  }

  /// @inheritdoc IAdditionalPoolDataProvider
  /// @dev Valid only if the data is abi.encode(bytes32 uid) and the factory returns the local token for that UID.
  function isAdditionalPoolDataValid(
    address localToken,
    uint64, // remoteChainSelector
    bytes calldata additionalPoolData
  ) public view virtual returns (bool isValid) {
    if (additionalPoolData.length != UID_POOL_DATA_LENGTH) return false;
    address factoryToken = i_tokenFactory.tokenAddress(abi.decode(additionalPoolData, (bytes32)));
    // The factory returns zero for a UID without a token. So zero is never a valid local token.
    return factoryToken != address(0) && factoryToken == localToken;
  }

  /// @notice Returns the token factory.
  /// @return tokenFactory The factory that maps each UID to its token on this chain.
  function getTokenFactory() public view virtual returns (IResolveTokenAddressFromUID tokenFactory) {
    return i_tokenFactory;
  }
}
