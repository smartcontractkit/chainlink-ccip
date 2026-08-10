// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {AdvancedPoolHooks} from "../../../pools/AdvancedPoolHooks.sol";
import {TokenPool} from "../../../pools/TokenPool.sol";
import {BaseERC20} from "../../../tokens/BaseERC20.sol";
import {CrossChainToken} from "../../../tokens/CrossChainToken.sol";
import {BaseTest} from "../../BaseTest.t.sol";
import {TokenPoolHelper} from "../../helpers/TokenPoolHelper.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

import {IERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/IERC20.sol";

contract AdvancedPoolHooksSetup is BaseTest {
  IERC20 internal s_token;
  TokenPoolHelper internal s_tokenPool;
  AdvancedPoolHooks internal s_advancedPoolHooks;
  uint16 internal constant BPS_DIVIDER = 10_000;
  uint256 internal constant CCV_THRESHOLD_AMOUNT = 1000e18;
  address public s_sender = makeAddr("sender");
  bytes public s_receiver = abi.encode(makeAddr("receiver"));

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

    address[] memory authorizedCallers = new address[](1);
    authorizedCallers[0] = OWNER;
    s_advancedPoolHooks = new AdvancedPoolHooks(new address[](0), CCV_THRESHOLD_AMOUNT, address(0), authorizedCallers);

    s_tokenPool = new TokenPoolHelper(
      s_token, DEFAULT_TOKEN_DECIMALS, address(s_advancedPoolHooks), address(s_mockRMNRemote), address(s_sourceRouter)
    );

    address[] memory poolCallers = new address[](1);
    poolCallers[0] = address(s_tokenPool);
    s_advancedPoolHooks.applyAuthorizedCallerUpdates(
      AuthorizedCallers.AuthorizedCallerArgs({addedCallers: poolCallers, removedCallers: new address[](0)})
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
}
