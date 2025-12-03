// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolV1} from "../../../interfaces/IPool.sol";

import {Router} from "../../../Router.sol";
import {Pool} from "../../../libraries/Pool.sol";
import {TokenPool} from "../../../pools/TokenPool.sol";
import {CCTPMessageTransmitterProxy} from "../../../pools/USDC/CCTPMessageTransmitterProxy.sol";
import {TokenAdminRegistry} from "../../../tokenAdminRegistry/TokenAdminRegistry.sol";
import {BaseTest} from "../../BaseTest.t.sol";
import {MockE2EUSDCTransmitter} from "../../mocks/MockE2EUSDCTransmitter.sol";
import {MockE2EUSDCTransmitterCCTPV2} from "../../mocks/MockE2EUSDCTransmitterCCTPV2.sol";
import {MockUSDCTokenMessenger} from "../../mocks/MockUSDCTokenMessenger.sol";

import {BurnMintERC20} from "@chainlink/contracts/src/v0.8/shared/token/ERC20/BurnMintERC20.sol";
import {IBurnMintERC20} from "@chainlink/contracts/src/v0.8/shared/token/ERC20/IBurnMintERC20.sol";
import {IERC165} from "@openzeppelin/contracts@5.0.2/utils/introspection/IERC165.sol";

contract USDCSetup is BaseTest {
  struct USDCMessage {
    uint32 version;
    uint32 sourceDomain;
    uint32 destinationDomain;
    uint64 nonce;
    bytes32 sender;
    bytes32 recipient;
    bytes32 destinationCaller;
    bytes messageBody;
  }

  // solhint-disable-next-line gas-struct-packing
  struct USDCMessageCCTPV2 {
    uint32 version;
    uint32 sourceDomain;
    uint32 destinationDomain;
    bytes32 nonce;
    bytes32 sender;
    bytes32 recipient;
    bytes32 destinationCaller;
    uint32 minFinalityThreshold;
    uint32 finalityThresholdExecuted;
    bytes messageBody;
  }

  uint32 internal constant USDC_DEST_TOKEN_GAS = 180_000;
  uint32 internal constant SOURCE_DOMAIN_IDENTIFIER = 0x02020202;
  uint32 internal constant DEST_DOMAIN_IDENTIFIER = 0;

  bytes32 internal constant SOURCE_CHAIN_TOKEN_SENDER = bytes32(uint256(uint160(0x01111111221)));
  address internal constant SOURCE_CHAIN_USDC_POOL = address(0x23789765456789);
  address internal constant DEST_CHAIN_USDC_POOL = address(0x987384873458734);
  address internal constant DEST_CHAIN_USDC_TOKEN = address(0x23598918358198766);

  MockUSDCTokenMessenger internal s_mockUSDCTokenMessenger;
  MockUSDCTokenMessenger internal s_mockUSDCTokenMessenger_CCTPV1;
  MockE2EUSDCTransmitter internal s_mockUSDCTransmitter;
  MockE2EUSDCTransmitterCCTPV2 internal s_mockUSDCTransmitterCCTPV2;

  CCTPMessageTransmitterProxy internal s_cctpMessageTransmitterProxy;

  address internal s_routerAllowedOnRamp = address(3456);
  address internal s_routerAllowedOffRamp = address(234);
  address internal s_previousPool = makeAddr("previousPool");
  address internal s_previousPoolMessageTransmitterProxy = makeAddr("previousPoolMessageTransmitterProxy");
  Router internal s_router;

  TokenAdminRegistry internal s_tokenAdminRegistry;

  IBurnMintERC20 internal s_USDCToken;

  function setUp() public virtual override {
    super.setUp();
    BurnMintERC20 usdcToken = new BurnMintERC20("USD Coin", "USDC", 6, 0, 0);
    s_USDCToken = usdcToken;

    s_tokenAdminRegistry = new TokenAdminRegistry();

    deal(address(s_USDCToken), OWNER, type(uint256).max);
    _setUpRamps();

    s_mockUSDCTransmitterCCTPV2 = new MockE2EUSDCTransmitterCCTPV2(1, DEST_DOMAIN_IDENTIFIER, address(s_USDCToken));
    s_mockUSDCTransmitter = new MockE2EUSDCTransmitter(0, DEST_DOMAIN_IDENTIFIER, address(s_USDCToken));

    // Create both of the mock token messengers, one for CCTP V1 and one for CCTP V2. The V1 messenger is
    // denoted by it's version being 0 and using the mock transmitter with the same version
    s_mockUSDCTokenMessenger = new MockUSDCTokenMessenger(1, address(s_mockUSDCTransmitterCCTPV2));
    s_mockUSDCTokenMessenger_CCTPV1 = new MockUSDCTokenMessenger(0, address(s_mockUSDCTransmitter));

    s_cctpMessageTransmitterProxy = new CCTPMessageTransmitterProxy(s_mockUSDCTokenMessenger);

    usdcToken.grantMintAndBurnRoles(address(s_mockUSDCTransmitterCCTPV2));
    usdcToken.grantMintAndBurnRoles(address(s_mockUSDCTokenMessenger));

    // Mock the previous pool's releaseOrMint function to return the input amount
    vm.mockCall(
      s_previousPool,
      abi.encodeWithSelector(TokenPool.releaseOrMint.selector),
      abi.encode(Pool.ReleaseOrMintOutV1({destinationAmount: 1}))
    );

    // Mock the previous pool's i_cctpMessageTransmitterProxy function to return an address
    // This is used to determine if the message was sent using CCTP V1 or V2
    vm.mockCall(
      s_previousPool,
      abi.encodeWithSelector(bytes4(keccak256("i_messageTransmitterProxy()"))),
      abi.encode(s_previousPoolMessageTransmitterProxy)
    );

    // Mock the previous pool's supportsInterface function to return true for IPoolV1 interface
    vm.mockCall(
      s_previousPool,
      abi.encodeWithSelector(IERC165.supportsInterface.selector, type(IPoolV1).interfaceId),
      abi.encode(true)
    );

    // Mock the previous pool's message transmitter proxy to return true for IPoolV1 interface
    vm.mockCall(
      s_previousPoolMessageTransmitterProxy,
      abi.encodeWithSelector(CCTPMessageTransmitterProxy.receiveMessage.selector),
      abi.encode(true)
    );
  }

  function _poolApplyChainUpdates(
    address pool
  ) internal {
    bytes[] memory sourcePoolAddresses = new bytes[](1);
    sourcePoolAddresses[0] = abi.encode(SOURCE_CHAIN_USDC_POOL);

    bytes[] memory destPoolAddresses = new bytes[](1);
    destPoolAddresses[0] = abi.encode(DEST_CHAIN_USDC_POOL);

    TokenPool.ChainUpdate[] memory chainUpdates = new TokenPool.ChainUpdate[](2);
    chainUpdates[0] = TokenPool.ChainUpdate({
      remoteChainSelector: SOURCE_CHAIN_SELECTOR,
      remotePoolAddresses: sourcePoolAddresses,
      remoteTokenAddress: abi.encode(address(s_USDCToken)),
      outboundRateLimiterConfig: _getOutboundRateLimiterConfig(),
      inboundRateLimiterConfig: _getInboundRateLimiterConfig()
    });
    chainUpdates[1] = TokenPool.ChainUpdate({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      remotePoolAddresses: destPoolAddresses,
      remoteTokenAddress: abi.encode(DEST_CHAIN_USDC_TOKEN),
      outboundRateLimiterConfig: _getOutboundRateLimiterConfig(),
      inboundRateLimiterConfig: _getInboundRateLimiterConfig()
    });

    TokenPool(pool).applyChainUpdates(new uint64[](0), chainUpdates);
  }

  function _setUpRamps() internal {
    s_router = new Router(address(s_USDCToken), address(s_mockRMNRemote));

    Router.OnRamp[] memory onRampUpdates = new Router.OnRamp[](1);
    onRampUpdates[0] = Router.OnRamp({destChainSelector: DEST_CHAIN_SELECTOR, onRamp: s_routerAllowedOnRamp});
    Router.OffRamp[] memory offRampUpdates = new Router.OffRamp[](1);
    address[] memory offRamps = new address[](1);
    offRamps[0] = s_routerAllowedOffRamp;
    offRampUpdates[0] = Router.OffRamp({sourceChainSelector: SOURCE_CHAIN_SELECTOR, offRamp: offRamps[0]});

    s_router.applyRampUpdates(onRampUpdates, new Router.OffRamp[](0), offRampUpdates);
  }

  function _generateUSDCMessage(
    USDCMessage memory usdcMessage
  ) internal pure returns (bytes memory) {
    return abi.encodePacked(
      usdcMessage.version,
      usdcMessage.sourceDomain,
      usdcMessage.destinationDomain,
      usdcMessage.nonce,
      usdcMessage.sender,
      usdcMessage.recipient,
      usdcMessage.destinationCaller,
      usdcMessage.messageBody
    );
  }

  function _generateUSDCMessageCCTPV2(
    USDCMessageCCTPV2 memory usdcMessage
  ) internal pure returns (bytes memory) {
    return abi.encodePacked(
      usdcMessage.version,
      usdcMessage.sourceDomain,
      usdcMessage.destinationDomain,
      usdcMessage.nonce,
      usdcMessage.sender,
      usdcMessage.recipient,
      usdcMessage.destinationCaller,
      usdcMessage.minFinalityThreshold,
      usdcMessage.finalityThresholdExecuted,
      usdcMessage.messageBody
    );
  }
}
