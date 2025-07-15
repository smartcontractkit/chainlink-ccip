// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolV1} from "../../../../interfaces/IPool.sol";
import {ITokenMessenger} from "../../../../pools/USDC/interfaces/ITokenMessenger.sol";
import {IERC165} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v5.0.2/contracts/utils/introspection/IERC165.sol";

import {USDCTokenPool} from "../../../../pools/USDC/USDCTokenPool.sol";
import {USDCTokenPoolCCTPV2} from "../../../../pools/USDC/USDCTokenPoolCCTPV2.sol";
import {USDCTokenPoolCCTPV2Setup} from "./USDCTokenPoolCCTPV2Setup.t.sol";

contract USDCTokenPoolCCTPV2_constructor is USDCTokenPoolCCTPV2Setup {
  function test_constructor() public {
    new USDCTokenPoolCCTPV2(
      s_mockLegacyUSDC,
      s_mockUSDC,
      s_cctpMessageTransmitterProxy,
      s_USDCToken,
      new address[](0),
      address(s_mockRMNRemote),
      address(s_router),
      s_previousPool
    );
  }

  function test_constructor_PreviousPoolZeroAddress() public {
    USDCTokenPoolCCTPV2 usdcTokenPool = new USDCTokenPoolCCTPV2(
      s_mockLegacyUSDC,
      s_mockUSDC,
      s_cctpMessageTransmitterProxy,
      s_USDCToken,
      new address[](0),
      address(s_mockRMNRemote),
      address(s_router),
      address(0)
    );

    assertEq(usdcTokenPool.i_previousPool(), address(0));
  }

  function test_constructor_RevertWhen_TokenMessangerAddressZero() public {
    vm.expectRevert(USDCTokenPool.InvalidConfig.selector);
    new USDCTokenPoolCCTPV2(
      s_mockLegacyUSDC,
      ITokenMessenger(address(0)),
      s_cctpMessageTransmitterProxy,
      s_USDCToken,
      new address[](0),
      address(s_mockRMNRemote),
      address(s_router),
      s_previousPool
    );
  }

  function test_constructor_RevertWhen_TransmitterVersionDoesNotMatchSupportedUSDCVersion() public {
    uint32 transmitterVersion = uint32(vm.randomUint());
    vm.mockCall(
      address(s_mockUSDCTransmitterCCTPV2),
      abi.encodeCall(s_mockUSDCTransmitter.version, ()),
      abi.encode(transmitterVersion)
    );
    vm.expectRevert(abi.encodeWithSelector(USDCTokenPool.InvalidMessageVersion.selector, transmitterVersion));
    new USDCTokenPoolCCTPV2(
      s_mockLegacyUSDC,
      s_mockUSDC,
      s_cctpMessageTransmitterProxy,
      s_USDCToken,
      new address[](0),
      address(s_mockRMNRemote),
      address(s_router),
      s_previousPool
    );
  }

  function test_constructor_RevertWhen_TokenMessengerVersionDoesNotMatchSupportedUSDCVersion() public {
    uint32 tokenMessengerVersion = s_mockUSDC.messageBodyVersion() + 1;
    vm.mockCall(
      address(s_mockUSDC), abi.encodeCall(s_mockUSDC.messageBodyVersion, ()), abi.encode(tokenMessengerVersion)
    );
    vm.expectRevert(abi.encodeWithSelector(USDCTokenPool.InvalidTokenMessengerVersion.selector, tokenMessengerVersion));
    new USDCTokenPoolCCTPV2(
      s_mockLegacyUSDC,
      s_mockUSDC,
      s_cctpMessageTransmitterProxy,
      s_USDCToken,
      new address[](0),
      address(s_mockRMNRemote),
      address(s_router),
      s_previousPool
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
      s_mockLegacyUSDC,
      s_mockUSDC,
      s_cctpMessageTransmitterProxy,
      s_USDCToken,
      new address[](0),
      address(s_mockRMNRemote),
      address(s_router),
      s_previousPool
    );
  }

  function test_constructor_RevertWhen_InvalidPreviousPool_AddressThis() public {
    // Deploy the pool using CREATE2 to predetermine its address, so that we can test the InvalidPreviousPool error
    bytes memory bytecode = abi.encodePacked(
      type(USDCTokenPoolCCTPV2).creationCode,
      abi.encode(
        s_mockLegacyUSDC,
        s_mockUSDC,
        s_cctpMessageTransmitterProxy,
        s_USDCToken,
        new address[](0),
        address(s_mockRMNRemote),
        address(s_router),
        s_previousPool // placeholder, will be replaced below
      )
    );

    bytes32 salt = keccak256("USDCTokenPoolSelfPreviousPool");
    // The constructor expects the previousPool address as the last argument, so we need to compute the address
    address predictedAddress = address(
      uint160( // downcast to address size
      uint256(keccak256(abi.encodePacked(bytes1(0xff), address(this), salt, keccak256(bytecode)))))
    );

    // Now, re-encode the constructor args with the predicted address as previousPool
    bytes memory constructorArgs = abi.encode(
      s_mockLegacyUSDC,
      s_mockUSDC,
      s_cctpMessageTransmitterProxy,
      s_USDCToken,
      new address[](0),
      address(s_mockRMNRemote),
      address(s_router),
      predictedAddress
    );
    // Concatenate the contract creation code and constructor arguments to form the full bytecode for deployment
    bytes memory fullBytecode = abi.encodePacked(type(USDCTokenPoolCCTPV2).creationCode, constructorArgs);

    // Expect the constructor to revert with InvalidPreviousPool, passing the predicted address as the argument
    vm.expectRevert(abi.encodeWithSelector(USDCTokenPool.InvalidPreviousPool.selector, predictedAddress));
    address deployed;
    // Deploy the contract using CREATE2 with the given salt and full bytecode
    // solhint-disable-next-line no-inline-assembly
    assembly {
      deployed := create2(0, add(fullBytecode, 0x20), mload(fullBytecode), salt)
    }
    // Assert that the deployment failed (address(0) means contract was not deployed)
    assertEq(deployed, address(0), "Pool should not have deployed");
  }

  function test_constructor_RevertWhen_InvalidPreviousPool_UnsupportedFunctions() public {
    // Create an address for an invalid previous pool that doesn't support IPoolV1 interface
    address invalidPreviousPool = makeAddr("INVALID_PREVIOUS_POOL");

    // Mock the supportsInterface call to return false for IPoolV1.interfaceId
    vm.mockCall(
      invalidPreviousPool,
      abi.encodeWithSelector(IERC165.supportsInterface.selector, type(IPoolV1).interfaceId),
      abi.encode(false)
    );

    // Expect the constructor to revert with InvalidPreviousPool error
    vm.expectRevert(USDCTokenPool.InvalidPreviousPool.selector);

    new USDCTokenPoolCCTPV2(
      s_mockLegacyUSDC,
      s_mockUSDC,
      s_cctpMessageTransmitterProxy,
      s_USDCToken,
      new address[](0),
      address(s_mockRMNRemote),
      address(s_router),
      invalidPreviousPool
    );
  }

  function test_constructor_RevertWhen_InvalidPreviousPool_MessageTransmitterProxyReverts() public {
    // Create a mock previous pool address
    address mockPreviousPool = makeAddr("MOCK_PREVIOUS_POOL");

    // Mock supportsInterface to return true so it passes the first check
    vm.mockCall(
      mockPreviousPool,
      abi.encodeWithSelector(IERC165.supportsInterface.selector, type(IPoolV1).interfaceId),
      abi.encode(true)
    );

    // Mock i_messageTransmitterProxy() to revert by returning an error
    vm.mockCallRevert(mockPreviousPool, abi.encodeWithSelector(bytes4(keccak256("i_messageTransmitterProxy()"))), "");

    // Expect the constructor to revert with InvalidPreviousPool
    vm.expectRevert(USDCTokenPool.InvalidPreviousPool.selector);

    new USDCTokenPoolCCTPV2(
      s_mockLegacyUSDC,
      s_mockUSDC,
      s_cctpMessageTransmitterProxy,
      s_USDCToken,
      new address[](0),
      address(s_mockRMNRemote),
      address(s_router),
      mockPreviousPool
    );
  }
}
