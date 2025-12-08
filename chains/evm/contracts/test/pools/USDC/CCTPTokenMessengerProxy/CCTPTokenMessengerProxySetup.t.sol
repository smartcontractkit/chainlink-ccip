// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {CCTPTokenMessengerProxy} from "../../../../pools/USDC/CCTPTokenMessengerProxy.sol";

import {BaseTest} from "../../../BaseTest.t.sol";
import {MockUSDCTokenMessenger} from "../../../mocks/MockUSDCTokenMessenger.sol";
import {BurnMintERC20} from "@chainlink/contracts/src/v0.8/shared/token/ERC20/BurnMintERC20.sol";

contract CCTPTokenMessengerProxySetup is BaseTest {
  uint32 internal constant MESSAGE_BODY_VERSION = 1;
  uint256 internal constant BURN_AMOUNT = 1000e6;

  MockUSDCTokenMessenger internal s_tokenMessenger;
  address internal s_messageTransmitter;
  address internal s_authorizedCaller;
  CCTPTokenMessengerProxy internal s_cctpTokenMessengerProxy;
  BurnMintERC20 internal s_USDCToken;

  uint32 internal constant DESTINATION_DOMAIN = 1;
  uint32 internal constant MAX_FEE = 100;
  uint32 internal constant MIN_FINALITY_THRESHOLD = 2000;

  bytes32 internal s_mintRecipient = bytes32(uint256(uint160(makeAddr("MINT_RECIPIENT"))));
  bytes32 internal s_destinationCaller = bytes32(uint256(uint160(makeAddr("DESTINATION_CALLER"))));

  function setUp() public virtual override {
    super.setUp();
    s_messageTransmitter = makeAddr("MESSAGE_TRANSMITTER");
    s_tokenMessenger = new MockUSDCTokenMessenger(MESSAGE_BODY_VERSION, s_messageTransmitter);
    s_authorizedCaller = makeAddr("AUTHORIZED_CALLER");
    s_USDCToken = new BurnMintERC20("USD Coin", "USDC", 6, 0, 0);

    address[] memory authorizedCallers = new address[](1);
    authorizedCallers[0] = s_authorizedCaller;
    s_cctpTokenMessengerProxy = new CCTPTokenMessengerProxy(s_tokenMessenger, s_USDCToken, authorizedCallers);

    deal(address(s_USDCToken), address(s_cctpTokenMessengerProxy), BURN_AMOUNT);

    // Grant mint and burn roles to the token messenger proxy and the message transmitter.
    BurnMintERC20(address(s_USDCToken)).grantMintAndBurnRoles(address(s_tokenMessenger));
  }
}
