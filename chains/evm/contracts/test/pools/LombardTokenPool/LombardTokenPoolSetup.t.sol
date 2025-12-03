// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Router} from "../../../Router.sol";
import {TokenPool} from "../../../pools/TokenPool.sol";
import {LombardTokenPoolHelper} from "../../helpers/LombardTokenPoolHelper.sol";
import {MockVerifier} from "../../mocks/MockVerifier.sol";
import {TokenPoolSetup} from "../TokenPool/TokenPoolSetup.t.sol";

contract LombardTokenPoolSetup is TokenPoolSetup {
  LombardTokenPoolHelper internal s_pool;
  MockVerifier internal s_verifierResolver;
  address internal constant VERIFIER_IMPL = address(0x2345);
  address internal s_remotePool = makeAddr("remotePool");
  address internal s_remoteToken = makeAddr("remoteToken");

  function setUp() public virtual override {
    super.setUp();

    s_verifierResolver = new MockVerifier("");

    s_pool = new LombardTokenPoolHelper(
      s_token, address(s_verifierResolver), address(s_mockRMNRemote), address(s_sourceRouter), DEFAULT_TOKEN_DECIMALS
    );

    // Configure remote chain.
    bytes[] memory remotePools = new bytes[](1);
    remotePools[0] = abi.encode(s_remotePool);

    TokenPool.ChainUpdate[] memory chainUpdate = new TokenPool.ChainUpdate[](1);
    chainUpdate[0] = TokenPool.ChainUpdate({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      remotePoolAddresses: remotePools,
      remoteTokenAddress: abi.encode(s_remoteToken),
      outboundRateLimiterConfig: _getOutboundRateLimiterConfig(),
      inboundRateLimiterConfig: _getInboundRateLimiterConfig()
    });

    vm.startPrank(OWNER);
    s_pool.applyChainUpdates(new uint64[](0), chainUpdate);

    Router.OnRamp[] memory onRampUpdates = new Router.OnRamp[](1);
    onRampUpdates[0] = Router.OnRamp({destChainSelector: DEST_CHAIN_SELECTOR, onRamp: s_allowedOnRamp});
    Router.OffRamp[] memory offRampUpdates = new Router.OffRamp[](1);
    offRampUpdates[0] = Router.OffRamp({sourceChainSelector: DEST_CHAIN_SELECTOR, offRamp: s_allowedOffRamp});
    s_sourceRouter.applyRampUpdates(onRampUpdates, new Router.OffRamp[](0), offRampUpdates);
  }
}
