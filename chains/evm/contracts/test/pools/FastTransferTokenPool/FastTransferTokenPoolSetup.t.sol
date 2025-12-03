// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Router} from "../../../Router.sol";
import {Client} from "../../../libraries/Client.sol";
import {Internal} from "../../../libraries/Internal.sol";
import {FastTransferTokenPoolAbstract} from "../../../pools/FastTransferTokenPoolAbstract.sol";
import {TokenPool} from "../../../pools/TokenPool.sol";
import {BaseTest} from "../../BaseTest.t.sol";
import {FastTransferTokenPoolHelper} from "../../helpers/FastTransferTokenPoolHelper.sol";
import {BurnMintERC20} from "@chainlink/contracts/src/v0.8/shared/token/ERC20/BurnMintERC20.sol";
import {WETH9} from "@chainlink/contracts/src/v0.8/vendor/canonical-weth/WETH9.sol";

import {IERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/IERC20.sol";

contract FastTransferTokenPoolSetup is BaseTest {
  uint256 public constant SOURCE_AMOUNT = 100 ether;
  uint8 public constant SOURCE_DECIMALS = 18;
  uint16 public constant FAST_FEE_FILLER_BPS = 100; // 1%
  uint16 public constant FAST_FEE_POOL_BPS = 100; // 1%
  address public constant RECEIVER = address(0x5);

  uint32 internal constant SVM_CHAIN_SELECTOR = uint32(uint256(keccak256("SVM_SELECTOR")));
  uint32 internal constant SETTLEMENT_GAS_OVERHEAD = 200_000;
  uint256 internal constant MAX_FILL_AMOUNT_PER_REQUEST = 1000 ether;

  bytes internal s_svmExtraArgsBytesEncoded = Client._svmArgsToBytes(
    Client.SVMExtraArgsV1({
      computeUnits: SETTLEMENT_GAS_OVERHEAD,
      accounts: new bytes32[](0),
      accountIsWritableBitmap: 2,
      allowOutOfOrderExecution: true,
      tokenReceiver: bytes32(0)
    })
  );
  IERC20 internal s_token;
  FastTransferTokenPoolHelper public s_pool;
  WETH9 public wrappedNative;
  bytes public destPoolAddress = abi.encode(makeAddr("destPool"));
  address public s_filler = makeAddr("filler");

  function setUp() public virtual override {
    super.setUp();

    Router.OnRamp[] memory onRampUpdates = new Router.OnRamp[](1);
    onRampUpdates[0] = Router.OnRamp({destChainSelector: DEST_CHAIN_SELECTOR, onRamp: makeAddr("onRamp")});
    s_sourceRouter.applyRampUpdates(onRampUpdates, new Router.OffRamp[](0), new Router.OffRamp[](0));

    s_token = new BurnMintERC20("LINK", "LNK", 18, 0, 0);

    deal(address(s_token), OWNER, type(uint256).max);
    wrappedNative = new WETH9();

    // Deploy pool
    FastTransferTokenPoolAbstract.DestChainConfigUpdateArgs[] memory laneConfigArgs =
      new FastTransferTokenPoolAbstract.DestChainConfigUpdateArgs[](2);
    laneConfigArgs[0] = FastTransferTokenPoolAbstract.DestChainConfigUpdateArgs({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      fastTransferFillerFeeBps: FAST_FEE_FILLER_BPS, // 1%
      fastTransferPoolFeeBps: 0, // No pool fee for this test
      fillerAllowlistEnabled: true,
      destinationPool: destPoolAddress,
      maxFillAmountPerRequest: MAX_FILL_AMOUNT_PER_REQUEST,
      settlementOverheadGas: SETTLEMENT_GAS_OVERHEAD,
      chainFamilySelector: Internal.CHAIN_FAMILY_SELECTOR_EVM,
      customExtraArgs: ""
    });
    laneConfigArgs[1] = FastTransferTokenPoolAbstract.DestChainConfigUpdateArgs({
      remoteChainSelector: SVM_CHAIN_SELECTOR,
      fastTransferFillerFeeBps: FAST_FEE_FILLER_BPS, // 1%
      fastTransferPoolFeeBps: 0, // No pool fee for this test
      fillerAllowlistEnabled: true,
      destinationPool: destPoolAddress,
      maxFillAmountPerRequest: MAX_FILL_AMOUNT_PER_REQUEST,
      settlementOverheadGas: 0,
      chainFamilySelector: Internal.CHAIN_FAMILY_SELECTOR_SVM,
      customExtraArgs: s_svmExtraArgsBytesEncoded
    });

    s_pool = new FastTransferTokenPoolHelper(
      s_token,
      18, // localTokenDecimals
      new address[](0), // allowlist
      address(s_mockRMNRemote), // rmnProxy
      address(s_sourceRouter), // router
      SOURCE_CHAIN_SELECTOR // sourceChainSelector
    );

    s_pool.updateDestChainConfig(laneConfigArgs);

    address[] memory fillersToAdd = new address[](1);
    fillersToAdd[0] = s_filler;

    s_pool.updateFillerAllowList(fillersToAdd, new address[](0));

    bytes[] memory remotePoolAddresses = new bytes[](1);
    remotePoolAddresses[0] = destPoolAddress;

    TokenPool.ChainUpdate[] memory chainUpdates = new TokenPool.ChainUpdate[](2);
    chainUpdates[0] = TokenPool.ChainUpdate({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      remotePoolAddresses: remotePoolAddresses,
      remoteTokenAddress: abi.encode(address(2)),
      outboundRateLimiterConfig: _getOutboundRateLimiterConfig(),
      inboundRateLimiterConfig: _getInboundRateLimiterConfig()
    });
    chainUpdates[1] = TokenPool.ChainUpdate({
      remoteChainSelector: SVM_CHAIN_SELECTOR,
      remotePoolAddresses: remotePoolAddresses,
      remoteTokenAddress: abi.encode(address(2)),
      outboundRateLimiterConfig: _getOutboundRateLimiterConfig(),
      inboundRateLimiterConfig: _getInboundRateLimiterConfig()
    });

    s_pool.applyChainUpdates(new uint64[](0), chainUpdates);

    // Approve tokens from the OWNER for the pool
    s_token.approve(address(s_pool), type(uint256).max);
  }

  function _singleConfigToList(
    FastTransferTokenPoolAbstract.DestChainConfigUpdateArgs memory config
  ) internal pure returns (FastTransferTokenPoolAbstract.DestChainConfigUpdateArgs[] memory) {
    FastTransferTokenPoolAbstract.DestChainConfigUpdateArgs[] memory list =
      new FastTransferTokenPoolAbstract.DestChainConfigUpdateArgs[](1);
    list[0] = config;
    return list;
  }
}
