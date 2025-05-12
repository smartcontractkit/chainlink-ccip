pragma solidity ^0.8.24;

import {IOwner} from "../../interfaces/IOwner.sol";

import {Router} from "../../Router.sol";

import {Client} from "../../libraries/Client.sol";
import {Internal} from "../../libraries/Internal.sol";

import {Pool} from "../../libraries/Pool.sol";
import {RateLimiter} from "../../libraries/RateLimiter.sol";
import {LockReleaseTokenPool} from "../../pools/LockReleaseTokenPool.sol";
import {TokenPool} from "../../pools/TokenPool.sol";
import {TokenAdminRegistry} from "../../tokenAdminRegistry/TokenAdminRegistry.sol";

import {IERC20} from "@chainlink/vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "@chainlink/vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/utils/SafeERC20.sol";
import {Test} from "forge-std/Test.sol";

interface LegacyTokenPool {
  // The current link token pool is legacy and uses a different struct than current pools
  struct RampUpdate {
    address ramp;
    bool allowed;
    RateLimiter.Config rateLimiterConfig;
  }

  // 1.0 LINK Token Pool Functions
  function applyRampUpdates(RampUpdate[] calldata onRamps, RampUpdate[] calldata offRamps) external;
  function releaseOrMint(bytes memory, address receiver, uint256 amount, uint64, bytes memory) external;
}

