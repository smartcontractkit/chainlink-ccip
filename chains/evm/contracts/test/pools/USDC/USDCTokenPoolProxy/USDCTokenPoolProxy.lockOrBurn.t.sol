// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ICrossChainVerifierResolver} from "../../../../interfaces/ICrossChainVerifierResolver.sol";
import {IPoolV1} from "../../../../interfaces/IPool.sol";
import {IPoolV2} from "../../../../interfaces/IPoolV2.sol";

import {Router} from "../../../../Router.sol";
import {Pool} from "../../../../libraries/Pool.sol";
import {USDCTokenPoolProxy} from "../../../../pools/USDC/USDCTokenPoolProxy.sol";
import {USDCTokenPoolProxySetup} from "./USDCTokenPoolProxySetup.t.sol";

import {IERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/IERC20.sol";
import {IERC165} from "@openzeppelin/contracts@5.3.0/utils/introspection/IERC165.sol";

contract USDCTokenPoolProxy_lockOrBurn is USDCTokenPoolProxySetup {
  address internal s_sender = makeAddr("sender");
  address internal s_receiver = makeAddr("receiver");
  bytes internal s_destPoolData = abi.encode(1, 2, 3);
  uint64 internal s_chainSelForV1 = SOURCE_CHAIN_SELECTOR;
  uint64 internal s_chainSelForV2 = DEST_CHAIN_SELECTOR;
  uint64 internal s_chainSelFoLockRelease = 12345;
  uint64 internal s_chainSelForCCV = 67890;

  function setUp() public virtual override {
    super.setUp();

    // Set the OnRamp on the router for each of the chain selectors.
    Router.OnRamp[] memory onRampUpdates = new Router.OnRamp[](4);
    onRampUpdates[0] = Router.OnRamp({destChainSelector: s_chainSelForV1, onRamp: s_routerAllowedOnRamp});
    onRampUpdates[1] = Router.OnRamp({destChainSelector: s_chainSelForV2, onRamp: s_routerAllowedOnRamp});
    onRampUpdates[2] = Router.OnRamp({destChainSelector: s_chainSelFoLockRelease, onRamp: s_routerAllowedOnRamp});
    onRampUpdates[3] = Router.OnRamp({destChainSelector: s_chainSelForCCV, onRamp: s_routerAllowedOnRamp});
    s_router.applyRampUpdates(onRampUpdates, new Router.OffRamp[](0), new Router.OffRamp[](0));

    // Enable ERC165 for the lock release pool.
    _enableERC165InterfaceChecks(s_lockReleasePool, type(IPoolV1).interfaceId);

    // Update pool addresses to include lock release pool.
    changePrank(OWNER);
    s_usdcTokenPoolProxy.updatePoolAddresses(
      USDCTokenPoolProxy.PoolAddresses({
        cctpV1Pool: s_cctpV1Pool,
        cctpV2Pool: s_cctpV2Pool,
        cctpV2PoolWithCCV: s_cctpThroughCCVTokenPool,
        siloedLockReleasePool: s_lockReleasePool
      })
    );

    // Configure lock or burn mechanisms for different chains.
    uint64[] memory chainSelectors = new uint64[](4);
    chainSelectors[0] = s_chainSelForV1;
    chainSelectors[1] = s_chainSelForV2;
    chainSelectors[2] = s_chainSelFoLockRelease;
    chainSelectors[3] = s_chainSelForCCV;

    USDCTokenPoolProxy.LockOrBurnMechanism[] memory mechanisms = new USDCTokenPoolProxy.LockOrBurnMechanism[](4);
    mechanisms[0] = USDCTokenPoolProxy.LockOrBurnMechanism.CCTP_V1;
    mechanisms[1] = USDCTokenPoolProxy.LockOrBurnMechanism.CCTP_V2;
    mechanisms[2] = USDCTokenPoolProxy.LockOrBurnMechanism.LOCK_RELEASE;
    mechanisms[3] = USDCTokenPoolProxy.LockOrBurnMechanism.CCV;

    s_usdcTokenPoolProxy.updateLockOrBurnMechanisms(chainSelectors, mechanisms);
  }

  function test_lockOrBurn_CCTPV1() public {
    uint256 amount = 100;
    bytes memory destTokenAddress = abi.encode(address(s_USDCToken));

    uint64[] memory selectors = new uint64[](1);
    selectors[0] = s_chainSelForV1;
    USDCTokenPoolProxy.LockOrBurnMechanism[] memory mechs = new USDCTokenPoolProxy.LockOrBurnMechanism[](1);
    mechs[0] = USDCTokenPoolProxy.LockOrBurnMechanism.CCTP_V1;
    s_usdcTokenPoolProxy.updateLockOrBurnMechanisms(selectors, mechs);

    Pool.LockOrBurnInV1 memory lockOrBurnIn = Pool.LockOrBurnInV1({
      receiver: abi.encode(s_receiver),
      remoteChainSelector: s_chainSelForV1,
      originalSender: s_sender,
      amount: amount,
      localToken: address(s_USDCToken)
    });

    // Mock the CCTP V1 pool's lockOrBurn to return expected output.
    Pool.LockOrBurnOutV1 memory expectedOutput =
      Pool.LockOrBurnOutV1({destTokenAddress: destTokenAddress, destPoolData: s_destPoolData});

    vm.mockCall(
      address(s_cctpV1Pool),
      abi.encodeWithSelector(IPoolV1.lockOrBurn.selector, lockOrBurnIn),
      abi.encode(expectedOutput)
    );

    vm.mockCall(
      address(s_cctpV1Pool),
      abi.encodeWithSelector(IERC165.supportsInterface.selector, type(IPoolV1).interfaceId),
      abi.encode(true)
    );

    vm.startPrank(s_routerAllowedOnRamp);

    vm.expectCall(address(s_cctpV1Pool), abi.encodeWithSelector(IPoolV1.lockOrBurn.selector, lockOrBurnIn));

    vm.expectCall(address(s_USDCToken), abi.encodeWithSelector(IERC20.transfer.selector, address(s_cctpV1Pool), amount));

    Pool.LockOrBurnOutV1 memory result = s_usdcTokenPoolProxy.lockOrBurn(lockOrBurnIn);

    assertEq(result.destTokenAddress, expectedOutput.destTokenAddress);
    assertEq(result.destPoolData, expectedOutput.destPoolData);
  }

  function test_lockOrBurn_CCTPV2() public {
    uint256 amount = 200;
    bytes memory destTokenAddress = abi.encode(address(s_USDCToken));

    Pool.LockOrBurnInV1 memory lockOrBurnIn = Pool.LockOrBurnInV1({
      receiver: abi.encode(s_receiver),
      remoteChainSelector: s_chainSelForV2,
      originalSender: s_sender,
      amount: amount,
      localToken: address(s_USDCToken)
    });

    vm.mockCall(
      address(s_cctpV2Pool),
      abi.encodeCall(IPoolV1.lockOrBurn, (lockOrBurnIn)),
      abi.encode(Pool.LockOrBurnOutV1({destTokenAddress: destTokenAddress, destPoolData: s_destPoolData}))
    );

    // Mock the CCTP V2 pool's lockOrBurn to return expected output.
    Pool.LockOrBurnOutV1 memory expectedOutput =
      Pool.LockOrBurnOutV1({destTokenAddress: destTokenAddress, destPoolData: s_destPoolData});

    vm.expectCall(address(s_cctpV2Pool), abi.encodeCall(IPoolV1.lockOrBurn, (lockOrBurnIn)));

    vm.expectCall(address(s_USDCToken), abi.encodeWithSelector(IERC20.transfer.selector, address(s_cctpV2Pool), amount));

    vm.startPrank(s_routerAllowedOnRamp);
    Pool.LockOrBurnOutV1 memory result = s_usdcTokenPoolProxy.lockOrBurn(lockOrBurnIn);
    assertEq(result.destTokenAddress, expectedOutput.destTokenAddress);
    assertEq(result.destPoolData, expectedOutput.destPoolData);
  }

  function test_lockOrBurn_LockRelease() public {
    uint256 amount = 300;
    bytes memory destTokenAddress = abi.encode(address(s_USDCToken));
    bytes[] memory remotePoolAddresses = new bytes[](1);
    remotePoolAddresses[0] = abi.encode(address(s_lockReleasePool));

    // The update function will check that the pool supports the IPoolV1 interface and IERC165 interface
    // so we need to mock those here.
    _enableERC165InterfaceChecks(address(s_lockReleasePool), type(IPoolV1).interfaceId);

    // Set the siloed pool via updatePoolAddresses - use a clean PoolAddresses struct.
    s_usdcTokenPoolProxy.updatePoolAddresses(
      USDCTokenPoolProxy.PoolAddresses({
        cctpV1Pool: address(0),
        cctpV2Pool: address(0),
        cctpV2PoolWithCCV: address(0),
        siloedLockReleasePool: address(s_lockReleasePool)
      })
    );

    vm.mockCall(
      address(s_router),
      abi.encodeWithSelector(bytes4(keccak256("getOnRamp(uint64)")), uint64(s_chainSelFoLockRelease)),
      abi.encode(s_routerAllowedOnRamp)
    );

    Pool.LockOrBurnInV1 memory lockOrBurnIn = Pool.LockOrBurnInV1({
      receiver: abi.encode(s_receiver),
      remoteChainSelector: s_chainSelFoLockRelease,
      originalSender: s_sender,
      amount: amount,
      localToken: address(s_USDCToken)
    });

    // Mock the lock release pool's lockOrBurn to return expected output.
    Pool.LockOrBurnOutV1 memory expectedOutput =
      Pool.LockOrBurnOutV1({destTokenAddress: destTokenAddress, destPoolData: s_destPoolData});

    vm.mockCall(
      address(s_lockReleasePool),
      abi.encodeWithSelector(IPoolV1.lockOrBurn.selector, lockOrBurnIn),
      abi.encode(expectedOutput)
    );

    vm.startPrank(s_routerAllowedOnRamp);

    vm.expectCall(address(s_lockReleasePool), abi.encodeWithSelector(IPoolV1.lockOrBurn.selector, lockOrBurnIn));
    vm.expectCall(
      address(s_USDCToken), abi.encodeWithSelector(IERC20.transfer.selector, address(s_lockReleasePool), amount)
    );

    Pool.LockOrBurnOutV1 memory result = s_usdcTokenPoolProxy.lockOrBurn(lockOrBurnIn);
    assertEq(result.destTokenAddress, expectedOutput.destTokenAddress);
    assertEq(result.destPoolData, expectedOutput.destPoolData);
  }

  function test_lockOrBurn_LockReleaseV2() public {
    uint256 amount = 500;
    bytes memory tokenArgs = abi.encode(1, 2, 3);
    uint16 blockConfirmationRequested = 2;
    bytes memory destTokenAddress = abi.encode(address(s_USDCToken));

    // Enable IPoolV2 interface for lock release pool.
    _enableERC165InterfaceChecks(address(s_lockReleasePool), type(IPoolV2).interfaceId);

    // Set the siloed pool via updatePoolAddresses.
    s_usdcTokenPoolProxy.updatePoolAddresses(
      USDCTokenPoolProxy.PoolAddresses({
        cctpV1Pool: address(0),
        cctpV2Pool: address(0),
        cctpV2PoolWithCCV: address(0),
        siloedLockReleasePool: address(s_lockReleasePool)
      })
    );

    Pool.LockOrBurnInV1 memory lockOrBurnIn = Pool.LockOrBurnInV1({
      receiver: abi.encode(s_receiver),
      remoteChainSelector: s_chainSelFoLockRelease,
      originalSender: s_sender,
      amount: amount,
      localToken: address(s_USDCToken)
    });

    Pool.LockOrBurnOutV1 memory expectedOutput =
      Pool.LockOrBurnOutV1({destTokenAddress: destTokenAddress, destPoolData: s_destPoolData});

    // Mock the lock release pool's IPoolV2.lockOrBurn.
    vm.mockCall(
      address(s_lockReleasePool),
      abi.encodeCall(IPoolV2.lockOrBurn, (lockOrBurnIn, blockConfirmationRequested, tokenArgs)),
      abi.encode(expectedOutput, amount)
    );

    vm.startPrank(s_routerAllowedOnRamp);

    vm.expectCall(
      address(s_lockReleasePool),
      abi.encodeCall(IPoolV2.lockOrBurn, (lockOrBurnIn, blockConfirmationRequested, tokenArgs))
    );
    vm.expectCall(address(s_USDCToken), abi.encodeCall(IERC20.transfer, (address(s_lockReleasePool), amount)));

    (Pool.LockOrBurnOutV1 memory result, uint256 destTokenAmount) =
      s_usdcTokenPoolProxy.lockOrBurn(lockOrBurnIn, blockConfirmationRequested, tokenArgs);
    assertEq(result.destTokenAddress, expectedOutput.destTokenAddress);
    assertEq(result.destPoolData, expectedOutput.destPoolData);
    assertEq(destTokenAmount, amount);
  }

  function test_lockOrBurn_CCTPV2WithCCV() public {
    uint256 amount = 400;
    bytes memory tokenArgs = abi.encode(abi.encode(1, 2, 3));
    uint16 blockConfirmationRequested = 1;
    bytes memory destTokenAddress = abi.encode(address(s_USDCToken));
    address verifierImpl = makeAddr("verifierImpl");

    Pool.LockOrBurnInV1 memory lockOrBurnIn = Pool.LockOrBurnInV1({
      receiver: abi.encode(s_receiver),
      remoteChainSelector: s_chainSelForCCV,
      originalSender: s_sender,
      amount: amount,
      localToken: address(s_USDCToken)
    });

    Pool.LockOrBurnOutV1 memory expectedOutput =
      Pool.LockOrBurnOutV1({destTokenAddress: destTokenAddress, destPoolData: s_destPoolData});

    vm.mockCall(
      address(s_cctpThroughCCVTokenPool),
      abi.encodeWithSelector(IPoolV2.lockOrBurn.selector, lockOrBurnIn, 1, tokenArgs),
      abi.encode(Pool.LockOrBurnOutV1({destTokenAddress: destTokenAddress, destPoolData: s_destPoolData}), amount)
    );

    vm.startPrank(s_routerAllowedOnRamp);

    vm.mockCall(
      address(s_cctpVerifier),
      abi.encodeWithSelector(
        ICrossChainVerifierResolver.getOutboundImplementation.selector, s_chainSelForCCV, tokenArgs
      ),
      abi.encode(verifierImpl)
    );

    vm.expectCall(
      address(s_cctpThroughCCVTokenPool),
      abi.encodeWithSelector(IPoolV2.lockOrBurn.selector, lockOrBurnIn, blockConfirmationRequested, tokenArgs)
    );
    vm.expectCall(address(s_USDCToken), abi.encodeWithSelector(IERC20.transfer.selector, verifierImpl, amount));

    (Pool.LockOrBurnOutV1 memory result, uint256 destTokenAmount) =
      s_usdcTokenPoolProxy.lockOrBurn(lockOrBurnIn, blockConfirmationRequested, tokenArgs);
    assertEq(result.destTokenAddress, expectedOutput.destTokenAddress);
    assertEq(result.destPoolData, expectedOutput.destPoolData);
    assertEq(destTokenAmount, amount);

    // Ensure that the balance of the verifier impl has been updated.
    assertEq(s_USDCToken.balanceOf(verifierImpl), amount);
  }

  // Reverts

  function test_lockOrBurn_CCTPV2WithCCV_RevertWhen_NoLockOrBurnMechanismSet() public {
    uint256 amount = 100;
    bytes[] memory remotePoolAddresses = new bytes[](1);
    remotePoolAddresses[0] = abi.encode(address(s_lockReleasePool));

    _enableERC165InterfaceChecks(s_cctpThroughCCVTokenPool, type(IPoolV1).interfaceId);
    _enableERC165InterfaceChecks(s_cctpV2Pool, type(IPoolV1).interfaceId);
    _enableERC165InterfaceChecks(s_cctpV1Pool, type(IPoolV1).interfaceId);

    // Remove the cctpV2PoolWithCCV pool from stored pools.
    s_usdcTokenPoolProxy.updatePoolAddresses(
      USDCTokenPoolProxy.PoolAddresses({
        cctpV1Pool: s_cctpV1Pool,
        cctpV2Pool: s_cctpV2Pool,
        cctpV2PoolWithCCV: address(0),
        siloedLockReleasePool: address(0)
      })
    );

    vm.mockCall(
      address(s_router), abi.encodeCall(Router.getOnRamp, s_chainSelForCCV), abi.encode(s_routerAllowedOnRamp)
    );

    vm.expectRevert(abi.encodeWithSelector(USDCTokenPoolProxy.NoLockOrBurnMechanismSet.selector, s_chainSelForCCV));

    vm.startPrank(s_routerAllowedOnRamp);
    s_usdcTokenPoolProxy.lockOrBurn(
      Pool.LockOrBurnInV1({
        receiver: abi.encode(s_receiver),
        remoteChainSelector: s_chainSelForCCV,
        originalSender: s_sender,
        amount: amount,
        localToken: address(s_USDCToken)
      }),
      0,
      ""
    );
  }

  function test_lockOrBurn_CCTPV2WithCCV_RevertWhen_ChainNotSupportedByVerifier() public {
    uint256 amount = 100;
    bytes memory tokenArgs = abi.encode(abi.encode(1, 2, 3));
    uint16 blockConfirmationRequested = 1;
    address verifierImpl = address(0);

    vm.mockCall(
      address(s_cctpVerifier),
      abi.encodeWithSelector(
        ICrossChainVerifierResolver.getOutboundImplementation.selector, s_chainSelForCCV, tokenArgs
      ),
      abi.encode(verifierImpl)
    );

    vm.expectRevert(abi.encodeWithSelector(USDCTokenPoolProxy.ChainNotSupportedByVerifier.selector, s_chainSelForCCV));

    vm.startPrank(s_routerAllowedOnRamp);
    s_usdcTokenPoolProxy.lockOrBurn(
      Pool.LockOrBurnInV1({
        receiver: abi.encode(s_receiver),
        remoteChainSelector: s_chainSelForCCV,
        originalSender: s_sender,
        amount: amount,
        localToken: address(s_USDCToken)
      }),
      blockConfirmationRequested,
      tokenArgs
    );
  }

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

  function test_lockOrBurn_RevertWhen_MustSetPoolForMechanism() public {
    // First, remove the lock release pool from the proxy.
    s_usdcTokenPoolProxy.updatePoolAddresses(
      USDCTokenPoolProxy.PoolAddresses({
        cctpV1Pool: s_cctpV1Pool,
        cctpV2Pool: s_cctpV2Pool,
        cctpV2PoolWithCCV: s_cctpThroughCCVTokenPool,
        siloedLockReleasePool: address(0)
      })
    );

    // Configure lock or burn mechanisms for different chains but do not set the lock release pool for the chain.
    uint64[] memory chainSelectors = new uint64[](1);
    chainSelectors[0] = s_chainSelFoLockRelease;

    USDCTokenPoolProxy.LockOrBurnMechanism[] memory mechanisms = new USDCTokenPoolProxy.LockOrBurnMechanism[](1);
    mechanisms[0] = USDCTokenPoolProxy.LockOrBurnMechanism.LOCK_RELEASE;

    // Now the mechanism cannot be set if the pool is not configured.
    vm.expectRevert(
      abi.encodeWithSelector(
        USDCTokenPoolProxy.MustSetPoolForMechanism.selector, s_chainSelFoLockRelease, mechanisms[0]
      )
    );
    s_usdcTokenPoolProxy.updateLockOrBurnMechanisms(chainSelectors, mechanisms);
  }

  function test_lockOrBurn_RevertWhen_Unauthorized() public {
    address unauthorized = makeAddr("unauthorized");

    vm.startPrank(unauthorized);

    Pool.LockOrBurnInV1 memory lockOrBurnIn = Pool.LockOrBurnInV1({
      receiver: abi.encode(s_receiver),
      remoteChainSelector: s_chainSelForV2,
      originalSender: s_sender,
      amount: 100,
      localToken: address(s_USDCToken)
    });

    vm.expectRevert(abi.encodeWithSelector(USDCTokenPoolProxy.CallerIsNotARampOnRouter.selector, unauthorized));
    s_usdcTokenPoolProxy.lockOrBurn(lockOrBurnIn);
  }

  function test_lockOrBurn_RevertWhen_NoLockOrBurnMechanismSet_PoolRemovedAfterMechanismConfigured() public {
    // This tests the path where a non-CCV mechanism (CCTP_V1, CCTP_V2, LOCK_RELEASE) is configured
    // but the corresponding pool is later removed (set to address(0)).

    // First, configure CCTP_V2 mechanism for a chain (already done in setUp).
    // Then remove the CCTP_V2 pool.
    s_usdcTokenPoolProxy.updatePoolAddresses(
      USDCTokenPoolProxy.PoolAddresses({
        cctpV1Pool: s_cctpV1Pool,
        cctpV2Pool: address(0), // Remove CCTP_V2 pool
        cctpV2PoolWithCCV: s_cctpThroughCCVTokenPool,
        siloedLockReleasePool: s_lockReleasePool
      })
    );

    Pool.LockOrBurnInV1 memory lockOrBurnIn = Pool.LockOrBurnInV1({
      receiver: abi.encode(s_receiver),
      remoteChainSelector: s_chainSelForV2, // Uses CCTP_V2 mechanism
      originalSender: s_sender,
      amount: 100,
      localToken: address(s_USDCToken)
    });

    vm.expectRevert(abi.encodeWithSelector(USDCTokenPoolProxy.NoLockOrBurnMechanismSet.selector, s_chainSelForV2));

    vm.startPrank(s_routerAllowedOnRamp);
    s_usdcTokenPoolProxy.lockOrBurn(lockOrBurnIn);
  }
}
