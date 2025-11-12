// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolV1} from "../../../interfaces/IPool.sol";
import {IPoolV2} from "../../../interfaces/IPoolV2.sol";
import {ITokenAdminRegistry} from "../../../interfaces/ITokenAdminRegistry.sol";

import {Client} from "../../../libraries/Client.sol";
import {MessageV1Codec} from "../../../libraries/MessageV1Codec.sol";
import {Pool} from "../../../libraries/Pool.sol";
import {OffRamp} from "../../../offRamp/OffRamp.sol";
import {BurnMintTokenPool} from "../../../pools/BurnMintTokenPool.sol";
import {OffRampHelper} from "../../helpers/OffRampHelper.sol";
import {TokenPoolSetup} from "../../pools/TokenPool/TokenPoolSetup.t.sol";

contract OffRamp_releaseOrMintSingleToken is TokenPoolSetup {
  BurnMintTokenPool internal s_pool;
  OffRampHelper internal s_offRamp;
  address internal s_tokenAdminRegistry = makeAddr("tokenAdminRegistry");
  address internal s_receiver = makeAddr("receiver");

  function setUp() public override {
    super.setUp();

    s_pool = new BurnMintTokenPool(
      s_token, DEFAULT_TOKEN_DECIMALS, new address[](0), address(s_mockRMNRemote), address(s_sourceRouter)
    );
    s_token.grantMintAndBurnRoles(address(s_pool));
    _applyChainUpdates(address(s_pool));

    s_offRamp = new OffRampHelper(
      OffRamp.StaticConfig({
        localChainSelector: SOURCE_CHAIN_SELECTOR,
        gasForCallExactCheck: GAS_FOR_CALL_EXACT_CHECK,
        rmnRemote: s_mockRMNRemote,
        tokenAdminRegistry: s_tokenAdminRegistry
      })
    );

    vm.mockCall(
      address(s_sourceRouter),
      abi.encodeCall(s_sourceRouter.isOffRamp, (DEST_CHAIN_SELECTOR, address(s_offRamp))),
      abi.encode(true)
    );
    vm.mockCall(
      address(s_tokenAdminRegistry),
      abi.encodeCall(ITokenAdminRegistry.getPool, address(s_token)),
      abi.encode(address(s_pool))
    );
  }

  function test_releaseOrMintSingleToken_CallsV2Function() public {
    Pool.ReleaseOrMintInV1 memory expectedInput = _buildReleaseInput();
    MessageV1Codec.TokenTransferV1 memory tokenTransfer = _buildTokenTransfer();
    uint16 finality = 2;

    vm.expectCall(address(s_pool), abi.encodeCall(IPoolV2.releaseOrMint, (expectedInput, finality)));

    Client.EVMTokenAmount memory dest =
      s_offRamp.releaseOrMintSingleToken(tokenTransfer, expectedInput.originalSender, DEST_CHAIN_SELECTOR, finality);

    assertEq(dest.token, address(s_token));
    assertEq(dest.amount, tokenTransfer.amount);
    assertEq(s_token.balanceOf(s_receiver), dest.amount);
  }

  function test_releaseOrMintSingleToken_CallsV1FunctionWhenNoV2Support() public {
    vm.mockCall(
      address(s_pool), abi.encodeCall(s_pool.supportsInterface, (type(IPoolV2).interfaceId)), abi.encode(false)
    );

    Pool.ReleaseOrMintInV1 memory expectedInput = _buildReleaseInput();
    MessageV1Codec.TokenTransferV1 memory tokenTransfer = _buildTokenTransfer();

    vm.expectCall(address(s_pool), abi.encodeCall(IPoolV1.releaseOrMint, (expectedInput)));

    Client.EVMTokenAmount memory dest =
      s_offRamp.releaseOrMintSingleToken(tokenTransfer, expectedInput.originalSender, DEST_CHAIN_SELECTOR, 0);

    assertEq(dest.token, address(s_token));
    assertEq(dest.amount, tokenTransfer.amount);
    assertEq(s_token.balanceOf(s_receiver), dest.amount);
  }

  function test_releaseOrMintSingleToken_PropagatesPoolError() public {
    Pool.ReleaseOrMintInV1 memory expectedInput = _buildReleaseInput();
    MessageV1Codec.TokenTransferV1 memory tokenTransfer = _buildTokenTransfer();

    bytes memory callData = abi.encodeWithSelector(IPoolV2.releaseOrMint.selector, expectedInput, uint16(2));
    vm.expectCall(address(s_pool), callData);
    bytes memory poolRevertData = abi.encode("pool-error");
    vm.mockCallRevert(address(s_pool), callData, poolRevertData);

    vm.expectRevert(abi.encodeWithSelector(OffRamp.TokenHandlingError.selector, address(s_token), poolRevertData));
    s_offRamp.releaseOrMintSingleToken(tokenTransfer, expectedInput.originalSender, DEST_CHAIN_SELECTOR, 2);
  }

  function test_releaseOrMintSingleToken_RevertWhen_NotACompatiblePool_PoolAddressZero() public {
    vm.mockCall(
      address(s_tokenAdminRegistry),
      abi.encodeCall(ITokenAdminRegistry.getPool, (address(s_token))),
      abi.encode(address(0))
    );

    MessageV1Codec.TokenTransferV1 memory tokenTransfer = _buildTokenTransfer();

    vm.expectRevert(abi.encodeWithSelector(OffRamp.NotACompatiblePool.selector, address(0)));
    s_offRamp.releaseOrMintSingleToken(tokenTransfer, abi.encodePacked(address(1)), DEST_CHAIN_SELECTOR, 0);
  }

  function test_releaseOrMintSingleToken_RevertWhen_NotACompatiblePool_UnsupportedInterface() public {
    vm.mockCall(address(s_pool), abi.encodeCall(s_pool.supportsInterface, (Pool.CCIP_POOL_V1)), abi.encode(false));
    vm.mockCall(
      address(s_pool), abi.encodeCall(s_pool.supportsInterface, (type(IPoolV2).interfaceId)), abi.encode(false)
    );

    MessageV1Codec.TokenTransferV1 memory tokenTransfer = _buildTokenTransfer();

    vm.expectRevert(abi.encodeWithSelector(OffRamp.NotACompatiblePool.selector, address(s_pool)));
    s_offRamp.releaseOrMintSingleToken(tokenTransfer, abi.encodePacked(address(1)), DEST_CHAIN_SELECTOR, 0);
  }

  function test_releaseOrMintSingleToken_RevertsWhen_ReleaseOrMintBalanceMismatch() public {
    vm.mockCall(
      address(s_pool), abi.encodeCall(s_pool.supportsInterface, (type(IPoolV2).interfaceId)), abi.encode(false)
    );
    Pool.ReleaseOrMintInV1 memory expectedInput = _buildReleaseInput();
    MessageV1Codec.TokenTransferV1 memory tokenTransfer = _buildTokenTransfer();

    vm.expectCall(address(s_pool), abi.encodeCall(IPoolV1.releaseOrMint, (expectedInput)));
    vm.mockCall(address(s_token), abi.encodeCall(s_token.balanceOf, (expectedInput.receiver)), abi.encode(0));

    vm.expectRevert(abi.encodeWithSelector(OffRamp.ReleaseOrMintBalanceMismatch.selector, tokenTransfer.amount, 0, 0));
    s_offRamp.releaseOrMintSingleToken(tokenTransfer, expectedInput.originalSender, DEST_CHAIN_SELECTOR, 1);
  }

  function _buildReleaseInput() internal returns (Pool.ReleaseOrMintInV1 memory) {
    return Pool.ReleaseOrMintInV1({
      originalSender: abi.encodePacked(makeAddr("originalSender")),
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      receiver: s_receiver,
      sourceDenominatedAmount: 1 ether,
      localToken: address(s_token),
      sourcePoolAddress: abi.encode(s_initialRemotePool),
      sourcePoolData: abi.encode(18),
      offchainTokenData: ""
    });
  }

  function _buildTokenTransfer() internal view returns (MessageV1Codec.TokenTransferV1 memory) {
    return MessageV1Codec.TokenTransferV1({
      amount: 1 ether,
      sourcePoolAddress: abi.encodePacked(s_initialRemotePool),
      sourceTokenAddress: abi.encodePacked(s_initialRemoteToken),
      destTokenAddress: abi.encodePacked(address(s_token)),
      tokenReceiver: abi.encodePacked(s_receiver),
      extraData: abi.encode(18)
    });
  }
}