contract LinkMigrationTest is Test {
  using SafeERC20 for IERC20;

  address public constant ROUTER = 0x80226fc0Ee2b096224EeAc085Bb9a8cba1146f7D;
  address public constant PREVIOUS_POOL = 0xC2291992A08eBFDfedfE248F2CCD34Da63570DF4;
  address public constant RMN_PROXY = 0x411dE17f12D1A34ecC7F45f49844626267c75e81;
  TokenAdminRegistry public constant TOKEN_ADMIN_REGISTRY =
    TokenAdminRegistry(0xb22764f98dD05c789929716D677382Df22C05Cb6);
  IERC20 public constant LINK_TOKEN = IERC20(0x514910771AF9Ca656af840dff83E8264EcF986CA);

  uint64 public constant MOCK_CHAIN_SELECTOR = type(uint64).max - 1;
  uint64 public constant BASE_CHAIN_SELECTOR = 15971525489660198786;
  address public constant ETHEREUM_BASE_ONRAMP = 0xb8a882f3B88bd52D1Ff56A873bfDB84b70431937;
  address public constant ETHEREUM_OFFRAMP_FROM_BASE = 0x6B4B6359Dd5B47Cdb030E5921456D2a0625a9EbD;
  address public constant BASE_LINK_POOL = 0x0A995a72D8346683c97514990F802F4778B7ac72;
  address public constant LINK_TOKEN_BASE = 0x88Fb150BDc53A65fe94Dea0c9BA0a6dAf8C6e196;

  address public alice = makeAddr("ALICE");
  address public currentPoolOwner;

  LockReleaseTokenPool public newTokenPool;

  function testLinkMigration() public {
    vm.createSelectFork(vm.envString("ETHEREUM_RPC_URL"));

    // Get the owner of the current pool, should just be MCMS
    currentPoolOwner = IOwner(PREVIOUS_POOL).owner();

    // Prank the owner to deploy a new pool
    vm.startPrank(currentPoolOwner);

    // Deploy the new pool
    newTokenPool = new LockReleaseTokenPool(
      LINK_TOKEN, // token
      18, // localTokenDecimals
      new address[](0), // allowList
      RMN_PROXY, // rmnProxy
      true, // acceptLiquidity
      ROUTER // router
    );

    // The Release Amount is the current balance of the Legacy Pool
    uint256 releaseAmount = LINK_TOKEN.balanceOf(PREVIOUS_POOL);

    // Craft a ramp update setting the timelock as an OffRamp
    LegacyTokenPool.RampUpdate[] memory offRampUpdates = new LegacyTokenPool.RampUpdate[](1);
    offRampUpdates[0] = LegacyTokenPool.RampUpdate({
      ramp: currentPoolOwner,
      allowed: true,
      rateLimiterConfig: RateLimiter.Config({isEnabled: false, capacity: 0, rate: 0})
    });

    // Apply the Ramp Update to the legacy pool so that a release TX can be sent
    LegacyTokenPool(PREVIOUS_POOL).applyRampUpdates(new LegacyTokenPool.RampUpdate[](0), offRampUpdates);

    // Release the tokens directly into the new token pool without an intermediary
    LegacyTokenPool(PREVIOUS_POOL).releaseOrMint("", address(newTokenPool), releaseAmount, 0, "");

    // Check that the tokens were moved from the legacy pool to the new one
    assertEq(LINK_TOKEN.balanceOf(address(newTokenPool)), releaseAmount, "NEW TOKEN POOL BALANCE INCORRECT");
    assertEq(LINK_TOKEN.balanceOf(PREVIOUS_POOL), 0, "PREVIOUS POOL BALANCE SHOULD BE 0");

    // Set the new token pool in the token admin registry
    address tokenAdmin = TOKEN_ADMIN_REGISTRY.getTokenConfig(address(LINK_TOKEN)).administrator;
    vm.startPrank(tokenAdmin);
    TOKEN_ADMIN_REGISTRY.setPool(address(LINK_TOKEN), address(newTokenPool));

    assertEq(
      TOKEN_ADMIN_REGISTRY.getTokenConfig(address(LINK_TOKEN)).tokenPool,
      address(newTokenPool),
      "Token pool not set in admin registry!"
    );

    // Set a new remote chain in the new token pool so that we can test messages from other chains
    TokenPool.ChainUpdate[] memory chainUpdates = new TokenPool.ChainUpdate[](1);
    bytes[] memory remotePools = new bytes[](1);
    remotePools[0] = abi.encode(BASE_LINK_POOL);
    chainUpdates[0] = TokenPool.ChainUpdate({
      remoteChainSelector: BASE_CHAIN_SELECTOR,
      remotePoolAddresses: remotePools,
      remoteTokenAddress: abi.encode(LINK_TOKEN_BASE),
      outboundRateLimiterConfig: RateLimiter.Config({isEnabled: false, capacity: 0, rate: 0}),
      inboundRateLimiterConfig: RateLimiter.Config({isEnabled: false, capacity: 0, rate: 0})
    });
    newTokenPool.applyChainUpdates(new uint64[](0), chainUpdates);

    // Test that the new token Pool is being used correctly for outgoing messages
    uint256 alicePreAmount = 100 ether;
    deal(alice, alicePreAmount);
    deal(address(LINK_TOKEN), alice, alicePreAmount);
    uint256 sendAmount = 1 ether;

    vm.startPrank(alice);
    LINK_TOKEN.safeIncreaseAllowance(ROUTER, type(uint256).max);
    Client.EVMTokenAmount[] memory tokens = new Client.EVMTokenAmount[](1);
    tokens[0] = Client.EVMTokenAmount({token: address(LINK_TOKEN), amount: sendAmount});

    {
      // Craft the message
      Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
        receiver: abi.encode(alice),
        data: "",
        tokenAmounts: tokens,
        feeToken: address(LINK_TOKEN),
        extraArgs: ""
      });
      uint256 fee = Router(ROUTER).getFee(BASE_CHAIN_SELECTOR, message);

      // Send the outgoing message
      vm.expectEmit();
      emit TokenPool.Locked(ETHEREUM_BASE_ONRAMP, sendAmount);
      Router(ROUTER).ccipSend(BASE_CHAIN_SELECTOR, message);

      // Test an incoming message from that same chain
      vm.startPrank(ETHEREUM_OFFRAMP_FROM_BASE);
      Pool.ReleaseOrMintInV1 memory testReleaseData = Pool.ReleaseOrMintInV1({
        originalSender: abi.encode(alice),
        remoteChainSelector: BASE_CHAIN_SELECTOR,
        receiver: alice,
        amount: sendAmount,
        localToken: address(LINK_TOKEN),
        sourcePoolAddress: abi.encode(BASE_LINK_POOL),
        sourcePoolData: "",
        offchainTokenData: ""
      });
      // Check that funds were removed from Alice's wallet for sending the message
      assertEq(LINK_TOKEN.balanceOf(alice), alicePreAmount - sendAmount - fee);

      vm.expectEmit();
      emit TokenPool.Released(ETHEREUM_OFFRAMP_FROM_BASE, alice, sendAmount);
      newTokenPool.releaseOrMint(testReleaseData);

      // Alice's balance after funds are returned should be the same minus the amount
      // used to pay for the initial outgoing message
      assertEq(LINK_TOKEN.balanceOf(alice), alicePreAmount - fee);
    }
  }
}
