// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {USDCTokenPool} from "../../../../pools/USDC/USDCTokenPool.sol";
import {ITokenMessenger} from "../../../../pools/USDC/interfaces/ITokenMessenger.sol";
import {MockE2EUSDCTransmitter} from "../../../mocks/MockE2EUSDCTransmitter.sol";
import {USDCTokenPoolSetup} from "./USDCTokenPoolSetup.t.sol";

contract USDCTokenPool_constructor is USDCTokenPoolSetup {
  function test_constructor() public {
    new USDCTokenPool(
      s_mockUSDCTokenMessenger,
      s_cctpMessageTransmitterProxy,
      s_USDCToken,
      new address[](0),
      address(s_mockRMNRemote),
      address(s_router),
      0
    );
  }

  function test_constructor_RevertWhen_TokenMessengerAddressZero() public {
    vm.expectRevert(USDCTokenPool.InvalidConfig.selector);
    new USDCTokenPool(
      ITokenMessenger(address(0)),
      s_cctpMessageTransmitterProxy,
      s_USDCToken,
      new address[](0),
      address(s_mockRMNRemote),
      address(s_router),
      0
    );
  }

  function test_constructor_RevertWhen_InvalidMessageVersion() public {
    // Should revert with InvalidMessageVersion error because the token messenger version is 0, but the token pool itself is being set with version of 1
    vm.expectRevert(abi.encodeWithSelector(USDCTokenPool.InvalidMessageVersion.selector, 1, 0));
    new USDCTokenPool(
      s_mockUSDCTokenMessenger,
      s_cctpMessageTransmitterProxy,
      s_USDCToken,
      new address[](0),
      address(s_mockRMNRemote),
      address(s_router),
      1
    );
  }

  function test_constructor_RevertWhen_InvalidTokenMessengerVersion() public {
    vm.mockCall(
      address(s_mockUSDCTransmitter), abi.encodeWithSelector(MockE2EUSDCTransmitter.version.selector), abi.encode(1)
    );

    // Should revert with InvalidTokenMessengerVersion error because the token messenger version is 0, but the transmitter version is 1
    vm.expectRevert(abi.encodeWithSelector(USDCTokenPool.InvalidTokenMessengerVersion.selector, 1, 0));
    new USDCTokenPool(
      s_mockUSDCTokenMessenger,
      s_cctpMessageTransmitterProxy,
      s_USDCToken,
      new address[](0),
      address(s_mockRMNRemote),
      address(s_router),
      1
    );
  }

  function test_constructor_RevertWhen_InvalidTransmitterVersionInProxy() public {
    address transmitterAddress = makeAddr("INVALID_TRANSMITTER");
    vm.mockCall(
      address(s_cctpMessageTransmitterProxy),
      abi.encodeCall(s_cctpMessageTransmitterProxy.i_cctpTransmitter, ()),
      abi.encode(transmitterAddress)
    );

    vm.expectRevert(abi.encodeWithSelector(USDCTokenPool.InvalidTransmitterInProxy.selector));
    new USDCTokenPool(
      s_mockUSDCTokenMessenger,
      s_cctpMessageTransmitterProxy,
      s_USDCToken,
      new address[](0),
      address(s_mockRMNRemote),
      address(s_router),
      0
    );
  }
}
