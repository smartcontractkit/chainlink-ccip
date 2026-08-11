// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {TokenPool} from "../../../pools/TokenPool.sol";
import {BaseERC20} from "../../../tokens/BaseERC20.sol";
import {CrossChainToken} from "../../../tokens/CrossChainToken.sol";
import {BaseTest} from "../../BaseTest.t.sol";
import {TokenPoolHelper} from "../../helpers/TokenPoolHelper.sol";

import {IERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/IERC20.sol";

contract TokenPoolSetup is BaseTest {
  IERC20 internal s_token;
  TokenPoolHelper internal s_tokenPool;

  address internal s_allowedOffRamp = makeAddr("allowed_offRamp");
  address internal s_allowedOnRamp = makeAddr("allowed_onRamp");

  address internal s_initialRemotePool = makeAddr("initialRemotePool");
  address internal s_initialRemoteToken = makeAddr("initialRemoteToken");

  function setUp() public virtual override {
    super.setUp();
    s_token = IERC20(
      address(
        new CrossChainToken(
          BaseERC20.ConstructorParams({
            name: "LINK",
            symbol: "LNK",
            decimals: 18,
            maxSupply: 0,
            preMint: 0,
            preMintRecipient: address(0),
            ccipAdmin: OWNER
          }),
          OWNER,
          OWNER
        )
      )
    );
    deal(address(s_token), OWNER, type(uint256).max);

    s_tokenPool = new TokenPoolHelper(
      s_token, DEFAULT_TOKEN_DECIMALS, address(0), address(s_mockRMNRemote), address(s_sourceRouter)
    );

    bytes[] memory remotePoolAddresses = new bytes[](1);
    remotePoolAddresses[0] = abi.encode(s_initialRemotePool);

    TokenPool.ChainUpdate[] memory chainUpdate = new TokenPool.ChainUpdate[](1);
    chainUpdate[0] = TokenPool.ChainUpdate({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      remotePoolAddresses: remotePoolAddresses,
      remoteTokenAddress: abi.encode(s_initialRemoteToken),
      outboundRateLimiterConfig: _getOutboundRateLimiterConfig(),
      inboundRateLimiterConfig: _getInboundRateLimiterConfig()
    });

    s_tokenPool.applyChainUpdates(new uint64[](0), chainUpdate);

    _setMockRouterOnRamp(address(s_sourceRouter), DEST_CHAIN_SELECTOR, s_allowedOnRamp);
    _setMockRouterOffRamp(address(s_sourceRouter), DEST_CHAIN_SELECTOR, s_allowedOffRamp, true);
  }

  function _applyChainUpdates(
    address pool
  ) internal {
    bytes[] memory remotePoolAddresses = new bytes[](1);
    remotePoolAddresses[0] = abi.encode(s_initialRemotePool);

    TokenPool.ChainUpdate[] memory chainsToAdd = new TokenPool.ChainUpdate[](1);
    chainsToAdd[0] = TokenPool.ChainUpdate({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      remotePoolAddresses: remotePoolAddresses,
      remoteTokenAddress: abi.encode(s_initialRemoteToken),
      outboundRateLimiterConfig: _getOutboundRateLimiterConfig(),
      inboundRateLimiterConfig: _getInboundRateLimiterConfig()
    });

    TokenPool(pool).applyChainUpdates(new uint64[](0), chainsToAdd);

    _setMockRouterOnRamp(address(s_sourceRouter), DEST_CHAIN_SELECTOR, s_allowedOnRamp);
    _setMockRouterOffRamp(address(s_sourceRouter), DEST_CHAIN_SELECTOR, s_allowedOffRamp, true);
  }
}
