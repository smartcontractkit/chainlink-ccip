// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolV1} from "../../../../interfaces/IPool.sol";
import {ITokenMessenger} from "../../../../pools/USDC/interfaces/ITokenMessenger.sol";
import {IERC165} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v5.0.2/contracts/utils/introspection/IERC165.sol";

import {USDCTokenPool} from "../../../../pools/USDC/USDCTokenPool.sol";
import {USDCSetup} from "../USDCSetup.t.sol";

contract USDCTokenPool_constructor is USDCSetup {
  function test_constructor() public {
    new USDCTokenPool(
      s_mockUSDC,
      s_cctpMessageTransmitterProxy,
      s_USDCToken,
      new address[](0),
      address(s_mockRMNRemote),
      address(s_router),
      s_previousPool
    );
  }

  function test_constructor_RevertWhen_TokenMessangerAddressZero() public {
    vm.expectRevert(USDCTokenPool.InvalidConfig.selector);
    new USDCTokenPool(
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
      address(s_mockUSDCTransmitter), abi.encodeCall(s_mockUSDCTransmitter.version, ()), abi.encode(transmitterVersion)
    );
    vm.expectRevert(abi.encodeWithSelector(USDCTokenPool.InvalidMessageVersion.selector, transmitterVersion));
    new USDCTokenPool(
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
    new USDCTokenPool(
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
    new USDCTokenPool(
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
      type(USDCTokenPool).creationCode,
      abi.encode(
        s_mockUSDC,
        s_cctpMessageTransmitterProxy,
        s_USDCToken,
        new address[](0),
        address(s_mockRMNRemote),
        address(s_router),
        bytes32(uint256(0)) // placeholder, will be replaced below
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
      s_mockUSDC,
      s_cctpMessageTransmitterProxy,
      s_USDCToken,
      new address[](0),
      address(s_mockRMNRemote),
      address(s_router),
      predictedAddress
    );
    // Concatenate the contract creation code and constructor arguments to form the full bytecode for deployment
    bytes memory fullBytecode = abi.encodePacked(type(USDCTokenPool).creationCode, constructorArgs);

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

    new USDCTokenPool(
      s_mockUSDC,
      s_cctpMessageTransmitterProxy,
      s_USDCToken,
      new address[](0),
      address(s_mockRMNRemote),
      address(s_router),
      invalidPreviousPool
    );
  }
}
