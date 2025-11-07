// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

/// forge-config: default.allow_internal_expect_revert = true

import {ExtraArgsCodec} from "../../../libraries/ExtraArgsCodec.sol";
import {BaseTest} from "../../BaseTest.t.sol";
import {ExtraArgsCodecHelper} from "../../helpers/ExtraArgsCodecHelpers.sol";

contract ExtraArgsCodecSui_Test is BaseTest {
  ExtraArgsCodecHelper internal s_helper;

  function setUp() public override {
    super.setUp();
    s_helper = new ExtraArgsCodecHelper();
  }

  function test_DecodeSuiExecutorArgsV1_NoObjectIds() public view {
    ExtraArgsCodec.SuiExecutorArgsV1 memory args =
      ExtraArgsCodec.SuiExecutorArgsV1({receiverObjectIds: new bytes32[](0)});

    bytes memory encoded = ExtraArgsCodec._encodeSuiExecutorArgsV1(args);
    ExtraArgsCodec.SuiExecutorArgsV1 memory decoded = s_helper._decodeSuiExecutorArgsV1(encoded);

    assertEq(decoded.receiverObjectIds.length, 0);
  }

  function test_DecodeSuiExecutorArgsV1_WithObjectIds() public view {
    bytes32[] memory objectIds = new bytes32[](2);
    objectIds[0] = keccak256("object1");
    objectIds[1] = keccak256("object2");

    ExtraArgsCodec.SuiExecutorArgsV1 memory args = ExtraArgsCodec.SuiExecutorArgsV1({receiverObjectIds: objectIds});

    bytes memory encoded = ExtraArgsCodec._encodeSuiExecutorArgsV1(args);
    ExtraArgsCodec.SuiExecutorArgsV1 memory decoded = s_helper._decodeSuiExecutorArgsV1(encoded);

    assertEq(decoded.receiverObjectIds.length, 2);
    assertEq(decoded.receiverObjectIds[0], keccak256("object1"));
  }

  function test_DecodeSuiExecutorArgsV1_ZeroValueObjectIds() public view {
    bytes32[] memory objectIds = new bytes32[](1);
    objectIds[0] = bytes32(0);

    ExtraArgsCodec.SuiExecutorArgsV1 memory args = ExtraArgsCodec.SuiExecutorArgsV1({receiverObjectIds: objectIds});

    bytes memory encoded = ExtraArgsCodec._encodeSuiExecutorArgsV1(args);
    ExtraArgsCodec.SuiExecutorArgsV1 memory decoded = s_helper._decodeSuiExecutorArgsV1(encoded);

    assertEq(decoded.receiverObjectIds[0], bytes32(0));
  }

  function test_DecodeSuiExecutorArgsV1_RevertWhen_DataTooShort() public {
    vm.expectRevert(
      abi.encodeWithSelector(
        ExtraArgsCodec.InvalidDataLength.selector, ExtraArgsCodec.EncodingErrorLocation.EXTRA_ARGS_STATIC_LENGTH_FIELDS
      )
    );
    s_helper._decodeSuiExecutorArgsV1(new bytes(3));
  }

  function test_DecodeSuiExecutorArgsV1_RevertWhen_ExtraBytes() public {
    ExtraArgsCodec.SuiExecutorArgsV1 memory args =
      ExtraArgsCodec.SuiExecutorArgsV1({receiverObjectIds: new bytes32[](0)});

    bytes memory encoded = ExtraArgsCodec._encodeSuiExecutorArgsV1(args);
    bytes memory withExtra = bytes.concat(encoded, bytes("extra"));

    vm.expectRevert(
      abi.encodeWithSelector(
        ExtraArgsCodec.InvalidDataLength.selector, ExtraArgsCodec.EncodingErrorLocation.SUI_EXECUTOR_FINAL_OFFSET
      )
    );
    s_helper._decodeSuiExecutorArgsV1(withExtra);
  }
}
