// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {CCIPReceiver} from "../../applications/CCIPReceiver.sol";
import {IRouterClient} from "../../interfaces/IRouterClient.sol";
import {Client} from "../../libraries/Client.sol";
import {ExtraArgsCodec} from "../../libraries/ExtraArgsCodec.sol";
import {Internal} from "../../libraries/Internal.sol";
import {OffRamp} from "../../offRamp/OffRamp.sol";
import {OnRamp} from "../../onRamp/OnRamp.sol";
import {BurnMintTokenPool} from "../../pools/BurnMintTokenPool.sol";
import {TokenPool} from "../../pools/TokenPool.sol";
import {BaseERC20} from "../../tmp/BaseERC20.sol";
import {CrossChainToken} from "../../tmp/CrossChainToken.sol";
import {RegistryModuleOwnerCustom} from "../../tokenAdminRegistry/RegistryModuleOwnerCustom.sol";
import {TokenPoolFactory} from "../../tokenAdminRegistry/TokenPoolFactory/TokenPoolFactory.sol";
import {e2e} from "./e2e.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

import {IERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/IERC20.sol";
import {VmSafe} from "forge-std/Vm.sol";

/// @notice End-to-end test that deploys a token and pool via the TokenPoolFactory,
/// claims the CCIP admin, and sends a cross-chain token transfer using the Router.
contract e2e_factoryDeployedPool is e2e {
  TokenPoolFactory internal s_factory;
  RegistryModuleOwnerCustom internal s_registryModule;

  address internal s_factoryToken;
  address internal s_factoryPool;
  address internal s_destToken;
  address internal s_destPool;

  uint256 internal constant PREMINT = 1000e18;
  uint256 internal constant TOKEN_TRANSFER_AMOUNT = 10e18;

  bytes internal constant POOL_INIT_CODE = type(BurnMintTokenPool).creationCode;
  bytes32 internal constant SALT = keccak256("e2e_factory_test");

  function setUp() public virtual override {
    super.setUp();

    s_registryModule = new RegistryModuleOwnerCustom(address(s_tokenAdminRegistry));
    s_tokenAdminRegistry.addRegistryModule(address(s_registryModule));

    s_factory =
      new TokenPoolFactory(s_tokenAdminRegistry, s_registryModule, address(s_mockRMNRemote), address(s_sourceRouter));

    // Deploy source token + pool via factory

    bytes memory tokenInitCode = abi.encodePacked(
      type(CrossChainToken).creationCode,
      abi.encode(
        BaseERC20.ConstructorParams({
          name: "FactoryToken", symbol: "FTK", decimals: 18, maxSupply: 0, preMint: 0, ccipAdmin: address(s_factory)
        }),
        address(s_factory),
        OWNER
      )
    );

    (s_factoryToken, s_factoryPool) = s_factory.deployTokenAndTokenPool(
      new TokenPoolFactory.RemoteTokenPoolInfo[](0),
      18,
      TokenPoolFactory.PoolType.BURN_MINT,
      tokenInitCode,
      POOL_INIT_CODE,
      address(0),
      SALT,
      address(0)
    );

    // Claim ownership: accept admin role and pool ownership

    s_tokenAdminRegistry.acceptAdminRole(s_factoryToken);
    Ownable2Step(s_factoryPool).acceptOwnership();

    // OWNER holds DEFAULT_ADMIN_ROLE which is admin of BURN_MINT_ADMIN_ROLE.
    // Grant BURN_MINT_ADMIN_ROLE to self, then use it to grant MINTER_ROLE and mint.
    CrossChainToken factoryToken = CrossChainToken(s_factoryToken);
    factoryToken.grantRole(factoryToken.BURN_MINT_ADMIN_ROLE(), OWNER);
    factoryToken.grantRole(factoryToken.MINTER_ROLE(), OWNER);
    factoryToken.mint(OWNER, PREMINT);

    // Deploy destination token manually, pool via a second factory (simulating dest chain)

    TokenPoolFactory destFactory =
      new TokenPoolFactory(s_tokenAdminRegistry, s_registryModule, address(s_mockRMNRemote), address(s_destRouter));

    CrossChainToken destToken = new CrossChainToken(
      BaseERC20.ConstructorParams({
        name: "FactoryToken", symbol: "FTK", decimals: 18, maxSupply: 0, preMint: 0, ccipAdmin: OWNER
      }),
      OWNER,
      OWNER
    );
    s_destToken = address(destToken);

    s_destPool = destFactory.deployTokenPoolWithExistingToken(
      s_destToken,
      18,
      TokenPoolFactory.PoolType.BURN_MINT,
      new TokenPoolFactory.RemoteTokenPoolInfo[](0),
      POOL_INIT_CODE,
      address(0),
      SALT,
      address(0)
    );
    Ownable2Step(s_destPool).acceptOwnership();

    destToken.grantMintAndBurnRoles(s_destPool);

    // Configure chain updates on both pools

    bytes[] memory remotePoolAddresses = new bytes[](1);
    TokenPool.ChainUpdate[] memory chainUpdates = new TokenPool.ChainUpdate[](1);

    // Source pool -> dest pool
    remotePoolAddresses[0] = abi.encode(s_destPool);
    chainUpdates[0] = TokenPool.ChainUpdate({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      remotePoolAddresses: remotePoolAddresses,
      remoteTokenAddress: abi.encode(s_destToken),
      outboundRateLimiterConfig: _getOutboundRateLimiterConfig(),
      inboundRateLimiterConfig: _getInboundRateLimiterConfig()
    });
    TokenPool(s_factoryPool).applyChainUpdates(new uint64[](0), chainUpdates);

    // Dest pool -> source pool
    remotePoolAddresses[0] = abi.encode(s_factoryPool);
    chainUpdates[0] = TokenPool.ChainUpdate({
      remoteChainSelector: SOURCE_CHAIN_SELECTOR,
      remotePoolAddresses: remotePoolAddresses,
      remoteTokenAddress: abi.encode(s_factoryToken),
      outboundRateLimiterConfig: _getOutboundRateLimiterConfig(),
      inboundRateLimiterConfig: _getInboundRateLimiterConfig()
    });
    TokenPool(s_destPool).applyChainUpdates(new uint64[](0), chainUpdates);

    // Register dest token+pool in the token admin registry (for OffRamp resolution)

    s_tokenAdminRegistry.proposeAdministrator(s_destToken, OWNER);
    s_tokenAdminRegistry.acceptAdminRole(s_destToken);
    s_tokenAdminRegistry.setPool(s_destToken, s_destPool);

    // Map source -> dest token for the event helper
    s_sourcePoolByToken[s_factoryToken] = s_factoryPool;
    s_destTokenBySourceToken[s_factoryToken] = s_destToken;
  }

  bytes32 internal constant CCIP_MESSAGE_SENT_TOPIC = keccak256(
    "CCIPMessageSent(uint64,address,bytes32,address,uint256,bytes,(address,uint32,uint32,uint256,bytes)[],bytes[])"
  );

  function test_e2e_factoryDeployedPool() public {
    uint256 senderBalanceBefore = IERC20(s_factoryToken).balanceOf(OWNER);
    uint256 sourceSupplyBefore = IERC20(s_factoryToken).totalSupply();
    uint256 destBalanceBefore = IERC20(s_destToken).balanceOf(OWNER);
    uint256 destSupplyBefore = IERC20(s_destToken).totalSupply();
    uint64 expectedMsgNum = s_onRamp.getDestChainConfig(DEST_CHAIN_SELECTOR).messageNumber + 1;

    assertEq(PREMINT, senderBalanceBefore);
    assertEq(PREMINT, sourceSupplyBefore);
    assertEq(0, destBalanceBefore);
    assertEq(0, destSupplyBefore);

    // Approve Router to spend tokens
    IERC20(s_factoryToken).approve(address(s_sourceRouter), TOKEN_TRANSFER_AMOUNT);
    IERC20(s_sourceFeeToken).approve(address(s_sourceRouter), type(uint256).max);

    Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
      receiver: abi.encode(OWNER),
      data: "",
      tokenAmounts: new Client.EVMTokenAmount[](1),
      feeToken: s_sourceFeeToken,
      extraArgs: ExtraArgsCodec._getBasicEncodedExtraArgsV3(0, 0)
    });
    message.tokenAmounts[0] = Client.EVMTokenAmount({token: s_factoryToken, amount: TOKEN_TRANSFER_AMOUNT});

    // Source: send via Router, tokens get burned
    vm.recordLogs();
    bytes32 messageId = s_sourceRouter.ccipSend(DEST_CHAIN_SELECTOR, message);

    assertEq(senderBalanceBefore - TOKEN_TRANSFER_AMOUNT, IERC20(s_factoryToken).balanceOf(OWNER));
    assertEq(sourceSupplyBefore - TOKEN_TRANSFER_AMOUNT, IERC20(s_factoryToken).totalSupply());

    // Extract encodedMessage from the CCIPMessageSent event
    bytes memory encodedMessage = _getEncodedMessageFromLogs(vm.getRecordedLogs());

    // Dest: execute on OffRamp, tokens get minted
    address[] memory ccvAddresses = new address[](1);
    ccvAddresses[0] = s_destVerifier;

    vm.expectEmit();
    emit OffRamp.ExecutionStateChanged({
      sourceChainSelector: SOURCE_CHAIN_SELECTOR,
      messageNumber: expectedMsgNum,
      messageId: messageId,
      state: Internal.MessageExecutionState.SUCCESS,
      returnData: ""
    });

    s_offRamp.execute(encodedMessage, ccvAddresses, new bytes[](1), 0);

    assertEq(TOKEN_TRANSFER_AMOUNT, IERC20(s_destToken).balanceOf(OWNER));
    assertEq(TOKEN_TRANSFER_AMOUNT, IERC20(s_destToken).totalSupply());
  }

  /// @notice End-to-end test: a sender/receiver pair enables cross-chain factory deployments via CCIP.
  /// The FactoryDeployerSender on the source chain sends a data message. The FactoryDeployerReceiver
  /// on the dest chain receives it, calls the factory with futureOwner from the payload, and the
  /// cross-chain user then accepts ownership.
  function test_e2e_remoteFactoryDeploymentViaCCIP() public {
    address alice = makeAddr("alice");

    TokenPoolFactory destFactory =
      new TokenPoolFactory(s_tokenAdminRegistry, s_registryModule, address(s_mockRMNRemote), address(s_destRouter));

    FactoryDeployer sourceDeployer = new FactoryDeployer(address(s_sourceRouter), s_factory, s_sourceFeeToken);
    FactoryDeployer destDeployer = new FactoryDeployer(address(s_destRouter), destFactory, s_sourceFeeToken);

    deal(s_sourceFeeToken, address(sourceDeployer), 5e18);

    bytes memory tokenInitCode = abi.encodePacked(
      type(CrossChainToken).creationCode,
      abi.encode(
        BaseERC20.ConstructorParams({
          name: "RemoteToken", symbol: "RMT", decimals: 18, maxSupply: 0, preMint: 0, ccipAdmin: address(destFactory)
        }),
        address(destFactory),
        alice
      )
    );

    vm.recordLogs();
    bytes32 messageId = sourceDeployer.requestRemoteDeployment(
      address(destDeployer),
      DEST_CHAIN_SELECTOR,
      alice,
      tokenInitCode,
      POOL_INIT_CODE,
      keccak256("remote_e2e_salt"),
      7_000_000
    );

    bytes memory encodedMessage = _getEncodedMessageFromLogs(vm.getRecordedLogs());

    address[] memory ccvAddresses = new address[](1);
    ccvAddresses[0] = s_destVerifier;

    vm.expectEmit();
    emit OffRamp.ExecutionStateChanged({
      sourceChainSelector: SOURCE_CHAIN_SELECTOR,
      messageNumber: s_onRamp.getDestChainConfig(DEST_CHAIN_SELECTOR).messageNumber,
      messageId: messageId,
      state: Internal.MessageExecutionState.SUCCESS,
      returnData: ""
    });

    s_offRamp.execute(encodedMessage, ccvAddresses, new bytes[](1), 0);

    (address deployedToken, address deployedPool) = destDeployer.lastDeployment();
    assertTrue(deployedToken != address(0), "Token should have been deployed");
    assertTrue(deployedPool != address(0), "Pool should have been deployed");

    vm.startPrank(alice);

    s_tokenAdminRegistry.acceptAdminRole(deployedToken);
    Ownable2Step(deployedPool).acceptOwnership();

    CrossChainToken token = CrossChainToken(deployedToken);

    // Verify: neither deployer contracts nor factory retain any permissions
    assertFalse(token.hasRole(token.DEFAULT_ADMIN_ROLE(), address(sourceDeployer)));
    assertFalse(token.hasRole(token.DEFAULT_ADMIN_ROLE(), address(destDeployer)));
    assertFalse(token.hasRole(token.DEFAULT_ADMIN_ROLE(), address(destFactory)));
    assertFalse(token.hasRole(token.BURN_MINT_ADMIN_ROLE(), address(destFactory)));
    assertNotEq(BurnMintTokenPool(deployedPool).owner(), address(sourceDeployer));
    assertNotEq(BurnMintTokenPool(deployedPool).owner(), address(destDeployer));
    assertNotEq(BurnMintTokenPool(deployedPool).owner(), address(destFactory));

    // Verify: alice has full control
    assertTrue(token.hasRole(token.DEFAULT_ADMIN_ROLE(), alice));
    assertEq(token.owner(), alice);
    assertEq(BurnMintTokenPool(deployedPool).owner(), alice);
    assertEq(s_tokenAdminRegistry.getTokenConfig(deployedToken).administrator, alice);

    // Verify: pool has mint and burn roles
    assertTrue(token.hasRole(token.MINTER_ROLE(), deployedPool));
    assertTrue(token.hasRole(token.BURNER_ROLE(), deployedPool));
  }

  function _getEncodedMessageFromLogs(
    VmSafe.Log[] memory logs
  ) private pure returns (bytes memory encodedMessage) {
    for (uint256 i = 0; i < logs.length; ++i) {
      if (logs[i].topics.length != 0 && logs[i].topics[0] == CCIP_MESSAGE_SENT_TOPIC) {
        (,, encodedMessage,,) = abi.decode(logs[i].data, (address, uint256, bytes, OnRamp.Receipt[], bytes[]));
        break;
      }
    }
    return encodedMessage;
  }
}

