// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Pool} from "../../../libraries/Pool.sol";
import {CCTTokenPool} from "../../../pools/CCTTokenPool.sol";
import {BaseERC20} from "../../../tmp/BaseERC20.sol";
import {CCTTokenPoolSetup} from "./CCTTokenPoolSetup.t.sol";

import {IERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/IERC20.sol";

contract CCTTokenPool_sendReceive is CCTTokenPoolSetup {
  CCTTokenPool internal s_destCCTPool;

  address internal s_sender = makeAddr("sender");
  address internal s_receiver = makeAddr("receiver");

  uint256 internal constant SEND_AMOUNT = 500e18;

  function setUp() public virtual override {
    super.setUp();

    s_destCCTPool = new CCTTokenPool(
      BaseERC20.ConstructorParams({
        name: "CCT Token",
        symbol: "CCT",
        decimals: DEFAULT_TOKEN_DECIMALS,
        maxSupply: MAX_SUPPLY,
        preMint: 0,
        ccipAdmin: OWNER
      }),
      address(0),
      address(s_mockRMNRemote),
      address(s_sourceRouter)
    );
    _applyChainUpdates(address(s_destCCTPool));

    deal(address(s_cctPool), s_sender, SEND_AMOUNT);
  }

  function test_sendReceive_FullLifecycle() public {
    vm.stopPrank();

    uint256 senderBalanceBefore = IERC20(address(s_cctPool)).balanceOf(s_sender);
    uint256 sourceSupplyBefore = IERC20(address(s_cctPool)).totalSupply();
    uint256 destSupplyBefore = IERC20(address(s_destCCTPool)).totalSupply();

    assertEq(SEND_AMOUNT, senderBalanceBefore);
    assertEq(0, destSupplyBefore);

    // --- Source chain: simulate the Router transferring tokens to the pool, then OnRamp calling lockOrBurn ---

    // The Router calls token.transferFrom(sender, pool, amount).
    // For CCTTokenPool the token IS the pool.
    vm.prank(s_sender);
    IERC20(address(s_cctPool)).approve(address(s_sourceRouter), SEND_AMOUNT);

    vm.prank(address(s_sourceRouter));
    IERC20(address(s_cctPool)).transferFrom(s_sender, address(s_cctPool), SEND_AMOUNT);

    assertEq(0, IERC20(address(s_cctPool)).balanceOf(s_sender));
    assertEq(SEND_AMOUNT, IERC20(address(s_cctPool)).balanceOf(address(s_cctPool)));

    // The OnRamp calls pool.lockOrBurn(), which burns from address(this) (the pool).
    vm.prank(s_allowedOnRamp);
    Pool.LockOrBurnOutV1 memory burnOut = s_cctPool.lockOrBurn(
      Pool.LockOrBurnInV1({
        originalSender: s_sender,
        receiver: abi.encode(s_receiver),
        amount: SEND_AMOUNT,
        remoteChainSelector: DEST_CHAIN_SELECTOR,
        localToken: address(s_cctPool)
      })
    );

    assertEq(0, IERC20(address(s_cctPool)).balanceOf(address(s_cctPool)));
    assertEq(sourceSupplyBefore - SEND_AMOUNT, IERC20(address(s_cctPool)).totalSupply());

    // --- Destination chain: OffRamp calls releaseOrMint on the dest pool ---

    vm.prank(s_allowedOffRamp);
    Pool.ReleaseOrMintOutV1 memory mintOut = s_destCCTPool.releaseOrMint(
      Pool.ReleaseOrMintInV1({
        originalSender: abi.encode(s_sender),
        receiver: s_receiver,
        sourceDenominatedAmount: SEND_AMOUNT,
        localToken: address(s_destCCTPool),
        remoteChainSelector: DEST_CHAIN_SELECTOR,
        sourcePoolAddress: abi.encode(s_initialRemotePool),
        sourcePoolData: burnOut.destPoolData,
        offchainTokenData: ""
      })
    );

    assertEq(SEND_AMOUNT, mintOut.destinationAmount);
    assertEq(SEND_AMOUNT, IERC20(address(s_destCCTPool)).balanceOf(s_receiver));
    assertEq(SEND_AMOUNT, IERC20(address(s_destCCTPool)).totalSupply());
  }

  function test_sendReceive_TransferToPoolAllowed() public {
    vm.startPrank(s_sender);
    IERC20(address(s_cctPool)).transfer(address(s_cctPool), SEND_AMOUNT);

    assertEq(0, IERC20(address(s_cctPool)).balanceOf(s_sender));
    assertEq(SEND_AMOUNT, IERC20(address(s_cctPool)).balanceOf(address(s_cctPool)));
  }

  function test_sendReceive_MintToPoolAllowed() public {
    vm.startPrank(s_allowedOffRamp);

    s_destCCTPool.releaseOrMint(
      Pool.ReleaseOrMintInV1({
        originalSender: abi.encode(s_sender),
        receiver: address(s_destCCTPool),
        sourceDenominatedAmount: 1e18,
        localToken: address(s_destCCTPool),
        remoteChainSelector: DEST_CHAIN_SELECTOR,
        sourcePoolAddress: abi.encode(s_initialRemotePool),
        sourcePoolData: abi.encode(DEFAULT_TOKEN_DECIMALS),
        offchainTokenData: ""
      })
    );
  }
}
