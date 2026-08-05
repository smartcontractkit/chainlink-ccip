// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IBurnMintERC20} from "../interfaces/IBurnMintERC20.sol";
import {BurnMintTokenPool} from "../pools/BurnMintTokenPool.sol";
import {ERC20LockBox} from "../pools/ERC20LockBox.sol";
import {LockReleaseTokenPool} from "../pools/LockReleaseTokenPool.sol";
import {TokenPool} from "../pools/TokenPool.sol";
import {TokenAdminRegistry} from "../tokenAdminRegistry/TokenAdminRegistry.sol";
import {CrossChainToken} from "../tokens/CrossChainToken.sol";
import {TokenFixture} from "./TokenFixture.t.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

import {IERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/IERC20.sol";

/// @notice Layers real token pools and registry wiring on top of the TokenFixture tokens. Only suites that transfer
/// tokens through pools should inherit this, others should inherit TokenFixture to keep pools out of their compile
/// scope.
contract TokenSetup is TokenFixture {
  mapping(address token => ERC20LockBox lockBox) internal s_lockBoxes;

  function _deployLockReleasePool(
    address token,
    bool isSourcePool
  ) internal {
    address router = address(s_sourceRouter);
    if (!isSourcePool) {
      router = address(s_destRouter);
    }

    ERC20LockBox lockBox = new ERC20LockBox(token);
    s_lockBoxes[token] = lockBox;

    LockReleaseTokenPool pool = new LockReleaseTokenPool(
      IERC20(token), DEFAULT_TOKEN_DECIMALS, address(0), address(s_mockRMNRemote), router, address(lockBox)
    );

    address[] memory authorizedCallers = new address[](1);
    authorizedCallers[0] = address(pool);
    AuthorizedCallers.AuthorizedCallerArgs memory args =
      AuthorizedCallers.AuthorizedCallerArgs({addedCallers: authorizedCallers, removedCallers: new address[](0)});
    lockBox.applyAuthorizedCallerUpdates(args);

    if (isSourcePool) {
      s_sourcePoolByToken[address(token)] = address(pool);
    } else {
      s_destPoolByToken[address(token)] = address(pool);
    }
  }

  function _deployBurnMintPool(
    address token,
    uint8 localTokenDecimals,
    bool isSourcePool
  ) internal {
    address router = address(s_sourceRouter);
    if (!isSourcePool) {
      router = address(s_destRouter);
    }

    BurnMintTokenPool pool = new BurnMintTokenPool(
      IBurnMintERC20(address(token)), localTokenDecimals, address(0), address(s_mockRMNRemote), router
    );
    CrossChainToken(token).grantMintAndBurnRoles(address(pool));

    if (isSourcePool) {
      s_sourcePoolByToken[address(token)] = address(pool);
    } else {
      s_destPoolByToken[address(token)] = address(pool);
    }
  }

  function setUp() public virtual override {
    super.setUp();

    // setUp is often called multiple times from tests' setUp due to inheritance.
    bool isSetup = s_sourcePoolByToken[s_sourceTokens[0]] != address(0);
    if (isSetup) {
      return;
    }

    // Source pools: lock/release for the fee token, burn/mint for the others.
    _deployLockReleasePool(s_sourceTokens[0], true);
    _deployBurnMintPool(s_sourceTokens[1], DEFAULT_TOKEN_DECIMALS, true);
    _deployBurnMintPool(s_sourceTokens[2], USDC_TOKEN_DECIMALS, true);

    // Destination pools mirror the source pool types.
    _deployLockReleasePool(s_destTokens[0], false);
    _deployBurnMintPool(s_destTokens[1], DEFAULT_TOKEN_DECIMALS, false);
    _deployBurnMintPool(s_destTokens[2], USDC_TOKEN_DECIMALS, false);

    s_destPoolBySourceToken[s_sourceTokens[0]] = s_destPoolByToken[s_destTokens[0]];
    s_destPoolBySourceToken[s_sourceTokens[1]] = s_destPoolByToken[s_destTokens[1]];
    s_destPoolBySourceToken[s_sourceTokens[2]] = s_destPoolByToken[s_destTokens[2]];

    // Float the dest link lock release pool with funds.
    IERC20(s_destTokens[0]).transfer(address(s_lockBoxes[s_destTokens[0]]), 1000 ether);

    // Set pools in the registry.
    for (uint256 i = 0; i < s_sourceTokens.length; ++i) {
      address token = s_sourceTokens[i];
      address pool = s_sourcePoolByToken[token];

      _setPool(
        s_tokenAdminRegistry, token, pool, DEST_CHAIN_SELECTOR, s_destPoolByToken[s_destTokens[i]], s_destTokens[i]
      );
    }

    for (uint256 i = 0; i < s_destTokens.length; ++i) {
      address token = s_destTokens[i];
      address pool = s_destPoolByToken[token];
      s_tokenAdminRegistry.proposeAdministrator(token, OWNER);
      s_tokenAdminRegistry.acceptAdminRole(token);
      s_tokenAdminRegistry.setPool(token, pool);

      _setPool(
        s_tokenAdminRegistry,
        token,
        pool,
        SOURCE_CHAIN_SELECTOR,
        s_sourcePoolByToken[s_sourceTokens[i]],
        s_sourceTokens[i]
      );
    }
  }

  function _setPool(
    TokenAdminRegistry tokenAdminRegistry,
    address token,
    address pool,
    uint64 remoteChainSelector,
    address remotePoolAddress,
    address remoteToken
  ) internal {
    if (!tokenAdminRegistry.isAdministrator(token, OWNER)) {
      tokenAdminRegistry.proposeAdministrator(token, OWNER);
      tokenAdminRegistry.acceptAdminRole(token);
    }

    tokenAdminRegistry.setPool(token, pool);

    bytes[] memory remotePoolAddresses = new bytes[](1);
    remotePoolAddresses[0] = abi.encode(remotePoolAddress);

    TokenPool.ChainUpdate[] memory chainUpdates = new TokenPool.ChainUpdate[](1);
    chainUpdates[0] = TokenPool.ChainUpdate({
      remoteChainSelector: remoteChainSelector,
      remotePoolAddresses: remotePoolAddresses,
      remoteTokenAddress: abi.encode(remoteToken),
      outboundRateLimiterConfig: _getOutboundRateLimiterConfig(),
      inboundRateLimiterConfig: _getInboundRateLimiterConfig()
    });

    TokenPool(pool).applyChainUpdates(new uint64[](0), chainUpdates);
  }
}
