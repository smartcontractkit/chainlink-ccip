// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IBridgeV2} from "../../../interfaces/lombard/IBridgeV2.sol";

import {Pool} from "../../../libraries/Pool.sol";
import {LombardTokenPool} from "../../../pools/Lombard/LombardTokenPool.sol";
import {TokenPool} from "../../../pools/TokenPool.sol";
import {LombardTokenPoolHelper} from "../../helpers/LombardTokenPoolHelper.sol";

import {MockLombardAdapter} from "../../mocks/MockLombardAdapter.sol";
import {LombardTokenPoolSetup} from "./LombardTokenPoolSetup.t.sol";

contract LombardTokenPool_lockOrBurn is LombardTokenPoolSetup {
  bytes32 internal constant L_CHAIN_ID = bytes32("LCHAIN");

  function setUp() public virtual override {
    super.setUp();
    vm.startPrank(s_allowedOnRamp);
  }

  function test_lockOrBurn_V2_ForwardsToVerifier() public {
    uint256 amount = 1e18;
    deal(address(s_token), address(s_pool), amount);

    (Pool.LockOrBurnOutV1 memory out, uint256 destAmount) = s_pool.lockOrBurn(
      Pool.LockOrBurnInV1({
        receiver: abi.encodePacked(s_receiver),
        remoteChainSelector: DEST_CHAIN_SELECTOR,
        originalSender: OWNER,
        amount: amount,
        localToken: address(s_token)
      }),
      0,
      ""
    );

    assertEq(destAmount, amount);
    assertEq(out.destTokenAddress, abi.encode(s_remoteToken));
    assertEq(out.destPoolData, abi.encode(uint8(DEFAULT_TOKEN_DECIMALS)));
    assertEq(s_token.balanceOf(s_verifierResolver.getOutboundImplementation(DEST_CHAIN_SELECTOR, "")), amount);
    assertEq(s_token.balanceOf(address(s_pool)), 0);
  }

  function test_lockOrBurn_V1() public {
    uint256 amount = 1e18;
    deal(address(s_token), address(s_pool), amount);

    _configurePathAndBridgeRemoteToken(bytes32(uint256(uint160(s_remoteToken))));

    vm.expectCall(
      address(s_bridge),
      abi.encodeCall(
        IBridgeV2.deposit,
        (
          L_CHAIN_ID,
          address(s_token),
          OWNER,
          bytes32(uint256(uint160(s_receiver))),
          amount,
          bytes32(uint256(uint160(s_remotePool)))
        )
      )
    );

    vm.expectEmit();
    emit TokenPool.LockedOrBurned({
      remoteChainSelector: DEST_CHAIN_SELECTOR, token: address(s_token), sender: OWNER, amount: amount
    });

    Pool.LockOrBurnOutV1 memory out = s_pool.lockOrBurn(
      Pool.LockOrBurnInV1({
        receiver: abi.encode(s_receiver),
        remoteChainSelector: DEST_CHAIN_SELECTOR,
        originalSender: OWNER,
        amount: amount,
        localToken: address(s_token)
      })
    );

    assertEq(out.destTokenAddress, abi.encode(s_remoteToken));
    assertEq(s_token.balanceOf(address(s_bridge)), amount);
  }

  function test_lockOrBurn_V1_UsesAdapterWhenConfigured() public {
    address tokenAdapter = address(new MockLombardAdapter(address(s_bridge), address(s_token)));
    uint256 amount = 1e18;

    changePrank(OWNER);
    LombardTokenPoolHelper adapterPool = new LombardTokenPoolHelper(
      s_token,
      address(s_verifierResolver),
      s_bridge,
      tokenAdapter,
      address(0),
      address(s_mockRMNRemote),
      address(s_sourceRouter),
      DEFAULT_TOKEN_DECIMALS,
      s_feeAggregator
    );
    _applyChainUpdates(address(adapterPool));

    bytes32 destToken = bytes32(uint256(uint160(s_initialRemoteToken)));
    adapterPool.setPath(DEST_CHAIN_SELECTOR, L_CHAIN_ID, abi.encode(s_initialRemotePool));
    s_bridge.setAllowedDestinationToken(L_CHAIN_ID, tokenAdapter, destToken);
    changePrank(s_allowedOnRamp);

    deal(address(s_token), address(adapterPool), amount);

    vm.expectCall(
      address(s_bridge),
      abi.encodeCall(
        IBridgeV2.deposit,
        (
          L_CHAIN_ID,
          tokenAdapter,
          OWNER,
          bytes32(uint256(uint160(s_adapterReceiver))),
          amount,
          bytes32(uint256(uint160(s_initialRemotePool)))
        )
      )
    );

    vm.expectEmit();
    emit TokenPool.LockedOrBurned({
      remoteChainSelector: DEST_CHAIN_SELECTOR, token: address(s_token), sender: OWNER, amount: amount
    });

    Pool.LockOrBurnOutV1 memory out = adapterPool.lockOrBurn(
      Pool.LockOrBurnInV1({
        receiver: abi.encode(s_adapterReceiver),
        remoteChainSelector: DEST_CHAIN_SELECTOR,
        originalSender: OWNER,
        amount: amount,
        localToken: address(s_token)
      })
    );

    assertEq(out.destTokenAddress, abi.encode(s_initialRemoteToken));
    assertEq(s_token.balanceOf(address(s_bridge)), amount);
  }

  function test_lockOrBurn_V2_RevertWhen_OutboundImplementationNotFoundForVerifier() public {
    uint256 amount = 1e18;
    deal(address(s_token), address(s_pool), amount);
    vm.mockCall(
      address(s_verifierResolver),
      abi.encodeCall(s_verifierResolver.getOutboundImplementation, (DEST_CHAIN_SELECTOR, "")),
      abi.encode(address(0))
    );

    vm.expectRevert(LombardTokenPool.OutboundImplementationNotFoundForVerifier.selector);
    s_pool.lockOrBurn(
      Pool.LockOrBurnInV1({
        receiver: abi.encodePacked(s_receiver),
        remoteChainSelector: DEST_CHAIN_SELECTOR,
        originalSender: OWNER,
        amount: amount,
        localToken: address(s_token)
      }),
      0,
      ""
    );
  }

  function test_lockOrBurn_V1_RevertWhen_PathNotExist() public {
    vm.expectRevert(abi.encodeWithSelector(LombardTokenPool.PathNotExist.selector, DEST_CHAIN_SELECTOR));
    s_pool.lockOrBurn(
      Pool.LockOrBurnInV1({
        receiver: abi.encodePacked(s_releaseRecipient),
        remoteChainSelector: DEST_CHAIN_SELECTOR,
        originalSender: OWNER,
        amount: 1e18,
        localToken: address(s_token)
      })
    );
  }

  function test_lockOrBurn_V1_RevertWhen_InvalidReceiver() public {
    _configurePathAndBridgeRemoteToken(bytes32(uint256(uint160(s_remoteToken))));

    vm.expectRevert(abi.encodeWithSelector(LombardTokenPool.InvalidReceiver.selector, hex"1234"));
    s_pool.lockOrBurn(
      Pool.LockOrBurnInV1({
        receiver: hex"1234",
        remoteChainSelector: DEST_CHAIN_SELECTOR,
        originalSender: OWNER,
        amount: 1e18,
        localToken: address(s_token)
      })
    );
  }

  function test_lockOrBurn_V1_RevertWhen_RemoteTokenMismatch() public {
    _configurePathAndBridgeRemoteToken(bytes32("differentToken"));

    vm.expectRevert(
      abi.encodeWithSelector(
        LombardTokenPool.RemoteTokenMismatch.selector,
        bytes32("differentToken"),
        bytes32(uint256(uint160(s_remoteToken)))
      )
    );
    s_pool.lockOrBurn(
      Pool.LockOrBurnInV1({
        receiver: abi.encode(s_receiver),
        remoteChainSelector: DEST_CHAIN_SELECTOR,
        originalSender: OWNER,
        amount: 1e18,
        localToken: address(s_token)
      })
    );
  }

  function _configurePathAndBridgeRemoteToken(
    bytes32 remoteTokenId
  ) internal {
    changePrank(OWNER);
    s_pool.setPath(DEST_CHAIN_SELECTOR, L_CHAIN_ID, abi.encode(s_remotePool));
    s_bridge.setAllowedDestinationToken(L_CHAIN_ID, address(s_token), remoteTokenId);
    changePrank(s_allowedOnRamp);
  }
}
