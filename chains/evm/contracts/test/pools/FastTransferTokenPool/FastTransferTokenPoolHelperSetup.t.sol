// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.10;

import {Router} from "../../../Router.sol";
import {FastTransferTokenPoolAbstract} from "../../../pools/FastTransferTokenPoolAbstract.sol";
import {FastTransferTokenPoolHelper} from "../../helpers/FastTransferTokenPoolHelper.sol";

import {BaseTest} from "../../BaseTest.t.sol";

import {BurnMintERC20} from "@chainlink/contracts/src/v0.8/shared/token/ERC20/BurnMintERC20.sol";
import {WETH9} from "@chainlink/contracts/src/v0.8/vendor/canonical-weth/WETH9.sol";

import {IERC20} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";

contract FastTransferTokenPoolHelperSetup is BaseTest {
  IERC20 internal s_token;
  FastTransferTokenPoolHelper public s_tokenPool;
  WETH9 public wrappedNative;

  address public s_filler;
  uint16 public constant FAST_FEE_BPS = 100; // 1%

  function setUp() public virtual override {
    super.setUp();
    address onRamp = makeAddr("onRamp");
    Router.OnRamp[] memory onRampUpdates = new Router.OnRamp[](1);
    onRampUpdates[0] = Router.OnRamp({destChainSelector: DEST_CHAIN_SELECTOR, onRamp: onRamp});
    s_sourceRouter.applyRampUpdates(onRampUpdates, new Router.OffRamp[](0), new Router.OffRamp[](0));
    s_filler = makeAddr("filler");
    s_token = new BurnMintERC20("LINK", "LNK", 18, 0, 0);
    deal(address(s_token), OWNER, type(uint256).max);
    wrappedNative = new WETH9();
    address[] memory addFillers = new address[](1);
    addFillers[0] = s_filler;
    // Deploy pool
    FastTransferTokenPoolAbstract.LaneConfigArgs[] memory laneConfigArgs =
      new FastTransferTokenPoolAbstract.LaneConfigArgs[](1);
    laneConfigArgs[0] = FastTransferTokenPoolAbstract.LaneConfigArgs({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      bpsFastFee: FAST_FEE_BPS, // 1%
      enabled: true,
      fillerAllowlistEnabled: true,
      destinationPool: address(0x4),
      fillAmountMaxPerRequest: 1000 ether,
      addFillers: addFillers,
      removeFillers: new address[](0)
    });

    s_tokenPool = new FastTransferTokenPoolHelper(s_token, wrappedNative, address(s_sourceRouter));
    s_tokenPool.updateLaneConfig(laneConfigArgs[0]);

    // Approve tokens
    s_token.approve(address(s_tokenPool), type(uint256).max);
  }
}
