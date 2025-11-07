// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

/// forge-config: default.allow_internal_expect_revert = true

import {ExtraArgsCodec} from "../../../libraries/ExtraArgsCodec.sol";
import {BaseTest} from "../../BaseTest.t.sol";
import {ExtraArgsCodecHelper} from "../../helpers/ExtraArgsCodecHelpers.sol";

contract ExtraArgsCodecSVM_Test is BaseTest {
  ExtraArgsCodecHelper internal s_helper;

  function setUp() public override {
    super.setUp();
    s_helper = new ExtraArgsCodecHelper();
  }

  function test_DecodeSVMExecutorArgsV1_NoAccounts() public view {
    ExtraArgsCodec.SVMExecutorArgsV1 memory args =
      ExtraArgsCodec.SVMExecutorArgsV1({useATA: false, accountIsWritableBitmap: 0, accounts: new bytes32[](0)});

    bytes memory encoded = ExtraArgsCodec._encodeSVMExecutorArgsV1(args);
    ExtraArgsCodec.SVMExecutorArgsV1 memory decoded = s_helper._decodeSVMExecutorArgsV1(encoded);

    assertEq(decoded.useATA, false);
    assertEq(decoded.accountIsWritableBitmap, 0);
    assertEq(decoded.accounts.length, 0);
  }

  function test_DecodeSVMExecutorArgsV1_WithAccounts() public view {
    bytes32[] memory accounts = new bytes32[](2);
    accounts[0] = bytes32(uint256(1));
    accounts[1] = bytes32(uint256(2));

    ExtraArgsCodec.SVMExecutorArgsV1 memory args =
      ExtraArgsCodec.SVMExecutorArgsV1({useATA: true, accountIsWritableBitmap: 0x03, accounts: accounts});

    bytes memory encoded = ExtraArgsCodec._encodeSVMExecutorArgsV1(args);
    ExtraArgsCodec.SVMExecutorArgsV1 memory decoded = s_helper._decodeSVMExecutorArgsV1(encoded);

    assertEq(decoded.useATA, true);
    assertEq(decoded.accountIsWritableBitmap, 0x03);
    assertEq(decoded.accounts.length, 2);
    assertEq(decoded.accounts[0], bytes32(uint256(1)));
  }

  function test_DecodeSVMExecutorArgsV1_MaxBitmap() public view {
    ExtraArgsCodec.SVMExecutorArgsV1 memory args = ExtraArgsCodec.SVMExecutorArgsV1({
      useATA: false,
      accountIsWritableBitmap: type(uint64).max,
      accounts: new bytes32[](0)
    });

    bytes memory encoded = ExtraArgsCodec._encodeSVMExecutorArgsV1(args);
    ExtraArgsCodec.SVMExecutorArgsV1 memory decoded = s_helper._decodeSVMExecutorArgsV1(encoded);

    assertEq(decoded.accountIsWritableBitmap, type(uint64).max);
  }

  function test_DecodeSVMExecutorArgsV1_RevertWhen_DataTooShort() public {
    vm.expectRevert(
      abi.encodeWithSelector(
        ExtraArgsCodec.InvalidDataLength.selector, ExtraArgsCodec.EncodingErrorLocation.EXTRA_ARGS_STATIC_LENGTH_FIELDS
      )
    );
    s_helper._decodeSVMExecutorArgsV1(new bytes(10));
  }

  function test_DecodeSVMExecutorArgsV1_RevertWhen_ExtraBytes() public {
    ExtraArgsCodec.SVMExecutorArgsV1 memory args =
      ExtraArgsCodec.SVMExecutorArgsV1({useATA: false, accountIsWritableBitmap: 0, accounts: new bytes32[](0)});

    bytes memory encoded = ExtraArgsCodec._encodeSVMExecutorArgsV1(args);
    bytes memory withExtra = bytes.concat(encoded, bytes("extra"));

    vm.expectRevert(
      abi.encodeWithSelector(
        ExtraArgsCodec.InvalidDataLength.selector, ExtraArgsCodec.EncodingErrorLocation.SVM_EXECUTOR_FINAL_OFFSET
      )
    );
    s_helper._decodeSVMExecutorArgsV1(withExtra);
  }
}
