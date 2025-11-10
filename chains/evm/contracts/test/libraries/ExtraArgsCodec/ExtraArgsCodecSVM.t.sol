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

  function test__decodeSVMExecutorArgsV1_Empty() public view {
    ExtraArgsCodec.SVMExecutorArgsV1 memory args = ExtraArgsCodec.SVMExecutorArgsV1({
      useATA: ExtraArgsCodec.SVMTokenReceiverUsage.DERIVE_ATA_AND_CREATE,
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

  function test__decodeSVMExecutorArgsV1_WithAccounts() public view {
    bytes32[] memory accounts = new bytes32[](2);
    accounts[0] = bytes32(uint256(1));
    accounts[1] = bytes32(uint256(2));

    ExtraArgsCodec.SVMExecutorArgsV1 memory args = ExtraArgsCodec.SVMExecutorArgsV1({
      useATA: ExtraArgsCodec.SVMTokenReceiverUsage.DERIVE_ATA_DONT_CREATE,
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

  function test__decodeSVMExecutorArgsV1_MaxBitmap() public view {
    ExtraArgsCodec.SVMExecutorArgsV1 memory args = ExtraArgsCodec.SVMExecutorArgsV1({
      useATA: ExtraArgsCodec.SVMTokenReceiverUsage.USE_AS_IS,
      accountIsWritableBitmap: type(uint64).max,
      accounts: new bytes32[](0)
    });

    bytes memory encoded = ExtraArgsCodec._encodeSVMExecutorArgsV1(args);
    ExtraArgsCodec.SVMExecutorArgsV1 memory decoded = s_helper._decodeSVMExecutorArgsV1(encoded);

    assertEq(decoded.accountIsWritableBitmap, type(uint64).max);
    assertEq(uint8(decoded.useATA), uint8(args.useATA));
  }

  function test__decodeSVMExecutorArgsV1_RevertWhen_EXTRA_ARGS_STATIC_LENGTH_FIELDS() public {
    vm.expectRevert(
      abi.encodeWithSelector(
        ExtraArgsCodec.InvalidDataLength.selector,
        ExtraArgsCodec.EncodingErrorLocation.EXTRA_ARGS_STATIC_LENGTH_FIELDS,
        10
      )
    );
    s_helper._decodeSVMExecutorArgsV1(new bytes(10));
  }

  function test__decodeSVMExecutorArgsV1_RevertWhen_SVM_EXECUTOR_FINAL_OFFSET() public {
    ExtraArgsCodec.SVMExecutorArgsV1 memory args = ExtraArgsCodec.SVMExecutorArgsV1({
      useATA: ExtraArgsCodec.SVMTokenReceiverUsage.DERIVE_ATA_AND_CREATE,
      accountIsWritableBitmap: 0,
      accounts: new bytes32[](0)
    });

    bytes memory encoded = ExtraArgsCodec._encodeSVMExecutorArgsV1(args);
    bytes memory withExtra = bytes.concat(encoded, bytes("extra"));

    vm.expectRevert(
      abi.encodeWithSelector(
        ExtraArgsCodec.InvalidDataLength.selector, ExtraArgsCodec.EncodingErrorLocation.SVM_EXECUTOR_FINAL_OFFSET, 14
      )
    );
    s_helper._decodeSVMExecutorArgsV1(withExtra);
  }

  function test__decodeSVMExecutorArgsV1_RevertWhen_SVM_EXECUTOR_ACCOUNTS_CONTENT() public {
    bytes memory invalidData = abi.encodePacked(
      ExtraArgsCodec.SVM_EXECUTOR_ARGS_V1_TAG,
      uint8(ExtraArgsCodec.SVMTokenReceiverUsage.DERIVE_ATA_AND_CREATE),
      uint64(0),
      uint8(2) // Claims 2 accounts but doesn't provide them.
    );

    vm.expectRevert(
      abi.encodeWithSelector(
        ExtraArgsCodec.InvalidDataLength.selector,
        ExtraArgsCodec.EncodingErrorLocation.SVM_EXECUTOR_ACCOUNTS_CONTENT,
        14
      )
    );
    s_helper._decodeSVMExecutorArgsV1(invalidData);
  }

  function test__encodeSVMExecutorArgsV1_RevertWhen_ENCODE_SVM_ACCOUNTS_LENGTH() public {
    bytes32[] memory accounts = new bytes32[](257);

    vm.expectRevert(
      abi.encodeWithSelector(
        ExtraArgsCodec.InvalidDataLength.selector, ExtraArgsCodec.EncodingErrorLocation.ENCODE_SVM_ACCOUNTS_LENGTH, 0
      )
    );
    ExtraArgsCodec._encodeSVMExecutorArgsV1(
      ExtraArgsCodec.SVMExecutorArgsV1({
        useATA: ExtraArgsCodec.SVMTokenReceiverUsage.DERIVE_ATA_AND_CREATE,
        accountIsWritableBitmap: 0,
        accounts: accounts
      })
    );
  }
}
