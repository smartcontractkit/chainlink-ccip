// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Router} from "../../../Router.sol";
import {Internal} from "../../../libraries/Internal.sol";
import {BurnMintFastTransferTokenPool} from "../../../pools/BurnMintFastTransferTokenPool.sol";
import {FastTransferTokenPoolAbstract} from "../../../pools/FastTransferTokenPoolAbstract.sol";
import {TokenPool} from "../../../pools/TokenPool.sol";
import {BaseTest} from "../../BaseTest.t.sol";
import {BurnMintERC20} from "@chainlink/contracts/src/v0.8/shared/token/ERC20/BurnMintERC20.sol";

contract BurnMintFastTransferTokenPoolSetup is BaseTest {
  BurnMintFastTransferTokenPool internal s_pool;
  BurnMintERC20 internal s_token;

  address internal s_burnMintOffRamp = makeAddr("burn_mint_offRamp");
  address internal s_burnMintOnRamp = makeAddr("burn_mint_onRamp");
  address internal s_remoteBurnMintPool = makeAddr("remote_burn_mint_pool");
  address internal s_remoteToken = makeAddr("remote_token");
  address internal s_filler = makeAddr("filler");

  uint256 internal constant TRANSFER_AMOUNT = 100 ether;
  address internal constant RECEIVER = address(0x1234);
  uint8 internal constant SOURCE_DECIMALS = 18;
  uint256 internal constant FILL_AMOUNT = 100 ether;

  uint16 internal constant FAST_FEE_FILLER_BPS = 100; // 1%
  uint256 internal constant FILL_AMOUNT_MAX = 1000 ether;
  uint32 internal constant SETTLEMENT_GAS_OVERHEAD = 200_000;

  function setUp() public virtual override {
    super.setUp();

    s_token = new BurnMintERC20("Chainlink Token", "LINK", 18, 0, 0);

    s_pool = new BurnMintFastTransferTokenPool(
      s_token,
      DEFAULT_TOKEN_DECIMALS,
      new address[](0),
      address(s_mockRMNRemote),
      address(s_sourceRouter),
      SOURCE_CHAIN_SELECTOR
    );

    s_token.grantMintAndBurnRoles(address(s_pool));

    _applyChainUpdates();
    _setupDestChainConfig();
  }

  function _applyChainUpdates() internal {
    bytes[] memory remotePoolAddresses = new bytes[](1);
    remotePoolAddresses[0] = abi.encode(s_remoteBurnMintPool);

    TokenPool.ChainUpdate[] memory chainsToAdd = new TokenPool.ChainUpdate[](1);
    chainsToAdd[0] = TokenPool.ChainUpdate({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      remotePoolAddresses: remotePoolAddresses,
      remoteTokenAddress: abi.encode(s_remoteToken),
      outboundRateLimiterConfig: _getOutboundRateLimiterConfig(),
      inboundRateLimiterConfig: _getInboundRateLimiterConfig()
    });

    s_pool.applyChainUpdates(new uint64[](0), chainsToAdd);

    Router.OnRamp[] memory onRampUpdates = new Router.OnRamp[](1);
    onRampUpdates[0] = Router.OnRamp({destChainSelector: DEST_CHAIN_SELECTOR, onRamp: s_burnMintOnRamp});
    Router.OffRamp[] memory offRampUpdates = new Router.OffRamp[](1);
    offRampUpdates[0] = Router.OffRamp({sourceChainSelector: DEST_CHAIN_SELECTOR, offRamp: s_burnMintOffRamp});
    s_sourceRouter.applyRampUpdates(onRampUpdates, new Router.OffRamp[](0), offRampUpdates);
  }

  function _setupDestChainConfig() internal {
    address[] memory addFillers = new address[](1);
    addFillers[0] = s_filler;

    FastTransferTokenPoolAbstract.DestChainConfigUpdateArgs memory laneConfigArgs = FastTransferTokenPoolAbstract
      .DestChainConfigUpdateArgs({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      fastTransferFillerFeeBps: FAST_FEE_FILLER_BPS,
      fastTransferPoolFeeBps: 0, // No pool fee for this test
      fillerAllowlistEnabled: true,
      destinationPool: abi.encode(s_remoteBurnMintPool),
      maxFillAmountPerRequest: FILL_AMOUNT_MAX,
      settlementOverheadGas: SETTLEMENT_GAS_OVERHEAD,
      chainFamilySelector: Internal.CHAIN_FAMILY_SELECTOR_EVM,
      customExtraArgs: ""
    });

    s_pool.updateDestChainConfig(_singleConfigToList(laneConfigArgs));
    s_pool.updateFillerAllowList(addFillers, new address[](0));
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