/// @notice A CCIP sender/receiver that enables cross-chain factory deployments.
/// On the source chain, call requestRemoteDeployment to send a CCIP message.
/// On the dest chain, the same contract receives the message and calls the factory.
contract FactoryDeployer is CCIPReceiver {
  TokenPoolFactory private immutable i_factory;
  IRouterClient private immutable i_routerClient;
  address private immutable i_feeToken;

  address public deployedToken;
  address public deployedPool;

  constructor(
    address router,
    TokenPoolFactory factory,
    address feeToken
  ) CCIPReceiver(router) {
    i_factory = factory;
    i_routerClient = IRouterClient(router);
    i_feeToken = feeToken;
    IERC20(feeToken).approve(router, type(uint256).max);
  }

  function requestRemoteDeployment(
    address remoteDeployer,
    uint64 destChainSelector,
    address futureOwner,
    bytes calldata tokenInitCode,
    bytes calldata poolInitCode,
    bytes32 salt,
    uint32 gasLimit
  ) external returns (bytes32 messageId) {
    Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
      receiver: abi.encode(remoteDeployer),
      data: abi.encode(futureOwner, tokenInitCode, poolInitCode, salt),
      tokenAmounts: new Client.EVMTokenAmount[](0),
      feeToken: i_feeToken,
      extraArgs: ExtraArgsCodec._getBasicEncodedExtraArgsV3(gasLimit, 0)
    });

    return i_routerClient.ccipSend(destChainSelector, message);
  }

  function _ccipReceive(
    Client.Any2EVMMessage memory message
  ) internal override {
    (address futureOwner, bytes memory tokenInitCode, bytes memory poolInitCode, bytes32 salt) =
      abi.decode(message.data, (address, bytes, bytes, bytes32));

    (deployedToken, deployedPool) = i_factory.deployTokenAndTokenPool(
      new TokenPoolFactory.RemoteTokenPoolInfo[](0),
      18,
      TokenPoolFactory.PoolType.BURN_MINT,
      tokenInitCode,
      poolInitCode,
      address(0),
      salt,
      futureOwner
    );
  }

  function lastDeployment() external view returns (address token, address pool) {
    return (deployedToken, deployedPool);
  }
}
