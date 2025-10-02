// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolV1} from "../../../../interfaces/IPool.sol";

import {Pool} from "../../../../libraries/Pool.sol";
import {TokenPool} from "../../../../pools/TokenPool.sol";
import {USDCTokenPool} from "../../../../pools/USDC/USDCTokenPool.sol";
import {USDCTokenPoolProxy} from "../../../../pools/USDC/USDCTokenPoolProxy.sol";
import {USDCTokenPoolProxySetup} from "./USDCTokenPoolProxySetup.t.sol";

import {IERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/IERC20.sol";
import {IERC165} from "@openzeppelin/contracts@5.0.2/utils/introspection/IERC165.sol";

contract USDCTokenPoolProxy_lockOrBurn is USDCTokenPoolProxySetup {
  address internal s_sender = makeAddr("sender");
  address internal s_receiver = makeAddr("receiver");
  bytes internal s_destPoolData = abi.encode(1, 2, 3);

  function setUp() public virtual override {
    super.setUp();

    // Configure lock or burn mechanisms for different chains
    uint64[] memory chainSelectors = new uint64[](3);
    chainSelectors[0] = SOURCE_CHAIN_SELECTOR;
    chainSelectors[1] = DEST_CHAIN_SELECTOR;
    chainSelectors[2] = 12345; // Another test chain

    USDCTokenPoolProxy.LockOrBurnMechanism[] memory mechanisms = new USDCTokenPoolProxy.LockOrBurnMechanism[](3);
    mechanisms[0] = USDCTokenPoolProxy.LockOrBurnMechanism.CCTP_V1;
    mechanisms[1] = USDCTokenPoolProxy.LockOrBurnMechanism.CCTP_V2;
    mechanisms[2] = USDCTokenPoolProxy.LockOrBurnMechanism.LOCK_RELEASE;

    s_usdcTokenPoolProxy.updateLockOrBurnMechanisms(chainSelectors, mechanisms);
  }

  function test_lockOrBurn_CCTPV1() public {
    uint256 amount = 100;
    bytes memory destTokenAddress = abi.encode(address(s_USDCToken));

    // Set the DEST_CHAIN_SELECTOR to use CCTP V1 using the update function
    uint64[] memory selectors = new uint64[](1);
    selectors[0] = DEST_CHAIN_SELECTOR;
    USDCTokenPoolProxy.LockOrBurnMechanism[] memory mechs = new USDCTokenPoolProxy.LockOrBurnMechanism[](1);
    mechs[0] = USDCTokenPoolProxy.LockOrBurnMechanism.CCTP_V1;
    s_usdcTokenPoolProxy.updateLockOrBurnMechanisms(selectors, mechs);

    Pool.LockOrBurnInV1 memory lockOrBurnIn = Pool.LockOrBurnInV1({
      receiver: abi.encode(s_receiver),
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      originalSender: s_sender,
      amount: amount,
      localToken: address(s_USDCToken)
    });

    // Mock the CCTP V1 pool's lockOrBurn to return expected output
    Pool.LockOrBurnOutV1 memory expectedOutput =
      Pool.LockOrBurnOutV1({destTokenAddress: destTokenAddress, destPoolData: s_destPoolData});

    vm.mockCall(
      address(s_cctpV1Pool),
      abi.encodeWithSelector(USDCTokenPool.lockOrBurn.selector, lockOrBurnIn),
      abi.encode(expectedOutput)
    );

    vm.mockCall(
      address(s_cctpV1Pool),
      abi.encodeWithSelector(IERC165.supportsInterface.selector, type(IPoolV1).interfaceId),
      abi.encode(true)
    );

    vm.startPrank(s_routerAllowedOnRamp);

    vm.expectCall(address(s_cctpV1Pool), abi.encodeWithSelector(USDCTokenPool.lockOrBurn.selector, lockOrBurnIn));

    vm.expectCall(address(s_USDCToken), abi.encodeWithSelector(IERC20.transfer.selector, address(s_cctpV1Pool), amount));

    Pool.LockOrBurnOutV1 memory result = s_usdcTokenPoolProxy.lockOrBurn(lockOrBurnIn);

    assertEq(result.destTokenAddress, expectedOutput.destTokenAddress);
    assertEq(result.destPoolData, expectedOutput.destPoolData);
  }

  function test_lockOrBurn_CCTPV2() public {
    // Arrange: Define test constants
    uint256 amount = 200;
    bytes memory destTokenAddress = abi.encode(address(s_USDCToken));

    Pool.LockOrBurnInV1 memory lockOrBurnIn = Pool.LockOrBurnInV1({
      receiver: abi.encode(s_receiver),
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      originalSender: s_sender,
      amount: amount,
      localToken: address(s_USDCToken)
    });

    vm.mockCall(
      address(s_cctpV2Pool),
      abi.encodeWithSelector(TokenPool.lockOrBurn.selector, lockOrBurnIn),
      abi.encode(Pool.LockOrBurnOutV1({destTokenAddress: destTokenAddress, destPoolData: s_destPoolData}))
    );

    // Mock the CCTP V2 pool's lockOrBurn to return expected output
    Pool.LockOrBurnOutV1 memory expectedOutput =
      Pool.LockOrBurnOutV1({destTokenAddress: destTokenAddress, destPoolData: s_destPoolData});

    vm.expectCall(address(s_cctpV2Pool), abi.encodeWithSelector(TokenPool.lockOrBurn.selector, lockOrBurnIn));

    vm.expectCall(address(s_USDCToken), abi.encodeWithSelector(IERC20.transfer.selector, address(s_cctpV2Pool), amount));

    vm.startPrank(s_routerAllowedOnRamp);
    Pool.LockOrBurnOutV1 memory result = s_usdcTokenPoolProxy.lockOrBurn(lockOrBurnIn);
    assertEq(result.destTokenAddress, expectedOutput.destTokenAddress);
    assertEq(result.destPoolData, expectedOutput.destPoolData);
  }

  function test_lockOrBurn_LockRelease() public {
    uint64 testChainSelector = 12345;
    uint256 amount = 300;
    bytes memory destTokenAddress = abi.encode(address(s_USDCToken));
    bytes[] memory remotePoolAddresses = new bytes[](1);
    remotePoolAddresses[0] = abi.encode(address(s_lockReleasePool));

    // The update function will check that the pool supports the IPoolV1 interface and IERC165 interface
    // so we need to mock those here.
    vm.mockCall(
      address(s_lockReleasePool),
      abi.encodeWithSelector(IERC165.supportsInterface.selector, type(IPoolV1).interfaceId),
      abi.encode(true)
    );

    vm.mockCall(
      address(s_lockReleasePool),
      abi.encodeWithSelector(IERC165.supportsInterface.selector, type(IERC165).interfaceId),
      abi.encode(true)
    );

    // Set the the s_lockRelease pool for the LockRelease mechanism
    uint64[] memory selectors = new uint64[](1);
    selectors[0] = testChainSelector;
    address[] memory lockReleasePools = new address[](1);
    lockReleasePools[0] = address(s_lockReleasePool);
    s_usdcTokenPoolProxy.updateLockReleasePoolAddresses(selectors, lockReleasePools);

    vm.mockCall(
      address(s_router),
      abi.encodeWithSelector(bytes4(keccak256("getOnRamp(uint64)")), uint64(testChainSelector)),
      abi.encode(s_routerAllowedOnRamp)
    );

    Pool.LockOrBurnInV1 memory lockOrBurnIn = Pool.LockOrBurnInV1({
      receiver: abi.encode(s_receiver),
      remoteChainSelector: testChainSelector,
      originalSender: s_sender,
      amount: amount,
      localToken: address(s_USDCToken)
    });

    // Mock the lock release pool's lockOrBurn to return expected output
    Pool.LockOrBurnOutV1 memory expectedOutput =
      Pool.LockOrBurnOutV1({destTokenAddress: destTokenAddress, destPoolData: s_destPoolData});

    vm.mockCall(
      address(s_lockReleasePool),
      abi.encodeWithSelector(USDCTokenPool.lockOrBurn.selector, lockOrBurnIn),
      abi.encode(expectedOutput)
    );

    vm.startPrank(s_routerAllowedOnRamp);

    Pool.LockOrBurnOutV1 memory result = s_usdcTokenPoolProxy.lockOrBurn(lockOrBurnIn);
    assertEq(result.destTokenAddress, expectedOutput.destTokenAddress);
    assertEq(result.destPoolData, expectedOutput.destPoolData);
  }

  // Reverts

  function test_lockOrBurn_RevertWhen_InvalidLockOrBurnMechanism() public {
    uint64 testChainSelector = 99999;
    uint256 amount = 100;
    bytes[] memory remotePoolAddresses = new bytes[](1);
    remotePoolAddresses[0] = abi.encode(address(s_lockReleasePool));

    vm.mockCall(
      address(s_router),
      abi.encodeWithSelector(bytes4(keccak256("getOnRamp(uint64)")), uint64(testChainSelector)),
      abi.encode(s_routerAllowedOnRamp)
    );

    assertFalse(s_usdcTokenPoolProxy.isSupportedChain(testChainSelector));

    vm.expectRevert(
      abi.encodeWithSelector(
        USDCTokenPoolProxy.InvalidLockOrBurnMechanism.selector, USDCTokenPoolProxy.LockOrBurnMechanism(0)
      )
    );

    vm.startPrank(s_routerAllowedOnRamp);
    s_usdcTokenPoolProxy.lockOrBurn(
      Pool.LockOrBurnInV1({
        receiver: abi.encode(s_receiver),
        remoteChainSelector: testChainSelector, // Chain with no configured mechanism
        originalSender: s_sender,
        amount: amount,
        localToken: address(s_USDCToken)
      })
    );
  }

  function test_lockOrBurn_RevertWhen_NoLockOrBurnMechanismSet() public {
    // Configure lock or burn mechanisms for different chains but do not set the lock release pool for the chain
    uint64[] memory chainSelectors = new uint64[](1);
    chainSelectors[0] = DEST_CHAIN_SELECTOR;

    USDCTokenPoolProxy.LockOrBurnMechanism[] memory mechanisms = new USDCTokenPoolProxy.LockOrBurnMechanism[](1);
    mechanisms[0] = USDCTokenPoolProxy.LockOrBurnMechanism.LOCK_RELEASE;

    vm.startPrank(OWNER);
    s_usdcTokenPoolProxy.updateLockOrBurnMechanisms(chainSelectors, mechanisms);

    vm.mockCall(
      address(s_router),
      abi.encodeWithSelector(bytes4(keccak256("getOnRamp(uint64)")), uint64(DEST_CHAIN_SELECTOR)),
      abi.encode(s_routerAllowedOnRamp)
    );

    Pool.LockOrBurnInV1 memory lockOrBurnIn = Pool.LockOrBurnInV1({
      receiver: abi.encode(s_receiver),
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      originalSender: s_sender,
      amount: 100,
      localToken: address(s_USDCToken)
    });

    vm.startPrank(s_routerAllowedOnRamp);

    vm.expectRevert(abi.encodeWithSelector(USDCTokenPoolProxy.NoLockOrBurnMechanismSet.selector, DEST_CHAIN_SELECTOR));
    s_usdcTokenPoolProxy.lockOrBurn(lockOrBurnIn);
  }

  function test_lockOrBurn_RevertWhen_Unauthorized() public {
    address unauthorized = makeAddr("unauthorized");

    vm.startPrank(unauthorized);

    Pool.LockOrBurnInV1 memory lockOrBurnIn = Pool.LockOrBurnInV1({
      receiver: abi.encode(s_receiver),
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      originalSender: s_sender,
      amount: 100,
      localToken: address(s_USDCToken)
    });

    vm.expectRevert(abi.encodeWithSelector(USDCTokenPoolProxy.CallerIsNotARampOnRouter.selector, unauthorized));
    s_usdcTokenPoolProxy.lockOrBurn(lockOrBurnIn);
  }
}
