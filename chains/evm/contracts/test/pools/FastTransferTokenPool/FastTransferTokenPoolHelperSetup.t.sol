// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Router} from "../../../Router.sol";
import {FastTransferTokenPoolAbstract} from "../../../pools/FastTransferTokenPoolAbstract.sol";

import {Client} from "../../../libraries/Client.sol";
import {TokenPool} from "../../../pools/TokenPool.sol";
import {BaseTest} from "../../BaseTest.t.sol";
import {FastTransferTokenPoolHelper} from "../../helpers/FastTransferTokenPoolHelper.sol";
import {BurnMintERC20} from "@chainlink/contracts/src/v0.8/shared/token/ERC20/BurnMintERC20.sol";
import {WETH9} from "@chainlink/contracts/src/v0.8/vendor/canonical-weth/WETH9.sol";

import {Internal} from "../../../libraries/Internal.sol";
import {IERC20} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";

contract FastTransferTokenPoolHelperSetup is BaseTest {
  uint16 public constant FAST_FEE_BPS = 100; // 1%
  uint32 internal constant SVM_CHAIN_SELECTOR = uint32(uint256(keccak256("SVM_SELECTOR")));
  uint32 internal constant SETTLEMENT_GAS_OVERHEAD = 200_000;
  uint256 internal constant MAX_FILL_AMOUNT_PER_REQUEST = 1000 ether;
  bytes svmExtraArgsBytesEncoded;
  IERC20 internal s_token;
  FastTransferTokenPoolHelper public s_tokenPool;
  WETH9 public wrappedNative;
  bytes public destPoolAddress;
  address public s_filler;

  function setUp() public virtual override {
    super.setUp();
    destPoolAddress = abi.encode(makeAddr("destPool"));
    address onRamp = makeAddr("onRamp");
    svmExtraArgsBytesEncoded = Client._svmArgsToBytes(
      Client.SVMExtraArgsV1({
        computeUnits: SETTLEMENT_GAS_OVERHEAD,
        accounts: new bytes32[](0),
        accountIsWritableBitmap: 2,
        allowOutOfOrderExecution: true,
        tokenReceiver: bytes32(0)
      })
    );
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
    FastTransferTokenPoolAbstract.DestChainConfigUpdateArgs[] memory laneConfigArgs =
      new FastTransferTokenPoolAbstract.DestChainConfigUpdateArgs[](2);
    laneConfigArgs[0] = FastTransferTokenPoolAbstract.DestChainConfigUpdateArgs({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      fastTransferBpsFee: FAST_FEE_BPS, // 1%
      fillerAllowlistEnabled: true,
      destinationPool: destPoolAddress,
      maxFillAmountPerRequest: MAX_FILL_AMOUNT_PER_REQUEST,
      settlementOverheadGas: SETTLEMENT_GAS_OVERHEAD,
      chainFamilySelector: Internal.CHAIN_FAMILY_SELECTOR_EVM,
      evmToAnyMessageExtraArgsBytes: "",
      addFillers: addFillers,
      removeFillers: new address[](0)
    });
    laneConfigArgs[1] = FastTransferTokenPoolAbstract.DestChainConfigUpdateArgs({
      remoteChainSelector: SVM_CHAIN_SELECTOR,
      fastTransferBpsFee: FAST_FEE_BPS, // 1%
      fillerAllowlistEnabled: true,
      destinationPool: destPoolAddress,
      maxFillAmountPerRequest: MAX_FILL_AMOUNT_PER_REQUEST,
      settlementOverheadGas: SETTLEMENT_GAS_OVERHEAD,
      chainFamilySelector: Internal.CHAIN_FAMILY_SELECTOR_SVM,
      addFillers: addFillers,
      removeFillers: new address[](0),
      evmToAnyMessageExtraArgsBytes: svmExtraArgsBytesEncoded
    });
    s_tokenPool = new FastTransferTokenPoolHelper(
      s_token,
      18, // localTokenDecimals
      new address[](0), // allowlist
      address(s_mockRMNRemote), // rmnProxy
      address(s_sourceRouter) // router
    );
    s_tokenPool.updateDestChainConfig(laneConfigArgs[0]);
    s_tokenPool.updateDestChainConfig(laneConfigArgs[1]);
    bytes[] memory remotePoolAddresses = new bytes[](1);
    remotePoolAddresses[0] = destPoolAddress;

    TokenPool.ChainUpdate[] memory chainUpdate = new TokenPool.ChainUpdate[](2);
    chainUpdate[0] = TokenPool.ChainUpdate({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      remotePoolAddresses: remotePoolAddresses,
      remoteTokenAddress: abi.encode(address(2)),
      outboundRateLimiterConfig: _getOutboundRateLimiterConfig(),
      inboundRateLimiterConfig: _getInboundRateLimiterConfig()
    });
    chainUpdate[1] = TokenPool.ChainUpdate({
      remoteChainSelector: SVM_CHAIN_SELECTOR,
      remotePoolAddresses: remotePoolAddresses,
      remoteTokenAddress: abi.encode(address(2)),
      outboundRateLimiterConfig: _getOutboundRateLimiterConfig(),
      inboundRateLimiterConfig: _getInboundRateLimiterConfig()
    });

    s_tokenPool.applyChainUpdates(new uint64[](0), chainUpdate);

    // Approve tokens
    s_token.approve(address(s_tokenPool), type(uint256).max);
  }
}
