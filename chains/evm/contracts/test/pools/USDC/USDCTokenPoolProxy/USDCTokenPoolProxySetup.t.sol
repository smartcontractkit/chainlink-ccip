// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {USDCTokenPoolProxy} from "../../../../pools/USDC/USDCTokenPoolProxy.sol";
import {USDCSetup} from "../USDCSetup.t.sol";

import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";
import {IERC165} from "@openzeppelin/contracts@5.3.0/utils/introspection/IERC165.sol";

contract USDCTokenPoolProxySetup is USDCSetup {
  address internal s_legacyCctpV1Pool = makeAddr("legacyCctpV1Pool");
  address internal s_cctpV1Pool = makeAddr("cctpV1Pool");
  address internal s_cctpV2Pool = makeAddr("cctpV2Pool");
  address internal s_cctpTokenPool = makeAddr("cctpTokenPool");
  address internal s_lockReleasePool = makeAddr("lockReleasePool");
  address internal s_mockTransmitterProxy = makeAddr("mockTransmitterProxy");
  address internal s_cctpVerifier = makeAddr("cctpVerifier");

  uint64 internal s_remoteLockReleaseChainSelector = 12345;
  uint64 internal s_remoteCCTPChainSelector = 12346;

  USDCTokenPoolProxy internal s_usdcTokenPoolProxy;

  function setUp() public virtual override {
    super.setUp();

    // Mock the transmitter proxy's receiveMessage function to return true
    vm.mockCall(
      address(s_cctpV1Pool),
      abi.encodeWithSelector(bytes4(keccak256("i_messageTransmitterProxy()"))),
      abi.encode(s_mockTransmitterProxy)
    );

    // Deploy the proxy
    s_usdcTokenPoolProxy = new USDCTokenPoolProxy(
      s_USDCToken,
      USDCTokenPoolProxy.PoolAddresses({
        legacyCctpV1Pool: s_legacyCctpV1Pool,
        cctpV1Pool: s_cctpV1Pool,
        cctpV2Pool: s_cctpV2Pool,
        cctpTokenPool: s_cctpTokenPool,
        siloedUsdCTokenPool: address(0)
      }),
      address(s_router),
      address(s_cctpVerifier)
    );

    // Deal some tokens to the proxy to test the transfer to the destination pool
    deal(address(s_USDCToken), address(s_usdcTokenPoolProxy), 1000e6);

    bytes[] memory sourcePoolAddresses = new bytes[](1);
    sourcePoolAddresses[0] = abi.encode(SOURCE_CHAIN_USDC_POOL);

    address[] memory addedCallers = new address[](4);
    addedCallers[0] = address(s_cctpV1Pool);
    addedCallers[1] = address(s_cctpV2Pool);
    addedCallers[2] = address(s_lockReleasePool);
    addedCallers[3] = address(s_cctpTokenPool);
    s_cctpMessageTransmitterProxy.applyAuthorizedCallerUpdates(
      AuthorizedCallers.AuthorizedCallerArgs({addedCallers: addedCallers, removedCallers: new address[](0)})
    );

    // Configure the lockOrBurn mechanism for the remote chain selectors.
    uint64[] memory remoteChainSelectors = new uint64[](2);
    remoteChainSelectors[0] = s_remoteLockReleaseChainSelector;
    remoteChainSelectors[1] = s_remoteCCTPChainSelector;

    USDCTokenPoolProxy.LockOrBurnMechanism[] memory mechanisms = new USDCTokenPoolProxy.LockOrBurnMechanism[](2);
    mechanisms[0] = USDCTokenPoolProxy.LockOrBurnMechanism.LOCK_RELEASE;
    mechanisms[1] = USDCTokenPoolProxy.LockOrBurnMechanism.CCTP_V2_WITH_CCV;
    s_usdcTokenPoolProxy.updateLockOrBurnMechanisms(remoteChainSelectors, mechanisms);
  }

  function _enableERC165InterfaceChecks(
    address pool,
    bytes4 interfaceId
  ) internal {
    vm.mockCall(
      address(pool), abi.encodeWithSelector(IERC165.supportsInterface.selector, interfaceId), abi.encode(true)
    );

    vm.mockCall(
      address(pool),
      abi.encodeWithSelector(IERC165.supportsInterface.selector, type(IERC165).interfaceId),
      abi.encode(true)
    );
  }
}
