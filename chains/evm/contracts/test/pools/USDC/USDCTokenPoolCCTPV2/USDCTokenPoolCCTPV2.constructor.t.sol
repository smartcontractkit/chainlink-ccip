// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ITokenMessenger} from "../../../../pools/USDC/interfaces/ITokenMessenger.sol";

import {USDCTokenPool} from "../../../../pools/USDC/USDCTokenPool.sol";
import {USDCTokenPoolCCTPV2} from "../../../../pools/USDC/USDCTokenPoolCCTPV2.sol";

import {MockUSDCTokenMessenger} from "../../../mocks/MockUSDCTokenMessenger.sol";
import {USDCTokenPoolCCTPV2Setup} from "./USDCTokenPoolCCTPV2Setup.t.sol";

contract USDCTokenPoolCCTPV2_constructor is USDCTokenPoolCCTPV2Setup {
  function test_constructor() public {
    new USDCTokenPoolCCTPV2(
      s_mockUSDC,
      s_cctpMessageTransmitterProxy,
      s_USDCToken,
      new address[](0),
      address(s_mockRMNRemote),
      address(s_router)
    );
  }

  function test_constructor_RevertWhen_TokenMessengerAddressZero() public {
    vm.expectRevert(USDCTokenPool.InvalidConfig.selector);
    new USDCTokenPoolCCTPV2(
      ITokenMessenger(address(0)),
      s_cctpMessageTransmitterProxy,
      s_USDCToken,
      new address[](0),
      address(s_mockRMNRemote),
      address(s_router)
    );
  }

  function test_constructor_RevertWhen_InvalidMessageVersion() public {
    vm.expectRevert(abi.encodeWithSelector(USDCTokenPool.InvalidMessageVersion.selector, 0, 1));

    new USDCTokenPoolCCTPV2(
      s_mockLegacyUSDC,
      s_cctpMessageTransmitterProxy,
      s_USDCToken,
      new address[](0),
      address(s_mockRMNRemote),
      address(s_router)
    );
  }

  function test_constructor_RevertWhen_InvalidTokenMessengerVersion() public {
    // The error we want to call is most likely unreachable because the token messenger version is 1, but we mock it to
    // 0 to test the error
    vm.mockCall(
      address(s_mockUSDC), abi.encodeWithSelector(MockUSDCTokenMessenger.messageBodyVersion.selector), abi.encode(0)
    );

    vm.expectRevert(abi.encodeWithSelector(USDCTokenPool.InvalidTokenMessengerVersion.selector, 0, 1));
    new USDCTokenPoolCCTPV2(
      s_mockUSDC,
      s_cctpMessageTransmitterProxy,
      s_USDCToken,
      new address[](0),
      address(s_mockRMNRemote),
      address(s_router)
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
    new USDCTokenPoolCCTPV2(
      s_mockUSDC,
      s_cctpMessageTransmitterProxy,
      s_USDCToken,
      new address[](0),
      address(s_mockRMNRemote),
      address(s_router)
    );
  }
}
