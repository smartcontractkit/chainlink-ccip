// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {TokenAdminRegistry} from "../tokenAdminRegistry/TokenAdminRegistry.sol";
import {BaseERC20} from "../tokens/BaseERC20.sol";
import {CrossChainToken} from "../tokens/CrossChainToken.sol";
import {BaseTest} from "./BaseTest.t.sol";

/// @notice Deploys the shared test tokens and the token admin registry without deploying any token pools. Suites that
/// only need tokens (e.g. fee accounting) inherit this fixture, suites that exercise token transfers through pools
/// inherit TokenSetup which layers the pools on top.
contract TokenFixture is BaseTest {
  address[] internal s_sourceTokens;
  address[] internal s_destTokens;

  address internal s_sourceFeeToken;
  address internal s_destFeeToken;

  TokenAdminRegistry internal s_tokenAdminRegistry;

  mapping(address sourceToken => address sourcePool) internal s_sourcePoolByToken;
  mapping(address sourceToken => address destPool) internal s_destPoolBySourceToken;
  mapping(address destToken => address destPool) internal s_destPoolByToken;
  mapping(address sourceToken => address destToken) internal s_destTokenBySourceToken;

  function setUp() public virtual override {
    super.setUp();

    // setUp is often called multiple times from tests' setUp due to inheritance.
    bool isSetup = s_sourceTokens.length != 0;
    if (isSetup) {
      return;
    }

    s_tokenAdminRegistry = new TokenAdminRegistry();

    address sourceLink = _deploySourceToken("sLINK", type(uint256).max, 18);
    s_sourceFeeToken = sourceLink;

    address sourceEth = _deploySourceToken("sETH", 2 ** 128, 18);

    // A 6 decimal token, like USDC, ensures decimal conversions are exercised.
    address sourceUsdc = _deploySourceToken("sUSDC", 2 ** 128, 6);

    address destLink = _deployDestToken("dLINK", type(uint256).max, 18);
    s_destFeeToken = destLink;
    s_destTokenBySourceToken[sourceLink] = destLink;

    address destEth = _deployDestToken("dETH", 2 ** 128, 18);
    s_destTokenBySourceToken[sourceEth] = destEth;

    address destUsdc = _deployDestToken("dUSDC", 2 ** 128, 6);
    s_destTokenBySourceToken[sourceUsdc] = destUsdc;
  }

  function _deployCrossChainToken(
    string memory tokenName,
    uint8 decimals
  ) internal returns (CrossChainToken) {
    return new CrossChainToken(
      BaseERC20.ConstructorParams({
        name: tokenName,
        symbol: tokenName,
        decimals: decimals,
        maxSupply: 0,
        preMint: 0,
        preMintRecipient: address(0),
        ccipAdmin: OWNER
      }),
      OWNER,
      OWNER
    );
  }

  function _deploySourceToken(
    string memory tokenName,
    uint256 dealAmount,
    uint8 decimals
  ) internal returns (address) {
    CrossChainToken token = _deployCrossChainToken(tokenName, decimals);
    s_sourceTokens.push(address(token));
    deal(address(token), OWNER, dealAmount);
    return address(token);
  }

  function _deployDestToken(
    string memory tokenName,
    uint256 dealAmount,
    uint8 decimals
  ) internal returns (address) {
    CrossChainToken token = _deployCrossChainToken(tokenName, decimals);
    s_destTokens.push(address(token));
    deal(address(token), OWNER, dealAmount);
    return address(token);
  }
}
