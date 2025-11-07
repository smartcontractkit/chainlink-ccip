// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ExtraArgsCodec} from "../../../libraries/ExtraArgsCodec.sol";
import {BaseTest} from "../../BaseTest.t.sol";
import {ExtraArgsCodecHelper} from "../../helpers/ExtraArgsCodecHelpers.sol";

/// forge-config: default.allow_internal_expect_revert = true
contract ExtraArgsCodecSVM_Test is BaseTest {
  ExtraArgsCodecHelper internal s_helper;

  function setUp() public override {
    super.setUp();
    s_helper = new ExtraArgsCodecHelper();
  }

  function test_DecodeSVMExecutorArgsV1_Empty() public view {
    ExtraArgsCodec.SVMExecutorArgsV1 memory args = ExtraArgsCodec.SVMExecutorArgsV1({
      useATA: ExtraArgsCodec.SVMATAUsage.DERIVE_ACCOUNT_AND_CREATE,
      accountIsWritableBitmap: 0,
      accounts: new bytes32[](0)
    });

    bytes memory encoded = ExtraArgsCodec._encodeSVMExecutorArgsV1(args);
    assertEq(encoded.length, ExtraArgsCodec.SVM_EXECUTOR_ARGS_V1_BASE_SIZE);

    ExtraArgsCodec.SVMExecutorArgsV1 memory decoded = s_helper._decodeSVMExecutorArgsV1(encoded);

    assertEq(uint8(decoded.useATA), uint8(args.useATA));
    assertEq(decoded.accountIsWritableBitmap, args.accountIsWritableBitmap);
    assertEq(decoded.accounts.length, args.accounts.length);
  }

  function test_DecodeSVMExecutorArgsV1_WithAccounts() public view {
    bytes32[] memory accounts = new bytes32[](2);
    accounts[0] = bytes32(uint256(1));
    accounts[1] = bytes32(uint256(2));

    ExtraArgsCodec.SVMExecutorArgsV1 memory args = ExtraArgsCodec.SVMExecutorArgsV1({
      useATA: ExtraArgsCodec.SVMATAUsage.DERIVE_ACCOUNT_DONT_CREATE,
      accountIsWritableBitmap: 0x03,
      accounts: accounts
    });

    bytes memory encoded = ExtraArgsCodec._encodeSVMExecutorArgsV1(args);
    ExtraArgsCodec.SVMExecutorArgsV1 memory decoded = s_helper._decodeSVMExecutorArgsV1(encoded);

    assertEq(uint8(decoded.useATA), uint8(args.useATA));
    assertEq(decoded.accountIsWritableBitmap, args.accountIsWritableBitmap);
    assertEq(decoded.accounts.length, args.accounts.length);
    assertEq(decoded.accounts[0], args.accounts[0]);
  }

  function test_DecodeSVMExecutorArgsV1_MaxBitmap() public view {
    ExtraArgsCodec.SVMExecutorArgsV1 memory args = ExtraArgsCodec.SVMExecutorArgsV1({
      useATA: ExtraArgsCodec.SVMATAUsage.DONT_DERIVE_ACCOUNT,
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
    ExtraArgsCodec.SVMExecutorArgsV1 memory args = ExtraArgsCodec.SVMExecutorArgsV1({
      useATA: ExtraArgsCodec.SVMATAUsage.DERIVE_ACCOUNT_AND_CREATE,
      accountIsWritableBitmap: 0,
      accounts: new bytes32[](0)
    });

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
